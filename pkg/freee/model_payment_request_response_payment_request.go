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

// PaymentRequestResponsePaymentRequest struct for PaymentRequestResponsePaymentRequest
type PaymentRequestResponsePaymentRequest struct {
	// 支払依頼ID
	Id int32 `json:"id"`
	// 事業所ID
	CompanyId int32 `json:"company_id"`
	// 申請タイトル
	Title string `json:"title"`
	// 申請日 (yyyy-mm-dd)
	ApplicationDate string `json:"application_date"`
	// 備考
	Description NullableString `json:"description"`
	// 合計金額
	TotalAmount int32 `json:"total_amount"`
	// 申請ステータス(draft:下書き, in_progress:申請中, approved:承認済, rejected:却下, feedback:差戻し)
	Status string `json:"status"`
	// 支払依頼の項目行一覧（配列）
	PaymentRequestLines []PaymentRequestResponsePaymentRequestPaymentRequestLines `json:"payment_request_lines"`
	// 取引ID (申請ステータス:statusがapprovedで、取引が存在する時のみdeal_idが表示されます)
	DealId NullableInt32 `json:"deal_id"`
	// 取引ステータス (申請ステータス:statusがapprovedで、取引が存在する時のみdeal_statusが表示されます settled:精算済み, unsettled:清算待ち)
	DealStatus NullableString `json:"deal_status"`
	// 申請者のユーザーID
	ApplicantId int32 `json:"applicant_id"`
	// 承認者（配列）   承認ステップのresource_typeがunspecified (指定なし)の場合はapproversはレスポンスに含まれません。   しかし、resource_typeがunspecifiedの承認ステップにおいて誰かが承認・却下・差し戻しのいずれかのアクションを取った後は、    approversはレスポンスに含まれるようになります。    その場合approversにはアクションを行ったステップのIDとアクションを行ったユーザーのIDが含まれます。
	Approvers []ExpenseApplicationResponseExpenseApplicationApprovers `json:"approvers"`
	// 申請No.
	ApplicationNumber string `json:"application_number"`
	// 申請経路ID
	ApprovalFlowRouteId int32 `json:"approval_flow_route_id"`
	// 支払依頼のコメント一覧（配列）
	Comments []ExpenseApplicationResponseExpenseApplicationComments `json:"comments"`
	// 支払依頼の承認履歴（配列）
	ApprovalFlowLogs []ExpenseApplicationResponseExpenseApplicationApprovalFlowLogs `json:"approval_flow_logs"`
	// 現在承認ステップID
	CurrentStepId NullableInt32 `json:"current_step_id"`
	// 現在のround。差し戻し等により申請がstepの最初からやり直しになるとroundの値が増えます。
	CurrentRound int32 `json:"current_round"`
	// 請求書番号
	DocumentCode string `json:"document_code"`
	// 証憑ID
	ReceiptId NullableInt32 `json:"receipt_id,omitempty"`
	// 発生日 (yyyy-mm-dd)
	IssueDate string `json:"issue_date"`
	// 支払期限 (yyyy-mm-dd)
	PaymentDate NullableString `json:"payment_date"`
	// 支払方法(none: 指定なし, domestic_bank_transfer: 国内振込, abroad_bank_transfer: 国外振込, account_transfer: 口座振替, credit_card: クレジットカード)
	PaymentMethod string `json:"payment_method"`
	// 取引先ID
	PartnerId NullableInt32 `json:"partner_id"`
	// 取引先コード
	PartnerCode NullableString `json:"partner_code,omitempty"`
	// 取引先名
	PartnerName NullableString `json:"partner_name"`
	// 銀行名
	BankName string `json:"bank_name"`
	// 銀行名（カナ）
	BankNameKana string `json:"bank_name_kana"`
	// 銀行コード
	BankCode string `json:"bank_code"`
	// 支店名
	BranchName string `json:"branch_name"`
	// 支店名（カナ）
	BranchKana string `json:"branch_kana"`
	// 支店番号
	BranchCode string `json:"branch_code"`
	// 口座種別(ordinary:普通、checking:当座、earmarked:納税準備預金、savings:貯蓄、other:その他)
	AccountType string `json:"account_type"`
	// 口座番号
	AccountNumber string `json:"account_number"`
	// 受取人名（カナ）
	AccountName string `json:"account_name"`
}

