# TrialBsResponseTrialBsBalances

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountItemId** | Pointer to **int32** | 勘定科目ID(勘定科目の時のみ含まれる) | [optional] 
**AccountItemName** | Pointer to **string** | 勘定科目名(勘定科目の時のみ含まれる) | [optional] 
**AccountGroupName** | Pointer to **string** | 決算書表示名(account_item_display_type:group指定時に決算書表示名の時のみ含まれる) | [optional] 
**Partners** | Pointer to [**[]TrialBsResponseTrialBsPartners**](TrialBsResponseTrialBsPartners.md) | breakdown_display_type:partner, account_item_display_type:account_item指定時のみ含まれる | [optional] 
**Items** | Pointer to [**[]TrialBsResponseTrialBsItems**](TrialBsResponseTrialBsItems.md) | breakdown_display_type:item, account_item_display_type:account_item指定時のみ含まれる | [optional] 
**AccountCategoryName** | Pointer to **string** | 勘定科目カテゴリー名 | [optional] 
**TotalLine** | Pointer to **bool** | 合計行(勘定科目カテゴリーの時のみ含まれる) | [optional] 
**HierarchyLevel** | Pointer to **int32** | 階層レベル | [optional] 
**ParentAccountCategoryName** | Pointer to **string** | 上位勘定科目カテゴリー名(勘定科目カテゴリーの時のみ、上層が存在する場合含まれる) | [optional] 
**OpeningBalance** | Pointer to **int32** | 期首残高 | [optional] 
**DebitAmount** | Pointer to **int32** | 借方金額 | [optional] 
**CreditAmount** | Pointer to **int32** | 貸方金額 | [optional] 
**ClosingBalance** | Pointer to **int32** | 期末残高 | [optional] 
**CompositionRatio** | Pointer to **float32** | 構成比 | [optional] 

## Methods

### NewTrialBsResponseTrialBsBalances

`func NewTrialBsResponseTrialBsBalances() *TrialBsResponseTrialBsBalances`

NewTrialBsResponseTrialBsBalances instantiates a new TrialBsResponseTrialBsBalances object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrialBsResponseTrialBsBalancesWithDefaults

`func NewTrialBsResponseTrialBsBalancesWithDefaults() *TrialBsResponseTrialBsBalances`

NewTrialBsResponseTrialBsBalancesWithDefaults instantiates a new TrialBsResponseTrialBsBalances object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountItemId

`func (o *TrialBsResponseTrialBsBalances) GetAccountItemId() int32`

GetAccountItemId returns the AccountItemId field if non-nil, zero value otherwise.

### GetAccountItemIdOk

`func (o *TrialBsResponseTrialBsBalances) GetAccountItemIdOk() (*int32, bool)`

GetAccountItemIdOk returns a tuple with the AccountItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountItemId

`func (o *TrialBsResponseTrialBsBalances) SetAccountItemId(v int32)`

SetAccountItemId sets AccountItemId field to given value.

### HasAccountItemId

`func (o *TrialBsResponseTrialBsBalances) HasAccountItemId() bool`

HasAccountItemId returns a boolean if a field has been set.

### GetAccountItemName

`func (o *TrialBsResponseTrialBsBalances) GetAccountItemName() string`

GetAccountItemName returns the AccountItemName field if non-nil, zero value otherwise.

### GetAccountItemNameOk

`func (o *TrialBsResponseTrialBsBalances) GetAccountItemNameOk() (*string, bool)`

GetAccountItemNameOk returns a tuple with the AccountItemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountItemName

`func (o *TrialBsResponseTrialBsBalances) SetAccountItemName(v string)`

SetAccountItemName sets AccountItemName field to given value.

### HasAccountItemName

`func (o *TrialBsResponseTrialBsBalances) HasAccountItemName() bool`

HasAccountItemName returns a boolean if a field has been set.

### GetAccountGroupName

`func (o *TrialBsResponseTrialBsBalances) GetAccountGroupName() string`

GetAccountGroupName returns the AccountGroupName field if non-nil, zero value otherwise.

### GetAccountGroupNameOk

`func (o *TrialBsResponseTrialBsBalances) GetAccountGroupNameOk() (*string, bool)`

