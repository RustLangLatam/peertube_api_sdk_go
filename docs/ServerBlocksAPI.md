# \ServerBlocksAPI

All URIs are relative to *https://peertube2.cpy.re*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV1ServerBlocklistServersGet**](ServerBlocksAPI.md#ApiV1ServerBlocklistServersGet) | **Get** /api/v1/server/blocklist/servers | List server blocks
[**ApiV1ServerBlocklistServersHostDelete**](ServerBlocksAPI.md#ApiV1ServerBlocklistServersHostDelete) | **Delete** /api/v1/server/blocklist/servers/{host} | Unblock a server by its domain
[**ApiV1ServerBlocklistServersPost**](ServerBlocksAPI.md#ApiV1ServerBlocklistServersPost) | **Post** /api/v1/server/blocklist/servers | Block a server



## ApiV1ServerBlocklistServersGet

> ApiV1ServerBlocklistServersGet(ctx).Start(start).Count(count).Sort(sort).Execute()

List server blocks

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
	sort := "-createdAt" // string | Sort column (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ServerBlocksAPI.ApiV1ServerBlocklistServersGet(context.Background()).Start(start).Count(count).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerBlocksAPI.ApiV1ServerBlocklistServersGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1ServerBlocklistServersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **int32** | Offset used to paginate results | 
 **count** | **int32** | Number of items to return | [default to 15]
 **sort** | **string** | Sort column | 

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


## ApiV1ServerBlocklistServersHostDelete

> ApiV1ServerBlocklistServersHostDelete(ctx, host).Execute()

Unblock a server by its domain

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
	host := "host_example" // string | server domain to unblock

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ServerBlocksAPI.ApiV1ServerBlocklistServersHostDelete(context.Background(), host).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerBlocksAPI.ApiV1ServerBlocklistServersHostDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**host** | **string** | server domain to unblock | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1ServerBlocklistServersHostDeleteRequest struct via the builder pattern


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


## ApiV1ServerBlocklistServersPost

> ApiV1ServerBlocklistServersPost(ctx).ApiV1ServerBlocklistServersPostRequest(apiV1ServerBlocklistServersPostRequest).Execute()

Block a server

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
	apiV1ServerBlocklistServersPostRequest := *openapiclient.NewApiV1ServerBlocklistServersPostRequest("Host_example") // ApiV1ServerBlocklistServersPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ServerBlocksAPI.ApiV1ServerBlocklistServersPost(context.Background()).ApiV1ServerBlocklistServersPostRequest(apiV1ServerBlocklistServersPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServerBlocksAPI.ApiV1ServerBlocklistServersPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1ServerBlocklistServersPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiV1ServerBlocklistServersPostRequest** | [**ApiV1ServerBlocklistServersPostRequest**](ApiV1ServerBlocklistServersPostRequest.md) |  | 

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

