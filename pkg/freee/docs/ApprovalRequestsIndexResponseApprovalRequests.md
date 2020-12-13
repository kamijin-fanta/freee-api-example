# ApprovalRequestsIndexResponseApprovalRequests

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | 各種申請ID | 
**CompanyId** | **int32** | 事業所ID | 
**ApplicationDate** | **string** | 申請日 (yyyy-mm-dd) | 
**Title** | **string** | 申請タイトル | 
**ApplicantId** | **int32** | 申請者のユーザーID | 
**ApplicationNumber** | **string** | 申請No. | 
**Status** | **string** | 申請ステータス(draft:下書き, in_progress:申請中, approved:承認済, rejected:却下, feedback:差戻し) | 
**RequestItems** | [**[]ApprovalRequestsIndexResponseRequestItems**](ApprovalRequestsIndexResponseRequestItems.md) | 各種申請の項目一覧（配列） | 
**FormId** | **int32** | 申請フォームID | 
**CurrentStepId** | **NullableInt32** | 現在承認ステップID | 
**CurrentRound** | **int32** | 現在のround。差し戻し等により申請がstepの最初からやり直しになるとroundの値が増えます。 | 

## Methods

### NewApprovalRequestsIndexResponseApprovalRequests

`func NewApprovalRequestsIndexResponseApprovalRequests(id int32, companyId int32, applicationDate string, title string, applicantId int32, applicationNumber string, status string, requestItems []ApprovalRequestsIndexResponseRequestItems, formId int32, currentStepId NullableInt32, currentRound int32, ) *ApprovalRequestsIndexResponseApprovalRequests`

NewApprovalRequestsIndexResponseApprovalRequests instantiates a new ApprovalRequestsIndexResponseApprovalRequests object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApprovalRequestsIndexResponseApprovalRequestsWithDefaults

`func NewApprovalRequestsIndexResponseApprovalRequestsWithDefaults() *ApprovalRequestsIndexResponseApprovalRequests`

NewApprovalRequestsIndexResponseApprovalRequestsWithDefaults instantiates a new ApprovalRequestsIndexResponseApprovalRequests object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApprovalRequestsIndexResponseApprovalRequests) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApprovalRequestsIndexResponseApprovalRequests) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApprovalRequestsIndexResponseApprovalRequests) SetId(v int32)`

SetId sets Id field to given value.


### GetCompanyId

`func (o *ApprovalRequestsIndexResponseApprovalRequests) GetCompanyId() int32`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *ApprovalRequestsIndexResponseApprovalRequests) GetCompanyIdOk() (*int32, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *ApprovalRequestsIndexResponseApprovalRequests) SetCompanyId(v int32)`

SetCompanyId sets CompanyId field to given value.


### GetApplicationDate

`func (o *ApprovalRequestsIndexResponseApprovalRequests) GetApplicationDate() string`

GetApplicationDate returns the ApplicationDate field if non-nil, zero value otherwise.

### GetApplicationDateOk

`func (o *ApprovalRequestsIndexResponseApprovalRequests) GetApplicationDateOk() (*string, bool)`

GetApplicationDateOk returns a tuple with the ApplicationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationDate

`func (o *ApprovalRequestsIndexResponseApprovalRequests) SetApplicationDate(v string)`

SetApplicationDate sets ApplicationDate field to given value.


### GetTitle

