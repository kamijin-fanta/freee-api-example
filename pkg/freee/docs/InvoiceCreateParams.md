# InvoiceCreateParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyId** | **int32** | 事業所ID | 
**IssueDate** | Pointer to **string** | 請求日 (yyyy-mm-dd) | [optional] 
**PartnerId** | Pointer to **NullableInt32** | 取引先ID | [optional] 
**PartnerCode** | Pointer to **NullableString** | 取引先コード | [optional] 
**InvoiceNumber** | Pointer to **string** | 請求書番号 (デフォルト: 自動採番されます) | [optional] 
**Title** | Pointer to **string** | タイトル (デフォルト: 請求書) | [optional] 
**DueDate** | Pointer to **string** | 期日 (yyyy-mm-dd) | [optional] 
**BookingDate** | Pointer to **string** | 売上計上日 | [optional] 
**Description** | Pointer to **string** | 概要 | [optional] 
**InvoiceStatus** | Pointer to **string** | 請求書ステータス&lt;br&gt; &lt;ul&gt;   &lt;li&gt;draft: 下書き (デフォルト)&lt;/li&gt;   &lt;li&gt;(廃止予定) issue: 発行 (送付待ち (unsubmitted) と同じです。)&lt;/li&gt;   &lt;li&gt;unsubmitted: 送付待ち&lt;/li&gt;   &lt;li&gt;submitted: 送付済み&lt;/li&gt; &lt;/ul&gt; issue, unsubmitted, submitted は請求書承認ワークフローを利用している場合は指定できません。  | [optional] 
**PartnerDisplayName** | **string** | 請求書に表示する取引先名 | 
**PartnerTitle** | **NullableString** | 敬称（御中、様、(空白)の3つから選択） | 
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
**PaymentType** | Pointer to **string** | 支払方法 (振込: transfer, 引き落とし: direct_debit) | [optional] 
**PaymentBankInfo** | Pointer to **string** | 支払口座 | [optional] 
**UseVirtualTransferAccount** | Pointer to **string** | 振込専用口座の利用(利用しない: not_use(デフォルト), 利用する: use) | [optional] 
**Message** | Pointer to **string** | メッセージ (デフォルト: 下記の通りご請求申し上げます。) | [optional] 
**Notes** | Pointer to **string** | 備考 | [optional] 
**InvoiceLayout** | Pointer to **string** | 請求書レイアウト * &#x60;default_classic&#x60; - レイアウト１/クラシック (デフォルト)  * &#x60;standard_classic&#x60; - レイアウト２/クラシック  * &#x60;envelope_classic&#x60; - 封筒１/クラシック  * &#x60;carried_forward_standard_classic&#x60; - レイアウト３（繰越金額欄あり）/クラシック  * &#x60;carried_forward_envelope_classic&#x60; - 封筒２（繰越金額欄あり）/クラシック  * &#x60;default_modern&#x60; - レイアウト１/モダン  * &#x60;standard_modern&#x60; - レイアウト２/モダン  * &#x60;envelope_modern&#x60; - 封筒/モダン | [optional] 
**TaxEntryMethod** | Pointer to **string** | 請求書の消費税計算方法(inclusive: 内税表示, exclusive: 外税表示 (デフォルト)) | [optional] 
**InvoiceContents** | Pointer to [**[]InvoiceCreateParamsInvoiceContents**](InvoiceCreateParamsInvoiceContents.md) | 請求内容 | [optional] 

## Methods

### NewInvoiceCreateParams

`func NewInvoiceCreateParams(companyId int32, partnerDisplayName string, partnerTitle NullableString, ) *InvoiceCreateParams`

NewInvoiceCreateParams instantiates a new InvoiceCreateParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceCreateParamsWithDefaults

`func NewInvoiceCreateParamsWithDefaults() *InvoiceCreateParams`

NewInvoiceCreateParamsWithDefaults instantiates a new InvoiceCreateParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyId

`func (o *InvoiceCreateParams) GetCompanyId() int32`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *InvoiceCreateParams) GetCompanyIdOk() (*int32, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *InvoiceCreateParams) SetCompanyId(v int32)`

SetCompanyId sets CompanyId field to given value.


### GetIssueDate

`func (o *InvoiceCreateParams) GetIssueDate() string`

GetIssueDate returns the IssueDate field if non-nil, zero value otherwise.

### GetIssueDateOk

`func (o *InvoiceCreateParams) GetIssueDateOk() (*string, bool)`

GetIssueDateOk returns a tuple with the IssueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueDate

`func (o *InvoiceCreateParams) SetIssueDate(v string)`

SetIssueDate sets IssueDate field to given value.

### HasIssueDate

`func (o *InvoiceCreateParams) HasIssueDate() bool`

HasIssueDate returns a boolean if a field has been set.

### GetPartnerId

`func (o *InvoiceCreateParams) GetPartnerId() int32`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *InvoiceCreateParams) GetPartnerIdOk() (*int32, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *InvoiceCreateParams) SetPartnerId(v int32)`

SetPartnerId sets PartnerId field to given value.

### HasPartnerId

`func (o *InvoiceCreateParams) HasPartnerId() bool`

HasPartnerId returns a boolean if a field has been set.

### SetPartnerIdNil

`func (o *InvoiceCreateParams) SetPartnerIdNil(b bool)`

 SetPartnerIdNil sets the value for PartnerId to be an explicit nil

### UnsetPartnerId
`func (o *InvoiceCreateParams) UnsetPartnerId()`

UnsetPartnerId ensures that no value is present for PartnerId, not even an explicit nil
### GetPartnerCode

`func (o *InvoiceCreateParams) GetPartnerCode() string`

GetPartnerCode returns the PartnerCode field if non-nil, zero value otherwise.

### GetPartnerCodeOk

`func (o *InvoiceCreateParams) GetPartnerCodeOk() (*string, bool)`

GetPartnerCodeOk returns a tuple with the PartnerCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerCode

`func (o *InvoiceCreateParams) SetPartnerCode(v string)`

SetPartnerCode sets PartnerCode field to given value.

### HasPartnerCode

`func (o *InvoiceCreateParams) HasPartnerCode() bool`

HasPartnerCode returns a boolean if a field has been set.

### SetPartnerCodeNil

`func (o *InvoiceCreateParams) SetPartnerCodeNil(b bool)`

 SetPartnerCodeNil sets the value for PartnerCode to be an explicit nil

### UnsetPartnerCode
`func (o *InvoiceCreateParams) UnsetPartnerCode()`

UnsetPartnerCode ensures that no value is present for PartnerCode, not even an explicit nil
### GetInvoiceNumber

`func (o *InvoiceCreateParams) GetInvoiceNumber() string`

GetInvoiceNumber returns the InvoiceNumber field if non-nil, zero value otherwise.

### GetInvoiceNumberOk

`func (o *InvoiceCreateParams) GetInvoiceNumberOk() (*string, bool)`

GetInvoiceNumberOk returns a tuple with the InvoiceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceNumber

`func (o *InvoiceCreateParams) SetInvoiceNumber(v string)`

SetInvoiceNumber sets InvoiceNumber field to given value.

### HasInvoiceNumber

`func (o *InvoiceCreateParams) HasInvoiceNumber() bool`

HasInvoiceNumber returns a boolean if a field has been set.

### GetTitle

`func (o *InvoiceCreateParams) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *InvoiceCreateParams) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *InvoiceCreateParams) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *InvoiceCreateParams) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDueDate

`func (o *InvoiceCreateParams) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *InvoiceCreateParams) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *InvoiceCreateParams) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *InvoiceCreateParams) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### GetBookingDate

`func (o *InvoiceCreateParams) GetBookingDate() string`

GetBookingDate returns the BookingDate field if non-nil, zero value otherwise.

### GetBookingDateOk

`func (o *InvoiceCreateParams) GetBookingDateOk() (*string, bool)`

GetBookingDateOk returns a tuple with the BookingDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBookingDate

`func (o *InvoiceCreateParams) SetBookingDate(v string)`

SetBookingDate sets BookingDate field to given value.

### HasBookingDate

`func (o *InvoiceCreateParams) HasBookingDate() bool`

HasBookingDate returns a boolean if a field has been set.

### GetDescription

