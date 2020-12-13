# QuotationIndexResponseQuotations

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | 見積書ID | 
**CompanyId** | **int32** | 事業所ID | 
**IssueDate** | **string** | 見積日 (yyyy-mm-dd) | 
**PartnerId** | **NullableInt32** | 取引先ID | 
**PartnerCode** | Pointer to **NullableString** | 取引先コード | [optional] 
**QuotationNumber** | **string** | 見積書番号 | 
**Title** | Pointer to **NullableString** | タイトル | [optional] 
**TotalAmount** | **int32** | 合計金額 | 
**TotalVat** | Pointer to **int32** | 消費税 | [optional] 
**SubTotal** | Pointer to **int32** | 小計 | [optional] 
**Description** | Pointer to **NullableString** | 概要 | [optional] 
**QuotationStatus** | **string** | 見積書ステータス  (unsubmitted: 送付待ち, submitted: 送付済み, all: 全て) | 
**WebPublishedAt** | Pointer to **NullableString** | Web共有日時(最新) | [optional] 
**WebDownloadedAt** | Pointer to **NullableString** | Web共有ダウンロード日時(最新) | [optional] 
**WebConfirmedAt** | Pointer to **NullableString** | Web共有取引先確認日時(最新) | [optional] 
**MailSentAt** | Pointer to **NullableString** | メール送信日時(最新) | [optional] 
**PartnerName** | Pointer to **NullableString** | 取引先名 | [optional] 
**PartnerDisplayName** | Pointer to **NullableString** | 見積書に表示する取引先名 | [optional] 
**PartnerTitle** | **NullableString** | 敬称（御中、様、(空白)の3つから選択） | 
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
**Message** | Pointer to **NullableString** | メッセージ | [optional] 
**Notes** | Pointer to **NullableString** | 備考 | [optional] 
**QuotationLayout** | **string** | 見積書レイアウト * &#x60;default_classic&#x60; - レイアウト１/クラシック (デフォルト)  * &#x60;standard_classic&#x60; - レイアウト２/クラシック  * &#x60;envelope_classic&#x60; - 封筒１/クラシック  * &#x60;default_modern&#x60; - レイアウト１/モダン  * &#x60;standard_modern&#x60; - レイアウト２/モダン  * &#x60;envelope_modern&#x60; - 封筒/モダン | 
**TaxEntryMethod** | **string** | 見積書の消費税計算方法(inclusive: 内税, exclusive: 外税) | 
**QuotationContents** | Pointer to [**[]QuotationIndexResponseQuotationContents**](QuotationIndexResponseQuotationContents.md) | 見積内容 | [optional] 
**TotalAmountPerVatRate** | [**InvoiceIndexResponseTotalAmountPerVatRate**](invoiceIndexResponse_total_amount_per_vat_rate.md) |  | 

## Methods

### NewQuotationIndexResponseQuotations

`func NewQuotationIndexResponseQuotations(id int32, companyId int32, issueDate string, partnerId NullableInt32, quotationNumber string, totalAmount int32, quotationStatus string, partnerTitle NullableString, companyName string, quotationLayout string, taxEntryMethod string, totalAmountPerVatRate InvoiceIndexResponseTotalAmountPerVatRate, ) *QuotationIndexResponseQuotations`

NewQuotationIndexResponseQuotations instantiates a new QuotationIndexResponseQuotations object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuotationIndexResponseQuotationsWithDefaults

`func NewQuotationIndexResponseQuotationsWithDefaults() *QuotationIndexResponseQuotations`

NewQuotationIndexResponseQuotationsWithDefaults instantiates a new QuotationIndexResponseQuotations object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QuotationIndexResponseQuotations) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QuotationIndexResponseQuotations) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QuotationIndexResponseQuotations) SetId(v int32)`

SetId sets Id field to given value.


### GetCompanyId

`func (o *QuotationIndexResponseQuotations) GetCompanyId() int32`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *QuotationIndexResponseQuotations) GetCompanyIdOk() (*int32, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *QuotationIndexResponseQuotations) SetCompanyId(v int32)`

SetCompanyId sets CompanyId field to given value.


### GetIssueDate

`func (o *QuotationIndexResponseQuotations) GetIssueDate() string`

GetIssueDate returns the IssueDate field if non-nil, zero value otherwise.

### GetIssueDateOk

`func (o *QuotationIndexResponseQuotations) GetIssueDateOk() (*string, bool)`

GetIssueDateOk returns a tuple with the IssueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueDate

`func (o *QuotationIndexResponseQuotations) SetIssueDate(v string)`

SetIssueDate sets IssueDate field to given value.


### GetPartnerId

