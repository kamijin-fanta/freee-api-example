# SectionParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyId** | **int32** | 事業所ID | 
**Name** | **string** | 部門名 (30文字以内) | 
**LongName** | Pointer to **string** | 正式名称 (255文字以内) | [optional] 
**Shortcut1** | Pointer to **string** | ショートカット１ (20文字以内) | [optional] 
**Shortcut2** | Pointer to **string** | ショートカット２ (20文字以内) | [optional] 
**ParentId** | Pointer to **int32** | 親部門ID (ビジネスプラン以上) | [optional] 

## Methods

### NewSectionParams

`func NewSectionParams(companyId int32, name string, ) *SectionParams`

NewSectionParams instantiates a new SectionParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSectionParamsWithDefaults

`func NewSectionParamsWithDefaults() *SectionParams`

NewSectionParamsWithDefaults instantiates a new SectionParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyId

`func (o *SectionParams) GetCompanyId() int32`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *SectionParams) GetCompanyIdOk() (*int32, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *SectionParams) SetCompanyId(v int32)`

SetCompanyId sets CompanyId field to given value.


### GetName

`func (o *SectionParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SectionParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SectionParams) SetName(v string)`

SetName sets Name field to given value.


### GetLongName

`func (o *SectionParams) GetLongName() string`

GetLongName returns the LongName field if non-nil, zero value otherwise.

### GetLongNameOk

`func (o *SectionParams) GetLongNameOk() (*string, bool)`

GetLongNameOk returns a tuple with the LongName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongName

`func (o *SectionParams) SetLongName(v string)`

SetLongName sets LongName field to given value.

### HasLongName

`func (o *SectionParams) HasLongName() bool`

HasLongName returns a boolean if a field has been set.

### GetShortcut1

`func (o *SectionParams) GetShortcut1() string`

GetShortcut1 returns the Shortcut1 field if non-nil, zero value otherwise.

### GetShortcut1Ok

`func (o *SectionParams) GetShortcut1Ok() (*string, bool)`

GetShortcut1Ok returns a tuple with the Shortcut1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortcut1

`func (o *SectionParams) SetShortcut1(v string)`

SetShortcut1 sets Shortcut1 field to given value.

### HasShortcut1

`func (o *SectionParams) HasShortcut1() bool`

HasShortcut1 returns a boolean if a field has been set.

### GetShortcut2

`func (o *SectionParams) GetShortcut2() string`

GetShortcut2 returns the Shortcut2 field if non-nil, zero value otherwise.

### GetShortcut2Ok

`func (o *SectionParams) GetShortcut2Ok() (*string, bool)`

GetShortcut2Ok returns a tuple with the Shortcut2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortcut2

`func (o *SectionParams) SetShortcut2(v string)`

SetShortcut2 sets Shortcut2 field to given value.

### HasShortcut2

`func (o *SectionParams) HasShortcut2() bool`

HasShortcut2 returns a boolean if a field has been set.

### GetParentId

`func (o *SectionParams) GetParentId() int32`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *SectionParams) GetParentIdOk() (*int32, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *SectionParams) SetParentId(v int32)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *SectionParams) HasParentId() bool`

HasParentId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


