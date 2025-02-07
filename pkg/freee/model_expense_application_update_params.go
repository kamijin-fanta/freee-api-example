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

// ExpenseApplicationUpdateParams struct for ExpenseApplicationUpdateParams
type ExpenseApplicationUpdateParams struct {
	// 事業所ID
	CompanyId int32 `json:"company_id"`
	// 申請タイトル (250文字以内)
	Title string `json:"title"`
	// 申請日 (yyyy-mm-dd)
	IssueDate string `json:"issue_date"`
	// 備考 (10000文字以内)
	Description *string `json:"description,omitempty"`
	// 会計freeeのWeb画面で申請内容を編集可能にするかどうか<br> falseの場合は、Web画面での項目行の追加／削除・金額の編集が出来なくなりますが、APIでの編集は可能です。<br> デフォルトはfalseです。 
	EditableOnWeb *bool `json:"editable_on_web,omitempty"`
	// 部門ID
	SectionId *int32 `json:"section_id,omitempty"`
	// メモタグID
	TagIds *[]int32 `json:"tag_ids,omitempty"`
	ExpenseApplicationLines []ExpenseApplicationUpdateParamsExpenseApplicationLines `json:"expense_application_lines"`
	// 申請経路ID<br> <ul>     <li>経費申請のステータスを申請中として作成する場合は、必ず指定してください。</li>     <li>指定する申請経路IDは、申請経路APIを利用して取得してください。</li>     <li>         未指定の場合は、基本経路を設定している事業所では基本経路が、基本経路を設定していない事業所では利用可能な申請経路の中から最初の申請経路が自動的に使用されます。         <ul>            <li>意図しない申請経路を持った経費申請の作成を防ぐために、使用する申請経路IDを指定することを推奨します。</li>         </ul>     </li>     <li>         ベーシックプランの事業所では以下のデフォルトで用意された申請経路のみ指定できます         <ul>         <li>指定なし</li>         <li>承認者を指定</li>         </ul>     </li> </ul> 
	ApprovalFlowRouteId *int32 `json:"approval_flow_route_id,omitempty"`
	// 承認者のユーザーID<br> 指定する承認者のユーザーIDは、申請経路APIを利用して取得してください。 
	ApproverId *int32 `json:"approver_id,omitempty"`
	// 経費申請のステータス<br> falseを指定した時は申請中（in_progress）で経費申請を更新します。<br> trueを指定した時は下書き（draft）で経費申請を更新します。<br> 未指定の時は下書きとみなして経費申請を更新します。 
	Draft *bool `json:"draft,omitempty"`
	// セグメント１ID(法人向けプロフェッショナル, 法人向けエンタープライズプラン)<br> セグメントタグ一覧APIを利用して取得してください。<br> <a href=\"https://support.freee.co.jp/hc/ja/articles/360020679611-%E3%82%BB%E3%82%B0%E3%83%A1%E3%83%B3%E3%83%88-%E5%88%86%E6%9E%90%E7%94%A8%E3%82%BF%E3%82%B0-%E3%81%AE%E8%A8%AD%E5%AE%9A\" target=\"_blank\">セグメント（分析用タグ）の設定</a><br> 
	Segment1TagId *int32 `json:"segment_1_tag_id,omitempty"`
	// セグメント２ID(法人向け エンタープライズプラン)<br> セグメントタグ一覧APIを利用して取得してください。<br> <a href=\"https://support.freee.co.jp/hc/ja/articles/360020679611-%E3%82%BB%E3%82%B0%E3%83%A1%E3%83%B3%E3%83%88-%E5%88%86%E6%9E%90%E7%94%A8%E3%82%BF%E3%82%B0-%E3%81%AE%E8%A8%AD%E5%AE%9A\" target=\"_blank\">セグメント（分析用タグ）の設定</a><br> 
	Segment2TagId *int32 `json:"segment_2_tag_id,omitempty"`
	// セグメント３ID(法人向け エンタープライズプラン)<br> セグメントタグ一覧APIを利用して取得してください。<br> <a href=\"https://support.freee.co.jp/hc/ja/articles/360020679611-%E3%82%BB%E3%82%B0%E3%83%A1%E3%83%B3%E3%83%88-%E5%88%86%E6%9E%90%E7%94%A8%E3%82%BF%E3%82%B0-%E3%81%AE%E8%A8%AD%E5%AE%9A\" target=\"_blank\">セグメント（分析用タグ）の設定</a><br> 
	Segment3TagId *int32 `json:"segment_3_tag_id,omitempty"`
}

