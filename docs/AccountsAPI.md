# \AccountsAPI

All URIs are relative to *https://peertube2.cpy.re*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV1AccountsNameRatingsGet**](AccountsAPI.md#ApiV1AccountsNameRatingsGet) | **Get** /api/v1/accounts/{name}/ratings | List ratings of an account
[**ApiV1AccountsNameVideoChannelSyncsGet**](AccountsAPI.md#ApiV1AccountsNameVideoChannelSyncsGet) | **Get** /api/v1/accounts/{name}/video-channel-syncs | List the synchronizations of video channels of an account
[**ApiV1AccountsNameVideoChannelsGet**](AccountsAPI.md#ApiV1AccountsNameVideoChannelsGet) | **Get** /api/v1/accounts/{name}/video-channels | List video channels of an account
[**ApiV1AccountsNameVideoPlaylistsGet**](AccountsAPI.md#ApiV1AccountsNameVideoPlaylistsGet) | **Get** /api/v1/accounts/{name}/video-playlists | List playlists of an account
[**GetAccount**](AccountsAPI.md#GetAccount) | **Get** /api/v1/accounts/{name} | Get an account
[**GetAccountFollowers**](AccountsAPI.md#GetAccountFollowers) | **Get** /api/v1/accounts/{name}/followers | List followers of an account
[**GetAccountVideos**](AccountsAPI.md#GetAccountVideos) | **Get** /api/v1/accounts/{name}/videos | List videos of an account
[**GetAccounts**](AccountsAPI.md#GetAccounts) | **Get** /api/v1/accounts | List accounts



## ApiV1AccountsNameRatingsGet

> []VideoRating ApiV1AccountsNameRatingsGet(ctx, name).Start(start).Count(count).Sort(sort).Rating(rating).Execute()

List ratings of an account

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
	name := "chocobozzz | chocobozzz@example.org" // string | The username or handle of the account
	start := int32(56) // int32 | Offset used to paginate results (optional)
	count := int32(56) // int32 | Number of items to return (optional) (default to 15)
	sort := "-createdAt" // string | Sort column (optional)
	rating := "rating_example" // string | Optionally filter which ratings to retrieve (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsAPI.ApiV1AccountsNameRatingsGet(context.Background(), name).Start(start).Count(count).Sort(sort).Rating(rating).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.ApiV1AccountsNameRatingsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1AccountsNameRatingsGet`: []VideoRating
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.ApiV1AccountsNameRatingsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The username or handle of the account | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1AccountsNameRatingsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **start** | **int32** | Offset used to paginate results | 
 **count** | **int32** | Number of items to return | [default to 15]
 **sort** | **string** | Sort column | 
 **rating** | **string** | Optionally filter which ratings to retrieve | 

### Return type

[**[]VideoRating**](VideoRating.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1AccountsNameVideoChannelSyncsGet

> VideoChannelSyncList ApiV1AccountsNameVideoChannelSyncsGet(ctx, name).Start(start).Count(count).Sort(sort).Execute()

List the synchronizations of video channels of an account

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
	name := "chocobozzz | chocobozzz@example.org" // string | The username or handle of the account
	start := int32(56) // int32 | Offset used to paginate results (optional)
	count := int32(56) // int32 | Number of items to return (optional) (default to 15)
	sort := "-createdAt" // string | Sort column (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsAPI.ApiV1AccountsNameVideoChannelSyncsGet(context.Background(), name).Start(start).Count(count).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.ApiV1AccountsNameVideoChannelSyncsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1AccountsNameVideoChannelSyncsGet`: VideoChannelSyncList
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.ApiV1AccountsNameVideoChannelSyncsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The username or handle of the account | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1AccountsNameVideoChannelSyncsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **start** | **int32** | Offset used to paginate results | 
 **count** | **int32** | Number of items to return | [default to 15]
 **sort** | **string** | Sort column | 

### Return type

[**VideoChannelSyncList**](VideoChannelSyncList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1AccountsNameVideoChannelsGet

> VideoChannelList ApiV1AccountsNameVideoChannelsGet(ctx, name).WithStats(withStats).Start(start).Count(count).Sort(sort).Execute()

List video channels of an account

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
	name := "chocobozzz | chocobozzz@example.org" // string | The username or handle of the account
	withStats := true // bool | include daily view statistics for the last 30 days and total views (only if authentified as the account user) (optional)
	start := int32(56) // int32 | Offset used to paginate results (optional)
	count := int32(56) // int32 | Number of items to return (optional) (default to 15)
	sort := "-createdAt" // string | Sort column (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsAPI.ApiV1AccountsNameVideoChannelsGet(context.Background(), name).WithStats(withStats).Start(start).Count(count).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.ApiV1AccountsNameVideoChannelsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1AccountsNameVideoChannelsGet`: VideoChannelList
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.ApiV1AccountsNameVideoChannelsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The username or handle of the account | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1AccountsNameVideoChannelsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **withStats** | **bool** | include daily view statistics for the last 30 days and total views (only if authentified as the account user) | 
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


## ApiV1AccountsNameVideoPlaylistsGet

> ApiV1VideoChannelsChannelHandleVideoPlaylistsGet200Response ApiV1AccountsNameVideoPlaylistsGet(ctx, name).Start(start).Count(count).Sort(sort).Search(search).PlaylistType(playlistType).Execute()

List playlists of an account

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
	name := "chocobozzz | chocobozzz@example.org" // string | The username or handle of the account
	start := int32(56) // int32 | Offset used to paginate results (optional)
	count := int32(56) // int32 | Number of items to return (optional) (default to 15)
	sort := "-createdAt" // string | Sort column (optional)
	search := "search_example" // string | Plain text search, applied to various parts of the model depending on endpoint (optional)
	playlistType := openapiclient.VideoPlaylistTypeSet(1) // VideoPlaylistTypeSet |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsAPI.ApiV1AccountsNameVideoPlaylistsGet(context.Background(), name).Start(start).Count(count).Sort(sort).Search(search).PlaylistType(playlistType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.ApiV1AccountsNameVideoPlaylistsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1AccountsNameVideoPlaylistsGet`: ApiV1VideoChannelsChannelHandleVideoPlaylistsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.ApiV1AccountsNameVideoPlaylistsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The username or handle of the account | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1AccountsNameVideoPlaylistsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **start** | **int32** | Offset used to paginate results | 
 **count** | **int32** | Number of items to return | [default to 15]
 **sort** | **string** | Sort column | 
 **search** | **string** | Plain text search, applied to various parts of the model depending on endpoint | 
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


## GetAccount

> Account GetAccount(ctx, name).Execute()

Get an account

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
	name := "chocobozzz | chocobozzz@example.org" // string | The username or handle of the account

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsAPI.GetAccount(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.GetAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccount`: Account
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.GetAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The username or handle of the account | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Account**](Account.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountFollowers

> GetAccountFollowers200Response GetAccountFollowers(ctx, name).Start(start).Count(count).Sort(sort).Search(search).Execute()

List followers of an account

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
	name := "chocobozzz | chocobozzz@example.org" // string | The username or handle of the account
	start := int32(56) // int32 | Offset used to paginate results (optional)
	count := int32(56) // int32 | Number of items to return (optional) (default to 15)
	sort := "sort_example" // string | Sort followers by criteria (optional)
	search := "search_example" // string | Plain text search, applied to various parts of the model depending on endpoint (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsAPI.GetAccountFollowers(context.Background(), name).Start(start).Count(count).Sort(sort).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.GetAccountFollowers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountFollowers`: GetAccountFollowers200Response
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.GetAccountFollowers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The username or handle of the account | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountFollowersRequest struct via the builder pattern


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


## GetAccountVideos

> VideoListResponse GetAccountVideos(ctx, name).CategoryOneOf(categoryOneOf).IsLive(isLive).TagsOneOf(tagsOneOf).TagsAllOf(tagsAllOf).LicenceOneOf(licenceOneOf).LanguageOneOf(languageOneOf).AutoTagOneOf(autoTagOneOf).Nsfw(nsfw).IsLocal(isLocal).Include(include).PrivacyOneOf(privacyOneOf).HasHLSFiles(hasHLSFiles).HasWebVideoFiles(hasWebVideoFiles).SkipCount(skipCount).Start(start).Count(count).Sort(sort).ExcludeAlreadyWatched(excludeAlreadyWatched).Search(search).Execute()

List videos of an account

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
	name := "chocobozzz | chocobozzz@example.org" // string | The username or handle of the account
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
	resp, r, err := apiClient.AccountsAPI.GetAccountVideos(context.Background(), name).CategoryOneOf(categoryOneOf).IsLive(isLive).TagsOneOf(tagsOneOf).TagsAllOf(tagsAllOf).LicenceOneOf(licenceOneOf).LanguageOneOf(languageOneOf).AutoTagOneOf(autoTagOneOf).Nsfw(nsfw).IsLocal(isLocal).Include(include).PrivacyOneOf(privacyOneOf).HasHLSFiles(hasHLSFiles).HasWebVideoFiles(hasWebVideoFiles).SkipCount(skipCount).Start(start).Count(count).Sort(sort).ExcludeAlreadyWatched(excludeAlreadyWatched).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.GetAccountVideos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountVideos`: VideoListResponse
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.GetAccountVideos`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The username or handle of the account | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountVideosRequest struct via the builder pattern


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


## GetAccounts

> GetAccounts200Response GetAccounts(ctx).Start(start).Count(count).Sort(sort).Execute()

List accounts

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
	resp, r, err := apiClient.AccountsAPI.GetAccounts(context.Background()).Start(start).Count(count).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.GetAccounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccounts`: GetAccounts200Response
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.GetAccounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **int32** | Offset used to paginate results | 
 **count** | **int32** | Number of items to return | [default to 15]
 **sort** | **string** | Sort column | 

### Return type

[**GetAccounts200Response**](GetAccounts200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

