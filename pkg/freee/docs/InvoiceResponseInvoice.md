# InvoiceResponseInvoice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | 請求書ID | 
**CompanyId** | **int32** | 事業所ID | 
**IssueDate** | **string** | 請求日 (yyyy-mm-dd) | 
**PartnerId** | **NullableInt32** | 取引先ID | 
**PartnerCode** | Pointer to **NullableString** | 取引先コード | [optional] 
**InvoiceNumber** | **string** | 請求書番号 | 
**Title** | Pointer to **NullableString** | タイトル | [optional] 
**DueDate** | Pointer to **NullableString** | 期日 (yyyy-mm-dd) | [optional] 
**TotalAmount** | **int32** | 合計金額 | 
**TotalVat** | Pointer to **int32** | 合計金額 | [optional] 
**SubTotal** | Pointer to **int32** | 小計 | [optional] 
**BookingDate** | Pointer to **NullableString** | 売上計上日 | [optional] 
**Description** | Pointer to **NullableString** | 概要 | [optional] 
**InvoiceStatus** | **string** | 請求書ステータス  (draft: 下書き, applying: 申請中, remanded: 差し戻し, rejected: 却下, approved: 承認済み, submitted: 送付済み, unsubmitted: 請求書の承認フローが無効の場合のみ、unsubmitted（送付待ち）の値をとります) | 
**PaymentStatus** | Pointer to **string** | 入金ステータス  (unsettled: 入金待ち, settled: 入金済み) | [optional] 
**PaymentDate** | Pointer to **NullableString** | 入金日 | [optional] 
**WebPublishedAt** | Pointer to **NullableString** | Web共有日時(最新) | [optional] 
**WebDownloadedAt** | Pointer to **NullableString** | Web共有ダウンロード日時(最新) | [optional] 
**WebConfirmedAt** | Pointer to **NullableString** | Web共有取引先確認日時(最新) | [optional] 
**MailSentAt** | Pointer to **NullableString** | メール送信日時(最新) | [optional] 
**PostingStatus** | **string** | 郵送ステータス(unrequested: リクエスト前, preview_registered: プレビュー登録, preview_failed: プレビュー登録失敗, ordered: 注文中, order_failed: 注文失敗, printing: 印刷中, canceled: キャンセル, posted: 投函済み) | 
**PartnerName** | Pointer to **NullableString** | 取引先名 | [optional] 
**PartnerDisplayName** | Pointer to **NullableString** | 請求書に表示する取引先名 | [optional] 
**PartnerTitle** | Pointer to **NullableString** | 敬称（御中、様、(空白)の3つから選択） | [optional] 
**PartnerZipcode** | Pointer to **NullableString** | 郵便番号 | [optional] 
**PartnerPrefectureCode** | Pointer to **NullableInt32** | 都道府県コード（0:北海道、1:青森、2:岩手、3:宮城、4:秋田、5:山形、6:福島、7:茨城、8:栃木、9:群馬、10:埼玉、11:千葉、12:東京、13:神奈川、14:新潟、15:富山、16:石川、17:福井、18:山梨、19:長野、20:岐阜、21:静岡、22:愛知、23:三重、24:滋賀、25:京都、26:大阪、27:兵庫、28:奈良、29:和歌山、30:鳥取、31:島根、32:岡山、33:広島、34:山口、35:徳島、36:香川、37:愛媛、38:高知、39:福岡、40:佐賀、41:長崎、42:熊本、43:大分、44:宮崎、45:鹿児島、46:沖縄 | [optional] 
**PartnerPrefectureName** | Pointer to **NullableString** | 都道府県 | [optional] 
**PartnerAddress1** | Pointer to **NullableString** | 市区町村・番地 | [optional] 
**PartnerAddress2** | Pointer to **NullableString** | 建物名・部屋番号など | [optional] 
**PartnerContactInfo** | Pointer to **NullableString** | 取引先担当者名 | [optional] 
**CompanyName** | **string** | 事業所名 | 
**CompanyZipcode** | Pointer to **NullableString** | 郵便番号 | [optional] 
**CompanyPrefectureCode** | Pointer to **NullableInt32** | 都道府県コード（0:北海道、1:青森、2:岩手、3:宮城、4:秋田、5:山形、6:福島、7:茨城、8:栃木、9:群馬、10:埼玉、11:千葉、12:東京、13:神奈川、14:新潟、15:富山、16:石川、17:福井、18:山梨、19:長野、20:岐阜、21:静岡、22:愛知、23:三重、24:滋賀、25:京都、26:大阪、27:兵庫、28:奈良、29:和歌山、30:鳥取、31:島根、32:岡山、33:広島、34:山口、35:徳島、36:香川、37:愛媛、38:高知、39:福岡、40:佐賀、41:長崎、42:熊本、43:大分、44:宮崎、45:鹿児島、46:沖縄 | [optional] 
**CompanyPrefectureName** | Pointer to **NullableString** | 都道府県 | [optional] 
**CompanyAddress1** | Pointer to **NullableString** | 市区町村・番地 | [optional] 
**CompanyAddress2** | Pointer to **NullableString** | 建物名・部屋番号など | [optional] 
**CompanyContactInfo** | Pointer to **NullableString** | 事業所担当者名 | [optional] 
**PaymentType** | **string** | 支払方法 (振込: transfer, 引き落とし: direct_debit) | 
**PaymentBankInfo** | Pointer to **NullableString** | 支払口座 | [optional] 
**Message** | Pointer to **NullableString** | メッセージ | [optional] 
**Notes** | Pointer to **NullableString** | 備考 | [optional] 
**InvoiceLayout** | **string** | 請求書レイアウト * &#x60;default_classic&#x60; - レイアウト１/クラシック (デフォルト)  * &#x60;standard_classic&#x60; - レイアウト２/クラシック  * &#x60;envelope_classic&#x60; - 封筒１/クラシック  * &#x60;carried_forward_standard_classic&#x60; - レイアウト３（繰越金額欄あり）/クラシック  * &#x60;carried_forward_envelope_classic&#x60; - 封筒２（繰越金額欄あり）/クラシック  * &#x60;default_modern&#x60; - レイアウト１/モダン  * &#x60;standard_modern&#x60; - レイアウト２/モダン  * &#x60;envelope_modern&#x60; - 封筒/モダン | 
**TaxEntryMethod** | **string** | 請求書の消費税計算方法(inclusive: 内税, exclusive: 外税) | 
**DealId** | Pointer to **NullableInt32** | 取引ID (invoice_statusがsubmitted, unsubmittedの時IDが表示されます) | [optional] 
**InvoiceContents** | Pointer to [**[]InvoiceIndexResponseInvoiceContents**](InvoiceIndexResponseInvoiceContents.md) | 請求内容 | [optional] 
**TotalAmountPerVatRate** | [**InvoiceIndexResponseTotalAmountPerVatRate**](invoiceIndexResponse_total_amount_per_vat_rate.md) |  | 
**RelatedQuotationIds** | Pointer to **[]int32** | 関連する見積書ID(配列)&lt;br&gt; 下記で作成したものが該当します。  &lt;a href&#x3D;\&quot;https://support.freee.co.jp/hc/ja/articles/203318410-請求書-納品書-見積書を作成する%231-2#1-2\&quot; target&#x3D;\&quot;_blank\&quot;&gt;見積書・納品書を納品書・請求書に変換する&lt;/a&gt;&lt;br&gt; &lt;a href&#x3D;\&quot;https://support.freee.co.jp/hc/ja/articles/209076226\&quot; target&#x3D;\&quot;_blank\&quot;&gt;複数の見積書・納品書から合算請求書を作成する&lt;/a&gt;&lt;br&gt;  | [optional] 

