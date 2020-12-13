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

// TrialBsTwoYearsResponseTrialBsTwoYears struct for TrialBsTwoYearsResponseTrialBsTwoYears
type TrialBsTwoYearsResponseTrialBsTwoYears struct {
	// 事業所ID
	CompanyId int32 `json:"company_id"`
	// 会計年度(条件に指定した時、または条件に月、日条件がない時のみ含まれる）
	FiscalYear *int32 `json:"fiscal_year,omitempty"`
	// 発生月で絞込：開始会計月(1-12)(条件に指定した時のみ含まれる）
	StartMonth *int32 `json:"start_month,omitempty"`
	// 発生月で絞込：終了会計月(1-12)(条件に指定した時のみ含まれる）
	EndMonth *int32 `json:"end_month,omitempty"`
	// 発生日で絞込：開始日(yyyy-mm-dd)(条件に指定した時のみ含まれる）
	StartDate *string `json:"start_date,omitempty"`
	// 発生日で絞込：終了日(yyyy-mm-dd)(条件に指定した時のみ含まれる）
	EndDate *string `json:"end_date,omitempty"`
	// 勘定科目の表示（勘定科目: account_item, 決算書表示:group）(条件に指定した時のみ含まれる）
	AccountItemDisplayType *string `json:"account_item_display_type,omitempty"`
	// 内訳の表示（取引先: partner, 品目: item, 勘定科目: account_item）(条件に指定した時のみ含まれる）
	BreakdownDisplayType *string `json:"breakdown_display_type,omitempty"`
	// 取引先ID(条件に指定した時のみ含まれる）
	PartnerId *int32 `json:"partner_id,omitempty"`
	// 取引先コード(条件に指定した時のみ含まれる）
	PartnerCode *string `json:"partner_code,omitempty"`
	// 品目ID(条件に指定した時のみ含まれる）
	ItemId *int32 `json:"item_id,omitempty"`
	// 決算整理仕訳のみ: only, 決算整理仕訳以外: without(条件に指定した時のみ含まれる）
	Adjustment *string `json:"adjustment,omitempty"`
	// 作成日時
	CreatedAt *string `json:"created_at,omitempty"`
	Balances []TrialBsTwoYearsResponseTrialBsTwoYearsBalances `json:"balances"`
}

// NewTrialBsTwoYearsResponseTrialBsTwoYears instantiates a new TrialBsTwoYearsResponseTrialBsTwoYears object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTrialBsTwoYearsResponseTrialBsTwoYears(companyId int32, balances []TrialBsTwoYearsResponseTrialBsTwoYearsBalances, ) *TrialBsTwoYearsResponseTrialBsTwoYears {
	this := TrialBsTwoYearsResponseTrialBsTwoYears{}
	this.CompanyId = companyId
	this.Balances = balances
	return &this
}

// NewTrialBsTwoYearsResponseTrialBsTwoYearsWithDefaults instantiates a new TrialBsTwoYearsResponseTrialBsTwoYears object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTrialBsTwoYearsResponseTrialBsTwoYearsWithDefaults() *TrialBsTwoYearsResponseTrialBsTwoYears {
	this := TrialBsTwoYearsResponseTrialBsTwoYears{}
	return &this
}

// GetCompanyId returns the CompanyId field value
func (o *TrialBsTwoYearsResponseTrialBsTwoYears) GetCompanyId() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.CompanyId
}

// GetCompanyIdOk returns a tuple with the CompanyId field value
// and a boolean to check if the value has been set.
func (o *TrialBsTwoYearsResponseTrialBsTwoYears) GetCompanyIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CompanyId, true
}

// SetCompanyId sets field value
func (o *TrialBsTwoYearsResponseTrialBsTwoYears) SetCompanyId(v int32) {
	o.CompanyId = v
}

// GetFiscalYear returns the FiscalYear field value if set, zero value otherwise.
func (o *TrialBsTwoYearsResponseTrialBsTwoYears) GetFiscalYear() int32 {
	if o == nil || o.FiscalYear == nil {
		var ret int32
		return ret
	}
	return *o.FiscalYear
}

// GetFiscalYearOk returns a tuple with the FiscalYear field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrialBsTwoYearsResponseTrialBsTwoYears) GetFiscalYearOk() (*int32, bool) {
	if o == nil || o.FiscalYear == nil {
		return nil, false
	}
	return o.FiscalYear, true
}

// HasFiscalYear returns a boolean if a field has been set.
func (o *TrialBsTwoYearsResponseTrialBsTwoYears) HasFiscalYear() bool {
	if o != nil && o.FiscalYear != nil {
		return true
	}

	return false
}

