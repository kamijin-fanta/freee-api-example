# PaymentRequestCreateParamsPaymentRequestLines

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LineType** | Pointer to **string** | 行の種類 (deal_line: 支払依頼, withholding_tax: 源泉徴収税) | [optional] 
**Description** | Pointer to **string** | 内容 | [optional] 
**Amount** | **int32** | 金額 | 
**AccountItemId** | Pointer to **int32** | 勘定科目ID | [optional] 
**TaxCode** | **int32** | 税区分コード | 
**ItemId** | Pointer to **int32** | 品目ID | [optional] 
**SectionId** | Pointer to **int32** | 部門ID | [optional] 
**TagIds** | Pointer to **[]int32** | メモタグID | [optional] 
**Segment1TagId** | Pointer to **int32** | セグメント１ID&lt;br&gt; セグメントタグ一覧APIを利用して取得してください。&lt;br&gt; &lt;a href&#x3D;\&quot;https://support.freee.co.jp/hc/ja/articles/360020679611\&quot; target&#x3D;\&quot;_blank\&quot;&gt;セグメント（分析用タグ）の設定&lt;/a&gt;&lt;br&gt;  | [optional] 
**Segment2TagId** | Pointer to **int32** | セグメント２ID(法人向けエンタープライズプラン)&lt;br&gt; セグメントタグ一覧APIを利用して取得してください。&lt;br&gt; &lt;a href&#x3D;\&quot;https://support.freee.co.jp/hc/ja/articles/360020679611\&quot; target&#x3D;\&quot;_blank\&quot;&gt;セグメント（分析用タグ）の設定&lt;/a&gt;&lt;br&gt;  | [optional] 
**Segment3TagId** | Pointer to **int32** | セグメント３ID(法人向けエンタープライズプラン)&lt;br&gt; セグメントタグ一覧APIを利用して取得してください。&lt;br&gt; &lt;a href&#x3D;\&quot;https://support.freee.co.jp/hc/ja/articles/360020679611\&quot; target&#x3D;\&quot;_blank\&quot;&gt;セグメント（分析用タグ）の設定&lt;/a&gt;&lt;br&gt;  | [optional] 

## Methods

### NewPaymentRequestCreateParamsPaymentRequestLines

`func NewPaymentRequestCreateParamsPaymentRequestLines(amount int32, taxCode int32, ) *PaymentRequestCreateParamsPaymentRequestLines`

NewPaymentRequestCreateParamsPaymentRequestLines instantiates a new PaymentRequestCreateParamsPaymentRequestLines object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentRequestCreateParamsPaymentRequestLinesWithDefaults

`func NewPaymentRequestCreateParamsPaymentRequestLinesWithDefaults() *PaymentRequestCreateParamsPaymentRequestLines`

NewPaymentRequestCreateParamsPaymentRequestLinesWithDefaults instantiates a new PaymentRequestCreateParamsPaymentRequestLines object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLineType

`func (o *PaymentRequestCreateParamsPaymentRequestLines) GetLineType() string`

GetLineType returns the LineType field if non-nil, zero value otherwise.

### GetLineTypeOk

`func (o *PaymentRequestCreateParamsPaymentRequestLines) GetLineTypeOk() (*string, bool)`

GetLineTypeOk returns a tuple with the LineType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineType

`func (o *PaymentRequestCreateParamsPaymentRequestLines) SetLineType(v string)`

SetLineType sets LineType field to given value.

### HasLineType

`func (o *PaymentRequestCreateParamsPaymentRequestLines) HasLineType() bool`

HasLineType returns a boolean if a field has been set.

### GetDescription