`func (o *QuotationIndexResponseQuotations) GetPartnerId() int32`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *QuotationIndexResponseQuotations) GetPartnerIdOk() (*int32, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *QuotationIndexResponseQuotations) SetPartnerId(v int32)`

SetPartnerId sets PartnerId field to given value.


### SetPartnerIdNil

`func (o *QuotationIndexResponseQuotations) SetPartnerIdNil(b bool)`

 SetPartnerIdNil sets the value for PartnerId to be an explicit nil

### UnsetPartnerId
`func (o *QuotationIndexResponseQuotations) UnsetPartnerId()`

UnsetPartnerId ensures that no value is present for PartnerId, not even an explicit nil
### GetPartnerCode

`func (o *QuotationIndexResponseQuotations) GetPartnerCode() string`

GetPartnerCode returns the PartnerCode field if non-nil, zero value otherwise.

### GetPartnerCodeOk

`func (o *QuotationIndexResponseQuotations) GetPartnerCodeOk() (*string, bool)`

GetPartnerCodeOk returns a tuple with the PartnerCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerCode

`func (o *QuotationIndexResponseQuotations) SetPartnerCode(v string)`

SetPartnerCode sets PartnerCode field to given value.

### HasPartnerCode

`func (o *QuotationIndexResponseQuotations) HasPartnerCode() bool`

HasPartnerCode returns a boolean if a field has been set.

### SetPartnerCodeNil

`func (o *QuotationIndexResponseQuotations) SetPartnerCodeNil(b bool)`

 SetPartnerCodeNil sets the value for PartnerCode to be an explicit nil

### UnsetPartnerCode
`func (o *QuotationIndexResponseQuotations) UnsetPartnerCode()`

UnsetPartnerCode ensures that no value is present for PartnerCode, not even an explicit nil
### GetQuotationNumber

`func (o *QuotationIndexResponseQuotations) GetQuotationNumber() string`

GetQuotationNumber returns the QuotationNumber field if non-nil, zero value otherwise.

### GetQuotationNumberOk

`func (o *QuotationIndexResponseQuotations) GetQuotationNumberOk() (*string, bool)`

GetQuotationNumberOk returns a tuple with the QuotationNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotationNumber

`func (o *QuotationIndexResponseQuotations) SetQuotationNumber(v string)`

SetQuotationNumber sets QuotationNumber field to given value.


### GetTitle

`func (o *QuotationIndexResponseQuotations) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *QuotationIndexResponseQuotations) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *QuotationIndexResponseQuotations) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *QuotationIndexResponseQuotations) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *QuotationIndexResponseQuotations) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *QuotationIndexResponseQuotations) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetTotalAmount

`func (o *QuotationIndexResponseQuotations) GetTotalAmount() int32`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *QuotationIndexResponseQuotations) GetTotalAmountOk() (*int32, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *QuotationIndexResponseQuotations) SetTotalAmount(v int32)`

SetTotalAmount sets TotalAmount field to given value.


### GetTotalVat

`func (o *QuotationIndexResponseQuotations) GetTotalVat() int32`

GetTotalVat returns the TotalVat field if non-nil, zero value otherwise.

### GetTotalVatOk

`func (o *QuotationIndexResponseQuotations) GetTotalVatOk() (*int32, bool)`

GetTotalVatOk returns a tuple with the TotalVat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalVat

`func (o *QuotationIndexResponseQuotations) SetTotalVat(v int32)`

SetTotalVat sets TotalVat field to given value.

### HasTotalVat

`func (o *QuotationIndexResponseQuotations) HasTotalVat() bool`

HasTotalVat returns a boolean if a field has been set.

### GetSubTotal

`func (o *QuotationIndexResponseQuotations) GetSubTotal() int32`

GetSubTotal returns the SubTotal field if non-nil, zero value otherwise.

### GetSubTotalOk

`func (o *QuotationIndexResponseQuotations) GetSubTotalOk() (*int32, bool)`

GetSubTotalOk returns a tuple with the SubTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubTotal

`func (o *QuotationIndexResponseQuotations) SetSubTotal(v int32)`

SetSubTotal sets SubTotal field to given value.

### HasSubTotal

`func (o *QuotationIndexResponseQuotations) HasSubTotal() bool`

HasSubTotal returns a boolean if a field has been set.

### GetDescription

`func (o *QuotationIndexResponseQuotations) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *QuotationIndexResponseQuotations) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *QuotationIndexResponseQuotations) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *QuotationIndexResponseQuotations) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *QuotationIndexResponseQuotations) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *QuotationIndexResponseQuotations) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetQuotationStatus

`func (o *QuotationIndexResponseQuotations) GetQuotationStatus() string`

GetQuotationStatus returns the QuotationStatus field if non-nil, zero value otherwise.

### GetQuotationStatusOk

`func (o *QuotationIndexResponseQuotations) GetQuotationStatusOk() (*string, bool)`

GetQuotationStatusOk returns a tuple with the QuotationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotationStatus

`func (o *QuotationIndexResponseQuotations) SetQuotationStatus(v string)`

SetQuotationStatus sets QuotationStatus field to given value.


### GetWebPublishedAt

`func (o *QuotationIndexResponseQuotations) GetWebPublishedAt() string`

GetWebPublishedAt returns the WebPublishedAt field if non-nil, zero value otherwise.

### GetWebPublishedAtOk

`func (o *QuotationIndexResponseQuotations) GetWebPublishedAtOk() (*string, bool)`

GetWebPublishedAtOk returns a tuple with the WebPublishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebPublishedAt

`func (o *QuotationIndexResponseQuotations) SetWebPublishedAt(v string)`

SetWebPublishedAt sets WebPublishedAt field to given value.

### HasWebPublishedAt

`func (o *QuotationIndexResponseQuotations) HasWebPublishedAt() bool`

HasWebPublishedAt returns a boolean if a field has been set.

### SetWebPublishedAtNil

`func (o *QuotationIndexResponseQuotations) SetWebPublishedAtNil(b bool)`

 SetWebPublishedAtNil sets the value for WebPublishedAt to be an explicit nil

### UnsetWebPublishedAt
`func (o *QuotationIndexResponseQuotations) UnsetWebPublishedAt()`

UnsetWebPublishedAt ensures that no value is present for WebPublishedAt, not even an explicit nil
### GetWebDownloadedAt

`func (o *QuotationIndexResponseQuotations) GetWebDownloadedAt() string`

GetWebDownloadedAt returns the WebDownloadedAt field if non-nil, zero value otherwise.

### GetWebDownloadedAtOk

`func (o *QuotationIndexResponseQuotations) GetWebDownloadedAtOk() (*string, bool)`

GetWebDownloadedAtOk returns a tuple with the WebDownloadedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebDownloadedAt

`func (o *QuotationIndexResponseQuotations) SetWebDownloadedAt(v string)`

SetWebDownloadedAt sets WebDownloadedAt field to given value.

### HasWebDownloadedAt

`func (o *QuotationIndexResponseQuotations) HasWebDownloadedAt() bool`

HasWebDownloadedAt returns a boolean if a field has been set.

### SetWebDownloadedAtNil

`func (o *QuotationIndexResponseQuotations) SetWebDownloadedAtNil(b bool)`

 SetWebDownloadedAtNil sets the value for WebDownloadedAt to be an explicit nil

### UnsetWebDownloadedAt
`func (o *QuotationIndexResponseQuotations) UnsetWebDownloadedAt()`

UnsetWebDownloadedAt ensures that no value is present for WebDownloadedAt, not even an explicit nil
### GetWebConfirmedAt

`func (o *QuotationIndexResponseQuotations) GetWebConfirmedAt() string`

GetWebConfirmedAt returns the WebConfirmedAt field if non-nil, zero value otherwise.

### GetWebConfirmedAtOk

`func (o *QuotationIndexResponseQuotations) GetWebConfirmedAtOk() (*string, bool)`

GetWebConfirmedAtOk returns a tuple with the WebConfirmedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebConfirmedAt

`func (o *QuotationIndexResponseQuotations) SetWebConfirmedAt(v string)`

SetWebConfirmedAt sets WebConfirmedAt field to given value.

### HasWebConfirmedAt

`func (o *QuotationIndexResponseQuotations) HasWebConfirmedAt() bool`

HasWebConfirmedAt returns a boolean if a field has been set.

### SetWebConfirmedAtNil

`func (o *QuotationIndexResponseQuotations) SetWebConfirmedAtNil(b bool)`

 SetWebConfirmedAtNil sets the value for WebConfirmedAt to be an explicit nil

### UnsetWebConfirmedAt
`func (o *QuotationIndexResponseQuotations) UnsetWebConfirmedAt()`

UnsetWebConfirmedAt ensures that no value is present for WebConfirmedAt, not even an explicit nil
### GetMailSentAt

`func (o *QuotationIndexResponseQuotations) GetMailSentAt() string`

GetMailSentAt returns the MailSentAt field if non-nil, zero value otherwise.

### GetMailSentAtOk

`func (o *QuotationIndexResponseQuotations) GetMailSentAtOk() (*string, bool)`

GetMailSentAtOk returns a tuple with the MailSentAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailSentAt

