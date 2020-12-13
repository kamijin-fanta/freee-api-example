# InvoiceIndexResponseInvoices

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

## Methods

### NewInvoiceIndexResponseInvoices

`func NewInvoiceIndexResponseInvoices(id int32, companyId int32, issueDate string, partnerId NullableInt32, invoiceNumber string, totalAmount int32, invoiceStatus string, postingStatus string, companyName string, paymentType string, invoiceLayout string, taxEntryMethod string, totalAmountPerVatRate InvoiceIndexResponseTotalAmountPerVatRate, ) *InvoiceIndexResponseInvoices`

NewInvoiceIndexResponseInvoices instantiates a new InvoiceIndexResponseInvoices object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceIndexResponseInvoicesWithDefaults

`func NewInvoiceIndexResponseInvoicesWithDefaults() *InvoiceIndexResponseInvoices`

NewInvoiceIndexResponseInvoicesWithDefaults instantiates a new InvoiceIndexResponseInvoices object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InvoiceIndexResponseInvoices) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InvoiceIndexResponseInvoices) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InvoiceIndexResponseInvoices) SetId(v int32)`

SetId sets Id field to given value.


### GetCompanyId

`func (o *InvoiceIndexResponseInvoices) GetCompanyId() int32`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *InvoiceIndexResponseInvoices) GetCompanyIdOk() (*int32, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *InvoiceIndexResponseInvoices) SetCompanyId(v int32)`

SetCompanyId sets CompanyId field to given value.


### GetIssueDate

`func (o *InvoiceIndexResponseInvoices) GetIssueDate() string`

GetIssueDate returns the IssueDate field if non-nil, zero value otherwise.

### GetIssueDateOk

`func (o *InvoiceIndexResponseInvoices) GetIssueDateOk() (*string, bool)`

GetIssueDateOk returns a tuple with the IssueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueDate

`func (o *InvoiceIndexResponseInvoices) SetIssueDate(v string)`

SetIssueDate sets IssueDate field to given value.


### GetPartnerId

`func (o *InvoiceIndexResponseInvoices) GetPartnerId() int32`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *InvoiceIndexResponseInvoices) GetPartnerIdOk() (*int32, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *InvoiceIndexResponseInvoices) SetPartnerId(v int32)`

SetPartnerId sets PartnerId field to given value.


### SetPartnerIdNil

`func (o *InvoiceIndexResponseInvoices) SetPartnerIdNil(b bool)`

 SetPartnerIdNil sets the value for PartnerId to be an explicit nil

### UnsetPartnerId
`func (o *InvoiceIndexResponseInvoices) UnsetPartnerId()`

UnsetPartnerId ensures that no value is present for PartnerId, not even an explicit nil
### GetPartnerCode

`func (o *InvoiceIndexResponseInvoices) GetPartnerCode() string`

GetPartnerCode returns the PartnerCode field if non-nil, zero value otherwise.

### GetPartnerCodeOk

`func (o *InvoiceIndexResponseInvoices) GetPartnerCodeOk() (*string, bool)`

GetPartnerCodeOk returns a tuple with the PartnerCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerCode

`func (o *InvoiceIndexResponseInvoices) SetPartnerCode(v string)`

SetPartnerCode sets PartnerCode field to given value.

### HasPartnerCode

`func (o *InvoiceIndexResponseInvoices) HasPartnerCode() bool`

HasPartnerCode returns a boolean if a field has been set.

### SetPartnerCodeNil

`func (o *InvoiceIndexResponseInvoices) SetPartnerCodeNil(b bool)`

 SetPartnerCodeNil sets the value for PartnerCode to be an explicit nil

### UnsetPartnerCode
`func (o *InvoiceIndexResponseInvoices) UnsetPartnerCode()`

UnsetPartnerCode ensures that no value is present for PartnerCode, not even an explicit nil
### GetInvoiceNumber

`func (o *InvoiceIndexResponseInvoices) GetInvoiceNumber() string`

GetInvoiceNumber returns the InvoiceNumber field if non-nil, zero value otherwise.

### GetInvoiceNumberOk

`func (o *InvoiceIndexResponseInvoices) GetInvoiceNumberOk() (*string, bool)`

GetInvoiceNumberOk returns a tuple with the InvoiceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceNumber

`func (o *InvoiceIndexResponseInvoices) SetInvoiceNumber(v string)`

SetInvoiceNumber sets InvoiceNumber field to given value.


### GetTitle

`func (o *InvoiceIndexResponseInvoices) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *InvoiceIndexResponseInvoices) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *InvoiceIndexResponseInvoices) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *InvoiceIndexResponseInvoices) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *InvoiceIndexResponseInvoices) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *InvoiceIndexResponseInvoices) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDueDate

`func (o *InvoiceIndexResponseInvoices) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *InvoiceIndexResponseInvoices) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *InvoiceIndexResponseInvoices) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *InvoiceIndexResponseInvoices) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### SetDueDateNil

`func (o *InvoiceIndexResponseInvoices) SetDueDateNil(b bool)`

 SetDueDateNil sets the value for DueDate to be an explicit nil

### UnsetDueDate
`func (o *InvoiceIndexResponseInvoices) UnsetDueDate()`

UnsetDueDate ensures that no value is present for DueDate, not even an explicit nil
### GetTotalAmount

`func (o *InvoiceIndexResponseInvoices) GetTotalAmount() int32`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *InvoiceIndexResponseInvoices) GetTotalAmountOk() (*int32, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *InvoiceIndexResponseInvoices) SetTotalAmount(v int32)`

