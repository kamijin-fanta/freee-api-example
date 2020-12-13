# QuotationUpdateParamsQuotationContents

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | 見積内容ID | [optional] 
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

### NewQuotationUpdateParamsQuotationContents

`func NewQuotationUpdateParamsQuotationContents(order int32, type_ string, ) *QuotationUpdateParamsQuotationContents`

NewQuotationUpdateParamsQuotationContents instantiates a new QuotationUpdateParamsQuotationContents object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuotationUpdateParamsQuotationContentsWithDefaults

`func NewQuotationUpdateParamsQuotationContentsWithDefaults() *QuotationUpdateParamsQuotationContents`

NewQuotationUpdateParamsQuotationContentsWithDefaults instantiates a new QuotationUpdateParamsQuotationContents object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QuotationUpdateParamsQuotationContents) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QuotationUpdateParamsQuotationContents) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QuotationUpdateParamsQuotationContents) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *QuotationUpdateParamsQuotationContents) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrder

`func (o *QuotationUpdateParamsQuotationContents) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *QuotationUpdateParamsQuotationContents) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *QuotationUpdateParamsQuotationContents) SetOrder(v int32)`

SetOrder sets Order field to given value.


### GetType

`func (o *QuotationUpdateParamsQuotationContents) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *QuotationUpdateParamsQuotationContents) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *QuotationUpdateParamsQuotationContents) SetType(v string)`

SetType sets Type field to given value.


### GetQty

`func (o *QuotationUpdateParamsQuotationContents) GetQty() float32`

GetQty returns the Qty field if non-nil, zero value otherwise.

### GetQtyOk

`func (o *QuotationUpdateParamsQuotationContents) GetQtyOk() (*float32, bool)`

GetQtyOk returns a tuple with the Qty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQty

`func (o *QuotationUpdateParamsQuotationContents) SetQty(v float32)`

SetQty sets Qty field to given value.

### HasQty

`func (o *QuotationUpdateParamsQuotationContents) HasQty() bool`

HasQty returns a boolean if a field has been set.

### GetUnit

`func (o *QuotationUpdateParamsQuotationContents) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *QuotationUpdateParamsQuotationContents) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *QuotationUpdateParamsQuotationContents) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *QuotationUpdateParamsQuotationContents) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetUnitPrice

`func (o *QuotationUpdateParamsQuotationContents) GetUnitPrice() float32`

GetUnitPrice returns the UnitPrice field if non-nil, zero value otherwise.

### GetUnitPriceOk

`func (o *QuotationUpdateParamsQuotationContents) GetUnitPriceOk() (*float32, bool)`

GetUnitPriceOk returns a tuple with the UnitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPrice

`func (o *QuotationUpdateParamsQuotationContents) SetUnitPrice(v float32)`

SetUnitPrice sets UnitPrice field to given value.

### HasUnitPrice

`func (o *QuotationUpdateParamsQuotationContents) HasUnitPrice() bool`

HasUnitPrice returns a boolean if a field has been set.

### GetVat

`func (o *QuotationUpdateParamsQuotationContents) GetVat() int32`

GetVat returns the Vat field if non-nil, zero value otherwise.

### GetVatOk

`func (o *QuotationUpdateParamsQuotationContents) GetVatOk() (*int32, bool)`

GetVatOk returns a tuple with the Vat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVat

`func (o *QuotationUpdateParamsQuotationContents) SetVat(v int32)`

SetVat sets Vat field to given value.

### HasVat

`func (o *QuotationUpdateParamsQuotationContents) HasVat() bool`

HasVat returns a boolean if a field has been set.

### SetVatNil

`func (o *QuotationUpdateParamsQuotationContents) SetVatNil(b bool)`

 SetVatNil sets the value for Vat to be an explicit nil

### UnsetVat
`func (o *QuotationUpdateParamsQuotationContents) UnsetVat()`

UnsetVat ensures that no value is present for Vat, not even an explicit nil
### GetDescription

`func (o *QuotationUpdateParamsQuotationContents) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *QuotationUpdateParamsQuotationContents) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *QuotationUpdateParamsQuotationContents) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *QuotationUpdateParamsQuotationContents) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAccountItemId

`func (o *QuotationUpdateParamsQuotationContents) GetAccountItemId() int32`

GetAccountItemId returns the AccountItemId field if non-nil, zero value otherwise.

### GetAccountItemIdOk

`func (o *QuotationUpdateParamsQuotationContents) GetAccountItemIdOk() (*int32, bool)`

GetAccountItemIdOk returns a tuple with the AccountItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountItemId

`func (o *QuotationUpdateParamsQuotationContents) SetAccountItemId(v int32)`

SetAccountItemId sets AccountItemId field to given value.

### HasAccountItemId

