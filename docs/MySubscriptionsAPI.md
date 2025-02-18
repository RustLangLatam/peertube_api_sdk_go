# \MySubscriptionsAPI

All URIs are relative to *https://peertube2.cpy.re*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV1UsersMeSubscriptionsExistGet**](MySubscriptionsAPI.md#ApiV1UsersMeSubscriptionsExistGet) | **Get** /api/v1/users/me/subscriptions/exist | Get if subscriptions exist for my user
[**ApiV1UsersMeSubscriptionsGet**](MySubscriptionsAPI.md#ApiV1UsersMeSubscriptionsGet) | **Get** /api/v1/users/me/subscriptions | Get my user subscriptions
[**ApiV1UsersMeSubscriptionsPost**](MySubscriptionsAPI.md#ApiV1UsersMeSubscriptionsPost) | **Post** /api/v1/users/me/subscriptions | Add subscription to my user
[**ApiV1UsersMeSubscriptionsSubscriptionHandleDelete**](MySubscriptionsAPI.md#ApiV1UsersMeSubscriptionsSubscriptionHandleDelete) | **Delete** /api/v1/users/me/subscriptions/{subscriptionHandle} | Delete subscription of my user
[**ApiV1UsersMeSubscriptionsSubscriptionHandleGet**](MySubscriptionsAPI.md#ApiV1UsersMeSubscriptionsSubscriptionHandleGet) | **Get** /api/v1/users/me/subscriptions/{subscriptionHandle} | Get subscription of my user
[**ApiV1UsersMeSubscriptionsVideosGet**](MySubscriptionsAPI.md#ApiV1UsersMeSubscriptionsVideosGet) | **Get** /api/v1/users/me/subscriptions/videos | List videos of subscriptions of my user



## ApiV1UsersMeSubscriptionsExistGet

> map[string]interface{} ApiV1UsersMeSubscriptionsExistGet(ctx).Uris(uris).Execute()

Get if subscriptions exist for my user

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
	uris := []string{"Inner_example"} // []string | list of uris to check if each is part of the user subscriptions

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MySubscriptionsAPI.ApiV1UsersMeSubscriptionsExistGet(context.Background()).Uris(uris).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MySubscriptionsAPI.ApiV1UsersMeSubscriptionsExistGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1UsersMeSubscriptionsExistGet`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `MySubscriptionsAPI.ApiV1UsersMeSubscriptionsExistGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1UsersMeSubscriptionsExistGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uris** | **[]string** | list of uris to check if each is part of the user subscriptions | 

### Return type

**map[string]interface{}**

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1UsersMeSubscriptionsGet

> VideoChannelList ApiV1UsersMeSubscriptionsGet(ctx).Start(start).Count(count).Sort(sort).Execute()

Get my user subscriptions

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
	start := int32(56) // int32 | Offset used to paginate results (optional)
	count := int32(56) // int32 | Number of items to return (optional) (default to 15)
	sort := "-createdAt" // string | Sort column (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MySubscriptionsAPI.ApiV1UsersMeSubscriptionsGet(context.Background()).Start(start).Count(count).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MySubscriptionsAPI.ApiV1UsersMeSubscriptionsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1UsersMeSubscriptionsGet`: VideoChannelList
	fmt.Fprintf(os.Stdout, "Response from `MySubscriptionsAPI.ApiV1UsersMeSubscriptionsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1UsersMeSubscriptionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **int32** | Offset used to paginate results | 
 **count** | **int32** | Number of items to return | [default to 15]
 **sort** | **string** | Sort column | 

### Return type

[**VideoChannelList**](VideoChannelList.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1UsersMeSubscriptionsPost

> ApiV1UsersMeSubscriptionsPost(ctx).ApiV1UsersMeSubscriptionsPostRequest(apiV1UsersMeSubscriptionsPostRequest).Execute()

Add subscription to my user

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
	apiV1UsersMeSubscriptionsPostRequest := *openapiclient.NewApiV1UsersMeSubscriptionsPostRequest("Uri_example") // ApiV1UsersMeSubscriptionsPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MySubscriptionsAPI.ApiV1UsersMeSubscriptionsPost(context.Background()).ApiV1UsersMeSubscriptionsPostRequest(apiV1UsersMeSubscriptionsPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MySubscriptionsAPI.ApiV1UsersMeSubscriptionsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1UsersMeSubscriptionsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiV1UsersMeSubscriptionsPostRequest** | [**ApiV1UsersMeSubscriptionsPostRequest**](ApiV1UsersMeSubscriptionsPostRequest.md) |  | 

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


## ApiV1UsersMeSubscriptionsSubscriptionHandleDelete

> ApiV1UsersMeSubscriptionsSubscriptionHandleDelete(ctx, subscriptionHandle).Execute()

Delete subscription of my user

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
	subscriptionHandle := "my_username | my_username@example.com" // string | The subscription handle

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MySubscriptionsAPI.ApiV1UsersMeSubscriptionsSubscriptionHandleDelete(context.Background(), subscriptionHandle).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MySubscriptionsAPI.ApiV1UsersMeSubscriptionsSubscriptionHandleDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionHandle** | **string** | The subscription handle | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1UsersMeSubscriptionsSubscriptionHandleDeleteRequest struct via the builder pattern


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


## ApiV1UsersMeSubscriptionsSubscriptionHandleGet

> VideoChannel ApiV1UsersMeSubscriptionsSubscriptionHandleGet(ctx, subscriptionHandle).Execute()

Get subscription of my user

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
	subscriptionHandle := "my_username | my_username@example.com" // string | The subscription handle

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MySubscriptionsAPI.ApiV1UsersMeSubscriptionsSubscriptionHandleGet(context.Background(), subscriptionHandle).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MySubscriptionsAPI.ApiV1UsersMeSubscriptionsSubscriptionHandleGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1UsersMeSubscriptionsSubscriptionHandleGet`: VideoChannel
	fmt.Fprintf(os.Stdout, "Response from `MySubscriptionsAPI.ApiV1UsersMeSubscriptionsSubscriptionHandleGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionHandle** | **string** | The subscription handle | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1UsersMeSubscriptionsSubscriptionHandleGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VideoChannel**](VideoChannel.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1UsersMeSubscriptionsVideosGet

> VideoListResponse ApiV1UsersMeSubscriptionsVideosGet(ctx).CategoryOneOf(categoryOneOf).IsLive(isLive).TagsOneOf(tagsOneOf).TagsAllOf(tagsAllOf).LicenceOneOf(licenceOneOf).LanguageOneOf(languageOneOf).AutoTagOneOf(autoTagOneOf).Nsfw(nsfw).IsLocal(isLocal).Include(include).PrivacyOneOf(privacyOneOf).HasHLSFiles(hasHLSFiles).HasWebVideoFiles(hasWebVideoFiles).SkipCount(skipCount).Start(start).Count(count).Sort(sort).ExcludeAlreadyWatched(excludeAlreadyWatched).Search(search).Execute()

List videos of subscriptions of my user

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
	resp, r, err := apiClient.MySubscriptionsAPI.ApiV1UsersMeSubscriptionsVideosGet(context.Background()).CategoryOneOf(categoryOneOf).IsLive(isLive).TagsOneOf(tagsOneOf).TagsAllOf(tagsAllOf).LicenceOneOf(licenceOneOf).LanguageOneOf(languageOneOf).AutoTagOneOf(autoTagOneOf).Nsfw(nsfw).IsLocal(isLocal).Include(include).PrivacyOneOf(privacyOneOf).HasHLSFiles(hasHLSFiles).HasWebVideoFiles(hasWebVideoFiles).SkipCount(skipCount).Start(start).Count(count).Sort(sort).ExcludeAlreadyWatched(excludeAlreadyWatched).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MySubscriptionsAPI.ApiV1UsersMeSubscriptionsVideosGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1UsersMeSubscriptionsVideosGet`: VideoListResponse
	fmt.Fprintf(os.Stdout, "Response from `MySubscriptionsAPI.ApiV1UsersMeSubscriptionsVideosGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1UsersMeSubscriptionsVideosGetRequest struct via the builder pattern


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

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