// NewPaymentRequestResponsePaymentRequest instantiates a new PaymentRequestResponsePaymentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentRequestResponsePaymentRequest(id int32, companyId int32, title string, applicationDate string, description NullableString, totalAmount int32, status string, paymentRequestLines []PaymentRequestResponsePaymentRequestPaymentRequestLines, dealId NullableInt32, dealStatus NullableString, applicantId int32, approvers []ExpenseApplicationResponseExpenseApplicationApprovers, applicationNumber string, approvalFlowRouteId int32, comments []ExpenseApplicationResponseExpenseApplicationComments, approvalFlowLogs []ExpenseApplicationResponseExpenseApplicationApprovalFlowLogs, currentStepId NullableInt32, currentRound int32, documentCode string, issueDate string, paymentDate NullableString, paymentMethod string, partnerId NullableInt32, partnerName NullableString, bankName string, bankNameKana string, bankCode string, branchName string, branchKana string, branchCode string, accountType string, accountNumber string, accountName string, ) *PaymentRequestResponsePaymentRequest {
	this := PaymentRequestResponsePaymentRequest{}
	this.Id = id
	this.CompanyId = companyId
	this.Title = title
	this.ApplicationDate = applicationDate
	this.Description = description
	this.TotalAmount = totalAmount
	this.Status = status
	this.PaymentRequestLines = paymentRequestLines
	this.DealId = dealId
	this.DealStatus = dealStatus
	this.ApplicantId = applicantId
	this.Approvers = approvers
	this.ApplicationNumber = applicationNumber
	this.ApprovalFlowRouteId = approvalFlowRouteId
	this.Comments = comments
	this.ApprovalFlowLogs = approvalFlowLogs
	this.CurrentStepId = currentStepId
	this.CurrentRound = currentRound
	this.DocumentCode = documentCode
	this.IssueDate = issueDate
	this.PaymentDate = paymentDate
	this.PaymentMethod = paymentMethod
	this.PartnerId = partnerId
	this.PartnerName = partnerName
	this.BankName = bankName
	this.BankNameKana = bankNameKana
	this.BankCode = bankCode
	this.BranchName = branchName
	this.BranchKana = branchKana
	this.BranchCode = branchCode
	this.AccountType = accountType
	this.AccountNumber = accountNumber
	this.AccountName = accountName
	return &this
}

// NewPaymentRequestResponsePaymentRequestWithDefaults instantiates a new PaymentRequestResponsePaymentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentRequestResponsePaymentRequestWithDefaults() *PaymentRequestResponsePaymentRequest {
	this := PaymentRequestResponsePaymentRequest{}
	return &this
}

// GetId returns the Id field value
func (o *PaymentRequestResponsePaymentRequest) GetId() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PaymentRequestResponsePaymentRequest) GetIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PaymentRequestResponsePaymentRequest) SetId(v int32) {
	o.Id = v
}

// GetCompanyId returns the CompanyId field value
func (o *PaymentRequestResponsePaymentRequest) GetCompanyId() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.CompanyId
}

// GetCompanyIdOk returns a tuple with the CompanyId field value
// and a boolean to check if the value has been set.
func (o *PaymentRequestResponsePaymentRequest) GetCompanyIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CompanyId, true
}

// SetCompanyId sets field value
func (o *PaymentRequestResponsePaymentRequest) SetCompanyId(v int32) {
	o.CompanyId = v
}

// GetTitle returns the Title field value
func (o *PaymentRequestResponsePaymentRequest) GetTitle() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *PaymentRequestResponsePaymentRequest) GetTitleOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *PaymentRequestResponsePaymentRequest) SetTitle(v string) {
	o.Title = v
}

// GetApplicationDate returns the ApplicationDate field value
func (o *PaymentRequestResponsePaymentRequest) GetApplicationDate() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.ApplicationDate
}

// GetApplicationDateOk returns a tuple with the ApplicationDate field value
// and a boolean to check if the value has been set.
func (o *PaymentRequestResponsePaymentRequest) GetApplicationDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ApplicationDate, true
}

// SetApplicationDate sets field value
func (o *PaymentRequestResponsePaymentRequest) SetApplicationDate(v string) {
	o.ApplicationDate = v
}

// GetDescription returns the Description field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PaymentRequestResponsePaymentRequest) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}

	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentRequestResponsePaymentRequest) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// SetDescription sets field value
func (o *PaymentRequestResponsePaymentRequest) SetDescription(v string) {
	o.Description.Set(&v)
}