`func (o *QuotationIndexResponseQuotations) SetMailSentAt(v string)`

SetMailSentAt sets MailSentAt field to given value.

### HasMailSentAt

`func (o *QuotationIndexResponseQuotations) HasMailSentAt() bool`

HasMailSentAt returns a boolean if a field has been set.

### SetMailSentAtNil

`func (o *QuotationIndexResponseQuotations) SetMailSentAtNil(b bool)`

 SetMailSentAtNil sets the value for MailSentAt to be an explicit nil

### UnsetMailSentAt
`func (o *QuotationIndexResponseQuotations) UnsetMailSentAt()`

UnsetMailSentAt ensures that no value is present for MailSentAt, not even an explicit nil
### GetPartnerName

`func (o *QuotationIndexResponseQuotations) GetPartnerName() string`

GetPartnerName returns the PartnerName field if non-nil, zero value otherwise.

### GetPartnerNameOk

`func (o *QuotationIndexResponseQuotations) GetPartnerNameOk() (*string, bool)`

GetPartnerNameOk returns a tuple with the PartnerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerName

`func (o *QuotationIndexResponseQuotations) SetPartnerName(v string)`

SetPartnerName sets PartnerName field to given value.

### HasPartnerName

`func (o *QuotationIndexResponseQuotations) HasPartnerName() bool`

HasPartnerName returns a boolean if a field has been set.

### SetPartnerNameNil

`func (o *QuotationIndexResponseQuotations) SetPartnerNameNil(b bool)`

 SetPartnerNameNil sets the value for PartnerName to be an explicit nil

### UnsetPartnerName
`func (o *QuotationIndexResponseQuotations) UnsetPartnerName()`

UnsetPartnerName ensures that no value is present for PartnerName, not even an explicit nil
### GetPartnerDisplayName

`func (o *QuotationIndexResponseQuotations) GetPartnerDisplayName() string`

GetPartnerDisplayName returns the PartnerDisplayName field if non-nil, zero value otherwise.

### GetPartnerDisplayNameOk

`func (o *QuotationIndexResponseQuotations) GetPartnerDisplayNameOk() (*string, bool)`

GetPartnerDisplayNameOk returns a tuple with the PartnerDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerDisplayName

`func (o *QuotationIndexResponseQuotations) SetPartnerDisplayName(v string)`

SetPartnerDisplayName sets PartnerDisplayName field to given value.

### HasPartnerDisplayName

`func (o *QuotationIndexResponseQuotations) HasPartnerDisplayName() bool`

HasPartnerDisplayName returns a boolean if a field has been set.

### SetPartnerDisplayNameNil

`func (o *QuotationIndexResponseQuotations) SetPartnerDisplayNameNil(b bool)`

 SetPartnerDisplayNameNil sets the value for PartnerDisplayName to be an explicit nil

### UnsetPartnerDisplayName
`func (o *QuotationIndexResponseQuotations) UnsetPartnerDisplayName()`

UnsetPartnerDisplayName ensures that no value is present for PartnerDisplayName, not even an explicit nil
### GetPartnerTitle

`func (o *QuotationIndexResponseQuotations) GetPartnerTitle() string`

GetPartnerTitle returns the PartnerTitle field if non-nil, zero value otherwise.

### GetPartnerTitleOk

`func (o *QuotationIndexResponseQuotations) GetPartnerTitleOk() (*string, bool)`

GetPartnerTitleOk returns a tuple with the PartnerTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerTitle

`func (o *QuotationIndexResponseQuotations) SetPartnerTitle(v string)`

SetPartnerTitle sets PartnerTitle field to given value.


### SetPartnerTitleNil

`func (o *QuotationIndexResponseQuotations) SetPartnerTitleNil(b bool)`

 SetPartnerTitleNil sets the value for PartnerTitle to be an explicit nil

### UnsetPartnerTitle
`func (o *QuotationIndexResponseQuotations) UnsetPartnerTitle()`

UnsetPartnerTitle ensures that no value is present for PartnerTitle, not even an explicit nil
### GetPartnerZipcode

`func (o *QuotationIndexResponseQuotations) GetPartnerZipcode() string`

GetPartnerZipcode returns the PartnerZipcode field if non-nil, zero value otherwise.

### GetPartnerZipcodeOk

`func (o *QuotationIndexResponseQuotations) GetPartnerZipcodeOk() (*string, bool)`

GetPartnerZipcodeOk returns a tuple with the PartnerZipcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerZipcode

`func (o *QuotationIndexResponseQuotations) SetPartnerZipcode(v string)`

SetPartnerZipcode sets PartnerZipcode field to given value.

