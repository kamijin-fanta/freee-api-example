# CompanyParamsFiscalYears

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UseIndustryTemplate** | Pointer to **bool** | 製造業向け機能（true: 使用する、false: 使用しない） | [optional] 
**IndirectWriteOffMethod** | Pointer to **bool** | 固定資産の控除法（true\\: 間接控除法、false\\: 直接控除法）&lt;br&gt; 間接控除法は法人のみ対応しています。  | [optional] 
**IndirectWriteOffMethodType** | Pointer to **bool** | 間接控除時の累計額(法人のみ)（true: 資産分類別、false: 共通）&#39;  | [optional] 
**StartDate** | Pointer to **string** | 期首日 | [optional] 
**EndDate** | Pointer to **string** | 期末日（決算日） | [optional] 
**AccountingPeriod** | Pointer to **int32** | 期 | [optional] 
**DepreciationFraction** | Pointer to **int32** | 減価償却端数処理法(法人のみ)(0: 切り捨て、1: 切り上げ) | [optional] 
**TaxFraction** | Pointer to **int32** | 消費税端数処理方法（0: 切り上げ、1: 切り捨て, 2: 四捨五入） | [optional] 
**ReturnCode** | Pointer to **int32** | 不動産所得使用区分（0: 一般、3: 一般/不動産） ※個人事業主のみ設定可能 | [optional] 

## Methods

### NewCompanyParamsFiscalYears

`func NewCompanyParamsFiscalYears() *CompanyParamsFiscalYears`

NewCompanyParamsFiscalYears instantiates a new CompanyParamsFiscalYears object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyParamsFiscalYearsWithDefaults

`func NewCompanyParamsFiscalYearsWithDefaults() *CompanyParamsFiscalYears`

NewCompanyParamsFiscalYearsWithDefaults instantiates a new CompanyParamsFiscalYears object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUseIndustryTemplate

`func (o *CompanyParamsFiscalYears) GetUseIndustryTemplate() bool`

GetUseIndustryTemplate returns the UseIndustryTemplate field if non-nil, zero value otherwise.

### GetUseIndustryTemplateOk

`func (o *CompanyParamsFiscalYears) GetUseIndustryTemplateOk() (*bool, bool)`

GetUseIndustryTemplateOk returns a tuple with the UseIndustryTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseIndustryTemplate

`func (o *CompanyParamsFiscalYears) SetUseIndustryTemplate(v bool)`

SetUseIndustryTemplate sets UseIndustryTemplate field to given value.

### HasUseIndustryTemplate

`func (o *CompanyParamsFiscalYears) HasUseIndustryTemplate() bool`

HasUseIndustryTemplate returns a boolean if a field has been set.

### GetIndirectWriteOffMethod

`func (o *CompanyParamsFiscalYears) GetIndirectWriteOffMethod() bool`

GetIndirectWriteOffMethod returns the IndirectWriteOffMethod field if non-nil, zero value otherwise.

### GetIndirectWriteOffMethodOk

`func (o *CompanyParamsFiscalYears) GetIndirectWriteOffMethodOk() (*bool, bool)`

GetIndirectWriteOffMethodOk returns a tuple with the IndirectWriteOffMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndirectWriteOffMethod

`func (o *CompanyParamsFiscalYears) SetIndirectWriteOffMethod(v bool)`

SetIndirectWriteOffMethod sets IndirectWriteOffMethod field to given value.

### HasIndirectWriteOffMethod

`func (o *CompanyParamsFiscalYears) HasIndirectWriteOffMethod() bool`

HasIndirectWriteOffMethod returns a boolean if a field has been set.

### GetIndirectWriteOffMethodType

`func (o *CompanyParamsFiscalYears) GetIndirectWriteOffMethodType() bool`

GetIndirectWriteOffMethodType returns the IndirectWriteOffMethodType field if non-nil, zero value otherwise.

### GetIndirectWriteOffMethodTypeOk

