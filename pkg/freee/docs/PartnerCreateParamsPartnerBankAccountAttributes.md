# PartnerCreateParamsPartnerBankAccountAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BankName** | Pointer to **string** | 銀行名 | [optional] 
**BankNameKana** | Pointer to **string** | 銀行名（カナ） | [optional] 
**BankCode** | Pointer to **string** | 銀行コード | [optional] 
**BranchName** | Pointer to **string** | 支店名 | [optional] 
**BranchKana** | Pointer to **string** | 支店名（カナ） | [optional] 
**BranchCode** | Pointer to **string** | 支店番号 | [optional] 
**AccountType** | Pointer to **string** | 口座種別(ordinary:普通、checking：当座、earmarked：納税準備預金、savings：貯蓄、other:その他) | [optional] 
**AccountNumber** | Pointer to **string** | 口座番号 | [optional] 
**LongAccountName** | Pointer to **string** | 受取人名 | [optional] 
**AccountName** | Pointer to **string** | 受取人名（カナ） | [optional] 

## Methods

### NewPartnerCreateParamsPartnerBankAccountAttributes

`func NewPartnerCreateParamsPartnerBankAccountAttributes() *PartnerCreateParamsPartnerBankAccountAttributes`

NewPartnerCreateParamsPartnerBankAccountAttributes instantiates a new PartnerCreateParamsPartnerBankAccountAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartnerCreateParamsPartnerBankAccountAttributesWithDefaults

`func NewPartnerCreateParamsPartnerBankAccountAttributesWithDefaults() *PartnerCreateParamsPartnerBankAccountAttributes`

NewPartnerCreateParamsPartnerBankAccountAttributesWithDefaults instantiates a new PartnerCreateParamsPartnerBankAccountAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBankName

`func (o *PartnerCreateParamsPartnerBankAccountAttributes) GetBankName() string`

GetBankName returns the BankName field if non-nil, zero value otherwise.

### GetBankNameOk

`func (o *PartnerCreateParamsPartnerBankAccountAttributes) GetBankNameOk() (*string, bool)`

GetBankNameOk returns a tuple with the BankName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankName

`func (o *PartnerCreateParamsPartnerBankAccountAttributes) SetBankName(v string)`

SetBankName sets BankName field to given value.

### HasBankName

`func (o *PartnerCreateParamsPartnerBankAccountAttributes) HasBankName() bool`

HasBankName returns a boolean if a field has been set.

### GetBankNameKana

`func (o *PartnerCreateParamsPartnerBankAccountAttributes) GetBankNameKana() string`

GetBankNameKana returns the BankNameKana field if non-nil, zero value otherwise.

### GetBankNameKanaOk

`func (o *PartnerCreateParamsPartnerBankAccountAttributes) GetBankNameKanaOk() (*string, bool)`

GetBankNameKanaOk returns a tuple with the BankNameKana field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankNameKana

`func (o *PartnerCreateParamsPartnerBankAccountAttributes) SetBankNameKana(v string)`

SetBankNameKana sets BankNameKana field to given value.

### HasBankNameKana

`func (o *PartnerCreateParamsPartnerBankAccountAttributes) HasBankNameKana() bool`

HasBankNameKana returns a boolean if a field has been set.

### GetBankCode

`func (o *PartnerCreateParamsPartnerBankAccountAttributes) GetBankCode() string`

GetBankCode returns the BankCode field if non-nil, zero value otherwise.

### GetBankCodeOk

`func (o *PartnerCreateParamsPartnerBankAccountAttributes) GetBankCodeOk() (*string, bool)`

GetBankCodeOk returns a tuple with the BankCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankCode

`func (o *PartnerCreateParamsPartnerBankAccountAttributes) SetBankCode(v string)`

SetBankCode sets BankCode field to given value.

### HasBankCode

`func (o *PartnerCreateParamsPartnerBankAccountAttributes) HasBankCode() bool`

HasBankCode returns a boolean if a field has been set.

### GetBranchName