// SetFiscalYear gets a reference to the given int32 and assigns it to the FiscalYear field.
func (o *TrialBsTwoYearsResponseTrialBsTwoYears) SetFiscalYear(v int32) {
	o.FiscalYear = &v
}

// GetStartMonth returns the StartMonth field value if set, zero value otherwise.
func (o *TrialBsTwoYearsResponseTrialBsTwoYears) GetStartMonth() int32 {
	if o == nil || o.StartMonth == nil {
		var ret int32
		return ret
	}
	return *o.StartMonth
}

// GetStartMonthOk returns a tuple with the StartMonth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrialBsTwoYearsResponseTrialBsTwoYears) GetStartMonthOk() (*int32, bool) {
	if o == nil || o.StartMonth == nil {
		return nil, false
	}
	return o.StartMonth, true
}

// HasStartMonth returns a boolean if a field has been set.
func (o *TrialBsTwoYearsResponseTrialBsTwoYears) HasStartMonth() bool {
	if o != nil && o.StartMonth != nil {
		return true
	}

	return false
}

// SetStartMonth gets a reference to the given int32 and assigns it to the StartMonth field.
func (o *TrialBsTwoYearsResponseTrialBsTwoYears) SetStartMonth(v int32) {
	o.StartMonth = &v
}

// GetEndMonth returns the EndMonth field value if set, zero value otherwise.
func (o *TrialBsTwoYearsResponseTrialBsTwoYears) GetEndMonth() int32 {
	if o == nil || o.EndMonth == nil {
		var ret int32
		return ret
	}
	return *o.EndMonth
}

// GetEndMonthOk returns a tuple with the EndMonth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrialBsTwoYearsResponseTrialBsTwoYears) GetEndMonthOk() (*int32, bool) {
	if o == nil || o.EndMonth == nil {
		return nil, false
	}
	return o.EndMonth, true
}

// HasEndMonth returns a boolean if a field has been set.
func (o *TrialBsTwoYearsResponseTrialBsTwoYears) HasEndMonth() bool {
	if o != nil && o.EndMonth != nil {
		return true
	}

	return false
}

