# SelectablesIndexResponseAccountCategories

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Balance** | **string** | 収支 | 
**OrgCode** | **string** | 事業形態（個人事業主: personal、法人: corporate） | 
**Role** | **string** | カテゴリーコード | 
**Title** | **string** | カテゴリー名 | 
**Desc** | Pointer to **string** | カテゴリーの説明 | [optional] 
**AccountItems** | [**[]SelectablesIndexResponseAccountItems**](SelectablesIndexResponseAccountItems.md) | 勘定科目の一覧 | 

## Methods

### NewSelectablesIndexResponseAccountCategories

`func NewSelectablesIndexResponseAccountCategories(balance string, orgCode string, role string, title string, accountItems []SelectablesIndexResponseAccountItems, ) *SelectablesIndexResponseAccountCategories`

NewSelectablesIndexResponseAccountCategories instantiates a new SelectablesIndexResponseAccountCategories object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSelectablesIndexResponseAccountCategoriesWithDefaults

`func NewSelectablesIndexResponseAccountCategoriesWithDefaults() *SelectablesIndexResponseAccountCategories`

NewSelectablesIndexResponseAccountCategoriesWithDefaults instantiates a new SelectablesIndexResponseAccountCategories object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBalance

`func (o *SelectablesIndexResponseAccountCategories) GetBalance() string`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *SelectablesIndexResponseAccountCategories) GetBalanceOk() (*string, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *SelectablesIndexResponseAccountCategories) SetBalance(v string)`

SetBalance sets Balance field to given value.


### GetOrgCode

`func (o *SelectablesIndexResponseAccountCategories) GetOrgCode() string`

GetOrgCode returns the OrgCode field if non-nil, zero value otherwise.

### GetOrgCodeOk

`func (o *SelectablesIndexResponseAccountCategories) GetOrgCodeOk() (*string, bool)`

GetOrgCodeOk returns a tuple with the OrgCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgCode

`func (o *SelectablesIndexResponseAccountCategories) SetOrgCode(v string)`

SetOrgCode sets OrgCode field to given value.


### GetRole

`func (o *SelectablesIndexResponseAccountCategories) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *SelectablesIndexResponseAccountCategories) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *SelectablesIndexResponseAccountCategories) SetRole(v string)`

SetRole sets Role field to given value.


### GetTitle

`func (o *SelectablesIndexResponseAccountCategories) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SelectablesIndexResponseAccountCategories) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SelectablesIndexResponseAccountCategories) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDesc

`func (o *SelectablesIndexResponseAccountCategories) GetDesc() string`

GetDesc returns the Desc field if non-nil, zero value otherwise.

### GetDescOk

`func (o *SelectablesIndexResponseAccountCategories) GetDescOk() (*string, bool)`

GetDescOk returns a tuple with the Desc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesc

`func (o *SelectablesIndexResponseAccountCategories) SetDesc(v string)`

SetDesc sets Desc field to given value.

### HasDesc

`func (o *SelectablesIndexResponseAccountCategories) HasDesc() bool`

HasDesc returns a boolean if a field has been set.

### GetAccountItems

`func (o *SelectablesIndexResponseAccountCategories) GetAccountItems() []SelectablesIndexResponseAccountItems`

GetAccountItems returns the AccountItems field if non-nil, zero value otherwise.

### GetAccountItemsOk

`func (o *SelectablesIndexResponseAccountCategories) GetAccountItemsOk() (*[]SelectablesIndexResponseAccountItems, bool)`

GetAccountItemsOk returns a tuple with the AccountItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountItems

`func (o *SelectablesIndexResponseAccountCategories) SetAccountItems(v []SelectablesIndexResponseAccountItems)`

SetAccountItems sets AccountItems field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


