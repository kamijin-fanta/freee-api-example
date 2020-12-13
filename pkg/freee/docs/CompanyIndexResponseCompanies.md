# CompanyIndexResponseCompanies

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | 事業所ID | 
**Name** | **NullableString** | 事業所名 | 
**NameKana** | **NullableString** | 事業所名（カナ） | 
**DisplayName** | **NullableString** | 事業所名 | 
**Role** | **string** | ユーザーの権限 | 

## Methods

### NewCompanyIndexResponseCompanies

`func NewCompanyIndexResponseCompanies(id int32, name NullableString, nameKana NullableString, displayName NullableString, role string, ) *CompanyIndexResponseCompanies`

NewCompanyIndexResponseCompanies instantiates a new CompanyIndexResponseCompanies object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyIndexResponseCompaniesWithDefaults

`func NewCompanyIndexResponseCompaniesWithDefaults() *CompanyIndexResponseCompanies`

NewCompanyIndexResponseCompaniesWithDefaults instantiates a new CompanyIndexResponseCompanies object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CompanyIndexResponseCompanies) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CompanyIndexResponseCompanies) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CompanyIndexResponseCompanies) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *CompanyIndexResponseCompanies) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CompanyIndexResponseCompanies) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CompanyIndexResponseCompanies) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *CompanyIndexResponseCompanies) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CompanyIndexResponseCompanies) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetNameKana

`func (o *CompanyIndexResponseCompanies) GetNameKana() string`

GetNameKana returns the NameKana field if non-nil, zero value otherwise.

### GetNameKanaOk

`func (o *CompanyIndexResponseCompanies) GetNameKanaOk() (*string, bool)`

GetNameKanaOk returns a tuple with the NameKana field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameKana

`func (o *CompanyIndexResponseCompanies) SetNameKana(v string)`

SetNameKana sets NameKana field to given value.


### SetNameKanaNil

`func (o *CompanyIndexResponseCompanies) SetNameKanaNil(b bool)`

 SetNameKanaNil sets the value for NameKana to be an explicit nil

### UnsetNameKana
`func (o *CompanyIndexResponseCompanies) UnsetNameKana()`

UnsetNameKana ensures that no value is present for NameKana, not even an explicit nil
### GetDisplayName

`func (o *CompanyIndexResponseCompanies) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CompanyIndexResponseCompanies) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CompanyIndexResponseCompanies) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### SetDisplayNameNil

`func (o *CompanyIndexResponseCompanies) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *CompanyIndexResponseCompanies) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetRole

`func (o *CompanyIndexResponseCompanies) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *CompanyIndexResponseCompanies) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *CompanyIndexResponseCompanies) SetRole(v string)`

SetRole sets Role field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