`func (o *PaymentRequestCreateParamsPaymentRequestLines) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PaymentRequestCreateParamsPaymentRequestLines) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PaymentRequestCreateParamsPaymentRequestLines) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PaymentRequestCreateParamsPaymentRequestLines) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAmount

`func (o *PaymentRequestCreateParamsPaymentRequestLines) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PaymentRequestCreateParamsPaymentRequestLines) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PaymentRequestCreateParamsPaymentRequestLines) SetAmount(v int32)`

SetAmount sets Amount field to given value.


### GetAccountItemId

`func (o *PaymentRequestCreateParamsPaymentRequestLines) GetAccountItemId() int32`

GetAccountItemId returns the AccountItemId field if non-nil, zero value otherwise.

### GetAccountItemIdOk

`func (o *PaymentRequestCreateParamsPaymentRequestLines) GetAccountItemIdOk() (*int32, bool)`

GetAccountItemIdOk returns a tuple with the AccountItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountItemId

`func (o *PaymentRequestCreateParamsPaymentRequestLines) SetAccountItemId(v int32)`

SetAccountItemId sets AccountItemId field to given value.

### HasAccountItemId

`func (o *PaymentRequestCreateParamsPaymentRequestLines) HasAccountItemId() bool`

HasAccountItemId returns a boolean if a field has been set.

### GetTaxCode

`func (o *PaymentRequestCreateParamsPaymentRequestLines) GetTaxCode() int32`

GetTaxCode returns the TaxCode field if non-nil, zero value otherwise.

### GetTaxCodeOk

`func (o *PaymentRequestCreateParamsPaymentRequestLines) GetTaxCodeOk() (*int32, bool)`

GetTaxCodeOk returns a tuple with the TaxCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCode

`func (o *PaymentRequestCreateParamsPaymentRequestLines) SetTaxCode(v int32)`

SetTaxCode sets TaxCode field to given value.


### GetItemId

`func (o *PaymentRequestCreateParamsPaymentRequestLines) GetItemId() int32`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *PaymentRequestCreateParamsPaymentRequestLines) GetItemIdOk() (*int32, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *PaymentRequestCreateParamsPaymentRequestLines) SetItemId(v int32)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *PaymentRequestCreateParamsPaymentRequestLines) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### GetSectionId

`func (o *PaymentRequestCreateParamsPaymentRequestLines) GetSectionId() int32`

GetSectionId returns the SectionId field if non-nil, zero value otherwise.

### GetSectionIdOk

`func (o *PaymentRequestCreateParamsPaymentRequestLines) GetSectionIdOk() (*int32, bool)`

GetSectionIdOk returns a tuple with the SectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionId

`func (o *PaymentRequestCreateParamsPaymentRequestLines) SetSectionId(v int32)`

SetSectionId sets SectionId field to given value.

### HasSectionId

`func (o *PaymentRequestCreateParamsPaymentRequestLines) HasSectionId() bool`

HasSectionId returns a boolean if a field has been set.

### GetTagIds

`func (o *PaymentRequestCreateParamsPaymentRequestLines) GetTagIds() []int32`

GetTagIds returns the TagIds field if non-nil, zero value otherwise.

### GetTagIdsOk

`func (o *PaymentRequestCreateParamsPaymentRequestLines) GetTagIdsOk() (*[]int32, bool)`

GetTagIdsOk returns a tuple with the TagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagIds

`func (o *PaymentRequestCreateParamsPaymentRequestLines) SetTagIds(v []int32)`

SetTagIds sets TagIds field to given value.

### HasTagIds

`func (o *PaymentRequestCreateParamsPaymentRequestLines) HasTagIds() bool`

HasTagIds returns a boolean if a field has been set.

### GetSegment1TagId

`func (o *PaymentRequestCreateParamsPaymentRequestLines) GetSegment1TagId() int32`

GetSegment1TagId returns the Segment1TagId field if non-nil, zero value otherwise.

### GetSegment1TagIdOk

`func (o *PaymentRequestCreateParamsPaymentRequestLines) GetSegment1TagIdOk() (*int32, bool)`

GetSegment1TagIdOk returns a tuple with the Segment1TagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment1TagId

`func (o *PaymentRequestCreateParamsPaymentRequestLines) SetSegment1TagId(v int32)`

SetSegment1TagId sets Segment1TagId field to given value.

### HasSegment1TagId

`func (o *PaymentRequestCreateParamsPaymentRequestLines) HasSegment1TagId() bool`

HasSegment1TagId returns a boolean if a field has been set.

### GetSegment2TagId

`func (o *PaymentRequestCreateParamsPaymentRequestLines) GetSegment2TagId() int32`

GetSegment2TagId returns the Segment2TagId field if non-nil, zero value otherwise.

### GetSegment2TagIdOk

`func (o *PaymentRequestCreateParamsPaymentRequestLines) GetSegment2TagIdOk() (*int32, bool)`

GetSegment2TagIdOk returns a tuple with the Segment2TagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment2TagId

`func (o *PaymentRequestCreateParamsPaymentRequestLines) SetSegment2TagId(v int32)`

SetSegment2TagId sets Segment2TagId field to given value.

### HasSegment2TagId

`func (o *PaymentRequestCreateParamsPaymentRequestLines) HasSegment2TagId() bool`

HasSegment2TagId returns a boolean if a field has been set.

### GetSegment3TagId

`func (o *PaymentRequestCreateParamsPaymentRequestLines) GetSegment3TagId() int32`

GetSegment3TagId returns the Segment3TagId field if non-nil, zero value otherwise.

### GetSegment3TagIdOk

`func (o *PaymentRequestCreateParamsPaymentRequestLines) GetSegment3TagIdOk() (*int32, bool)`

GetSegment3TagIdOk returns a tuple with the Segment3TagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment3TagId

`func (o *PaymentRequestCreateParamsPaymentRequestLines) SetSegment3TagId(v int32)`

SetSegment3TagId sets Segment3TagId field to given value.

### HasSegment3TagId

`func (o *PaymentRequestCreateParamsPaymentRequestLines) HasSegment3TagId() bool`

HasSegment3TagId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


