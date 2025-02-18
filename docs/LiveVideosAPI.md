# \LiveVideosAPI

All URIs are relative to *https://peertube2.cpy.re*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddLive**](LiveVideosAPI.md#AddLive) | **Post** /api/v1/videos/live | Create a live
[**ApiV1VideosIdLiveSessionGet**](LiveVideosAPI.md#ApiV1VideosIdLiveSessionGet) | **Get** /api/v1/videos/{id}/live-session | Get live session of a replay
[**ApiV1VideosLiveIdSessionsGet**](LiveVideosAPI.md#ApiV1VideosLiveIdSessionsGet) | **Get** /api/v1/videos/live/{id}/sessions | List live sessions
[**GetLiveId**](LiveVideosAPI.md#GetLiveId) | **Get** /api/v1/videos/live/{id} | Get information about a live
[**UpdateLiveId**](LiveVideosAPI.md#UpdateLiveId) | **Put** /api/v1/videos/live/{id} | Update information about a live



## AddLive

> VideoUploadResponse AddLive(ctx).ChannelId(channelId).Name(name).SaveReplay(saveReplay).ReplaySettings(replaySettings).PermanentLive(permanentLive).LatencyMode(latencyMode).Thumbnailfile(thumbnailfile).Previewfile(previewfile).Privacy(privacy).Category(category).Licence(licence).Language(language).Description(description).Support(support).Nsfw(nsfw).Tags(tags).CommentsEnabled(commentsEnabled).CommentsPolicy(commentsPolicy).DownloadEnabled(downloadEnabled).Execute()

Create a live

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
	channelId := int32(56) // int32 | Channel id that will contain this live video
	name := "name_example" // string | Live video/replay name
	saveReplay := true // bool |  (optional)
	replaySettings := *openapiclient.NewLiveVideoReplaySettings() // LiveVideoReplaySettings |  (optional)
	permanentLive := true // bool | User can stream multiple times in a permanent live (optional)
	latencyMode := openapiclient.LiveVideoLatencyMode(1) // LiveVideoLatencyMode |  (optional)
	thumbnailfile := os.NewFile(1234, "some_file") // *os.File | Live video/replay thumbnail file (optional)
	previewfile := os.NewFile(1234, "some_file") // *os.File | Live video/replay preview file (optional)
	privacy := openapiclient.VideoPrivacySet(1) // VideoPrivacySet |  (optional)
	category := int32(56) // int32 | category id of the video (see [/videos/categories](#operation/getCategories)) (optional)
	licence := int32(56) // int32 | licence id of the video (see [/videos/licences](#operation/getLicences)) (optional)
	language := "language_example" // string | language id of the video (see [/videos/languages](#operation/getLanguages)) (optional)
	description := "description_example" // string | Live video/replay description (optional)
	support := "support_example" // string | A text tell the audience how to support the creator (optional)
	nsfw := true // bool | Whether or not this live video/replay contains sensitive content (optional)
	tags := []string{"Inner_example"} // []string | Live video/replay tags (maximum 5 tags each between 2 and 30 characters) (optional)
	commentsEnabled := true // bool | Deprecated in 6.2, use commentsPolicy instead (optional)
	commentsPolicy := openapiclient.VideoCommentsPolicySet(1) // VideoCommentsPolicySet |  (optional)
	downloadEnabled := true // bool | Enable or disable downloading for the replay of this live video (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LiveVideosAPI.AddLive(context.Background()).ChannelId(channelId).Name(name).SaveReplay(saveReplay).ReplaySettings(replaySettings).PermanentLive(permanentLive).LatencyMode(latencyMode).Thumbnailfile(thumbnailfile).Previewfile(previewfile).Privacy(privacy).Category(category).Licence(licence).Language(language).Description(description).Support(support).Nsfw(nsfw).Tags(tags).CommentsEnabled(commentsEnabled).CommentsPolicy(commentsPolicy).DownloadEnabled(downloadEnabled).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveVideosAPI.AddLive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddLive`: VideoUploadResponse
	fmt.Fprintf(os.Stdout, "Response from `LiveVideosAPI.AddLive`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddLiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **channelId** | **int32** | Channel id that will contain this live video | 
 **name** | **string** | Live video/replay name | 
 **saveReplay** | **bool** |  | 
 **replaySettings** | [**LiveVideoReplaySettings**](LiveVideoReplaySettings.md) |  | 
 **permanentLive** | **bool** | User can stream multiple times in a permanent live | 
 **latencyMode** | [**LiveVideoLatencyMode**](LiveVideoLatencyMode.md) |  | 
 **thumbnailfile** | ***os.File** | Live video/replay thumbnail file | 
 **previewfile** | ***os.File** | Live video/replay preview file | 
 **privacy** | [**VideoPrivacySet**](VideoPrivacySet.md) |  | 
 **category** | **int32** | category id of the video (see [/videos/categories](#operation/getCategories)) | 
 **licence** | **int32** | licence id of the video (see [/videos/licences](#operation/getLicences)) | 
 **language** | **string** | language id of the video (see [/videos/languages](#operation/getLanguages)) | 
 **description** | **string** | Live video/replay description | 
 **support** | **string** | A text tell the audience how to support the creator | 
 **nsfw** | **bool** | Whether or not this live video/replay contains sensitive content | 
 **tags** | **[]string** | Live video/replay tags (maximum 5 tags each between 2 and 30 characters) | 
 **commentsEnabled** | **bool** | Deprecated in 6.2, use commentsPolicy instead | 
 **commentsPolicy** | [**VideoCommentsPolicySet**](VideoCommentsPolicySet.md) |  | 
 **downloadEnabled** | **bool** | Enable or disable downloading for the replay of this live video | 

### Return type

[**VideoUploadResponse**](VideoUploadResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1VideosIdLiveSessionGet

> LiveVideoSessionResponse ApiV1VideosIdLiveSessionGet(ctx, id).XPeertubeVideoPassword(xPeertubeVideoPassword).Execute()

Get live session of a replay



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
	resp, r, err := apiClient.LiveVideosAPI.ApiV1VideosIdLiveSessionGet(context.Background(), id).XPeertubeVideoPassword(xPeertubeVideoPassword).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveVideosAPI.ApiV1VideosIdLiveSessionGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1VideosIdLiveSessionGet`: LiveVideoSessionResponse
	fmt.Fprintf(os.Stdout, "Response from `LiveVideosAPI.ApiV1VideosIdLiveSessionGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**ApiV1VideosOwnershipIdAcceptPostIdParameter**](.md) | The object id, uuid or short uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1VideosIdLiveSessionGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xPeertubeVideoPassword** | **string** | Required on password protected video | 

### Return type

[**LiveVideoSessionResponse**](LiveVideoSessionResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1VideosLiveIdSessionsGet

> ApiV1VideosLiveIdSessionsGet200Response ApiV1VideosLiveIdSessionsGet(ctx, id).Execute()

List live sessions



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LiveVideosAPI.ApiV1VideosLiveIdSessionsGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveVideosAPI.ApiV1VideosLiveIdSessionsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1VideosLiveIdSessionsGet`: ApiV1VideosLiveIdSessionsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `LiveVideosAPI.ApiV1VideosLiveIdSessionsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**ApiV1VideosOwnershipIdAcceptPostIdParameter**](.md) | The object id, uuid or short uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1VideosLiveIdSessionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiV1VideosLiveIdSessionsGet200Response**](ApiV1VideosLiveIdSessionsGet200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLiveId

> LiveVideoResponse GetLiveId(ctx, id).Execute()

Get information about a live

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LiveVideosAPI.GetLiveId(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveVideosAPI.GetLiveId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLiveId`: LiveVideoResponse
	fmt.Fprintf(os.Stdout, "Response from `LiveVideosAPI.GetLiveId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**ApiV1VideosOwnershipIdAcceptPostIdParameter**](.md) | The object id, uuid or short uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLiveIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LiveVideoResponse**](LiveVideoResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLiveId

> UpdateLiveId(ctx, id).LiveVideoUpdate(liveVideoUpdate).Execute()

Update information about a live

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
	liveVideoUpdate := *openapiclient.NewLiveVideoUpdate() // LiveVideoUpdate |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LiveVideosAPI.UpdateLiveId(context.Background(), id).LiveVideoUpdate(liveVideoUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveVideosAPI.UpdateLiveId``: %v\n", err)
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

Other parameters are passed through a pointer to a apiUpdateLiveIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **liveVideoUpdate** | [**LiveVideoUpdate**](LiveVideoUpdate.md) |  | 

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

