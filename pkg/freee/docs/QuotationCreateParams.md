# QuotationCreateParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyId** | **int32** | 事業所ID | 
**IssueDate** | Pointer to **string** | 見積日 (yyyy-mm-dd) | [optional] 
**PartnerId** | Pointer to **NullableInt32** | 取引先ID | [optional] 
**PartnerCode** | Pointer to **NullableString** | 取引先コード | [optional] 
**QuotationNumber** | Pointer to **string** | 見積書番号 (デフォルト: 自動採番されます) | [optional] 
**Title** | Pointer to **string** | タイトル (デフォルト: 見積書) | [optional] 
**Description** | Pointer to **string** | 概要 | [optional] 
**QuotationStatus** | Pointer to **string** | 見積書ステータス  (unsubmitted: 送付待ち, submitted: 送付済み) | [optional] 
**PartnerDisplayName** | **string** | 見積書に表示する取引先名 | 
**PartnerTitle** | **string** | 敬称（御中、様、(空白)の3つから選択） | 
**PartnerContactInfo** | Pointer to **NullableString** | 取引先担当者名 | [optional] 
**PartnerZipcode** | Pointer to **NullableString** | 取引先郵便番号 (デフォルトはpartner_idもしくはpartner_codeで指定された取引先設定情報が補完されます) | [optional] 
**PartnerPrefectureCode** | Pointer to **NullableInt32** | 取引先都道府県コード（0:北海道、1:青森、2:岩手、3:宮城、4:秋田、5:山形、6:福島、7:茨城、8:栃木、9:群馬、10:埼玉、11:千葉、12:東京、13:神奈川、14:新潟、15:富山、16:石川、17:福井、18:山梨、19:長野、20:岐阜、21:静岡、22:愛知、23:三重、24:滋賀、25:京都、26:大阪、27:兵庫、28:奈良、29:和歌山、30:鳥取、31:島根、32:岡山、33:広島、34:山口、35:徳島、36:香川、37:愛媛、38:高知、39:福岡、40:佐賀、41:長崎、42:熊本、43:大分、44:宮崎、45:鹿児島、46:沖縄) (デフォルトはpartner_idもしくはpartner_codeで指定された取引先設定情報が補完されます) | [optional] 
**PartnerAddress1** | Pointer to **NullableString** | 取引先市区町村・番地 (デフォルトはpartner_idもしくはpartner_codeで指定された取引先設定情報が補完されます) | [optional] 
**PartnerAddress2** | Pointer to **NullableString** | 取引先建物名・部屋番号など (デフォルトはpartner_idもしくはpartner_codeで指定された取引先設定情報が補完されます) | [optional] 
**CompanyName** | Pointer to **string** | 事業所名 (デフォルトは事業所設定情報が補完されます) | [optional] 
**CompanyZipcode** | Pointer to **string** | 郵便番号 (デフォルトは事業所設定情報が補完されます) | [optional] 
**CompanyPrefectureCode** | Pointer to **int32** | 都道府県コード（0:北海道、1:青森、2:岩手、3:宮城、4:秋田、5:山形、6:福島、7:茨城、8:栃木、9:群馬、10:埼玉、11:千葉、12:東京、13:神奈川、14:新潟、15:富山、16:石川、17:福井、18:山梨、19:長野、20:岐阜、21:静岡、22:愛知、23:三重、24:滋賀、25:京都、26:大阪、27:兵庫、28:奈良、29:和歌山、30:鳥取、31:島根、32:岡山、33:広島、34:山口、35:徳島、36:香川、37:愛媛、38:高知、39:福岡、40:佐賀、41:長崎、42:熊本、43:大分、44:宮崎、45:鹿児島、46:沖縄) (デフォルトは事業所設定情報が補完されます) | [optional] 
**CompanyAddress1** | Pointer to **string** | 市区町村・番地 (デフォルトは事業所設定情報が補完されます) | [optional] 
**CompanyAddress2** | Pointer to **string** | 建物名・部屋番号など (デフォルトは事業所設定情報が補完されます) | [optional] 
**CompanyContactInfo** | Pointer to **string** | 事業所担当者名 (デフォルトは事業所設定情報が補完されます) | [optional] 
**Message** | Pointer to **string** | メッセージ (デフォルト: 下記の通り御見積申し上げます。) | [optional] 
**Notes** | Pointer to **string** | 備考 | [optional] 
**QuotationLayout** | Pointer to **string** | 見積書レイアウト * &#x60;default_classic&#x60; - レイアウト１/クラシック (デフォルト)  * &#x60;standard_classic&#x60; - レイアウト２/クラシック  * &#x60;envelope_classic&#x60; - 封筒１/クラシック  * &#x60;default_modern&#x60; - レイアウト１/モダン  * &#x60;standard_modern&#x60; - レイアウト２/モダン  * &#x60;envelope_modern&#x60; - 封筒/モダン | [optional] 
**TaxEntryMethod** | Pointer to **string** | 見積書の消費税計算方法(inclusive: 内税表示, exclusive: 外税表示 (デフォルト)) | [optional] 
**QuotationContents** | Pointer to [**[]InvoiceCreateParamsInvoiceContents**](InvoiceCreateParamsInvoiceContents.md) | 見積内容 | [optional] 

