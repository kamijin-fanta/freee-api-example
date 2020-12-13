# TrialBsTwoYearsResponseTrialBsTwoYearsBalances

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountItemId** | Pointer to **int32** | 勘定科目ID(勘定科目の時のみ含まれる) | [optional] 
**AccountItemName** | Pointer to **string** | 勘定科目名(勘定科目の時のみ含まれる) | [optional] 
**AccountGroupName** | Pointer to **string** | 決算書表示名(account_item_display_type:group指定時に決算書表示名の時のみ含まれる) | [optional] 
**Partners** | Pointer to [**[]TrialBsTwoYearsResponseTrialBsTwoYearsPartners**](TrialBsTwoYearsResponseTrialBsTwoYearsPartners.md) | breakdown_display_type:partner, account_item_display_type:account_item指定時のみ含まれる | [optional] 
**Items** | Pointer to [**[]TrialBsTwoYearsResponseTrialBsTwoYearsItems**](TrialBsTwoYearsResponseTrialBsTwoYearsItems.md) | breakdown_display_type:item, account_item_display_type:account_item指定時のみ含まれる | [optional] 
**AccountCategoryName** | Pointer to **string** | 勘定科目カテゴリー名 | [optional] 
**TotalLine** | Pointer to **bool** | 合計行(勘定科目カテゴリーの時のみ含まれる) | [optional] 
**HierarchyLevel** | Pointer to **int32** | 階層レベル | [optional] 
**ParentAccountCategoryName** | Pointer to **string** | 上位勘定科目カテゴリー名(勘定科目カテゴリーの時のみ、上層が存在する場合含まれる) | [optional] 
**LastYearClosingBalance** | Pointer to **int32** | 前年度期末残高 | [optional] 
**ClosingBalance** | Pointer to **int32** | 期末残高 | [optional] 
**YearOnYear** | Pointer to **float32** | 前年比 | [optional] 

## Methods

### NewTrialBsTwoYearsResponseTrialBsTwoYearsBalances

`func NewTrialBsTwoYearsResponseTrialBsTwoYearsBalances() *TrialBsTwoYearsResponseTrialBsTwoYearsBalances`

NewTrialBsTwoYearsResponseTrialBsTwoYearsBalances instantiates a new TrialBsTwoYearsResponseTrialBsTwoYearsBalances object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrialBsTwoYearsResponseTrialBsTwoYearsBalancesWithDefaults

`func NewTrialBsTwoYearsResponseTrialBsTwoYearsBalancesWithDefaults() *TrialBsTwoYearsResponseTrialBsTwoYearsBalances`

NewTrialBsTwoYearsResponseTrialBsTwoYearsBalancesWithDefaults instantiates a new TrialBsTwoYearsResponseTrialBsTwoYearsBalances object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountItemId

`func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalances) GetAccountItemId() int32`

GetAccountItemId returns the AccountItemId field if non-nil, zero value otherwise.

### GetAccountItemIdOk

`func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalances) GetAccountItemIdOk() (*int32, bool)`

GetAccountItemIdOk returns a tuple with the AccountItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountItemId

`func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalances) SetAccountItemId(v int32)`

SetAccountItemId sets AccountItemId field to given value.

### HasAccountItemId

`func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalances) HasAccountItemId() bool`

HasAccountItemId returns a boolean if a field has been set.

### GetAccountItemName

`func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalances) GetAccountItemName() string`

GetAccountItemName returns the AccountItemName field if non-nil, zero value otherwise.

### GetAccountItemNameOk

`func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalances) GetAccountItemNameOk() (*string, bool)`

GetAccountItemNameOk returns a tuple with the AccountItemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountItemName

`func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalances) SetAccountItemName(v string)`

SetAccountItemName sets AccountItemName field to given value.

### HasAccountItemName

`func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalances) HasAccountItemName() bool`

HasAccountItemName returns a boolean if a field has been set.

### GetAccountGroupName

`func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalances) GetAccountGroupName() string`

GetAccountGroupName returns the AccountGroupName field if non-nil, zero value otherwise.

### GetAccountGroupNameOk

`func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalances) GetAccountGroupNameOk() (*string, bool)`

GetAccountGroupNameOk returns a tuple with the AccountGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountGroupName

`func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalances) SetAccountGroupName(v string)`

SetAccountGroupName sets AccountGroupName field to given value.

### HasAccountGroupName

`func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalances) HasAccountGroupName() bool`

HasAccountGroupName returns a boolean if a field has been set.

### GetPartners

`func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalances) GetPartners() []TrialBsTwoYearsResponseTrialBsTwoYearsPartners`

GetPartners returns the Partners field if non-nil, zero value otherwise.

### GetPartnersOk

`func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalances) GetPartnersOk() (*[]TrialBsTwoYearsResponseTrialBsTwoYearsPartners, bool)`

GetPartnersOk returns a tuple with the Partners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartners

`func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalances) SetPartners(v []TrialBsTwoYearsResponseTrialBsTwoYearsPartners)`

