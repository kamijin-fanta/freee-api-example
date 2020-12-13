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

// ApprovalFlowRoutesIndexResponseApprovalFlowRoutes struct for ApprovalFlowRoutesIndexResponseApprovalFlowRoutes
type ApprovalFlowRoutesIndexResponseApprovalFlowRoutes struct {
	// 申請経路ID
	Id int32 `json:"id"`
	// 申請経路名
	Name *string `json:"name,omitempty"`
	// 申請経路の説明
	Description *string `json:"description,omitempty"`
	// 更新したユーザーのユーザーID
	UserId NullableInt32 `json:"user_id,omitempty"`
	// システム作成の申請経路かどうか
	DefinitionSystem *bool `json:"definition_system,omitempty"`
	// 最初の承認ステップのID
	FirstStepId *int32 `json:"first_step_id,omitempty"`
	// 申請種別（申請経路を使用できる申請種別を示します。例えば、ApprovalRequest の場合は、各種申請で使用できる申請経路です。） * `TxnApproval` - 仕訳承認 * `ExpenseApplication` - 経費精算 * `PaymentRequest` - 支払依頼 * `ApprovalRequest` - 各種申請 * `DocApproval` - 請求書等 (見積書・納品書・請求書・発注書)
	Usages *[]string `json:"usages,omitempty"`
	// 申請経路で利用できる申請フォームID配列
	RequestFormIds *[]int32 `json:"request_form_ids,omitempty"`
	// 基本経路として設定されているかどうか<br><br> リクエストパラメータusageに下記のいずれかが指定され、かつ、基本経路の場合はtrueになります。 * `TxnApproval` - 仕訳承認 * `ExpenseApplication` - 経費精算 * `PaymentRequest` - 支払依頼 * `ApprovalRequest`(リクエストパラメータrequest_form_idを同時に指定) - 各種申請 * `DocApproval` - 請求書等 (見積書・納品書・請求書・発注書)  <a href=\"https://support.freee.co.jp/hc/ja/articles/900000507963-%E7%94%B3%E8%AB%8B%E3%83%95%E3%82%A9%E3%83%BC%E3%83%A0%E3%81%AE%E5%9F%BA%E6%9C%AC%E7%B5%8C%E8%B7%AF%E8%A8%AD%E5%AE%9A\" target=\"_blank\">申請フォームの基本経路設定</a> 
	DefaultRoute bool `json:"default_route"`
}

// NewApprovalFlowRoutesIndexResponseApprovalFlowRoutes instantiates a new ApprovalFlowRoutesIndexResponseApprovalFlowRoutes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApprovalFlowRoutesIndexResponseApprovalFlowRoutes(id int32, defaultRoute bool, ) *ApprovalFlowRoutesIndexResponseApprovalFlowRoutes {
	this := ApprovalFlowRoutesIndexResponseApprovalFlowRoutes{}
	this.Id = id
	this.DefaultRoute = defaultRoute
	return &this
}

// NewApprovalFlowRoutesIndexResponseApprovalFlowRoutesWithDefaults instantiates a new ApprovalFlowRoutesIndexResponseApprovalFlowRoutes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApprovalFlowRoutesIndexResponseApprovalFlowRoutesWithDefaults() *ApprovalFlowRoutesIndexResponseApprovalFlowRoutes {
	this := ApprovalFlowRoutesIndexResponseApprovalFlowRoutes{}
	return &this
}

// GetId returns the Id field value
func (o *ApprovalFlowRoutesIndexResponseApprovalFlowRoutes) GetId() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ApprovalFlowRoutesIndexResponseApprovalFlowRoutes) GetIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ApprovalFlowRoutesIndexResponseApprovalFlowRoutes) SetId(v int32) {
	o.Id = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ApprovalFlowRoutesIndexResponseApprovalFlowRoutes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApprovalFlowRoutesIndexResponseApprovalFlowRoutes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ApprovalFlowRoutesIndexResponseApprovalFlowRoutes) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ApprovalFlowRoutesIndexResponseApprovalFlowRoutes) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ApprovalFlowRoutesIndexResponseApprovalFlowRoutes) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApprovalFlowRoutesIndexResponseApprovalFlowRoutes) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ApprovalFlowRoutesIndexResponseApprovalFlowRoutes) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ApprovalFlowRoutesIndexResponseApprovalFlowRoutes) SetDescription(v string) {
	o.Description = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApprovalFlowRoutesIndexResponseApprovalFlowRoutes) GetUserId() int32 {
	if o == nil || o.UserId.Get() == nil {
		var ret int32
		return ret
	}
	return *o.UserId.Get()
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApprovalFlowRoutesIndexResponseApprovalFlowRoutes) GetUserIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UserId.Get(), o.UserId.IsSet()
}

