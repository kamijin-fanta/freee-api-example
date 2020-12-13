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

// TrialBsThreeYearsResponseTrialBsThreeYearsItems struct for TrialBsThreeYearsResponseTrialBsThreeYearsItems
type TrialBsThreeYearsResponseTrialBsThreeYearsItems struct {
	// 品目ID
	Id int32 `json:"id"`
	// 品目
	Name *string `json:"name,omitempty"`
	// 前々年度期末残高
	TwoYearsBeforeClosingBalance *int32 `json:"two_years_before_closing_balance,omitempty"`
	// 前年度期末残高
	LastYearClosingBalance *int32 `json:"last_year_closing_balance,omitempty"`
	// 期末残高
	ClosingBalance *int32 `json:"closing_balance,omitempty"`
	// 前年比
	YearOnYear *float32 `json:"year_on_year,omitempty"`
}

// NewTrialBsThreeYearsResponseTrialBsThreeYearsItems instantiates a new TrialBsThreeYearsResponseTrialBsThreeYearsItems object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTrialBsThreeYearsResponseTrialBsThreeYearsItems(id int32, ) *TrialBsThreeYearsResponseTrialBsThreeYearsItems {
	this := TrialBsThreeYearsResponseTrialBsThreeYearsItems{}
	this.Id = id
	return &this
}

// NewTrialBsThreeYearsResponseTrialBsThreeYearsItemsWithDefaults instantiates a new TrialBsThreeYearsResponseTrialBsThreeYearsItems object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTrialBsThreeYearsResponseTrialBsThreeYearsItemsWithDefaults() *TrialBsThreeYearsResponseTrialBsThreeYearsItems {
	this := TrialBsThreeYearsResponseTrialBsThreeYearsItems{}
	return &this
}

// GetId returns the Id field value
func (o *TrialBsThreeYearsResponseTrialBsThreeYearsItems) GetId() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TrialBsThreeYearsResponseTrialBsThreeYearsItems) GetIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TrialBsThreeYearsResponseTrialBsThreeYearsItems) SetId(v int32) {
	o.Id = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TrialBsThreeYearsResponseTrialBsThreeYearsItems) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrialBsThreeYearsResponseTrialBsThreeYearsItems) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TrialBsThreeYearsResponseTrialBsThreeYearsItems) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TrialBsThreeYearsResponseTrialBsThreeYearsItems) SetName(v string) {
	o.Name = &v
}

// GetTwoYearsBeforeClosingBalance returns the TwoYearsBeforeClosingBalance field value if set, zero value otherwise.
func (o *TrialBsThreeYearsResponseTrialBsThreeYearsItems) GetTwoYearsBeforeClosingBalance() int32 {
	if o == nil || o.TwoYearsBeforeClosingBalance == nil {
		var ret int32
		return ret
	}
	return *o.TwoYearsBeforeClosingBalance
}

// GetTwoYearsBeforeClosingBalanceOk returns a tuple with the TwoYearsBeforeClosingBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrialBsThreeYearsResponseTrialBsThreeYearsItems) GetTwoYearsBeforeClosingBalanceOk() (*int32, bool) {
	if o == nil || o.TwoYearsBeforeClosingBalance == nil {
		return nil, false
	}
	return o.TwoYearsBeforeClosingBalance, true
}

// HasTwoYearsBeforeClosingBalance returns a boolean if a field has been set.
func (o *TrialBsThreeYearsResponseTrialBsThreeYearsItems) HasTwoYearsBeforeClosingBalance() bool {
	if o != nil && o.TwoYearsBeforeClosingBalance != nil {
		return true
	}

	return false
}

// SetTwoYearsBeforeClosingBalance gets a reference to the given int32 and assigns it to the TwoYearsBeforeClosingBalance field.
func (o *TrialBsThreeYearsResponseTrialBsThreeYearsItems) SetTwoYearsBeforeClosingBalance(v int32) {
	o.TwoYearsBeforeClosingBalance = &v
}

// GetLastYearClosingBalance returns the LastYearClosingBalance field value if set, zero value otherwise.
func (o *TrialBsThreeYearsResponseTrialBsThreeYearsItems) GetLastYearClosingBalance() int32 {
	if o == nil || o.LastYearClosingBalance == nil {
		var ret int32
		return ret
	}
	return *o.LastYearClosingBalance
}

// GetLastYearClosingBalanceOk returns a tuple with the LastYearClosingBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrialBsThreeYearsResponseTrialBsThreeYearsItems) GetLastYearClosingBalanceOk() (*int32, bool) {
	if o == nil || o.LastYearClosingBalance == nil {
		return nil, false
	}
	return o.LastYearClosingBalance, true
}

