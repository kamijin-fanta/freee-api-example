# PartnerResponsePartner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | 取引先ID | 
**Code** | **NullableString** | 取引先コード | 
**CompanyId** | **int32** | 事業所ID | 
**Name** | **string** | 取引先名 | 
**Shortcut1** | Pointer to **NullableString** | ショートカット1 (20文字以内) | [optional] 
**Shortcut2** | Pointer to **NullableString** | ショートカット2 (20文字以内) | [optional] 
**OrgCode** | Pointer to **NullableInt32** | 事業所種別（null: 未設定、1: 法人、2: 個人） | [optional] 
**CountryCode** | Pointer to **string** | 地域（JP: 国内、ZZ:国外） | [optional] 
**LongName** | Pointer to **NullableString** | 正式名称（255文字以内） | [optional] 
**NameKana** | Pointer to **NullableString** | カナ名称（255文字以内） | [optional] 
**DefaultTitle** | Pointer to **NullableString** | 敬称（御中、様、(空白)の3つから選択） | [optional] 
**Phone** | Pointer to **NullableString** | 電話番号 | [optional] 
**ContactName** | Pointer to **NullableString** | 担当者 氏名 | [optional] 
**Email** | Pointer to **NullableString** | 担当者 メールアドレス | [optional] 
**PayerWalletableId** | Pointer to **NullableInt32** | 振込元口座ID（一括振込ファイル用）:（未設定の場合は、nullです。） | [optional] 
**TransferFeeHandlingSide** | Pointer to **string** | 振込手数料負担（一括振込ファイル用）: (振込元(当方): payer, 振込先(先方): payee) | [optional] 
**AddressAttributesZipcode** | Pointer to **NullableString** | 郵便番号 | [optional] 
**AddressAttributesPrefectureCode** | Pointer to **int32** | 都道府県コード（0:北海道、1:青森、2:岩手、3:宮城、4:秋田、5:山形、6:福島、7:茨城、8:栃木、9:群馬、10:埼玉、11:千葉、12:東京、13:神奈川、14:新潟、15:富山、16:石川、17:福井、18:山梨、19:長野、20:岐阜、21:静岡、22:愛知、23:三重、24:滋賀、25:京都、26:大阪、27:兵庫、28:奈良、29:和歌山、30:鳥取、31:島根、32:岡山、33:広島、34:山口、35:徳島、36:香川、37:愛媛、38:高知、39:福岡、40:佐賀、41:長崎、42:熊本、43:大分、44:宮崎、45:鹿児島、46:沖縄 | [optional] 
**AddressAttributesStreetName1** | Pointer to **NullableString** | 市区町村・番地 | [optional] 
**AddressAttributesStreetName2** | Pointer to **NullableString** | 建物名・部屋番号など | [optional] 
**PartnerDocSettingAttributesSendingMethod** | Pointer to **NullableString** | 請求書送付方法(email:メール、posting:郵送、email_and_posting:メールと郵送) | [optional] 
**PartnerBankAccountAttributesBankName** | Pointer to **NullableString** | 銀行名 | [optional] 
**PartnerBankAccountAttributesBankNameKana** | Pointer to **NullableString** | 銀行名（カナ） | [optional] 
**PartnerBankAccountAttributesBankCode** | Pointer to **NullableString** | 銀行コード | [optional] 
**PartnerBankAccountAttributesBranchName** | Pointer to **NullableString** | 支店名 | [optional] 
**PartnerBankAccountAttributesBranchKana** | Pointer to **NullableString** | 支店名（カナ） | [optional] 
**PartnerBankAccountAttributesBranchCode** | Pointer to **NullableString** | 支店番号 | [optional] 
**PartnerBankAccountAttributesAccountType** | Pointer to **NullableString** | 口座種別(ordinary:普通、checking:当座、earmarked:納税準備預金、savings:貯蓄、other:その他) | [optional] 
**PartnerBankAccountAttributesAccountNumber** | Pointer to **NullableString** | 口座番号 | [optional] 
**PartnerBankAccountAttributesAccountName** | Pointer to **NullableString** | 受取人名（カナ） | [optional] 
**PartnerBankAccountAttributesLongAccountName** | Pointer to **NullableString** | 受取人名 | [optional] 
**PaymentTermAttributesCutoffDay** | Pointer to **NullableInt32** | 締め日（29, 30, 31日の末日を指定する場合は、32。） | [optional] 
**PaymentTermAttributesAdditionalMonths** | Pointer to **NullableInt32** | 支払月 | [optional] 
**PaymentTermAttributesFixedDay** | Pointer to **NullableInt32** | 支払日（29, 30, 31日の末日を指定する場合は、32。） | [optional] 
**InvoicePaymentTermAttributesCutoffDay** | Pointer to **NullableInt32** | 締め日（29, 30, 31日の末日を指定する場合は、32。） | [optional] 
**InvoicePaymentTermAttributesAdditionalMonths** | Pointer to **NullableInt32** | 支払月 | [optional] 
**InvoicePaymentTermAttributesFixedDay** | Pointer to **NullableInt32** | 支払日（29, 30, 31日の末日を指定する場合は、32。） | [optional] 

## Methods

### NewPartnerResponsePartner

`func NewPartnerResponsePartner(id int32, code NullableString, companyId int32, name string, ) *PartnerResponsePartner`

NewPartnerResponsePartner instantiates a new PartnerResponsePartner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartnerResponsePartnerWithDefaults

`func NewPartnerResponsePartnerWithDefaults() *PartnerResponsePartner`

NewPartnerResponsePartnerWithDefaults instantiates a new PartnerResponsePartner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PartnerResponsePartner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PartnerResponsePartner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PartnerResponsePartner) SetId(v int32)`

SetId sets Id field to given value.


### GetCode

`func (o *PartnerResponsePartner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *PartnerResponsePartner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *PartnerResponsePartner) SetCode(v string)`

SetCode sets Code field to given value.


### SetCodeNil

`func (o *PartnerResponsePartner) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *PartnerResponsePartner) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetCompanyId

