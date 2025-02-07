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

// SelectablesIndexResponseAccountGroups struct for SelectablesIndexResponseAccountGroups
type SelectablesIndexResponseAccountGroups struct {
	// 決算書表示名（小カテゴリー）ID
	Id int32 `json:"id"`
	// 決算書表示名
	Name string `json:"name"`
	// 年度ID
	AccountStructureId int32 `json:"account_structure_id"`
	// 勘定科目カテゴリーID
	AccountCategoryId int32 `json:"account_category_id"`
	// 詳細パラメータの種類
	DetailType *int32 `json:"detail_type,omitempty"`
	// 並び順
	Index int32 `json:"index"`
	// 作成日時
	CreatedAt *string `json:"created_at,omitempty"`
	// 更新日時
	UpdatedAt *string `json:"updated_at,omitempty"`
}

// NewSelectablesIndexResponseAccountGroups instantiates a new SelectablesIndexResponseAccountGroups object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSelectablesIndexResponseAccountGroups(id int32, name string, accountStructureId int32, accountCategoryId int32, index int32, ) *SelectablesIndexResponseAccountGroups {
	this := SelectablesIndexResponseAccountGroups{}
	this.Id = id
	this.Name = name
	this.AccountStructureId = accountStructureId
	this.AccountCategoryId = accountCategoryId
	this.Index = index
	return &this
}

// NewSelectablesIndexResponseAccountGroupsWithDefaults instantiates a new SelectablesIndexResponseAccountGroups object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSelectablesIndexResponseAccountGroupsWithDefaults() *SelectablesIndexResponseAccountGroups {
	this := SelectablesIndexResponseAccountGroups{}
	return &this
}

// GetId returns the Id field value
func (o *SelectablesIndexResponseAccountGroups) GetId() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SelectablesIndexResponseAccountGroups) GetIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SelectablesIndexResponseAccountGroups) SetId(v int32) {
	o.Id = v
}

// GetName returns the Name field value
func (o *SelectablesIndexResponseAccountGroups) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SelectablesIndexResponseAccountGroups) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SelectablesIndexResponseAccountGroups) SetName(v string) {
	o.Name = v
}

// GetAccountStructureId returns the AccountStructureId field value
func (o *SelectablesIndexResponseAccountGroups) GetAccountStructureId() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.AccountStructureId
}

// GetAccountStructureIdOk returns a tuple with the AccountStructureId field value
// and a boolean to check if the value has been set.
func (o *SelectablesIndexResponseAccountGroups) GetAccountStructureIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccountStructureId, true
}

// SetAccountStructureId sets field value
func (o *SelectablesIndexResponseAccountGroups) SetAccountStructureId(v int32) {
	o.AccountStructureId = v
}

// GetAccountCategoryId returns the AccountCategoryId field value
func (o *SelectablesIndexResponseAccountGroups) GetAccountCategoryId() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.AccountCategoryId
}

// GetAccountCategoryIdOk returns a tuple with the AccountCategoryId field value
// and a boolean to check if the value has been set.
func (o *SelectablesIndexResponseAccountGroups) GetAccountCategoryIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccountCategoryId, true
}

// SetAccountCategoryId sets field value
func (o *SelectablesIndexResponseAccountGroups) SetAccountCategoryId(v int32) {
	o.AccountCategoryId = v
}

// GetDetailType returns the DetailType field value if set, zero value otherwise.
func (o *SelectablesIndexResponseAccountGroups) GetDetailType() int32 {
	if o == nil || o.DetailType == nil {
		var ret int32
		return ret
	}
	return *o.DetailType
}

// GetDetailTypeOk returns a tuple with the DetailType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelectablesIndexResponseAccountGroups) GetDetailTypeOk() (*int32, bool) {
	if o == nil || o.DetailType == nil {
		return nil, false
	}
	return o.DetailType, true
}

// HasDetailType returns a boolean if a field has been set.
func (o *SelectablesIndexResponseAccountGroups) HasDetailType() bool {
	if o != nil && o.DetailType != nil {
		return true
	}

	return false
}

// SetDetailType gets a reference to the given int32 and assigns it to the DetailType field.
func (o *SelectablesIndexResponseAccountGroups) SetDetailType(v int32) {
	o.DetailType = &v
}

// GetIndex returns the Index field value
func (o *SelectablesIndexResponseAccountGroups) GetIndex() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *SelectablesIndexResponseAccountGroups) GetIndexOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *SelectablesIndexResponseAccountGroups) SetIndex(v int32) {
	o.Index = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *SelectablesIndexResponseAccountGroups) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelectablesIndexResponseAccountGroups) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *SelectablesIndexResponseAccountGroups) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *SelectablesIndexResponseAccountGroups) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *SelectablesIndexResponseAccountGroups) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelectablesIndexResponseAccountGroups) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *SelectablesIndexResponseAccountGroups) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *SelectablesIndexResponseAccountGroups) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

func (o SelectablesIndexResponseAccountGroups) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["account_structure_id"] = o.AccountStructureId
	}
	if true {
		toSerialize["account_category_id"] = o.AccountCategoryId
	}
	if o.DetailType != nil {
		toSerialize["detail_type"] = o.DetailType
	}
	if true {
		toSerialize["index"] = o.Index
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableSelectablesIndexResponseAccountGroups struct {
	value *SelectablesIndexResponseAccountGroups
	isSet bool
}

func (v NullableSelectablesIndexResponseAccountGroups) Get() *SelectablesIndexResponseAccountGroups {
	return v.value
}

func (v *NullableSelectablesIndexResponseAccountGroups) Set(val *SelectablesIndexResponseAccountGroups) {
	v.value = val
	v.isSet = true
}

func (v NullableSelectablesIndexResponseAccountGroups) IsSet() bool {
	return v.isSet
}

func (v *NullableSelectablesIndexResponseAccountGroups) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSelectablesIndexResponseAccountGroups(val *SelectablesIndexResponseAccountGroups) *NullableSelectablesIndexResponseAccountGroups {
	return &NullableSelectablesIndexResponseAccountGroups{value: val, isSet: true}
}

func (v NullableSelectablesIndexResponseAccountGroups) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSelectablesIndexResponseAccountGroups) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