`func (o *InvoiceCreateParams) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InvoiceCreateParams) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InvoiceCreateParams) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InvoiceCreateParams) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetInvoiceStatus

`func (o *InvoiceCreateParams) GetInvoiceStatus() string`

GetInvoiceStatus returns the InvoiceStatus field if non-nil, zero value otherwise.

### GetInvoiceStatusOk

`func (o *InvoiceCreateParams) GetInvoiceStatusOk() (*string, bool)`

GetInvoiceStatusOk returns a tuple with the InvoiceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceStatus

`func (o *InvoiceCreateParams) SetInvoiceStatus(v string)`

SetInvoiceStatus sets InvoiceStatus field to given value.

### HasInvoiceStatus

`func (o *InvoiceCreateParams) HasInvoiceStatus() bool`

HasInvoiceStatus returns a boolean if a field has been set.

### GetPartnerDisplayName

`func (o *InvoiceCreateParams) GetPartnerDisplayName() string`

GetPartnerDisplayName returns the PartnerDisplayName field if non-nil, zero value otherwise.

### GetPartnerDisplayNameOk

`func (o *InvoiceCreateParams) GetPartnerDisplayNameOk() (*string, bool)`

GetPartnerDisplayNameOk returns a tuple with the PartnerDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerDisplayName

`func (o *InvoiceCreateParams) SetPartnerDisplayName(v string)`

SetPartnerDisplayName sets PartnerDisplayName field to given value.


### GetPartnerTitle

`func (o *InvoiceCreateParams) GetPartnerTitle() string`

GetPartnerTitle returns the PartnerTitle field if non-nil, zero value otherwise.

### GetPartnerTitleOk

`func (o *InvoiceCreateParams) GetPartnerTitleOk() (*string, bool)`

GetPartnerTitleOk returns a tuple with the PartnerTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerTitle

`func (o *InvoiceCreateParams) SetPartnerTitle(v string)`

SetPartnerTitle sets PartnerTitle field to given value.


### SetPartnerTitleNil

`func (o *InvoiceCreateParams) SetPartnerTitleNil(b bool)`

 SetPartnerTitleNil sets the value for PartnerTitle to be an explicit nil

### UnsetPartnerTitle
`func (o *InvoiceCreateParams) UnsetPartnerTitle()`

UnsetPartnerTitle ensures that no value is present for PartnerTitle, not even an explicit nil
### GetPartnerContactInfo

`func (o *InvoiceCreateParams) GetPartnerContactInfo() string`

GetPartnerContactInfo returns the PartnerContactInfo field if non-nil, zero value otherwise.

### GetPartnerContactInfoOk

`func (o *InvoiceCreateParams) GetPartnerContactInfoOk() (*string, bool)`

GetPartnerContactInfoOk returns a tuple with the PartnerContactInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerContactInfo

`func (o *InvoiceCreateParams) SetPartnerContactInfo(v string)`

SetPartnerContactInfo sets PartnerContactInfo field to given value.

### HasPartnerContactInfo

`func (o *InvoiceCreateParams) HasPartnerContactInfo() bool`

HasPartnerContactInfo returns a boolean if a field has been set.

### SetPartnerContactInfoNil

`func (o *InvoiceCreateParams) SetPartnerContactInfoNil(b bool)`

 SetPartnerContactInfoNil sets the value for PartnerContactInfo to be an explicit nil

### UnsetPartnerContactInfo
`func (o *InvoiceCreateParams) UnsetPartnerContactInfo()`

UnsetPartnerContactInfo ensures that no value is present for PartnerContactInfo, not even an explicit nil
### GetPartnerZipcode

`func (o *InvoiceCreateParams) GetPartnerZipcode() string`

GetPartnerZipcode returns the PartnerZipcode field if non-nil, zero value otherwise.

### GetPartnerZipcodeOk

`func (o *InvoiceCreateParams) GetPartnerZipcodeOk() (*string, bool)`

GetPartnerZipcodeOk returns a tuple with the PartnerZipcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerZipcode

`func (o *InvoiceCreateParams) SetPartnerZipcode(v string)`

SetPartnerZipcode sets PartnerZipcode field to given value.

### HasPartnerZipcode

`func (o *InvoiceCreateParams) HasPartnerZipcode() bool`

HasPartnerZipcode returns a boolean if a field has been set.

### SetPartnerZipcodeNil

`func (o *InvoiceCreateParams) SetPartnerZipcodeNil(b bool)`

 SetPartnerZipcodeNil sets the value for PartnerZipcode to be an explicit nil

### UnsetPartnerZipcode
`func (o *InvoiceCreateParams) UnsetPartnerZipcode()`

UnsetPartnerZipcode ensures that no value is present for PartnerZipcode, not even an explicit nil
### GetPartnerPrefectureCode

`func (o *InvoiceCreateParams) GetPartnerPrefectureCode() int32`

GetPartnerPrefectureCode returns the PartnerPrefectureCode field if non-nil, zero value otherwise.

### GetPartnerPrefectureCodeOk

`func (o *InvoiceCreateParams) GetPartnerPrefectureCodeOk() (*int32, bool)`

GetPartnerPrefectureCodeOk returns a tuple with the PartnerPrefectureCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerPrefectureCode

`func (o *InvoiceCreateParams) SetPartnerPrefectureCode(v int32)`

SetPartnerPrefectureCode sets PartnerPrefectureCode field to given value.

### HasPartnerPrefectureCode

`func (o *InvoiceCreateParams) HasPartnerPrefectureCode() bool`

HasPartnerPrefectureCode returns a boolean if a field has been set.

### SetPartnerPrefectureCodeNil

`func (o *InvoiceCreateParams) SetPartnerPrefectureCodeNil(b bool)`

 SetPartnerPrefectureCodeNil sets the value for PartnerPrefectureCode to be an explicit nil

### UnsetPartnerPrefectureCode
`func (o *InvoiceCreateParams) UnsetPartnerPrefectureCode()`

UnsetPartnerPrefectureCode ensures that no value is present for PartnerPrefectureCode, not even an explicit nil
### GetPartnerAddress1

`func (o *InvoiceCreateParams) GetPartnerAddress1() string`

GetPartnerAddress1 returns the PartnerAddress1 field if non-nil, zero value otherwise.

### GetPartnerAddress1Ok

`func (o *InvoiceCreateParams) GetPartnerAddress1Ok() (*string, bool)`

GetPartnerAddress1Ok returns a tuple with the PartnerAddress1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerAddress1

`func (o *InvoiceCreateParams) SetPartnerAddress1(v string)`

SetPartnerAddress1 sets PartnerAddress1 field to given value.

### HasPartnerAddress1

`func (o *InvoiceCreateParams) HasPartnerAddress1() bool`

HasPartnerAddress1 returns a boolean if a field has been set.

### SetPartnerAddress1Nil

`func (o *InvoiceCreateParams) SetPartnerAddress1Nil(b bool)`

 SetPartnerAddress1Nil sets the value for PartnerAddress1 to be an explicit nil

### UnsetPartnerAddress1
`func (o *InvoiceCreateParams) UnsetPartnerAddress1()`

UnsetPartnerAddress1 ensures that no value is present for PartnerAddress1, not even an explicit nil
### GetPartnerAddress2

`func (o *InvoiceCreateParams) GetPartnerAddress2() string`

GetPartnerAddress2 returns the PartnerAddress2 field if non-nil, zero value otherwise.

### GetPartnerAddress2Ok

`func (o *InvoiceCreateParams) GetPartnerAddress2Ok() (*string, bool)`

GetPartnerAddress2Ok returns a tuple with the PartnerAddress2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerAddress2

`func (o *InvoiceCreateParams) SetPartnerAddress2(v string)`

SetPartnerAddress2 sets PartnerAddress2 field to given value.

### HasPartnerAddress2

`func (o *InvoiceCreateParams) HasPartnerAddress2() bool`

HasPartnerAddress2 returns a boolean if a field has been set.

### SetPartnerAddress2Nil

`func (o *InvoiceCreateParams) SetPartnerAddress2Nil(b bool)`

 SetPartnerAddress2Nil sets the value for PartnerAddress2 to be an explicit nil

### UnsetPartnerAddress2
`func (o *InvoiceCreateParams) UnsetPartnerAddress2()`

UnsetPartnerAddress2 ensures that no value is present for PartnerAddress2, not even an explicit nil
### GetCompanyName

`func (o *InvoiceCreateParams) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *InvoiceCreateParams) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *InvoiceCreateParams) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *InvoiceCreateParams) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetCompanyZipcode

`func (o *InvoiceCreateParams) GetCompanyZipcode() string`

GetCompanyZipcode returns the CompanyZipcode field if non-nil, zero value otherwise.

### GetCompanyZipcodeOk

`func (o *InvoiceCreateParams) GetCompanyZipcodeOk() (*string, bool)`

GetCompanyZipcodeOk returns a tuple with the CompanyZipcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyZipcode

`func (o *InvoiceCreateParams) SetCompanyZipcode(v string)`

SetCompanyZipcode sets CompanyZipcode field to given value.

### HasCompanyZipcode

`func (o *InvoiceCreateParams) HasCompanyZipcode() bool`

HasCompanyZipcode returns a boolean if a field has been set.

### GetCompanyPrefectureCode

`func (o *InvoiceCreateParams) GetCompanyPrefectureCode() int32`

GetCompanyPrefectureCode returns the CompanyPrefectureCode field if non-nil, zero value otherwise.

### GetCompanyPrefectureCodeOk

`func (o *InvoiceCreateParams) GetCompanyPrefectureCodeOk() (*int32, bool)`

GetCompanyPrefectureCodeOk returns a tuple with the CompanyPrefectureCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyPrefectureCode

`func (o *InvoiceCreateParams) SetCompanyPrefectureCode(v int32)`