## Methods

### NewQuotationCreateParams

`func NewQuotationCreateParams(companyId int32, partnerDisplayName string, partnerTitle string, ) *QuotationCreateParams`

NewQuotationCreateParams instantiates a new QuotationCreateParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuotationCreateParamsWithDefaults

`func NewQuotationCreateParamsWithDefaults() *QuotationCreateParams`

NewQuotationCreateParamsWithDefaults instantiates a new QuotationCreateParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyId

`func (o *QuotationCreateParams) GetCompanyId() int32`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *QuotationCreateParams) GetCompanyIdOk() (*int32, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *QuotationCreateParams) SetCompanyId(v int32)`

SetCompanyId sets CompanyId field to given value.


### GetIssueDate

`func (o *QuotationCreateParams) GetIssueDate() string`

GetIssueDate returns the IssueDate field if non-nil, zero value otherwise.

### GetIssueDateOk

`func (o *QuotationCreateParams) GetIssueDateOk() (*string, bool)`

GetIssueDateOk returns a tuple with the IssueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueDate

`func (o *QuotationCreateParams) SetIssueDate(v string)`

SetIssueDate sets IssueDate field to given value.

### HasIssueDate

`func (o *QuotationCreateParams) HasIssueDate() bool`

HasIssueDate returns a boolean if a field has been set.

### GetPartnerId

`func (o *QuotationCreateParams) GetPartnerId() int32`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *QuotationCreateParams) GetPartnerIdOk() (*int32, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *QuotationCreateParams) SetPartnerId(v int32)`

SetPartnerId sets PartnerId field to given value.

### HasPartnerId

`func (o *QuotationCreateParams) HasPartnerId() bool`

HasPartnerId returns a boolean if a field has been set.

### SetPartnerIdNil

`func (o *QuotationCreateParams) SetPartnerIdNil(b bool)`

 SetPartnerIdNil sets the value for PartnerId to be an explicit nil

### UnsetPartnerId
`func (o *QuotationCreateParams) UnsetPartnerId()`

UnsetPartnerId ensures that no value is present for PartnerId, not even an explicit nil
### GetPartnerCode

`func (o *QuotationCreateParams) GetPartnerCode() string`

GetPartnerCode returns the PartnerCode field if non-nil, zero value otherwise.

### GetPartnerCodeOk

`func (o *QuotationCreateParams) GetPartnerCodeOk() (*string, bool)`

GetPartnerCodeOk returns a tuple with the PartnerCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerCode

`func (o *QuotationCreateParams) SetPartnerCode(v string)`

SetPartnerCode sets PartnerCode field to given value.