## Methods

### NewInvoiceResponseInvoice

`func NewInvoiceResponseInvoice(id int32, companyId int32, issueDate string, partnerId NullableInt32, invoiceNumber string, totalAmount int32, invoiceStatus string, postingStatus string, companyName string, paymentType string, invoiceLayout string, taxEntryMethod string, totalAmountPerVatRate InvoiceIndexResponseTotalAmountPerVatRate, ) *InvoiceResponseInvoice`

NewInvoiceResponseInvoice instantiates a new InvoiceResponseInvoice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceResponseInvoiceWithDefaults

`func NewInvoiceResponseInvoiceWithDefaults() *InvoiceResponseInvoice`

NewInvoiceResponseInvoiceWithDefaults instantiates a new InvoiceResponseInvoice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InvoiceResponseInvoice) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InvoiceResponseInvoice) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InvoiceResponseInvoice) SetId(v int32)`

SetId sets Id field to given value.


### GetCompanyId

`func (o *InvoiceResponseInvoice) GetCompanyId() int32`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *InvoiceResponseInvoice) GetCompanyIdOk() (*int32, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *InvoiceResponseInvoice) SetCompanyId(v int32)`

SetCompanyId sets CompanyId field to given value.


### GetIssueDate

`func (o *InvoiceResponseInvoice) GetIssueDate() string`

GetIssueDate returns the IssueDate field if non-nil, zero value otherwise.

### GetIssueDateOk

`func (o *InvoiceResponseInvoice) GetIssueDateOk() (*string, bool)`

GetIssueDateOk returns a tuple with the IssueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueDate

`func (o *InvoiceResponseInvoice) SetIssueDate(v string)`

SetIssueDate sets IssueDate field to given value.


### GetPartnerId

`func (o *InvoiceResponseInvoice) GetPartnerId() int32`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *InvoiceResponseInvoice) GetPartnerIdOk() (*int32, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *InvoiceResponseInvoice) SetPartnerId(v int32)`

SetPartnerId sets PartnerId field to given value.


### SetPartnerIdNil

`func (o *InvoiceResponseInvoice) SetPartnerIdNil(b bool)`

 SetPartnerIdNil sets the value for PartnerId to be an explicit nil

### UnsetPartnerId
`func (o *InvoiceResponseInvoice) UnsetPartnerId()`

UnsetPartnerId ensures that no value is present for PartnerId, not even an explicit nil
### GetPartnerCode

`func (o *InvoiceResponseInvoice) GetPartnerCode() string`

GetPartnerCode returns the PartnerCode field if non-nil, zero value otherwise.

### GetPartnerCodeOk

`func (o *InvoiceResponseInvoice) GetPartnerCodeOk() (*string, bool)`

GetPartnerCodeOk returns a tuple with the PartnerCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerCode

`func (o *InvoiceResponseInvoice) SetPartnerCode(v string)`

SetPartnerCode sets PartnerCode field to given value.

### HasPartnerCode

`func (o *InvoiceResponseInvoice) HasPartnerCode() bool`

HasPartnerCode returns a boolean if a field has been set.

### SetPartnerCodeNil

`func (o *InvoiceResponseInvoice) SetPartnerCodeNil(b bool)`

 SetPartnerCodeNil sets the value for PartnerCode to be an explicit nil

### UnsetPartnerCode
`func (o *InvoiceResponseInvoice) UnsetPartnerCode()`

UnsetPartnerCode ensures that no value is present for PartnerCode, not even an explicit nil
### GetInvoiceNumber

`func (o *InvoiceResponseInvoice) GetInvoiceNumber() string`

GetInvoiceNumber returns the InvoiceNumber field if non-nil, zero value otherwise.

### GetInvoiceNumberOk

`func (o *InvoiceResponseInvoice) GetInvoiceNumberOk() (*string, bool)`

GetInvoiceNumberOk returns a tuple with the InvoiceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceNumber

`func (o *InvoiceResponseInvoice) SetInvoiceNumber(v string)`

SetInvoiceNumber sets InvoiceNumber field to given value.


### GetTitle

`func (o *InvoiceResponseInvoice) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *InvoiceResponseInvoice) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *InvoiceResponseInvoice) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *InvoiceResponseInvoice) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *InvoiceResponseInvoice) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *InvoiceResponseInvoice) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDueDate

`func (o *InvoiceResponseInvoice) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *InvoiceResponseInvoice) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *InvoiceResponseInvoice) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *InvoiceResponseInvoice) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### SetDueDateNil

`func (o *InvoiceResponseInvoice) SetDueDateNil(b bool)`

 SetDueDateNil sets the value for DueDate to be an explicit nil

### UnsetDueDate
`func (o *InvoiceResponseInvoice) UnsetDueDate()`

UnsetDueDate ensures that no value is present for DueDate, not even an explicit nil
### GetTotalAmount

`func (o *InvoiceResponseInvoice) GetTotalAmount() int32`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *InvoiceResponseInvoice) GetTotalAmountOk() (*int32, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *InvoiceResponseInvoice) SetTotalAmount(v int32)`

SetTotalAmount sets TotalAmount field to given value.


### GetTotalVat

`func (o *InvoiceResponseInvoice) GetTotalVat() int32`

GetTotalVat returns the TotalVat field if non-nil, zero value otherwise.

### GetTotalVatOk

`func (o *InvoiceResponseInvoice) GetTotalVatOk() (*int32, bool)`

GetTotalVatOk returns a tuple with the TotalVat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalVat

`func (o *InvoiceResponseInvoice) SetTotalVat(v int32)`

SetTotalVat sets TotalVat field to given value.

### HasTotalVat

`func (o *InvoiceResponseInvoice) HasTotalVat() bool`

HasTotalVat returns a boolean if a field has been set.

### GetSubTotal

`func (o *InvoiceResponseInvoice) GetSubTotal() int32`

GetSubTotal returns the SubTotal field if non-nil, zero value otherwise.

### GetSubTotalOk

`func (o *InvoiceResponseInvoice) GetSubTotalOk() (*int32, bool)`

GetSubTotalOk returns a tuple with the SubTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubTotal

`func (o *InvoiceResponseInvoice) SetSubTotal(v int32)`

SetSubTotal sets SubTotal field to given value.

### HasSubTotal

`func (o *InvoiceResponseInvoice) HasSubTotal() bool`

HasSubTotal returns a boolean if a field has been set.

### GetBookingDate

`func (o *InvoiceResponseInvoice) GetBookingDate() string`

GetBookingDate returns the BookingDate field if non-nil, zero value otherwise.

### GetBookingDateOk

`func (o *InvoiceResponseInvoice) GetBookingDateOk() (*string, bool)`

GetBookingDateOk returns a tuple with the BookingDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBookingDate

`func (o *InvoiceResponseInvoice) SetBookingDate(v string)`

SetBookingDate sets BookingDate field to given value.

### HasBookingDate

`func (o *InvoiceResponseInvoice) HasBookingDate() bool`

HasBookingDate returns a boolean if a field has been set.

### SetBookingDateNil

`func (o *InvoiceResponseInvoice) SetBookingDateNil(b bool)`

 SetBookingDateNil sets the value for BookingDate to be an explicit nil

### UnsetBookingDate
`func (o *InvoiceResponseInvoice) UnsetBookingDate()`

UnsetBookingDate ensures that no value is present for BookingDate, not even an explicit nil
### GetDescription