`func (o *ApprovalRequestsIndexResponseApprovalRequests) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ApprovalRequestsIndexResponseApprovalRequests) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ApprovalRequestsIndexResponseApprovalRequests) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetApplicantId

`func (o *ApprovalRequestsIndexResponseApprovalRequests) GetApplicantId() int32`

GetApplicantId returns the ApplicantId field if non-nil, zero value otherwise.

### GetApplicantIdOk

`func (o *ApprovalRequestsIndexResponseApprovalRequests) GetApplicantIdOk() (*int32, bool)`

GetApplicantIdOk returns a tuple with the ApplicantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicantId

`func (o *ApprovalRequestsIndexResponseApprovalRequests) SetApplicantId(v int32)`

SetApplicantId sets ApplicantId field to given value.


### GetApplicationNumber

`func (o *ApprovalRequestsIndexResponseApprovalRequests) GetApplicationNumber() string`

GetApplicationNumber returns the ApplicationNumber field if non-nil, zero value otherwise.

### GetApplicationNumberOk

`func (o *ApprovalRequestsIndexResponseApprovalRequests) GetApplicationNumberOk() (*string, bool)`

GetApplicationNumberOk returns a tuple with the ApplicationNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationNumber

`func (o *ApprovalRequestsIndexResponseApprovalRequests) SetApplicationNumber(v string)`

SetApplicationNumber sets ApplicationNumber field to given value.


### GetStatus

`func (o *ApprovalRequestsIndexResponseApprovalRequests) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApprovalRequestsIndexResponseApprovalRequests) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApprovalRequestsIndexResponseApprovalRequests) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetRequestItems

`func (o *ApprovalRequestsIndexResponseApprovalRequests) GetRequestItems() []ApprovalRequestsIndexResponseRequestItems`

GetRequestItems returns the RequestItems field if non-nil, zero value otherwise.

### GetRequestItemsOk

`func (o *ApprovalRequestsIndexResponseApprovalRequests) GetRequestItemsOk() (*[]ApprovalRequestsIndexResponseRequestItems, bool)`

GetRequestItemsOk returns a tuple with the RequestItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestItems

`func (o *ApprovalRequestsIndexResponseApprovalRequests) SetRequestItems(v []ApprovalRequestsIndexResponseRequestItems)`

SetRequestItems sets RequestItems field to given value.


### GetFormId

`func (o *ApprovalRequestsIndexResponseApprovalRequests) GetFormId() int32`

GetFormId returns the FormId field if non-nil, zero value otherwise.

### GetFormIdOk

`func (o *ApprovalRequestsIndexResponseApprovalRequests) GetFormIdOk() (*int32, bool)`

GetFormIdOk returns a tuple with the FormId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormId

`func (o *ApprovalRequestsIndexResponseApprovalRequests) SetFormId(v int32)`

SetFormId sets FormId field to given value.


### GetCurrentStepId

`func (o *ApprovalRequestsIndexResponseApprovalRequests) GetCurrentStepId() int32`

GetCurrentStepId returns the CurrentStepId field if non-nil, zero value otherwise.

### GetCurrentStepIdOk

`func (o *ApprovalRequestsIndexResponseApprovalRequests) GetCurrentStepIdOk() (*int32, bool)`

GetCurrentStepIdOk returns a tuple with the CurrentStepId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentStepId

`func (o *ApprovalRequestsIndexResponseApprovalRequests) SetCurrentStepId(v int32)`

SetCurrentStepId sets CurrentStepId field to given value.


### SetCurrentStepIdNil

`func (o *ApprovalRequestsIndexResponseApprovalRequests) SetCurrentStepIdNil(b bool)`

 SetCurrentStepIdNil sets the value for CurrentStepId to be an explicit nil

### UnsetCurrentStepId
`func (o *ApprovalRequestsIndexResponseApprovalRequests) UnsetCurrentStepId()`

UnsetCurrentStepId ensures that no value is present for CurrentStepId, not even an explicit nil
### GetCurrentRound

`func (o *ApprovalRequestsIndexResponseApprovalRequests) GetCurrentRound() int32`

GetCurrentRound returns the CurrentRound field if non-nil, zero value otherwise.

### GetCurrentRoundOk

`func (o *ApprovalRequestsIndexResponseApprovalRequests) GetCurrentRoundOk() (*int32, bool)`

GetCurrentRoundOk returns a tuple with the CurrentRound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentRound

`func (o *ApprovalRequestsIndexResponseApprovalRequests) SetCurrentRound(v int32)`

SetCurrentRound sets CurrentRound field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


