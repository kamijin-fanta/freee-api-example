# TagResponseTag

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | タグID | 
**CompanyId** | **int32** | 事業所ID | 
**Name** | **NullableString** | 名前(30文字以内) | 
**Shortcut1** | Pointer to **NullableString** | ショートカット1 (255文字以内) | [optional] 
**Shortcut2** | Pointer to **NullableString** | ショートカット2 (255文字以内) | [optional] 

## Methods

### NewTagResponseTag

`func NewTagResponseTag(id int32, companyId int32, name NullableString, ) *TagResponseTag`

NewTagResponseTag instantiates a new TagResponseTag object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagResponseTagWithDefaults

`func NewTagResponseTagWithDefaults() *TagResponseTag`

NewTagResponseTagWithDefaults instantiates a new TagResponseTag object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TagResponseTag) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TagResponseTag) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TagResponseTag) SetId(v int32)`

SetId sets Id field to given value.


### GetCompanyId

`func (o *TagResponseTag) GetCompanyId() int32`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *TagResponseTag) GetCompanyIdOk() (*int32, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *TagResponseTag) SetCompanyId(v int32)`

SetCompanyId sets CompanyId field to given value.


### GetName

`func (o *TagResponseTag) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TagResponseTag) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TagResponseTag) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *TagResponseTag) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *TagResponseTag) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetShortcut1

`func (o *TagResponseTag) GetShortcut1() string`

GetShortcut1 returns the Shortcut1 field if non-nil, zero value otherwise.

### GetShortcut1Ok

`func (o *TagResponseTag) GetShortcut1Ok() (*string, bool)`

GetShortcut1Ok returns a tuple with the Shortcut1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortcut1

`func (o *TagResponseTag) SetShortcut1(v string)`

SetShortcut1 sets Shortcut1 field to given value.

### HasShortcut1

`func (o *TagResponseTag) HasShortcut1() bool`

HasShortcut1 returns a boolean if a field has been set.

### SetShortcut1Nil

`func (o *TagResponseTag) SetShortcut1Nil(b bool)`

 SetShortcut1Nil sets the value for Shortcut1 to be an explicit nil

### UnsetShortcut1
`func (o *TagResponseTag) UnsetShortcut1()`

UnsetShortcut1 ensures that no value is present for Shortcut1, not even an explicit nil
### GetShortcut2

`func (o *TagResponseTag) GetShortcut2() string`

GetShortcut2 returns the Shortcut2 field if non-nil, zero value otherwise.

### GetShortcut2Ok

`func (o *TagResponseTag) GetShortcut2Ok() (*string, bool)`

GetShortcut2Ok returns a tuple with the Shortcut2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortcut2

`func (o *TagResponseTag) SetShortcut2(v string)`

SetShortcut2 sets Shortcut2 field to given value.

### HasShortcut2

`func (o *TagResponseTag) HasShortcut2() bool`

HasShortcut2 returns a boolean if a field has been set.

### SetShortcut2Nil

`func (o *TagResponseTag) SetShortcut2Nil(b bool)`

 SetShortcut2Nil sets the value for Shortcut2 to be an explicit nil

### UnsetShortcut2
`func (o *TagResponseTag) UnsetShortcut2()`

UnsetShortcut2 ensures that no value is present for Shortcut2, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


