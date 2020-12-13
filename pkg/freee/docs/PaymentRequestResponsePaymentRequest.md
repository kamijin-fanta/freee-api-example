# PaymentRequestResponsePaymentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | 支払依頼ID | 
**CompanyId** | **int32** | 事業所ID | 
**Title** | **string** | 申請タイトル | 
**ApplicationDate** | **string** | 申請日 (yyyy-mm-dd) | 
**Description** | **NullableString** | 備考 | 
**TotalAmount** | **int32** | 合計金額 | 
**Status** | **string** | 申請ステータス(draft:下書き, in_progress:申請中, approved:承認済, rejected:却下, feedback:差戻し) | 
**PaymentRequestLines** | [**[]PaymentRequestResponsePaymentRequestPaymentRequestLines**](PaymentRequestResponsePaymentRequestPaymentRequestLines.md) | 支払依頼の項目行一覧（配列） | 
**DealId** | **NullableInt32** | 取引ID (申請ステータス:statusがapprovedで、取引が存在する時のみdeal_idが表示されます) | 
**DealStatus** | **NullableString** | 取引ステータス (申請ステータス:statusがapprovedで、取引が存在する時のみdeal_statusが表示されます settled:精算済み, unsettled:清算待ち) | 
**ApplicantId** | **int32** | 申請者のユーザーID | 
**Approvers** | [**[]ExpenseApplicationResponseExpenseApplicationApprovers**](ExpenseApplicationResponseExpenseApplicationApprovers.md) | 承認者（配列）   承認ステップのresource_typeがunspecified (指定なし)の場合はapproversはレスポンスに含まれません。   しかし、resource_typeがunspecifiedの承認ステップにおいて誰かが承認・却下・差し戻しのいずれかのアクションを取った後は、    approversはレスポンスに含まれるようになります。    その場合approversにはアクションを行ったステップのIDとアクションを行ったユーザーのIDが含まれます。 | 
**ApplicationNumber** | **string** | 申請No. | 
**ApprovalFlowRouteId** | **int32** | 申請経路ID | 
**Comments** | [**[]ExpenseApplicationResponseExpenseApplicationComments**](ExpenseApplicationResponseExpenseApplicationComments.md) | 支払依頼のコメント一覧（配列） | 
**ApprovalFlowLogs** | [**[]ExpenseApplicationResponseExpenseApplicationApprovalFlowLogs**](ExpenseApplicationResponseExpenseApplicationApprovalFlowLogs.md) | 支払依頼の承認履歴（配列） | 
**CurrentStepId** | **NullableInt32** | 現在承認ステップID | 
**CurrentRound** | **int32** | 現在のround。差し戻し等により申請がstepの最初からやり直しになるとroundの値が増えます。 | 
**DocumentCode** | **string** | 請求書番号 | 
**ReceiptId** | Pointer to **NullableInt32** | 証憑ID | [optional] 
**IssueDate** | **string** | 発生日 (yyyy-mm-dd) | 
**PaymentDate** | **NullableString** | 支払期限 (yyyy-mm-dd) | 
**PaymentMethod** | **string** | 支払方法(none: 指定なし, domestic_bank_transfer: 国内振込, abroad_bank_transfer: 国外振込, account_transfer: 口座振替, credit_card: クレジットカード) | 
**PartnerId** | **NullableInt32** | 取引先ID | 
**PartnerCode** | Pointer to **NullableString** | 取引先コード | [optional] 
**PartnerName** | **NullableString** | 取引先名 | 
**BankName** | **string** | 銀行名 | 
**BankNameKana** | **string** | 銀行名（カナ） | 
**BankCode** | **string** | 銀行コード | 
**BranchName** | **string** | 支店名 | 
**BranchKana** | **string** | 支店名（カナ） | 
**BranchCode** | **string** | 支店番号 | 
**AccountType** | **string** | 口座種別(ordinary:普通、checking:当座、earmarked:納税準備預金、savings:貯蓄、other:その他) | 
**AccountNumber** | **string** | 口座番号 | 
**AccountName** | **string** | 受取人名（カナ） | 

## Methods

### NewPaymentRequestResponsePaymentRequest

