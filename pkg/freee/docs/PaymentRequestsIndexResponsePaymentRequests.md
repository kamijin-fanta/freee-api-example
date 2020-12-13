# PaymentRequestsIndexResponsePaymentRequests

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | 支払依頼ID | 
**CompanyId** | **int32** | 事業所ID | 
**Title** | **string** | 申請タイトル | 
**ApplicationDate** | **string** | 申請日 (yyyy-mm-dd) | 
**TotalAmount** | **int32** | 合計金額 | 
**Status** | **string** | 申請ステータス(draft:下書き, in_progress:申請中, approved:承認済, rejected:却下, feedback:差戻し) | 
**DealId** | **NullableInt32** | 取引ID (申請ステータス:statusがapprovedで、取引が存在する時のみdeal_idが表示されます) | 
**DealStatus** | **NullableString** | 取引ステータス (申請ステータス:statusがapprovedで、取引が存在する時のみdeal_statusが表示されます settled:精算済み, unsettled:清算待ち) | 
**ApplicantId** | **int32** | 申請者のユーザーID | 
**ApplicationNumber** | **string** | 申請No. | 
**CurrentStepId** | **NullableInt32** | 現在承認ステップID | 
**CurrentRound** | **int32** | 現在のround。差し戻し等により申請がstepの最初からやり直しになるとroundの値が増えます。 | 
**IssueDate** | **string** | 発生日 (yyyy-mm-dd) | 
**PaymentDate** | **NullableString** | 支払期限 (yyyy-mm-dd) | 
**PaymentMethod** | **string** | 支払方法(none: 指定なし, domestic_bank_transfer: 国内振込, abroad_bank_transfer: 国外振込, account_transfer: 口座振替, credit_card: クレジットカード) | 
**PartnerId** | **NullableInt32** | 取引先ID | 
**PartnerCode** | **NullableString** | 取引先コード | 
**PartnerName** | **NullableString** | 取引先名 | 

## Methods

### NewPaymentRequestsIndexResponsePaymentRequests

`func NewPaymentRequestsIndexResponsePaymentRequests(id int32, companyId int32, title string, applicationDate string, totalAmount int32, status string, dealId NullableInt32, dealStatus NullableString, applicantId int32, applicationNumber string, currentStepId NullableInt32, currentRound int32, issueDate string, paymentDate NullableString, paymentMethod string, partnerId NullableInt32, partnerCode NullableString, partnerName NullableString, ) *PaymentRequestsIndexResponsePaymentRequests`

NewPaymentRequestsIndexResponsePaymentRequests instantiates a new PaymentRequestsIndexResponsePaymentRequests object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentRequestsIndexResponsePaymentRequestsWithDefaults

`func NewPaymentRequestsIndexResponsePaymentRequestsWithDefaults() *PaymentRequestsIndexResponsePaymentRequests`

NewPaymentRequestsIndexResponsePaymentRequestsWithDefaults instantiates a new PaymentRequestsIndexResponsePaymentRequests object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PaymentRequestsIndexResponsePaymentRequests) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaymentRequestsIndexResponsePaymentRequests) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaymentRequestsIndexResponsePaymentRequests) SetId(v int32)`

SetId sets Id field to given value.


### GetCompanyId

`func (o *PaymentRequestsIndexResponsePaymentRequests) GetCompanyId() int32`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *PaymentRequestsIndexResponsePaymentRequests) GetCompanyIdOk() (*int32, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *PaymentRequestsIndexResponsePaymentRequests) SetCompanyId(v int32)`

SetCompanyId sets CompanyId field to given value.


### GetTitle

`func (o *PaymentRequestsIndexResponsePaymentRequests) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PaymentRequestsIndexResponsePaymentRequests) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PaymentRequestsIndexResponsePaymentRequests) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetApplicationDate

`func (o *PaymentRequestsIndexResponsePaymentRequests) GetApplicationDate() string`

GetApplicationDate returns the ApplicationDate field if non-nil, zero value otherwise.

### GetApplicationDateOk

`func (o *PaymentRequestsIndexResponsePaymentRequests) GetApplicationDateOk() (*string, bool)`

GetApplicationDateOk returns a tuple with the ApplicationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationDate

`func (o *PaymentRequestsIndexResponsePaymentRequests) SetApplicationDate(v string)`

SetApplicationDate sets ApplicationDate field to given value.


### GetTotalAmount

`func (o *PaymentRequestsIndexResponsePaymentRequests) GetTotalAmount() int32`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *PaymentRequestsIndexResponsePaymentRequests) GetTotalAmountOk() (*int32, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *PaymentRequestsIndexResponsePaymentRequests) SetTotalAmount(v int32)`

SetTotalAmount sets TotalAmount field to given value.


### GetStatus

