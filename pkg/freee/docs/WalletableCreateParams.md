# WalletableCreateParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | 口座名 (255文字以内) | 
**Type** | **string** | 口座種別（bank_account : 銀行口座, credit_card : クレジットカード, wallet : その他の決済口座） | 
**CompanyId** | **int32** | 事業所ID | 
**BankId** | Pointer to **int32** | サービスID | [optional] 
**GroupName** | Pointer to **string** | 決算書表示名（小カテゴリー）　例：売掛金, 受取手形, 未収入金（法人のみ）, 買掛金, 支払手形, 未払金, 預り金, 前受金 | [optional] 

## Methods

### NewWalletableCreateParams

`func NewWalletableCreateParams(name string, type_ string, companyId int32, ) *WalletableCreateParams`

NewWalletableCreateParams instantiates a new WalletableCreateParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWalletableCreateParamsWithDefaults

`func NewWalletableCreateParamsWithDefaults() *WalletableCreateParams`

NewWalletableCreateParamsWithDefaults instantiates a new WalletableCreateParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *WalletableCreateParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WalletableCreateParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WalletableCreateParams) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *WalletableCreateParams) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WalletableCreateParams) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WalletableCreateParams) SetType(v string)`

SetType sets Type field to given value.


### GetCompanyId

`func (o *WalletableCreateParams) GetCompanyId() int32`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *WalletableCreateParams) GetCompanyIdOk() (*int32, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *WalletableCreateParams) SetCompanyId(v int32)`

SetCompanyId sets CompanyId field to given value.


### GetBankId

`func (o *WalletableCreateParams) GetBankId() int32`

GetBankId returns the BankId field if non-nil, zero value otherwise.

### GetBankIdOk

`func (o *WalletableCreateParams) GetBankIdOk() (*int32, bool)`

GetBankIdOk returns a tuple with the BankId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankId

`func (o *WalletableCreateParams) SetBankId(v int32)`

SetBankId sets BankId field to given value.

### HasBankId

`func (o *WalletableCreateParams) HasBankId() bool`

HasBankId returns a boolean if a field has been set.

### GetGroupName

`func (o *WalletableCreateParams) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *WalletableCreateParams) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *WalletableCreateParams) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *WalletableCreateParams) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


