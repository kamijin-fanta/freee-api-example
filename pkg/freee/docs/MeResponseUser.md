# MeResponseUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | ユーザーID | 
**Email** | **string** | メールアドレス | 
**DisplayName** | Pointer to **NullableString** | 表示ユーザー名 | [optional] 
**FirstName** | Pointer to **NullableString** | 名 | [optional] 
**LastName** | Pointer to **NullableString** | 姓 | [optional] 
**FirstNameKana** | Pointer to **NullableString** | 名（カナ） | [optional] 
**LastNameKana** | Pointer to **NullableString** | 姓（カナ） | [optional] 
**Companies** | Pointer to [**[]MeResponseUserCompanies**](MeResponseUserCompanies.md) |  | [optional] 

## Methods

### NewMeResponseUser

`func NewMeResponseUser(id int32, email string, ) *MeResponseUser`

NewMeResponseUser instantiates a new MeResponseUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMeResponseUserWithDefaults

`func NewMeResponseUserWithDefaults() *MeResponseUser`

NewMeResponseUserWithDefaults instantiates a new MeResponseUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MeResponseUser) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MeResponseUser) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MeResponseUser) SetId(v int32)`

SetId sets Id field to given value.


### GetEmail

`func (o *MeResponseUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *MeResponseUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *MeResponseUser) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetDisplayName

`func (o *MeResponseUser) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MeResponseUser) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MeResponseUser) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MeResponseUser) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *MeResponseUser) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *MeResponseUser) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetFirstName

`func (o *MeResponseUser) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *MeResponseUser) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *MeResponseUser) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *MeResponseUser) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *MeResponseUser) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *MeResponseUser) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetLastName

`func (o *MeResponseUser) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *MeResponseUser) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *MeResponseUser) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *MeResponseUser) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *MeResponseUser) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *MeResponseUser) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetFirstNameKana

`func (o *MeResponseUser) GetFirstNameKana() string`

GetFirstNameKana returns the FirstNameKana field if non-nil, zero value otherwise.

### GetFirstNameKanaOk

`func (o *MeResponseUser) GetFirstNameKanaOk() (*string, bool)`

GetFirstNameKanaOk returns a tuple with the FirstNameKana field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstNameKana

`func (o *MeResponseUser) SetFirstNameKana(v string)`

SetFirstNameKana sets FirstNameKana field to given value.

### HasFirstNameKana

`func (o *MeResponseUser) HasFirstNameKana() bool`

HasFirstNameKana returns a boolean if a field has been set.

### SetFirstNameKanaNil

`func (o *MeResponseUser) SetFirstNameKanaNil(b bool)`

 SetFirstNameKanaNil sets the value for FirstNameKana to be an explicit nil

### UnsetFirstNameKana
`func (o *MeResponseUser) UnsetFirstNameKana()`

UnsetFirstNameKana ensures that no value is present for FirstNameKana, not even an explicit nil
### GetLastNameKana

`func (o *MeResponseUser) GetLastNameKana() string`

GetLastNameKana returns the LastNameKana field if non-nil, zero value otherwise.

### GetLastNameKanaOk

`func (o *MeResponseUser) GetLastNameKanaOk() (*string, bool)`

GetLastNameKanaOk returns a tuple with the LastNameKana field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastNameKana

`func (o *MeResponseUser) SetLastNameKana(v string)`

SetLastNameKana sets LastNameKana field to given value.

### HasLastNameKana

`func (o *MeResponseUser) HasLastNameKana() bool`

HasLastNameKana returns a boolean if a field has been set.

### SetLastNameKanaNil

`func (o *MeResponseUser) SetLastNameKanaNil(b bool)`

 SetLastNameKanaNil sets the value for LastNameKana to be an explicit nil

### UnsetLastNameKana
`func (o *MeResponseUser) UnsetLastNameKana()`

UnsetLastNameKana ensures that no value is present for LastNameKana, not even an explicit nil
### GetCompanies

`func (o *MeResponseUser) GetCompanies() []MeResponseUserCompanies`

GetCompanies returns the Companies field if non-nil, zero value otherwise.

### GetCompaniesOk

`func (o *MeResponseUser) GetCompaniesOk() (*[]MeResponseUserCompanies, bool)`

GetCompaniesOk returns a tuple with the Companies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanies

`func (o *MeResponseUser) SetCompanies(v []MeResponseUserCompanies)`

SetCompanies sets Companies field to given value.

### HasCompanies

`func (o *MeResponseUser) HasCompanies() bool`

HasCompanies returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