`func (o *InvoiceResponseInvoice) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InvoiceResponseInvoice) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InvoiceResponseInvoice) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InvoiceResponseInvoice) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *InvoiceResponseInvoice) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *InvoiceResponseInvoice) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetInvoiceStatus

`func (o *InvoiceResponseInvoice) GetInvoiceStatus() string`

GetInvoiceStatus returns the InvoiceStatus field if non-nil, zero value otherwise.

### GetInvoiceStatusOk

`func (o *InvoiceResponseInvoice) GetInvoiceStatusOk() (*string, bool)`

GetInvoiceStatusOk returns a tuple with the InvoiceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceStatus

`func (o *InvoiceResponseInvoice) SetInvoiceStatus(v string)`

SetInvoiceStatus sets InvoiceStatus field to given value.


### GetPaymentStatus

`func (o *InvoiceResponseInvoice) GetPaymentStatus() string`

GetPaymentStatus returns the PaymentStatus field if non-nil, zero value otherwise.

### GetPaymentStatusOk

`func (o *InvoiceResponseInvoice) GetPaymentStatusOk() (*string, bool)`

GetPaymentStatusOk returns a tuple with the PaymentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentStatus

`func (o *InvoiceResponseInvoice) SetPaymentStatus(v string)`

SetPaymentStatus sets PaymentStatus field to given value.

### HasPaymentStatus

`func (o *InvoiceResponseInvoice) HasPaymentStatus() bool`

HasPaymentStatus returns a boolean if a field has been set.

### GetPaymentDate

`func (o *InvoiceResponseInvoice) GetPaymentDate() string`

GetPaymentDate returns the PaymentDate field if non-nil, zero value otherwise.

### GetPaymentDateOk

`func (o *InvoiceResponseInvoice) GetPaymentDateOk() (*string, bool)`

GetPaymentDateOk returns a tuple with the PaymentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentDate

`func (o *InvoiceResponseInvoice) SetPaymentDate(v string)`

SetPaymentDate sets PaymentDate field to given value.

### HasPaymentDate

`func (o *InvoiceResponseInvoice) HasPaymentDate() bool`

HasPaymentDate returns a boolean if a field has been set.

### SetPaymentDateNil

`func (o *InvoiceResponseInvoice) SetPaymentDateNil(b bool)`

 SetPaymentDateNil sets the value for PaymentDate to be an explicit nil

### UnsetPaymentDate
`func (o *InvoiceResponseInvoice) UnsetPaymentDate()`

UnsetPaymentDate ensures that no value is present for PaymentDate, not even an explicit nil
### GetWebPublishedAt

`func (o *InvoiceResponseInvoice) GetWebPublishedAt() string`

GetWebPublishedAt returns the WebPublishedAt field if non-nil, zero value otherwise.

### GetWebPublishedAtOk

`func (o *InvoiceResponseInvoice) GetWebPublishedAtOk() (*string, bool)`

GetWebPublishedAtOk returns a tuple with the WebPublishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebPublishedAt

`func (o *InvoiceResponseInvoice) SetWebPublishedAt(v string)`

SetWebPublishedAt sets WebPublishedAt field to given value.

### HasWebPublishedAt

`func (o *InvoiceResponseInvoice) HasWebPublishedAt() bool`

HasWebPublishedAt returns a boolean if a field has been set.

### SetWebPublishedAtNil

`func (o *InvoiceResponseInvoice) SetWebPublishedAtNil(b bool)`

 SetWebPublishedAtNil sets the value for WebPublishedAt to be an explicit nil

### UnsetWebPublishedAt
`func (o *InvoiceResponseInvoice) UnsetWebPublishedAt()`

UnsetWebPublishedAt ensures that no value is present for WebPublishedAt, not even an explicit nil
### GetWebDownloadedAt

`func (o *InvoiceResponseInvoice) GetWebDownloadedAt() string`

GetWebDownloadedAt returns the WebDownloadedAt field if non-nil, zero value otherwise.

### GetWebDownloadedAtOk

`func (o *InvoiceResponseInvoice) GetWebDownloadedAtOk() (*string, bool)`

GetWebDownloadedAtOk returns a tuple with the WebDownloadedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebDownloadedAt

`func (o *InvoiceResponseInvoice) SetWebDownloadedAt(v string)`

SetWebDownloadedAt sets WebDownloadedAt field to given value.

### HasWebDownloadedAt

`func (o *InvoiceResponseInvoice) HasWebDownloadedAt() bool`

HasWebDownloadedAt returns a boolean if a field has been set.

### SetWebDownloadedAtNil

`func (o *InvoiceResponseInvoice) SetWebDownloadedAtNil(b bool)`

 SetWebDownloadedAtNil sets the value for WebDownloadedAt to be an explicit nil

### UnsetWebDownloadedAt
`func (o *InvoiceResponseInvoice) UnsetWebDownloadedAt()`

UnsetWebDownloadedAt ensures that no value is present for WebDownloadedAt, not even an explicit nil
### GetWebConfirmedAt

`func (o *InvoiceResponseInvoice) GetWebConfirmedAt() string`

GetWebConfirmedAt returns the WebConfirmedAt field if non-nil, zero value otherwise.

### GetWebConfirmedAtOk

`func (o *InvoiceResponseInvoice) GetWebConfirmedAtOk() (*string, bool)`

GetWebConfirmedAtOk returns a tuple with the WebConfirmedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebConfirmedAt

`func (o *InvoiceResponseInvoice) SetWebConfirmedAt(v string)`

SetWebConfirmedAt sets WebConfirmedAt field to given value.

### HasWebConfirmedAt

`func (o *InvoiceResponseInvoice) HasWebConfirmedAt() bool`

HasWebConfirmedAt returns a boolean if a field has been set.

### SetWebConfirmedAtNil

`func (o *InvoiceResponseInvoice) SetWebConfirmedAtNil(b bool)`

 SetWebConfirmedAtNil sets the value for WebConfirmedAt to be an explicit nil

### UnsetWebConfirmedAt
`func (o *InvoiceResponseInvoice) UnsetWebConfirmedAt()`

UnsetWebConfirmedAt ensures that no value is present for WebConfirmedAt, not even an explicit nil
### GetMailSentAt

`func (o *InvoiceResponseInvoice) GetMailSentAt() string`

GetMailSentAt returns the MailSentAt field if non-nil, zero value otherwise.

### GetMailSentAtOk

`func (o *InvoiceResponseInvoice) GetMailSentAtOk() (*string, bool)`

GetMailSentAtOk returns a tuple with the MailSentAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailSentAt

`func (o *InvoiceResponseInvoice) SetMailSentAt(v string)`

SetMailSentAt sets MailSentAt field to given value.

### HasMailSentAt

`func (o *InvoiceResponseInvoice) HasMailSentAt() bool`

HasMailSentAt returns a boolean if a field has been set.

### SetMailSentAtNil

`func (o *InvoiceResponseInvoice) SetMailSentAtNil(b bool)`

 SetMailSentAtNil sets the value for MailSentAt to be an explicit nil

### UnsetMailSentAt
`func (o *InvoiceResponseInvoice) UnsetMailSentAt()`

UnsetMailSentAt ensures that no value is present for MailSentAt, not even an explicit nil
### GetPostingStatus

`func (o *InvoiceResponseInvoice) GetPostingStatus() string`

GetPostingStatus returns the PostingStatus field if non-nil, zero value otherwise.

### GetPostingStatusOk

`func (o *InvoiceResponseInvoice) GetPostingStatusOk() (*string, bool)`

GetPostingStatusOk returns a tuple with the PostingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostingStatus

`func (o *InvoiceResponseInvoice) SetPostingStatus(v string)`

SetPostingStatus sets PostingStatus field to given value.


### GetPartnerName

`func (o *InvoiceResponseInvoice) GetPartnerName() string`

GetPartnerName returns the PartnerName field if non-nil, zero value otherwise.

### GetPartnerNameOk

`func (o *InvoiceResponseInvoice) GetPartnerNameOk() (*string, bool)`

GetPartnerNameOk returns a tuple with the PartnerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerName

`func (o *InvoiceResponseInvoice) SetPartnerName(v string)`

SetPartnerName sets PartnerName field to given value.

### HasPartnerName

`func (o *InvoiceResponseInvoice) HasPartnerName() bool`

HasPartnerName returns a boolean if a field has been set.

### SetPartnerNameNil

`func (o *InvoiceResponseInvoice) SetPartnerNameNil(b bool)`

 SetPartnerNameNil sets the value for PartnerName to be an explicit nil

### UnsetPartnerName
`func (o *InvoiceResponseInvoice) UnsetPartnerName()`

UnsetPartnerName ensures that no value is present for PartnerName, not even an explicit nil
### GetPartnerDisplayName

`func (o *InvoiceResponseInvoice) GetPartnerDisplayName() string`

GetPartnerDisplayName returns the PartnerDisplayName field if non-nil, zero value otherwise.

### GetPartnerDisplayNameOk

`func (o *InvoiceResponseInvoice) GetPartnerDisplayNameOk() (*string, bool)`

GetPartnerDisplayNameOk returns a tuple with the PartnerDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerDisplayName

`func (o *InvoiceResponseInvoice) SetPartnerDisplayName(v string)`

SetPartnerDisplayName sets PartnerDisplayName field to given value.

### HasPartnerDisplayName

`func (o *InvoiceResponseInvoice) HasPartnerDisplayName() bool`

HasPartnerDisplayName returns a boolean if a field has been set.

### SetPartnerDisplayNameNil

`func (o *InvoiceResponseInvoice) SetPartnerDisplayNameNil(b bool)`

 SetPartnerDisplayNameNil sets the value for PartnerDisplayName to be an explicit nil

### UnsetPartnerDisplayName
`func (o *InvoiceResponseInvoice) UnsetPartnerDisplayName()`

UnsetPartnerDisplayName ensures that no value is present for PartnerDisplayName, not even an explicit nil
### GetPartnerTitle

`func (o *InvoiceResponseInvoice) GetPartnerTitle() string`

GetPartnerTitle returns the PartnerTitle field if non-nil, zero value otherwise.

### GetPartnerTitleOk

`func (o *InvoiceResponseInvoice) GetPartnerTitleOk() (*string, bool)`

GetPartnerTitleOk returns a tuple with the PartnerTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerTitle

`func (o *InvoiceResponseInvoice) SetPartnerTitle(v string)`

SetPartnerTitle sets PartnerTitle field to given value.

### HasPartnerTitle

`func (o *InvoiceResponseInvoice) HasPartnerTitle() bool`

HasPartnerTitle returns a boolean if a field has been set.

### SetPartnerTitleNil

`func (o *InvoiceResponseInvoice) SetPartnerTitleNil(b bool)`

 SetPartnerTitleNil sets the value for PartnerTitle to be an explicit nil

### UnsetPartnerTitle
`func (o *InvoiceResponseInvoice) UnsetPartnerTitle()`

UnsetPartnerTitle ensures that no value is present for PartnerTitle, not even an explicit nil
### GetPartnerZipcode

`func (o *InvoiceResponseInvoice) GetPartnerZipcode() string`

GetPartnerZipcode returns the PartnerZipcode field if non-nil, zero value otherwise.

### GetPartnerZipcodeOk

`func (o *InvoiceResponseInvoice) GetPartnerZipcodeOk() (*string, bool)`

GetPartnerZipcodeOk returns a tuple with the PartnerZipcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerZipcode

`func (o *InvoiceResponseInvoice) SetPartnerZipcode(v string)`

SetPartnerZipcode sets PartnerZipcode field to given value.

### HasPartnerZipcode

`func (o *InvoiceResponseInvoice) HasPartnerZipcode() bool`

HasPartnerZipcode returns a boolean if a field has been set.

### SetPartnerZipcodeNil

`func (o *InvoiceResponseInvoice) SetPartnerZipcodeNil(b bool)`

 SetPartnerZipcodeNil sets the value for PartnerZipcode to be an explicit nil

### UnsetPartnerZipcode
`func (o *InvoiceResponseInvoice) UnsetPartnerZipcode()`

UnsetPartnerZipcode ensures that no value is present for PartnerZipcode, not even an explicit nil
### GetPartnerPrefectureCode

`func (o *InvoiceResponseInvoice) GetPartnerPrefectureCode() int32`

GetPartnerPrefectureCode returns the PartnerPrefectureCode field if non-nil, zero value otherwise.

### GetPartnerPrefectureCodeOk

`func (o *InvoiceResponseInvoice) GetPartnerPrefectureCodeOk() (*int32, bool)`

GetPartnerPrefectureCodeOk returns a tuple with the PartnerPrefectureCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerPrefectureCode

`func (o *InvoiceResponseInvoice) SetPartnerPrefectureCode(v int32)`

SetPartnerPrefectureCode sets PartnerPrefectureCode field to given value.

### HasPartnerPrefectureCode

`func (o *InvoiceResponseInvoice) HasPartnerPrefectureCode() bool`

HasPartnerPrefectureCode returns a boolean if a field has been set.

### SetPartnerPrefectureCodeNil

`func (o *InvoiceResponseInvoice) SetPartnerPrefectureCodeNil(b bool)`

 SetPartnerPrefectureCodeNil sets the value for PartnerPrefectureCode to be an explicit nil

### UnsetPartnerPrefectureCode
`func (o *InvoiceResponseInvoice) UnsetPartnerPrefectureCode()`

UnsetPartnerPrefectureCode ensures that no value is present for PartnerPrefectureCode, not even an explicit nil
### GetPartnerPrefectureName

`func (o *InvoiceResponseInvoice) GetPartnerPrefectureName() string`

GetPartnerPrefectureName returns the PartnerPrefectureName field if non-nil, zero value otherwise.

### GetPartnerPrefectureNameOk

`func (o *InvoiceResponseInvoice) GetPartnerPrefectureNameOk() (*string, bool)`

GetPartnerPrefectureNameOk returns a tuple with the PartnerPrefectureName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerPrefectureName

`func (o *InvoiceResponseInvoice) SetPartnerPrefectureName(v string)`

SetPartnerPrefectureName sets PartnerPrefectureName field to given value.

### HasPartnerPrefectureName

`func (o *InvoiceResponseInvoice) HasPartnerPrefectureName() bool`

HasPartnerPrefectureName returns a boolean if a field has been set.

### SetPartnerPrefectureNameNil

`func (o *InvoiceResponseInvoice) SetPartnerPrefectureNameNil(b bool)`

 SetPartnerPrefectureNameNil sets the value for PartnerPrefectureName to be an explicit nil

### UnsetPartnerPrefectureName
`func (o *InvoiceResponseInvoice) UnsetPartnerPrefectureName()`

UnsetPartnerPrefectureName ensures that no value is present for PartnerPrefectureName, not even an explicit nil
### GetPartnerAddress1

`func (o *InvoiceResponseInvoice) GetPartnerAddress1() string`

GetPartnerAddress1 returns the PartnerAddress1 field if non-nil, zero value otherwise.

### GetPartnerAddress1Ok

`func (o *InvoiceResponseInvoice) GetPartnerAddress1Ok() (*string, bool)`

GetPartnerAddress1Ok returns a tuple with the PartnerAddress1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerAddress1

`func (o *InvoiceResponseInvoice) SetPartnerAddress1(v string)`

SetPartnerAddress1 sets PartnerAddress1 field to given value.

### HasPartnerAddress1

`func (o *InvoiceResponseInvoice) HasPartnerAddress1() bool`

HasPartnerAddress1 returns a boolean if a field has been set.

### SetPartnerAddress1Nil

`func (o *InvoiceResponseInvoice) SetPartnerAddress1Nil(b bool)`

 SetPartnerAddress1Nil sets the value for PartnerAddress1 to be an explicit nil

### UnsetPartnerAddress1
`func (o *InvoiceResponseInvoice) UnsetPartnerAddress1()`

UnsetPartnerAddress1 ensures that no value is present for PartnerAddress1, not even an explicit nil
### GetPartnerAddress2

`func (o *InvoiceResponseInvoice) GetPartnerAddress2() string`

GetPartnerAddress2 returns the PartnerAddress2 field if non-nil, zero value otherwise.

### GetPartnerAddress2Ok

`func (o *InvoiceResponseInvoice) GetPartnerAddress2Ok() (*string, bool)`

GetPartnerAddress2Ok returns a tuple with the PartnerAddress2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerAddress2

`func (o *InvoiceResponseInvoice) SetPartnerAddress2(v string)`

SetPartnerAddress2 sets PartnerAddress2 field to given value.

### HasPartnerAddress2

`func (o *InvoiceResponseInvoice) HasPartnerAddress2() bool`

HasPartnerAddress2 returns a boolean if a field has been set.

### SetPartnerAddress2Nil

`func (o *InvoiceResponseInvoice) SetPartnerAddress2Nil(b bool)`

 SetPartnerAddress2Nil sets the value for PartnerAddress2 to be an explicit nil

### UnsetPartnerAddress2
`func (o *InvoiceResponseInvoice) UnsetPartnerAddress2()`

UnsetPartnerAddress2 ensures that no value is present for PartnerAddress2, not even an explicit nil
### GetPartnerContactInfo

`func (o *InvoiceResponseInvoice) GetPartnerContactInfo() string`

GetPartnerContactInfo returns the PartnerContactInfo field if non-nil, zero value otherwise.

### GetPartnerContactInfoOk

`func (o *InvoiceResponseInvoice) GetPartnerContactInfoOk() (*string, bool)`

GetPartnerContactInfoOk returns a tuple with the PartnerContactInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerContactInfo

`func (o *InvoiceResponseInvoice) SetPartnerContactInfo(v string)`

SetPartnerContactInfo sets PartnerContactInfo field to given value.

### HasPartnerContactInfo

`func (o *InvoiceResponseInvoice) HasPartnerContactInfo() bool`

HasPartnerContactInfo returns a boolean if a field has been set.

### SetPartnerContactInfoNil

`func (o *InvoiceResponseInvoice) SetPartnerContactInfoNil(b bool)`

 SetPartnerContactInfoNil sets the value for PartnerContactInfo to be an explicit nil

### UnsetPartnerContactInfo
`func (o *InvoiceResponseInvoice) UnsetPartnerContactInfo()`

UnsetPartnerContactInfo ensures that no value is present for PartnerContactInfo, not even an explicit nil
### GetCompanyName

`func (o *InvoiceResponseInvoice) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *InvoiceResponseInvoice) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *InvoiceResponseInvoice) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.


