# DealResponseDealReceipts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | 証憑ID | 
**Status** | **string** | ステータス(unconfirmed:確認待ち、confirmed:確認済み、deleted:削除済み、ignored:無視) | 
**Description** | Pointer to **string** | メモ | [optional] 
**MimeType** | **string** | MIMEタイプ | 
**IssueDate** | Pointer to **string** | 発生日 | [optional] 
**Origin** | **string** | アップロード元種別 | 
**CreatedAt** | **string** | 作成日時（ISO8601形式） | 
**FileSrc** | **string** | ファイルのダウンロードURL（freeeにログインした状態でのみ閲覧可能です。） &lt;br&gt; &lt;br&gt; file_srcは廃止予定の属性になります。&lt;br&gt; file_srcに替わり、証憑ファイルのダウンロード APIをご利用ください。&lt;br&gt; 証憑ファイルのダウンロードAPIを利用することで、以下のようになります。 &lt;ul&gt;   &lt;li&gt;アプリケーション利用者はfreee APIアプリケーションにログインしていれば、証憑ダウンロード毎にfreeeに改めてログインすることなくファイルが参照できるようになります。&lt;/li&gt; &lt;/ul&gt; | 
**User** | [**DealResponseDealUser**](dealResponse_deal_user.md) |  | 

## Methods

### NewDealResponseDealReceipts

`func NewDealResponseDealReceipts(id int32, status string, mimeType string, origin string, createdAt string, fileSrc string, user DealResponseDealUser, ) *DealResponseDealReceipts`

NewDealResponseDealReceipts instantiates a new DealResponseDealReceipts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDealResponseDealReceiptsWithDefaults

`func NewDealResponseDealReceiptsWithDefaults() *DealResponseDealReceipts`

NewDealResponseDealReceiptsWithDefaults instantiates a new DealResponseDealReceipts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DealResponseDealReceipts) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DealResponseDealReceipts) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DealResponseDealReceipts) SetId(v int32)`

SetId sets Id field to given value.


### GetStatus

`func (o *DealResponseDealReceipts) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DealResponseDealReceipts) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DealResponseDealReceipts) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetDescription

`func (o *DealResponseDealReceipts) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DealResponseDealReceipts) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DealResponseDealReceipts) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DealResponseDealReceipts) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMimeType

`func (o *DealResponseDealReceipts) GetMimeType() string`

GetMimeType returns the MimeType field if non-nil, zero value otherwise.

### GetMimeTypeOk

`func (o *DealResponseDealReceipts) GetMimeTypeOk() (*string, bool)`

GetMimeTypeOk returns a tuple with the MimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimeType

`func (o *DealResponseDealReceipts) SetMimeType(v string)`

SetMimeType sets MimeType field to given value.


### GetIssueDate

`func (o *DealResponseDealReceipts) GetIssueDate() string`

GetIssueDate returns the IssueDate field if non-nil, zero value otherwise.

### GetIssueDateOk

`func (o *DealResponseDealReceipts) GetIssueDateOk() (*string, bool)`

GetIssueDateOk returns a tuple with the IssueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueDate

`func (o *DealResponseDealReceipts) SetIssueDate(v string)`

SetIssueDate sets IssueDate field to given value.

### HasIssueDate

`func (o *DealResponseDealReceipts) HasIssueDate() bool`

HasIssueDate returns a boolean if a field has been set.

### GetOrigin

`func (o *DealResponseDealReceipts) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *DealResponseDealReceipts) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *DealResponseDealReceipts) SetOrigin(v string)`

SetOrigin sets Origin field to given value.


### GetCreatedAt

`func (o *DealResponseDealReceipts) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DealResponseDealReceipts) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DealResponseDealReceipts) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetFileSrc

`func (o *DealResponseDealReceipts) GetFileSrc() string`

GetFileSrc returns the FileSrc field if non-nil, zero value otherwise.

### GetFileSrcOk

`func (o *DealResponseDealReceipts) GetFileSrcOk() (*string, bool)`

GetFileSrcOk returns a tuple with the FileSrc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSrc

`func (o *DealResponseDealReceipts) SetFileSrc(v string)`

SetFileSrc sets FileSrc field to given value.


### GetUser

`func (o *DealResponseDealReceipts) GetUser() DealResponseDealUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *DealResponseDealReceipts) GetUserOk() (*DealResponseDealUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *DealResponseDealReceipts) SetUser(v DealResponseDealUser)`

SetUser sets User field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


