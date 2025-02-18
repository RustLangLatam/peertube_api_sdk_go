# \VideoCaptionsAPI

All URIs are relative to *https://peertube2.cpy.re*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddVideoCaption**](VideoCaptionsAPI.md#AddVideoCaption) | **Put** /api/v1/videos/{id}/captions/{captionLanguage} | Add or replace a video caption
[**DelVideoCaption**](VideoCaptionsAPI.md#DelVideoCaption) | **Delete** /api/v1/videos/{id}/captions/{captionLanguage} | Delete a video caption
[**GenerateVideoCaption**](VideoCaptionsAPI.md#GenerateVideoCaption) | **Post** /api/v1/videos/{id}/captions/generate | Generate a video caption
[**GetVideoCaptions**](VideoCaptionsAPI.md#GetVideoCaptions) | **Get** /api/v1/videos/{id}/captions | List captions of a video



## AddVideoCaption

> AddVideoCaption(ctx, id, captionLanguage).Captionfile(captionfile).Execute()

Add or replace a video caption

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
	captionLanguage := "captionLanguage_example" // string | The caption language
	captionfile := os.NewFile(1234, "some_file") // *os.File | The file to upload. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VideoCaptionsAPI.AddVideoCaption(context.Background(), id, captionLanguage).Captionfile(captionfile).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideoCaptionsAPI.AddVideoCaption``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**ApiV1VideosOwnershipIdAcceptPostIdParameter**](.md) | The object id, uuid or short uuid | 
**captionLanguage** | **string** | The caption language | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddVideoCaptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **captionfile** | ***os.File** | The file to upload. | 

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


## DelVideoCaption

> DelVideoCaption(ctx, id, captionLanguage).Execute()

Delete a video caption

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
	captionLanguage := "captionLanguage_example" // string | The caption language

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VideoCaptionsAPI.DelVideoCaption(context.Background(), id, captionLanguage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideoCaptionsAPI.DelVideoCaption``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**ApiV1VideosOwnershipIdAcceptPostIdParameter**](.md) | The object id, uuid or short uuid | 
**captionLanguage** | **string** | The caption language | 

### Other Parameters

Other parameters are passed through a pointer to a apiDelVideoCaptionRequest struct via the builder pattern


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


## GenerateVideoCaption

> GenerateVideoCaption(ctx, id).GenerateVideoCaptionRequest(generateVideoCaptionRequest).Execute()

Generate a video caption



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
	generateVideoCaptionRequest := *openapiclient.NewGenerateVideoCaptionRequest() // GenerateVideoCaptionRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VideoCaptionsAPI.GenerateVideoCaption(context.Background(), id).GenerateVideoCaptionRequest(generateVideoCaptionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideoCaptionsAPI.GenerateVideoCaption``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGenerateVideoCaptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **generateVideoCaptionRequest** | [**GenerateVideoCaptionRequest**](GenerateVideoCaptionRequest.md) |  | 

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


## GetVideoCaptions

> GetVideoCaptions200Response GetVideoCaptions(ctx, id).XPeertubeVideoPassword(xPeertubeVideoPassword).Execute()

List captions of a video

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
	resp, r, err := apiClient.VideoCaptionsAPI.GetVideoCaptions(context.Background(), id).XPeertubeVideoPassword(xPeertubeVideoPassword).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideoCaptionsAPI.GetVideoCaptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVideoCaptions`: GetVideoCaptions200Response
	fmt.Fprintf(os.Stdout, "Response from `VideoCaptionsAPI.GetVideoCaptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**ApiV1VideosOwnershipIdAcceptPostIdParameter**](.md) | The object id, uuid or short uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVideoCaptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xPeertubeVideoPassword** | **string** | Required on password protected video | 

### Return type

[**GetVideoCaptions200Response**](GetVideoCaptions200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

