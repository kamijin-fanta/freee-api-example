# ExpenseApplicationUpdateParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyId** | **int32** | 事業所ID | 
**Title** | **string** | 申請タイトル (250文字以内) | 
**IssueDate** | **string** | 申請日 (yyyy-mm-dd) | 
**Description** | Pointer to **string** | 備考 (10000文字以内) | [optional] 
**EditableOnWeb** | Pointer to **bool** | 会計freeeのWeb画面で申請内容を編集可能にするかどうか&lt;br&gt; falseの場合は、Web画面での項目行の追加／削除・金額の編集が出来なくなりますが、APIでの編集は可能です。&lt;br&gt; デフォルトはfalseです。  | [optional] 
**SectionId** | Pointer to **int32** | 部門ID | [optional] 
**TagIds** | Pointer to **[]int32** | メモタグID | [optional] 
**ExpenseApplicationLines** | [**[]ExpenseApplicationUpdateParamsExpenseApplicationLines**](ExpenseApplicationUpdateParamsExpenseApplicationLines.md) |  | 
**ApprovalFlowRouteId** | Pointer to **int32** | 申請経路ID&lt;br&gt; &lt;ul&gt;     &lt;li&gt;経費申請のステータスを申請中として作成する場合は、必ず指定してください。&lt;/li&gt;     &lt;li&gt;指定する申請経路IDは、申請経路APIを利用して取得してください。&lt;/li&gt;     &lt;li&gt;         未指定の場合は、基本経路を設定している事業所では基本経路が、基本経路を設定していない事業所では利用可能な申請経路の中から最初の申請経路が自動的に使用されます。         &lt;ul&gt;            &lt;li&gt;意図しない申請経路を持った経費申請の作成を防ぐために、使用する申請経路IDを指定することを推奨します。&lt;/li&gt;         &lt;/ul&gt;     &lt;/li&gt;     &lt;li&gt;         ベーシックプランの事業所では以下のデフォルトで用意された申請経路のみ指定できます         &lt;ul&gt;         &lt;li&gt;指定なし&lt;/li&gt;         &lt;li&gt;承認者を指定&lt;/li&gt;         &lt;/ul&gt;     &lt;/li&gt; &lt;/ul&gt;  | [optional] 
**ApproverId** | Pointer to **int32** | 承認者のユーザーID&lt;br&gt; 指定する承認者のユーザーIDは、申請経路APIを利用して取得してください。  | [optional] 
**Draft** | Pointer to **bool** | 経費申請のステータス&lt;br&gt; falseを指定した時は申請中（in_progress）で経費申請を更新します。&lt;br&gt; trueを指定した時は下書き（draft）で経費申請を更新します。&lt;br&gt; 未指定の時は下書きとみなして経費申請を更新します。  | [optional] 
**Segment1TagId** | Pointer to **int32** | セグメント１ID(法人向けプロフェッショナル, 法人向けエンタープライズプラン)&lt;br&gt; セグメントタグ一覧APIを利用して取得してください。&lt;br&gt; &lt;a href&#x3D;\&quot;https://support.freee.co.jp/hc/ja/articles/360020679611-%E3%82%BB%E3%82%B0%E3%83%A1%E3%83%B3%E3%83%88-%E5%88%86%E6%9E%90%E7%94%A8%E3%82%BF%E3%82%B0-%E3%81%AE%E8%A8%AD%E5%AE%9A\&quot; target&#x3D;\&quot;_blank\&quot;&gt;セグメント（分析用タグ）の設定&lt;/a&gt;&lt;br&gt;  | [optional] 
**Segment2TagId** | Pointer to **int32** | セグメント２ID(法人向け エンタープライズプラン)&lt;br&gt; セグメントタグ一覧APIを利用して取得してください。&lt;br&gt; &lt;a href&#x3D;\&quot;https://support.freee.co.jp/hc/ja/articles/360020679611-%E3%82%BB%E3%82%B0%E3%83%A1%E3%83%B3%E3%83%88-%E5%88%86%E6%9E%90%E7%94%A8%E3%82%BF%E3%82%B0-%E3%81%AE%E8%A8%AD%E5%AE%9A\&quot; target&#x3D;\&quot;_blank\&quot;&gt;セグメント（分析用タグ）の設定&lt;/a&gt;&lt;br&gt;  | [optional] 
**Segment3TagId** | Pointer to **int32** | セグメント３ID(法人向け エンタープライズプラン)&lt;br&gt; セグメントタグ一覧APIを利用して取得してください。&lt;br&gt; &lt;a href&#x3D;\&quot;https://support.freee.co.jp/hc/ja/articles/360020679611-%E3%82%BB%E3%82%B0%E3%83%A1%E3%83%B3%E3%83%88-%E5%88%86%E6%9E%90%E7%94%A8%E3%82%BF%E3%82%B0-%E3%81%AE%E8%A8%AD%E5%AE%9A\&quot; target&#x3D;\&quot;_blank\&quot;&gt;セグメント（分析用タグ）の設定&lt;/a&gt;&lt;br&gt;  | [optional] 

