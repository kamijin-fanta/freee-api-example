# ExpenseApplicationsIndexResponseExpenseApplications

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | 経費申請ID | 
**CompanyId** | **int32** | 事業所ID | 
**Title** | **string** | 申請タイトル | 
**IssueDate** | **string** | 申請日 (yyyy-mm-dd) | 
**Description** | Pointer to **string** | 備考 | [optional] 
**EditableOnWeb** | **bool** | 会計freeeのWeb画面から申請内容を編集可能：falseの場合、Web上からの項目行の追加／削除・金額の編集が出来なくなります。APIでの編集は可能です。 | 
**TotalAmount** | Pointer to **int32** | 合計金額 | [optional] 
**Status** | **string** | 申請ステータス(draft:下書き, in_progress:申請中, approved:承認済, rejected:却下, feedback:差戻し) | 
**SectionId** | Pointer to **NullableInt32** | 部門ID | [optional] 
**TagIds** | Pointer to **[]int32** | メモタグID | [optional] 
**ExpenseApplicationLines** | [**[]ExpenseApplicationsIndexResponseExpenseApplicationLines**](ExpenseApplicationsIndexResponseExpenseApplicationLines.md) | 経費申請の項目行一覧（配列） | 
**DealId** | **NullableInt32** | 取引ID (申請ステータス:statusがapprovedで、取引が存在する時のみdeal_idが表示されます) | 
**DealStatus** | **NullableString** | 取引ステータス (申請ステータス:statusがapprovedで、取引が存在する時のみdeal_statusが表示されます settled:精算済み, unsettled:清算待ち) | 
**ApplicantId** | **int32** | 申請者のユーザーID | 
**ApplicationNumber** | **string** | 申請No. | 
**CurrentStepId** | Pointer to **NullableInt32** | 現在承認ステップID | [optional] 
**CurrentRound** | Pointer to **int32** | 現在のround。差し戻し等により申請がstepの最初からやり直しになるとroundの値が増えます。 | [optional] 
**Segment1TagId** | Pointer to **NullableInt32** | セグメント１ID | [optional] 
**Segment2TagId** | Pointer to **NullableInt32** | セグメント２ID | [optional] 
**Segment3TagId** | Pointer to **NullableInt32** | セグメント３ID | [optional] 

## Methods

### NewExpenseApplicationsIndexResponseExpenseApplications

`func NewExpenseApplicationsIndexResponseExpenseApplications(id int32, companyId int32, title string, issueDate string, editableOnWeb bool, status string, expenseApplicationLines []ExpenseApplicationsIndexResponseExpenseApplicationLines, dealId NullableInt32, dealStatus NullableString, applicantId int32, applicationNumber string, ) *ExpenseApplicationsIndexResponseExpenseApplications`

NewExpenseApplicationsIndexResponseExpenseApplications instantiates a new ExpenseApplicationsIndexResponseExpenseApplications object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExpenseApplicationsIndexResponseExpenseApplicationsWithDefaults

`func NewExpenseApplicationsIndexResponseExpenseApplicationsWithDefaults() *ExpenseApplicationsIndexResponseExpenseApplications`

NewExpenseApplicationsIndexResponseExpenseApplicationsWithDefaults instantiates a new ExpenseApplicationsIndexResponseExpenseApplications object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ExpenseApplicationsIndexResponseExpenseApplications) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExpenseApplicationsIndexResponseExpenseApplications) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExpenseApplicationsIndexResponseExpenseApplications) SetId(v int32)`

SetId sets Id field to given value.


### GetCompanyId

`func (o *ExpenseApplicationsIndexResponseExpenseApplications) GetCompanyId() int32`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *ExpenseApplicationsIndexResponseExpenseApplications) GetCompanyIdOk() (*int32, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *ExpenseApplicationsIndexResponseExpenseApplications) SetCompanyId(v int32)`

SetCompanyId sets CompanyId field to given value.


### GetTitle

`func (o *ExpenseApplicationsIndexResponseExpenseApplications) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ExpenseApplicationsIndexResponseExpenseApplications) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ExpenseApplicationsIndexResponseExpenseApplications) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetIssueDate

`func (o *ExpenseApplicationsIndexResponseExpenseApplications) GetIssueDate() string`

GetIssueDate returns the IssueDate field if non-nil, zero value otherwise.

### GetIssueDateOk

`func (o *ExpenseApplicationsIndexResponseExpenseApplications) GetIssueDateOk() (*string, bool)`

GetIssueDateOk returns a tuple with the IssueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueDate

`func (o *ExpenseApplicationsIndexResponseExpenseApplications) SetIssueDate(v string)`

SetIssueDate sets IssueDate field to given value.


### GetDescription

`func (o *ExpenseApplicationsIndexResponseExpenseApplications) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ExpenseApplicationsIndexResponseExpenseApplications) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ExpenseApplicationsIndexResponseExpenseApplications) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ExpenseApplicationsIndexResponseExpenseApplications) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEditableOnWeb

`func (o *ExpenseApplicationsIndexResponseExpenseApplications) GetEditableOnWeb() bool`

GetEditableOnWeb returns the EditableOnWeb field if non-nil, zero value otherwise.

### GetEditableOnWebOk

`func (o *ExpenseApplicationsIndexResponseExpenseApplications) GetEditableOnWebOk() (*bool, bool)`

GetEditableOnWebOk returns a tuple with the EditableOnWeb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditableOnWeb

`func (o *ExpenseApplicationsIndexResponseExpenseApplications) SetEditableOnWeb(v bool)`

SetEditableOnWeb sets EditableOnWeb field to given value.


### GetTotalAmount

`func (o *ExpenseApplicationsIndexResponseExpenseApplications) GetTotalAmount() int32`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *ExpenseApplicationsIndexResponseExpenseApplications) GetTotalAmountOk() (*int32, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *ExpenseApplicationsIndexResponseExpenseApplications) SetTotalAmount(v int32)`

SetTotalAmount sets TotalAmount field to given value.

### HasTotalAmount

`func (o *ExpenseApplicationsIndexResponseExpenseApplications) HasTotalAmount() bool`

HasTotalAmount returns a boolean if a field has been set.

### GetStatus

`func (o *ExpenseApplicationsIndexResponseExpenseApplications) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ExpenseApplicationsIndexResponseExpenseApplications) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ExpenseApplicationsIndexResponseExpenseApplications) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetSectionId

`func (o *ExpenseApplicationsIndexResponseExpenseApplications) GetSectionId() int32`

GetSectionId returns the SectionId field if non-nil, zero value otherwise.

### GetSectionIdOk

`func (o *ExpenseApplicationsIndexResponseExpenseApplications) GetSectionIdOk() (*int32, bool)`

GetSectionIdOk returns a tuple with the SectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionId

`func (o *ExpenseApplicationsIndexResponseExpenseApplications) SetSectionId(v int32)`

SetSectionId sets SectionId field to given value.

### HasSectionId

`func (o *ExpenseApplicationsIndexResponseExpenseApplications) HasSectionId() bool`

HasSectionId returns a boolean if a field has been set.

### SetSectionIdNil

`func (o *ExpenseApplicationsIndexResponseExpenseApplications) SetSectionIdNil(b bool)`

 SetSectionIdNil sets the value for SectionId to be an explicit nil

### UnsetSectionId
`func (o *ExpenseApplicationsIndexResponseExpenseApplications) UnsetSectionId()`

UnsetSectionId ensures that no value is present for SectionId, not even an explicit nil
### GetTagIds

`func (o *ExpenseApplicationsIndexResponseExpenseApplications) GetTagIds() []int32`

GetTagIds returns the TagIds field if non-nil, zero value otherwise.

### GetTagIdsOk

`func (o *ExpenseApplicationsIndexResponseExpenseApplications) GetTagIdsOk() (*[]int32, bool)`

GetTagIdsOk returns a tuple with the TagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagIds

`func (o *ExpenseApplicationsIndexResponseExpenseApplications) SetTagIds(v []int32)`

SetTagIds sets TagIds field to given value.

### HasTagIds

`func (o *ExpenseApplicationsIndexResponseExpenseApplications) HasTagIds() bool`

HasTagIds returns a boolean if a field has been set.

### GetExpenseApplicationLines

`func (o *ExpenseApplicationsIndexResponseExpenseApplications) GetExpenseApplicationLines() []ExpenseApplicationsIndexResponseExpenseApplicationLines`

GetExpenseApplicationLines returns the ExpenseApplicationLines field if non-nil, zero value otherwise.

### GetExpenseApplicationLinesOk

`func (o *ExpenseApplicationsIndexResponseExpenseApplications) GetExpenseApplicationLinesOk() (*[]ExpenseApplicationsIndexResponseExpenseApplicationLines, bool)`

GetExpenseApplicationLinesOk returns a tuple with the ExpenseApplicationLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpenseApplicationLines

`func (o *ExpenseApplicationsIndexResponseExpenseApplications) SetExpenseApplicationLines(v []ExpenseApplicationsIndexResponseExpenseApplicationLines)`

SetExpenseApplicationLines sets ExpenseApplicationLines field to given value.


