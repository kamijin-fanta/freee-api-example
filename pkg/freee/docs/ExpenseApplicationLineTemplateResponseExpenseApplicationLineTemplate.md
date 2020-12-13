# ExpenseApplicationLineTemplateResponseExpenseApplicationLineTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | 経費科目ID | 
**Name** | **string** | 経費科目名 | 
**AccountItemId** | Pointer to **int32** | 勘定科目ID | [optional] 
**AccountItemName** | **string** | 勘定科目名 | 
**TaxCode** | Pointer to **int32** | 税区分コード | [optional] 
**TaxName** | **string** | 税区分名 | 
**Description** | Pointer to **string** | 経費科目の説明 | [optional] 
**LineDescription** | Pointer to **string** | 内容の補足 | [optional] 

## Methods

### NewExpenseApplicationLineTemplateResponseExpenseApplicationLineTemplate

`func NewExpenseApplicationLineTemplateResponseExpenseApplicationLineTemplate(id int32, name string, accountItemName string, taxName string, ) *ExpenseApplicationLineTemplateResponseExpenseApplicationLineTemplate`

NewExpenseApplicationLineTemplateResponseExpenseApplicationLineTemplate instantiates a new ExpenseApplicationLineTemplateResponseExpenseApplicationLineTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExpenseApplicationLineTemplateResponseExpenseApplicationLineTemplateWithDefaults

`func NewExpenseApplicationLineTemplateResponseExpenseApplicationLineTemplateWithDefaults() *ExpenseApplicationLineTemplateResponseExpenseApplicationLineTemplate`

NewExpenseApplicationLineTemplateResponseExpenseApplicationLineTemplateWithDefaults instantiates a new ExpenseApplicationLineTemplateResponseExpenseApplicationLineTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ExpenseApplicationLineTemplateResponseExpenseApplicationLineTemplate) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExpenseApplicationLineTemplateResponseExpenseApplicationLineTemplate) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExpenseApplicationLineTemplateResponseExpenseApplicationLineTemplate) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *ExpenseApplicationLineTemplateResponseExpenseApplicationLineTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExpenseApplicationLineTemplateResponseExpenseApplicationLineTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExpenseApplicationLineTemplateResponseExpenseApplicationLineTemplate) SetName(v string)`

SetName sets Name field to given value.


### GetAccountItemId

`func (o *ExpenseApplicationLineTemplateResponseExpenseApplicationLineTemplate) GetAccountItemId() int32`

GetAccountItemId returns the AccountItemId field if non-nil, zero value otherwise.

### GetAccountItemIdOk

`func (o *ExpenseApplicationLineTemplateResponseExpenseApplicationLineTemplate) GetAccountItemIdOk() (*int32, bool)`

GetAccountItemIdOk returns a tuple with the AccountItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountItemId

`func (o *ExpenseApplicationLineTemplateResponseExpenseApplicationLineTemplate) SetAccountItemId(v int32)`

SetAccountItemId sets AccountItemId field to given value.

### HasAccountItemId

`func (o *ExpenseApplicationLineTemplateResponseExpenseApplicationLineTemplate) HasAccountItemId() bool`

HasAccountItemId returns a boolean if a field has been set.

### GetAccountItemName

`func (o *ExpenseApplicationLineTemplateResponseExpenseApplicationLineTemplate) GetAccountItemName() string`

GetAccountItemName returns the AccountItemName field if non-nil, zero value otherwise.

### GetAccountItemNameOk

`func (o *ExpenseApplicationLineTemplateResponseExpenseApplicationLineTemplate) GetAccountItemNameOk() (*string, bool)`

GetAccountItemNameOk returns a tuple with the AccountItemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountItemName

`func (o *ExpenseApplicationLineTemplateResponseExpenseApplicationLineTemplate) SetAccountItemName(v string)`

SetAccountItemName sets AccountItemName field to given value.


### GetTaxCode

`func (o *ExpenseApplicationLineTemplateResponseExpenseApplicationLineTemplate) GetTaxCode() int32`

GetTaxCode returns the TaxCode field if non-nil, zero value otherwise.

### GetTaxCodeOk

`func (o *ExpenseApplicationLineTemplateResponseExpenseApplicationLineTemplate) GetTaxCodeOk() (*int32, bool)`

GetTaxCodeOk returns a tuple with the TaxCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCode

`func (o *ExpenseApplicationLineTemplateResponseExpenseApplicationLineTemplate) SetTaxCode(v int32)`

SetTaxCode sets TaxCode field to given value.

### HasTaxCode

`func (o *ExpenseApplicationLineTemplateResponseExpenseApplicationLineTemplate) HasTaxCode() bool`

HasTaxCode returns a boolean if a field has been set.

### GetTaxName

`func (o *ExpenseApplicationLineTemplateResponseExpenseApplicationLineTemplate) GetTaxName() string`

GetTaxName returns the TaxName field if non-nil, zero value otherwise.

### GetTaxNameOk

`func (o *ExpenseApplicationLineTemplateResponseExpenseApplicationLineTemplate) GetTaxNameOk() (*string, bool)`

GetTaxNameOk returns a tuple with the TaxName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxName

`func (o *ExpenseApplicationLineTemplateResponseExpenseApplicationLineTemplate) SetTaxName(v string)`

SetTaxName sets TaxName field to given value.


### GetDescription

`func (o *ExpenseApplicationLineTemplateResponseExpenseApplicationLineTemplate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ExpenseApplicationLineTemplateResponseExpenseApplicationLineTemplate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ExpenseApplicationLineTemplateResponseExpenseApplicationLineTemplate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ExpenseApplicationLineTemplateResponseExpenseApplicationLineTemplate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLineDescription

`func (o *ExpenseApplicationLineTemplateResponseExpenseApplicationLineTemplate) GetLineDescription() string`

GetLineDescription returns the LineDescription field if non-nil, zero value otherwise.

### GetLineDescriptionOk

`func (o *ExpenseApplicationLineTemplateResponseExpenseApplicationLineTemplate) GetLineDescriptionOk() (*string, bool)`

GetLineDescriptionOk returns a tuple with the LineDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineDescription

`func (o *ExpenseApplicationLineTemplateResponseExpenseApplicationLineTemplate) SetLineDescription(v string)`

SetLineDescription sets LineDescription field to given value.

### HasLineDescription

`func (o *ExpenseApplicationLineTemplateResponseExpenseApplicationLineTemplate) HasLineDescription() bool`

HasLineDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


