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

// DealCreateResponseDeal struct for DealCreateResponseDeal
type DealCreateResponseDeal struct {
	// 取引ID
	Id int32 `json:"id"`
	// 事業所ID
	CompanyId int32 `json:"company_id"`
	// 発生日 (yyyy-mm-dd)
	IssueDate string `json:"issue_date"`
	// 支払期日 (yyyy-mm-dd)
	DueDate *string `json:"due_date,omitempty"`
	// 金額
	Amount int32 `json:"amount"`
	// 支払金額
	DueAmount *int32 `json:"due_amount,omitempty"`
	// 収支区分 (収入: income, 支出: expense)
	Type *string `json:"type,omitempty"`
	// 取引先ID
	PartnerId int32 `json:"partner_id"`
	// 取引先コード
	PartnerCode NullableString `json:"partner_code,omitempty"`
	// 管理番号
	RefNumber *string `json:"ref_number,omitempty"`
	// 決済状況 (未決済: unsettled, 完了: settled)
	Status string `json:"status"`
	// 取引の明細行
	Details *[]DealCreateResponseDealDetails `json:"details,omitempty"`
	// 取引の支払行
	Payments *[]DealCreateResponseDealPayments `json:"payments,omitempty"`
}

// NewDealCreateResponseDeal instantiates a new DealCreateResponseDeal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDealCreateResponseDeal(id int32, companyId int32, issueDate string, amount int32, partnerId int32, status string, ) *DealCreateResponseDeal {
	this := DealCreateResponseDeal{}
	this.Id = id
	this.CompanyId = companyId
	this.IssueDate = issueDate
	this.Amount = amount
	this.PartnerId = partnerId
	this.Status = status
	return &this
}

// NewDealCreateResponseDealWithDefaults instantiates a new DealCreateResponseDeal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDealCreateResponseDealWithDefaults() *DealCreateResponseDeal {
	this := DealCreateResponseDeal{}
	return &this
}

// GetId returns the Id field value
func (o *DealCreateResponseDeal) GetId() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DealCreateResponseDeal) GetIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DealCreateResponseDeal) SetId(v int32) {
	o.Id = v
}

// GetCompanyId returns the CompanyId field value
func (o *DealCreateResponseDeal) GetCompanyId() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.CompanyId
}

// GetCompanyIdOk returns a tuple with the CompanyId field value
// and a boolean to check if the value has been set.
func (o *DealCreateResponseDeal) GetCompanyIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CompanyId, true
}

// SetCompanyId sets field value
func (o *DealCreateResponseDeal) SetCompanyId(v int32) {
	o.CompanyId = v
}

// GetIssueDate returns the IssueDate field value
func (o *DealCreateResponseDeal) GetIssueDate() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.IssueDate
}

// GetIssueDateOk returns a tuple with the IssueDate field value
// and a boolean to check if the value has been set.
func (o *DealCreateResponseDeal) GetIssueDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IssueDate, true
}

// SetIssueDate sets field value
func (o *DealCreateResponseDeal) SetIssueDate(v string) {
	o.IssueDate = v
}

// GetDueDate returns the DueDate field value if set, zero value otherwise.
func (o *DealCreateResponseDeal) GetDueDate() string {
	if o == nil || o.DueDate == nil {
		var ret string
		return ret
	}
	return *o.DueDate
}

// GetDueDateOk returns a tuple with the DueDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DealCreateResponseDeal) GetDueDateOk() (*string, bool) {
	if o == nil || o.DueDate == nil {
		return nil, false
	}
	return o.DueDate, true
}

// HasDueDate returns a boolean if a field has been set.
func (o *DealCreateResponseDeal) HasDueDate() bool {
	if o != nil && o.DueDate != nil {
		return true
	}

	return false
}

// SetDueDate gets a reference to the given string and assigns it to the DueDate field.
func (o *DealCreateResponseDeal) SetDueDate(v string) {
	o.DueDate = &v
}

// GetAmount returns the Amount field value
func (o *DealCreateResponseDeal) GetAmount() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *DealCreateResponseDeal) GetAmountOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *DealCreateResponseDeal) SetAmount(v int32) {
	o.Amount = v
}

// GetDueAmount returns the DueAmount field value if set, zero value otherwise.
func (o *DealCreateResponseDeal) GetDueAmount() int32 {
	if o == nil || o.DueAmount == nil {
		var ret int32
		return ret
	}
	return *o.DueAmount
}

// GetDueAmountOk returns a tuple with the DueAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DealCreateResponseDeal) GetDueAmountOk() (*int32, bool) {
	if o == nil || o.DueAmount == nil {
		return nil, false
	}
	return o.DueAmount, true
}

// HasDueAmount returns a boolean if a field has been set.
func (o *DealCreateResponseDeal) HasDueAmount() bool {
	if o != nil && o.DueAmount != nil {
		return true
	}

	return false
}

// SetDueAmount gets a reference to the given int32 and assigns it to the DueAmount field.
func (o *DealCreateResponseDeal) SetDueAmount(v int32) {
	o.DueAmount = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DealCreateResponseDeal) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DealCreateResponseDeal) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DealCreateResponseDeal) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *DealCreateResponseDeal) SetType(v string) {
	o.Type = &v
}

// GetPartnerId returns the PartnerId field value
func (o *DealCreateResponseDeal) GetPartnerId() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.PartnerId
}

// GetPartnerIdOk returns a tuple with the PartnerId field value
// and a boolean to check if the value has been set.
func (o *DealCreateResponseDeal) GetPartnerIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PartnerId, true
}

// SetPartnerId sets field value
func (o *DealCreateResponseDeal) SetPartnerId(v int32) {
	o.PartnerId = v
}

