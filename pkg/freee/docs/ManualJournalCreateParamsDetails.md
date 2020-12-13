# ManualJournalCreateParamsDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntrySide** | **string** | 貸借（貸方: credit, 借方: debit） | 
**TaxCode** | **int32** | 税区分コード | 
**AccountItemId** | **int32** | 勘定科目ID | 
**Amount** | **int32** | 取引金額（税込で指定してください） | 
**Vat** | Pointer to **int32** | 消費税額（指定しない場合は自動で計算されます） | [optional] 
**PartnerId** | Pointer to **int32** | 取引先ID | [optional] 
**PartnerCode** | Pointer to **string** | 取引先コード | [optional] 
**ItemId** | Pointer to **int32** | 品目ID | [optional] 
**SectionId** | Pointer to **int32** | 部門ID | [optional] 
**TagIds** | Pointer to **[]int32** | メモタグID | [optional] 
**Segment1TagId** | Pointer to **int32** | セグメント１ID | [optional] 
**Segment2TagId** | Pointer to **int32** | セグメント２ID | [optional] 
**Segment3TagId** | Pointer to **int32** | セグメント３ID | [optional] 
**Description** | Pointer to **string** | 備考 | [optional] 

## Methods

### NewManualJournalCreateParamsDetails

`func NewManualJournalCreateParamsDetails(entrySide string, taxCode int32, accountItemId int32, amount int32, ) *ManualJournalCreateParamsDetails`

NewManualJournalCreateParamsDetails instantiates a new ManualJournalCreateParamsDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManualJournalCreateParamsDetailsWithDefaults

`func NewManualJournalCreateParamsDetailsWithDefaults() *ManualJournalCreateParamsDetails`

NewManualJournalCreateParamsDetailsWithDefaults instantiates a new ManualJournalCreateParamsDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntrySide

`func (o *ManualJournalCreateParamsDetails) GetEntrySide() string`

GetEntrySide returns the EntrySide field if non-nil, zero value otherwise.

### GetEntrySideOk

`func (o *ManualJournalCreateParamsDetails) GetEntrySideOk() (*string, bool)`

GetEntrySideOk returns a tuple with the EntrySide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntrySide

`func (o *ManualJournalCreateParamsDetails) SetEntrySide(v string)`

SetEntrySide sets EntrySide field to given value.


### GetTaxCode

`func (o *ManualJournalCreateParamsDetails) GetTaxCode() int32`

GetTaxCode returns the TaxCode field if non-nil, zero value otherwise.

### GetTaxCodeOk

`func (o *ManualJournalCreateParamsDetails) GetTaxCodeOk() (*int32, bool)`

GetTaxCodeOk returns a tuple with the TaxCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCode

`func (o *ManualJournalCreateParamsDetails) SetTaxCode(v int32)`

SetTaxCode sets TaxCode field to given value.


### GetAccountItemId

`func (o *ManualJournalCreateParamsDetails) GetAccountItemId() int32`

GetAccountItemId returns the AccountItemId field if non-nil, zero value otherwise.

### GetAccountItemIdOk

`func (o *ManualJournalCreateParamsDetails) GetAccountItemIdOk() (*int32, bool)`

GetAccountItemIdOk returns a tuple with the AccountItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountItemId

`func (o *ManualJournalCreateParamsDetails) SetAccountItemId(v int32)`

SetAccountItemId sets AccountItemId field to given value.


### GetAmount

`func (o *ManualJournalCreateParamsDetails) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ManualJournalCreateParamsDetails) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ManualJournalCreateParamsDetails) SetAmount(v int32)`

SetAmount sets Amount field to given value.


### GetVat

`func (o *ManualJournalCreateParamsDetails) GetVat() int32`

GetVat returns the Vat field if non-nil, zero value otherwise.

### GetVatOk

`func (o *ManualJournalCreateParamsDetails) GetVatOk() (*int32, bool)`

GetVatOk returns a tuple with the Vat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVat

`func (o *ManualJournalCreateParamsDetails) SetVat(v int32)`

SetVat sets Vat field to given value.

### HasVat

`func (o *ManualJournalCreateParamsDetails) HasVat() bool`

HasVat returns a boolean if a field has been set.

### GetPartnerId

