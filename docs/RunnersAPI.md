# \RunnersAPI

All URIs are relative to *https://peertube2.cpy.re*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV1RunnersGet**](RunnersAPI.md#ApiV1RunnersGet) | **Get** /api/v1/runners | List runners
[**ApiV1RunnersRegisterPost**](RunnersAPI.md#ApiV1RunnersRegisterPost) | **Post** /api/v1/runners/register | Register a new runner
[**ApiV1RunnersRunnerIdDelete**](RunnersAPI.md#ApiV1RunnersRunnerIdDelete) | **Delete** /api/v1/runners/{runnerId} | Delete a runner
[**ApiV1RunnersUnregisterPost**](RunnersAPI.md#ApiV1RunnersUnregisterPost) | **Post** /api/v1/runners/unregister | Unregister a runner



## ApiV1RunnersGet

> ApiV1RunnersGet200Response ApiV1RunnersGet(ctx).Start(start).Count(count).Sort(sort).Execute()

List runners

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
	start := int32(56) // int32 | Offset used to paginate results (optional)
	count := int32(56) // int32 | Number of items to return (optional) (default to 15)
	sort := "sort_example" // string | Sort runners by criteria (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RunnersAPI.ApiV1RunnersGet(context.Background()).Start(start).Count(count).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunnersAPI.ApiV1RunnersGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1RunnersGet`: ApiV1RunnersGet200Response
	fmt.Fprintf(os.Stdout, "Response from `RunnersAPI.ApiV1RunnersGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1RunnersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **int32** | Offset used to paginate results | 
 **count** | **int32** | Number of items to return | [default to 15]
 **sort** | **string** | Sort runners by criteria | 

### Return type

[**ApiV1RunnersGet200Response**](ApiV1RunnersGet200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1RunnersRegisterPost

> ApiV1RunnersRegisterPost200Response ApiV1RunnersRegisterPost(ctx).ApiV1RunnersRegisterPostRequest(apiV1RunnersRegisterPostRequest).Execute()

Register a new runner



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
	apiV1RunnersRegisterPostRequest := *openapiclient.NewApiV1RunnersRegisterPostRequest("RegistrationToken_example", "Name_example") // ApiV1RunnersRegisterPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RunnersAPI.ApiV1RunnersRegisterPost(context.Background()).ApiV1RunnersRegisterPostRequest(apiV1RunnersRegisterPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunnersAPI.ApiV1RunnersRegisterPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1RunnersRegisterPost`: ApiV1RunnersRegisterPost200Response
	fmt.Fprintf(os.Stdout, "Response from `RunnersAPI.ApiV1RunnersRegisterPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1RunnersRegisterPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiV1RunnersRegisterPostRequest** | [**ApiV1RunnersRegisterPostRequest**](ApiV1RunnersRegisterPostRequest.md) |  | 

### Return type

[**ApiV1RunnersRegisterPost200Response**](ApiV1RunnersRegisterPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1RunnersRunnerIdDelete

> ApiV1RunnersRunnerIdDelete(ctx, runnerId).ApiV1RunnersUnregisterPostRequest(apiV1RunnersUnregisterPostRequest).Execute()

Delete a runner

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
	runnerId := int32(56) // int32 | 
	apiV1RunnersUnregisterPostRequest := *openapiclient.NewApiV1RunnersUnregisterPostRequest("RunnerToken_example") // ApiV1RunnersUnregisterPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RunnersAPI.ApiV1RunnersRunnerIdDelete(context.Background(), runnerId).ApiV1RunnersUnregisterPostRequest(apiV1RunnersUnregisterPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunnersAPI.ApiV1RunnersRunnerIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**runnerId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1RunnersRunnerIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiV1RunnersUnregisterPostRequest** | [**ApiV1RunnersUnregisterPostRequest**](ApiV1RunnersUnregisterPostRequest.md) |  | 

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


## ApiV1RunnersUnregisterPost

> ApiV1RunnersUnregisterPost(ctx).ApiV1RunnersUnregisterPostRequest(apiV1RunnersUnregisterPostRequest).Execute()

Unregister a runner



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
	apiV1RunnersUnregisterPostRequest := *openapiclient.NewApiV1RunnersUnregisterPostRequest("RunnerToken_example") // ApiV1RunnersUnregisterPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RunnersAPI.ApiV1RunnersUnregisterPost(context.Background()).ApiV1RunnersUnregisterPostRequest(apiV1RunnersUnregisterPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunnersAPI.ApiV1RunnersUnregisterPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1RunnersUnregisterPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiV1RunnersUnregisterPostRequest** | [**ApiV1RunnersUnregisterPostRequest**](ApiV1RunnersUnregisterPostRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

