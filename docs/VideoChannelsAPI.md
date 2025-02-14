# \VideoChannelsAPI

All URIs are relative to *https://peertube2.cpy.re*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddVideoChannel**](VideoChannelsAPI.md#AddVideoChannel) | **Post** /api/v1/video-channels | Create a video channel
[**ApiV1VideoChannelsChannelHandleAvatarDelete**](VideoChannelsAPI.md#ApiV1VideoChannelsChannelHandleAvatarDelete) | **Delete** /api/v1/video-channels/{channelHandle}/avatar | Delete channel avatar
[**ApiV1VideoChannelsChannelHandleAvatarPickPost**](VideoChannelsAPI.md#ApiV1VideoChannelsChannelHandleAvatarPickPost) | **Post** /api/v1/video-channels/{channelHandle}/avatar/pick | Update channel avatar
[**ApiV1VideoChannelsChannelHandleBannerDelete**](VideoChannelsAPI.md#ApiV1VideoChannelsChannelHandleBannerDelete) | **Delete** /api/v1/video-channels/{channelHandle}/banner | Delete channel banner
[**ApiV1VideoChannelsChannelHandleBannerPickPost**](VideoChannelsAPI.md#ApiV1VideoChannelsChannelHandleBannerPickPost) | **Post** /api/v1/video-channels/{channelHandle}/banner/pick | Update channel banner
[**ApiV1VideoChannelsChannelHandleImportVideosPost**](VideoChannelsAPI.md#ApiV1VideoChannelsChannelHandleImportVideosPost) | **Post** /api/v1/video-channels/{channelHandle}/import-videos | Import videos in channel
[**ApiV1VideoChannelsChannelHandleVideoPlaylistsGet**](VideoChannelsAPI.md#ApiV1VideoChannelsChannelHandleVideoPlaylistsGet) | **Get** /api/v1/video-channels/{channelHandle}/video-playlists | List playlists of a channel
[**DelVideoChannel**](VideoChannelsAPI.md#DelVideoChannel) | **Delete** /api/v1/video-channels/{channelHandle} | Delete a video channel
[**GetVideoChannel**](VideoChannelsAPI.md#GetVideoChannel) | **Get** /api/v1/video-channels/{channelHandle} | Get a video channel
[**GetVideoChannelFollowers**](VideoChannelsAPI.md#GetVideoChannelFollowers) | **Get** /api/v1/video-channels/{channelHandle}/followers | List followers of a video channel
[**GetVideoChannelVideos**](VideoChannelsAPI.md#GetVideoChannelVideos) | **Get** /api/v1/video-channels/{channelHandle}/videos | List videos of a video channel
[**GetVideoChannels**](VideoChannelsAPI.md#GetVideoChannels) | **Get** /api/v1/video-channels | List video channels
[**PutVideoChannel**](VideoChannelsAPI.md#PutVideoChannel) | **Put** /api/v1/video-channels/{channelHandle} | Update a video channel



## AddVideoChannel

> AddVideoChannel200Response AddVideoChannel(ctx).VideoChannelCreate(videoChannelCreate).Execute()

Create a video channel

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
	videoChannelCreate := *openapiclient.NewVideoChannelCreate(interface{}(123), "framasoft_videos") // VideoChannelCreate |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VideoChannelsAPI.AddVideoChannel(context.Background()).VideoChannelCreate(videoChannelCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideoChannelsAPI.AddVideoChannel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddVideoChannel`: AddVideoChannel200Response
	fmt.Fprintf(os.Stdout, "Response from `VideoChannelsAPI.AddVideoChannel`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddVideoChannelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **videoChannelCreate** | [**VideoChannelCreate**](VideoChannelCreate.md) |  | 

### Return type

[**AddVideoChannel200Response**](AddVideoChannel200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1VideoChannelsChannelHandleAvatarDelete

> ApiV1VideoChannelsChannelHandleAvatarDelete(ctx, channelHandle).Execute()

Delete channel avatar

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
	channelHandle := "my_username | my_username@example.com" // string | The video channel handle

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VideoChannelsAPI.ApiV1VideoChannelsChannelHandleAvatarDelete(context.Background(), channelHandle).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideoChannelsAPI.ApiV1VideoChannelsChannelHandleAvatarDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelHandle** | **string** | The video channel handle | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1VideoChannelsChannelHandleAvatarDeleteRequest struct via the builder pattern


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


## ApiV1VideoChannelsChannelHandleAvatarPickPost

> ApiV1UsersMeAvatarPickPost200Response ApiV1VideoChannelsChannelHandleAvatarPickPost(ctx, channelHandle).Avatarfile(avatarfile).Execute()

Update channel avatar

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
	channelHandle := "my_username | my_username@example.com" // string | The video channel handle
	avatarfile := os.NewFile(1234, "some_file") // *os.File | The file to upload. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VideoChannelsAPI.ApiV1VideoChannelsChannelHandleAvatarPickPost(context.Background(), channelHandle).Avatarfile(avatarfile).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideoChannelsAPI.ApiV1VideoChannelsChannelHandleAvatarPickPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1VideoChannelsChannelHandleAvatarPickPost`: ApiV1UsersMeAvatarPickPost200Response
	fmt.Fprintf(os.Stdout, "Response from `VideoChannelsAPI.ApiV1VideoChannelsChannelHandleAvatarPickPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelHandle** | **string** | The video channel handle | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1VideoChannelsChannelHandleAvatarPickPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **avatarfile** | ***os.File** | The file to upload. | 

### Return type

[**ApiV1UsersMeAvatarPickPost200Response**](ApiV1UsersMeAvatarPickPost200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1VideoChannelsChannelHandleBannerDelete

> ApiV1VideoChannelsChannelHandleBannerDelete(ctx, channelHandle).Execute()

Delete channel banner

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
	channelHandle := "my_username | my_username@example.com" // string | The video channel handle

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VideoChannelsAPI.ApiV1VideoChannelsChannelHandleBannerDelete(context.Background(), channelHandle).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideoChannelsAPI.ApiV1VideoChannelsChannelHandleBannerDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelHandle** | **string** | The video channel handle | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1VideoChannelsChannelHandleBannerDeleteRequest struct via the builder pattern


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


## ApiV1VideoChannelsChannelHandleBannerPickPost

> ApiV1VideoChannelsChannelHandleBannerPickPost200Response ApiV1VideoChannelsChannelHandleBannerPickPost(ctx, channelHandle).Bannerfile(bannerfile).Execute()

Update channel banner

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
	channelHandle := "my_username | my_username@example.com" // string | The video channel handle
	bannerfile := os.NewFile(1234, "some_file") // *os.File | The file to upload. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VideoChannelsAPI.ApiV1VideoChannelsChannelHandleBannerPickPost(context.Background(), channelHandle).Bannerfile(bannerfile).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideoChannelsAPI.ApiV1VideoChannelsChannelHandleBannerPickPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1VideoChannelsChannelHandleBannerPickPost`: ApiV1VideoChannelsChannelHandleBannerPickPost200Response
	fmt.Fprintf(os.Stdout, "Response from `VideoChannelsAPI.ApiV1VideoChannelsChannelHandleBannerPickPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelHandle** | **string** | The video channel handle | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1VideoChannelsChannelHandleBannerPickPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bannerfile** | ***os.File** | The file to upload. | 

### Return type

[**ApiV1VideoChannelsChannelHandleBannerPickPost200Response**](ApiV1VideoChannelsChannelHandleBannerPickPost200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1VideoChannelsChannelHandleImportVideosPost

> ApiV1VideoChannelsChannelHandleImportVideosPost(ctx, channelHandle).ImportVideosInChannelCreate(importVideosInChannelCreate).Execute()

Import videos in channel



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
	channelHandle := "my_username | my_username@example.com" // string | The video channel handle
	importVideosInChannelCreate := *openapiclient.NewImportVideosInChannelCreate("https://youtube.com/c/UC_myfancychannel") // ImportVideosInChannelCreate |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VideoChannelsAPI.ApiV1VideoChannelsChannelHandleImportVideosPost(context.Background(), channelHandle).ImportVideosInChannelCreate(importVideosInChannelCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideoChannelsAPI.ApiV1VideoChannelsChannelHandleImportVideosPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelHandle** | **string** | The video channel handle | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1VideoChannelsChannelHandleImportVideosPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **importVideosInChannelCreate** | [**ImportVideosInChannelCreate**](ImportVideosInChannelCreate.md) |  | 

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


## ApiV1VideoChannelsChannelHandleVideoPlaylistsGet

> ApiV1VideoChannelsChannelHandleVideoPlaylistsGet200Response ApiV1VideoChannelsChannelHandleVideoPlaylistsGet(ctx, channelHandle).Start(start).Count(count).Sort(sort).PlaylistType(playlistType).Execute()

List playlists of a channel

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
	channelHandle := "my_username | my_username@example.com" // string | The video channel handle
	start := int32(56) // int32 | Offset used to paginate results (optional)
	count := int32(56) // int32 | Number of items to return (optional) (default to 15)
	sort := "-createdAt" // string | Sort column (optional)
	playlistType := openapiclient.VideoPlaylistTypeSet(1) // VideoPlaylistTypeSet |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VideoChannelsAPI.ApiV1VideoChannelsChannelHandleVideoPlaylistsGet(context.Background(), channelHandle).Start(start).Count(count).Sort(sort).PlaylistType(playlistType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideoChannelsAPI.ApiV1VideoChannelsChannelHandleVideoPlaylistsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1VideoChannelsChannelHandleVideoPlaylistsGet`: ApiV1VideoChannelsChannelHandleVideoPlaylistsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `VideoChannelsAPI.ApiV1VideoChannelsChannelHandleVideoPlaylistsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelHandle** | **string** | The video channel handle | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1VideoChannelsChannelHandleVideoPlaylistsGetRequest struct via the builder pattern


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


## DelVideoChannel

> DelVideoChannel(ctx, channelHandle).Execute()

Delete a video channel

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
	channelHandle := "my_username | my_username@example.com" // string | The video channel handle

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VideoChannelsAPI.DelVideoChannel(context.Background(), channelHandle).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideoChannelsAPI.DelVideoChannel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelHandle** | **string** | The video channel handle | 

### Other Parameters

Other parameters are passed through a pointer to a apiDelVideoChannelRequest struct via the builder pattern


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


## GetVideoChannel

> VideoChannel GetVideoChannel(ctx, channelHandle).Execute()

Get a video channel

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
	channelHandle := "my_username | my_username@example.com" // string | The video channel handle

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VideoChannelsAPI.GetVideoChannel(context.Background(), channelHandle).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideoChannelsAPI.GetVideoChannel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVideoChannel`: VideoChannel
	fmt.Fprintf(os.Stdout, "Response from `VideoChannelsAPI.GetVideoChannel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelHandle** | **string** | The video channel handle | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVideoChannelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VideoChannel**](VideoChannel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVideoChannelFollowers

> GetAccountFollowers200Response GetVideoChannelFollowers(ctx, channelHandle).Start(start).Count(count).Sort(sort).Search(search).Execute()

List followers of a video channel

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
	channelHandle := "my_username | my_username@example.com" // string | The video channel handle
	start := int32(56) // int32 | Offset used to paginate results (optional)
	count := int32(56) // int32 | Number of items to return (optional) (default to 15)
	sort := "sort_example" // string | Sort followers by criteria (optional)
	search := "search_example" // string | Plain text search, applied to various parts of the model depending on endpoint (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VideoChannelsAPI.GetVideoChannelFollowers(context.Background(), channelHandle).Start(start).Count(count).Sort(sort).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideoChannelsAPI.GetVideoChannelFollowers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVideoChannelFollowers`: GetAccountFollowers200Response
	fmt.Fprintf(os.Stdout, "Response from `VideoChannelsAPI.GetVideoChannelFollowers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelHandle** | **string** | The video channel handle | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVideoChannelFollowersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **start** | **int32** | Offset used to paginate results | 
 **count** | **int32** | Number of items to return | [default to 15]
 **sort** | **string** | Sort followers by criteria | 
 **search** | **string** | Plain text search, applied to various parts of the model depending on endpoint | 

### Return type

[**GetAccountFollowers200Response**](GetAccountFollowers200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVideoChannelVideos

> VideoListResponse GetVideoChannelVideos(ctx, channelHandle).CategoryOneOf(categoryOneOf).IsLive(isLive).TagsOneOf(tagsOneOf).TagsAllOf(tagsAllOf).LicenceOneOf(licenceOneOf).LanguageOneOf(languageOneOf).AutoTagOneOf(autoTagOneOf).Nsfw(nsfw).IsLocal(isLocal).Include(include).PrivacyOneOf(privacyOneOf).HasHLSFiles(hasHLSFiles).HasWebVideoFiles(hasWebVideoFiles).SkipCount(skipCount).Start(start).Count(count).Sort(sort).ExcludeAlreadyWatched(excludeAlreadyWatched).Search(search).Execute()

List videos of a video channel

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
	channelHandle := "my_username | my_username@example.com" // string | The video channel handle
	categoryOneOf := openapiclient.getAccountVideos_categoryOneOf_parameter{ArrayOfInt32: new([]int32)} // GetAccountVideosCategoryOneOfParameter | category id of the video (see [/videos/categories](#operation/getCategories)) (optional)
	isLive := true // bool | whether or not the video is a live (optional)
	tagsOneOf := openapiclient.getAccountVideos_tagsOneOf_parameter{ArrayOfString: new([]string)} // GetAccountVideosTagsOneOfParameter | tag(s) of the video (optional)
	tagsAllOf := openapiclient.getAccountVideos_tagsAllOf_parameter{ArrayOfString: new([]string)} // GetAccountVideosTagsAllOfParameter | tag(s) of the video, where all should be present in the video (optional)
	licenceOneOf := openapiclient.getAccountVideos_licenceOneOf_parameter{ArrayOfInt32: new([]int32)} // GetAccountVideosLicenceOneOfParameter | licence id of the video (see [/videos/licences](#operation/getLicences)) (optional)
	languageOneOf := openapiclient.getAccountVideos_languageOneOf_parameter{ArrayOfString: new([]string)} // GetAccountVideosLanguageOneOfParameter | language id of the video (see [/videos/languages](#operation/getLanguages)). Use `_unknown` to filter on videos that don't have a video language (optional)
	autoTagOneOf := openapiclient.getAccountVideos_tagsAllOf_parameter{ArrayOfString: new([]string)} // GetAccountVideosTagsAllOfParameter | **PeerTube >= 6.2** **Admins and moderators only** filter on videos that contain one of these automatic tags (optional)
	nsfw := "nsfw_example" // string | whether to include nsfw videos, if any (optional)
	isLocal := true // bool | **PeerTube >= 4.0** Display only local or remote objects (optional)
	include := int32(56) // int32 | **Only administrators and moderators can use this parameter**  Include additional videos in results (can be combined using bitwise or operator) - `0` NONE - `1` NOT_PUBLISHED_STATE - `2` BLACKLISTED - `4` BLOCKED_OWNER - `8` FILES - `16` CAPTIONS - `32` VIDEO SOURCE  (optional)
	privacyOneOf := openapiclient.VideoPrivacySet(1) // VideoPrivacySet | **PeerTube >= 4.0** Display only videos in this specific privacy/privacies (optional)
	hasHLSFiles := true // bool | **PeerTube >= 4.0** Display only videos that have HLS files (optional)
	hasWebVideoFiles := true // bool | **PeerTube >= 6.0** Display only videos that have Web Video files (optional)
	skipCount := "skipCount_example" // string | if you don't need the `total` in the response (optional) (default to "false")
	start := int32(56) // int32 | Offset used to paginate results (optional)
	count := int32(56) // int32 | Number of items to return (optional) (default to 15)
	sort := "sort_example" // string |  (optional)
	excludeAlreadyWatched := true // bool | Whether or not to exclude videos that are in the user's video history (optional)
	search := "search_example" // string | Plain text search, applied to various parts of the model depending on endpoint (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VideoChannelsAPI.GetVideoChannelVideos(context.Background(), channelHandle).CategoryOneOf(categoryOneOf).IsLive(isLive).TagsOneOf(tagsOneOf).TagsAllOf(tagsAllOf).LicenceOneOf(licenceOneOf).LanguageOneOf(languageOneOf).AutoTagOneOf(autoTagOneOf).Nsfw(nsfw).IsLocal(isLocal).Include(include).PrivacyOneOf(privacyOneOf).HasHLSFiles(hasHLSFiles).HasWebVideoFiles(hasWebVideoFiles).SkipCount(skipCount).Start(start).Count(count).Sort(sort).ExcludeAlreadyWatched(excludeAlreadyWatched).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideoChannelsAPI.GetVideoChannelVideos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVideoChannelVideos`: VideoListResponse
	fmt.Fprintf(os.Stdout, "Response from `VideoChannelsAPI.GetVideoChannelVideos`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelHandle** | **string** | The video channel handle | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVideoChannelVideosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **categoryOneOf** | [**GetAccountVideosCategoryOneOfParameter**](GetAccountVideosCategoryOneOfParameter.md) | category id of the video (see [/videos/categories](#operation/getCategories)) | 
 **isLive** | **bool** | whether or not the video is a live | 
 **tagsOneOf** | [**GetAccountVideosTagsOneOfParameter**](GetAccountVideosTagsOneOfParameter.md) | tag(s) of the video | 
 **tagsAllOf** | [**GetAccountVideosTagsAllOfParameter**](GetAccountVideosTagsAllOfParameter.md) | tag(s) of the video, where all should be present in the video | 
 **licenceOneOf** | [**GetAccountVideosLicenceOneOfParameter**](GetAccountVideosLicenceOneOfParameter.md) | licence id of the video (see [/videos/licences](#operation/getLicences)) | 
 **languageOneOf** | [**GetAccountVideosLanguageOneOfParameter**](GetAccountVideosLanguageOneOfParameter.md) | language id of the video (see [/videos/languages](#operation/getLanguages)). Use &#x60;_unknown&#x60; to filter on videos that don&#39;t have a video language | 
 **autoTagOneOf** | [**GetAccountVideosTagsAllOfParameter**](GetAccountVideosTagsAllOfParameter.md) | **PeerTube &gt;&#x3D; 6.2** **Admins and moderators only** filter on videos that contain one of these automatic tags | 
 **nsfw** | **string** | whether to include nsfw videos, if any | 
 **isLocal** | **bool** | **PeerTube &gt;&#x3D; 4.0** Display only local or remote objects | 
 **include** | **int32** | **Only administrators and moderators can use this parameter**  Include additional videos in results (can be combined using bitwise or operator) - &#x60;0&#x60; NONE - &#x60;1&#x60; NOT_PUBLISHED_STATE - &#x60;2&#x60; BLACKLISTED - &#x60;4&#x60; BLOCKED_OWNER - &#x60;8&#x60; FILES - &#x60;16&#x60; CAPTIONS - &#x60;32&#x60; VIDEO SOURCE  | 
 **privacyOneOf** | [**VideoPrivacySet**](VideoPrivacySet.md) | **PeerTube &gt;&#x3D; 4.0** Display only videos in this specific privacy/privacies | 
 **hasHLSFiles** | **bool** | **PeerTube &gt;&#x3D; 4.0** Display only videos that have HLS files | 
 **hasWebVideoFiles** | **bool** | **PeerTube &gt;&#x3D; 6.0** Display only videos that have Web Video files | 
 **skipCount** | **string** | if you don&#39;t need the &#x60;total&#x60; in the response | [default to &quot;false&quot;]
 **start** | **int32** | Offset used to paginate results | 
 **count** | **int32** | Number of items to return | [default to 15]
 **sort** | **string** |  | 
 **excludeAlreadyWatched** | **bool** | Whether or not to exclude videos that are in the user&#39;s video history | 
 **search** | **string** | Plain text search, applied to various parts of the model depending on endpoint | 

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


## GetVideoChannels

> VideoChannelList GetVideoChannels(ctx).Start(start).Count(count).Sort(sort).Execute()

List video channels

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VideoChannelsAPI.GetVideoChannels(context.Background()).Start(start).Count(count).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideoChannelsAPI.GetVideoChannels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVideoChannels`: VideoChannelList
	fmt.Fprintf(os.Stdout, "Response from `VideoChannelsAPI.GetVideoChannels`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetVideoChannelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **int32** | Offset used to paginate results | 
 **count** | **int32** | Number of items to return | [default to 15]
 **sort** | **string** | Sort column | 

### Return type

[**VideoChannelList**](VideoChannelList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutVideoChannel

> PutVideoChannel(ctx, channelHandle).VideoChannelUpdate(videoChannelUpdate).Execute()

Update a video channel

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
	channelHandle := "my_username | my_username@example.com" // string | The video channel handle
	videoChannelUpdate := *openapiclient.NewVideoChannelUpdate() // VideoChannelUpdate |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VideoChannelsAPI.PutVideoChannel(context.Background(), channelHandle).VideoChannelUpdate(videoChannelUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideoChannelsAPI.PutVideoChannel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelHandle** | **string** | The video channel handle | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutVideoChannelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **videoChannelUpdate** | [**VideoChannelUpdate**](VideoChannelUpdate.md) |  | 

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

