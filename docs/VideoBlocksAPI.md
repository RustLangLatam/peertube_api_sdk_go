# \VideoBlocksAPI

All URIs are relative to *https://peertube2.cpy.re*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddVideoBlock**](VideoBlocksAPI.md#AddVideoBlock) | **Post** /api/v1/videos/{id}/blacklist | Block a video
[**DelVideoBlock**](VideoBlocksAPI.md#DelVideoBlock) | **Delete** /api/v1/videos/{id}/blacklist | Unblock a video by its id
[**GetVideoBlocks**](VideoBlocksAPI.md#GetVideoBlocks) | **Get** /api/v1/videos/blacklist | List video blocks



## AddVideoBlock

> AddVideoBlock(ctx, id).Execute()

Block a video

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
	id := openapiclient._api_v1_videos_ownership__id__accept_post_id_parameter{Int32: new(int32)} // ApiV1VideosOwnershipIdAcceptPostIdParameter | The object id, uuid or short uuid

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VideoBlocksAPI.AddVideoBlock(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideoBlocksAPI.AddVideoBlock``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**ApiV1VideosOwnershipIdAcceptPostIdParameter**](.md) | The object id, uuid or short uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddVideoBlockRequest struct via the builder pattern


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


## DelVideoBlock

> DelVideoBlock(ctx, id).Execute()

Unblock a video by its id

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
	id := openapiclient._api_v1_videos_ownership__id__accept_post_id_parameter{Int32: new(int32)} // ApiV1VideosOwnershipIdAcceptPostIdParameter | The object id, uuid or short uuid

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VideoBlocksAPI.DelVideoBlock(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideoBlocksAPI.DelVideoBlock``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**ApiV1VideosOwnershipIdAcceptPostIdParameter**](.md) | The object id, uuid or short uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiDelVideoBlockRequest struct via the builder pattern


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


## GetVideoBlocks

> GetVideoBlocks200Response GetVideoBlocks(ctx).Type_(type_).Search(search).Start(start).Count(count).Sort(sort).Execute()

List video blocks

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
	type_ := int32(56) // int32 | list only blocks that match this type: - `1`: manual block - `2`: automatic block that needs review  (optional)
	search := "search_example" // string | plain search that will match with video titles, and more (optional)
	start := int32(56) // int32 | Offset used to paginate results (optional)
	count := int32(56) // int32 | Number of items to return (optional) (default to 15)
	sort := "sort_example" // string | Sort blocklists by criteria (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VideoBlocksAPI.GetVideoBlocks(context.Background()).Type_(type_).Search(search).Start(start).Count(count).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideoBlocksAPI.GetVideoBlocks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVideoBlocks`: GetVideoBlocks200Response
	fmt.Fprintf(os.Stdout, "Response from `VideoBlocksAPI.GetVideoBlocks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetVideoBlocksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **int32** | list only blocks that match this type: - &#x60;1&#x60;: manual block - &#x60;2&#x60;: automatic block that needs review  | 
 **search** | **string** | plain search that will match with video titles, and more | 
 **start** | **int32** | Offset used to paginate results | 
 **count** | **int32** | Number of items to return | [default to 15]
 **sort** | **string** | Sort blocklists by criteria | 

### Return type

[**GetVideoBlocks200Response**](GetVideoBlocks200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

