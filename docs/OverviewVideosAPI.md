# \OverviewVideosAPI

All URIs are relative to *https://peertube2.cpy.re*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOverviewVideos**](OverviewVideosAPI.md#GetOverviewVideos) | **Get** /api/v1/overviews/videos | Get overview of videos



## GetOverviewVideos

> OverviewVideosResponse GetOverviewVideos(ctx).Page(page).Execute()

Get overview of videos

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
	page := int32(56) // int32 | Offset used to paginate results (optional) (default to 1)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OverviewVideosAPI.GetOverviewVideos(context.Background()).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OverviewVideosAPI.GetOverviewVideos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOverviewVideos`: OverviewVideosResponse
	fmt.Fprintf(os.Stdout, "Response from `OverviewVideosAPI.GetOverviewVideos`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOverviewVideosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Offset used to paginate results | [default to 1]

### Return type

[**OverviewVideosResponse**](OverviewVideosResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

