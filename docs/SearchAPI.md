# \SearchAPI

All URIs are relative to *https://peertube2.cpy.re*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchChannels**](SearchAPI.md#SearchChannels) | **Get** /api/v1/search/video-channels | Search channels
[**SearchPlaylists**](SearchAPI.md#SearchPlaylists) | **Get** /api/v1/search/video-playlists | Search playlists
[**SearchVideos**](SearchAPI.md#SearchVideos) | **Get** /api/v1/search/videos | Search videos



## SearchChannels

> VideoChannelList SearchChannels(ctx).Search(search).Start(start).Count(count).SearchTarget(searchTarget).Sort(sort).Host(host).Handles(handles).Execute()

Search channels

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
	search := "search_example" // string | String to search. If the user can make a remote URI search, and the string is an URI then the PeerTube instance will fetch the remote object and add it to its database. Then, you can use the REST API to fetch the complete channel information and interact with it. 
	start := int32(56) // int32 | Offset used to paginate results (optional)
	count := int32(56) // int32 | Number of items to return (optional) (default to 15)
	searchTarget := "searchTarget_example" // string | If the administrator enabled search index support, you can override the default search target.  **Warning**: If you choose to make an index search, PeerTube will get results from a third party service. It means the instance may not yet know the objects you fetched. If you want to load video/channel information:   * If the current user has the ability to make a remote URI search (this information is available in the config endpoint),   then reuse the search API to make a search using the object URI so PeerTube instance fetches the remote object and fill its database.   After that, you can use the classic REST API endpoints to fetch the complete object or interact with it   * If the current user doesn't have the ability to make a remote URI search, then redirect the user on the origin instance or fetch   the data from the origin instance API  (optional)
	sort := "-createdAt" // string | Sort column (optional)
	host := "host_example" // string | Find elements owned by this host (optional)
	handles := []string{"Inner_example"} // []string | Find elements with these handles (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAPI.SearchChannels(context.Background()).Search(search).Start(start).Count(count).SearchTarget(searchTarget).Sort(sort).Host(host).Handles(handles).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.SearchChannels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchChannels`: VideoChannelList
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.SearchChannels`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchChannelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **string** | String to search. If the user can make a remote URI search, and the string is an URI then the PeerTube instance will fetch the remote object and add it to its database. Then, you can use the REST API to fetch the complete channel information and interact with it.  | 
 **start** | **int32** | Offset used to paginate results | 
 **count** | **int32** | Number of items to return | [default to 15]
 **searchTarget** | **string** | If the administrator enabled search index support, you can override the default search target.  **Warning**: If you choose to make an index search, PeerTube will get results from a third party service. It means the instance may not yet know the objects you fetched. If you want to load video/channel information:   * If the current user has the ability to make a remote URI search (this information is available in the config endpoint),   then reuse the search API to make a search using the object URI so PeerTube instance fetches the remote object and fill its database.   After that, you can use the classic REST API endpoints to fetch the complete object or interact with it   * If the current user doesn&#39;t have the ability to make a remote URI search, then redirect the user on the origin instance or fetch   the data from the origin instance API  | 
 **sort** | **string** | Sort column | 
 **host** | **string** | Find elements owned by this host | 
 **handles** | **[]string** | Find elements with these handles | 

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


## SearchPlaylists

> ApiV1VideoChannelsChannelHandleVideoPlaylistsGet200Response SearchPlaylists(ctx).Search(search).Start(start).Count(count).SearchTarget(searchTarget).Sort(sort).Host(host).Uuids(uuids).Execute()

Search playlists

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
	search := "search_example" // string | String to search. If the user can make a remote URI search, and the string is an URI then the PeerTube instance will fetch the remote object and add it to its database. Then, you can use the REST API to fetch the complete playlist information and interact with it. 
	start := int32(56) // int32 | Offset used to paginate results (optional)
	count := int32(56) // int32 | Number of items to return (optional) (default to 15)
	searchTarget := "searchTarget_example" // string | If the administrator enabled search index support, you can override the default search target.  **Warning**: If you choose to make an index search, PeerTube will get results from a third party service. It means the instance may not yet know the objects you fetched. If you want to load video/channel information:   * If the current user has the ability to make a remote URI search (this information is available in the config endpoint),   then reuse the search API to make a search using the object URI so PeerTube instance fetches the remote object and fill its database.   After that, you can use the classic REST API endpoints to fetch the complete object or interact with it   * If the current user doesn't have the ability to make a remote URI search, then redirect the user on the origin instance or fetch   the data from the origin instance API  (optional)
	sort := "-createdAt" // string | Sort column (optional)
	host := "host_example" // string | Find elements owned by this host (optional)
	uuids := []string{"Inner_example"} // []string | Find elements with specific UUIDs (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAPI.SearchPlaylists(context.Background()).Search(search).Start(start).Count(count).SearchTarget(searchTarget).Sort(sort).Host(host).Uuids(uuids).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.SearchPlaylists``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchPlaylists`: ApiV1VideoChannelsChannelHandleVideoPlaylistsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.SearchPlaylists`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchPlaylistsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **string** | String to search. If the user can make a remote URI search, and the string is an URI then the PeerTube instance will fetch the remote object and add it to its database. Then, you can use the REST API to fetch the complete playlist information and interact with it.  | 
 **start** | **int32** | Offset used to paginate results | 
 **count** | **int32** | Number of items to return | [default to 15]
 **searchTarget** | **string** | If the administrator enabled search index support, you can override the default search target.  **Warning**: If you choose to make an index search, PeerTube will get results from a third party service. It means the instance may not yet know the objects you fetched. If you want to load video/channel information:   * If the current user has the ability to make a remote URI search (this information is available in the config endpoint),   then reuse the search API to make a search using the object URI so PeerTube instance fetches the remote object and fill its database.   After that, you can use the classic REST API endpoints to fetch the complete object or interact with it   * If the current user doesn&#39;t have the ability to make a remote URI search, then redirect the user on the origin instance or fetch   the data from the origin instance API  | 
 **sort** | **string** | Sort column | 
 **host** | **string** | Find elements owned by this host | 
 **uuids** | **[]string** | Find elements with specific UUIDs | 

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


## SearchVideos

> VideoListResponse SearchVideos(ctx).Search(search).CategoryOneOf(categoryOneOf).IsLive(isLive).TagsOneOf(tagsOneOf).TagsAllOf(tagsAllOf).LicenceOneOf(licenceOneOf).LanguageOneOf(languageOneOf).AutoTagOneOf(autoTagOneOf).Nsfw(nsfw).IsLocal(isLocal).Include(include).PrivacyOneOf(privacyOneOf).Uuids(uuids).HasHLSFiles(hasHLSFiles).HasWebVideoFiles(hasWebVideoFiles).SkipCount(skipCount).Start(start).Count(count).SearchTarget(searchTarget).Sort(sort).ExcludeAlreadyWatched(excludeAlreadyWatched).Host(host).StartDate(startDate).EndDate(endDate).OriginallyPublishedStartDate(originallyPublishedStartDate).OriginallyPublishedEndDate(originallyPublishedEndDate).DurationMin(durationMin).DurationMax(durationMax).Execute()

Search videos

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
	search := "search_example" // string | String to search. If the user can make a remote URI search, and the string is an URI then the PeerTube instance will fetch the remote object and add it to its database. Then, you can use the REST API to fetch the complete video information and interact with it. 
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
	uuids := []string{"Inner_example"} // []string | Find elements with specific UUIDs (optional)
	hasHLSFiles := true // bool | **PeerTube >= 4.0** Display only videos that have HLS files (optional)
	hasWebVideoFiles := true // bool | **PeerTube >= 6.0** Display only videos that have Web Video files (optional)
	skipCount := "skipCount_example" // string | if you don't need the `total` in the response (optional) (default to "false")
	start := int32(56) // int32 | Offset used to paginate results (optional)
	count := int32(56) // int32 | Number of items to return (optional) (default to 15)
	searchTarget := "searchTarget_example" // string | If the administrator enabled search index support, you can override the default search target.  **Warning**: If you choose to make an index search, PeerTube will get results from a third party service. It means the instance may not yet know the objects you fetched. If you want to load video/channel information:   * If the current user has the ability to make a remote URI search (this information is available in the config endpoint),   then reuse the search API to make a search using the object URI so PeerTube instance fetches the remote object and fill its database.   After that, you can use the classic REST API endpoints to fetch the complete object or interact with it   * If the current user doesn't have the ability to make a remote URI search, then redirect the user on the origin instance or fetch   the data from the origin instance API  (optional)
	sort := "sort_example" // string | Sort videos by criteria (prefixing with `-` means `DESC` order):  (optional)
	excludeAlreadyWatched := true // bool | Whether or not to exclude videos that are in the user's video history (optional)
	host := "host_example" // string | Find elements owned by this host (optional)
	startDate := time.Now() // time.Time | Get videos that are published after this date (optional)
	endDate := time.Now() // time.Time | Get videos that are published before this date (optional)
	originallyPublishedStartDate := time.Now() // time.Time | Get videos that are originally published after this date (optional)
	originallyPublishedEndDate := time.Now() // time.Time | Get videos that are originally published before this date (optional)
	durationMin := int32(56) // int32 | Get videos that have this minimum duration (optional)
	durationMax := int32(56) // int32 | Get videos that have this maximum duration (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAPI.SearchVideos(context.Background()).Search(search).CategoryOneOf(categoryOneOf).IsLive(isLive).TagsOneOf(tagsOneOf).TagsAllOf(tagsAllOf).LicenceOneOf(licenceOneOf).LanguageOneOf(languageOneOf).AutoTagOneOf(autoTagOneOf).Nsfw(nsfw).IsLocal(isLocal).Include(include).PrivacyOneOf(privacyOneOf).Uuids(uuids).HasHLSFiles(hasHLSFiles).HasWebVideoFiles(hasWebVideoFiles).SkipCount(skipCount).Start(start).Count(count).SearchTarget(searchTarget).Sort(sort).ExcludeAlreadyWatched(excludeAlreadyWatched).Host(host).StartDate(startDate).EndDate(endDate).OriginallyPublishedStartDate(originallyPublishedStartDate).OriginallyPublishedEndDate(originallyPublishedEndDate).DurationMin(durationMin).DurationMax(durationMax).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.SearchVideos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchVideos`: VideoListResponse
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.SearchVideos`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchVideosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **string** | String to search. If the user can make a remote URI search, and the string is an URI then the PeerTube instance will fetch the remote object and add it to its database. Then, you can use the REST API to fetch the complete video information and interact with it.  | 
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
 **uuids** | **[]string** | Find elements with specific UUIDs | 
 **hasHLSFiles** | **bool** | **PeerTube &gt;&#x3D; 4.0** Display only videos that have HLS files | 
 **hasWebVideoFiles** | **bool** | **PeerTube &gt;&#x3D; 6.0** Display only videos that have Web Video files | 
 **skipCount** | **string** | if you don&#39;t need the &#x60;total&#x60; in the response | [default to &quot;false&quot;]
 **start** | **int32** | Offset used to paginate results | 
 **count** | **int32** | Number of items to return | [default to 15]
 **searchTarget** | **string** | If the administrator enabled search index support, you can override the default search target.  **Warning**: If you choose to make an index search, PeerTube will get results from a third party service. It means the instance may not yet know the objects you fetched. If you want to load video/channel information:   * If the current user has the ability to make a remote URI search (this information is available in the config endpoint),   then reuse the search API to make a search using the object URI so PeerTube instance fetches the remote object and fill its database.   After that, you can use the classic REST API endpoints to fetch the complete object or interact with it   * If the current user doesn&#39;t have the ability to make a remote URI search, then redirect the user on the origin instance or fetch   the data from the origin instance API  | 
 **sort** | **string** | Sort videos by criteria (prefixing with &#x60;-&#x60; means &#x60;DESC&#x60; order):  | 
 **excludeAlreadyWatched** | **bool** | Whether or not to exclude videos that are in the user&#39;s video history | 
 **host** | **string** | Find elements owned by this host | 
 **startDate** | **time.Time** | Get videos that are published after this date | 
 **endDate** | **time.Time** | Get videos that are published before this date | 
 **originallyPublishedStartDate** | **time.Time** | Get videos that are originally published after this date | 
 **originallyPublishedEndDate** | **time.Time** | Get videos that are originally published before this date | 
 **durationMin** | **int32** | Get videos that have this minimum duration | 
 **durationMax** | **int32** | Get videos that have this maximum duration | 

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

