# ApprovalRequestCreateParamsRequestItems

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | 項目ID | [optional] 
**Type** | Pointer to **string** | 項目タイプ(title: 申請タイトル, single_line: 自由記述形式 1行, multi_line: 自由記述形式 複数行, select: プルダウン, date: 日付, amount: 金額, receipt: 添付ファイル) | [optional] 
**Value** | Pointer to **string** | 項目の値 | [optional] 

## Methods

### NewApprovalRequestCreateParamsRequestItems

`func NewApprovalRequestCreateParamsRequestItems() *ApprovalRequestCreateParamsRequestItems`

NewApprovalRequestCreateParamsRequestItems instantiates a new ApprovalRequestCreateParamsRequestItems object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApprovalRequestCreateParamsRequestItemsWithDefaults

`func NewApprovalRequestCreateParamsRequestItemsWithDefaults() *ApprovalRequestCreateParamsRequestItems`

NewApprovalRequestCreateParamsRequestItemsWithDefaults instantiates a new ApprovalRequestCreateParamsRequestItems object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApprovalRequestCreateParamsRequestItems) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApprovalRequestCreateParamsRequestItems) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApprovalRequestCreateParamsRequestItems) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ApprovalRequestCreateParamsRequestItems) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *ApprovalRequestCreateParamsRequestItems) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApprovalRequestCreateParamsRequestItems) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApprovalRequestCreateParamsRequestItems) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ApprovalRequestCreateParamsRequestItems) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValue

`func (o *ApprovalRequestCreateParamsRequestItems) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ApprovalRequestCreateParamsRequestItems) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ApprovalRequestCreateParamsRequestItems) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ApprovalRequestCreateParamsRequestItems) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