`func (o *PartnerResponsePartner) GetCompanyId() int32`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *PartnerResponsePartner) GetCompanyIdOk() (*int32, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *PartnerResponsePartner) SetCompanyId(v int32)`

SetCompanyId sets CompanyId field to given value.


### GetName

`func (o *PartnerResponsePartner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PartnerResponsePartner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PartnerResponsePartner) SetName(v string)`

SetName sets Name field to given value.


### GetShortcut1

`func (o *PartnerResponsePartner) GetShortcut1() string`

GetShortcut1 returns the Shortcut1 field if non-nil, zero value otherwise.

### GetShortcut1Ok

`func (o *PartnerResponsePartner) GetShortcut1Ok() (*string, bool)`

GetShortcut1Ok returns a tuple with the Shortcut1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortcut1

`func (o *PartnerResponsePartner) SetShortcut1(v string)`

SetShortcut1 sets Shortcut1 field to given value.

### HasShortcut1

`func (o *PartnerResponsePartner) HasShortcut1() bool`

HasShortcut1 returns a boolean if a field has been set.

### SetShortcut1Nil

`func (o *PartnerResponsePartner) SetShortcut1Nil(b bool)`

 SetShortcut1Nil sets the value for Shortcut1 to be an explicit nil

### UnsetShortcut1
`func (o *PartnerResponsePartner) UnsetShortcut1()`

UnsetShortcut1 ensures that no value is present for Shortcut1, not even an explicit nil
### GetShortcut2

`func (o *PartnerResponsePartner) GetShortcut2() string`

GetShortcut2 returns the Shortcut2 field if non-nil, zero value otherwise.

### GetShortcut2Ok

`func (o *PartnerResponsePartner) GetShortcut2Ok() (*string, bool)`

GetShortcut2Ok returns a tuple with the Shortcut2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortcut2

`func (o *PartnerResponsePartner) SetShortcut2(v string)`

SetShortcut2 sets Shortcut2 field to given value.

### HasShortcut2

`func (o *PartnerResponsePartner) HasShortcut2() bool`

HasShortcut2 returns a boolean if a field has been set.

### SetShortcut2Nil

`func (o *PartnerResponsePartner) SetShortcut2Nil(b bool)`

 SetShortcut2Nil sets the value for Shortcut2 to be an explicit nil

### UnsetShortcut2
`func (o *PartnerResponsePartner) UnsetShortcut2()`

UnsetShortcut2 ensures that no value is present for Shortcut2, not even an explicit nil
### GetOrgCode

`func (o *PartnerResponsePartner) GetOrgCode() int32`

GetOrgCode returns the OrgCode field if non-nil, zero value otherwise.

### GetOrgCodeOk

`func (o *PartnerResponsePartner) GetOrgCodeOk() (*int32, bool)`

GetOrgCodeOk returns a tuple with the OrgCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgCode

`func (o *PartnerResponsePartner) SetOrgCode(v int32)`

SetOrgCode sets OrgCode field to given value.

### HasOrgCode

`func (o *PartnerResponsePartner) HasOrgCode() bool`

HasOrgCode returns a boolean if a field has been set.

### SetOrgCodeNil

`func (o *PartnerResponsePartner) SetOrgCodeNil(b bool)`

 SetOrgCodeNil sets the value for OrgCode to be an explicit nil

### UnsetOrgCode
`func (o *PartnerResponsePartner) UnsetOrgCode()`

UnsetOrgCode ensures that no value is present for OrgCode, not even an explicit nil
### GetCountryCode

`func (o *PartnerResponsePartner) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *PartnerResponsePartner) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *PartnerResponsePartner) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *PartnerResponsePartner) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetLongName

`func (o *PartnerResponsePartner) GetLongName() string`

GetLongName returns the LongName field if non-nil, zero value otherwise.

### GetLongNameOk

`func (o *PartnerResponsePartner) GetLongNameOk() (*string, bool)`

GetLongNameOk returns a tuple with the LongName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongName

`func (o *PartnerResponsePartner) SetLongName(v string)`

SetLongName sets LongName field to given value.

### HasLongName

`func (o *PartnerResponsePartner) HasLongName() bool`

HasLongName returns a boolean if a field has been set.

### SetLongNameNil

`func (o *PartnerResponsePartner) SetLongNameNil(b bool)`

 SetLongNameNil sets the value for LongName to be an explicit nil

### UnsetLongName
`func (o *PartnerResponsePartner) UnsetLongName()`

UnsetLongName ensures that no value is present for LongName, not even an explicit nil
### GetNameKana

`func (o *PartnerResponsePartner) GetNameKana() string`

GetNameKana returns the NameKana field if non-nil, zero value otherwise.

### GetNameKanaOk

`func (o *PartnerResponsePartner) GetNameKanaOk() (*string, bool)`

GetNameKanaOk returns a tuple with the NameKana field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameKana

`func (o *PartnerResponsePartner) SetNameKana(v string)`

SetNameKana sets NameKana field to given value.

### HasNameKana

`func (o *PartnerResponsePartner) HasNameKana() bool`

HasNameKana returns a boolean if a field has been set.

### SetNameKanaNil

`func (o *PartnerResponsePartner) SetNameKanaNil(b bool)`

 SetNameKanaNil sets the value for NameKana to be an explicit nil

### UnsetNameKana
`func (o *PartnerResponsePartner) UnsetNameKana()`

UnsetNameKana ensures that no value is present for NameKana, not even an explicit nil
### GetDefaultTitle

`func (o *PartnerResponsePartner) GetDefaultTitle() string`

GetDefaultTitle returns the DefaultTitle field if non-nil, zero value otherwise.

### GetDefaultTitleOk

`func (o *PartnerResponsePartner) GetDefaultTitleOk() (*string, bool)`

GetDefaultTitleOk returns a tuple with the DefaultTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTitle

`func (o *PartnerResponsePartner) SetDefaultTitle(v string)`

SetDefaultTitle sets DefaultTitle field to given value.

### HasDefaultTitle

`func (o *PartnerResponsePartner) HasDefaultTitle() bool`

HasDefaultTitle returns a boolean if a field has been set.

### SetDefaultTitleNil

`func (o *PartnerResponsePartner) SetDefaultTitleNil(b bool)`

 SetDefaultTitleNil sets the value for DefaultTitle to be an explicit nil

### UnsetDefaultTitle
`func (o *PartnerResponsePartner) UnsetDefaultTitle()`

UnsetDefaultTitle ensures that no value is present for DefaultTitle, not even an explicit nil
### GetPhone

