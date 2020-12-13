# PartnersResponsePartners

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
**PartnerDocSettingAttributesSendingMethod** | Pointer to **NullableString** | 請求書送付方法(mail:メール、posting:郵送、mail_and_posting:メールと郵送) | [optional] 
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

## Methods

### NewPartnersResponsePartners

`func NewPartnersResponsePartners(id int32, code NullableString, companyId int32, name string, ) *PartnersResponsePartners`

NewPartnersResponsePartners instantiates a new PartnersResponsePartners object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartnersResponsePartnersWithDefaults

`func NewPartnersResponsePartnersWithDefaults() *PartnersResponsePartners`

NewPartnersResponsePartnersWithDefaults instantiates a new PartnersResponsePartners object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PartnersResponsePartners) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PartnersResponsePartners) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PartnersResponsePartners) SetId(v int32)`

SetId sets Id field to given value.


### GetCode

`func (o *PartnersResponsePartners) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *PartnersResponsePartners) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *PartnersResponsePartners) SetCode(v string)`

SetCode sets Code field to given value.


### SetCodeNil

`func (o *PartnersResponsePartners) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *PartnersResponsePartners) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetCompanyId

`func (o *PartnersResponsePartners) GetCompanyId() int32`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *PartnersResponsePartners) GetCompanyIdOk() (*int32, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *PartnersResponsePartners) SetCompanyId(v int32)`

SetCompanyId sets CompanyId field to given value.


### GetName

`func (o *PartnersResponsePartners) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PartnersResponsePartners) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PartnersResponsePartners) SetName(v string)`

SetName sets Name field to given value.


### GetShortcut1

`func (o *PartnersResponsePartners) GetShortcut1() string`

GetShortcut1 returns the Shortcut1 field if non-nil, zero value otherwise.

### GetShortcut1Ok

`func (o *PartnersResponsePartners) GetShortcut1Ok() (*string, bool)`

GetShortcut1Ok returns a tuple with the Shortcut1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortcut1

`func (o *PartnersResponsePartners) SetShortcut1(v string)`

SetShortcut1 sets Shortcut1 field to given value.

### HasShortcut1

`func (o *PartnersResponsePartners) HasShortcut1() bool`

HasShortcut1 returns a boolean if a field has been set.

### SetShortcut1Nil

`func (o *PartnersResponsePartners) SetShortcut1Nil(b bool)`

 SetShortcut1Nil sets the value for Shortcut1 to be an explicit nil

### UnsetShortcut1
`func (o *PartnersResponsePartners) UnsetShortcut1()`

UnsetShortcut1 ensures that no value is present for Shortcut1, not even an explicit nil
### GetShortcut2

`func (o *PartnersResponsePartners) GetShortcut2() string`

GetShortcut2 returns the Shortcut2 field if non-nil, zero value otherwise.

### GetShortcut2Ok

`func (o *PartnersResponsePartners) GetShortcut2Ok() (*string, bool)`

GetShortcut2Ok returns a tuple with the Shortcut2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortcut2

`func (o *PartnersResponsePartners) SetShortcut2(v string)`

SetShortcut2 sets Shortcut2 field to given value.

### HasShortcut2

`func (o *PartnersResponsePartners) HasShortcut2() bool`

HasShortcut2 returns a boolean if a field has been set.

### SetShortcut2Nil

`func (o *PartnersResponsePartners) SetShortcut2Nil(b bool)`

 SetShortcut2Nil sets the value for Shortcut2 to be an explicit nil

### UnsetShortcut2
`func (o *PartnersResponsePartners) UnsetShortcut2()`

UnsetShortcut2 ensures that no value is present for Shortcut2, not even an explicit nil
### GetOrgCode

`func (o *PartnersResponsePartners) GetOrgCode() int32`

GetOrgCode returns the OrgCode field if non-nil, zero value otherwise.

### GetOrgCodeOk

`func (o *PartnersResponsePartners) GetOrgCodeOk() (*int32, bool)`

GetOrgCodeOk returns a tuple with the OrgCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgCode

`func (o *PartnersResponsePartners) SetOrgCode(v int32)`

SetOrgCode sets OrgCode field to given value.

### HasOrgCode

`func (o *PartnersResponsePartners) HasOrgCode() bool`

HasOrgCode returns a boolean if a field has been set.

### SetOrgCodeNil

`func (o *PartnersResponsePartners) SetOrgCodeNil(b bool)`

 SetOrgCodeNil sets the value for OrgCode to be an explicit nil

### UnsetOrgCode
`func (o *PartnersResponsePartners) UnsetOrgCode()`

UnsetOrgCode ensures that no value is present for OrgCode, not even an explicit nil
### GetCountryCode

`func (o *PartnersResponsePartners) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *PartnersResponsePartners) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *PartnersResponsePartners) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *PartnersResponsePartners) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetLongName

