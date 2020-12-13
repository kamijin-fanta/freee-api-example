# DeprecatedApprovalRequestParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyId** | **int32** | 事業所ID | 
**RequestFormId** | **int32** | 申請フォームID | 
**ApprovalFlowRouteId** | **int32** | 経路申請ID | 
**ApproverId** | Pointer to **int32** | 承認者のユーザーID | [optional] 
**Title** | **string** | 申請タイトル (255文字以内) | 
**RequestItems** | Pointer to [**[]DeprecatedApprovalRequestParamsRequestItems**](DeprecatedApprovalRequestParamsRequestItems.md) |  | [optional] 

## Methods

### NewDeprecatedApprovalRequestParams

`func NewDeprecatedApprovalRequestParams(companyId int32, requestFormId int32, approvalFlowRouteId int32, title string, ) *DeprecatedApprovalRequestParams`

NewDeprecatedApprovalRequestParams instantiates a new DeprecatedApprovalRequestParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeprecatedApprovalRequestParamsWithDefaults

`func NewDeprecatedApprovalRequestParamsWithDefaults() *DeprecatedApprovalRequestParams`

NewDeprecatedApprovalRequestParamsWithDefaults instantiates a new DeprecatedApprovalRequestParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyId

`func (o *DeprecatedApprovalRequestParams) GetCompanyId() int32`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *DeprecatedApprovalRequestParams) GetCompanyIdOk() (*int32, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *DeprecatedApprovalRequestParams) SetCompanyId(v int32)`

SetCompanyId sets CompanyId field to given value.


### GetRequestFormId

`func (o *DeprecatedApprovalRequestParams) GetRequestFormId() int32`

GetRequestFormId returns the RequestFormId field if non-nil, zero value otherwise.

### GetRequestFormIdOk

`func (o *DeprecatedApprovalRequestParams) GetRequestFormIdOk() (*int32, bool)`

GetRequestFormIdOk returns a tuple with the RequestFormId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestFormId

`func (o *DeprecatedApprovalRequestParams) SetRequestFormId(v int32)`

SetRequestFormId sets RequestFormId field to given value.


### GetApprovalFlowRouteId

`func (o *DeprecatedApprovalRequestParams) GetApprovalFlowRouteId() int32`

GetApprovalFlowRouteId returns the ApprovalFlowRouteId field if non-nil, zero value otherwise.

### GetApprovalFlowRouteIdOk

`func (o *DeprecatedApprovalRequestParams) GetApprovalFlowRouteIdOk() (*int32, bool)`

GetApprovalFlowRouteIdOk returns a tuple with the ApprovalFlowRouteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalFlowRouteId

`func (o *DeprecatedApprovalRequestParams) SetApprovalFlowRouteId(v int32)`

SetApprovalFlowRouteId sets ApprovalFlowRouteId field to given value.


### GetApproverId

`func (o *DeprecatedApprovalRequestParams) GetApproverId() int32`

GetApproverId returns the ApproverId field if non-nil, zero value otherwise.

### GetApproverIdOk

`func (o *DeprecatedApprovalRequestParams) GetApproverIdOk() (*int32, bool)`

GetApproverIdOk returns a tuple with the ApproverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproverId

`func (o *DeprecatedApprovalRequestParams) SetApproverId(v int32)`

SetApproverId sets ApproverId field to given value.

### HasApproverId

`func (o *DeprecatedApprovalRequestParams) HasApproverId() bool`

HasApproverId returns a boolean if a field has been set.

### GetTitle

`func (o *DeprecatedApprovalRequestParams) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *DeprecatedApprovalRequestParams) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *DeprecatedApprovalRequestParams) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetRequestItems

`func (o *DeprecatedApprovalRequestParams) GetRequestItems() []DeprecatedApprovalRequestParamsRequestItems`

GetRequestItems returns the RequestItems field if non-nil, zero value otherwise.

### GetRequestItemsOk

`func (o *DeprecatedApprovalRequestParams) GetRequestItemsOk() (*[]DeprecatedApprovalRequestParamsRequestItems, bool)`

GetRequestItemsOk returns a tuple with the RequestItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestItems

`func (o *DeprecatedApprovalRequestParams) SetRequestItems(v []DeprecatedApprovalRequestParamsRequestItems)`

SetRequestItems sets RequestItems field to given value.

### HasRequestItems

`func (o *DeprecatedApprovalRequestParams) HasRequestItems() bool`

HasRequestItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


