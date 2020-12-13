# WalletableCreateResponseWalletable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | 口座ID | 
**Name** | **string** | 口座名, 最大255文字 | 
**BankId** | **int32** | サービスID | 
**Type** | **string** | 口座区分 (銀行口座: bank_account, クレジットカード: credit_card, 現金: wallet) | 

## Methods

### NewWalletableCreateResponseWalletable

`func NewWalletableCreateResponseWalletable(id int32, name string, bankId int32, type_ string, ) *WalletableCreateResponseWalletable`

NewWalletableCreateResponseWalletable instantiates a new WalletableCreateResponseWalletable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWalletableCreateResponseWalletableWithDefaults

`func NewWalletableCreateResponseWalletableWithDefaults() *WalletableCreateResponseWalletable`

NewWalletableCreateResponseWalletableWithDefaults instantiates a new WalletableCreateResponseWalletable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WalletableCreateResponseWalletable) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WalletableCreateResponseWalletable) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WalletableCreateResponseWalletable) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *WalletableCreateResponseWalletable) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WalletableCreateResponseWalletable) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WalletableCreateResponseWalletable) SetName(v string)`

SetName sets Name field to given value.


### GetBankId

`func (o *WalletableCreateResponseWalletable) GetBankId() int32`

GetBankId returns the BankId field if non-nil, zero value otherwise.

### GetBankIdOk

`func (o *WalletableCreateResponseWalletable) GetBankIdOk() (*int32, bool)`

GetBankIdOk returns a tuple with the BankId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankId

`func (o *WalletableCreateResponseWalletable) SetBankId(v int32)`

SetBankId sets BankId field to given value.


### GetType

`func (o *WalletableCreateResponseWalletable) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WalletableCreateResponseWalletable) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WalletableCreateResponseWalletable) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


