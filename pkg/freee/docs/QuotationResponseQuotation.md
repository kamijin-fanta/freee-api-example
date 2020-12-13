# QuotationResponseQuotation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | 見積書ID | 
**CompanyId** | **int32** | 事業所ID | 
**IssueDate** | **string** | 見積日 (yyyy-mm-dd) | 
**PartnerId** | **NullableInt32** | 取引先ID | 
**PartnerCode** | Pointer to **NullableString** | 取引先コード | [optional] 
**QuotationNumber** | **string** | 見積書番号 | 
**Title** | Pointer to **NullableString** | タイトル | [optional] 
**TotalAmount** | **int32** | 合計金額 | 
**TotalVat** | Pointer to **int32** | 消費税 | [optional] 
**SubTotal** | Pointer to **int32** | 小計 | [optional] 
**Description** | Pointer to **NullableString** | 概要 | [optional] 
**QuotationStatus** | **string** | 見積書ステータス  (unsubmitted: 送付待ち, submitted: 送付済み, all: 全て) | 
**WebPublishedAt** | Pointer to **NullableString** | Web共有日時(最新) | [optional] 
**WebDownloadedAt** | Pointer to **NullableString** | Web共有ダウンロード日時(最新) | [optional] 
**WebConfirmedAt** | Pointer to **NullableString** | Web共有取引先確認日時(最新) | [optional] 
**MailSentAt** | Pointer to **NullableString** | メール送信日時(最新) | [optional] 
**PartnerName** | Pointer to **NullableString** | 取引先名 | [optional] 
**PartnerDisplayName** | Pointer to **NullableString** | 見積書に表示する取引先名 | [optional] 
**PartnerTitle** | **NullableString** | 敬称（御中、様、(空白)の3つから選択） | 
**PartnerZipcode** | Pointer to **NullableString** | 郵便番号 | [optional] 
**PartnerPrefectureCode** | Pointer to **NullableInt32** | 都道府県コード（0:北海道、1:青森、2:岩手、3:宮城、4:秋田、5:山形、6:福島、7:茨城、8:栃木、9:群馬、10:埼玉、11:千葉、12:東京、13:神奈川、14:新潟、15:富山、16:石川、17:福井、18:山梨、19:長野、20:岐阜、21:静岡、22:愛知、23:三重、24:滋賀、25:京都、26:大阪、27:兵庫、28:奈良、29:和歌山、30:鳥取、31:島根、32:岡山、33:広島、34:山口、35:徳島、36:香川、37:愛媛、38:高知、39:福岡、40:佐賀、41:長崎、42:熊本、43:大分、44:宮崎、45:鹿児島、46:沖縄 | [optional] 
**PartnerPrefectureName** | Pointer to **NullableString** | 都道府県 | [optional] 
**PartnerAddress1** | Pointer to **NullableString** | 市区町村・番地 | [optional] 
**PartnerAddress2** | Pointer to **NullableString** | 建物名・部屋番号など | [optional] 
**PartnerContactInfo** | Pointer to **NullableString** | 取引先担当者名 | [optional] 
**CompanyName** | **string** | 事業所名 | 
**CompanyZipcode** | Pointer to **NullableString** | 郵便番号 | [optional] 
**CompanyPrefectureCode** | Pointer to **NullableInt32** | 都道府県コード（0:北海道、1:青森、2:岩手、3:宮城、4:秋田、5:山形、6:福島、7:茨城、8:栃木、9:群馬、10:埼玉、11:千葉、12:東京、13:神奈川、14:新潟、15:富山、16:石川、17:福井、18:山梨、19:長野、20:岐阜、21:静岡、22:愛知、23:三重、24:滋賀、25:京都、26:大阪、27:兵庫、28:奈良、29:和歌山、30:鳥取、31:島根、32:岡山、33:広島、34:山口、35:徳島、36:香川、37:愛媛、38:高知、39:福岡、40:佐賀、41:長崎、42:熊本、43:大分、44:宮崎、45:鹿児島、46:沖縄 | [optional] 
**CompanyPrefectureName** | Pointer to **NullableString** | 都道府県 | [optional] 
**CompanyAddress1** | Pointer to **NullableString** | 市区町村・番地 | [optional] 
**CompanyAddress2** | Pointer to **NullableString** | 建物名・部屋番号など | [optional] 
**CompanyContactInfo** | Pointer to **NullableString** | 事業所担当者名 | [optional] 
**Message** | Pointer to **NullableString** | メッセージ | [optional] 
**Notes** | Pointer to **NullableString** | 備考 | [optional] 
**QuotationLayout** | **string** | 見積書レイアウト * &#x60;default_classic&#x60; - レイアウト１/クラシック (デフォルト)  * &#x60;standard_classic&#x60; - レイアウト２/クラシック  * &#x60;envelope_classic&#x60; - 封筒１/クラシック  * &#x60;default_modern&#x60; - レイアウト１/モダン  * &#x60;standard_modern&#x60; - レイアウト２/モダン  * &#x60;envelope_modern&#x60; - 封筒/モダン | 
**TaxEntryMethod** | **string** | 見積書の消費税計算方法(inclusive: 内税, exclusive: 外税) | 
**QuotationContents** | Pointer to [**[]QuotationIndexResponseQuotationContents**](QuotationIndexResponseQuotationContents.md) | 見積内容 | [optional] 
**TotalAmountPerVatRate** | [**InvoiceIndexResponseTotalAmountPerVatRate**](invoiceIndexResponse_total_amount_per_vat_rate.md) |  | 
**RelatedInvoiceId** | Pointer to **NullableInt32** | 関連する請求書ID&lt;br&gt; 下記で作成したものが該当します。  &lt;a href&#x3D;\&quot;https://support.freee.co.jp/hc/ja/articles/203318410-請求書-納品書-見積書を作成する%231-2#1-2\&quot; target&#x3D;\&quot;_blank\&quot;&gt;見積書・納品書を納品書・請求書に変換する&lt;/a&gt;&lt;br&gt; &lt;a href&#x3D;\&quot;https://support.freee.co.jp/hc/ja/articles/209076226\&quot; target&#x3D;\&quot;_blank\&quot;&gt;複数の見積書・納品書から合算請求書を作成する&lt;/a&gt;&lt;br&gt;  | [optional] 
**RelatedQuotationIds** | Pointer to **[]int32** | 関連する見積書ID(配列)&lt;br&gt; 下記で作成したものが該当します。  &lt;a href&#x3D;\&quot;https://support.freee.co.jp/hc/ja/articles/203318410-請求書-納品書-見積書を作成する%231-2#1-2\&quot; target&#x3D;\&quot;_blank\&quot;&gt;見積書・納品書を納品書・請求書に変換する&lt;/a&gt;&lt;br&gt; &lt;a href&#x3D;\&quot;https://support.freee.co.jp/hc/ja/articles/209076226\&quot; target&#x3D;\&quot;_blank\&quot;&gt;複数の見積書・納品書から合算請求書を作成する&lt;/a&gt;&lt;br&gt;  | [optional] 

