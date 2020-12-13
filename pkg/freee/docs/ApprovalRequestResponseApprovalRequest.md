# ApprovalRequestResponseApprovalRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | 各種申請ID | 
**CompanyId** | **int32** | 事業所ID | 
**ApplicationDate** | **string** | 申請日 (yyyy-mm-dd) | 
**Title** | **string** | 申請タイトル | 
**ApplicantId** | **int32** | 申請者のユーザーID | 
**Approvers** | [**[]ExpenseApplicationResponseExpenseApplicationApprovers**](ExpenseApplicationResponseExpenseApplicationApprovers.md) | 承認者（配列）   承認ステップのresource_typeがunspecified (指定なし)の場合はapproversはレスポンスに含まれません。   しかし、resource_typeがunspecifiedの承認ステップにおいて誰かが承認・却下・差し戻しのいずれかのアクションを取った後は、    approversはレスポンスに含まれるようになります。    その場合approversにはアクションを行ったステップのIDとアクションを行ったユーザーのIDが含まれます。 | 
**ApplicationNumber** | **string** | 申請No. | 
**Status** | **string** | 申請ステータス(draft:下書き, in_progress:申請中, approved:承認済, rejected:却下, feedback:差戻し) | 
**RequestItems** | [**[]ApprovalRequestsIndexResponseRequestItems**](ApprovalRequestsIndexResponseRequestItems.md) | 各種申請の項目一覧（配列） | 
**FormId** | **int32** | 申請フォームID | 
**ApprovalFlowRouteId** | **int32** | 申請経路ID | 
**Comments** | [**[]ExpenseApplicationResponseExpenseApplicationComments**](ExpenseApplicationResponseExpenseApplicationComments.md) | 各種申請のコメント一覧（配列） | 
**ApprovalFlowLogs** | [**[]ExpenseApplicationResponseExpenseApplicationApprovalFlowLogs**](ExpenseApplicationResponseExpenseApplicationApprovalFlowLogs.md) | 各種申請の承認履歴（配列） | 
**CurrentStepId** | **NullableInt32** | 現在承認ステップID | 
**CurrentRound** | **int32** | 現在のround。差し戻し等により申請がstepの最初からやり直しになるとroundの値が増えます。 | 
**ApprovalRequestForm** | [**ApprovalRequestResponseApprovalRequestApprovalRequestForm**](approvalRequestResponse_approval_request_approval_request_form.md) |  | 

## Methods

### NewApprovalRequestResponseApprovalRequest

`func NewApprovalRequestResponseApprovalRequest(id int32, companyId int32, applicationDate string, title string, applicantId int32, approvers []ExpenseApplicationResponseExpenseApplicationApprovers, applicationNumber string, status string, requestItems []ApprovalRequestsIndexResponseRequestItems, formId int32, approvalFlowRouteId int32, comments []ExpenseApplicationResponseExpenseApplicationComments, approvalFlowLogs []ExpenseApplicationResponseExpenseApplicationApprovalFlowLogs, currentStepId NullableInt32, currentRound int32, approvalRequestForm ApprovalRequestResponseApprovalRequestApprovalRequestForm, ) *ApprovalRequestResponseApprovalRequest`

NewApprovalRequestResponseApprovalRequest instantiates a new ApprovalRequestResponseApprovalRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApprovalRequestResponseApprovalRequestWithDefaults

`func NewApprovalRequestResponseApprovalRequestWithDefaults() *ApprovalRequestResponseApprovalRequest`

NewApprovalRequestResponseApprovalRequestWithDefaults instantiates a new ApprovalRequestResponseApprovalRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApprovalRequestResponseApprovalRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApprovalRequestResponseApprovalRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApprovalRequestResponseApprovalRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetCompanyId

`func (o *ApprovalRequestResponseApprovalRequest) GetCompanyId() int32`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *ApprovalRequestResponseApprovalRequest) GetCompanyIdOk() (*int32, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *ApprovalRequestResponseApprovalRequest) SetCompanyId(v int32)`

SetCompanyId sets CompanyId field to given value.


### GetApplicationDate

`func (o *ApprovalRequestResponseApprovalRequest) GetApplicationDate() string`

GetApplicationDate returns the ApplicationDate field if non-nil, zero value otherwise.

### GetApplicationDateOk

`func (o *ApprovalRequestResponseApprovalRequest) GetApplicationDateOk() (*string, bool)`

GetApplicationDateOk returns a tuple with the ApplicationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationDate

`func (o *ApprovalRequestResponseApprovalRequest) SetApplicationDate(v string)`

SetApplicationDate sets ApplicationDate field to given value.


### GetTitle

`func (o *ApprovalRequestResponseApprovalRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ApprovalRequestResponseApprovalRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ApprovalRequestResponseApprovalRequest) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetApplicantId

`func (o *ApprovalRequestResponseApprovalRequest) GetApplicantId() int32`

GetApplicantId returns the ApplicantId field if non-nil, zero value otherwise.

### GetApplicantIdOk

`func (o *ApprovalRequestResponseApprovalRequest) GetApplicantIdOk() (*int32, bool)`

GetApplicantIdOk returns a tuple with the ApplicantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicantId

`func (o *ApprovalRequestResponseApprovalRequest) SetApplicantId(v int32)`

SetApplicantId sets ApplicantId field to given value.


### GetApprovers

`func (o *ApprovalRequestResponseApprovalRequest) GetApprovers() []ExpenseApplicationResponseExpenseApplicationApprovers`

GetApprovers returns the Approvers field if non-nil, zero value otherwise.

### GetApproversOk

`func (o *ApprovalRequestResponseApprovalRequest) GetApproversOk() (*[]ExpenseApplicationResponseExpenseApplicationApprovers, bool)`

GetApproversOk returns a tuple with the Approvers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovers

`func (o *ApprovalRequestResponseApprovalRequest) SetApprovers(v []ExpenseApplicationResponseExpenseApplicationApprovers)`

SetApprovers sets Approvers field to given value.


### GetApplicationNumber

`func (o *ApprovalRequestResponseApprovalRequest) GetApplicationNumber() string`

GetApplicationNumber returns the ApplicationNumber field if non-nil, zero value otherwise.

### GetApplicationNumberOk

`func (o *ApprovalRequestResponseApprovalRequest) GetApplicationNumberOk() (*string, bool)`

GetApplicationNumberOk returns a tuple with the ApplicationNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationNumber

`func (o *ApprovalRequestResponseApprovalRequest) SetApplicationNumber(v string)`

SetApplicationNumber sets ApplicationNumber field to given value.


### GetStatus

`func (o *ApprovalRequestResponseApprovalRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApprovalRequestResponseApprovalRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApprovalRequestResponseApprovalRequest) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetRequestItems

`func (o *ApprovalRequestResponseApprovalRequest) GetRequestItems() []ApprovalRequestsIndexResponseRequestItems`

GetRequestItems returns the RequestItems field if non-nil, zero value otherwise.

### GetRequestItemsOk

`func (o *ApprovalRequestResponseApprovalRequest) GetRequestItemsOk() (*[]ApprovalRequestsIndexResponseRequestItems, bool)`

GetRequestItemsOk returns a tuple with the RequestItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestItems

`func (o *ApprovalRequestResponseApprovalRequest) SetRequestItems(v []ApprovalRequestsIndexResponseRequestItems)`

SetRequestItems sets RequestItems field to given value.


### GetFormId

`func (o *ApprovalRequestResponseApprovalRequest) GetFormId() int32`

GetFormId returns the FormId field if non-nil, zero value otherwise.

### GetFormIdOk

`func (o *ApprovalRequestResponseApprovalRequest) GetFormIdOk() (*int32, bool)`

GetFormIdOk returns a tuple with the FormId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormId

`func (o *ApprovalRequestResponseApprovalRequest) SetFormId(v int32)`

SetFormId sets FormId field to given value.


### GetApprovalFlowRouteId

`func (o *ApprovalRequestResponseApprovalRequest) GetApprovalFlowRouteId() int32`

GetApprovalFlowRouteId returns the ApprovalFlowRouteId field if non-nil, zero value otherwise.

### GetApprovalFlowRouteIdOk

`func (o *ApprovalRequestResponseApprovalRequest) GetApprovalFlowRouteIdOk() (*int32, bool)`

GetApprovalFlowRouteIdOk returns a tuple with the ApprovalFlowRouteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalFlowRouteId

`func (o *ApprovalRequestResponseApprovalRequest) SetApprovalFlowRouteId(v int32)`

SetApprovalFlowRouteId sets ApprovalFlowRouteId field to given value.


### GetComments

`func (o *ApprovalRequestResponseApprovalRequest) GetComments() []ExpenseApplicationResponseExpenseApplicationComments`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *ApprovalRequestResponseApprovalRequest) GetCommentsOk() (*[]ExpenseApplicationResponseExpenseApplicationComments, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *ApprovalRequestResponseApprovalRequest) SetComments(v []ExpenseApplicationResponseExpenseApplicationComments)`

