# ReceiptCreateParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyId** | **int32** | 事業所ID | 
**Description** | Pointer to **string** | メモ (255文字以内) | [optional] 
**IssueDate** | Pointer to **string** | 取引日 (yyyy-mm-dd) | [optional] 
**Receipt** | ***os.File** | 証憑ファイル | 

## Methods

### NewReceiptCreateParams

`func NewReceiptCreateParams(companyId int32, receipt *os.File, ) *ReceiptCreateParams`

NewReceiptCreateParams instantiates a new ReceiptCreateParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReceiptCreateParamsWithDefaults

`func NewReceiptCreateParamsWithDefaults() *ReceiptCreateParams`

NewReceiptCreateParamsWithDefaults instantiates a new ReceiptCreateParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyId

`func (o *ReceiptCreateParams) GetCompanyId() int32`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *ReceiptCreateParams) GetCompanyIdOk() (*int32, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *ReceiptCreateParams) SetCompanyId(v int32)`

SetCompanyId sets CompanyId field to given value.


### GetDescription

`func (o *ReceiptCreateParams) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ReceiptCreateParams) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ReceiptCreateParams) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ReceiptCreateParams) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIssueDate

`func (o *ReceiptCreateParams) GetIssueDate() string`

GetIssueDate returns the IssueDate field if non-nil, zero value otherwise.

### GetIssueDateOk

`func (o *ReceiptCreateParams) GetIssueDateOk() (*string, bool)`

GetIssueDateOk returns a tuple with the IssueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueDate

`func (o *ReceiptCreateParams) SetIssueDate(v string)`

SetIssueDate sets IssueDate field to given value.

### HasIssueDate

`func (o *ReceiptCreateParams) HasIssueDate() bool`

HasIssueDate returns a boolean if a field has been set.

### GetReceipt

`func (o *ReceiptCreateParams) GetReceipt() *os.File`

GetReceipt returns the Receipt field if non-nil, zero value otherwise.

### GetReceiptOk

`func (o *ReceiptCreateParams) GetReceiptOk() (**os.File, bool)`

GetReceiptOk returns a tuple with the Receipt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceipt

`func (o *ReceiptCreateParams) SetReceipt(v *os.File)`

SetReceipt sets Receipt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


