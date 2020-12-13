# ExpenseApplicationCreateParams

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
**ExpenseApplicationLines** | [**[]ExpenseApplicationCreateParamsExpenseApplicationLines**](ExpenseApplicationCreateParamsExpenseApplicationLines.md) |  | 
**ApprovalFlowRouteId** | Pointer to **int32** | 申請経路ID&lt;br&gt; &lt;ul&gt;     &lt;li&gt;経費申請のステータスを申請中として作成する場合は、必ず指定してください。&lt;/li&gt;     &lt;li&gt;指定する申請経路IDは、申請経路APIを利用して取得してください。&lt;/li&gt;     &lt;li&gt;         未指定の場合は、基本経路を設定している事業所では基本経路が、基本経路を設定していない事業所では利用可能な申請経路の中から最初の申請経路が自動的に使用されます。         &lt;ul&gt;            &lt;li&gt;意図しない申請経路を持った経費申請の作成を防ぐために、使用する申請経路IDを指定することを推奨します。&lt;/li&gt;         &lt;/ul&gt;     &lt;/li&gt;     &lt;li&gt;         ベーシックプランの事業所では以下のデフォルトで用意された申請経路のみ指定できます         &lt;ul&gt;         &lt;li&gt;指定なし&lt;/li&gt;         &lt;li&gt;承認者を指定&lt;/li&gt;         &lt;/ul&gt;     &lt;/li&gt; &lt;/ul&gt;  | [optional] 
**ApproverId** | Pointer to **int32** | 承認者のユーザーID&lt;br&gt; 「承認者を指定」の経路を申請経路として使用する場合に指定してください。&lt;br&gt; 指定する承認者のユーザーIDは、申請経路APIを利用して取得してください。  | [optional] 
**Draft** | Pointer to **bool** | 経費申請のステータス&lt;br&gt; falseを指定した時は申請中（in_progress）で経費申請を作成します。&lt;br&gt; trueを指定した時は下書き（draft）で経費申請を作成します。&lt;br&gt; 未指定の時は下書きとみなして経費申請を作成します。  | [optional] 
**ParentId** | Pointer to **int32** | 親申請ID(法人向け エンタープライズプラン)&lt;br&gt; &lt;ul&gt;   &lt;li&gt;承認済みの既存各種申請IDのみ指定可能です。&lt;/li&gt;   &lt;li&gt;各種申請一覧APIを利用して取得してください。&lt;/li&gt; &lt;/ul&gt;  | [optional] 
**Segment1TagId** | Pointer to **int32** | セグメント１ID(法人向けプロフェッショナル, 法人向けエンタープライズプラン)&lt;br&gt; セグメントタグ一覧APIを利用して取得してください。&lt;br&gt; &lt;a href&#x3D;\&quot;https://support.freee.co.jp/hc/ja/articles/360020679611-%E3%82%BB%E3%82%B0%E3%83%A1%E3%83%B3%E3%83%88-%E5%88%86%E6%9E%90%E7%94%A8%E3%82%BF%E3%82%B0-%E3%81%AE%E8%A8%AD%E5%AE%9A\&quot; target&#x3D;\&quot;_blank\&quot;&gt;セグメント（分析用タグ）の設定&lt;/a&gt;&lt;br&gt;  | [optional] 
**Segment2TagId** | Pointer to **int32** | セグメント２ID(法人向け エンタープライズプラン)&lt;br&gt; セグメントタグ一覧APIを利用して取得してください。&lt;br&gt; &lt;a href&#x3D;\&quot;https://support.freee.co.jp/hc/ja/articles/360020679611-%E3%82%BB%E3%82%B0%E3%83%A1%E3%83%B3%E3%83%88-%E5%88%86%E6%9E%90%E7%94%A8%E3%82%BF%E3%82%B0-%E3%81%AE%E8%A8%AD%E5%AE%9A\&quot; target&#x3D;\&quot;_blank\&quot;&gt;セグメント（分析用タグ）の設定&lt;/a&gt;&lt;br&gt;  | [optional] 
**Segment3TagId** | Pointer to **int32** | セグメント３ID(法人向け エンタープライズプラン)&lt;br&gt; セグメントタグ一覧APIを利用して取得してください。&lt;br&gt; &lt;a href&#x3D;\&quot;https://support.freee.co.jp/hc/ja/articles/360020679611-%E3%82%BB%E3%82%B0%E3%83%A1%E3%83%B3%E3%83%88-%E5%88%86%E6%9E%90%E7%94%A8%E3%82%BF%E3%82%B0-%E3%81%AE%E8%A8%AD%E5%AE%9A\&quot; target&#x3D;\&quot;_blank\&quot;&gt;セグメント（分析用タグ）の設定&lt;/a&gt;&lt;br&gt;  | [optional] 

