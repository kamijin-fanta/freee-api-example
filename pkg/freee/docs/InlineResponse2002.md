# InlineResponse2002

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Deals** | [**[]Deal**](Deal.md) |  | 
**Meta** | [**InlineResponse2002Meta**](inline_response_200_2_meta.md) |  | 

## Methods

### NewInlineResponse2002

`func NewInlineResponse2002(deals []Deal, meta InlineResponse2002Meta, ) *InlineResponse2002`

NewInlineResponse2002 instantiates a new InlineResponse2002 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse2002WithDefaults

`func NewInlineResponse2002WithDefaults() *InlineResponse2002`

NewInlineResponse2002WithDefaults instantiates a new InlineResponse2002 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeals

`func (o *InlineResponse2002) GetDeals() []Deal`

GetDeals returns the Deals field if non-nil, zero value otherwise.

### GetDealsOk

`func (o *InlineResponse2002) GetDealsOk() (*[]Deal, bool)`

GetDealsOk returns a tuple with the Deals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeals

`func (o *InlineResponse2002) SetDeals(v []Deal)`

SetDeals sets Deals field to given value.


### GetMeta

`func (o *InlineResponse2002) GetMeta() InlineResponse2002Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *InlineResponse2002) GetMetaOk() (*InlineResponse2002Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *InlineResponse2002) SetMeta(v InlineResponse2002Meta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


