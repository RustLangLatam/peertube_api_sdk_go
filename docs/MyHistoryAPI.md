# \MyHistoryAPI

All URIs are relative to *https://peertube2.cpy.re*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV1UsersMeHistoryVideosGet**](MyHistoryAPI.md#ApiV1UsersMeHistoryVideosGet) | **Get** /api/v1/users/me/history/videos | List watched videos history
[**ApiV1UsersMeHistoryVideosRemovePost**](MyHistoryAPI.md#ApiV1UsersMeHistoryVideosRemovePost) | **Post** /api/v1/users/me/history/videos/remove | Clear video history
[**ApiV1UsersMeHistoryVideosVideoIdDelete**](MyHistoryAPI.md#ApiV1UsersMeHistoryVideosVideoIdDelete) | **Delete** /api/v1/users/me/history/videos/{videoId} | Delete history element



## ApiV1UsersMeHistoryVideosGet

> VideoListResponse ApiV1UsersMeHistoryVideosGet(ctx).Start(start).Count(count).Search(search).Execute()

List watched videos history

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
	search := "search_example" // string | Plain text search, applied to various parts of the model depending on endpoint (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MyHistoryAPI.ApiV1UsersMeHistoryVideosGet(context.Background()).Start(start).Count(count).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MyHistoryAPI.ApiV1UsersMeHistoryVideosGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1UsersMeHistoryVideosGet`: VideoListResponse
	fmt.Fprintf(os.Stdout, "Response from `MyHistoryAPI.ApiV1UsersMeHistoryVideosGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1UsersMeHistoryVideosGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **int32** | Offset used to paginate results | 
 **count** | **int32** | Number of items to return | [default to 15]
 **search** | **string** | Plain text search, applied to various parts of the model depending on endpoint | 

### Return type

[**VideoListResponse**](VideoListResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1UsersMeHistoryVideosRemovePost

> ApiV1UsersMeHistoryVideosRemovePost(ctx).BeforeDate(beforeDate).Execute()

Clear video history

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/peertube_api_sdk_go"
)

func main() {
	beforeDate := time.Now() // time.Time | history before this date will be deleted (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MyHistoryAPI.ApiV1UsersMeHistoryVideosRemovePost(context.Background()).BeforeDate(beforeDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MyHistoryAPI.ApiV1UsersMeHistoryVideosRemovePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1UsersMeHistoryVideosRemovePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **beforeDate** | **time.Time** | history before this date will be deleted | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1UsersMeHistoryVideosVideoIdDelete

> ApiV1UsersMeHistoryVideosVideoIdDelete(ctx, videoId).Execute()

Delete history element

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
	videoId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MyHistoryAPI.ApiV1UsersMeHistoryVideosVideoIdDelete(context.Background(), videoId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MyHistoryAPI.ApiV1UsersMeHistoryVideosVideoIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**videoId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1UsersMeHistoryVideosVideoIdDeleteRequest struct via the builder pattern


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