// GetTotalAmount returns the TotalAmount field value
func (o *PaymentRequestResponsePaymentRequest) GetTotalAmount() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.TotalAmount
}

// GetTotalAmountOk returns a tuple with the TotalAmount field value
// and a boolean to check if the value has been set.
func (o *PaymentRequestResponsePaymentRequest) GetTotalAmountOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TotalAmount, true
}

// SetTotalAmount sets field value
func (o *PaymentRequestResponsePaymentRequest) SetTotalAmount(v int32) {
	o.TotalAmount = v
}

// GetStatus returns the Status field value
func (o *PaymentRequestResponsePaymentRequest) GetStatus() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *PaymentRequestResponsePaymentRequest) GetStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *PaymentRequestResponsePaymentRequest) SetStatus(v string) {
	o.Status = v
}

// GetPaymentRequestLines returns the PaymentRequestLines field value
func (o *PaymentRequestResponsePaymentRequest) GetPaymentRequestLines() []PaymentRequestResponsePaymentRequestPaymentRequestLines {
	if o == nil  {
		var ret []PaymentRequestResponsePaymentRequestPaymentRequestLines
		return ret
	}

	return o.PaymentRequestLines
}

// GetPaymentRequestLinesOk returns a tuple with the PaymentRequestLines field value
// and a boolean to check if the value has been set.
func (o *PaymentRequestResponsePaymentRequest) GetPaymentRequestLinesOk() (*[]PaymentRequestResponsePaymentRequestPaymentRequestLines, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PaymentRequestLines, true
}

// SetPaymentRequestLines sets field value
func (o *PaymentRequestResponsePaymentRequest) SetPaymentRequestLines(v []PaymentRequestResponsePaymentRequestPaymentRequestLines) {
	o.PaymentRequestLines = v
}

// GetDealId returns the DealId field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *PaymentRequestResponsePaymentRequest) GetDealId() int32 {
	if o == nil || o.DealId.Get() == nil {
		var ret int32
		return ret
	}

	return *o.DealId.Get()
}

// GetDealIdOk returns a tuple with the DealId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentRequestResponsePaymentRequest) GetDealIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DealId.Get(), o.DealId.IsSet()
}

// SetDealId sets field value
func (o *PaymentRequestResponsePaymentRequest) SetDealId(v int32) {
	o.DealId.Set(&v)
}

// GetDealStatus returns the DealStatus field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PaymentRequestResponsePaymentRequest) GetDealStatus() string {
	if o == nil || o.DealStatus.Get() == nil {
		var ret string
		return ret
	}

	return *o.DealStatus.Get()
}

// GetDealStatusOk returns a tuple with the DealStatus field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentRequestResponsePaymentRequest) GetDealStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DealStatus.Get(), o.DealStatus.IsSet()
}

// SetDealStatus sets field value
func (o *PaymentRequestResponsePaymentRequest) SetDealStatus(v string) {
	o.DealStatus.Set(&v)
}

// GetApplicantId returns the ApplicantId field value
func (o *PaymentRequestResponsePaymentRequest) GetApplicantId() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.ApplicantId
}

// GetApplicantIdOk returns a tuple with the ApplicantId field value
// and a boolean to check if the value has been set.
func (o *PaymentRequestResponsePaymentRequest) GetApplicantIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ApplicantId, true
}

// SetApplicantId sets field value
func (o *PaymentRequestResponsePaymentRequest) SetApplicantId(v int32) {
	o.ApplicantId = v
}

// GetApprovers returns the Approvers field value
func (o *PaymentRequestResponsePaymentRequest) GetApprovers() []ExpenseApplicationResponseExpenseApplicationApprovers {
	if o == nil  {
		var ret []ExpenseApplicationResponseExpenseApplicationApprovers
		return ret
	}

	return o.Approvers
}

// GetApproversOk returns a tuple with the Approvers field value
// and a boolean to check if the value has been set.
func (o *PaymentRequestResponsePaymentRequest) GetApproversOk() (*[]ExpenseApplicationResponseExpenseApplicationApprovers, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Approvers, true
}

// SetApprovers sets field value
func (o *PaymentRequestResponsePaymentRequest) SetApprovers(v []ExpenseApplicationResponseExpenseApplicationApprovers) {
	o.Approvers = v
}

// GetApplicationNumber returns the ApplicationNumber field value
func (o *PaymentRequestResponsePaymentRequest) GetApplicationNumber() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.ApplicationNumber
}