GetAccountGroupNameOk returns a tuple with the AccountGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountGroupName

`func (o *TrialBsResponseTrialBsBalances) SetAccountGroupName(v string)`

SetAccountGroupName sets AccountGroupName field to given value.

### HasAccountGroupName

`func (o *TrialBsResponseTrialBsBalances) HasAccountGroupName() bool`

HasAccountGroupName returns a boolean if a field has been set.

### GetPartners

`func (o *TrialBsResponseTrialBsBalances) GetPartners() []TrialBsResponseTrialBsPartners`

GetPartners returns the Partners field if non-nil, zero value otherwise.

### GetPartnersOk

`func (o *TrialBsResponseTrialBsBalances) GetPartnersOk() (*[]TrialBsResponseTrialBsPartners, bool)`

GetPartnersOk returns a tuple with the Partners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartners

`func (o *TrialBsResponseTrialBsBalances) SetPartners(v []TrialBsResponseTrialBsPartners)`

SetPartners sets Partners field to given value.

### HasPartners

`func (o *TrialBsResponseTrialBsBalances) HasPartners() bool`

HasPartners returns a boolean if a field has been set.

### GetItems

`func (o *TrialBsResponseTrialBsBalances) GetItems() []TrialBsResponseTrialBsItems`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *TrialBsResponseTrialBsBalances) GetItemsOk() (*[]TrialBsResponseTrialBsItems, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *TrialBsResponseTrialBsBalances) SetItems(v []TrialBsResponseTrialBsItems)`

SetItems sets Items field to given value.

### HasItems

`func (o *TrialBsResponseTrialBsBalances) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetAccountCategoryName

`func (o *TrialBsResponseTrialBsBalances) GetAccountCategoryName() string`

GetAccountCategoryName returns the AccountCategoryName field if non-nil, zero value otherwise.

### GetAccountCategoryNameOk

`func (o *TrialBsResponseTrialBsBalances) GetAccountCategoryNameOk() (*string, bool)`

GetAccountCategoryNameOk returns a tuple with the AccountCategoryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountCategoryName

`func (o *TrialBsResponseTrialBsBalances) SetAccountCategoryName(v string)`

SetAccountCategoryName sets AccountCategoryName field to given value.

### HasAccountCategoryName

`func (o *TrialBsResponseTrialBsBalances) HasAccountCategoryName() bool`

HasAccountCategoryName returns a boolean if a field has been set.

### GetTotalLine

`func (o *TrialBsResponseTrialBsBalances) GetTotalLine() bool`

GetTotalLine returns the TotalLine field if non-nil, zero value otherwise.

### GetTotalLineOk

`func (o *TrialBsResponseTrialBsBalances) GetTotalLineOk() (*bool, bool)`

GetTotalLineOk returns a tuple with the TotalLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalLine

`func (o *TrialBsResponseTrialBsBalances) SetTotalLine(v bool)`

SetTotalLine sets TotalLine field to given value.

### HasTotalLine

`func (o *TrialBsResponseTrialBsBalances) HasTotalLine() bool`

HasTotalLine returns a boolean if a field has been set.

### GetHierarchyLevel

`func (o *TrialBsResponseTrialBsBalances) GetHierarchyLevel() int32`

GetHierarchyLevel returns the HierarchyLevel field if non-nil, zero value otherwise.

### GetHierarchyLevelOk

`func (o *TrialBsResponseTrialBsBalances) GetHierarchyLevelOk() (*int32, bool)`

GetHierarchyLevelOk returns a tuple with the HierarchyLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHierarchyLevel

`func (o *TrialBsResponseTrialBsBalances) SetHierarchyLevel(v int32)`

SetHierarchyLevel sets HierarchyLevel field to given value.

### HasHierarchyLevel

`func (o *TrialBsResponseTrialBsBalances) HasHierarchyLevel() bool`

HasHierarchyLevel returns a boolean if a field has been set.

### GetParentAccountCategoryName

`func (o *TrialBsResponseTrialBsBalances) GetParentAccountCategoryName() string`

GetParentAccountCategoryName returns the ParentAccountCategoryName field if non-nil, zero value otherwise.

### GetParentAccountCategoryNameOk

`func (o *TrialBsResponseTrialBsBalances) GetParentAccountCategoryNameOk() (*string, bool)`