// NewExpenseApplicationUpdateParams instantiates a new ExpenseApplicationUpdateParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExpenseApplicationUpdateParams(companyId int32, title string, issueDate string, expenseApplicationLines []ExpenseApplicationUpdateParamsExpenseApplicationLines, ) *ExpenseApplicationUpdateParams {
	this := ExpenseApplicationUpdateParams{}
	this.CompanyId = companyId
	this.Title = title
	this.IssueDate = issueDate
	this.ExpenseApplicationLines = expenseApplicationLines
	return &this
}

// NewExpenseApplicationUpdateParamsWithDefaults instantiates a new ExpenseApplicationUpdateParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExpenseApplicationUpdateParamsWithDefaults() *ExpenseApplicationUpdateParams {
	this := ExpenseApplicationUpdateParams{}
	return &this
}

// GetCompanyId returns the CompanyId field value
func (o *ExpenseApplicationUpdateParams) GetCompanyId() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.CompanyId
}

// GetCompanyIdOk returns a tuple with the CompanyId field value
// and a boolean to check if the value has been set.
func (o *ExpenseApplicationUpdateParams) GetCompanyIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CompanyId, true
}

// SetCompanyId sets field value
func (o *ExpenseApplicationUpdateParams) SetCompanyId(v int32) {
	o.CompanyId = v
}

// GetTitle returns the Title field value
func (o *ExpenseApplicationUpdateParams) GetTitle() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *ExpenseApplicationUpdateParams) GetTitleOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *ExpenseApplicationUpdateParams) SetTitle(v string) {
	o.Title = v
}

// GetIssueDate returns the IssueDate field value
func (o *ExpenseApplicationUpdateParams) GetIssueDate() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.IssueDate
}

// GetIssueDateOk returns a tuple with the IssueDate field value
// and a boolean to check if the value has been set.
func (o *ExpenseApplicationUpdateParams) GetIssueDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IssueDate, true
}

// SetIssueDate sets field value
func (o *ExpenseApplicationUpdateParams) SetIssueDate(v string) {
	o.IssueDate = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ExpenseApplicationUpdateParams) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpenseApplicationUpdateParams) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ExpenseApplicationUpdateParams) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ExpenseApplicationUpdateParams) SetDescription(v string) {
	o.Description = &v
}

// GetEditableOnWeb returns the EditableOnWeb field value if set, zero value otherwise.
func (o *ExpenseApplicationUpdateParams) GetEditableOnWeb() bool {
	if o == nil || o.EditableOnWeb == nil {
		var ret bool
		return ret
	}
	return *o.EditableOnWeb
}

// GetEditableOnWebOk returns a tuple with the EditableOnWeb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpenseApplicationUpdateParams) GetEditableOnWebOk() (*bool, bool) {
	if o == nil || o.EditableOnWeb == nil {
		return nil, false
	}
	return o.EditableOnWeb, true
}

// HasEditableOnWeb returns a boolean if a field has been set.
func (o *ExpenseApplicationUpdateParams) HasEditableOnWeb() bool {
	if o != nil && o.EditableOnWeb != nil {
		return true
	}

	return false
}

// SetEditableOnWeb gets a reference to the given bool and assigns it to the EditableOnWeb field.
func (o *ExpenseApplicationUpdateParams) SetEditableOnWeb(v bool) {
	o.EditableOnWeb = &v
}

// GetSectionId returns the SectionId field value if set, zero value otherwise.
func (o *ExpenseApplicationUpdateParams) GetSectionId() int32 {
	if o == nil || o.SectionId == nil {
		var ret int32
		return ret
	}
	return *o.SectionId
}

// GetSectionIdOk returns a tuple with the SectionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpenseApplicationUpdateParams) GetSectionIdOk() (*int32, bool) {
	if o == nil || o.SectionId == nil {
		return nil, false
	}
	return o.SectionId, true
}

// HasSectionId returns a boolean if a field has been set.
func (o *ExpenseApplicationUpdateParams) HasSectionId() bool {
	if o != nil && o.SectionId != nil {
		return true
	}

	return false
}

// SetSectionId gets a reference to the given int32 and assigns it to the SectionId field.
func (o *ExpenseApplicationUpdateParams) SetSectionId(v int32) {
	o.SectionId = &v
}

// GetTagIds returns the TagIds field value if set, zero value otherwise.
func (o *ExpenseApplicationUpdateParams) GetTagIds() []int32 {
	if o == nil || o.TagIds == nil {
		var ret []int32
		return ret
	}
	return *o.TagIds
}

// GetTagIdsOk returns a tuple with the TagIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpenseApplicationUpdateParams) GetTagIdsOk() (*[]int32, bool) {
	if o == nil || o.TagIds == nil {
		return nil, false
	}
	return o.TagIds, true
}

// HasTagIds returns a boolean if a field has been set.
func (o *ExpenseApplicationUpdateParams) HasTagIds() bool {
	if o != nil && o.TagIds != nil {
		return true
	}

	return false
}

