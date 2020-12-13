# DealCreateResponseDealDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | 取引行ID | 
**AccountItemId** | **int32** | 勘定科目ID | 
**TaxCode** | **int32** | 税区分コード | 
**ItemId** | Pointer to **NullableInt32** | 品目ID | [optional] 
**SectionId** | Pointer to **NullableInt32** | 部門ID | [optional] 
**TagIds** | Pointer to **[]int32** | メモタグID | [optional] 
**Segment1TagId** | Pointer to **NullableInt32** | セグメント１ID | [optional] 
**Segment2TagId** | Pointer to **NullableInt32** | セグメント２ID | [optional] 
**Segment3TagId** | Pointer to **NullableInt32** | セグメント３ID | [optional] 
**Amount** | **int32** | 取引金額 | 
**Vat** | **int32** | 消費税額 | 
**Description** | Pointer to **string** | 備考 | [optional] 
**EntrySide** | **string** | 貸借（貸方: credit, 借方: debit） | 

## Methods

### NewDealCreateResponseDealDetails

`func NewDealCreateResponseDealDetails(id int32, accountItemId int32, taxCode int32, amount int32, vat int32, entrySide string, ) *DealCreateResponseDealDetails`

NewDealCreateResponseDealDetails instantiates a new DealCreateResponseDealDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDealCreateResponseDealDetailsWithDefaults

`func NewDealCreateResponseDealDetailsWithDefaults() *DealCreateResponseDealDetails`