// HasUserId returns a boolean if a field has been set.
func (o *ApprovalFlowRoutesIndexResponseApprovalFlowRoutes) HasUserId() bool {
	if o != nil && o.UserId.IsSet() {
		return true
	}

	return false
}

// SetUserId gets a reference to the given NullableInt32 and assigns it to the UserId field.
func (o *ApprovalFlowRoutesIndexResponseApprovalFlowRoutes) SetUserId(v int32) {
	o.UserId.Set(&v)
}
// SetUserIdNil sets the value for UserId to be an explicit nil
func (o *ApprovalFlowRoutesIndexResponseApprovalFlowRoutes) SetUserIdNil() {
	o.UserId.Set(nil)
}

// UnsetUserId ensures that no value is present for UserId, not even an explicit nil
func (o *ApprovalFlowRoutesIndexResponseApprovalFlowRoutes) UnsetUserId() {
	o.UserId.Unset()
}

// GetDefinitionSystem returns the DefinitionSystem field value if set, zero value otherwise.
func (o *ApprovalFlowRoutesIndexResponseApprovalFlowRoutes) GetDefinitionSystem() bool {
	if o == nil || o.DefinitionSystem == nil {
		var ret bool
		return ret
	}
	return *o.DefinitionSystem
}

// GetDefinitionSystemOk returns a tuple with the DefinitionSystem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApprovalFlowRoutesIndexResponseApprovalFlowRoutes) GetDefinitionSystemOk() (*bool, bool) {
	if o == nil || o.DefinitionSystem == nil {
		return nil, false
	}
	return o.DefinitionSystem, true
}

// HasDefinitionSystem returns a boolean if a field has been set.
func (o *ApprovalFlowRoutesIndexResponseApprovalFlowRoutes) HasDefinitionSystem() bool {
	if o != nil && o.DefinitionSystem != nil {
		return true
	}

	return false
}

// SetDefinitionSystem gets a reference to the given bool and assigns it to the DefinitionSystem field.
func (o *ApprovalFlowRoutesIndexResponseApprovalFlowRoutes) SetDefinitionSystem(v bool) {
	o.DefinitionSystem = &v
}

// GetFirstStepId returns the FirstStepId field value if set, zero value otherwise.
func (o *ApprovalFlowRoutesIndexResponseApprovalFlowRoutes) GetFirstStepId() int32 {
	if o == nil || o.FirstStepId == nil {
		var ret int32
		return ret
	}
	return *o.FirstStepId
}

// GetFirstStepIdOk returns a tuple with the FirstStepId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApprovalFlowRoutesIndexResponseApprovalFlowRoutes) GetFirstStepIdOk() (*int32, bool) {
	if o == nil || o.FirstStepId == nil {
		return nil, false
	}
	return o.FirstStepId, true
}

// HasFirstStepId returns a boolean if a field has been set.
func (o *ApprovalFlowRoutesIndexResponseApprovalFlowRoutes) HasFirstStepId() bool {
	if o != nil && o.FirstStepId != nil {
		return true
	}

	return false
}

// SetFirstStepId gets a reference to the given int32 and assigns it to the FirstStepId field.
func (o *ApprovalFlowRoutesIndexResponseApprovalFlowRoutes) SetFirstStepId(v int32) {
	o.FirstStepId = &v
}

// GetUsages returns the Usages field value if set, zero value otherwise.
func (o *ApprovalFlowRoutesIndexResponseApprovalFlowRoutes) GetUsages() []string {
	if o == nil || o.Usages == nil {
		var ret []string
		return ret
	}
	return *o.Usages
}

// GetUsagesOk returns a tuple with the Usages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApprovalFlowRoutesIndexResponseApprovalFlowRoutes) GetUsagesOk() (*[]string, bool) {
	if o == nil || o.Usages == nil {
		return nil, false
	}
	return o.Usages, true
}