// GetApplicationNumberOk returns a tuple with the ApplicationNumber field value
// and a boolean to check if the value has been set.
func (o *PaymentRequestResponsePaymentRequest) GetApplicationNumberOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ApplicationNumber, true
}

// SetApplicationNumber sets field value
func (o *PaymentRequestResponsePaymentRequest) SetApplicationNumber(v string) {
	o.ApplicationNumber = v
}

// GetApprovalFlowRouteId returns the ApprovalFlowRouteId field value
func (o *PaymentRequestResponsePaymentRequest) GetApprovalFlowRouteId() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.ApprovalFlowRouteId
}

// GetApprovalFlowRouteIdOk returns a tuple with the ApprovalFlowRouteId field value
// and a boolean to check if the value has been set.
func (o *PaymentRequestResponsePaymentRequest) GetApprovalFlowRouteIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ApprovalFlowRouteId, true
}

// SetApprovalFlowRouteId sets field value
func (o *PaymentRequestResponsePaymentRequest) SetApprovalFlowRouteId(v int32) {
	o.ApprovalFlowRouteId = v
}

// GetComments returns the Comments field value
func (o *PaymentRequestResponsePaymentRequest) GetComments() []ExpenseApplicationResponseExpenseApplicationComments {
	if o == nil  {
		var ret []ExpenseApplicationResponseExpenseApplicationComments
		return ret
	}

	return o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value
// and a boolean to check if the value has been set.
func (o *PaymentRequestResponsePaymentRequest) GetCommentsOk() (*[]ExpenseApplicationResponseExpenseApplicationComments, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Comments, true
}

// SetComments sets field value
func (o *PaymentRequestResponsePaymentRequest) SetComments(v []ExpenseApplicationResponseExpenseApplicationComments) {
	o.Comments = v
}

// GetApprovalFlowLogs returns the ApprovalFlowLogs field value
func (o *PaymentRequestResponsePaymentRequest) GetApprovalFlowLogs() []ExpenseApplicationResponseExpenseApplicationApprovalFlowLogs {
	if o == nil  {
		var ret []ExpenseApplicationResponseExpenseApplicationApprovalFlowLogs
		return ret
	}

	return o.ApprovalFlowLogs
}

// GetApprovalFlowLogsOk returns a tuple with the ApprovalFlowLogs field value
// and a boolean to check if the value has been set.
func (o *PaymentRequestResponsePaymentRequest) GetApprovalFlowLogsOk() (*[]ExpenseApplicationResponseExpenseApplicationApprovalFlowLogs, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ApprovalFlowLogs, true
}

// SetApprovalFlowLogs sets field value
func (o *PaymentRequestResponsePaymentRequest) SetApprovalFlowLogs(v []ExpenseApplicationResponseExpenseApplicationApprovalFlowLogs) {
	o.ApprovalFlowLogs = v
}

// GetCurrentStepId returns the CurrentStepId field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *PaymentRequestResponsePaymentRequest) GetCurrentStepId() int32 {
	if o == nil || o.CurrentStepId.Get() == nil {
		var ret int32
		return ret
	}

	return *o.CurrentStepId.Get()
}

// GetCurrentStepIdOk returns a tuple with the CurrentStepId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentRequestResponsePaymentRequest) GetCurrentStepIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CurrentStepId.Get(), o.CurrentStepId.IsSet()
}

// SetCurrentStepId sets field value
func (o *PaymentRequestResponsePaymentRequest) SetCurrentStepId(v int32) {
	o.CurrentStepId.Set(&v)
}

// GetCurrentRound returns the CurrentRound field value
func (o *PaymentRequestResponsePaymentRequest) GetCurrentRound() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.CurrentRound
}

// GetCurrentRoundOk returns a tuple with the CurrentRound field value
// and a boolean to check if the value has been set.
func (o *PaymentRequestResponsePaymentRequest) GetCurrentRoundOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CurrentRound, true
}

// SetCurrentRound sets field value
func (o *PaymentRequestResponsePaymentRequest) SetCurrentRound(v int32) {
	o.CurrentRound = v
}

// GetDocumentCode returns the DocumentCode field value
func (o *PaymentRequestResponsePaymentRequest) GetDocumentCode() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.DocumentCode
}

// GetDocumentCodeOk returns a tuple with the DocumentCode field value
// and a boolean to check if the value has been set.
func (o *PaymentRequestResponsePaymentRequest) GetDocumentCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DocumentCode, true
}