`func (o *PartnersResponsePartners) GetLongName() string`

GetLongName returns the LongName field if non-nil, zero value otherwise.

### GetLongNameOk

`func (o *PartnersResponsePartners) GetLongNameOk() (*string, bool)`

GetLongNameOk returns a tuple with the LongName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongName

`func (o *PartnersResponsePartners) SetLongName(v string)`

SetLongName sets LongName field to given value.

### HasLongName

`func (o *PartnersResponsePartners) HasLongName() bool`

HasLongName returns a boolean if a field has been set.

### SetLongNameNil

`func (o *PartnersResponsePartners) SetLongNameNil(b bool)`

 SetLongNameNil sets the value for LongName to be an explicit nil

### UnsetLongName
`func (o *PartnersResponsePartners) UnsetLongName()`

UnsetLongName ensures that no value is present for LongName, not even an explicit nil
### GetNameKana

`func (o *PartnersResponsePartners) GetNameKana() string`

GetNameKana returns the NameKana field if non-nil, zero value otherwise.

### GetNameKanaOk

`func (o *PartnersResponsePartners) GetNameKanaOk() (*string, bool)`

GetNameKanaOk returns a tuple with the NameKana field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameKana

`func (o *PartnersResponsePartners) SetNameKana(v string)`

SetNameKana sets NameKana field to given value.

### HasNameKana

`func (o *PartnersResponsePartners) HasNameKana() bool`

HasNameKana returns a boolean if a field has been set.

### SetNameKanaNil

`func (o *PartnersResponsePartners) SetNameKanaNil(b bool)`

 SetNameKanaNil sets the value for NameKana to be an explicit nil

### UnsetNameKana
`func (o *PartnersResponsePartners) UnsetNameKana()`

UnsetNameKana ensures that no value is present for NameKana, not even an explicit nil
### GetDefaultTitle

`func (o *PartnersResponsePartners) GetDefaultTitle() string`

GetDefaultTitle returns the DefaultTitle field if non-nil, zero value otherwise.

### GetDefaultTitleOk

`func (o *PartnersResponsePartners) GetDefaultTitleOk() (*string, bool)`

GetDefaultTitleOk returns a tuple with the DefaultTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTitle

`func (o *PartnersResponsePartners) SetDefaultTitle(v string)`

SetDefaultTitle sets DefaultTitle field to given value.

### HasDefaultTitle

`func (o *PartnersResponsePartners) HasDefaultTitle() bool`

HasDefaultTitle returns a boolean if a field has been set.

### SetDefaultTitleNil

`func (o *PartnersResponsePartners) SetDefaultTitleNil(b bool)`

 SetDefaultTitleNil sets the value for DefaultTitle to be an explicit nil

### UnsetDefaultTitle
`func (o *PartnersResponsePartners) UnsetDefaultTitle()`

UnsetDefaultTitle ensures that no value is present for DefaultTitle, not even an explicit nil
### GetPhone

`func (o *PartnersResponsePartners) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *PartnersResponsePartners) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *PartnersResponsePartners) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *PartnersResponsePartners) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### SetPhoneNil

`func (o *PartnersResponsePartners) SetPhoneNil(b bool)`

 SetPhoneNil sets the value for Phone to be an explicit nil

### UnsetPhone
`func (o *PartnersResponsePartners) UnsetPhone()`

UnsetPhone ensures that no value is present for Phone, not even an explicit nil
### GetContactName

`func (o *PartnersResponsePartners) GetContactName() string`

GetContactName returns the ContactName field if non-nil, zero value otherwise.

### GetContactNameOk

`func (o *PartnersResponsePartners) GetContactNameOk() (*string, bool)`

GetContactNameOk returns a tuple with the ContactName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactName

`func (o *PartnersResponsePartners) SetContactName(v string)`

SetContactName sets ContactName field to given value.

### HasContactName

`func (o *PartnersResponsePartners) HasContactName() bool`

HasContactName returns a boolean if a field has been set.

### SetContactNameNil

`func (o *PartnersResponsePartners) SetContactNameNil(b bool)`

 SetContactNameNil sets the value for ContactName to be an explicit nil

### UnsetContactName
`func (o *PartnersResponsePartners) UnsetContactName()`

UnsetContactName ensures that no value is present for ContactName, not even an explicit nil
### GetEmail

`func (o *PartnersResponsePartners) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *PartnersResponsePartners) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *PartnersResponsePartners) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *PartnersResponsePartners) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *PartnersResponsePartners) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *PartnersResponsePartners) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetPayerWalletableId

`func (o *PartnersResponsePartners) GetPayerWalletableId() int32`

GetPayerWalletableId returns the PayerWalletableId field if non-nil, zero value otherwise.

### GetPayerWalletableIdOk

`func (o *PartnersResponsePartners) GetPayerWalletableIdOk() (*int32, bool)`

GetPayerWalletableIdOk returns a tuple with the PayerWalletableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayerWalletableId

`func (o *PartnersResponsePartners) SetPayerWalletableId(v int32)`

SetPayerWalletableId sets PayerWalletableId field to given value.

### HasPayerWalletableId

`func (o *PartnersResponsePartners) HasPayerWalletableId() bool`

HasPayerWalletableId returns a boolean if a field has been set.

### SetPayerWalletableIdNil

`func (o *PartnersResponsePartners) SetPayerWalletableIdNil(b bool)`

 SetPayerWalletableIdNil sets the value for PayerWalletableId to be an explicit nil

### UnsetPayerWalletableId
`func (o *PartnersResponsePartners) UnsetPayerWalletableId()`

UnsetPayerWalletableId ensures that no value is present for PayerWalletableId, not even an explicit nil
### GetTransferFeeHandlingSide

`func (o *PartnersResponsePartners) GetTransferFeeHandlingSide() string`

GetTransferFeeHandlingSide returns the TransferFeeHandlingSide field if non-nil, zero value otherwise.

### GetTransferFeeHandlingSideOk

`func (o *PartnersResponsePartners) GetTransferFeeHandlingSideOk() (*string, bool)`

GetTransferFeeHandlingSideOk returns a tuple with the TransferFeeHandlingSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferFeeHandlingSide

`func (o *PartnersResponsePartners) SetTransferFeeHandlingSide(v string)`

SetTransferFeeHandlingSide sets TransferFeeHandlingSide field to given value.

### HasTransferFeeHandlingSide

`func (o *PartnersResponsePartners) HasTransferFeeHandlingSide() bool`

HasTransferFeeHandlingSide returns a boolean if a field has been set.

### GetAddressAttributesZipcode

`func (o *PartnersResponsePartners) GetAddressAttributesZipcode() string`

GetAddressAttributesZipcode returns the AddressAttributesZipcode field if non-nil, zero value otherwise.

### GetAddressAttributesZipcodeOk

`func (o *PartnersResponsePartners) GetAddressAttributesZipcodeOk() (*string, bool)`

GetAddressAttributesZipcodeOk returns a tuple with the AddressAttributesZipcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressAttributesZipcode

`func (o *PartnersResponsePartners) SetAddressAttributesZipcode(v string)`

SetAddressAttributesZipcode sets AddressAttributesZipcode field to given value.

### HasAddressAttributesZipcode

`func (o *PartnersResponsePartners) HasAddressAttributesZipcode() bool`

HasAddressAttributesZipcode returns a boolean if a field has been set.

### SetAddressAttributesZipcodeNil

`func (o *PartnersResponsePartners) SetAddressAttributesZipcodeNil(b bool)`

 SetAddressAttributesZipcodeNil sets the value for AddressAttributesZipcode to be an explicit nil

### UnsetAddressAttributesZipcode
`func (o *PartnersResponsePartners) UnsetAddressAttributesZipcode()`

UnsetAddressAttributesZipcode ensures that no value is present for AddressAttributesZipcode, not even an explicit nil
### GetAddressAttributesPrefectureCode

`func (o *PartnersResponsePartners) GetAddressAttributesPrefectureCode() int32`

GetAddressAttributesPrefectureCode returns the AddressAttributesPrefectureCode field if non-nil, zero value otherwise.

### GetAddressAttributesPrefectureCodeOk

`func (o *PartnersResponsePartners) GetAddressAttributesPrefectureCodeOk() (*int32, bool)`

GetAddressAttributesPrefectureCodeOk returns a tuple with the AddressAttributesPrefectureCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressAttributesPrefectureCode

`func (o *PartnersResponsePartners) SetAddressAttributesPrefectureCode(v int32)`

SetAddressAttributesPrefectureCode sets AddressAttributesPrefectureCode field to given value.

### HasAddressAttributesPrefectureCode

`func (o *PartnersResponsePartners) HasAddressAttributesPrefectureCode() bool`

HasAddressAttributesPrefectureCode returns a boolean if a field has been set.

### GetAddressAttributesStreetName1

`func (o *PartnersResponsePartners) GetAddressAttributesStreetName1() string`

GetAddressAttributesStreetName1 returns the AddressAttributesStreetName1 field if non-nil, zero value otherwise.

### GetAddressAttributesStreetName1Ok

`func (o *PartnersResponsePartners) GetAddressAttributesStreetName1Ok() (*string, bool)`

GetAddressAttributesStreetName1Ok returns a tuple with the AddressAttributesStreetName1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressAttributesStreetName1

`func (o *PartnersResponsePartners) SetAddressAttributesStreetName1(v string)`

SetAddressAttributesStreetName1 sets AddressAttributesStreetName1 field to given value.

### HasAddressAttributesStreetName1

`func (o *PartnersResponsePartners) HasAddressAttributesStreetName1() bool`

HasAddressAttributesStreetName1 returns a boolean if a field has been set.

### SetAddressAttributesStreetName1Nil

`func (o *PartnersResponsePartners) SetAddressAttributesStreetName1Nil(b bool)`

 SetAddressAttributesStreetName1Nil sets the value for AddressAttributesStreetName1 to be an explicit nil

### UnsetAddressAttributesStreetName1
`func (o *PartnersResponsePartners) UnsetAddressAttributesStreetName1()`

UnsetAddressAttributesStreetName1 ensures that no value is present for AddressAttributesStreetName1, not even an explicit nil
### GetAddressAttributesStreetName2

`func (o *PartnersResponsePartners) GetAddressAttributesStreetName2() string`

GetAddressAttributesStreetName2 returns the AddressAttributesStreetName2 field if non-nil, zero value otherwise.

### GetAddressAttributesStreetName2Ok

`func (o *PartnersResponsePartners) GetAddressAttributesStreetName2Ok() (*string, bool)`

GetAddressAttributesStreetName2Ok returns a tuple with the AddressAttributesStreetName2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressAttributesStreetName2

`func (o *PartnersResponsePartners) SetAddressAttributesStreetName2(v string)`

SetAddressAttributesStreetName2 sets AddressAttributesStreetName2 field to given value.

### HasAddressAttributesStreetName2

`func (o *PartnersResponsePartners) HasAddressAttributesStreetName2() bool`

HasAddressAttributesStreetName2 returns a boolean if a field has been set.

### SetAddressAttributesStreetName2Nil

`func (o *PartnersResponsePartners) SetAddressAttributesStreetName2Nil(b bool)`

 SetAddressAttributesStreetName2Nil sets the value for AddressAttributesStreetName2 to be an explicit nil

### UnsetAddressAttributesStreetName2
`func (o *PartnersResponsePartners) UnsetAddressAttributesStreetName2()`

UnsetAddressAttributesStreetName2 ensures that no value is present for AddressAttributesStreetName2, not even an explicit nil
### GetPartnerDocSettingAttributesSendingMethod

`func (o *PartnersResponsePartners) GetPartnerDocSettingAttributesSendingMethod() string`

GetPartnerDocSettingAttributesSendingMethod returns the PartnerDocSettingAttributesSendingMethod field if non-nil, zero value otherwise.

### GetPartnerDocSettingAttributesSendingMethodOk

`func (o *PartnersResponsePartners) GetPartnerDocSettingAttributesSendingMethodOk() (*string, bool)`

GetPartnerDocSettingAttributesSendingMethodOk returns a tuple with the PartnerDocSettingAttributesSendingMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerDocSettingAttributesSendingMethod

`func (o *PartnersResponsePartners) SetPartnerDocSettingAttributesSendingMethod(v string)`

SetPartnerDocSettingAttributesSendingMethod sets PartnerDocSettingAttributesSendingMethod field to given value.

### HasPartnerDocSettingAttributesSendingMethod

`func (o *PartnersResponsePartners) HasPartnerDocSettingAttributesSendingMethod() bool`

HasPartnerDocSettingAttributesSendingMethod returns a boolean if a field has been set.

### SetPartnerDocSettingAttributesSendingMethodNil