## Methods

### NewExpenseApplicationCreateParams

`func NewExpenseApplicationCreateParams(companyId int32, title string, issueDate string, expenseApplicationLines []ExpenseApplicationCreateParamsExpenseApplicationLines, ) *ExpenseApplicationCreateParams`

NewExpenseApplicationCreateParams instantiates a new ExpenseApplicationCreateParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExpenseApplicationCreateParamsWithDefaults

`func NewExpenseApplicationCreateParamsWithDefaults() *ExpenseApplicationCreateParams`

NewExpenseApplicationCreateParamsWithDefaults instantiates a new ExpenseApplicationCreateParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyId

`func (o *ExpenseApplicationCreateParams) GetCompanyId() int32`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *ExpenseApplicationCreateParams) GetCompanyIdOk() (*int32, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *ExpenseApplicationCreateParams) SetCompanyId(v int32)`

SetCompanyId sets CompanyId field to given value.


### GetTitle

`func (o *ExpenseApplicationCreateParams) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ExpenseApplicationCreateParams) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ExpenseApplicationCreateParams) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetIssueDate

`func (o *ExpenseApplicationCreateParams) GetIssueDate() string`

GetIssueDate returns the IssueDate field if non-nil, zero value otherwise.

### GetIssueDateOk

`func (o *ExpenseApplicationCreateParams) GetIssueDateOk() (*string, bool)`

GetIssueDateOk returns a tuple with the IssueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueDate

`func (o *ExpenseApplicationCreateParams) SetIssueDate(v string)`

SetIssueDate sets IssueDate field to given value.


### GetDescription

