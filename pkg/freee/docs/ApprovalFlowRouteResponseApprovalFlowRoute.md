# ApprovalFlowRouteResponseApprovalFlowRoute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | 申請経路ID | 
**Name** | Pointer to **string** | 申請経路名 | [optional] 
**Description** | Pointer to **string** | 申請経路の説明 | [optional] 
**UserId** | Pointer to **NullableInt32** | 更新したユーザーのユーザーID | [optional] 
**DefinitionSystem** | Pointer to **bool** | システム作成の申請経路かどうか | [optional] 
**FirstStepId** | Pointer to **int32** | 最初の承認ステップのID | [optional] 
**Usages** | Pointer to **[]string** | 申請種別（申請経路を使用できる申請種別を示します。例えば、ApprovalRequest の場合は、各種申請で使用できる申請経路です。） * &#x60;TxnApproval&#x60; - 仕訳承認 * &#x60;ExpenseApplication&#x60; - 経費精算 * &#x60;PaymentRequest&#x60; - 支払依頼 * &#x60;ApprovalRequest&#x60; - 各種申請 * &#x60;DocApproval&#x60; - 請求書等 (見積書・納品書・請求書・発注書) | [optional] 
**RequestFormIds** | **[]int32** | 申請経路で利用できる申請フォームID配列 | 
**Steps** | Pointer to [**[]ApprovalFlowRouteResponseApprovalFlowRouteSteps**](ApprovalFlowRouteResponseApprovalFlowRouteSteps.md) | 承認ステップ（配列） | [optional] 

## Methods

### NewApprovalFlowRouteResponseApprovalFlowRoute

`func NewApprovalFlowRouteResponseApprovalFlowRoute(id int32, requestFormIds []int32, ) *ApprovalFlowRouteResponseApprovalFlowRoute`

NewApprovalFlowRouteResponseApprovalFlowRoute instantiates a new ApprovalFlowRouteResponseApprovalFlowRoute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApprovalFlowRouteResponseApprovalFlowRouteWithDefaults

`func NewApprovalFlowRouteResponseApprovalFlowRouteWithDefaults() *ApprovalFlowRouteResponseApprovalFlowRoute`

NewApprovalFlowRouteResponseApprovalFlowRouteWithDefaults instantiates a new ApprovalFlowRouteResponseApprovalFlowRoute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApprovalFlowRouteResponseApprovalFlowRoute) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApprovalFlowRouteResponseApprovalFlowRoute) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApprovalFlowRouteResponseApprovalFlowRoute) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *ApprovalFlowRouteResponseApprovalFlowRoute) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApprovalFlowRouteResponseApprovalFlowRoute) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApprovalFlowRouteResponseApprovalFlowRoute) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApprovalFlowRouteResponseApprovalFlowRoute) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ApprovalFlowRouteResponseApprovalFlowRoute) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApprovalFlowRouteResponseApprovalFlowRoute) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApprovalFlowRouteResponseApprovalFlowRoute) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApprovalFlowRouteResponseApprovalFlowRoute) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetUserId

`func (o *ApprovalFlowRouteResponseApprovalFlowRoute) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ApprovalFlowRouteResponseApprovalFlowRoute) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ApprovalFlowRouteResponseApprovalFlowRoute) SetUserId(v int32)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ApprovalFlowRouteResponseApprovalFlowRoute) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *ApprovalFlowRouteResponseApprovalFlowRoute) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *ApprovalFlowRouteResponseApprovalFlowRoute) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetDefinitionSystem

`func (o *ApprovalFlowRouteResponseApprovalFlowRoute) GetDefinitionSystem() bool`

GetDefinitionSystem returns the DefinitionSystem field if non-nil, zero value otherwise.

### GetDefinitionSystemOk

`func (o *ApprovalFlowRouteResponseApprovalFlowRoute) GetDefinitionSystemOk() (*bool, bool)`

GetDefinitionSystemOk returns a tuple with the DefinitionSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinitionSystem

`func (o *ApprovalFlowRouteResponseApprovalFlowRoute) SetDefinitionSystem(v bool)`

SetDefinitionSystem sets DefinitionSystem field to given value.

### HasDefinitionSystem

`func (o *ApprovalFlowRouteResponseApprovalFlowRoute) HasDefinitionSystem() bool`

HasDefinitionSystem returns a boolean if a field has been set.

### GetFirstStepId

`func (o *ApprovalFlowRouteResponseApprovalFlowRoute) GetFirstStepId() int32`

GetFirstStepId returns the FirstStepId field if non-nil, zero value otherwise.

### GetFirstStepIdOk

`func (o *ApprovalFlowRouteResponseApprovalFlowRoute) GetFirstStepIdOk() (*int32, bool)`

GetFirstStepIdOk returns a tuple with the FirstStepId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstStepId

`func (o *ApprovalFlowRouteResponseApprovalFlowRoute) SetFirstStepId(v int32)`

SetFirstStepId sets FirstStepId field to given value.

### HasFirstStepId

`func (o *ApprovalFlowRouteResponseApprovalFlowRoute) HasFirstStepId() bool`

HasFirstStepId returns a boolean if a field has been set.

### GetUsages

`func (o *ApprovalFlowRouteResponseApprovalFlowRoute) GetUsages() []string`

GetUsages returns the Usages field if non-nil, zero value otherwise.

### GetUsagesOk

`func (o *ApprovalFlowRouteResponseApprovalFlowRoute) GetUsagesOk() (*[]string, bool)`

GetUsagesOk returns a tuple with the Usages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsages

`func (o *ApprovalFlowRouteResponseApprovalFlowRoute) SetUsages(v []string)`

SetUsages sets Usages field to given value.

### HasUsages

`func (o *ApprovalFlowRouteResponseApprovalFlowRoute) HasUsages() bool`

HasUsages returns a boolean if a field has been set.

### GetRequestFormIds

`func (o *ApprovalFlowRouteResponseApprovalFlowRoute) GetRequestFormIds() []int32`

GetRequestFormIds returns the RequestFormIds field if non-nil, zero value otherwise.

### GetRequestFormIdsOk

`func (o *ApprovalFlowRouteResponseApprovalFlowRoute) GetRequestFormIdsOk() (*[]int32, bool)`

GetRequestFormIdsOk returns a tuple with the RequestFormIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestFormIds

`func (o *ApprovalFlowRouteResponseApprovalFlowRoute) SetRequestFormIds(v []int32)`

SetRequestFormIds sets RequestFormIds field to given value.


### GetSteps

`func (o *ApprovalFlowRouteResponseApprovalFlowRoute) GetSteps() []ApprovalFlowRouteResponseApprovalFlowRouteSteps`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *ApprovalFlowRouteResponseApprovalFlowRoute) GetStepsOk() (*[]ApprovalFlowRouteResponseApprovalFlowRouteSteps, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *ApprovalFlowRouteResponseApprovalFlowRoute) SetSteps(v []ApprovalFlowRouteResponseApprovalFlowRouteSteps)`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *ApprovalFlowRouteResponseApprovalFlowRoute) HasSteps() bool`

HasSteps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