`func (o *PaymentRequestsIndexResponsePaymentRequests) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PaymentRequestsIndexResponsePaymentRequests) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PaymentRequestsIndexResponsePaymentRequests) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetDealId

`func (o *PaymentRequestsIndexResponsePaymentRequests) GetDealId() int32`

GetDealId returns the DealId field if non-nil, zero value otherwise.

### GetDealIdOk

`func (o *PaymentRequestsIndexResponsePaymentRequests) GetDealIdOk() (*int32, bool)`

GetDealIdOk returns a tuple with the DealId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDealId

`func (o *PaymentRequestsIndexResponsePaymentRequests) SetDealId(v int32)`

SetDealId sets DealId field to given value.


### SetDealIdNil

`func (o *PaymentRequestsIndexResponsePaymentRequests) SetDealIdNil(b bool)`

 SetDealIdNil sets the value for DealId to be an explicit nil

### UnsetDealId
`func (o *PaymentRequestsIndexResponsePaymentRequests) UnsetDealId()`

UnsetDealId ensures that no value is present for DealId, not even an explicit nil
### GetDealStatus

`func (o *PaymentRequestsIndexResponsePaymentRequests) GetDealStatus() string`

GetDealStatus returns the DealStatus field if non-nil, zero value otherwise.

### GetDealStatusOk

`func (o *PaymentRequestsIndexResponsePaymentRequests) GetDealStatusOk() (*string, bool)`

GetDealStatusOk returns a tuple with the DealStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDealStatus

`func (o *PaymentRequestsIndexResponsePaymentRequests) SetDealStatus(v string)`

SetDealStatus sets DealStatus field to given value.


### SetDealStatusNil

`func (o *PaymentRequestsIndexResponsePaymentRequests) SetDealStatusNil(b bool)`

 SetDealStatusNil sets the value for DealStatus to be an explicit nil

### UnsetDealStatus
`func (o *PaymentRequestsIndexResponsePaymentRequests) UnsetDealStatus()`

UnsetDealStatus ensures that no value is present for DealStatus, not even an explicit nil
### GetApplicantId

`func (o *PaymentRequestsIndexResponsePaymentRequests) GetApplicantId() int32`

GetApplicantId returns the ApplicantId field if non-nil, zero value otherwise.

### GetApplicantIdOk

`func (o *PaymentRequestsIndexResponsePaymentRequests) GetApplicantIdOk() (*int32, bool)`

GetApplicantIdOk returns a tuple with the ApplicantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicantId

`func (o *PaymentRequestsIndexResponsePaymentRequests) SetApplicantId(v int32)`

SetApplicantId sets ApplicantId field to given value.


### GetApplicationNumber

`func (o *PaymentRequestsIndexResponsePaymentRequests) GetApplicationNumber() string`

GetApplicationNumber returns the ApplicationNumber field if non-nil, zero value otherwise.

### GetApplicationNumberOk

`func (o *PaymentRequestsIndexResponsePaymentRequests) GetApplicationNumberOk() (*string, bool)`

GetApplicationNumberOk returns a tuple with the ApplicationNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationNumber

`func (o *PaymentRequestsIndexResponsePaymentRequests) SetApplicationNumber(v string)`

SetApplicationNumber sets ApplicationNumber field to given value.


### GetCurrentStepId

`func (o *PaymentRequestsIndexResponsePaymentRequests) GetCurrentStepId() int32`

GetCurrentStepId returns the CurrentStepId field if non-nil, zero value otherwise.

### GetCurrentStepIdOk

`func (o *PaymentRequestsIndexResponsePaymentRequests) GetCurrentStepIdOk() (*int32, bool)`

GetCurrentStepIdOk returns a tuple with the CurrentStepId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentStepId

`func (o *PaymentRequestsIndexResponsePaymentRequests) SetCurrentStepId(v int32)`

SetCurrentStepId sets CurrentStepId field to given value.


### SetCurrentStepIdNil

`func (o *PaymentRequestsIndexResponsePaymentRequests) SetCurrentStepIdNil(b bool)`

 SetCurrentStepIdNil sets the value for CurrentStepId to be an explicit nil

### UnsetCurrentStepId
`func (o *PaymentRequestsIndexResponsePaymentRequests) UnsetCurrentStepId()`

UnsetCurrentStepId ensures that no value is present for CurrentStepId, not even an explicit nil
### GetCurrentRound

`func (o *PaymentRequestsIndexResponsePaymentRequests) GetCurrentRound() int32`

GetCurrentRound returns the CurrentRound field if non-nil, zero value otherwise.

### GetCurrentRoundOk

`func (o *PaymentRequestsIndexResponsePaymentRequests) GetCurrentRoundOk() (*int32, bool)`

GetCurrentRoundOk returns a tuple with the CurrentRound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentRound

`func (o *PaymentRequestsIndexResponsePaymentRequests) SetCurrentRound(v int32)`

SetCurrentRound sets CurrentRound field to given value.


### GetIssueDate

`func (o *PaymentRequestsIndexResponsePaymentRequests) GetIssueDate() string`

GetIssueDate returns the IssueDate field if non-nil, zero value otherwise.

### GetIssueDateOk

`func (o *PaymentRequestsIndexResponsePaymentRequests) GetIssueDateOk() (*string, bool)`

GetIssueDateOk returns a tuple with the IssueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueDate

`func (o *PaymentRequestsIndexResponsePaymentRequests) SetIssueDate(v string)`

SetIssueDate sets IssueDate field to given value.


### GetPaymentDate

`func (o *PaymentRequestsIndexResponsePaymentRequests) GetPaymentDate() string`

GetPaymentDate returns the PaymentDate field if non-nil, zero value otherwise.

### GetPaymentDateOk

`func (o *PaymentRequestsIndexResponsePaymentRequests) GetPaymentDateOk() (*string, bool)`

GetPaymentDateOk returns a tuple with the PaymentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentDate

`func (o *PaymentRequestsIndexResponsePaymentRequests) SetPaymentDate(v string)`

SetPaymentDate sets PaymentDate field to given value.


### SetPaymentDateNil

`func (o *PaymentRequestsIndexResponsePaymentRequests) SetPaymentDateNil(b bool)`

 SetPaymentDateNil sets the value for PaymentDate to be an explicit nil

### UnsetPaymentDate
`func (o *PaymentRequestsIndexResponsePaymentRequests) UnsetPaymentDate()`

UnsetPaymentDate ensures that no value is present for PaymentDate, not even an explicit nil
### GetPaymentMethod

`func (o *PaymentRequestsIndexResponsePaymentRequests) GetPaymentMethod() string`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *PaymentRequestsIndexResponsePaymentRequests) GetPaymentMethodOk() (*string, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *PaymentRequestsIndexResponsePaymentRequests) SetPaymentMethod(v string)`