`func (o *ManualJournalCreateParamsDetails) GetPartnerId() int32`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *ManualJournalCreateParamsDetails) GetPartnerIdOk() (*int32, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *ManualJournalCreateParamsDetails) SetPartnerId(v int32)`

SetPartnerId sets PartnerId field to given value.

### HasPartnerId

`func (o *ManualJournalCreateParamsDetails) HasPartnerId() bool`

HasPartnerId returns a boolean if a field has been set.

### GetPartnerCode

`func (o *ManualJournalCreateParamsDetails) GetPartnerCode() string`

GetPartnerCode returns the PartnerCode field if non-nil, zero value otherwise.

### GetPartnerCodeOk

`func (o *ManualJournalCreateParamsDetails) GetPartnerCodeOk() (*string, bool)`

GetPartnerCodeOk returns a tuple with the PartnerCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerCode

`func (o *ManualJournalCreateParamsDetails) SetPartnerCode(v string)`

SetPartnerCode sets PartnerCode field to given value.

### HasPartnerCode

`func (o *ManualJournalCreateParamsDetails) HasPartnerCode() bool`

HasPartnerCode returns a boolean if a field has been set.

### GetItemId

`func (o *ManualJournalCreateParamsDetails) GetItemId() int32`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *ManualJournalCreateParamsDetails) GetItemIdOk() (*int32, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *ManualJournalCreateParamsDetails) SetItemId(v int32)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *ManualJournalCreateParamsDetails) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### GetSectionId

`func (o *ManualJournalCreateParamsDetails) GetSectionId() int32`

GetSectionId returns the SectionId field if non-nil, zero value otherwise.

### GetSectionIdOk

`func (o *ManualJournalCreateParamsDetails) GetSectionIdOk() (*int32, bool)`

GetSectionIdOk returns a tuple with the SectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionId

`func (o *ManualJournalCreateParamsDetails) SetSectionId(v int32)`

SetSectionId sets SectionId field to given value.

### HasSectionId

`func (o *ManualJournalCreateParamsDetails) HasSectionId() bool`

HasSectionId returns a boolean if a field has been set.

### GetTagIds

`func (o *ManualJournalCreateParamsDetails) GetTagIds() []int32`

GetTagIds returns the TagIds field if non-nil, zero value otherwise.

### GetTagIdsOk

`func (o *ManualJournalCreateParamsDetails) GetTagIdsOk() (*[]int32, bool)`

GetTagIdsOk returns a tuple with the TagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagIds

`func (o *ManualJournalCreateParamsDetails) SetTagIds(v []int32)`

SetTagIds sets TagIds field to given value.

### HasTagIds

`func (o *ManualJournalCreateParamsDetails) HasTagIds() bool`

HasTagIds returns a boolean if a field has been set.

### GetSegment1TagId

`func (o *ManualJournalCreateParamsDetails) GetSegment1TagId() int32`

GetSegment1TagId returns the Segment1TagId field if non-nil, zero value otherwise.

### GetSegment1TagIdOk

`func (o *ManualJournalCreateParamsDetails) GetSegment1TagIdOk() (*int32, bool)`

GetSegment1TagIdOk returns a tuple with the Segment1TagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment1TagId

`func (o *ManualJournalCreateParamsDetails) SetSegment1TagId(v int32)`

SetSegment1TagId sets Segment1TagId field to given value.

### HasSegment1TagId

`func (o *ManualJournalCreateParamsDetails) HasSegment1TagId() bool`

HasSegment1TagId returns a boolean if a field has been set.

### GetSegment2TagId

`func (o *ManualJournalCreateParamsDetails) GetSegment2TagId() int32`

GetSegment2TagId returns the Segment2TagId field if non-nil, zero value otherwise.

### GetSegment2TagIdOk

`func (o *ManualJournalCreateParamsDetails) GetSegment2TagIdOk() (*int32, bool)`

GetSegment2TagIdOk returns a tuple with the Segment2TagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment2TagId

`func (o *ManualJournalCreateParamsDetails) SetSegment2TagId(v int32)`

SetSegment2TagId sets Segment2TagId field to given value.

### HasSegment2TagId

`func (o *ManualJournalCreateParamsDetails) HasSegment2TagId() bool`

HasSegment2TagId returns a boolean if a field has been set.

### GetSegment3TagId

`func (o *ManualJournalCreateParamsDetails) GetSegment3TagId() int32`

GetSegment3TagId returns the Segment3TagId field if non-nil, zero value otherwise.

### GetSegment3TagIdOk

`func (o *ManualJournalCreateParamsDetails) GetSegment3TagIdOk() (*int32, bool)`

GetSegment3TagIdOk returns a tuple with the Segment3TagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment3TagId

`func (o *ManualJournalCreateParamsDetails) SetSegment3TagId(v int32)`

SetSegment3TagId sets Segment3TagId field to given value.

### HasSegment3TagId

`func (o *ManualJournalCreateParamsDetails) HasSegment3TagId() bool`

HasSegment3TagId returns a boolean if a field has been set.

### GetDescription

`func (o *ManualJournalCreateParamsDetails) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ManualJournalCreateParamsDetails) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ManualJournalCreateParamsDetails) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ManualJournalCreateParamsDetails) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