### GetCompanyZipcode

`func (o *InvoiceResponseInvoice) GetCompanyZipcode() string`

GetCompanyZipcode returns the CompanyZipcode field if non-nil, zero value otherwise.

### GetCompanyZipcodeOk

`func (o *InvoiceResponseInvoice) GetCompanyZipcodeOk() (*string, bool)`

GetCompanyZipcodeOk returns a tuple with the CompanyZipcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyZipcode

`func (o *InvoiceResponseInvoice) SetCompanyZipcode(v string)`

SetCompanyZipcode sets CompanyZipcode field to given value.

### HasCompanyZipcode

`func (o *InvoiceResponseInvoice) HasCompanyZipcode() bool`

HasCompanyZipcode returns a boolean if a field has been set.

### SetCompanyZipcodeNil

`func (o *InvoiceResponseInvoice) SetCompanyZipcodeNil(b bool)`

 SetCompanyZipcodeNil sets the value for CompanyZipcode to be an explicit nil

### UnsetCompanyZipcode
`func (o *InvoiceResponseInvoice) UnsetCompanyZipcode()`

UnsetCompanyZipcode ensures that no value is present for CompanyZipcode, not even an explicit nil
### GetCompanyPrefectureCode

`func (o *InvoiceResponseInvoice) GetCompanyPrefectureCode() int32`

