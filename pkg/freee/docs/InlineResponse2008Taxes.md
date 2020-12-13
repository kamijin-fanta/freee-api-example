# InlineResponse2008Taxes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **int32** | 税区分コード | 
**Name** | **string** | 税区分名 | 
**NameJa** | **string** | 税区分名（日本語表示用） | 
**DisplayCategory** | **NullableString** | 税区分の表示カテゴリ（tax_5: 5%表示の税区分、tax_8: 8%表示の税区分、tax_r8: 軽減税率8%表示の税区分、tax_10: 10%表示の税区分、null: 税率未設定税区分） | 
**Available** | **bool** | true: 使用する、false: 使用しない | 

## Methods

### NewInlineResponse2008Taxes

`func NewInlineResponse2008Taxes(code int32, name string, nameJa string, displayCategory NullableString, available bool, ) *InlineResponse2008Taxes`

NewInlineResponse2008Taxes instantiates a new InlineResponse2008Taxes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse2008TaxesWithDefaults

`func NewInlineResponse2008TaxesWithDefaults() *InlineResponse2008Taxes`

NewInlineResponse2008TaxesWithDefaults instantiates a new InlineResponse2008Taxes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *InlineResponse2008Taxes) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *InlineResponse2008Taxes) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *InlineResponse2008Taxes) SetCode(v int32)`

SetCode sets Code field to given value.


### GetName

`func (o *InlineResponse2008Taxes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse2008Taxes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse2008Taxes) SetName(v string)`

SetName sets Name field to given value.


### GetNameJa

`func (o *InlineResponse2008Taxes) GetNameJa() string`

GetNameJa returns the NameJa field if non-nil, zero value otherwise.

### GetNameJaOk

`func (o *InlineResponse2008Taxes) GetNameJaOk() (*string, bool)`

GetNameJaOk returns a tuple with the NameJa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameJa

`func (o *InlineResponse2008Taxes) SetNameJa(v string)`

SetNameJa sets NameJa field to given value.


### GetDisplayCategory

`func (o *InlineResponse2008Taxes) GetDisplayCategory() string`

GetDisplayCategory returns the DisplayCategory field if non-nil, zero value otherwise.

### GetDisplayCategoryOk

`func (o *InlineResponse2008Taxes) GetDisplayCategoryOk() (*string, bool)`

GetDisplayCategoryOk returns a tuple with the DisplayCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayCategory

`func (o *InlineResponse2008Taxes) SetDisplayCategory(v string)`

SetDisplayCategory sets DisplayCategory field to given value.


### SetDisplayCategoryNil

`func (o *InlineResponse2008Taxes) SetDisplayCategoryNil(b bool)`

 SetDisplayCategoryNil sets the value for DisplayCategory to be an explicit nil

### UnsetDisplayCategory
`func (o *InlineResponse2008Taxes) UnsetDisplayCategory()`

UnsetDisplayCategory ensures that no value is present for DisplayCategory, not even an explicit nil
### GetAvailable

`func (o *InlineResponse2008Taxes) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *InlineResponse2008Taxes) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *InlineResponse2008Taxes) SetAvailable(v bool)`

SetAvailable sets Available field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


