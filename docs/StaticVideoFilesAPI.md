# \StaticVideoFilesAPI

All URIs are relative to *https://peertube2.cpy.re*

Method | HTTP request | Description
------------- | ------------- | -------------
[**StaticStreamingPlaylistsHlsFilenameGet**](StaticVideoFilesAPI.md#StaticStreamingPlaylistsHlsFilenameGet) | **Get** /static/streaming-playlists/hls/{filename} | Get public HLS video file
[**StaticStreamingPlaylistsHlsPrivateFilenameGet**](StaticVideoFilesAPI.md#StaticStreamingPlaylistsHlsPrivateFilenameGet) | **Get** /static/streaming-playlists/hls/private/{filename} | Get private HLS video file
[**StaticWebVideosFilenameGet**](StaticVideoFilesAPI.md#StaticWebVideosFilenameGet) | **Get** /static/web-videos/{filename} | Get public Web Video file
[**StaticWebVideosPrivateFilenameGet**](StaticVideoFilesAPI.md#StaticWebVideosPrivateFilenameGet) | **Get** /static/web-videos/private/{filename} | Get private Web Video file



## StaticStreamingPlaylistsHlsFilenameGet

> StaticStreamingPlaylistsHlsFilenameGet(ctx, filename).Execute()

Get public HLS video file

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
	filename := "filename_example" // string | Filename

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StaticVideoFilesAPI.StaticStreamingPlaylistsHlsFilenameGet(context.Background(), filename).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StaticVideoFilesAPI.StaticStreamingPlaylistsHlsFilenameGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**filename** | **string** | Filename | 

### Other Parameters

Other parameters are passed through a pointer to a apiStaticStreamingPlaylistsHlsFilenameGetRequest struct via the builder pattern


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


## StaticStreamingPlaylistsHlsPrivateFilenameGet

> StaticStreamingPlaylistsHlsPrivateFilenameGet(ctx, filename).VideoFileToken(videoFileToken).ReinjectVideoFileToken(reinjectVideoFileToken).Execute()

Get private HLS video file

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
	filename := "filename_example" // string | Filename
	videoFileToken := "videoFileToken_example" // string | Video file token [generated](#operation/requestVideoToken) by PeerTube so you don't need to provide an OAuth token in the request header. (optional)
	reinjectVideoFileToken := true // bool | Ask the server to reinject videoFileToken in URLs in m3u8 playlist (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StaticVideoFilesAPI.StaticStreamingPlaylistsHlsPrivateFilenameGet(context.Background(), filename).VideoFileToken(videoFileToken).ReinjectVideoFileToken(reinjectVideoFileToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StaticVideoFilesAPI.StaticStreamingPlaylistsHlsPrivateFilenameGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**filename** | **string** | Filename | 

### Other Parameters

Other parameters are passed through a pointer to a apiStaticStreamingPlaylistsHlsPrivateFilenameGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **videoFileToken** | **string** | Video file token [generated](#operation/requestVideoToken) by PeerTube so you don&#39;t need to provide an OAuth token in the request header. | 
 **reinjectVideoFileToken** | **bool** | Ask the server to reinject videoFileToken in URLs in m3u8 playlist | 

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


## StaticWebVideosFilenameGet

> StaticWebVideosFilenameGet(ctx, filename).Execute()

Get public Web Video file



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
	filename := "filename_example" // string | Filename

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StaticVideoFilesAPI.StaticWebVideosFilenameGet(context.Background(), filename).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StaticVideoFilesAPI.StaticWebVideosFilenameGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**filename** | **string** | Filename | 

### Other Parameters

Other parameters are passed through a pointer to a apiStaticWebVideosFilenameGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## StaticWebVideosPrivateFilenameGet

> StaticWebVideosPrivateFilenameGet(ctx, filename).VideoFileToken(videoFileToken).Execute()

Get private Web Video file



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
	filename := "filename_example" // string | Filename
	videoFileToken := "videoFileToken_example" // string | Video file token [generated](#operation/requestVideoToken) by PeerTube so you don't need to provide an OAuth token in the request header. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StaticVideoFilesAPI.StaticWebVideosPrivateFilenameGet(context.Background(), filename).VideoFileToken(videoFileToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StaticVideoFilesAPI.StaticWebVideosPrivateFilenameGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**filename** | **string** | Filename | 

### Other Parameters

Other parameters are passed through a pointer to a apiStaticWebVideosPrivateFilenameGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **videoFileToken** | **string** | Video file token [generated](#operation/requestVideoToken) by PeerTube so you don&#39;t need to provide an OAuth token in the request header. | 

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

