# \SettingsServiceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SettingsServiceGet**](SettingsServiceApi.md#SettingsServiceGet) | **Get** /api/v1/settings | Get returns Argo CD settings



## SettingsServiceGet

> ClusterSettings SettingsServiceGet(ctx).Execute()

Get returns Argo CD settings

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
    resp, r, err := apiClient.SettingsServiceApi.SettingsServiceGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsServiceApi.SettingsServiceGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SettingsServiceGet`: ClusterSettings
    fmt.Fprintf(os.Stdout, "Response from `SettingsServiceApi.SettingsServiceGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSettingsServiceGetRequest struct via the builder pattern


### Return type

[**ClusterSettings**](ClusterSettings.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

