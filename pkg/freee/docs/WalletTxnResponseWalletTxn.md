# WalletTxnResponseWalletTxn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | 明細ID | 
**CompanyId** | **int32** | 事業所ID | 
**Date** | **string** | 取引日(yyyy-mm-dd) | 
**Amount** | **int32** | 取引金額 | 
**DueAmount** | **int32** | 未決済金額 | 
**Balance** | **int32** | 残高(銀行口座等) | 
**EntrySide** | **string** | 入金／出金 (入金: income, 出金: expense) | 
**WalletableType** | **string** | 口座区分 (銀行口座: bank_account, クレジットカード: credit_card, 現金: wallet) | 
**WalletableId** | **int32** | 口座ID | 
**Description** | **string** | 取引内容 | 
**Status** | **int32** | 明細のステータス（消込待ち: 1, 消込済み: 2, 無視: 3, 消込中: 4） | 

## Methods

### NewWalletTxnResponseWalletTxn

`func NewWalletTxnResponseWalletTxn(id int32, companyId int32, date string, amount int32, dueAmount int32, balance int32, entrySide string, walletableType string, walletableId int32, description string, status int32, ) *WalletTxnResponseWalletTxn`

NewWalletTxnResponseWalletTxn instantiates a new WalletTxnResponseWalletTxn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWalletTxnResponseWalletTxnWithDefaults

`func NewWalletTxnResponseWalletTxnWithDefaults() *WalletTxnResponseWalletTxn`

NewWalletTxnResponseWalletTxnWithDefaults instantiates a new WalletTxnResponseWalletTxn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WalletTxnResponseWalletTxn) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WalletTxnResponseWalletTxn) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WalletTxnResponseWalletTxn) SetId(v int32)`

SetId sets Id field to given value.


### GetCompanyId

`func (o *WalletTxnResponseWalletTxn) GetCompanyId() int32`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *WalletTxnResponseWalletTxn) GetCompanyIdOk() (*int32, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *WalletTxnResponseWalletTxn) SetCompanyId(v int32)`

SetCompanyId sets CompanyId field to given value.


### GetDate

`func (o *WalletTxnResponseWalletTxn) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *WalletTxnResponseWalletTxn) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *WalletTxnResponseWalletTxn) SetDate(v string)`

SetDate sets Date field to given value.


### GetAmount

`func (o *WalletTxnResponseWalletTxn) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *WalletTxnResponseWalletTxn) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *WalletTxnResponseWalletTxn) SetAmount(v int32)`

SetAmount sets Amount field to given value.


### GetDueAmount

`func (o *WalletTxnResponseWalletTxn) GetDueAmount() int32`

GetDueAmount returns the DueAmount field if non-nil, zero value otherwise.

### GetDueAmountOk

`func (o *WalletTxnResponseWalletTxn) GetDueAmountOk() (*int32, bool)`

GetDueAmountOk returns a tuple with the DueAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueAmount

`func (o *WalletTxnResponseWalletTxn) SetDueAmount(v int32)`

SetDueAmount sets DueAmount field to given value.


### GetBalance

`func (o *WalletTxnResponseWalletTxn) GetBalance() int32`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *WalletTxnResponseWalletTxn) GetBalanceOk() (*int32, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *WalletTxnResponseWalletTxn) SetBalance(v int32)`

SetBalance sets Balance field to given value.


### GetEntrySide

`func (o *WalletTxnResponseWalletTxn) GetEntrySide() string`

GetEntrySide returns the EntrySide field if non-nil, zero value otherwise.

### GetEntrySideOk

`func (o *WalletTxnResponseWalletTxn) GetEntrySideOk() (*string, bool)`

GetEntrySideOk returns a tuple with the EntrySide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntrySide

`func (o *WalletTxnResponseWalletTxn) SetEntrySide(v string)`

SetEntrySide sets EntrySide field to given value.


### GetWalletableType

`func (o *WalletTxnResponseWalletTxn) GetWalletableType() string`

GetWalletableType returns the WalletableType field if non-nil, zero value otherwise.

### GetWalletableTypeOk

`func (o *WalletTxnResponseWalletTxn) GetWalletableTypeOk() (*string, bool)`

GetWalletableTypeOk returns a tuple with the WalletableType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletableType

`func (o *WalletTxnResponseWalletTxn) SetWalletableType(v string)`

SetWalletableType sets WalletableType field to given value.


### GetWalletableId

`func (o *WalletTxnResponseWalletTxn) GetWalletableId() int32`

GetWalletableId returns the WalletableId field if non-nil, zero value otherwise.

### GetWalletableIdOk

`func (o *WalletTxnResponseWalletTxn) GetWalletableIdOk() (*int32, bool)`

GetWalletableIdOk returns a tuple with the WalletableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletableId

`func (o *WalletTxnResponseWalletTxn) SetWalletableId(v int32)`

SetWalletableId sets WalletableId field to given value.


### GetDescription

`func (o *WalletTxnResponseWalletTxn) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WalletTxnResponseWalletTxn) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WalletTxnResponseWalletTxn) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetStatus

`func (o *WalletTxnResponseWalletTxn) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WalletTxnResponseWalletTxn) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WalletTxnResponseWalletTxn) SetStatus(v int32)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


