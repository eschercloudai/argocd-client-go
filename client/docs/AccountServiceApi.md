# \AccountServiceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AccountServiceCanI**](AccountServiceApi.md#AccountServiceCanI) | **Get** /api/v1/account/can-i/{resource}/{action}/{subresource} | CanI checks if the current account has permission to perform an action
[**AccountServiceCreateToken**](AccountServiceApi.md#AccountServiceCreateToken) | **Post** /api/v1/account/{name}/token | CreateToken creates a token
[**AccountServiceDeleteToken**](AccountServiceApi.md#AccountServiceDeleteToken) | **Delete** /api/v1/account/{name}/token/{id} | DeleteToken deletes a token
[**AccountServiceGetAccount**](AccountServiceApi.md#AccountServiceGetAccount) | **Get** /api/v1/account/{name} | GetAccount returns an account
[**AccountServiceListAccounts**](AccountServiceApi.md#AccountServiceListAccounts) | **Get** /api/v1/account | ListAccounts returns the list of accounts
[**AccountServiceUpdatePassword**](AccountServiceApi.md#AccountServiceUpdatePassword) | **Put** /api/v1/account/password | UpdatePassword updates an account&#39;s password to a new value



## AccountServiceCanI

> AccountCanIResponse AccountServiceCanI(ctx, resource, action, subresource).Execute()

CanI checks if the current account has permission to perform an action

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
    resource := "resource_example" // string | 
    action := "action_example" // string | 
    subresource := "subresource_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountServiceApi.AccountServiceCanI(context.Background(), resource, action, subresource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountServiceApi.AccountServiceCanI``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccountServiceCanI`: AccountCanIResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountServiceApi.AccountServiceCanI`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resource** | **string** |  | 
**action** | **string** |  | 
**subresource** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccountServiceCanIRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**AccountCanIResponse**](AccountCanIResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountServiceCreateToken

> AccountCreateTokenResponse AccountServiceCreateToken(ctx, name).Body(body).Execute()

CreateToken creates a token

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
    name := "name_example" // string | 
    body := *openapiclient.NewAccountCreateTokenRequest() // AccountCreateTokenRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountServiceApi.AccountServiceCreateToken(context.Background(), name).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountServiceApi.AccountServiceCreateToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccountServiceCreateToken`: AccountCreateTokenResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountServiceApi.AccountServiceCreateToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccountServiceCreateTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**AccountCreateTokenRequest**](AccountCreateTokenRequest.md) |  | 

### Return type

[**AccountCreateTokenResponse**](AccountCreateTokenResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountServiceDeleteToken

> map[string]interface{} AccountServiceDeleteToken(ctx, name, id).Execute()

DeleteToken deletes a token

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
    name := "name_example" // string | 
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountServiceApi.AccountServiceDeleteToken(context.Background(), name, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountServiceApi.AccountServiceDeleteToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccountServiceDeleteToken`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AccountServiceApi.AccountServiceDeleteToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccountServiceDeleteTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountServiceGetAccount

> AccountAccount AccountServiceGetAccount(ctx, name).Execute()

GetAccount returns an account

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
    name := "name_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountServiceApi.AccountServiceGetAccount(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountServiceApi.AccountServiceGetAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccountServiceGetAccount`: AccountAccount
    fmt.Fprintf(os.Stdout, "Response from `AccountServiceApi.AccountServiceGetAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccountServiceGetAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AccountAccount**](AccountAccount.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountServiceListAccounts

> AccountAccountsList AccountServiceListAccounts(ctx).Execute()

ListAccounts returns the list of accounts

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountServiceApi.AccountServiceListAccounts(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountServiceApi.AccountServiceListAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccountServiceListAccounts`: AccountAccountsList
    fmt.Fprintf(os.Stdout, "Response from `AccountServiceApi.AccountServiceListAccounts`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAccountServiceListAccountsRequest struct via the builder pattern


### Return type

[**AccountAccountsList**](AccountAccountsList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountServiceUpdatePassword

> map[string]interface{} AccountServiceUpdatePassword(ctx).Body(body).Execute()

UpdatePassword updates an account's password to a new value

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
    body := *openapiclient.NewAccountUpdatePasswordRequest() // AccountUpdatePasswordRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountServiceApi.AccountServiceUpdatePassword(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountServiceApi.AccountServiceUpdatePassword``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccountServiceUpdatePassword`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AccountServiceApi.AccountServiceUpdatePassword`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAccountServiceUpdatePasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AccountUpdatePasswordRequest**](AccountUpdatePasswordRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

