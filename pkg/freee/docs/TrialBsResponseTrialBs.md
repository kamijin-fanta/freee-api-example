# TrialBsResponseTrialBs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyId** | **int32** | 事業所ID | 
**FiscalYear** | Pointer to **int32** | 会計年度(条件に指定した時、または条件に月、日条件がない時のみ含まれる） | [optional] 
**StartMonth** | Pointer to **int32** | 発生月で絞込：開始会計月(1-12)(条件に指定した時のみ含まれる） | [optional] 
**EndMonth** | Pointer to **int32** | 発生月で絞込：終了会計月(1-12)(条件に指定した時のみ含まれる） | [optional] 
**StartDate** | Pointer to **string** | 発生日で絞込：開始日(yyyy-mm-dd)(条件に指定した時のみ含まれる） | [optional] 
**EndDate** | Pointer to **string** | 発生日で絞込：終了日(yyyy-mm-dd)(条件に指定した時のみ含まれる） | [optional] 
**AccountItemDisplayType** | Pointer to **string** | 勘定科目の表示（勘定科目: account_item, 決算書表示:group）(条件に指定した時のみ含まれる） | [optional] 
**BreakdownDisplayType** | Pointer to **string** | 内訳の表示（取引先: partner, 品目: item, 勘定科目: account_item）(条件に指定した時のみ含まれる） | [optional] 
**PartnerId** | Pointer to **int32** | 取引先ID(条件に指定した時のみ含まれる） | [optional] 
**PartnerCode** | Pointer to **string** | 取引先コード(条件に指定した時のみ含まれる） | [optional] 
**ItemId** | Pointer to **int32** | 品目ID(条件に指定した時のみ含まれる） | [optional] 
**Adjustment** | Pointer to **string** | 決算整理仕訳のみ: only, 決算整理仕訳以外: without(条件に指定した時のみ含まれる） | [optional] 
**CreatedAt** | Pointer to **string** | 作成日時 | [optional] 
**Balances** | [**[]TrialBsResponseTrialBsBalances**](TrialBsResponseTrialBsBalances.md) |  | 

## Methods

### NewTrialBsResponseTrialBs

`func NewTrialBsResponseTrialBs(companyId int32, balances []TrialBsResponseTrialBsBalances, ) *TrialBsResponseTrialBs`

NewTrialBsResponseTrialBs instantiates a new TrialBsResponseTrialBs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrialBsResponseTrialBsWithDefaults

`func NewTrialBsResponseTrialBsWithDefaults() *TrialBsResponseTrialBs`

NewTrialBsResponseTrialBsWithDefaults instantiates a new TrialBsResponseTrialBs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyId

`func (o *TrialBsResponseTrialBs) GetCompanyId() int32`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *TrialBsResponseTrialBs) GetCompanyIdOk() (*int32, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *TrialBsResponseTrialBs) SetCompanyId(v int32)`

SetCompanyId sets CompanyId field to given value.


### GetFiscalYear

`func (o *TrialBsResponseTrialBs) GetFiscalYear() int32`

GetFiscalYear returns the FiscalYear field if non-nil, zero value otherwise.

### GetFiscalYearOk

`func (o *TrialBsResponseTrialBs) GetFiscalYearOk() (*int32, bool)`

GetFiscalYearOk returns a tuple with the FiscalYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiscalYear

`func (o *TrialBsResponseTrialBs) SetFiscalYear(v int32)`

SetFiscalYear sets FiscalYear field to given value.

### HasFiscalYear

`func (o *TrialBsResponseTrialBs) HasFiscalYear() bool`

HasFiscalYear returns a boolean if a field has been set.

### GetStartMonth

`func (o *TrialBsResponseTrialBs) GetStartMonth() int32`

GetStartMonth returns the StartMonth field if non-nil, zero value otherwise.

### GetStartMonthOk

`func (o *TrialBsResponseTrialBs) GetStartMonthOk() (*int32, bool)`

GetStartMonthOk returns a tuple with the StartMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartMonth

`func (o *TrialBsResponseTrialBs) SetStartMonth(v int32)`

SetStartMonth sets StartMonth field to given value.

### HasStartMonth

`func (o *TrialBsResponseTrialBs) HasStartMonth() bool`

HasStartMonth returns a boolean if a field has been set.

### GetEndMonth

`func (o *TrialBsResponseTrialBs) GetEndMonth() int32`

GetEndMonth returns the EndMonth field if non-nil, zero value otherwise.

### GetEndMonthOk

`func (o *TrialBsResponseTrialBs) GetEndMonthOk() (*int32, bool)`

GetEndMonthOk returns a tuple with the EndMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndMonth

`func (o *TrialBsResponseTrialBs) SetEndMonth(v int32)`

SetEndMonth sets EndMonth field to given value.

### HasEndMonth

`func (o *TrialBsResponseTrialBs) HasEndMonth() bool`

HasEndMonth returns a boolean if a field has been set.

### GetStartDate

`func (o *TrialBsResponseTrialBs) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *TrialBsResponseTrialBs) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *TrialBsResponseTrialBs) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *TrialBsResponseTrialBs) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *TrialBsResponseTrialBs) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *TrialBsResponseTrialBs) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *TrialBsResponseTrialBs) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *TrialBsResponseTrialBs) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetAccountItemDisplayType

`func (o *TrialBsResponseTrialBs) GetAccountItemDisplayType() string`

GetAccountItemDisplayType returns the AccountItemDisplayType field if non-nil, zero value otherwise.

### GetAccountItemDisplayTypeOk

`func (o *TrialBsResponseTrialBs) GetAccountItemDisplayTypeOk() (*string, bool)`

GetAccountItemDisplayTypeOk returns a tuple with the AccountItemDisplayType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountItemDisplayType

`func (o *TrialBsResponseTrialBs) SetAccountItemDisplayType(v string)`

SetAccountItemDisplayType sets AccountItemDisplayType field to given value.

### HasAccountItemDisplayType

`func (o *TrialBsResponseTrialBs) HasAccountItemDisplayType() bool`

HasAccountItemDisplayType returns a boolean if a field has been set.

### GetBreakdownDisplayType

`func (o *TrialBsResponseTrialBs) GetBreakdownDisplayType() string`

GetBreakdownDisplayType returns the BreakdownDisplayType field if non-nil, zero value otherwise.

### GetBreakdownDisplayTypeOk

`func (o *TrialBsResponseTrialBs) GetBreakdownDisplayTypeOk() (*string, bool)`

GetBreakdownDisplayTypeOk returns a tuple with the BreakdownDisplayType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakdownDisplayType

`func (o *TrialBsResponseTrialBs) SetBreakdownDisplayType(v string)`

SetBreakdownDisplayType sets BreakdownDisplayType field to given value.

### HasBreakdownDisplayType

`func (o *TrialBsResponseTrialBs) HasBreakdownDisplayType() bool`

HasBreakdownDisplayType returns a boolean if a field has been set.

### GetPartnerId

`func (o *TrialBsResponseTrialBs) GetPartnerId() int32`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *TrialBsResponseTrialBs) GetPartnerIdOk() (*int32, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *TrialBsResponseTrialBs) SetPartnerId(v int32)`

SetPartnerId sets PartnerId field to given value.

### HasPartnerId

`func (o *TrialBsResponseTrialBs) HasPartnerId() bool`

HasPartnerId returns a boolean if a field has been set.

### GetPartnerCode

`func (o *TrialBsResponseTrialBs) GetPartnerCode() string`

GetPartnerCode returns the PartnerCode field if non-nil, zero value otherwise.

### GetPartnerCodeOk

`func (o *TrialBsResponseTrialBs) GetPartnerCodeOk() (*string, bool)`

GetPartnerCodeOk returns a tuple with the PartnerCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerCode

`func (o *TrialBsResponseTrialBs) SetPartnerCode(v string)`

SetPartnerCode sets PartnerCode field to given value.

### HasPartnerCode

`func (o *TrialBsResponseTrialBs) HasPartnerCode() bool`

HasPartnerCode returns a boolean if a field has been set.

### GetItemId

`func (o *TrialBsResponseTrialBs) GetItemId() int32`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *TrialBsResponseTrialBs) GetItemIdOk() (*int32, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *TrialBsResponseTrialBs) SetItemId(v int32)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *TrialBsResponseTrialBs) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### GetAdjustment

`func (o *TrialBsResponseTrialBs) GetAdjustment() string`

GetAdjustment returns the Adjustment field if non-nil, zero value otherwise.

### GetAdjustmentOk

`func (o *TrialBsResponseTrialBs) GetAdjustmentOk() (*string, bool)`

GetAdjustmentOk returns a tuple with the Adjustment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustment

`func (o *TrialBsResponseTrialBs) SetAdjustment(v string)`

SetAdjustment sets Adjustment field to given value.

### HasAdjustment

`func (o *TrialBsResponseTrialBs) HasAdjustment() bool`

HasAdjustment returns a boolean if a field has been set.

### GetCreatedAt

`func (o *TrialBsResponseTrialBs) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TrialBsResponseTrialBs) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TrialBsResponseTrialBs) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *TrialBsResponseTrialBs) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetBalances

`func (o *TrialBsResponseTrialBs) GetBalances() []TrialBsResponseTrialBsBalances`

GetBalances returns the Balances field if non-nil, zero value otherwise.

### GetBalancesOk

`func (o *TrialBsResponseTrialBs) GetBalancesOk() (*[]TrialBsResponseTrialBsBalances, bool)`

GetBalancesOk returns a tuple with the Balances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalances

`func (o *TrialBsResponseTrialBs) SetBalances(v []TrialBsResponseTrialBsBalances)`

SetBalances sets Balances field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