## Methods

### NewQuotationResponseQuotation

`func NewQuotationResponseQuotation(id int32, companyId int32, issueDate string, partnerId NullableInt32, quotationNumber string, totalAmount int32, quotationStatus string, partnerTitle NullableString, companyName string, quotationLayout string, taxEntryMethod string, totalAmountPerVatRate InvoiceIndexResponseTotalAmountPerVatRate, ) *QuotationResponseQuotation`

NewQuotationResponseQuotation instantiates a new QuotationResponseQuotation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuotationResponseQuotationWithDefaults

`func NewQuotationResponseQuotationWithDefaults() *QuotationResponseQuotation`

NewQuotationResponseQuotationWithDefaults instantiates a new QuotationResponseQuotation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QuotationResponseQuotation) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QuotationResponseQuotation) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QuotationResponseQuotation) SetId(v int32)`

SetId sets Id field to given value.


### GetCompanyId

`func (o *QuotationResponseQuotation) GetCompanyId() int32`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *QuotationResponseQuotation) GetCompanyIdOk() (*int32, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *QuotationResponseQuotation) SetCompanyId(v int32)`

SetCompanyId sets CompanyId field to given value.


### GetIssueDate

`func (o *QuotationResponseQuotation) GetIssueDate() string`

GetIssueDate returns the IssueDate field if non-nil, zero value otherwise.

### GetIssueDateOk

`func (o *QuotationResponseQuotation) GetIssueDateOk() (*string, bool)`

GetIssueDateOk returns a tuple with the IssueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueDate

`func (o *QuotationResponseQuotation) SetIssueDate(v string)`

SetIssueDate sets IssueDate field to given value.


### GetPartnerId

`func (o *QuotationResponseQuotation) GetPartnerId() int32`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *QuotationResponseQuotation) GetPartnerIdOk() (*int32, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *QuotationResponseQuotation) SetPartnerId(v int32)`

SetPartnerId sets PartnerId field to given value.


### SetPartnerIdNil

`func (o *QuotationResponseQuotation) SetPartnerIdNil(b bool)`

 SetPartnerIdNil sets the value for PartnerId to be an explicit nil

### UnsetPartnerId
`func (o *QuotationResponseQuotation) UnsetPartnerId()`

UnsetPartnerId ensures that no value is present for PartnerId, not even an explicit nil
### GetPartnerCode

`func (o *QuotationResponseQuotation) GetPartnerCode() string`

GetPartnerCode returns the PartnerCode field if non-nil, zero value otherwise.

### GetPartnerCodeOk

`func (o *QuotationResponseQuotation) GetPartnerCodeOk() (*string, bool)`

GetPartnerCodeOk returns a tuple with the PartnerCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerCode

`func (o *QuotationResponseQuotation) SetPartnerCode(v string)`

SetPartnerCode sets PartnerCode field to given value.

### HasPartnerCode

`func (o *QuotationResponseQuotation) HasPartnerCode() bool`

HasPartnerCode returns a boolean if a field has been set.

### SetPartnerCodeNil

`func (o *QuotationResponseQuotation) SetPartnerCodeNil(b bool)`

 SetPartnerCodeNil sets the value for PartnerCode to be an explicit nil

### UnsetPartnerCode
`func (o *QuotationResponseQuotation) UnsetPartnerCode()`

UnsetPartnerCode ensures that no value is present for PartnerCode, not even an explicit nil
### GetQuotationNumber

`func (o *QuotationResponseQuotation) GetQuotationNumber() string`

GetQuotationNumber returns the QuotationNumber field if non-nil, zero value otherwise.

### GetQuotationNumberOk

`func (o *QuotationResponseQuotation) GetQuotationNumberOk() (*string, bool)`

GetQuotationNumberOk returns a tuple with the QuotationNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotationNumber

`func (o *QuotationResponseQuotation) SetQuotationNumber(v string)`

SetQuotationNumber sets QuotationNumber field to given value.


### GetTitle

`func (o *QuotationResponseQuotation) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *QuotationResponseQuotation) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *QuotationResponseQuotation) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *QuotationResponseQuotation) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *QuotationResponseQuotation) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *QuotationResponseQuotation) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetTotalAmount

