# AccountItemsResponseAccountItems

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | 勘定科目ID | 
**Name** | **string** | 勘定科目名 (30文字以内) | 
**TaxCode** | **int32** | 税区分コード | 
**Shortcut** | Pointer to **NullableString** | ショートカット1 (20文字以内) | [optional] 
**ShortcutNum** | Pointer to **NullableString** | ショートカット2(勘定科目コード) (20文字以内) | [optional] 
**DefaultTaxId** | Pointer to **int32** | デフォルト設定がされている税区分ID | [optional] 
**DefaultTaxCode** | **int32** | デフォルト設定がされている税区分コード | 
**AccountCategory** | **string** | 勘定科目カテゴリー | 
**AccountCategoryId** | **int32** | 勘定科目のカテゴリーID | 
**Categories** | **[]string** |  | 
**Available** | **bool** | 勘定科目の使用設定（true: 使用する、false: 使用しない） | 
**WalletableId** | **NullableInt32** | 口座ID | 
**GroupName** | Pointer to **NullableString** | 決算書表示名（小カテゴリー） | [optional] 
**CorrespondingIncomeName** | Pointer to **NullableString** | 収入取引相手勘定科目名 | [optional] 
**CorrespondingIncomeId** | Pointer to **NullableInt32** | 収入取引相手勘定科目ID | [optional] 
**CorrespondingExpenseName** | Pointer to **NullableString** | 支出取引相手勘定科目名 | [optional] 
**CorrespondingExpenseId** | Pointer to **NullableInt32** | 支出取引相手勘定科目ID | [optional] 

## Methods

### NewAccountItemsResponseAccountItems

`func NewAccountItemsResponseAccountItems(id int32, name string, taxCode int32, defaultTaxCode int32, accountCategory string, accountCategoryId int32, categories []string, available bool, walletableId NullableInt32, ) *AccountItemsResponseAccountItems`

NewAccountItemsResponseAccountItems instantiates a new AccountItemsResponseAccountItems object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountItemsResponseAccountItemsWithDefaults

`func NewAccountItemsResponseAccountItemsWithDefaults() *AccountItemsResponseAccountItems`

NewAccountItemsResponseAccountItemsWithDefaults instantiates a new AccountItemsResponseAccountItems object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AccountItemsResponseAccountItems) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccountItemsResponseAccountItems) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccountItemsResponseAccountItems) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *AccountItemsResponseAccountItems) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccountItemsResponseAccountItems) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccountItemsResponseAccountItems) SetName(v string)`

SetName sets Name field to given value.


### GetTaxCode

`func (o *AccountItemsResponseAccountItems) GetTaxCode() int32`

GetTaxCode returns the TaxCode field if non-nil, zero value otherwise.

### GetTaxCodeOk

`func (o *AccountItemsResponseAccountItems) GetTaxCodeOk() (*int32, bool)`

GetTaxCodeOk returns a tuple with the TaxCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCode

`func (o *AccountItemsResponseAccountItems) SetTaxCode(v int32)`

SetTaxCode sets TaxCode field to given value.


### GetShortcut

`func (o *AccountItemsResponseAccountItems) GetShortcut() string`

GetShortcut returns the Shortcut field if non-nil, zero value otherwise.

### GetShortcutOk

`func (o *AccountItemsResponseAccountItems) GetShortcutOk() (*string, bool)`

GetShortcutOk returns a tuple with the Shortcut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortcut

`func (o *AccountItemsResponseAccountItems) SetShortcut(v string)`

SetShortcut sets Shortcut field to given value.

### HasShortcut

`func (o *AccountItemsResponseAccountItems) HasShortcut() bool`

HasShortcut returns a boolean if a field has been set.

### SetShortcutNil

`func (o *AccountItemsResponseAccountItems) SetShortcutNil(b bool)`

 SetShortcutNil sets the value for Shortcut to be an explicit nil

### UnsetShortcut
`func (o *AccountItemsResponseAccountItems) UnsetShortcut()`

UnsetShortcut ensures that no value is present for Shortcut, not even an explicit nil
### GetShortcutNum

`func (o *AccountItemsResponseAccountItems) GetShortcutNum() string`

GetShortcutNum returns the ShortcutNum field if non-nil, zero value otherwise.

### GetShortcutNumOk

`func (o *AccountItemsResponseAccountItems) GetShortcutNumOk() (*string, bool)`

GetShortcutNumOk returns a tuple with the ShortcutNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortcutNum

`func (o *AccountItemsResponseAccountItems) SetShortcutNum(v string)`

SetShortcutNum sets ShortcutNum field to given value.

### HasShortcutNum

`func (o *AccountItemsResponseAccountItems) HasShortcutNum() bool`

HasShortcutNum returns a boolean if a field has been set.

### SetShortcutNumNil

`func (o *AccountItemsResponseAccountItems) SetShortcutNumNil(b bool)`

 SetShortcutNumNil sets the value for ShortcutNum to be an explicit nil

### UnsetShortcutNum
`func (o *AccountItemsResponseAccountItems) UnsetShortcutNum()`

UnsetShortcutNum ensures that no value is present for ShortcutNum, not even an explicit nil
### GetDefaultTaxId

`func (o *AccountItemsResponseAccountItems) GetDefaultTaxId() int32`

GetDefaultTaxId returns the DefaultTaxId field if non-nil, zero value otherwise.

### GetDefaultTaxIdOk

`func (o *AccountItemsResponseAccountItems) GetDefaultTaxIdOk() (*int32, bool)`

GetDefaultTaxIdOk returns a tuple with the DefaultTaxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTaxId

`func (o *AccountItemsResponseAccountItems) SetDefaultTaxId(v int32)`

SetDefaultTaxId sets DefaultTaxId field to given value.

### HasDefaultTaxId

`func (o *AccountItemsResponseAccountItems) HasDefaultTaxId() bool`

HasDefaultTaxId returns a boolean if a field has been set.

### GetDefaultTaxCode

`func (o *AccountItemsResponseAccountItems) GetDefaultTaxCode() int32`

GetDefaultTaxCode returns the DefaultTaxCode field if non-nil, zero value otherwise.

### GetDefaultTaxCodeOk

`func (o *AccountItemsResponseAccountItems) GetDefaultTaxCodeOk() (*int32, bool)`

GetDefaultTaxCodeOk returns a tuple with the DefaultTaxCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTaxCode

`func (o *AccountItemsResponseAccountItems) SetDefaultTaxCode(v int32)`

SetDefaultTaxCode sets DefaultTaxCode field to given value.


### GetAccountCategory

`func (o *AccountItemsResponseAccountItems) GetAccountCategory() string`

GetAccountCategory returns the AccountCategory field if non-nil, zero value otherwise.

### GetAccountCategoryOk

`func (o *AccountItemsResponseAccountItems) GetAccountCategoryOk() (*string, bool)`

GetAccountCategoryOk returns a tuple with the AccountCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountCategory

`func (o *AccountItemsResponseAccountItems) SetAccountCategory(v string)`

SetAccountCategory sets AccountCategory field to given value.


### GetAccountCategoryId

`func (o *AccountItemsResponseAccountItems) GetAccountCategoryId() int32`

GetAccountCategoryId returns the AccountCategoryId field if non-nil, zero value otherwise.

### GetAccountCategoryIdOk

`func (o *AccountItemsResponseAccountItems) GetAccountCategoryIdOk() (*int32, bool)`

GetAccountCategoryIdOk returns a tuple with the AccountCategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountCategoryId

`func (o *AccountItemsResponseAccountItems) SetAccountCategoryId(v int32)`

SetAccountCategoryId sets AccountCategoryId field to given value.


### GetCategories

`func (o *AccountItemsResponseAccountItems) GetCategories() []string`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *AccountItemsResponseAccountItems) GetCategoriesOk() (*[]string, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *AccountItemsResponseAccountItems) SetCategories(v []string)`

