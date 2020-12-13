# DealCreateParams

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
**Details** | [**[]DealCreateParamsDetails**](DealCreateParamsDetails.md) |  | 
**Payments** | Pointer to [**[]DealCreateParamsPayments**](DealCreateParamsPayments.md) | 支払行一覧（配列）：未指定の場合、未決済の取引を作成します。 | [optional] 
**ReceiptIds** | Pointer to **[]int32** | 証憑ファイルID（配列） | [optional] 

## Methods

### NewDealCreateParams

`func NewDealCreateParams(issueDate string, type_ string, companyId int32, details []DealCreateParamsDetails, ) *DealCreateParams`

NewDealCreateParams instantiates a new DealCreateParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDealCreateParamsWithDefaults

`func NewDealCreateParamsWithDefaults() *DealCreateParams`

NewDealCreateParamsWithDefaults instantiates a new DealCreateParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIssueDate

`func (o *DealCreateParams) GetIssueDate() string`

GetIssueDate returns the IssueDate field if non-nil, zero value otherwise.

### GetIssueDateOk

`func (o *DealCreateParams) GetIssueDateOk() (*string, bool)`

GetIssueDateOk returns a tuple with the IssueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueDate

`func (o *DealCreateParams) SetIssueDate(v string)`

SetIssueDate sets IssueDate field to given value.


### GetType

`func (o *DealCreateParams) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DealCreateParams) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DealCreateParams) SetType(v string)`

SetType sets Type field to given value.


### GetCompanyId

`func (o *DealCreateParams) GetCompanyId() int32`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *DealCreateParams) GetCompanyIdOk() (*int32, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *DealCreateParams) SetCompanyId(v int32)`

SetCompanyId sets CompanyId field to given value.


### GetDueDate

`func (o *DealCreateParams) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *DealCreateParams) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *DealCreateParams) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *DealCreateParams) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### GetPartnerId

`func (o *DealCreateParams) GetPartnerId() int32`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *DealCreateParams) GetPartnerIdOk() (*int32, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *DealCreateParams) SetPartnerId(v int32)`

SetPartnerId sets PartnerId field to given value.

### HasPartnerId

`func (o *DealCreateParams) HasPartnerId() bool`

HasPartnerId returns a boolean if a field has been set.

### GetPartnerCode

`func (o *DealCreateParams) GetPartnerCode() string`

GetPartnerCode returns the PartnerCode field if non-nil, zero value otherwise.

### GetPartnerCodeOk

`func (o *DealCreateParams) GetPartnerCodeOk() (*string, bool)`

GetPartnerCodeOk returns a tuple with the PartnerCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerCode

`func (o *DealCreateParams) SetPartnerCode(v string)`

SetPartnerCode sets PartnerCode field to given value.

### HasPartnerCode

`func (o *DealCreateParams) HasPartnerCode() bool`

HasPartnerCode returns a boolean if a field has been set.

### GetRefNumber

`func (o *DealCreateParams) GetRefNumber() string`

GetRefNumber returns the RefNumber field if non-nil, zero value otherwise.

### GetRefNumberOk

`func (o *DealCreateParams) GetRefNumberOk() (*string, bool)`

GetRefNumberOk returns a tuple with the RefNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefNumber

`func (o *DealCreateParams) SetRefNumber(v string)`

SetRefNumber sets RefNumber field to given value.

### HasRefNumber

`func (o *DealCreateParams) HasRefNumber() bool`

HasRefNumber returns a boolean if a field has been set.

### GetDetails

`func (o *DealCreateParams) GetDetails() []DealCreateParamsDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *DealCreateParams) GetDetailsOk() (*[]DealCreateParamsDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *DealCreateParams) SetDetails(v []DealCreateParamsDetails)`

SetDetails sets Details field to given value.


### GetPayments

`func (o *DealCreateParams) GetPayments() []DealCreateParamsPayments`

GetPayments returns the Payments field if non-nil, zero value otherwise.

### GetPaymentsOk

`func (o *DealCreateParams) GetPaymentsOk() (*[]DealCreateParamsPayments, bool)`

GetPaymentsOk returns a tuple with the Payments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayments

`func (o *DealCreateParams) SetPayments(v []DealCreateParamsPayments)`

SetPayments sets Payments field to given value.

### HasPayments

`func (o *DealCreateParams) HasPayments() bool`

HasPayments returns a boolean if a field has been set.

### GetReceiptIds

`func (o *DealCreateParams) GetReceiptIds() []int32`

GetReceiptIds returns the ReceiptIds field if non-nil, zero value otherwise.

### GetReceiptIdsOk

`func (o *DealCreateParams) GetReceiptIdsOk() (*[]int32, bool)`

GetReceiptIdsOk returns a tuple with the ReceiptIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiptIds

`func (o *DealCreateParams) SetReceiptIds(v []int32)`

SetReceiptIds sets ReceiptIds field to given value.

### HasReceiptIds

`func (o *DealCreateParams) HasReceiptIds() bool`

HasReceiptIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


