# SegmentTagParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyId** | **int32** | 事業所ID | 
**Name** | **string** | セグメントタグ名 (30文字以内) | 
**Description** | Pointer to **string** | 備考 (30文字以内) | [optional] 
**Shortcut1** | Pointer to **string** | ショートカット１ (20文字以内) | [optional] 
**Shortcut2** | Pointer to **string** | ショートカット２ (20文字以内) | [optional] 

## Methods

### NewSegmentTagParams

`func NewSegmentTagParams(companyId int32, name string, ) *SegmentTagParams`

NewSegmentTagParams instantiates a new SegmentTagParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSegmentTagParamsWithDefaults

`func NewSegmentTagParamsWithDefaults() *SegmentTagParams`

NewSegmentTagParamsWithDefaults instantiates a new SegmentTagParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyId

`func (o *SegmentTagParams) GetCompanyId() int32`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *SegmentTagParams) GetCompanyIdOk() (*int32, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *SegmentTagParams) SetCompanyId(v int32)`

SetCompanyId sets CompanyId field to given value.


### GetName

`func (o *SegmentTagParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SegmentTagParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SegmentTagParams) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *SegmentTagParams) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SegmentTagParams) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SegmentTagParams) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SegmentTagParams) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetShortcut1

`func (o *SegmentTagParams) GetShortcut1() string`

GetShortcut1 returns the Shortcut1 field if non-nil, zero value otherwise.

### GetShortcut1Ok

`func (o *SegmentTagParams) GetShortcut1Ok() (*string, bool)`

GetShortcut1Ok returns a tuple with the Shortcut1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortcut1

`func (o *SegmentTagParams) SetShortcut1(v string)`

SetShortcut1 sets Shortcut1 field to given value.

### HasShortcut1

`func (o *SegmentTagParams) HasShortcut1() bool`

HasShortcut1 returns a boolean if a field has been set.

### GetShortcut2

`func (o *SegmentTagParams) GetShortcut2() string`

GetShortcut2 returns the Shortcut2 field if non-nil, zero value otherwise.

### GetShortcut2Ok

`func (o *SegmentTagParams) GetShortcut2Ok() (*string, bool)`

GetShortcut2Ok returns a tuple with the Shortcut2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortcut2

`func (o *SegmentTagParams) SetShortcut2(v string)`

SetShortcut2 sets Shortcut2 field to given value.

### HasShortcut2

`func (o *SegmentTagParams) HasShortcut2() bool`

HasShortcut2 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


