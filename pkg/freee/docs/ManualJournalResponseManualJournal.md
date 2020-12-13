# ManualJournalResponseManualJournal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | 振替伝票ID | 
**CompanyId** | **int32** | 事業所ID | 
**IssueDate** | **string** | 発生日 (yyyy-mm-dd) | 
**Adjustment** | **bool** | 決算整理仕訳フラグ（falseまたは未指定の場合: 日常仕訳） | 
**TxnNumber** | **NullableString** | 仕訳番号 | 
**Details** | [**[]ManualJournalResponseManualJournalDetails**](ManualJournalResponseManualJournalDetails.md) | 貸借行一覧（配列）: 貸借合わせて100行まで登録できます。 | 

## Methods

### NewManualJournalResponseManualJournal

`func NewManualJournalResponseManualJournal(id int32, companyId int32, issueDate string, adjustment bool, txnNumber NullableString, details []ManualJournalResponseManualJournalDetails, ) *ManualJournalResponseManualJournal`

NewManualJournalResponseManualJournal instantiates a new ManualJournalResponseManualJournal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManualJournalResponseManualJournalWithDefaults

`func NewManualJournalResponseManualJournalWithDefaults() *ManualJournalResponseManualJournal`

NewManualJournalResponseManualJournalWithDefaults instantiates a new ManualJournalResponseManualJournal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ManualJournalResponseManualJournal) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ManualJournalResponseManualJournal) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ManualJournalResponseManualJournal) SetId(v int32)`

SetId sets Id field to given value.


### GetCompanyId

`func (o *ManualJournalResponseManualJournal) GetCompanyId() int32`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *ManualJournalResponseManualJournal) GetCompanyIdOk() (*int32, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *ManualJournalResponseManualJournal) SetCompanyId(v int32)`

SetCompanyId sets CompanyId field to given value.


### GetIssueDate

`func (o *ManualJournalResponseManualJournal) GetIssueDate() string`

GetIssueDate returns the IssueDate field if non-nil, zero value otherwise.

### GetIssueDateOk

`func (o *ManualJournalResponseManualJournal) GetIssueDateOk() (*string, bool)`

GetIssueDateOk returns a tuple with the IssueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueDate

`func (o *ManualJournalResponseManualJournal) SetIssueDate(v string)`

SetIssueDate sets IssueDate field to given value.


### GetAdjustment

`func (o *ManualJournalResponseManualJournal) GetAdjustment() bool`

GetAdjustment returns the Adjustment field if non-nil, zero value otherwise.

### GetAdjustmentOk

`func (o *ManualJournalResponseManualJournal) GetAdjustmentOk() (*bool, bool)`

GetAdjustmentOk returns a tuple with the Adjustment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustment

`func (o *ManualJournalResponseManualJournal) SetAdjustment(v bool)`

SetAdjustment sets Adjustment field to given value.


### GetTxnNumber

`func (o *ManualJournalResponseManualJournal) GetTxnNumber() string`

GetTxnNumber returns the TxnNumber field if non-nil, zero value otherwise.

### GetTxnNumberOk

`func (o *ManualJournalResponseManualJournal) GetTxnNumberOk() (*string, bool)`

GetTxnNumberOk returns a tuple with the TxnNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxnNumber

`func (o *ManualJournalResponseManualJournal) SetTxnNumber(v string)`

SetTxnNumber sets TxnNumber field to given value.


### SetTxnNumberNil

`func (o *ManualJournalResponseManualJournal) SetTxnNumberNil(b bool)`

 SetTxnNumberNil sets the value for TxnNumber to be an explicit nil

### UnsetTxnNumber
`func (o *ManualJournalResponseManualJournal) UnsetTxnNumber()`

UnsetTxnNumber ensures that no value is present for TxnNumber, not even an explicit nil
### GetDetails

`func (o *ManualJournalResponseManualJournal) GetDetails() []ManualJournalResponseManualJournalDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *ManualJournalResponseManualJournal) GetDetailsOk() (*[]ManualJournalResponseManualJournalDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *ManualJournalResponseManualJournal) SetDetails(v []ManualJournalResponseManualJournalDetails)`

SetDetails sets Details field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