// SetDocumentCode sets field value
func (o *PaymentRequestResponsePaymentRequest) SetDocumentCode(v string) {
	o.DocumentCode = v
}

// GetReceiptId returns the ReceiptId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentRequestResponsePaymentRequest) GetReceiptId() int32 {
	if o == nil || o.ReceiptId.Get() == nil {
		var ret int32
		return ret
	}
	return *o.ReceiptId.Get()
}

// GetReceiptIdOk returns a tuple with the ReceiptId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentRequestResponsePaymentRequest) GetReceiptIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ReceiptId.Get(), o.ReceiptId.IsSet()
}

// HasReceiptId returns a boolean if a field has been set.
func (o *PaymentRequestResponsePaymentRequest) HasReceiptId() bool {
	if o != nil && o.ReceiptId.IsSet() {
		return true
	}

	return false
}

// SetReceiptId gets a reference to the given NullableInt32 and assigns it to the ReceiptId field.
func (o *PaymentRequestResponsePaymentRequest) SetReceiptId(v int32) {
	o.ReceiptId.Set(&v)
}
// SetReceiptIdNil sets the value for ReceiptId to be an explicit nil
func (o *PaymentRequestResponsePaymentRequest) SetReceiptIdNil() {
	o.ReceiptId.Set(nil)
}

// UnsetReceiptId ensures that no value is present for ReceiptId, not even an explicit nil
func (o *PaymentRequestResponsePaymentRequest) UnsetReceiptId() {
	o.ReceiptId.Unset()
}

// GetIssueDate returns the IssueDate field value
func (o *PaymentRequestResponsePaymentRequest) GetIssueDate() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.IssueDate
}

// GetIssueDateOk returns a tuple with the IssueDate field value
// and a boolean to check if the value has been set.
func (o *PaymentRequestResponsePaymentRequest) GetIssueDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IssueDate, true
}

// SetIssueDate sets field value
func (o *PaymentRequestResponsePaymentRequest) SetIssueDate(v string) {
	o.IssueDate = v
}

// GetPaymentDate returns the PaymentDate field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PaymentRequestResponsePaymentRequest) GetPaymentDate() string {
	if o == nil || o.PaymentDate.Get() == nil {
		var ret string
		return ret
	}

	return *o.PaymentDate.Get()
}

// GetPaymentDateOk returns a tuple with the PaymentDate field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentRequestResponsePaymentRequest) GetPaymentDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PaymentDate.Get(), o.PaymentDate.IsSet()
}

// SetPaymentDate sets field value
func (o *PaymentRequestResponsePaymentRequest) SetPaymentDate(v string) {
	o.PaymentDate.Set(&v)
}

// GetPaymentMethod returns the PaymentMethod field value
func (o *PaymentRequestResponsePaymentRequest) GetPaymentMethod() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.PaymentMethod
}

// GetPaymentMethodOk returns a tuple with the PaymentMethod field value
// and a boolean to check if the value has been set.
func (o *PaymentRequestResponsePaymentRequest) GetPaymentMethodOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PaymentMethod, true
}

// SetPaymentMethod sets field value
func (o *PaymentRequestResponsePaymentRequest) SetPaymentMethod(v string) {
	o.PaymentMethod = v
}

// GetPartnerId returns the PartnerId field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *PaymentRequestResponsePaymentRequest) GetPartnerId() int32 {
	if o == nil || o.PartnerId.Get() == nil {
		var ret int32
		return ret
	}

	return *o.PartnerId.Get()
}

// GetPartnerIdOk returns a tuple with the PartnerId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentRequestResponsePaymentRequest) GetPartnerIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PartnerId.Get(), o.PartnerId.IsSet()
}

// SetPartnerId sets field value
func (o *PaymentRequestResponsePaymentRequest) SetPartnerId(v int32) {
	o.PartnerId.Set(&v)
}

// GetPartnerCode returns the PartnerCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentRequestResponsePaymentRequest) GetPartnerCode() string {
	if o == nil || o.PartnerCode.Get() == nil {
		var ret string
		return ret
	}
	return *o.PartnerCode.Get()
}

// GetPartnerCodeOk returns a tuple with the PartnerCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentRequestResponsePaymentRequest) GetPartnerCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PartnerCode.Get(), o.PartnerCode.IsSet()
}

