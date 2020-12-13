# DealUpdateParamsDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | 取引行ID: 既存取引行を更新する場合に指定します。IDを指定しない取引行は、新規行として扱われ追加されます。また、detailsに含まれない既存の取引行は削除されます。更新後も残したい行は、必ず取引行IDを指定してdetailsに含めてください。 | [optional] 
**TaxCode** | **int32** | 税区分コード | 
**AccountItemId** | **int32** | 勘定科目ID | 
**Amount** | **int32** | 取引金額（税込で指定してください） | 
**ItemId** | Pointer to **int32** | 品目ID | [optional] 
**SectionId** | Pointer to **int32** | 部門ID | [optional] 
**TagIds** | Pointer to **[]int32** | メモタグID | [optional] 
**Segment1TagId** | Pointer to **int32** | セグメント１ID | [optional] 
**Segment2TagId** | Pointer to **int32** | セグメント２ID | [optional] 
**Segment3TagId** | Pointer to **int32** | セグメント３ID | [optional] 
**Description** | Pointer to **string** | 備考 | [optional] 
**Vat** | Pointer to **int32** | 消費税額（指定しない場合は自動で計算されます） | [optional] 

## Methods

### NewDealUpdateParamsDetails

`func NewDealUpdateParamsDetails(taxCode int32, accountItemId int32, amount int32, ) *DealUpdateParamsDetails`

NewDealUpdateParamsDetails instantiates a new DealUpdateParamsDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDealUpdateParamsDetailsWithDefaults

`func NewDealUpdateParamsDetailsWithDefaults() *DealUpdateParamsDetails`

NewDealUpdateParamsDetailsWithDefaults instantiates a new DealUpdateParamsDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DealUpdateParamsDetails) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DealUpdateParamsDetails) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DealUpdateParamsDetails) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *DealUpdateParamsDetails) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTaxCode

`func (o *DealUpdateParamsDetails) GetTaxCode() int32`

GetTaxCode returns the TaxCode field if non-nil, zero value otherwise.

### GetTaxCodeOk

`func (o *DealUpdateParamsDetails) GetTaxCodeOk() (*int32, bool)`

GetTaxCodeOk returns a tuple with the TaxCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCode

`func (o *DealUpdateParamsDetails) SetTaxCode(v int32)`

SetTaxCode sets TaxCode field to given value.


### GetAccountItemId

`func (o *DealUpdateParamsDetails) GetAccountItemId() int32`

GetAccountItemId returns the AccountItemId field if non-nil, zero value otherwise.

### GetAccountItemIdOk

`func (o *DealUpdateParamsDetails) GetAccountItemIdOk() (*int32, bool)`

GetAccountItemIdOk returns a tuple with the AccountItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountItemId

`func (o *DealUpdateParamsDetails) SetAccountItemId(v int32)`

SetAccountItemId sets AccountItemId field to given value.


### GetAmount

`func (o *DealUpdateParamsDetails) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *DealUpdateParamsDetails) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *DealUpdateParamsDetails) SetAmount(v int32)`

SetAmount sets Amount field to given value.


### GetItemId

`func (o *DealUpdateParamsDetails) GetItemId() int32`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *DealUpdateParamsDetails) GetItemIdOk() (*int32, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *DealUpdateParamsDetails) SetItemId(v int32)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *DealUpdateParamsDetails) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### GetSectionId

`func (o *DealUpdateParamsDetails) GetSectionId() int32`

GetSectionId returns the SectionId field if non-nil, zero value otherwise.

### GetSectionIdOk

`func (o *DealUpdateParamsDetails) GetSectionIdOk() (*int32, bool)`

GetSectionIdOk returns a tuple with the SectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionId

`func (o *DealUpdateParamsDetails) SetSectionId(v int32)`

SetSectionId sets SectionId field to given value.

### HasSectionId

`func (o *DealUpdateParamsDetails) HasSectionId() bool`

HasSectionId returns a boolean if a field has been set.

### GetTagIds

`func (o *DealUpdateParamsDetails) GetTagIds() []int32`

GetTagIds returns the TagIds field if non-nil, zero value otherwise.

### GetTagIdsOk

`func (o *DealUpdateParamsDetails) GetTagIdsOk() (*[]int32, bool)`

GetTagIdsOk returns a tuple with the TagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagIds

`func (o *DealUpdateParamsDetails) SetTagIds(v []int32)`

SetTagIds sets TagIds field to given value.

### HasTagIds

`func (o *DealUpdateParamsDetails) HasTagIds() bool`

HasTagIds returns a boolean if a field has been set.

### GetSegment1TagId

`func (o *DealUpdateParamsDetails) GetSegment1TagId() int32`

GetSegment1TagId returns the Segment1TagId field if non-nil, zero value otherwise.

### GetSegment1TagIdOk

`func (o *DealUpdateParamsDetails) GetSegment1TagIdOk() (*int32, bool)`

GetSegment1TagIdOk returns a tuple with the Segment1TagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment1TagId

`func (o *DealUpdateParamsDetails) SetSegment1TagId(v int32)`

SetSegment1TagId sets Segment1TagId field to given value.

### HasSegment1TagId

`func (o *DealUpdateParamsDetails) HasSegment1TagId() bool`

HasSegment1TagId returns a boolean if a field has been set.

### GetSegment2TagId

`func (o *DealUpdateParamsDetails) GetSegment2TagId() int32`

GetSegment2TagId returns the Segment2TagId field if non-nil, zero value otherwise.

### GetSegment2TagIdOk

`func (o *DealUpdateParamsDetails) GetSegment2TagIdOk() (*int32, bool)`

GetSegment2TagIdOk returns a tuple with the Segment2TagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment2TagId

`func (o *DealUpdateParamsDetails) SetSegment2TagId(v int32)`

SetSegment2TagId sets Segment2TagId field to given value.

### HasSegment2TagId

`func (o *DealUpdateParamsDetails) HasSegment2TagId() bool`

HasSegment2TagId returns a boolean if a field has been set.

### GetSegment3TagId

`func (o *DealUpdateParamsDetails) GetSegment3TagId() int32`

GetSegment3TagId returns the Segment3TagId field if non-nil, zero value otherwise.

### GetSegment3TagIdOk

`func (o *DealUpdateParamsDetails) GetSegment3TagIdOk() (*int32, bool)`

GetSegment3TagIdOk returns a tuple with the Segment3TagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment3TagId

`func (o *DealUpdateParamsDetails) SetSegment3TagId(v int32)`

SetSegment3TagId sets Segment3TagId field to given value.

### HasSegment3TagId

`func (o *DealUpdateParamsDetails) HasSegment3TagId() bool`

HasSegment3TagId returns a boolean if a field has been set.

### GetDescription

`func (o *DealUpdateParamsDetails) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DealUpdateParamsDetails) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DealUpdateParamsDetails) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DealUpdateParamsDetails) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVat

`func (o *DealUpdateParamsDetails) GetVat() int32`

GetVat returns the Vat field if non-nil, zero value otherwise.

### GetVatOk

`func (o *DealUpdateParamsDetails) GetVatOk() (*int32, bool)`

GetVatOk returns a tuple with the Vat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVat

`func (o *DealUpdateParamsDetails) SetVat(v int32)`

SetVat sets Vat field to given value.

### HasVat

`func (o *DealUpdateParamsDetails) HasVat() bool`

HasVat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


