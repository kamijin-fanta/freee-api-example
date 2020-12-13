# DeprecatedApprovalRequestResponseApprovalRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | 各種申請ID | 
**CompanyId** | **int32** | 事業所ID | 
**ApplicationDate** | **string** | 申請日 (yyyy-mm-dd) | 
**Title** | **string** | 申請タイトル | 
**ApplicantId** | **int32** | 申請者のユーザーID | 
**ApproverId** | **NullableInt32** | 承認者のユーザーID | 
**ApplicationNumber** | **string** | 申請No. | 
**Status** | **string** | 申請ステータス | 
**RequestItems** | [**[]DeprecatedApprovalRequestResponseApprovalRequestRequestItems**](DeprecatedApprovalRequestResponseApprovalRequestRequestItems.md) | 各種申請の項目一覧（配列） | 

## Methods

### NewDeprecatedApprovalRequestResponseApprovalRequest

`func NewDeprecatedApprovalRequestResponseApprovalRequest(id int32, companyId int32, applicationDate string, title string, applicantId int32, approverId NullableInt32, applicationNumber string, status string, requestItems []DeprecatedApprovalRequestResponseApprovalRequestRequestItems, ) *DeprecatedApprovalRequestResponseApprovalRequest`

NewDeprecatedApprovalRequestResponseApprovalRequest instantiates a new DeprecatedApprovalRequestResponseApprovalRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeprecatedApprovalRequestResponseApprovalRequestWithDefaults

`func NewDeprecatedApprovalRequestResponseApprovalRequestWithDefaults() *DeprecatedApprovalRequestResponseApprovalRequest`

NewDeprecatedApprovalRequestResponseApprovalRequestWithDefaults instantiates a new DeprecatedApprovalRequestResponseApprovalRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeprecatedApprovalRequestResponseApprovalRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeprecatedApprovalRequestResponseApprovalRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeprecatedApprovalRequestResponseApprovalRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetCompanyId

`func (o *DeprecatedApprovalRequestResponseApprovalRequest) GetCompanyId() int32`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *DeprecatedApprovalRequestResponseApprovalRequest) GetCompanyIdOk() (*int32, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *DeprecatedApprovalRequestResponseApprovalRequest) SetCompanyId(v int32)`

SetCompanyId sets CompanyId field to given value.


### GetApplicationDate

`func (o *DeprecatedApprovalRequestResponseApprovalRequest) GetApplicationDate() string`

GetApplicationDate returns the ApplicationDate field if non-nil, zero value otherwise.

### GetApplicationDateOk

`func (o *DeprecatedApprovalRequestResponseApprovalRequest) GetApplicationDateOk() (*string, bool)`

GetApplicationDateOk returns a tuple with the ApplicationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationDate

`func (o *DeprecatedApprovalRequestResponseApprovalRequest) SetApplicationDate(v string)`

SetApplicationDate sets ApplicationDate field to given value.


### GetTitle

`func (o *DeprecatedApprovalRequestResponseApprovalRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *DeprecatedApprovalRequestResponseApprovalRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *DeprecatedApprovalRequestResponseApprovalRequest) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetApplicantId

`func (o *DeprecatedApprovalRequestResponseApprovalRequest) GetApplicantId() int32`

GetApplicantId returns the ApplicantId field if non-nil, zero value otherwise.

### GetApplicantIdOk

`func (o *DeprecatedApprovalRequestResponseApprovalRequest) GetApplicantIdOk() (*int32, bool)`

GetApplicantIdOk returns a tuple with the ApplicantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicantId

`func (o *DeprecatedApprovalRequestResponseApprovalRequest) SetApplicantId(v int32)`

SetApplicantId sets ApplicantId field to given value.


### GetApproverId

`func (o *DeprecatedApprovalRequestResponseApprovalRequest) GetApproverId() int32`

GetApproverId returns the ApproverId field if non-nil, zero value otherwise.

### GetApproverIdOk

`func (o *DeprecatedApprovalRequestResponseApprovalRequest) GetApproverIdOk() (*int32, bool)`

GetApproverIdOk returns a tuple with the ApproverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproverId

`func (o *DeprecatedApprovalRequestResponseApprovalRequest) SetApproverId(v int32)`

SetApproverId sets ApproverId field to given value.


### SetApproverIdNil

`func (o *DeprecatedApprovalRequestResponseApprovalRequest) SetApproverIdNil(b bool)`

 SetApproverIdNil sets the value for ApproverId to be an explicit nil

### UnsetApproverId
`func (o *DeprecatedApprovalRequestResponseApprovalRequest) UnsetApproverId()`

UnsetApproverId ensures that no value is present for ApproverId, not even an explicit nil
### GetApplicationNumber

`func (o *DeprecatedApprovalRequestResponseApprovalRequest) GetApplicationNumber() string`

GetApplicationNumber returns the ApplicationNumber field if non-nil, zero value otherwise.

### GetApplicationNumberOk

`func (o *DeprecatedApprovalRequestResponseApprovalRequest) GetApplicationNumberOk() (*string, bool)`

GetApplicationNumberOk returns a tuple with the ApplicationNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationNumber

`func (o *DeprecatedApprovalRequestResponseApprovalRequest) SetApplicationNumber(v string)`

SetApplicationNumber sets ApplicationNumber field to given value.


### GetStatus

`func (o *DeprecatedApprovalRequestResponseApprovalRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DeprecatedApprovalRequestResponseApprovalRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DeprecatedApprovalRequestResponseApprovalRequest) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetRequestItems

`func (o *DeprecatedApprovalRequestResponseApprovalRequest) GetRequestItems() []DeprecatedApprovalRequestResponseApprovalRequestRequestItems`

GetRequestItems returns the RequestItems field if non-nil, zero value otherwise.

### GetRequestItemsOk

`func (o *DeprecatedApprovalRequestResponseApprovalRequest) GetRequestItemsOk() (*[]DeprecatedApprovalRequestResponseApprovalRequestRequestItems, bool)`

GetRequestItemsOk returns a tuple with the RequestItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestItems

`func (o *DeprecatedApprovalRequestResponseApprovalRequest) SetRequestItems(v []DeprecatedApprovalRequestResponseApprovalRequestRequestItems)`

SetRequestItems sets RequestItems field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


