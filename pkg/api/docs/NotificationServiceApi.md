# \NotificationServiceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NotificationServiceListServices**](NotificationServiceApi.md#NotificationServiceListServices) | **Get** /api/v1/notifications/services | List returns list of services
[**NotificationServiceListTemplates**](NotificationServiceApi.md#NotificationServiceListTemplates) | **Get** /api/v1/notifications/templates | List returns list of templates
[**NotificationServiceListTriggers**](NotificationServiceApi.md#NotificationServiceListTriggers) | **Get** /api/v1/notifications/triggers | List returns list of triggers



## NotificationServiceListServices

> NotificationServiceList NotificationServiceListServices(ctx).Execute()

List returns list of services

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
    resp, r, err := apiClient.NotificationServiceApi.NotificationServiceListServices(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationServiceApi.NotificationServiceListServices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NotificationServiceListServices`: NotificationServiceList
    fmt.Fprintf(os.Stdout, "Response from `NotificationServiceApi.NotificationServiceListServices`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiNotificationServiceListServicesRequest struct via the builder pattern


### Return type

[**NotificationServiceList**](NotificationServiceList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NotificationServiceListTemplates

> NotificationTemplateList NotificationServiceListTemplates(ctx).Execute()

List returns list of templates

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
    resp, r, err := apiClient.NotificationServiceApi.NotificationServiceListTemplates(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationServiceApi.NotificationServiceListTemplates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NotificationServiceListTemplates`: NotificationTemplateList
    fmt.Fprintf(os.Stdout, "Response from `NotificationServiceApi.NotificationServiceListTemplates`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiNotificationServiceListTemplatesRequest struct via the builder pattern


### Return type

[**NotificationTemplateList**](NotificationTemplateList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NotificationServiceListTriggers

> NotificationTriggerList NotificationServiceListTriggers(ctx).Execute()

List returns list of triggers

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
    resp, r, err := apiClient.NotificationServiceApi.NotificationServiceListTriggers(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationServiceApi.NotificationServiceListTriggers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NotificationServiceListTriggers`: NotificationTriggerList
    fmt.Fprintf(os.Stdout, "Response from `NotificationServiceApi.NotificationServiceListTriggers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiNotificationServiceListTriggersRequest struct via the builder pattern


### Return type

[**NotificationTriggerList**](NotificationTriggerList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

