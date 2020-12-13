# \ApprovalFlowRoutesApi

All URIs are relative to *https://api.freee.co.jp*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApprovalFlowRoute**](ApprovalFlowRoutesApi.md#GetApprovalFlowRoute) | **Get** /api/1/approval_flow_routes/{id} | 申請経路の取得
[**GetApprovalFlowRoutes**](ApprovalFlowRoutesApi.md#GetApprovalFlowRoutes) | **Get** /api/1/approval_flow_routes | 申請経路一覧の取得



## GetApprovalFlowRoute

> ApprovalFlowRouteResponse GetApprovalFlowRoute(ctx, id).CompanyId(companyId).Execute()

申請経路の取得



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := int32(56) // int32 | 経路申請ID
    companyId := int32(56) // int32 | 事業所ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ApprovalFlowRoutesApi.GetApprovalFlowRoute(context.Background(), id).CompanyId(companyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApprovalFlowRoutesApi.GetApprovalFlowRoute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApprovalFlowRoute`: ApprovalFlowRouteResponse
    fmt.Fprintf(os.Stdout, "Response from `ApprovalFlowRoutesApi.GetApprovalFlowRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | 経路申請ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApprovalFlowRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **companyId** | **int32** | 事業所ID | 

### Return type

[**ApprovalFlowRouteResponse**](approvalFlowRouteResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApprovalFlowRoutes

> ApprovalFlowRoutesIndexResponse GetApprovalFlowRoutes(ctx).CompanyId(companyId).IncludedUserId(includedUserId).Usage(usage).RequestFormId(requestFormId).Execute()

申請経路一覧の取得



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    companyId := int32(56) // int32 | 事業所ID
    includedUserId := int32(56) // int32 | 経路に含まれるユーザーのユーザーID (optional)
    usage := "usage_example" // string | 申請種別（各申請種別が使用できる申請経路に絞り込めます。例えば、ApprovalRequest を指定すると、各種申請が使用できる申請経路に絞り込めます。） * `TxnApproval` - 仕訳承認 * `ExpenseApplication` - 経費精算 * `PaymentRequest` - 支払依頼 * `ApprovalRequest` - 各種申請 * `DocApproval` - 請求書等 (見積書・納品書・請求書・発注書) (optional)
    requestFormId := int32(56) // int32 | 申請フォームID request_form_id指定時はusage条件をApprovalRequestに指定してください。指定しない場合無効になります。 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ApprovalFlowRoutesApi.GetApprovalFlowRoutes(context.Background()).CompanyId(companyId).IncludedUserId(includedUserId).Usage(usage).RequestFormId(requestFormId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApprovalFlowRoutesApi.GetApprovalFlowRoutes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApprovalFlowRoutes`: ApprovalFlowRoutesIndexResponse
    fmt.Fprintf(os.Stdout, "Response from `ApprovalFlowRoutesApi.GetApprovalFlowRoutes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetApprovalFlowRoutesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyId** | **int32** | 事業所ID | 
 **includedUserId** | **int32** | 経路に含まれるユーザーのユーザーID | 
 **usage** | **string** | 申請種別（各申請種別が使用できる申請経路に絞り込めます。例えば、ApprovalRequest を指定すると、各種申請が使用できる申請経路に絞り込めます。） * &#x60;TxnApproval&#x60; - 仕訳承認 * &#x60;ExpenseApplication&#x60; - 経費精算 * &#x60;PaymentRequest&#x60; - 支払依頼 * &#x60;ApprovalRequest&#x60; - 各種申請 * &#x60;DocApproval&#x60; - 請求書等 (見積書・納品書・請求書・発注書) | 
 **requestFormId** | **int32** | 申請フォームID request_form_id指定時はusage条件をApprovalRequestに指定してください。指定しない場合無効になります。 | 

### Return type

[**ApprovalFlowRoutesIndexResponse**](approvalFlowRoutesIndexResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

