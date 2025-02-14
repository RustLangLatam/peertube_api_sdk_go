# \MyUserAPI

All URIs are relative to *https://peertube2.cpy.re*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV1UsersMeAvatarDelete**](MyUserAPI.md#ApiV1UsersMeAvatarDelete) | **Delete** /api/v1/users/me/avatar | Delete my avatar
[**ApiV1UsersMeAvatarPickPost**](MyUserAPI.md#ApiV1UsersMeAvatarPickPost) | **Post** /api/v1/users/me/avatar/pick | Update my user avatar
[**ApiV1UsersMeVideoQuotaUsedGet**](MyUserAPI.md#ApiV1UsersMeVideoQuotaUsedGet) | **Get** /api/v1/users/me/video-quota-used | Get my user used quota
[**ApiV1UsersMeVideosGet**](MyUserAPI.md#ApiV1UsersMeVideosGet) | **Get** /api/v1/users/me/videos | List videos of my user
[**ApiV1UsersMeVideosImportsGet**](MyUserAPI.md#ApiV1UsersMeVideosImportsGet) | **Get** /api/v1/users/me/videos/imports | Get video imports of my user
[**ApiV1UsersMeVideosVideoIdRatingGet**](MyUserAPI.md#ApiV1UsersMeVideosVideoIdRatingGet) | **Get** /api/v1/users/me/videos/{videoId}/rating | Get rate of my user for a video
[**GetUserInfo**](MyUserAPI.md#GetUserInfo) | **Get** /api/v1/users/me | Get my user information
[**PutUserInfo**](MyUserAPI.md#PutUserInfo) | **Put** /api/v1/users/me | Update my user information



## ApiV1UsersMeAvatarDelete

> ApiV1UsersMeAvatarDelete(ctx).Execute()

Delete my avatar

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
	r, err := apiClient.MyUserAPI.ApiV1UsersMeAvatarDelete(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MyUserAPI.ApiV1UsersMeAvatarDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1UsersMeAvatarDeleteRequest struct via the builder pattern


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


## ApiV1UsersMeAvatarPickPost

> ApiV1UsersMeAvatarPickPost200Response ApiV1UsersMeAvatarPickPost(ctx).Avatarfile(avatarfile).Execute()

Update my user avatar

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
	avatarfile := os.NewFile(1234, "some_file") // *os.File | The file to upload (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MyUserAPI.ApiV1UsersMeAvatarPickPost(context.Background()).Avatarfile(avatarfile).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MyUserAPI.ApiV1UsersMeAvatarPickPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1UsersMeAvatarPickPost`: ApiV1UsersMeAvatarPickPost200Response
	fmt.Fprintf(os.Stdout, "Response from `MyUserAPI.ApiV1UsersMeAvatarPickPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1UsersMeAvatarPickPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **avatarfile** | ***os.File** | The file to upload | 

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


## ApiV1UsersMeVideoQuotaUsedGet

> ApiV1UsersMeVideoQuotaUsedGet200Response ApiV1UsersMeVideoQuotaUsedGet(ctx).Execute()

Get my user used quota

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
	resp, r, err := apiClient.MyUserAPI.ApiV1UsersMeVideoQuotaUsedGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MyUserAPI.ApiV1UsersMeVideoQuotaUsedGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1UsersMeVideoQuotaUsedGet`: ApiV1UsersMeVideoQuotaUsedGet200Response
	fmt.Fprintf(os.Stdout, "Response from `MyUserAPI.ApiV1UsersMeVideoQuotaUsedGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1UsersMeVideoQuotaUsedGetRequest struct via the builder pattern


### Return type

[**ApiV1UsersMeVideoQuotaUsedGet200Response**](ApiV1UsersMeVideoQuotaUsedGet200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1UsersMeVideosGet

> VideoListResponse ApiV1UsersMeVideosGet(ctx).Start(start).Count(count).Sort(sort).Execute()

List videos of my user

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
	resp, r, err := apiClient.MyUserAPI.ApiV1UsersMeVideosGet(context.Background()).Start(start).Count(count).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MyUserAPI.ApiV1UsersMeVideosGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1UsersMeVideosGet`: VideoListResponse
	fmt.Fprintf(os.Stdout, "Response from `MyUserAPI.ApiV1UsersMeVideosGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1UsersMeVideosGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **int32** | Offset used to paginate results | 
 **count** | **int32** | Number of items to return | [default to 15]
 **sort** | **string** | Sort column | 

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


## ApiV1UsersMeVideosImportsGet

> VideoImportsList ApiV1UsersMeVideosImportsGet(ctx).Start(start).Count(count).Sort(sort).TargetUrl(targetUrl).VideoChannelSyncId(videoChannelSyncId).Search(search).Execute()

Get video imports of my user

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
	targetUrl := "targetUrl_example" // string | Filter on import target URL (optional)
	videoChannelSyncId := float32(8.14) // float32 | Filter on imports created by a specific channel synchronization (optional)
	search := "search_example" // string | Search in video names (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MyUserAPI.ApiV1UsersMeVideosImportsGet(context.Background()).Start(start).Count(count).Sort(sort).TargetUrl(targetUrl).VideoChannelSyncId(videoChannelSyncId).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MyUserAPI.ApiV1UsersMeVideosImportsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1UsersMeVideosImportsGet`: VideoImportsList
	fmt.Fprintf(os.Stdout, "Response from `MyUserAPI.ApiV1UsersMeVideosImportsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1UsersMeVideosImportsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **int32** | Offset used to paginate results | 
 **count** | **int32** | Number of items to return | [default to 15]
 **sort** | **string** | Sort column | 
 **targetUrl** | **string** | Filter on import target URL | 
 **videoChannelSyncId** | **float32** | Filter on imports created by a specific channel synchronization | 
 **search** | **string** | Search in video names | 

### Return type

[**VideoImportsList**](VideoImportsList.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1UsersMeVideosVideoIdRatingGet

> GetMeVideoRating ApiV1UsersMeVideosVideoIdRatingGet(ctx, videoId).Execute()

Get rate of my user for a video

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
	videoId := int32(56) // int32 | The video id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MyUserAPI.ApiV1UsersMeVideosVideoIdRatingGet(context.Background(), videoId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MyUserAPI.ApiV1UsersMeVideosVideoIdRatingGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1UsersMeVideosVideoIdRatingGet`: GetMeVideoRating
	fmt.Fprintf(os.Stdout, "Response from `MyUserAPI.ApiV1UsersMeVideosVideoIdRatingGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**videoId** | **int32** | The video id | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1UsersMeVideosVideoIdRatingGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetMeVideoRating**](GetMeVideoRating.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserInfo

> []User GetUserInfo(ctx).Execute()

Get my user information

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
	resp, r, err := apiClient.MyUserAPI.GetUserInfo(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MyUserAPI.GetUserInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserInfo`: []User
	fmt.Fprintf(os.Stdout, "Response from `MyUserAPI.GetUserInfo`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserInfoRequest struct via the builder pattern


### Return type

[**[]User**](User.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutUserInfo

> PutUserInfo(ctx).UpdateMe(updateMe).Execute()

Update my user information

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
	updateMe := *openapiclient.NewUpdateMe() // UpdateMe | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MyUserAPI.PutUserInfo(context.Background()).UpdateMe(updateMe).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MyUserAPI.PutUserInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutUserInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateMe** | [**UpdateMe**](UpdateMe.md) |  | 

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