SetCompanyPrefectureCode sets CompanyPrefectureCode field to given value.

### HasCompanyPrefectureCode

`func (o *InvoiceCreateParams) HasCompanyPrefectureCode() bool`

HasCompanyPrefectureCode returns a boolean if a field has been set.

### GetCompanyAddress1

`func (o *InvoiceCreateParams) GetCompanyAddress1() string`

GetCompanyAddress1 returns the CompanyAddress1 field if non-nil, zero value otherwise.

### GetCompanyAddress1Ok

`func (o *InvoiceCreateParams) GetCompanyAddress1Ok() (*string, bool)`

GetCompanyAddress1Ok returns a tuple with the CompanyAddress1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyAddress1

`func (o *InvoiceCreateParams) SetCompanyAddress1(v string)`

SetCompanyAddress1 sets CompanyAddress1 field to given value.

### HasCompanyAddress1

`func (o *InvoiceCreateParams) HasCompanyAddress1() bool`

HasCompanyAddress1 returns a boolean if a field has been set.

### GetCompanyAddress2

`func (o *InvoiceCreateParams) GetCompanyAddress2() string`

GetCompanyAddress2 returns the CompanyAddress2 field if non-nil, zero value otherwise.

### GetCompanyAddress2Ok

`func (o *InvoiceCreateParams) GetCompanyAddress2Ok() (*string, bool)`

GetCompanyAddress2Ok returns a tuple with the CompanyAddress2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyAddress2

`func (o *InvoiceCreateParams) SetCompanyAddress2(v string)`

SetCompanyAddress2 sets CompanyAddress2 field to given value.

### HasCompanyAddress2

`func (o *InvoiceCreateParams) HasCompanyAddress2() bool`

HasCompanyAddress2 returns a boolean if a field has been set.

### GetCompanyContactInfo

`func (o *InvoiceCreateParams) GetCompanyContactInfo() string`

GetCompanyContactInfo returns the CompanyContactInfo field if non-nil, zero value otherwise.

### GetCompanyContactInfoOk

`func (o *InvoiceCreateParams) GetCompanyContactInfoOk() (*string, bool)`

GetCompanyContactInfoOk returns a tuple with the CompanyContactInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyContactInfo

`func (o *InvoiceCreateParams) SetCompanyContactInfo(v string)`

SetCompanyContactInfo sets CompanyContactInfo field to given value.

### HasCompanyContactInfo

`func (o *InvoiceCreateParams) HasCompanyContactInfo() bool`

HasCompanyContactInfo returns a boolean if a field has been set.

### GetPaymentType

`func (o *InvoiceCreateParams) GetPaymentType() string`

GetPaymentType returns the PaymentType field if non-nil, zero value otherwise.

### GetPaymentTypeOk

`func (o *InvoiceCreateParams) GetPaymentTypeOk() (*string, bool)`

GetPaymentTypeOk returns a tuple with the PaymentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentType

`func (o *InvoiceCreateParams) SetPaymentType(v string)`

SetPaymentType sets PaymentType field to given value.

### HasPaymentType

`func (o *InvoiceCreateParams) HasPaymentType() bool`

HasPaymentType returns a boolean if a field has been set.

