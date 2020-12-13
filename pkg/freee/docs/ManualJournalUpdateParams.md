# ManualJournalUpdateParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyId** | **int32** | 事業所ID | 
**IssueDate** | **string** | 発生日 (yyyy-mm-dd) | 
**Adjustment** | Pointer to **bool** | 決算整理仕訳フラグ（falseまたは未指定の場合: 日常仕訳） | [optional] 
**Details** | [**[]ManualJournalUpdateParamsDetails**](ManualJournalUpdateParamsDetails.md) |  | 

## Methods

### NewManualJournalUpdateParams

`func NewManualJournalUpdateParams(companyId int32, issueDate string, details []ManualJournalUpdateParamsDetails, ) *ManualJournalUpdateParams`

NewManualJournalUpdateParams instantiates a new ManualJournalUpdateParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManualJournalUpdateParamsWithDefaults

`func NewManualJournalUpdateParamsWithDefaults() *ManualJournalUpdateParams`

NewManualJournalUpdateParamsWithDefaults instantiates a new ManualJournalUpdateParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyId

`func (o *ManualJournalUpdateParams) GetCompanyId() int32`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *ManualJournalUpdateParams) GetCompanyIdOk() (*int32, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *ManualJournalUpdateParams) SetCompanyId(v int32)`

SetCompanyId sets CompanyId field to given value.


### GetIssueDate

`func (o *ManualJournalUpdateParams) GetIssueDate() string`

GetIssueDate returns the IssueDate field if non-nil, zero value otherwise.

### GetIssueDateOk

`func (o *ManualJournalUpdateParams) GetIssueDateOk() (*string, bool)`

GetIssueDateOk returns a tuple with the IssueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueDate

`func (o *ManualJournalUpdateParams) SetIssueDate(v string)`

SetIssueDate sets IssueDate field to given value.


### GetAdjustment

`func (o *ManualJournalUpdateParams) GetAdjustment() bool`

GetAdjustment returns the Adjustment field if non-nil, zero value otherwise.

### GetAdjustmentOk

`func (o *ManualJournalUpdateParams) GetAdjustmentOk() (*bool, bool)`

GetAdjustmentOk returns a tuple with the Adjustment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustment

`func (o *ManualJournalUpdateParams) SetAdjustment(v bool)`

SetAdjustment sets Adjustment field to given value.

### HasAdjustment

`func (o *ManualJournalUpdateParams) HasAdjustment() bool`

HasAdjustment returns a boolean if a field has been set.

### GetDetails

`func (o *ManualJournalUpdateParams) GetDetails() []ManualJournalUpdateParamsDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *ManualJournalUpdateParams) GetDetailsOk() (*[]ManualJournalUpdateParamsDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *ManualJournalUpdateParams) SetDetails(v []ManualJournalUpdateParamsDetails)`

SetDetails sets Details field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