`func (o *QuotationResponseQuotation) GetTotalAmount() int32`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *QuotationResponseQuotation) GetTotalAmountOk() (*int32, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *QuotationResponseQuotation) SetTotalAmount(v int32)`

SetTotalAmount sets TotalAmount field to given value.


### GetTotalVat

`func (o *QuotationResponseQuotation) GetTotalVat() int32`

GetTotalVat returns the TotalVat field if non-nil, zero value otherwise.

### GetTotalVatOk

`func (o *QuotationResponseQuotation) GetTotalVatOk() (*int32, bool)`

GetTotalVatOk returns a tuple with the TotalVat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalVat

`func (o *QuotationResponseQuotation) SetTotalVat(v int32)`

SetTotalVat sets TotalVat field to given value.

### HasTotalVat

`func (o *QuotationResponseQuotation) HasTotalVat() bool`

HasTotalVat returns a boolean if a field has been set.

### GetSubTotal

`func (o *QuotationResponseQuotation) GetSubTotal() int32`

GetSubTotal returns the SubTotal field if non-nil, zero value otherwise.

### GetSubTotalOk

`func (o *QuotationResponseQuotation) GetSubTotalOk() (*int32, bool)`

GetSubTotalOk returns a tuple with the SubTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubTotal

`func (o *QuotationResponseQuotation) SetSubTotal(v int32)`

SetSubTotal sets SubTotal field to given value.

### HasSubTotal

`func (o *QuotationResponseQuotation) HasSubTotal() bool`

HasSubTotal returns a boolean if a field has been set.

### GetDescription

`func (o *QuotationResponseQuotation) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *QuotationResponseQuotation) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *QuotationResponseQuotation) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *QuotationResponseQuotation) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *QuotationResponseQuotation) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *QuotationResponseQuotation) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetQuotationStatus

`func (o *QuotationResponseQuotation) GetQuotationStatus() string`

GetQuotationStatus returns the QuotationStatus field if non-nil, zero value otherwise.

### GetQuotationStatusOk

`func (o *QuotationResponseQuotation) GetQuotationStatusOk() (*string, bool)`

GetQuotationStatusOk returns a tuple with the QuotationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotationStatus

`func (o *QuotationResponseQuotation) SetQuotationStatus(v string)`

SetQuotationStatus sets QuotationStatus field to given value.


### GetWebPublishedAt

`func (o *QuotationResponseQuotation) GetWebPublishedAt() string`

GetWebPublishedAt returns the WebPublishedAt field if non-nil, zero value otherwise.

### GetWebPublishedAtOk

`func (o *QuotationResponseQuotation) GetWebPublishedAtOk() (*string, bool)`

GetWebPublishedAtOk returns a tuple with the WebPublishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebPublishedAt

`func (o *QuotationResponseQuotation) SetWebPublishedAt(v string)`

SetWebPublishedAt sets WebPublishedAt field to given value.

### HasWebPublishedAt

`func (o *QuotationResponseQuotation) HasWebPublishedAt() bool`

HasWebPublishedAt returns a boolean if a field has been set.

### SetWebPublishedAtNil

`func (o *QuotationResponseQuotation) SetWebPublishedAtNil(b bool)`

 SetWebPublishedAtNil sets the value for WebPublishedAt to be an explicit nil

### UnsetWebPublishedAt
`func (o *QuotationResponseQuotation) UnsetWebPublishedAt()`

UnsetWebPublishedAt ensures that no value is present for WebPublishedAt, not even an explicit nil
### GetWebDownloadedAt

`func (o *QuotationResponseQuotation) GetWebDownloadedAt() string`

GetWebDownloadedAt returns the WebDownloadedAt field if non-nil, zero value otherwise.

### GetWebDownloadedAtOk

`func (o *QuotationResponseQuotation) GetWebDownloadedAtOk() (*string, bool)`

GetWebDownloadedAtOk returns a tuple with the WebDownloadedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebDownloadedAt

`func (o *QuotationResponseQuotation) SetWebDownloadedAt(v string)`

SetWebDownloadedAt sets WebDownloadedAt field to given value.

### HasWebDownloadedAt

`func (o *QuotationResponseQuotation) HasWebDownloadedAt() bool`

HasWebDownloadedAt returns a boolean if a field has been set.

### SetWebDownloadedAtNil

`func (o *QuotationResponseQuotation) SetWebDownloadedAtNil(b bool)`

 SetWebDownloadedAtNil sets the value for WebDownloadedAt to be an explicit nil