SetPartners sets Partners field to given value.

### HasPartners

`func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalances) HasPartners() bool`

HasPartners returns a boolean if a field has been set.

### GetItems

`func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalances) GetItems() []TrialBsTwoYearsResponseTrialBsTwoYearsItems`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalances) GetItemsOk() (*[]TrialBsTwoYearsResponseTrialBsTwoYearsItems, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalances) SetItems(v []TrialBsTwoYearsResponseTrialBsTwoYearsItems)`

SetItems sets Items field to given value.

### HasItems

`func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalances) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetAccountCategoryName

`func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalances) GetAccountCategoryName() string`

GetAccountCategoryName returns the AccountCategoryName field if non-nil, zero value otherwise.

### GetAccountCategoryNameOk

`func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalances) GetAccountCategoryNameOk() (*string, bool)`

GetAccountCategoryNameOk returns a tuple with the AccountCategoryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountCategoryName

`func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalances) SetAccountCategoryName(v string)`

SetAccountCategoryName sets AccountCategoryName field to given value.

### HasAccountCategoryName

`func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalances) HasAccountCategoryName() bool`

HasAccountCategoryName returns a boolean if a field has been set.

### GetTotalLine

`func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalances) GetTotalLine() bool`

GetTotalLine returns the TotalLine field if non-nil, zero value otherwise.

### GetTotalLineOk

`func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalances) GetTotalLineOk() (*bool, bool)`

GetTotalLineOk returns a tuple with the TotalLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalLine

`func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalances) SetTotalLine(v bool)`

SetTotalLine sets TotalLine field to given value.

### HasTotalLine

`func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalances) HasTotalLine() bool`

HasTotalLine returns a boolean if a field has been set.

### GetHierarchyLevel

`func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalances) GetHierarchyLevel() int32`

GetHierarchyLevel returns the HierarchyLevel field if non-nil, zero value otherwise.

### GetHierarchyLevelOk

`func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalances) GetHierarchyLevelOk() (*int32, bool)`

GetHierarchyLevelOk returns a tuple with the HierarchyLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHierarchyLevel

`func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalances) SetHierarchyLevel(v int32)`

SetHierarchyLevel sets HierarchyLevel field to given value.

### HasHierarchyLevel

`func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalances) HasHierarchyLevel() bool`

HasHierarchyLevel returns a boolean if a field has been set.

### GetParentAccountCategoryName

`func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalances) GetParentAccountCategoryName() string`

GetParentAccountCategoryName returns the ParentAccountCategoryName field if non-nil, zero value otherwise.

### GetParentAccountCategoryNameOk

`func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalances) GetParentAccountCategoryNameOk() (*string, bool)`

GetParentAccountCategoryNameOk returns a tuple with the ParentAccountCategoryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentAccountCategoryName

`func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalances) SetParentAccountCategoryName(v string)`

SetParentAccountCategoryName sets ParentAccountCategoryName field to given value.

### HasParentAccountCategoryName

`func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalances) HasParentAccountCategoryName() bool`

HasParentAccountCategoryName returns a boolean if a field has been set.

### GetLastYearClosingBalance

`func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalances) GetLastYearClosingBalance() int32`

GetLastYearClosingBalance returns the LastYearClosingBalance field if non-nil, zero value otherwise.

### GetLastYearClosingBalanceOk

`func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalances) GetLastYearClosingBalanceOk() (*int32, bool)`

GetLastYearClosingBalanceOk returns a tuple with the LastYearClosingBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastYearClosingBalance

`func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalances) SetLastYearClosingBalance(v int32)`

SetLastYearClosingBalance sets LastYearClosingBalance field to given value.

### HasLastYearClosingBalance

`func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalances) HasLastYearClosingBalance() bool`

HasLastYearClosingBalance returns a boolean if a field has been set.

### GetClosingBalance

`func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalances) GetClosingBalance() int32`

GetClosingBalance returns the ClosingBalance field if non-nil, zero value otherwise.

### GetClosingBalanceOk

`func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalances) GetClosingBalanceOk() (*int32, bool)`

GetClosingBalanceOk returns a tuple with the ClosingBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosingBalance

`func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalances) SetClosingBalance(v int32)`

SetClosingBalance sets ClosingBalance field to given value.

### HasClosingBalance

`func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalances) HasClosingBalance() bool`

HasClosingBalance returns a boolean if a field has been set.

### GetYearOnYear

`func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalances) GetYearOnYear() float32`

GetYearOnYear returns the YearOnYear field if non-nil, zero value otherwise.

### GetYearOnYearOk

`func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalances) GetYearOnYearOk() (*float32, bool)`

GetYearOnYearOk returns a tuple with the YearOnYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYearOnYear

`func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalances) SetYearOnYear(v float32)`

SetYearOnYear sets YearOnYear field to given value.

### HasYearOnYear

`func (o *TrialBsTwoYearsResponseTrialBsTwoYearsBalances) HasYearOnYear() bool`

HasYearOnYear returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


