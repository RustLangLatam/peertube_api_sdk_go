# \VideoFeedsAPI

All URIs are relative to *https://peertube2.cpy.re*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSyndicatedComments**](VideoFeedsAPI.md#GetSyndicatedComments) | **Get** /feeds/video-comments.{format} | Comments on videos feeds
[**GetSyndicatedSubscriptionVideos**](VideoFeedsAPI.md#GetSyndicatedSubscriptionVideos) | **Get** /feeds/subscriptions.{format} | Videos of subscriptions feeds
[**GetSyndicatedVideos**](VideoFeedsAPI.md#GetSyndicatedVideos) | **Get** /feeds/videos.{format} | Common videos feeds
[**GetVideosPodcastFeed**](VideoFeedsAPI.md#GetVideosPodcastFeed) | **Get** /feeds/podcast/videos.xml | Videos podcast feed



## GetSyndicatedComments

> []VideoCommentsForXMLInner GetSyndicatedComments(ctx, format).VideoId(videoId).AccountId(accountId).AccountName(accountName).VideoChannelId(videoChannelId).VideoChannelName(videoChannelName).Execute()

Comments on videos feeds

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
	format := "format_example" // string | format expected (we focus on making `rss` the most featureful ; it serves [Media RSS](https://www.rssboard.org/media-rss))
	videoId := "videoId_example" // string | limit listing comments to a specific video (optional)
	accountId := "accountId_example" // string | limit listing comments to videos of a specific account (optional)
	accountName := "accountName_example" // string | limit listing comments to videos of a specific account (optional)
	videoChannelId := "videoChannelId_example" // string | limit listing comments to videos of a specific video channel (optional)
	videoChannelName := "videoChannelName_example" // string | limit listing comments to videos of a specific video channel (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VideoFeedsAPI.GetSyndicatedComments(context.Background(), format).VideoId(videoId).AccountId(accountId).AccountName(accountName).VideoChannelId(videoChannelId).VideoChannelName(videoChannelName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideoFeedsAPI.GetSyndicatedComments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSyndicatedComments`: []VideoCommentsForXMLInner
	fmt.Fprintf(os.Stdout, "Response from `VideoFeedsAPI.GetSyndicatedComments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**format** | **string** | format expected (we focus on making &#x60;rss&#x60; the most featureful ; it serves [Media RSS](https://www.rssboard.org/media-rss)) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSyndicatedCommentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **videoId** | **string** | limit listing comments to a specific video | 
 **accountId** | **string** | limit listing comments to videos of a specific account | 
 **accountName** | **string** | limit listing comments to videos of a specific account | 
 **videoChannelId** | **string** | limit listing comments to videos of a specific video channel | 
 **videoChannelName** | **string** | limit listing comments to videos of a specific video channel | 

### Return type

[**[]VideoCommentsForXMLInner**](VideoCommentsForXMLInner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/xml, application/rss+xml, text/xml, application/atom+xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSyndicatedSubscriptionVideos

> []VideosForXMLInner GetSyndicatedSubscriptionVideos(ctx, format).AccountId(accountId).Token(token).Sort(sort).Nsfw(nsfw).IsLocal(isLocal).Include(include).PrivacyOneOf(privacyOneOf).HasHLSFiles(hasHLSFiles).HasWebVideoFiles(hasWebVideoFiles).Execute()

Videos of subscriptions feeds

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
	format := "format_example" // string | format expected (we focus on making `rss` the most featureful ; it serves [Media RSS](https://www.rssboard.org/media-rss))
	accountId := "accountId_example" // string | limit listing to a specific account
	token := "token_example" // string | private token allowing access
	sort := "-createdAt" // string | Sort column (optional)
	nsfw := "nsfw_example" // string | whether to include nsfw videos, if any (optional)
	isLocal := true // bool | **PeerTube >= 4.0** Display only local or remote objects (optional)
	include := int32(56) // int32 | **Only administrators and moderators can use this parameter**  Include additional videos in results (can be combined using bitwise or operator) - `0` NONE - `1` NOT_PUBLISHED_STATE - `2` BLACKLISTED - `4` BLOCKED_OWNER - `8` FILES - `16` CAPTIONS - `32` VIDEO SOURCE  (optional)
	privacyOneOf := openapiclient.VideoPrivacySet(1) // VideoPrivacySet | **PeerTube >= 4.0** Display only videos in this specific privacy/privacies (optional)
	hasHLSFiles := true // bool | **PeerTube >= 4.0** Display only videos that have HLS files (optional)
	hasWebVideoFiles := true // bool | **PeerTube >= 6.0** Display only videos that have Web Video files (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VideoFeedsAPI.GetSyndicatedSubscriptionVideos(context.Background(), format).AccountId(accountId).Token(token).Sort(sort).Nsfw(nsfw).IsLocal(isLocal).Include(include).PrivacyOneOf(privacyOneOf).HasHLSFiles(hasHLSFiles).HasWebVideoFiles(hasWebVideoFiles).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideoFeedsAPI.GetSyndicatedSubscriptionVideos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSyndicatedSubscriptionVideos`: []VideosForXMLInner
	fmt.Fprintf(os.Stdout, "Response from `VideoFeedsAPI.GetSyndicatedSubscriptionVideos`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**format** | **string** | format expected (we focus on making &#x60;rss&#x60; the most featureful ; it serves [Media RSS](https://www.rssboard.org/media-rss)) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSyndicatedSubscriptionVideosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountId** | **string** | limit listing to a specific account | 
 **token** | **string** | private token allowing access | 
 **sort** | **string** | Sort column | 
 **nsfw** | **string** | whether to include nsfw videos, if any | 
 **isLocal** | **bool** | **PeerTube &gt;&#x3D; 4.0** Display only local or remote objects | 
 **include** | **int32** | **Only administrators and moderators can use this parameter**  Include additional videos in results (can be combined using bitwise or operator) - &#x60;0&#x60; NONE - &#x60;1&#x60; NOT_PUBLISHED_STATE - &#x60;2&#x60; BLACKLISTED - &#x60;4&#x60; BLOCKED_OWNER - &#x60;8&#x60; FILES - &#x60;16&#x60; CAPTIONS - &#x60;32&#x60; VIDEO SOURCE  | 
 **privacyOneOf** | [**VideoPrivacySet**](VideoPrivacySet.md) | **PeerTube &gt;&#x3D; 4.0** Display only videos in this specific privacy/privacies | 
 **hasHLSFiles** | **bool** | **PeerTube &gt;&#x3D; 4.0** Display only videos that have HLS files | 
 **hasWebVideoFiles** | **bool** | **PeerTube &gt;&#x3D; 6.0** Display only videos that have Web Video files | 

### Return type

[**[]VideosForXMLInner**](VideosForXMLInner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/xml, application/rss+xml, text/xml, application/atom+xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSyndicatedVideos

> []VideosForXMLInner GetSyndicatedVideos(ctx, format).AccountId(accountId).AccountName(accountName).VideoChannelId(videoChannelId).VideoChannelName(videoChannelName).Sort(sort).Nsfw(nsfw).IsLocal(isLocal).Include(include).PrivacyOneOf(privacyOneOf).HasHLSFiles(hasHLSFiles).HasWebVideoFiles(hasWebVideoFiles).Execute()

Common videos feeds

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
	format := "format_example" // string | format expected (we focus on making `rss` the most featureful ; it serves [Media RSS](https://www.rssboard.org/media-rss))
	accountId := "accountId_example" // string | limit listing to a specific account (optional)
	accountName := "accountName_example" // string | limit listing to a specific account (optional)
	videoChannelId := "videoChannelId_example" // string | limit listing to a specific video channel (optional)
	videoChannelName := "videoChannelName_example" // string | limit listing to a specific video channel (optional)
	sort := "-createdAt" // string | Sort column (optional)
	nsfw := "nsfw_example" // string | whether to include nsfw videos, if any (optional)
	isLocal := true // bool | **PeerTube >= 4.0** Display only local or remote objects (optional)
	include := int32(56) // int32 | **Only administrators and moderators can use this parameter**  Include additional videos in results (can be combined using bitwise or operator) - `0` NONE - `1` NOT_PUBLISHED_STATE - `2` BLACKLISTED - `4` BLOCKED_OWNER - `8` FILES - `16` CAPTIONS - `32` VIDEO SOURCE  (optional)
	privacyOneOf := openapiclient.VideoPrivacySet(1) // VideoPrivacySet | **PeerTube >= 4.0** Display only videos in this specific privacy/privacies (optional)
	hasHLSFiles := true // bool | **PeerTube >= 4.0** Display only videos that have HLS files (optional)
	hasWebVideoFiles := true // bool | **PeerTube >= 6.0** Display only videos that have Web Video files (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VideoFeedsAPI.GetSyndicatedVideos(context.Background(), format).AccountId(accountId).AccountName(accountName).VideoChannelId(videoChannelId).VideoChannelName(videoChannelName).Sort(sort).Nsfw(nsfw).IsLocal(isLocal).Include(include).PrivacyOneOf(privacyOneOf).HasHLSFiles(hasHLSFiles).HasWebVideoFiles(hasWebVideoFiles).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideoFeedsAPI.GetSyndicatedVideos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSyndicatedVideos`: []VideosForXMLInner
	fmt.Fprintf(os.Stdout, "Response from `VideoFeedsAPI.GetSyndicatedVideos`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**format** | **string** | format expected (we focus on making &#x60;rss&#x60; the most featureful ; it serves [Media RSS](https://www.rssboard.org/media-rss)) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSyndicatedVideosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountId** | **string** | limit listing to a specific account | 
 **accountName** | **string** | limit listing to a specific account | 
 **videoChannelId** | **string** | limit listing to a specific video channel | 
 **videoChannelName** | **string** | limit listing to a specific video channel | 
 **sort** | **string** | Sort column | 
 **nsfw** | **string** | whether to include nsfw videos, if any | 
 **isLocal** | **bool** | **PeerTube &gt;&#x3D; 4.0** Display only local or remote objects | 
 **include** | **int32** | **Only administrators and moderators can use this parameter**  Include additional videos in results (can be combined using bitwise or operator) - &#x60;0&#x60; NONE - &#x60;1&#x60; NOT_PUBLISHED_STATE - &#x60;2&#x60; BLACKLISTED - &#x60;4&#x60; BLOCKED_OWNER - &#x60;8&#x60; FILES - &#x60;16&#x60; CAPTIONS - &#x60;32&#x60; VIDEO SOURCE  | 
 **privacyOneOf** | [**VideoPrivacySet**](VideoPrivacySet.md) | **PeerTube &gt;&#x3D; 4.0** Display only videos in this specific privacy/privacies | 
 **hasHLSFiles** | **bool** | **PeerTube &gt;&#x3D; 4.0** Display only videos that have HLS files | 
 **hasWebVideoFiles** | **bool** | **PeerTube &gt;&#x3D; 6.0** Display only videos that have Web Video files | 

### Return type

[**[]VideosForXMLInner**](VideosForXMLInner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/xml, application/rss+xml, text/xml, application/atom+xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVideosPodcastFeed

> GetVideosPodcastFeed(ctx).VideoChannelId(videoChannelId).Execute()

Videos podcast feed

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
	videoChannelId := "videoChannelId_example" // string | Limit listing to a specific video channel

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VideoFeedsAPI.GetVideosPodcastFeed(context.Background()).VideoChannelId(videoChannelId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideoFeedsAPI.GetVideosPodcastFeed``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetVideosPodcastFeedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **videoChannelId** | **string** | Limit listing to a specific video channel | 

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