### HasPartnerCode

`func (o *QuotationCreateParams) HasPartnerCode() bool`

HasPartnerCode returns a boolean if a field has been set.

### SetPartnerCodeNil

`func (o *QuotationCreateParams) SetPartnerCodeNil(b bool)`

 SetPartnerCodeNil sets the value for PartnerCode to be an explicit nil

### UnsetPartnerCode
`func (o *QuotationCreateParams) UnsetPartnerCode()`

UnsetPartnerCode ensures that no value is present for PartnerCode, not even an explicit nil
### GetQuotationNumber

`func (o *QuotationCreateParams) GetQuotationNumber() string`

GetQuotationNumber returns the QuotationNumber field if non-nil, zero value otherwise.

### GetQuotationNumberOk

`func (o *QuotationCreateParams) GetQuotationNumberOk() (*string, bool)`

GetQuotationNumberOk returns a tuple with the QuotationNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotationNumber

`func (o *QuotationCreateParams) SetQuotationNumber(v string)`

SetQuotationNumber sets QuotationNumber field to given value.

### HasQuotationNumber

`func (o *QuotationCreateParams) HasQuotationNumber() bool`

HasQuotationNumber returns a boolean if a field has been set.

### GetTitle

`func (o *QuotationCreateParams) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *QuotationCreateParams) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *QuotationCreateParams) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *QuotationCreateParams) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *QuotationCreateParams) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *QuotationCreateParams) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *QuotationCreateParams) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *QuotationCreateParams) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetQuotationStatus

`func (o *QuotationCreateParams) GetQuotationStatus() string`

GetQuotationStatus returns the QuotationStatus field if non-nil, zero value otherwise.

### GetQuotationStatusOk

`func (o *QuotationCreateParams) GetQuotationStatusOk() (*string, bool)`

GetQuotationStatusOk returns a tuple with the QuotationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotationStatus

`func (o *QuotationCreateParams) SetQuotationStatus(v string)`

SetQuotationStatus sets QuotationStatus field to given value.

### HasQuotationStatus

`func (o *QuotationCreateParams) HasQuotationStatus() bool`

HasQuotationStatus returns a boolean if a field has been set.

### GetPartnerDisplayName

`func (o *QuotationCreateParams) GetPartnerDisplayName() string`

GetPartnerDisplayName returns the PartnerDisplayName field if non-nil, zero value otherwise.

### GetPartnerDisplayNameOk

`func (o *QuotationCreateParams) GetPartnerDisplayNameOk() (*string, bool)`

GetPartnerDisplayNameOk returns a tuple with the PartnerDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerDisplayName

`func (o *QuotationCreateParams) SetPartnerDisplayName(v string)`

SetPartnerDisplayName sets PartnerDisplayName field to given value.


### GetPartnerTitle

`func (o *QuotationCreateParams) GetPartnerTitle() string`

GetPartnerTitle returns the PartnerTitle field if non-nil, zero value otherwise.

### GetPartnerTitleOk

`func (o *QuotationCreateParams) GetPartnerTitleOk() (*string, bool)`

GetPartnerTitleOk returns a tuple with the PartnerTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerTitle

`func (o *QuotationCreateParams) SetPartnerTitle(v string)`

SetPartnerTitle sets PartnerTitle field to given value.


### GetPartnerContactInfo

`func (o *QuotationCreateParams) GetPartnerContactInfo() string`

GetPartnerContactInfo returns the PartnerContactInfo field if non-nil, zero value otherwise.

### GetPartnerContactInfoOk

`func (o *QuotationCreateParams) GetPartnerContactInfoOk() (*string, bool)`

GetPartnerContactInfoOk returns a tuple with the PartnerContactInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerContactInfo

`func (o *QuotationCreateParams) SetPartnerContactInfo(v string)`

SetPartnerContactInfo sets PartnerContactInfo field to given value.

### HasPartnerContactInfo

`func (o *QuotationCreateParams) HasPartnerContactInfo() bool`

HasPartnerContactInfo returns a boolean if a field has been set.

### SetPartnerContactInfoNil

`func (o *QuotationCreateParams) SetPartnerContactInfoNil(b bool)`

 SetPartnerContactInfoNil sets the value for PartnerContactInfo to be an explicit nil

### UnsetPartnerContactInfo
`func (o *QuotationCreateParams) UnsetPartnerContactInfo()`

UnsetPartnerContactInfo ensures that no value is present for PartnerContactInfo, not even an explicit nil
### GetPartnerZipcode

`func (o *QuotationCreateParams) GetPartnerZipcode() string`

GetPartnerZipcode returns the PartnerZipcode field if non-nil, zero value otherwise.

### GetPartnerZipcodeOk

`func (o *QuotationCreateParams) GetPartnerZipcodeOk() (*string, bool)`

GetPartnerZipcodeOk returns a tuple with the PartnerZipcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerZipcode

`func (o *QuotationCreateParams) SetPartnerZipcode(v string)`

SetPartnerZipcode sets PartnerZipcode field to given value.

### HasPartnerZipcode

`func (o *QuotationCreateParams) HasPartnerZipcode() bool`

HasPartnerZipcode returns a boolean if a field has been set.

### SetPartnerZipcodeNil

`func (o *QuotationCreateParams) SetPartnerZipcodeNil(b bool)`

 SetPartnerZipcodeNil sets the value for PartnerZipcode to be an explicit nil

### UnsetPartnerZipcode
`func (o *QuotationCreateParams) UnsetPartnerZipcode()`

UnsetPartnerZipcode ensures that no value is present for PartnerZipcode, not even an explicit nil
### GetPartnerPrefectureCode

`func (o *QuotationCreateParams) GetPartnerPrefectureCode() int32`

GetPartnerPrefectureCode returns the PartnerPrefectureCode field if non-nil, zero value otherwise.

### GetPartnerPrefectureCodeOk

`func (o *QuotationCreateParams) GetPartnerPrefectureCodeOk() (*int32, bool)`

GetPartnerPrefectureCodeOk returns a tuple with the PartnerPrefectureCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerPrefectureCode

`func (o *QuotationCreateParams) SetPartnerPrefectureCode(v int32)`

SetPartnerPrefectureCode sets PartnerPrefectureCode field to given value.

### HasPartnerPrefectureCode

`func (o *QuotationCreateParams) HasPartnerPrefectureCode() bool`

HasPartnerPrefectureCode returns a boolean if a field has been set.

### SetPartnerPrefectureCodeNil

`func (o *QuotationCreateParams) SetPartnerPrefectureCodeNil(b bool)`

 SetPartnerPrefectureCodeNil sets the value for PartnerPrefectureCode to be an explicit nil

### UnsetPartnerPrefectureCode
`func (o *QuotationCreateParams) UnsetPartnerPrefectureCode()`

UnsetPartnerPrefectureCode ensures that no value is present for PartnerPrefectureCode, not even an explicit nil
### GetPartnerAddress1

`func (o *QuotationCreateParams) GetPartnerAddress1() string`

GetPartnerAddress1 returns the PartnerAddress1 field if non-nil, zero value otherwise.

### GetPartnerAddress1Ok

`func (o *QuotationCreateParams) GetPartnerAddress1Ok() (*string, bool)`

GetPartnerAddress1Ok returns a tuple with the PartnerAddress1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerAddress1

`func (o *QuotationCreateParams) SetPartnerAddress1(v string)`

SetPartnerAddress1 sets PartnerAddress1 field to given value.

### HasPartnerAddress1

`func (o *QuotationCreateParams) HasPartnerAddress1() bool`

HasPartnerAddress1 returns a boolean if a field has been set.

### SetPartnerAddress1Nil

`func (o *QuotationCreateParams) SetPartnerAddress1Nil(b bool)`

 SetPartnerAddress1Nil sets the value for PartnerAddress1 to be an explicit nil

### UnsetPartnerAddress1
`func (o *QuotationCreateParams) UnsetPartnerAddress1()`

UnsetPartnerAddress1 ensures that no value is present for PartnerAddress1, not even an explicit nil
### GetPartnerAddress2

`func (o *QuotationCreateParams) GetPartnerAddress2() string`

GetPartnerAddress2 returns the PartnerAddress2 field if non-nil, zero value otherwise.

### GetPartnerAddress2Ok

`func (o *QuotationCreateParams) GetPartnerAddress2Ok() (*string, bool)`

GetPartnerAddress2Ok returns a tuple with the PartnerAddress2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerAddress2

`func (o *QuotationCreateParams) SetPartnerAddress2(v string)`

SetPartnerAddress2 sets PartnerAddress2 field to given value.

### HasPartnerAddress2

`func (o *QuotationCreateParams) HasPartnerAddress2() bool`

HasPartnerAddress2 returns a boolean if a field has been set.

### SetPartnerAddress2Nil

`func (o *QuotationCreateParams) SetPartnerAddress2Nil(b bool)`

 SetPartnerAddress2Nil sets the value for PartnerAddress2 to be an explicit nil

### UnsetPartnerAddress2
`func (o *QuotationCreateParams) UnsetPartnerAddress2()`

UnsetPartnerAddress2 ensures that no value is present for PartnerAddress2, not even an explicit nil
### GetCompanyName

`func (o *QuotationCreateParams) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *QuotationCreateParams) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *QuotationCreateParams) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *QuotationCreateParams) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetCompanyZipcode

`func (o *QuotationCreateParams) GetCompanyZipcode() string`

GetCompanyZipcode returns the CompanyZipcode field if non-nil, zero value otherwise.

### GetCompanyZipcodeOk

`func (o *QuotationCreateParams) GetCompanyZipcodeOk() (*string, bool)`

GetCompanyZipcodeOk returns a tuple with the CompanyZipcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyZipcode

`func (o *QuotationCreateParams) SetCompanyZipcode(v string)`

SetCompanyZipcode sets CompanyZipcode field to given value.

### HasCompanyZipcode

`func (o *QuotationCreateParams) HasCompanyZipcode() bool`

HasCompanyZipcode returns a boolean if a field has been set.

### GetCompanyPrefectureCode

`func (o *QuotationCreateParams) GetCompanyPrefectureCode() int32`

GetCompanyPrefectureCode returns the CompanyPrefectureCode field if non-nil, zero value otherwise.

### GetCompanyPrefectureCodeOk

`func (o *QuotationCreateParams) GetCompanyPrefectureCodeOk() (*int32, bool)`

GetCompanyPrefectureCodeOk returns a tuple with the CompanyPrefectureCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyPrefectureCode

`func (o *QuotationCreateParams) SetCompanyPrefectureCode(v int32)`

SetCompanyPrefectureCode sets CompanyPrefectureCode field to given value.