`func (o *PartnerResponsePartner) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *PartnerResponsePartner) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *PartnerResponsePartner) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *PartnerResponsePartner) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### SetPhoneNil

`func (o *PartnerResponsePartner) SetPhoneNil(b bool)`

 SetPhoneNil sets the value for Phone to be an explicit nil

### UnsetPhone
`func (o *PartnerResponsePartner) UnsetPhone()`

UnsetPhone ensures that no value is present for Phone, not even an explicit nil
### GetContactName

`func (o *PartnerResponsePartner) GetContactName() string`

GetContactName returns the ContactName field if non-nil, zero value otherwise.

### GetContactNameOk

`func (o *PartnerResponsePartner) GetContactNameOk() (*string, bool)`

GetContactNameOk returns a tuple with the ContactName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactName

`func (o *PartnerResponsePartner) SetContactName(v string)`

SetContactName sets ContactName field to given value.

### HasContactName

`func (o *PartnerResponsePartner) HasContactName() bool`

HasContactName returns a boolean if a field has been set.

### SetContactNameNil

`func (o *PartnerResponsePartner) SetContactNameNil(b bool)`

 SetContactNameNil sets the value for ContactName to be an explicit nil

### UnsetContactName
`func (o *PartnerResponsePartner) UnsetContactName()`

UnsetContactName ensures that no value is present for ContactName, not even an explicit nil
### GetEmail

`func (o *PartnerResponsePartner) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *PartnerResponsePartner) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *PartnerResponsePartner) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *PartnerResponsePartner) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *PartnerResponsePartner) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *PartnerResponsePartner) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetPayerWalletableId

`func (o *PartnerResponsePartner) GetPayerWalletableId() int32`

GetPayerWalletableId returns the PayerWalletableId field if non-nil, zero value otherwise.

### GetPayerWalletableIdOk

`func (o *PartnerResponsePartner) GetPayerWalletableIdOk() (*int32, bool)`

GetPayerWalletableIdOk returns a tuple with the PayerWalletableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayerWalletableId

`func (o *PartnerResponsePartner) SetPayerWalletableId(v int32)`

SetPayerWalletableId sets PayerWalletableId field to given value.

### HasPayerWalletableId

`func (o *PartnerResponsePartner) HasPayerWalletableId() bool`

HasPayerWalletableId returns a boolean if a field has been set.

### SetPayerWalletableIdNil

`func (o *PartnerResponsePartner) SetPayerWalletableIdNil(b bool)`

 SetPayerWalletableIdNil sets the value for PayerWalletableId to be an explicit nil

### UnsetPayerWalletableId
`func (o *PartnerResponsePartner) UnsetPayerWalletableId()`

UnsetPayerWalletableId ensures that no value is present for PayerWalletableId, not even an explicit nil
### GetTransferFeeHandlingSide

`func (o *PartnerResponsePartner) GetTransferFeeHandlingSide() string`

GetTransferFeeHandlingSide returns the TransferFeeHandlingSide field if non-nil, zero value otherwise.

### GetTransferFeeHandlingSideOk

`func (o *PartnerResponsePartner) GetTransferFeeHandlingSideOk() (*string, bool)`

GetTransferFeeHandlingSideOk returns a tuple with the TransferFeeHandlingSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferFeeHandlingSide

`func (o *PartnerResponsePartner) SetTransferFeeHandlingSide(v string)`

SetTransferFeeHandlingSide sets TransferFeeHandlingSide field to given value.

### HasTransferFeeHandlingSide

`func (o *PartnerResponsePartner) HasTransferFeeHandlingSide() bool`

HasTransferFeeHandlingSide returns a boolean if a field has been set.

### GetAddressAttributesZipcode

`func (o *PartnerResponsePartner) GetAddressAttributesZipcode() string`

GetAddressAttributesZipcode returns the AddressAttributesZipcode field if non-nil, zero value otherwise.

### GetAddressAttributesZipcodeOk

`func (o *PartnerResponsePartner) GetAddressAttributesZipcodeOk() (*string, bool)`

GetAddressAttributesZipcodeOk returns a tuple with the AddressAttributesZipcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressAttributesZipcode

`func (o *PartnerResponsePartner) SetAddressAttributesZipcode(v string)`

SetAddressAttributesZipcode sets AddressAttributesZipcode field to given value.

### HasAddressAttributesZipcode

`func (o *PartnerResponsePartner) HasAddressAttributesZipcode() bool`

HasAddressAttributesZipcode returns a boolean if a field has been set.

### SetAddressAttributesZipcodeNil

`func (o *PartnerResponsePartner) SetAddressAttributesZipcodeNil(b bool)`

 SetAddressAttributesZipcodeNil sets the value for AddressAttributesZipcode to be an explicit nil

### UnsetAddressAttributesZipcode
`func (o *PartnerResponsePartner) UnsetAddressAttributesZipcode()`

UnsetAddressAttributesZipcode ensures that no value is present for AddressAttributesZipcode, not even an explicit nil
### GetAddressAttributesPrefectureCode

`func (o *PartnerResponsePartner) GetAddressAttributesPrefectureCode() int32`

GetAddressAttributesPrefectureCode returns the AddressAttributesPrefectureCode field if non-nil, zero value otherwise.

### GetAddressAttributesPrefectureCodeOk

`func (o *PartnerResponsePartner) GetAddressAttributesPrefectureCodeOk() (*int32, bool)`

GetAddressAttributesPrefectureCodeOk returns a tuple with the AddressAttributesPrefectureCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressAttributesPrefectureCode

`func (o *PartnerResponsePartner) SetAddressAttributesPrefectureCode(v int32)`

SetAddressAttributesPrefectureCode sets AddressAttributesPrefectureCode field to given value.

### HasAddressAttributesPrefectureCode

`func (o *PartnerResponsePartner) HasAddressAttributesPrefectureCode() bool`

HasAddressAttributesPrefectureCode returns a boolean if a field has been set.

### GetAddressAttributesStreetName1

`func (o *PartnerResponsePartner) GetAddressAttributesStreetName1() string`

GetAddressAttributesStreetName1 returns the AddressAttributesStreetName1 field if non-nil, zero value otherwise.

### GetAddressAttributesStreetName1Ok

`func (o *PartnerResponsePartner) GetAddressAttributesStreetName1Ok() (*string, bool)`

GetAddressAttributesStreetName1Ok returns a tuple with the AddressAttributesStreetName1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressAttributesStreetName1

