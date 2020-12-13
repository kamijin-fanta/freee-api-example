# TrialPlSectionsResponseTrialPlSectionsBalances

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountItemId** | Pointer to **int32** | 勘定科目ID(勘定科目の時のみ含まれる) | [optional] 
**AccountItemName** | Pointer to **string** | 勘定科目名(勘定科目の時のみ含まれる) | [optional] 
**AccountGroupName** | Pointer to **string** | 決算書表示名(account_item_display_type:group指定時に決算書表示名の時のみ含まれる) | [optional] 
**Sections** | Pointer to [**[]TrialPlSectionsResponseTrialPlSectionsSections**](TrialPlSectionsResponseTrialPlSectionsSections.md) | 部門 | [optional] 
**AccountCategoryName** | Pointer to **string** | 勘定科目カテゴリー名 | [optional] 
**TotalLine** | Pointer to **bool** | 合計行(勘定科目カテゴリーの時のみ含まれる) | [optional] 
**HierarchyLevel** | Pointer to **int32** | 階層レベル | [optional] 
**ParentAccountCategoryName** | Pointer to **string** | 上位勘定科目カテゴリー名(勘定科目カテゴリーの時のみ、上層が存在する場合含まれる) | [optional] 
**ClosingBalance** | Pointer to **int32** | 期末残高 | [optional] 

## Methods

### NewTrialPlSectionsResponseTrialPlSectionsBalances

`func NewTrialPlSectionsResponseTrialPlSectionsBalances() *TrialPlSectionsResponseTrialPlSectionsBalances`

NewTrialPlSectionsResponseTrialPlSectionsBalances instantiates a new TrialPlSectionsResponseTrialPlSectionsBalances object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrialPlSectionsResponseTrialPlSectionsBalancesWithDefaults

`func NewTrialPlSectionsResponseTrialPlSectionsBalancesWithDefaults() *TrialPlSectionsResponseTrialPlSectionsBalances`

NewTrialPlSectionsResponseTrialPlSectionsBalancesWithDefaults instantiates a new TrialPlSectionsResponseTrialPlSectionsBalances object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountItemId

`func (o *TrialPlSectionsResponseTrialPlSectionsBalances) GetAccountItemId() int32`

GetAccountItemId returns the AccountItemId field if non-nil, zero value otherwise.

### GetAccountItemIdOk

`func (o *TrialPlSectionsResponseTrialPlSectionsBalances) GetAccountItemIdOk() (*int32, bool)`

GetAccountItemIdOk returns a tuple with the AccountItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountItemId

`func (o *TrialPlSectionsResponseTrialPlSectionsBalances) SetAccountItemId(v int32)`

SetAccountItemId sets AccountItemId field to given value.

### HasAccountItemId

`func (o *TrialPlSectionsResponseTrialPlSectionsBalances) HasAccountItemId() bool`

HasAccountItemId returns a boolean if a field has been set.

### GetAccountItemName

`func (o *TrialPlSectionsResponseTrialPlSectionsBalances) GetAccountItemName() string`

GetAccountItemName returns the AccountItemName field if non-nil, zero value otherwise.

### GetAccountItemNameOk

`func (o *TrialPlSectionsResponseTrialPlSectionsBalances) GetAccountItemNameOk() (*string, bool)`

GetAccountItemNameOk returns a tuple with the AccountItemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountItemName

`func (o *TrialPlSectionsResponseTrialPlSectionsBalances) SetAccountItemName(v string)`

SetAccountItemName sets AccountItemName field to given value.

### HasAccountItemName

`func (o *TrialPlSectionsResponseTrialPlSectionsBalances) HasAccountItemName() bool`

HasAccountItemName returns a boolean if a field has been set.

### GetAccountGroupName

`func (o *TrialPlSectionsResponseTrialPlSectionsBalances) GetAccountGroupName() string`

GetAccountGroupName returns the AccountGroupName field if non-nil, zero value otherwise.

### GetAccountGroupNameOk

`func (o *TrialPlSectionsResponseTrialPlSectionsBalances) GetAccountGroupNameOk() (*string, bool)`

GetAccountGroupNameOk returns a tuple with the AccountGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountGroupName

`func (o *TrialPlSectionsResponseTrialPlSectionsBalances) SetAccountGroupName(v string)`

SetAccountGroupName sets AccountGroupName field to given value.

### HasAccountGroupName

`func (o *TrialPlSectionsResponseTrialPlSectionsBalances) HasAccountGroupName() bool`

HasAccountGroupName returns a boolean if a field has been set.

### GetSections

`func (o *TrialPlSectionsResponseTrialPlSectionsBalances) GetSections() []TrialPlSectionsResponseTrialPlSectionsSections`

GetSections returns the Sections field if non-nil, zero value otherwise.

### GetSectionsOk