GetCompanyPrefectureCode returns the CompanyPrefectureCode field if non-nil, zero value otherwise.

### GetCompanyPrefectureCodeOk

`func (o *InvoiceResponseInvoice) GetCompanyPrefectureCodeOk() (*int32, bool)`

GetCompanyPrefectureCodeOk returns a tuple with the CompanyPrefectureCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyPrefectureCode

`func (o *InvoiceResponseInvoice) SetCompanyPrefectureCode(v int32)`

SetCompanyPrefectureCode sets CompanyPrefectureCode field to given value.

### HasCompanyPrefectureCode

`func (o *InvoiceResponseInvoice) HasCompanyPrefectureCode() bool`

HasCompanyPrefectureCode returns a boolean if a field has been set.

### SetCompanyPrefectureCodeNil

`func (o *InvoiceResponseInvoice) SetCompanyPrefectureCodeNil(b bool)`

 SetCompanyPrefectureCodeNil sets the value for CompanyPrefectureCode to be an explicit nil

### UnsetCompanyPrefectureCode
`func (o *InvoiceResponseInvoice) UnsetCompanyPrefectureCode()`

UnsetCompanyPrefectureCode ensures that no value is present for CompanyPrefectureCode, not even an explicit nil
### GetCompanyPrefectureName

`func (o *InvoiceResponseInvoice) GetCompanyPrefectureName() string`