### HasPartnerZipcode

`func (o *QuotationIndexResponseQuotations) HasPartnerZipcode() bool`

HasPartnerZipcode returns a boolean if a field has been set.

### SetPartnerZipcodeNil

`func (o *QuotationIndexResponseQuotations) SetPartnerZipcodeNil(b bool)`

 SetPartnerZipcodeNil sets the value for PartnerZipcode to be an explicit nil

### UnsetPartnerZipcode
`func (o *QuotationIndexResponseQuotations) UnsetPartnerZipcode()`

UnsetPartnerZipcode ensures that no value is present for PartnerZipcode, not even an explicit nil
### GetPartnerPrefectureCode

`func (o *QuotationIndexResponseQuotations) GetPartnerPrefectureCode() int32`

GetPartnerPrefectureCode returns the PartnerPrefectureCode field if non-nil, zero value otherwise.

### GetPartnerPrefectureCodeOk

`func (o *QuotationIndexResponseQuotations) GetPartnerPrefectureCodeOk() (*int32, bool)`

GetPartnerPrefectureCodeOk returns a tuple with the PartnerPrefectureCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerPrefectureCode

`func (o *QuotationIndexResponseQuotations) SetPartnerPrefectureCode(v int32)`

SetPartnerPrefectureCode sets PartnerPrefectureCode field to given value.

### HasPartnerPrefectureCode

`func (o *QuotationIndexResponseQuotations) HasPartnerPrefectureCode() bool`

HasPartnerPrefectureCode returns a boolean if a field has been set.

### SetPartnerPrefectureCodeNil

`func (o *QuotationIndexResponseQuotations) SetPartnerPrefectureCodeNil(b bool)`

 SetPartnerPrefectureCodeNil sets the value for PartnerPrefectureCode to be an explicit nil

### UnsetPartnerPrefectureCode
`func (o *QuotationIndexResponseQuotations) UnsetPartnerPrefectureCode()`

UnsetPartnerPrefectureCode ensures that no value is present for PartnerPrefectureCode, not even an explicit nil
### GetPartnerPrefectureName

`func (o *QuotationIndexResponseQuotations) GetPartnerPrefectureName() string`

GetPartnerPrefectureName returns the PartnerPrefectureName field if non-nil, zero value otherwise.

### GetPartnerPrefectureNameOk

`func (o *QuotationIndexResponseQuotations) GetPartnerPrefectureNameOk() (*string, bool)`

GetPartnerPrefectureNameOk returns a tuple with the PartnerPrefectureName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerPrefectureName

`func (o *QuotationIndexResponseQuotations) SetPartnerPrefectureName(v string)`

SetPartnerPrefectureName sets PartnerPrefectureName field to given value.

### HasPartnerPrefectureName

`func (o *QuotationIndexResponseQuotations) HasPartnerPrefectureName() bool`

HasPartnerPrefectureName returns a boolean if a field has been set.

### SetPartnerPrefectureNameNil

`func (o *QuotationIndexResponseQuotations) SetPartnerPrefectureNameNil(b bool)`

 SetPartnerPrefectureNameNil sets the value for PartnerPrefectureName to be an explicit nil

### UnsetPartnerPrefectureName
`func (o *QuotationIndexResponseQuotations) UnsetPartnerPrefectureName()`

UnsetPartnerPrefectureName ensures that no value is present for PartnerPrefectureName, not even an explicit nil
### GetPartnerAddress1

`func (o *QuotationIndexResponseQuotations) GetPartnerAddress1() string`

GetPartnerAddress1 returns the PartnerAddress1 field if non-nil, zero value otherwise.

### GetPartnerAddress1Ok

`func (o *QuotationIndexResponseQuotations) GetPartnerAddress1Ok() (*string, bool)`

GetPartnerAddress1Ok returns a tuple with the PartnerAddress1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerAddress1

`func (o *QuotationIndexResponseQuotations) SetPartnerAddress1(v string)`

SetPartnerAddress1 sets PartnerAddress1 field to given value.

### HasPartnerAddress1

`func (o *QuotationIndexResponseQuotations) HasPartnerAddress1() bool`

HasPartnerAddress1 returns a boolean if a field has been set.

### SetPartnerAddress1Nil

`func (o *QuotationIndexResponseQuotations) SetPartnerAddress1Nil(b bool)`

 SetPartnerAddress1Nil sets the value for PartnerAddress1 to be an explicit nil

### UnsetPartnerAddress1
`func (o *QuotationIndexResponseQuotations) UnsetPartnerAddress1()`

UnsetPartnerAddress1 ensures that no value is present for PartnerAddress1, not even an explicit nil
### GetPartnerAddress2