SetTotalAmount sets TotalAmount field to given value.


### GetTotalVat

`func (o *InvoiceIndexResponseInvoices) GetTotalVat() int32`

GetTotalVat returns the TotalVat field if non-nil, zero value otherwise.

### GetTotalVatOk

`func (o *InvoiceIndexResponseInvoices) GetTotalVatOk() (*int32, bool)`

GetTotalVatOk returns a tuple with the TotalVat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalVat

`func (o *InvoiceIndexResponseInvoices) SetTotalVat(v int32)`

SetTotalVat sets TotalVat field to given value.

### HasTotalVat

`func (o *InvoiceIndexResponseInvoices) HasTotalVat() bool`

HasTotalVat returns a boolean if a field has been set.

### GetSubTotal

`func (o *InvoiceIndexResponseInvoices) GetSubTotal() int32`

GetSubTotal returns the SubTotal field if non-nil, zero value otherwise.

### GetSubTotalOk

`func (o *InvoiceIndexResponseInvoices) GetSubTotalOk() (*int32, bool)`

GetSubTotalOk returns a tuple with the SubTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubTotal

`func (o *InvoiceIndexResponseInvoices) SetSubTotal(v int32)`

SetSubTotal sets SubTotal field to given value.

### HasSubTotal

`func (o *InvoiceIndexResponseInvoices) HasSubTotal() bool`

HasSubTotal returns a boolean if a field has been set.

### GetBookingDate

`func (o *InvoiceIndexResponseInvoices) GetBookingDate() string`

GetBookingDate returns the BookingDate field if non-nil, zero value otherwise.

### GetBookingDateOk

`func (o *InvoiceIndexResponseInvoices) GetBookingDateOk() (*string, bool)`

GetBookingDateOk returns a tuple with the BookingDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBookingDate

`func (o *InvoiceIndexResponseInvoices) SetBookingDate(v string)`

SetBookingDate sets BookingDate field to given value.

### HasBookingDate

`func (o *InvoiceIndexResponseInvoices) HasBookingDate() bool`

HasBookingDate returns a boolean if a field has been set.

### SetBookingDateNil

`func (o *InvoiceIndexResponseInvoices) SetBookingDateNil(b bool)`

 SetBookingDateNil sets the value for BookingDate to be an explicit nil

### UnsetBookingDate
`func (o *InvoiceIndexResponseInvoices) UnsetBookingDate()`

UnsetBookingDate ensures that no value is present for BookingDate, not even an explicit nil
### GetDescription

