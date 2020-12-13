# SelectablesIndexResponseAccountGroups

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | 決算書表示名（小カテゴリー）ID | 
**Name** | **string** | 決算書表示名 | 
**AccountStructureId** | **int32** | 年度ID | 
**AccountCategoryId** | **int32** | 勘定科目カテゴリーID | 
**DetailType** | Pointer to **int32** | 詳細パラメータの種類 | [optional] 
**Index** | **int32** | 並び順 | 
**CreatedAt** | Pointer to **string** | 作成日時 | [optional] 
**UpdatedAt** | Pointer to **string** | 更新日時 | [optional] 

## Methods

### NewSelectablesIndexResponseAccountGroups

`func NewSelectablesIndexResponseAccountGroups(id int32, name string, accountStructureId int32, accountCategoryId int32, index int32, ) *SelectablesIndexResponseAccountGroups`

NewSelectablesIndexResponseAccountGroups instantiates a new SelectablesIndexResponseAccountGroups object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSelectablesIndexResponseAccountGroupsWithDefaults

`func NewSelectablesIndexResponseAccountGroupsWithDefaults() *SelectablesIndexResponseAccountGroups`

NewSelectablesIndexResponseAccountGroupsWithDefaults instantiates a new SelectablesIndexResponseAccountGroups object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SelectablesIndexResponseAccountGroups) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SelectablesIndexResponseAccountGroups) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SelectablesIndexResponseAccountGroups) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *SelectablesIndexResponseAccountGroups) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SelectablesIndexResponseAccountGroups) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SelectablesIndexResponseAccountGroups) SetName(v string)`

SetName sets Name field to given value.


### GetAccountStructureId

`func (o *SelectablesIndexResponseAccountGroups) GetAccountStructureId() int32`

GetAccountStructureId returns the AccountStructureId field if non-nil, zero value otherwise.

### GetAccountStructureIdOk

`func (o *SelectablesIndexResponseAccountGroups) GetAccountStructureIdOk() (*int32, bool)`

GetAccountStructureIdOk returns a tuple with the AccountStructureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountStructureId

`func (o *SelectablesIndexResponseAccountGroups) SetAccountStructureId(v int32)`

SetAccountStructureId sets AccountStructureId field to given value.


### GetAccountCategoryId

`func (o *SelectablesIndexResponseAccountGroups) GetAccountCategoryId() int32`

GetAccountCategoryId returns the AccountCategoryId field if non-nil, zero value otherwise.

### GetAccountCategoryIdOk

`func (o *SelectablesIndexResponseAccountGroups) GetAccountCategoryIdOk() (*int32, bool)`

GetAccountCategoryIdOk returns a tuple with the AccountCategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountCategoryId

`func (o *SelectablesIndexResponseAccountGroups) SetAccountCategoryId(v int32)`

SetAccountCategoryId sets AccountCategoryId field to given value.


### GetDetailType

`func (o *SelectablesIndexResponseAccountGroups) GetDetailType() int32`

GetDetailType returns the DetailType field if non-nil, zero value otherwise.

### GetDetailTypeOk

`func (o *SelectablesIndexResponseAccountGroups) GetDetailTypeOk() (*int32, bool)`

GetDetailTypeOk returns a tuple with the DetailType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetailType

`func (o *SelectablesIndexResponseAccountGroups) SetDetailType(v int32)`

SetDetailType sets DetailType field to given value.

### HasDetailType

`func (o *SelectablesIndexResponseAccountGroups) HasDetailType() bool`

HasDetailType returns a boolean if a field has been set.

### GetIndex

`func (o *SelectablesIndexResponseAccountGroups) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *SelectablesIndexResponseAccountGroups) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *SelectablesIndexResponseAccountGroups) SetIndex(v int32)`

SetIndex sets Index field to given value.


### GetCreatedAt

`func (o *SelectablesIndexResponseAccountGroups) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SelectablesIndexResponseAccountGroups) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SelectablesIndexResponseAccountGroups) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SelectablesIndexResponseAccountGroups) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *SelectablesIndexResponseAccountGroups) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SelectablesIndexResponseAccountGroups) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SelectablesIndexResponseAccountGroups) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *SelectablesIndexResponseAccountGroups) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


