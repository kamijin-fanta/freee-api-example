# QuotationIndexResponseQuotationContents

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | 見積内容ID | 
**Order** | **int32** | 順序 | 
**Type** | **string** | 行の種類 | 
**Qty** | **float32** | 数量 | 
**Unit** | **NullableString** | 単位 | 
**UnitPrice** | **float32** | 単価 | 
**Amount** | **int32** | 金額 | 
**Vat** | **int32** | 消費税額 | 
**ReducedVat** | **bool** | 軽減税率税区分（true: 対象、false: 対象外） | 
**Description** | **NullableString** | 備考 | 
**AccountItemId** | **NullableInt32** | 勘定科目ID | 
**AccountItemName** | **NullableString** | 勘定科目名 | 
**TaxCode** | **NullableInt32** | 税区分コード | 
**ItemId** | **NullableInt32** | 品目ID | 
**ItemName** | **NullableString** | 品目 | 
**SectionId** | **NullableInt32** | 部門ID | 
**SectionName** | **NullableString** | 部門 | 
**TagIds** | **[]int32** |  | 
**TagNames** | **[]string** |  | 
**Segment1TagId** | Pointer to **NullableInt32** | セグメント１ID | [optional] 
**Segment1TagName** | Pointer to **NullableString** | セグメント１ID | [optional] 
**Segment2TagId** | Pointer to **NullableInt32** | セグメント２ID | [optional] 
**Segment2TagName** | Pointer to **NullableString** | セグメント２ | [optional] 
**Segment3TagId** | Pointer to **NullableInt32** | セグメント３ID | [optional] 
**Segment3TagName** | Pointer to **NullableString** | セグメント３ | [optional] 

## Methods

### NewQuotationIndexResponseQuotationContents

`func NewQuotationIndexResponseQuotationContents(id int32, order int32, type_ string, qty float32, unit NullableString, unitPrice float32, amount int32, vat int32, reducedVat bool, description NullableString, accountItemId NullableInt32, accountItemName NullableString, taxCode NullableInt32, itemId NullableInt32, itemName NullableString, sectionId NullableInt32, sectionName NullableString, tagIds []int32, tagNames []string, ) *QuotationIndexResponseQuotationContents`

NewQuotationIndexResponseQuotationContents instantiates a new QuotationIndexResponseQuotationContents object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuotationIndexResponseQuotationContentsWithDefaults

`func NewQuotationIndexResponseQuotationContentsWithDefaults() *QuotationIndexResponseQuotationContents`