`func (o *InvoiceIndexResponseInvoices) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InvoiceIndexResponseInvoices) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InvoiceIndexResponseInvoices) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InvoiceIndexResponseInvoices) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *InvoiceIndexResponseInvoices) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *InvoiceIndexResponseInvoices) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetInvoiceStatus

`func (o *InvoiceIndexResponseInvoices) GetInvoiceStatus() string`

GetInvoiceStatus returns the InvoiceStatus field if non-nil, zero value otherwise.

### GetInvoiceStatusOk

`func (o *InvoiceIndexResponseInvoices) GetInvoiceStatusOk() (*string, bool)`

GetInvoiceStatusOk returns a tuple with the InvoiceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceStatus

`func (o *InvoiceIndexResponseInvoices) SetInvoiceStatus(v string)`

SetInvoiceStatus sets InvoiceStatus field to given value.


### GetPaymentStatus

`func (o *InvoiceIndexResponseInvoices) GetPaymentStatus() string`

GetPaymentStatus returns the PaymentStatus field if non-nil, zero value otherwise.

### GetPaymentStatusOk

`func (o *InvoiceIndexResponseInvoices) GetPaymentStatusOk() (*string, bool)`

GetPaymentStatusOk returns a tuple with the PaymentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentStatus

`func (o *InvoiceIndexResponseInvoices) SetPaymentStatus(v string)`

SetPaymentStatus sets PaymentStatus field to given value.

### HasPaymentStatus

`func (o *InvoiceIndexResponseInvoices) HasPaymentStatus() bool`

HasPaymentStatus returns a boolean if a field has been set.

### GetPaymentDate

`func (o *InvoiceIndexResponseInvoices) GetPaymentDate() string`

GetPaymentDate returns the PaymentDate field if non-nil, zero value otherwise.

### GetPaymentDateOk

`func (o *InvoiceIndexResponseInvoices) GetPaymentDateOk() (*string, bool)`

GetPaymentDateOk returns a tuple with the PaymentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentDate

`func (o *InvoiceIndexResponseInvoices) SetPaymentDate(v string)`

SetPaymentDate sets PaymentDate field to given value.

### HasPaymentDate

`func (o *InvoiceIndexResponseInvoices) HasPaymentDate() bool`

HasPaymentDate returns a boolean if a field has been set.

### SetPaymentDateNil

`func (o *InvoiceIndexResponseInvoices) SetPaymentDateNil(b bool)`

 SetPaymentDateNil sets the value for PaymentDate to be an explicit nil

### UnsetPaymentDate
`func (o *InvoiceIndexResponseInvoices) UnsetPaymentDate()`

UnsetPaymentDate ensures that no value is present for PaymentDate, not even an explicit nil
### GetWebPublishedAt

`func (o *InvoiceIndexResponseInvoices) GetWebPublishedAt() string`

GetWebPublishedAt returns the WebPublishedAt field if non-nil, zero value otherwise.

### GetWebPublishedAtOk

`func (o *InvoiceIndexResponseInvoices) GetWebPublishedAtOk() (*string, bool)`

GetWebPublishedAtOk returns a tuple with the WebPublishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebPublishedAt

`func (o *InvoiceIndexResponseInvoices) SetWebPublishedAt(v string)`

SetWebPublishedAt sets WebPublishedAt field to given value.

### HasWebPublishedAt

`func (o *InvoiceIndexResponseInvoices) HasWebPublishedAt() bool`

HasWebPublishedAt returns a boolean if a field has been set.

### SetWebPublishedAtNil

`func (o *InvoiceIndexResponseInvoices) SetWebPublishedAtNil(b bool)`

 SetWebPublishedAtNil sets the value for WebPublishedAt to be an explicit nil

### UnsetWebPublishedAt
`func (o *InvoiceIndexResponseInvoices) UnsetWebPublishedAt()`

UnsetWebPublishedAt ensures that no value is present for WebPublishedAt, not even an explicit nil
### GetWebDownloadedAt

`func (o *InvoiceIndexResponseInvoices) GetWebDownloadedAt() string`

GetWebDownloadedAt returns the WebDownloadedAt field if non-nil, zero value otherwise.

### GetWebDownloadedAtOk

`func (o *InvoiceIndexResponseInvoices) GetWebDownloadedAtOk() (*string, bool)`

GetWebDownloadedAtOk returns a tuple with the WebDownloadedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebDownloadedAt

`func (o *InvoiceIndexResponseInvoices) SetWebDownloadedAt(v string)`

SetWebDownloadedAt sets WebDownloadedAt field to given value.

### HasWebDownloadedAt

`func (o *InvoiceIndexResponseInvoices) HasWebDownloadedAt() bool`

HasWebDownloadedAt returns a boolean if a field has been set.

### SetWebDownloadedAtNil

`func (o *InvoiceIndexResponseInvoices) SetWebDownloadedAtNil(b bool)`

 SetWebDownloadedAtNil sets the value for WebDownloadedAt to be an explicit nil

### UnsetWebDownloadedAt
`func (o *InvoiceIndexResponseInvoices) UnsetWebDownloadedAt()`

UnsetWebDownloadedAt ensures that no value is present for WebDownloadedAt, not even an explicit nil
### GetWebConfirmedAt

`func (o *InvoiceIndexResponseInvoices) GetWebConfirmedAt() string`

GetWebConfirmedAt returns the WebConfirmedAt field if non-nil, zero value otherwise.

### GetWebConfirmedAtOk

`func (o *InvoiceIndexResponseInvoices) GetWebConfirmedAtOk() (*string, bool)`

GetWebConfirmedAtOk returns a tuple with the WebConfirmedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebConfirmedAt

`func (o *InvoiceIndexResponseInvoices) SetWebConfirmedAt(v string)`

SetWebConfirmedAt sets WebConfirmedAt field to given value.

### HasWebConfirmedAt

`func (o *InvoiceIndexResponseInvoices) HasWebConfirmedAt() bool`

HasWebConfirmedAt returns a boolean if a field has been set.

### SetWebConfirmedAtNil

`func (o *InvoiceIndexResponseInvoices) SetWebConfirmedAtNil(b bool)`

 SetWebConfirmedAtNil sets the value for WebConfirmedAt to be an explicit nil

### UnsetWebConfirmedAt
`func (o *InvoiceIndexResponseInvoices) UnsetWebConfirmedAt()`

UnsetWebConfirmedAt ensures that no value is present for WebConfirmedAt, not even an explicit nil
### GetMailSentAt

`func (o *InvoiceIndexResponseInvoices) GetMailSentAt() string`

GetMailSentAt returns the MailSentAt field if non-nil, zero value otherwise.

### GetMailSentAtOk

`func (o *InvoiceIndexResponseInvoices) GetMailSentAtOk() (*string, bool)`

GetMailSentAtOk returns a tuple with the MailSentAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailSentAt

`func (o *InvoiceIndexResponseInvoices) SetMailSentAt(v string)`

SetMailSentAt sets MailSentAt field to given value.

### HasMailSentAt

`func (o *InvoiceIndexResponseInvoices) HasMailSentAt() bool`

HasMailSentAt returns a boolean if a field has been set.

### SetMailSentAtNil

`func (o *InvoiceIndexResponseInvoices) SetMailSentAtNil(b bool)`

 SetMailSentAtNil sets the value for MailSentAt to be an explicit nil

### UnsetMailSentAt
`func (o *InvoiceIndexResponseInvoices) UnsetMailSentAt()`

UnsetMailSentAt ensures that no value is present for MailSentAt, not even an explicit nil
### GetPostingStatus

`func (o *InvoiceIndexResponseInvoices) GetPostingStatus() string`

GetPostingStatus returns the PostingStatus field if non-nil, zero value otherwise.

### GetPostingStatusOk

`func (o *InvoiceIndexResponseInvoices) GetPostingStatusOk() (*string, bool)`

GetPostingStatusOk returns a tuple with the PostingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostingStatus

`func (o *InvoiceIndexResponseInvoices) SetPostingStatus(v string)`

SetPostingStatus sets PostingStatus field to given value.


### GetPartnerName

`func (o *InvoiceIndexResponseInvoices) GetPartnerName() string`

GetPartnerName returns the PartnerName field if non-nil, zero value otherwise.

### GetPartnerNameOk

`func (o *InvoiceIndexResponseInvoices) GetPartnerNameOk() (*string, bool)`

GetPartnerNameOk returns a tuple with the PartnerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerName

`func (o *InvoiceIndexResponseInvoices) SetPartnerName(v string)`

SetPartnerName sets PartnerName field to given value.

### HasPartnerName

`func (o *InvoiceIndexResponseInvoices) HasPartnerName() bool`

HasPartnerName returns a boolean if a field has been set.

### SetPartnerNameNil

`func (o *InvoiceIndexResponseInvoices) SetPartnerNameNil(b bool)`

 SetPartnerNameNil sets the value for PartnerName to be an explicit nil

### UnsetPartnerName
`func (o *InvoiceIndexResponseInvoices) UnsetPartnerName()`

UnsetPartnerName ensures that no value is present for PartnerName, not even an explicit nil
### GetPartnerDisplayName

`func (o *InvoiceIndexResponseInvoices) GetPartnerDisplayName() string`

GetPartnerDisplayName returns the PartnerDisplayName field if non-nil, zero value otherwise.

### GetPartnerDisplayNameOk

`func (o *InvoiceIndexResponseInvoices) GetPartnerDisplayNameOk() (*string, bool)`

GetPartnerDisplayNameOk returns a tuple with the PartnerDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerDisplayName

`func (o *InvoiceIndexResponseInvoices) SetPartnerDisplayName(v string)`

SetPartnerDisplayName sets PartnerDisplayName field to given value.

### HasPartnerDisplayName

`func (o *InvoiceIndexResponseInvoices) HasPartnerDisplayName() bool`

HasPartnerDisplayName returns a boolean if a field has been set.

### SetPartnerDisplayNameNil

`func (o *InvoiceIndexResponseInvoices) SetPartnerDisplayNameNil(b bool)`

 SetPartnerDisplayNameNil sets the value for PartnerDisplayName to be an explicit nil

### UnsetPartnerDisplayName
`func (o *InvoiceIndexResponseInvoices) UnsetPartnerDisplayName()`

UnsetPartnerDisplayName ensures that no value is present for PartnerDisplayName, not even an explicit nil
### GetPartnerTitle

`func (o *InvoiceIndexResponseInvoices) GetPartnerTitle() string`

GetPartnerTitle returns the PartnerTitle field if non-nil, zero value otherwise.

### GetPartnerTitleOk

`func (o *InvoiceIndexResponseInvoices) GetPartnerTitleOk() (*string, bool)`

GetPartnerTitleOk returns a tuple with the PartnerTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerTitle

`func (o *InvoiceIndexResponseInvoices) SetPartnerTitle(v string)`

SetPartnerTitle sets PartnerTitle field to given value.

### HasPartnerTitle

`func (o *InvoiceIndexResponseInvoices) HasPartnerTitle() bool`

HasPartnerTitle returns a boolean if a field has been set.

### SetPartnerTitleNil

`func (o *InvoiceIndexResponseInvoices) SetPartnerTitleNil(b bool)`

 SetPartnerTitleNil sets the value for PartnerTitle to be an explicit nil

### UnsetPartnerTitle
`func (o *InvoiceIndexResponseInvoices) UnsetPartnerTitle()`

UnsetPartnerTitle ensures that no value is present for PartnerTitle, not even an explicit nil
### GetPartnerZipcode

`func (o *InvoiceIndexResponseInvoices) GetPartnerZipcode() string`

GetPartnerZipcode returns the PartnerZipcode field if non-nil, zero value otherwise.

### GetPartnerZipcodeOk

`func (o *InvoiceIndexResponseInvoices) GetPartnerZipcodeOk() (*string, bool)`

GetPartnerZipcodeOk returns a tuple with the PartnerZipcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerZipcode

`func (o *InvoiceIndexResponseInvoices) SetPartnerZipcode(v string)`

SetPartnerZipcode sets PartnerZipcode field to given value.

### HasPartnerZipcode

`func (o *InvoiceIndexResponseInvoices) HasPartnerZipcode() bool`

HasPartnerZipcode returns a boolean if a field has been set.

### SetPartnerZipcodeNil

`func (o *InvoiceIndexResponseInvoices) SetPartnerZipcodeNil(b bool)`

 SetPartnerZipcodeNil sets the value for PartnerZipcode to be an explicit nil

### UnsetPartnerZipcode
`func (o *InvoiceIndexResponseInvoices) UnsetPartnerZipcode()`

UnsetPartnerZipcode ensures that no value is present for PartnerZipcode, not even an explicit nil
### GetPartnerPrefectureCode

`func (o *InvoiceIndexResponseInvoices) GetPartnerPrefectureCode() int32`

GetPartnerPrefectureCode returns the PartnerPrefectureCode field if non-nil, zero value otherwise.

### GetPartnerPrefectureCodeOk

`func (o *InvoiceIndexResponseInvoices) GetPartnerPrefectureCodeOk() (*int32, bool)`

GetPartnerPrefectureCodeOk returns a tuple with the PartnerPrefectureCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerPrefectureCode

`func (o *InvoiceIndexResponseInvoices) SetPartnerPrefectureCode(v int32)`

SetPartnerPrefectureCode sets PartnerPrefectureCode field to given value.

### HasPartnerPrefectureCode

`func (o *InvoiceIndexResponseInvoices) HasPartnerPrefectureCode() bool`

HasPartnerPrefectureCode returns a boolean if a field has been set.

### SetPartnerPrefectureCodeNil

`func (o *InvoiceIndexResponseInvoices) SetPartnerPrefectureCodeNil(b bool)`

 SetPartnerPrefectureCodeNil sets the value for PartnerPrefectureCode to be an explicit nil

### UnsetPartnerPrefectureCode
`func (o *InvoiceIndexResponseInvoices) UnsetPartnerPrefectureCode()`

UnsetPartnerPrefectureCode ensures that no value is present for PartnerPrefectureCode, not even an explicit nil
### GetPartnerPrefectureName

`func (o *InvoiceIndexResponseInvoices) GetPartnerPrefectureName() string`

GetPartnerPrefectureName returns the PartnerPrefectureName field if non-nil, zero value otherwise.

### GetPartnerPrefectureNameOk

`func (o *InvoiceIndexResponseInvoices) GetPartnerPrefectureNameOk() (*string, bool)`

GetPartnerPrefectureNameOk returns a tuple with the PartnerPrefectureName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerPrefectureName

`func (o *InvoiceIndexResponseInvoices) SetPartnerPrefectureName(v string)`

SetPartnerPrefectureName sets PartnerPrefectureName field to given value.

### HasPartnerPrefectureName

`func (o *InvoiceIndexResponseInvoices) HasPartnerPrefectureName() bool`

HasPartnerPrefectureName returns a boolean if a field has been set.

### SetPartnerPrefectureNameNil

`func (o *InvoiceIndexResponseInvoices) SetPartnerPrefectureNameNil(b bool)`

 SetPartnerPrefectureNameNil sets the value for PartnerPrefectureName to be an explicit nil

### UnsetPartnerPrefectureName
`func (o *InvoiceIndexResponseInvoices) UnsetPartnerPrefectureName()`

UnsetPartnerPrefectureName ensures that no value is present for PartnerPrefectureName, not even an explicit nil
### GetPartnerAddress1

`func (o *InvoiceIndexResponseInvoices) GetPartnerAddress1() string`

GetPartnerAddress1 returns the PartnerAddress1 field if non-nil, zero value otherwise.

### GetPartnerAddress1Ok

`func (o *InvoiceIndexResponseInvoices) GetPartnerAddress1Ok() (*string, bool)`

GetPartnerAddress1Ok returns a tuple with the PartnerAddress1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerAddress1

`func (o *InvoiceIndexResponseInvoices) SetPartnerAddress1(v string)`

SetPartnerAddress1 sets PartnerAddress1 field to given value.

### HasPartnerAddress1

`func (o *InvoiceIndexResponseInvoices) HasPartnerAddress1() bool`

HasPartnerAddress1 returns a boolean if a field has been set.

### SetPartnerAddress1Nil

`func (o *InvoiceIndexResponseInvoices) SetPartnerAddress1Nil(b bool)`

 SetPartnerAddress1Nil sets the value for PartnerAddress1 to be an explicit nil

### UnsetPartnerAddress1
`func (o *InvoiceIndexResponseInvoices) UnsetPartnerAddress1()`

UnsetPartnerAddress1 ensures that no value is present for PartnerAddress1, not even an explicit nil
### GetPartnerAddress2

`func (o *InvoiceIndexResponseInvoices) GetPartnerAddress2() string`

GetPartnerAddress2 returns the PartnerAddress2 field if non-nil, zero value otherwise.

### GetPartnerAddress2Ok

`func (o *InvoiceIndexResponseInvoices) GetPartnerAddress2Ok() (*string, bool)`

GetPartnerAddress2Ok returns a tuple with the PartnerAddress2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerAddress2

`func (o *InvoiceIndexResponseInvoices) SetPartnerAddress2(v string)`

SetPartnerAddress2 sets PartnerAddress2 field to given value.

### HasPartnerAddress2

`func (o *InvoiceIndexResponseInvoices) HasPartnerAddress2() bool`

HasPartnerAddress2 returns a boolean if a field has been set.

### SetPartnerAddress2Nil

`func (o *InvoiceIndexResponseInvoices) SetPartnerAddress2Nil(b bool)`

 SetPartnerAddress2Nil sets the value for PartnerAddress2 to be an explicit nil

### UnsetPartnerAddress2
`func (o *InvoiceIndexResponseInvoices) UnsetPartnerAddress2()`

UnsetPartnerAddress2 ensures that no value is present for PartnerAddress2, not even an explicit nil
### GetPartnerContactInfo

`func (o *InvoiceIndexResponseInvoices) GetPartnerContactInfo() string`

GetPartnerContactInfo returns the PartnerContactInfo field if non-nil, zero value otherwise.

### GetPartnerContactInfoOk

`func (o *InvoiceIndexResponseInvoices) GetPartnerContactInfoOk() (*string, bool)`

GetPartnerContactInfoOk returns a tuple with the PartnerContactInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerContactInfo

`func (o *InvoiceIndexResponseInvoices) SetPartnerContactInfo(v string)`

SetPartnerContactInfo sets PartnerContactInfo field to given value.

### HasPartnerContactInfo

`func (o *InvoiceIndexResponseInvoices) HasPartnerContactInfo() bool`

HasPartnerContactInfo returns a boolean if a field has been set.

### SetPartnerContactInfoNil

`func (o *InvoiceIndexResponseInvoices) SetPartnerContactInfoNil(b bool)`

 SetPartnerContactInfoNil sets the value for PartnerContactInfo to be an explicit nil

### UnsetPartnerContactInfo
`func (o *InvoiceIndexResponseInvoices) UnsetPartnerContactInfo()`

UnsetPartnerContactInfo ensures that no value is present for PartnerContactInfo, not even an explicit nil
### GetCompanyName

`func (o *InvoiceIndexResponseInvoices) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *InvoiceIndexResponseInvoices) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *InvoiceIndexResponseInvoices) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.


