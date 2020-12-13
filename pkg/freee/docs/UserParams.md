# UserParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | 表示名 (20文字以内) | [optional] 
**FirstName** | Pointer to **string** | 氏名（名） (20文字以内) | [optional] 
**LastName** | Pointer to **string** | 氏名（姓） (20文字以内) | [optional] 
**FirstNameKana** | Pointer to **string** | 氏名（カナ・名） (20文字以内) | [optional] 
**LastNameKana** | Pointer to **string** | 氏名（カナ・姓） (20文字以内) | [optional] 

## Methods

### NewUserParams

`func NewUserParams() *UserParams`

NewUserParams instantiates a new UserParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserParamsWithDefaults

`func NewUserParamsWithDefaults() *UserParams`

NewUserParamsWithDefaults instantiates a new UserParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *UserParams) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *UserParams) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *UserParams) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *UserParams) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetFirstName

`func (o *UserParams) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *UserParams) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *UserParams) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *UserParams) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *UserParams) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *UserParams) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *UserParams) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *UserParams) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetFirstNameKana

`func (o *UserParams) GetFirstNameKana() string`

GetFirstNameKana returns the FirstNameKana field if non-nil, zero value otherwise.

### GetFirstNameKanaOk

`func (o *UserParams) GetFirstNameKanaOk() (*string, bool)`

GetFirstNameKanaOk returns a tuple with the FirstNameKana field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstNameKana

`func (o *UserParams) SetFirstNameKana(v string)`

SetFirstNameKana sets FirstNameKana field to given value.

### HasFirstNameKana

`func (o *UserParams) HasFirstNameKana() bool`

HasFirstNameKana returns a boolean if a field has been set.

### GetLastNameKana

`func (o *UserParams) GetLastNameKana() string`

GetLastNameKana returns the LastNameKana field if non-nil, zero value otherwise.

### GetLastNameKanaOk

`func (o *UserParams) GetLastNameKanaOk() (*string, bool)`

GetLastNameKanaOk returns a tuple with the LastNameKana field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastNameKana

`func (o *UserParams) SetLastNameKana(v string)`

SetLastNameKana sets LastNameKana field to given value.

### HasLastNameKana

`func (o *UserParams) HasLastNameKana() bool`

HasLastNameKana returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


