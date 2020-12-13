# BankResponseBank

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | 連携サービスID | 
**Name** | Pointer to **string** | 連携サービス名 | [optional] 
**Type** | Pointer to **string** | 連携サービス種別: (銀行口座: bank_account, クレジットカード: credit_card, 現金: wallet) | [optional] 
**NameKana** | Pointer to **NullableString** | 連携サービス名(カナ) | [optional] 

## Methods

### NewBankResponseBank

`func NewBankResponseBank(id int32, ) *BankResponseBank`

NewBankResponseBank instantiates a new BankResponseBank object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBankResponseBankWithDefaults

`func NewBankResponseBankWithDefaults() *BankResponseBank`

NewBankResponseBankWithDefaults instantiates a new BankResponseBank object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BankResponseBank) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BankResponseBank) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BankResponseBank) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *BankResponseBank) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BankResponseBank) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BankResponseBank) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BankResponseBank) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *BankResponseBank) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BankResponseBank) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BankResponseBank) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *BankResponseBank) HasType() bool`

HasType returns a boolean if a field has been set.

### GetNameKana

`func (o *BankResponseBank) GetNameKana() string`

GetNameKana returns the NameKana field if non-nil, zero value otherwise.

### GetNameKanaOk

`func (o *BankResponseBank) GetNameKanaOk() (*string, bool)`

GetNameKanaOk returns a tuple with the NameKana field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameKana

`func (o *BankResponseBank) SetNameKana(v string)`

SetNameKana sets NameKana field to given value.

### HasNameKana

`func (o *BankResponseBank) HasNameKana() bool`

HasNameKana returns a boolean if a field has been set.

### SetNameKanaNil

`func (o *BankResponseBank) SetNameKanaNil(b bool)`

 SetNameKanaNil sets the value for NameKana to be an explicit nil

### UnsetNameKana
`func (o *BankResponseBank) UnsetNameKana()`

UnsetNameKana ensures that no value is present for NameKana, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