### GetCompanyZipcode

`func (o *InvoiceIndexResponseInvoices) GetCompanyZipcode() string`

GetCompanyZipcode returns the CompanyZipcode field if non-nil, zero value otherwise.

### GetCompanyZipcodeOk

`func (o *InvoiceIndexResponseInvoices) GetCompanyZipcodeOk() (*string, bool)`

GetCompanyZipcodeOk returns a tuple with the CompanyZipcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyZipcode

`func (o *InvoiceIndexResponseInvoices) SetCompanyZipcode(v string)`

SetCompanyZipcode sets CompanyZipcode field to given value.

### HasCompanyZipcode

`func (o *InvoiceIndexResponseInvoices) HasCompanyZipcode() bool`

HasCompanyZipcode returns a boolean if a field has been set.

### SetCompanyZipcodeNil

`func (o *InvoiceIndexResponseInvoices) SetCompanyZipcodeNil(b bool)`

 SetCompanyZipcodeNil sets the value for CompanyZipcode to be an explicit nil

### UnsetCompanyZipcode
`func (o *InvoiceIndexResponseInvoices) UnsetCompanyZipcode()`

UnsetCompanyZipcode ensures that no value is present for CompanyZipcode, not even an explicit nil
### GetCompanyPrefectureCode

`func (o *InvoiceIndexResponseInvoices) GetCompanyPrefectureCode() int32`

