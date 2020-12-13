package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/kb"
	"github.com/joho/godotenv"
	"github.com/kamijin-fanta/freee-api-example/pkg/freee"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"os/signal"
	"path"
	"strconv"
	"strings"
	"syscall"
)

func main() {
	godotenv.Load()

	app := &cli.App{
		Name: "freee-booth-exporter",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "pixiv-username",
				Usage:    "Pixivへログインする際のユーザ名",
				EnvVars:  []string{"PIXIV_USERNAME"},
				Required: true,
			},
			&cli.StringFlag{
				Name:     "pixiv-password",
				Usage:    "Pixivへログインする際のパスワード",
				EnvVars:  []string{"PIXIV_PASSWORD"},
				Required: true,
			},
			&cli.IntFlag{
				Name:     "freee-company-id",
				Usage:    "freee上での事業所ID",
				EnvVars:  []string{"FREEE_COMPANY_ID"},
				Required: true,
			},
			&cli.IntFlag{
				Name:     "freee-walletable-id",
				Usage:    "freee上で作成したPixivBooth用の口座ID",
				EnvVars:  []string{"FREEE_WALLETABLE_ID"},
				Required: true,
			},
			&cli.StringFlag{
				Name:     "freee-access-token",
				Usage:    "freeeにアクセスするためのトークン文字列",
				EnvVars:  []string{"FREEE_ACCESS_TOKEN"},
				Required: true,
			},
		},
		Action: func(c *cli.Context) error {
			username := c.String("pixiv-username")
			password := c.String("pixiv-password")
			companyId := int32(c.Int("freee-company-id"))
			walletableId := int32(c.Int("freee-walletable-id"))
			accessToken := c.String("freee-access-token")

			// freeeのOpenAPIクライアントを作成
			apiConfig := freee.NewConfiguration()
			client := freee.NewAPIClient(apiConfig)
			apiClientCtx := context.TODO()
			apiClientCtx = context.WithValue(apiClientCtx, freee.ContextAccessToken, accessToken)

			// Chromedpクライアントの作成
			opts := append(
				chromedp.DefaultExecAllocatorOptions[:],
				chromedp.Flag("headless", false), // ブラウザが見えるようにオプションの追加
			)
			browserCtx, _ := chromedp.NewExecAllocator(context.Background(), opts...)
			browserCtx, cancel := chromedp.NewContext(browserCtx)
			defer cancel()
			// Ctrl+CでChromeが終了するようシグナルを受ける
			sigs := make(chan os.Signal, 1)
			signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
			go func() {
				<-sigs
				cancel()
				os.Exit(1)
			}()

			// freee会計から取引一覧を取得、最も大きいPixivBoothの注文IDを検出する
			dealHistory, err := FetchFreeeDealHistory(apiClientCtx, client, companyId, walletableId)
			if err != nil {
				panic(err)
			}
			latestOrderId := ExtractLatestOrderId(dealHistory)
			fmt.Printf("latest pxiv booth order id %d\n", latestOrderId)

			// PixivBoothの注文一覧を取得
			orders, err := FetchBoothOrders(browserCtx, username, password, latestOrderId)
			if err != nil {
				panic(err)
			}

			// 注文一覧を元に、freeeの取引を作成する
			err = CreateFreeeDeals(apiClientCtx, client, companyId, walletableId, orders)
			if err != nil {
				panic(err)
			}

			by, _ := json.MarshalIndent(orders, "", "  ")
			fmt.Printf("%s\n", by)
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}

type Order struct {
	Detail OrderDetail
	Date   string
	Id     string
}
type OrderDetail struct {
	Fee    int
	Income int
}

func FetchFreeeDealHistory(ctx context.Context, client *freee.APIClient, companyId, walletableId int32) ([]*freee.Deal, error) {
	dealHistory, _, err := client.DealsApi.GetDeals(ctx).
		//RegisteredFrom("me"). // 本来は叩いているAPIクライアントからの取引をフィルタ出来るが、何らかの理由で行えないので以下にて口座でフィルタを行う
		CompanyId(companyId).
		Limit(int32(100)).
		Execute()
	if err != nil {
		return nil, err
	}
	filteredDealHistories := make([]*freee.Deal, 0, len(dealHistory.Deals))
	for i := range dealHistory.Deals {
		found := false
		if dealHistory.Deals[i].Payments != nil {
			for _, pay := range *dealHistory.Deals[i].Payments {
				if *pay.FromWalletableId == walletableId {
					found = true
				}
			}
		}
		if found {
			filteredDealHistories = append(filteredDealHistories, &dealHistory.Deals[i])
		}
	}
	return filteredDealHistories, nil
}

