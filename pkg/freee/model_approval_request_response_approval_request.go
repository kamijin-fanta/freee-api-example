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

// ApprovalRequestResponseApprovalRequest struct for ApprovalRequestResponseApprovalRequest
type ApprovalRequestResponseApprovalRequest struct {
	// 各種申請ID
	Id int32 `json:"id"`
	// 事業所ID
	CompanyId int32 `json:"company_id"`
	// 申請日 (yyyy-mm-dd)
	ApplicationDate string `json:"application_date"`
	// 申請タイトル
	Title string `json:"title"`
	// 申請者のユーザーID
	ApplicantId int32 `json:"applicant_id"`
	// 承認者（配列）   承認ステップのresource_typeがunspecified (指定なし)の場合はapproversはレスポンスに含まれません。   しかし、resource_typeがunspecifiedの承認ステップにおいて誰かが承認・却下・差し戻しのいずれかのアクションを取った後は、    approversはレスポンスに含まれるようになります。    その場合approversにはアクションを行ったステップのIDとアクションを行ったユーザーのIDが含まれます。
	Approvers []ExpenseApplicationResponseExpenseApplicationApprovers `json:"approvers"`
	// 申請No.
	ApplicationNumber string `json:"application_number"`
	// 申請ステータス(draft:下書き, in_progress:申請中, approved:承認済, rejected:却下, feedback:差戻し)
	Status string `json:"status"`
	// 各種申請の項目一覧（配列）
	RequestItems []ApprovalRequestsIndexResponseRequestItems `json:"request_items"`
	// 申請フォームID
	FormId int32 `json:"form_id"`
	// 申請経路ID
	ApprovalFlowRouteId int32 `json:"approval_flow_route_id"`
	// 各種申請のコメント一覧（配列）
	Comments []ExpenseApplicationResponseExpenseApplicationComments `json:"comments"`
	// 各種申請の承認履歴（配列）
	ApprovalFlowLogs []ExpenseApplicationResponseExpenseApplicationApprovalFlowLogs `json:"approval_flow_logs"`
	// 現在承認ステップID
	CurrentStepId NullableInt32 `json:"current_step_id"`
	// 現在のround。差し戻し等により申請がstepの最初からやり直しになるとroundの値が増えます。
	CurrentRound int32 `json:"current_round"`
	ApprovalRequestForm ApprovalRequestResponseApprovalRequestApprovalRequestForm `json:"approval_request_form"`
}

// NewApprovalRequestResponseApprovalRequest instantiates a new ApprovalRequestResponseApprovalRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApprovalRequestResponseApprovalRequest(id int32, companyId int32, applicationDate string, title string, applicantId int32, approvers []ExpenseApplicationResponseExpenseApplicationApprovers, applicationNumber string, status string, requestItems []ApprovalRequestsIndexResponseRequestItems, formId int32, approvalFlowRouteId int32, comments []ExpenseApplicationResponseExpenseApplicationComments, approvalFlowLogs []ExpenseApplicationResponseExpenseApplicationApprovalFlowLogs, currentStepId NullableInt32, currentRound int32, approvalRequestForm ApprovalRequestResponseApprovalRequestApprovalRequestForm, ) *ApprovalRequestResponseApprovalRequest {
	this := ApprovalRequestResponseApprovalRequest{}
	this.Id = id
	this.CompanyId = companyId
	this.ApplicationDate = applicationDate
	this.Title = title
	this.ApplicantId = applicantId
	this.Approvers = approvers
	this.ApplicationNumber = applicationNumber
	this.Status = status
	this.RequestItems = requestItems
	this.FormId = formId
	this.ApprovalFlowRouteId = approvalFlowRouteId
	this.Comments = comments
	this.ApprovalFlowLogs = approvalFlowLogs
	this.CurrentStepId = currentStepId
	this.CurrentRound = currentRound
	this.ApprovalRequestForm = approvalRequestForm
	return &this
}

