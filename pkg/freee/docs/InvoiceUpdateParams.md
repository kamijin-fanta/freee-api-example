# InvoiceUpdateParams

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
**InvoiceStatus** | Pointer to **string** | 請求書ステータス&lt;br&gt; &lt;ul&gt;   &lt;li&gt;draft: 下書き (デフォルト)&lt;/li&gt;   &lt;li&gt;(廃止予定) issue: 発行 (送付待ち (unsubmitted) と同じです。)&lt;/li&gt;   &lt;li&gt;unsubmitted: 送付待ち&lt;/li&gt;   &lt;li&gt;submitted: 送付済み&lt;/li&gt; &lt;/ul&gt; issue, unsubmitted は請求書承認ワークフローを利用している場合は、承認済みの請求書にのみ指定できます。&lt;br&gt; submitted は請求書承認ワークフローを利用している場合は、送付待ちの請求書にのみ指定できます。  | [optional] 
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
**InvoiceContents** | Pointer to [**[]InvoiceUpdateParamsInvoiceContents**](InvoiceUpdateParamsInvoiceContents.md) | 請求内容 | [optional] 

## Methods

### NewInvoiceUpdateParams

`func NewInvoiceUpdateParams(companyId int32, partnerDisplayName string, partnerTitle NullableString, ) *InvoiceUpdateParams`

NewInvoiceUpdateParams instantiates a new InvoiceUpdateParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceUpdateParamsWithDefaults

`func NewInvoiceUpdateParamsWithDefaults() *InvoiceUpdateParams`

NewInvoiceUpdateParamsWithDefaults instantiates a new InvoiceUpdateParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyId

`func (o *InvoiceUpdateParams) GetCompanyId() int32`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *InvoiceUpdateParams) GetCompanyIdOk() (*int32, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *InvoiceUpdateParams) SetCompanyId(v int32)`

SetCompanyId sets CompanyId field to given value.


### GetIssueDate

`func (o *InvoiceUpdateParams) GetIssueDate() string`

GetIssueDate returns the IssueDate field if non-nil, zero value otherwise.

### GetIssueDateOk

`func (o *InvoiceUpdateParams) GetIssueDateOk() (*string, bool)`

GetIssueDateOk returns a tuple with the IssueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueDate

`func (o *InvoiceUpdateParams) SetIssueDate(v string)`

SetIssueDate sets IssueDate field to given value.

### HasIssueDate

`func (o *InvoiceUpdateParams) HasIssueDate() bool`

HasIssueDate returns a boolean if a field has been set.

### GetPartnerId

`func (o *InvoiceUpdateParams) GetPartnerId() int32`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *InvoiceUpdateParams) GetPartnerIdOk() (*int32, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *InvoiceUpdateParams) SetPartnerId(v int32)`

SetPartnerId sets PartnerId field to given value.

### HasPartnerId

`func (o *InvoiceUpdateParams) HasPartnerId() bool`

HasPartnerId returns a boolean if a field has been set.

### SetPartnerIdNil

`func (o *InvoiceUpdateParams) SetPartnerIdNil(b bool)`

 SetPartnerIdNil sets the value for PartnerId to be an explicit nil

### UnsetPartnerId
`func (o *InvoiceUpdateParams) UnsetPartnerId()`

UnsetPartnerId ensures that no value is present for PartnerId, not even an explicit nil
### GetPartnerCode

`func (o *InvoiceUpdateParams) GetPartnerCode() string`

GetPartnerCode returns the PartnerCode field if non-nil, zero value otherwise.

### GetPartnerCodeOk

`func (o *InvoiceUpdateParams) GetPartnerCodeOk() (*string, bool)`

GetPartnerCodeOk returns a tuple with the PartnerCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerCode

`func (o *InvoiceUpdateParams) SetPartnerCode(v string)`

SetPartnerCode sets PartnerCode field to given value.

### HasPartnerCode

`func (o *InvoiceUpdateParams) HasPartnerCode() bool`