NewQuotationIndexResponseQuotationContentsWithDefaults instantiates a new QuotationIndexResponseQuotationContents object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QuotationIndexResponseQuotationContents) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QuotationIndexResponseQuotationContents) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QuotationIndexResponseQuotationContents) SetId(v int32)`

SetId sets Id field to given value.


### GetOrder

`func (o *QuotationIndexResponseQuotationContents) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *QuotationIndexResponseQuotationContents) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *QuotationIndexResponseQuotationContents) SetOrder(v int32)`

SetOrder sets Order field to given value.


### GetType

`func (o *QuotationIndexResponseQuotationContents) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *QuotationIndexResponseQuotationContents) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *QuotationIndexResponseQuotationContents) SetType(v string)`

SetType sets Type field to given value.


### GetQty

`func (o *QuotationIndexResponseQuotationContents) GetQty() float32`

GetQty returns the Qty field if non-nil, zero value otherwise.

### GetQtyOk

`func (o *QuotationIndexResponseQuotationContents) GetQtyOk() (*float32, bool)`

GetQtyOk returns a tuple with the Qty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQty

`func (o *QuotationIndexResponseQuotationContents) SetQty(v float32)`

SetQty sets Qty field to given value.


### GetUnit

`func (o *QuotationIndexResponseQuotationContents) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *QuotationIndexResponseQuotationContents) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *QuotationIndexResponseQuotationContents) SetUnit(v string)`

SetUnit sets Unit field to given value.


### SetUnitNil

`func (o *QuotationIndexResponseQuotationContents) SetUnitNil(b bool)`

 SetUnitNil sets the value for Unit to be an explicit nil

### UnsetUnit
`func (o *QuotationIndexResponseQuotationContents) UnsetUnit()`

UnsetUnit ensures that no value is present for Unit, not even an explicit nil
### GetUnitPrice

`func (o *QuotationIndexResponseQuotationContents) GetUnitPrice() float32`

GetUnitPrice returns the UnitPrice field if non-nil, zero value otherwise.

### GetUnitPriceOk

`func (o *QuotationIndexResponseQuotationContents) GetUnitPriceOk() (*float32, bool)`

GetUnitPriceOk returns a tuple with the UnitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPrice

`func (o *QuotationIndexResponseQuotationContents) SetUnitPrice(v float32)`

SetUnitPrice sets UnitPrice field to given value.


### GetAmount

`func (o *QuotationIndexResponseQuotationContents) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *QuotationIndexResponseQuotationContents) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *QuotationIndexResponseQuotationContents) SetAmount(v int32)`

SetAmount sets Amount field to given value.


### GetVat

`func (o *QuotationIndexResponseQuotationContents) GetVat() int32`

GetVat returns the Vat field if non-nil, zero value otherwise.

### GetVatOk

`func (o *QuotationIndexResponseQuotationContents) GetVatOk() (*int32, bool)`

GetVatOk returns a tuple with the Vat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVat

`func (o *QuotationIndexResponseQuotationContents) SetVat(v int32)`

SetVat sets Vat field to given value.


### GetReducedVat

`func (o *QuotationIndexResponseQuotationContents) GetReducedVat() bool`

GetReducedVat returns the ReducedVat field if non-nil, zero value otherwise.

### GetReducedVatOk

`func (o *QuotationIndexResponseQuotationContents) GetReducedVatOk() (*bool, bool)`

GetReducedVatOk returns a tuple with the ReducedVat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReducedVat

`func (o *QuotationIndexResponseQuotationContents) SetReducedVat(v bool)`

SetReducedVat sets ReducedVat field to given value.


### GetDescription

`func (o *QuotationIndexResponseQuotationContents) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *QuotationIndexResponseQuotationContents) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *QuotationIndexResponseQuotationContents) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *QuotationIndexResponseQuotationContents) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *QuotationIndexResponseQuotationContents) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetAccountItemId

`func (o *QuotationIndexResponseQuotationContents) GetAccountItemId() int32`

GetAccountItemId returns the AccountItemId field if non-nil, zero value otherwise.

### GetAccountItemIdOk

`func (o *QuotationIndexResponseQuotationContents) GetAccountItemIdOk() (*int32, bool)`

GetAccountItemIdOk returns a tuple with the AccountItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountItemId

`func (o *QuotationIndexResponseQuotationContents) SetAccountItemId(v int32)`

SetAccountItemId sets AccountItemId field to given value.


### SetAccountItemIdNil

`func (o *QuotationIndexResponseQuotationContents) SetAccountItemIdNil(b bool)`

 SetAccountItemIdNil sets the value for AccountItemId to be an explicit nil

### UnsetAccountItemId
`func (o *QuotationIndexResponseQuotationContents) UnsetAccountItemId()`

UnsetAccountItemId ensures that no value is present for AccountItemId, not even an explicit nil
### GetAccountItemName

`func (o *QuotationIndexResponseQuotationContents) GetAccountItemName() string`

GetAccountItemName returns the AccountItemName field if non-nil, zero value otherwise.

### GetAccountItemNameOk

`func (o *QuotationIndexResponseQuotationContents) GetAccountItemNameOk() (*string, bool)`

GetAccountItemNameOk returns a tuple with the AccountItemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountItemName

`func (o *QuotationIndexResponseQuotationContents) SetAccountItemName(v string)`

SetAccountItemName sets AccountItemName field to given value.


### SetAccountItemNameNil

`func (o *QuotationIndexResponseQuotationContents) SetAccountItemNameNil(b bool)`

 SetAccountItemNameNil sets the value for AccountItemName to be an explicit nil

### UnsetAccountItemName
`func (o *QuotationIndexResponseQuotationContents) UnsetAccountItemName()`

UnsetAccountItemName ensures that no value is present for AccountItemName, not even an explicit nil
### GetTaxCode