// NewApprovalRequestResponseApprovalRequestWithDefaults instantiates a new ApprovalRequestResponseApprovalRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApprovalRequestResponseApprovalRequestWithDefaults() *ApprovalRequestResponseApprovalRequest {
	this := ApprovalRequestResponseApprovalRequest{}
	return &this
}

// GetId returns the Id field value
func (o *ApprovalRequestResponseApprovalRequest) GetId() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ApprovalRequestResponseApprovalRequest) GetIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ApprovalRequestResponseApprovalRequest) SetId(v int32) {
	o.Id = v
}

// GetCompanyId returns the CompanyId field value
func (o *ApprovalRequestResponseApprovalRequest) GetCompanyId() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.CompanyId
}

// GetCompanyIdOk returns a tuple with the CompanyId field value
// and a boolean to check if the value has been set.
func (o *ApprovalRequestResponseApprovalRequest) GetCompanyIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CompanyId, true
}

// SetCompanyId sets field value
func (o *ApprovalRequestResponseApprovalRequest) SetCompanyId(v int32) {
	o.CompanyId = v
}

// GetApplicationDate returns the ApplicationDate field value
func (o *ApprovalRequestResponseApprovalRequest) GetApplicationDate() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.ApplicationDate
}

// GetApplicationDateOk returns a tuple with the ApplicationDate field value
// and a boolean to check if the value has been set.
func (o *ApprovalRequestResponseApprovalRequest) GetApplicationDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ApplicationDate, true
}

// SetApplicationDate sets field value
func (o *ApprovalRequestResponseApprovalRequest) SetApplicationDate(v string) {
	o.ApplicationDate = v
}

// GetTitle returns the Title field value
func (o *ApprovalRequestResponseApprovalRequest) GetTitle() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *ApprovalRequestResponseApprovalRequest) GetTitleOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *ApprovalRequestResponseApprovalRequest) SetTitle(v string) {
	o.Title = v
}

// GetApplicantId returns the ApplicantId field value
func (o *ApprovalRequestResponseApprovalRequest) GetApplicantId() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.ApplicantId
}

// GetApplicantIdOk returns a tuple with the ApplicantId field value
// and a boolean to check if the value has been set.
func (o *ApprovalRequestResponseApprovalRequest) GetApplicantIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ApplicantId, true
}

// SetApplicantId sets field value
func (o *ApprovalRequestResponseApprovalRequest) SetApplicantId(v int32) {
	o.ApplicantId = v
}

// GetApprovers returns the Approvers field value
func (o *ApprovalRequestResponseApprovalRequest) GetApprovers() []ExpenseApplicationResponseExpenseApplicationApprovers {
	if o == nil  {
		var ret []ExpenseApplicationResponseExpenseApplicationApprovers
		return ret
	}

	return o.Approvers
}

// GetApproversOk returns a tuple with the Approvers field value
// and a boolean to check if the value has been set.
func (o *ApprovalRequestResponseApprovalRequest) GetApproversOk() (*[]ExpenseApplicationResponseExpenseApplicationApprovers, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Approvers, true
}

// SetApprovers sets field value
func (o *ApprovalRequestResponseApprovalRequest) SetApprovers(v []ExpenseApplicationResponseExpenseApplicationApprovers) {
	o.Approvers = v
}

// GetApplicationNumber returns the ApplicationNumber field value
func (o *ApprovalRequestResponseApprovalRequest) GetApplicationNumber() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.ApplicationNumber
}

// GetApplicationNumberOk returns a tuple with the ApplicationNumber field value
// and a boolean to check if the value has been set.
func (o *ApprovalRequestResponseApprovalRequest) GetApplicationNumberOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ApplicationNumber, true
}

// SetApplicationNumber sets field value
func (o *ApprovalRequestResponseApprovalRequest) SetApplicationNumber(v string) {
	o.ApplicationNumber = v
}

// GetStatus returns the Status field value
func (o *ApprovalRequestResponseApprovalRequest) GetStatus() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ApprovalRequestResponseApprovalRequest) GetStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ApprovalRequestResponseApprovalRequest) SetStatus(v string) {
	o.Status = v
}

