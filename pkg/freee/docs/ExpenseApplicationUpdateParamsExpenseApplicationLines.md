# ExpenseApplicationUpdateParamsExpenseApplicationLines

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | 経費申請の項目行ID: 既存項目行を更新する場合に指定します。IDを指定しない項目行は、新規行として扱われ追加されます。また、expense_application_linesに含まれない既存の項目行は削除されます。更新後も残したい行は、必ず経費申請の項目行IDを指定してexpense_application_linesに含めてください。 | [optional] 
**TransactionDate** | Pointer to **string** | 日付 (yyyy-mm-dd) | [optional] 
**Description** | Pointer to **string** | 内容 (250文字以内) | [optional] 
**Amount** | Pointer to **int32** | 金額 | [optional] 
**ExpenseApplicationLineTemplateId** | Pointer to **int32** | 経費科目ID | [optional] 
**ReceiptId** | Pointer to **int32** | 証憑ID | [optional] 

## Methods

### NewExpenseApplicationUpdateParamsExpenseApplicationLines

`func NewExpenseApplicationUpdateParamsExpenseApplicationLines() *ExpenseApplicationUpdateParamsExpenseApplicationLines`

NewExpenseApplicationUpdateParamsExpenseApplicationLines instantiates a new ExpenseApplicationUpdateParamsExpenseApplicationLines object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExpenseApplicationUpdateParamsExpenseApplicationLinesWithDefaults

`func NewExpenseApplicationUpdateParamsExpenseApplicationLinesWithDefaults() *ExpenseApplicationUpdateParamsExpenseApplicationLines`

NewExpenseApplicationUpdateParamsExpenseApplicationLinesWithDefaults instantiates a new ExpenseApplicationUpdateParamsExpenseApplicationLines object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ExpenseApplicationUpdateParamsExpenseApplicationLines) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExpenseApplicationUpdateParamsExpenseApplicationLines) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExpenseApplicationUpdateParamsExpenseApplicationLines) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ExpenseApplicationUpdateParamsExpenseApplicationLines) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTransactionDate

`func (o *ExpenseApplicationUpdateParamsExpenseApplicationLines) GetTransactionDate() string`

GetTransactionDate returns the TransactionDate field if non-nil, zero value otherwise.

### GetTransactionDateOk

`func (o *ExpenseApplicationUpdateParamsExpenseApplicationLines) GetTransactionDateOk() (*string, bool)`

GetTransactionDateOk returns a tuple with the TransactionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionDate

`func (o *ExpenseApplicationUpdateParamsExpenseApplicationLines) SetTransactionDate(v string)`

SetTransactionDate sets TransactionDate field to given value.

### HasTransactionDate

`func (o *ExpenseApplicationUpdateParamsExpenseApplicationLines) HasTransactionDate() bool`

HasTransactionDate returns a boolean if a field has been set.

### GetDescription

`func (o *ExpenseApplicationUpdateParamsExpenseApplicationLines) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ExpenseApplicationUpdateParamsExpenseApplicationLines) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ExpenseApplicationUpdateParamsExpenseApplicationLines) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ExpenseApplicationUpdateParamsExpenseApplicationLines) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAmount

`func (o *ExpenseApplicationUpdateParamsExpenseApplicationLines) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ExpenseApplicationUpdateParamsExpenseApplicationLines) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ExpenseApplicationUpdateParamsExpenseApplicationLines) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *ExpenseApplicationUpdateParamsExpenseApplicationLines) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetExpenseApplicationLineTemplateId

`func (o *ExpenseApplicationUpdateParamsExpenseApplicationLines) GetExpenseApplicationLineTemplateId() int32`

GetExpenseApplicationLineTemplateId returns the ExpenseApplicationLineTemplateId field if non-nil, zero value otherwise.

### GetExpenseApplicationLineTemplateIdOk

`func (o *ExpenseApplicationUpdateParamsExpenseApplicationLines) GetExpenseApplicationLineTemplateIdOk() (*int32, bool)`

GetExpenseApplicationLineTemplateIdOk returns a tuple with the ExpenseApplicationLineTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpenseApplicationLineTemplateId

`func (o *ExpenseApplicationUpdateParamsExpenseApplicationLines) SetExpenseApplicationLineTemplateId(v int32)`

SetExpenseApplicationLineTemplateId sets ExpenseApplicationLineTemplateId field to given value.

### HasExpenseApplicationLineTemplateId

`func (o *ExpenseApplicationUpdateParamsExpenseApplicationLines) HasExpenseApplicationLineTemplateId() bool`

HasExpenseApplicationLineTemplateId returns a boolean if a field has been set.

### GetReceiptId

`func (o *ExpenseApplicationUpdateParamsExpenseApplicationLines) GetReceiptId() int32`

GetReceiptId returns the ReceiptId field if non-nil, zero value otherwise.

### GetReceiptIdOk

`func (o *ExpenseApplicationUpdateParamsExpenseApplicationLines) GetReceiptIdOk() (*int32, bool)`

GetReceiptIdOk returns a tuple with the ReceiptId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiptId

`func (o *ExpenseApplicationUpdateParamsExpenseApplicationLines) SetReceiptId(v int32)`

SetReceiptId sets ReceiptId field to given value.

### HasReceiptId

`func (o *ExpenseApplicationUpdateParamsExpenseApplicationLines) HasReceiptId() bool`

HasReceiptId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


