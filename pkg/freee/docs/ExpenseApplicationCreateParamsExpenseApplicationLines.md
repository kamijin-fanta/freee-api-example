# ExpenseApplicationCreateParamsExpenseApplicationLines

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransactionDate** | Pointer to **string** | 日付 (yyyy-mm-dd) | [optional] 
**Description** | Pointer to **string** | 内容 (250文字以内) | [optional] 
**Amount** | Pointer to **int32** | 金額 | [optional] 
**ExpenseApplicationLineTemplateId** | Pointer to **int32** | 経費科目ID | [optional] 
**ReceiptId** | Pointer to **int32** | 証憑ID | [optional] 

## Methods

### NewExpenseApplicationCreateParamsExpenseApplicationLines

`func NewExpenseApplicationCreateParamsExpenseApplicationLines() *ExpenseApplicationCreateParamsExpenseApplicationLines`

NewExpenseApplicationCreateParamsExpenseApplicationLines instantiates a new ExpenseApplicationCreateParamsExpenseApplicationLines object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExpenseApplicationCreateParamsExpenseApplicationLinesWithDefaults

`func NewExpenseApplicationCreateParamsExpenseApplicationLinesWithDefaults() *ExpenseApplicationCreateParamsExpenseApplicationLines`

NewExpenseApplicationCreateParamsExpenseApplicationLinesWithDefaults instantiates a new ExpenseApplicationCreateParamsExpenseApplicationLines object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactionDate

`func (o *ExpenseApplicationCreateParamsExpenseApplicationLines) GetTransactionDate() string`

GetTransactionDate returns the TransactionDate field if non-nil, zero value otherwise.

### GetTransactionDateOk

`func (o *ExpenseApplicationCreateParamsExpenseApplicationLines) GetTransactionDateOk() (*string, bool)`

GetTransactionDateOk returns a tuple with the TransactionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionDate

`func (o *ExpenseApplicationCreateParamsExpenseApplicationLines) SetTransactionDate(v string)`

SetTransactionDate sets TransactionDate field to given value.

### HasTransactionDate

`func (o *ExpenseApplicationCreateParamsExpenseApplicationLines) HasTransactionDate() bool`

HasTransactionDate returns a boolean if a field has been set.

### GetDescription

`func (o *ExpenseApplicationCreateParamsExpenseApplicationLines) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ExpenseApplicationCreateParamsExpenseApplicationLines) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ExpenseApplicationCreateParamsExpenseApplicationLines) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ExpenseApplicationCreateParamsExpenseApplicationLines) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAmount

`func (o *ExpenseApplicationCreateParamsExpenseApplicationLines) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ExpenseApplicationCreateParamsExpenseApplicationLines) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ExpenseApplicationCreateParamsExpenseApplicationLines) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *ExpenseApplicationCreateParamsExpenseApplicationLines) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetExpenseApplicationLineTemplateId

`func (o *ExpenseApplicationCreateParamsExpenseApplicationLines) GetExpenseApplicationLineTemplateId() int32`

GetExpenseApplicationLineTemplateId returns the ExpenseApplicationLineTemplateId field if non-nil, zero value otherwise.

### GetExpenseApplicationLineTemplateIdOk

`func (o *ExpenseApplicationCreateParamsExpenseApplicationLines) GetExpenseApplicationLineTemplateIdOk() (*int32, bool)`

GetExpenseApplicationLineTemplateIdOk returns a tuple with the ExpenseApplicationLineTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpenseApplicationLineTemplateId

`func (o *ExpenseApplicationCreateParamsExpenseApplicationLines) SetExpenseApplicationLineTemplateId(v int32)`

SetExpenseApplicationLineTemplateId sets ExpenseApplicationLineTemplateId field to given value.

### HasExpenseApplicationLineTemplateId

`func (o *ExpenseApplicationCreateParamsExpenseApplicationLines) HasExpenseApplicationLineTemplateId() bool`

HasExpenseApplicationLineTemplateId returns a boolean if a field has been set.

### GetReceiptId

`func (o *ExpenseApplicationCreateParamsExpenseApplicationLines) GetReceiptId() int32`

GetReceiptId returns the ReceiptId field if non-nil, zero value otherwise.

### GetReceiptIdOk

`func (o *ExpenseApplicationCreateParamsExpenseApplicationLines) GetReceiptIdOk() (*int32, bool)`

GetReceiptIdOk returns a tuple with the ReceiptId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiptId

`func (o *ExpenseApplicationCreateParamsExpenseApplicationLines) SetReceiptId(v int32)`

SetReceiptId sets ReceiptId field to given value.

### HasReceiptId

`func (o *ExpenseApplicationCreateParamsExpenseApplicationLines) HasReceiptId() bool`

HasReceiptId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