`func (o *QuotationIndexResponseQuotationContents) GetTaxCode() int32`

GetTaxCode returns the TaxCode field if non-nil, zero value otherwise.

### GetTaxCodeOk

`func (o *QuotationIndexResponseQuotationContents) GetTaxCodeOk() (*int32, bool)`

GetTaxCodeOk returns a tuple with the TaxCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCode

`func (o *QuotationIndexResponseQuotationContents) SetTaxCode(v int32)`

SetTaxCode sets TaxCode field to given value.


### SetTaxCodeNil

`func (o *QuotationIndexResponseQuotationContents) SetTaxCodeNil(b bool)`

 SetTaxCodeNil sets the value for TaxCode to be an explicit nil

### UnsetTaxCode
`func (o *QuotationIndexResponseQuotationContents) UnsetTaxCode()`

UnsetTaxCode ensures that no value is present for TaxCode, not even an explicit nil
### GetItemId

`func (o *QuotationIndexResponseQuotationContents) GetItemId() int32`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *QuotationIndexResponseQuotationContents) GetItemIdOk() (*int32, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *QuotationIndexResponseQuotationContents) SetItemId(v int32)`

SetItemId sets ItemId field to given value.


### SetItemIdNil

`func (o *QuotationIndexResponseQuotationContents) SetItemIdNil(b bool)`

 SetItemIdNil sets the value for ItemId to be an explicit nil

### UnsetItemId
`func (o *QuotationIndexResponseQuotationContents) UnsetItemId()`

UnsetItemId ensures that no value is present for ItemId, not even an explicit nil
### GetItemName

`func (o *QuotationIndexResponseQuotationContents) GetItemName() string`

GetItemName returns the ItemName field if non-nil, zero value otherwise.

### GetItemNameOk

`func (o *QuotationIndexResponseQuotationContents) GetItemNameOk() (*string, bool)`

GetItemNameOk returns a tuple with the ItemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemName

`func (o *QuotationIndexResponseQuotationContents) SetItemName(v string)`

SetItemName sets ItemName field to given value.


### SetItemNameNil

`func (o *QuotationIndexResponseQuotationContents) SetItemNameNil(b bool)`

 SetItemNameNil sets the value for ItemName to be an explicit nil

### UnsetItemName
`func (o *QuotationIndexResponseQuotationContents) UnsetItemName()`

UnsetItemName ensures that no value is present for ItemName, not even an explicit nil
### GetSectionId

`func (o *QuotationIndexResponseQuotationContents) GetSectionId() int32`

GetSectionId returns the SectionId field if non-nil, zero value otherwise.

### GetSectionIdOk

`func (o *QuotationIndexResponseQuotationContents) GetSectionIdOk() (*int32, bool)`

GetSectionIdOk returns a tuple with the SectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionId

`func (o *QuotationIndexResponseQuotationContents) SetSectionId(v int32)`

SetSectionId sets SectionId field to given value.


### SetSectionIdNil

`func (o *QuotationIndexResponseQuotationContents) SetSectionIdNil(b bool)`

 SetSectionIdNil sets the value for SectionId to be an explicit nil

### UnsetSectionId
`func (o *QuotationIndexResponseQuotationContents) UnsetSectionId()`

UnsetSectionId ensures that no value is present for SectionId, not even an explicit nil
### GetSectionName

`func (o *QuotationIndexResponseQuotationContents) GetSectionName() string`

GetSectionName returns the SectionName field if non-nil, zero value otherwise.

### GetSectionNameOk

`func (o *QuotationIndexResponseQuotationContents) GetSectionNameOk() (*string, bool)`

GetSectionNameOk returns a tuple with the SectionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionName

`func (o *QuotationIndexResponseQuotationContents) SetSectionName(v string)`

SetSectionName sets SectionName field to given value.


### SetSectionNameNil

`func (o *QuotationIndexResponseQuotationContents) SetSectionNameNil(b bool)`

 SetSectionNameNil sets the value for SectionName to be an explicit nil

### UnsetSectionName
`func (o *QuotationIndexResponseQuotationContents) UnsetSectionName()`

UnsetSectionName ensures that no value is present for SectionName, not even an explicit nil
### GetTagIds

`func (o *QuotationIndexResponseQuotationContents) GetTagIds() []int32`

GetTagIds returns the TagIds field if non-nil, zero value otherwise.

### GetTagIdsOk

`func (o *QuotationIndexResponseQuotationContents) GetTagIdsOk() (*[]int32, bool)`

GetTagIdsOk returns a tuple with the TagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagIds

`func (o *QuotationIndexResponseQuotationContents) SetTagIds(v []int32)`

SetTagIds sets TagIds field to given value.


### GetTagNames

`func (o *QuotationIndexResponseQuotationContents) GetTagNames() []string`

GetTagNames returns the TagNames field if non-nil, zero value otherwise.

### GetTagNamesOk

`func (o *QuotationIndexResponseQuotationContents) GetTagNamesOk() (*[]string, bool)`

GetTagNamesOk returns a tuple with the TagNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagNames

`func (o *QuotationIndexResponseQuotationContents) SetTagNames(v []string)`

SetTagNames sets TagNames field to given value.


### GetSegment1TagId

`func (o *QuotationIndexResponseQuotationContents) GetSegment1TagId() int32`

GetSegment1TagId returns the Segment1TagId field if non-nil, zero value otherwise.

### GetSegment1TagIdOk

`func (o *QuotationIndexResponseQuotationContents) GetSegment1TagIdOk() (*int32, bool)`

GetSegment1TagIdOk returns a tuple with the Segment1TagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment1TagId

`func (o *QuotationIndexResponseQuotationContents) SetSegment1TagId(v int32)`

SetSegment1TagId sets Segment1TagId field to given value.

### HasSegment1TagId

`func (o *QuotationIndexResponseQuotationContents) HasSegment1TagId() bool`

HasSegment1TagId returns a boolean if a field has been set.

### SetSegment1TagIdNil

`func (o *QuotationIndexResponseQuotationContents) SetSegment1TagIdNil(b bool)`

 SetSegment1TagIdNil sets the value for Segment1TagId to be an explicit nil

### UnsetSegment1TagId
`func (o *QuotationIndexResponseQuotationContents) UnsetSegment1TagId()`

UnsetSegment1TagId ensures that no value is present for Segment1TagId, not even an explicit nil
### GetSegment1TagName

`func (o *QuotationIndexResponseQuotationContents) GetSegment1TagName() string`

GetSegment1TagName returns the Segment1TagName field if non-nil, zero value otherwise.

### GetSegment1TagNameOk

`func (o *QuotationIndexResponseQuotationContents) GetSegment1TagNameOk() (*string, bool)`

GetSegment1TagNameOk returns a tuple with the Segment1TagName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment1TagName

`func (o *QuotationIndexResponseQuotationContents) SetSegment1TagName(v string)`

SetSegment1TagName sets Segment1TagName field to given value.

### HasSegment1TagName

`func (o *QuotationIndexResponseQuotationContents) HasSegment1TagName() bool`

HasSegment1TagName returns a boolean if a field has been set.

### SetSegment1TagNameNil

`func (o *QuotationIndexResponseQuotationContents) SetSegment1TagNameNil(b bool)`

 SetSegment1TagNameNil sets the value for Segment1TagName to be an explicit nil

### UnsetSegment1TagName
`func (o *QuotationIndexResponseQuotationContents) UnsetSegment1TagName()`

UnsetSegment1TagName ensures that no value is present for Segment1TagName, not even an explicit nil
### GetSegment2TagId

`func (o *QuotationIndexResponseQuotationContents) GetSegment2TagId() int32`

GetSegment2TagId returns the Segment2TagId field if non-nil, zero value otherwise.

### GetSegment2TagIdOk

`func (o *QuotationIndexResponseQuotationContents) GetSegment2TagIdOk() (*int32, bool)`

GetSegment2TagIdOk returns a tuple with the Segment2TagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment2TagId

`func (o *QuotationIndexResponseQuotationContents) SetSegment2TagId(v int32)`

SetSegment2TagId sets Segment2TagId field to given value.

### HasSegment2TagId

`func (o *QuotationIndexResponseQuotationContents) HasSegment2TagId() bool`

HasSegment2TagId returns a boolean if a field has been set.

### SetSegment2TagIdNil

`func (o *QuotationIndexResponseQuotationContents) SetSegment2TagIdNil(b bool)`

 SetSegment2TagIdNil sets the value for Segment2TagId to be an explicit nil

### UnsetSegment2TagId
`func (o *QuotationIndexResponseQuotationContents) UnsetSegment2TagId()`

UnsetSegment2TagId ensures that no value is present for Segment2TagId, not even an explicit nil
### GetSegment2TagName

`func (o *QuotationIndexResponseQuotationContents) GetSegment2TagName() string`

GetSegment2TagName returns the Segment2TagName field if non-nil, zero value otherwise.

### GetSegment2TagNameOk

`func (o *QuotationIndexResponseQuotationContents) GetSegment2TagNameOk() (*string, bool)`

GetSegment2TagNameOk returns a tuple with the Segment2TagName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment2TagName

`func (o *QuotationIndexResponseQuotationContents) SetSegment2TagName(v string)`

SetSegment2TagName sets Segment2TagName field to given value.

### HasSegment2TagName

`func (o *QuotationIndexResponseQuotationContents) HasSegment2TagName() bool`

HasSegment2TagName returns a boolean if a field has been set.

### SetSegment2TagNameNil

`func (o *QuotationIndexResponseQuotationContents) SetSegment2TagNameNil(b bool)`

 SetSegment2TagNameNil sets the value for Segment2TagName to be an explicit nil

### UnsetSegment2TagName
`func (o *QuotationIndexResponseQuotationContents) UnsetSegment2TagName()`

UnsetSegment2TagName ensures that no value is present for Segment2TagName, not even an explicit nil
### GetSegment3TagId

`func (o *QuotationIndexResponseQuotationContents) GetSegment3TagId() int32`

GetSegment3TagId returns the Segment3TagId field if non-nil, zero value otherwise.

### GetSegment3TagIdOk

`func (o *QuotationIndexResponseQuotationContents) GetSegment3TagIdOk() (*int32, bool)`

GetSegment3TagIdOk returns a tuple with the Segment3TagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment3TagId

`func (o *QuotationIndexResponseQuotationContents) SetSegment3TagId(v int32)`

SetSegment3TagId sets Segment3TagId field to given value.

### HasSegment3TagId

`func (o *QuotationIndexResponseQuotationContents) HasSegment3TagId() bool`

HasSegment3TagId returns a boolean if a field has been set.

### SetSegment3TagIdNil

`func (o *QuotationIndexResponseQuotationContents) SetSegment3TagIdNil(b bool)`

 SetSegment3TagIdNil sets the value for Segment3TagId to be an explicit nil

### UnsetSegment3TagId
`func (o *QuotationIndexResponseQuotationContents) UnsetSegment3TagId()`

UnsetSegment3TagId ensures that no value is present for Segment3TagId, not even an explicit nil
### GetSegment3TagName

`func (o *QuotationIndexResponseQuotationContents) GetSegment3TagName() string`

GetSegment3TagName returns the Segment3TagName field if non-nil, zero value otherwise.

### GetSegment3TagNameOk

`func (o *QuotationIndexResponseQuotationContents) GetSegment3TagNameOk() (*string, bool)`

GetSegment3TagNameOk returns a tuple with the Segment3TagName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment3TagName

`func (o *QuotationIndexResponseQuotationContents) SetSegment3TagName(v string)`

SetSegment3TagName sets Segment3TagName field to given value.

### HasSegment3TagName

`func (o *QuotationIndexResponseQuotationContents) HasSegment3TagName() bool`

HasSegment3TagName returns a boolean if a field has been set.

### SetSegment3TagNameNil

`func (o *QuotationIndexResponseQuotationContents) SetSegment3TagNameNil(b bool)`

 SetSegment3TagNameNil sets the value for Segment3TagName to be an explicit nil

### UnsetSegment3TagName
`func (o *QuotationIndexResponseQuotationContents) UnsetSegment3TagName()`

UnsetSegment3TagName ensures that no value is present for Segment3TagName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


