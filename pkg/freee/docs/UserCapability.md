# UserCapability

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Read** | Pointer to **bool** | 閲覧 | [optional] 
**Create** | Pointer to **bool** | 作成 | [optional] 
**Update** | Pointer to **bool** | 更新 | [optional] 
**Destroy** | Pointer to **bool** | 削除 | [optional] 

## Methods

### NewUserCapability

`func NewUserCapability() *UserCapability`

NewUserCapability instantiates a new UserCapability object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserCapabilityWithDefaults

`func NewUserCapabilityWithDefaults() *UserCapability`

NewUserCapabilityWithDefaults instantiates a new UserCapability object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRead

`func (o *UserCapability) GetRead() bool`

GetRead returns the Read field if non-nil, zero value otherwise.

### GetReadOk

`func (o *UserCapability) GetReadOk() (*bool, bool)`

GetReadOk returns a tuple with the Read field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRead

`func (o *UserCapability) SetRead(v bool)`

SetRead sets Read field to given value.

### HasRead

`func (o *UserCapability) HasRead() bool`

HasRead returns a boolean if a field has been set.

### GetCreate

`func (o *UserCapability) GetCreate() bool`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *UserCapability) GetCreateOk() (*bool, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *UserCapability) SetCreate(v bool)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *UserCapability) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetUpdate

`func (o *UserCapability) GetUpdate() bool`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *UserCapability) GetUpdateOk() (*bool, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *UserCapability) SetUpdate(v bool)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *UserCapability) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetDestroy

`func (o *UserCapability) GetDestroy() bool`

GetDestroy returns the Destroy field if non-nil, zero value otherwise.

### GetDestroyOk

`func (o *UserCapability) GetDestroyOk() (*bool, bool)`

GetDestroyOk returns a tuple with the Destroy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestroy

`func (o *UserCapability) SetDestroy(v bool)`

SetDestroy sets Destroy field to given value.

### HasDestroy

`func (o *UserCapability) HasDestroy() bool`

HasDestroy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


