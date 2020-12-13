# AccountItemParamsAccountItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | 勘定科目名 (30文字以内) | 
**Shortcut** | Pointer to **string** | ショートカット1 (20文字以内) | [optional] 
**ShortcutNum** | Pointer to **string** | ショートカット2(勘定科目コード)(20文字以内) | [optional] 
**TaxCode** | **int32** | 税区分コード | 
**GroupName** | **string** | 決算書表示名（小カテゴリー） Selectablesフォーム用選択項目情報エンドポイント(account_groups.name)で取得可能です | 
**AccountCategoryId** | **int32** | 勘定科目カテゴリーID Selectablesフォーム用選択項目情報エンドポイント(account_groups.account_category_id)で取得可能です | 
**CorrespondingIncomeId** | **int32** | 収入取引相手勘定科目ID | 
**CorrespondingExpenseId** | **int32** | 支出取引相手勘定科目ID | 
**AccumulatedDepAccountItemId** | Pointer to **int32** | 減価償却累計額勘定科目ID（法人のみ利用可能） | [optional] 
**Searchable** | Pointer to **int32** | 検索可能:2, 検索不可：3(登録時未指定の場合は2で登録されます。更新時未指定の場合はsearchableは変更されません。) | [optional] 
**Items** | Pointer to [**[]AccountItemParamsAccountItemItems**](AccountItemParamsAccountItemItems.md) | 品目 | [optional] 
**Partners** | Pointer to [**[]AccountItemParamsAccountItemItems**](AccountItemParamsAccountItemItems.md) | 取引先 | [optional] 

## Methods

### NewAccountItemParamsAccountItem

`func NewAccountItemParamsAccountItem(name string, taxCode int32, groupName string, accountCategoryId int32, correspondingIncomeId int32, correspondingExpenseId int32, ) *AccountItemParamsAccountItem`

NewAccountItemParamsAccountItem instantiates a new AccountItemParamsAccountItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountItemParamsAccountItemWithDefaults

`func NewAccountItemParamsAccountItemWithDefaults() *AccountItemParamsAccountItem`

NewAccountItemParamsAccountItemWithDefaults instantiates a new AccountItemParamsAccountItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AccountItemParamsAccountItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccountItemParamsAccountItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccountItemParamsAccountItem) SetName(v string)`

SetName sets Name field to given value.


### GetShortcut

`func (o *AccountItemParamsAccountItem) GetShortcut() string`

GetShortcut returns the Shortcut field if non-nil, zero value otherwise.

### GetShortcutOk

`func (o *AccountItemParamsAccountItem) GetShortcutOk() (*string, bool)`

GetShortcutOk returns a tuple with the Shortcut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortcut

`func (o *AccountItemParamsAccountItem) SetShortcut(v string)`

SetShortcut sets Shortcut field to given value.

### HasShortcut

`func (o *AccountItemParamsAccountItem) HasShortcut() bool`

HasShortcut returns a boolean if a field has been set.

### GetShortcutNum

`func (o *AccountItemParamsAccountItem) GetShortcutNum() string`

GetShortcutNum returns the ShortcutNum field if non-nil, zero value otherwise.

### GetShortcutNumOk

`func (o *AccountItemParamsAccountItem) GetShortcutNumOk() (*string, bool)`

GetShortcutNumOk returns a tuple with the ShortcutNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortcutNum

`func (o *AccountItemParamsAccountItem) SetShortcutNum(v string)`

SetShortcutNum sets ShortcutNum field to given value.

### HasShortcutNum

`func (o *AccountItemParamsAccountItem) HasShortcutNum() bool`

HasShortcutNum returns a boolean if a field has been set.

### GetTaxCode

`func (o *AccountItemParamsAccountItem) GetTaxCode() int32`

GetTaxCode returns the TaxCode field if non-nil, zero value otherwise.

### GetTaxCodeOk

`func (o *AccountItemParamsAccountItem) GetTaxCodeOk() (*int32, bool)`

GetTaxCodeOk returns a tuple with the TaxCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCode

`func (o *AccountItemParamsAccountItem) SetTaxCode(v int32)`

SetTaxCode sets TaxCode field to given value.


### GetGroupName

`func (o *AccountItemParamsAccountItem) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *AccountItemParamsAccountItem) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *AccountItemParamsAccountItem) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.


### GetAccountCategoryId

`func (o *AccountItemParamsAccountItem) GetAccountCategoryId() int32`

GetAccountCategoryId returns the AccountCategoryId field if non-nil, zero value otherwise.

