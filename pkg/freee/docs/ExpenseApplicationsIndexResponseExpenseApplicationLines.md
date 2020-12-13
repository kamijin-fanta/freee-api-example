# ExpenseApplicationsIndexResponseExpenseApplicationLines

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | 経費申請の項目行ID | 
**TransactionDate** | Pointer to **string** | 日付 (yyyy-mm-dd) | [optional] 
**Description** | Pointer to **string** | 内容 | [optional] 
**Amount** | Pointer to **int32** | 金額 | [optional] 
**ExpenseApplicationLineTemplateId** | Pointer to **int32** | 経費科目ID | [optional] 
**ReceiptId** | Pointer to **int32** | 証憑ID | [optional] 

## Methods

### NewExpenseApplicationsIndexResponseExpenseApplicationLines

`func NewExpenseApplicationsIndexResponseExpenseApplicationLines(id int32, ) *ExpenseApplicationsIndexResponseExpenseApplicationLines`

NewExpenseApplicationsIndexResponseExpenseApplicationLines instantiates a new ExpenseApplicationsIndexResponseExpenseApplicationLines object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExpenseApplicationsIndexResponseExpenseApplicationLinesWithDefaults

`func NewExpenseApplicationsIndexResponseExpenseApplicationLinesWithDefaults() *ExpenseApplicationsIndexResponseExpenseApplicationLines`

NewExpenseApplicationsIndexResponseExpenseApplicationLinesWithDefaults instantiates a new ExpenseApplicationsIndexResponseExpenseApplicationLines object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ExpenseApplicationsIndexResponseExpenseApplicationLines) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExpenseApplicationsIndexResponseExpenseApplicationLines) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExpenseApplicationsIndexResponseExpenseApplicationLines) SetId(v int32)`

SetId sets Id field to given value.


### GetTransactionDate

`func (o *ExpenseApplicationsIndexResponseExpenseApplicationLines) GetTransactionDate() string`

GetTransactionDate returns the TransactionDate field if non-nil, zero value otherwise.

### GetTransactionDateOk

`func (o *ExpenseApplicationsIndexResponseExpenseApplicationLines) GetTransactionDateOk() (*string, bool)`

GetTransactionDateOk returns a tuple with the TransactionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionDate

`func (o *ExpenseApplicationsIndexResponseExpenseApplicationLines) SetTransactionDate(v string)`

SetTransactionDate sets TransactionDate field to given value.

### HasTransactionDate

`func (o *ExpenseApplicationsIndexResponseExpenseApplicationLines) HasTransactionDate() bool`

HasTransactionDate returns a boolean if a field has been set.

### GetDescription

`func (o *ExpenseApplicationsIndexResponseExpenseApplicationLines) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ExpenseApplicationsIndexResponseExpenseApplicationLines) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ExpenseApplicationsIndexResponseExpenseApplicationLines) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ExpenseApplicationsIndexResponseExpenseApplicationLines) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAmount

`func (o *ExpenseApplicationsIndexResponseExpenseApplicationLines) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ExpenseApplicationsIndexResponseExpenseApplicationLines) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ExpenseApplicationsIndexResponseExpenseApplicationLines) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *ExpenseApplicationsIndexResponseExpenseApplicationLines) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetExpenseApplicationLineTemplateId

`func (o *ExpenseApplicationsIndexResponseExpenseApplicationLines) GetExpenseApplicationLineTemplateId() int32`

GetExpenseApplicationLineTemplateId returns the ExpenseApplicationLineTemplateId field if non-nil, zero value otherwise.

### GetExpenseApplicationLineTemplateIdOk

`func (o *ExpenseApplicationsIndexResponseExpenseApplicationLines) GetExpenseApplicationLineTemplateIdOk() (*int32, bool)`

GetExpenseApplicationLineTemplateIdOk returns a tuple with the ExpenseApplicationLineTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpenseApplicationLineTemplateId

`func (o *ExpenseApplicationsIndexResponseExpenseApplicationLines) SetExpenseApplicationLineTemplateId(v int32)`

SetExpenseApplicationLineTemplateId sets ExpenseApplicationLineTemplateId field to given value.

### HasExpenseApplicationLineTemplateId

`func (o *ExpenseApplicationsIndexResponseExpenseApplicationLines) HasExpenseApplicationLineTemplateId() bool`

HasExpenseApplicationLineTemplateId returns a boolean if a field has been set.

### GetReceiptId

`func (o *ExpenseApplicationsIndexResponseExpenseApplicationLines) GetReceiptId() int32`

GetReceiptId returns the ReceiptId field if non-nil, zero value otherwise.

### GetReceiptIdOk

`func (o *ExpenseApplicationsIndexResponseExpenseApplicationLines) GetReceiptIdOk() (*int32, bool)`

GetReceiptIdOk returns a tuple with the ReceiptId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiptId

`func (o *ExpenseApplicationsIndexResponseExpenseApplicationLines) SetReceiptId(v int32)`

SetReceiptId sets ReceiptId field to given value.

### HasReceiptId

`func (o *ExpenseApplicationsIndexResponseExpenseApplicationLines) HasReceiptId() bool`

HasReceiptId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