HasPartnerCode returns a boolean if a field has been set.

### SetPartnerCodeNil

`func (o *InvoiceUpdateParams) SetPartnerCodeNil(b bool)`

 SetPartnerCodeNil sets the value for PartnerCode to be an explicit nil

### UnsetPartnerCode
`func (o *InvoiceUpdateParams) UnsetPartnerCode()`

UnsetPartnerCode ensures that no value is present for PartnerCode, not even an explicit nil
### GetInvoiceNumber

`func (o *InvoiceUpdateParams) GetInvoiceNumber() string`

GetInvoiceNumber returns the InvoiceNumber field if non-nil, zero value otherwise.

### GetInvoiceNumberOk

`func (o *InvoiceUpdateParams) GetInvoiceNumberOk() (*string, bool)`

GetInvoiceNumberOk returns a tuple with the InvoiceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceNumber

`func (o *InvoiceUpdateParams) SetInvoiceNumber(v string)`

SetInvoiceNumber sets InvoiceNumber field to given value.

### HasInvoiceNumber

`func (o *InvoiceUpdateParams) HasInvoiceNumber() bool`

HasInvoiceNumber returns a boolean if a field has been set.

### GetTitle

`func (o *InvoiceUpdateParams) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *InvoiceUpdateParams) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *InvoiceUpdateParams) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *InvoiceUpdateParams) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDueDate

`func (o *InvoiceUpdateParams) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *InvoiceUpdateParams) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *InvoiceUpdateParams) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *InvoiceUpdateParams) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### GetBookingDate

`func (o *InvoiceUpdateParams) GetBookingDate() string`

GetBookingDate returns the BookingDate field if non-nil, zero value otherwise.

### GetBookingDateOk

`func (o *InvoiceUpdateParams) GetBookingDateOk() (*string, bool)`

GetBookingDateOk returns a tuple with the BookingDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBookingDate

`func (o *InvoiceUpdateParams) SetBookingDate(v string)`

SetBookingDate sets BookingDate field to given value.

### HasBookingDate

`func (o *InvoiceUpdateParams) HasBookingDate() bool`

HasBookingDate returns a boolean if a field has been set.

### GetDescription