// HasPartnerCode returns a boolean if a field has been set.
func (o *PaymentRequestResponsePaymentRequest) HasPartnerCode() bool {
	if o != nil && o.PartnerCode.IsSet() {
		return true
	}

	return false
}

// SetPartnerCode gets a reference to the given NullableString and assigns it to the PartnerCode field.
func (o *PaymentRequestResponsePaymentRequest) SetPartnerCode(v string) {
	o.PartnerCode.Set(&v)
}
// SetPartnerCodeNil sets the value for PartnerCode to be an explicit nil
func (o *PaymentRequestResponsePaymentRequest) SetPartnerCodeNil() {
	o.PartnerCode.Set(nil)
}

// UnsetPartnerCode ensures that no value is present for PartnerCode, not even an explicit nil
func (o *PaymentRequestResponsePaymentRequest) UnsetPartnerCode() {
	o.PartnerCode.Unset()
}

// GetPartnerName returns the PartnerName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PaymentRequestResponsePaymentRequest) GetPartnerName() string {
	if o == nil || o.PartnerName.Get() == nil {
		var ret string
		return ret
	}

	return *o.PartnerName.Get()
}

// GetPartnerNameOk returns a tuple with the PartnerName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentRequestResponsePaymentRequest) GetPartnerNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PartnerName.Get(), o.PartnerName.IsSet()
}

// SetPartnerName sets field value
func (o *PaymentRequestResponsePaymentRequest) SetPartnerName(v string) {
	o.PartnerName.Set(&v)
}

// GetBankName returns the BankName field value
func (o *PaymentRequestResponsePaymentRequest) GetBankName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.BankName
}

// GetBankNameOk returns a tuple with the BankName field value
// and a boolean to check if the value has been set.
func (o *PaymentRequestResponsePaymentRequest) GetBankNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.BankName, true
}

// SetBankName sets field value
func (o *PaymentRequestResponsePaymentRequest) SetBankName(v string) {
	o.BankName = v
}

// GetBankNameKana returns the BankNameKana field value
func (o *PaymentRequestResponsePaymentRequest) GetBankNameKana() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.BankNameKana
}

// GetBankNameKanaOk returns a tuple with the BankNameKana field value
// and a boolean to check if the value has been set.
func (o *PaymentRequestResponsePaymentRequest) GetBankNameKanaOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.BankNameKana, true
}

// SetBankNameKana sets field value
func (o *PaymentRequestResponsePaymentRequest) SetBankNameKana(v string) {
	o.BankNameKana = v
}

// GetBankCode returns the BankCode field value
func (o *PaymentRequestResponsePaymentRequest) GetBankCode() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.BankCode
}

// GetBankCodeOk returns a tuple with the BankCode field value
// and a boolean to check if the value has been set.
func (o *PaymentRequestResponsePaymentRequest) GetBankCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.BankCode, true
}

// SetBankCode sets field value
func (o *PaymentRequestResponsePaymentRequest) SetBankCode(v string) {
	o.BankCode = v
}

// GetBranchName returns the BranchName field value
func (o *PaymentRequestResponsePaymentRequest) GetBranchName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.BranchName
}

// GetBranchNameOk returns a tuple with the BranchName field value
// and a boolean to check if the value has been set.
func (o *PaymentRequestResponsePaymentRequest) GetBranchNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.BranchName, true
}

// SetBranchName sets field value
func (o *PaymentRequestResponsePaymentRequest) SetBranchName(v string) {
	o.BranchName = v
}

// GetBranchKana returns the BranchKana field value
func (o *PaymentRequestResponsePaymentRequest) GetBranchKana() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.BranchKana
}

// GetBranchKanaOk returns a tuple with the BranchKana field value
// and a boolean to check if the value has been set.
func (o *PaymentRequestResponsePaymentRequest) GetBranchKanaOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.BranchKana, true
}

// SetBranchKana sets field value
func (o *PaymentRequestResponsePaymentRequest) SetBranchKana(v string) {
	o.BranchKana = v
}

// GetBranchCode returns the BranchCode field value
func (o *PaymentRequestResponsePaymentRequest) GetBranchCode() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.BranchCode
}

// GetBranchCodeOk returns a tuple with the BranchCode field value
// and a boolean to check if the value has been set.
func (o *PaymentRequestResponsePaymentRequest) GetBranchCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.BranchCode, true
}

