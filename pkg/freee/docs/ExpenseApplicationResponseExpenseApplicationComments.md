# ExpenseApplicationResponseExpenseApplicationComments

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | **string** | コメント内容 | 
**UserId** | **int32** | ユーザーID | 
**PostedAt** | **string** | コメント日時(ISO8601形式) | 

## Methods

### NewExpenseApplicationResponseExpenseApplicationComments

`func NewExpenseApplicationResponseExpenseApplicationComments(comment string, userId int32, postedAt string, ) *ExpenseApplicationResponseExpenseApplicationComments`

NewExpenseApplicationResponseExpenseApplicationComments instantiates a new ExpenseApplicationResponseExpenseApplicationComments object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExpenseApplicationResponseExpenseApplicationCommentsWithDefaults

`func NewExpenseApplicationResponseExpenseApplicationCommentsWithDefaults() *ExpenseApplicationResponseExpenseApplicationComments`

NewExpenseApplicationResponseExpenseApplicationCommentsWithDefaults instantiates a new ExpenseApplicationResponseExpenseApplicationComments object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *ExpenseApplicationResponseExpenseApplicationComments) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ExpenseApplicationResponseExpenseApplicationComments) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ExpenseApplicationResponseExpenseApplicationComments) SetComment(v string)`

SetComment sets Comment field to given value.


### GetUserId

`func (o *ExpenseApplicationResponseExpenseApplicationComments) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ExpenseApplicationResponseExpenseApplicationComments) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ExpenseApplicationResponseExpenseApplicationComments) SetUserId(v int32)`

SetUserId sets UserId field to given value.


### GetPostedAt

`func (o *ExpenseApplicationResponseExpenseApplicationComments) GetPostedAt() string`

GetPostedAt returns the PostedAt field if non-nil, zero value otherwise.

### GetPostedAtOk

`func (o *ExpenseApplicationResponseExpenseApplicationComments) GetPostedAtOk() (*string, bool)`

GetPostedAtOk returns a tuple with the PostedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostedAt

`func (o *ExpenseApplicationResponseExpenseApplicationComments) SetPostedAt(v string)`

SetPostedAt sets PostedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