### GetDealId

`func (o *ExpenseApplicationsIndexResponseExpenseApplications) GetDealId() int32`

GetDealId returns the DealId field if non-nil, zero value otherwise.

### GetDealIdOk

`func (o *ExpenseApplicationsIndexResponseExpenseApplications) GetDealIdOk() (*int32, bool)`

GetDealIdOk returns a tuple with the DealId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDealId

`func (o *ExpenseApplicationsIndexResponseExpenseApplications) SetDealId(v int32)`

SetDealId sets DealId field to given value.


### SetDealIdNil

`func (o *ExpenseApplicationsIndexResponseExpenseApplications) SetDealIdNil(b bool)`

 SetDealIdNil sets the value for DealId to be an explicit nil

### UnsetDealId
`func (o *ExpenseApplicationsIndexResponseExpenseApplications) UnsetDealId()`

UnsetDealId ensures that no value is present for DealId, not even an explicit nil
### GetDealStatus

`func (o *ExpenseApplicationsIndexResponseExpenseApplications) GetDealStatus() string`

GetDealStatus returns the DealStatus field if non-nil, zero value otherwise.

### GetDealStatusOk

`func (o *ExpenseApplicationsIndexResponseExpenseApplications) GetDealStatusOk() (*string, bool)`

GetDealStatusOk returns a tuple with the DealStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDealStatus

`func (o *ExpenseApplicationsIndexResponseExpenseApplications) SetDealStatus(v string)`

SetDealStatus sets DealStatus field to given value.


### SetDealStatusNil

`func (o *ExpenseApplicationsIndexResponseExpenseApplications) SetDealStatusNil(b bool)`

 SetDealStatusNil sets the value for DealStatus to be an explicit nil

### UnsetDealStatus
`func (o *ExpenseApplicationsIndexResponseExpenseApplications) UnsetDealStatus()`

UnsetDealStatus ensures that no value is present for DealStatus, not even an explicit nil
### GetApplicantId

`func (o *ExpenseApplicationsIndexResponseExpenseApplications) GetApplicantId() int32`

GetApplicantId returns the ApplicantId field if non-nil, zero value otherwise.

### GetApplicantIdOk

`func (o *ExpenseApplicationsIndexResponseExpenseApplications) GetApplicantIdOk() (*int32, bool)`

GetApplicantIdOk returns a tuple with the ApplicantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicantId

`func (o *ExpenseApplicationsIndexResponseExpenseApplications) SetApplicantId(v int32)`

SetApplicantId sets ApplicantId field to given value.


### GetApplicationNumber

`func (o *ExpenseApplicationsIndexResponseExpenseApplications) GetApplicationNumber() string`

GetApplicationNumber returns the ApplicationNumber field if non-nil, zero value otherwise.

### GetApplicationNumberOk

`func (o *ExpenseApplicationsIndexResponseExpenseApplications) GetApplicationNumberOk() (*string, bool)`

GetApplicationNumberOk returns a tuple with the ApplicationNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationNumber

`func (o *ExpenseApplicationsIndexResponseExpenseApplications) SetApplicationNumber(v string)`

SetApplicationNumber sets ApplicationNumber field to given value.


### GetCurrentStepId

`func (o *ExpenseApplicationsIndexResponseExpenseApplications) GetCurrentStepId() int32`

GetCurrentStepId returns the CurrentStepId field if non-nil, zero value otherwise.

### GetCurrentStepIdOk

`func (o *ExpenseApplicationsIndexResponseExpenseApplications) GetCurrentStepIdOk() (*int32, bool)`

GetCurrentStepIdOk returns a tuple with the CurrentStepId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentStepId

`func (o *ExpenseApplicationsIndexResponseExpenseApplications) SetCurrentStepId(v int32)`

SetCurrentStepId sets CurrentStepId field to given value.

### HasCurrentStepId

`func (o *ExpenseApplicationsIndexResponseExpenseApplications) HasCurrentStepId() bool`

HasCurrentStepId returns a boolean if a field has been set.

### SetCurrentStepIdNil

`func (o *ExpenseApplicationsIndexResponseExpenseApplications) SetCurrentStepIdNil(b bool)`

 SetCurrentStepIdNil sets the value for CurrentStepId to be an explicit nil

### UnsetCurrentStepId
`func (o *ExpenseApplicationsIndexResponseExpenseApplications) UnsetCurrentStepId()`

UnsetCurrentStepId ensures that no value is present for CurrentStepId, not even an explicit nil
### GetCurrentRound

`func (o *ExpenseApplicationsIndexResponseExpenseApplications) GetCurrentRound() int32`

