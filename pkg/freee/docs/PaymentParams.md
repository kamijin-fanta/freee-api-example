# PaymentParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyId** | **int32** | 事業所ID | 
**Date** | **string** | 支払日 | 
**FromWalletableType** | **string** | 口座区分 (銀行口座: bank_account, クレジットカード: credit_card, 現金: wallet, プライベート資金（法人の場合は役員借入金もしくは役員借入金、個人の場合は事業主貸もしくは事業主借）: private_account_item)：payments指定時は必須 | 
**FromWalletableId** | **int32** | 口座ID（from_walletable_typeがprivate_account_itemの場合は勘定科目ID）：payments指定時は必須 | 
**Amount** | **int32** | 金額 | 

## Methods

### NewPaymentParams

`func NewPaymentParams(companyId int32, date string, fromWalletableType string, fromWalletableId int32, amount int32, ) *PaymentParams`

NewPaymentParams instantiates a new PaymentParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentParamsWithDefaults

`func NewPaymentParamsWithDefaults() *PaymentParams`

NewPaymentParamsWithDefaults instantiates a new PaymentParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyId

`func (o *PaymentParams) GetCompanyId() int32`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *PaymentParams) GetCompanyIdOk() (*int32, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *PaymentParams) SetCompanyId(v int32)`

SetCompanyId sets CompanyId field to given value.


### GetDate

`func (o *PaymentParams) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *PaymentParams) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *PaymentParams) SetDate(v string)`

SetDate sets Date field to given value.


### GetFromWalletableType

`func (o *PaymentParams) GetFromWalletableType() string`

GetFromWalletableType returns the FromWalletableType field if non-nil, zero value otherwise.

### GetFromWalletableTypeOk

`func (o *PaymentParams) GetFromWalletableTypeOk() (*string, bool)`

GetFromWalletableTypeOk returns a tuple with the FromWalletableType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromWalletableType

`func (o *PaymentParams) SetFromWalletableType(v string)`

SetFromWalletableType sets FromWalletableType field to given value.


### GetFromWalletableId

`func (o *PaymentParams) GetFromWalletableId() int32`

GetFromWalletableId returns the FromWalletableId field if non-nil, zero value otherwise.

### GetFromWalletableIdOk

`func (o *PaymentParams) GetFromWalletableIdOk() (*int32, bool)`

GetFromWalletableIdOk returns a tuple with the FromWalletableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromWalletableId

`func (o *PaymentParams) SetFromWalletableId(v int32)`

SetFromWalletableId sets FromWalletableId field to given value.


### GetAmount

`func (o *PaymentParams) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PaymentParams) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PaymentParams) SetAmount(v int32)`

SetAmount sets Amount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