// HasLastYearClosingBalance returns a boolean if a field has been set.
func (o *TrialBsThreeYearsResponseTrialBsThreeYearsItems) HasLastYearClosingBalance() bool {
	if o != nil && o.LastYearClosingBalance != nil {
		return true
	}

	return false
}

// SetLastYearClosingBalance gets a reference to the given int32 and assigns it to the LastYearClosingBalance field.
func (o *TrialBsThreeYearsResponseTrialBsThreeYearsItems) SetLastYearClosingBalance(v int32) {
	o.LastYearClosingBalance = &v
}

// GetClosingBalance returns the ClosingBalance field value if set, zero value otherwise.
func (o *TrialBsThreeYearsResponseTrialBsThreeYearsItems) GetClosingBalance() int32 {
	if o == nil || o.ClosingBalance == nil {
		var ret int32
		return ret
	}
	return *o.ClosingBalance
}

// GetClosingBalanceOk returns a tuple with the ClosingBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrialBsThreeYearsResponseTrialBsThreeYearsItems) GetClosingBalanceOk() (*int32, bool) {
	if o == nil || o.ClosingBalance == nil {
		return nil, false
	}
	return o.ClosingBalance, true
}

// HasClosingBalance returns a boolean if a field has been set.
func (o *TrialBsThreeYearsResponseTrialBsThreeYearsItems) HasClosingBalance() bool {
	if o != nil && o.ClosingBalance != nil {
		return true
	}

	return false
}

// SetClosingBalance gets a reference to the given int32 and assigns it to the ClosingBalance field.
func (o *TrialBsThreeYearsResponseTrialBsThreeYearsItems) SetClosingBalance(v int32) {
	o.ClosingBalance = &v
}

// GetYearOnYear returns the YearOnYear field value if set, zero value otherwise.
func (o *TrialBsThreeYearsResponseTrialBsThreeYearsItems) GetYearOnYear() float32 {
	if o == nil || o.YearOnYear == nil {
		var ret float32
		return ret
	}
	return *o.YearOnYear
}

// GetYearOnYearOk returns a tuple with the YearOnYear field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrialBsThreeYearsResponseTrialBsThreeYearsItems) GetYearOnYearOk() (*float32, bool) {
	if o == nil || o.YearOnYear == nil {
		return nil, false
	}
	return o.YearOnYear, true
}

// HasYearOnYear returns a boolean if a field has been set.
func (o *TrialBsThreeYearsResponseTrialBsThreeYearsItems) HasYearOnYear() bool {
	if o != nil && o.YearOnYear != nil {
		return true
	}

	return false
}

// SetYearOnYear gets a reference to the given float32 and assigns it to the YearOnYear field.
func (o *TrialBsThreeYearsResponseTrialBsThreeYearsItems) SetYearOnYear(v float32) {
	o.YearOnYear = &v
}

func (o TrialBsThreeYearsResponseTrialBsThreeYearsItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.TwoYearsBeforeClosingBalance != nil {
		toSerialize["two_years_before_closing_balance"] = o.TwoYearsBeforeClosingBalance
	}
	if o.LastYearClosingBalance != nil {
		toSerialize["last_year_closing_balance"] = o.LastYearClosingBalance
	}
	if o.ClosingBalance != nil {
		toSerialize["closing_balance"] = o.ClosingBalance
	}
	if o.YearOnYear != nil {
		toSerialize["year_on_year"] = o.YearOnYear
	}
	return json.Marshal(toSerialize)
}

type NullableTrialBsThreeYearsResponseTrialBsThreeYearsItems struct {
	value *TrialBsThreeYearsResponseTrialBsThreeYearsItems
	isSet bool
}

func (v NullableTrialBsThreeYearsResponseTrialBsThreeYearsItems) Get() *TrialBsThreeYearsResponseTrialBsThreeYearsItems {
	return v.value
}

func (v *NullableTrialBsThreeYearsResponseTrialBsThreeYearsItems) Set(val *TrialBsThreeYearsResponseTrialBsThreeYearsItems) {
	v.value = val
	v.isSet = true
}

func (v NullableTrialBsThreeYearsResponseTrialBsThreeYearsItems) IsSet() bool {
	return v.isSet
}

func (v *NullableTrialBsThreeYearsResponseTrialBsThreeYearsItems) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrialBsThreeYearsResponseTrialBsThreeYearsItems(val *TrialBsThreeYearsResponseTrialBsThreeYearsItems) *NullableTrialBsThreeYearsResponseTrialBsThreeYearsItems {
	return &NullableTrialBsThreeYearsResponseTrialBsThreeYearsItems{value: val, isSet: true}
}

func (v NullableTrialBsThreeYearsResponseTrialBsThreeYearsItems) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrialBsThreeYearsResponseTrialBsThreeYearsItems) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


