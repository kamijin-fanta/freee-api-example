# CompanyUpdateResponseCompanyFiscalYears

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UseIndustryTemplate** | **bool** | 製造業向け機能（true: 使用する、false: 使用しない） | 
**IndirectWriteOffMethod** | **bool** | 固定資産の控除法（true: 間接控除法、false: 直接控除法） | 
**IndirectWriteOffMethodType** | **bool** | 間接控除時の累計額（true: 資産分類別、false: 共通） | 
**StartDate** | Pointer to **string** | 期首日 | [optional] 
**EndDate** | Pointer to **string** | 期末日 | [optional] 
**AccountingPeriod** | Pointer to **int32** | 期 | [optional] 
**DepreciationRecordMethod** | **int32** | 月次償却（0: しない、1: する） | 
**TaxMethod** | **int32** | 課税区分（0: 免税、1: 簡易課税、2: 本則課税（個別対応方式）、3: 本則課税（一括比例配分方式）、4: 本則課税（全額控除）） | 
**SalesTaxBusinessCode** | **int32** | 簡易課税用事業区分（0: 第一種：卸売業、1: 第二種：小売業、2: 第三種：農林水産業、工業、建設業、製造業など、3: 第四種：飲食店業など、4: 第五種：金融・保険業、運輸通信業、サービス業など、5: 第六種：不動産業など | 
**TaxFraction** | **int32** | 消費税端数処理方法（0: 切り捨て、1: 切り上げ、2: 四捨五入） | 
**TaxAccountMethod** | **int32** | 消費税経理処理方法（0: 税込経理、1: 旧税抜経理、2: 税抜経理） | 
**ReturnCode** | **int32** | 不動産所得使用区分（0: 一般、3: 一般/不動産） ※個人事業主のみ設定可能 | 

## Methods

### NewCompanyUpdateResponseCompanyFiscalYears

`func NewCompanyUpdateResponseCompanyFiscalYears(useIndustryTemplate bool, indirectWriteOffMethod bool, indirectWriteOffMethodType bool, depreciationRecordMethod int32, taxMethod int32, salesTaxBusinessCode int32, taxFraction int32, taxAccountMethod int32, returnCode int32, ) *CompanyUpdateResponseCompanyFiscalYears`

NewCompanyUpdateResponseCompanyFiscalYears instantiates a new CompanyUpdateResponseCompanyFiscalYears object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyUpdateResponseCompanyFiscalYearsWithDefaults

`func NewCompanyUpdateResponseCompanyFiscalYearsWithDefaults() *CompanyUpdateResponseCompanyFiscalYears`

NewCompanyUpdateResponseCompanyFiscalYearsWithDefaults instantiates a new CompanyUpdateResponseCompanyFiscalYears object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUseIndustryTemplate

`func (o *CompanyUpdateResponseCompanyFiscalYears) GetUseIndustryTemplate() bool`

GetUseIndustryTemplate returns the UseIndustryTemplate field if non-nil, zero value otherwise.

### GetUseIndustryTemplateOk

`func (o *CompanyUpdateResponseCompanyFiscalYears) GetUseIndustryTemplateOk() (*bool, bool)`

GetUseIndustryTemplateOk returns a tuple with the UseIndustryTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseIndustryTemplate

`func (o *CompanyUpdateResponseCompanyFiscalYears) SetUseIndustryTemplate(v bool)`

SetUseIndustryTemplate sets UseIndustryTemplate field to given value.


### GetIndirectWriteOffMethod

`func (o *CompanyUpdateResponseCompanyFiscalYears) GetIndirectWriteOffMethod() bool`

GetIndirectWriteOffMethod returns the IndirectWriteOffMethod field if non-nil, zero value otherwise.

### GetIndirectWriteOffMethodOk

`func (o *CompanyUpdateResponseCompanyFiscalYears) GetIndirectWriteOffMethodOk() (*bool, bool)`

GetIndirectWriteOffMethodOk returns a tuple with the IndirectWriteOffMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndirectWriteOffMethod

`func (o *CompanyUpdateResponseCompanyFiscalYears) SetIndirectWriteOffMethod(v bool)`

SetIndirectWriteOffMethod sets IndirectWriteOffMethod field to given value.


### GetIndirectWriteOffMethodType

`func (o *CompanyUpdateResponseCompanyFiscalYears) GetIndirectWriteOffMethodType() bool`

GetIndirectWriteOffMethodType returns the IndirectWriteOffMethodType field if non-nil, zero value otherwise.

### GetIndirectWriteOffMethodTypeOk

`func (o *CompanyUpdateResponseCompanyFiscalYears) GetIndirectWriteOffMethodTypeOk() (*bool, bool)`

GetIndirectWriteOffMethodTypeOk returns a tuple with the IndirectWriteOffMethodType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndirectWriteOffMethodType

`func (o *CompanyUpdateResponseCompanyFiscalYears) SetIndirectWriteOffMethodType(v bool)`

SetIndirectWriteOffMethodType sets IndirectWriteOffMethodType field to given value.


### GetStartDate