`func (o *InvoiceUpdateParams) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InvoiceUpdateParams) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InvoiceUpdateParams) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InvoiceUpdateParams) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetInvoiceStatus

`func (o *InvoiceUpdateParams) GetInvoiceStatus() string`

GetInvoiceStatus returns the InvoiceStatus field if non-nil, zero value otherwise.

### GetInvoiceStatusOk

`func (o *InvoiceUpdateParams) GetInvoiceStatusOk() (*string, bool)`

GetInvoiceStatusOk returns a tuple with the InvoiceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceStatus

`func (o *InvoiceUpdateParams) SetInvoiceStatus(v string)`

SetInvoiceStatus sets InvoiceStatus field to given value.

### HasInvoiceStatus

`func (o *InvoiceUpdateParams) HasInvoiceStatus() bool`

HasInvoiceStatus returns a boolean if a field has been set.

### GetPartnerDisplayName

`func (o *InvoiceUpdateParams) GetPartnerDisplayName() string`

GetPartnerDisplayName returns the PartnerDisplayName field if non-nil, zero value otherwise.

### GetPartnerDisplayNameOk

`func (o *InvoiceUpdateParams) GetPartnerDisplayNameOk() (*string, bool)`

GetPartnerDisplayNameOk returns a tuple with the PartnerDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerDisplayName

`func (o *InvoiceUpdateParams) SetPartnerDisplayName(v string)`

SetPartnerDisplayName sets PartnerDisplayName field to given value.


### GetPartnerTitle

`func (o *InvoiceUpdateParams) GetPartnerTitle() string`

GetPartnerTitle returns the PartnerTitle field if non-nil, zero value otherwise.

### GetPartnerTitleOk

`func (o *InvoiceUpdateParams) GetPartnerTitleOk() (*string, bool)`

GetPartnerTitleOk returns a tuple with the PartnerTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerTitle

`func (o *InvoiceUpdateParams) SetPartnerTitle(v string)`

SetPartnerTitle sets PartnerTitle field to given value.


### SetPartnerTitleNil

`func (o *InvoiceUpdateParams) SetPartnerTitleNil(b bool)`

 SetPartnerTitleNil sets the value for PartnerTitle to be an explicit nil

### UnsetPartnerTitle
`func (o *InvoiceUpdateParams) UnsetPartnerTitle()`

UnsetPartnerTitle ensures that no value is present for PartnerTitle, not even an explicit nil
### GetPartnerContactInfo

`func (o *InvoiceUpdateParams) GetPartnerContactInfo() string`

GetPartnerContactInfo returns the PartnerContactInfo field if non-nil, zero value otherwise.

### GetPartnerContactInfoOk

`func (o *InvoiceUpdateParams) GetPartnerContactInfoOk() (*string, bool)`

GetPartnerContactInfoOk returns a tuple with the PartnerContactInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerContactInfo

`func (o *InvoiceUpdateParams) SetPartnerContactInfo(v string)`

SetPartnerContactInfo sets PartnerContactInfo field to given value.

### HasPartnerContactInfo

`func (o *InvoiceUpdateParams) HasPartnerContactInfo() bool`

HasPartnerContactInfo returns a boolean if a field has been set.

### SetPartnerContactInfoNil

`func (o *InvoiceUpdateParams) SetPartnerContactInfoNil(b bool)`

 SetPartnerContactInfoNil sets the value for PartnerContactInfo to be an explicit nil

### UnsetPartnerContactInfo
`func (o *InvoiceUpdateParams) UnsetPartnerContactInfo()`

UnsetPartnerContactInfo ensures that no value is present for PartnerContactInfo, not even an explicit nil
### GetPartnerZipcode

`func (o *InvoiceUpdateParams) GetPartnerZipcode() string`

GetPartnerZipcode returns the PartnerZipcode field if non-nil, zero value otherwise.

### GetPartnerZipcodeOk

`func (o *InvoiceUpdateParams) GetPartnerZipcodeOk() (*string, bool)`

GetPartnerZipcodeOk returns a tuple with the PartnerZipcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerZipcode

`func (o *InvoiceUpdateParams) SetPartnerZipcode(v string)`

SetPartnerZipcode sets PartnerZipcode field to given value.

### HasPartnerZipcode

`func (o *InvoiceUpdateParams) HasPartnerZipcode() bool`

HasPartnerZipcode returns a boolean if a field has been set.

### SetPartnerZipcodeNil

`func (o *InvoiceUpdateParams) SetPartnerZipcodeNil(b bool)`

 SetPartnerZipcodeNil sets the value for PartnerZipcode to be an explicit nil

### UnsetPartnerZipcode
`func (o *InvoiceUpdateParams) UnsetPartnerZipcode()`

UnsetPartnerZipcode ensures that no value is present for PartnerZipcode, not even an explicit nil
### GetPartnerPrefectureCode

`func (o *InvoiceUpdateParams) GetPartnerPrefectureCode() int32`

GetPartnerPrefectureCode returns the PartnerPrefectureCode field if non-nil, zero value otherwise.

### GetPartnerPrefectureCodeOk

`func (o *InvoiceUpdateParams) GetPartnerPrefectureCodeOk() (*int32, bool)`

GetPartnerPrefectureCodeOk returns a tuple with the PartnerPrefectureCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerPrefectureCode

`func (o *InvoiceUpdateParams) SetPartnerPrefectureCode(v int32)`

SetPartnerPrefectureCode sets PartnerPrefectureCode field to given value.

### HasPartnerPrefectureCode

`func (o *InvoiceUpdateParams) HasPartnerPrefectureCode() bool`

HasPartnerPrefectureCode returns a boolean if a field has been set.

### SetPartnerPrefectureCodeNil

`func (o *InvoiceUpdateParams) SetPartnerPrefectureCodeNil(b bool)`

 SetPartnerPrefectureCodeNil sets the value for PartnerPrefectureCode to be an explicit nil

### UnsetPartnerPrefectureCode
`func (o *InvoiceUpdateParams) UnsetPartnerPrefectureCode()`

UnsetPartnerPrefectureCode ensures that no value is present for PartnerPrefectureCode, not even an explicit nil
### GetPartnerAddress1

`func (o *InvoiceUpdateParams) GetPartnerAddress1() string`

GetPartnerAddress1 returns the PartnerAddress1 field if non-nil, zero value otherwise.

### GetPartnerAddress1Ok

`func (o *InvoiceUpdateParams) GetPartnerAddress1Ok() (*string, bool)`

GetPartnerAddress1Ok returns a tuple with the PartnerAddress1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerAddress1

`func (o *InvoiceUpdateParams) SetPartnerAddress1(v string)`

SetPartnerAddress1 sets PartnerAddress1 field to given value.

### HasPartnerAddress1

`func (o *InvoiceUpdateParams) HasPartnerAddress1() bool`

HasPartnerAddress1 returns a boolean if a field has been set.

### SetPartnerAddress1Nil

`func (o *InvoiceUpdateParams) SetPartnerAddress1Nil(b bool)`

 SetPartnerAddress1Nil sets the value for PartnerAddress1 to be an explicit nil

### UnsetPartnerAddress1
`func (o *InvoiceUpdateParams) UnsetPartnerAddress1()`

UnsetPartnerAddress1 ensures that no value is present for PartnerAddress1, not even an explicit nil
### GetPartnerAddress2

`func (o *InvoiceUpdateParams) GetPartnerAddress2() string`

GetPartnerAddress2 returns the PartnerAddress2 field if non-nil, zero value otherwise.

### GetPartnerAddress2Ok

`func (o *InvoiceUpdateParams) GetPartnerAddress2Ok() (*string, bool)`

GetPartnerAddress2Ok returns a tuple with the PartnerAddress2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerAddress2

`func (o *InvoiceUpdateParams) SetPartnerAddress2(v string)`

SetPartnerAddress2 sets PartnerAddress2 field to given value.

### HasPartnerAddress2

`func (o *InvoiceUpdateParams) HasPartnerAddress2() bool`

HasPartnerAddress2 returns a boolean if a field has been set.

### SetPartnerAddress2Nil

`func (o *InvoiceUpdateParams) SetPartnerAddress2Nil(b bool)`

 SetPartnerAddress2Nil sets the value for PartnerAddress2 to be an explicit nil

### UnsetPartnerAddress2
`func (o *InvoiceUpdateParams) UnsetPartnerAddress2()`

UnsetPartnerAddress2 ensures that no value is present for PartnerAddress2, not even an explicit nil
### GetCompanyName

`func (o *InvoiceUpdateParams) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *InvoiceUpdateParams) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *InvoiceUpdateParams) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *InvoiceUpdateParams) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetCompanyZipcode

