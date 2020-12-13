# TrialPlResponseTrialPlBalances

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountItemId** | Pointer to **int32** | 勘定科目ID(勘定科目の時のみ含まれる) | [optional] 
**AccountItemName** | Pointer to **string** | 勘定科目名(勘定科目の時のみ含まれる) | [optional] 
**AccountGroupName** | Pointer to **string** | 決算書表示名(account_item_display_type:group指定時に決算書表示名の時のみ含まれる) | [optional] 
**Partners** | Pointer to [**[]TrialBsResponseTrialBsPartners**](TrialBsResponseTrialBsPartners.md) | breakdown_display_type:partner, account_item_display_type:account_item指定時のみ含まれる | [optional] 
**Items** | Pointer to [**[]TrialBsResponseTrialBsItems**](TrialBsResponseTrialBsItems.md) | breakdown_display_type:item, account_item_display_type:account_item指定時のみ含まれる | [optional] 
**Sections** | Pointer to [**[]TrialPlResponseTrialPlSections**](TrialPlResponseTrialPlSections.md) | breakdown_display_type:section, account_item_display_type:account_item指定時のみ含まれる | [optional] 
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

### NewTrialPlResponseTrialPlBalances

`func NewTrialPlResponseTrialPlBalances() *TrialPlResponseTrialPlBalances`

NewTrialPlResponseTrialPlBalances instantiates a new TrialPlResponseTrialPlBalances object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrialPlResponseTrialPlBalancesWithDefaults

`func NewTrialPlResponseTrialPlBalancesWithDefaults() *TrialPlResponseTrialPlBalances`

NewTrialPlResponseTrialPlBalancesWithDefaults instantiates a new TrialPlResponseTrialPlBalances object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountItemId

`func (o *TrialPlResponseTrialPlBalances) GetAccountItemId() int32`

GetAccountItemId returns the AccountItemId field if non-nil, zero value otherwise.

### GetAccountItemIdOk

`func (o *TrialPlResponseTrialPlBalances) GetAccountItemIdOk() (*int32, bool)`

GetAccountItemIdOk returns a tuple with the AccountItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountItemId

`func (o *TrialPlResponseTrialPlBalances) SetAccountItemId(v int32)`

SetAccountItemId sets AccountItemId field to given value.

### HasAccountItemId

`func (o *TrialPlResponseTrialPlBalances) HasAccountItemId() bool`

HasAccountItemId returns a boolean if a field has been set.

### GetAccountItemName

`func (o *TrialPlResponseTrialPlBalances) GetAccountItemName() string`

GetAccountItemName returns the AccountItemName field if non-nil, zero value otherwise.

### GetAccountItemNameOk

`func (o *TrialPlResponseTrialPlBalances) GetAccountItemNameOk() (*string, bool)`

GetAccountItemNameOk returns a tuple with the AccountItemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountItemName

`func (o *TrialPlResponseTrialPlBalances) SetAccountItemName(v string)`

SetAccountItemName sets AccountItemName field to given value.

### HasAccountItemName

`func (o *TrialPlResponseTrialPlBalances) HasAccountItemName() bool`

HasAccountItemName returns a boolean if a field has been set.

### GetAccountGroupName

`func (o *TrialPlResponseTrialPlBalances) GetAccountGroupName() string`

GetAccountGroupName returns the AccountGroupName field if non-nil, zero value otherwise.

### GetAccountGroupNameOk

`func (o *TrialPlResponseTrialPlBalances) GetAccountGroupNameOk() (*string, bool)`

GetAccountGroupNameOk returns a tuple with the AccountGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountGroupName

`func (o *TrialPlResponseTrialPlBalances) SetAccountGroupName(v string)`

SetAccountGroupName sets AccountGroupName field to given value.

### HasAccountGroupName

`func (o *TrialPlResponseTrialPlBalances) HasAccountGroupName() bool`

HasAccountGroupName returns a boolean if a field has been set.

### GetPartners

`func (o *TrialPlResponseTrialPlBalances) GetPartners() []TrialBsResponseTrialBsPartners`

GetPartners returns the Partners field if non-nil, zero value otherwise.

### GetPartnersOk

`func (o *TrialPlResponseTrialPlBalances) GetPartnersOk() (*[]TrialBsResponseTrialBsPartners, bool)`