func CreateFreeeDeals(ctx context.Context, client *freee.APIClient, companyId, walletableId int32, orders []*Order) error {
	for _, order := range orders {
		// 取引の詳細欄に注文IDを記述することで重複した登録を行うことを回避する
		description := fmt.Sprintf("pixivBoothOrder=%s", order.Id)

		// 取引登録を行う
		params := freee.DealCreateParams{
			IssueDate: order.Date,
			Type:      "income",
			CompanyId: companyId,
			Details: []freee.DealCreateParamsDetails{
				// 売上
				{
					TaxCode:       129,       // 課税売上10%
					AccountItemId: 429799547, // 売上高
					Amount:        int32(order.Detail.Income),
					Description:   &description,
				},

				// サービス利用料金の支出登録を行う
				{
					TaxCode:       136,       // 課対仕入10%
					AccountItemId: 429799593, // 販売手数料
					Amount:        int32(-order.Detail.Fee),
				},
			},
			Payments: &[]freee.DealCreateParamsPayments{
				{
					Amount:             int32(order.Detail.Income - order.Detail.Fee),
					FromWalletableType: "wallet",
					FromWalletableId:   walletableId,
					Date:               order.Date,
				},
			},
		}
		_, _, err := client.DealsApi.CreateDeal(ctx).
			DealCreateParams(params).
			Execute()
		if err != nil {
			return err
		}
	}
	return nil
}

func ExtractLatestOrderId(orders []*freee.Deal) int64 {
	latest := int64(0)
	for _, order := range orders {
		if order.Details != nil {
			for _, order := range *order.Details {
				if order.Description != nil {
					// 説明文に記述されている pixivBoothOrder=12345678,others=filed 形式の文字列から注文IDを抜き出す
					desc := *order.Description
					attrs := strings.Split(desc, ",")
					for _, attr := range attrs {
						pair := strings.Split(attr, "=")
						if len(pair) == 2 && pair[0] == "pixivBoothOrder" {
							orderId, _ := strconv.ParseInt(pair[1], 10, 64)
							if latest < orderId {
								latest = orderId
							}
						}
					}
				}
			}
		}
	}
	return latest
}

func FetchBoothOrders(ctx context.Context, username, password string, latestOrderId int64) ([]*Order, error) {
	var orderLinks []string
	err := chromedp.Run(ctx,
		// Pixiv Boothへのログイン
		chromedp.Navigate(`https://accounts.pixiv.net/login?prompt=select_account&return_to=https%3A%2F%2Fmanage.booth.pm%2Fusers%2Fauth%2Fpixiv&source=booth&view_type=page`),
		chromedp.SendKeys(`input[autocomplete="username"]`, username),
		chromedp.SendKeys(`input[autocomplete="current-password"]`, password),
		chromedp.SendKeys(`input[autocomplete="current-password"]`, kb.Enter),
		chromedp.WaitVisible("img[alt=BOOTH]"),

		// 注文一覧ページへ移動
		chromedp.Navigate("https://manage.booth.pm/orders"),

		// 注文詳細ページのURLを収集
		chromedp.Evaluate(`[...document.querySelectorAll('.link-ext')].map(link => link.href)`, &orderLinks),
	)
	if err != nil {
		return nil, err
	}

	fmt.Printf("found %d orders\n", len(orderLinks))

	// 注文詳細ページを巡回する
	orders := make([]*Order, 0, len(orderLinks))
	for _, orderUrl := range orderLinks {
		// URLから注文IDを抽出し、latestOrderIdと比較・既に取得済みの注文はスキップする
		orderIdStr := path.Base(orderUrl)
		orderId, _ := strconv.ParseInt(orderIdStr, 10, 64)
		if orderId <= latestOrderId {
			continue
		}

		order := Order{}
		err := chromedp.Run(ctx,
			chromedp.Navigate(orderUrl),
			chromedp.WaitVisible(".manage-page-head-title"),
			// 注文金額・サービス利用手数料・取引日・注文IDをブラウザ上で取得し、order構造体へ代入する
			chromedp.Evaluate(`
				({
				  Detail: [...document.querySelector('div.co-breakdown-table.u-mt-200.u-pt-200.u-tpg-footnote > div > div.lo-grid-cell').querySelectorAll('.lo-grid-cell .lo-grid-cell .lo-grid')].map(cell => ({title: cell.children[0].textContent, value: +cell.children[1].textContent.substr(2).replace(/,/g, '')})).filter(row => row.title !== '商品合計').reduce((res, row) => { const key = {"合計金額": "Income", "サービス利用料": "Fee"}[row.title]; res[key] = row.value; return res  },{}),
				  Date: document.querySelector('body > div.page-wrap > main > div.manage-page-body > div > div:nth-child(1) > div > div > div:nth-child(2) > div:nth-child(2)').textContent.substring(0, 10).replace(/\//g, '-'),
				  Id: document.querySelector('body > div.page-wrap > main > div.manage-page-body > div > div:nth-child(1) > div > div > div:nth-child(1) > div:nth-child(2)').textContent,
				})
			`, &order),
		)
		if err != nil {
			return nil, err
		}

		orders = append(orders, &order)
	}

	return orders, nil
}
