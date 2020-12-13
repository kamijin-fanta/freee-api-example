# DealResponseDealUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | ユーザーID | 
**Email** | **string** | メールアドレス | 
**DisplayName** | Pointer to **NullableString** | 表示名 | [optional] 

## Methods

### NewDealResponseDealUser

`func NewDealResponseDealUser(id int32, email string, ) *DealResponseDealUser`

NewDealResponseDealUser instantiates a new DealResponseDealUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDealResponseDealUserWithDefaults

`func NewDealResponseDealUserWithDefaults() *DealResponseDealUser`

NewDealResponseDealUserWithDefaults instantiates a new DealResponseDealUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DealResponseDealUser) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DealResponseDealUser) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DealResponseDealUser) SetId(v int32)`

SetId sets Id field to given value.


### GetEmail

`func (o *DealResponseDealUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *DealResponseDealUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *DealResponseDealUser) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetDisplayName

`func (o *DealResponseDealUser) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *DealResponseDealUser) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *DealResponseDealUser) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *DealResponseDealUser) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *DealResponseDealUser) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *DealResponseDealUser) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


