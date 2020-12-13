# ManualJournalUpdateParamsDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | 貸借行ID: 既存貸借行を更新または削除する場合に指定します。IDを指定しない貸借行は、新規行として扱われ追加されます。 | [optional] 
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

### NewManualJournalUpdateParamsDetails

`func NewManualJournalUpdateParamsDetails(entrySide string, taxCode int32, accountItemId int32, amount int32, ) *ManualJournalUpdateParamsDetails`

NewManualJournalUpdateParamsDetails instantiates a new ManualJournalUpdateParamsDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManualJournalUpdateParamsDetailsWithDefaults

`func NewManualJournalUpdateParamsDetailsWithDefaults() *ManualJournalUpdateParamsDetails`

NewManualJournalUpdateParamsDetailsWithDefaults instantiates a new ManualJournalUpdateParamsDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ManualJournalUpdateParamsDetails) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ManualJournalUpdateParamsDetails) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ManualJournalUpdateParamsDetails) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ManualJournalUpdateParamsDetails) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEntrySide

`func (o *ManualJournalUpdateParamsDetails) GetEntrySide() string`

GetEntrySide returns the EntrySide field if non-nil, zero value otherwise.

### GetEntrySideOk

`func (o *ManualJournalUpdateParamsDetails) GetEntrySideOk() (*string, bool)`

GetEntrySideOk returns a tuple with the EntrySide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntrySide

`func (o *ManualJournalUpdateParamsDetails) SetEntrySide(v string)`

SetEntrySide sets EntrySide field to given value.


### GetTaxCode

`func (o *ManualJournalUpdateParamsDetails) GetTaxCode() int32`

GetTaxCode returns the TaxCode field if non-nil, zero value otherwise.

### GetTaxCodeOk

`func (o *ManualJournalUpdateParamsDetails) GetTaxCodeOk() (*int32, bool)`

GetTaxCodeOk returns a tuple with the TaxCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCode

`func (o *ManualJournalUpdateParamsDetails) SetTaxCode(v int32)`

SetTaxCode sets TaxCode field to given value.


### GetAccountItemId

`func (o *ManualJournalUpdateParamsDetails) GetAccountItemId() int32`

GetAccountItemId returns the AccountItemId field if non-nil, zero value otherwise.

### GetAccountItemIdOk

`func (o *ManualJournalUpdateParamsDetails) GetAccountItemIdOk() (*int32, bool)`

GetAccountItemIdOk returns a tuple with the AccountItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountItemId

`func (o *ManualJournalUpdateParamsDetails) SetAccountItemId(v int32)`

SetAccountItemId sets AccountItemId field to given value.


### GetAmount

`func (o *ManualJournalUpdateParamsDetails) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ManualJournalUpdateParamsDetails) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ManualJournalUpdateParamsDetails) SetAmount(v int32)`

SetAmount sets Amount field to given value.


### GetVat

`func (o *ManualJournalUpdateParamsDetails) GetVat() int32`

GetVat returns the Vat field if non-nil, zero value otherwise.

### GetVatOk

`func (o *ManualJournalUpdateParamsDetails) GetVatOk() (*int32, bool)`

GetVatOk returns a tuple with the Vat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVat

`func (o *ManualJournalUpdateParamsDetails) SetVat(v int32)`

SetVat sets Vat field to given value.

### HasVat

`func (o *ManualJournalUpdateParamsDetails) HasVat() bool`

HasVat returns a boolean if a field has been set.

### GetPartnerId