`func (o *PartnersResponsePartners) SetPartnerDocSettingAttributesSendingMethodNil(b bool)`

 SetPartnerDocSettingAttributesSendingMethodNil sets the value for PartnerDocSettingAttributesSendingMethod to be an explicit nil

### UnsetPartnerDocSettingAttributesSendingMethod
`func (o *PartnersResponsePartners) UnsetPartnerDocSettingAttributesSendingMethod()`

UnsetPartnerDocSettingAttributesSendingMethod ensures that no value is present for PartnerDocSettingAttributesSendingMethod, not even an explicit nil
### GetPartnerBankAccountAttributesBankName

`func (o *PartnersResponsePartners) GetPartnerBankAccountAttributesBankName() string`

GetPartnerBankAccountAttributesBankName returns the PartnerBankAccountAttributesBankName field if non-nil, zero value otherwise.

### GetPartnerBankAccountAttributesBankNameOk

`func (o *PartnersResponsePartners) GetPartnerBankAccountAttributesBankNameOk() (*string, bool)`

GetPartnerBankAccountAttributesBankNameOk returns a tuple with the PartnerBankAccountAttributesBankName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerBankAccountAttributesBankName

`func (o *PartnersResponsePartners) SetPartnerBankAccountAttributesBankName(v string)`

SetPartnerBankAccountAttributesBankName sets PartnerBankAccountAttributesBankName field to given value.

### HasPartnerBankAccountAttributesBankName

`func (o *PartnersResponsePartners) HasPartnerBankAccountAttributesBankName() bool`

HasPartnerBankAccountAttributesBankName returns a boolean if a field has been set.

### SetPartnerBankAccountAttributesBankNameNil

`func (o *PartnersResponsePartners) SetPartnerBankAccountAttributesBankNameNil(b bool)`

 SetPartnerBankAccountAttributesBankNameNil sets the value for PartnerBankAccountAttributesBankName to be an explicit nil

### UnsetPartnerBankAccountAttributesBankName
`func (o *PartnersResponsePartners) UnsetPartnerBankAccountAttributesBankName()`

UnsetPartnerBankAccountAttributesBankName ensures that no value is present for PartnerBankAccountAttributesBankName, not even an explicit nil
### GetPartnerBankAccountAttributesBankNameKana

`func (o *PartnersResponsePartners) GetPartnerBankAccountAttributesBankNameKana() string`

GetPartnerBankAccountAttributesBankNameKana returns the PartnerBankAccountAttributesBankNameKana field if non-nil, zero value otherwise.

### GetPartnerBankAccountAttributesBankNameKanaOk

`func (o *PartnersResponsePartners) GetPartnerBankAccountAttributesBankNameKanaOk() (*string, bool)`

GetPartnerBankAccountAttributesBankNameKanaOk returns a tuple with the PartnerBankAccountAttributesBankNameKana field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerBankAccountAttributesBankNameKana

`func (o *PartnersResponsePartners) SetPartnerBankAccountAttributesBankNameKana(v string)`

SetPartnerBankAccountAttributesBankNameKana sets PartnerBankAccountAttributesBankNameKana field to given value.

### HasPartnerBankAccountAttributesBankNameKana

`func (o *PartnersResponsePartners) HasPartnerBankAccountAttributesBankNameKana() bool`

HasPartnerBankAccountAttributesBankNameKana returns a boolean if a field has been set.

### SetPartnerBankAccountAttributesBankNameKanaNil

`func (o *PartnersResponsePartners) SetPartnerBankAccountAttributesBankNameKanaNil(b bool)`

 SetPartnerBankAccountAttributesBankNameKanaNil sets the value for PartnerBankAccountAttributesBankNameKana to be an explicit nil

### UnsetPartnerBankAccountAttributesBankNameKana
`func (o *PartnersResponsePartners) UnsetPartnerBankAccountAttributesBankNameKana()`

UnsetPartnerBankAccountAttributesBankNameKana ensures that no value is present for PartnerBankAccountAttributesBankNameKana, not even an explicit nil
### GetPartnerBankAccountAttributesBankCode

`func (o *PartnersResponsePartners) GetPartnerBankAccountAttributesBankCode() string`

GetPartnerBankAccountAttributesBankCode returns the PartnerBankAccountAttributesBankCode field if non-nil, zero value otherwise.

### GetPartnerBankAccountAttributesBankCodeOk

`func (o *PartnersResponsePartners) GetPartnerBankAccountAttributesBankCodeOk() (*string, bool)`

