# \VideoMirroringAPI

All URIs are relative to *https://peertube2.cpy.re*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DelMirroredVideo**](VideoMirroringAPI.md#DelMirroredVideo) | **Delete** /api/v1/server/redundancy/videos/{redundancyId} | Delete a mirror done on a video
[**GetMirroredVideos**](VideoMirroringAPI.md#GetMirroredVideos) | **Get** /api/v1/server/redundancy/videos | List videos being mirrored
[**PutMirroredVideo**](VideoMirroringAPI.md#PutMirroredVideo) | **Post** /api/v1/server/redundancy/videos | Mirror a video



## DelMirroredVideo

> DelMirroredVideo(ctx, redundancyId).Execute()

Delete a mirror done on a video

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
	redundancyId := "redundancyId_example" // string | id of an existing redundancy on a video

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VideoMirroringAPI.DelMirroredVideo(context.Background(), redundancyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideoMirroringAPI.DelMirroredVideo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**redundancyId** | **string** | id of an existing redundancy on a video | 

### Other Parameters

Other parameters are passed through a pointer to a apiDelMirroredVideoRequest struct via the builder pattern


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


## GetMirroredVideos

> []VideoRedundancy GetMirroredVideos(ctx).Target(target).Start(start).Count(count).Sort(sort).Execute()

List videos being mirrored

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
	target := "target_example" // string | direction of the mirror
	start := int32(56) // int32 | Offset used to paginate results (optional)
	count := int32(56) // int32 | Number of items to return (optional) (default to 15)
	sort := "sort_example" // string | Sort abuses by criteria (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VideoMirroringAPI.GetMirroredVideos(context.Background()).Target(target).Start(start).Count(count).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideoMirroringAPI.GetMirroredVideos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMirroredVideos`: []VideoRedundancy
	fmt.Fprintf(os.Stdout, "Response from `VideoMirroringAPI.GetMirroredVideos`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMirroredVideosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **target** | **string** | direction of the mirror | 
 **start** | **int32** | Offset used to paginate results | 
 **count** | **int32** | Number of items to return | [default to 15]
 **sort** | **string** | Sort abuses by criteria | 

### Return type

[**[]VideoRedundancy**](VideoRedundancy.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutMirroredVideo

> PutMirroredVideo(ctx).PutMirroredVideoRequest(putMirroredVideoRequest).Execute()

Mirror a video

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
	putMirroredVideoRequest := *openapiclient.NewPutMirroredVideoRequest(int32(42)) // PutMirroredVideoRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VideoMirroringAPI.PutMirroredVideo(context.Background()).PutMirroredVideoRequest(putMirroredVideoRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideoMirroringAPI.PutMirroredVideo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutMirroredVideoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **putMirroredVideoRequest** | [**PutMirroredVideoRequest**](PutMirroredVideoRequest.md) |  | 

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

