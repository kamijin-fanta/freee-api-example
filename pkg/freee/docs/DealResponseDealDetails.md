# DealResponseDealDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | +更新の明細行ID | 
**EntrySide** | **string** | 貸借(貸方: credit, 借方: debit) | 
**AccountItemId** | **int32** | 勘定科目ID | 
**TaxCode** | **int32** | 税区分コード | 
**ItemId** | Pointer to **NullableInt32** | 品目ID | [optional] 
**SectionId** | Pointer to **NullableInt32** | 部門ID | [optional] 
**TagIds** | **[]int32** |  | 
**Segment1TagId** | Pointer to **NullableInt32** | セグメント１ID | [optional] 
**Segment2TagId** | Pointer to **NullableInt32** | セグメント２ID | [optional] 
**Segment3TagId** | Pointer to **NullableInt32** | セグメント３ID | [optional] 
**Amount** | **int32** | 金額（税込で指定してください） | 
**Vat** | **int32** | 消費税額（指定しない場合は自動で計算されます） | 
**Description** | Pointer to **NullableString** | 備考 | [optional] 

## Methods

### NewDealResponseDealDetails

`func NewDealResponseDealDetails(id int32, entrySide string, accountItemId int32, taxCode int32, tagIds []int32, amount int32, vat int32, ) *DealResponseDealDetails`

NewDealResponseDealDetails instantiates a new DealResponseDealDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDealResponseDealDetailsWithDefaults

`func NewDealResponseDealDetailsWithDefaults() *DealResponseDealDetails`

