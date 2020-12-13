# ExpenseApplicationResponseExpenseApplicationExpenseApplicationLines

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | 経費申請の項目行ID | 
**TransactionDate** | Pointer to **NullableString** | 日付 (yyyy-mm-dd) | [optional] 
**Description** | Pointer to **NullableString** | 内容 | [optional] 
**Amount** | Pointer to **int32** | 金額 | [optional] 
**ExpenseApplicationLineTemplateId** | Pointer to **NullableInt32** | 経費科目ID | [optional] 
**ReceiptId** | Pointer to **NullableInt32** | 証憑ID | [optional] 

## Methods

### NewExpenseApplicationResponseExpenseApplicationExpenseApplicationLines

`func NewExpenseApplicationResponseExpenseApplicationExpenseApplicationLines(id int32, ) *ExpenseApplicationResponseExpenseApplicationExpenseApplicationLines`

NewExpenseApplicationResponseExpenseApplicationExpenseApplicationLines instantiates a new ExpenseApplicationResponseExpenseApplicationExpenseApplicationLines object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExpenseApplicationResponseExpenseApplicationExpenseApplicationLinesWithDefaults

`func NewExpenseApplicationResponseExpenseApplicationExpenseApplicationLinesWithDefaults() *ExpenseApplicationResponseExpenseApplicationExpenseApplicationLines`

NewExpenseApplicationResponseExpenseApplicationExpenseApplicationLinesWithDefaults instantiates a new ExpenseApplicationResponseExpenseApplicationExpenseApplicationLines object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ExpenseApplicationResponseExpenseApplicationExpenseApplicationLines) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExpenseApplicationResponseExpenseApplicationExpenseApplicationLines) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExpenseApplicationResponseExpenseApplicationExpenseApplicationLines) SetId(v int32)`

SetId sets Id field to given value.


### GetTransactionDate

`func (o *ExpenseApplicationResponseExpenseApplicationExpenseApplicationLines) GetTransactionDate() string`

GetTransactionDate returns the TransactionDate field if non-nil, zero value otherwise.

### GetTransactionDateOk

`func (o *ExpenseApplicationResponseExpenseApplicationExpenseApplicationLines) GetTransactionDateOk() (*string, bool)`

GetTransactionDateOk returns a tuple with the TransactionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionDate

`func (o *ExpenseApplicationResponseExpenseApplicationExpenseApplicationLines) SetTransactionDate(v string)`

SetTransactionDate sets TransactionDate field to given value.

### HasTransactionDate

`func (o *ExpenseApplicationResponseExpenseApplicationExpenseApplicationLines) HasTransactionDate() bool`

HasTransactionDate returns a boolean if a field has been set.

### SetTransactionDateNil

`func (o *ExpenseApplicationResponseExpenseApplicationExpenseApplicationLines) SetTransactionDateNil(b bool)`

 SetTransactionDateNil sets the value for TransactionDate to be an explicit nil

### UnsetTransactionDate
`func (o *ExpenseApplicationResponseExpenseApplicationExpenseApplicationLines) UnsetTransactionDate()`

UnsetTransactionDate ensures that no value is present for TransactionDate, not even an explicit nil
### GetDescription

`func (o *ExpenseApplicationResponseExpenseApplicationExpenseApplicationLines) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ExpenseApplicationResponseExpenseApplicationExpenseApplicationLines) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ExpenseApplicationResponseExpenseApplicationExpenseApplicationLines) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ExpenseApplicationResponseExpenseApplicationExpenseApplicationLines) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ExpenseApplicationResponseExpenseApplicationExpenseApplicationLines) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ExpenseApplicationResponseExpenseApplicationExpenseApplicationLines) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetAmount

`func (o *ExpenseApplicationResponseExpenseApplicationExpenseApplicationLines) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ExpenseApplicationResponseExpenseApplicationExpenseApplicationLines) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ExpenseApplicationResponseExpenseApplicationExpenseApplicationLines) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *ExpenseApplicationResponseExpenseApplicationExpenseApplicationLines) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetExpenseApplicationLineTemplateId

`func (o *ExpenseApplicationResponseExpenseApplicationExpenseApplicationLines) GetExpenseApplicationLineTemplateId() int32`

GetExpenseApplicationLineTemplateId returns the ExpenseApplicationLineTemplateId field if non-nil, zero value otherwise.

### GetExpenseApplicationLineTemplateIdOk

`func (o *ExpenseApplicationResponseExpenseApplicationExpenseApplicationLines) GetExpenseApplicationLineTemplateIdOk() (*int32, bool)`

GetExpenseApplicationLineTemplateIdOk returns a tuple with the ExpenseApplicationLineTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpenseApplicationLineTemplateId

`func (o *ExpenseApplicationResponseExpenseApplicationExpenseApplicationLines) SetExpenseApplicationLineTemplateId(v int32)`

SetExpenseApplicationLineTemplateId sets ExpenseApplicationLineTemplateId field to given value.

### HasExpenseApplicationLineTemplateId

`func (o *ExpenseApplicationResponseExpenseApplicationExpenseApplicationLines) HasExpenseApplicationLineTemplateId() bool`

HasExpenseApplicationLineTemplateId returns a boolean if a field has been set.

### SetExpenseApplicationLineTemplateIdNil

`func (o *ExpenseApplicationResponseExpenseApplicationExpenseApplicationLines) SetExpenseApplicationLineTemplateIdNil(b bool)`

 SetExpenseApplicationLineTemplateIdNil sets the value for ExpenseApplicationLineTemplateId to be an explicit nil

### UnsetExpenseApplicationLineTemplateId
`func (o *ExpenseApplicationResponseExpenseApplicationExpenseApplicationLines) UnsetExpenseApplicationLineTemplateId()`

UnsetExpenseApplicationLineTemplateId ensures that no value is present for ExpenseApplicationLineTemplateId, not even an explicit nil
### GetReceiptId

`func (o *ExpenseApplicationResponseExpenseApplicationExpenseApplicationLines) GetReceiptId() int32`

GetReceiptId returns the ReceiptId field if non-nil, zero value otherwise.

### GetReceiptIdOk

`func (o *ExpenseApplicationResponseExpenseApplicationExpenseApplicationLines) GetReceiptIdOk() (*int32, bool)`

GetReceiptIdOk returns a tuple with the ReceiptId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiptId

`func (o *ExpenseApplicationResponseExpenseApplicationExpenseApplicationLines) SetReceiptId(v int32)`

SetReceiptId sets ReceiptId field to given value.

### HasReceiptId

`func (o *ExpenseApplicationResponseExpenseApplicationExpenseApplicationLines) HasReceiptId() bool`

HasReceiptId returns a boolean if a field has been set.

### SetReceiptIdNil

`func (o *ExpenseApplicationResponseExpenseApplicationExpenseApplicationLines) SetReceiptIdNil(b bool)`

 SetReceiptIdNil sets the value for ReceiptId to be an explicit nil

### UnsetReceiptId
`func (o *ExpenseApplicationResponseExpenseApplicationExpenseApplicationLines) UnsetReceiptId()`

UnsetReceiptId ensures that no value is present for ReceiptId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