### UnsetWebDownloadedAt
`func (o *QuotationResponseQuotation) UnsetWebDownloadedAt()`

UnsetWebDownloadedAt ensures that no value is present for WebDownloadedAt, not even an explicit nil
### GetWebConfirmedAt

`func (o *QuotationResponseQuotation) GetWebConfirmedAt() string`

GetWebConfirmedAt returns the WebConfirmedAt field if non-nil, zero value otherwise.

### GetWebConfirmedAtOk

`func (o *QuotationResponseQuotation) GetWebConfirmedAtOk() (*string, bool)`

GetWebConfirmedAtOk returns a tuple with the WebConfirmedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebConfirmedAt

`func (o *QuotationResponseQuotation) SetWebConfirmedAt(v string)`

SetWebConfirmedAt sets WebConfirmedAt field to given value.

### HasWebConfirmedAt

`func (o *QuotationResponseQuotation) HasWebConfirmedAt() bool`

HasWebConfirmedAt returns a boolean if a field has been set.

### SetWebConfirmedAtNil

`func (o *QuotationResponseQuotation) SetWebConfirmedAtNil(b bool)`

 SetWebConfirmedAtNil sets the value for WebConfirmedAt to be an explicit nil

### UnsetWebConfirmedAt
`func (o *QuotationResponseQuotation) UnsetWebConfirmedAt()`

UnsetWebConfirmedAt ensures that no value is present for WebConfirmedAt, not even an explicit nil
### GetMailSentAt

`func (o *QuotationResponseQuotation) GetMailSentAt() string`

GetMailSentAt returns the MailSentAt field if non-nil, zero value otherwise.

### GetMailSentAtOk

`func (o *QuotationResponseQuotation) GetMailSentAtOk() (*string, bool)`

GetMailSentAtOk returns a tuple with the MailSentAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailSentAt

`func (o *QuotationResponseQuotation) SetMailSentAt(v string)`

SetMailSentAt sets MailSentAt field to given value.

### HasMailSentAt

`func (o *QuotationResponseQuotation) HasMailSentAt() bool`

HasMailSentAt returns a boolean if a field has been set.

### SetMailSentAtNil

`func (o *QuotationResponseQuotation) SetMailSentAtNil(b bool)`

 SetMailSentAtNil sets the value for MailSentAt to be an explicit nil

### UnsetMailSentAt
`func (o *QuotationResponseQuotation) UnsetMailSentAt()`

UnsetMailSentAt ensures that no value is present for MailSentAt, not even an explicit nil
### GetPartnerName

`func (o *QuotationResponseQuotation) GetPartnerName() string`

GetPartnerName returns the PartnerName field if non-nil, zero value otherwise.

### GetPartnerNameOk

`func (o *QuotationResponseQuotation) GetPartnerNameOk() (*string, bool)`

GetPartnerNameOk returns a tuple with the PartnerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerName

`func (o *QuotationResponseQuotation) SetPartnerName(v string)`

SetPartnerName sets PartnerName field to given value.

### HasPartnerName

`func (o *QuotationResponseQuotation) HasPartnerName() bool`

HasPartnerName returns a boolean if a field has been set.

### SetPartnerNameNil

`func (o *QuotationResponseQuotation) SetPartnerNameNil(b bool)`

 SetPartnerNameNil sets the value for PartnerName to be an explicit nil

### UnsetPartnerName
`func (o *QuotationResponseQuotation) UnsetPartnerName()`

UnsetPartnerName ensures that no value is present for PartnerName, not even an explicit nil
### GetPartnerDisplayName

`func (o *QuotationResponseQuotation) GetPartnerDisplayName() string`

GetPartnerDisplayName returns the PartnerDisplayName field if non-nil, zero value otherwise.

### GetPartnerDisplayNameOk

`func (o *QuotationResponseQuotation) GetPartnerDisplayNameOk() (*string, bool)`

GetPartnerDisplayNameOk returns a tuple with the PartnerDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerDisplayName

`func (o *QuotationResponseQuotation) SetPartnerDisplayName(v string)`

SetPartnerDisplayName sets PartnerDisplayName field to given value.

### HasPartnerDisplayName

`func (o *QuotationResponseQuotation) HasPartnerDisplayName() bool`

HasPartnerDisplayName returns a boolean if a field has been set.

### SetPartnerDisplayNameNil

`func (o *QuotationResponseQuotation) SetPartnerDisplayNameNil(b bool)`

 SetPartnerDisplayNameNil sets the value for PartnerDisplayName to be an explicit nil

### UnsetPartnerDisplayName
`func (o *QuotationResponseQuotation) UnsetPartnerDisplayName()`

UnsetPartnerDisplayName ensures that no value is present for PartnerDisplayName, not even an explicit nil
### GetPartnerTitle

`func (o *QuotationResponseQuotation) GetPartnerTitle() string`

GetPartnerTitle returns the PartnerTitle field if non-nil, zero value otherwise.

### GetPartnerTitleOk

`func (o *QuotationResponseQuotation) GetPartnerTitleOk() (*string, bool)`

GetPartnerTitleOk returns a tuple with the PartnerTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerTitle

`func (o *QuotationResponseQuotation) SetPartnerTitle(v string)`