`func (o *PartnerCreateParamsPartnerBankAccountAttributes) GetBranchName() string`

GetBranchName returns the BranchName field if non-nil, zero value otherwise.

### GetBranchNameOk

`func (o *PartnerCreateParamsPartnerBankAccountAttributes) GetBranchNameOk() (*string, bool)`

GetBranchNameOk returns a tuple with the BranchName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchName

`func (o *PartnerCreateParamsPartnerBankAccountAttributes) SetBranchName(v string)`

SetBranchName sets BranchName field to given value.

### HasBranchName

`func (o *PartnerCreateParamsPartnerBankAccountAttributes) HasBranchName() bool`

HasBranchName returns a boolean if a field has been set.

### GetBranchKana

`func (o *PartnerCreateParamsPartnerBankAccountAttributes) GetBranchKana() string`

GetBranchKana returns the BranchKana field if non-nil, zero value otherwise.

### GetBranchKanaOk

`func (o *PartnerCreateParamsPartnerBankAccountAttributes) GetBranchKanaOk() (*string, bool)`

GetBranchKanaOk returns a tuple with the BranchKana field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchKana

`func (o *PartnerCreateParamsPartnerBankAccountAttributes) SetBranchKana(v string)`

SetBranchKana sets BranchKana field to given value.

### HasBranchKana

`func (o *PartnerCreateParamsPartnerBankAccountAttributes) HasBranchKana() bool`

HasBranchKana returns a boolean if a field has been set.

### GetBranchCode

`func (o *PartnerCreateParamsPartnerBankAccountAttributes) GetBranchCode() string`

GetBranchCode returns the BranchCode field if non-nil, zero value otherwise.

### GetBranchCodeOk

`func (o *PartnerCreateParamsPartnerBankAccountAttributes) GetBranchCodeOk() (*string, bool)`

GetBranchCodeOk returns a tuple with the BranchCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchCode

`func (o *PartnerCreateParamsPartnerBankAccountAttributes) SetBranchCode(v string)`

SetBranchCode sets BranchCode field to given value.

### HasBranchCode

`func (o *PartnerCreateParamsPartnerBankAccountAttributes) HasBranchCode() bool`

HasBranchCode returns a boolean if a field has been set.

### GetAccountType

`func (o *PartnerCreateParamsPartnerBankAccountAttributes) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *PartnerCreateParamsPartnerBankAccountAttributes) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *PartnerCreateParamsPartnerBankAccountAttributes) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *PartnerCreateParamsPartnerBankAccountAttributes) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### GetAccountNumber

`func (o *PartnerCreateParamsPartnerBankAccountAttributes) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *PartnerCreateParamsPartnerBankAccountAttributes) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *PartnerCreateParamsPartnerBankAccountAttributes) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *PartnerCreateParamsPartnerBankAccountAttributes) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### GetLongAccountName

`func (o *PartnerCreateParamsPartnerBankAccountAttributes) GetLongAccountName() string`

GetLongAccountName returns the LongAccountName field if non-nil, zero value otherwise.

### GetLongAccountNameOk

`func (o *PartnerCreateParamsPartnerBankAccountAttributes) GetLongAccountNameOk() (*string, bool)`

GetLongAccountNameOk returns a tuple with the LongAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongAccountName

`func (o *PartnerCreateParamsPartnerBankAccountAttributes) SetLongAccountName(v string)`

SetLongAccountName sets LongAccountName field to given value.

### HasLongAccountName

`func (o *PartnerCreateParamsPartnerBankAccountAttributes) HasLongAccountName() bool`

HasLongAccountName returns a boolean if a field has been set.

### GetAccountName

`func (o *PartnerCreateParamsPartnerBankAccountAttributes) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *PartnerCreateParamsPartnerBankAccountAttributes) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *PartnerCreateParamsPartnerBankAccountAttributes) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.

### HasAccountName

`func (o *PartnerCreateParamsPartnerBankAccountAttributes) HasAccountName() bool`

HasAccountName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