NewDealResponseDealDetailsWithDefaults instantiates a new DealResponseDealDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DealResponseDealDetails) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DealResponseDealDetails) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DealResponseDealDetails) SetId(v int32)`

SetId sets Id field to given value.


### GetEntrySide

`func (o *DealResponseDealDetails) GetEntrySide() string`

GetEntrySide returns the EntrySide field if non-nil, zero value otherwise.

### GetEntrySideOk

`func (o *DealResponseDealDetails) GetEntrySideOk() (*string, bool)`

GetEntrySideOk returns a tuple with the EntrySide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntrySide

`func (o *DealResponseDealDetails) SetEntrySide(v string)`

SetEntrySide sets EntrySide field to given value.


### GetAccountItemId

`func (o *DealResponseDealDetails) GetAccountItemId() int32`

GetAccountItemId returns the AccountItemId field if non-nil, zero value otherwise.

### GetAccountItemIdOk

`func (o *DealResponseDealDetails) GetAccountItemIdOk() (*int32, bool)`

GetAccountItemIdOk returns a tuple with the AccountItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountItemId

`func (o *DealResponseDealDetails) SetAccountItemId(v int32)`

SetAccountItemId sets AccountItemId field to given value.


### GetTaxCode

`func (o *DealResponseDealDetails) GetTaxCode() int32`

GetTaxCode returns the TaxCode field if non-nil, zero value otherwise.

### GetTaxCodeOk

`func (o *DealResponseDealDetails) GetTaxCodeOk() (*int32, bool)`

GetTaxCodeOk returns a tuple with the TaxCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCode

`func (o *DealResponseDealDetails) SetTaxCode(v int32)`

SetTaxCode sets TaxCode field to given value.


### GetItemId

`func (o *DealResponseDealDetails) GetItemId() int32`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *DealResponseDealDetails) GetItemIdOk() (*int32, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *DealResponseDealDetails) SetItemId(v int32)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *DealResponseDealDetails) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### SetItemIdNil

`func (o *DealResponseDealDetails) SetItemIdNil(b bool)`

 SetItemIdNil sets the value for ItemId to be an explicit nil

### UnsetItemId
`func (o *DealResponseDealDetails) UnsetItemId()`

UnsetItemId ensures that no value is present for ItemId, not even an explicit nil
### GetSectionId

`func (o *DealResponseDealDetails) GetSectionId() int32`

GetSectionId returns the SectionId field if non-nil, zero value otherwise.

### GetSectionIdOk

`func (o *DealResponseDealDetails) GetSectionIdOk() (*int32, bool)`

GetSectionIdOk returns a tuple with the SectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionId

`func (o *DealResponseDealDetails) SetSectionId(v int32)`

SetSectionId sets SectionId field to given value.

### HasSectionId

`func (o *DealResponseDealDetails) HasSectionId() bool`

HasSectionId returns a boolean if a field has been set.

### SetSectionIdNil

`func (o *DealResponseDealDetails) SetSectionIdNil(b bool)`

 SetSectionIdNil sets the value for SectionId to be an explicit nil

### UnsetSectionId
`func (o *DealResponseDealDetails) UnsetSectionId()`

UnsetSectionId ensures that no value is present for SectionId, not even an explicit nil
### GetTagIds

`func (o *DealResponseDealDetails) GetTagIds() []int32`

GetTagIds returns the TagIds field if non-nil, zero value otherwise.

### GetTagIdsOk

`func (o *DealResponseDealDetails) GetTagIdsOk() (*[]int32, bool)`

GetTagIdsOk returns a tuple with the TagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagIds

`func (o *DealResponseDealDetails) SetTagIds(v []int32)`

SetTagIds sets TagIds field to given value.


### GetSegment1TagId

`func (o *DealResponseDealDetails) GetSegment1TagId() int32`

GetSegment1TagId returns the Segment1TagId field if non-nil, zero value otherwise.

### GetSegment1TagIdOk

`func (o *DealResponseDealDetails) GetSegment1TagIdOk() (*int32, bool)`

GetSegment1TagIdOk returns a tuple with the Segment1TagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment1TagId

`func (o *DealResponseDealDetails) SetSegment1TagId(v int32)`

SetSegment1TagId sets Segment1TagId field to given value.

### HasSegment1TagId

`func (o *DealResponseDealDetails) HasSegment1TagId() bool`

HasSegment1TagId returns a boolean if a field has been set.

### SetSegment1TagIdNil

`func (o *DealResponseDealDetails) SetSegment1TagIdNil(b bool)`

 SetSegment1TagIdNil sets the value for Segment1TagId to be an explicit nil

### UnsetSegment1TagId
`func (o *DealResponseDealDetails) UnsetSegment1TagId()`

UnsetSegment1TagId ensures that no value is present for Segment1TagId, not even an explicit nil
### GetSegment2TagId

`func (o *DealResponseDealDetails) GetSegment2TagId() int32`

GetSegment2TagId returns the Segment2TagId field if non-nil, zero value otherwise.

### GetSegment2TagIdOk

`func (o *DealResponseDealDetails) GetSegment2TagIdOk() (*int32, bool)`

GetSegment2TagIdOk returns a tuple with the Segment2TagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment2TagId

`func (o *DealResponseDealDetails) SetSegment2TagId(v int32)`

SetSegment2TagId sets Segment2TagId field to given value.

### HasSegment2TagId

`func (o *DealResponseDealDetails) HasSegment2TagId() bool`

HasSegment2TagId returns a boolean if a field has been set.

### SetSegment2TagIdNil

`func (o *DealResponseDealDetails) SetSegment2TagIdNil(b bool)`

 SetSegment2TagIdNil sets the value for Segment2TagId to be an explicit nil

### UnsetSegment2TagId
`func (o *DealResponseDealDetails) UnsetSegment2TagId()`

UnsetSegment2TagId ensures that no value is present for Segment2TagId, not even an explicit nil
### GetSegment3TagId

`func (o *DealResponseDealDetails) GetSegment3TagId() int32`

GetSegment3TagId returns the Segment3TagId field if non-nil, zero value otherwise.

### GetSegment3TagIdOk

`func (o *DealResponseDealDetails) GetSegment3TagIdOk() (*int32, bool)`

GetSegment3TagIdOk returns a tuple with the Segment3TagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment3TagId

`func (o *DealResponseDealDetails) SetSegment3TagId(v int32)`

SetSegment3TagId sets Segment3TagId field to given value.

### HasSegment3TagId

`func (o *DealResponseDealDetails) HasSegment3TagId() bool`

HasSegment3TagId returns a boolean if a field has been set.

### SetSegment3TagIdNil

`func (o *DealResponseDealDetails) SetSegment3TagIdNil(b bool)`

 SetSegment3TagIdNil sets the value for Segment3TagId to be an explicit nil

### UnsetSegment3TagId
`func (o *DealResponseDealDetails) UnsetSegment3TagId()`

UnsetSegment3TagId ensures that no value is present for Segment3TagId, not even an explicit nil
### GetAmount

`func (o *DealResponseDealDetails) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *DealResponseDealDetails) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *DealResponseDealDetails) SetAmount(v int32)`

SetAmount sets Amount field to given value.


### GetVat

`func (o *DealResponseDealDetails) GetVat() int32`

GetVat returns the Vat field if non-nil, zero value otherwise.

### GetVatOk

`func (o *DealResponseDealDetails) GetVatOk() (*int32, bool)`

GetVatOk returns a tuple with the Vat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVat

`func (o *DealResponseDealDetails) SetVat(v int32)`

SetVat sets Vat field to given value.


### GetDescription

`func (o *DealResponseDealDetails) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DealResponseDealDetails) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DealResponseDealDetails) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DealResponseDealDetails) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *DealResponseDealDetails) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *DealResponseDealDetails) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


