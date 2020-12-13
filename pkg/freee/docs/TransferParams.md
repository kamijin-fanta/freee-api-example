# TransferParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ToWalletableId** | **int32** | 振替先口座ID | 
**ToWalletableType** | **string** | 振替先口座区分 (銀行口座: bank_account, クレジットカード: credit_card, 現金: wallet) | 
**FromWalletableId** | **int32** | 振替元口座ID | 
**FromWalletableType** | **string** | 振替元口座区分 (銀行口座: bank_account, クレジットカード: credit_card, 現金: wallet) | 
**Amount** | **int32** | 金額 | 
**Date** | **string** | 振替日 (yyyy-mm-dd) | 
**CompanyId** | **int32** | 事業所ID | 
**Description** | Pointer to **string** | 備考 | [optional] 

## Methods

### NewTransferParams

`func NewTransferParams(toWalletableId int32, toWalletableType string, fromWalletableId int32, fromWalletableType string, amount int32, date string, companyId int32, ) *TransferParams`

NewTransferParams instantiates a new TransferParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferParamsWithDefaults

`func NewTransferParamsWithDefaults() *TransferParams`

NewTransferParamsWithDefaults instantiates a new TransferParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToWalletableId

`func (o *TransferParams) GetToWalletableId() int32`

GetToWalletableId returns the ToWalletableId field if non-nil, zero value otherwise.

### GetToWalletableIdOk

`func (o *TransferParams) GetToWalletableIdOk() (*int32, bool)`

GetToWalletableIdOk returns a tuple with the ToWalletableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToWalletableId

`func (o *TransferParams) SetToWalletableId(v int32)`

SetToWalletableId sets ToWalletableId field to given value.


### GetToWalletableType

`func (o *TransferParams) GetToWalletableType() string`

GetToWalletableType returns the ToWalletableType field if non-nil, zero value otherwise.

### GetToWalletableTypeOk

`func (o *TransferParams) GetToWalletableTypeOk() (*string, bool)`

GetToWalletableTypeOk returns a tuple with the ToWalletableType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToWalletableType

`func (o *TransferParams) SetToWalletableType(v string)`

SetToWalletableType sets ToWalletableType field to given value.


### GetFromWalletableId

`func (o *TransferParams) GetFromWalletableId() int32`

GetFromWalletableId returns the FromWalletableId field if non-nil, zero value otherwise.

### GetFromWalletableIdOk

`func (o *TransferParams) GetFromWalletableIdOk() (*int32, bool)`

GetFromWalletableIdOk returns a tuple with the FromWalletableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromWalletableId

`func (o *TransferParams) SetFromWalletableId(v int32)`

SetFromWalletableId sets FromWalletableId field to given value.


### GetFromWalletableType

`func (o *TransferParams) GetFromWalletableType() string`

GetFromWalletableType returns the FromWalletableType field if non-nil, zero value otherwise.

### GetFromWalletableTypeOk

`func (o *TransferParams) GetFromWalletableTypeOk() (*string, bool)`

GetFromWalletableTypeOk returns a tuple with the FromWalletableType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromWalletableType

`func (o *TransferParams) SetFromWalletableType(v string)`

SetFromWalletableType sets FromWalletableType field to given value.


### GetAmount

`func (o *TransferParams) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TransferParams) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TransferParams) SetAmount(v int32)`

SetAmount sets Amount field to given value.


### GetDate

`func (o *TransferParams) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *TransferParams) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *TransferParams) SetDate(v string)`

SetDate sets Date field to given value.


### GetCompanyId

`func (o *TransferParams) GetCompanyId() int32`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *TransferParams) GetCompanyIdOk() (*int32, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *TransferParams) SetCompanyId(v int32)`

SetCompanyId sets CompanyId field to given value.


### GetDescription

`func (o *TransferParams) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TransferParams) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TransferParams) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TransferParams) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