GetCompanyPrefectureCode returns the CompanyPrefectureCode field if non-nil, zero value otherwise.

### GetCompanyPrefectureCodeOk

`func (o *InvoiceIndexResponseInvoices) GetCompanyPrefectureCodeOk() (*int32, bool)`

GetCompanyPrefectureCodeOk returns a tuple with the CompanyPrefectureCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyPrefectureCode

`func (o *InvoiceIndexResponseInvoices) SetCompanyPrefectureCode(v int32)`

SetCompanyPrefectureCode sets CompanyPrefectureCode field to given value.

### HasCompanyPrefectureCode

`func (o *InvoiceIndexResponseInvoices) HasCompanyPrefectureCode() bool`

HasCompanyPrefectureCode returns a boolean if a field has been set.

### SetCompanyPrefectureCodeNil

`func (o *InvoiceIndexResponseInvoices) SetCompanyPrefectureCodeNil(b bool)`

 SetCompanyPrefectureCodeNil sets the value for CompanyPrefectureCode to be an explicit nil

### UnsetCompanyPrefectureCode
`func (o *InvoiceIndexResponseInvoices) UnsetCompanyPrefectureCode()`

UnsetCompanyPrefectureCode ensures that no value is present for CompanyPrefectureCode, not even an explicit nil
### GetCompanyPrefectureName