// SetTagIds gets a reference to the given []int32 and assigns it to the TagIds field.
func (o *ExpenseApplicationUpdateParams) SetTagIds(v []int32) {
	o.TagIds = &v
}

// GetExpenseApplicationLines returns the ExpenseApplicationLines field value
func (o *ExpenseApplicationUpdateParams) GetExpenseApplicationLines() []ExpenseApplicationUpdateParamsExpenseApplicationLines {
	if o == nil  {
		var ret []ExpenseApplicationUpdateParamsExpenseApplicationLines
		return ret
	}

	return o.ExpenseApplicationLines
}

// GetExpenseApplicationLinesOk returns a tuple with the ExpenseApplicationLines field value
// and a boolean to check if the value has been set.
func (o *ExpenseApplicationUpdateParams) GetExpenseApplicationLinesOk() (*[]ExpenseApplicationUpdateParamsExpenseApplicationLines, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ExpenseApplicationLines, true
}

// SetExpenseApplicationLines sets field value
func (o *ExpenseApplicationUpdateParams) SetExpenseApplicationLines(v []ExpenseApplicationUpdateParamsExpenseApplicationLines) {
	o.ExpenseApplicationLines = v
}

// GetApprovalFlowRouteId returns the ApprovalFlowRouteId field value if set, zero value otherwise.
func (o *ExpenseApplicationUpdateParams) GetApprovalFlowRouteId() int32 {
	if o == nil || o.ApprovalFlowRouteId == nil {
		var ret int32
		return ret
	}
	return *o.ApprovalFlowRouteId
}

// GetApprovalFlowRouteIdOk returns a tuple with the ApprovalFlowRouteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpenseApplicationUpdateParams) GetApprovalFlowRouteIdOk() (*int32, bool) {
	if o == nil || o.ApprovalFlowRouteId == nil {
		return nil, false
	}
	return o.ApprovalFlowRouteId, true
}

// HasApprovalFlowRouteId returns a boolean if a field has been set.
func (o *ExpenseApplicationUpdateParams) HasApprovalFlowRouteId() bool {
	if o != nil && o.ApprovalFlowRouteId != nil {
		return true
	}

	return false
}

// SetApprovalFlowRouteId gets a reference to the given int32 and assigns it to the ApprovalFlowRouteId field.
func (o *ExpenseApplicationUpdateParams) SetApprovalFlowRouteId(v int32) {
	o.ApprovalFlowRouteId = &v
}

// GetApproverId returns the ApproverId field value if set, zero value otherwise.
func (o *ExpenseApplicationUpdateParams) GetApproverId() int32 {
	if o == nil || o.ApproverId == nil {
		var ret int32
		return ret
	}
	return *o.ApproverId
}

// GetApproverIdOk returns a tuple with the ApproverId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpenseApplicationUpdateParams) GetApproverIdOk() (*int32, bool) {
	if o == nil || o.ApproverId == nil {
		return nil, false
	}
	return o.ApproverId, true
}

// HasApproverId returns a boolean if a field has been set.
func (o *ExpenseApplicationUpdateParams) HasApproverId() bool {
	if o != nil && o.ApproverId != nil {
		return true
	}

	return false
}

// SetApproverId gets a reference to the given int32 and assigns it to the ApproverId field.
func (o *ExpenseApplicationUpdateParams) SetApproverId(v int32) {
	o.ApproverId = &v
}

// GetDraft returns the Draft field value if set, zero value otherwise.
func (o *ExpenseApplicationUpdateParams) GetDraft() bool {
	if o == nil || o.Draft == nil {
		var ret bool
		return ret
	}
	return *o.Draft
}

// GetDraftOk returns a tuple with the Draft field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpenseApplicationUpdateParams) GetDraftOk() (*bool, bool) {
	if o == nil || o.Draft == nil {
		return nil, false
	}
	return o.Draft, true
}

// HasDraft returns a boolean if a field has been set.
func (o *ExpenseApplicationUpdateParams) HasDraft() bool {
	if o != nil && o.Draft != nil {
		return true
	}

	return false
}

// SetDraft gets a reference to the given bool and assigns it to the Draft field.
func (o *ExpenseApplicationUpdateParams) SetDraft(v bool) {
	o.Draft = &v
}

// GetSegment1TagId returns the Segment1TagId field value if set, zero value otherwise.
func (o *ExpenseApplicationUpdateParams) GetSegment1TagId() int32 {
	if o == nil || o.Segment1TagId == nil {
		var ret int32
		return ret
	}
	return *o.Segment1TagId
}

// GetSegment1TagIdOk returns a tuple with the Segment1TagId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpenseApplicationUpdateParams) GetSegment1TagIdOk() (*int32, bool) {
	if o == nil || o.Segment1TagId == nil {
		return nil, false
	}
	return o.Segment1TagId, true
}