SetPartnerTitle sets PartnerTitle field to given value.


### SetPartnerTitleNil

`func (o *QuotationResponseQuotation) SetPartnerTitleNil(b bool)`

 SetPartnerTitleNil sets the value for PartnerTitle to be an explicit nil

### UnsetPartnerTitle
`func (o *QuotationResponseQuotation) UnsetPartnerTitle()`

UnsetPartnerTitle ensures that no value is present for PartnerTitle, not even an explicit nil
### GetPartnerZipcode

`func (o *QuotationResponseQuotation) GetPartnerZipcode() string`

GetPartnerZipcode returns the PartnerZipcode field if non-nil, zero value otherwise.

### GetPartnerZipcodeOk

`func (o *QuotationResponseQuotation) GetPartnerZipcodeOk() (*string, bool)`

GetPartnerZipcodeOk returns a tuple with the PartnerZipcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerZipcode

`func (o *QuotationResponseQuotation) SetPartnerZipcode(v string)`

SetPartnerZipcode sets PartnerZipcode field to given value.

### HasPartnerZipcode

`func (o *QuotationResponseQuotation) HasPartnerZipcode() bool`

HasPartnerZipcode returns a boolean if a field has been set.

### SetPartnerZipcodeNil

`func (o *QuotationResponseQuotation) SetPartnerZipcodeNil(b bool)`

 SetPartnerZipcodeNil sets the value for PartnerZipcode to be an explicit nil

### UnsetPartnerZipcode
`func (o *QuotationResponseQuotation) UnsetPartnerZipcode()`

UnsetPartnerZipcode ensures that no value is present for PartnerZipcode, not even an explicit nil
### GetPartnerPrefectureCode

`func (o *QuotationResponseQuotation) GetPartnerPrefectureCode() int32`

GetPartnerPrefectureCode returns the PartnerPrefectureCode field if non-nil, zero value otherwise.

### GetPartnerPrefectureCodeOk

`func (o *QuotationResponseQuotation) GetPartnerPrefectureCodeOk() (*int32, bool)`

GetPartnerPrefectureCodeOk returns a tuple with the PartnerPrefectureCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerPrefectureCode

`func (o *QuotationResponseQuotation) SetPartnerPrefectureCode(v int32)`

SetPartnerPrefectureCode sets PartnerPrefectureCode field to given value.

### HasPartnerPrefectureCode

`func (o *QuotationResponseQuotation) HasPartnerPrefectureCode() bool`

HasPartnerPrefectureCode returns a boolean if a field has been set.

### SetPartnerPrefectureCodeNil

`func (o *QuotationResponseQuotation) SetPartnerPrefectureCodeNil(b bool)`

 SetPartnerPrefectureCodeNil sets the value for PartnerPrefectureCode to be an explicit nil

### UnsetPartnerPrefectureCode
`func (o *QuotationResponseQuotation) UnsetPartnerPrefectureCode()`

UnsetPartnerPrefectureCode ensures that no value is present for PartnerPrefectureCode, not even an explicit nil
### GetPartnerPrefectureName

`func (o *QuotationResponseQuotation) GetPartnerPrefectureName() string`

GetPartnerPrefectureName returns the PartnerPrefectureName field if non-nil, zero value otherwise.

### GetPartnerPrefectureNameOk

`func (o *QuotationResponseQuotation) GetPartnerPrefectureNameOk() (*string, bool)`

GetPartnerPrefectureNameOk returns a tuple with the PartnerPrefectureName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerPrefectureName

`func (o *QuotationResponseQuotation) SetPartnerPrefectureName(v string)`

SetPartnerPrefectureName sets PartnerPrefectureName field to given value.

### HasPartnerPrefectureName

`func (o *QuotationResponseQuotation) HasPartnerPrefectureName() bool`

HasPartnerPrefectureName returns a boolean if a field has been set.

### SetPartnerPrefectureNameNil

`func (o *QuotationResponseQuotation) SetPartnerPrefectureNameNil(b bool)`

 SetPartnerPrefectureNameNil sets the value for PartnerPrefectureName to be an explicit nil

### UnsetPartnerPrefectureName
`func (o *QuotationResponseQuotation) UnsetPartnerPrefectureName()`

UnsetPartnerPrefectureName ensures that no value is present for PartnerPrefectureName, not even an explicit nil
### GetPartnerAddress1

`func (o *QuotationResponseQuotation) GetPartnerAddress1() string`

GetPartnerAddress1 returns the PartnerAddress1 field if non-nil, zero value otherwise.

### GetPartnerAddress1Ok

`func (o *QuotationResponseQuotation) GetPartnerAddress1Ok() (*string, bool)`

GetPartnerAddress1Ok returns a tuple with the PartnerAddress1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerAddress1

`func (o *QuotationResponseQuotation) SetPartnerAddress1(v string)`

SetPartnerAddress1 sets PartnerAddress1 field to given value.

### HasPartnerAddress1

`func (o *QuotationResponseQuotation) HasPartnerAddress1() bool`

HasPartnerAddress1 returns a boolean if a field has been set.

### SetPartnerAddress1Nil

