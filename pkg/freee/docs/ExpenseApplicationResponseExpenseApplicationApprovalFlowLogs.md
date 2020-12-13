# ExpenseApplicationResponseExpenseApplicationApprovalFlowLogs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** | 操作(apply: 申請, approve: 承認, force_approve: 代理承認, cancel: 取消, reject: 却下, feedback: 差戻し) | 
**UserId** | **int32** | ユーザーID | 
**UpdatedAt** | **string** | 更新日時(ISO8601形式) | 

## Methods

### NewExpenseApplicationResponseExpenseApplicationApprovalFlowLogs

`func NewExpenseApplicationResponseExpenseApplicationApprovalFlowLogs(action string, userId int32, updatedAt string, ) *ExpenseApplicationResponseExpenseApplicationApprovalFlowLogs`

NewExpenseApplicationResponseExpenseApplicationApprovalFlowLogs instantiates a new ExpenseApplicationResponseExpenseApplicationApprovalFlowLogs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExpenseApplicationResponseExpenseApplicationApprovalFlowLogsWithDefaults

`func NewExpenseApplicationResponseExpenseApplicationApprovalFlowLogsWithDefaults() *ExpenseApplicationResponseExpenseApplicationApprovalFlowLogs`

NewExpenseApplicationResponseExpenseApplicationApprovalFlowLogsWithDefaults instantiates a new ExpenseApplicationResponseExpenseApplicationApprovalFlowLogs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *ExpenseApplicationResponseExpenseApplicationApprovalFlowLogs) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ExpenseApplicationResponseExpenseApplicationApprovalFlowLogs) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ExpenseApplicationResponseExpenseApplicationApprovalFlowLogs) SetAction(v string)`

SetAction sets Action field to given value.


### GetUserId

`func (o *ExpenseApplicationResponseExpenseApplicationApprovalFlowLogs) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ExpenseApplicationResponseExpenseApplicationApprovalFlowLogs) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ExpenseApplicationResponseExpenseApplicationApprovalFlowLogs) SetUserId(v int32)`

SetUserId sets UserId field to given value.


### GetUpdatedAt

`func (o *ExpenseApplicationResponseExpenseApplicationApprovalFlowLogs) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ExpenseApplicationResponseExpenseApplicationApprovalFlowLogs) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ExpenseApplicationResponseExpenseApplicationApprovalFlowLogs) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