NewDealCreateResponseDealDetailsWithDefaults instantiates a new DealCreateResponseDealDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DealCreateResponseDealDetails) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DealCreateResponseDealDetails) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DealCreateResponseDealDetails) SetId(v int32)`

SetId sets Id field to given value.


### GetAccountItemId

`func (o *DealCreateResponseDealDetails) GetAccountItemId() int32`

GetAccountItemId returns the AccountItemId field if non-nil, zero value otherwise.

### GetAccountItemIdOk

`func (o *DealCreateResponseDealDetails) GetAccountItemIdOk() (*int32, bool)`

GetAccountItemIdOk returns a tuple with the AccountItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountItemId

`func (o *DealCreateResponseDealDetails) SetAccountItemId(v int32)`

SetAccountItemId sets AccountItemId field to given value.


### GetTaxCode

`func (o *DealCreateResponseDealDetails) GetTaxCode() int32`

GetTaxCode returns the TaxCode field if non-nil, zero value otherwise.

### GetTaxCodeOk

`func (o *DealCreateResponseDealDetails) GetTaxCodeOk() (*int32, bool)`

GetTaxCodeOk returns a tuple with the TaxCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCode

`func (o *DealCreateResponseDealDetails) SetTaxCode(v int32)`

SetTaxCode sets TaxCode field to given value.


### GetItemId

`func (o *DealCreateResponseDealDetails) GetItemId() int32`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *DealCreateResponseDealDetails) GetItemIdOk() (*int32, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *DealCreateResponseDealDetails) SetItemId(v int32)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *DealCreateResponseDealDetails) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### SetItemIdNil

`func (o *DealCreateResponseDealDetails) SetItemIdNil(b bool)`

 SetItemIdNil sets the value for ItemId to be an explicit nil

### UnsetItemId
`func (o *DealCreateResponseDealDetails) UnsetItemId()`

UnsetItemId ensures that no value is present for ItemId, not even an explicit nil
### GetSectionId

`func (o *DealCreateResponseDealDetails) GetSectionId() int32`

GetSectionId returns the SectionId field if non-nil, zero value otherwise.

### GetSectionIdOk

`func (o *DealCreateResponseDealDetails) GetSectionIdOk() (*int32, bool)`

GetSectionIdOk returns a tuple with the SectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionId

`func (o *DealCreateResponseDealDetails) SetSectionId(v int32)`

SetSectionId sets SectionId field to given value.

### HasSectionId

`func (o *DealCreateResponseDealDetails) HasSectionId() bool`

HasSectionId returns a boolean if a field has been set.

### SetSectionIdNil

`func (o *DealCreateResponseDealDetails) SetSectionIdNil(b bool)`

 SetSectionIdNil sets the value for SectionId to be an explicit nil

### UnsetSectionId
`func (o *DealCreateResponseDealDetails) UnsetSectionId()`

UnsetSectionId ensures that no value is present for SectionId, not even an explicit nil
### GetTagIds

`func (o *DealCreateResponseDealDetails) GetTagIds() []int32`

GetTagIds returns the TagIds field if non-nil, zero value otherwise.

### GetTagIdsOk

`func (o *DealCreateResponseDealDetails) GetTagIdsOk() (*[]int32, bool)`

GetTagIdsOk returns a tuple with the TagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagIds

`func (o *DealCreateResponseDealDetails) SetTagIds(v []int32)`

SetTagIds sets TagIds field to given value.

### HasTagIds

`func (o *DealCreateResponseDealDetails) HasTagIds() bool`

HasTagIds returns a boolean if a field has been set.

### GetSegment1TagId

`func (o *DealCreateResponseDealDetails) GetSegment1TagId() int32`

GetSegment1TagId returns the Segment1TagId field if non-nil, zero value otherwise.

### GetSegment1TagIdOk

`func (o *DealCreateResponseDealDetails) GetSegment1TagIdOk() (*int32, bool)`

GetSegment1TagIdOk returns a tuple with the Segment1TagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment1TagId

`func (o *DealCreateResponseDealDetails) SetSegment1TagId(v int32)`

SetSegment1TagId sets Segment1TagId field to given value.

### HasSegment1TagId

`func (o *DealCreateResponseDealDetails) HasSegment1TagId() bool`

HasSegment1TagId returns a boolean if a field has been set.

### SetSegment1TagIdNil

`func (o *DealCreateResponseDealDetails) SetSegment1TagIdNil(b bool)`

 SetSegment1TagIdNil sets the value for Segment1TagId to be an explicit nil

### UnsetSegment1TagId
`func (o *DealCreateResponseDealDetails) UnsetSegment1TagId()`

UnsetSegment1TagId ensures that no value is present for Segment1TagId, not even an explicit nil
### GetSegment2TagId

`func (o *DealCreateResponseDealDetails) GetSegment2TagId() int32`

GetSegment2TagId returns the Segment2TagId field if non-nil, zero value otherwise.

### GetSegment2TagIdOk

`func (o *DealCreateResponseDealDetails) GetSegment2TagIdOk() (*int32, bool)`

GetSegment2TagIdOk returns a tuple with the Segment2TagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment2TagId

`func (o *DealCreateResponseDealDetails) SetSegment2TagId(v int32)`

SetSegment2TagId sets Segment2TagId field to given value.

### HasSegment2TagId

`func (o *DealCreateResponseDealDetails) HasSegment2TagId() bool`

HasSegment2TagId returns a boolean if a field has been set.

### SetSegment2TagIdNil

`func (o *DealCreateResponseDealDetails) SetSegment2TagIdNil(b bool)`

 SetSegment2TagIdNil sets the value for Segment2TagId to be an explicit nil

### UnsetSegment2TagId
`func (o *DealCreateResponseDealDetails) UnsetSegment2TagId()`

UnsetSegment2TagId ensures that no value is present for Segment2TagId, not even an explicit nil
### GetSegment3TagId

`func (o *DealCreateResponseDealDetails) GetSegment3TagId() int32`

GetSegment3TagId returns the Segment3TagId field if non-nil, zero value otherwise.

### GetSegment3TagIdOk

`func (o *DealCreateResponseDealDetails) GetSegment3TagIdOk() (*int32, bool)`

GetSegment3TagIdOk returns a tuple with the Segment3TagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment3TagId

`func (o *DealCreateResponseDealDetails) SetSegment3TagId(v int32)`

SetSegment3TagId sets Segment3TagId field to given value.

### HasSegment3TagId

`func (o *DealCreateResponseDealDetails) HasSegment3TagId() bool`

HasSegment3TagId returns a boolean if a field has been set.

### SetSegment3TagIdNil

`func (o *DealCreateResponseDealDetails) SetSegment3TagIdNil(b bool)`

 SetSegment3TagIdNil sets the value for Segment3TagId to be an explicit nil

### UnsetSegment3TagId
`func (o *DealCreateResponseDealDetails) UnsetSegment3TagId()`

UnsetSegment3TagId ensures that no value is present for Segment3TagId, not even an explicit nil
### GetAmount

`func (o *DealCreateResponseDealDetails) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *DealCreateResponseDealDetails) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *DealCreateResponseDealDetails) SetAmount(v int32)`

SetAmount sets Amount field to given value.


### GetVat

`func (o *DealCreateResponseDealDetails) GetVat() int32`

GetVat returns the Vat field if non-nil, zero value otherwise.

### GetVatOk

`func (o *DealCreateResponseDealDetails) GetVatOk() (*int32, bool)`

GetVatOk returns a tuple with the Vat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVat

`func (o *DealCreateResponseDealDetails) SetVat(v int32)`

SetVat sets Vat field to given value.


### GetDescription

`func (o *DealCreateResponseDealDetails) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DealCreateResponseDealDetails) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DealCreateResponseDealDetails) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DealCreateResponseDealDetails) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEntrySide

`func (o *DealCreateResponseDealDetails) GetEntrySide() string`

GetEntrySide returns the EntrySide field if non-nil, zero value otherwise.

### GetEntrySideOk

`func (o *DealCreateResponseDealDetails) GetEntrySideOk() (*string, bool)`

GetEntrySideOk returns a tuple with the EntrySide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntrySide

`func (o *DealCreateResponseDealDetails) SetEntrySide(v string)`

SetEntrySide sets EntrySide field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


