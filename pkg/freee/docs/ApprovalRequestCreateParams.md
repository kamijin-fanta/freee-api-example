# ApprovalRequestCreateParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyId** | **int32** | 事業所ID | 
**ApplicationDate** | **string** | 申請日 (yyyy-mm-dd) | 
**ApprovalFlowRouteId** | **int32** | 申請経路ID | 
**FormId** | **int32** | 申請フォームID | 
**ApproverId** | Pointer to **int32** | 承認者のユーザーID | [optional] 
**Draft** | **bool** | falseの時、in_progress:申請中で作成する。それ以外の時はdraft:下書きで作成する | 
**ParentId** | Pointer to **int32** | 親申請ID(既存各種申請IDのみ指定可能です。) | [optional] 
**RequestItems** | [**[]ApprovalRequestCreateParamsRequestItems**](ApprovalRequestCreateParamsRequestItems.md) |  | 

## Methods

### NewApprovalRequestCreateParams

`func NewApprovalRequestCreateParams(companyId int32, applicationDate string, approvalFlowRouteId int32, formId int32, draft bool, requestItems []ApprovalRequestCreateParamsRequestItems, ) *ApprovalRequestCreateParams`

NewApprovalRequestCreateParams instantiates a new ApprovalRequestCreateParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApprovalRequestCreateParamsWithDefaults

`func NewApprovalRequestCreateParamsWithDefaults() *ApprovalRequestCreateParams`

NewApprovalRequestCreateParamsWithDefaults instantiates a new ApprovalRequestCreateParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyId

`func (o *ApprovalRequestCreateParams) GetCompanyId() int32`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *ApprovalRequestCreateParams) GetCompanyIdOk() (*int32, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *ApprovalRequestCreateParams) SetCompanyId(v int32)`

SetCompanyId sets CompanyId field to given value.


### GetApplicationDate

`func (o *ApprovalRequestCreateParams) GetApplicationDate() string`

GetApplicationDate returns the ApplicationDate field if non-nil, zero value otherwise.

### GetApplicationDateOk

`func (o *ApprovalRequestCreateParams) GetApplicationDateOk() (*string, bool)`

GetApplicationDateOk returns a tuple with the ApplicationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationDate

`func (o *ApprovalRequestCreateParams) SetApplicationDate(v string)`

SetApplicationDate sets ApplicationDate field to given value.


### GetApprovalFlowRouteId

`func (o *ApprovalRequestCreateParams) GetApprovalFlowRouteId() int32`

GetApprovalFlowRouteId returns the ApprovalFlowRouteId field if non-nil, zero value otherwise.

### GetApprovalFlowRouteIdOk

`func (o *ApprovalRequestCreateParams) GetApprovalFlowRouteIdOk() (*int32, bool)`

GetApprovalFlowRouteIdOk returns a tuple with the ApprovalFlowRouteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalFlowRouteId

`func (o *ApprovalRequestCreateParams) SetApprovalFlowRouteId(v int32)`

SetApprovalFlowRouteId sets ApprovalFlowRouteId field to given value.


### GetFormId

`func (o *ApprovalRequestCreateParams) GetFormId() int32`

GetFormId returns the FormId field if non-nil, zero value otherwise.

### GetFormIdOk

`func (o *ApprovalRequestCreateParams) GetFormIdOk() (*int32, bool)`

GetFormIdOk returns a tuple with the FormId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormId

`func (o *ApprovalRequestCreateParams) SetFormId(v int32)`

SetFormId sets FormId field to given value.


### GetApproverId

`func (o *ApprovalRequestCreateParams) GetApproverId() int32`

GetApproverId returns the ApproverId field if non-nil, zero value otherwise.

### GetApproverIdOk

`func (o *ApprovalRequestCreateParams) GetApproverIdOk() (*int32, bool)`

GetApproverIdOk returns a tuple with the ApproverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproverId

`func (o *ApprovalRequestCreateParams) SetApproverId(v int32)`

SetApproverId sets ApproverId field to given value.

### HasApproverId

`func (o *ApprovalRequestCreateParams) HasApproverId() bool`

HasApproverId returns a boolean if a field has been set.

### GetDraft

`func (o *ApprovalRequestCreateParams) GetDraft() bool`

GetDraft returns the Draft field if non-nil, zero value otherwise.

### GetDraftOk

`func (o *ApprovalRequestCreateParams) GetDraftOk() (*bool, bool)`

GetDraftOk returns a tuple with the Draft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraft

`func (o *ApprovalRequestCreateParams) SetDraft(v bool)`

SetDraft sets Draft field to given value.


### GetParentId

`func (o *ApprovalRequestCreateParams) GetParentId() int32`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *ApprovalRequestCreateParams) GetParentIdOk() (*int32, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *ApprovalRequestCreateParams) SetParentId(v int32)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *ApprovalRequestCreateParams) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetRequestItems

`func (o *ApprovalRequestCreateParams) GetRequestItems() []ApprovalRequestCreateParamsRequestItems`

GetRequestItems returns the RequestItems field if non-nil, zero value otherwise.

### GetRequestItemsOk

`func (o *ApprovalRequestCreateParams) GetRequestItemsOk() (*[]ApprovalRequestCreateParamsRequestItems, bool)`

GetRequestItemsOk returns a tuple with the RequestItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestItems

`func (o *ApprovalRequestCreateParams) SetRequestItems(v []ApprovalRequestCreateParamsRequestItems)`

SetRequestItems sets RequestItems field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