### GetAccountCategoryIdOk

`func (o *AccountItemParamsAccountItem) GetAccountCategoryIdOk() (*int32, bool)`

GetAccountCategoryIdOk returns a tuple with the AccountCategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountCategoryId

`func (o *AccountItemParamsAccountItem) SetAccountCategoryId(v int32)`

SetAccountCategoryId sets AccountCategoryId field to given value.


### GetCorrespondingIncomeId

`func (o *AccountItemParamsAccountItem) GetCorrespondingIncomeId() int32`

GetCorrespondingIncomeId returns the CorrespondingIncomeId field if non-nil, zero value otherwise.

### GetCorrespondingIncomeIdOk

`func (o *AccountItemParamsAccountItem) GetCorrespondingIncomeIdOk() (*int32, bool)`

GetCorrespondingIncomeIdOk returns a tuple with the CorrespondingIncomeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrespondingIncomeId

`func (o *AccountItemParamsAccountItem) SetCorrespondingIncomeId(v int32)`

SetCorrespondingIncomeId sets CorrespondingIncomeId field to given value.


### GetCorrespondingExpenseId

`func (o *AccountItemParamsAccountItem) GetCorrespondingExpenseId() int32`

GetCorrespondingExpenseId returns the CorrespondingExpenseId field if non-nil, zero value otherwise.

### GetCorrespondingExpenseIdOk

`func (o *AccountItemParamsAccountItem) GetCorrespondingExpenseIdOk() (*int32, bool)`

GetCorrespondingExpenseIdOk returns a tuple with the CorrespondingExpenseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrespondingExpenseId

`func (o *AccountItemParamsAccountItem) SetCorrespondingExpenseId(v int32)`

SetCorrespondingExpenseId sets CorrespondingExpenseId field to given value.


### GetAccumulatedDepAccountItemId

`func (o *AccountItemParamsAccountItem) GetAccumulatedDepAccountItemId() int32`

GetAccumulatedDepAccountItemId returns the AccumulatedDepAccountItemId field if non-nil, zero value otherwise.

### GetAccumulatedDepAccountItemIdOk

`func (o *AccountItemParamsAccountItem) GetAccumulatedDepAccountItemIdOk() (*int32, bool)`

GetAccumulatedDepAccountItemIdOk returns a tuple with the AccumulatedDepAccountItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccumulatedDepAccountItemId

`func (o *AccountItemParamsAccountItem) SetAccumulatedDepAccountItemId(v int32)`

SetAccumulatedDepAccountItemId sets AccumulatedDepAccountItemId field to given value.

### HasAccumulatedDepAccountItemId

`func (o *AccountItemParamsAccountItem) HasAccumulatedDepAccountItemId() bool`

HasAccumulatedDepAccountItemId returns a boolean if a field has been set.

### GetSearchable

`func (o *AccountItemParamsAccountItem) GetSearchable() int32`

GetSearchable returns the Searchable field if non-nil, zero value otherwise.

### GetSearchableOk

`func (o *AccountItemParamsAccountItem) GetSearchableOk() (*int32, bool)`

GetSearchableOk returns a tuple with the Searchable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchable

`func (o *AccountItemParamsAccountItem) SetSearchable(v int32)`

SetSearchable sets Searchable field to given value.

### HasSearchable

`func (o *AccountItemParamsAccountItem) HasSearchable() bool`

HasSearchable returns a boolean if a field has been set.

### GetItems

`func (o *AccountItemParamsAccountItem) GetItems() []AccountItemParamsAccountItemItems`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *AccountItemParamsAccountItem) GetItemsOk() (*[]AccountItemParamsAccountItemItems, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *AccountItemParamsAccountItem) SetItems(v []AccountItemParamsAccountItemItems)`

SetItems sets Items field to given value.

### HasItems

`func (o *AccountItemParamsAccountItem) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetPartners

`func (o *AccountItemParamsAccountItem) GetPartners() []AccountItemParamsAccountItemItems`

GetPartners returns the Partners field if non-nil, zero value otherwise.

### GetPartnersOk

`func (o *AccountItemParamsAccountItem) GetPartnersOk() (*[]AccountItemParamsAccountItemItems, bool)`

GetPartnersOk returns a tuple with the Partners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartners

`func (o *AccountItemParamsAccountItem) SetPartners(v []AccountItemParamsAccountItemItems)`

SetPartners sets Partners field to given value.

### HasPartners

`func (o *AccountItemParamsAccountItem) HasPartners() bool`

HasPartners returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


