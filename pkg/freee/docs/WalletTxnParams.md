# WalletTxnParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntrySide** | **string** | 入金／出金 (入金: income, 出金: expense) | 
**Description** | Pointer to **string** | 取引内容 | [optional] 
**Amount** | **int32** | 取引金額 | 
**WalletableId** | **int32** | 口座ID | 
**WalletableType** | **string** | 口座区分 (銀行口座: bank_account, クレジットカード: credit_card, 現金: wallet) | 
**Date** | **string** | 取引日 (yyyy-mm-dd) | 
**CompanyId** | **int32** | 事業所ID | 
**Balance** | Pointer to **int32** | 残高 (銀行口座等) | [optional] 

## Methods

### NewWalletTxnParams

`func NewWalletTxnParams(entrySide string, amount int32, walletableId int32, walletableType string, date string, companyId int32, ) *WalletTxnParams`

NewWalletTxnParams instantiates a new WalletTxnParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWalletTxnParamsWithDefaults

`func NewWalletTxnParamsWithDefaults() *WalletTxnParams`

NewWalletTxnParamsWithDefaults instantiates a new WalletTxnParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntrySide

`func (o *WalletTxnParams) GetEntrySide() string`

GetEntrySide returns the EntrySide field if non-nil, zero value otherwise.

### GetEntrySideOk

`func (o *WalletTxnParams) GetEntrySideOk() (*string, bool)`

GetEntrySideOk returns a tuple with the EntrySide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntrySide

`func (o *WalletTxnParams) SetEntrySide(v string)`

SetEntrySide sets EntrySide field to given value.


### GetDescription

`func (o *WalletTxnParams) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WalletTxnParams) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WalletTxnParams) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WalletTxnParams) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAmount

`func (o *WalletTxnParams) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *WalletTxnParams) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *WalletTxnParams) SetAmount(v int32)`

SetAmount sets Amount field to given value.


### GetWalletableId

`func (o *WalletTxnParams) GetWalletableId() int32`

GetWalletableId returns the WalletableId field if non-nil, zero value otherwise.

### GetWalletableIdOk

`func (o *WalletTxnParams) GetWalletableIdOk() (*int32, bool)`

GetWalletableIdOk returns a tuple with the WalletableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletableId

`func (o *WalletTxnParams) SetWalletableId(v int32)`

SetWalletableId sets WalletableId field to given value.


### GetWalletableType

`func (o *WalletTxnParams) GetWalletableType() string`

GetWalletableType returns the WalletableType field if non-nil, zero value otherwise.

### GetWalletableTypeOk

`func (o *WalletTxnParams) GetWalletableTypeOk() (*string, bool)`

GetWalletableTypeOk returns a tuple with the WalletableType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletableType

`func (o *WalletTxnParams) SetWalletableType(v string)`

SetWalletableType sets WalletableType field to given value.


### GetDate

`func (o *WalletTxnParams) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *WalletTxnParams) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *WalletTxnParams) SetDate(v string)`

SetDate sets Date field to given value.


### GetCompanyId

`func (o *WalletTxnParams) GetCompanyId() int32`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *WalletTxnParams) GetCompanyIdOk() (*int32, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *WalletTxnParams) SetCompanyId(v int32)`

SetCompanyId sets CompanyId field to given value.


### GetBalance

`func (o *WalletTxnParams) GetBalance() int32`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *WalletTxnParams) GetBalanceOk() (*int32, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *WalletTxnParams) SetBalance(v int32)`

SetBalance sets Balance field to given value.

### HasBalance

`func (o *WalletTxnParams) HasBalance() bool`

HasBalance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


