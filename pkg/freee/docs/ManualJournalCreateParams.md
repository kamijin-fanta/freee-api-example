# ManualJournalCreateParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyId** | **int32** | 事業所ID | 
**IssueDate** | **string** | 発生日 (yyyy-mm-dd) | 
**Adjustment** | Pointer to **bool** | 決算整理仕訳フラグ（falseまたは未指定の場合: 日常仕訳） | [optional] 
**Details** | [**[]ManualJournalCreateParamsDetails**](ManualJournalCreateParamsDetails.md) |  | 

## Methods

### NewManualJournalCreateParams

`func NewManualJournalCreateParams(companyId int32, issueDate string, details []ManualJournalCreateParamsDetails, ) *ManualJournalCreateParams`

NewManualJournalCreateParams instantiates a new ManualJournalCreateParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManualJournalCreateParamsWithDefaults

`func NewManualJournalCreateParamsWithDefaults() *ManualJournalCreateParams`

NewManualJournalCreateParamsWithDefaults instantiates a new ManualJournalCreateParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyId

`func (o *ManualJournalCreateParams) GetCompanyId() int32`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *ManualJournalCreateParams) GetCompanyIdOk() (*int32, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *ManualJournalCreateParams) SetCompanyId(v int32)`

SetCompanyId sets CompanyId field to given value.


### GetIssueDate

`func (o *ManualJournalCreateParams) GetIssueDate() string`

GetIssueDate returns the IssueDate field if non-nil, zero value otherwise.

### GetIssueDateOk

`func (o *ManualJournalCreateParams) GetIssueDateOk() (*string, bool)`

GetIssueDateOk returns a tuple with the IssueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueDate

`func (o *ManualJournalCreateParams) SetIssueDate(v string)`

SetIssueDate sets IssueDate field to given value.


### GetAdjustment

`func (o *ManualJournalCreateParams) GetAdjustment() bool`

GetAdjustment returns the Adjustment field if non-nil, zero value otherwise.

### GetAdjustmentOk

`func (o *ManualJournalCreateParams) GetAdjustmentOk() (*bool, bool)`

GetAdjustmentOk returns a tuple with the Adjustment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustment

`func (o *ManualJournalCreateParams) SetAdjustment(v bool)`

SetAdjustment sets Adjustment field to given value.

### HasAdjustment

`func (o *ManualJournalCreateParams) HasAdjustment() bool`

HasAdjustment returns a boolean if a field has been set.

### GetDetails

`func (o *ManualJournalCreateParams) GetDetails() []ManualJournalCreateParamsDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *ManualJournalCreateParams) GetDetailsOk() (*[]ManualJournalCreateParamsDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *ManualJournalCreateParams) SetDetails(v []ManualJournalCreateParamsDetails)`

SetDetails sets Details field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


