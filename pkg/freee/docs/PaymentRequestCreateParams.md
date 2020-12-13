# PaymentRequestCreateParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyId** | **int32** | 事業所ID | 
**Title** | **string** | 申請タイトル | 
**ApplicationDate** | **string** | 申請日 (yyyy-mm-dd) | 
**Description** | Pointer to **string** | 備考 | [optional] 
**PaymentRequestLines** | [**[]PaymentRequestCreateParamsPaymentRequestLines**](PaymentRequestCreateParamsPaymentRequestLines.md) | 支払依頼の項目行一覧（配列） | 
**ApproverId** | Pointer to **int32** | 承認者のユーザーID&lt;br&gt; 「承認者を指定」の経路を申請経路として使用する場合に指定してください。&lt;br&gt; 指定する承認者のユーザーIDは、申請経路APIを利用して取得してください。  | [optional] 
**ApprovalFlowRouteId** | **int32** | 申請経路ID&lt;br&gt; &lt;ul&gt;     &lt;li&gt;支払依頼のステータスを申請中として作成する場合は、必ず指定してください。&lt;/li&gt;     &lt;li&gt;指定する申請経路IDは、申請経路APIを利用して取得してください。&lt;/li&gt;     &lt;li&gt;         未指定の場合は、基本経路を設定している事業所では基本経路が、基本経路を設定していない事業所では利用可能な申請経路の中から最初の申請経路が自動的に使用されます。         &lt;ul&gt;            &lt;li&gt;意図しない申請経路を持った支払依頼の作成を防ぐために、使用する申請経路IDを指定することを推奨します。&lt;/li&gt;         &lt;/ul&gt;     &lt;/li&gt; &lt;/ul&gt;  | 
**ParentId** | Pointer to **int32** | 親申請ID(法人向け エンタープライズプラン、プロフェッショナルプラン)&lt;br&gt; &lt;ul&gt;   &lt;li&gt;承認済みの既存各種申請IDのみ指定可能です。&lt;/li&gt;   &lt;li&gt;各種申請一覧APIを利用して取得してください。&lt;/li&gt; &lt;/ul&gt;  | [optional] 
**Draft** | **bool** | 支払依頼のステータス&lt;br&gt; falseを指定した時は申請中（in_progress）で支払依頼を作成します。&lt;br&gt; trueを指定した時は下書き（draft）で支払依頼を作成します。&lt;br&gt; 未指定の時は下書きとみなして支払依頼を作成します。  | 
**DocumentCode** | Pointer to **string** | 請求書番号 | [optional] 
**ReceiptIds** | Pointer to **[]int32** | 証憑ファイルID（配列） | [optional] 
**IssueDate** | **string** | 発生日 (yyyy-mm-dd) | 
**PaymentDate** | Pointer to **NullableString** | 支払期限 (yyyy-mm-dd) | [optional] 
**PaymentMethod** | **string** | 支払方法(none: 指定なし, domestic_bank_transfer: 国内振込, abroad_bank_transfer: 国外振込, account_transfer: 口座振替, credit_card: クレジットカード) | 
**PartnerId** | Pointer to **NullableInt32** | 取引先ID | [optional] 
**PartnerCode** | Pointer to **NullableString** | 取引先コード | [optional] 
**BankCode** | Pointer to **string** | 銀行番号 | [optional] 
**BankName** | Pointer to **string** | 銀行名 | [optional] 
**BankNameKana** | Pointer to **string** | 銀行名（カナ）&lt;br&gt; 支払先指定時には無効  | [optional] 
**BranchCode** | Pointer to **string** | 支店番号&lt;br&gt; 支払先指定時には無効  | [optional] 
**BranchName** | Pointer to **string** | 支店名&lt;br&gt; 支払先指定時には無効  | [optional] 
**BranchKana** | Pointer to **string** | 支店名（カナ）&lt;br&gt; 支払先指定時には無効  | [optional] 
**AccountName** | Pointer to **string** | 受取人名（カナ）&lt;br&gt; 支払先指定時には無効  | [optional] 
**AccountNumber** | Pointer to **string** | 口座番号&lt;br&gt; 支払先指定時には無効  | [optional] 
**AccountType** | Pointer to **string** | &#39;口座種別(ordinary:普通、checking:当座、earmarked:納税準備預金、savings:貯蓄、other:その他)&#39;&lt;br&gt; 支払先指定時には無効&lt;br&gt; 支払先未指定時には必須  | [optional] 

## Methods

### NewPaymentRequestCreateParams

`func NewPaymentRequestCreateParams(companyId int32, title string, applicationDate string, paymentRequestLines []PaymentRequestCreateParamsPaymentRequestLines, approvalFlowRouteId int32, draft bool, issueDate string, paymentMethod string, ) *PaymentRequestCreateParams`