`func (o *QuotationIndexResponseQuotations) GetPartnerAddress2() string`

GetPartnerAddress2 returns the PartnerAddress2 field if non-nil, zero value otherwise.

### GetPartnerAddress2Ok

`func (o *QuotationIndexResponseQuotations) GetPartnerAddress2Ok() (*string, bool)`

GetPartnerAddress2Ok returns a tuple with the PartnerAddress2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerAddress2

`func (o *QuotationIndexResponseQuotations) SetPartnerAddress2(v string)`

SetPartnerAddress2 sets PartnerAddress2 field to given value.

### HasPartnerAddress2

`func (o *QuotationIndexResponseQuotations) HasPartnerAddress2() bool`

HasPartnerAddress2 returns a boolean if a field has been set.

### SetPartnerAddress2Nil

`func (o *QuotationIndexResponseQuotations) SetPartnerAddress2Nil(b bool)`

 SetPartnerAddress2Nil sets the value for PartnerAddress2 to be an explicit nil

### UnsetPartnerAddress2
`func (o *QuotationIndexResponseQuotations) UnsetPartnerAddress2()`

UnsetPartnerAddress2 ensures that no value is present for PartnerAddress2, not even an explicit nil
### GetPartnerContactInfo

`func (o *QuotationIndexResponseQuotations) GetPartnerContactInfo() string`

GetPartnerContactInfo returns the PartnerContactInfo field if non-nil, zero value otherwise.

### GetPartnerContactInfoOk

`func (o *QuotationIndexResponseQuotations) GetPartnerContactInfoOk() (*string, bool)`

GetPartnerContactInfoOk returns a tuple with the PartnerContactInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerContactInfo

`func (o *QuotationIndexResponseQuotations) SetPartnerContactInfo(v string)`

SetPartnerContactInfo sets PartnerContactInfo field to given value.

### HasPartnerContactInfo

`func (o *QuotationIndexResponseQuotations) HasPartnerContactInfo() bool`

HasPartnerContactInfo returns a boolean if a field has been set.

### SetPartnerContactInfoNil

`func (o *QuotationIndexResponseQuotations) SetPartnerContactInfoNil(b bool)`

 SetPartnerContactInfoNil sets the value for PartnerContactInfo to be an explicit nil

### UnsetPartnerContactInfo
`func (o *QuotationIndexResponseQuotations) UnsetPartnerContactInfo()`

UnsetPartnerContactInfo ensures that no value is present for PartnerContactInfo, not even an explicit nil
### GetCompanyName

