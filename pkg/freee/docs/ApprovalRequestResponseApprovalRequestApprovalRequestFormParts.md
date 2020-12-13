# ApprovalRequestResponseApprovalRequestApprovalRequestFormParts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | 項目ID | 
**Order** | Pointer to **int32** | 順序 | [optional] 
**Type** | Pointer to **string** | 項目種別 (title: 申請タイトル, single_line: 自由記述形式 1行, multi_line: 自由記述形式 複数行, select: プルダウン, date: 日付, amount: 金額, receipt: 添付ファイル) | [optional] 
**Label** | Pointer to **string** | 項目名 | [optional] 
**Annotation** | Pointer to **NullableString** | 追加説明 | [optional] 
**Required** | Pointer to **NullableBool** | 必須かどうか | [optional] 
**Values** | Pointer to [**[]ApprovalRequestResponseApprovalRequestApprovalRequestFormValues**](ApprovalRequestResponseApprovalRequestApprovalRequestFormValues.md) | 選択項目 | [optional] 
**MaxAmount** | Pointer to **NullableInt32** | 上限金額 | [optional] 
**MinAmount** | Pointer to **NullableInt32** | 下限金額 | [optional] 

## Methods

### NewApprovalRequestResponseApprovalRequestApprovalRequestFormParts

`func NewApprovalRequestResponseApprovalRequestApprovalRequestFormParts(id int32, ) *ApprovalRequestResponseApprovalRequestApprovalRequestFormParts`

NewApprovalRequestResponseApprovalRequestApprovalRequestFormParts instantiates a new ApprovalRequestResponseApprovalRequestApprovalRequestFormParts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApprovalRequestResponseApprovalRequestApprovalRequestFormPartsWithDefaults

`func NewApprovalRequestResponseApprovalRequestApprovalRequestFormPartsWithDefaults() *ApprovalRequestResponseApprovalRequestApprovalRequestFormParts`

NewApprovalRequestResponseApprovalRequestApprovalRequestFormPartsWithDefaults instantiates a new ApprovalRequestResponseApprovalRequestApprovalRequestFormParts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApprovalRequestResponseApprovalRequestApprovalRequestFormParts) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApprovalRequestResponseApprovalRequestApprovalRequestFormParts) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApprovalRequestResponseApprovalRequestApprovalRequestFormParts) SetId(v int32)`

SetId sets Id field to given value.


### GetOrder

`func (o *ApprovalRequestResponseApprovalRequestApprovalRequestFormParts) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *ApprovalRequestResponseApprovalRequestApprovalRequestFormParts) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *ApprovalRequestResponseApprovalRequestApprovalRequestFormParts) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *ApprovalRequestResponseApprovalRequestApprovalRequestFormParts) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetType

`func (o *ApprovalRequestResponseApprovalRequestApprovalRequestFormParts) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApprovalRequestResponseApprovalRequestApprovalRequestFormParts) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApprovalRequestResponseApprovalRequestApprovalRequestFormParts) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ApprovalRequestResponseApprovalRequestApprovalRequestFormParts) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLabel

`func (o *ApprovalRequestResponseApprovalRequestApprovalRequestFormParts) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ApprovalRequestResponseApprovalRequestApprovalRequestFormParts) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ApprovalRequestResponseApprovalRequestApprovalRequestFormParts) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *ApprovalRequestResponseApprovalRequestApprovalRequestFormParts) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetAnnotation

`func (o *ApprovalRequestResponseApprovalRequestApprovalRequestFormParts) GetAnnotation() string`

GetAnnotation returns the Annotation field if non-nil, zero value otherwise.

### GetAnnotationOk

`func (o *ApprovalRequestResponseApprovalRequestApprovalRequestFormParts) GetAnnotationOk() (*string, bool)`

GetAnnotationOk returns a tuple with the Annotation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotation

`func (o *ApprovalRequestResponseApprovalRequestApprovalRequestFormParts) SetAnnotation(v string)`

SetAnnotation sets Annotation field to given value.

### HasAnnotation

`func (o *ApprovalRequestResponseApprovalRequestApprovalRequestFormParts) HasAnnotation() bool`

HasAnnotation returns a boolean if a field has been set.

### SetAnnotationNil

`func (o *ApprovalRequestResponseApprovalRequestApprovalRequestFormParts) SetAnnotationNil(b bool)`

 SetAnnotationNil sets the value for Annotation to be an explicit nil

### UnsetAnnotation
`func (o *ApprovalRequestResponseApprovalRequestApprovalRequestFormParts) UnsetAnnotation()`

