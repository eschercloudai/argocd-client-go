# \GPGKeyServiceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GPGKeyServiceCreate**](GPGKeyServiceApi.md#GPGKeyServiceCreate) | **Post** /api/v1/gpgkeys | Create one or more GPG public keys in the server&#39;s configuration
[**GPGKeyServiceDelete**](GPGKeyServiceApi.md#GPGKeyServiceDelete) | **Delete** /api/v1/gpgkeys | Delete specified GPG public key from the server&#39;s configuration
[**GPGKeyServiceGet**](GPGKeyServiceApi.md#GPGKeyServiceGet) | **Get** /api/v1/gpgkeys/{keyID} | Get information about specified GPG public key from the server
[**GPGKeyServiceList**](GPGKeyServiceApi.md#GPGKeyServiceList) | **Get** /api/v1/gpgkeys | List all available repository certificates



## GPGKeyServiceCreate

> GpgkeyGnuPGPublicKeyCreateResponse GPGKeyServiceCreate(ctx).Body(body).Upsert(upsert).Execute()

Create one or more GPG public keys in the server's configuration

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
    body := *openapiclient.NewV1alpha1GnuPGPublicKey() // V1alpha1GnuPGPublicKey | Raw key data of the GPG key(s) to create
    upsert := true // bool | Whether to upsert already existing public keys. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GPGKeyServiceApi.GPGKeyServiceCreate(context.Background()).Body(body).Upsert(upsert).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GPGKeyServiceApi.GPGKeyServiceCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GPGKeyServiceCreate`: GpgkeyGnuPGPublicKeyCreateResponse
    fmt.Fprintf(os.Stdout, "Response from `GPGKeyServiceApi.GPGKeyServiceCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGPGKeyServiceCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**V1alpha1GnuPGPublicKey**](V1alpha1GnuPGPublicKey.md) | Raw key data of the GPG key(s) to create | 
 **upsert** | **bool** | Whether to upsert already existing public keys. | 

### Return type

[**GpgkeyGnuPGPublicKeyCreateResponse**](GpgkeyGnuPGPublicKeyCreateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GPGKeyServiceDelete

> map[string]interface{} GPGKeyServiceDelete(ctx).KeyID(keyID).Execute()

Delete specified GPG public key from the server's configuration

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
    keyID := "keyID_example" // string | The GPG key ID to query for. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GPGKeyServiceApi.GPGKeyServiceDelete(context.Background()).KeyID(keyID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GPGKeyServiceApi.GPGKeyServiceDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GPGKeyServiceDelete`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `GPGKeyServiceApi.GPGKeyServiceDelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGPGKeyServiceDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **keyID** | **string** | The GPG key ID to query for. | 

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


## GPGKeyServiceGet

> V1alpha1GnuPGPublicKey GPGKeyServiceGet(ctx, keyID).Execute()

Get information about specified GPG public key from the server

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
    keyID := "keyID_example" // string | The GPG key ID to query for

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GPGKeyServiceApi.GPGKeyServiceGet(context.Background(), keyID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GPGKeyServiceApi.GPGKeyServiceGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GPGKeyServiceGet`: V1alpha1GnuPGPublicKey
    fmt.Fprintf(os.Stdout, "Response from `GPGKeyServiceApi.GPGKeyServiceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyID** | **string** | The GPG key ID to query for | 

### Other Parameters

Other parameters are passed through a pointer to a apiGPGKeyServiceGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V1alpha1GnuPGPublicKey**](V1alpha1GnuPGPublicKey.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GPGKeyServiceList

> V1alpha1GnuPGPublicKeyList GPGKeyServiceList(ctx).KeyID(keyID).Execute()

List all available repository certificates

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
    keyID := "keyID_example" // string | The GPG key ID to query for. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GPGKeyServiceApi.GPGKeyServiceList(context.Background()).KeyID(keyID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GPGKeyServiceApi.GPGKeyServiceList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GPGKeyServiceList`: V1alpha1GnuPGPublicKeyList
    fmt.Fprintf(os.Stdout, "Response from `GPGKeyServiceApi.GPGKeyServiceList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGPGKeyServiceListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **keyID** | **string** | The GPG key ID to query for. | 

### Return type

[**V1alpha1GnuPGPublicKeyList**](V1alpha1GnuPGPublicKeyList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