// SetEndMonth gets a reference to the given int32 and assigns it to the EndMonth field.
func (o *TrialBsTwoYearsResponseTrialBsTwoYears) SetEndMonth(v int32) {
	o.EndMonth = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *TrialBsTwoYearsResponseTrialBsTwoYears) GetStartDate() string {
	if o == nil || o.StartDate == nil {
		var ret string
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrialBsTwoYearsResponseTrialBsTwoYears) GetStartDateOk() (*string, bool) {
	if o == nil || o.StartDate == nil {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *TrialBsTwoYearsResponseTrialBsTwoYears) HasStartDate() bool {
	if o != nil && o.StartDate != nil {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given string and assigns it to the StartDate field.
func (o *TrialBsTwoYearsResponseTrialBsTwoYears) SetStartDate(v string) {
	o.StartDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *TrialBsTwoYearsResponseTrialBsTwoYears) GetEndDate() string {
	if o == nil || o.EndDate == nil {
		var ret string
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrialBsTwoYearsResponseTrialBsTwoYears) GetEndDateOk() (*string, bool) {
	if o == nil || o.EndDate == nil {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *TrialBsTwoYearsResponseTrialBsTwoYears) HasEndDate() bool {
	if o != nil && o.EndDate != nil {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given string and assigns it to the EndDate field.
func (o *TrialBsTwoYearsResponseTrialBsTwoYears) SetEndDate(v string) {
	o.EndDate = &v
}

// GetAccountItemDisplayType returns the AccountItemDisplayType field value if set, zero value otherwise.
func (o *TrialBsTwoYearsResponseTrialBsTwoYears) GetAccountItemDisplayType() string {
	if o == nil || o.AccountItemDisplayType == nil {
		var ret string
		return ret
	}
	return *o.AccountItemDisplayType
}

// GetAccountItemDisplayTypeOk returns a tuple with the AccountItemDisplayType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrialBsTwoYearsResponseTrialBsTwoYears) GetAccountItemDisplayTypeOk() (*string, bool) {
	if o == nil || o.AccountItemDisplayType == nil {
		return nil, false
	}
	return o.AccountItemDisplayType, true
}

// HasAccountItemDisplayType returns a boolean if a field has been set.
func (o *TrialBsTwoYearsResponseTrialBsTwoYears) HasAccountItemDisplayType() bool {
	if o != nil && o.AccountItemDisplayType != nil {
		return true
	}

	return false
}

// SetAccountItemDisplayType gets a reference to the given string and assigns it to the AccountItemDisplayType field.
func (o *TrialBsTwoYearsResponseTrialBsTwoYears) SetAccountItemDisplayType(v string) {
	o.AccountItemDisplayType = &v
}

// GetBreakdownDisplayType returns the BreakdownDisplayType field value if set, zero value otherwise.
func (o *TrialBsTwoYearsResponseTrialBsTwoYears) GetBreakdownDisplayType() string {
	if o == nil || o.BreakdownDisplayType == nil {
		var ret string
		return ret
	}
	return *o.BreakdownDisplayType
}

// GetBreakdownDisplayTypeOk returns a tuple with the BreakdownDisplayType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrialBsTwoYearsResponseTrialBsTwoYears) GetBreakdownDisplayTypeOk() (*string, bool) {
	if o == nil || o.BreakdownDisplayType == nil {
		return nil, false
	}
	return o.BreakdownDisplayType, true
}

// HasBreakdownDisplayType returns a boolean if a field has been set.
func (o *TrialBsTwoYearsResponseTrialBsTwoYears) HasBreakdownDisplayType() bool {
	if o != nil && o.BreakdownDisplayType != nil {
		return true
	}

	return false
}

// SetBreakdownDisplayType gets a reference to the given string and assigns it to the BreakdownDisplayType field.
func (o *TrialBsTwoYearsResponseTrialBsTwoYears) SetBreakdownDisplayType(v string) {
	o.BreakdownDisplayType = &v
}

// GetPartnerId returns the PartnerId field value if set, zero value otherwise.
func (o *TrialBsTwoYearsResponseTrialBsTwoYears) GetPartnerId() int32 {
	if o == nil || o.PartnerId == nil {
		var ret int32
		return ret
	}
	return *o.PartnerId
}

// GetPartnerIdOk returns a tuple with the PartnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrialBsTwoYearsResponseTrialBsTwoYears) GetPartnerIdOk() (*int32, bool) {
	if o == nil || o.PartnerId == nil {
		return nil, false
	}
	return o.PartnerId, true
}

// HasPartnerId returns a boolean if a field has been set.
func (o *TrialBsTwoYearsResponseTrialBsTwoYears) HasPartnerId() bool {
	if o != nil && o.PartnerId != nil {
		return true
	}

	return false
}

// SetPartnerId gets a reference to the given int32 and assigns it to the PartnerId field.
func (o *TrialBsTwoYearsResponseTrialBsTwoYears) SetPartnerId(v int32) {
	o.PartnerId = &v
}

// GetPartnerCode returns the PartnerCode field value if set, zero value otherwise.
func (o *TrialBsTwoYearsResponseTrialBsTwoYears) GetPartnerCode() string {
	if o == nil || o.PartnerCode == nil {
		var ret string
		return ret
	}
	return *o.PartnerCode
}

// GetPartnerCodeOk returns a tuple with the PartnerCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrialBsTwoYearsResponseTrialBsTwoYears) GetPartnerCodeOk() (*string, bool) {
	if o == nil || o.PartnerCode == nil {
		return nil, false
	}
	return o.PartnerCode, true
}

// HasPartnerCode returns a boolean if a field has been set.
func (o *TrialBsTwoYearsResponseTrialBsTwoYears) HasPartnerCode() bool {
	if o != nil && o.PartnerCode != nil {
		return true
	}

	return false
}

// SetPartnerCode gets a reference to the given string and assigns it to the PartnerCode field.
func (o *TrialBsTwoYearsResponseTrialBsTwoYears) SetPartnerCode(v string) {
	o.PartnerCode = &v
}

// GetItemId returns the ItemId field value if set, zero value otherwise.
func (o *TrialBsTwoYearsResponseTrialBsTwoYears) GetItemId() int32 {
	if o == nil || o.ItemId == nil {
		var ret int32
		return ret
	}
	return *o.ItemId
}

// GetItemIdOk returns a tuple with the ItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrialBsTwoYearsResponseTrialBsTwoYears) GetItemIdOk() (*int32, bool) {
	if o == nil || o.ItemId == nil {
		return nil, false
	}
	return o.ItemId, true
}

// HasItemId returns a boolean if a field has been set.
func (o *TrialBsTwoYearsResponseTrialBsTwoYears) HasItemId() bool {
	if o != nil && o.ItemId != nil {
		return true
	}

	return false
}

// SetItemId gets a reference to the given int32 and assigns it to the ItemId field.
func (o *TrialBsTwoYearsResponseTrialBsTwoYears) SetItemId(v int32) {
	o.ItemId = &v
}

// GetAdjustment returns the Adjustment field value if set, zero value otherwise.
func (o *TrialBsTwoYearsResponseTrialBsTwoYears) GetAdjustment() string {
	if o == nil || o.Adjustment == nil {
		var ret string
		return ret
	}
	return *o.Adjustment
}

// GetAdjustmentOk returns a tuple with the Adjustment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrialBsTwoYearsResponseTrialBsTwoYears) GetAdjustmentOk() (*string, bool) {
	if o == nil || o.Adjustment == nil {
		return nil, false
	}
	return o.Adjustment, true
}

// HasAdjustment returns a boolean if a field has been set.
func (o *TrialBsTwoYearsResponseTrialBsTwoYears) HasAdjustment() bool {
	if o != nil && o.Adjustment != nil {
		return true
	}

	return false
}

// SetAdjustment gets a reference to the given string and assigns it to the Adjustment field.
func (o *TrialBsTwoYearsResponseTrialBsTwoYears) SetAdjustment(v string) {
	o.Adjustment = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *TrialBsTwoYearsResponseTrialBsTwoYears) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrialBsTwoYearsResponseTrialBsTwoYears) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *TrialBsTwoYearsResponseTrialBsTwoYears) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *TrialBsTwoYearsResponseTrialBsTwoYears) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetBalances returns the Balances field value
