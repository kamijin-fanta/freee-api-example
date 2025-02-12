/*
 * freee API
 *
 *  <h1 id=\"freee_api\">freee API</h1> <hr /> <h2 id=\"start_guide\">スタートガイド</h2>  <p>freee API開発がはじめての方は<a href=\"https://developer.freee.co.jp/getting-started\">freee API スタートガイド</a>を参照してください。</p>  <hr /> <h2 id=\"specification\">仕様</h2>  <pre><code>【重要】会計freee APIの新バージョンについて 2020年12月まで、2つのバージョンが利用できる状態です。古いものは2020年12月に利用不可となります。<br> 新しいAPIを利用するにはリクエストヘッダーに以下を指定します。 X-Api-Version: 2020-06-15<br> 指定がない場合は2020年12月に廃止予定のAPIを利用することとなります。<br> 【重要】APIのバージョン指定をせずに利用し続ける場合 2020年12月に新しいバージョンのAPIに自動的に切り替わります。 詳細は、<a href=\"https://developer.freee.co.jp/release-note/2948\" target=\"_blank\">リリースノート</a>をご覧ください。<br> 旧バージョンのAPIリファレンスを確認したい場合は、<a href=\"https://freee.github.io/freee-api-schema/\" target=\"_blank\">旧バージョンのAPIリファレンスページ</a>をご覧ください。 </code></pre>  <h3 id=\"api_endpoint\">APIエンドポイント</h3>  <p>https://api.freee.co.jp/ (httpsのみ)</p>  <h3 id=\"about_authorize\">認証について</h3> <p>OAuth2.0を利用します。詳細は<a href=\"https://developer.freee.co.jp/docs\" target=\"_blank\">ドキュメントの認証</a>パートを参照してください。</p>  <h3 id=\"data_format\">データフォーマット</h3>  <p>リクエスト、レスポンスともにJSON形式をサポートしていますが、詳細は、API毎の説明欄（application/jsonなど）を確認してください。</p>  <h3 id=\"compatibility\">後方互換性ありの変更</h3>  <p>freeeでは、APIを改善していくために以下のような変更は後方互換性ありとして通知なく変更を入れることがあります。アプリケーション実装者は以下を踏まえて開発を行ってください。</p>  <ul> <li>新しいAPIリソース・エンドポイントの追加</li> <li>既存のAPIに対して必須ではない新しいリクエストパラメータの追加</li> <li>既存のAPIレスポンスに対する新しいプロパティの追加</li> <li>既存のAPIレスポンスに対するプロパティの順番の入れ変え</li> <li>keyとなっているidやcodeの長さの変更（長くする）</li> </ul>  <h3 id=\"common_response_header\">共通レスポンスヘッダー</h3>  <p>すべてのAPIのレスポンスには以下のHTTPヘッダーが含まれます。</p>  <ul> <li> <p>X-Freee-Request-ID</p> <ul> <li>各リクエスト毎に発行されるID</li> </ul> </li> </ul>  <h3 id=\"common_error_response\">共通エラーレスポンス</h3>  <ul> <li> <p>ステータスコードはレスポンス内のJSONに含まれる他、HTTPヘッダにも含まれる</p> </li> <li> <p>一部のエラーレスポンスにはエラーコードが含まれます。<br>詳細は、<a href=\"https://developer.freee.co.jp/tips/faq/40x-checkpoint\">HTTPステータスコード400台エラー時のチェックポイント</a>を参照してください</p> </li> <p>type</p>  <ul> <li>status : HTTPステータスコードの説明</li>  <li>validation : エラーの詳細の説明（開発者向け）</li> </ul> </li> </ul>  <p>レスポンスの例</p>  <pre><code>  {     &quot;status_code&quot; : 400,     &quot;errors&quot; : [       {         &quot;type&quot; : &quot;status&quot;,         &quot;messages&quot; : [&quot;不正なリクエストです。&quot;]       },       {         &quot;type&quot; : &quot;validation&quot;,         &quot;messages&quot; : [&quot;Date は不正な日付フォーマットです。入力例：2013-01-01&quot;]       }     ]   }</code></pre>  </br>  <h3 id=\"api_rate_limit\">API使用制限</h3>    <p>freeeは一定期間に過度のアクセスを検知した場合、APIアクセスをコントロールする場合があります。</p>   <p>その際のhttp status codeは403となります。制限がかかってから10分程度が過ぎると再度使用することができるようになります。</p>  <h4 id=\"reports_api_endpoint\">/reportsと/receipts/{id}/downloadエンドポイント</h4>  <p>freeeはエンドポイント毎に一定頻度以上のアクセスを検知した場合、APIアクセスをコントロールする場合があります。その際のhttp status codeは429（too many requests）となります。</p>  <ul>   <li>/reports:1秒に10回まで</li>   <li>/receipts/{id}/download:1秒に3回まで</li> </ul>  <p>レスポンスボディのmetaプロパティに以下を含めます。</p>  <ul>   <li>設定されている上限値</li>   <li>上限に達するまでの使用可能回数</li>   <li>（上限値に達した場合）使用回数がリセットされる時刻</li> </ul>  <h3 id=\"plan_api_rate_limit\">プラン別のAPI Rate Limit</h3>   <table border=\"1\">     <tbody>       <tr>         <th style=\"padding: 10px\"><strong>会計freeeプラン名</strong></th>         <th style=\"padding: 10px\"><strong>事業所とアプリケーション毎に1日でのAPIコール数</strong></th>       </tr>       <tr>         <td style=\"padding: 10px\">エンタープライズ</td>         <td style=\"padding: 10px\">10,000</td>       </tr>       <tr>         <td style=\"padding: 10px\">プロフェッショナル</td>         <td style=\"padding: 10px\">5,000</td>       </tr>       <tr>         <td style=\"padding: 10px\">ベーシック</td>         <td style=\"padding: 10px\">3,000</td>       </tr>       <tr>         <td style=\"padding: 10px\">ミニマム</td>         <td style=\"padding: 10px\">3,000</td>       </tr>       <tr>         <td style=\"padding: 10px\">上記以外</td>         <td style=\"padding: 10px\">3,000</td>       </tr>     </tbody>   </table>  <h3 id=\"webhook\">Webhookについて</h3>  <p>詳細は<a href=\"https://developer.freee.co.jp/docs/accounting/webhook\" target=\"_blank\">会計Webhook概要</a>を参照してください。</p>  <hr /> <h2 id=\"contact\">連絡先</h2>  <p>ご不明点、ご要望等は <a href=\"https://support.freee.co.jp/hc/ja/requests/new\">freee サポートデスクへのお問い合わせフォーム</a> からご連絡ください。</p> <hr />&copy; Since 2013 freee K.K.
 *
 * API version: v1.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package freee

import (
	"encoding/json"
)

// DealCreateParamsPayments struct for DealCreateParamsPayments
type DealCreateParamsPayments struct {
	// 支払金額：payments指定時は必須
	Amount int32 `json:"amount"`
	// 口座ID（from_walletable_typeがprivate_account_itemの場合は勘定科目ID）：payments指定時は必須
	FromWalletableId int32 `json:"from_walletable_id"`
	// 口座区分 (銀行口座: bank_account, クレジットカード: credit_card, 現金: wallet, プライベート資金（法人の場合は役員借入金もしくは役員借入金、個人の場合は事業主貸もしくは事業主借）: private_account_item)：payments指定時は必須
	FromWalletableType string `json:"from_walletable_type"`
	// 支払日：payments指定時は必須
	Date string `json:"date"`
}

// NewDealCreateParamsPayments instantiates a new DealCreateParamsPayments object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDealCreateParamsPayments(amount int32, fromWalletableId int32, fromWalletableType string, date string, ) *DealCreateParamsPayments {
	this := DealCreateParamsPayments{}
	this.Amount = amount
	this.FromWalletableId = fromWalletableId
	this.FromWalletableType = fromWalletableType
	this.Date = date
	return &this
}

// NewDealCreateParamsPaymentsWithDefaults instantiates a new DealCreateParamsPayments object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDealCreateParamsPaymentsWithDefaults() *DealCreateParamsPayments {
	this := DealCreateParamsPayments{}
	return &this
}

// GetAmount returns the Amount field value
func (o *DealCreateParamsPayments) GetAmount() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *DealCreateParamsPayments) GetAmountOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *DealCreateParamsPayments) SetAmount(v int32) {
	o.Amount = v
}

// GetFromWalletableId returns the FromWalletableId field value
func (o *DealCreateParamsPayments) GetFromWalletableId() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.FromWalletableId
}

// GetFromWalletableIdOk returns a tuple with the FromWalletableId field value
// and a boolean to check if the value has been set.
func (o *DealCreateParamsPayments) GetFromWalletableIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FromWalletableId, true
}

// SetFromWalletableId sets field value
func (o *DealCreateParamsPayments) SetFromWalletableId(v int32) {
	o.FromWalletableId = v
}

// GetFromWalletableType returns the FromWalletableType field value
func (o *DealCreateParamsPayments) GetFromWalletableType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.FromWalletableType
}

// GetFromWalletableTypeOk returns a tuple with the FromWalletableType field value
// and a boolean to check if the value has been set.
func (o *DealCreateParamsPayments) GetFromWalletableTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FromWalletableType, true
}

// SetFromWalletableType sets field value
func (o *DealCreateParamsPayments) SetFromWalletableType(v string) {
	o.FromWalletableType = v
}

// GetDate returns the Date field value
func (o *DealCreateParamsPayments) GetDate() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Date
}

// GetDateOk returns a tuple with the Date field value
// and a boolean to check if the value has been set.
func (o *DealCreateParamsPayments) GetDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Date, true
}

// SetDate sets field value
func (o *DealCreateParamsPayments) SetDate(v string) {
	o.Date = v
}

func (o DealCreateParamsPayments) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["from_walletable_id"] = o.FromWalletableId
	}
	if true {
		toSerialize["from_walletable_type"] = o.FromWalletableType
	}
	if true {
		toSerialize["date"] = o.Date
	}
	return json.Marshal(toSerialize)
}

type NullableDealCreateParamsPayments struct {
	value *DealCreateParamsPayments
	isSet bool
}

func (v NullableDealCreateParamsPayments) Get() *DealCreateParamsPayments {
	return v.value
}

func (v *NullableDealCreateParamsPayments) Set(val *DealCreateParamsPayments) {
	v.value = val
	v.isSet = true
}

func (v NullableDealCreateParamsPayments) IsSet() bool {
	return v.isSet
}

func (v *NullableDealCreateParamsPayments) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDealCreateParamsPayments(val *DealCreateParamsPayments) *NullableDealCreateParamsPayments {
	return &NullableDealCreateParamsPayments{value: val, isSet: true}
}

func (v NullableDealCreateParamsPayments) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDealCreateParamsPayments) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


