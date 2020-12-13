# AccountItemParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyId** | **int32** | 事業所ID | 
**AccountItem** | [**AccountItemParamsAccountItem**](accountItemParams_account_item.md) |  | 

## Methods

### NewAccountItemParams

`func NewAccountItemParams(companyId int32, accountItem AccountItemParamsAccountItem, ) *AccountItemParams`

NewAccountItemParams instantiates a new AccountItemParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountItemParamsWithDefaults

`func NewAccountItemParamsWithDefaults() *AccountItemParams`

NewAccountItemParamsWithDefaults instantiates a new AccountItemParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyId

`func (o *AccountItemParams) GetCompanyId() int32`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *AccountItemParams) GetCompanyIdOk() (*int32, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *AccountItemParams) SetCompanyId(v int32)`

SetCompanyId sets CompanyId field to given value.


### GetAccountItem

`func (o *AccountItemParams) GetAccountItem() AccountItemParamsAccountItem`

GetAccountItem returns the AccountItem field if non-nil, zero value otherwise.

### GetAccountItemOk

`func (o *AccountItemParams) GetAccountItemOk() (*AccountItemParamsAccountItem, bool)`

GetAccountItemOk returns a tuple with the AccountItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountItem

`func (o *AccountItemParams) SetAccountItem(v AccountItemParamsAccountItem)`

SetAccountItem sets AccountItem field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


