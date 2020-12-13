# InvoiceCreateParamsInvoiceContents

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Order** | **int32** | 順序 | 
**Type** | **string** | 行の種類 &lt;ul&gt; &lt;li&gt;normal、discountを指定する場合、account_item_id,tax_codeとunit_priceが必須となります。&lt;/li&gt; &lt;li&gt;normalを指定した場合、qtyが必須となります。&lt;/li&gt; &lt;/ul&gt; | 
**Qty** | Pointer to **float32** | 数量 | [optional] 
**Unit** | Pointer to **string** | 単位 | [optional] 
**UnitPrice** | Pointer to **float32** | 単価 (tax_entry_method: inclusiveの場合は税込価格、tax_entry_method: exclusiveの場合は税抜価格となります) | [optional] 
**Vat** | Pointer to **NullableInt32** | 消費税額 | [optional] 
**Description** | Pointer to **string** | 備考 | [optional] 
**AccountItemId** | Pointer to **int32** | 勘定科目ID | [optional] 
**TaxCode** | Pointer to **int32** | 税区分コード | [optional] 
**ItemId** | Pointer to **int32** | 品目ID | [optional] 
**SectionId** | Pointer to **int32** | 部門ID | [optional] 
**TagIds** | Pointer to **[]int32** |  | [optional] 
**Segment1TagId** | Pointer to **int32** | セグメント１ID | [optional] 
**Segment2TagId** | Pointer to **int32** | セグメント２ID | [optional] 
**Segment3TagId** | Pointer to **int32** | セグメント３ID | [optional] 

## Methods

### NewInvoiceCreateParamsInvoiceContents

`func NewInvoiceCreateParamsInvoiceContents(order int32, type_ string, ) *InvoiceCreateParamsInvoiceContents`

NewInvoiceCreateParamsInvoiceContents instantiates a new InvoiceCreateParamsInvoiceContents object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceCreateParamsInvoiceContentsWithDefaults

`func NewInvoiceCreateParamsInvoiceContentsWithDefaults() *InvoiceCreateParamsInvoiceContents`

NewInvoiceCreateParamsInvoiceContentsWithDefaults instantiates a new InvoiceCreateParamsInvoiceContents object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrder

`func (o *InvoiceCreateParamsInvoiceContents) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *InvoiceCreateParamsInvoiceContents) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *InvoiceCreateParamsInvoiceContents) SetOrder(v int32)`

SetOrder sets Order field to given value.


### GetType

`func (o *InvoiceCreateParamsInvoiceContents) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InvoiceCreateParamsInvoiceContents) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InvoiceCreateParamsInvoiceContents) SetType(v string)`

SetType sets Type field to given value.


### GetQty

`func (o *InvoiceCreateParamsInvoiceContents) GetQty() float32`

GetQty returns the Qty field if non-nil, zero value otherwise.

### GetQtyOk

`func (o *InvoiceCreateParamsInvoiceContents) GetQtyOk() (*float32, bool)`

GetQtyOk returns a tuple with the Qty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQty

`func (o *InvoiceCreateParamsInvoiceContents) SetQty(v float32)`

SetQty sets Qty field to given value.

### HasQty

`func (o *InvoiceCreateParamsInvoiceContents) HasQty() bool`

HasQty returns a boolean if a field has been set.

### GetUnit

`func (o *InvoiceCreateParamsInvoiceContents) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *InvoiceCreateParamsInvoiceContents) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *InvoiceCreateParamsInvoiceContents) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *InvoiceCreateParamsInvoiceContents) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetUnitPrice

`func (o *InvoiceCreateParamsInvoiceContents) GetUnitPrice() float32`

GetUnitPrice returns the UnitPrice field if non-nil, zero value otherwise.

### GetUnitPriceOk

`func (o *InvoiceCreateParamsInvoiceContents) GetUnitPriceOk() (*float32, bool)`

GetUnitPriceOk returns a tuple with the UnitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPrice

`func (o *InvoiceCreateParamsInvoiceContents) SetUnitPrice(v float32)`

SetUnitPrice sets UnitPrice field to given value.

### HasUnitPrice

`func (o *InvoiceCreateParamsInvoiceContents) HasUnitPrice() bool`

HasUnitPrice returns a boolean if a field has been set.

### GetVat

`func (o *InvoiceCreateParamsInvoiceContents) GetVat() int32`

GetVat returns the Vat field if non-nil, zero value otherwise.

### GetVatOk

`func (o *InvoiceCreateParamsInvoiceContents) GetVatOk() (*int32, bool)`

GetVatOk returns a tuple with the Vat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVat

`func (o *InvoiceCreateParamsInvoiceContents) SetVat(v int32)`

SetVat sets Vat field to given value.

### HasVat

`func (o *InvoiceCreateParamsInvoiceContents) HasVat() bool`

HasVat returns a boolean if a field has been set.

### SetVatNil

`func (o *InvoiceCreateParamsInvoiceContents) SetVatNil(b bool)`

 SetVatNil sets the value for Vat to be an explicit nil

### UnsetVat
`func (o *InvoiceCreateParamsInvoiceContents) UnsetVat()`

UnsetVat ensures that no value is present for Vat, not even an explicit nil
### GetDescription

`func (o *InvoiceCreateParamsInvoiceContents) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InvoiceCreateParamsInvoiceContents) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InvoiceCreateParamsInvoiceContents) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InvoiceCreateParamsInvoiceContents) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAccountItemId

`func (o *InvoiceCreateParamsInvoiceContents) GetAccountItemId() int32`

GetAccountItemId returns the AccountItemId field if non-nil, zero value otherwise.

### GetAccountItemIdOk

`func (o *InvoiceCreateParamsInvoiceContents) GetAccountItemIdOk() (*int32, bool)`

GetAccountItemIdOk returns a tuple with the AccountItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountItemId

`func (o *InvoiceCreateParamsInvoiceContents) SetAccountItemId(v int32)`

SetAccountItemId sets AccountItemId field to given value.

### HasAccountItemId

`func (o *InvoiceCreateParamsInvoiceContents) HasAccountItemId() bool`

HasAccountItemId returns a boolean if a field has been set.

### GetTaxCode

`func (o *InvoiceCreateParamsInvoiceContents) GetTaxCode() int32`

GetTaxCode returns the TaxCode field if non-nil, zero value otherwise.

### GetTaxCodeOk

`func (o *InvoiceCreateParamsInvoiceContents) GetTaxCodeOk() (*int32, bool)`

GetTaxCodeOk returns a tuple with the TaxCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCode

`func (o *InvoiceCreateParamsInvoiceContents) SetTaxCode(v int32)`

SetTaxCode sets TaxCode field to given value.

### HasTaxCode

`func (o *InvoiceCreateParamsInvoiceContents) HasTaxCode() bool`

HasTaxCode returns a boolean if a field has been set.

### GetItemId

`func (o *InvoiceCreateParamsInvoiceContents) GetItemId() int32`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *InvoiceCreateParamsInvoiceContents) GetItemIdOk() (*int32, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *InvoiceCreateParamsInvoiceContents) SetItemId(v int32)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *InvoiceCreateParamsInvoiceContents) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### GetSectionId

`func (o *InvoiceCreateParamsInvoiceContents) GetSectionId() int32`

GetSectionId returns the SectionId field if non-nil, zero value otherwise.

### GetSectionIdOk

`func (o *InvoiceCreateParamsInvoiceContents) GetSectionIdOk() (*int32, bool)`

GetSectionIdOk returns a tuple with the SectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionId

`func (o *InvoiceCreateParamsInvoiceContents) SetSectionId(v int32)`

SetSectionId sets SectionId field to given value.

### HasSectionId

`func (o *InvoiceCreateParamsInvoiceContents) HasSectionId() bool`

HasSectionId returns a boolean if a field has been set.

### GetTagIds

`func (o *InvoiceCreateParamsInvoiceContents) GetTagIds() []int32`

GetTagIds returns the TagIds field if non-nil, zero value otherwise.

### GetTagIdsOk

`func (o *InvoiceCreateParamsInvoiceContents) GetTagIdsOk() (*[]int32, bool)`

GetTagIdsOk returns a tuple with the TagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagIds

`func (o *InvoiceCreateParamsInvoiceContents) SetTagIds(v []int32)`

SetTagIds sets TagIds field to given value.

### HasTagIds

`func (o *InvoiceCreateParamsInvoiceContents) HasTagIds() bool`

HasTagIds returns a boolean if a field has been set.

### GetSegment1TagId

`func (o *InvoiceCreateParamsInvoiceContents) GetSegment1TagId() int32`

GetSegment1TagId returns the Segment1TagId field if non-nil, zero value otherwise.

### GetSegment1TagIdOk

`func (o *InvoiceCreateParamsInvoiceContents) GetSegment1TagIdOk() (*int32, bool)`

GetSegment1TagIdOk returns a tuple with the Segment1TagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment1TagId

`func (o *InvoiceCreateParamsInvoiceContents) SetSegment1TagId(v int32)`

SetSegment1TagId sets Segment1TagId field to given value.

### HasSegment1TagId

`func (o *InvoiceCreateParamsInvoiceContents) HasSegment1TagId() bool`

HasSegment1TagId returns a boolean if a field has been set.

### GetSegment2TagId

`func (o *InvoiceCreateParamsInvoiceContents) GetSegment2TagId() int32`

GetSegment2TagId returns the Segment2TagId field if non-nil, zero value otherwise.

### GetSegment2TagIdOk

`func (o *InvoiceCreateParamsInvoiceContents) GetSegment2TagIdOk() (*int32, bool)`

GetSegment2TagIdOk returns a tuple with the Segment2TagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment2TagId

`func (o *InvoiceCreateParamsInvoiceContents) SetSegment2TagId(v int32)`

SetSegment2TagId sets Segment2TagId field to given value.

### HasSegment2TagId

`func (o *InvoiceCreateParamsInvoiceContents) HasSegment2TagId() bool`

HasSegment2TagId returns a boolean if a field has been set.

### GetSegment3TagId

`func (o *InvoiceCreateParamsInvoiceContents) GetSegment3TagId() int32`

GetSegment3TagId returns the Segment3TagId field if non-nil, zero value otherwise.

### GetSegment3TagIdOk

`func (o *InvoiceCreateParamsInvoiceContents) GetSegment3TagIdOk() (*int32, bool)`

GetSegment3TagIdOk returns a tuple with the Segment3TagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment3TagId

`func (o *InvoiceCreateParamsInvoiceContents) SetSegment3TagId(v int32)`

SetSegment3TagId sets Segment3TagId field to given value.

### HasSegment3TagId

`func (o *InvoiceCreateParamsInvoiceContents) HasSegment3TagId() bool`

HasSegment3TagId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