`func (o *InvoiceUpdateParams) GetCompanyZipcode() string`

GetCompanyZipcode returns the CompanyZipcode field if non-nil, zero value otherwise.

### GetCompanyZipcodeOk

`func (o *InvoiceUpdateParams) GetCompanyZipcodeOk() (*string, bool)`

GetCompanyZipcodeOk returns a tuple with the CompanyZipcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyZipcode

`func (o *InvoiceUpdateParams) SetCompanyZipcode(v string)`

SetCompanyZipcode sets CompanyZipcode field to given value.

### HasCompanyZipcode

`func (o *InvoiceUpdateParams) HasCompanyZipcode() bool`

HasCompanyZipcode returns a boolean if a field has been set.

### GetCompanyPrefectureCode

`func (o *InvoiceUpdateParams) GetCompanyPrefectureCode() int32`

GetCompanyPrefectureCode returns the CompanyPrefectureCode field if non-nil, zero value otherwise.

### GetCompanyPrefectureCodeOk

`func (o *InvoiceUpdateParams) GetCompanyPrefectureCodeOk() (*int32, bool)`

GetCompanyPrefectureCodeOk returns a tuple with the CompanyPrefectureCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyPrefectureCode

`func (o *InvoiceUpdateParams) SetCompanyPrefectureCode(v int32)`