NewPaymentRequestCreateParams instantiates a new PaymentRequestCreateParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentRequestCreateParamsWithDefaults

`func NewPaymentRequestCreateParamsWithDefaults() *PaymentRequestCreateParams`

NewPaymentRequestCreateParamsWithDefaults instantiates a new PaymentRequestCreateParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyId

`func (o *PaymentRequestCreateParams) GetCompanyId() int32`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *PaymentRequestCreateParams) GetCompanyIdOk() (*int32, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *PaymentRequestCreateParams) SetCompanyId(v int32)`

SetCompanyId sets CompanyId field to given value.


### GetTitle

`func (o *PaymentRequestCreateParams) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PaymentRequestCreateParams) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PaymentRequestCreateParams) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetApplicationDate

`func (o *PaymentRequestCreateParams) GetApplicationDate() string`

GetApplicationDate returns the ApplicationDate field if non-nil, zero value otherwise.

### GetApplicationDateOk

`func (o *PaymentRequestCreateParams) GetApplicationDateOk() (*string, bool)`

GetApplicationDateOk returns a tuple with the ApplicationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationDate

`func (o *PaymentRequestCreateParams) SetApplicationDate(v string)`

SetApplicationDate sets ApplicationDate field to given value.


### GetDescription

`func (o *PaymentRequestCreateParams) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PaymentRequestCreateParams) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PaymentRequestCreateParams) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PaymentRequestCreateParams) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPaymentRequestLines

`func (o *PaymentRequestCreateParams) GetPaymentRequestLines() []PaymentRequestCreateParamsPaymentRequestLines`

GetPaymentRequestLines returns the PaymentRequestLines field if non-nil, zero value otherwise.

### GetPaymentRequestLinesOk

`func (o *PaymentRequestCreateParams) GetPaymentRequestLinesOk() (*[]PaymentRequestCreateParamsPaymentRequestLines, bool)`

GetPaymentRequestLinesOk returns a tuple with the PaymentRequestLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentRequestLines

`func (o *PaymentRequestCreateParams) SetPaymentRequestLines(v []PaymentRequestCreateParamsPaymentRequestLines)`

SetPaymentRequestLines sets PaymentRequestLines field to given value.


### GetApproverId

`func (o *PaymentRequestCreateParams) GetApproverId() int32`

GetApproverId returns the ApproverId field if non-nil, zero value otherwise.

### GetApproverIdOk

`func (o *PaymentRequestCreateParams) GetApproverIdOk() (*int32, bool)`

GetApproverIdOk returns a tuple with the ApproverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproverId

`func (o *PaymentRequestCreateParams) SetApproverId(v int32)`

SetApproverId sets ApproverId field to given value.

### HasApproverId

`func (o *PaymentRequestCreateParams) HasApproverId() bool`

HasApproverId returns a boolean if a field has been set.

### GetApprovalFlowRouteId

`func (o *PaymentRequestCreateParams) GetApprovalFlowRouteId() int32`

GetApprovalFlowRouteId returns the ApprovalFlowRouteId field if non-nil, zero value otherwise.

### GetApprovalFlowRouteIdOk

`func (o *PaymentRequestCreateParams) GetApprovalFlowRouteIdOk() (*int32, bool)`

GetApprovalFlowRouteIdOk returns a tuple with the ApprovalFlowRouteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalFlowRouteId

`func (o *PaymentRequestCreateParams) SetApprovalFlowRouteId(v int32)`

SetApprovalFlowRouteId sets ApprovalFlowRouteId field to given value.


### GetParentId

`func (o *PaymentRequestCreateParams) GetParentId() int32`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *PaymentRequestCreateParams) GetParentIdOk() (*int32, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *PaymentRequestCreateParams) SetParentId(v int32)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *PaymentRequestCreateParams) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetDraft

`func (o *PaymentRequestCreateParams) GetDraft() bool`

GetDraft returns the Draft field if non-nil, zero value otherwise.

### GetDraftOk

`func (o *PaymentRequestCreateParams) GetDraftOk() (*bool, bool)`

GetDraftOk returns a tuple with the Draft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraft

`func (o *PaymentRequestCreateParams) SetDraft(v bool)`

SetDraft sets Draft field to given value.


### GetDocumentCode

`func (o *PaymentRequestCreateParams) GetDocumentCode() string`

GetDocumentCode returns the DocumentCode field if non-nil, zero value otherwise.

### GetDocumentCodeOk

`func (o *PaymentRequestCreateParams) GetDocumentCodeOk() (*string, bool)`

GetDocumentCodeOk returns a tuple with the DocumentCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentCode

`func (o *PaymentRequestCreateParams) SetDocumentCode(v string)`

SetDocumentCode sets DocumentCode field to given value.

### HasDocumentCode

`func (o *PaymentRequestCreateParams) HasDocumentCode() bool`

HasDocumentCode returns a boolean if a field has been set.

### GetReceiptIds

`func (o *PaymentRequestCreateParams) GetReceiptIds() []int32`

GetReceiptIds returns the ReceiptIds field if non-nil, zero value otherwise.

### GetReceiptIdsOk

`func (o *PaymentRequestCreateParams) GetReceiptIdsOk() (*[]int32, bool)`

GetReceiptIdsOk returns a tuple with the ReceiptIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiptIds

`func (o *PaymentRequestCreateParams) SetReceiptIds(v []int32)`

SetReceiptIds sets ReceiptIds field to given value.

### HasReceiptIds

`func (o *PaymentRequestCreateParams) HasReceiptIds() bool`

HasReceiptIds returns a boolean if a field has been set.

### GetIssueDate

`func (o *PaymentRequestCreateParams) GetIssueDate() string`

GetIssueDate returns the IssueDate field if non-nil, zero value otherwise.

### GetIssueDateOk

`func (o *PaymentRequestCreateParams) GetIssueDateOk() (*string, bool)`

GetIssueDateOk returns a tuple with the IssueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueDate

`func (o *PaymentRequestCreateParams) SetIssueDate(v string)`

SetIssueDate sets IssueDate field to given value.


### GetPaymentDate

`func (o *PaymentRequestCreateParams) GetPaymentDate() string`

GetPaymentDate returns the PaymentDate field if non-nil, zero value otherwise.

### GetPaymentDateOk

`func (o *PaymentRequestCreateParams) GetPaymentDateOk() (*string, bool)`

GetPaymentDateOk returns a tuple with the PaymentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentDate

`func (o *PaymentRequestCreateParams) SetPaymentDate(v string)`

SetPaymentDate sets PaymentDate field to given value.

### HasPaymentDate

`func (o *PaymentRequestCreateParams) HasPaymentDate() bool`

HasPaymentDate returns a boolean if a field has been set.

### SetPaymentDateNil

`func (o *PaymentRequestCreateParams) SetPaymentDateNil(b bool)`

 SetPaymentDateNil sets the value for PaymentDate to be an explicit nil

### UnsetPaymentDate
`func (o *PaymentRequestCreateParams) UnsetPaymentDate()`

UnsetPaymentDate ensures that no value is present for PaymentDate, not even an explicit nil
### GetPaymentMethod

`func (o *PaymentRequestCreateParams) GetPaymentMethod() string`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *PaymentRequestCreateParams) GetPaymentMethodOk() (*string, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *PaymentRequestCreateParams) SetPaymentMethod(v string)`