// HasSegment1TagId returns a boolean if a field has been set.
func (o *ExpenseApplicationUpdateParams) HasSegment1TagId() bool {
	if o != nil && o.Segment1TagId != nil {
		return true
	}

	return false
}

// SetSegment1TagId gets a reference to the given int32 and assigns it to the Segment1TagId field.
func (o *ExpenseApplicationUpdateParams) SetSegment1TagId(v int32) {
	o.Segment1TagId = &v
}

// GetSegment2TagId returns the Segment2TagId field value if set, zero value otherwise.
func (o *ExpenseApplicationUpdateParams) GetSegment2TagId() int32 {
	if o == nil || o.Segment2TagId == nil {
		var ret int32
		return ret
	}
	return *o.Segment2TagId
}

// GetSegment2TagIdOk returns a tuple with the Segment2TagId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpenseApplicationUpdateParams) GetSegment2TagIdOk() (*int32, bool) {
	if o == nil || o.Segment2TagId == nil {
		return nil, false
	}
	return o.Segment2TagId, true
}

// HasSegment2TagId returns a boolean if a field has been set.
func (o *ExpenseApplicationUpdateParams) HasSegment2TagId() bool {
	if o != nil && o.Segment2TagId != nil {
		return true
	}

	return false
}

// SetSegment2TagId gets a reference to the given int32 and assigns it to the Segment2TagId field.
func (o *ExpenseApplicationUpdateParams) SetSegment2TagId(v int32) {
	o.Segment2TagId = &v
}

// GetSegment3TagId returns the Segment3TagId field value if set, zero value otherwise.
func (o *ExpenseApplicationUpdateParams) GetSegment3TagId() int32 {
	if o == nil || o.Segment3TagId == nil {
		var ret int32
		return ret
	}
	return *o.Segment3TagId
}

// GetSegment3TagIdOk returns a tuple with the Segment3TagId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpenseApplicationUpdateParams) GetSegment3TagIdOk() (*int32, bool) {
	if o == nil || o.Segment3TagId == nil {
		return nil, false
	}
	return o.Segment3TagId, true
}

// HasSegment3TagId returns a boolean if a field has been set.
func (o *ExpenseApplicationUpdateParams) HasSegment3TagId() bool {
	if o != nil && o.Segment3TagId != nil {
		return true
	}

	return false
}

// SetSegment3TagId gets a reference to the given int32 and assigns it to the Segment3TagId field.
func (o *ExpenseApplicationUpdateParams) SetSegment3TagId(v int32) {
	o.Segment3TagId = &v
}

func (o ExpenseApplicationUpdateParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["company_id"] = o.CompanyId
	}
	if true {
		toSerialize["title"] = o.Title
	}
	if true {
		toSerialize["issue_date"] = o.IssueDate
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.EditableOnWeb != nil {
		toSerialize["editable_on_web"] = o.EditableOnWeb
	}
	if o.SectionId != nil {
		toSerialize["section_id"] = o.SectionId
	}
	if o.TagIds != nil {
		toSerialize["tag_ids"] = o.TagIds
	}
	if true {
		toSerialize["expense_application_lines"] = o.ExpenseApplicationLines
	}
	if o.ApprovalFlowRouteId != nil {
		toSerialize["approval_flow_route_id"] = o.ApprovalFlowRouteId
	}
	if o.ApproverId != nil {
		toSerialize["approver_id"] = o.ApproverId
	}
	if o.Draft != nil {
		toSerialize["draft"] = o.Draft
	}
	if o.Segment1TagId != nil {
		toSerialize["segment_1_tag_id"] = o.Segment1TagId
	}
	if o.Segment2TagId != nil {
		toSerialize["segment_2_tag_id"] = o.Segment2TagId
	}
	if o.Segment3TagId != nil {
		toSerialize["segment_3_tag_id"] = o.Segment3TagId
	}
	return json.Marshal(toSerialize)
}

type NullableExpenseApplicationUpdateParams struct {
	value *ExpenseApplicationUpdateParams
	isSet bool
}

func (v NullableExpenseApplicationUpdateParams) Get() *ExpenseApplicationUpdateParams {
	return v.value
}

func (v *NullableExpenseApplicationUpdateParams) Set(val *ExpenseApplicationUpdateParams) {
	v.value = val
	v.isSet = true
}

func (v NullableExpenseApplicationUpdateParams) IsSet() bool {
	return v.isSet
}

func (v *NullableExpenseApplicationUpdateParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExpenseApplicationUpdateParams(val *ExpenseApplicationUpdateParams) *NullableExpenseApplicationUpdateParams {
	return &NullableExpenseApplicationUpdateParams{value: val, isSet: true}
}

func (v NullableExpenseApplicationUpdateParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExpenseApplicationUpdateParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