SetCompanyPrefectureCode sets CompanyPrefectureCode field to given value.

### HasCompanyPrefectureCode

`func (o *InvoiceUpdateParams) HasCompanyPrefectureCode() bool`

HasCompanyPrefectureCode returns a boolean if a field has been set.

### GetCompanyAddress1

`func (o *InvoiceUpdateParams) GetCompanyAddress1() string`

GetCompanyAddress1 returns the CompanyAddress1 field if non-nil, zero value otherwise.

### GetCompanyAddress1Ok

`func (o *InvoiceUpdateParams) GetCompanyAddress1Ok() (*string, bool)`

GetCompanyAddress1Ok returns a tuple with the CompanyAddress1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyAddress1

`func (o *InvoiceUpdateParams) SetCompanyAddress1(v string)`

SetCompanyAddress1 sets CompanyAddress1 field to given value.

### HasCompanyAddress1

`func (o *InvoiceUpdateParams) HasCompanyAddress1() bool`

HasCompanyAddress1 returns a boolean if a field has been set.

### GetCompanyAddress2

`func (o *InvoiceUpdateParams) GetCompanyAddress2() string`

GetCompanyAddress2 returns the CompanyAddress2 field if non-nil, zero value otherwise.

### GetCompanyAddress2Ok

`func (o *InvoiceUpdateParams) GetCompanyAddress2Ok() (*string, bool)`

GetCompanyAddress2Ok returns a tuple with the CompanyAddress2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyAddress2

`func (o *InvoiceUpdateParams) SetCompanyAddress2(v string)`

SetCompanyAddress2 sets CompanyAddress2 field to given value.

### HasCompanyAddress2

`func (o *InvoiceUpdateParams) HasCompanyAddress2() bool`

HasCompanyAddress2 returns a boolean if a field has been set.

### GetCompanyContactInfo

`func (o *InvoiceUpdateParams) GetCompanyContactInfo() string`

GetCompanyContactInfo returns the CompanyContactInfo field if non-nil, zero value otherwise.

### GetCompanyContactInfoOk

`func (o *InvoiceUpdateParams) GetCompanyContactInfoOk() (*string, bool)`

GetCompanyContactInfoOk returns a tuple with the CompanyContactInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyContactInfo

`func (o *InvoiceUpdateParams) SetCompanyContactInfo(v string)`

SetCompanyContactInfo sets CompanyContactInfo field to given value.

### HasCompanyContactInfo

`func (o *InvoiceUpdateParams) HasCompanyContactInfo() bool`

HasCompanyContactInfo returns a boolean if a field has been set.

### GetPaymentType

`func (o *InvoiceUpdateParams) GetPaymentType() string`

GetPaymentType returns the PaymentType field if non-nil, zero value otherwise.

### GetPaymentTypeOk

`func (o *InvoiceUpdateParams) GetPaymentTypeOk() (*string, bool)`

GetPaymentTypeOk returns a tuple with the PaymentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentType

`func (o *InvoiceUpdateParams) SetPaymentType(v string)`

SetPaymentType sets PaymentType field to given value.

### HasPaymentType

`func (o *InvoiceUpdateParams) HasPaymentType() bool`

HasPaymentType returns a boolean if a field has been set.

### GetPaymentBankInfo

`func (o *InvoiceUpdateParams) GetPaymentBankInfo() string`

GetPaymentBankInfo returns the PaymentBankInfo field if non-nil, zero value otherwise.

### GetPaymentBankInfoOk

`func (o *InvoiceUpdateParams) GetPaymentBankInfoOk() (*string, bool)`