`func NewPaymentRequestResponsePaymentRequest(id int32, companyId int32, title string, applicationDate string, description NullableString, totalAmount int32, status string, paymentRequestLines []PaymentRequestResponsePaymentRequestPaymentRequestLines, dealId NullableInt32, dealStatus NullableString, applicantId int32, approvers []ExpenseApplicationResponseExpenseApplicationApprovers, applicationNumber string, approvalFlowRouteId int32, comments []ExpenseApplicationResponseExpenseApplicationComments, approvalFlowLogs []ExpenseApplicationResponseExpenseApplicationApprovalFlowLogs, currentStepId NullableInt32, currentRound int32, documentCode string, issueDate string, paymentDate NullableString, paymentMethod string, partnerId NullableInt32, partnerName NullableString, bankName string, bankNameKana string, bankCode string, branchName string, branchKana string, branchCode string, accountType string, accountNumber string, accountName string, ) *PaymentRequestResponsePaymentRequest`

NewPaymentRequestResponsePaymentRequest instantiates a new PaymentRequestResponsePaymentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentRequestResponsePaymentRequestWithDefaults

`func NewPaymentRequestResponsePaymentRequestWithDefaults() *PaymentRequestResponsePaymentRequest`

NewPaymentRequestResponsePaymentRequestWithDefaults instantiates a new PaymentRequestResponsePaymentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PaymentRequestResponsePaymentRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaymentRequestResponsePaymentRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaymentRequestResponsePaymentRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetCompanyId

`func (o *PaymentRequestResponsePaymentRequest) GetCompanyId() int32`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *PaymentRequestResponsePaymentRequest) GetCompanyIdOk() (*int32, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *PaymentRequestResponsePaymentRequest) SetCompanyId(v int32)`

SetCompanyId sets CompanyId field to given value.


### GetTitle

`func (o *PaymentRequestResponsePaymentRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PaymentRequestResponsePaymentRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PaymentRequestResponsePaymentRequest) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetApplicationDate

`func (o *PaymentRequestResponsePaymentRequest) GetApplicationDate() string`

GetApplicationDate returns the ApplicationDate field if non-nil, zero value otherwise.

### GetApplicationDateOk

`func (o *PaymentRequestResponsePaymentRequest) GetApplicationDateOk() (*string, bool)`

GetApplicationDateOk returns a tuple with the ApplicationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationDate

`func (o *PaymentRequestResponsePaymentRequest) SetApplicationDate(v string)`

SetApplicationDate sets ApplicationDate field to given value.


### GetDescription

`func (o *PaymentRequestResponsePaymentRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PaymentRequestResponsePaymentRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PaymentRequestResponsePaymentRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *PaymentRequestResponsePaymentRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PaymentRequestResponsePaymentRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetTotalAmount

`func (o *PaymentRequestResponsePaymentRequest) GetTotalAmount() int32`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *PaymentRequestResponsePaymentRequest) GetTotalAmountOk() (*int32, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *PaymentRequestResponsePaymentRequest) SetTotalAmount(v int32)`

SetTotalAmount sets TotalAmount field to given value.


### GetStatus

`func (o *PaymentRequestResponsePaymentRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PaymentRequestResponsePaymentRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PaymentRequestResponsePaymentRequest) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetPaymentRequestLines

`func (o *PaymentRequestResponsePaymentRequest) GetPaymentRequestLines() []PaymentRequestResponsePaymentRequestPaymentRequestLines`

GetPaymentRequestLines returns the PaymentRequestLines field if non-nil, zero value otherwise.

### GetPaymentRequestLinesOk

`func (o *PaymentRequestResponsePaymentRequest) GetPaymentRequestLinesOk() (*[]PaymentRequestResponsePaymentRequestPaymentRequestLines, bool)`

GetPaymentRequestLinesOk returns a tuple with the PaymentRequestLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentRequestLines

`func (o *PaymentRequestResponsePaymentRequest) SetPaymentRequestLines(v []PaymentRequestResponsePaymentRequestPaymentRequestLines)`

SetPaymentRequestLines sets PaymentRequestLines field to given value.


### GetDealId

`func (o *PaymentRequestResponsePaymentRequest) GetDealId() int32`

GetDealId returns the DealId field if non-nil, zero value otherwise.

### GetDealIdOk

`func (o *PaymentRequestResponsePaymentRequest) GetDealIdOk() (*int32, bool)`

GetDealIdOk returns a tuple with the DealId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDealId

`func (o *PaymentRequestResponsePaymentRequest) SetDealId(v int32)`

SetDealId sets DealId field to given value.


### SetDealIdNil

`func (o *PaymentRequestResponsePaymentRequest) SetDealIdNil(b bool)`

 SetDealIdNil sets the value for DealId to be an explicit nil

### UnsetDealId
`func (o *PaymentRequestResponsePaymentRequest) UnsetDealId()`

UnsetDealId ensures that no value is present for DealId, not even an explicit nil
### GetDealStatus

`func (o *PaymentRequestResponsePaymentRequest) GetDealStatus() string`

GetDealStatus returns the DealStatus field if non-nil, zero value otherwise.

### GetDealStatusOk

`func (o *PaymentRequestResponsePaymentRequest) GetDealStatusOk() (*string, bool)`

GetDealStatusOk returns a tuple with the DealStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDealStatus

`func (o *PaymentRequestResponsePaymentRequest) SetDealStatus(v string)`

SetDealStatus sets DealStatus field to given value.


### SetDealStatusNil

`func (o *PaymentRequestResponsePaymentRequest) SetDealStatusNil(b bool)`

 SetDealStatusNil sets the value for DealStatus to be an explicit nil

### UnsetDealStatus
`func (o *PaymentRequestResponsePaymentRequest) UnsetDealStatus()`

UnsetDealStatus ensures that no value is present for DealStatus, not even an explicit nil
### GetApplicantId

`func (o *PaymentRequestResponsePaymentRequest) GetApplicantId() int32`

GetApplicantId returns the ApplicantId field if non-nil, zero value otherwise.

### GetApplicantIdOk

`func (o *PaymentRequestResponsePaymentRequest) GetApplicantIdOk() (*int32, bool)`

GetApplicantIdOk returns a tuple with the ApplicantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicantId

`func (o *PaymentRequestResponsePaymentRequest) SetApplicantId(v int32)`

SetApplicantId sets ApplicantId field to given value.


### GetApprovers

`func (o *PaymentRequestResponsePaymentRequest) GetApprovers() []ExpenseApplicationResponseExpenseApplicationApprovers`

GetApprovers returns the Approvers field if non-nil, zero value otherwise.

### GetApproversOk

`func (o *PaymentRequestResponsePaymentRequest) GetApproversOk() (*[]ExpenseApplicationResponseExpenseApplicationApprovers, bool)`

GetApproversOk returns a tuple with the Approvers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovers

`func (o *PaymentRequestResponsePaymentRequest) SetApprovers(v []ExpenseApplicationResponseExpenseApplicationApprovers)`

SetApprovers sets Approvers field to given value.


### GetApplicationNumber

`func (o *PaymentRequestResponsePaymentRequest) GetApplicationNumber() string`

GetApplicationNumber returns the ApplicationNumber field if non-nil, zero value otherwise.

### GetApplicationNumberOk

`func (o *PaymentRequestResponsePaymentRequest) GetApplicationNumberOk() (*string, bool)`

GetApplicationNumberOk returns a tuple with the ApplicationNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationNumber

`func (o *PaymentRequestResponsePaymentRequest) SetApplicationNumber(v string)`

SetApplicationNumber sets ApplicationNumber field to given value.


### GetApprovalFlowRouteId

`func (o *PaymentRequestResponsePaymentRequest) GetApprovalFlowRouteId() int32`

GetApprovalFlowRouteId returns the ApprovalFlowRouteId field if non-nil, zero value otherwise.

### GetApprovalFlowRouteIdOk

`func (o *PaymentRequestResponsePaymentRequest) GetApprovalFlowRouteIdOk() (*int32, bool)`

GetApprovalFlowRouteIdOk returns a tuple with the ApprovalFlowRouteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalFlowRouteId

`func (o *PaymentRequestResponsePaymentRequest) SetApprovalFlowRouteId(v int32)`

SetApprovalFlowRouteId sets ApprovalFlowRouteId field to given value.


### GetComments

`func (o *PaymentRequestResponsePaymentRequest) GetComments() []ExpenseApplicationResponseExpenseApplicationComments`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *PaymentRequestResponsePaymentRequest) GetCommentsOk() (*[]ExpenseApplicationResponseExpenseApplicationComments, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *PaymentRequestResponsePaymentRequest) SetComments(v []ExpenseApplicationResponseExpenseApplicationComments)`