`func (o *CompanyParamsFiscalYears) GetIndirectWriteOffMethodTypeOk() (*bool, bool)`

GetIndirectWriteOffMethodTypeOk returns a tuple with the IndirectWriteOffMethodType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndirectWriteOffMethodType

`func (o *CompanyParamsFiscalYears) SetIndirectWriteOffMethodType(v bool)`

SetIndirectWriteOffMethodType sets IndirectWriteOffMethodType field to given value.

### HasIndirectWriteOffMethodType

`func (o *CompanyParamsFiscalYears) HasIndirectWriteOffMethodType() bool`

HasIndirectWriteOffMethodType returns a boolean if a field has been set.

### GetStartDate

`func (o *CompanyParamsFiscalYears) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *CompanyParamsFiscalYears) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *CompanyParamsFiscalYears) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *CompanyParamsFiscalYears) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *CompanyParamsFiscalYears) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *CompanyParamsFiscalYears) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *CompanyParamsFiscalYears) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *CompanyParamsFiscalYears) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetAccountingPeriod

`func (o *CompanyParamsFiscalYears) GetAccountingPeriod() int32`

GetAccountingPeriod returns the AccountingPeriod field if non-nil, zero value otherwise.

### GetAccountingPeriodOk

`func (o *CompanyParamsFiscalYears) GetAccountingPeriodOk() (*int32, bool)`

GetAccountingPeriodOk returns a tuple with the AccountingPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountingPeriod

`func (o *CompanyParamsFiscalYears) SetAccountingPeriod(v int32)`

SetAccountingPeriod sets AccountingPeriod field to given value.

### HasAccountingPeriod

`func (o *CompanyParamsFiscalYears) HasAccountingPeriod() bool`

HasAccountingPeriod returns a boolean if a field has been set.

### GetDepreciationFraction

`func (o *CompanyParamsFiscalYears) GetDepreciationFraction() int32`

GetDepreciationFraction returns the DepreciationFraction field if non-nil, zero value otherwise.

### GetDepreciationFractionOk

`func (o *CompanyParamsFiscalYears) GetDepreciationFractionOk() (*int32, bool)`

GetDepreciationFractionOk returns a tuple with the DepreciationFraction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepreciationFraction

`func (o *CompanyParamsFiscalYears) SetDepreciationFraction(v int32)`

SetDepreciationFraction sets DepreciationFraction field to given value.

### HasDepreciationFraction

`func (o *CompanyParamsFiscalYears) HasDepreciationFraction() bool`

HasDepreciationFraction returns a boolean if a field has been set.

### GetTaxFraction

`func (o *CompanyParamsFiscalYears) GetTaxFraction() int32`

GetTaxFraction returns the TaxFraction field if non-nil, zero value otherwise.

### GetTaxFractionOk

`func (o *CompanyParamsFiscalYears) GetTaxFractionOk() (*int32, bool)`

GetTaxFractionOk returns a tuple with the TaxFraction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxFraction

`func (o *CompanyParamsFiscalYears) SetTaxFraction(v int32)`

SetTaxFraction sets TaxFraction field to given value.

### HasTaxFraction

`func (o *CompanyParamsFiscalYears) HasTaxFraction() bool`

HasTaxFraction returns a boolean if a field has been set.

### GetReturnCode

`func (o *CompanyParamsFiscalYears) GetReturnCode() int32`

GetReturnCode returns the ReturnCode field if non-nil, zero value otherwise.

### GetReturnCodeOk

`func (o *CompanyParamsFiscalYears) GetReturnCodeOk() (*int32, bool)`

GetReturnCodeOk returns a tuple with the ReturnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnCode

`func (o *CompanyParamsFiscalYears) SetReturnCode(v int32)`

SetReturnCode sets ReturnCode field to given value.

### HasReturnCode

`func (o *CompanyParamsFiscalYears) HasReturnCode() bool`

HasReturnCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