## Methods

### NewExpenseApplicationUpdateParams

`func NewExpenseApplicationUpdateParams(companyId int32, title string, issueDate string, expenseApplicationLines []ExpenseApplicationUpdateParamsExpenseApplicationLines, ) *ExpenseApplicationUpdateParams`

NewExpenseApplicationUpdateParams instantiates a new ExpenseApplicationUpdateParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExpenseApplicationUpdateParamsWithDefaults

`func NewExpenseApplicationUpdateParamsWithDefaults() *ExpenseApplicationUpdateParams`

NewExpenseApplicationUpdateParamsWithDefaults instantiates a new ExpenseApplicationUpdateParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyId

`func (o *ExpenseApplicationUpdateParams) GetCompanyId() int32`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *ExpenseApplicationUpdateParams) GetCompanyIdOk() (*int32, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *ExpenseApplicationUpdateParams) SetCompanyId(v int32)`

SetCompanyId sets CompanyId field to given value.


### GetTitle

`func (o *ExpenseApplicationUpdateParams) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ExpenseApplicationUpdateParams) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ExpenseApplicationUpdateParams) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetIssueDate

`func (o *ExpenseApplicationUpdateParams) GetIssueDate() string`

GetIssueDate returns the IssueDate field if non-nil, zero value otherwise.

### GetIssueDateOk

`func (o *ExpenseApplicationUpdateParams) GetIssueDateOk() (*string, bool)`

GetIssueDateOk returns a tuple with the IssueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueDate

`func (o *ExpenseApplicationUpdateParams) SetIssueDate(v string)`

SetIssueDate sets IssueDate field to given value.


### GetDescription

