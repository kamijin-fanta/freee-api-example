# InvoiceUpdateParamsInvoiceContents

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | 請求内容ID | [optional] 
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

### NewInvoiceUpdateParamsInvoiceContents

`func NewInvoiceUpdateParamsInvoiceContents(order int32, type_ string, ) *InvoiceUpdateParamsInvoiceContents`

NewInvoiceUpdateParamsInvoiceContents instantiates a new InvoiceUpdateParamsInvoiceContents object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceUpdateParamsInvoiceContentsWithDefaults

`func NewInvoiceUpdateParamsInvoiceContentsWithDefaults() *InvoiceUpdateParamsInvoiceContents`

NewInvoiceUpdateParamsInvoiceContentsWithDefaults instantiates a new InvoiceUpdateParamsInvoiceContents object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InvoiceUpdateParamsInvoiceContents) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InvoiceUpdateParamsInvoiceContents) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InvoiceUpdateParamsInvoiceContents) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *InvoiceUpdateParamsInvoiceContents) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrder

`func (o *InvoiceUpdateParamsInvoiceContents) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *InvoiceUpdateParamsInvoiceContents) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *InvoiceUpdateParamsInvoiceContents) SetOrder(v int32)`

SetOrder sets Order field to given value.


### GetType

`func (o *InvoiceUpdateParamsInvoiceContents) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InvoiceUpdateParamsInvoiceContents) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InvoiceUpdateParamsInvoiceContents) SetType(v string)`

SetType sets Type field to given value.


### GetQty

`func (o *InvoiceUpdateParamsInvoiceContents) GetQty() float32`

GetQty returns the Qty field if non-nil, zero value otherwise.

### GetQtyOk

`func (o *InvoiceUpdateParamsInvoiceContents) GetQtyOk() (*float32, bool)`

GetQtyOk returns a tuple with the Qty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQty

`func (o *InvoiceUpdateParamsInvoiceContents) SetQty(v float32)`

SetQty sets Qty field to given value.

### HasQty

`func (o *InvoiceUpdateParamsInvoiceContents) HasQty() bool`

HasQty returns a boolean if a field has been set.

### GetUnit

`func (o *InvoiceUpdateParamsInvoiceContents) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *InvoiceUpdateParamsInvoiceContents) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *InvoiceUpdateParamsInvoiceContents) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *InvoiceUpdateParamsInvoiceContents) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetUnitPrice

`func (o *InvoiceUpdateParamsInvoiceContents) GetUnitPrice() float32`

GetUnitPrice returns the UnitPrice field if non-nil, zero value otherwise.

### GetUnitPriceOk

`func (o *InvoiceUpdateParamsInvoiceContents) GetUnitPriceOk() (*float32, bool)`

GetUnitPriceOk returns a tuple with the UnitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPrice

`func (o *InvoiceUpdateParamsInvoiceContents) SetUnitPrice(v float32)`

SetUnitPrice sets UnitPrice field to given value.

### HasUnitPrice

`func (o *InvoiceUpdateParamsInvoiceContents) HasUnitPrice() bool`

HasUnitPrice returns a boolean if a field has been set.

### GetVat

`func (o *InvoiceUpdateParamsInvoiceContents) GetVat() int32`

GetVat returns the Vat field if non-nil, zero value otherwise.

### GetVatOk

`func (o *InvoiceUpdateParamsInvoiceContents) GetVatOk() (*int32, bool)`

GetVatOk returns a tuple with the Vat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVat

`func (o *InvoiceUpdateParamsInvoiceContents) SetVat(v int32)`

SetVat sets Vat field to given value.

### HasVat

`func (o *InvoiceUpdateParamsInvoiceContents) HasVat() bool`

HasVat returns a boolean if a field has been set.

### SetVatNil

`func (o *InvoiceUpdateParamsInvoiceContents) SetVatNil(b bool)`

 SetVatNil sets the value for Vat to be an explicit nil

### UnsetVat
`func (o *InvoiceUpdateParamsInvoiceContents) UnsetVat()`

UnsetVat ensures that no value is present for Vat, not even an explicit nil
### GetDescription

`func (o *InvoiceUpdateParamsInvoiceContents) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InvoiceUpdateParamsInvoiceContents) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InvoiceUpdateParamsInvoiceContents) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InvoiceUpdateParamsInvoiceContents) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAccountItemId

`func (o *InvoiceUpdateParamsInvoiceContents) GetAccountItemId() int32`

GetAccountItemId returns the AccountItemId field if non-nil, zero value otherwise.

### GetAccountItemIdOk

`func (o *InvoiceUpdateParamsInvoiceContents) GetAccountItemIdOk() (*int32, bool)`

GetAccountItemIdOk returns a tuple with the AccountItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountItemId

`func (o *InvoiceUpdateParamsInvoiceContents) SetAccountItemId(v int32)`

SetAccountItemId sets AccountItemId field to given value.

### HasAccountItemId

