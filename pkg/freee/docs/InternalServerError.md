# InternalServerError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StatusCode** | Pointer to **int32** |  | [optional] 
**Errors** | Pointer to [**[]InternalServerErrorErrors**](InternalServerErrorErrors.md) |  | [optional] 

## Methods

### NewInternalServerError

`func NewInternalServerError() *InternalServerError`

NewInternalServerError instantiates a new InternalServerError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternalServerErrorWithDefaults

`func NewInternalServerErrorWithDefaults() *InternalServerError`

NewInternalServerErrorWithDefaults instantiates a new InternalServerError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatusCode

`func (o *InternalServerError) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *InternalServerError) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *InternalServerError) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *InternalServerError) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.

### GetErrors

`func (o *InternalServerError) GetErrors() []InternalServerErrorErrors`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *InternalServerError) GetErrorsOk() (*[]InternalServerErrorErrors, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *InternalServerError) SetErrors(v []InternalServerErrorErrors)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *InternalServerError) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