UnsetAnnotation ensures that no value is present for Annotation, not even an explicit nil
### GetRequired

`func (o *ApprovalRequestResponseApprovalRequestApprovalRequestFormParts) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *ApprovalRequestResponseApprovalRequestApprovalRequestFormParts) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *ApprovalRequestResponseApprovalRequestApprovalRequestFormParts) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *ApprovalRequestResponseApprovalRequestApprovalRequestFormParts) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### SetRequiredNil

`func (o *ApprovalRequestResponseApprovalRequestApprovalRequestFormParts) SetRequiredNil(b bool)`

 SetRequiredNil sets the value for Required to be an explicit nil

### UnsetRequired
`func (o *ApprovalRequestResponseApprovalRequestApprovalRequestFormParts) UnsetRequired()`

UnsetRequired ensures that no value is present for Required, not even an explicit nil
### GetValues

`func (o *ApprovalRequestResponseApprovalRequestApprovalRequestFormParts) GetValues() []ApprovalRequestResponseApprovalRequestApprovalRequestFormValues`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *ApprovalRequestResponseApprovalRequestApprovalRequestFormParts) GetValuesOk() (*[]ApprovalRequestResponseApprovalRequestApprovalRequestFormValues, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *ApprovalRequestResponseApprovalRequestApprovalRequestFormParts) SetValues(v []ApprovalRequestResponseApprovalRequestApprovalRequestFormValues)`

SetValues sets Values field to given value.

### HasValues

`func (o *ApprovalRequestResponseApprovalRequestApprovalRequestFormParts) HasValues() bool`

HasValues returns a boolean if a field has been set.

### SetValuesNil

`func (o *ApprovalRequestResponseApprovalRequestApprovalRequestFormParts) SetValuesNil(b bool)`

 SetValuesNil sets the value for Values to be an explicit nil

### UnsetValues
`func (o *ApprovalRequestResponseApprovalRequestApprovalRequestFormParts) UnsetValues()`

UnsetValues ensures that no value is present for Values, not even an explicit nil
### GetMaxAmount

`func (o *ApprovalRequestResponseApprovalRequestApprovalRequestFormParts) GetMaxAmount() int32`

GetMaxAmount returns the MaxAmount field if non-nil, zero value otherwise.

### GetMaxAmountOk

`func (o *ApprovalRequestResponseApprovalRequestApprovalRequestFormParts) GetMaxAmountOk() (*int32, bool)`

GetMaxAmountOk returns a tuple with the MaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAmount

`func (o *ApprovalRequestResponseApprovalRequestApprovalRequestFormParts) SetMaxAmount(v int32)`

SetMaxAmount sets MaxAmount field to given value.

### HasMaxAmount

`func (o *ApprovalRequestResponseApprovalRequestApprovalRequestFormParts) HasMaxAmount() bool`

HasMaxAmount returns a boolean if a field has been set.

### SetMaxAmountNil

`func (o *ApprovalRequestResponseApprovalRequestApprovalRequestFormParts) SetMaxAmountNil(b bool)`

 SetMaxAmountNil sets the value for MaxAmount to be an explicit nil

### UnsetMaxAmount
`func (o *ApprovalRequestResponseApprovalRequestApprovalRequestFormParts) UnsetMaxAmount()`

UnsetMaxAmount ensures that no value is present for MaxAmount, not even an explicit nil
### GetMinAmount

`func (o *ApprovalRequestResponseApprovalRequestApprovalRequestFormParts) GetMinAmount() int32`

GetMinAmount returns the MinAmount field if non-nil, zero value otherwise.

### GetMinAmountOk

`func (o *ApprovalRequestResponseApprovalRequestApprovalRequestFormParts) GetMinAmountOk() (*int32, bool)`

GetMinAmountOk returns a tuple with the MinAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinAmount

`func (o *ApprovalRequestResponseApprovalRequestApprovalRequestFormParts) SetMinAmount(v int32)`

SetMinAmount sets MinAmount field to given value.

### HasMinAmount

`func (o *ApprovalRequestResponseApprovalRequestApprovalRequestFormParts) HasMinAmount() bool`

HasMinAmount returns a boolean if a field has been set.

### SetMinAmountNil

`func (o *ApprovalRequestResponseApprovalRequestApprovalRequestFormParts) SetMinAmountNil(b bool)`

 SetMinAmountNil sets the value for MinAmount to be an explicit nil

### UnsetMinAmount
`func (o *ApprovalRequestResponseApprovalRequestApprovalRequestFormParts) UnsetMinAmount()`

UnsetMinAmount ensures that no value is present for MinAmount, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