### HasCompanyPrefectureCode

`func (o *QuotationCreateParams) HasCompanyPrefectureCode() bool`

HasCompanyPrefectureCode returns a boolean if a field has been set.

### GetCompanyAddress1

`func (o *QuotationCreateParams) GetCompanyAddress1() string`

GetCompanyAddress1 returns the CompanyAddress1 field if non-nil, zero value otherwise.

### GetCompanyAddress1Ok

`func (o *QuotationCreateParams) GetCompanyAddress1Ok() (*string, bool)`

GetCompanyAddress1Ok returns a tuple with the CompanyAddress1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyAddress1

`func (o *QuotationCreateParams) SetCompanyAddress1(v string)`

SetCompanyAddress1 sets CompanyAddress1 field to given value.

### HasCompanyAddress1

`func (o *QuotationCreateParams) HasCompanyAddress1() bool`

HasCompanyAddress1 returns a boolean if a field has been set.

### GetCompanyAddress2

`func (o *QuotationCreateParams) GetCompanyAddress2() string`

GetCompanyAddress2 returns the CompanyAddress2 field if non-nil, zero value otherwise.

### GetCompanyAddress2Ok

`func (o *QuotationCreateParams) GetCompanyAddress2Ok() (*string, bool)`

GetCompanyAddress2Ok returns a tuple with the CompanyAddress2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyAddress2