`func (o *QuotationResponseQuotation) SetPartnerAddress1Nil(b bool)`

 SetPartnerAddress1Nil sets the value for PartnerAddress1 to be an explicit nil

### UnsetPartnerAddress1
`func (o *QuotationResponseQuotation) UnsetPartnerAddress1()`

UnsetPartnerAddress1 ensures that no value is present for PartnerAddress1, not even an explicit nil
### GetPartnerAddress2

`func (o *QuotationResponseQuotation) GetPartnerAddress2() string`

GetPartnerAddress2 returns the PartnerAddress2 field if non-nil, zero value otherwise.

### GetPartnerAddress2Ok

`func (o *QuotationResponseQuotation) GetPartnerAddress2Ok() (*string, bool)`

GetPartnerAddress2Ok returns a tuple with the PartnerAddress2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerAddress2

`func (o *QuotationResponseQuotation) SetPartnerAddress2(v string)`

SetPartnerAddress2 sets PartnerAddress2 field to given value.

### HasPartnerAddress2

`func (o *QuotationResponseQuotation) HasPartnerAddress2() bool`

HasPartnerAddress2 returns a boolean if a field has been set.

### SetPartnerAddress2Nil

`func (o *QuotationResponseQuotation) SetPartnerAddress2Nil(b bool)`

 SetPartnerAddress2Nil sets the value for PartnerAddress2 to be an explicit nil

### UnsetPartnerAddress2
`func (o *QuotationResponseQuotation) UnsetPartnerAddress2()`

UnsetPartnerAddress2 ensures that no value is present for PartnerAddress2, not even an explicit nil
### GetPartnerContactInfo

`func (o *QuotationResponseQuotation) GetPartnerContactInfo() string`

GetPartnerContactInfo returns the PartnerContactInfo field if non-nil, zero value otherwise.

### GetPartnerContactInfoOk

`func (o *QuotationResponseQuotation) GetPartnerContactInfoOk() (*string, bool)`

GetPartnerContactInfoOk returns a tuple with the PartnerContactInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerContactInfo

`func (o *QuotationResponseQuotation) SetPartnerContactInfo(v string)`

SetPartnerContactInfo sets PartnerContactInfo field to given value.

### HasPartnerContactInfo

`func (o *QuotationResponseQuotation) HasPartnerContactInfo() bool`

HasPartnerContactInfo returns a boolean if a field has been set.

### SetPartnerContactInfoNil

`func (o *QuotationResponseQuotation) SetPartnerContactInfoNil(b bool)`

 SetPartnerContactInfoNil sets the value for PartnerContactInfo to be an explicit nil

### UnsetPartnerContactInfo
`func (o *QuotationResponseQuotation) UnsetPartnerContactInfo()`

UnsetPartnerContactInfo ensures that no value is present for PartnerContactInfo, not even an explicit nil
### GetCompanyName

