# BadRequestError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StatusCode** | Pointer to **int32** |  | [optional] 
**Errors** | Pointer to [**[]BadRequestErrorErrors**](BadRequestErrorErrors.md) |  | [optional] 

## Methods

### NewBadRequestError

`func NewBadRequestError() *BadRequestError`

NewBadRequestError instantiates a new BadRequestError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBadRequestErrorWithDefaults

`func NewBadRequestErrorWithDefaults() *BadRequestError`

NewBadRequestErrorWithDefaults instantiates a new BadRequestError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatusCode

`func (o *BadRequestError) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *BadRequestError) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *BadRequestError) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *BadRequestError) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.

### GetErrors

`func (o *BadRequestError) GetErrors() []BadRequestErrorErrors`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *BadRequestError) GetErrorsOk() (*[]BadRequestErrorErrors, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *BadRequestError) SetErrors(v []BadRequestErrorErrors)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *BadRequestError) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


