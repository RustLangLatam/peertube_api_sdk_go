# \MyNotificationsAPI

All URIs are relative to *https://peertube2.cpy.re*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV1UsersMeNotificationSettingsPut**](MyNotificationsAPI.md#ApiV1UsersMeNotificationSettingsPut) | **Put** /api/v1/users/me/notification-settings | Update my notification settings
[**ApiV1UsersMeNotificationsGet**](MyNotificationsAPI.md#ApiV1UsersMeNotificationsGet) | **Get** /api/v1/users/me/notifications | List my notifications
[**ApiV1UsersMeNotificationsReadAllPost**](MyNotificationsAPI.md#ApiV1UsersMeNotificationsReadAllPost) | **Post** /api/v1/users/me/notifications/read-all | Mark all my notification as read
[**ApiV1UsersMeNotificationsReadPost**](MyNotificationsAPI.md#ApiV1UsersMeNotificationsReadPost) | **Post** /api/v1/users/me/notifications/read | Mark notifications as read by their id



## ApiV1UsersMeNotificationSettingsPut

> ApiV1UsersMeNotificationSettingsPut(ctx).ApiV1UsersMeNotificationSettingsPutRequest(apiV1UsersMeNotificationSettingsPutRequest).Execute()

Update my notification settings

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/RustLangLatam/peertube_api_sdk"
)

func main() {
	apiV1UsersMeNotificationSettingsPutRequest := *openapiclient.NewApiV1UsersMeNotificationSettingsPutRequest() // ApiV1UsersMeNotificationSettingsPutRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MyNotificationsAPI.ApiV1UsersMeNotificationSettingsPut(context.Background()).ApiV1UsersMeNotificationSettingsPutRequest(apiV1UsersMeNotificationSettingsPutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MyNotificationsAPI.ApiV1UsersMeNotificationSettingsPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1UsersMeNotificationSettingsPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiV1UsersMeNotificationSettingsPutRequest** | [**ApiV1UsersMeNotificationSettingsPutRequest**](ApiV1UsersMeNotificationSettingsPutRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1UsersMeNotificationsGet

> NotificationListResponse ApiV1UsersMeNotificationsGet(ctx).Unread(unread).Start(start).Count(count).Sort(sort).Execute()

List my notifications

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/RustLangLatam/peertube_api_sdk"
)

func main() {
	unread := true // bool | only list unread notifications (optional)
	start := int32(56) // int32 | Offset used to paginate results (optional)
	count := int32(56) // int32 | Number of items to return (optional) (default to 15)
	sort := "-createdAt" // string | Sort column (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MyNotificationsAPI.ApiV1UsersMeNotificationsGet(context.Background()).Unread(unread).Start(start).Count(count).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MyNotificationsAPI.ApiV1UsersMeNotificationsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1UsersMeNotificationsGet`: NotificationListResponse
	fmt.Fprintf(os.Stdout, "Response from `MyNotificationsAPI.ApiV1UsersMeNotificationsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1UsersMeNotificationsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unread** | **bool** | only list unread notifications | 
 **start** | **int32** | Offset used to paginate results | 
 **count** | **int32** | Number of items to return | [default to 15]
 **sort** | **string** | Sort column | 

### Return type

[**NotificationListResponse**](NotificationListResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1UsersMeNotificationsReadAllPost

> ApiV1UsersMeNotificationsReadAllPost(ctx).Execute()

Mark all my notification as read

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/RustLangLatam/peertube_api_sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MyNotificationsAPI.ApiV1UsersMeNotificationsReadAllPost(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MyNotificationsAPI.ApiV1UsersMeNotificationsReadAllPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1UsersMeNotificationsReadAllPostRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1UsersMeNotificationsReadPost

> ApiV1UsersMeNotificationsReadPost(ctx).ApiV1UsersMeNotificationsReadPostRequest(apiV1UsersMeNotificationsReadPostRequest).Execute()

Mark notifications as read by their id

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/RustLangLatam/peertube_api_sdk"
)

func main() {
	apiV1UsersMeNotificationsReadPostRequest := *openapiclient.NewApiV1UsersMeNotificationsReadPostRequest([]int32{int32(123)}) // ApiV1UsersMeNotificationsReadPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MyNotificationsAPI.ApiV1UsersMeNotificationsReadPost(context.Background()).ApiV1UsersMeNotificationsReadPostRequest(apiV1UsersMeNotificationsReadPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MyNotificationsAPI.ApiV1UsersMeNotificationsReadPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1UsersMeNotificationsReadPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiV1UsersMeNotificationsReadPostRequest** | [**ApiV1UsersMeNotificationsReadPostRequest**](ApiV1UsersMeNotificationsReadPostRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