GetParentAccountCategoryNameOk returns a tuple with the ParentAccountCategoryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentAccountCategoryName

`func (o *TrialBsResponseTrialBsBalances) SetParentAccountCategoryName(v string)`

SetParentAccountCategoryName sets ParentAccountCategoryName field to given value.

### HasParentAccountCategoryName

`func (o *TrialBsResponseTrialBsBalances) HasParentAccountCategoryName() bool`

HasParentAccountCategoryName returns a boolean if a field has been set.

### GetOpeningBalance

`func (o *TrialBsResponseTrialBsBalances) GetOpeningBalance() int32`

GetOpeningBalance returns the OpeningBalance field if non-nil, zero value otherwise.

### GetOpeningBalanceOk

`func (o *TrialBsResponseTrialBsBalances) GetOpeningBalanceOk() (*int32, bool)`

GetOpeningBalanceOk returns a tuple with the OpeningBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpeningBalance

`func (o *TrialBsResponseTrialBsBalances) SetOpeningBalance(v int32)`

SetOpeningBalance sets OpeningBalance field to given value.

### HasOpeningBalance

`func (o *TrialBsResponseTrialBsBalances) HasOpeningBalance() bool`

HasOpeningBalance returns a boolean if a field has been set.

### GetDebitAmount

`func (o *TrialBsResponseTrialBsBalances) GetDebitAmount() int32`

GetDebitAmount returns the DebitAmount field if non-nil, zero value otherwise.

### GetDebitAmountOk

`func (o *TrialBsResponseTrialBsBalances) GetDebitAmountOk() (*int32, bool)`

GetDebitAmountOk returns a tuple with the DebitAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebitAmount

`func (o *TrialBsResponseTrialBsBalances) SetDebitAmount(v int32)`

SetDebitAmount sets DebitAmount field to given value.

### HasDebitAmount

`func (o *TrialBsResponseTrialBsBalances) HasDebitAmount() bool`

HasDebitAmount returns a boolean if a field has been set.

### GetCreditAmount

`func (o *TrialBsResponseTrialBsBalances) GetCreditAmount() int32`

GetCreditAmount returns the CreditAmount field if non-nil, zero value otherwise.

### GetCreditAmountOk

`func (o *TrialBsResponseTrialBsBalances) GetCreditAmountOk() (*int32, bool)`

GetCreditAmountOk returns a tuple with the CreditAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditAmount

`func (o *TrialBsResponseTrialBsBalances) SetCreditAmount(v int32)`

SetCreditAmount sets CreditAmount field to given value.

### HasCreditAmount

`func (o *TrialBsResponseTrialBsBalances) HasCreditAmount() bool`

HasCreditAmount returns a boolean if a field has been set.

### GetClosingBalance

`func (o *TrialBsResponseTrialBsBalances) GetClosingBalance() int32`

GetClosingBalance returns the ClosingBalance field if non-nil, zero value otherwise.

### GetClosingBalanceOk

`func (o *TrialBsResponseTrialBsBalances) GetClosingBalanceOk() (*int32, bool)`

GetClosingBalanceOk returns a tuple with the ClosingBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosingBalance

`func (o *TrialBsResponseTrialBsBalances) SetClosingBalance(v int32)`

SetClosingBalance sets ClosingBalance field to given value.

### HasClosingBalance

`func (o *TrialBsResponseTrialBsBalances) HasClosingBalance() bool`

HasClosingBalance returns a boolean if a field has been set.

### GetCompositionRatio

`func (o *TrialBsResponseTrialBsBalances) GetCompositionRatio() float32`

GetCompositionRatio returns the CompositionRatio field if non-nil, zero value otherwise.

### GetCompositionRatioOk

`func (o *TrialBsResponseTrialBsBalances) GetCompositionRatioOk() (*float32, bool)`

GetCompositionRatioOk returns a tuple with the CompositionRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompositionRatio

`func (o *TrialBsResponseTrialBsBalances) SetCompositionRatio(v float32)`

SetCompositionRatio sets CompositionRatio field to given value.

### HasCompositionRatio

`func (o *TrialBsResponseTrialBsBalances) HasCompositionRatio() bool`

HasCompositionRatio returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