// SetBranchCode sets field value
func (o *PaymentRequestResponsePaymentRequest) SetBranchCode(v string) {
	o.BranchCode = v
}

// GetAccountType returns the AccountType field value
func (o *PaymentRequestResponsePaymentRequest) GetAccountType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.AccountType
}

// GetAccountTypeOk returns a tuple with the AccountType field value
// and a boolean to check if the value has been set.
func (o *PaymentRequestResponsePaymentRequest) GetAccountTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccountType, true
}

// SetAccountType sets field value
func (o *PaymentRequestResponsePaymentRequest) SetAccountType(v string) {
	o.AccountType = v
}

// GetAccountNumber returns the AccountNumber field value
func (o *PaymentRequestResponsePaymentRequest) GetAccountNumber() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.AccountNumber
}

// GetAccountNumberOk returns a tuple with the AccountNumber field value
// and a boolean to check if the value has been set.
func (o *PaymentRequestResponsePaymentRequest) GetAccountNumberOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccountNumber, true
}

// SetAccountNumber sets field value
func (o *PaymentRequestResponsePaymentRequest) SetAccountNumber(v string) {
	o.AccountNumber = v
}

// GetAccountName returns the AccountName field value
func (o *PaymentRequestResponsePaymentRequest) GetAccountName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.AccountName
}

// GetAccountNameOk returns a tuple with the AccountName field value
// and a boolean to check if the value has been set.
func (o *PaymentRequestResponsePaymentRequest) GetAccountNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccountName, true
}

// SetAccountName sets field value
func (o *PaymentRequestResponsePaymentRequest) SetAccountName(v string) {
	o.AccountName = v
}

func (o PaymentRequestResponsePaymentRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["company_id"] = o.CompanyId
	}
	if true {
		toSerialize["title"] = o.Title
	}
	if true {
		toSerialize["application_date"] = o.ApplicationDate
	}
	if true {
		toSerialize["description"] = o.Description.Get()
	}
	if true {
		toSerialize["total_amount"] = o.TotalAmount
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["payment_request_lines"] = o.PaymentRequestLines
	}
	if true {
		toSerialize["deal_id"] = o.DealId.Get()
	}
	if true {
		toSerialize["deal_status"] = o.DealStatus.Get()
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
		toSerialize["document_code"] = o.DocumentCode
	}
	if o.ReceiptId.IsSet() {
		toSerialize["receipt_id"] = o.ReceiptId.Get()
	}
	if true {
		toSerialize["issue_date"] = o.IssueDate
	}
	if true {
		toSerialize["payment_date"] = o.PaymentDate.Get()
	}
	if true {
		toSerialize["payment_method"] = o.PaymentMethod
	}
	if true {
		toSerialize["partner_id"] = o.PartnerId.Get()
	}
	if o.PartnerCode.IsSet() {
		toSerialize["partner_code"] = o.PartnerCode.Get()
	}
	if true {
		toSerialize["partner_name"] = o.PartnerName.Get()
	}
	if true {
		toSerialize["bank_name"] = o.BankName
	}
	if true {
		toSerialize["bank_name_kana"] = o.BankNameKana
	}
	if true {
		toSerialize["bank_code"] = o.BankCode
	}
	if true {
		toSerialize["branch_name"] = o.BranchName
	}
	if true {
		toSerialize["branch_kana"] = o.BranchKana
	}
	if true {
		toSerialize["branch_code"] = o.BranchCode
	}
	if true {
		toSerialize["account_type"] = o.AccountType
	}
	if true {
		toSerialize["account_number"] = o.AccountNumber
	}
	if true {
		toSerialize["account_name"] = o.AccountName
	}
	return json.Marshal(toSerialize)
}

type NullablePaymentRequestResponsePaymentRequest struct {
	value *PaymentRequestResponsePaymentRequest
	isSet bool
}

func (v NullablePaymentRequestResponsePaymentRequest) Get() *PaymentRequestResponsePaymentRequest {
	return v.value
}

func (v *NullablePaymentRequestResponsePaymentRequest) Set(val *PaymentRequestResponsePaymentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentRequestResponsePaymentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentRequestResponsePaymentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentRequestResponsePaymentRequest(val *PaymentRequestResponsePaymentRequest) *NullablePaymentRequestResponsePaymentRequest {
	return &NullablePaymentRequestResponsePaymentRequest{value: val, isSet: true}
}

func (v NullablePaymentRequestResponsePaymentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentRequestResponsePaymentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


