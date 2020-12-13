# TrialPlTwoYearsResponseTrialPlTwoYearsBalances

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountItemId** | Pointer to **int32** | 勘定科目ID(勘定科目の時のみ含まれる) | [optional] 
**AccountItemName** | Pointer to **string** | 勘定科目名(勘定科目の時のみ含まれる) | [optional] 
**AccountGroupName** | Pointer to **string** | 決算書表示名(account_item_display_type:group指定時に決算書表示名の時のみ含まれる) | [optional] 
**Partners** | Pointer to [**[]TrialBsTwoYearsResponseTrialBsTwoYearsPartners**](TrialBsTwoYearsResponseTrialBsTwoYearsPartners.md) | breakdown_display_type:partner, account_item_display_type:account_item指定時のみ含まれる | [optional] 
**Items** | Pointer to [**[]TrialBsTwoYearsResponseTrialBsTwoYearsItems**](TrialBsTwoYearsResponseTrialBsTwoYearsItems.md) | breakdown_display_type:item, account_item_display_type:account_item指定時のみ含まれる | [optional] 
**Sections** | Pointer to [**[]TrialPlTwoYearsResponseTrialPlTwoYearsSections**](TrialPlTwoYearsResponseTrialPlTwoYearsSections.md) | breakdown_display_type:section, account_item_display_type:account_item指定時のみ含まれる | [optional] 
**AccountCategoryName** | Pointer to **string** | 勘定科目カテゴリー名 | [optional] 
**TotalLine** | Pointer to **bool** | 合計行(勘定科目カテゴリーの時のみ含まれる) | [optional] 
**HierarchyLevel** | Pointer to **int32** | 階層レベル | [optional] 
**ParentAccountCategoryName** | Pointer to **string** | 上位勘定科目カテゴリー名(勘定科目カテゴリーの時のみ、上層が存在する場合含まれる) | [optional] 
**LastYearClosingBalance** | Pointer to **int32** | 前年度期末残高 | [optional] 
**ClosingBalance** | Pointer to **int32** | 期末残高 | [optional] 
**YearOnYear** | Pointer to **float32** | 前年比 | [optional] 

## Methods

### NewTrialPlTwoYearsResponseTrialPlTwoYearsBalances

`func NewTrialPlTwoYearsResponseTrialPlTwoYearsBalances() *TrialPlTwoYearsResponseTrialPlTwoYearsBalances`

NewTrialPlTwoYearsResponseTrialPlTwoYearsBalances instantiates a new TrialPlTwoYearsResponseTrialPlTwoYearsBalances object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrialPlTwoYearsResponseTrialPlTwoYearsBalancesWithDefaults

`func NewTrialPlTwoYearsResponseTrialPlTwoYearsBalancesWithDefaults() *TrialPlTwoYearsResponseTrialPlTwoYearsBalances`

NewTrialPlTwoYearsResponseTrialPlTwoYearsBalancesWithDefaults instantiates a new TrialPlTwoYearsResponseTrialPlTwoYearsBalances object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountItemId

`func (o *TrialPlTwoYearsResponseTrialPlTwoYearsBalances) GetAccountItemId() int32`

GetAccountItemId returns the AccountItemId field if non-nil, zero value otherwise.

### GetAccountItemIdOk

`func (o *TrialPlTwoYearsResponseTrialPlTwoYearsBalances) GetAccountItemIdOk() (*int32, bool)`

GetAccountItemIdOk returns a tuple with the AccountItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountItemId

`func (o *TrialPlTwoYearsResponseTrialPlTwoYearsBalances) SetAccountItemId(v int32)`

SetAccountItemId sets AccountItemId field to given value.

### HasAccountItemId

`func (o *TrialPlTwoYearsResponseTrialPlTwoYearsBalances) HasAccountItemId() bool`

HasAccountItemId returns a boolean if a field has been set.

### GetAccountItemName

`func (o *TrialPlTwoYearsResponseTrialPlTwoYearsBalances) GetAccountItemName() string`

GetAccountItemName returns the AccountItemName field if non-nil, zero value otherwise.

### GetAccountItemNameOk

`func (o *TrialPlTwoYearsResponseTrialPlTwoYearsBalances) GetAccountItemNameOk() (*string, bool)`

GetAccountItemNameOk returns a tuple with the AccountItemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountItemName

`func (o *TrialPlTwoYearsResponseTrialPlTwoYearsBalances) SetAccountItemName(v string)`

SetAccountItemName sets AccountItemName field to given value.

### HasAccountItemName

`func (o *TrialPlTwoYearsResponseTrialPlTwoYearsBalances) HasAccountItemName() bool`

HasAccountItemName returns a boolean if a field has been set.

### GetAccountGroupName

`func (o *TrialPlTwoYearsResponseTrialPlTwoYearsBalances) GetAccountGroupName() string`

