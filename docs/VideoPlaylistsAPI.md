# \VideoPlaylistsAPI

All URIs are relative to *https://peertube2.cpy.re*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddPlaylist**](VideoPlaylistsAPI.md#AddPlaylist) | **Post** /api/v1/video-playlists | Create a video playlist
[**AddVideoPlaylistVideo**](VideoPlaylistsAPI.md#AddVideoPlaylistVideo) | **Post** /api/v1/video-playlists/{playlistId}/videos | Add a video in a playlist
[**ApiV1UsersMeVideoPlaylistsVideosExistGet**](VideoPlaylistsAPI.md#ApiV1UsersMeVideoPlaylistsVideosExistGet) | **Get** /api/v1/users/me/video-playlists/videos-exist | Check video exists in my playlists
[**ApiV1VideoPlaylistsPlaylistIdDelete**](VideoPlaylistsAPI.md#ApiV1VideoPlaylistsPlaylistIdDelete) | **Delete** /api/v1/video-playlists/{playlistId} | Delete a video playlist
[**ApiV1VideoPlaylistsPlaylistIdGet**](VideoPlaylistsAPI.md#ApiV1VideoPlaylistsPlaylistIdGet) | **Get** /api/v1/video-playlists/{playlistId} | Get a video playlist
[**ApiV1VideoPlaylistsPlaylistIdPut**](VideoPlaylistsAPI.md#ApiV1VideoPlaylistsPlaylistIdPut) | **Put** /api/v1/video-playlists/{playlistId} | Update a video playlist
[**DelVideoPlaylistVideo**](VideoPlaylistsAPI.md#DelVideoPlaylistVideo) | **Delete** /api/v1/video-playlists/{playlistId}/videos/{playlistElementId} | Delete an element from a playlist
[**GetPlaylistPrivacyPolicies**](VideoPlaylistsAPI.md#GetPlaylistPrivacyPolicies) | **Get** /api/v1/video-playlists/privacies | List available playlist privacy policies
[**GetPlaylists**](VideoPlaylistsAPI.md#GetPlaylists) | **Get** /api/v1/video-playlists | List video playlists
[**GetVideoPlaylistVideos**](VideoPlaylistsAPI.md#GetVideoPlaylistVideos) | **Get** /api/v1/video-playlists/{playlistId}/videos | List videos of a playlist
[**PutVideoPlaylistVideo**](VideoPlaylistsAPI.md#PutVideoPlaylistVideo) | **Put** /api/v1/video-playlists/{playlistId}/videos/{playlistElementId} | Update a playlist element
[**ReorderVideoPlaylist**](VideoPlaylistsAPI.md#ReorderVideoPlaylist) | **Post** /api/v1/video-playlists/{playlistId}/videos/reorder | Reorder a playlist



## AddPlaylist

> AddPlaylist200Response AddPlaylist(ctx).DisplayName(displayName).Thumbnailfile(thumbnailfile).Privacy(privacy).Description(description).VideoChannelId(videoChannelId).Execute()

Create a video playlist



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
	displayName := "displayName_example" // string | Video playlist display name
	thumbnailfile := os.NewFile(1234, "some_file") // *os.File | Video playlist thumbnail file (optional)
	privacy := openapiclient.VideoPlaylistPrivacySet(1) // VideoPlaylistPrivacySet |  (optional)
	description := "description_example" // string | Video playlist description (optional)
	videoChannelId := TODO // int32 | Video channel in which the playlist will be published (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VideoPlaylistsAPI.AddPlaylist(context.Background()).DisplayName(displayName).Thumbnailfile(thumbnailfile).Privacy(privacy).Description(description).VideoChannelId(videoChannelId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideoPlaylistsAPI.AddPlaylist``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddPlaylist`: AddPlaylist200Response
	fmt.Fprintf(os.Stdout, "Response from `VideoPlaylistsAPI.AddPlaylist`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddPlaylistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **displayName** | **string** | Video playlist display name | 
 **thumbnailfile** | ***os.File** | Video playlist thumbnail file | 
 **privacy** | [**VideoPlaylistPrivacySet**](VideoPlaylistPrivacySet.md) |  | 
 **description** | **string** | Video playlist description | 
 **videoChannelId** | [**int32**](int32.md) | Video channel in which the playlist will be published | 

### Return type

[**AddPlaylist200Response**](AddPlaylist200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddVideoPlaylistVideo

> AddVideoPlaylistVideo200Response AddVideoPlaylistVideo(ctx, playlistId).AddVideoPlaylistVideoRequest(addVideoPlaylistVideoRequest).Execute()

Add a video in a playlist

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
	playlistId := int32(56) // int32 | Playlist id
	addVideoPlaylistVideoRequest := *openapiclient.NewAddVideoPlaylistVideoRequest(openapiclient.addVideoPlaylistVideo_request_videoId{Int32: new(int32)}) // AddVideoPlaylistVideoRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VideoPlaylistsAPI.AddVideoPlaylistVideo(context.Background(), playlistId).AddVideoPlaylistVideoRequest(addVideoPlaylistVideoRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideoPlaylistsAPI.AddVideoPlaylistVideo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddVideoPlaylistVideo`: AddVideoPlaylistVideo200Response
	fmt.Fprintf(os.Stdout, "Response from `VideoPlaylistsAPI.AddVideoPlaylistVideo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**playlistId** | **int32** | Playlist id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddVideoPlaylistVideoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addVideoPlaylistVideoRequest** | [**AddVideoPlaylistVideoRequest**](AddVideoPlaylistVideoRequest.md) |  | 

### Return type

[**AddVideoPlaylistVideo200Response**](AddVideoPlaylistVideo200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1UsersMeVideoPlaylistsVideosExistGet

> ApiV1UsersMeVideoPlaylistsVideosExistGet200Response ApiV1UsersMeVideoPlaylistsVideosExistGet(ctx).VideoIds(videoIds).Execute()

Check video exists in my playlists

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
	videoIds := []int32{int32(42)} // []int32 | The video ids to check

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VideoPlaylistsAPI.ApiV1UsersMeVideoPlaylistsVideosExistGet(context.Background()).VideoIds(videoIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideoPlaylistsAPI.ApiV1UsersMeVideoPlaylistsVideosExistGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1UsersMeVideoPlaylistsVideosExistGet`: ApiV1UsersMeVideoPlaylistsVideosExistGet200Response
	fmt.Fprintf(os.Stdout, "Response from `VideoPlaylistsAPI.ApiV1UsersMeVideoPlaylistsVideosExistGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1UsersMeVideoPlaylistsVideosExistGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **videoIds** | **[]int32** | The video ids to check | 

### Return type

[**ApiV1UsersMeVideoPlaylistsVideosExistGet200Response**](ApiV1UsersMeVideoPlaylistsVideosExistGet200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1VideoPlaylistsPlaylistIdDelete

> ApiV1VideoPlaylistsPlaylistIdDelete(ctx, playlistId).Execute()

Delete a video playlist

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
	playlistId := int32(56) // int32 | Playlist id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VideoPlaylistsAPI.ApiV1VideoPlaylistsPlaylistIdDelete(context.Background(), playlistId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideoPlaylistsAPI.ApiV1VideoPlaylistsPlaylistIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**playlistId** | **int32** | Playlist id | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1VideoPlaylistsPlaylistIdDeleteRequest struct via the builder pattern


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


## ApiV1VideoPlaylistsPlaylistIdGet

> VideoPlaylist ApiV1VideoPlaylistsPlaylistIdGet(ctx, playlistId).Execute()

Get a video playlist

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
	playlistId := int32(56) // int32 | Playlist id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VideoPlaylistsAPI.ApiV1VideoPlaylistsPlaylistIdGet(context.Background(), playlistId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideoPlaylistsAPI.ApiV1VideoPlaylistsPlaylistIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1VideoPlaylistsPlaylistIdGet`: VideoPlaylist
	fmt.Fprintf(os.Stdout, "Response from `VideoPlaylistsAPI.ApiV1VideoPlaylistsPlaylistIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**playlistId** | **int32** | Playlist id | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1VideoPlaylistsPlaylistIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VideoPlaylist**](VideoPlaylist.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1VideoPlaylistsPlaylistIdPut

> ApiV1VideoPlaylistsPlaylistIdPut(ctx, playlistId).DisplayName(displayName).Thumbnailfile(thumbnailfile).Privacy(privacy).Description(description).VideoChannelId(videoChannelId).Execute()

Update a video playlist



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
	playlistId := int32(56) // int32 | Playlist id
	displayName := "displayName_example" // string | Video playlist display name (optional)
	thumbnailfile := os.NewFile(1234, "some_file") // *os.File | Video playlist thumbnail file (optional)
	privacy := openapiclient.VideoPlaylistPrivacySet(1) // VideoPlaylistPrivacySet |  (optional)
	description := "description_example" // string | Video playlist description (optional)
	videoChannelId := TODO // int32 | Video channel in which the playlist will be published (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VideoPlaylistsAPI.ApiV1VideoPlaylistsPlaylistIdPut(context.Background(), playlistId).DisplayName(displayName).Thumbnailfile(thumbnailfile).Privacy(privacy).Description(description).VideoChannelId(videoChannelId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideoPlaylistsAPI.ApiV1VideoPlaylistsPlaylistIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**playlistId** | **int32** | Playlist id | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1VideoPlaylistsPlaylistIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **displayName** | **string** | Video playlist display name | 
 **thumbnailfile** | ***os.File** | Video playlist thumbnail file | 
 **privacy** | [**VideoPlaylistPrivacySet**](VideoPlaylistPrivacySet.md) |  | 
 **description** | **string** | Video playlist description | 
 **videoChannelId** | [**int32**](int32.md) | Video channel in which the playlist will be published | 

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


## DelVideoPlaylistVideo

> DelVideoPlaylistVideo(ctx, playlistId, playlistElementId).Execute()

Delete an element from a playlist

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
	playlistId := int32(56) // int32 | Playlist id
	playlistElementId := int32(56) // int32 | Playlist element id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VideoPlaylistsAPI.DelVideoPlaylistVideo(context.Background(), playlistId, playlistElementId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideoPlaylistsAPI.DelVideoPlaylistVideo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**playlistId** | **int32** | Playlist id | 
**playlistElementId** | **int32** | Playlist element id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDelVideoPlaylistVideoRequest struct via the builder pattern


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


## GetPlaylistPrivacyPolicies

> []string GetPlaylistPrivacyPolicies(ctx).Execute()

List available playlist privacy policies

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VideoPlaylistsAPI.GetPlaylistPrivacyPolicies(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideoPlaylistsAPI.GetPlaylistPrivacyPolicies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlaylistPrivacyPolicies`: []string
	fmt.Fprintf(os.Stdout, "Response from `VideoPlaylistsAPI.GetPlaylistPrivacyPolicies`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlaylistPrivacyPoliciesRequest struct via the builder pattern


### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlaylists

> ApiV1VideoChannelsChannelHandleVideoPlaylistsGet200Response GetPlaylists(ctx).Start(start).Count(count).Sort(sort).PlaylistType(playlistType).Execute()

List video playlists

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
	start := int32(56) // int32 | Offset used to paginate results (optional)
	count := int32(56) // int32 | Number of items to return (optional) (default to 15)
	sort := "-createdAt" // string | Sort column (optional)
	playlistType := openapiclient.VideoPlaylistTypeSet(1) // VideoPlaylistTypeSet |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VideoPlaylistsAPI.GetPlaylists(context.Background()).Start(start).Count(count).Sort(sort).PlaylistType(playlistType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideoPlaylistsAPI.GetPlaylists``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlaylists`: ApiV1VideoChannelsChannelHandleVideoPlaylistsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `VideoPlaylistsAPI.GetPlaylists`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPlaylistsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **int32** | Offset used to paginate results | 
 **count** | **int32** | Number of items to return | [default to 15]
 **sort** | **string** | Sort column | 
 **playlistType** | [**VideoPlaylistTypeSet**](VideoPlaylistTypeSet.md) |  | 

### Return type

[**ApiV1VideoChannelsChannelHandleVideoPlaylistsGet200Response**](ApiV1VideoChannelsChannelHandleVideoPlaylistsGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVideoPlaylistVideos

> VideoListResponse GetVideoPlaylistVideos(ctx, playlistId).Start(start).Count(count).Execute()

List videos of a playlist

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
	playlistId := int32(56) // int32 | Playlist id
	start := int32(56) // int32 | Offset used to paginate results (optional)
	count := int32(56) // int32 | Number of items to return (optional) (default to 15)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VideoPlaylistsAPI.GetVideoPlaylistVideos(context.Background(), playlistId).Start(start).Count(count).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideoPlaylistsAPI.GetVideoPlaylistVideos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVideoPlaylistVideos`: VideoListResponse
	fmt.Fprintf(os.Stdout, "Response from `VideoPlaylistsAPI.GetVideoPlaylistVideos`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**playlistId** | **int32** | Playlist id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVideoPlaylistVideosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **start** | **int32** | Offset used to paginate results | 
 **count** | **int32** | Number of items to return | [default to 15]

### Return type

[**VideoListResponse**](VideoListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutVideoPlaylistVideo

> PutVideoPlaylistVideo(ctx, playlistId, playlistElementId).PutVideoPlaylistVideoRequest(putVideoPlaylistVideoRequest).Execute()

Update a playlist element

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
	playlistId := int32(56) // int32 | Playlist id
	playlistElementId := int32(56) // int32 | Playlist element id
	putVideoPlaylistVideoRequest := *openapiclient.NewPutVideoPlaylistVideoRequest() // PutVideoPlaylistVideoRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VideoPlaylistsAPI.PutVideoPlaylistVideo(context.Background(), playlistId, playlistElementId).PutVideoPlaylistVideoRequest(putVideoPlaylistVideoRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideoPlaylistsAPI.PutVideoPlaylistVideo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**playlistId** | **int32** | Playlist id | 
**playlistElementId** | **int32** | Playlist element id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutVideoPlaylistVideoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **putVideoPlaylistVideoRequest** | [**PutVideoPlaylistVideoRequest**](PutVideoPlaylistVideoRequest.md) |  | 

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


## ReorderVideoPlaylist

> ReorderVideoPlaylist(ctx, playlistId).ReorderVideoPlaylistRequest(reorderVideoPlaylistRequest).Execute()

Reorder a playlist

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
	playlistId := int32(56) // int32 | Playlist id
	reorderVideoPlaylistRequest := *openapiclient.NewReorderVideoPlaylistRequest(int32(123), int32(123)) // ReorderVideoPlaylistRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VideoPlaylistsAPI.ReorderVideoPlaylist(context.Background(), playlistId).ReorderVideoPlaylistRequest(reorderVideoPlaylistRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideoPlaylistsAPI.ReorderVideoPlaylist``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**playlistId** | **int32** | Playlist id | 

### Other Parameters

Other parameters are passed through a pointer to a apiReorderVideoPlaylistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **reorderVideoPlaylistRequest** | [**ReorderVideoPlaylistRequest**](ReorderVideoPlaylistRequest.md) |  | 

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