SetComments sets Comments field to given value.


### GetApprovalFlowLogs

`func (o *ApprovalRequestResponseApprovalRequest) GetApprovalFlowLogs() []ExpenseApplicationResponseExpenseApplicationApprovalFlowLogs`

GetApprovalFlowLogs returns the ApprovalFlowLogs field if non-nil, zero value otherwise.

### GetApprovalFlowLogsOk

`func (o *ApprovalRequestResponseApprovalRequest) GetApprovalFlowLogsOk() (*[]ExpenseApplicationResponseExpenseApplicationApprovalFlowLogs, bool)`

GetApprovalFlowLogsOk returns a tuple with the ApprovalFlowLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalFlowLogs

`func (o *ApprovalRequestResponseApprovalRequest) SetApprovalFlowLogs(v []ExpenseApplicationResponseExpenseApplicationApprovalFlowLogs)`

SetApprovalFlowLogs sets ApprovalFlowLogs field to given value.


### GetCurrentStepId

`func (o *ApprovalRequestResponseApprovalRequest) GetCurrentStepId() int32`

GetCurrentStepId returns the CurrentStepId field if non-nil, zero value otherwise.

### GetCurrentStepIdOk

`func (o *ApprovalRequestResponseApprovalRequest) GetCurrentStepIdOk() (*int32, bool)`

GetCurrentStepIdOk returns a tuple with the CurrentStepId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentStepId

`func (o *ApprovalRequestResponseApprovalRequest) SetCurrentStepId(v int32)`

SetCurrentStepId sets CurrentStepId field to given value.


### SetCurrentStepIdNil

`func (o *ApprovalRequestResponseApprovalRequest) SetCurrentStepIdNil(b bool)`

 SetCurrentStepIdNil sets the value for CurrentStepId to be an explicit nil

### UnsetCurrentStepId
`func (o *ApprovalRequestResponseApprovalRequest) UnsetCurrentStepId()`

UnsetCurrentStepId ensures that no value is present for CurrentStepId, not even an explicit nil
### GetCurrentRound

`func (o *ApprovalRequestResponseApprovalRequest) GetCurrentRound() int32`

GetCurrentRound returns the CurrentRound field if non-nil, zero value otherwise.

### GetCurrentRoundOk

`func (o *ApprovalRequestResponseApprovalRequest) GetCurrentRoundOk() (*int32, bool)`

GetCurrentRoundOk returns a tuple with the CurrentRound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentRound

`func (o *ApprovalRequestResponseApprovalRequest) SetCurrentRound(v int32)`

SetCurrentRound sets CurrentRound field to given value.


### GetApprovalRequestForm

`func (o *ApprovalRequestResponseApprovalRequest) GetApprovalRequestForm() ApprovalRequestResponseApprovalRequestApprovalRequestForm`

GetApprovalRequestForm returns the ApprovalRequestForm field if non-nil, zero value otherwise.

### GetApprovalRequestFormOk

`func (o *ApprovalRequestResponseApprovalRequest) GetApprovalRequestFormOk() (*ApprovalRequestResponseApprovalRequestApprovalRequestForm, bool)`

GetApprovalRequestFormOk returns a tuple with the ApprovalRequestForm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalRequestForm

`func (o *ApprovalRequestResponseApprovalRequest) SetApprovalRequestForm(v ApprovalRequestResponseApprovalRequestApprovalRequestForm)`

SetApprovalRequestForm sets ApprovalRequestForm field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


