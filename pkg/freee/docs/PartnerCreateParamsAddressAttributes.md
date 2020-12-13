# PartnerCreateParamsAddressAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Zipcode** | Pointer to **string** | 郵便番号（8文字以内） | [optional] 
**PrefectureCode** | Pointer to **int32** | 都道府県コード（0: 北海道、1:青森、2:岩手、3:宮城、4:秋田、5:山形、6:福島、7:茨城、8:栃木、9:群馬、10:埼玉、11:千葉、12:東京、13:神奈川、14:新潟、15:富山、16:石川、17:福井、18:山梨、19:長野、20:岐阜、21:静岡、22:愛知、23:三重、24:滋賀、25:京都、26:大阪、27:兵庫、28:奈良、29:和歌山、30:鳥取、31:島根、32:岡山、33:広島、34:山口、35:徳島、36:香川、37:愛媛、38:高知、39:福岡、40:佐賀、41:長崎、42:熊本、43:大分、44:宮崎、45:鹿児島、46:沖縄 | [optional] 
**StreetName1** | Pointer to **string** | 市区町村・番地（255文字以内） | [optional] 
**StreetName2** | Pointer to **string** | 建物名・部屋番号など（255文字以内） | [optional] 

## Methods

### NewPartnerCreateParamsAddressAttributes

`func NewPartnerCreateParamsAddressAttributes() *PartnerCreateParamsAddressAttributes`

NewPartnerCreateParamsAddressAttributes instantiates a new PartnerCreateParamsAddressAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartnerCreateParamsAddressAttributesWithDefaults

`func NewPartnerCreateParamsAddressAttributesWithDefaults() *PartnerCreateParamsAddressAttributes`

NewPartnerCreateParamsAddressAttributesWithDefaults instantiates a new PartnerCreateParamsAddressAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetZipcode

`func (o *PartnerCreateParamsAddressAttributes) GetZipcode() string`

GetZipcode returns the Zipcode field if non-nil, zero value otherwise.

### GetZipcodeOk

`func (o *PartnerCreateParamsAddressAttributes) GetZipcodeOk() (*string, bool)`

GetZipcodeOk returns a tuple with the Zipcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipcode

`func (o *PartnerCreateParamsAddressAttributes) SetZipcode(v string)`

SetZipcode sets Zipcode field to given value.

### HasZipcode

`func (o *PartnerCreateParamsAddressAttributes) HasZipcode() bool`

HasZipcode returns a boolean if a field has been set.

### GetPrefectureCode

`func (o *PartnerCreateParamsAddressAttributes) GetPrefectureCode() int32`

GetPrefectureCode returns the PrefectureCode field if non-nil, zero value otherwise.

### GetPrefectureCodeOk

`func (o *PartnerCreateParamsAddressAttributes) GetPrefectureCodeOk() (*int32, bool)`

GetPrefectureCodeOk returns a tuple with the PrefectureCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefectureCode

`func (o *PartnerCreateParamsAddressAttributes) SetPrefectureCode(v int32)`

SetPrefectureCode sets PrefectureCode field to given value.

### HasPrefectureCode

`func (o *PartnerCreateParamsAddressAttributes) HasPrefectureCode() bool`

HasPrefectureCode returns a boolean if a field has been set.

### GetStreetName1

`func (o *PartnerCreateParamsAddressAttributes) GetStreetName1() string`

GetStreetName1 returns the StreetName1 field if non-nil, zero value otherwise.

### GetStreetName1Ok

`func (o *PartnerCreateParamsAddressAttributes) GetStreetName1Ok() (*string, bool)`

GetStreetName1Ok returns a tuple with the StreetName1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetName1

`func (o *PartnerCreateParamsAddressAttributes) SetStreetName1(v string)`

SetStreetName1 sets StreetName1 field to given value.

### HasStreetName1

`func (o *PartnerCreateParamsAddressAttributes) HasStreetName1() bool`

HasStreetName1 returns a boolean if a field has been set.

### GetStreetName2

`func (o *PartnerCreateParamsAddressAttributes) GetStreetName2() string`

GetStreetName2 returns the StreetName2 field if non-nil, zero value otherwise.

### GetStreetName2Ok

`func (o *PartnerCreateParamsAddressAttributes) GetStreetName2Ok() (*string, bool)`

GetStreetName2Ok returns a tuple with the StreetName2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetName2

`func (o *PartnerCreateParamsAddressAttributes) SetStreetName2(v string)`

SetStreetName2 sets StreetName2 field to given value.

### HasStreetName2

`func (o *PartnerCreateParamsAddressAttributes) HasStreetName2() bool`

HasStreetName2 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