GetPartnersOk returns a tuple with the Partners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartners

`func (o *TrialPlResponseTrialPlBalances) SetPartners(v []TrialBsResponseTrialBsPartners)`

SetPartners sets Partners field to given value.

### HasPartners

`func (o *TrialPlResponseTrialPlBalances) HasPartners() bool`

HasPartners returns a boolean if a field has been set.

### GetItems

`func (o *TrialPlResponseTrialPlBalances) GetItems() []TrialBsResponseTrialBsItems`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *TrialPlResponseTrialPlBalances) GetItemsOk() (*[]TrialBsResponseTrialBsItems, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *TrialPlResponseTrialPlBalances) SetItems(v []TrialBsResponseTrialBsItems)`

SetItems sets Items field to given value.

### HasItems

`func (o *TrialPlResponseTrialPlBalances) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetSections

`func (o *TrialPlResponseTrialPlBalances) GetSections() []TrialPlResponseTrialPlSections`

GetSections returns the Sections field if non-nil, zero value otherwise.

### GetSectionsOk

`func (o *TrialPlResponseTrialPlBalances) GetSectionsOk() (*[]TrialPlResponseTrialPlSections, bool)`

GetSectionsOk returns a tuple with the Sections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSections

`func (o *TrialPlResponseTrialPlBalances) SetSections(v []TrialPlResponseTrialPlSections)`

SetSections sets Sections field to given value.

### HasSections

`func (o *TrialPlResponseTrialPlBalances) HasSections() bool`

HasSections returns a boolean if a field has been set.

### GetAccountCategoryName

`func (o *TrialPlResponseTrialPlBalances) GetAccountCategoryName() string`

GetAccountCategoryName returns the AccountCategoryName field if non-nil, zero value otherwise.

### GetAccountCategoryNameOk

`func (o *TrialPlResponseTrialPlBalances) GetAccountCategoryNameOk() (*string, bool)`

GetAccountCategoryNameOk returns a tuple with the AccountCategoryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountCategoryName

`func (o *TrialPlResponseTrialPlBalances) SetAccountCategoryName(v string)`

SetAccountCategoryName sets AccountCategoryName field to given value.

### HasAccountCategoryName

`func (o *TrialPlResponseTrialPlBalances) HasAccountCategoryName() bool`

HasAccountCategoryName returns a boolean if a field has been set.

### GetTotalLine

`func (o *TrialPlResponseTrialPlBalances) GetTotalLine() bool`

GetTotalLine returns the TotalLine field if non-nil, zero value otherwise.

### GetTotalLineOk

`func (o *TrialPlResponseTrialPlBalances) GetTotalLineOk() (*bool, bool)`

GetTotalLineOk returns a tuple with the TotalLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalLine

`func (o *TrialPlResponseTrialPlBalances) SetTotalLine(v bool)`

SetTotalLine sets TotalLine field to given value.

### HasTotalLine

`func (o *TrialPlResponseTrialPlBalances) HasTotalLine() bool`

HasTotalLine returns a boolean if a field has been set.

### GetHierarchyLevel

`func (o *TrialPlResponseTrialPlBalances) GetHierarchyLevel() int32`

GetHierarchyLevel returns the HierarchyLevel field if non-nil, zero value otherwise.

### GetHierarchyLevelOk

`func (o *TrialPlResponseTrialPlBalances) GetHierarchyLevelOk() (*int32, bool)`

GetHierarchyLevelOk returns a tuple with the HierarchyLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHierarchyLevel

`func (o *TrialPlResponseTrialPlBalances) SetHierarchyLevel(v int32)`

SetHierarchyLevel sets HierarchyLevel field to given value.

### HasHierarchyLevel

`func (o *TrialPlResponseTrialPlBalances) HasHierarchyLevel() bool`

HasHierarchyLevel returns a boolean if a field has been set.

### GetParentAccountCategoryName

`func (o *TrialPlResponseTrialPlBalances) GetParentAccountCategoryName() string`

GetParentAccountCategoryName returns the ParentAccountCategoryName field if non-nil, zero value otherwise.

### GetParentAccountCategoryNameOk

`func (o *TrialPlResponseTrialPlBalances) GetParentAccountCategoryNameOk() (*string, bool)`