GetCompanyPrefectureName returns the CompanyPrefectureName field if non-nil, zero value otherwise.

### GetCompanyPrefectureNameOk

`func (o *InvoiceResponseInvoice) GetCompanyPrefectureNameOk() (*string, bool)`

GetCompanyPrefectureNameOk returns a tuple with the CompanyPrefectureName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyPrefectureName

`func (o *InvoiceResponseInvoice) SetCompanyPrefectureName(v string)`

SetCompanyPrefectureName sets CompanyPrefectureName field to given value.

### HasCompanyPrefectureName

`func (o *InvoiceResponseInvoice) HasCompanyPrefectureName() bool`

HasCompanyPrefectureName returns a boolean if a field has been set.

### SetCompanyPrefectureNameNil

`func (o *InvoiceResponseInvoice) SetCompanyPrefectureNameNil(b bool)`

 SetCompanyPrefectureNameNil sets the value for CompanyPrefectureName to be an explicit nil

### UnsetCompanyPrefectureName
`func (o *InvoiceResponseInvoice) UnsetCompanyPrefectureName()`

UnsetCompanyPrefectureName ensures that no value is present for CompanyPrefectureName, not even an explicit nil
### GetCompanyAddress1

`func (o *InvoiceResponseInvoice) GetCompanyAddress1() string`

GetCompanyAddress1 returns the CompanyAddress1 field if non-nil, zero value otherwise.

### GetCompanyAddress1Ok

`func (o *InvoiceResponseInvoice) GetCompanyAddress1Ok() (*string, bool)`

GetCompanyAddress1Ok returns a tuple with the CompanyAddress1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyAddress1

`func (o *InvoiceResponseInvoice) SetCompanyAddress1(v string)`

SetCompanyAddress1 sets CompanyAddress1 field to given value.

### HasCompanyAddress1

`func (o *InvoiceResponseInvoice) HasCompanyAddress1() bool`

HasCompanyAddress1 returns a boolean if a field has been set.

### SetCompanyAddress1Nil

`func (o *InvoiceResponseInvoice) SetCompanyAddress1Nil(b bool)`

 SetCompanyAddress1Nil sets the value for CompanyAddress1 to be an explicit nil

### UnsetCompanyAddress1
`func (o *InvoiceResponseInvoice) UnsetCompanyAddress1()`

UnsetCompanyAddress1 ensures that no value is present for CompanyAddress1, not even an explicit nil
### GetCompanyAddress2

`func (o *InvoiceResponseInvoice) GetCompanyAddress2() string`

GetCompanyAddress2 returns the CompanyAddress2 field if non-nil, zero value otherwise.

### GetCompanyAddress2Ok

`func (o *InvoiceResponseInvoice) GetCompanyAddress2Ok() (*string, bool)`

GetCompanyAddress2Ok returns a tuple with the CompanyAddress2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyAddress2

`func (o *InvoiceResponseInvoice) SetCompanyAddress2(v string)`

SetCompanyAddress2 sets CompanyAddress2 field to given value.

### HasCompanyAddress2

`func (o *InvoiceResponseInvoice) HasCompanyAddress2() bool`

HasCompanyAddress2 returns a boolean if a field has been set.

### SetCompanyAddress2Nil

`func (o *InvoiceResponseInvoice) SetCompanyAddress2Nil(b bool)`

 SetCompanyAddress2Nil sets the value for CompanyAddress2 to be an explicit nil

### UnsetCompanyAddress2
`func (o *InvoiceResponseInvoice) UnsetCompanyAddress2()`

UnsetCompanyAddress2 ensures that no value is present for CompanyAddress2, not even an explicit nil
### GetCompanyContactInfo

`func (o *InvoiceResponseInvoice) GetCompanyContactInfo() string`

GetCompanyContactInfo returns the CompanyContactInfo field if non-nil, zero value otherwise.

### GetCompanyContactInfoOk

`func (o *InvoiceResponseInvoice) GetCompanyContactInfoOk() (*string, bool)`

GetCompanyContactInfoOk returns a tuple with the CompanyContactInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyContactInfo

`func (o *InvoiceResponseInvoice) SetCompanyContactInfo(v string)`

SetCompanyContactInfo sets CompanyContactInfo field to given value.

### HasCompanyContactInfo

`func (o *InvoiceResponseInvoice) HasCompanyContactInfo() bool`

HasCompanyContactInfo returns a boolean if a field has been set.

### SetCompanyContactInfoNil

`func (o *InvoiceResponseInvoice) SetCompanyContactInfoNil(b bool)`

 SetCompanyContactInfoNil sets the value for CompanyContactInfo to be an explicit nil

### UnsetCompanyContactInfo
`func (o *InvoiceResponseInvoice) UnsetCompanyContactInfo()`

UnsetCompanyContactInfo ensures that no value is present for CompanyContactInfo, not even an explicit nil
### GetPaymentType

`func (o *InvoiceResponseInvoice) GetPaymentType() string`

GetPaymentType returns the PaymentType field if non-nil, zero value otherwise.

### GetPaymentTypeOk

`func (o *InvoiceResponseInvoice) GetPaymentTypeOk() (*string, bool)`

GetPaymentTypeOk returns a tuple with the PaymentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentType

`func (o *InvoiceResponseInvoice) SetPaymentType(v string)`

SetPaymentType sets PaymentType field to given value.


### GetPaymentBankInfo

`func (o *InvoiceResponseInvoice) GetPaymentBankInfo() string`

GetPaymentBankInfo returns the PaymentBankInfo field if non-nil, zero value otherwise.

### GetPaymentBankInfoOk

`func (o *InvoiceResponseInvoice) GetPaymentBankInfoOk() (*string, bool)`