// HasUsages returns a boolean if a field has been set.
func (o *ApprovalFlowRoutesIndexResponseApprovalFlowRoutes) HasUsages() bool {
	if o != nil && o.Usages != nil {
		return true
	}

	return false
}

// SetUsages gets a reference to the given []string and assigns it to the Usages field.
func (o *ApprovalFlowRoutesIndexResponseApprovalFlowRoutes) SetUsages(v []string) {
	o.Usages = &v
}

// GetRequestFormIds returns the RequestFormIds field value if set, zero value otherwise.
func (o *ApprovalFlowRoutesIndexResponseApprovalFlowRoutes) GetRequestFormIds() []int32 {
	if o == nil || o.RequestFormIds == nil {
		var ret []int32
		return ret
	}
	return *o.RequestFormIds
}

// GetRequestFormIdsOk returns a tuple with the RequestFormIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApprovalFlowRoutesIndexResponseApprovalFlowRoutes) GetRequestFormIdsOk() (*[]int32, bool) {
	if o == nil || o.RequestFormIds == nil {
		return nil, false
	}
	return o.RequestFormIds, true
}

// HasRequestFormIds returns a boolean if a field has been set.
func (o *ApprovalFlowRoutesIndexResponseApprovalFlowRoutes) HasRequestFormIds() bool {
	if o != nil && o.RequestFormIds != nil {
		return true
	}

	return false
}

// SetRequestFormIds gets a reference to the given []int32 and assigns it to the RequestFormIds field.
func (o *ApprovalFlowRoutesIndexResponseApprovalFlowRoutes) SetRequestFormIds(v []int32) {
	o.RequestFormIds = &v
}

// GetDefaultRoute returns the DefaultRoute field value
func (o *ApprovalFlowRoutesIndexResponseApprovalFlowRoutes) GetDefaultRoute() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.DefaultRoute
}

// GetDefaultRouteOk returns a tuple with the DefaultRoute field value
// and a boolean to check if the value has been set.
func (o *ApprovalFlowRoutesIndexResponseApprovalFlowRoutes) GetDefaultRouteOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DefaultRoute, true
}

// SetDefaultRoute sets field value
func (o *ApprovalFlowRoutesIndexResponseApprovalFlowRoutes) SetDefaultRoute(v bool) {
	o.DefaultRoute = v
}

func (o ApprovalFlowRoutesIndexResponseApprovalFlowRoutes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.UserId.IsSet() {
		toSerialize["user_id"] = o.UserId.Get()
	}
	if o.DefinitionSystem != nil {
		toSerialize["definition_system"] = o.DefinitionSystem
	}
	if o.FirstStepId != nil {
		toSerialize["first_step_id"] = o.FirstStepId
	}
	if o.Usages != nil {
		toSerialize["usages"] = o.Usages
	}
	if o.RequestFormIds != nil {
		toSerialize["request_form_ids"] = o.RequestFormIds
	}
	if true {
		toSerialize["default_route"] = o.DefaultRoute
	}
	return json.Marshal(toSerialize)
}

type NullableApprovalFlowRoutesIndexResponseApprovalFlowRoutes struct {
	value *ApprovalFlowRoutesIndexResponseApprovalFlowRoutes
	isSet bool
}

func (v NullableApprovalFlowRoutesIndexResponseApprovalFlowRoutes) Get() *ApprovalFlowRoutesIndexResponseApprovalFlowRoutes {
	return v.value
}

func (v *NullableApprovalFlowRoutesIndexResponseApprovalFlowRoutes) Set(val *ApprovalFlowRoutesIndexResponseApprovalFlowRoutes) {
	v.value = val
	v.isSet = true
}

func (v NullableApprovalFlowRoutesIndexResponseApprovalFlowRoutes) IsSet() bool {
	return v.isSet
}

func (v *NullableApprovalFlowRoutesIndexResponseApprovalFlowRoutes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApprovalFlowRoutesIndexResponseApprovalFlowRoutes(val *ApprovalFlowRoutesIndexResponseApprovalFlowRoutes) *NullableApprovalFlowRoutesIndexResponseApprovalFlowRoutes {
	return &NullableApprovalFlowRoutesIndexResponseApprovalFlowRoutes{value: val, isSet: true}
}

func (v NullableApprovalFlowRoutesIndexResponseApprovalFlowRoutes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApprovalFlowRoutesIndexResponseApprovalFlowRoutes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


