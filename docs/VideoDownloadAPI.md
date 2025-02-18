# \VideoDownloadAPI

All URIs are relative to *https://peertube2.cpy.re*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DownloadVideosGenerateVideoIdGet**](VideoDownloadAPI.md#DownloadVideosGenerateVideoIdGet) | **Get** /download/videos/generate/:videoId | Download video file



## DownloadVideosGenerateVideoIdGet

> DownloadVideosGenerateVideoIdGet(ctx, videoId).VideoFileIds(videoFileIds).VideoFileToken(videoFileToken).Execute()

Download video file



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
	videoId := int32(56) // int32 | The video id
	videoFileIds := []int32{int32(123)} // []int32 | streams of video files to mux in the output
	videoFileToken := "videoFileToken_example" // string | Video file token [generated](#operation/requestVideoToken) by PeerTube so you don't need to provide an OAuth token in the request header. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VideoDownloadAPI.DownloadVideosGenerateVideoIdGet(context.Background(), videoId).VideoFileIds(videoFileIds).VideoFileToken(videoFileToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideoDownloadAPI.DownloadVideosGenerateVideoIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**videoId** | **int32** | The video id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadVideosGenerateVideoIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **videoFileIds** | **[]int32** | streams of video files to mux in the output | 
 **videoFileToken** | **string** | Video file token [generated](#operation/requestVideoToken) by PeerTube so you don&#39;t need to provide an OAuth token in the request header. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

