# ReceiptUpdateParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyId** | **int32** | 事業所ID | 
**Description** | Pointer to **string** | メモ (255文字以内) | [optional] 
**IssueDate** | **string** | 取引日 (yyyy-mm-dd) | 

## Methods

### NewReceiptUpdateParams

`func NewReceiptUpdateParams(companyId int32, issueDate string, ) *ReceiptUpdateParams`

NewReceiptUpdateParams instantiates a new ReceiptUpdateParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReceiptUpdateParamsWithDefaults

`func NewReceiptUpdateParamsWithDefaults() *ReceiptUpdateParams`

NewReceiptUpdateParamsWithDefaults instantiates a new ReceiptUpdateParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyId

`func (o *ReceiptUpdateParams) GetCompanyId() int32`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *ReceiptUpdateParams) GetCompanyIdOk() (*int32, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *ReceiptUpdateParams) SetCompanyId(v int32)`

SetCompanyId sets CompanyId field to given value.


### GetDescription

`func (o *ReceiptUpdateParams) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ReceiptUpdateParams) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ReceiptUpdateParams) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ReceiptUpdateParams) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIssueDate

`func (o *ReceiptUpdateParams) GetIssueDate() string`

GetIssueDate returns the IssueDate field if non-nil, zero value otherwise.

### GetIssueDateOk

`func (o *ReceiptUpdateParams) GetIssueDateOk() (*string, bool)`

GetIssueDateOk returns a tuple with the IssueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueDate

`func (o *ReceiptUpdateParams) SetIssueDate(v string)`

SetIssueDate sets IssueDate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