`func (o *PartnerResponsePartner) SetAddressAttributesStreetName1(v string)`

SetAddressAttributesStreetName1 sets AddressAttributesStreetName1 field to given value.

### HasAddressAttributesStreetName1

`func (o *PartnerResponsePartner) HasAddressAttributesStreetName1() bool`

HasAddressAttributesStreetName1 returns a boolean if a field has been set.

### SetAddressAttributesStreetName1Nil

`func (o *PartnerResponsePartner) SetAddressAttributesStreetName1Nil(b bool)`

 SetAddressAttributesStreetName1Nil sets the value for AddressAttributesStreetName1 to be an explicit nil

### UnsetAddressAttributesStreetName1
`func (o *PartnerResponsePartner) UnsetAddressAttributesStreetName1()`

UnsetAddressAttributesStreetName1 ensures that no value is present for AddressAttributesStreetName1, not even an explicit nil
### GetAddressAttributesStreetName2

`func (o *PartnerResponsePartner) GetAddressAttributesStreetName2() string`

GetAddressAttributesStreetName2 returns the AddressAttributesStreetName2 field if non-nil, zero value otherwise.

### GetAddressAttributesStreetName2Ok

`func (o *PartnerResponsePartner) GetAddressAttributesStreetName2Ok() (*string, bool)`

GetAddressAttributesStreetName2Ok returns a tuple with the AddressAttributesStreetName2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressAttributesStreetName2

`func (o *PartnerResponsePartner) SetAddressAttributesStreetName2(v string)`

SetAddressAttributesStreetName2 sets AddressAttributesStreetName2 field to given value.

### HasAddressAttributesStreetName2

`func (o *PartnerResponsePartner) HasAddressAttributesStreetName2() bool`

HasAddressAttributesStreetName2 returns a boolean if a field has been set.

### SetAddressAttributesStreetName2Nil

`func (o *PartnerResponsePartner) SetAddressAttributesStreetName2Nil(b bool)`

 SetAddressAttributesStreetName2Nil sets the value for AddressAttributesStreetName2 to be an explicit nil

### UnsetAddressAttributesStreetName2
`func (o *PartnerResponsePartner) UnsetAddressAttributesStreetName2()`

UnsetAddressAttributesStreetName2 ensures that no value is present for AddressAttributesStreetName2, not even an explicit nil
### GetPartnerDocSettingAttributesSendingMethod

`func (o *PartnerResponsePartner) GetPartnerDocSettingAttributesSendingMethod() string`

GetPartnerDocSettingAttributesSendingMethod returns the PartnerDocSettingAttributesSendingMethod field if non-nil, zero value otherwise.

### GetPartnerDocSettingAttributesSendingMethodOk

`func (o *PartnerResponsePartner) GetPartnerDocSettingAttributesSendingMethodOk() (*string, bool)`

GetPartnerDocSettingAttributesSendingMethodOk returns a tuple with the PartnerDocSettingAttributesSendingMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerDocSettingAttributesSendingMethod

`func (o *PartnerResponsePartner) SetPartnerDocSettingAttributesSendingMethod(v string)`

SetPartnerDocSettingAttributesSendingMethod sets PartnerDocSettingAttributesSendingMethod field to given value.

### HasPartnerDocSettingAttributesSendingMethod

`func (o *PartnerResponsePartner) HasPartnerDocSettingAttributesSendingMethod() bool`

HasPartnerDocSettingAttributesSendingMethod returns a boolean if a field has been set.

### SetPartnerDocSettingAttributesSendingMethodNil

`func (o *PartnerResponsePartner) SetPartnerDocSettingAttributesSendingMethodNil(b bool)`

 SetPartnerDocSettingAttributesSendingMethodNil sets the value for PartnerDocSettingAttributesSendingMethod to be an explicit nil

### UnsetPartnerDocSettingAttributesSendingMethod
`func (o *PartnerResponsePartner) UnsetPartnerDocSettingAttributesSendingMethod()`

UnsetPartnerDocSettingAttributesSendingMethod ensures that no value is present for PartnerDocSettingAttributesSendingMethod, not even an explicit nil
### GetPartnerBankAccountAttributesBankName

`func (o *PartnerResponsePartner) GetPartnerBankAccountAttributesBankName() string`

GetPartnerBankAccountAttributesBankName returns the PartnerBankAccountAttributesBankName field if non-nil, zero value otherwise.

### GetPartnerBankAccountAttributesBankNameOk

`func (o *PartnerResponsePartner) GetPartnerBankAccountAttributesBankNameOk() (*string, bool)`

GetPartnerBankAccountAttributesBankNameOk returns a tuple with the PartnerBankAccountAttributesBankName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerBankAccountAttributesBankName

`func (o *PartnerResponsePartner) SetPartnerBankAccountAttributesBankName(v string)`

SetPartnerBankAccountAttributesBankName sets PartnerBankAccountAttributesBankName field to given value.

### HasPartnerBankAccountAttributesBankName

`func (o *PartnerResponsePartner) HasPartnerBankAccountAttributesBankName() bool`

HasPartnerBankAccountAttributesBankName returns a boolean if a field has been set.

### SetPartnerBankAccountAttributesBankNameNil

`func (o *PartnerResponsePartner) SetPartnerBankAccountAttributesBankNameNil(b bool)`

 SetPartnerBankAccountAttributesBankNameNil sets the value for PartnerBankAccountAttributesBankName to be an explicit nil

### UnsetPartnerBankAccountAttributesBankName
`func (o *PartnerResponsePartner) UnsetPartnerBankAccountAttributesBankName()`

UnsetPartnerBankAccountAttributesBankName ensures that no value is present for PartnerBankAccountAttributesBankName, not even an explicit nil
### GetPartnerBankAccountAttributesBankNameKana

`func (o *PartnerResponsePartner) GetPartnerBankAccountAttributesBankNameKana() string`

GetPartnerBankAccountAttributesBankNameKana returns the PartnerBankAccountAttributesBankNameKana field if non-nil, zero value otherwise.

### GetPartnerBankAccountAttributesBankNameKanaOk

`func (o *PartnerResponsePartner) GetPartnerBankAccountAttributesBankNameKanaOk() (*string, bool)`

GetPartnerBankAccountAttributesBankNameKanaOk returns a tuple with the PartnerBankAccountAttributesBankNameKana field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerBankAccountAttributesBankNameKana

`func (o *PartnerResponsePartner) SetPartnerBankAccountAttributesBankNameKana(v string)`

SetPartnerBankAccountAttributesBankNameKana sets PartnerBankAccountAttributesBankNameKana field to given value.

### HasPartnerBankAccountAttributesBankNameKana

`func (o *PartnerResponsePartner) HasPartnerBankAccountAttributesBankNameKana() bool`

HasPartnerBankAccountAttributesBankNameKana returns a boolean if a field has been set.

### SetPartnerBankAccountAttributesBankNameKanaNil

`func (o *PartnerResponsePartner) SetPartnerBankAccountAttributesBankNameKanaNil(b bool)`

 SetPartnerBankAccountAttributesBankNameKanaNil sets the value for PartnerBankAccountAttributesBankNameKana to be an explicit nil

### UnsetPartnerBankAccountAttributesBankNameKana
`func (o *PartnerResponsePartner) UnsetPartnerBankAccountAttributesBankNameKana()`

UnsetPartnerBankAccountAttributesBankNameKana ensures that no value is present for PartnerBankAccountAttributesBankNameKana, not even an explicit nil
### GetPartnerBankAccountAttributesBankCode

`func (o *PartnerResponsePartner) GetPartnerBankAccountAttributesBankCode() string`

GetPartnerBankAccountAttributesBankCode returns the PartnerBankAccountAttributesBankCode field if non-nil, zero value otherwise.

### GetPartnerBankAccountAttributesBankCodeOk

`func (o *PartnerResponsePartner) GetPartnerBankAccountAttributesBankCodeOk() (*string, bool)`

GetPartnerBankAccountAttributesBankCodeOk returns a tuple with the PartnerBankAccountAttributesBankCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerBankAccountAttributesBankCode

`func (o *PartnerResponsePartner) SetPartnerBankAccountAttributesBankCode(v string)`

SetPartnerBankAccountAttributesBankCode sets PartnerBankAccountAttributesBankCode field to given value.

### HasPartnerBankAccountAttributesBankCode

`func (o *PartnerResponsePartner) HasPartnerBankAccountAttributesBankCode() bool`

HasPartnerBankAccountAttributesBankCode returns a boolean if a field has been set.

### SetPartnerBankAccountAttributesBankCodeNil

`func (o *PartnerResponsePartner) SetPartnerBankAccountAttributesBankCodeNil(b bool)`

 SetPartnerBankAccountAttributesBankCodeNil sets the value for PartnerBankAccountAttributesBankCode to be an explicit nil

### UnsetPartnerBankAccountAttributesBankCode
`func (o *PartnerResponsePartner) UnsetPartnerBankAccountAttributesBankCode()`

UnsetPartnerBankAccountAttributesBankCode ensures that no value is present for PartnerBankAccountAttributesBankCode, not even an explicit nil
### GetPartnerBankAccountAttributesBranchName

`func (o *PartnerResponsePartner) GetPartnerBankAccountAttributesBranchName() string`

GetPartnerBankAccountAttributesBranchName returns the PartnerBankAccountAttributesBranchName field if non-nil, zero value otherwise.

### GetPartnerBankAccountAttributesBranchNameOk

`func (o *PartnerResponsePartner) GetPartnerBankAccountAttributesBranchNameOk() (*string, bool)`

GetPartnerBankAccountAttributesBranchNameOk returns a tuple with the PartnerBankAccountAttributesBranchName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerBankAccountAttributesBranchName

`func (o *PartnerResponsePartner) SetPartnerBankAccountAttributesBranchName(v string)`

SetPartnerBankAccountAttributesBranchName sets PartnerBankAccountAttributesBranchName field to given value.

### HasPartnerBankAccountAttributesBranchName

`func (o *PartnerResponsePartner) HasPartnerBankAccountAttributesBranchName() bool`

HasPartnerBankAccountAttributesBranchName returns a boolean if a field has been set.

### SetPartnerBankAccountAttributesBranchNameNil

`func (o *PartnerResponsePartner) SetPartnerBankAccountAttributesBranchNameNil(b bool)`

 SetPartnerBankAccountAttributesBranchNameNil sets the value for PartnerBankAccountAttributesBranchName to be an explicit nil

### UnsetPartnerBankAccountAttributesBranchName
`func (o *PartnerResponsePartner) UnsetPartnerBankAccountAttributesBranchName()`

UnsetPartnerBankAccountAttributesBranchName ensures that no value is present for PartnerBankAccountAttributesBranchName, not even an explicit nil
### GetPartnerBankAccountAttributesBranchKana

`func (o *PartnerResponsePartner) GetPartnerBankAccountAttributesBranchKana() string`

GetPartnerBankAccountAttributesBranchKana returns the PartnerBankAccountAttributesBranchKana field if non-nil, zero value otherwise.

### GetPartnerBankAccountAttributesBranchKanaOk

`func (o *PartnerResponsePartner) GetPartnerBankAccountAttributesBranchKanaOk() (*string, bool)`

GetPartnerBankAccountAttributesBranchKanaOk returns a tuple with the PartnerBankAccountAttributesBranchKana field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerBankAccountAttributesBranchKana

`func (o *PartnerResponsePartner) SetPartnerBankAccountAttributesBranchKana(v string)`

SetPartnerBankAccountAttributesBranchKana sets PartnerBankAccountAttributesBranchKana field to given value.

### HasPartnerBankAccountAttributesBranchKana

`func (o *PartnerResponsePartner) HasPartnerBankAccountAttributesBranchKana() bool`

HasPartnerBankAccountAttributesBranchKana returns a boolean if a field has been set.

### SetPartnerBankAccountAttributesBranchKanaNil

`func (o *PartnerResponsePartner) SetPartnerBankAccountAttributesBranchKanaNil(b bool)`

 SetPartnerBankAccountAttributesBranchKanaNil sets the value for PartnerBankAccountAttributesBranchKana to be an explicit nil

### UnsetPartnerBankAccountAttributesBranchKana
`func (o *PartnerResponsePartner) UnsetPartnerBankAccountAttributesBranchKana()`

UnsetPartnerBankAccountAttributesBranchKana ensures that no value is present for PartnerBankAccountAttributesBranchKana, not even an explicit nil
### GetPartnerBankAccountAttributesBranchCode

`func (o *PartnerResponsePartner) GetPartnerBankAccountAttributesBranchCode() string`

