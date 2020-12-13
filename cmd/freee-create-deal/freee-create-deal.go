package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/kamijin-fanta/freee-api-example/pkg/freee"
)

func main() {
	accessToken := flag.String("access-token", "", "freee access token")
	companyId := flag.Int("company-id", 0, "freee company id")
	walletableId := flag.Int("walletable-id", 0, "freee walletable id")
	flag.Parse()

	apiConfig := freee.NewConfiguration()
	client := freee.NewAPIClient(apiConfig)

	ctx := context.TODO()
	ctx = context.WithValue(ctx, freee.ContextAccessToken, *accessToken)

	description := "APIテスト"
	params := freee.DealCreateParams{
		IssueDate: "2020-12-11",
		Type:      "income",
		CompanyId: int32(*companyId),
		Details: []freee.DealCreateParamsDetails{
			{
				TaxCode:       129,       // 課税売上10%
				AccountItemId: 429799547, // 売上高
				Amount:        800,
				Description:   &description,
			},
		},
		Payments: &[]freee.DealCreateParamsPayments{
			{
				Amount:             800,
				FromWalletableId:   int32(*walletableId),
				FromWalletableType: "wallet",
				Date:               "2020-12-11",
			},
		},
	}
	req := client.DealsApi.CreateDeal(ctx).DealCreateParams(params)

	res, _, err := req.Execute()
	if err != nil {
		panic(err)
	}

	by, _ := json.MarshalIndent(res, "", "  ")
	fmt.Printf("%s\n", by)
}
