# SectionResponseSection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | 部門ID | 
**Name** | **string** | 部門名 (30文字以内) | 
**LongName** | Pointer to **NullableString** | 正式名称（255文字以内） | [optional] 
**CompanyId** | **int32** | 事業所ID | 
**Shortcut1** | Pointer to **NullableString** | ショートカット１ (20文字以内) | [optional] 
**Shortcut2** | Pointer to **NullableString** | ショートカット２ (20文字以内) | [optional] 

## Methods

### NewSectionResponseSection

`func NewSectionResponseSection(id int32, name string, companyId int32, ) *SectionResponseSection`

NewSectionResponseSection instantiates a new SectionResponseSection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSectionResponseSectionWithDefaults

`func NewSectionResponseSectionWithDefaults() *SectionResponseSection`

NewSectionResponseSectionWithDefaults instantiates a new SectionResponseSection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SectionResponseSection) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SectionResponseSection) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SectionResponseSection) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *SectionResponseSection) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SectionResponseSection) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SectionResponseSection) SetName(v string)`

SetName sets Name field to given value.


### GetLongName

`func (o *SectionResponseSection) GetLongName() string`

GetLongName returns the LongName field if non-nil, zero value otherwise.

### GetLongNameOk

`func (o *SectionResponseSection) GetLongNameOk() (*string, bool)`

GetLongNameOk returns a tuple with the LongName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongName

`func (o *SectionResponseSection) SetLongName(v string)`

SetLongName sets LongName field to given value.

### HasLongName

`func (o *SectionResponseSection) HasLongName() bool`

HasLongName returns a boolean if a field has been set.

### SetLongNameNil

`func (o *SectionResponseSection) SetLongNameNil(b bool)`

 SetLongNameNil sets the value for LongName to be an explicit nil

### UnsetLongName
`func (o *SectionResponseSection) UnsetLongName()`

UnsetLongName ensures that no value is present for LongName, not even an explicit nil
### GetCompanyId

`func (o *SectionResponseSection) GetCompanyId() int32`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *SectionResponseSection) GetCompanyIdOk() (*int32, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *SectionResponseSection) SetCompanyId(v int32)`

SetCompanyId sets CompanyId field to given value.


### GetShortcut1

`func (o *SectionResponseSection) GetShortcut1() string`

GetShortcut1 returns the Shortcut1 field if non-nil, zero value otherwise.

### GetShortcut1Ok

`func (o *SectionResponseSection) GetShortcut1Ok() (*string, bool)`

GetShortcut1Ok returns a tuple with the Shortcut1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortcut1

`func (o *SectionResponseSection) SetShortcut1(v string)`

SetShortcut1 sets Shortcut1 field to given value.

### HasShortcut1

`func (o *SectionResponseSection) HasShortcut1() bool`

HasShortcut1 returns a boolean if a field has been set.

### SetShortcut1Nil

`func (o *SectionResponseSection) SetShortcut1Nil(b bool)`

 SetShortcut1Nil sets the value for Shortcut1 to be an explicit nil

### UnsetShortcut1
`func (o *SectionResponseSection) UnsetShortcut1()`

UnsetShortcut1 ensures that no value is present for Shortcut1, not even an explicit nil
### GetShortcut2

`func (o *SectionResponseSection) GetShortcut2() string`

GetShortcut2 returns the Shortcut2 field if non-nil, zero value otherwise.

### GetShortcut2Ok

`func (o *SectionResponseSection) GetShortcut2Ok() (*string, bool)`

GetShortcut2Ok returns a tuple with the Shortcut2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortcut2

`func (o *SectionResponseSection) SetShortcut2(v string)`

SetShortcut2 sets Shortcut2 field to given value.

### HasShortcut2

`func (o *SectionResponseSection) HasShortcut2() bool`

HasShortcut2 returns a boolean if a field has been set.

### SetShortcut2Nil

`func (o *SectionResponseSection) SetShortcut2Nil(b bool)`

 SetShortcut2Nil sets the value for Shortcut2 to be an explicit nil

### UnsetShortcut2
`func (o *SectionResponseSection) UnsetShortcut2()`

UnsetShortcut2 ensures that no value is present for Shortcut2, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