GetAccountGroupName returns the AccountGroupName field if non-nil, zero value otherwise.

### GetAccountGroupNameOk

`func (o *TrialPlTwoYearsResponseTrialPlTwoYearsBalances) GetAccountGroupNameOk() (*string, bool)`

GetAccountGroupNameOk returns a tuple with the AccountGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountGroupName

`func (o *TrialPlTwoYearsResponseTrialPlTwoYearsBalances) SetAccountGroupName(v string)`

SetAccountGroupName sets AccountGroupName field to given value.

### HasAccountGroupName

`func (o *TrialPlTwoYearsResponseTrialPlTwoYearsBalances) HasAccountGroupName() bool`

HasAccountGroupName returns a boolean if a field has been set.

### GetPartners

`func (o *TrialPlTwoYearsResponseTrialPlTwoYearsBalances) GetPartners() []TrialBsTwoYearsResponseTrialBsTwoYearsPartners`

GetPartners returns the Partners field if non-nil, zero value otherwise.

### GetPartnersOk

`func (o *TrialPlTwoYearsResponseTrialPlTwoYearsBalances) GetPartnersOk() (*[]TrialBsTwoYearsResponseTrialBsTwoYearsPartners, bool)`

GetPartnersOk returns a tuple with the Partners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartners

`func (o *TrialPlTwoYearsResponseTrialPlTwoYearsBalances) SetPartners(v []TrialBsTwoYearsResponseTrialBsTwoYearsPartners)`

SetPartners sets Partners field to given value.

### HasPartners

`func (o *TrialPlTwoYearsResponseTrialPlTwoYearsBalances) HasPartners() bool`

HasPartners returns a boolean if a field has been set.

### GetItems

`func (o *TrialPlTwoYearsResponseTrialPlTwoYearsBalances) GetItems() []TrialBsTwoYearsResponseTrialBsTwoYearsItems`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *TrialPlTwoYearsResponseTrialPlTwoYearsBalances) GetItemsOk() (*[]TrialBsTwoYearsResponseTrialBsTwoYearsItems, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *TrialPlTwoYearsResponseTrialPlTwoYearsBalances) SetItems(v []TrialBsTwoYearsResponseTrialBsTwoYearsItems)`

SetItems sets Items field to given value.

### HasItems

`func (o *TrialPlTwoYearsResponseTrialPlTwoYearsBalances) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetSections

`func (o *TrialPlTwoYearsResponseTrialPlTwoYearsBalances) GetSections() []TrialPlTwoYearsResponseTrialPlTwoYearsSections`

GetSections returns the Sections field if non-nil, zero value otherwise.

### GetSectionsOk

`func (o *TrialPlTwoYearsResponseTrialPlTwoYearsBalances) GetSectionsOk() (*[]TrialPlTwoYearsResponseTrialPlTwoYearsSections, bool)`

GetSectionsOk returns a tuple with the Sections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSections

`func (o *TrialPlTwoYearsResponseTrialPlTwoYearsBalances) SetSections(v []TrialPlTwoYearsResponseTrialPlTwoYearsSections)`

SetSections sets Sections field to given value.

### HasSections

`func (o *TrialPlTwoYearsResponseTrialPlTwoYearsBalances) HasSections() bool`

HasSections returns a boolean if a field has been set.

### GetAccountCategoryName

`func (o *TrialPlTwoYearsResponseTrialPlTwoYearsBalances) GetAccountCategoryName() string`

GetAccountCategoryName returns the AccountCategoryName field if non-nil, zero value otherwise.

### GetAccountCategoryNameOk

`func (o *TrialPlTwoYearsResponseTrialPlTwoYearsBalances) GetAccountCategoryNameOk() (*string, bool)`

GetAccountCategoryNameOk returns a tuple with the AccountCategoryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountCategoryName

`func (o *TrialPlTwoYearsResponseTrialPlTwoYearsBalances) SetAccountCategoryName(v string)`

SetAccountCategoryName sets AccountCategoryName field to given value.

### HasAccountCategoryName

`func (o *TrialPlTwoYearsResponseTrialPlTwoYearsBalances) HasAccountCategoryName() bool`

HasAccountCategoryName returns a boolean if a field has been set.

### GetTotalLine

`func (o *TrialPlTwoYearsResponseTrialPlTwoYearsBalances) GetTotalLine() bool`

GetTotalLine returns the TotalLine field if non-nil, zero value otherwise.

### GetTotalLineOk

`func (o *TrialPlTwoYearsResponseTrialPlTwoYearsBalances) GetTotalLineOk() (*bool, bool)`

GetTotalLineOk returns a tuple with the TotalLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalLine

`func (o *TrialPlTwoYearsResponseTrialPlTwoYearsBalances) SetTotalLine(v bool)`

SetTotalLine sets TotalLine field to given value.