`func (o *QuotationResponseQuotation) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *QuotationResponseQuotation) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *QuotationResponseQuotation) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.


### GetCompanyZipcode

`func (o *QuotationResponseQuotation) GetCompanyZipcode() string`

GetCompanyZipcode returns the CompanyZipcode field if non-nil, zero value otherwise.

### GetCompanyZipcodeOk

`func (o *QuotationResponseQuotation) GetCompanyZipcodeOk() (*string, bool)`

GetCompanyZipcodeOk returns a tuple with the CompanyZipcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyZipcode

`func (o *QuotationResponseQuotation) SetCompanyZipcode(v string)`

SetCompanyZipcode sets CompanyZipcode field to given value.

### HasCompanyZipcode

`func (o *QuotationResponseQuotation) HasCompanyZipcode() bool`

HasCompanyZipcode returns a boolean if a field has been set.

### SetCompanyZipcodeNil

`func (o *QuotationResponseQuotation) SetCompanyZipcodeNil(b bool)`

 SetCompanyZipcodeNil sets the value for CompanyZipcode to be an explicit nil

### UnsetCompanyZipcode
`func (o *QuotationResponseQuotation) UnsetCompanyZipcode()`

UnsetCompanyZipcode ensures that no value is present for CompanyZipcode, not even an explicit nil
### GetCompanyPrefectureCode

`func (o *QuotationResponseQuotation) GetCompanyPrefectureCode() int32`

GetCompanyPrefectureCode returns the CompanyPrefectureCode field if non-nil, zero value otherwise.

### GetCompanyPrefectureCodeOk

`func (o *QuotationResponseQuotation) GetCompanyPrefectureCodeOk() (*int32, bool)`

GetCompanyPrefectureCodeOk returns a tuple with the CompanyPrefectureCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyPrefectureCode

`func (o *QuotationResponseQuotation) SetCompanyPrefectureCode(v int32)`

SetCompanyPrefectureCode sets CompanyPrefectureCode field to given value.

### HasCompanyPrefectureCode

`func (o *QuotationResponseQuotation) HasCompanyPrefectureCode() bool`

HasCompanyPrefectureCode returns a boolean if a field has been set.

### SetCompanyPrefectureCodeNil

`func (o *QuotationResponseQuotation) SetCompanyPrefectureCodeNil(b bool)`

 SetCompanyPrefectureCodeNil sets the value for CompanyPrefectureCode to be an explicit nil

### UnsetCompanyPrefectureCode
`func (o *QuotationResponseQuotation) UnsetCompanyPrefectureCode()`

UnsetCompanyPrefectureCode ensures that no value is present for CompanyPrefectureCode, not even an explicit nil
### GetCompanyPrefectureName

`func (o *QuotationResponseQuotation) GetCompanyPrefectureName() string`

GetCompanyPrefectureName returns the CompanyPrefectureName field if non-nil, zero value otherwise.

### GetCompanyPrefectureNameOk

`func (o *QuotationResponseQuotation) GetCompanyPrefectureNameOk() (*string, bool)`

GetCompanyPrefectureNameOk returns a tuple with the CompanyPrefectureName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyPrefectureName

`func (o *QuotationResponseQuotation) SetCompanyPrefectureName(v string)`

SetCompanyPrefectureName sets CompanyPrefectureName field to given value.

### HasCompanyPrefectureName

`func (o *QuotationResponseQuotation) HasCompanyPrefectureName() bool`

HasCompanyPrefectureName returns a boolean if a field has been set.

### SetCompanyPrefectureNameNil

`func (o *QuotationResponseQuotation) SetCompanyPrefectureNameNil(b bool)`

 SetCompanyPrefectureNameNil sets the value for CompanyPrefectureName to be an explicit nil

### UnsetCompanyPrefectureName
`func (o *QuotationResponseQuotation) UnsetCompanyPrefectureName()`

UnsetCompanyPrefectureName ensures that no value is present for CompanyPrefectureName, not even an explicit nil
### GetCompanyAddress1

`func (o *QuotationResponseQuotation) GetCompanyAddress1() string`

GetCompanyAddress1 returns the CompanyAddress1 field if non-nil, zero value otherwise.

### GetCompanyAddress1Ok

`func (o *QuotationResponseQuotation) GetCompanyAddress1Ok() (*string, bool)`

GetCompanyAddress1Ok returns a tuple with the CompanyAddress1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyAddress1

`func (o *QuotationResponseQuotation) SetCompanyAddress1(v string)`

SetCompanyAddress1 sets CompanyAddress1 field to given value.

### HasCompanyAddress1

`func (o *QuotationResponseQuotation) HasCompanyAddress1() bool`

HasCompanyAddress1 returns a boolean if a field has been set.

### SetCompanyAddress1Nil

`func (o *QuotationResponseQuotation) SetCompanyAddress1Nil(b bool)`

 SetCompanyAddress1Nil sets the value for CompanyAddress1 to be an explicit nil

### UnsetCompanyAddress1
`func (o *QuotationResponseQuotation) UnsetCompanyAddress1()`

UnsetCompanyAddress1 ensures that no value is present for CompanyAddress1, not even an explicit nil
### GetCompanyAddress2

`func (o *QuotationResponseQuotation) GetCompanyAddress2() string`

GetCompanyAddress2 returns the CompanyAddress2 field if non-nil, zero value otherwise.

### GetCompanyAddress2Ok

`func (o *QuotationResponseQuotation) GetCompanyAddress2Ok() (*string, bool)`

GetCompanyAddress2Ok returns a tuple with the CompanyAddress2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyAddress2

`func (o *QuotationResponseQuotation) SetCompanyAddress2(v string)`

SetCompanyAddress2 sets CompanyAddress2 field to given value.

### HasCompanyAddress2

`func (o *QuotationResponseQuotation) HasCompanyAddress2() bool`

HasCompanyAddress2 returns a boolean if a field has been set.

### SetCompanyAddress2Nil

`func (o *QuotationResponseQuotation) SetCompanyAddress2Nil(b bool)`

 SetCompanyAddress2Nil sets the value for CompanyAddress2 to be an explicit nil

### UnsetCompanyAddress2
`func (o *QuotationResponseQuotation) UnsetCompanyAddress2()`

UnsetCompanyAddress2 ensures that no value is present for CompanyAddress2, not even an explicit nil
### GetCompanyContactInfo

`func (o *QuotationResponseQuotation) GetCompanyContactInfo() string`

GetCompanyContactInfo returns the CompanyContactInfo field if non-nil, zero value otherwise.

### GetCompanyContactInfoOk

`func (o *QuotationResponseQuotation) GetCompanyContactInfoOk() (*string, bool)`

GetCompanyContactInfoOk returns a tuple with the CompanyContactInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyContactInfo

`func (o *QuotationResponseQuotation) SetCompanyContactInfo(v string)`

SetCompanyContactInfo sets CompanyContactInfo field to given value.

### HasCompanyContactInfo

`func (o *QuotationResponseQuotation) HasCompanyContactInfo() bool`

HasCompanyContactInfo returns a boolean if a field has been set.

### SetCompanyContactInfoNil

`func (o *QuotationResponseQuotation) SetCompanyContactInfoNil(b bool)`

 SetCompanyContactInfoNil sets the value for CompanyContactInfo to be an explicit nil

### UnsetCompanyContactInfo
`func (o *QuotationResponseQuotation) UnsetCompanyContactInfo()`

UnsetCompanyContactInfo ensures that no value is present for CompanyContactInfo, not even an explicit nil
### GetMessage

`func (o *QuotationResponseQuotation) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *QuotationResponseQuotation) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *QuotationResponseQuotation) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *QuotationResponseQuotation) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *QuotationResponseQuotation) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *QuotationResponseQuotation) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetNotes

`func (o *QuotationResponseQuotation) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *QuotationResponseQuotation) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *QuotationResponseQuotation) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *QuotationResponseQuotation) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *QuotationResponseQuotation) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *QuotationResponseQuotation) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetQuotationLayout

`func (o *QuotationResponseQuotation) GetQuotationLayout() string`

GetQuotationLayout returns the QuotationLayout field if non-nil, zero value otherwise.

### GetQuotationLayoutOk

`func (o *QuotationResponseQuotation) GetQuotationLayoutOk() (*string, bool)`

GetQuotationLayoutOk returns a tuple with the QuotationLayout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotationLayout

`func (o *QuotationResponseQuotation) SetQuotationLayout(v string)`

SetQuotationLayout sets QuotationLayout field to given value.


### GetTaxEntryMethod

`func (o *QuotationResponseQuotation) GetTaxEntryMethod() string`

GetTaxEntryMethod returns the TaxEntryMethod field if non-nil, zero value otherwise.

### GetTaxEntryMethodOk

`func (o *QuotationResponseQuotation) GetTaxEntryMethodOk() (*string, bool)`

GetTaxEntryMethodOk returns a tuple with the TaxEntryMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxEntryMethod

`func (o *QuotationResponseQuotation) SetTaxEntryMethod(v string)`

SetTaxEntryMethod sets TaxEntryMethod field to given value.


### GetQuotationContents

`func (o *QuotationResponseQuotation) GetQuotationContents() []QuotationIndexResponseQuotationContents`

GetQuotationContents returns the QuotationContents field if non-nil, zero value otherwise.

### GetQuotationContentsOk

`func (o *QuotationResponseQuotation) GetQuotationContentsOk() (*[]QuotationIndexResponseQuotationContents, bool)`

GetQuotationContentsOk returns a tuple with the QuotationContents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotationContents

`func (o *QuotationResponseQuotation) SetQuotationContents(v []QuotationIndexResponseQuotationContents)`

SetQuotationContents sets QuotationContents field to given value.

### HasQuotationContents

`func (o *QuotationResponseQuotation) HasQuotationContents() bool`

HasQuotationContents returns a boolean if a field has been set.

### GetTotalAmountPerVatRate

`func (o *QuotationResponseQuotation) GetTotalAmountPerVatRate() InvoiceIndexResponseTotalAmountPerVatRate`

GetTotalAmountPerVatRate returns the TotalAmountPerVatRate field if non-nil, zero value otherwise.

### GetTotalAmountPerVatRateOk

`func (o *QuotationResponseQuotation) GetTotalAmountPerVatRateOk() (*InvoiceIndexResponseTotalAmountPerVatRate, bool)`

GetTotalAmountPerVatRateOk returns a tuple with the TotalAmountPerVatRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmountPerVatRate

`func (o *QuotationResponseQuotation) SetTotalAmountPerVatRate(v InvoiceIndexResponseTotalAmountPerVatRate)`

SetTotalAmountPerVatRate sets TotalAmountPerVatRate field to given value.


### GetRelatedInvoiceId

`func (o *QuotationResponseQuotation) GetRelatedInvoiceId() int32`

GetRelatedInvoiceId returns the RelatedInvoiceId field if non-nil, zero value otherwise.

### GetRelatedInvoiceIdOk

`func (o *QuotationResponseQuotation) GetRelatedInvoiceIdOk() (*int32, bool)`

GetRelatedInvoiceIdOk returns a tuple with the RelatedInvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedInvoiceId

`func (o *QuotationResponseQuotation) SetRelatedInvoiceId(v int32)`

SetRelatedInvoiceId sets RelatedInvoiceId field to given value.

### HasRelatedInvoiceId

`func (o *QuotationResponseQuotation) HasRelatedInvoiceId() bool`

HasRelatedInvoiceId returns a boolean if a field has been set.

### SetRelatedInvoiceIdNil

`func (o *QuotationResponseQuotation) SetRelatedInvoiceIdNil(b bool)`

 SetRelatedInvoiceIdNil sets the value for RelatedInvoiceId to be an explicit nil

### UnsetRelatedInvoiceId
`func (o *QuotationResponseQuotation) UnsetRelatedInvoiceId()`

UnsetRelatedInvoiceId ensures that no value is present for RelatedInvoiceId, not even an explicit nil
### GetRelatedQuotationIds

`func (o *QuotationResponseQuotation) GetRelatedQuotationIds() []int32`

GetRelatedQuotationIds returns the RelatedQuotationIds field if non-nil, zero value otherwise.

### GetRelatedQuotationIdsOk

`func (o *QuotationResponseQuotation) GetRelatedQuotationIdsOk() (*[]int32, bool)`

GetRelatedQuotationIdsOk returns a tuple with the RelatedQuotationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedQuotationIds

`func (o *QuotationResponseQuotation) SetRelatedQuotationIds(v []int32)`

SetRelatedQuotationIds sets RelatedQuotationIds field to given value.

### HasRelatedQuotationIds

`func (o *QuotationResponseQuotation) HasRelatedQuotationIds() bool`

HasRelatedQuotationIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