GetPartnerBankAccountAttributesBranchCode returns the PartnerBankAccountAttributesBranchCode field if non-nil, zero value otherwise.

### GetPartnerBankAccountAttributesBranchCodeOk

`func (o *PartnerResponsePartner) GetPartnerBankAccountAttributesBranchCodeOk() (*string, bool)`

GetPartnerBankAccountAttributesBranchCodeOk returns a tuple with the PartnerBankAccountAttributesBranchCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerBankAccountAttributesBranchCode

`func (o *PartnerResponsePartner) SetPartnerBankAccountAttributesBranchCode(v string)`

SetPartnerBankAccountAttributesBranchCode sets PartnerBankAccountAttributesBranchCode field to given value.

### HasPartnerBankAccountAttributesBranchCode

`func (o *PartnerResponsePartner) HasPartnerBankAccountAttributesBranchCode() bool`

HasPartnerBankAccountAttributesBranchCode returns a boolean if a field has been set.

### SetPartnerBankAccountAttributesBranchCodeNil

`func (o *PartnerResponsePartner) SetPartnerBankAccountAttributesBranchCodeNil(b bool)`

 SetPartnerBankAccountAttributesBranchCodeNil sets the value for PartnerBankAccountAttributesBranchCode to be an explicit nil

### UnsetPartnerBankAccountAttributesBranchCode
`func (o *PartnerResponsePartner) UnsetPartnerBankAccountAttributesBranchCode()`

UnsetPartnerBankAccountAttributesBranchCode ensures that no value is present for PartnerBankAccountAttributesBranchCode, not even an explicit nil
### GetPartnerBankAccountAttributesAccountType

`func (o *PartnerResponsePartner) GetPartnerBankAccountAttributesAccountType() string`

GetPartnerBankAccountAttributesAccountType returns the PartnerBankAccountAttributesAccountType field if non-nil, zero value otherwise.

### GetPartnerBankAccountAttributesAccountTypeOk

`func (o *PartnerResponsePartner) GetPartnerBankAccountAttributesAccountTypeOk() (*string, bool)`

GetPartnerBankAccountAttributesAccountTypeOk returns a tuple with the PartnerBankAccountAttributesAccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerBankAccountAttributesAccountType

`func (o *PartnerResponsePartner) SetPartnerBankAccountAttributesAccountType(v string)`

SetPartnerBankAccountAttributesAccountType sets PartnerBankAccountAttributesAccountType field to given value.

### HasPartnerBankAccountAttributesAccountType

`func (o *PartnerResponsePartner) HasPartnerBankAccountAttributesAccountType() bool`

HasPartnerBankAccountAttributesAccountType returns a boolean if a field has been set.

### SetPartnerBankAccountAttributesAccountTypeNil

`func (o *PartnerResponsePartner) SetPartnerBankAccountAttributesAccountTypeNil(b bool)`

 SetPartnerBankAccountAttributesAccountTypeNil sets the value for PartnerBankAccountAttributesAccountType to be an explicit nil

### UnsetPartnerBankAccountAttributesAccountType
`func (o *PartnerResponsePartner) UnsetPartnerBankAccountAttributesAccountType()`

UnsetPartnerBankAccountAttributesAccountType ensures that no value is present for PartnerBankAccountAttributesAccountType, not even an explicit nil
### GetPartnerBankAccountAttributesAccountNumber

`func (o *PartnerResponsePartner) GetPartnerBankAccountAttributesAccountNumber() string`

GetPartnerBankAccountAttributesAccountNumber returns the PartnerBankAccountAttributesAccountNumber field if non-nil, zero value otherwise.

### GetPartnerBankAccountAttributesAccountNumberOk

`func (o *PartnerResponsePartner) GetPartnerBankAccountAttributesAccountNumberOk() (*string, bool)`

GetPartnerBankAccountAttributesAccountNumberOk returns a tuple with the PartnerBankAccountAttributesAccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerBankAccountAttributesAccountNumber

`func (o *PartnerResponsePartner) SetPartnerBankAccountAttributesAccountNumber(v string)`

SetPartnerBankAccountAttributesAccountNumber sets PartnerBankAccountAttributesAccountNumber field to given value.

### HasPartnerBankAccountAttributesAccountNumber

`func (o *PartnerResponsePartner) HasPartnerBankAccountAttributesAccountNumber() bool`

HasPartnerBankAccountAttributesAccountNumber returns a boolean if a field has been set.

### SetPartnerBankAccountAttributesAccountNumberNil

`func (o *PartnerResponsePartner) SetPartnerBankAccountAttributesAccountNumberNil(b bool)`

 SetPartnerBankAccountAttributesAccountNumberNil sets the value for PartnerBankAccountAttributesAccountNumber to be an explicit nil

### UnsetPartnerBankAccountAttributesAccountNumber
`func (o *PartnerResponsePartner) UnsetPartnerBankAccountAttributesAccountNumber()`

UnsetPartnerBankAccountAttributesAccountNumber ensures that no value is present for PartnerBankAccountAttributesAccountNumber, not even an explicit nil
### GetPartnerBankAccountAttributesAccountName

`func (o *PartnerResponsePartner) GetPartnerBankAccountAttributesAccountName() string`

GetPartnerBankAccountAttributesAccountName returns the PartnerBankAccountAttributesAccountName field if non-nil, zero value otherwise.

### GetPartnerBankAccountAttributesAccountNameOk

`func (o *PartnerResponsePartner) GetPartnerBankAccountAttributesAccountNameOk() (*string, bool)`

GetPartnerBankAccountAttributesAccountNameOk returns a tuple with the PartnerBankAccountAttributesAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerBankAccountAttributesAccountName

`func (o *PartnerResponsePartner) SetPartnerBankAccountAttributesAccountName(v string)`

SetPartnerBankAccountAttributesAccountName sets PartnerBankAccountAttributesAccountName field to given value.

### HasPartnerBankAccountAttributesAccountName

`func (o *PartnerResponsePartner) HasPartnerBankAccountAttributesAccountName() bool`

HasPartnerBankAccountAttributesAccountName returns a boolean if a field has been set.

### SetPartnerBankAccountAttributesAccountNameNil

`func (o *PartnerResponsePartner) SetPartnerBankAccountAttributesAccountNameNil(b bool)`

 SetPartnerBankAccountAttributesAccountNameNil sets the value for PartnerBankAccountAttributesAccountName to be an explicit nil

