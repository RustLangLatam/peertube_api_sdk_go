# \VideoAPI

All URIs are relative to *https://peertube2.cpy.re*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddView**](VideoAPI.md#AddView) | **Post** /api/v1/videos/{id}/views | Notify user is watching a video
[**ApiV1VideosIdStudioEditPost**](VideoAPI.md#ApiV1VideosIdStudioEditPost) | **Post** /api/v1/videos/{id}/studio/edit | Create a studio task
[**ApiV1VideosIdWatchingPut**](VideoAPI.md#ApiV1VideosIdWatchingPut) | **Put** /api/v1/videos/{id}/watching | Set watching progress of a video
[**DelVideo**](VideoAPI.md#DelVideo) | **Delete** /api/v1/videos/{id} | Delete a video
[**DeleteVideoSourceFile**](VideoAPI.md#DeleteVideoSourceFile) | **Delete** /api/v1/videos/{id}/source/file | Delete video source file
[**GetCategories**](VideoAPI.md#GetCategories) | **Get** /api/v1/videos/categories | List available video categories
[**GetLanguages**](VideoAPI.md#GetLanguages) | **Get** /api/v1/videos/languages | List available video languages
[**GetLicences**](VideoAPI.md#GetLicences) | **Get** /api/v1/videos/licences | List available video licences
[**GetVideo**](VideoAPI.md#GetVideo) | **Get** /api/v1/videos/{id} | Get a video
[**GetVideoDesc**](VideoAPI.md#GetVideoDesc) | **Get** /api/v1/videos/{id}/description | Get complete video description
[**GetVideoPrivacyPolicies**](VideoAPI.md#GetVideoPrivacyPolicies) | **Get** /api/v1/videos/privacies | List available video privacy policies
[**GetVideoSource**](VideoAPI.md#GetVideoSource) | **Get** /api/v1/videos/{id}/source | Get video source file metadata
[**GetVideos**](VideoAPI.md#GetVideos) | **Get** /api/v1/videos | List videos
[**ListVideoStoryboards**](VideoAPI.md#ListVideoStoryboards) | **Get** /api/v1/videos/{id}/storyboards | List storyboards of a video
[**PutVideo**](VideoAPI.md#PutVideo) | **Put** /api/v1/videos/{id} | Update a video
[**RequestVideoToken**](VideoAPI.md#RequestVideoToken) | **Post** /api/v1/videos/{id}/token | Request video token
[**UploadLegacy**](VideoAPI.md#UploadLegacy) | **Post** /api/v1/videos/upload | Upload a video
[**UploadResumable**](VideoAPI.md#UploadResumable) | **Put** /api/v1/videos/upload-resumable | Send chunk for the resumable upload of a video
[**UploadResumableCancel**](VideoAPI.md#UploadResumableCancel) | **Delete** /api/v1/videos/upload-resumable | Cancel the resumable upload of a video, deleting any data uploaded so far
[**UploadResumableInit**](VideoAPI.md#UploadResumableInit) | **Post** /api/v1/videos/upload-resumable | Initialize the resumable upload of a video



## AddView

> AddView(ctx, id).UserViewingVideo(userViewingVideo).Execute()

Notify user is watching a video



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
	userViewingVideo := *openapiclient.NewUserViewingVideo(int32(5)) // UserViewingVideo | 
	id := openapiclient._api_v1_videos_ownership__id__accept_post_id_parameter{Int32: new(int32)} // ApiV1VideosOwnershipIdAcceptPostIdParameter | The object id, uuid or short uuid

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VideoAPI.AddView(context.Background(), id).UserViewingVideo(userViewingVideo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideoAPI.AddView``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAddViewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userViewingVideo** | [**UserViewingVideo**](UserViewingVideo.md) |  | 


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1VideosIdStudioEditPost

> ApiV1VideosIdStudioEditPost(ctx, id).Execute()

Create a studio task



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
	r, err := apiClient.VideoAPI.ApiV1VideosIdStudioEditPost(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideoAPI.ApiV1VideosIdStudioEditPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiApiV1VideosIdStudioEditPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1VideosIdWatchingPut

> ApiV1VideosIdWatchingPut(ctx, id).UserViewingVideo(userViewingVideo).Execute()

Set watching progress of a video



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
	userViewingVideo := *openapiclient.NewUserViewingVideo(int32(5)) // UserViewingVideo | 
	id := openapiclient._api_v1_videos_ownership__id__accept_post_id_parameter{Int32: new(int32)} // ApiV1VideosOwnershipIdAcceptPostIdParameter | The object id, uuid or short uuid

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VideoAPI.ApiV1VideosIdWatchingPut(context.Background(), id).UserViewingVideo(userViewingVideo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideoAPI.ApiV1VideosIdWatchingPut``: %v\n", err)
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

Other parameters are passed through a pointer to a apiApiV1VideosIdWatchingPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userViewingVideo** | [**UserViewingVideo**](UserViewingVideo.md) |  | 


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


## DelVideo

> DelVideo(ctx, id).Execute()

Delete a video

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
	r, err := apiClient.VideoAPI.DelVideo(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideoAPI.DelVideo``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDelVideoRequest struct via the builder pattern


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


## DeleteVideoSourceFile

> DeleteVideoSourceFile(ctx, id).Execute()

Delete video source file

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
	r, err := apiClient.VideoAPI.DeleteVideoSourceFile(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideoAPI.DeleteVideoSourceFile``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteVideoSourceFileRequest struct via the builder pattern


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


## GetCategories

> []string GetCategories(ctx).Execute()

List available video categories

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VideoAPI.GetCategories(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideoAPI.GetCategories``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCategories`: []string
	fmt.Fprintf(os.Stdout, "Response from `VideoAPI.GetCategories`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCategoriesRequest struct via the builder pattern


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


## GetLanguages

> []string GetLanguages(ctx).Execute()

List available video languages

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VideoAPI.GetLanguages(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideoAPI.GetLanguages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLanguages`: []string
	fmt.Fprintf(os.Stdout, "Response from `VideoAPI.GetLanguages`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLanguagesRequest struct via the builder pattern


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


## GetLicences

> []string GetLicences(ctx).Execute()

List available video licences

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VideoAPI.GetLicences(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideoAPI.GetLicences``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLicences`: []string
	fmt.Fprintf(os.Stdout, "Response from `VideoAPI.GetLicences`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLicencesRequest struct via the builder pattern


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


## GetVideo

> VideoDetails GetVideo(ctx, id).XPeertubeVideoPassword(xPeertubeVideoPassword).Execute()

Get a video

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
	resp, r, err := apiClient.VideoAPI.GetVideo(context.Background(), id).XPeertubeVideoPassword(xPeertubeVideoPassword).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideoAPI.GetVideo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVideo`: VideoDetails
	fmt.Fprintf(os.Stdout, "Response from `VideoAPI.GetVideo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**ApiV1VideosOwnershipIdAcceptPostIdParameter**](.md) | The object id, uuid or short uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVideoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xPeertubeVideoPassword** | **string** | Required on password protected video | 

### Return type

[**VideoDetails**](VideoDetails.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVideoDesc

> string GetVideoDesc(ctx, id).XPeertubeVideoPassword(xPeertubeVideoPassword).Execute()

Get complete video description

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
	resp, r, err := apiClient.VideoAPI.GetVideoDesc(context.Background(), id).XPeertubeVideoPassword(xPeertubeVideoPassword).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideoAPI.GetVideoDesc``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVideoDesc`: string
	fmt.Fprintf(os.Stdout, "Response from `VideoAPI.GetVideoDesc`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**ApiV1VideosOwnershipIdAcceptPostIdParameter**](.md) | The object id, uuid or short uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVideoDescRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xPeertubeVideoPassword** | **string** | Required on password protected video | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVideoPrivacyPolicies

> []string GetVideoPrivacyPolicies(ctx).Execute()

List available video privacy policies

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VideoAPI.GetVideoPrivacyPolicies(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideoAPI.GetVideoPrivacyPolicies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVideoPrivacyPolicies`: []string
	fmt.Fprintf(os.Stdout, "Response from `VideoAPI.GetVideoPrivacyPolicies`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVideoPrivacyPoliciesRequest struct via the builder pattern


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


## GetVideoSource

> VideoSource GetVideoSource(ctx, id).Execute()

Get video source file metadata

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
	resp, r, err := apiClient.VideoAPI.GetVideoSource(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideoAPI.GetVideoSource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVideoSource`: VideoSource
	fmt.Fprintf(os.Stdout, "Response from `VideoAPI.GetVideoSource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**ApiV1VideosOwnershipIdAcceptPostIdParameter**](.md) | The object id, uuid or short uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVideoSourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VideoSource**](VideoSource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVideos

> VideoListResponse GetVideos(ctx).CategoryOneOf(categoryOneOf).IsLive(isLive).TagsOneOf(tagsOneOf).TagsAllOf(tagsAllOf).LicenceOneOf(licenceOneOf).LanguageOneOf(languageOneOf).AutoTagOneOf(autoTagOneOf).Nsfw(nsfw).IsLocal(isLocal).Include(include).PrivacyOneOf(privacyOneOf).HasHLSFiles(hasHLSFiles).HasWebVideoFiles(hasWebVideoFiles).SkipCount(skipCount).Start(start).Count(count).Sort(sort).ExcludeAlreadyWatched(excludeAlreadyWatched).Search(search).Execute()

List videos

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
	resp, r, err := apiClient.VideoAPI.GetVideos(context.Background()).CategoryOneOf(categoryOneOf).IsLive(isLive).TagsOneOf(tagsOneOf).TagsAllOf(tagsAllOf).LicenceOneOf(licenceOneOf).LanguageOneOf(languageOneOf).AutoTagOneOf(autoTagOneOf).Nsfw(nsfw).IsLocal(isLocal).Include(include).PrivacyOneOf(privacyOneOf).HasHLSFiles(hasHLSFiles).HasWebVideoFiles(hasWebVideoFiles).SkipCount(skipCount).Start(start).Count(count).Sort(sort).ExcludeAlreadyWatched(excludeAlreadyWatched).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideoAPI.GetVideos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVideos`: VideoListResponse
	fmt.Fprintf(os.Stdout, "Response from `VideoAPI.GetVideos`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetVideosRequest struct via the builder pattern


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


## ListVideoStoryboards

> ListVideoStoryboards200Response ListVideoStoryboards(ctx, id).Execute()

List storyboards of a video



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
	resp, r, err := apiClient.VideoAPI.ListVideoStoryboards(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideoAPI.ListVideoStoryboards``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListVideoStoryboards`: ListVideoStoryboards200Response
	fmt.Fprintf(os.Stdout, "Response from `VideoAPI.ListVideoStoryboards`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**ApiV1VideosOwnershipIdAcceptPostIdParameter**](.md) | The object id, uuid or short uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiListVideoStoryboardsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListVideoStoryboards200Response**](ListVideoStoryboards200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutVideo

> PutVideo(ctx, id).Thumbnailfile(thumbnailfile).Previewfile(previewfile).Category(category).Licence(licence).Language(language).Privacy(privacy).Description(description).WaitTranscoding(waitTranscoding).Support(support).Nsfw(nsfw).Name(name).Tags(tags).CommentsEnabled(commentsEnabled).CommentsPolicy(commentsPolicy).DownloadEnabled(downloadEnabled).OriginallyPublishedAt(originallyPublishedAt).ScheduleUpdate(scheduleUpdate).VideoPasswords(videoPasswords).Execute()

Update a video

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
	id := openapiclient._api_v1_videos_ownership__id__accept_post_id_parameter{Int32: new(int32)} // ApiV1VideosOwnershipIdAcceptPostIdParameter | The object id, uuid or short uuid
	thumbnailfile := os.NewFile(1234, "some_file") // *os.File | Video thumbnail file (optional)
	previewfile := os.NewFile(1234, "some_file") // *os.File | Video preview file (optional)
	category := int32(56) // int32 | category id of the video (see [/videos/categories](#operation/getCategories)) (optional)
	licence := int32(56) // int32 | licence id of the video (see [/videos/licences](#operation/getLicences)) (optional)
	language := "language_example" // string | language id of the video (see [/videos/languages](#operation/getLanguages)) (optional)
	privacy := openapiclient.VideoPrivacySet(1) // VideoPrivacySet |  (optional)
	description := "description_example" // string | Video description (optional)
	waitTranscoding := "waitTranscoding_example" // string | Whether or not we wait transcoding before publish the video (optional)
	support := "support_example" // string | A text tell the audience how to support the video creator (optional)
	nsfw := true // bool | Whether or not this video contains sensitive content (optional)
	name := "name_example" // string | Video name (optional)
	tags := []string{"Inner_example"} // []string | Video tags (maximum 5 tags each between 2 and 30 characters) (optional)
	commentsEnabled := true // bool | Deprecated in 6.2, use commentsPolicy instead (optional)
	commentsPolicy := openapiclient.VideoCommentsPolicySet(1) // VideoCommentsPolicySet |  (optional)
	downloadEnabled := true // bool | Enable or disable downloading for this video (optional)
	originallyPublishedAt := time.Now() // time.Time | Date when the content was originally published (optional)
	scheduleUpdate := *openapiclient.NewVideoScheduledUpdate(time.Now()) // VideoScheduledUpdate |  (optional)
	videoPasswords := *openapiclient.NewVideoPassword() // VideoPassword |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VideoAPI.PutVideo(context.Background(), id).Thumbnailfile(thumbnailfile).Previewfile(previewfile).Category(category).Licence(licence).Language(language).Privacy(privacy).Description(description).WaitTranscoding(waitTranscoding).Support(support).Nsfw(nsfw).Name(name).Tags(tags).CommentsEnabled(commentsEnabled).CommentsPolicy(commentsPolicy).DownloadEnabled(downloadEnabled).OriginallyPublishedAt(originallyPublishedAt).ScheduleUpdate(scheduleUpdate).VideoPasswords(videoPasswords).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideoAPI.PutVideo``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPutVideoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **thumbnailfile** | ***os.File** | Video thumbnail file | 
 **previewfile** | ***os.File** | Video preview file | 
 **category** | **int32** | category id of the video (see [/videos/categories](#operation/getCategories)) | 
 **licence** | **int32** | licence id of the video (see [/videos/licences](#operation/getLicences)) | 
 **language** | **string** | language id of the video (see [/videos/languages](#operation/getLanguages)) | 
 **privacy** | [**VideoPrivacySet**](VideoPrivacySet.md) |  | 
 **description** | **string** | Video description | 
 **waitTranscoding** | **string** | Whether or not we wait transcoding before publish the video | 
 **support** | **string** | A text tell the audience how to support the video creator | 
 **nsfw** | **bool** | Whether or not this video contains sensitive content | 
 **name** | **string** | Video name | 
 **tags** | **[]string** | Video tags (maximum 5 tags each between 2 and 30 characters) | 
 **commentsEnabled** | **bool** | Deprecated in 6.2, use commentsPolicy instead | 
 **commentsPolicy** | [**VideoCommentsPolicySet**](VideoCommentsPolicySet.md) |  | 
 **downloadEnabled** | **bool** | Enable or disable downloading for this video | 
 **originallyPublishedAt** | **time.Time** | Date when the content was originally published | 
 **scheduleUpdate** | [**VideoScheduledUpdate**](VideoScheduledUpdate.md) |  | 
 **videoPasswords** | [**VideoPassword**](VideoPassword.md) |  | 

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


## RequestVideoToken

> VideoTokenResponse RequestVideoToken(ctx, id).XPeertubeVideoPassword(xPeertubeVideoPassword).Execute()

Request video token



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
	resp, r, err := apiClient.VideoAPI.RequestVideoToken(context.Background(), id).XPeertubeVideoPassword(xPeertubeVideoPassword).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideoAPI.RequestVideoToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RequestVideoToken`: VideoTokenResponse
	fmt.Fprintf(os.Stdout, "Response from `VideoAPI.RequestVideoToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**ApiV1VideosOwnershipIdAcceptPostIdParameter**](.md) | The object id, uuid or short uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiRequestVideoTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xPeertubeVideoPassword** | **string** | Required on password protected video | 

### Return type

[**VideoTokenResponse**](VideoTokenResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadLegacy

> VideoUploadResponse UploadLegacy(ctx).Name(name).ChannelId(channelId).Videofile(videofile).Privacy(privacy).Category(category).Licence(licence).Language(language).Description(description).WaitTranscoding(waitTranscoding).GenerateTranscription(generateTranscription).Support(support).Nsfw(nsfw).Tags(tags).CommentsEnabled(commentsEnabled).CommentsPolicy(commentsPolicy).DownloadEnabled(downloadEnabled).OriginallyPublishedAt(originallyPublishedAt).ScheduleUpdate(scheduleUpdate).Thumbnailfile(thumbnailfile).Previewfile(previewfile).VideoPasswords(videoPasswords).Execute()

Upload a video



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
	name := "name_example" // string | Video name
	channelId := int32(56) // int32 | Channel id that will contain this video
	videofile := os.NewFile(1234, "some_file") // *os.File | Video file
	privacy := openapiclient.VideoPrivacySet(1) // VideoPrivacySet |  (optional)
	category := int32(56) // int32 | category id of the video (see [/videos/categories](#operation/getCategories)) (optional)
	licence := int32(56) // int32 | licence id of the video (see [/videos/licences](#operation/getLicences)) (optional)
	language := "language_example" // string | language id of the video (see [/videos/languages](#operation/getLanguages)) (optional)
	description := "description_example" // string | Video description (optional)
	waitTranscoding := true // bool | Whether or not we wait transcoding before publish the video (optional)
	generateTranscription := true // bool | **PeerTube >= 6.2** If enabled by the admin, automatically generate a subtitle of the video (optional)
	support := "support_example" // string | A text tell the audience how to support the video creator (optional)
	nsfw := true // bool | Whether or not this video contains sensitive content (optional)
	tags := []string{"Inner_example"} // []string | Video tags (maximum 5 tags each between 2 and 30 characters) (optional)
	commentsEnabled := true // bool | Deprecated in 6.2, use commentsPolicy instead (optional)
	commentsPolicy := openapiclient.VideoCommentsPolicySet(1) // VideoCommentsPolicySet |  (optional)
	downloadEnabled := true // bool | Enable or disable downloading for this video (optional)
	originallyPublishedAt := time.Now() // time.Time | Date when the content was originally published (optional)
	scheduleUpdate := *openapiclient.NewVideoScheduledUpdate(time.Now()) // VideoScheduledUpdate |  (optional)
	thumbnailfile := os.NewFile(1234, "some_file") // *os.File | Video thumbnail file (optional)
	previewfile := os.NewFile(1234, "some_file") // *os.File | Video preview file (optional)
	videoPasswords := []string{"Inner_example"} // []string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VideoAPI.UploadLegacy(context.Background()).Name(name).ChannelId(channelId).Videofile(videofile).Privacy(privacy).Category(category).Licence(licence).Language(language).Description(description).WaitTranscoding(waitTranscoding).GenerateTranscription(generateTranscription).Support(support).Nsfw(nsfw).Tags(tags).CommentsEnabled(commentsEnabled).CommentsPolicy(commentsPolicy).DownloadEnabled(downloadEnabled).OriginallyPublishedAt(originallyPublishedAt).ScheduleUpdate(scheduleUpdate).Thumbnailfile(thumbnailfile).Previewfile(previewfile).VideoPasswords(videoPasswords).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideoAPI.UploadLegacy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UploadLegacy`: VideoUploadResponse
	fmt.Fprintf(os.Stdout, "Response from `VideoAPI.UploadLegacy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUploadLegacyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | Video name | 
 **channelId** | **int32** | Channel id that will contain this video | 
 **videofile** | ***os.File** | Video file | 
 **privacy** | [**VideoPrivacySet**](VideoPrivacySet.md) |  | 
 **category** | **int32** | category id of the video (see [/videos/categories](#operation/getCategories)) | 
 **licence** | **int32** | licence id of the video (see [/videos/licences](#operation/getLicences)) | 
 **language** | **string** | language id of the video (see [/videos/languages](#operation/getLanguages)) | 
 **description** | **string** | Video description | 
 **waitTranscoding** | **bool** | Whether or not we wait transcoding before publish the video | 
 **generateTranscription** | **bool** | **PeerTube &gt;&#x3D; 6.2** If enabled by the admin, automatically generate a subtitle of the video | 
 **support** | **string** | A text tell the audience how to support the video creator | 
 **nsfw** | **bool** | Whether or not this video contains sensitive content | 
 **tags** | **[]string** | Video tags (maximum 5 tags each between 2 and 30 characters) | 
 **commentsEnabled** | **bool** | Deprecated in 6.2, use commentsPolicy instead | 
 **commentsPolicy** | [**VideoCommentsPolicySet**](VideoCommentsPolicySet.md) |  | 
 **downloadEnabled** | **bool** | Enable or disable downloading for this video | 
 **originallyPublishedAt** | **time.Time** | Date when the content was originally published | 
 **scheduleUpdate** | [**VideoScheduledUpdate**](VideoScheduledUpdate.md) |  | 
 **thumbnailfile** | ***os.File** | Video thumbnail file | 
 **previewfile** | ***os.File** | Video preview file | 
 **videoPasswords** | **[]string** |  | 

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


## UploadResumable

> VideoUploadResponse UploadResumable(ctx).UploadId(uploadId).ContentRange(contentRange).ContentLength(contentLength).Body(body).Execute()

Send chunk for the resumable upload of a video



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
	uploadId := "uploadId_example" // string | Created session id to proceed with. If you didn't send chunks in the last hour, it is not valid anymore and you need to initialize a new upload. 
	contentRange := "bytes 0-262143/2469036" // string | Specifies the bytes in the file that the request is uploading.  For example, a value of `bytes 0-262143/1000000` shows that the request is sending the first 262144 bytes (256 x 1024) in a 2,469,036 byte file. 
	contentLength := float32(262144) // float32 | Size of the chunk that the request is sending.  Remember that larger chunks are more efficient. PeerTube's web client uses chunks varying from 1048576 bytes (~1MB) and increases or reduces size depending on connection health. 
	body := os.NewFile(1234, "some_file") // *os.File |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VideoAPI.UploadResumable(context.Background()).UploadId(uploadId).ContentRange(contentRange).ContentLength(contentLength).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideoAPI.UploadResumable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UploadResumable`: VideoUploadResponse
	fmt.Fprintf(os.Stdout, "Response from `VideoAPI.UploadResumable`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUploadResumableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uploadId** | **string** | Created session id to proceed with. If you didn&#39;t send chunks in the last hour, it is not valid anymore and you need to initialize a new upload.  | 
 **contentRange** | **string** | Specifies the bytes in the file that the request is uploading.  For example, a value of &#x60;bytes 0-262143/1000000&#x60; shows that the request is sending the first 262144 bytes (256 x 1024) in a 2,469,036 byte file.  | 
 **contentLength** | **float32** | Size of the chunk that the request is sending.  Remember that larger chunks are more efficient. PeerTube&#39;s web client uses chunks varying from 1048576 bytes (~1MB) and increases or reduces size depending on connection health.  | 
 **body** | ***os.File** |  | 

### Return type

[**VideoUploadResponse**](VideoUploadResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadResumableCancel

> UploadResumableCancel(ctx).UploadId(uploadId).ContentLength(contentLength).Execute()

Cancel the resumable upload of a video, deleting any data uploaded so far



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
	uploadId := "uploadId_example" // string | Created session id to proceed with. If you didn't send chunks in the last hour, it is not valid anymore and you need to initialize a new upload. 
	contentLength := float32(0) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VideoAPI.UploadResumableCancel(context.Background()).UploadId(uploadId).ContentLength(contentLength).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideoAPI.UploadResumableCancel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUploadResumableCancelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uploadId** | **string** | Created session id to proceed with. If you didn&#39;t send chunks in the last hour, it is not valid anymore and you need to initialize a new upload.  | 
 **contentLength** | **float32** |  | 

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


## UploadResumableInit

> UploadResumableInit(ctx).XUploadContentLength(xUploadContentLength).XUploadContentType(xUploadContentType).VideoUploadRequestResumable(videoUploadRequestResumable).Execute()

Initialize the resumable upload of a video



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
	xUploadContentLength := float32(2469036) // float32 | Number of bytes that will be uploaded in subsequent requests. Set this value to the size of the file you are uploading.
	xUploadContentType := "video/mp4" // string | MIME type of the file that you are uploading. Depending on your instance settings, acceptable values might vary.
	videoUploadRequestResumable := *openapiclient.NewVideoUploadRequestResumable("What is PeerTube?", int32(3), "what_is_peertube.mp4") // VideoUploadRequestResumable |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VideoAPI.UploadResumableInit(context.Background()).XUploadContentLength(xUploadContentLength).XUploadContentType(xUploadContentType).VideoUploadRequestResumable(videoUploadRequestResumable).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideoAPI.UploadResumableInit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUploadResumableInitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xUploadContentLength** | **float32** | Number of bytes that will be uploaded in subsequent requests. Set this value to the size of the file you are uploading. | 
 **xUploadContentType** | **string** | MIME type of the file that you are uploading. Depending on your instance settings, acceptable values might vary. | 
 **videoUploadRequestResumable** | [**VideoUploadRequestResumable**](VideoUploadRequestResumable.md) |  | 

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

