## refs

- 最初のPOSTリクエストを行う https://developer.freee.co.jp//tutorials/first-call
- ECサイトでの入力例 https://www.freee.co.jp/houjin2/netshop/
- Booth 手数料: https://booth.pm/guide
- Baseの例: https://support.freee.co.jp/hc/ja/articles/203274964-BASE-%E3%83%99%E3%82%A4%E3%82%B9-%E5%A3%B2%E4%B8%8A%E3%83%87%E3%83%BC%E3%82%BF%E3%82%92%E5%8F%96%E3%82%8A%E8%BE%BC%E3%82%80
- oapi https://openapi-generator.tech/docs/generators/go/


## 流れ

- 最初は手で入力する
  - 1件の注文毎に、収入と手数料の支出を記入する
  - 「決済サービス・電子マネー」にてその他の口座を作成 https://secure.freee.co.jp/walletables
- 
- 
- 
- 


## 気付き

- ドキュメントにてoapiのモデルが表示されていないので、フィールドに入れるべき値が分かりにくい
- `from_walletable_type: wallet` が分からなかった
- 会計知識がないと少し手間取るかなと感じる
- freeeのWebUI上で操作した上でChromeDevToolsなどで通信を覗くと、どのようなデータを投げれば良いか非常に参考になる
- 金額にマイナスが使えることが分からずかなり悩んだ箇所がある・もう少し詳細なドキュメントが有れば良いなと感じた
- 
- 
- 

```json
{
  "status_code": 400,
  "errors": [
    {
      "type": "status",
      "messages": [
        "不正なリクエストです。"
      ]
    },
    {
      "type": "validation",
      "messages": [
        "貸借が一致していません。",
        "貸借が一致していません。",
        "更新金額と決済金額の合計が取引の合計金額を上回らないように入力して下さい。",
        "決済金額が取引金額を上回らないように入力してください。",
        "取引単位は不正な値です。"
      ]
    }
  ]
}
```

```json
// company_id: 事業所ID
// account_item_id: 勘定科目ID 429799547は「売上高」
// tax_code: 税区分 1は「課税売上10%」

// POST /api/1/deals
{
  "issue_date": "2020-02-23",
  "type": "income",
  "company_id": 2686190,
  "due_date": "",
  "ref_number": "",
  "details": [
    {
      "tax_code": 129,
      "account_item_id": 429799547,
      "amount": 500,
      "description": "pixivBoothOrder=123456789"
    },
    {
      "tax_code": 136,
      "account_item_id": 429799593,
      "amount": -18
    }
  ],
  "payments": [
    {
      "amount": 482,
      "from_walletable_type": "wallet",
      "from_walletable_id": 3099244,
      "date": "2020-02-23"
    }
  ]
}
```