### UnsetPartnerBankAccountAttributesAccountName
`func (o *PartnerResponsePartner) UnsetPartnerBankAccountAttributesAccountName()`

UnsetPartnerBankAccountAttributesAccountName ensures that no value is present for PartnerBankAccountAttributesAccountName, not even an explicit nil
### GetPartnerBankAccountAttributesLongAccountName

`func (o *PartnerResponsePartner) GetPartnerBankAccountAttributesLongAccountName() string`

GetPartnerBankAccountAttributesLongAccountName returns the PartnerBankAccountAttributesLongAccountName field if non-nil, zero value otherwise.

### GetPartnerBankAccountAttributesLongAccountNameOk

`func (o *PartnerResponsePartner) GetPartnerBankAccountAttributesLongAccountNameOk() (*string, bool)`

GetPartnerBankAccountAttributesLongAccountNameOk returns a tuple with the PartnerBankAccountAttributesLongAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerBankAccountAttributesLongAccountName

`func (o *PartnerResponsePartner) SetPartnerBankAccountAttributesLongAccountName(v string)`

SetPartnerBankAccountAttributesLongAccountName sets PartnerBankAccountAttributesLongAccountName field to given value.

### HasPartnerBankAccountAttributesLongAccountName

`func (o *PartnerResponsePartner) HasPartnerBankAccountAttributesLongAccountName() bool`

HasPartnerBankAccountAttributesLongAccountName returns a boolean if a field has been set.

### SetPartnerBankAccountAttributesLongAccountNameNil

`func (o *PartnerResponsePartner) SetPartnerBankAccountAttributesLongAccountNameNil(b bool)`

 SetPartnerBankAccountAttributesLongAccountNameNil sets the value for PartnerBankAccountAttributesLongAccountName to be an explicit nil

### UnsetPartnerBankAccountAttributesLongAccountName
`func (o *PartnerResponsePartner) UnsetPartnerBankAccountAttributesLongAccountName()`

UnsetPartnerBankAccountAttributesLongAccountName ensures that no value is present for PartnerBankAccountAttributesLongAccountName, not even an explicit nil
### GetPaymentTermAttributesCutoffDay

`func (o *PartnerResponsePartner) GetPaymentTermAttributesCutoffDay() int32`

GetPaymentTermAttributesCutoffDay returns the PaymentTermAttributesCutoffDay field if non-nil, zero value otherwise.

### GetPaymentTermAttributesCutoffDayOk

`func (o *PartnerResponsePartner) GetPaymentTermAttributesCutoffDayOk() (*int32, bool)`

GetPaymentTermAttributesCutoffDayOk returns a tuple with the PaymentTermAttributesCutoffDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentTermAttributesCutoffDay

`func (o *PartnerResponsePartner) SetPaymentTermAttributesCutoffDay(v int32)`

SetPaymentTermAttributesCutoffDay sets PaymentTermAttributesCutoffDay field to given value.

### HasPaymentTermAttributesCutoffDay

`func (o *PartnerResponsePartner) HasPaymentTermAttributesCutoffDay() bool`

HasPaymentTermAttributesCutoffDay returns a boolean if a field has been set.

### SetPaymentTermAttributesCutoffDayNil

`func (o *PartnerResponsePartner) SetPaymentTermAttributesCutoffDayNil(b bool)`

 SetPaymentTermAttributesCutoffDayNil sets the value for PaymentTermAttributesCutoffDay to be an explicit nil

### UnsetPaymentTermAttributesCutoffDay
`func (o *PartnerResponsePartner) UnsetPaymentTermAttributesCutoffDay()`

UnsetPaymentTermAttributesCutoffDay ensures that no value is present for PaymentTermAttributesCutoffDay, not even an explicit nil
### GetPaymentTermAttributesAdditionalMonths

`func (o *PartnerResponsePartner) GetPaymentTermAttributesAdditionalMonths() int32`

GetPaymentTermAttributesAdditionalMonths returns the PaymentTermAttributesAdditionalMonths field if non-nil, zero value otherwise.

### GetPaymentTermAttributesAdditionalMonthsOk

`func (o *PartnerResponsePartner) GetPaymentTermAttributesAdditionalMonthsOk() (*int32, bool)`

GetPaymentTermAttributesAdditionalMonthsOk returns a tuple with the PaymentTermAttributesAdditionalMonths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentTermAttributesAdditionalMonths

`func (o *PartnerResponsePartner) SetPaymentTermAttributesAdditionalMonths(v int32)`

SetPaymentTermAttributesAdditionalMonths sets PaymentTermAttributesAdditionalMonths field to given value.

### HasPaymentTermAttributesAdditionalMonths

`func (o *PartnerResponsePartner) HasPaymentTermAttributesAdditionalMonths() bool`

HasPaymentTermAttributesAdditionalMonths returns a boolean if a field has been set.

### SetPaymentTermAttributesAdditionalMonthsNil

`func (o *PartnerResponsePartner) SetPaymentTermAttributesAdditionalMonthsNil(b bool)`

 SetPaymentTermAttributesAdditionalMonthsNil sets the value for PaymentTermAttributesAdditionalMonths to be an explicit nil

### UnsetPaymentTermAttributesAdditionalMonths
`func (o *PartnerResponsePartner) UnsetPaymentTermAttributesAdditionalMonths()`

UnsetPaymentTermAttributesAdditionalMonths ensures that no value is present for PaymentTermAttributesAdditionalMonths, not even an explicit nil
### GetPaymentTermAttributesFixedDay

`func (o *PartnerResponsePartner) GetPaymentTermAttributesFixedDay() int32`

GetPaymentTermAttributesFixedDay returns the PaymentTermAttributesFixedDay field if non-nil, zero value otherwise.

### GetPaymentTermAttributesFixedDayOk

`func (o *PartnerResponsePartner) GetPaymentTermAttributesFixedDayOk() (*int32, bool)`

GetPaymentTermAttributesFixedDayOk returns a tuple with the PaymentTermAttributesFixedDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentTermAttributesFixedDay

`func (o *PartnerResponsePartner) SetPaymentTermAttributesFixedDay(v int32)`

SetPaymentTermAttributesFixedDay sets PaymentTermAttributesFixedDay field to given value.

