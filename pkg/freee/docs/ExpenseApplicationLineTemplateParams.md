# ExpenseApplicationLineTemplateParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyId** | **int32** | 事業所ID | 
**Name** | **string** | 経費科目名 (100文字以内) | 
**AccountItemId** | **int32** | 勘定科目ID | 
**ItemId** | Pointer to **int32** | 品目ID | [optional] 
**TaxCode** | **int32** | 税区分コード（税区分のdisplay_categoryがtax_5: 5%表示の税区分, tax_r8: 軽減税率8%表示の税区分に該当するtax_codeのみ利用可能です。税区分のdisplay_categoryは /taxes/companies/{:company_id}のAPIから取得可能です。） | 
**Description** | Pointer to **string** | 経費科目の説明 (1000文字以内) | [optional] 
**LineDescription** | Pointer to **string** | 内容の補足 (1000文字以内) | [optional] 

## Methods

### NewExpenseApplicationLineTemplateParams

`func NewExpenseApplicationLineTemplateParams(companyId int32, name string, accountItemId int32, taxCode int32, ) *ExpenseApplicationLineTemplateParams`

NewExpenseApplicationLineTemplateParams instantiates a new ExpenseApplicationLineTemplateParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExpenseApplicationLineTemplateParamsWithDefaults

`func NewExpenseApplicationLineTemplateParamsWithDefaults() *ExpenseApplicationLineTemplateParams`

NewExpenseApplicationLineTemplateParamsWithDefaults instantiates a new ExpenseApplicationLineTemplateParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyId

`func (o *ExpenseApplicationLineTemplateParams) GetCompanyId() int32`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *ExpenseApplicationLineTemplateParams) GetCompanyIdOk() (*int32, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *ExpenseApplicationLineTemplateParams) SetCompanyId(v int32)`

SetCompanyId sets CompanyId field to given value.


### GetName

`func (o *ExpenseApplicationLineTemplateParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExpenseApplicationLineTemplateParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExpenseApplicationLineTemplateParams) SetName(v string)`

SetName sets Name field to given value.


### GetAccountItemId

`func (o *ExpenseApplicationLineTemplateParams) GetAccountItemId() int32`

GetAccountItemId returns the AccountItemId field if non-nil, zero value otherwise.

### GetAccountItemIdOk

`func (o *ExpenseApplicationLineTemplateParams) GetAccountItemIdOk() (*int32, bool)`

GetAccountItemIdOk returns a tuple with the AccountItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountItemId

`func (o *ExpenseApplicationLineTemplateParams) SetAccountItemId(v int32)`

SetAccountItemId sets AccountItemId field to given value.


### GetItemId

`func (o *ExpenseApplicationLineTemplateParams) GetItemId() int32`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *ExpenseApplicationLineTemplateParams) GetItemIdOk() (*int32, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *ExpenseApplicationLineTemplateParams) SetItemId(v int32)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *ExpenseApplicationLineTemplateParams) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### GetTaxCode

`func (o *ExpenseApplicationLineTemplateParams) GetTaxCode() int32`

GetTaxCode returns the TaxCode field if non-nil, zero value otherwise.

### GetTaxCodeOk

`func (o *ExpenseApplicationLineTemplateParams) GetTaxCodeOk() (*int32, bool)`

GetTaxCodeOk returns a tuple with the TaxCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCode

`func (o *ExpenseApplicationLineTemplateParams) SetTaxCode(v int32)`

SetTaxCode sets TaxCode field to given value.


### GetDescription

`func (o *ExpenseApplicationLineTemplateParams) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ExpenseApplicationLineTemplateParams) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ExpenseApplicationLineTemplateParams) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ExpenseApplicationLineTemplateParams) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLineDescription

`func (o *ExpenseApplicationLineTemplateParams) GetLineDescription() string`

GetLineDescription returns the LineDescription field if non-nil, zero value otherwise.

### GetLineDescriptionOk

`func (o *ExpenseApplicationLineTemplateParams) GetLineDescriptionOk() (*string, bool)`

GetLineDescriptionOk returns a tuple with the LineDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineDescription

`func (o *ExpenseApplicationLineTemplateParams) SetLineDescription(v string)`

SetLineDescription sets LineDescription field to given value.

### HasLineDescription

`func (o *ExpenseApplicationLineTemplateParams) HasLineDescription() bool`

HasLineDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