SetComments sets Comments field to given value.


### GetApprovalFlowLogs

`func (o *PaymentRequestResponsePaymentRequest) GetApprovalFlowLogs() []ExpenseApplicationResponseExpenseApplicationApprovalFlowLogs`

GetApprovalFlowLogs returns the ApprovalFlowLogs field if non-nil, zero value otherwise.

### GetApprovalFlowLogsOk

`func (o *PaymentRequestResponsePaymentRequest) GetApprovalFlowLogsOk() (*[]ExpenseApplicationResponseExpenseApplicationApprovalFlowLogs, bool)`

GetApprovalFlowLogsOk returns a tuple with the ApprovalFlowLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalFlowLogs

`func (o *PaymentRequestResponsePaymentRequest) SetApprovalFlowLogs(v []ExpenseApplicationResponseExpenseApplicationApprovalFlowLogs)`

SetApprovalFlowLogs sets ApprovalFlowLogs field to given value.


### GetCurrentStepId

`func (o *PaymentRequestResponsePaymentRequest) GetCurrentStepId() int32`

GetCurrentStepId returns the CurrentStepId field if non-nil, zero value otherwise.

### GetCurrentStepIdOk

`func (o *PaymentRequestResponsePaymentRequest) GetCurrentStepIdOk() (*int32, bool)`

GetCurrentStepIdOk returns a tuple with the CurrentStepId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentStepId

`func (o *PaymentRequestResponsePaymentRequest) SetCurrentStepId(v int32)`

SetCurrentStepId sets CurrentStepId field to given value.


### SetCurrentStepIdNil

`func (o *PaymentRequestResponsePaymentRequest) SetCurrentStepIdNil(b bool)`

 SetCurrentStepIdNil sets the value for CurrentStepId to be an explicit nil

### UnsetCurrentStepId
`func (o *PaymentRequestResponsePaymentRequest) UnsetCurrentStepId()`

UnsetCurrentStepId ensures that no value is present for CurrentStepId, not even an explicit nil
### GetCurrentRound

`func (o *PaymentRequestResponsePaymentRequest) GetCurrentRound() int32`

GetCurrentRound returns the CurrentRound field if non-nil, zero value otherwise.

### GetCurrentRoundOk

`func (o *PaymentRequestResponsePaymentRequest) GetCurrentRoundOk() (*int32, bool)`

GetCurrentRoundOk returns a tuple with the CurrentRound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentRound

`func (o *PaymentRequestResponsePaymentRequest) SetCurrentRound(v int32)`

SetCurrentRound sets CurrentRound field to given value.


### GetDocumentCode

`func (o *PaymentRequestResponsePaymentRequest) GetDocumentCode() string`

GetDocumentCode returns the DocumentCode field if non-nil, zero value otherwise.

### GetDocumentCodeOk

`func (o *PaymentRequestResponsePaymentRequest) GetDocumentCodeOk() (*string, bool)`

GetDocumentCodeOk returns a tuple with the DocumentCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentCode

`func (o *PaymentRequestResponsePaymentRequest) SetDocumentCode(v string)`

SetDocumentCode sets DocumentCode field to given value.


### GetReceiptId

`func (o *PaymentRequestResponsePaymentRequest) GetReceiptId() int32`

GetReceiptId returns the ReceiptId field if non-nil, zero value otherwise.

### GetReceiptIdOk

`func (o *PaymentRequestResponsePaymentRequest) GetReceiptIdOk() (*int32, bool)`

GetReceiptIdOk returns a tuple with the ReceiptId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiptId

`func (o *PaymentRequestResponsePaymentRequest) SetReceiptId(v int32)`

SetReceiptId sets ReceiptId field to given value.

### HasReceiptId

`func (o *PaymentRequestResponsePaymentRequest) HasReceiptId() bool`

HasReceiptId returns a boolean if a field has been set.

### SetReceiptIdNil

`func (o *PaymentRequestResponsePaymentRequest) SetReceiptIdNil(b bool)`

 SetReceiptIdNil sets the value for ReceiptId to be an explicit nil

