# UnauthorizedError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] 
**Messages** | Pointer to **string** |  | [optional] 

## Methods

### NewUnauthorizedError

`func NewUnauthorizedError() *UnauthorizedError`

NewUnauthorizedError instantiates a new UnauthorizedError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnauthorizedErrorWithDefaults

`func NewUnauthorizedErrorWithDefaults() *UnauthorizedError`

NewUnauthorizedErrorWithDefaults instantiates a new UnauthorizedError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *UnauthorizedError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *UnauthorizedError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *UnauthorizedError) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *UnauthorizedError) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetMessages

`func (o *UnauthorizedError) GetMessages() string`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *UnauthorizedError) GetMessagesOk() (*string, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *UnauthorizedError) SetMessages(v string)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *UnauthorizedError) HasMessages() bool`

HasMessages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