`func (o *QuotationIndexResponseQuotations) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *QuotationIndexResponseQuotations) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *QuotationIndexResponseQuotations) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.


### GetCompanyZipcode

`func (o *QuotationIndexResponseQuotations) GetCompanyZipcode() string`

GetCompanyZipcode returns the CompanyZipcode field if non-nil, zero value otherwise.

### GetCompanyZipcodeOk

`func (o *QuotationIndexResponseQuotations) GetCompanyZipcodeOk() (*string, bool)`

GetCompanyZipcodeOk returns a tuple with the CompanyZipcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyZipcode

`func (o *QuotationIndexResponseQuotations) SetCompanyZipcode(v string)`

SetCompanyZipcode sets CompanyZipcode field to given value.

### HasCompanyZipcode

`func (o *QuotationIndexResponseQuotations) HasCompanyZipcode() bool`

HasCompanyZipcode returns a boolean if a field has been set.

### SetCompanyZipcodeNil

`func (o *QuotationIndexResponseQuotations) SetCompanyZipcodeNil(b bool)`

 SetCompanyZipcodeNil sets the value for CompanyZipcode to be an explicit nil

### UnsetCompanyZipcode
`func (o *QuotationIndexResponseQuotations) UnsetCompanyZipcode()`

UnsetCompanyZipcode ensures that no value is present for CompanyZipcode, not even an explicit nil
### GetCompanyPrefectureCode

`func (o *QuotationIndexResponseQuotations) GetCompanyPrefectureCode() int32`

GetCompanyPrefectureCode returns the CompanyPrefectureCode field if non-nil, zero value otherwise.

### GetCompanyPrefectureCodeOk

`func (o *QuotationIndexResponseQuotations) GetCompanyPrefectureCodeOk() (*int32, bool)`

GetCompanyPrefectureCodeOk returns a tuple with the CompanyPrefectureCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyPrefectureCode

`func (o *QuotationIndexResponseQuotations) SetCompanyPrefectureCode(v int32)`

SetCompanyPrefectureCode sets CompanyPrefectureCode field to given value.

### HasCompanyPrefectureCode

`func (o *QuotationIndexResponseQuotations) HasCompanyPrefectureCode() bool`

HasCompanyPrefectureCode returns a boolean if a field has been set.

### SetCompanyPrefectureCodeNil

`func (o *QuotationIndexResponseQuotations) SetCompanyPrefectureCodeNil(b bool)`

 SetCompanyPrefectureCodeNil sets the value for CompanyPrefectureCode to be an explicit nil

### UnsetCompanyPrefectureCode
`func (o *QuotationIndexResponseQuotations) UnsetCompanyPrefectureCode()`

UnsetCompanyPrefectureCode ensures that no value is present for CompanyPrefectureCode, not even an explicit nil
### GetCompanyPrefectureName

`func (o *QuotationIndexResponseQuotations) GetCompanyPrefectureName() string`

GetCompanyPrefectureName returns the CompanyPrefectureName field if non-nil, zero value otherwise.

### GetCompanyPrefectureNameOk

`func (o *QuotationIndexResponseQuotations) GetCompanyPrefectureNameOk() (*string, bool)`

GetCompanyPrefectureNameOk returns a tuple with the CompanyPrefectureName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyPrefectureName

`func (o *QuotationIndexResponseQuotations) SetCompanyPrefectureName(v string)`

SetCompanyPrefectureName sets CompanyPrefectureName field to given value.

### HasCompanyPrefectureName

`func (o *QuotationIndexResponseQuotations) HasCompanyPrefectureName() bool`

HasCompanyPrefectureName returns a boolean if a field has been set.

### SetCompanyPrefectureNameNil

`func (o *QuotationIndexResponseQuotations) SetCompanyPrefectureNameNil(b bool)`

 SetCompanyPrefectureNameNil sets the value for CompanyPrefectureName to be an explicit nil

### UnsetCompanyPrefectureName
`func (o *QuotationIndexResponseQuotations) UnsetCompanyPrefectureName()`

UnsetCompanyPrefectureName ensures that no value is present for CompanyPrefectureName, not even an explicit nil
### GetCompanyAddress1

`func (o *QuotationIndexResponseQuotations) GetCompanyAddress1() string`

GetCompanyAddress1 returns the CompanyAddress1 field if non-nil, zero value otherwise.

### GetCompanyAddress1Ok

`func (o *QuotationIndexResponseQuotations) GetCompanyAddress1Ok() (*string, bool)`

GetCompanyAddress1Ok returns a tuple with the CompanyAddress1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyAddress1

`func (o *QuotationIndexResponseQuotations) SetCompanyAddress1(v string)`

SetCompanyAddress1 sets CompanyAddress1 field to given value.

### HasCompanyAddress1

`func (o *QuotationIndexResponseQuotations) HasCompanyAddress1() bool`

HasCompanyAddress1 returns a boolean if a field has been set.

### SetCompanyAddress1Nil

`func (o *QuotationIndexResponseQuotations) SetCompanyAddress1Nil(b bool)`

 SetCompanyAddress1Nil sets the value for CompanyAddress1 to be an explicit nil

### UnsetCompanyAddress1
`func (o *QuotationIndexResponseQuotations) UnsetCompanyAddress1()`

UnsetCompanyAddress1 ensures that no value is present for CompanyAddress1, not even an explicit nil
### GetCompanyAddress2

`func (o *QuotationIndexResponseQuotations) GetCompanyAddress2() string`

GetCompanyAddress2 returns the CompanyAddress2 field if non-nil, zero value otherwise.

### GetCompanyAddress2Ok

`func (o *QuotationIndexResponseQuotations) GetCompanyAddress2Ok() (*string, bool)`

GetCompanyAddress2Ok returns a tuple with the CompanyAddress2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyAddress2

`func (o *QuotationIndexResponseQuotations) SetCompanyAddress2(v string)`

SetCompanyAddress2 sets CompanyAddress2 field to given value.

### HasCompanyAddress2

`func (o *QuotationIndexResponseQuotations) HasCompanyAddress2() bool`

HasCompanyAddress2 returns a boolean if a field has been set.

### SetCompanyAddress2Nil

`func (o *QuotationIndexResponseQuotations) SetCompanyAddress2Nil(b bool)`

 SetCompanyAddress2Nil sets the value for CompanyAddress2 to be an explicit nil

### UnsetCompanyAddress2
`func (o *QuotationIndexResponseQuotations) UnsetCompanyAddress2()`

UnsetCompanyAddress2 ensures that no value is present for CompanyAddress2, not even an explicit nil
### GetCompanyContactInfo

`func (o *QuotationIndexResponseQuotations) GetCompanyContactInfo() string`

GetCompanyContactInfo returns the CompanyContactInfo field if non-nil, zero value otherwise.

### GetCompanyContactInfoOk

`func (o *QuotationIndexResponseQuotations) GetCompanyContactInfoOk() (*string, bool)`

GetCompanyContactInfoOk returns a tuple with the CompanyContactInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyContactInfo

`func (o *QuotationIndexResponseQuotations) SetCompanyContactInfo(v string)`

SetCompanyContactInfo sets CompanyContactInfo field to given value.

### HasCompanyContactInfo

`func (o *QuotationIndexResponseQuotations) HasCompanyContactInfo() bool`

HasCompanyContactInfo returns a boolean if a field has been set.

### SetCompanyContactInfoNil

`func (o *QuotationIndexResponseQuotations) SetCompanyContactInfoNil(b bool)`

 SetCompanyContactInfoNil sets the value for CompanyContactInfo to be an explicit nil

### UnsetCompanyContactInfo
`func (o *QuotationIndexResponseQuotations) UnsetCompanyContactInfo()`

UnsetCompanyContactInfo ensures that no value is present for CompanyContactInfo, not even an explicit nil
### GetMessage

`func (o *QuotationIndexResponseQuotations) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *QuotationIndexResponseQuotations) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *QuotationIndexResponseQuotations) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *QuotationIndexResponseQuotations) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *QuotationIndexResponseQuotations) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *QuotationIndexResponseQuotations) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetNotes

`func (o *QuotationIndexResponseQuotations) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *QuotationIndexResponseQuotations) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *QuotationIndexResponseQuotations) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *QuotationIndexResponseQuotations) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *QuotationIndexResponseQuotations) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *QuotationIndexResponseQuotations) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetQuotationLayout

`func (o *QuotationIndexResponseQuotations) GetQuotationLayout() string`

GetQuotationLayout returns the QuotationLayout field if non-nil, zero value otherwise.

### GetQuotationLayoutOk

`func (o *QuotationIndexResponseQuotations) GetQuotationLayoutOk() (*string, bool)`

GetQuotationLayoutOk returns a tuple with the QuotationLayout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotationLayout

`func (o *QuotationIndexResponseQuotations) SetQuotationLayout(v string)`

SetQuotationLayout sets QuotationLayout field to given value.


### GetTaxEntryMethod

`func (o *QuotationIndexResponseQuotations) GetTaxEntryMethod() string`

GetTaxEntryMethod returns the TaxEntryMethod field if non-nil, zero value otherwise.

### GetTaxEntryMethodOk

`func (o *QuotationIndexResponseQuotations) GetTaxEntryMethodOk() (*string, bool)`

GetTaxEntryMethodOk returns a tuple with the TaxEntryMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxEntryMethod

`func (o *QuotationIndexResponseQuotations) SetTaxEntryMethod(v string)`

SetTaxEntryMethod sets TaxEntryMethod field to given value.


### GetQuotationContents

`func (o *QuotationIndexResponseQuotations) GetQuotationContents() []QuotationIndexResponseQuotationContents`

GetQuotationContents returns the QuotationContents field if non-nil, zero value otherwise.

### GetQuotationContentsOk

`func (o *QuotationIndexResponseQuotations) GetQuotationContentsOk() (*[]QuotationIndexResponseQuotationContents, bool)`

GetQuotationContentsOk returns a tuple with the QuotationContents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotationContents

`func (o *QuotationIndexResponseQuotations) SetQuotationContents(v []QuotationIndexResponseQuotationContents)`

SetQuotationContents sets QuotationContents field to given value.

### HasQuotationContents

`func (o *QuotationIndexResponseQuotations) HasQuotationContents() bool`

HasQuotationContents returns a boolean if a field has been set.

### GetTotalAmountPerVatRate

`func (o *QuotationIndexResponseQuotations) GetTotalAmountPerVatRate() InvoiceIndexResponseTotalAmountPerVatRate`

GetTotalAmountPerVatRate returns the TotalAmountPerVatRate field if non-nil, zero value otherwise.

### GetTotalAmountPerVatRateOk

`func (o *QuotationIndexResponseQuotations) GetTotalAmountPerVatRateOk() (*InvoiceIndexResponseTotalAmountPerVatRate, bool)`

GetTotalAmountPerVatRateOk returns a tuple with the TotalAmountPerVatRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmountPerVatRate

`func (o *QuotationIndexResponseQuotations) SetTotalAmountPerVatRate(v InvoiceIndexResponseTotalAmountPerVatRate)`

SetTotalAmountPerVatRate sets TotalAmountPerVatRate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