`func (o *ExpenseApplicationUpdateParams) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ExpenseApplicationUpdateParams) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ExpenseApplicationUpdateParams) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ExpenseApplicationUpdateParams) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEditableOnWeb

`func (o *ExpenseApplicationUpdateParams) GetEditableOnWeb() bool`

GetEditableOnWeb returns the EditableOnWeb field if non-nil, zero value otherwise.

### GetEditableOnWebOk

`func (o *ExpenseApplicationUpdateParams) GetEditableOnWebOk() (*bool, bool)`

GetEditableOnWebOk returns a tuple with the EditableOnWeb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditableOnWeb

`func (o *ExpenseApplicationUpdateParams) SetEditableOnWeb(v bool)`

SetEditableOnWeb sets EditableOnWeb field to given value.

### HasEditableOnWeb

`func (o *ExpenseApplicationUpdateParams) HasEditableOnWeb() bool`

HasEditableOnWeb returns a boolean if a field has been set.

### GetSectionId

`func (o *ExpenseApplicationUpdateParams) GetSectionId() int32`

GetSectionId returns the SectionId field if non-nil, zero value otherwise.

### GetSectionIdOk

`func (o *ExpenseApplicationUpdateParams) GetSectionIdOk() (*int32, bool)`

GetSectionIdOk returns a tuple with the SectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionId

`func (o *ExpenseApplicationUpdateParams) SetSectionId(v int32)`

SetSectionId sets SectionId field to given value.

### HasSectionId

`func (o *ExpenseApplicationUpdateParams) HasSectionId() bool`

HasSectionId returns a boolean if a field has been set.

### GetTagIds

`func (o *ExpenseApplicationUpdateParams) GetTagIds() []int32`

GetTagIds returns the TagIds field if non-nil, zero value otherwise.

### GetTagIdsOk

`func (o *ExpenseApplicationUpdateParams) GetTagIdsOk() (*[]int32, bool)`

GetTagIdsOk returns a tuple with the TagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagIds

`func (o *ExpenseApplicationUpdateParams) SetTagIds(v []int32)`

SetTagIds sets TagIds field to given value.

### HasTagIds

`func (o *ExpenseApplicationUpdateParams) HasTagIds() bool`

HasTagIds returns a boolean if a field has been set.

### GetExpenseApplicationLines

`func (o *ExpenseApplicationUpdateParams) GetExpenseApplicationLines() []ExpenseApplicationUpdateParamsExpenseApplicationLines`

GetExpenseApplicationLines returns the ExpenseApplicationLines field if non-nil, zero value otherwise.

### GetExpenseApplicationLinesOk

`func (o *ExpenseApplicationUpdateParams) GetExpenseApplicationLinesOk() (*[]ExpenseApplicationUpdateParamsExpenseApplicationLines, bool)`

GetExpenseApplicationLinesOk returns a tuple with the ExpenseApplicationLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpenseApplicationLines

`func (o *ExpenseApplicationUpdateParams) SetExpenseApplicationLines(v []ExpenseApplicationUpdateParamsExpenseApplicationLines)`

SetExpenseApplicationLines sets ExpenseApplicationLines field to given value.


### GetApprovalFlowRouteId

`func (o *ExpenseApplicationUpdateParams) GetApprovalFlowRouteId() int32`

GetApprovalFlowRouteId returns the ApprovalFlowRouteId field if non-nil, zero value otherwise.

### GetApprovalFlowRouteIdOk

`func (o *ExpenseApplicationUpdateParams) GetApprovalFlowRouteIdOk() (*int32, bool)`

GetApprovalFlowRouteIdOk returns a tuple with the ApprovalFlowRouteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalFlowRouteId

`func (o *ExpenseApplicationUpdateParams) SetApprovalFlowRouteId(v int32)`

SetApprovalFlowRouteId sets ApprovalFlowRouteId field to given value.

### HasApprovalFlowRouteId

`func (o *ExpenseApplicationUpdateParams) HasApprovalFlowRouteId() bool`

HasApprovalFlowRouteId returns a boolean if a field has been set.

### GetApproverId

`func (o *ExpenseApplicationUpdateParams) GetApproverId() int32`

GetApproverId returns the ApproverId field if non-nil, zero value otherwise.

### GetApproverIdOk

`func (o *ExpenseApplicationUpdateParams) GetApproverIdOk() (*int32, bool)`

GetApproverIdOk returns a tuple with the ApproverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproverId

`func (o *ExpenseApplicationUpdateParams) SetApproverId(v int32)`

SetApproverId sets ApproverId field to given value.

### HasApproverId

`func (o *ExpenseApplicationUpdateParams) HasApproverId() bool`

HasApproverId returns a boolean if a field has been set.

### GetDraft

`func (o *ExpenseApplicationUpdateParams) GetDraft() bool`

GetDraft returns the Draft field if non-nil, zero value otherwise.

### GetDraftOk

`func (o *ExpenseApplicationUpdateParams) GetDraftOk() (*bool, bool)`

GetDraftOk returns a tuple with the Draft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraft

`func (o *ExpenseApplicationUpdateParams) SetDraft(v bool)`

SetDraft sets Draft field to given value.

### HasDraft

`func (o *ExpenseApplicationUpdateParams) HasDraft() bool`

HasDraft returns a boolean if a field has been set.

### GetSegment1TagId

`func (o *ExpenseApplicationUpdateParams) GetSegment1TagId() int32`

GetSegment1TagId returns the Segment1TagId field if non-nil, zero value otherwise.

### GetSegment1TagIdOk

`func (o *ExpenseApplicationUpdateParams) GetSegment1TagIdOk() (*int32, bool)`

GetSegment1TagIdOk returns a tuple with the Segment1TagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment1TagId

`func (o *ExpenseApplicationUpdateParams) SetSegment1TagId(v int32)`

SetSegment1TagId sets Segment1TagId field to given value.

### HasSegment1TagId

`func (o *ExpenseApplicationUpdateParams) HasSegment1TagId() bool`

HasSegment1TagId returns a boolean if a field has been set.

### GetSegment2TagId

`func (o *ExpenseApplicationUpdateParams) GetSegment2TagId() int32`

GetSegment2TagId returns the Segment2TagId field if non-nil, zero value otherwise.

### GetSegment2TagIdOk

`func (o *ExpenseApplicationUpdateParams) GetSegment2TagIdOk() (*int32, bool)`

GetSegment2TagIdOk returns a tuple with the Segment2TagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment2TagId

`func (o *ExpenseApplicationUpdateParams) SetSegment2TagId(v int32)`

SetSegment2TagId sets Segment2TagId field to given value.

### HasSegment2TagId

`func (o *ExpenseApplicationUpdateParams) HasSegment2TagId() bool`

HasSegment2TagId returns a boolean if a field has been set.

### GetSegment3TagId

`func (o *ExpenseApplicationUpdateParams) GetSegment3TagId() int32`

GetSegment3TagId returns the Segment3TagId field if non-nil, zero value otherwise.

### GetSegment3TagIdOk

`func (o *ExpenseApplicationUpdateParams) GetSegment3TagIdOk() (*int32, bool)`

GetSegment3TagIdOk returns a tuple with the Segment3TagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment3TagId

`func (o *ExpenseApplicationUpdateParams) SetSegment3TagId(v int32)`

SetSegment3TagId sets Segment3TagId field to given value.

### HasSegment3TagId

`func (o *ExpenseApplicationUpdateParams) HasSegment3TagId() bool`

HasSegment3TagId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