### GetPaymentBankInfo

`func (o *InvoiceCreateParams) GetPaymentBankInfo() string`

GetPaymentBankInfo returns the PaymentBankInfo field if non-nil, zero value otherwise.

### GetPaymentBankInfoOk

`func (o *InvoiceCreateParams) GetPaymentBankInfoOk() (*string, bool)`

GetPaymentBankInfoOk returns a tuple with the PaymentBankInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentBankInfo

`func (o *InvoiceCreateParams) SetPaymentBankInfo(v string)`

SetPaymentBankInfo sets PaymentBankInfo field to given value.

### HasPaymentBankInfo

`func (o *InvoiceCreateParams) HasPaymentBankInfo() bool`

HasPaymentBankInfo returns a boolean if a field has been set.

### GetUseVirtualTransferAccount

`func (o *InvoiceCreateParams) GetUseVirtualTransferAccount() string`

GetUseVirtualTransferAccount returns the UseVirtualTransferAccount field if non-nil, zero value otherwise.

### GetUseVirtualTransferAccountOk

`func (o *InvoiceCreateParams) GetUseVirtualTransferAccountOk() (*string, bool)`

GetUseVirtualTransferAccountOk returns a tuple with the UseVirtualTransferAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseVirtualTransferAccount

`func (o *InvoiceCreateParams) SetUseVirtualTransferAccount(v string)`

SetUseVirtualTransferAccount sets UseVirtualTransferAccount field to given value.

### HasUseVirtualTransferAccount

`func (o *InvoiceCreateParams) HasUseVirtualTransferAccount() bool`

HasUseVirtualTransferAccount returns a boolean if a field has been set.

### GetMessage

`func (o *InvoiceCreateParams) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *InvoiceCreateParams) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *InvoiceCreateParams) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *InvoiceCreateParams) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetNotes

`func (o *InvoiceCreateParams) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *InvoiceCreateParams) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *InvoiceCreateParams) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *InvoiceCreateParams) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetInvoiceLayout

`func (o *InvoiceCreateParams) GetInvoiceLayout() string`

GetInvoiceLayout returns the InvoiceLayout field if non-nil, zero value otherwise.

### GetInvoiceLayoutOk

`func (o *InvoiceCreateParams) GetInvoiceLayoutOk() (*string, bool)`

GetInvoiceLayoutOk returns a tuple with the InvoiceLayout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceLayout

`func (o *InvoiceCreateParams) SetInvoiceLayout(v string)`

SetInvoiceLayout sets InvoiceLayout field to given value.

### HasInvoiceLayout

`func (o *InvoiceCreateParams) HasInvoiceLayout() bool`

HasInvoiceLayout returns a boolean if a field has been set.

### GetTaxEntryMethod

`func (o *InvoiceCreateParams) GetTaxEntryMethod() string`

GetTaxEntryMethod returns the TaxEntryMethod field if non-nil, zero value otherwise.

### GetTaxEntryMethodOk

`func (o *InvoiceCreateParams) GetTaxEntryMethodOk() (*string, bool)`

GetTaxEntryMethodOk returns a tuple with the TaxEntryMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxEntryMethod

`func (o *InvoiceCreateParams) SetTaxEntryMethod(v string)`

SetTaxEntryMethod sets TaxEntryMethod field to given value.

### HasTaxEntryMethod

`func (o *InvoiceCreateParams) HasTaxEntryMethod() bool`

HasTaxEntryMethod returns a boolean if a field has been set.

### GetInvoiceContents

`func (o *InvoiceCreateParams) GetInvoiceContents() []InvoiceCreateParamsInvoiceContents`

GetInvoiceContents returns the InvoiceContents field if non-nil, zero value otherwise.

### GetInvoiceContentsOk

`func (o *InvoiceCreateParams) GetInvoiceContentsOk() (*[]InvoiceCreateParamsInvoiceContents, bool)`

GetInvoiceContentsOk returns a tuple with the InvoiceContents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceContents

`func (o *InvoiceCreateParams) SetInvoiceContents(v []InvoiceCreateParamsInvoiceContents)`

SetInvoiceContents sets InvoiceContents field to given value.

### HasInvoiceContents

`func (o *InvoiceCreateParams) HasInvoiceContents() bool`

HasInvoiceContents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


