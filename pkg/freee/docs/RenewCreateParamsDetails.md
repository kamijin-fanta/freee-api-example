# RenewCreateParamsDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountItemId** | **int32** | 勘定科目ID | 
**TaxCode** | **int32** | 税区分コード | 
**Amount** | **int32** | 取引金額（税込で指定してください） | 
**Vat** | Pointer to **int32** | 消費税額（指定しない場合は自動で計算されます） | [optional] 
**ItemId** | Pointer to **int32** | 品目ID | [optional] 
**SectionId** | Pointer to **int32** | 部門ID | [optional] 
**TagIds** | Pointer to **[]int32** | メモタグID | [optional] 
**Segment1TagId** | Pointer to **int32** | セグメント１ID | [optional] 
**Segment2TagId** | Pointer to **int32** | セグメント２ID | [optional] 
**Segment3TagId** | Pointer to **int32** | セグメント３ID | [optional] 
**Description** | Pointer to **string** | 備考 | [optional] 

## Methods

### NewRenewCreateParamsDetails

`func NewRenewCreateParamsDetails(accountItemId int32, taxCode int32, amount int32, ) *RenewCreateParamsDetails`

NewRenewCreateParamsDetails instantiates a new RenewCreateParamsDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRenewCreateParamsDetailsWithDefaults

`func NewRenewCreateParamsDetailsWithDefaults() *RenewCreateParamsDetails`

NewRenewCreateParamsDetailsWithDefaults instantiates a new RenewCreateParamsDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountItemId

`func (o *RenewCreateParamsDetails) GetAccountItemId() int32`

GetAccountItemId returns the AccountItemId field if non-nil, zero value otherwise.

### GetAccountItemIdOk

`func (o *RenewCreateParamsDetails) GetAccountItemIdOk() (*int32, bool)`

GetAccountItemIdOk returns a tuple with the AccountItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountItemId

`func (o *RenewCreateParamsDetails) SetAccountItemId(v int32)`

SetAccountItemId sets AccountItemId field to given value.


### GetTaxCode

`func (o *RenewCreateParamsDetails) GetTaxCode() int32`

GetTaxCode returns the TaxCode field if non-nil, zero value otherwise.

### GetTaxCodeOk

`func (o *RenewCreateParamsDetails) GetTaxCodeOk() (*int32, bool)`

GetTaxCodeOk returns a tuple with the TaxCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCode

`func (o *RenewCreateParamsDetails) SetTaxCode(v int32)`

SetTaxCode sets TaxCode field to given value.


### GetAmount

`func (o *RenewCreateParamsDetails) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *RenewCreateParamsDetails) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *RenewCreateParamsDetails) SetAmount(v int32)`

SetAmount sets Amount field to given value.


### GetVat

`func (o *RenewCreateParamsDetails) GetVat() int32`

GetVat returns the Vat field if non-nil, zero value otherwise.

### GetVatOk

`func (o *RenewCreateParamsDetails) GetVatOk() (*int32, bool)`

GetVatOk returns a tuple with the Vat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVat

`func (o *RenewCreateParamsDetails) SetVat(v int32)`

SetVat sets Vat field to given value.

### HasVat

`func (o *RenewCreateParamsDetails) HasVat() bool`

HasVat returns a boolean if a field has been set.

### GetItemId

`func (o *RenewCreateParamsDetails) GetItemId() int32`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *RenewCreateParamsDetails) GetItemIdOk() (*int32, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *RenewCreateParamsDetails) SetItemId(v int32)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *RenewCreateParamsDetails) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### GetSectionId

`func (o *RenewCreateParamsDetails) GetSectionId() int32`

GetSectionId returns the SectionId field if non-nil, zero value otherwise.

### GetSectionIdOk

`func (o *RenewCreateParamsDetails) GetSectionIdOk() (*int32, bool)`

GetSectionIdOk returns a tuple with the SectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionId

`func (o *RenewCreateParamsDetails) SetSectionId(v int32)`

SetSectionId sets SectionId field to given value.

### HasSectionId

`func (o *RenewCreateParamsDetails) HasSectionId() bool`

HasSectionId returns a boolean if a field has been set.

### GetTagIds

`func (o *RenewCreateParamsDetails) GetTagIds() []int32`

GetTagIds returns the TagIds field if non-nil, zero value otherwise.

### GetTagIdsOk

`func (o *RenewCreateParamsDetails) GetTagIdsOk() (*[]int32, bool)`

GetTagIdsOk returns a tuple with the TagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagIds

`func (o *RenewCreateParamsDetails) SetTagIds(v []int32)`

SetTagIds sets TagIds field to given value.

### HasTagIds

`func (o *RenewCreateParamsDetails) HasTagIds() bool`

HasTagIds returns a boolean if a field has been set.

### GetSegment1TagId

`func (o *RenewCreateParamsDetails) GetSegment1TagId() int32`

GetSegment1TagId returns the Segment1TagId field if non-nil, zero value otherwise.

### GetSegment1TagIdOk

`func (o *RenewCreateParamsDetails) GetSegment1TagIdOk() (*int32, bool)`

GetSegment1TagIdOk returns a tuple with the Segment1TagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment1TagId

`func (o *RenewCreateParamsDetails) SetSegment1TagId(v int32)`

SetSegment1TagId sets Segment1TagId field to given value.

### HasSegment1TagId

`func (o *RenewCreateParamsDetails) HasSegment1TagId() bool`

HasSegment1TagId returns a boolean if a field has been set.

### GetSegment2TagId

`func (o *RenewCreateParamsDetails) GetSegment2TagId() int32`

GetSegment2TagId returns the Segment2TagId field if non-nil, zero value otherwise.

### GetSegment2TagIdOk

`func (o *RenewCreateParamsDetails) GetSegment2TagIdOk() (*int32, bool)`

GetSegment2TagIdOk returns a tuple with the Segment2TagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment2TagId

`func (o *RenewCreateParamsDetails) SetSegment2TagId(v int32)`

SetSegment2TagId sets Segment2TagId field to given value.

### HasSegment2TagId

`func (o *RenewCreateParamsDetails) HasSegment2TagId() bool`

HasSegment2TagId returns a boolean if a field has been set.

### GetSegment3TagId

`func (o *RenewCreateParamsDetails) GetSegment3TagId() int32`

GetSegment3TagId returns the Segment3TagId field if non-nil, zero value otherwise.

### GetSegment3TagIdOk

`func (o *RenewCreateParamsDetails) GetSegment3TagIdOk() (*int32, bool)`

GetSegment3TagIdOk returns a tuple with the Segment3TagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment3TagId

`func (o *RenewCreateParamsDetails) SetSegment3TagId(v int32)`

SetSegment3TagId sets Segment3TagId field to given value.

### HasSegment3TagId

`func (o *RenewCreateParamsDetails) HasSegment3TagId() bool`

HasSegment3TagId returns a boolean if a field has been set.

### GetDescription

`func (o *RenewCreateParamsDetails) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RenewCreateParamsDetails) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RenewCreateParamsDetails) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RenewCreateParamsDetails) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