`func (o *QuotationUpdateParamsQuotationContents) HasAccountItemId() bool`

HasAccountItemId returns a boolean if a field has been set.

### GetTaxCode

`func (o *QuotationUpdateParamsQuotationContents) GetTaxCode() int32`

GetTaxCode returns the TaxCode field if non-nil, zero value otherwise.

### GetTaxCodeOk

`func (o *QuotationUpdateParamsQuotationContents) GetTaxCodeOk() (*int32, bool)`

GetTaxCodeOk returns a tuple with the TaxCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCode

`func (o *QuotationUpdateParamsQuotationContents) SetTaxCode(v int32)`

SetTaxCode sets TaxCode field to given value.

### HasTaxCode

`func (o *QuotationUpdateParamsQuotationContents) HasTaxCode() bool`

HasTaxCode returns a boolean if a field has been set.

### GetItemId

`func (o *QuotationUpdateParamsQuotationContents) GetItemId() int32`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *QuotationUpdateParamsQuotationContents) GetItemIdOk() (*int32, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *QuotationUpdateParamsQuotationContents) SetItemId(v int32)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *QuotationUpdateParamsQuotationContents) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### GetSectionId

`func (o *QuotationUpdateParamsQuotationContents) GetSectionId() int32`

GetSectionId returns the SectionId field if non-nil, zero value otherwise.

### GetSectionIdOk

`func (o *QuotationUpdateParamsQuotationContents) GetSectionIdOk() (*int32, bool)`

GetSectionIdOk returns a tuple with the SectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionId

`func (o *QuotationUpdateParamsQuotationContents) SetSectionId(v int32)`

SetSectionId sets SectionId field to given value.

### HasSectionId

`func (o *QuotationUpdateParamsQuotationContents) HasSectionId() bool`

HasSectionId returns a boolean if a field has been set.

### GetTagIds

`func (o *QuotationUpdateParamsQuotationContents) GetTagIds() []int32`

GetTagIds returns the TagIds field if non-nil, zero value otherwise.

### GetTagIdsOk

`func (o *QuotationUpdateParamsQuotationContents) GetTagIdsOk() (*[]int32, bool)`

GetTagIdsOk returns a tuple with the TagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagIds

`func (o *QuotationUpdateParamsQuotationContents) SetTagIds(v []int32)`

SetTagIds sets TagIds field to given value.

### HasTagIds

`func (o *QuotationUpdateParamsQuotationContents) HasTagIds() bool`

HasTagIds returns a boolean if a field has been set.

### GetSegment1TagId

`func (o *QuotationUpdateParamsQuotationContents) GetSegment1TagId() int32`

GetSegment1TagId returns the Segment1TagId field if non-nil, zero value otherwise.

### GetSegment1TagIdOk

`func (o *QuotationUpdateParamsQuotationContents) GetSegment1TagIdOk() (*int32, bool)`

GetSegment1TagIdOk returns a tuple with the Segment1TagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment1TagId

`func (o *QuotationUpdateParamsQuotationContents) SetSegment1TagId(v int32)`

SetSegment1TagId sets Segment1TagId field to given value.

### HasSegment1TagId

`func (o *QuotationUpdateParamsQuotationContents) HasSegment1TagId() bool`

HasSegment1TagId returns a boolean if a field has been set.

### GetSegment2TagId

`func (o *QuotationUpdateParamsQuotationContents) GetSegment2TagId() int32`

GetSegment2TagId returns the Segment2TagId field if non-nil, zero value otherwise.

### GetSegment2TagIdOk

`func (o *QuotationUpdateParamsQuotationContents) GetSegment2TagIdOk() (*int32, bool)`

GetSegment2TagIdOk returns a tuple with the Segment2TagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment2TagId

`func (o *QuotationUpdateParamsQuotationContents) SetSegment2TagId(v int32)`

SetSegment2TagId sets Segment2TagId field to given value.

### HasSegment2TagId

`func (o *QuotationUpdateParamsQuotationContents) HasSegment2TagId() bool`

HasSegment2TagId returns a boolean if a field has been set.

### GetSegment3TagId

`func (o *QuotationUpdateParamsQuotationContents) GetSegment3TagId() int32`

GetSegment3TagId returns the Segment3TagId field if non-nil, zero value otherwise.

### GetSegment3TagIdOk

`func (o *QuotationUpdateParamsQuotationContents) GetSegment3TagIdOk() (*int32, bool)`

GetSegment3TagIdOk returns a tuple with the Segment3TagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment3TagId

`func (o *QuotationUpdateParamsQuotationContents) SetSegment3TagId(v int32)`

SetSegment3TagId sets Segment3TagId field to given value.

### HasSegment3TagId

`func (o *QuotationUpdateParamsQuotationContents) HasSegment3TagId() bool`

HasSegment3TagId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


