# DealResponseDeal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | 取引ID | 
**CompanyId** | **int32** | 事業所ID | 
**IssueDate** | **string** | 発生日 (yyyy-mm-dd) | 
**DueDate** | Pointer to **string** | 支払期日 (yyyy-mm-dd) | [optional] 
**Amount** | **int32** | 金額 | 
**DueAmount** | Pointer to **int32** | 支払金額 | [optional] 
**Type** | Pointer to **string** | 収支区分 (収入: income, 支出: expense) | [optional] 
**PartnerId** | **int32** | 取引先ID | 
**PartnerCode** | Pointer to **NullableString** | 取引先コード | [optional] 
**RefNumber** | Pointer to **string** | 管理番号 | [optional] 
**Status** | **string** | 決済状況 (未決済: unsettled, 完了: settled) | 
**Details** | Pointer to [**[]DealCreateResponseDealDetails**](DealCreateResponseDealDetails.md) | 取引の明細行 | [optional] 
**Renews** | Pointer to [**[]DealResponseDealRenews**](DealResponseDealRenews.md) | 取引の+更新行 | [optional] 
**Payments** | Pointer to [**[]DealCreateResponseDealPayments**](DealCreateResponseDealPayments.md) | 取引の支払行 | [optional] 
**Receipts** | Pointer to [**[]DealResponseDealReceipts**](DealResponseDealReceipts.md) | 証憑ファイル | [optional] 

## Methods

### NewDealResponseDeal

`func NewDealResponseDeal(id int32, companyId int32, issueDate string, amount int32, partnerId int32, status string, ) *DealResponseDeal`

NewDealResponseDeal instantiates a new DealResponseDeal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDealResponseDealWithDefaults

`func NewDealResponseDealWithDefaults() *DealResponseDeal`

NewDealResponseDealWithDefaults instantiates a new DealResponseDeal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DealResponseDeal) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DealResponseDeal) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DealResponseDeal) SetId(v int32)`

SetId sets Id field to given value.


### GetCompanyId

`func (o *DealResponseDeal) GetCompanyId() int32`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *DealResponseDeal) GetCompanyIdOk() (*int32, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *DealResponseDeal) SetCompanyId(v int32)`

SetCompanyId sets CompanyId field to given value.


### GetIssueDate

`func (o *DealResponseDeal) GetIssueDate() string`

GetIssueDate returns the IssueDate field if non-nil, zero value otherwise.

### GetIssueDateOk

`func (o *DealResponseDeal) GetIssueDateOk() (*string, bool)`

GetIssueDateOk returns a tuple with the IssueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueDate

`func (o *DealResponseDeal) SetIssueDate(v string)`

SetIssueDate sets IssueDate field to given value.


### GetDueDate

`func (o *DealResponseDeal) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *DealResponseDeal) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *DealResponseDeal) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *DealResponseDeal) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### GetAmount

`func (o *DealResponseDeal) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *DealResponseDeal) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *DealResponseDeal) SetAmount(v int32)`

SetAmount sets Amount field to given value.


### GetDueAmount

`func (o *DealResponseDeal) GetDueAmount() int32`

GetDueAmount returns the DueAmount field if non-nil, zero value otherwise.

### GetDueAmountOk

`func (o *DealResponseDeal) GetDueAmountOk() (*int32, bool)`

GetDueAmountOk returns a tuple with the DueAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueAmount

`func (o *DealResponseDeal) SetDueAmount(v int32)`

SetDueAmount sets DueAmount field to given value.

### HasDueAmount

`func (o *DealResponseDeal) HasDueAmount() bool`

HasDueAmount returns a boolean if a field has been set.

### GetType

`func (o *DealResponseDeal) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DealResponseDeal) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DealResponseDeal) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DealResponseDeal) HasType() bool`

HasType returns a boolean if a field has been set.

### GetPartnerId

`func (o *DealResponseDeal) GetPartnerId() int32`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *DealResponseDeal) GetPartnerIdOk() (*int32, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *DealResponseDeal) SetPartnerId(v int32)`

