# DealCreateResponseDealPayments

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | 取引行ID | 
**Date** | **string** | 支払日 | 
**FromWalletableType** | Pointer to **string** | 口座区分 (銀行口座: bank_account, クレジットカード: credit_card, 現金: wallet, プライベート資金（法人の場合は役員借入金もしくは役員借入金、個人の場合は事業主貸もしくは事業主借）: private_account_item) | [optional] 
**FromWalletableId** | Pointer to **int32** | 口座ID（from_walletable_typeがprivate_account_itemの場合は勘定科目ID） | [optional] 
**Amount** | **int32** | 支払金額 | 

## Methods

### NewDealCreateResponseDealPayments

`func NewDealCreateResponseDealPayments(id int32, date string, amount int32, ) *DealCreateResponseDealPayments`

NewDealCreateResponseDealPayments instantiates a new DealCreateResponseDealPayments object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDealCreateResponseDealPaymentsWithDefaults

`func NewDealCreateResponseDealPaymentsWithDefaults() *DealCreateResponseDealPayments`

NewDealCreateResponseDealPaymentsWithDefaults instantiates a new DealCreateResponseDealPayments object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DealCreateResponseDealPayments) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DealCreateResponseDealPayments) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DealCreateResponseDealPayments) SetId(v int32)`

SetId sets Id field to given value.


### GetDate

`func (o *DealCreateResponseDealPayments) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *DealCreateResponseDealPayments) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *DealCreateResponseDealPayments) SetDate(v string)`

SetDate sets Date field to given value.


### GetFromWalletableType

`func (o *DealCreateResponseDealPayments) GetFromWalletableType() string`

GetFromWalletableType returns the FromWalletableType field if non-nil, zero value otherwise.

### GetFromWalletableTypeOk

`func (o *DealCreateResponseDealPayments) GetFromWalletableTypeOk() (*string, bool)`

GetFromWalletableTypeOk returns a tuple with the FromWalletableType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromWalletableType

`func (o *DealCreateResponseDealPayments) SetFromWalletableType(v string)`

SetFromWalletableType sets FromWalletableType field to given value.

### HasFromWalletableType

`func (o *DealCreateResponseDealPayments) HasFromWalletableType() bool`

HasFromWalletableType returns a boolean if a field has been set.

### GetFromWalletableId

`func (o *DealCreateResponseDealPayments) GetFromWalletableId() int32`

GetFromWalletableId returns the FromWalletableId field if non-nil, zero value otherwise.

### GetFromWalletableIdOk

`func (o *DealCreateResponseDealPayments) GetFromWalletableIdOk() (*int32, bool)`

GetFromWalletableIdOk returns a tuple with the FromWalletableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromWalletableId

`func (o *DealCreateResponseDealPayments) SetFromWalletableId(v int32)`

SetFromWalletableId sets FromWalletableId field to given value.

### HasFromWalletableId

`func (o *DealCreateResponseDealPayments) HasFromWalletableId() bool`

HasFromWalletableId returns a boolean if a field has been set.

### GetAmount

`func (o *DealCreateResponseDealPayments) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *DealCreateResponseDealPayments) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *DealCreateResponseDealPayments) SetAmount(v int32)`

SetAmount sets Amount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