// GetPartnerCode returns the PartnerCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DealCreateResponseDeal) GetPartnerCode() string {
	if o == nil || o.PartnerCode.Get() == nil {
		var ret string
		return ret
	}
	return *o.PartnerCode.Get()
}

// GetPartnerCodeOk returns a tuple with the PartnerCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DealCreateResponseDeal) GetPartnerCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PartnerCode.Get(), o.PartnerCode.IsSet()
}

// HasPartnerCode returns a boolean if a field has been set.
func (o *DealCreateResponseDeal) HasPartnerCode() bool {
	if o != nil && o.PartnerCode.IsSet() {
		return true
	}

	return false
}

// SetPartnerCode gets a reference to the given NullableString and assigns it to the PartnerCode field.
func (o *DealCreateResponseDeal) SetPartnerCode(v string) {
	o.PartnerCode.Set(&v)
}
// SetPartnerCodeNil sets the value for PartnerCode to be an explicit nil
func (o *DealCreateResponseDeal) SetPartnerCodeNil() {
	o.PartnerCode.Set(nil)
}

// UnsetPartnerCode ensures that no value is present for PartnerCode, not even an explicit nil
func (o *DealCreateResponseDeal) UnsetPartnerCode() {
	o.PartnerCode.Unset()
}

// GetRefNumber returns the RefNumber field value if set, zero value otherwise.
func (o *DealCreateResponseDeal) GetRefNumber() string {
	if o == nil || o.RefNumber == nil {
		var ret string
		return ret
	}
	return *o.RefNumber
}

// GetRefNumberOk returns a tuple with the RefNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DealCreateResponseDeal) GetRefNumberOk() (*string, bool) {
	if o == nil || o.RefNumber == nil {
		return nil, false
	}
	return o.RefNumber, true
}

// HasRefNumber returns a boolean if a field has been set.
func (o *DealCreateResponseDeal) HasRefNumber() bool {
	if o != nil && o.RefNumber != nil {
		return true
	}

	return false
}

// SetRefNumber gets a reference to the given string and assigns it to the RefNumber field.
func (o *DealCreateResponseDeal) SetRefNumber(v string) {
	o.RefNumber = &v
}

// GetStatus returns the Status field value
func (o *DealCreateResponseDeal) GetStatus() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *DealCreateResponseDeal) GetStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *DealCreateResponseDeal) SetStatus(v string) {
	o.Status = v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *DealCreateResponseDeal) GetDetails() []DealCreateResponseDealDetails {
	if o == nil || o.Details == nil {
		var ret []DealCreateResponseDealDetails
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DealCreateResponseDeal) GetDetailsOk() (*[]DealCreateResponseDealDetails, bool) {
	if o == nil || o.Details == nil {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *DealCreateResponseDeal) HasDetails() bool {
	if o != nil && o.Details != nil {
		return true
	}

	return false
}

// SetDetails gets a reference to the given []DealCreateResponseDealDetails and assigns it to the Details field.
func (o *DealCreateResponseDeal) SetDetails(v []DealCreateResponseDealDetails) {
	o.Details = &v
}

// GetPayments returns the Payments field value if set, zero value otherwise.
func (o *DealCreateResponseDeal) GetPayments() []DealCreateResponseDealPayments {
	if o == nil || o.Payments == nil {
		var ret []DealCreateResponseDealPayments
		return ret
	}
	return *o.Payments
}

// GetPaymentsOk returns a tuple with the Payments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DealCreateResponseDeal) GetPaymentsOk() (*[]DealCreateResponseDealPayments, bool) {
	if o == nil || o.Payments == nil {
		return nil, false
	}
	return o.Payments, true
}

// HasPayments returns a boolean if a field has been set.
func (o *DealCreateResponseDeal) HasPayments() bool {
	if o != nil && o.Payments != nil {
		return true
	}

	return false
}

// SetPayments gets a reference to the given []DealCreateResponseDealPayments and assigns it to the Payments field.
func (o *DealCreateResponseDeal) SetPayments(v []DealCreateResponseDealPayments) {
	o.Payments = &v
}

func (o DealCreateResponseDeal) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["company_id"] = o.CompanyId
	}
	if true {
		toSerialize["issue_date"] = o.IssueDate
	}
	if o.DueDate != nil {
		toSerialize["due_date"] = o.DueDate
	}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if o.DueAmount != nil {
		toSerialize["due_amount"] = o.DueAmount
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["partner_id"] = o.PartnerId
	}
	if o.PartnerCode.IsSet() {
		toSerialize["partner_code"] = o.PartnerCode.Get()
	}
	if o.RefNumber != nil {
		toSerialize["ref_number"] = o.RefNumber
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if o.Details != nil {
		toSerialize["details"] = o.Details
	}
	if o.Payments != nil {
		toSerialize["payments"] = o.Payments
	}
	return json.Marshal(toSerialize)
}

type NullableDealCreateResponseDeal struct {
	value *DealCreateResponseDeal
	isSet bool
}

func (v NullableDealCreateResponseDeal) Get() *DealCreateResponseDeal {
	return v.value
}

func (v *NullableDealCreateResponseDeal) Set(val *DealCreateResponseDeal) {
	v.value = val
	v.isSet = true
}

func (v NullableDealCreateResponseDeal) IsSet() bool {
	return v.isSet
}

func (v *NullableDealCreateResponseDeal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDealCreateResponseDeal(val *DealCreateResponseDeal) *NullableDealCreateResponseDeal {
	return &NullableDealCreateResponseDeal{value: val, isSet: true}
}

func (v NullableDealCreateResponseDeal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDealCreateResponseDeal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


