# JournalStatusResponseJournals

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | 受け付けID | 
**CompanyId** | **int32** | 事業所ID | 
**DownloadType** | **string** | ダウンロード形式 | 
**Status** | **string** | 事業所ID | 
**StartDate** | **string** | 取得開始日 (yyyy-mm-dd) | 
**EndDate** | **string** | 取得終了日 (yyyy-mm-dd) | 
**VisibleTags** | Pointer to **[]string** |  | [optional] 
**DownloadUrl** | Pointer to **string** | ダウンロードURL | [optional] 

## Methods

### NewJournalStatusResponseJournals

`func NewJournalStatusResponseJournals(id int32, companyId int32, downloadType string, status string, startDate string, endDate string, ) *JournalStatusResponseJournals`

NewJournalStatusResponseJournals instantiates a new JournalStatusResponseJournals object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJournalStatusResponseJournalsWithDefaults

`func NewJournalStatusResponseJournalsWithDefaults() *JournalStatusResponseJournals`

NewJournalStatusResponseJournalsWithDefaults instantiates a new JournalStatusResponseJournals object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *JournalStatusResponseJournals) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *JournalStatusResponseJournals) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *JournalStatusResponseJournals) SetId(v int32)`

SetId sets Id field to given value.


### GetCompanyId

`func (o *JournalStatusResponseJournals) GetCompanyId() int32`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *JournalStatusResponseJournals) GetCompanyIdOk() (*int32, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *JournalStatusResponseJournals) SetCompanyId(v int32)`

SetCompanyId sets CompanyId field to given value.


### GetDownloadType

`func (o *JournalStatusResponseJournals) GetDownloadType() string`

GetDownloadType returns the DownloadType field if non-nil, zero value otherwise.

### GetDownloadTypeOk

`func (o *JournalStatusResponseJournals) GetDownloadTypeOk() (*string, bool)`

GetDownloadTypeOk returns a tuple with the DownloadType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadType

`func (o *JournalStatusResponseJournals) SetDownloadType(v string)`

SetDownloadType sets DownloadType field to given value.


### GetStatus

`func (o *JournalStatusResponseJournals) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *JournalStatusResponseJournals) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *JournalStatusResponseJournals) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetStartDate

`func (o *JournalStatusResponseJournals) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *JournalStatusResponseJournals) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *JournalStatusResponseJournals) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.


### GetEndDate

`func (o *JournalStatusResponseJournals) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *JournalStatusResponseJournals) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *JournalStatusResponseJournals) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.


### GetVisibleTags

`func (o *JournalStatusResponseJournals) GetVisibleTags() []string`

GetVisibleTags returns the VisibleTags field if non-nil, zero value otherwise.

### GetVisibleTagsOk

`func (o *JournalStatusResponseJournals) GetVisibleTagsOk() (*[]string, bool)`

GetVisibleTagsOk returns a tuple with the VisibleTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibleTags

`func (o *JournalStatusResponseJournals) SetVisibleTags(v []string)`

SetVisibleTags sets VisibleTags field to given value.

### HasVisibleTags

`func (o *JournalStatusResponseJournals) HasVisibleTags() bool`

HasVisibleTags returns a boolean if a field has been set.

### GetDownloadUrl

`func (o *JournalStatusResponseJournals) GetDownloadUrl() string`

GetDownloadUrl returns the DownloadUrl field if non-nil, zero value otherwise.

### GetDownloadUrlOk

`func (o *JournalStatusResponseJournals) GetDownloadUrlOk() (*string, bool)`

GetDownloadUrlOk returns a tuple with the DownloadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadUrl

`func (o *JournalStatusResponseJournals) SetDownloadUrl(v string)`

SetDownloadUrl sets DownloadUrl field to given value.

### HasDownloadUrl

`func (o *JournalStatusResponseJournals) HasDownloadUrl() bool`

HasDownloadUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