`func (o *QuotationCreateParams) SetCompanyAddress2(v string)`

SetCompanyAddress2 sets CompanyAddress2 field to given value.

### HasCompanyAddress2

`func (o *QuotationCreateParams) HasCompanyAddress2() bool`

HasCompanyAddress2 returns a boolean if a field has been set.

### GetCompanyContactInfo

`func (o *QuotationCreateParams) GetCompanyContactInfo() string`

GetCompanyContactInfo returns the CompanyContactInfo field if non-nil, zero value otherwise.

### GetCompanyContactInfoOk

`func (o *QuotationCreateParams) GetCompanyContactInfoOk() (*string, bool)`

GetCompanyContactInfoOk returns a tuple with the CompanyContactInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyContactInfo

`func (o *QuotationCreateParams) SetCompanyContactInfo(v string)`

SetCompanyContactInfo sets CompanyContactInfo field to given value.

### HasCompanyContactInfo

`func (o *QuotationCreateParams) HasCompanyContactInfo() bool`

HasCompanyContactInfo returns a boolean if a field has been set.

### GetMessage

`func (o *QuotationCreateParams) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *QuotationCreateParams) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *QuotationCreateParams) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *QuotationCreateParams) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetNotes

`func (o *QuotationCreateParams) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *QuotationCreateParams) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *QuotationCreateParams) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *QuotationCreateParams) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetQuotationLayout

`func (o *QuotationCreateParams) GetQuotationLayout() string`

GetQuotationLayout returns the QuotationLayout field if non-nil, zero value otherwise.

### GetQuotationLayoutOk

`func (o *QuotationCreateParams) GetQuotationLayoutOk() (*string, bool)`

GetQuotationLayoutOk returns a tuple with the QuotationLayout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotationLayout

`func (o *QuotationCreateParams) SetQuotationLayout(v string)`

SetQuotationLayout sets QuotationLayout field to given value.

### HasQuotationLayout

`func (o *QuotationCreateParams) HasQuotationLayout() bool`

HasQuotationLayout returns a boolean if a field has been set.

### GetTaxEntryMethod

`func (o *QuotationCreateParams) GetTaxEntryMethod() string`

GetTaxEntryMethod returns the TaxEntryMethod field if non-nil, zero value otherwise.

### GetTaxEntryMethodOk

`func (o *QuotationCreateParams) GetTaxEntryMethodOk() (*string, bool)`

GetTaxEntryMethodOk returns a tuple with the TaxEntryMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxEntryMethod

`func (o *QuotationCreateParams) SetTaxEntryMethod(v string)`

SetTaxEntryMethod sets TaxEntryMethod field to given value.

### HasTaxEntryMethod

`func (o *QuotationCreateParams) HasTaxEntryMethod() bool`

HasTaxEntryMethod returns a boolean if a field has been set.

### GetQuotationContents

`func (o *QuotationCreateParams) GetQuotationContents() []InvoiceCreateParamsInvoiceContents`

GetQuotationContents returns the QuotationContents field if non-nil, zero value otherwise.

### GetQuotationContentsOk

`func (o *QuotationCreateParams) GetQuotationContentsOk() (*[]InvoiceCreateParamsInvoiceContents, bool)`

GetQuotationContentsOk returns a tuple with the QuotationContents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotationContents

`func (o *QuotationCreateParams) SetQuotationContents(v []InvoiceCreateParamsInvoiceContents)`

SetQuotationContents sets QuotationContents field to given value.

### HasQuotationContents

`func (o *QuotationCreateParams) HasQuotationContents() bool`

HasQuotationContents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