GetPaymentBankInfoOk returns a tuple with the PaymentBankInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentBankInfo

`func (o *InvoiceUpdateParams) SetPaymentBankInfo(v string)`

SetPaymentBankInfo sets PaymentBankInfo field to given value.

### HasPaymentBankInfo

`func (o *InvoiceUpdateParams) HasPaymentBankInfo() bool`

HasPaymentBankInfo returns a boolean if a field has been set.

### GetUseVirtualTransferAccount

`func (o *InvoiceUpdateParams) GetUseVirtualTransferAccount() string`

GetUseVirtualTransferAccount returns the UseVirtualTransferAccount field if non-nil, zero value otherwise.

### GetUseVirtualTransferAccountOk

`func (o *InvoiceUpdateParams) GetUseVirtualTransferAccountOk() (*string, bool)`

GetUseVirtualTransferAccountOk returns a tuple with the UseVirtualTransferAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseVirtualTransferAccount

`func (o *InvoiceUpdateParams) SetUseVirtualTransferAccount(v string)`

SetUseVirtualTransferAccount sets UseVirtualTransferAccount field to given value.

### HasUseVirtualTransferAccount

`func (o *InvoiceUpdateParams) HasUseVirtualTransferAccount() bool`

HasUseVirtualTransferAccount returns a boolean if a field has been set.

### GetMessage

`func (o *InvoiceUpdateParams) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *InvoiceUpdateParams) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *InvoiceUpdateParams) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *InvoiceUpdateParams) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetNotes

`func (o *InvoiceUpdateParams) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *InvoiceUpdateParams) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *InvoiceUpdateParams) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *InvoiceUpdateParams) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetInvoiceLayout

`func (o *InvoiceUpdateParams) GetInvoiceLayout() string`

GetInvoiceLayout returns the InvoiceLayout field if non-nil, zero value otherwise.

### GetInvoiceLayoutOk

`func (o *InvoiceUpdateParams) GetInvoiceLayoutOk() (*string, bool)`

GetInvoiceLayoutOk returns a tuple with the InvoiceLayout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceLayout

`func (o *InvoiceUpdateParams) SetInvoiceLayout(v string)`

SetInvoiceLayout sets InvoiceLayout field to given value.

### HasInvoiceLayout

`func (o *InvoiceUpdateParams) HasInvoiceLayout() bool`

HasInvoiceLayout returns a boolean if a field has been set.

### GetTaxEntryMethod

`func (o *InvoiceUpdateParams) GetTaxEntryMethod() string`

GetTaxEntryMethod returns the TaxEntryMethod field if non-nil, zero value otherwise.

### GetTaxEntryMethodOk

`func (o *InvoiceUpdateParams) GetTaxEntryMethodOk() (*string, bool)`

GetTaxEntryMethodOk returns a tuple with the TaxEntryMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxEntryMethod

`func (o *InvoiceUpdateParams) SetTaxEntryMethod(v string)`

SetTaxEntryMethod sets TaxEntryMethod field to given value.

### HasTaxEntryMethod

`func (o *InvoiceUpdateParams) HasTaxEntryMethod() bool`

HasTaxEntryMethod returns a boolean if a field has been set.

### GetInvoiceContents

`func (o *InvoiceUpdateParams) GetInvoiceContents() []InvoiceUpdateParamsInvoiceContents`

GetInvoiceContents returns the InvoiceContents field if non-nil, zero value otherwise.

### GetInvoiceContentsOk

`func (o *InvoiceUpdateParams) GetInvoiceContentsOk() (*[]InvoiceUpdateParamsInvoiceContents, bool)`

GetInvoiceContentsOk returns a tuple with the InvoiceContents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceContents

`func (o *InvoiceUpdateParams) SetInvoiceContents(v []InvoiceUpdateParamsInvoiceContents)`

SetInvoiceContents sets InvoiceContents field to given value.

### HasInvoiceContents

`func (o *InvoiceUpdateParams) HasInvoiceContents() bool`

HasInvoiceContents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


