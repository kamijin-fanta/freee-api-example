# PaymentRequestResponsePaymentRequestPaymentRequestLines

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | 支払依頼の項目行ID | 
**LineType** | **string** | 行の種類 (deal_line: 支払依頼, withholding_tax: 源泉徴収税) | 
**Description** | **string** | 内容 | 
**Amount** | **int32** | 金額 | 
**AccountItemId** | **int32** | 勘定科目ID | 
**TaxCode** | **int32** | 税区分コード | 
**ItemId** | **int32** | 品目ID | 
**SectionId** | **int32** | 部門ID | 
**TagIds** | **[]int32** | メモタグID | 
**Segment1TagId** | Pointer to **int32** | セグメント１ID | [optional] 
**Segment2TagId** | Pointer to **int32** | セグメント２ID | [optional] 
**Segment3TagId** | Pointer to **int32** | セグメント３ID | [optional] 

## Methods

### NewPaymentRequestResponsePaymentRequestPaymentRequestLines

`func NewPaymentRequestResponsePaymentRequestPaymentRequestLines(id int32, lineType string, description string, amount int32, accountItemId int32, taxCode int32, itemId int32, sectionId int32, tagIds []int32, ) *PaymentRequestResponsePaymentRequestPaymentRequestLines`

NewPaymentRequestResponsePaymentRequestPaymentRequestLines instantiates a new PaymentRequestResponsePaymentRequestPaymentRequestLines object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentRequestResponsePaymentRequestPaymentRequestLinesWithDefaults

`func NewPaymentRequestResponsePaymentRequestPaymentRequestLinesWithDefaults() *PaymentRequestResponsePaymentRequestPaymentRequestLines`

NewPaymentRequestResponsePaymentRequestPaymentRequestLinesWithDefaults instantiates a new PaymentRequestResponsePaymentRequestPaymentRequestLines object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PaymentRequestResponsePaymentRequestPaymentRequestLines) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaymentRequestResponsePaymentRequestPaymentRequestLines) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaymentRequestResponsePaymentRequestPaymentRequestLines) SetId(v int32)`

SetId sets Id field to given value.


### GetLineType

`func (o *PaymentRequestResponsePaymentRequestPaymentRequestLines) GetLineType() string`

GetLineType returns the LineType field if non-nil, zero value otherwise.

### GetLineTypeOk

`func (o *PaymentRequestResponsePaymentRequestPaymentRequestLines) GetLineTypeOk() (*string, bool)`

GetLineTypeOk returns a tuple with the LineType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineType

`func (o *PaymentRequestResponsePaymentRequestPaymentRequestLines) SetLineType(v string)`

SetLineType sets LineType field to given value.


### GetDescription

`func (o *PaymentRequestResponsePaymentRequestPaymentRequestLines) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PaymentRequestResponsePaymentRequestPaymentRequestLines) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PaymentRequestResponsePaymentRequestPaymentRequestLines) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetAmount

`func (o *PaymentRequestResponsePaymentRequestPaymentRequestLines) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PaymentRequestResponsePaymentRequestPaymentRequestLines) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PaymentRequestResponsePaymentRequestPaymentRequestLines) SetAmount(v int32)`

SetAmount sets Amount field to given value.


### GetAccountItemId

`func (o *PaymentRequestResponsePaymentRequestPaymentRequestLines) GetAccountItemId() int32`

GetAccountItemId returns the AccountItemId field if non-nil, zero value otherwise.

### GetAccountItemIdOk

`func (o *PaymentRequestResponsePaymentRequestPaymentRequestLines) GetAccountItemIdOk() (*int32, bool)`

GetAccountItemIdOk returns a tuple with the AccountItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountItemId

`func (o *PaymentRequestResponsePaymentRequestPaymentRequestLines) SetAccountItemId(v int32)`

SetAccountItemId sets AccountItemId field to given value.


### GetTaxCode

`func (o *PaymentRequestResponsePaymentRequestPaymentRequestLines) GetTaxCode() int32`

GetTaxCode returns the TaxCode field if non-nil, zero value otherwise.

### GetTaxCodeOk

`func (o *PaymentRequestResponsePaymentRequestPaymentRequestLines) GetTaxCodeOk() (*int32, bool)`

GetTaxCodeOk returns a tuple with the TaxCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCode

`func (o *PaymentRequestResponsePaymentRequestPaymentRequestLines) SetTaxCode(v int32)`

SetTaxCode sets TaxCode field to given value.


### GetItemId

`func (o *PaymentRequestResponsePaymentRequestPaymentRequestLines) GetItemId() int32`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *PaymentRequestResponsePaymentRequestPaymentRequestLines) GetItemIdOk() (*int32, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *PaymentRequestResponsePaymentRequestPaymentRequestLines) SetItemId(v int32)`

SetItemId sets ItemId field to given value.


### GetSectionId

`func (o *PaymentRequestResponsePaymentRequestPaymentRequestLines) GetSectionId() int32`

GetSectionId returns the SectionId field if non-nil, zero value otherwise.

### GetSectionIdOk

`func (o *PaymentRequestResponsePaymentRequestPaymentRequestLines) GetSectionIdOk() (*int32, bool)`

GetSectionIdOk returns a tuple with the SectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionId

`func (o *PaymentRequestResponsePaymentRequestPaymentRequestLines) SetSectionId(v int32)`

SetSectionId sets SectionId field to given value.


### GetTagIds

`func (o *PaymentRequestResponsePaymentRequestPaymentRequestLines) GetTagIds() []int32`

GetTagIds returns the TagIds field if non-nil, zero value otherwise.

### GetTagIdsOk

`func (o *PaymentRequestResponsePaymentRequestPaymentRequestLines) GetTagIdsOk() (*[]int32, bool)`

GetTagIdsOk returns a tuple with the TagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagIds

`func (o *PaymentRequestResponsePaymentRequestPaymentRequestLines) SetTagIds(v []int32)`

SetTagIds sets TagIds field to given value.


### GetSegment1TagId

`func (o *PaymentRequestResponsePaymentRequestPaymentRequestLines) GetSegment1TagId() int32`

GetSegment1TagId returns the Segment1TagId field if non-nil, zero value otherwise.

### GetSegment1TagIdOk

`func (o *PaymentRequestResponsePaymentRequestPaymentRequestLines) GetSegment1TagIdOk() (*int32, bool)`

GetSegment1TagIdOk returns a tuple with the Segment1TagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment1TagId

`func (o *PaymentRequestResponsePaymentRequestPaymentRequestLines) SetSegment1TagId(v int32)`

SetSegment1TagId sets Segment1TagId field to given value.

### HasSegment1TagId

`func (o *PaymentRequestResponsePaymentRequestPaymentRequestLines) HasSegment1TagId() bool`

HasSegment1TagId returns a boolean if a field has been set.

### GetSegment2TagId

`func (o *PaymentRequestResponsePaymentRequestPaymentRequestLines) GetSegment2TagId() int32`

GetSegment2TagId returns the Segment2TagId field if non-nil, zero value otherwise.

### GetSegment2TagIdOk

`func (o *PaymentRequestResponsePaymentRequestPaymentRequestLines) GetSegment2TagIdOk() (*int32, bool)`

GetSegment2TagIdOk returns a tuple with the Segment2TagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment2TagId

`func (o *PaymentRequestResponsePaymentRequestPaymentRequestLines) SetSegment2TagId(v int32)`

SetSegment2TagId sets Segment2TagId field to given value.

### HasSegment2TagId

`func (o *PaymentRequestResponsePaymentRequestPaymentRequestLines) HasSegment2TagId() bool`

HasSegment2TagId returns a boolean if a field has been set.

### GetSegment3TagId

`func (o *PaymentRequestResponsePaymentRequestPaymentRequestLines) GetSegment3TagId() int32`

GetSegment3TagId returns the Segment3TagId field if non-nil, zero value otherwise.

### GetSegment3TagIdOk

`func (o *PaymentRequestResponsePaymentRequestPaymentRequestLines) GetSegment3TagIdOk() (*int32, bool)`

GetSegment3TagIdOk returns a tuple with the Segment3TagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment3TagId

`func (o *PaymentRequestResponsePaymentRequestPaymentRequestLines) SetSegment3TagId(v int32)`

SetSegment3TagId sets Segment3TagId field to given value.

### HasSegment3TagId

`func (o *PaymentRequestResponsePaymentRequestPaymentRequestLines) HasSegment3TagId() bool`

HasSegment3TagId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