// GetRequestItems returns the RequestItems field value
func (o *ApprovalRequestResponseApprovalRequest) GetRequestItems() []ApprovalRequestsIndexResponseRequestItems {
	if o == nil  {
		var ret []ApprovalRequestsIndexResponseRequestItems
		return ret
	}

	return o.RequestItems
}

// GetRequestItemsOk returns a tuple with the RequestItems field value
// and a boolean to check if the value has been set.
func (o *ApprovalRequestResponseApprovalRequest) GetRequestItemsOk() (*[]ApprovalRequestsIndexResponseRequestItems, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestItems, true
}

// SetRequestItems sets field value
func (o *ApprovalRequestResponseApprovalRequest) SetRequestItems(v []ApprovalRequestsIndexResponseRequestItems) {
	o.RequestItems = v
}

// GetFormId returns the FormId field value
func (o *ApprovalRequestResponseApprovalRequest) GetFormId() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.FormId
}

// GetFormIdOk returns a tuple with the FormId field value
// and a boolean to check if the value has been set.
func (o *ApprovalRequestResponseApprovalRequest) GetFormIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FormId, true
}

// SetFormId sets field value
func (o *ApprovalRequestResponseApprovalRequest) SetFormId(v int32) {
	o.FormId = v
}

// GetApprovalFlowRouteId returns the ApprovalFlowRouteId field value
func (o *ApprovalRequestResponseApprovalRequest) GetApprovalFlowRouteId() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.ApprovalFlowRouteId
}

// GetApprovalFlowRouteIdOk returns a tuple with the ApprovalFlowRouteId field value
// and a boolean to check if the value has been set.
func (o *ApprovalRequestResponseApprovalRequest) GetApprovalFlowRouteIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ApprovalFlowRouteId, true
}

// SetApprovalFlowRouteId sets field value
func (o *ApprovalRequestResponseApprovalRequest) SetApprovalFlowRouteId(v int32) {
	o.ApprovalFlowRouteId = v
}

// GetComments returns the Comments field value
func (o *ApprovalRequestResponseApprovalRequest) GetComments() []ExpenseApplicationResponseExpenseApplicationComments {
	if o == nil  {
		var ret []ExpenseApplicationResponseExpenseApplicationComments
		return ret
	}

	return o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value
// and a boolean to check if the value has been set.
func (o *ApprovalRequestResponseApprovalRequest) GetCommentsOk() (*[]ExpenseApplicationResponseExpenseApplicationComments, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Comments, true
}

// SetComments sets field value
func (o *ApprovalRequestResponseApprovalRequest) SetComments(v []ExpenseApplicationResponseExpenseApplicationComments) {
	o.Comments = v
}

// GetApprovalFlowLogs returns the ApprovalFlowLogs field value
func (o *ApprovalRequestResponseApprovalRequest) GetApprovalFlowLogs() []ExpenseApplicationResponseExpenseApplicationApprovalFlowLogs {
	if o == nil  {
		var ret []ExpenseApplicationResponseExpenseApplicationApprovalFlowLogs
		return ret
	}

	return o.ApprovalFlowLogs
}

// GetApprovalFlowLogsOk returns a tuple with the ApprovalFlowLogs field value
// and a boolean to check if the value has been set.
func (o *ApprovalRequestResponseApprovalRequest) GetApprovalFlowLogsOk() (*[]ExpenseApplicationResponseExpenseApplicationApprovalFlowLogs, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ApprovalFlowLogs, true
}

// SetApprovalFlowLogs sets field value
func (o *ApprovalRequestResponseApprovalRequest) SetApprovalFlowLogs(v []ExpenseApplicationResponseExpenseApplicationApprovalFlowLogs) {
	o.ApprovalFlowLogs = v
}

// GetCurrentStepId returns the CurrentStepId field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *ApprovalRequestResponseApprovalRequest) GetCurrentStepId() int32 {
	if o == nil || o.CurrentStepId.Get() == nil {
		var ret int32
		return ret
	}

	return *o.CurrentStepId.Get()
}

