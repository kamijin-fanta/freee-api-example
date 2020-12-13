# UserResponseUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | ユーザーID | 
**Email** | **string** | メールアドレス | 
**DisplayName** | Pointer to **NullableString** | 表示名 | [optional] 
**FirstName** | Pointer to **NullableString** | 氏名（名） | [optional] 
**LastName** | Pointer to **NullableString** | 氏名（姓） | [optional] 
**FirstNameKana** | Pointer to **NullableString** | 氏名（カナ・名） | [optional] 
**LastNameKana** | Pointer to **NullableString** | 氏名（カナ・姓） | [optional] 

## Methods

### NewUserResponseUser

`func NewUserResponseUser(id int32, email string, ) *UserResponseUser`

NewUserResponseUser instantiates a new UserResponseUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserResponseUserWithDefaults

`func NewUserResponseUserWithDefaults() *UserResponseUser`

NewUserResponseUserWithDefaults instantiates a new UserResponseUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserResponseUser) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserResponseUser) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserResponseUser) SetId(v int32)`

SetId sets Id field to given value.


### GetEmail

`func (o *UserResponseUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserResponseUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserResponseUser) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetDisplayName

`func (o *UserResponseUser) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *UserResponseUser) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *UserResponseUser) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *UserResponseUser) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *UserResponseUser) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *UserResponseUser) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetFirstName

`func (o *UserResponseUser) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *UserResponseUser) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *UserResponseUser) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *UserResponseUser) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *UserResponseUser) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *UserResponseUser) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetLastName

`func (o *UserResponseUser) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *UserResponseUser) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *UserResponseUser) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *UserResponseUser) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *UserResponseUser) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *UserResponseUser) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetFirstNameKana

`func (o *UserResponseUser) GetFirstNameKana() string`

GetFirstNameKana returns the FirstNameKana field if non-nil, zero value otherwise.

### GetFirstNameKanaOk

`func (o *UserResponseUser) GetFirstNameKanaOk() (*string, bool)`

GetFirstNameKanaOk returns a tuple with the FirstNameKana field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstNameKana

`func (o *UserResponseUser) SetFirstNameKana(v string)`

SetFirstNameKana sets FirstNameKana field to given value.

### HasFirstNameKana

`func (o *UserResponseUser) HasFirstNameKana() bool`

HasFirstNameKana returns a boolean if a field has been set.

### SetFirstNameKanaNil

`func (o *UserResponseUser) SetFirstNameKanaNil(b bool)`

 SetFirstNameKanaNil sets the value for FirstNameKana to be an explicit nil

### UnsetFirstNameKana
`func (o *UserResponseUser) UnsetFirstNameKana()`

UnsetFirstNameKana ensures that no value is present for FirstNameKana, not even an explicit nil
### GetLastNameKana

`func (o *UserResponseUser) GetLastNameKana() string`

GetLastNameKana returns the LastNameKana field if non-nil, zero value otherwise.

### GetLastNameKanaOk

`func (o *UserResponseUser) GetLastNameKanaOk() (*string, bool)`

GetLastNameKanaOk returns a tuple with the LastNameKana field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastNameKana

`func (o *UserResponseUser) SetLastNameKana(v string)`

SetLastNameKana sets LastNameKana field to given value.

### HasLastNameKana

`func (o *UserResponseUser) HasLastNameKana() bool`

HasLastNameKana returns a boolean if a field has been set.

### SetLastNameKanaNil

`func (o *UserResponseUser) SetLastNameKanaNil(b bool)`

 SetLastNameKanaNil sets the value for LastNameKana to be an explicit nil

### UnsetLastNameKana
`func (o *UserResponseUser) UnsetLastNameKana()`

UnsetLastNameKana ensures that no value is present for LastNameKana, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


