# ApprovalRequestActionCreateParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyId** | **int32** | 事業所ID | 
**ApprovalAction** | **string** | 操作(approve: 承認する、force_approve: 代理承認する、cancel: 申請を取り消す、reject: 却下する、feedback: 申請者へ差し戻す、force_feedback: 承認済み・却下済みを取り消す) | 
**TargetStepId** | **int32** | 対象承認ステップID 各種申請の取得APIレスポンス.current_step_idを送信してください。 | 
**TargetRound** | **int32** | 対象round。差し戻し等により申請がstepの最初からやり直しになるとroundの値が増えます。各種申請の取得APIレスポンス.current_roundを送信してください。 | 
**NextApproverId** | Pointer to **NullableInt32** | 次ステップの承認者のユーザーID | [optional] 

## Methods

### NewApprovalRequestActionCreateParams

`func NewApprovalRequestActionCreateParams(companyId int32, approvalAction string, targetStepId int32, targetRound int32, ) *ApprovalRequestActionCreateParams`

NewApprovalRequestActionCreateParams instantiates a new ApprovalRequestActionCreateParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApprovalRequestActionCreateParamsWithDefaults

`func NewApprovalRequestActionCreateParamsWithDefaults() *ApprovalRequestActionCreateParams`

NewApprovalRequestActionCreateParamsWithDefaults instantiates a new ApprovalRequestActionCreateParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyId

`func (o *ApprovalRequestActionCreateParams) GetCompanyId() int32`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *ApprovalRequestActionCreateParams) GetCompanyIdOk() (*int32, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *ApprovalRequestActionCreateParams) SetCompanyId(v int32)`

SetCompanyId sets CompanyId field to given value.


### GetApprovalAction

`func (o *ApprovalRequestActionCreateParams) GetApprovalAction() string`

GetApprovalAction returns the ApprovalAction field if non-nil, zero value otherwise.

### GetApprovalActionOk

`func (o *ApprovalRequestActionCreateParams) GetApprovalActionOk() (*string, bool)`

GetApprovalActionOk returns a tuple with the ApprovalAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalAction

`func (o *ApprovalRequestActionCreateParams) SetApprovalAction(v string)`

SetApprovalAction sets ApprovalAction field to given value.


### GetTargetStepId

`func (o *ApprovalRequestActionCreateParams) GetTargetStepId() int32`

GetTargetStepId returns the TargetStepId field if non-nil, zero value otherwise.

### GetTargetStepIdOk

`func (o *ApprovalRequestActionCreateParams) GetTargetStepIdOk() (*int32, bool)`

GetTargetStepIdOk returns a tuple with the TargetStepId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetStepId

`func (o *ApprovalRequestActionCreateParams) SetTargetStepId(v int32)`

SetTargetStepId sets TargetStepId field to given value.


### GetTargetRound

`func (o *ApprovalRequestActionCreateParams) GetTargetRound() int32`

GetTargetRound returns the TargetRound field if non-nil, zero value otherwise.

### GetTargetRoundOk

`func (o *ApprovalRequestActionCreateParams) GetTargetRoundOk() (*int32, bool)`

GetTargetRoundOk returns a tuple with the TargetRound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetRound

`func (o *ApprovalRequestActionCreateParams) SetTargetRound(v int32)`

SetTargetRound sets TargetRound field to given value.


### GetNextApproverId

`func (o *ApprovalRequestActionCreateParams) GetNextApproverId() int32`

GetNextApproverId returns the NextApproverId field if non-nil, zero value otherwise.

### GetNextApproverIdOk

`func (o *ApprovalRequestActionCreateParams) GetNextApproverIdOk() (*int32, bool)`

GetNextApproverIdOk returns a tuple with the NextApproverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextApproverId

`func (o *ApprovalRequestActionCreateParams) SetNextApproverId(v int32)`

SetNextApproverId sets NextApproverId field to given value.

### HasNextApproverId

`func (o *ApprovalRequestActionCreateParams) HasNextApproverId() bool`

HasNextApproverId returns a boolean if a field has been set.

### SetNextApproverIdNil

`func (o *ApprovalRequestActionCreateParams) SetNextApproverIdNil(b bool)`

 SetNextApproverIdNil sets the value for NextApproverId to be an explicit nil

### UnsetNextApproverId
`func (o *ApprovalRequestActionCreateParams) UnsetNextApproverId()`

UnsetNextApproverId ensures that no value is present for NextApproverId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


