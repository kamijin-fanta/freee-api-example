# ReceiptResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Receipt** | [**DealResponseDealReceipts**](dealResponse_deal_receipts.md) |  | 

## Methods

### NewReceiptResponse

`func NewReceiptResponse(receipt DealResponseDealReceipts, ) *ReceiptResponse`

NewReceiptResponse instantiates a new ReceiptResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReceiptResponseWithDefaults

`func NewReceiptResponseWithDefaults() *ReceiptResponse`

NewReceiptResponseWithDefaults instantiates a new ReceiptResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReceipt

`func (o *ReceiptResponse) GetReceipt() DealResponseDealReceipts`

GetReceipt returns the Receipt field if non-nil, zero value otherwise.

### GetReceiptOk

`func (o *ReceiptResponse) GetReceiptOk() (*DealResponseDealReceipts, bool)`

GetReceiptOk returns a tuple with the Receipt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceipt

`func (o *ReceiptResponse) SetReceipt(v DealResponseDealReceipts)`

SetReceipt sets Receipt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


