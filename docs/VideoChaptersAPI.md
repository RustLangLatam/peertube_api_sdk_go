# \VideoChaptersAPI

All URIs are relative to *https://peertube2.cpy.re*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetVideoChapters**](VideoChaptersAPI.md#GetVideoChapters) | **Get** /api/v1/videos/{id}/chapters | Get chapters of a video
[**ReplaceVideoChapters**](VideoChaptersAPI.md#ReplaceVideoChapters) | **Put** /api/v1/videos/{id}/chapters | Replace video chapters



## GetVideoChapters

> VideoChapters GetVideoChapters(ctx, id).XPeertubeVideoPassword(xPeertubeVideoPassword).Execute()

Get chapters of a video



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
	id := openapiclient._api_v1_videos_ownership__id__accept_post_id_parameter{Int32: new(int32)} // ApiV1VideosOwnershipIdAcceptPostIdParameter | The object id, uuid or short uuid
	xPeertubeVideoPassword := "xPeertubeVideoPassword_example" // string | Required on password protected video (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VideoChaptersAPI.GetVideoChapters(context.Background(), id).XPeertubeVideoPassword(xPeertubeVideoPassword).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideoChaptersAPI.GetVideoChapters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVideoChapters`: VideoChapters
	fmt.Fprintf(os.Stdout, "Response from `VideoChaptersAPI.GetVideoChapters`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**ApiV1VideosOwnershipIdAcceptPostIdParameter**](.md) | The object id, uuid or short uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVideoChaptersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xPeertubeVideoPassword** | **string** | Required on password protected video | 

### Return type

[**VideoChapters**](VideoChapters.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceVideoChapters

> ReplaceVideoChapters(ctx, id).ReplaceVideoChaptersRequest(replaceVideoChaptersRequest).Execute()

Replace video chapters



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
	id := openapiclient._api_v1_videos_ownership__id__accept_post_id_parameter{Int32: new(int32)} // ApiV1VideosOwnershipIdAcceptPostIdParameter | The object id, uuid or short uuid
	replaceVideoChaptersRequest := *openapiclient.NewReplaceVideoChaptersRequest() // ReplaceVideoChaptersRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VideoChaptersAPI.ReplaceVideoChapters(context.Background(), id).ReplaceVideoChaptersRequest(replaceVideoChaptersRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideoChaptersAPI.ReplaceVideoChapters``: %v\n", err)
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

Other parameters are passed through a pointer to a apiReplaceVideoChaptersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **replaceVideoChaptersRequest** | [**ReplaceVideoChaptersRequest**](ReplaceVideoChaptersRequest.md) |  | 

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