GetCurrentRound returns the CurrentRound field if non-nil, zero value otherwise.

### GetCurrentRoundOk

`func (o *ExpenseApplicationsIndexResponseExpenseApplications) GetCurrentRoundOk() (*int32, bool)`

GetCurrentRoundOk returns a tuple with the CurrentRound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentRound

`func (o *ExpenseApplicationsIndexResponseExpenseApplications) SetCurrentRound(v int32)`

SetCurrentRound sets CurrentRound field to given value.

### HasCurrentRound

`func (o *ExpenseApplicationsIndexResponseExpenseApplications) HasCurrentRound() bool`

HasCurrentRound returns a boolean if a field has been set.

### GetSegment1TagId

`func (o *ExpenseApplicationsIndexResponseExpenseApplications) GetSegment1TagId() int32`

GetSegment1TagId returns the Segment1TagId field if non-nil, zero value otherwise.

### GetSegment1TagIdOk

`func (o *ExpenseApplicationsIndexResponseExpenseApplications) GetSegment1TagIdOk() (*int32, bool)`

GetSegment1TagIdOk returns a tuple with the Segment1TagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment1TagId

`func (o *ExpenseApplicationsIndexResponseExpenseApplications) SetSegment1TagId(v int32)`

SetSegment1TagId sets Segment1TagId field to given value.

### HasSegment1TagId

`func (o *ExpenseApplicationsIndexResponseExpenseApplications) HasSegment1TagId() bool`

HasSegment1TagId returns a boolean if a field has been set.

### SetSegment1TagIdNil

`func (o *ExpenseApplicationsIndexResponseExpenseApplications) SetSegment1TagIdNil(b bool)`

 SetSegment1TagIdNil sets the value for Segment1TagId to be an explicit nil

### UnsetSegment1TagId
`func (o *ExpenseApplicationsIndexResponseExpenseApplications) UnsetSegment1TagId()`

UnsetSegment1TagId ensures that no value is present for Segment1TagId, not even an explicit nil
### GetSegment2TagId

`func (o *ExpenseApplicationsIndexResponseExpenseApplications) GetSegment2TagId() int32`

GetSegment2TagId returns the Segment2TagId field if non-nil, zero value otherwise.

### GetSegment2TagIdOk

`func (o *ExpenseApplicationsIndexResponseExpenseApplications) GetSegment2TagIdOk() (*int32, bool)`

GetSegment2TagIdOk returns a tuple with the Segment2TagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment2TagId

`func (o *ExpenseApplicationsIndexResponseExpenseApplications) SetSegment2TagId(v int32)`

SetSegment2TagId sets Segment2TagId field to given value.

### HasSegment2TagId

`func (o *ExpenseApplicationsIndexResponseExpenseApplications) HasSegment2TagId() bool`

HasSegment2TagId returns a boolean if a field has been set.

### SetSegment2TagIdNil

`func (o *ExpenseApplicationsIndexResponseExpenseApplications) SetSegment2TagIdNil(b bool)`

 SetSegment2TagIdNil sets the value for Segment2TagId to be an explicit nil

### UnsetSegment2TagId
`func (o *ExpenseApplicationsIndexResponseExpenseApplications) UnsetSegment2TagId()`

UnsetSegment2TagId ensures that no value is present for Segment2TagId, not even an explicit nil
### GetSegment3TagId

`func (o *ExpenseApplicationsIndexResponseExpenseApplications) GetSegment3TagId() int32`

GetSegment3TagId returns the Segment3TagId field if non-nil, zero value otherwise.

### GetSegment3TagIdOk

`func (o *ExpenseApplicationsIndexResponseExpenseApplications) GetSegment3TagIdOk() (*int32, bool)`

GetSegment3TagIdOk returns a tuple with the Segment3TagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment3TagId

`func (o *ExpenseApplicationsIndexResponseExpenseApplications) SetSegment3TagId(v int32)`

SetSegment3TagId sets Segment3TagId field to given value.

### HasSegment3TagId

`func (o *ExpenseApplicationsIndexResponseExpenseApplications) HasSegment3TagId() bool`

HasSegment3TagId returns a boolean if a field has been set.

### SetSegment3TagIdNil

`func (o *ExpenseApplicationsIndexResponseExpenseApplications) SetSegment3TagIdNil(b bool)`

 SetSegment3TagIdNil sets the value for Segment3TagId to be an explicit nil

### UnsetSegment3TagId
`func (o *ExpenseApplicationsIndexResponseExpenseApplications) UnsetSegment3TagId()`

UnsetSegment3TagId ensures that no value is present for Segment3TagId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


