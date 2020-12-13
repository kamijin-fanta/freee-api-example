# TagParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyId** | **int32** | 事業所ID | 
**Name** | **string** | メモタグ名 (30文字以内) | 
**Shortcut1** | Pointer to **string** | メモタグ検索用 (20文字以内) | [optional] 
**Shortcut2** | Pointer to **string** | メモタグ検索用 (20文字以内) | [optional] 

## Methods

### NewTagParams

`func NewTagParams(companyId int32, name string, ) *TagParams`

NewTagParams instantiates a new TagParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagParamsWithDefaults

`func NewTagParamsWithDefaults() *TagParams`

NewTagParamsWithDefaults instantiates a new TagParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyId

`func (o *TagParams) GetCompanyId() int32`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *TagParams) GetCompanyIdOk() (*int32, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *TagParams) SetCompanyId(v int32)`

SetCompanyId sets CompanyId field to given value.


### GetName

`func (o *TagParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TagParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TagParams) SetName(v string)`

SetName sets Name field to given value.


### GetShortcut1

`func (o *TagParams) GetShortcut1() string`

GetShortcut1 returns the Shortcut1 field if non-nil, zero value otherwise.

### GetShortcut1Ok

`func (o *TagParams) GetShortcut1Ok() (*string, bool)`

GetShortcut1Ok returns a tuple with the Shortcut1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortcut1

`func (o *TagParams) SetShortcut1(v string)`

SetShortcut1 sets Shortcut1 field to given value.

### HasShortcut1

`func (o *TagParams) HasShortcut1() bool`

HasShortcut1 returns a boolean if a field has been set.

### GetShortcut2

`func (o *TagParams) GetShortcut2() string`

GetShortcut2 returns the Shortcut2 field if non-nil, zero value otherwise.

### GetShortcut2Ok

`func (o *TagParams) GetShortcut2Ok() (*string, bool)`

GetShortcut2Ok returns a tuple with the Shortcut2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortcut2

`func (o *TagParams) SetShortcut2(v string)`

SetShortcut2 sets Shortcut2 field to given value.

### HasShortcut2

`func (o *TagParams) HasShortcut2() bool`

HasShortcut2 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