GetPartnerBankAccountAttributesBankCodeOk returns a tuple with the PartnerBankAccountAttributesBankCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerBankAccountAttributesBankCode

`func (o *PartnersResponsePartners) SetPartnerBankAccountAttributesBankCode(v string)`

SetPartnerBankAccountAttributesBankCode sets PartnerBankAccountAttributesBankCode field to given value.

### HasPartnerBankAccountAttributesBankCode

`func (o *PartnersResponsePartners) HasPartnerBankAccountAttributesBankCode() bool`

HasPartnerBankAccountAttributesBankCode returns a boolean if a field has been set.

### SetPartnerBankAccountAttributesBankCodeNil

`func (o *PartnersResponsePartners) SetPartnerBankAccountAttributesBankCodeNil(b bool)`

 SetPartnerBankAccountAttributesBankCodeNil sets the value for PartnerBankAccountAttributesBankCode to be an explicit nil

### UnsetPartnerBankAccountAttributesBankCode
`func (o *PartnersResponsePartners) UnsetPartnerBankAccountAttributesBankCode()`

UnsetPartnerBankAccountAttributesBankCode ensures that no value is present for PartnerBankAccountAttributesBankCode, not even an explicit nil
### GetPartnerBankAccountAttributesBranchName

`func (o *PartnersResponsePartners) GetPartnerBankAccountAttributesBranchName() string`

GetPartnerBankAccountAttributesBranchName returns the PartnerBankAccountAttributesBranchName field if non-nil, zero value otherwise.

### GetPartnerBankAccountAttributesBranchNameOk

`func (o *PartnersResponsePartners) GetPartnerBankAccountAttributesBranchNameOk() (*string, bool)`

GetPartnerBankAccountAttributesBranchNameOk returns a tuple with the PartnerBankAccountAttributesBranchName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerBankAccountAttributesBranchName

`func (o *PartnersResponsePartners) SetPartnerBankAccountAttributesBranchName(v string)`

SetPartnerBankAccountAttributesBranchName sets PartnerBankAccountAttributesBranchName field to given value.

### HasPartnerBankAccountAttributesBranchName

`func (o *PartnersResponsePartners) HasPartnerBankAccountAttributesBranchName() bool`

HasPartnerBankAccountAttributesBranchName returns a boolean if a field has been set.

### SetPartnerBankAccountAttributesBranchNameNil

`func (o *PartnersResponsePartners) SetPartnerBankAccountAttributesBranchNameNil(b bool)`

 SetPartnerBankAccountAttributesBranchNameNil sets the value for PartnerBankAccountAttributesBranchName to be an explicit nil

### UnsetPartnerBankAccountAttributesBranchName
`func (o *PartnersResponsePartners) UnsetPartnerBankAccountAttributesBranchName()`

UnsetPartnerBankAccountAttributesBranchName ensures that no value is present for PartnerBankAccountAttributesBranchName, not even an explicit nil
### GetPartnerBankAccountAttributesBranchKana

`func (o *PartnersResponsePartners) GetPartnerBankAccountAttributesBranchKana() string`

GetPartnerBankAccountAttributesBranchKana returns the PartnerBankAccountAttributesBranchKana field if non-nil, zero value otherwise.

### GetPartnerBankAccountAttributesBranchKanaOk

`func (o *PartnersResponsePartners) GetPartnerBankAccountAttributesBranchKanaOk() (*string, bool)`

GetPartnerBankAccountAttributesBranchKanaOk returns a tuple with the PartnerBankAccountAttributesBranchKana field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerBankAccountAttributesBranchKana

`func (o *PartnersResponsePartners) SetPartnerBankAccountAttributesBranchKana(v string)`

SetPartnerBankAccountAttributesBranchKana sets PartnerBankAccountAttributesBranchKana field to given value.

### HasPartnerBankAccountAttributesBranchKana

`func (o *PartnersResponsePartners) HasPartnerBankAccountAttributesBranchKana() bool`

HasPartnerBankAccountAttributesBranchKana returns a boolean if a field has been set.

### SetPartnerBankAccountAttributesBranchKanaNil

`func (o *PartnersResponsePartners) SetPartnerBankAccountAttributesBranchKanaNil(b bool)`

 SetPartnerBankAccountAttributesBranchKanaNil sets the value for PartnerBankAccountAttributesBranchKana to be an explicit nil

### UnsetPartnerBankAccountAttributesBranchKana
`func (o *PartnersResponsePartners) UnsetPartnerBankAccountAttributesBranchKana()`

