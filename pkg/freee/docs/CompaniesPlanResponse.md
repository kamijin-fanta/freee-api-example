# CompaniesPlanResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | 事業所ID | 
**Plan** | **string** | 会計プラン 個人用(non_charged: 無料プラン、starter: スターター、standard: スタンダード、premium: プレミアム) 法人用(non_charged: 無料プラン、minimum: ミニマム、basic: ベーシック、professional: プロフェッショナル、enterprise: エンタープライズ) | 
**OrgCode** | **string** | 事業形態（個人事業主: personal、法人: corporate） | 

## Methods

### NewCompaniesPlanResponse

`func NewCompaniesPlanResponse(id int32, plan string, orgCode string, ) *CompaniesPlanResponse`

NewCompaniesPlanResponse instantiates a new CompaniesPlanResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompaniesPlanResponseWithDefaults

`func NewCompaniesPlanResponseWithDefaults() *CompaniesPlanResponse`

NewCompaniesPlanResponseWithDefaults instantiates a new CompaniesPlanResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CompaniesPlanResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CompaniesPlanResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CompaniesPlanResponse) SetId(v int32)`

SetId sets Id field to given value.


### GetPlan

`func (o *CompaniesPlanResponse) GetPlan() string`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *CompaniesPlanResponse) GetPlanOk() (*string, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *CompaniesPlanResponse) SetPlan(v string)`

SetPlan sets Plan field to given value.


### GetOrgCode

`func (o *CompaniesPlanResponse) GetOrgCode() string`

GetOrgCode returns the OrgCode field if non-nil, zero value otherwise.

### GetOrgCodeOk

`func (o *CompaniesPlanResponse) GetOrgCodeOk() (*string, bool)`

GetOrgCodeOk returns a tuple with the OrgCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgCode

`func (o *CompaniesPlanResponse) SetOrgCode(v string)`

SetOrgCode sets OrgCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


