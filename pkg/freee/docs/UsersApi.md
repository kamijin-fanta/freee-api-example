# \UsersApi

All URIs are relative to *https://api.freee.co.jp*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUsers**](UsersApi.md#GetUsers) | **Get** /api/1/users | 事業所に所属するユーザー一覧の取得
[**GetUsersCapabilities**](UsersApi.md#GetUsersCapabilities) | **Get** /api/1/users/capabilities | ログインユーザーの権限の取得
[**GetUsersMe**](UsersApi.md#GetUsersMe) | **Get** /api/1/users/me | ログインユーザー情報の取得
[**UpdateUser**](UsersApi.md#UpdateUser) | **Put** /api/1/users/me | ユーザー情報の更新



## GetUsers

> InlineResponse2004 GetUsers(ctx).CompanyId(companyId).Limit(limit).Execute()

事業所に所属するユーザー一覧の取得



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
    limit := int32(56) // int32 | 取得レコードの件数 (デフォルト: 50, 最小: 1, 最大: 3000) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersApi.GetUsers(context.Background()).CompanyId(companyId).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.GetUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUsers`: InlineResponse2004
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.GetUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyId** | **int32** | 事業所ID | 
 **limit** | **int32** | 取得レコードの件数 (デフォルト: 50, 最小: 1, 最大: 3000) | 

### Return type

[**InlineResponse2004**](inline_response_200_4.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsersCapabilities

> InlineResponse2005 GetUsersCapabilities(ctx).CompanyId(companyId).Execute()

ログインユーザーの権限の取得



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersApi.GetUsersCapabilities(context.Background()).CompanyId(companyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.GetUsersCapabilities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUsersCapabilities`: InlineResponse2005
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.GetUsersCapabilities`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUsersCapabilitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companyId** | **int32** | 事業所ID | 

### Return type

[**InlineResponse2005**](inline_response_200_5.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsersMe

> MeResponse GetUsersMe(ctx).Companies(companies).Execute()

ログインユーザー情報の取得



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
    companies := true // bool | 取得情報にユーザーが所属する事業所一覧を含める (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersApi.GetUsersMe(context.Background()).Companies(companies).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.GetUsersMe``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUsersMe`: MeResponse
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.GetUsersMe`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUsersMeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **companies** | **bool** | 取得情報にユーザーが所属する事業所一覧を含める | 

### Return type

[**MeResponse**](meResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUser

> UserResponse UpdateUser(ctx).UserParams(userParams).Execute()

ユーザー情報の更新



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
    userParams := *openapiclient.NewUserParams() // UserParams | ユーザー情報の更新 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersApi.UpdateUser(context.Background()).UserParams(userParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UpdateUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateUser`: UserResponse
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UpdateUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userParams** | [**UserParams**](UserParams.md) | ユーザー情報の更新 | 

### Return type

[**UserResponse**](userResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

