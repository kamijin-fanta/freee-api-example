# SelectablesIndexResponseAccountItems

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | 勘定科目ID | 
**Name** | Pointer to **string** | 勘定科目 | [optional] 
**Desc** | Pointer to **string** | 勘定科目の説明 | [optional] 
**Help** | Pointer to **string** | 勘定科目の説明（詳細） | [optional] 
**Shortcut** | Pointer to **string** | ショートカット | [optional] 
**DefaultTax** | Pointer to [**SelectablesIndexResponseDefaultTax**](selectablesIndexResponse_default_tax.md) |  | [optional] 

## Methods

### NewSelectablesIndexResponseAccountItems

`func NewSelectablesIndexResponseAccountItems(id int32, ) *SelectablesIndexResponseAccountItems`

NewSelectablesIndexResponseAccountItems instantiates a new SelectablesIndexResponseAccountItems object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSelectablesIndexResponseAccountItemsWithDefaults

`func NewSelectablesIndexResponseAccountItemsWithDefaults() *SelectablesIndexResponseAccountItems`

NewSelectablesIndexResponseAccountItemsWithDefaults instantiates a new SelectablesIndexResponseAccountItems object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SelectablesIndexResponseAccountItems) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SelectablesIndexResponseAccountItems) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SelectablesIndexResponseAccountItems) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *SelectablesIndexResponseAccountItems) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SelectablesIndexResponseAccountItems) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SelectablesIndexResponseAccountItems) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SelectablesIndexResponseAccountItems) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDesc

`func (o *SelectablesIndexResponseAccountItems) GetDesc() string`

GetDesc returns the Desc field if non-nil, zero value otherwise.

### GetDescOk

`func (o *SelectablesIndexResponseAccountItems) GetDescOk() (*string, bool)`

GetDescOk returns a tuple with the Desc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesc

`func (o *SelectablesIndexResponseAccountItems) SetDesc(v string)`

SetDesc sets Desc field to given value.

### HasDesc

`func (o *SelectablesIndexResponseAccountItems) HasDesc() bool`

HasDesc returns a boolean if a field has been set.

### GetHelp

`func (o *SelectablesIndexResponseAccountItems) GetHelp() string`

GetHelp returns the Help field if non-nil, zero value otherwise.

### GetHelpOk

`func (o *SelectablesIndexResponseAccountItems) GetHelpOk() (*string, bool)`

GetHelpOk returns a tuple with the Help field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelp

`func (o *SelectablesIndexResponseAccountItems) SetHelp(v string)`

SetHelp sets Help field to given value.

### HasHelp

`func (o *SelectablesIndexResponseAccountItems) HasHelp() bool`

HasHelp returns a boolean if a field has been set.

### GetShortcut

`func (o *SelectablesIndexResponseAccountItems) GetShortcut() string`

GetShortcut returns the Shortcut field if non-nil, zero value otherwise.

### GetShortcutOk

`func (o *SelectablesIndexResponseAccountItems) GetShortcutOk() (*string, bool)`

GetShortcutOk returns a tuple with the Shortcut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortcut

`func (o *SelectablesIndexResponseAccountItems) SetShortcut(v string)`

SetShortcut sets Shortcut field to given value.

### HasShortcut

`func (o *SelectablesIndexResponseAccountItems) HasShortcut() bool`

HasShortcut returns a boolean if a field has been set.

### GetDefaultTax

`func (o *SelectablesIndexResponseAccountItems) GetDefaultTax() SelectablesIndexResponseDefaultTax`

GetDefaultTax returns the DefaultTax field if non-nil, zero value otherwise.

### GetDefaultTaxOk

`func (o *SelectablesIndexResponseAccountItems) GetDefaultTaxOk() (*SelectablesIndexResponseDefaultTax, bool)`

GetDefaultTaxOk returns a tuple with the DefaultTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTax

`func (o *SelectablesIndexResponseAccountItems) SetDefaultTax(v SelectablesIndexResponseDefaultTax)`

SetDefaultTax sets DefaultTax field to given value.

### HasDefaultTax

`func (o *SelectablesIndexResponseAccountItems) HasDefaultTax() bool`

HasDefaultTax returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


