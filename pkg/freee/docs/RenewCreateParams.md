# RenewCreateParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyId** | **int32** | 事業所ID | 
**UpdateDate** | **string** | 更新日 (yyyy-mm-dd) | 
**RenewTargetId** | **int32** | +更新対象行ID (details(取引の明細行), accruals(債権債務行), renewsのdetails(+更新の明細行)のIDを指定)  | 
**Details** | [**[]RenewCreateParamsDetails**](RenewCreateParamsDetails.md) | +更新の明細行 | 

## Methods

### NewRenewCreateParams

`func NewRenewCreateParams(companyId int32, updateDate string, renewTargetId int32, details []RenewCreateParamsDetails, ) *RenewCreateParams`

NewRenewCreateParams instantiates a new RenewCreateParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRenewCreateParamsWithDefaults

`func NewRenewCreateParamsWithDefaults() *RenewCreateParams`

NewRenewCreateParamsWithDefaults instantiates a new RenewCreateParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyId

`func (o *RenewCreateParams) GetCompanyId() int32`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *RenewCreateParams) GetCompanyIdOk() (*int32, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *RenewCreateParams) SetCompanyId(v int32)`

SetCompanyId sets CompanyId field to given value.


### GetUpdateDate

`func (o *RenewCreateParams) GetUpdateDate() string`

GetUpdateDate returns the UpdateDate field if non-nil, zero value otherwise.

### GetUpdateDateOk

`func (o *RenewCreateParams) GetUpdateDateOk() (*string, bool)`

GetUpdateDateOk returns a tuple with the UpdateDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateDate

`func (o *RenewCreateParams) SetUpdateDate(v string)`

SetUpdateDate sets UpdateDate field to given value.


### GetRenewTargetId

`func (o *RenewCreateParams) GetRenewTargetId() int32`

GetRenewTargetId returns the RenewTargetId field if non-nil, zero value otherwise.

### GetRenewTargetIdOk

`func (o *RenewCreateParams) GetRenewTargetIdOk() (*int32, bool)`

GetRenewTargetIdOk returns a tuple with the RenewTargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenewTargetId

`func (o *RenewCreateParams) SetRenewTargetId(v int32)`

SetRenewTargetId sets RenewTargetId field to given value.


### GetDetails

`func (o *RenewCreateParams) GetDetails() []RenewCreateParamsDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *RenewCreateParams) GetDetailsOk() (*[]RenewCreateParamsDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *RenewCreateParams) SetDetails(v []RenewCreateParamsDetails)`

SetDetails sets Details field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


