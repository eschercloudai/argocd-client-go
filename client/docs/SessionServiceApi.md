# \SessionServiceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SessionServiceCreate**](SessionServiceApi.md#SessionServiceCreate) | **Post** /api/v1/session | Create a new JWT for authentication and set a cookie if using HTTP
[**SessionServiceDelete**](SessionServiceApi.md#SessionServiceDelete) | **Delete** /api/v1/session | Delete an existing JWT cookie if using HTTP
[**SessionServiceGetUserInfo**](SessionServiceApi.md#SessionServiceGetUserInfo) | **Get** /api/v1/session/userinfo | Get the current user&#39;s info



## SessionServiceCreate

> SessionSessionResponse SessionServiceCreate(ctx).Body(body).Execute()

Create a new JWT for authentication and set a cookie if using HTTP

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
    body := *openapiclient.NewSessionSessionCreateRequest() // SessionSessionCreateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SessionServiceApi.SessionServiceCreate(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionServiceApi.SessionServiceCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SessionServiceCreate`: SessionSessionResponse
    fmt.Fprintf(os.Stdout, "Response from `SessionServiceApi.SessionServiceCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSessionServiceCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**SessionSessionCreateRequest**](SessionSessionCreateRequest.md) |  | 

### Return type

[**SessionSessionResponse**](SessionSessionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SessionServiceDelete

> SessionSessionResponse SessionServiceDelete(ctx).Execute()

Delete an existing JWT cookie if using HTTP

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
    resp, r, err := apiClient.SessionServiceApi.SessionServiceDelete(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionServiceApi.SessionServiceDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SessionServiceDelete`: SessionSessionResponse
    fmt.Fprintf(os.Stdout, "Response from `SessionServiceApi.SessionServiceDelete`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSessionServiceDeleteRequest struct via the builder pattern


### Return type

[**SessionSessionResponse**](SessionSessionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SessionServiceGetUserInfo

> SessionGetUserInfoResponse SessionServiceGetUserInfo(ctx).Execute()

Get the current user's info

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
    resp, r, err := apiClient.SessionServiceApi.SessionServiceGetUserInfo(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionServiceApi.SessionServiceGetUserInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SessionServiceGetUserInfo`: SessionGetUserInfoResponse
    fmt.Fprintf(os.Stdout, "Response from `SessionServiceApi.SessionServiceGetUserInfo`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSessionServiceGetUserInfoRequest struct via the builder pattern


### Return type

[**SessionGetUserInfoResponse**](SessionGetUserInfoResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

