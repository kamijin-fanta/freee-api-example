# DealUpdateParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IssueDate** | **string** | 発生日 (yyyy-mm-dd) | 
**Type** | **string** | 収支区分 (収入: income, 支出: expense) | 
**CompanyId** | **int32** | 事業所ID | 
**DueDate** | Pointer to **string** | 支払期日(yyyy-mm-dd) | [optional] 
**PartnerId** | Pointer to **int32** | 取引先ID | [optional] 
**PartnerCode** | Pointer to **string** | 取引先コード | [optional] 
**RefNumber** | Pointer to **string** | 管理番号 | [optional] 
**Details** | [**[]DealUpdateParamsDetails**](DealUpdateParamsDetails.md) |  | 
**ReceiptIds** | Pointer to **[]int32** | 証憑ファイルID（配列） | [optional] 

## Methods

### NewDealUpdateParams

`func NewDealUpdateParams(issueDate string, type_ string, companyId int32, details []DealUpdateParamsDetails, ) *DealUpdateParams`

NewDealUpdateParams instantiates a new DealUpdateParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDealUpdateParamsWithDefaults

`func NewDealUpdateParamsWithDefaults() *DealUpdateParams`

NewDealUpdateParamsWithDefaults instantiates a new DealUpdateParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIssueDate

`func (o *DealUpdateParams) GetIssueDate() string`

GetIssueDate returns the IssueDate field if non-nil, zero value otherwise.

### GetIssueDateOk

`func (o *DealUpdateParams) GetIssueDateOk() (*string, bool)`

GetIssueDateOk returns a tuple with the IssueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueDate

`func (o *DealUpdateParams) SetIssueDate(v string)`

SetIssueDate sets IssueDate field to given value.


### GetType

`func (o *DealUpdateParams) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DealUpdateParams) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DealUpdateParams) SetType(v string)`

SetType sets Type field to given value.


### GetCompanyId

`func (o *DealUpdateParams) GetCompanyId() int32`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *DealUpdateParams) GetCompanyIdOk() (*int32, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *DealUpdateParams) SetCompanyId(v int32)`

SetCompanyId sets CompanyId field to given value.


### GetDueDate

`func (o *DealUpdateParams) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *DealUpdateParams) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *DealUpdateParams) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *DealUpdateParams) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### GetPartnerId

`func (o *DealUpdateParams) GetPartnerId() int32`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *DealUpdateParams) GetPartnerIdOk() (*int32, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *DealUpdateParams) SetPartnerId(v int32)`

SetPartnerId sets PartnerId field to given value.

### HasPartnerId

`func (o *DealUpdateParams) HasPartnerId() bool`

HasPartnerId returns a boolean if a field has been set.

### GetPartnerCode

`func (o *DealUpdateParams) GetPartnerCode() string`

GetPartnerCode returns the PartnerCode field if non-nil, zero value otherwise.

### GetPartnerCodeOk

`func (o *DealUpdateParams) GetPartnerCodeOk() (*string, bool)`

GetPartnerCodeOk returns a tuple with the PartnerCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerCode

`func (o *DealUpdateParams) SetPartnerCode(v string)`

SetPartnerCode sets PartnerCode field to given value.

### HasPartnerCode

`func (o *DealUpdateParams) HasPartnerCode() bool`

HasPartnerCode returns a boolean if a field has been set.

### GetRefNumber

`func (o *DealUpdateParams) GetRefNumber() string`

GetRefNumber returns the RefNumber field if non-nil, zero value otherwise.

### GetRefNumberOk

`func (o *DealUpdateParams) GetRefNumberOk() (*string, bool)`

GetRefNumberOk returns a tuple with the RefNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefNumber

`func (o *DealUpdateParams) SetRefNumber(v string)`

SetRefNumber sets RefNumber field to given value.

### HasRefNumber

`func (o *DealUpdateParams) HasRefNumber() bool`

HasRefNumber returns a boolean if a field has been set.

### GetDetails

`func (o *DealUpdateParams) GetDetails() []DealUpdateParamsDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *DealUpdateParams) GetDetailsOk() (*[]DealUpdateParamsDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *DealUpdateParams) SetDetails(v []DealUpdateParamsDetails)`

SetDetails sets Details field to given value.


### GetReceiptIds

`func (o *DealUpdateParams) GetReceiptIds() []int32`

GetReceiptIds returns the ReceiptIds field if non-nil, zero value otherwise.

### GetReceiptIdsOk

`func (o *DealUpdateParams) GetReceiptIdsOk() (*[]int32, bool)`

GetReceiptIdsOk returns a tuple with the ReceiptIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiptIds

`func (o *DealUpdateParams) SetReceiptIds(v []int32)`

SetReceiptIds sets ReceiptIds field to given value.

### HasReceiptIds

`func (o *DealUpdateParams) HasReceiptIds() bool`

HasReceiptIds returns a boolean if a field has been set.

### SetReceiptIdsNil

`func (o *DealUpdateParams) SetReceiptIdsNil(b bool)`

 SetReceiptIdsNil sets the value for ReceiptIds to be an explicit nil

### UnsetReceiptIds
`func (o *DealUpdateParams) UnsetReceiptIds()`

UnsetReceiptIds ensures that no value is present for ReceiptIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