UnsetPartnerBankAccountAttributesBranchKana ensures that no value is present for PartnerBankAccountAttributesBranchKana, not even an explicit nil
### GetPartnerBankAccountAttributesBranchCode

`func (o *PartnersResponsePartners) GetPartnerBankAccountAttributesBranchCode() string`

GetPartnerBankAccountAttributesBranchCode returns the PartnerBankAccountAttributesBranchCode field if non-nil, zero value otherwise.

### GetPartnerBankAccountAttributesBranchCodeOk

`func (o *PartnersResponsePartners) GetPartnerBankAccountAttributesBranchCodeOk() (*string, bool)`

GetPartnerBankAccountAttributesBranchCodeOk returns a tuple with the PartnerBankAccountAttributesBranchCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerBankAccountAttributesBranchCode

`func (o *PartnersResponsePartners) SetPartnerBankAccountAttributesBranchCode(v string)`

SetPartnerBankAccountAttributesBranchCode sets PartnerBankAccountAttributesBranchCode field to given value.

### HasPartnerBankAccountAttributesBranchCode

`func (o *PartnersResponsePartners) HasPartnerBankAccountAttributesBranchCode() bool`

HasPartnerBankAccountAttributesBranchCode returns a boolean if a field has been set.

### SetPartnerBankAccountAttributesBranchCodeNil

`func (o *PartnersResponsePartners) SetPartnerBankAccountAttributesBranchCodeNil(b bool)`

 SetPartnerBankAccountAttributesBranchCodeNil sets the value for PartnerBankAccountAttributesBranchCode to be an explicit nil

### UnsetPartnerBankAccountAttributesBranchCode
`func (o *PartnersResponsePartners) UnsetPartnerBankAccountAttributesBranchCode()`

UnsetPartnerBankAccountAttributesBranchCode ensures that no value is present for PartnerBankAccountAttributesBranchCode, not even an explicit nil
### GetPartnerBankAccountAttributesAccountType

`func (o *PartnersResponsePartners) GetPartnerBankAccountAttributesAccountType() string`

GetPartnerBankAccountAttributesAccountType returns the PartnerBankAccountAttributesAccountType field if non-nil, zero value otherwise.

### GetPartnerBankAccountAttributesAccountTypeOk

`func (o *PartnersResponsePartners) GetPartnerBankAccountAttributesAccountTypeOk() (*string, bool)`

GetPartnerBankAccountAttributesAccountTypeOk returns a tuple with the PartnerBankAccountAttributesAccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerBankAccountAttributesAccountType

`func (o *PartnersResponsePartners) SetPartnerBankAccountAttributesAccountType(v string)`

SetPartnerBankAccountAttributesAccountType sets PartnerBankAccountAttributesAccountType field to given value.

### HasPartnerBankAccountAttributesAccountType

`func (o *PartnersResponsePartners) HasPartnerBankAccountAttributesAccountType() bool`

HasPartnerBankAccountAttributesAccountType returns a boolean if a field has been set.

### SetPartnerBankAccountAttributesAccountTypeNil

`func (o *PartnersResponsePartners) SetPartnerBankAccountAttributesAccountTypeNil(b bool)`

 SetPartnerBankAccountAttributesAccountTypeNil sets the value for PartnerBankAccountAttributesAccountType to be an explicit nil

### UnsetPartnerBankAccountAttributesAccountType
`func (o *PartnersResponsePartners) UnsetPartnerBankAccountAttributesAccountType()`

UnsetPartnerBankAccountAttributesAccountType ensures that no value is present for PartnerBankAccountAttributesAccountType, not even an explicit nil
### GetPartnerBankAccountAttributesAccountNumber

`func (o *PartnersResponsePartners) GetPartnerBankAccountAttributesAccountNumber() string`

GetPartnerBankAccountAttributesAccountNumber returns the PartnerBankAccountAttributesAccountNumber field if non-nil, zero value otherwise.

### GetPartnerBankAccountAttributesAccountNumberOk

`func (o *PartnersResponsePartners) GetPartnerBankAccountAttributesAccountNumberOk() (*string, bool)`

GetPartnerBankAccountAttributesAccountNumberOk returns a tuple with the PartnerBankAccountAttributesAccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerBankAccountAttributesAccountNumber

`func (o *PartnersResponsePartners) SetPartnerBankAccountAttributesAccountNumber(v string)`