`func (o *InvoiceIndexResponseInvoices) GetCompanyPrefectureName() string`

GetCompanyPrefectureName returns the CompanyPrefectureName field if non-nil, zero value otherwise.

### GetCompanyPrefectureNameOk

`func (o *InvoiceIndexResponseInvoices) GetCompanyPrefectureNameOk() (*string, bool)`

GetCompanyPrefectureNameOk returns a tuple with the CompanyPrefectureName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyPrefectureName

`func (o *InvoiceIndexResponseInvoices) SetCompanyPrefectureName(v string)`

SetCompanyPrefectureName sets CompanyPrefectureName field to given value.

### HasCompanyPrefectureName

`func (o *InvoiceIndexResponseInvoices) HasCompanyPrefectureName() bool`

HasCompanyPrefectureName returns a boolean if a field has been set.

### SetCompanyPrefectureNameNil

`func (o *InvoiceIndexResponseInvoices) SetCompanyPrefectureNameNil(b bool)`

 SetCompanyPrefectureNameNil sets the value for CompanyPrefectureName to be an explicit nil

### UnsetCompanyPrefectureName
`func (o *InvoiceIndexResponseInvoices) UnsetCompanyPrefectureName()`

UnsetCompanyPrefectureName ensures that no value is present for CompanyPrefectureName, not even an explicit nil
### GetCompanyAddress1

