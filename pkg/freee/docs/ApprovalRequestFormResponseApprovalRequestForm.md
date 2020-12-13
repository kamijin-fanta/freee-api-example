# ApprovalRequestFormResponseApprovalRequestForm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | 申請フォームID | 
**CompanyId** | **int32** | 事業所ID | 
**Name** | **string** | 申請フォームの名前 | 
**Description** | **string** | 申請フォームの説明 | 
**Status** | **string** | ステータス(draft: 申請で使用しない、active: 申請で使用する、deleted: 削除済み) | 
**CreatedDate** | **string** | 作成日時 | 
**FormOrder** | **NullableInt32** | 表示順（申請者が選択する申請フォームの表示順を設定できます。小さい数ほど上位に表示されます。（0を除く整数のみ。マイナス不可）未入力の場合、表示順が後ろになります。同じ数字が入力された場合、登録順で表示されます。） | 
**Parts** | Pointer to [**[]ApprovalRequestResponseApprovalRequestApprovalRequestFormParts**](ApprovalRequestResponseApprovalRequestApprovalRequestFormParts.md) | 申請フォームの項目 | [optional] 
**RouteSettingCount** | **int32** | 適用された経路数 | 

## Methods

### NewApprovalRequestFormResponseApprovalRequestForm

`func NewApprovalRequestFormResponseApprovalRequestForm(id int32, companyId int32, name string, description string, status string, createdDate string, formOrder NullableInt32, routeSettingCount int32, ) *ApprovalRequestFormResponseApprovalRequestForm`

NewApprovalRequestFormResponseApprovalRequestForm instantiates a new ApprovalRequestFormResponseApprovalRequestForm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApprovalRequestFormResponseApprovalRequestFormWithDefaults

`func NewApprovalRequestFormResponseApprovalRequestFormWithDefaults() *ApprovalRequestFormResponseApprovalRequestForm`

NewApprovalRequestFormResponseApprovalRequestFormWithDefaults instantiates a new ApprovalRequestFormResponseApprovalRequestForm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApprovalRequestFormResponseApprovalRequestForm) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApprovalRequestFormResponseApprovalRequestForm) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApprovalRequestFormResponseApprovalRequestForm) SetId(v int32)`

SetId sets Id field to given value.


### GetCompanyId

`func (o *ApprovalRequestFormResponseApprovalRequestForm) GetCompanyId() int32`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *ApprovalRequestFormResponseApprovalRequestForm) GetCompanyIdOk() (*int32, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *ApprovalRequestFormResponseApprovalRequestForm) SetCompanyId(v int32)`

SetCompanyId sets CompanyId field to given value.


### GetName

`func (o *ApprovalRequestFormResponseApprovalRequestForm) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApprovalRequestFormResponseApprovalRequestForm) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApprovalRequestFormResponseApprovalRequestForm) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ApprovalRequestFormResponseApprovalRequestForm) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApprovalRequestFormResponseApprovalRequestForm) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApprovalRequestFormResponseApprovalRequestForm) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetStatus

`func (o *ApprovalRequestFormResponseApprovalRequestForm) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApprovalRequestFormResponseApprovalRequestForm) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApprovalRequestFormResponseApprovalRequestForm) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCreatedDate

`func (o *ApprovalRequestFormResponseApprovalRequestForm) GetCreatedDate() string`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *ApprovalRequestFormResponseApprovalRequestForm) GetCreatedDateOk() (*string, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *ApprovalRequestFormResponseApprovalRequestForm) SetCreatedDate(v string)`

SetCreatedDate sets CreatedDate field to given value.


### GetFormOrder

`func (o *ApprovalRequestFormResponseApprovalRequestForm) GetFormOrder() int32`

GetFormOrder returns the FormOrder field if non-nil, zero value otherwise.

### GetFormOrderOk

`func (o *ApprovalRequestFormResponseApprovalRequestForm) GetFormOrderOk() (*int32, bool)`

GetFormOrderOk returns a tuple with the FormOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormOrder

`func (o *ApprovalRequestFormResponseApprovalRequestForm) SetFormOrder(v int32)`

SetFormOrder sets FormOrder field to given value.


### SetFormOrderNil

`func (o *ApprovalRequestFormResponseApprovalRequestForm) SetFormOrderNil(b bool)`

 SetFormOrderNil sets the value for FormOrder to be an explicit nil

### UnsetFormOrder
`func (o *ApprovalRequestFormResponseApprovalRequestForm) UnsetFormOrder()`

UnsetFormOrder ensures that no value is present for FormOrder, not even an explicit nil
### GetParts

`func (o *ApprovalRequestFormResponseApprovalRequestForm) GetParts() []ApprovalRequestResponseApprovalRequestApprovalRequestFormParts`

GetParts returns the Parts field if non-nil, zero value otherwise.

### GetPartsOk

`func (o *ApprovalRequestFormResponseApprovalRequestForm) GetPartsOk() (*[]ApprovalRequestResponseApprovalRequestApprovalRequestFormParts, bool)`

GetPartsOk returns a tuple with the Parts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParts

`func (o *ApprovalRequestFormResponseApprovalRequestForm) SetParts(v []ApprovalRequestResponseApprovalRequestApprovalRequestFormParts)`

SetParts sets Parts field to given value.

### HasParts

`func (o *ApprovalRequestFormResponseApprovalRequestForm) HasParts() bool`

HasParts returns a boolean if a field has been set.

### GetRouteSettingCount

`func (o *ApprovalRequestFormResponseApprovalRequestForm) GetRouteSettingCount() int32`

GetRouteSettingCount returns the RouteSettingCount field if non-nil, zero value otherwise.

### GetRouteSettingCountOk

`func (o *ApprovalRequestFormResponseApprovalRequestForm) GetRouteSettingCountOk() (*int32, bool)`

GetRouteSettingCountOk returns a tuple with the RouteSettingCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteSettingCount

`func (o *ApprovalRequestFormResponseApprovalRequestForm) SetRouteSettingCount(v int32)`

SetRouteSettingCount sets RouteSettingCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