### UnsetReceiptId
`func (o *PaymentRequestResponsePaymentRequest) UnsetReceiptId()`

UnsetReceiptId ensures that no value is present for ReceiptId, not even an explicit nil
### GetIssueDate

`func (o *PaymentRequestResponsePaymentRequest) GetIssueDate() string`

GetIssueDate returns the IssueDate field if non-nil, zero value otherwise.

### GetIssueDateOk

`func (o *PaymentRequestResponsePaymentRequest) GetIssueDateOk() (*string, bool)`

GetIssueDateOk returns a tuple with the IssueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueDate

`func (o *PaymentRequestResponsePaymentRequest) SetIssueDate(v string)`

SetIssueDate sets IssueDate field to given value.


### GetPaymentDate

`func (o *PaymentRequestResponsePaymentRequest) GetPaymentDate() string`

GetPaymentDate returns the PaymentDate field if non-nil, zero value otherwise.

### GetPaymentDateOk

`func (o *PaymentRequestResponsePaymentRequest) GetPaymentDateOk() (*string, bool)`

GetPaymentDateOk returns a tuple with the PaymentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentDate

`func (o *PaymentRequestResponsePaymentRequest) SetPaymentDate(v string)`

SetPaymentDate sets PaymentDate field to given value.


### SetPaymentDateNil

`func (o *PaymentRequestResponsePaymentRequest) SetPaymentDateNil(b bool)`

 SetPaymentDateNil sets the value for PaymentDate to be an explicit nil

### UnsetPaymentDate
`func (o *PaymentRequestResponsePaymentRequest) UnsetPaymentDate()`

UnsetPaymentDate ensures that no value is present for PaymentDate, not even an explicit nil
### GetPaymentMethod

`func (o *PaymentRequestResponsePaymentRequest) GetPaymentMethod() string`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *PaymentRequestResponsePaymentRequest) GetPaymentMethodOk() (*string, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *PaymentRequestResponsePaymentRequest) SetPaymentMethod(v string)`

SetPaymentMethod sets PaymentMethod field to given value.


### GetPartnerId

`func (o *PaymentRequestResponsePaymentRequest) GetPartnerId() int32`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *PaymentRequestResponsePaymentRequest) GetPartnerIdOk() (*int32, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *PaymentRequestResponsePaymentRequest) SetPartnerId(v int32)`

SetPartnerId sets PartnerId field to given value.


### SetPartnerIdNil

`func (o *PaymentRequestResponsePaymentRequest) SetPartnerIdNil(b bool)`

 SetPartnerIdNil sets the value for PartnerId to be an explicit nil

### UnsetPartnerId
`func (o *PaymentRequestResponsePaymentRequest) UnsetPartnerId()`

UnsetPartnerId ensures that no value is present for PartnerId, not even an explicit nil
### GetPartnerCode

`func (o *PaymentRequestResponsePaymentRequest) GetPartnerCode() string`

GetPartnerCode returns the PartnerCode field if non-nil, zero value otherwise.

### GetPartnerCodeOk

`func (o *PaymentRequestResponsePaymentRequest) GetPartnerCodeOk() (*string, bool)`

GetPartnerCodeOk returns a tuple with the PartnerCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerCode

`func (o *PaymentRequestResponsePaymentRequest) SetPartnerCode(v string)`

SetPartnerCode sets PartnerCode field to given value.

### HasPartnerCode

`func (o *PaymentRequestResponsePaymentRequest) HasPartnerCode() bool`

HasPartnerCode returns a boolean if a field has been set.

### SetPartnerCodeNil

`func (o *PaymentRequestResponsePaymentRequest) SetPartnerCodeNil(b bool)`

 SetPartnerCodeNil sets the value for PartnerCode to be an explicit nil

### UnsetPartnerCode
`func (o *PaymentRequestResponsePaymentRequest) UnsetPartnerCode()`

UnsetPartnerCode ensures that no value is present for PartnerCode, not even an explicit nil
### GetPartnerName

`func (o *PaymentRequestResponsePaymentRequest) GetPartnerName() string`

GetPartnerName returns the PartnerName field if non-nil, zero value otherwise.

### GetPartnerNameOk

`func (o *PaymentRequestResponsePaymentRequest) GetPartnerNameOk() (*string, bool)`

GetPartnerNameOk returns a tuple with the PartnerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerName

`func (o *PaymentRequestResponsePaymentRequest) SetPartnerName(v string)`

SetPartnerName sets PartnerName field to given value.


### SetPartnerNameNil

`func (o *PaymentRequestResponsePaymentRequest) SetPartnerNameNil(b bool)`

 SetPartnerNameNil sets the value for PartnerName to be an explicit nil

### UnsetPartnerName
`func (o *PaymentRequestResponsePaymentRequest) UnsetPartnerName()`

UnsetPartnerName ensures that no value is present for PartnerName, not even an explicit nil
### GetBankName

`func (o *PaymentRequestResponsePaymentRequest) GetBankName() string`

GetBankName returns the BankName field if non-nil, zero value otherwise.

### GetBankNameOk

`func (o *PaymentRequestResponsePaymentRequest) GetBankNameOk() (*string, bool)`

GetBankNameOk returns a tuple with the BankName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankName

`func (o *PaymentRequestResponsePaymentRequest) SetBankName(v string)`

SetBankName sets BankName field to given value.


### GetBankNameKana

`func (o *PaymentRequestResponsePaymentRequest) GetBankNameKana() string`

GetBankNameKana returns the BankNameKana field if non-nil, zero value otherwise.

### GetBankNameKanaOk

`func (o *PaymentRequestResponsePaymentRequest) GetBankNameKanaOk() (*string, bool)`

GetBankNameKanaOk returns a tuple with the BankNameKana field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankNameKana

`func (o *PaymentRequestResponsePaymentRequest) SetBankNameKana(v string)`

SetBankNameKana sets BankNameKana field to given value.


### GetBankCode

`func (o *PaymentRequestResponsePaymentRequest) GetBankCode() string`

GetBankCode returns the BankCode field if non-nil, zero value otherwise.

### GetBankCodeOk

`func (o *PaymentRequestResponsePaymentRequest) GetBankCodeOk() (*string, bool)`

GetBankCodeOk returns a tuple with the BankCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankCode

`func (o *PaymentRequestResponsePaymentRequest) SetBankCode(v string)`

SetBankCode sets BankCode field to given value.


### GetBranchName

`func (o *PaymentRequestResponsePaymentRequest) GetBranchName() string`

GetBranchName returns the BranchName field if non-nil, zero value otherwise.

### GetBranchNameOk

`func (o *PaymentRequestResponsePaymentRequest) GetBranchNameOk() (*string, bool)`

GetBranchNameOk returns a tuple with the BranchName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchName

`func (o *PaymentRequestResponsePaymentRequest) SetBranchName(v string)`

SetBranchName sets BranchName field to given value.


### GetBranchKana

`func (o *PaymentRequestResponsePaymentRequest) GetBranchKana() string`

GetBranchKana returns the BranchKana field if non-nil, zero value otherwise.

### GetBranchKanaOk

`func (o *PaymentRequestResponsePaymentRequest) GetBranchKanaOk() (*string, bool)`

GetBranchKanaOk returns a tuple with the BranchKana field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchKana

`func (o *PaymentRequestResponsePaymentRequest) SetBranchKana(v string)`

SetBranchKana sets BranchKana field to given value.


### GetBranchCode

`func (o *PaymentRequestResponsePaymentRequest) GetBranchCode() string`

GetBranchCode returns the BranchCode field if non-nil, zero value otherwise.

### GetBranchCodeOk

`func (o *PaymentRequestResponsePaymentRequest) GetBranchCodeOk() (*string, bool)`

GetBranchCodeOk returns a tuple with the BranchCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchCode

`func (o *PaymentRequestResponsePaymentRequest) SetBranchCode(v string)`

SetBranchCode sets BranchCode field to given value.


### GetAccountType

`func (o *PaymentRequestResponsePaymentRequest) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *PaymentRequestResponsePaymentRequest) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *PaymentRequestResponsePaymentRequest) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.


### GetAccountNumber

`func (o *PaymentRequestResponsePaymentRequest) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *PaymentRequestResponsePaymentRequest) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *PaymentRequestResponsePaymentRequest) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.


### GetAccountName

`func (o *PaymentRequestResponsePaymentRequest) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *PaymentRequestResponsePaymentRequest) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *PaymentRequestResponsePaymentRequest) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


