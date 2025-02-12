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

// JournalStatusResponseJournals struct for JournalStatusResponseJournals
type JournalStatusResponseJournals struct {
	// 受け付けID
	Id int32 `json:"id"`
	// 事業所ID
	CompanyId int32 `json:"company_id"`
	// ダウンロード形式
	DownloadType string `json:"download_type"`
	// 事業所ID
	Status string `json:"status"`
	// 取得開始日 (yyyy-mm-dd)
	StartDate string `json:"start_date"`
	// 取得終了日 (yyyy-mm-dd)
	EndDate string `json:"end_date"`
	VisibleTags *[]string `json:"visible_tags,omitempty"`
	// ダウンロードURL
	DownloadUrl *string `json:"download_url,omitempty"`
}

// NewJournalStatusResponseJournals instantiates a new JournalStatusResponseJournals object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJournalStatusResponseJournals(id int32, companyId int32, downloadType string, status string, startDate string, endDate string, ) *JournalStatusResponseJournals {
	this := JournalStatusResponseJournals{}
	this.Id = id
	this.CompanyId = companyId
	this.DownloadType = downloadType
	this.Status = status
	this.StartDate = startDate
	this.EndDate = endDate
	return &this
}

// NewJournalStatusResponseJournalsWithDefaults instantiates a new JournalStatusResponseJournals object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJournalStatusResponseJournalsWithDefaults() *JournalStatusResponseJournals {
	this := JournalStatusResponseJournals{}
	return &this
}

// GetId returns the Id field value
func (o *JournalStatusResponseJournals) GetId() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *JournalStatusResponseJournals) GetIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *JournalStatusResponseJournals) SetId(v int32) {
	o.Id = v
}

// GetCompanyId returns the CompanyId field value
func (o *JournalStatusResponseJournals) GetCompanyId() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.CompanyId
}

// GetCompanyIdOk returns a tuple with the CompanyId field value
// and a boolean to check if the value has been set.
func (o *JournalStatusResponseJournals) GetCompanyIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CompanyId, true
}

// SetCompanyId sets field value
func (o *JournalStatusResponseJournals) SetCompanyId(v int32) {
	o.CompanyId = v
}

// GetDownloadType returns the DownloadType field value
func (o *JournalStatusResponseJournals) GetDownloadType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.DownloadType
}

// GetDownloadTypeOk returns a tuple with the DownloadType field value
// and a boolean to check if the value has been set.
func (o *JournalStatusResponseJournals) GetDownloadTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DownloadType, true
}

// SetDownloadType sets field value
func (o *JournalStatusResponseJournals) SetDownloadType(v string) {
	o.DownloadType = v
}

// GetStatus returns the Status field value
func (o *JournalStatusResponseJournals) GetStatus() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *JournalStatusResponseJournals) GetStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *JournalStatusResponseJournals) SetStatus(v string) {
	o.Status = v
}

// GetStartDate returns the StartDate field value
func (o *JournalStatusResponseJournals) GetStartDate() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value
// and a boolean to check if the value has been set.
func (o *JournalStatusResponseJournals) GetStartDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.StartDate, true
}

// SetStartDate sets field value
func (o *JournalStatusResponseJournals) SetStartDate(v string) {
	o.StartDate = v
}

// GetEndDate returns the EndDate field value
func (o *JournalStatusResponseJournals) GetEndDate() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value
// and a boolean to check if the value has been set.
func (o *JournalStatusResponseJournals) GetEndDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EndDate, true
}

// SetEndDate sets field value
func (o *JournalStatusResponseJournals) SetEndDate(v string) {
	o.EndDate = v
}

// GetVisibleTags returns the VisibleTags field value if set, zero value otherwise.
func (o *JournalStatusResponseJournals) GetVisibleTags() []string {
	if o == nil || o.VisibleTags == nil {
		var ret []string
		return ret
	}
	return *o.VisibleTags
}

// GetVisibleTagsOk returns a tuple with the VisibleTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JournalStatusResponseJournals) GetVisibleTagsOk() (*[]string, bool) {
	if o == nil || o.VisibleTags == nil {
		return nil, false
	}
	return o.VisibleTags, true
}

// HasVisibleTags returns a boolean if a field has been set.
func (o *JournalStatusResponseJournals) HasVisibleTags() bool {
	if o != nil && o.VisibleTags != nil {
		return true
	}

	return false
}

// SetVisibleTags gets a reference to the given []string and assigns it to the VisibleTags field.
func (o *JournalStatusResponseJournals) SetVisibleTags(v []string) {
	o.VisibleTags = &v
}

// GetDownloadUrl returns the DownloadUrl field value if set, zero value otherwise.
func (o *JournalStatusResponseJournals) GetDownloadUrl() string {
	if o == nil || o.DownloadUrl == nil {
		var ret string
		return ret
	}
	return *o.DownloadUrl
}

// GetDownloadUrlOk returns a tuple with the DownloadUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JournalStatusResponseJournals) GetDownloadUrlOk() (*string, bool) {
	if o == nil || o.DownloadUrl == nil {
		return nil, false
	}
	return o.DownloadUrl, true
}

// HasDownloadUrl returns a boolean if a field has been set.
func (o *JournalStatusResponseJournals) HasDownloadUrl() bool {
	if o != nil && o.DownloadUrl != nil {
		return true
	}

	return false
}

// SetDownloadUrl gets a reference to the given string and assigns it to the DownloadUrl field.
func (o *JournalStatusResponseJournals) SetDownloadUrl(v string) {
	o.DownloadUrl = &v
}

func (o JournalStatusResponseJournals) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["company_id"] = o.CompanyId
	}
	if true {
		toSerialize["download_type"] = o.DownloadType
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["start_date"] = o.StartDate
	}
	if true {
		toSerialize["end_date"] = o.EndDate
	}
	if o.VisibleTags != nil {
		toSerialize["visible_tags"] = o.VisibleTags
	}
	if o.DownloadUrl != nil {
		toSerialize["download_url"] = o.DownloadUrl
	}
	return json.Marshal(toSerialize)
}

type NullableJournalStatusResponseJournals struct {
	value *JournalStatusResponseJournals
	isSet bool
}

func (v NullableJournalStatusResponseJournals) Get() *JournalStatusResponseJournals {
	return v.value
}

func (v *NullableJournalStatusResponseJournals) Set(val *JournalStatusResponseJournals) {
	v.value = val
	v.isSet = true
}

func (v NullableJournalStatusResponseJournals) IsSet() bool {
	return v.isSet
}

func (v *NullableJournalStatusResponseJournals) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJournalStatusResponseJournals(val *JournalStatusResponseJournals) *NullableJournalStatusResponseJournals {
	return &NullableJournalStatusResponseJournals{value: val, isSet: true}
}

func (v NullableJournalStatusResponseJournals) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJournalStatusResponseJournals) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


