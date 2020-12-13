# ApprovalFlowRoutesIndexResponseApprovalFlowRoutes

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
**RequestFormIds** | Pointer to **[]int32** | 申請経路で利用できる申請フォームID配列 | [optional] 
**DefaultRoute** | **bool** | 基本経路として設定されているかどうか&lt;br&gt;&lt;br&gt; リクエストパラメータusageに下記のいずれかが指定され、かつ、基本経路の場合はtrueになります。 * &#x60;TxnApproval&#x60; - 仕訳承認 * &#x60;ExpenseApplication&#x60; - 経費精算 * &#x60;PaymentRequest&#x60; - 支払依頼 * &#x60;ApprovalRequest&#x60;(リクエストパラメータrequest_form_idを同時に指定) - 各種申請 * &#x60;DocApproval&#x60; - 請求書等 (見積書・納品書・請求書・発注書)  &lt;a href&#x3D;\&quot;https://support.freee.co.jp/hc/ja/articles/900000507963-%E7%94%B3%E8%AB%8B%E3%83%95%E3%82%A9%E3%83%BC%E3%83%A0%E3%81%AE%E5%9F%BA%E6%9C%AC%E7%B5%8C%E8%B7%AF%E8%A8%AD%E5%AE%9A\&quot; target&#x3D;\&quot;_blank\&quot;&gt;申請フォームの基本経路設定&lt;/a&gt;  | 

## Methods

### NewApprovalFlowRoutesIndexResponseApprovalFlowRoutes

`func NewApprovalFlowRoutesIndexResponseApprovalFlowRoutes(id int32, defaultRoute bool, ) *ApprovalFlowRoutesIndexResponseApprovalFlowRoutes`

NewApprovalFlowRoutesIndexResponseApprovalFlowRoutes instantiates a new ApprovalFlowRoutesIndexResponseApprovalFlowRoutes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApprovalFlowRoutesIndexResponseApprovalFlowRoutesWithDefaults

`func NewApprovalFlowRoutesIndexResponseApprovalFlowRoutesWithDefaults() *ApprovalFlowRoutesIndexResponseApprovalFlowRoutes`

NewApprovalFlowRoutesIndexResponseApprovalFlowRoutesWithDefaults instantiates a new ApprovalFlowRoutesIndexResponseApprovalFlowRoutes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApprovalFlowRoutesIndexResponseApprovalFlowRoutes) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApprovalFlowRoutesIndexResponseApprovalFlowRoutes) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApprovalFlowRoutesIndexResponseApprovalFlowRoutes) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *ApprovalFlowRoutesIndexResponseApprovalFlowRoutes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApprovalFlowRoutesIndexResponseApprovalFlowRoutes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApprovalFlowRoutesIndexResponseApprovalFlowRoutes) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApprovalFlowRoutesIndexResponseApprovalFlowRoutes) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ApprovalFlowRoutesIndexResponseApprovalFlowRoutes) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApprovalFlowRoutesIndexResponseApprovalFlowRoutes) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApprovalFlowRoutesIndexResponseApprovalFlowRoutes) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApprovalFlowRoutesIndexResponseApprovalFlowRoutes) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetUserId

`func (o *ApprovalFlowRoutesIndexResponseApprovalFlowRoutes) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ApprovalFlowRoutesIndexResponseApprovalFlowRoutes) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ApprovalFlowRoutesIndexResponseApprovalFlowRoutes) SetUserId(v int32)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ApprovalFlowRoutesIndexResponseApprovalFlowRoutes) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *ApprovalFlowRoutesIndexResponseApprovalFlowRoutes) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *ApprovalFlowRoutesIndexResponseApprovalFlowRoutes) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetDefinitionSystem

`func (o *ApprovalFlowRoutesIndexResponseApprovalFlowRoutes) GetDefinitionSystem() bool`

GetDefinitionSystem returns the DefinitionSystem field if non-nil, zero value otherwise.

### GetDefinitionSystemOk

`func (o *ApprovalFlowRoutesIndexResponseApprovalFlowRoutes) GetDefinitionSystemOk() (*bool, bool)`

GetDefinitionSystemOk returns a tuple with the DefinitionSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinitionSystem

`func (o *ApprovalFlowRoutesIndexResponseApprovalFlowRoutes) SetDefinitionSystem(v bool)`

SetDefinitionSystem sets DefinitionSystem field to given value.

### HasDefinitionSystem

`func (o *ApprovalFlowRoutesIndexResponseApprovalFlowRoutes) HasDefinitionSystem() bool`

HasDefinitionSystem returns a boolean if a field has been set.

### GetFirstStepId

`func (o *ApprovalFlowRoutesIndexResponseApprovalFlowRoutes) GetFirstStepId() int32`

GetFirstStepId returns the FirstStepId field if non-nil, zero value otherwise.

### GetFirstStepIdOk

`func (o *ApprovalFlowRoutesIndexResponseApprovalFlowRoutes) GetFirstStepIdOk() (*int32, bool)`

GetFirstStepIdOk returns a tuple with the FirstStepId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstStepId

`func (o *ApprovalFlowRoutesIndexResponseApprovalFlowRoutes) SetFirstStepId(v int32)`

SetFirstStepId sets FirstStepId field to given value.

### HasFirstStepId

`func (o *ApprovalFlowRoutesIndexResponseApprovalFlowRoutes) HasFirstStepId() bool`

HasFirstStepId returns a boolean if a field has been set.

### GetUsages

`func (o *ApprovalFlowRoutesIndexResponseApprovalFlowRoutes) GetUsages() []string`

GetUsages returns the Usages field if non-nil, zero value otherwise.

### GetUsagesOk

`func (o *ApprovalFlowRoutesIndexResponseApprovalFlowRoutes) GetUsagesOk() (*[]string, bool)`

GetUsagesOk returns a tuple with the Usages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsages

`func (o *ApprovalFlowRoutesIndexResponseApprovalFlowRoutes) SetUsages(v []string)`

SetUsages sets Usages field to given value.

### HasUsages

`func (o *ApprovalFlowRoutesIndexResponseApprovalFlowRoutes) HasUsages() bool`

HasUsages returns a boolean if a field has been set.

### GetRequestFormIds

`func (o *ApprovalFlowRoutesIndexResponseApprovalFlowRoutes) GetRequestFormIds() []int32`

GetRequestFormIds returns the RequestFormIds field if non-nil, zero value otherwise.

### GetRequestFormIdsOk

`func (o *ApprovalFlowRoutesIndexResponseApprovalFlowRoutes) GetRequestFormIdsOk() (*[]int32, bool)`

GetRequestFormIdsOk returns a tuple with the RequestFormIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestFormIds

`func (o *ApprovalFlowRoutesIndexResponseApprovalFlowRoutes) SetRequestFormIds(v []int32)`

SetRequestFormIds sets RequestFormIds field to given value.

### HasRequestFormIds

`func (o *ApprovalFlowRoutesIndexResponseApprovalFlowRoutes) HasRequestFormIds() bool`

HasRequestFormIds returns a boolean if a field has been set.

### GetDefaultRoute

`func (o *ApprovalFlowRoutesIndexResponseApprovalFlowRoutes) GetDefaultRoute() bool`

GetDefaultRoute returns the DefaultRoute field if non-nil, zero value otherwise.

### GetDefaultRouteOk

`func (o *ApprovalFlowRoutesIndexResponseApprovalFlowRoutes) GetDefaultRouteOk() (*bool, bool)`

GetDefaultRouteOk returns a tuple with the DefaultRoute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRoute

`func (o *ApprovalFlowRoutesIndexResponseApprovalFlowRoutes) SetDefaultRoute(v bool)`

SetDefaultRoute sets DefaultRoute field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