GetPaymentBankInfoOk returns a tuple with the PaymentBankInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentBankInfo

`func (o *InvoiceResponseInvoice) SetPaymentBankInfo(v string)`

SetPaymentBankInfo sets PaymentBankInfo field to given value.

### HasPaymentBankInfo

`func (o *InvoiceResponseInvoice) HasPaymentBankInfo() bool`

HasPaymentBankInfo returns a boolean if a field has been set.

### SetPaymentBankInfoNil

`func (o *InvoiceResponseInvoice) SetPaymentBankInfoNil(b bool)`

 SetPaymentBankInfoNil sets the value for PaymentBankInfo to be an explicit nil

### UnsetPaymentBankInfo
`func (o *InvoiceResponseInvoice) UnsetPaymentBankInfo()`

UnsetPaymentBankInfo ensures that no value is present for PaymentBankInfo, not even an explicit nil
### GetMessage

`func (o *InvoiceResponseInvoice) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *InvoiceResponseInvoice) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *InvoiceResponseInvoice) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *InvoiceResponseInvoice) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *InvoiceResponseInvoice) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *InvoiceResponseInvoice) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetNotes

`func (o *InvoiceResponseInvoice) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *InvoiceResponseInvoice) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *InvoiceResponseInvoice) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *InvoiceResponseInvoice) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *InvoiceResponseInvoice) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *InvoiceResponseInvoice) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetInvoiceLayout

`func (o *InvoiceResponseInvoice) GetInvoiceLayout() string`

GetInvoiceLayout returns the InvoiceLayout field if non-nil, zero value otherwise.

### GetInvoiceLayoutOk

`func (o *InvoiceResponseInvoice) GetInvoiceLayoutOk() (*string, bool)`

GetInvoiceLayoutOk returns a tuple with the InvoiceLayout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceLayout

`func (o *InvoiceResponseInvoice) SetInvoiceLayout(v string)`

SetInvoiceLayout sets InvoiceLayout field to given value.


### GetTaxEntryMethod

`func (o *InvoiceResponseInvoice) GetTaxEntryMethod() string`

GetTaxEntryMethod returns the TaxEntryMethod field if non-nil, zero value otherwise.

### GetTaxEntryMethodOk

`func (o *InvoiceResponseInvoice) GetTaxEntryMethodOk() (*string, bool)`

GetTaxEntryMethodOk returns a tuple with the TaxEntryMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxEntryMethod

`func (o *InvoiceResponseInvoice) SetTaxEntryMethod(v string)`

SetTaxEntryMethod sets TaxEntryMethod field to given value.


### GetDealId

`func (o *InvoiceResponseInvoice) GetDealId() int32`

GetDealId returns the DealId field if non-nil, zero value otherwise.

### GetDealIdOk

`func (o *InvoiceResponseInvoice) GetDealIdOk() (*int32, bool)`

GetDealIdOk returns a tuple with the DealId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDealId

`func (o *InvoiceResponseInvoice) SetDealId(v int32)`

SetDealId sets DealId field to given value.

### HasDealId

`func (o *InvoiceResponseInvoice) HasDealId() bool`

HasDealId returns a boolean if a field has been set.

### SetDealIdNil

`func (o *InvoiceResponseInvoice) SetDealIdNil(b bool)`

 SetDealIdNil sets the value for DealId to be an explicit nil

### UnsetDealId
`func (o *InvoiceResponseInvoice) UnsetDealId()`

UnsetDealId ensures that no value is present for DealId, not even an explicit nil
### GetInvoiceContents

`func (o *InvoiceResponseInvoice) GetInvoiceContents() []InvoiceIndexResponseInvoiceContents`

GetInvoiceContents returns the InvoiceContents field if non-nil, zero value otherwise.

### GetInvoiceContentsOk

`func (o *InvoiceResponseInvoice) GetInvoiceContentsOk() (*[]InvoiceIndexResponseInvoiceContents, bool)`

GetInvoiceContentsOk returns a tuple with the InvoiceContents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceContents

`func (o *InvoiceResponseInvoice) SetInvoiceContents(v []InvoiceIndexResponseInvoiceContents)`

SetInvoiceContents sets InvoiceContents field to given value.

### HasInvoiceContents

`func (o *InvoiceResponseInvoice) HasInvoiceContents() bool`

HasInvoiceContents returns a boolean if a field has been set.

### GetTotalAmountPerVatRate

`func (o *InvoiceResponseInvoice) GetTotalAmountPerVatRate() InvoiceIndexResponseTotalAmountPerVatRate`

GetTotalAmountPerVatRate returns the TotalAmountPerVatRate field if non-nil, zero value otherwise.

### GetTotalAmountPerVatRateOk

`func (o *InvoiceResponseInvoice) GetTotalAmountPerVatRateOk() (*InvoiceIndexResponseTotalAmountPerVatRate, bool)`

GetTotalAmountPerVatRateOk returns a tuple with the TotalAmountPerVatRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmountPerVatRate

`func (o *InvoiceResponseInvoice) SetTotalAmountPerVatRate(v InvoiceIndexResponseTotalAmountPerVatRate)`

SetTotalAmountPerVatRate sets TotalAmountPerVatRate field to given value.


### GetRelatedQuotationIds

`func (o *InvoiceResponseInvoice) GetRelatedQuotationIds() []int32`

GetRelatedQuotationIds returns the RelatedQuotationIds field if non-nil, zero value otherwise.

### GetRelatedQuotationIdsOk

`func (o *InvoiceResponseInvoice) GetRelatedQuotationIdsOk() (*[]int32, bool)`

GetRelatedQuotationIdsOk returns a tuple with the RelatedQuotationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedQuotationIds

`func (o *InvoiceResponseInvoice) SetRelatedQuotationIds(v []int32)`

SetRelatedQuotationIds sets RelatedQuotationIds field to given value.

### HasRelatedQuotationIds

`func (o *InvoiceResponseInvoice) HasRelatedQuotationIds() bool`

HasRelatedQuotationIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