### HasTotalLine

`func (o *TrialPlTwoYearsResponseTrialPlTwoYearsBalances) HasTotalLine() bool`

HasTotalLine returns a boolean if a field has been set.

### GetHierarchyLevel

`func (o *TrialPlTwoYearsResponseTrialPlTwoYearsBalances) GetHierarchyLevel() int32`

GetHierarchyLevel returns the HierarchyLevel field if non-nil, zero value otherwise.

### GetHierarchyLevelOk

`func (o *TrialPlTwoYearsResponseTrialPlTwoYearsBalances) GetHierarchyLevelOk() (*int32, bool)`

GetHierarchyLevelOk returns a tuple with the HierarchyLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHierarchyLevel

`func (o *TrialPlTwoYearsResponseTrialPlTwoYearsBalances) SetHierarchyLevel(v int32)`

SetHierarchyLevel sets HierarchyLevel field to given value.

### HasHierarchyLevel

`func (o *TrialPlTwoYearsResponseTrialPlTwoYearsBalances) HasHierarchyLevel() bool`

HasHierarchyLevel returns a boolean if a field has been set.

### GetParentAccountCategoryName

`func (o *TrialPlTwoYearsResponseTrialPlTwoYearsBalances) GetParentAccountCategoryName() string`

GetParentAccountCategoryName returns the ParentAccountCategoryName field if non-nil, zero value otherwise.

### GetParentAccountCategoryNameOk

`func (o *TrialPlTwoYearsResponseTrialPlTwoYearsBalances) GetParentAccountCategoryNameOk() (*string, bool)`

GetParentAccountCategoryNameOk returns a tuple with the ParentAccountCategoryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentAccountCategoryName

`func (o *TrialPlTwoYearsResponseTrialPlTwoYearsBalances) SetParentAccountCategoryName(v string)`

SetParentAccountCategoryName sets ParentAccountCategoryName field to given value.

### HasParentAccountCategoryName

`func (o *TrialPlTwoYearsResponseTrialPlTwoYearsBalances) HasParentAccountCategoryName() bool`

HasParentAccountCategoryName returns a boolean if a field has been set.

### GetLastYearClosingBalance

`func (o *TrialPlTwoYearsResponseTrialPlTwoYearsBalances) GetLastYearClosingBalance() int32`

GetLastYearClosingBalance returns the LastYearClosingBalance field if non-nil, zero value otherwise.

### GetLastYearClosingBalanceOk

`func (o *TrialPlTwoYearsResponseTrialPlTwoYearsBalances) GetLastYearClosingBalanceOk() (*int32, bool)`

GetLastYearClosingBalanceOk returns a tuple with the LastYearClosingBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastYearClosingBalance

`func (o *TrialPlTwoYearsResponseTrialPlTwoYearsBalances) SetLastYearClosingBalance(v int32)`

SetLastYearClosingBalance sets LastYearClosingBalance field to given value.

### HasLastYearClosingBalance

`func (o *TrialPlTwoYearsResponseTrialPlTwoYearsBalances) HasLastYearClosingBalance() bool`

HasLastYearClosingBalance returns a boolean if a field has been set.

### GetClosingBalance

`func (o *TrialPlTwoYearsResponseTrialPlTwoYearsBalances) GetClosingBalance() int32`

GetClosingBalance returns the ClosingBalance field if non-nil, zero value otherwise.

### GetClosingBalanceOk

`func (o *TrialPlTwoYearsResponseTrialPlTwoYearsBalances) GetClosingBalanceOk() (*int32, bool)`

GetClosingBalanceOk returns a tuple with the ClosingBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosingBalance

`func (o *TrialPlTwoYearsResponseTrialPlTwoYearsBalances) SetClosingBalance(v int32)`

SetClosingBalance sets ClosingBalance field to given value.

### HasClosingBalance

`func (o *TrialPlTwoYearsResponseTrialPlTwoYearsBalances) HasClosingBalance() bool`

HasClosingBalance returns a boolean if a field has been set.

### GetYearOnYear

`func (o *TrialPlTwoYearsResponseTrialPlTwoYearsBalances) GetYearOnYear() float32`

GetYearOnYear returns the YearOnYear field if non-nil, zero value otherwise.

### GetYearOnYearOk

`func (o *TrialPlTwoYearsResponseTrialPlTwoYearsBalances) GetYearOnYearOk() (*float32, bool)`

GetYearOnYearOk returns a tuple with the YearOnYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYearOnYear

`func (o *TrialPlTwoYearsResponseTrialPlTwoYearsBalances) SetYearOnYear(v float32)`

SetYearOnYear sets YearOnYear field to given value.

### HasYearOnYear

`func (o *TrialPlTwoYearsResponseTrialPlTwoYearsBalances) HasYearOnYear() bool`

HasYearOnYear returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


