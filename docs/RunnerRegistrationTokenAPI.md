# \RunnerRegistrationTokenAPI

All URIs are relative to *https://peertube2.cpy.re*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV1RunnersRegistrationTokensGeneratePost**](RunnerRegistrationTokenAPI.md#ApiV1RunnersRegistrationTokensGeneratePost) | **Post** /api/v1/runners/registration-tokens/generate | Generate registration token
[**ApiV1RunnersRegistrationTokensGet**](RunnerRegistrationTokenAPI.md#ApiV1RunnersRegistrationTokensGet) | **Get** /api/v1/runners/registration-tokens | List registration tokens
[**ApiV1RunnersRegistrationTokensRegistrationTokenIdDelete**](RunnerRegistrationTokenAPI.md#ApiV1RunnersRegistrationTokensRegistrationTokenIdDelete) | **Delete** /api/v1/runners/registration-tokens/{registrationTokenId} | Remove registration token



## ApiV1RunnersRegistrationTokensGeneratePost

> ApiV1RunnersRegistrationTokensGeneratePost(ctx).Execute()

Generate registration token



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RunnerRegistrationTokenAPI.ApiV1RunnersRegistrationTokensGeneratePost(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunnerRegistrationTokenAPI.ApiV1RunnersRegistrationTokensGeneratePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1RunnersRegistrationTokensGeneratePostRequest struct via the builder pattern


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


## ApiV1RunnersRegistrationTokensGet

> ApiV1RunnersRegistrationTokensGet200Response ApiV1RunnersRegistrationTokensGet(ctx).Start(start).Count(count).Sort(sort).Execute()

List registration tokens

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
	sort := "sort_example" // string | Sort registration tokens by criteria (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RunnerRegistrationTokenAPI.ApiV1RunnersRegistrationTokensGet(context.Background()).Start(start).Count(count).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunnerRegistrationTokenAPI.ApiV1RunnersRegistrationTokensGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1RunnersRegistrationTokensGet`: ApiV1RunnersRegistrationTokensGet200Response
	fmt.Fprintf(os.Stdout, "Response from `RunnerRegistrationTokenAPI.ApiV1RunnersRegistrationTokensGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1RunnersRegistrationTokensGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **int32** | Offset used to paginate results | 
 **count** | **int32** | Number of items to return | [default to 15]
 **sort** | **string** | Sort registration tokens by criteria | 

### Return type

[**ApiV1RunnersRegistrationTokensGet200Response**](ApiV1RunnersRegistrationTokensGet200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1RunnersRegistrationTokensRegistrationTokenIdDelete

> ApiV1RunnersRegistrationTokensRegistrationTokenIdDelete(ctx, registrationTokenId).Execute()

Remove registration token



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
	registrationTokenId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RunnerRegistrationTokenAPI.ApiV1RunnersRegistrationTokensRegistrationTokenIdDelete(context.Background(), registrationTokenId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunnerRegistrationTokenAPI.ApiV1RunnersRegistrationTokensRegistrationTokenIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registrationTokenId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1RunnersRegistrationTokensRegistrationTokenIdDeleteRequest struct via the builder pattern


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