`func (o *TrialPlSectionsResponseTrialPlSectionsBalances) GetSectionsOk() (*[]TrialPlSectionsResponseTrialPlSectionsSections, bool)`

GetSectionsOk returns a tuple with the Sections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSections

`func (o *TrialPlSectionsResponseTrialPlSectionsBalances) SetSections(v []TrialPlSectionsResponseTrialPlSectionsSections)`

SetSections sets Sections field to given value.

### HasSections

`func (o *TrialPlSectionsResponseTrialPlSectionsBalances) HasSections() bool`

HasSections returns a boolean if a field has been set.

### GetAccountCategoryName

`func (o *TrialPlSectionsResponseTrialPlSectionsBalances) GetAccountCategoryName() string`

GetAccountCategoryName returns the AccountCategoryName field if non-nil, zero value otherwise.

### GetAccountCategoryNameOk

`func (o *TrialPlSectionsResponseTrialPlSectionsBalances) GetAccountCategoryNameOk() (*string, bool)`

GetAccountCategoryNameOk returns a tuple with the AccountCategoryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountCategoryName

`func (o *TrialPlSectionsResponseTrialPlSectionsBalances) SetAccountCategoryName(v string)`

SetAccountCategoryName sets AccountCategoryName field to given value.

### HasAccountCategoryName

`func (o *TrialPlSectionsResponseTrialPlSectionsBalances) HasAccountCategoryName() bool`

HasAccountCategoryName returns a boolean if a field has been set.

### GetTotalLine

`func (o *TrialPlSectionsResponseTrialPlSectionsBalances) GetTotalLine() bool`

GetTotalLine returns the TotalLine field if non-nil, zero value otherwise.

### GetTotalLineOk

`func (o *TrialPlSectionsResponseTrialPlSectionsBalances) GetTotalLineOk() (*bool, bool)`

GetTotalLineOk returns a tuple with the TotalLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalLine

`func (o *TrialPlSectionsResponseTrialPlSectionsBalances) SetTotalLine(v bool)`

SetTotalLine sets TotalLine field to given value.

### HasTotalLine

`func (o *TrialPlSectionsResponseTrialPlSectionsBalances) HasTotalLine() bool`

HasTotalLine returns a boolean if a field has been set.

### GetHierarchyLevel

`func (o *TrialPlSectionsResponseTrialPlSectionsBalances) GetHierarchyLevel() int32`

GetHierarchyLevel returns the HierarchyLevel field if non-nil, zero value otherwise.

### GetHierarchyLevelOk

`func (o *TrialPlSectionsResponseTrialPlSectionsBalances) GetHierarchyLevelOk() (*int32, bool)`

GetHierarchyLevelOk returns a tuple with the HierarchyLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHierarchyLevel

`func (o *TrialPlSectionsResponseTrialPlSectionsBalances) SetHierarchyLevel(v int32)`

SetHierarchyLevel sets HierarchyLevel field to given value.

### HasHierarchyLevel

`func (o *TrialPlSectionsResponseTrialPlSectionsBalances) HasHierarchyLevel() bool`

HasHierarchyLevel returns a boolean if a field has been set.

### GetParentAccountCategoryName

`func (o *TrialPlSectionsResponseTrialPlSectionsBalances) GetParentAccountCategoryName() string`

GetParentAccountCategoryName returns the ParentAccountCategoryName field if non-nil, zero value otherwise.

### GetParentAccountCategoryNameOk

`func (o *TrialPlSectionsResponseTrialPlSectionsBalances) GetParentAccountCategoryNameOk() (*string, bool)`

GetParentAccountCategoryNameOk returns a tuple with the ParentAccountCategoryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentAccountCategoryName

`func (o *TrialPlSectionsResponseTrialPlSectionsBalances) SetParentAccountCategoryName(v string)`

SetParentAccountCategoryName sets ParentAccountCategoryName field to given value.

### HasParentAccountCategoryName

`func (o *TrialPlSectionsResponseTrialPlSectionsBalances) HasParentAccountCategoryName() bool`

HasParentAccountCategoryName returns a boolean if a field has been set.

### GetClosingBalance

`func (o *TrialPlSectionsResponseTrialPlSectionsBalances) GetClosingBalance() int32`

GetClosingBalance returns the ClosingBalance field if non-nil, zero value otherwise.

### GetClosingBalanceOk

`func (o *TrialPlSectionsResponseTrialPlSectionsBalances) GetClosingBalanceOk() (*int32, bool)`

GetClosingBalanceOk returns a tuple with the ClosingBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosingBalance

`func (o *TrialPlSectionsResponseTrialPlSectionsBalances) SetClosingBalance(v int32)`

SetClosingBalance sets ClosingBalance field to given value.

### HasClosingBalance

`func (o *TrialPlSectionsResponseTrialPlSectionsBalances) HasClosingBalance() bool`

HasClosingBalance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