SetPaymentMethod sets PaymentMethod field to given value.


### GetPartnerId

`func (o *PaymentRequestsIndexResponsePaymentRequests) GetPartnerId() int32`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *PaymentRequestsIndexResponsePaymentRequests) GetPartnerIdOk() (*int32, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *PaymentRequestsIndexResponsePaymentRequests) SetPartnerId(v int32)`

SetPartnerId sets PartnerId field to given value.


### SetPartnerIdNil

`func (o *PaymentRequestsIndexResponsePaymentRequests) SetPartnerIdNil(b bool)`

 SetPartnerIdNil sets the value for PartnerId to be an explicit nil

### UnsetPartnerId
`func (o *PaymentRequestsIndexResponsePaymentRequests) UnsetPartnerId()`

UnsetPartnerId ensures that no value is present for PartnerId, not even an explicit nil
### GetPartnerCode

`func (o *PaymentRequestsIndexResponsePaymentRequests) GetPartnerCode() string`

GetPartnerCode returns the PartnerCode field if non-nil, zero value otherwise.

### GetPartnerCodeOk

`func (o *PaymentRequestsIndexResponsePaymentRequests) GetPartnerCodeOk() (*string, bool)`

GetPartnerCodeOk returns a tuple with the PartnerCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerCode

`func (o *PaymentRequestsIndexResponsePaymentRequests) SetPartnerCode(v string)`

SetPartnerCode sets PartnerCode field to given value.


### SetPartnerCodeNil

`func (o *PaymentRequestsIndexResponsePaymentRequests) SetPartnerCodeNil(b bool)`

 SetPartnerCodeNil sets the value for PartnerCode to be an explicit nil

### UnsetPartnerCode
`func (o *PaymentRequestsIndexResponsePaymentRequests) UnsetPartnerCode()`

UnsetPartnerCode ensures that no value is present for PartnerCode, not even an explicit nil
### GetPartnerName

`func (o *PaymentRequestsIndexResponsePaymentRequests) GetPartnerName() string`

GetPartnerName returns the PartnerName field if non-nil, zero value otherwise.

### GetPartnerNameOk

`func (o *PaymentRequestsIndexResponsePaymentRequests) GetPartnerNameOk() (*string, bool)`

GetPartnerNameOk returns a tuple with the PartnerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerName

`func (o *PaymentRequestsIndexResponsePaymentRequests) SetPartnerName(v string)`

SetPartnerName sets PartnerName field to given value.


### SetPartnerNameNil

`func (o *PaymentRequestsIndexResponsePaymentRequests) SetPartnerNameNil(b bool)`

 SetPartnerNameNil sets the value for PartnerName to be an explicit nil

### UnsetPartnerName
`func (o *PaymentRequestsIndexResponsePaymentRequests) UnsetPartnerName()`

UnsetPartnerName ensures that no value is present for PartnerName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