`func (o *ManualJournalUpdateParamsDetails) GetPartnerId() int32`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *ManualJournalUpdateParamsDetails) GetPartnerIdOk() (*int32, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *ManualJournalUpdateParamsDetails) SetPartnerId(v int32)`

SetPartnerId sets PartnerId field to given value.

### HasPartnerId

`func (o *ManualJournalUpdateParamsDetails) HasPartnerId() bool`

HasPartnerId returns a boolean if a field has been set.

### GetPartnerCode

`func (o *ManualJournalUpdateParamsDetails) GetPartnerCode() string`

GetPartnerCode returns the PartnerCode field if non-nil, zero value otherwise.

### GetPartnerCodeOk

`func (o *ManualJournalUpdateParamsDetails) GetPartnerCodeOk() (*string, bool)`

GetPartnerCodeOk returns a tuple with the PartnerCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerCode

`func (o *ManualJournalUpdateParamsDetails) SetPartnerCode(v string)`

SetPartnerCode sets PartnerCode field to given value.

### HasPartnerCode

`func (o *ManualJournalUpdateParamsDetails) HasPartnerCode() bool`

HasPartnerCode returns a boolean if a field has been set.

### GetItemId

`func (o *ManualJournalUpdateParamsDetails) GetItemId() int32`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *ManualJournalUpdateParamsDetails) GetItemIdOk() (*int32, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *ManualJournalUpdateParamsDetails) SetItemId(v int32)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *ManualJournalUpdateParamsDetails) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### GetSectionId

`func (o *ManualJournalUpdateParamsDetails) GetSectionId() int32`

GetSectionId returns the SectionId field if non-nil, zero value otherwise.

### GetSectionIdOk

`func (o *ManualJournalUpdateParamsDetails) GetSectionIdOk() (*int32, bool)`

GetSectionIdOk returns a tuple with the SectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionId

`func (o *ManualJournalUpdateParamsDetails) SetSectionId(v int32)`

SetSectionId sets SectionId field to given value.

### HasSectionId

`func (o *ManualJournalUpdateParamsDetails) HasSectionId() bool`

HasSectionId returns a boolean if a field has been set.

### GetTagIds

`func (o *ManualJournalUpdateParamsDetails) GetTagIds() []int32`

GetTagIds returns the TagIds field if non-nil, zero value otherwise.

### GetTagIdsOk

`func (o *ManualJournalUpdateParamsDetails) GetTagIdsOk() (*[]int32, bool)`

GetTagIdsOk returns a tuple with the TagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagIds

`func (o *ManualJournalUpdateParamsDetails) SetTagIds(v []int32)`

SetTagIds sets TagIds field to given value.

### HasTagIds

`func (o *ManualJournalUpdateParamsDetails) HasTagIds() bool`

HasTagIds returns a boolean if a field has been set.

### GetSegment1TagId

`func (o *ManualJournalUpdateParamsDetails) GetSegment1TagId() int32`

GetSegment1TagId returns the Segment1TagId field if non-nil, zero value otherwise.

### GetSegment1TagIdOk

`func (o *ManualJournalUpdateParamsDetails) GetSegment1TagIdOk() (*int32, bool)`

GetSegment1TagIdOk returns a tuple with the Segment1TagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment1TagId

`func (o *ManualJournalUpdateParamsDetails) SetSegment1TagId(v int32)`

SetSegment1TagId sets Segment1TagId field to given value.

### HasSegment1TagId

`func (o *ManualJournalUpdateParamsDetails) HasSegment1TagId() bool`

HasSegment1TagId returns a boolean if a field has been set.

### GetSegment2TagId

`func (o *ManualJournalUpdateParamsDetails) GetSegment2TagId() int32`

GetSegment2TagId returns the Segment2TagId field if non-nil, zero value otherwise.

### GetSegment2TagIdOk

`func (o *ManualJournalUpdateParamsDetails) GetSegment2TagIdOk() (*int32, bool)`

GetSegment2TagIdOk returns a tuple with the Segment2TagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment2TagId

`func (o *ManualJournalUpdateParamsDetails) SetSegment2TagId(v int32)`

SetSegment2TagId sets Segment2TagId field to given value.

### HasSegment2TagId

`func (o *ManualJournalUpdateParamsDetails) HasSegment2TagId() bool`

HasSegment2TagId returns a boolean if a field has been set.

### GetSegment3TagId

`func (o *ManualJournalUpdateParamsDetails) GetSegment3TagId() int32`

GetSegment3TagId returns the Segment3TagId field if non-nil, zero value otherwise.

### GetSegment3TagIdOk

`func (o *ManualJournalUpdateParamsDetails) GetSegment3TagIdOk() (*int32, bool)`

GetSegment3TagIdOk returns a tuple with the Segment3TagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment3TagId

`func (o *ManualJournalUpdateParamsDetails) SetSegment3TagId(v int32)`

SetSegment3TagId sets Segment3TagId field to given value.

### HasSegment3TagId

`func (o *ManualJournalUpdateParamsDetails) HasSegment3TagId() bool`

HasSegment3TagId returns a boolean if a field has been set.

### GetDescription

`func (o *ManualJournalUpdateParamsDetails) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ManualJournalUpdateParamsDetails) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ManualJournalUpdateParamsDetails) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ManualJournalUpdateParamsDetails) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