SetCategories sets Categories field to given value.


### GetAvailable

`func (o *AccountItemsResponseAccountItems) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *AccountItemsResponseAccountItems) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *AccountItemsResponseAccountItems) SetAvailable(v bool)`

SetAvailable sets Available field to given value.


### GetWalletableId

`func (o *AccountItemsResponseAccountItems) GetWalletableId() int32`

GetWalletableId returns the WalletableId field if non-nil, zero value otherwise.

### GetWalletableIdOk

`func (o *AccountItemsResponseAccountItems) GetWalletableIdOk() (*int32, bool)`

GetWalletableIdOk returns a tuple with the WalletableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletableId

`func (o *AccountItemsResponseAccountItems) SetWalletableId(v int32)`

SetWalletableId sets WalletableId field to given value.


### SetWalletableIdNil

`func (o *AccountItemsResponseAccountItems) SetWalletableIdNil(b bool)`

 SetWalletableIdNil sets the value for WalletableId to be an explicit nil

### UnsetWalletableId
`func (o *AccountItemsResponseAccountItems) UnsetWalletableId()`

UnsetWalletableId ensures that no value is present for WalletableId, not even an explicit nil
### GetGroupName

`func (o *AccountItemsResponseAccountItems) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *AccountItemsResponseAccountItems) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *AccountItemsResponseAccountItems) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *AccountItemsResponseAccountItems) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### SetGroupNameNil

`func (o *AccountItemsResponseAccountItems) SetGroupNameNil(b bool)`

 SetGroupNameNil sets the value for GroupName to be an explicit nil

### UnsetGroupName
`func (o *AccountItemsResponseAccountItems) UnsetGroupName()`

UnsetGroupName ensures that no value is present for GroupName, not even an explicit nil
### GetCorrespondingIncomeName

`func (o *AccountItemsResponseAccountItems) GetCorrespondingIncomeName() string`

GetCorrespondingIncomeName returns the CorrespondingIncomeName field if non-nil, zero value otherwise.

### GetCorrespondingIncomeNameOk

`func (o *AccountItemsResponseAccountItems) GetCorrespondingIncomeNameOk() (*string, bool)`

GetCorrespondingIncomeNameOk returns a tuple with the CorrespondingIncomeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrespondingIncomeName

`func (o *AccountItemsResponseAccountItems) SetCorrespondingIncomeName(v string)`

SetCorrespondingIncomeName sets CorrespondingIncomeName field to given value.

### HasCorrespondingIncomeName

`func (o *AccountItemsResponseAccountItems) HasCorrespondingIncomeName() bool`

HasCorrespondingIncomeName returns a boolean if a field has been set.

### SetCorrespondingIncomeNameNil

`func (o *AccountItemsResponseAccountItems) SetCorrespondingIncomeNameNil(b bool)`

 SetCorrespondingIncomeNameNil sets the value for CorrespondingIncomeName to be an explicit nil

### UnsetCorrespondingIncomeName
`func (o *AccountItemsResponseAccountItems) UnsetCorrespondingIncomeName()`

UnsetCorrespondingIncomeName ensures that no value is present for CorrespondingIncomeName, not even an explicit nil
### GetCorrespondingIncomeId

`func (o *AccountItemsResponseAccountItems) GetCorrespondingIncomeId() int32`

GetCorrespondingIncomeId returns the CorrespondingIncomeId field if non-nil, zero value otherwise.

### GetCorrespondingIncomeIdOk

`func (o *AccountItemsResponseAccountItems) GetCorrespondingIncomeIdOk() (*int32, bool)`

GetCorrespondingIncomeIdOk returns a tuple with the CorrespondingIncomeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrespondingIncomeId

`func (o *AccountItemsResponseAccountItems) SetCorrespondingIncomeId(v int32)`

SetCorrespondingIncomeId sets CorrespondingIncomeId field to given value.

### HasCorrespondingIncomeId

`func (o *AccountItemsResponseAccountItems) HasCorrespondingIncomeId() bool`

HasCorrespondingIncomeId returns a boolean if a field has been set.

### SetCorrespondingIncomeIdNil

`func (o *AccountItemsResponseAccountItems) SetCorrespondingIncomeIdNil(b bool)`

 SetCorrespondingIncomeIdNil sets the value for CorrespondingIncomeId to be an explicit nil

### UnsetCorrespondingIncomeId
`func (o *AccountItemsResponseAccountItems) UnsetCorrespondingIncomeId()`

UnsetCorrespondingIncomeId ensures that no value is present for CorrespondingIncomeId, not even an explicit nil
### GetCorrespondingExpenseName

`func (o *AccountItemsResponseAccountItems) GetCorrespondingExpenseName() string`

GetCorrespondingExpenseName returns the CorrespondingExpenseName field if non-nil, zero value otherwise.

### GetCorrespondingExpenseNameOk

`func (o *AccountItemsResponseAccountItems) GetCorrespondingExpenseNameOk() (*string, bool)`

GetCorrespondingExpenseNameOk returns a tuple with the CorrespondingExpenseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrespondingExpenseName

`func (o *AccountItemsResponseAccountItems) SetCorrespondingExpenseName(v string)`

SetCorrespondingExpenseName sets CorrespondingExpenseName field to given value.

### HasCorrespondingExpenseName

`func (o *AccountItemsResponseAccountItems) HasCorrespondingExpenseName() bool`

HasCorrespondingExpenseName returns a boolean if a field has been set.

### SetCorrespondingExpenseNameNil

`func (o *AccountItemsResponseAccountItems) SetCorrespondingExpenseNameNil(b bool)`

 SetCorrespondingExpenseNameNil sets the value for CorrespondingExpenseName to be an explicit nil

### UnsetCorrespondingExpenseName
`func (o *AccountItemsResponseAccountItems) UnsetCorrespondingExpenseName()`

UnsetCorrespondingExpenseName ensures that no value is present for CorrespondingExpenseName, not even an explicit nil
### GetCorrespondingExpenseId

`func (o *AccountItemsResponseAccountItems) GetCorrespondingExpenseId() int32`

GetCorrespondingExpenseId returns the CorrespondingExpenseId field if non-nil, zero value otherwise.

### GetCorrespondingExpenseIdOk

`func (o *AccountItemsResponseAccountItems) GetCorrespondingExpenseIdOk() (*int32, bool)`

GetCorrespondingExpenseIdOk returns a tuple with the CorrespondingExpenseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrespondingExpenseId

`func (o *AccountItemsResponseAccountItems) SetCorrespondingExpenseId(v int32)`

SetCorrespondingExpenseId sets CorrespondingExpenseId field to given value.

### HasCorrespondingExpenseId

`func (o *AccountItemsResponseAccountItems) HasCorrespondingExpenseId() bool`

HasCorrespondingExpenseId returns a boolean if a field has been set.

### SetCorrespondingExpenseIdNil

`func (o *AccountItemsResponseAccountItems) SetCorrespondingExpenseIdNil(b bool)`

 SetCorrespondingExpenseIdNil sets the value for CorrespondingExpenseId to be an explicit nil

### UnsetCorrespondingExpenseId
`func (o *AccountItemsResponseAccountItems) UnsetCorrespondingExpenseId()`

UnsetCorrespondingExpenseId ensures that no value is present for CorrespondingExpenseId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