SetPartnerId sets PartnerId field to given value.


### GetPartnerCode

`func (o *DealResponseDeal) GetPartnerCode() string`

GetPartnerCode returns the PartnerCode field if non-nil, zero value otherwise.

### GetPartnerCodeOk

`func (o *DealResponseDeal) GetPartnerCodeOk() (*string, bool)`

GetPartnerCodeOk returns a tuple with the PartnerCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerCode

`func (o *DealResponseDeal) SetPartnerCode(v string)`

SetPartnerCode sets PartnerCode field to given value.

### HasPartnerCode

`func (o *DealResponseDeal) HasPartnerCode() bool`

HasPartnerCode returns a boolean if a field has been set.

### SetPartnerCodeNil

`func (o *DealResponseDeal) SetPartnerCodeNil(b bool)`

 SetPartnerCodeNil sets the value for PartnerCode to be an explicit nil

### UnsetPartnerCode
`func (o *DealResponseDeal) UnsetPartnerCode()`

UnsetPartnerCode ensures that no value is present for PartnerCode, not even an explicit nil
### GetRefNumber

`func (o *DealResponseDeal) GetRefNumber() string`

GetRefNumber returns the RefNumber field if non-nil, zero value otherwise.

### GetRefNumberOk

`func (o *DealResponseDeal) GetRefNumberOk() (*string, bool)`

GetRefNumberOk returns a tuple with the RefNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefNumber

`func (o *DealResponseDeal) SetRefNumber(v string)`

SetRefNumber sets RefNumber field to given value.

### HasRefNumber

`func (o *DealResponseDeal) HasRefNumber() bool`

HasRefNumber returns a boolean if a field has been set.

### GetStatus

`func (o *DealResponseDeal) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DealResponseDeal) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DealResponseDeal) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetDetails

`func (o *DealResponseDeal) GetDetails() []DealCreateResponseDealDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *DealResponseDeal) GetDetailsOk() (*[]DealCreateResponseDealDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *DealResponseDeal) SetDetails(v []DealCreateResponseDealDetails)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *DealResponseDeal) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetRenews

`func (o *DealResponseDeal) GetRenews() []DealResponseDealRenews`

GetRenews returns the Renews field if non-nil, zero value otherwise.

### GetRenewsOk

`func (o *DealResponseDeal) GetRenewsOk() (*[]DealResponseDealRenews, bool)`

GetRenewsOk returns a tuple with the Renews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenews

`func (o *DealResponseDeal) SetRenews(v []DealResponseDealRenews)`

SetRenews sets Renews field to given value.

### HasRenews

`func (o *DealResponseDeal) HasRenews() bool`

HasRenews returns a boolean if a field has been set.

### GetPayments

`func (o *DealResponseDeal) GetPayments() []DealCreateResponseDealPayments`

GetPayments returns the Payments field if non-nil, zero value otherwise.

### GetPaymentsOk

`func (o *DealResponseDeal) GetPaymentsOk() (*[]DealCreateResponseDealPayments, bool)`

GetPaymentsOk returns a tuple with the Payments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayments

`func (o *DealResponseDeal) SetPayments(v []DealCreateResponseDealPayments)`

SetPayments sets Payments field to given value.

### HasPayments

`func (o *DealResponseDeal) HasPayments() bool`

HasPayments returns a boolean if a field has been set.

### GetReceipts

`func (o *DealResponseDeal) GetReceipts() []DealResponseDealReceipts`

GetReceipts returns the Receipts field if non-nil, zero value otherwise.

### GetReceiptsOk

`func (o *DealResponseDeal) GetReceiptsOk() (*[]DealResponseDealReceipts, bool)`

GetReceiptsOk returns a tuple with the Receipts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceipts

`func (o *DealResponseDeal) SetReceipts(v []DealResponseDealReceipts)`

SetReceipts sets Receipts field to given value.

### HasReceipts

`func (o *DealResponseDeal) HasReceipts() bool`

HasReceipts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