SetPaymentMethod sets PaymentMethod field to given value.


### GetPartnerId

`func (o *PaymentRequestCreateParams) GetPartnerId() int32`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *PaymentRequestCreateParams) GetPartnerIdOk() (*int32, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *PaymentRequestCreateParams) SetPartnerId(v int32)`

SetPartnerId sets PartnerId field to given value.

### HasPartnerId

`func (o *PaymentRequestCreateParams) HasPartnerId() bool`

HasPartnerId returns a boolean if a field has been set.

### SetPartnerIdNil

`func (o *PaymentRequestCreateParams) SetPartnerIdNil(b bool)`

 SetPartnerIdNil sets the value for PartnerId to be an explicit nil

### UnsetPartnerId
`func (o *PaymentRequestCreateParams) UnsetPartnerId()`

UnsetPartnerId ensures that no value is present for PartnerId, not even an explicit nil
### GetPartnerCode

`func (o *PaymentRequestCreateParams) GetPartnerCode() string`

GetPartnerCode returns the PartnerCode field if non-nil, zero value otherwise.

### GetPartnerCodeOk

`func (o *PaymentRequestCreateParams) GetPartnerCodeOk() (*string, bool)`

GetPartnerCodeOk returns a tuple with the PartnerCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerCode

`func (o *PaymentRequestCreateParams) SetPartnerCode(v string)`

SetPartnerCode sets PartnerCode field to given value.

### HasPartnerCode

`func (o *PaymentRequestCreateParams) HasPartnerCode() bool`

HasPartnerCode returns a boolean if a field has been set.

### SetPartnerCodeNil

`func (o *PaymentRequestCreateParams) SetPartnerCodeNil(b bool)`

 SetPartnerCodeNil sets the value for PartnerCode to be an explicit nil

### UnsetPartnerCode
`func (o *PaymentRequestCreateParams) UnsetPartnerCode()`

UnsetPartnerCode ensures that no value is present for PartnerCode, not even an explicit nil
### GetBankCode

`func (o *PaymentRequestCreateParams) GetBankCode() string`

GetBankCode returns the BankCode field if non-nil, zero value otherwise.

### GetBankCodeOk

`func (o *PaymentRequestCreateParams) GetBankCodeOk() (*string, bool)`

GetBankCodeOk returns a tuple with the BankCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankCode

`func (o *PaymentRequestCreateParams) SetBankCode(v string)`

SetBankCode sets BankCode field to given value.

### HasBankCode

`func (o *PaymentRequestCreateParams) HasBankCode() bool`

HasBankCode returns a boolean if a field has been set.

### GetBankName

`func (o *PaymentRequestCreateParams) GetBankName() string`

GetBankName returns the BankName field if non-nil, zero value otherwise.

### GetBankNameOk

`func (o *PaymentRequestCreateParams) GetBankNameOk() (*string, bool)`

GetBankNameOk returns a tuple with the BankName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankName

`func (o *PaymentRequestCreateParams) SetBankName(v string)`

SetBankName sets BankName field to given value.

### HasBankName

`func (o *PaymentRequestCreateParams) HasBankName() bool`

HasBankName returns a boolean if a field has been set.

### GetBankNameKana

`func (o *PaymentRequestCreateParams) GetBankNameKana() string`

GetBankNameKana returns the BankNameKana field if non-nil, zero value otherwise.

### GetBankNameKanaOk

`func (o *PaymentRequestCreateParams) GetBankNameKanaOk() (*string, bool)`

GetBankNameKanaOk returns a tuple with the BankNameKana field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankNameKana

`func (o *PaymentRequestCreateParams) SetBankNameKana(v string)`

SetBankNameKana sets BankNameKana field to given value.

### HasBankNameKana

`func (o *PaymentRequestCreateParams) HasBankNameKana() bool`

HasBankNameKana returns a boolean if a field has been set.

### GetBranchCode

`func (o *PaymentRequestCreateParams) GetBranchCode() string`

GetBranchCode returns the BranchCode field if non-nil, zero value otherwise.

### GetBranchCodeOk

`func (o *PaymentRequestCreateParams) GetBranchCodeOk() (*string, bool)`

GetBranchCodeOk returns a tuple with the BranchCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchCode

`func (o *PaymentRequestCreateParams) SetBranchCode(v string)`

SetBranchCode sets BranchCode field to given value.

### HasBranchCode

`func (o *PaymentRequestCreateParams) HasBranchCode() bool`

HasBranchCode returns a boolean if a field has been set.

### GetBranchName

`func (o *PaymentRequestCreateParams) GetBranchName() string`

GetBranchName returns the BranchName field if non-nil, zero value otherwise.

### GetBranchNameOk

`func (o *PaymentRequestCreateParams) GetBranchNameOk() (*string, bool)`

GetBranchNameOk returns a tuple with the BranchName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchName

`func (o *PaymentRequestCreateParams) SetBranchName(v string)`

SetBranchName sets BranchName field to given value.

### HasBranchName

`func (o *PaymentRequestCreateParams) HasBranchName() bool`

HasBranchName returns a boolean if a field has been set.

### GetBranchKana

`func (o *PaymentRequestCreateParams) GetBranchKana() string`

GetBranchKana returns the BranchKana field if non-nil, zero value otherwise.

### GetBranchKanaOk

`func (o *PaymentRequestCreateParams) GetBranchKanaOk() (*string, bool)`

GetBranchKanaOk returns a tuple with the BranchKana field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchKana

`func (o *PaymentRequestCreateParams) SetBranchKana(v string)`

SetBranchKana sets BranchKana field to given value.

### HasBranchKana

`func (o *PaymentRequestCreateParams) HasBranchKana() bool`

HasBranchKana returns a boolean if a field has been set.

### GetAccountName

`func (o *PaymentRequestCreateParams) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *PaymentRequestCreateParams) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *PaymentRequestCreateParams) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.

### HasAccountName

`func (o *PaymentRequestCreateParams) HasAccountName() bool`

HasAccountName returns a boolean if a field has been set.

### GetAccountNumber

`func (o *PaymentRequestCreateParams) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *PaymentRequestCreateParams) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *PaymentRequestCreateParams) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *PaymentRequestCreateParams) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### GetAccountType

`func (o *PaymentRequestCreateParams) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *PaymentRequestCreateParams) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *PaymentRequestCreateParams) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *PaymentRequestCreateParams) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


