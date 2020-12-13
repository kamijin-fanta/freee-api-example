# RenewUpdateParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyId** | **int32** | 事業所ID | 
**UpdateDate** | **string** | 更新日 (yyyy-mm-dd) | 
**Details** | [**[]RenewUpdateParamsDetails**](RenewUpdateParamsDetails.md) | +更新の明細行 | 

## Methods

### NewRenewUpdateParams

`func NewRenewUpdateParams(companyId int32, updateDate string, details []RenewUpdateParamsDetails, ) *RenewUpdateParams`

NewRenewUpdateParams instantiates a new RenewUpdateParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRenewUpdateParamsWithDefaults

`func NewRenewUpdateParamsWithDefaults() *RenewUpdateParams`

NewRenewUpdateParamsWithDefaults instantiates a new RenewUpdateParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyId

`func (o *RenewUpdateParams) GetCompanyId() int32`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *RenewUpdateParams) GetCompanyIdOk() (*int32, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *RenewUpdateParams) SetCompanyId(v int32)`

SetCompanyId sets CompanyId field to given value.


### GetUpdateDate

`func (o *RenewUpdateParams) GetUpdateDate() string`

GetUpdateDate returns the UpdateDate field if non-nil, zero value otherwise.

### GetUpdateDateOk

`func (o *RenewUpdateParams) GetUpdateDateOk() (*string, bool)`

GetUpdateDateOk returns a tuple with the UpdateDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateDate

`func (o *RenewUpdateParams) SetUpdateDate(v string)`

SetUpdateDate sets UpdateDate field to given value.


### GetDetails

`func (o *RenewUpdateParams) GetDetails() []RenewUpdateParamsDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *RenewUpdateParams) GetDetailsOk() (*[]RenewUpdateParamsDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *RenewUpdateParams) SetDetails(v []RenewUpdateParamsDetails)`

SetDetails sets Details field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