GetParentAccountCategoryNameOk returns a tuple with the ParentAccountCategoryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentAccountCategoryName

`func (o *TrialPlResponseTrialPlBalances) SetParentAccountCategoryName(v string)`

SetParentAccountCategoryName sets ParentAccountCategoryName field to given value.

### HasParentAccountCategoryName

`func (o *TrialPlResponseTrialPlBalances) HasParentAccountCategoryName() bool`

HasParentAccountCategoryName returns a boolean if a field has been set.

### GetOpeningBalance

`func (o *TrialPlResponseTrialPlBalances) GetOpeningBalance() int32`

GetOpeningBalance returns the OpeningBalance field if non-nil, zero value otherwise.

### GetOpeningBalanceOk

`func (o *TrialPlResponseTrialPlBalances) GetOpeningBalanceOk() (*int32, bool)`

GetOpeningBalanceOk returns a tuple with the OpeningBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpeningBalance

`func (o *TrialPlResponseTrialPlBalances) SetOpeningBalance(v int32)`

SetOpeningBalance sets OpeningBalance field to given value.

### HasOpeningBalance

`func (o *TrialPlResponseTrialPlBalances) HasOpeningBalance() bool`

HasOpeningBalance returns a boolean if a field has been set.

### GetDebitAmount

`func (o *TrialPlResponseTrialPlBalances) GetDebitAmount() int32`

GetDebitAmount returns the DebitAmount field if non-nil, zero value otherwise.

### GetDebitAmountOk

`func (o *TrialPlResponseTrialPlBalances) GetDebitAmountOk() (*int32, bool)`

GetDebitAmountOk returns a tuple with the DebitAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebitAmount

`func (o *TrialPlResponseTrialPlBalances) SetDebitAmount(v int32)`

SetDebitAmount sets DebitAmount field to given value.

### HasDebitAmount

`func (o *TrialPlResponseTrialPlBalances) HasDebitAmount() bool`

HasDebitAmount returns a boolean if a field has been set.

### GetCreditAmount

`func (o *TrialPlResponseTrialPlBalances) GetCreditAmount() int32`

GetCreditAmount returns the CreditAmount field if non-nil, zero value otherwise.

### GetCreditAmountOk

`func (o *TrialPlResponseTrialPlBalances) GetCreditAmountOk() (*int32, bool)`

GetCreditAmountOk returns a tuple with the CreditAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditAmount

`func (o *TrialPlResponseTrialPlBalances) SetCreditAmount(v int32)`

SetCreditAmount sets CreditAmount field to given value.

### HasCreditAmount

`func (o *TrialPlResponseTrialPlBalances) HasCreditAmount() bool`

HasCreditAmount returns a boolean if a field has been set.

### GetClosingBalance

`func (o *TrialPlResponseTrialPlBalances) GetClosingBalance() int32`

GetClosingBalance returns the ClosingBalance field if non-nil, zero value otherwise.

### GetClosingBalanceOk

`func (o *TrialPlResponseTrialPlBalances) GetClosingBalanceOk() (*int32, bool)`

GetClosingBalanceOk returns a tuple with the ClosingBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosingBalance

`func (o *TrialPlResponseTrialPlBalances) SetClosingBalance(v int32)`

SetClosingBalance sets ClosingBalance field to given value.

### HasClosingBalance

`func (o *TrialPlResponseTrialPlBalances) HasClosingBalance() bool`

HasClosingBalance returns a boolean if a field has been set.

### GetCompositionRatio

`func (o *TrialPlResponseTrialPlBalances) GetCompositionRatio() float32`

GetCompositionRatio returns the CompositionRatio field if non-nil, zero value otherwise.

### GetCompositionRatioOk

`func (o *TrialPlResponseTrialPlBalances) GetCompositionRatioOk() (*float32, bool)`

GetCompositionRatioOk returns a tuple with the CompositionRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompositionRatio

`func (o *TrialPlResponseTrialPlBalances) SetCompositionRatio(v float32)`

SetCompositionRatio sets CompositionRatio field to given value.

### HasCompositionRatio

`func (o *TrialPlResponseTrialPlBalances) HasCompositionRatio() bool`

HasCompositionRatio returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