func (o *TrialBsTwoYearsResponseTrialBsTwoYears) GetBalances() []TrialBsTwoYearsResponseTrialBsTwoYearsBalances {
	if o == nil  {
		var ret []TrialBsTwoYearsResponseTrialBsTwoYearsBalances
		return ret
	}

	return o.Balances
}

// GetBalancesOk returns a tuple with the Balances field value
// and a boolean to check if the value has been set.
func (o *TrialBsTwoYearsResponseTrialBsTwoYears) GetBalancesOk() (*[]TrialBsTwoYearsResponseTrialBsTwoYearsBalances, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Balances, true
}

// SetBalances sets field value
func (o *TrialBsTwoYearsResponseTrialBsTwoYears) SetBalances(v []TrialBsTwoYearsResponseTrialBsTwoYearsBalances) {
	o.Balances = v
}

func (o TrialBsTwoYearsResponseTrialBsTwoYears) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["company_id"] = o.CompanyId
	}
	if o.FiscalYear != nil {
		toSerialize["fiscal_year"] = o.FiscalYear
	}
	if o.StartMonth != nil {
		toSerialize["start_month"] = o.StartMonth
	}
	if o.EndMonth != nil {
		toSerialize["end_month"] = o.EndMonth
	}
	if o.StartDate != nil {
		toSerialize["start_date"] = o.StartDate
	}
	if o.EndDate != nil {
		toSerialize["end_date"] = o.EndDate
	}
	if o.AccountItemDisplayType != nil {
		toSerialize["account_item_display_type"] = o.AccountItemDisplayType
	}
	if o.BreakdownDisplayType != nil {
		toSerialize["breakdown_display_type"] = o.BreakdownDisplayType
	}
	if o.PartnerId != nil {
		toSerialize["partner_id"] = o.PartnerId
	}
	if o.PartnerCode != nil {
		toSerialize["partner_code"] = o.PartnerCode
	}
	if o.ItemId != nil {
		toSerialize["item_id"] = o.ItemId
	}
	if o.Adjustment != nil {
		toSerialize["adjustment"] = o.Adjustment
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if true {
		toSerialize["balances"] = o.Balances
	}
	return json.Marshal(toSerialize)
}

type NullableTrialBsTwoYearsResponseTrialBsTwoYears struct {
	value *TrialBsTwoYearsResponseTrialBsTwoYears
	isSet bool
}

func (v NullableTrialBsTwoYearsResponseTrialBsTwoYears) Get() *TrialBsTwoYearsResponseTrialBsTwoYears {
	return v.value
}

func (v *NullableTrialBsTwoYearsResponseTrialBsTwoYears) Set(val *TrialBsTwoYearsResponseTrialBsTwoYears) {
	v.value = val
	v.isSet = true
}

func (v NullableTrialBsTwoYearsResponseTrialBsTwoYears) IsSet() bool {
	return v.isSet
}

func (v *NullableTrialBsTwoYearsResponseTrialBsTwoYears) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrialBsTwoYearsResponseTrialBsTwoYears(val *TrialBsTwoYearsResponseTrialBsTwoYears) *NullableTrialBsTwoYearsResponseTrialBsTwoYears {
	return &NullableTrialBsTwoYearsResponseTrialBsTwoYears{value: val, isSet: true}
}

func (v NullableTrialBsTwoYearsResponseTrialBsTwoYears) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrialBsTwoYearsResponseTrialBsTwoYears) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