### HasPaymentTermAttributesFixedDay

`func (o *PartnerResponsePartner) HasPaymentTermAttributesFixedDay() bool`

HasPaymentTermAttributesFixedDay returns a boolean if a field has been set.

### SetPaymentTermAttributesFixedDayNil

`func (o *PartnerResponsePartner) SetPaymentTermAttributesFixedDayNil(b bool)`

 SetPaymentTermAttributesFixedDayNil sets the value for PaymentTermAttributesFixedDay to be an explicit nil

### UnsetPaymentTermAttributesFixedDay
`func (o *PartnerResponsePartner) UnsetPaymentTermAttributesFixedDay()`

UnsetPaymentTermAttributesFixedDay ensures that no value is present for PaymentTermAttributesFixedDay, not even an explicit nil
### GetInvoicePaymentTermAttributesCutoffDay

`func (o *PartnerResponsePartner) GetInvoicePaymentTermAttributesCutoffDay() int32`

GetInvoicePaymentTermAttributesCutoffDay returns the InvoicePaymentTermAttributesCutoffDay field if non-nil, zero value otherwise.

### GetInvoicePaymentTermAttributesCutoffDayOk

`func (o *PartnerResponsePartner) GetInvoicePaymentTermAttributesCutoffDayOk() (*int32, bool)`

GetInvoicePaymentTermAttributesCutoffDayOk returns a tuple with the InvoicePaymentTermAttributesCutoffDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoicePaymentTermAttributesCutoffDay

`func (o *PartnerResponsePartner) SetInvoicePaymentTermAttributesCutoffDay(v int32)`

SetInvoicePaymentTermAttributesCutoffDay sets InvoicePaymentTermAttributesCutoffDay field to given value.

### HasInvoicePaymentTermAttributesCutoffDay

`func (o *PartnerResponsePartner) HasInvoicePaymentTermAttributesCutoffDay() bool`

HasInvoicePaymentTermAttributesCutoffDay returns a boolean if a field has been set.

### SetInvoicePaymentTermAttributesCutoffDayNil

`func (o *PartnerResponsePartner) SetInvoicePaymentTermAttributesCutoffDayNil(b bool)`

 SetInvoicePaymentTermAttributesCutoffDayNil sets the value for InvoicePaymentTermAttributesCutoffDay to be an explicit nil

### UnsetInvoicePaymentTermAttributesCutoffDay
`func (o *PartnerResponsePartner) UnsetInvoicePaymentTermAttributesCutoffDay()`

UnsetInvoicePaymentTermAttributesCutoffDay ensures that no value is present for InvoicePaymentTermAttributesCutoffDay, not even an explicit nil
### GetInvoicePaymentTermAttributesAdditionalMonths

`func (o *PartnerResponsePartner) GetInvoicePaymentTermAttributesAdditionalMonths() int32`

GetInvoicePaymentTermAttributesAdditionalMonths returns the InvoicePaymentTermAttributesAdditionalMonths field if non-nil, zero value otherwise.

### GetInvoicePaymentTermAttributesAdditionalMonthsOk

`func (o *PartnerResponsePartner) GetInvoicePaymentTermAttributesAdditionalMonthsOk() (*int32, bool)`

GetInvoicePaymentTermAttributesAdditionalMonthsOk returns a tuple with the InvoicePaymentTermAttributesAdditionalMonths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoicePaymentTermAttributesAdditionalMonths

`func (o *PartnerResponsePartner) SetInvoicePaymentTermAttributesAdditionalMonths(v int32)`

SetInvoicePaymentTermAttributesAdditionalMonths sets InvoicePaymentTermAttributesAdditionalMonths field to given value.

### HasInvoicePaymentTermAttributesAdditionalMonths

`func (o *PartnerResponsePartner) HasInvoicePaymentTermAttributesAdditionalMonths() bool`

HasInvoicePaymentTermAttributesAdditionalMonths returns a boolean if a field has been set.

### SetInvoicePaymentTermAttributesAdditionalMonthsNil

`func (o *PartnerResponsePartner) SetInvoicePaymentTermAttributesAdditionalMonthsNil(b bool)`

 SetInvoicePaymentTermAttributesAdditionalMonthsNil sets the value for InvoicePaymentTermAttributesAdditionalMonths to be an explicit nil

### UnsetInvoicePaymentTermAttributesAdditionalMonths
`func (o *PartnerResponsePartner) UnsetInvoicePaymentTermAttributesAdditionalMonths()`

UnsetInvoicePaymentTermAttributesAdditionalMonths ensures that no value is present for InvoicePaymentTermAttributesAdditionalMonths, not even an explicit nil
### GetInvoicePaymentTermAttributesFixedDay

`func (o *PartnerResponsePartner) GetInvoicePaymentTermAttributesFixedDay() int32`

GetInvoicePaymentTermAttributesFixedDay returns the InvoicePaymentTermAttributesFixedDay field if non-nil, zero value otherwise.

### GetInvoicePaymentTermAttributesFixedDayOk

`func (o *PartnerResponsePartner) GetInvoicePaymentTermAttributesFixedDayOk() (*int32, bool)`

GetInvoicePaymentTermAttributesFixedDayOk returns a tuple with the InvoicePaymentTermAttributesFixedDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoicePaymentTermAttributesFixedDay

`func (o *PartnerResponsePartner) SetInvoicePaymentTermAttributesFixedDay(v int32)`

SetInvoicePaymentTermAttributesFixedDay sets InvoicePaymentTermAttributesFixedDay field to given value.

### HasInvoicePaymentTermAttributesFixedDay

`func (o *PartnerResponsePartner) HasInvoicePaymentTermAttributesFixedDay() bool`

HasInvoicePaymentTermAttributesFixedDay returns a boolean if a field has been set.

### SetInvoicePaymentTermAttributesFixedDayNil

`func (o *PartnerResponsePartner) SetInvoicePaymentTermAttributesFixedDayNil(b bool)`

 SetInvoicePaymentTermAttributesFixedDayNil sets the value for InvoicePaymentTermAttributesFixedDay to be an explicit nil

### UnsetInvoicePaymentTermAttributesFixedDay
`func (o *PartnerResponsePartner) UnsetInvoicePaymentTermAttributesFixedDay()`

UnsetInvoicePaymentTermAttributesFixedDay ensures that no value is present for InvoicePaymentTermAttributesFixedDay, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


