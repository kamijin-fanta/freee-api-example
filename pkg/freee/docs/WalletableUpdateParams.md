# WalletableUpdateParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | 口座名 (255文字以内) | 
**CompanyId** | **int32** | 事業所ID | 

## Methods

### NewWalletableUpdateParams

`func NewWalletableUpdateParams(name string, companyId int32, ) *WalletableUpdateParams`

NewWalletableUpdateParams instantiates a new WalletableUpdateParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWalletableUpdateParamsWithDefaults

`func NewWalletableUpdateParamsWithDefaults() *WalletableUpdateParams`

NewWalletableUpdateParamsWithDefaults instantiates a new WalletableUpdateParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *WalletableUpdateParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WalletableUpdateParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WalletableUpdateParams) SetName(v string)`

SetName sets Name field to given value.


### GetCompanyId

`func (o *WalletableUpdateParams) GetCompanyId() int32`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *WalletableUpdateParams) GetCompanyIdOk() (*int32, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *WalletableUpdateParams) SetCompanyId(v int32)`

SetCompanyId sets CompanyId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