`func (o *ExpenseApplicationCreateParams) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ExpenseApplicationCreateParams) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ExpenseApplicationCreateParams) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ExpenseApplicationCreateParams) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEditableOnWeb

`func (o *ExpenseApplicationCreateParams) GetEditableOnWeb() bool`

GetEditableOnWeb returns the EditableOnWeb field if non-nil, zero value otherwise.

### GetEditableOnWebOk

`func (o *ExpenseApplicationCreateParams) GetEditableOnWebOk() (*bool, bool)`

GetEditableOnWebOk returns a tuple with the EditableOnWeb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditableOnWeb

`func (o *ExpenseApplicationCreateParams) SetEditableOnWeb(v bool)`

SetEditableOnWeb sets EditableOnWeb field to given value.

### HasEditableOnWeb

`func (o *ExpenseApplicationCreateParams) HasEditableOnWeb() bool`

HasEditableOnWeb returns a boolean if a field has been set.

### GetSectionId

`func (o *ExpenseApplicationCreateParams) GetSectionId() int32`

GetSectionId returns the SectionId field if non-nil, zero value otherwise.

### GetSectionIdOk

`func (o *ExpenseApplicationCreateParams) GetSectionIdOk() (*int32, bool)`

GetSectionIdOk returns a tuple with the SectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionId

`func (o *ExpenseApplicationCreateParams) SetSectionId(v int32)`

SetSectionId sets SectionId field to given value.

### HasSectionId

`func (o *ExpenseApplicationCreateParams) HasSectionId() bool`

HasSectionId returns a boolean if a field has been set.

### GetTagIds

`func (o *ExpenseApplicationCreateParams) GetTagIds() []int32`

GetTagIds returns the TagIds field if non-nil, zero value otherwise.

### GetTagIdsOk

`func (o *ExpenseApplicationCreateParams) GetTagIdsOk() (*[]int32, bool)`

GetTagIdsOk returns a tuple with the TagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagIds

`func (o *ExpenseApplicationCreateParams) SetTagIds(v []int32)`

SetTagIds sets TagIds field to given value.

### HasTagIds

`func (o *ExpenseApplicationCreateParams) HasTagIds() bool`

HasTagIds returns a boolean if a field has been set.

### GetExpenseApplicationLines

`func (o *ExpenseApplicationCreateParams) GetExpenseApplicationLines() []ExpenseApplicationCreateParamsExpenseApplicationLines`

GetExpenseApplicationLines returns the ExpenseApplicationLines field if non-nil, zero value otherwise.

### GetExpenseApplicationLinesOk

`func (o *ExpenseApplicationCreateParams) GetExpenseApplicationLinesOk() (*[]ExpenseApplicationCreateParamsExpenseApplicationLines, bool)`

GetExpenseApplicationLinesOk returns a tuple with the ExpenseApplicationLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpenseApplicationLines

`func (o *ExpenseApplicationCreateParams) SetExpenseApplicationLines(v []ExpenseApplicationCreateParamsExpenseApplicationLines)`

SetExpenseApplicationLines sets ExpenseApplicationLines field to given value.


### GetApprovalFlowRouteId

`func (o *ExpenseApplicationCreateParams) GetApprovalFlowRouteId() int32`

GetApprovalFlowRouteId returns the ApprovalFlowRouteId field if non-nil, zero value otherwise.

### GetApprovalFlowRouteIdOk

`func (o *ExpenseApplicationCreateParams) GetApprovalFlowRouteIdOk() (*int32, bool)`

GetApprovalFlowRouteIdOk returns a tuple with the ApprovalFlowRouteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalFlowRouteId

`func (o *ExpenseApplicationCreateParams) SetApprovalFlowRouteId(v int32)`

SetApprovalFlowRouteId sets ApprovalFlowRouteId field to given value.

### HasApprovalFlowRouteId

`func (o *ExpenseApplicationCreateParams) HasApprovalFlowRouteId() bool`

HasApprovalFlowRouteId returns a boolean if a field has been set.

### GetApproverId

`func (o *ExpenseApplicationCreateParams) GetApproverId() int32`

GetApproverId returns the ApproverId field if non-nil, zero value otherwise.

### GetApproverIdOk

`func (o *ExpenseApplicationCreateParams) GetApproverIdOk() (*int32, bool)`

GetApproverIdOk returns a tuple with the ApproverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproverId

`func (o *ExpenseApplicationCreateParams) SetApproverId(v int32)`

SetApproverId sets ApproverId field to given value.

### HasApproverId

`func (o *ExpenseApplicationCreateParams) HasApproverId() bool`

HasApproverId returns a boolean if a field has been set.

### GetDraft

`func (o *ExpenseApplicationCreateParams) GetDraft() bool`

GetDraft returns the Draft field if non-nil, zero value otherwise.

### GetDraftOk

`func (o *ExpenseApplicationCreateParams) GetDraftOk() (*bool, bool)`

GetDraftOk returns a tuple with the Draft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraft

`func (o *ExpenseApplicationCreateParams) SetDraft(v bool)`

SetDraft sets Draft field to given value.

### HasDraft

`func (o *ExpenseApplicationCreateParams) HasDraft() bool`

HasDraft returns a boolean if a field has been set.

### GetParentId

`func (o *ExpenseApplicationCreateParams) GetParentId() int32`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *ExpenseApplicationCreateParams) GetParentIdOk() (*int32, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *ExpenseApplicationCreateParams) SetParentId(v int32)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *ExpenseApplicationCreateParams) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetSegment1TagId

`func (o *ExpenseApplicationCreateParams) GetSegment1TagId() int32`

GetSegment1TagId returns the Segment1TagId field if non-nil, zero value otherwise.

### GetSegment1TagIdOk

`func (o *ExpenseApplicationCreateParams) GetSegment1TagIdOk() (*int32, bool)`

GetSegment1TagIdOk returns a tuple with the Segment1TagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment1TagId

`func (o *ExpenseApplicationCreateParams) SetSegment1TagId(v int32)`

SetSegment1TagId sets Segment1TagId field to given value.

### HasSegment1TagId

`func (o *ExpenseApplicationCreateParams) HasSegment1TagId() bool`

HasSegment1TagId returns a boolean if a field has been set.

### GetSegment2TagId

`func (o *ExpenseApplicationCreateParams) GetSegment2TagId() int32`

GetSegment2TagId returns the Segment2TagId field if non-nil, zero value otherwise.

### GetSegment2TagIdOk

`func (o *ExpenseApplicationCreateParams) GetSegment2TagIdOk() (*int32, bool)`

GetSegment2TagIdOk returns a tuple with the Segment2TagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment2TagId

`func (o *ExpenseApplicationCreateParams) SetSegment2TagId(v int32)`

SetSegment2TagId sets Segment2TagId field to given value.

### HasSegment2TagId

`func (o *ExpenseApplicationCreateParams) HasSegment2TagId() bool`

HasSegment2TagId returns a boolean if a field has been set.

### GetSegment3TagId

`func (o *ExpenseApplicationCreateParams) GetSegment3TagId() int32`

GetSegment3TagId returns the Segment3TagId field if non-nil, zero value otherwise.

### GetSegment3TagIdOk

`func (o *ExpenseApplicationCreateParams) GetSegment3TagIdOk() (*int32, bool)`

GetSegment3TagIdOk returns a tuple with the Segment3TagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment3TagId

`func (o *ExpenseApplicationCreateParams) SetSegment3TagId(v int32)`

SetSegment3TagId sets Segment3TagId field to given value.

### HasSegment3TagId

`func (o *ExpenseApplicationCreateParams) HasSegment3TagId() bool`

HasSegment3TagId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


