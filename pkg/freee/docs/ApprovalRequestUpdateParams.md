# ApprovalRequestUpdateParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyId** | **int32** | 事業所ID | 
**ApplicationDate** | **string** | 申請日 (yyyy-mm-dd) | 
**ApprovalFlowRouteId** | **int32** | 申請経路ID | 
**ApproverId** | Pointer to **int32** | 承認者のユーザーID | [optional] 
**Draft** | **bool** | falseの時、in_progress:申請中で更新する。それ以外の時はdraft:下書きで更新する | 
**RequestItems** | [**[]ApprovalRequestCreateParamsRequestItems**](ApprovalRequestCreateParamsRequestItems.md) |  | 

## Methods

### NewApprovalRequestUpdateParams

`func NewApprovalRequestUpdateParams(companyId int32, applicationDate string, approvalFlowRouteId int32, draft bool, requestItems []ApprovalRequestCreateParamsRequestItems, ) *ApprovalRequestUpdateParams`

NewApprovalRequestUpdateParams instantiates a new ApprovalRequestUpdateParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApprovalRequestUpdateParamsWithDefaults

`func NewApprovalRequestUpdateParamsWithDefaults() *ApprovalRequestUpdateParams`

NewApprovalRequestUpdateParamsWithDefaults instantiates a new ApprovalRequestUpdateParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyId

`func (o *ApprovalRequestUpdateParams) GetCompanyId() int32`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *ApprovalRequestUpdateParams) GetCompanyIdOk() (*int32, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *ApprovalRequestUpdateParams) SetCompanyId(v int32)`

SetCompanyId sets CompanyId field to given value.


### GetApplicationDate

`func (o *ApprovalRequestUpdateParams) GetApplicationDate() string`

GetApplicationDate returns the ApplicationDate field if non-nil, zero value otherwise.

### GetApplicationDateOk

`func (o *ApprovalRequestUpdateParams) GetApplicationDateOk() (*string, bool)`

GetApplicationDateOk returns a tuple with the ApplicationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationDate

`func (o *ApprovalRequestUpdateParams) SetApplicationDate(v string)`

SetApplicationDate sets ApplicationDate field to given value.


### GetApprovalFlowRouteId

`func (o *ApprovalRequestUpdateParams) GetApprovalFlowRouteId() int32`

GetApprovalFlowRouteId returns the ApprovalFlowRouteId field if non-nil, zero value otherwise.

### GetApprovalFlowRouteIdOk

`func (o *ApprovalRequestUpdateParams) GetApprovalFlowRouteIdOk() (*int32, bool)`

GetApprovalFlowRouteIdOk returns a tuple with the ApprovalFlowRouteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalFlowRouteId

`func (o *ApprovalRequestUpdateParams) SetApprovalFlowRouteId(v int32)`

SetApprovalFlowRouteId sets ApprovalFlowRouteId field to given value.


### GetApproverId

`func (o *ApprovalRequestUpdateParams) GetApproverId() int32`

GetApproverId returns the ApproverId field if non-nil, zero value otherwise.

### GetApproverIdOk

`func (o *ApprovalRequestUpdateParams) GetApproverIdOk() (*int32, bool)`

GetApproverIdOk returns a tuple with the ApproverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproverId

`func (o *ApprovalRequestUpdateParams) SetApproverId(v int32)`

SetApproverId sets ApproverId field to given value.

### HasApproverId

`func (o *ApprovalRequestUpdateParams) HasApproverId() bool`

HasApproverId returns a boolean if a field has been set.

### GetDraft

`func (o *ApprovalRequestUpdateParams) GetDraft() bool`

GetDraft returns the Draft field if non-nil, zero value otherwise.

### GetDraftOk

`func (o *ApprovalRequestUpdateParams) GetDraftOk() (*bool, bool)`

GetDraftOk returns a tuple with the Draft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraft

`func (o *ApprovalRequestUpdateParams) SetDraft(v bool)`

SetDraft sets Draft field to given value.


### GetRequestItems

`func (o *ApprovalRequestUpdateParams) GetRequestItems() []ApprovalRequestCreateParamsRequestItems`

GetRequestItems returns the RequestItems field if non-nil, zero value otherwise.

### GetRequestItemsOk

`func (o *ApprovalRequestUpdateParams) GetRequestItemsOk() (*[]ApprovalRequestCreateParamsRequestItems, bool)`

GetRequestItemsOk returns a tuple with the RequestItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestItems

`func (o *ApprovalRequestUpdateParams) SetRequestItems(v []ApprovalRequestCreateParamsRequestItems)`

SetRequestItems sets RequestItems field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