SetPartnerBankAccountAttributesAccountNumber sets PartnerBankAccountAttributesAccountNumber field to given value.

### HasPartnerBankAccountAttributesAccountNumber

`func (o *PartnersResponsePartners) HasPartnerBankAccountAttributesAccountNumber() bool`

HasPartnerBankAccountAttributesAccountNumber returns a boolean if a field has been set.

### SetPartnerBankAccountAttributesAccountNumberNil

`func (o *PartnersResponsePartners) SetPartnerBankAccountAttributesAccountNumberNil(b bool)`

 SetPartnerBankAccountAttributesAccountNumberNil sets the value for PartnerBankAccountAttributesAccountNumber to be an explicit nil

### UnsetPartnerBankAccountAttributesAccountNumber
`func (o *PartnersResponsePartners) UnsetPartnerBankAccountAttributesAccountNumber()`

UnsetPartnerBankAccountAttributesAccountNumber ensures that no value is present for PartnerBankAccountAttributesAccountNumber, not even an explicit nil
### GetPartnerBankAccountAttributesAccountName

`func (o *PartnersResponsePartners) GetPartnerBankAccountAttributesAccountName() string`

GetPartnerBankAccountAttributesAccountName returns the PartnerBankAccountAttributesAccountName field if non-nil, zero value otherwise.

### GetPartnerBankAccountAttributesAccountNameOk

`func (o *PartnersResponsePartners) GetPartnerBankAccountAttributesAccountNameOk() (*string, bool)`

GetPartnerBankAccountAttributesAccountNameOk returns a tuple with the PartnerBankAccountAttributesAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerBankAccountAttributesAccountName

`func (o *PartnersResponsePartners) SetPartnerBankAccountAttributesAccountName(v string)`

SetPartnerBankAccountAttributesAccountName sets PartnerBankAccountAttributesAccountName field to given value.

### HasPartnerBankAccountAttributesAccountName

`func (o *PartnersResponsePartners) HasPartnerBankAccountAttributesAccountName() bool`

HasPartnerBankAccountAttributesAccountName returns a boolean if a field has been set.

### SetPartnerBankAccountAttributesAccountNameNil

`func (o *PartnersResponsePartners) SetPartnerBankAccountAttributesAccountNameNil(b bool)`

 SetPartnerBankAccountAttributesAccountNameNil sets the value for PartnerBankAccountAttributesAccountName to be an explicit nil

### UnsetPartnerBankAccountAttributesAccountName
`func (o *PartnersResponsePartners) UnsetPartnerBankAccountAttributesAccountName()`

UnsetPartnerBankAccountAttributesAccountName ensures that no value is present for PartnerBankAccountAttributesAccountName, not even an explicit nil
### GetPartnerBankAccountAttributesLongAccountName

`func (o *PartnersResponsePartners) GetPartnerBankAccountAttributesLongAccountName() string`

GetPartnerBankAccountAttributesLongAccountName returns the PartnerBankAccountAttributesLongAccountName field if non-nil, zero value otherwise.

### GetPartnerBankAccountAttributesLongAccountNameOk

`func (o *PartnersResponsePartners) GetPartnerBankAccountAttributesLongAccountNameOk() (*string, bool)`

GetPartnerBankAccountAttributesLongAccountNameOk returns a tuple with the PartnerBankAccountAttributesLongAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerBankAccountAttributesLongAccountName

`func (o *PartnersResponsePartners) SetPartnerBankAccountAttributesLongAccountName(v string)`

SetPartnerBankAccountAttributesLongAccountName sets PartnerBankAccountAttributesLongAccountName field to given value.

### HasPartnerBankAccountAttributesLongAccountName

`func (o *PartnersResponsePartners) HasPartnerBankAccountAttributesLongAccountName() bool`

HasPartnerBankAccountAttributesLongAccountName returns a boolean if a field has been set.

### SetPartnerBankAccountAttributesLongAccountNameNil

`func (o *PartnersResponsePartners) SetPartnerBankAccountAttributesLongAccountNameNil(b bool)`

 SetPartnerBankAccountAttributesLongAccountNameNil sets the value for PartnerBankAccountAttributesLongAccountName to be an explicit nil

### UnsetPartnerBankAccountAttributesLongAccountName
`func (o *PartnersResponsePartners) UnsetPartnerBankAccountAttributesLongAccountName()`

UnsetPartnerBankAccountAttributesLongAccountName ensures that no value is present for PartnerBankAccountAttributesLongAccountName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


