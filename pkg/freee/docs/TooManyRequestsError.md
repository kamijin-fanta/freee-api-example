# TooManyRequestsError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StatusCode** | **int32** |  | 
**Meta** | [**TooManyRequestsErrorMeta**](tooManyRequestsError_meta.md) |  | 

## Methods

### NewTooManyRequestsError

`func NewTooManyRequestsError(statusCode int32, meta TooManyRequestsErrorMeta, ) *TooManyRequestsError`

NewTooManyRequestsError instantiates a new TooManyRequestsError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTooManyRequestsErrorWithDefaults

`func NewTooManyRequestsErrorWithDefaults() *TooManyRequestsError`

NewTooManyRequestsErrorWithDefaults instantiates a new TooManyRequestsError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatusCode

`func (o *TooManyRequestsError) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *TooManyRequestsError) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *TooManyRequestsError) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.


### GetMeta

`func (o *TooManyRequestsError) GetMeta() TooManyRequestsErrorMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *TooManyRequestsError) GetMetaOk() (*TooManyRequestsErrorMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *TooManyRequestsError) SetMeta(v TooManyRequestsErrorMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