// GetCurrentStepIdOk returns a tuple with the CurrentStepId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApprovalRequestResponseApprovalRequest) GetCurrentStepIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CurrentStepId.Get(), o.CurrentStepId.IsSet()
}

// SetCurrentStepId sets field value
func (o *ApprovalRequestResponseApprovalRequest) SetCurrentStepId(v int32) {
	o.CurrentStepId.Set(&v)
}

// GetCurrentRound returns the CurrentRound field value
func (o *ApprovalRequestResponseApprovalRequest) GetCurrentRound() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.CurrentRound
}

// GetCurrentRoundOk returns a tuple with the CurrentRound field value
// and a boolean to check if the value has been set.
func (o *ApprovalRequestResponseApprovalRequest) GetCurrentRoundOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CurrentRound, true
}

// SetCurrentRound sets field value
func (o *ApprovalRequestResponseApprovalRequest) SetCurrentRound(v int32) {
	o.CurrentRound = v
}

// GetApprovalRequestForm returns the ApprovalRequestForm field value
func (o *ApprovalRequestResponseApprovalRequest) GetApprovalRequestForm() ApprovalRequestResponseApprovalRequestApprovalRequestForm {
	if o == nil  {
		var ret ApprovalRequestResponseApprovalRequestApprovalRequestForm
		return ret
	}

	return o.ApprovalRequestForm
}

// GetApprovalRequestFormOk returns a tuple with the ApprovalRequestForm field value
// and a boolean to check if the value has been set.
func (o *ApprovalRequestResponseApprovalRequest) GetApprovalRequestFormOk() (*ApprovalRequestResponseApprovalRequestApprovalRequestForm, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ApprovalRequestForm, true
}

// SetApprovalRequestForm sets field value
func (o *ApprovalRequestResponseApprovalRequest) SetApprovalRequestForm(v ApprovalRequestResponseApprovalRequestApprovalRequestForm) {
	o.ApprovalRequestForm = v
}

func (o ApprovalRequestResponseApprovalRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["company_id"] = o.CompanyId
	}
	if true {
		toSerialize["application_date"] = o.ApplicationDate
	}
	if true {
		toSerialize["title"] = o.Title
	}
	if true {
		toSerialize["applicant_id"] = o.ApplicantId
	}
	if true {
		toSerialize["approvers"] = o.Approvers
	}
	if true {
		toSerialize["application_number"] = o.ApplicationNumber
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["request_items"] = o.RequestItems
	}
	if true {
		toSerialize["form_id"] = o.FormId
	}
	if true {
		toSerialize["approval_flow_route_id"] = o.ApprovalFlowRouteId
	}
	if true {
		toSerialize["comments"] = o.Comments
	}
	if true {
		toSerialize["approval_flow_logs"] = o.ApprovalFlowLogs
	}
	if true {
		toSerialize["current_step_id"] = o.CurrentStepId.Get()
	}
	if true {
		toSerialize["current_round"] = o.CurrentRound
	}
	if true {
		toSerialize["approval_request_form"] = o.ApprovalRequestForm
	}
	return json.Marshal(toSerialize)
}

type NullableApprovalRequestResponseApprovalRequest struct {
	value *ApprovalRequestResponseApprovalRequest
	isSet bool
}

func (v NullableApprovalRequestResponseApprovalRequest) Get() *ApprovalRequestResponseApprovalRequest {
	return v.value
}

func (v *NullableApprovalRequestResponseApprovalRequest) Set(val *ApprovalRequestResponseApprovalRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApprovalRequestResponseApprovalRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApprovalRequestResponseApprovalRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApprovalRequestResponseApprovalRequest(val *ApprovalRequestResponseApprovalRequest) *NullableApprovalRequestResponseApprovalRequest {
	return &NullableApprovalRequestResponseApprovalRequest{value: val, isSet: true}
}

func (v NullableApprovalRequestResponseApprovalRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApprovalRequestResponseApprovalRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