`func (o *InvoiceIndexResponseInvoices) GetCompanyAddress1() string`

GetCompanyAddress1 returns the CompanyAddress1 field if non-nil, zero value otherwise.

### GetCompanyAddress1Ok

`func (o *InvoiceIndexResponseInvoices) GetCompanyAddress1Ok() (*string, bool)`

GetCompanyAddress1Ok returns a tuple with the CompanyAddress1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyAddress1

`func (o *InvoiceIndexResponseInvoices) SetCompanyAddress1(v string)`

SetCompanyAddress1 sets CompanyAddress1 field to given value.

### HasCompanyAddress1

`func (o *InvoiceIndexResponseInvoices) HasCompanyAddress1() bool`

HasCompanyAddress1 returns a boolean if a field has been set.

### SetCompanyAddress1Nil

`func (o *InvoiceIndexResponseInvoices) SetCompanyAddress1Nil(b bool)`

 SetCompanyAddress1Nil sets the value for CompanyAddress1 to be an explicit nil

### UnsetCompanyAddress1
`func (o *InvoiceIndexResponseInvoices) UnsetCompanyAddress1()`

UnsetCompanyAddress1 ensures that no value is present for CompanyAddress1, not even an explicit nil
### GetCompanyAddress2

`func (o *InvoiceIndexResponseInvoices) GetCompanyAddress2() string`

GetCompanyAddress2 returns the CompanyAddress2 field if non-nil, zero value otherwise.

### GetCompanyAddress2Ok

`func (o *InvoiceIndexResponseInvoices) GetCompanyAddress2Ok() (*string, bool)`

GetCompanyAddress2Ok returns a tuple with the CompanyAddress2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyAddress2

`func (o *InvoiceIndexResponseInvoices) SetCompanyAddress2(v string)`

SetCompanyAddress2 sets CompanyAddress2 field to given value.

### HasCompanyAddress2

`func (o *InvoiceIndexResponseInvoices) HasCompanyAddress2() bool`

HasCompanyAddress2 returns a boolean if a field has been set.

### SetCompanyAddress2Nil

`func (o *InvoiceIndexResponseInvoices) SetCompanyAddress2Nil(b bool)`

 SetCompanyAddress2Nil sets the value for CompanyAddress2 to be an explicit nil

### UnsetCompanyAddress2
`func (o *InvoiceIndexResponseInvoices) UnsetCompanyAddress2()`

UnsetCompanyAddress2 ensures that no value is present for CompanyAddress2, not even an explicit nil
### GetCompanyContactInfo

`func (o *InvoiceIndexResponseInvoices) GetCompanyContactInfo() string`

GetCompanyContactInfo returns the CompanyContactInfo field if non-nil, zero value otherwise.

### GetCompanyContactInfoOk

`func (o *InvoiceIndexResponseInvoices) GetCompanyContactInfoOk() (*string, bool)`

GetCompanyContactInfoOk returns a tuple with the CompanyContactInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyContactInfo

`func (o *InvoiceIndexResponseInvoices) SetCompanyContactInfo(v string)`

SetCompanyContactInfo sets CompanyContactInfo field to given value.

### HasCompanyContactInfo

`func (o *InvoiceIndexResponseInvoices) HasCompanyContactInfo() bool`

HasCompanyContactInfo returns a boolean if a field has been set.

### SetCompanyContactInfoNil

`func (o *InvoiceIndexResponseInvoices) SetCompanyContactInfoNil(b bool)`

 SetCompanyContactInfoNil sets the value for CompanyContactInfo to be an explicit nil

### UnsetCompanyContactInfo
`func (o *InvoiceIndexResponseInvoices) UnsetCompanyContactInfo()`

UnsetCompanyContactInfo ensures that no value is present for CompanyContactInfo, not even an explicit nil
### GetPaymentType

`func (o *InvoiceIndexResponseInvoices) GetPaymentType() string`

GetPaymentType returns the PaymentType field if non-nil, zero value otherwise.

### GetPaymentTypeOk

`func (o *InvoiceIndexResponseInvoices) GetPaymentTypeOk() (*string, bool)`

GetPaymentTypeOk returns a tuple with the PaymentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentType

`func (o *InvoiceIndexResponseInvoices) SetPaymentType(v string)`

SetPaymentType sets PaymentType field to given value.


### GetPaymentBankInfo

`func (o *InvoiceIndexResponseInvoices) GetPaymentBankInfo() string`

GetPaymentBankInfo returns the PaymentBankInfo field if non-nil, zero value otherwise.

### GetPaymentBankInfoOk

`func (o *InvoiceIndexResponseInvoices) GetPaymentBankInfoOk() (*string, bool)`

GetPaymentBankInfoOk returns a tuple with the PaymentBankInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentBankInfo

`func (o *InvoiceIndexResponseInvoices) SetPaymentBankInfo(v string)`

SetPaymentBankInfo sets PaymentBankInfo field to given value.