`func (o *CompanyUpdateResponseCompanyFiscalYears) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *CompanyUpdateResponseCompanyFiscalYears) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *CompanyUpdateResponseCompanyFiscalYears) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *CompanyUpdateResponseCompanyFiscalYears) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *CompanyUpdateResponseCompanyFiscalYears) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *CompanyUpdateResponseCompanyFiscalYears) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *CompanyUpdateResponseCompanyFiscalYears) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *CompanyUpdateResponseCompanyFiscalYears) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetAccountingPeriod

`func (o *CompanyUpdateResponseCompanyFiscalYears) GetAccountingPeriod() int32`

GetAccountingPeriod returns the AccountingPeriod field if non-nil, zero value otherwise.

### GetAccountingPeriodOk

`func (o *CompanyUpdateResponseCompanyFiscalYears) GetAccountingPeriodOk() (*int32, bool)`

GetAccountingPeriodOk returns a tuple with the AccountingPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountingPeriod

`func (o *CompanyUpdateResponseCompanyFiscalYears) SetAccountingPeriod(v int32)`

SetAccountingPeriod sets AccountingPeriod field to given value.

### HasAccountingPeriod

`func (o *CompanyUpdateResponseCompanyFiscalYears) HasAccountingPeriod() bool`

HasAccountingPeriod returns a boolean if a field has been set.

### GetDepreciationRecordMethod

`func (o *CompanyUpdateResponseCompanyFiscalYears) GetDepreciationRecordMethod() int32`

GetDepreciationRecordMethod returns the DepreciationRecordMethod field if non-nil, zero value otherwise.

### GetDepreciationRecordMethodOk

`func (o *CompanyUpdateResponseCompanyFiscalYears) GetDepreciationRecordMethodOk() (*int32, bool)`

GetDepreciationRecordMethodOk returns a tuple with the DepreciationRecordMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepreciationRecordMethod

`func (o *CompanyUpdateResponseCompanyFiscalYears) SetDepreciationRecordMethod(v int32)`

SetDepreciationRecordMethod sets DepreciationRecordMethod field to given value.


### GetTaxMethod

`func (o *CompanyUpdateResponseCompanyFiscalYears) GetTaxMethod() int32`

GetTaxMethod returns the TaxMethod field if non-nil, zero value otherwise.

### GetTaxMethodOk

`func (o *CompanyUpdateResponseCompanyFiscalYears) GetTaxMethodOk() (*int32, bool)`

GetTaxMethodOk returns a tuple with the TaxMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxMethod

`func (o *CompanyUpdateResponseCompanyFiscalYears) SetTaxMethod(v int32)`

SetTaxMethod sets TaxMethod field to given value.


### GetSalesTaxBusinessCode

`func (o *CompanyUpdateResponseCompanyFiscalYears) GetSalesTaxBusinessCode() int32`

GetSalesTaxBusinessCode returns the SalesTaxBusinessCode field if non-nil, zero value otherwise.

### GetSalesTaxBusinessCodeOk

`func (o *CompanyUpdateResponseCompanyFiscalYears) GetSalesTaxBusinessCodeOk() (*int32, bool)`

GetSalesTaxBusinessCodeOk returns a tuple with the SalesTaxBusinessCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTaxBusinessCode

`func (o *CompanyUpdateResponseCompanyFiscalYears) SetSalesTaxBusinessCode(v int32)`

SetSalesTaxBusinessCode sets SalesTaxBusinessCode field to given value.


### GetTaxFraction

`func (o *CompanyUpdateResponseCompanyFiscalYears) GetTaxFraction() int32`

GetTaxFraction returns the TaxFraction field if non-nil, zero value otherwise.

### GetTaxFractionOk

`func (o *CompanyUpdateResponseCompanyFiscalYears) GetTaxFractionOk() (*int32, bool)`

GetTaxFractionOk returns a tuple with the TaxFraction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxFraction

`func (o *CompanyUpdateResponseCompanyFiscalYears) SetTaxFraction(v int32)`

SetTaxFraction sets TaxFraction field to given value.


### GetTaxAccountMethod

`func (o *CompanyUpdateResponseCompanyFiscalYears) GetTaxAccountMethod() int32`

GetTaxAccountMethod returns the TaxAccountMethod field if non-nil, zero value otherwise.

### GetTaxAccountMethodOk

`func (o *CompanyUpdateResponseCompanyFiscalYears) GetTaxAccountMethodOk() (*int32, bool)`

GetTaxAccountMethodOk returns a tuple with the TaxAccountMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxAccountMethod

`func (o *CompanyUpdateResponseCompanyFiscalYears) SetTaxAccountMethod(v int32)`

SetTaxAccountMethod sets TaxAccountMethod field to given value.


### GetReturnCode

`func (o *CompanyUpdateResponseCompanyFiscalYears) GetReturnCode() int32`

GetReturnCode returns the ReturnCode field if non-nil, zero value otherwise.

### GetReturnCodeOk

`func (o *CompanyUpdateResponseCompanyFiscalYears) GetReturnCodeOk() (*int32, bool)`

GetReturnCodeOk returns a tuple with the ReturnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnCode

`func (o *CompanyUpdateResponseCompanyFiscalYears) SetReturnCode(v int32)`

SetReturnCode sets ReturnCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


