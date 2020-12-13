# TrialPlSectionsResponseTrialPlSectionsSections

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | 部門ID | 
**Name** | Pointer to **string** | 部門名 | [optional] 
**ClosingBalance** | Pointer to **int32** | 期末残高 | [optional] 
**Partners** | Pointer to [**[]TrialPlSectionsResponseTrialPlSectionsPartners**](TrialPlSectionsResponseTrialPlSectionsPartners.md) | breakdown_display_type:partner, account_item_display_type:account_item指定時のみ含まれる | [optional] 
**Items** | Pointer to [**[]TrialPlSectionsResponseTrialPlSectionsItems**](TrialPlSectionsResponseTrialPlSectionsItems.md) | breakdown_display_type:item, account_item_display_type:account_item指定時のみ含まれる | [optional] 

## Methods

### NewTrialPlSectionsResponseTrialPlSectionsSections

`func NewTrialPlSectionsResponseTrialPlSectionsSections(id int32, ) *TrialPlSectionsResponseTrialPlSectionsSections`

NewTrialPlSectionsResponseTrialPlSectionsSections instantiates a new TrialPlSectionsResponseTrialPlSectionsSections object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrialPlSectionsResponseTrialPlSectionsSectionsWithDefaults

`func NewTrialPlSectionsResponseTrialPlSectionsSectionsWithDefaults() *TrialPlSectionsResponseTrialPlSectionsSections`

NewTrialPlSectionsResponseTrialPlSectionsSectionsWithDefaults instantiates a new TrialPlSectionsResponseTrialPlSectionsSections object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TrialPlSectionsResponseTrialPlSectionsSections) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TrialPlSectionsResponseTrialPlSectionsSections) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TrialPlSectionsResponseTrialPlSectionsSections) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *TrialPlSectionsResponseTrialPlSectionsSections) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TrialPlSectionsResponseTrialPlSectionsSections) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TrialPlSectionsResponseTrialPlSectionsSections) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TrialPlSectionsResponseTrialPlSectionsSections) HasName() bool`

HasName returns a boolean if a field has been set.

### GetClosingBalance

`func (o *TrialPlSectionsResponseTrialPlSectionsSections) GetClosingBalance() int32`

GetClosingBalance returns the ClosingBalance field if non-nil, zero value otherwise.

### GetClosingBalanceOk

`func (o *TrialPlSectionsResponseTrialPlSectionsSections) GetClosingBalanceOk() (*int32, bool)`

GetClosingBalanceOk returns a tuple with the ClosingBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosingBalance

`func (o *TrialPlSectionsResponseTrialPlSectionsSections) SetClosingBalance(v int32)`

SetClosingBalance sets ClosingBalance field to given value.

### HasClosingBalance

`func (o *TrialPlSectionsResponseTrialPlSectionsSections) HasClosingBalance() bool`

HasClosingBalance returns a boolean if a field has been set.

### GetPartners

`func (o *TrialPlSectionsResponseTrialPlSectionsSections) GetPartners() []TrialPlSectionsResponseTrialPlSectionsPartners`

GetPartners returns the Partners field if non-nil, zero value otherwise.

### GetPartnersOk

`func (o *TrialPlSectionsResponseTrialPlSectionsSections) GetPartnersOk() (*[]TrialPlSectionsResponseTrialPlSectionsPartners, bool)`

GetPartnersOk returns a tuple with the Partners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartners

`func (o *TrialPlSectionsResponseTrialPlSectionsSections) SetPartners(v []TrialPlSectionsResponseTrialPlSectionsPartners)`

SetPartners sets Partners field to given value.

### HasPartners

`func (o *TrialPlSectionsResponseTrialPlSectionsSections) HasPartners() bool`

HasPartners returns a boolean if a field has been set.

### GetItems

`func (o *TrialPlSectionsResponseTrialPlSectionsSections) GetItems() []TrialPlSectionsResponseTrialPlSectionsItems`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *TrialPlSectionsResponseTrialPlSectionsSections) GetItemsOk() (*[]TrialPlSectionsResponseTrialPlSectionsItems, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *TrialPlSectionsResponseTrialPlSectionsSections) SetItems(v []TrialPlSectionsResponseTrialPlSectionsItems)`

SetItems sets Items field to given value.

### HasItems

`func (o *TrialPlSectionsResponseTrialPlSectionsSections) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


