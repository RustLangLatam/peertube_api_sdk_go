# \VideoCommentsAPI

All URIs are relative to *https://peertube2.cpy.re*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV1UsersMeVideosCommentsGet**](VideoCommentsAPI.md#ApiV1UsersMeVideosCommentsGet) | **Get** /api/v1/users/me/videos/comments | List comments on user&#39;s videos
[**ApiV1VideosCommentsGet**](VideoCommentsAPI.md#ApiV1VideosCommentsGet) | **Get** /api/v1/videos/comments | List instance comments
[**ApiV1VideosIdCommentThreadsGet**](VideoCommentsAPI.md#ApiV1VideosIdCommentThreadsGet) | **Get** /api/v1/videos/{id}/comment-threads | List threads of a video
[**ApiV1VideosIdCommentThreadsPost**](VideoCommentsAPI.md#ApiV1VideosIdCommentThreadsPost) | **Post** /api/v1/videos/{id}/comment-threads | Create a thread
[**ApiV1VideosIdCommentThreadsThreadIdGet**](VideoCommentsAPI.md#ApiV1VideosIdCommentThreadsThreadIdGet) | **Get** /api/v1/videos/{id}/comment-threads/{threadId} | Get a thread
[**ApiV1VideosIdCommentsCommentIdApprovePost**](VideoCommentsAPI.md#ApiV1VideosIdCommentsCommentIdApprovePost) | **Post** /api/v1/videos/{id}/comments/{commentId}/approve | Approve a comment
[**ApiV1VideosIdCommentsCommentIdDelete**](VideoCommentsAPI.md#ApiV1VideosIdCommentsCommentIdDelete) | **Delete** /api/v1/videos/{id}/comments/{commentId} | Delete a comment or a reply
[**ApiV1VideosIdCommentsCommentIdPost**](VideoCommentsAPI.md#ApiV1VideosIdCommentsCommentIdPost) | **Post** /api/v1/videos/{id}/comments/{commentId} | Reply to a thread of a video



## ApiV1UsersMeVideosCommentsGet

> ApiV1UsersMeVideosCommentsGet200Response ApiV1UsersMeVideosCommentsGet(ctx).Search(search).SearchAccount(searchAccount).SearchVideo(searchVideo).VideoId(videoId).VideoChannelId(videoChannelId).AutoTagOneOf(autoTagOneOf).IsHeldForReview(isHeldForReview).Execute()

List comments on user's videos



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
	search := "search_example" // string | Plain text search, applied to various parts of the model depending on endpoint (optional)
	searchAccount := "searchAccount_example" // string | Filter comments by searching on the account (optional)
	searchVideo := "searchVideo_example" // string | Filter comments by searching on the video (optional)
	videoId := int32(56) // int32 | Limit results on this specific video (optional)
	videoChannelId := int32(56) // int32 | Limit results on this specific video channel (optional)
	autoTagOneOf := openapiclient.getAccountVideos_tagsAllOf_parameter{ArrayOfString: new([]string)} // GetAccountVideosTagsAllOfParameter | **PeerTube >= 6.2** filter on comments that contain one of these automatic tags (optional)
	isHeldForReview := true // bool | only display comments that are held for review (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VideoCommentsAPI.ApiV1UsersMeVideosCommentsGet(context.Background()).Search(search).SearchAccount(searchAccount).SearchVideo(searchVideo).VideoId(videoId).VideoChannelId(videoChannelId).AutoTagOneOf(autoTagOneOf).IsHeldForReview(isHeldForReview).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideoCommentsAPI.ApiV1UsersMeVideosCommentsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1UsersMeVideosCommentsGet`: ApiV1UsersMeVideosCommentsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `VideoCommentsAPI.ApiV1UsersMeVideosCommentsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1UsersMeVideosCommentsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **string** | Plain text search, applied to various parts of the model depending on endpoint | 
 **searchAccount** | **string** | Filter comments by searching on the account | 
 **searchVideo** | **string** | Filter comments by searching on the video | 
 **videoId** | **int32** | Limit results on this specific video | 
 **videoChannelId** | **int32** | Limit results on this specific video channel | 
 **autoTagOneOf** | [**GetAccountVideosTagsAllOfParameter**](GetAccountVideosTagsAllOfParameter.md) | **PeerTube &gt;&#x3D; 6.2** filter on comments that contain one of these automatic tags | 
 **isHeldForReview** | **bool** | only display comments that are held for review | 

### Return type

[**ApiV1UsersMeVideosCommentsGet200Response**](ApiV1UsersMeVideosCommentsGet200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1VideosCommentsGet

> ApiV1UsersMeVideosCommentsGet200Response ApiV1VideosCommentsGet(ctx).Search(search).SearchAccount(searchAccount).SearchVideo(searchVideo).VideoId(videoId).VideoChannelId(videoChannelId).AutoTagOneOf(autoTagOneOf).IsLocal(isLocal).OnLocalVideo(onLocalVideo).Execute()

List instance comments

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
	search := "search_example" // string | Plain text search, applied to various parts of the model depending on endpoint (optional)
	searchAccount := "searchAccount_example" // string | Filter comments by searching on the account (optional)
	searchVideo := "searchVideo_example" // string | Filter comments by searching on the video (optional)
	videoId := int32(56) // int32 | Limit results on this specific video (optional)
	videoChannelId := int32(56) // int32 | Limit results on this specific video channel (optional)
	autoTagOneOf := openapiclient.getAccountVideos_tagsAllOf_parameter{ArrayOfString: new([]string)} // GetAccountVideosTagsAllOfParameter | **PeerTube >= 6.2** filter on comments that contain one of these automatic tags (optional)
	isLocal := true // bool | **PeerTube >= 4.0** Display only local or remote objects (optional)
	onLocalVideo := true // bool | Display only objects of local or remote videos (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VideoCommentsAPI.ApiV1VideosCommentsGet(context.Background()).Search(search).SearchAccount(searchAccount).SearchVideo(searchVideo).VideoId(videoId).VideoChannelId(videoChannelId).AutoTagOneOf(autoTagOneOf).IsLocal(isLocal).OnLocalVideo(onLocalVideo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideoCommentsAPI.ApiV1VideosCommentsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1VideosCommentsGet`: ApiV1UsersMeVideosCommentsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `VideoCommentsAPI.ApiV1VideosCommentsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1VideosCommentsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **string** | Plain text search, applied to various parts of the model depending on endpoint | 
 **searchAccount** | **string** | Filter comments by searching on the account | 
 **searchVideo** | **string** | Filter comments by searching on the video | 
 **videoId** | **int32** | Limit results on this specific video | 
 **videoChannelId** | **int32** | Limit results on this specific video channel | 
 **autoTagOneOf** | [**GetAccountVideosTagsAllOfParameter**](GetAccountVideosTagsAllOfParameter.md) | **PeerTube &gt;&#x3D; 6.2** filter on comments that contain one of these automatic tags | 
 **isLocal** | **bool** | **PeerTube &gt;&#x3D; 4.0** Display only local or remote objects | 
 **onLocalVideo** | **bool** | Display only objects of local or remote videos | 

### Return type

[**ApiV1UsersMeVideosCommentsGet200Response**](ApiV1UsersMeVideosCommentsGet200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1VideosIdCommentThreadsGet

> CommentThreadResponse ApiV1VideosIdCommentThreadsGet(ctx, id).Start(start).Count(count).Sort(sort).XPeertubeVideoPassword(xPeertubeVideoPassword).Execute()

List threads of a video

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
	start := int32(56) // int32 | Offset used to paginate results (optional)
	count := int32(56) // int32 | Number of items to return (optional) (default to 15)
	sort := "sort_example" // string | Sort comments by criteria (optional)
	xPeertubeVideoPassword := "xPeertubeVideoPassword_example" // string | Required on password protected video (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VideoCommentsAPI.ApiV1VideosIdCommentThreadsGet(context.Background(), id).Start(start).Count(count).Sort(sort).XPeertubeVideoPassword(xPeertubeVideoPassword).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideoCommentsAPI.ApiV1VideosIdCommentThreadsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1VideosIdCommentThreadsGet`: CommentThreadResponse
	fmt.Fprintf(os.Stdout, "Response from `VideoCommentsAPI.ApiV1VideosIdCommentThreadsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**ApiV1VideosOwnershipIdAcceptPostIdParameter**](.md) | The object id, uuid or short uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1VideosIdCommentThreadsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **start** | **int32** | Offset used to paginate results | 
 **count** | **int32** | Number of items to return | [default to 15]
 **sort** | **string** | Sort comments by criteria | 
 **xPeertubeVideoPassword** | **string** | Required on password protected video | 

### Return type

[**CommentThreadResponse**](CommentThreadResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1VideosIdCommentThreadsPost

> CommentThreadPostResponse ApiV1VideosIdCommentThreadsPost(ctx, id).ApiV1VideosIdCommentThreadsPostRequest(apiV1VideosIdCommentThreadsPostRequest).Execute()

Create a thread

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
	apiV1VideosIdCommentThreadsPostRequest := *openapiclient.NewApiV1VideosIdCommentThreadsPostRequest(string(123)) // ApiV1VideosIdCommentThreadsPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VideoCommentsAPI.ApiV1VideosIdCommentThreadsPost(context.Background(), id).ApiV1VideosIdCommentThreadsPostRequest(apiV1VideosIdCommentThreadsPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideoCommentsAPI.ApiV1VideosIdCommentThreadsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1VideosIdCommentThreadsPost`: CommentThreadPostResponse
	fmt.Fprintf(os.Stdout, "Response from `VideoCommentsAPI.ApiV1VideosIdCommentThreadsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**ApiV1VideosOwnershipIdAcceptPostIdParameter**](.md) | The object id, uuid or short uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1VideosIdCommentThreadsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiV1VideosIdCommentThreadsPostRequest** | [**ApiV1VideosIdCommentThreadsPostRequest**](ApiV1VideosIdCommentThreadsPostRequest.md) |  | 

### Return type

[**CommentThreadPostResponse**](CommentThreadPostResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1VideosIdCommentThreadsThreadIdGet

> VideoCommentThreadTree ApiV1VideosIdCommentThreadsThreadIdGet(ctx, id, threadId).XPeertubeVideoPassword(xPeertubeVideoPassword).Execute()

Get a thread

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
	threadId := int32(56) // int32 | The thread id (root comment id)
	xPeertubeVideoPassword := "xPeertubeVideoPassword_example" // string | Required on password protected video (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VideoCommentsAPI.ApiV1VideosIdCommentThreadsThreadIdGet(context.Background(), id, threadId).XPeertubeVideoPassword(xPeertubeVideoPassword).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideoCommentsAPI.ApiV1VideosIdCommentThreadsThreadIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1VideosIdCommentThreadsThreadIdGet`: VideoCommentThreadTree
	fmt.Fprintf(os.Stdout, "Response from `VideoCommentsAPI.ApiV1VideosIdCommentThreadsThreadIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**ApiV1VideosOwnershipIdAcceptPostIdParameter**](.md) | The object id, uuid or short uuid | 
**threadId** | **int32** | The thread id (root comment id) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1VideosIdCommentThreadsThreadIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xPeertubeVideoPassword** | **string** | Required on password protected video | 

### Return type

[**VideoCommentThreadTree**](VideoCommentThreadTree.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1VideosIdCommentsCommentIdApprovePost

> ApiV1VideosIdCommentsCommentIdApprovePost(ctx, id, commentId).Execute()

Approve a comment



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
	commentId := int32(56) // int32 | The comment id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VideoCommentsAPI.ApiV1VideosIdCommentsCommentIdApprovePost(context.Background(), id, commentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideoCommentsAPI.ApiV1VideosIdCommentsCommentIdApprovePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**ApiV1VideosOwnershipIdAcceptPostIdParameter**](.md) | The object id, uuid or short uuid | 
**commentId** | **int32** | The comment id | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1VideosIdCommentsCommentIdApprovePostRequest struct via the builder pattern


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


## ApiV1VideosIdCommentsCommentIdDelete

> ApiV1VideosIdCommentsCommentIdDelete(ctx, id, commentId).Execute()

Delete a comment or a reply

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
	commentId := int32(56) // int32 | The comment id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VideoCommentsAPI.ApiV1VideosIdCommentsCommentIdDelete(context.Background(), id, commentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideoCommentsAPI.ApiV1VideosIdCommentsCommentIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**ApiV1VideosOwnershipIdAcceptPostIdParameter**](.md) | The object id, uuid or short uuid | 
**commentId** | **int32** | The comment id | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1VideosIdCommentsCommentIdDeleteRequest struct via the builder pattern


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


## ApiV1VideosIdCommentsCommentIdPost

> CommentThreadPostResponse ApiV1VideosIdCommentsCommentIdPost(ctx, id, commentId).ApiV1VideosIdCommentThreadsPostRequest(apiV1VideosIdCommentThreadsPostRequest).XPeertubeVideoPassword(xPeertubeVideoPassword).Execute()

Reply to a thread of a video

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
	commentId := int32(56) // int32 | The comment id
	apiV1VideosIdCommentThreadsPostRequest := *openapiclient.NewApiV1VideosIdCommentThreadsPostRequest(string(123)) // ApiV1VideosIdCommentThreadsPostRequest |  (optional)
	xPeertubeVideoPassword := "xPeertubeVideoPassword_example" // string | Required on password protected video (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VideoCommentsAPI.ApiV1VideosIdCommentsCommentIdPost(context.Background(), id, commentId).ApiV1VideosIdCommentThreadsPostRequest(apiV1VideosIdCommentThreadsPostRequest).XPeertubeVideoPassword(xPeertubeVideoPassword).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideoCommentsAPI.ApiV1VideosIdCommentsCommentIdPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1VideosIdCommentsCommentIdPost`: CommentThreadPostResponse
	fmt.Fprintf(os.Stdout, "Response from `VideoCommentsAPI.ApiV1VideosIdCommentsCommentIdPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**ApiV1VideosOwnershipIdAcceptPostIdParameter**](.md) | The object id, uuid or short uuid | 
**commentId** | **int32** | The comment id | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1VideosIdCommentsCommentIdPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apiV1VideosIdCommentThreadsPostRequest** | [**ApiV1VideosIdCommentThreadsPostRequest**](ApiV1VideosIdCommentThreadsPostRequest.md) |  | 
 **xPeertubeVideoPassword** | **string** | Required on password protected video | 

### Return type

[**CommentThreadPostResponse**](CommentThreadPostResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