`func (o *InvoiceUpdateParamsInvoiceContents) HasAccountItemId() bool`

HasAccountItemId returns a boolean if a field has been set.

### GetTaxCode

`func (o *InvoiceUpdateParamsInvoiceContents) GetTaxCode() int32`

GetTaxCode returns the TaxCode field if non-nil, zero value otherwise.

### GetTaxCodeOk

`func (o *InvoiceUpdateParamsInvoiceContents) GetTaxCodeOk() (*int32, bool)`

GetTaxCodeOk returns a tuple with the TaxCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCode

`func (o *InvoiceUpdateParamsInvoiceContents) SetTaxCode(v int32)`

SetTaxCode sets TaxCode field to given value.

### HasTaxCode

`func (o *InvoiceUpdateParamsInvoiceContents) HasTaxCode() bool`

HasTaxCode returns a boolean if a field has been set.

### GetItemId

`func (o *InvoiceUpdateParamsInvoiceContents) GetItemId() int32`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *InvoiceUpdateParamsInvoiceContents) GetItemIdOk() (*int32, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *InvoiceUpdateParamsInvoiceContents) SetItemId(v int32)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *InvoiceUpdateParamsInvoiceContents) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### GetSectionId

`func (o *InvoiceUpdateParamsInvoiceContents) GetSectionId() int32`

GetSectionId returns the SectionId field if non-nil, zero value otherwise.

### GetSectionIdOk

`func (o *InvoiceUpdateParamsInvoiceContents) GetSectionIdOk() (*int32, bool)`

GetSectionIdOk returns a tuple with the SectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionId

`func (o *InvoiceUpdateParamsInvoiceContents) SetSectionId(v int32)`

SetSectionId sets SectionId field to given value.

### HasSectionId

`func (o *InvoiceUpdateParamsInvoiceContents) HasSectionId() bool`

HasSectionId returns a boolean if a field has been set.

### GetTagIds

`func (o *InvoiceUpdateParamsInvoiceContents) GetTagIds() []int32`

GetTagIds returns the TagIds field if non-nil, zero value otherwise.

### GetTagIdsOk

`func (o *InvoiceUpdateParamsInvoiceContents) GetTagIdsOk() (*[]int32, bool)`

GetTagIdsOk returns a tuple with the TagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagIds

`func (o *InvoiceUpdateParamsInvoiceContents) SetTagIds(v []int32)`

SetTagIds sets TagIds field to given value.

### HasTagIds

`func (o *InvoiceUpdateParamsInvoiceContents) HasTagIds() bool`

HasTagIds returns a boolean if a field has been set.

### GetSegment1TagId

`func (o *InvoiceUpdateParamsInvoiceContents) GetSegment1TagId() int32`

GetSegment1TagId returns the Segment1TagId field if non-nil, zero value otherwise.

### GetSegment1TagIdOk

`func (o *InvoiceUpdateParamsInvoiceContents) GetSegment1TagIdOk() (*int32, bool)`

GetSegment1TagIdOk returns a tuple with the Segment1TagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment1TagId

`func (o *InvoiceUpdateParamsInvoiceContents) SetSegment1TagId(v int32)`

SetSegment1TagId sets Segment1TagId field to given value.

### HasSegment1TagId

`func (o *InvoiceUpdateParamsInvoiceContents) HasSegment1TagId() bool`

HasSegment1TagId returns a boolean if a field has been set.

### GetSegment2TagId

`func (o *InvoiceUpdateParamsInvoiceContents) GetSegment2TagId() int32`

GetSegment2TagId returns the Segment2TagId field if non-nil, zero value otherwise.

### GetSegment2TagIdOk

`func (o *InvoiceUpdateParamsInvoiceContents) GetSegment2TagIdOk() (*int32, bool)`

GetSegment2TagIdOk returns a tuple with the Segment2TagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment2TagId

`func (o *InvoiceUpdateParamsInvoiceContents) SetSegment2TagId(v int32)`

SetSegment2TagId sets Segment2TagId field to given value.

### HasSegment2TagId

`func (o *InvoiceUpdateParamsInvoiceContents) HasSegment2TagId() bool`

HasSegment2TagId returns a boolean if a field has been set.

### GetSegment3TagId

`func (o *InvoiceUpdateParamsInvoiceContents) GetSegment3TagId() int32`

GetSegment3TagId returns the Segment3TagId field if non-nil, zero value otherwise.

### GetSegment3TagIdOk

`func (o *InvoiceUpdateParamsInvoiceContents) GetSegment3TagIdOk() (*int32, bool)`

GetSegment3TagIdOk returns a tuple with the Segment3TagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment3TagId

`func (o *InvoiceUpdateParamsInvoiceContents) SetSegment3TagId(v int32)`

SetSegment3TagId sets Segment3TagId field to given value.

### HasSegment3TagId

`func (o *InvoiceUpdateParamsInvoiceContents) HasSegment3TagId() bool`

HasSegment3TagId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


