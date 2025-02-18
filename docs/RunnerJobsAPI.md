# \RunnerJobsAPI

All URIs are relative to *https://peertube2.cpy.re*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV1RunnersJobsGet**](RunnerJobsAPI.md#ApiV1RunnersJobsGet) | **Get** /api/v1/runners/jobs | List jobs
[**ApiV1RunnersJobsJobUUIDAbortPost**](RunnerJobsAPI.md#ApiV1RunnersJobsJobUUIDAbortPost) | **Post** /api/v1/runners/jobs/{jobUUID}/abort | Abort job
[**ApiV1RunnersJobsJobUUIDAcceptPost**](RunnerJobsAPI.md#ApiV1RunnersJobsJobUUIDAcceptPost) | **Post** /api/v1/runners/jobs/{jobUUID}/accept | Accept job
[**ApiV1RunnersJobsJobUUIDCancelGet**](RunnerJobsAPI.md#ApiV1RunnersJobsJobUUIDCancelGet) | **Get** /api/v1/runners/jobs/{jobUUID}/cancel | Cancel a job
[**ApiV1RunnersJobsJobUUIDDelete**](RunnerJobsAPI.md#ApiV1RunnersJobsJobUUIDDelete) | **Delete** /api/v1/runners/jobs/{jobUUID} | Delete a job
[**ApiV1RunnersJobsJobUUIDErrorPost**](RunnerJobsAPI.md#ApiV1RunnersJobsJobUUIDErrorPost) | **Post** /api/v1/runners/jobs/{jobUUID}/error | Post job error
[**ApiV1RunnersJobsJobUUIDSuccessPost**](RunnerJobsAPI.md#ApiV1RunnersJobsJobUUIDSuccessPost) | **Post** /api/v1/runners/jobs/{jobUUID}/success | Post job success
[**ApiV1RunnersJobsJobUUIDUpdatePost**](RunnerJobsAPI.md#ApiV1RunnersJobsJobUUIDUpdatePost) | **Post** /api/v1/runners/jobs/{jobUUID}/update | Update job
[**ApiV1RunnersJobsRequestPost**](RunnerJobsAPI.md#ApiV1RunnersJobsRequestPost) | **Post** /api/v1/runners/jobs/request | Request a new job



## ApiV1RunnersJobsGet

> ApiV1RunnersJobsGet200Response ApiV1RunnersJobsGet(ctx).Start(start).Count(count).Sort(sort).Search(search).StateOneOf(stateOneOf).Execute()

List jobs

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/peertube_api_sdk_go"
)

func main() {
	start := int32(56) // int32 | Offset used to paginate results (optional)
	count := int32(56) // int32 | Number of items to return (optional) (default to 15)
	sort := "sort_example" // string | Sort runner jobs by criteria (optional)
	search := "search_example" // string | Plain text search, applied to various parts of the model depending on endpoint (optional)
	stateOneOf := []openapiclient.RunnerJobState{openapiclient.RunnerJobState(1)} // []RunnerJobState |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RunnerJobsAPI.ApiV1RunnersJobsGet(context.Background()).Start(start).Count(count).Sort(sort).Search(search).StateOneOf(stateOneOf).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunnerJobsAPI.ApiV1RunnersJobsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1RunnersJobsGet`: ApiV1RunnersJobsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `RunnerJobsAPI.ApiV1RunnersJobsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1RunnersJobsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **int32** | Offset used to paginate results | 
 **count** | **int32** | Number of items to return | [default to 15]
 **sort** | **string** | Sort runner jobs by criteria | 
 **search** | **string** | Plain text search, applied to various parts of the model depending on endpoint | 
 **stateOneOf** | [**[]RunnerJobState**](RunnerJobState.md) |  | 

### Return type

[**ApiV1RunnersJobsGet200Response**](ApiV1RunnersJobsGet200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1RunnersJobsJobUUIDAbortPost

> ApiV1RunnersJobsJobUUIDAbortPost(ctx, jobUUID).ApiV1RunnersJobsJobUUIDAbortPostRequest(apiV1RunnersJobsJobUUIDAbortPostRequest).Execute()

Abort job



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/peertube_api_sdk_go"
)

func main() {
	jobUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiV1RunnersJobsJobUUIDAbortPostRequest := *openapiclient.NewApiV1RunnersJobsJobUUIDAbortPostRequest("RunnerToken_example", "JobToken_example", "Reason_example") // ApiV1RunnersJobsJobUUIDAbortPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RunnerJobsAPI.ApiV1RunnersJobsJobUUIDAbortPost(context.Background(), jobUUID).ApiV1RunnersJobsJobUUIDAbortPostRequest(apiV1RunnersJobsJobUUIDAbortPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunnerJobsAPI.ApiV1RunnersJobsJobUUIDAbortPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobUUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1RunnersJobsJobUUIDAbortPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiV1RunnersJobsJobUUIDAbortPostRequest** | [**ApiV1RunnersJobsJobUUIDAbortPostRequest**](ApiV1RunnersJobsJobUUIDAbortPostRequest.md) |  | 

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


## ApiV1RunnersJobsJobUUIDAcceptPost

> ApiV1RunnersJobsJobUUIDAcceptPost200Response ApiV1RunnersJobsJobUUIDAcceptPost(ctx, jobUUID).ApiV1RunnersUnregisterPostRequest(apiV1RunnersUnregisterPostRequest).Execute()

Accept job



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/peertube_api_sdk_go"
)

func main() {
	jobUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiV1RunnersUnregisterPostRequest := *openapiclient.NewApiV1RunnersUnregisterPostRequest("RunnerToken_example") // ApiV1RunnersUnregisterPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RunnerJobsAPI.ApiV1RunnersJobsJobUUIDAcceptPost(context.Background(), jobUUID).ApiV1RunnersUnregisterPostRequest(apiV1RunnersUnregisterPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunnerJobsAPI.ApiV1RunnersJobsJobUUIDAcceptPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1RunnersJobsJobUUIDAcceptPost`: ApiV1RunnersJobsJobUUIDAcceptPost200Response
	fmt.Fprintf(os.Stdout, "Response from `RunnerJobsAPI.ApiV1RunnersJobsJobUUIDAcceptPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobUUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1RunnersJobsJobUUIDAcceptPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiV1RunnersUnregisterPostRequest** | [**ApiV1RunnersUnregisterPostRequest**](ApiV1RunnersUnregisterPostRequest.md) |  | 

### Return type

[**ApiV1RunnersJobsJobUUIDAcceptPost200Response**](ApiV1RunnersJobsJobUUIDAcceptPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1RunnersJobsJobUUIDCancelGet

> ApiV1RunnersJobsJobUUIDCancelGet(ctx, jobUUID).Execute()

Cancel a job

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/peertube_api_sdk_go"
)

func main() {
	jobUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RunnerJobsAPI.ApiV1RunnersJobsJobUUIDCancelGet(context.Background(), jobUUID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunnerJobsAPI.ApiV1RunnersJobsJobUUIDCancelGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobUUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1RunnersJobsJobUUIDCancelGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## ApiV1RunnersJobsJobUUIDDelete

> ApiV1RunnersJobsJobUUIDDelete(ctx, jobUUID).Execute()

Delete a job



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/peertube_api_sdk_go"
)

func main() {
	jobUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RunnerJobsAPI.ApiV1RunnersJobsJobUUIDDelete(context.Background(), jobUUID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunnerJobsAPI.ApiV1RunnersJobsJobUUIDDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobUUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1RunnersJobsJobUUIDDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## ApiV1RunnersJobsJobUUIDErrorPost

> ApiV1RunnersJobsJobUUIDErrorPost(ctx, jobUUID).ApiV1RunnersJobsJobUUIDErrorPostRequest(apiV1RunnersJobsJobUUIDErrorPostRequest).Execute()

Post job error



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/peertube_api_sdk_go"
)

func main() {
	jobUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiV1RunnersJobsJobUUIDErrorPostRequest := *openapiclient.NewApiV1RunnersJobsJobUUIDErrorPostRequest("RunnerToken_example", "JobToken_example", "Message_example") // ApiV1RunnersJobsJobUUIDErrorPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RunnerJobsAPI.ApiV1RunnersJobsJobUUIDErrorPost(context.Background(), jobUUID).ApiV1RunnersJobsJobUUIDErrorPostRequest(apiV1RunnersJobsJobUUIDErrorPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunnerJobsAPI.ApiV1RunnersJobsJobUUIDErrorPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobUUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1RunnersJobsJobUUIDErrorPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiV1RunnersJobsJobUUIDErrorPostRequest** | [**ApiV1RunnersJobsJobUUIDErrorPostRequest**](ApiV1RunnersJobsJobUUIDErrorPostRequest.md) |  | 

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


## ApiV1RunnersJobsJobUUIDSuccessPost

> ApiV1RunnersJobsJobUUIDSuccessPost(ctx, jobUUID).ApiV1RunnersJobsJobUUIDSuccessPostRequest(apiV1RunnersJobsJobUUIDSuccessPostRequest).Execute()

Post job success



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/peertube_api_sdk_go"
)

func main() {
	jobUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiV1RunnersJobsJobUUIDSuccessPostRequest := *openapiclient.NewApiV1RunnersJobsJobUUIDSuccessPostRequest("RunnerToken_example", "JobToken_example", *openapiclient.NewApiV1RunnersJobsJobUUIDSuccessPostRequestPayload()) // ApiV1RunnersJobsJobUUIDSuccessPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RunnerJobsAPI.ApiV1RunnersJobsJobUUIDSuccessPost(context.Background(), jobUUID).ApiV1RunnersJobsJobUUIDSuccessPostRequest(apiV1RunnersJobsJobUUIDSuccessPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunnerJobsAPI.ApiV1RunnersJobsJobUUIDSuccessPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobUUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1RunnersJobsJobUUIDSuccessPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiV1RunnersJobsJobUUIDSuccessPostRequest** | [**ApiV1RunnersJobsJobUUIDSuccessPostRequest**](ApiV1RunnersJobsJobUUIDSuccessPostRequest.md) |  | 

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


## ApiV1RunnersJobsJobUUIDUpdatePost

> ApiV1RunnersJobsJobUUIDUpdatePost(ctx, jobUUID).ApiV1RunnersJobsJobUUIDUpdatePostRequest(apiV1RunnersJobsJobUUIDUpdatePostRequest).Execute()

Update job



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/peertube_api_sdk_go"
)

func main() {
	jobUUID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	apiV1RunnersJobsJobUUIDUpdatePostRequest := *openapiclient.NewApiV1RunnersJobsJobUUIDUpdatePostRequest("RunnerToken_example", "JobToken_example") // ApiV1RunnersJobsJobUUIDUpdatePostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RunnerJobsAPI.ApiV1RunnersJobsJobUUIDUpdatePost(context.Background(), jobUUID).ApiV1RunnersJobsJobUUIDUpdatePostRequest(apiV1RunnersJobsJobUUIDUpdatePostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunnerJobsAPI.ApiV1RunnersJobsJobUUIDUpdatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobUUID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1RunnersJobsJobUUIDUpdatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiV1RunnersJobsJobUUIDUpdatePostRequest** | [**ApiV1RunnersJobsJobUUIDUpdatePostRequest**](ApiV1RunnersJobsJobUUIDUpdatePostRequest.md) |  | 

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


## ApiV1RunnersJobsRequestPost

> ApiV1RunnersJobsRequestPost200Response ApiV1RunnersJobsRequestPost(ctx).ApiV1RunnersJobsRequestPostRequest(apiV1RunnersJobsRequestPostRequest).Execute()

Request a new job



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/peertube_api_sdk_go"
)

func main() {
	apiV1RunnersJobsRequestPostRequest := *openapiclient.NewApiV1RunnersJobsRequestPostRequest("RunnerToken_example") // ApiV1RunnersJobsRequestPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RunnerJobsAPI.ApiV1RunnersJobsRequestPost(context.Background()).ApiV1RunnersJobsRequestPostRequest(apiV1RunnersJobsRequestPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunnerJobsAPI.ApiV1RunnersJobsRequestPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1RunnersJobsRequestPost`: ApiV1RunnersJobsRequestPost200Response
	fmt.Fprintf(os.Stdout, "Response from `RunnerJobsAPI.ApiV1RunnersJobsRequestPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1RunnersJobsRequestPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiV1RunnersJobsRequestPostRequest** | [**ApiV1RunnersJobsRequestPostRequest**](ApiV1RunnersJobsRequestPostRequest.md) |  | 

### Return type

[**ApiV1RunnersJobsRequestPost200Response**](ApiV1RunnersJobsRequestPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