### HasPaymentBankInfo

`func (o *InvoiceIndexResponseInvoices) HasPaymentBankInfo() bool`

HasPaymentBankInfo returns a boolean if a field has been set.

### SetPaymentBankInfoNil

`func (o *InvoiceIndexResponseInvoices) SetPaymentBankInfoNil(b bool)`

 SetPaymentBankInfoNil sets the value for PaymentBankInfo to be an explicit nil

### UnsetPaymentBankInfo
`func (o *InvoiceIndexResponseInvoices) UnsetPaymentBankInfo()`

UnsetPaymentBankInfo ensures that no value is present for PaymentBankInfo, not even an explicit nil
### GetMessage

`func (o *InvoiceIndexResponseInvoices) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *InvoiceIndexResponseInvoices) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *InvoiceIndexResponseInvoices) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *InvoiceIndexResponseInvoices) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *InvoiceIndexResponseInvoices) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *InvoiceIndexResponseInvoices) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetNotes

`func (o *InvoiceIndexResponseInvoices) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *InvoiceIndexResponseInvoices) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *InvoiceIndexResponseInvoices) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *InvoiceIndexResponseInvoices) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *InvoiceIndexResponseInvoices) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *InvoiceIndexResponseInvoices) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetInvoiceLayout

`func (o *InvoiceIndexResponseInvoices) GetInvoiceLayout() string`

GetInvoiceLayout returns the InvoiceLayout field if non-nil, zero value otherwise.

### GetInvoiceLayoutOk

`func (o *InvoiceIndexResponseInvoices) GetInvoiceLayoutOk() (*string, bool)`

GetInvoiceLayoutOk returns a tuple with the InvoiceLayout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceLayout

`func (o *InvoiceIndexResponseInvoices) SetInvoiceLayout(v string)`

SetInvoiceLayout sets InvoiceLayout field to given value.


### GetTaxEntryMethod

`func (o *InvoiceIndexResponseInvoices) GetTaxEntryMethod() string`

GetTaxEntryMethod returns the TaxEntryMethod field if non-nil, zero value otherwise.

### GetTaxEntryMethodOk

`func (o *InvoiceIndexResponseInvoices) GetTaxEntryMethodOk() (*string, bool)`

GetTaxEntryMethodOk returns a tuple with the TaxEntryMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxEntryMethod

`func (o *InvoiceIndexResponseInvoices) SetTaxEntryMethod(v string)`

SetTaxEntryMethod sets TaxEntryMethod field to given value.


### GetDealId

`func (o *InvoiceIndexResponseInvoices) GetDealId() int32`

GetDealId returns the DealId field if non-nil, zero value otherwise.

### GetDealIdOk

`func (o *InvoiceIndexResponseInvoices) GetDealIdOk() (*int32, bool)`

GetDealIdOk returns a tuple with the DealId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDealId

`func (o *InvoiceIndexResponseInvoices) SetDealId(v int32)`

SetDealId sets DealId field to given value.

### HasDealId

`func (o *InvoiceIndexResponseInvoices) HasDealId() bool`

HasDealId returns a boolean if a field has been set.

### SetDealIdNil

`func (o *InvoiceIndexResponseInvoices) SetDealIdNil(b bool)`

 SetDealIdNil sets the value for DealId to be an explicit nil

### UnsetDealId
`func (o *InvoiceIndexResponseInvoices) UnsetDealId()`

UnsetDealId ensures that no value is present for DealId, not even an explicit nil
### GetInvoiceContents

`func (o *InvoiceIndexResponseInvoices) GetInvoiceContents() []InvoiceIndexResponseInvoiceContents`

GetInvoiceContents returns the InvoiceContents field if non-nil, zero value otherwise.

### GetInvoiceContentsOk

`func (o *InvoiceIndexResponseInvoices) GetInvoiceContentsOk() (*[]InvoiceIndexResponseInvoiceContents, bool)`

GetInvoiceContentsOk returns a tuple with the InvoiceContents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceContents

`func (o *InvoiceIndexResponseInvoices) SetInvoiceContents(v []InvoiceIndexResponseInvoiceContents)`

SetInvoiceContents sets InvoiceContents field to given value.

### HasInvoiceContents

`func (o *InvoiceIndexResponseInvoices) HasInvoiceContents() bool`

HasInvoiceContents returns a boolean if a field has been set.

### GetTotalAmountPerVatRate

`func (o *InvoiceIndexResponseInvoices) GetTotalAmountPerVatRate() InvoiceIndexResponseTotalAmountPerVatRate`

GetTotalAmountPerVatRate returns the TotalAmountPerVatRate field if non-nil, zero value otherwise.

### GetTotalAmountPerVatRateOk

`func (o *InvoiceIndexResponseInvoices) GetTotalAmountPerVatRateOk() (*InvoiceIndexResponseTotalAmountPerVatRate, bool)`

GetTotalAmountPerVatRateOk returns a tuple with the TotalAmountPerVatRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmountPerVatRate

`func (o *InvoiceIndexResponseInvoices) SetTotalAmountPerVatRate(v InvoiceIndexResponseTotalAmountPerVatRate)`

SetTotalAmountPerVatRate sets TotalAmountPerVatRate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


