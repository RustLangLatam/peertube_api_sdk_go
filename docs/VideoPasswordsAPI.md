# \VideoPasswordsAPI

All URIs are relative to *https://peertube2.cpy.re*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV1VideosIdPasswordsGet**](VideoPasswordsAPI.md#ApiV1VideosIdPasswordsGet) | **Get** /api/v1/videos/{id}/passwords | List video passwords
[**ApiV1VideosIdPasswordsPut**](VideoPasswordsAPI.md#ApiV1VideosIdPasswordsPut) | **Put** /api/v1/videos/{id}/passwords | Update video passwords
[**ApiV1VideosIdPasswordsVideoPasswordIdDelete**](VideoPasswordsAPI.md#ApiV1VideosIdPasswordsVideoPasswordIdDelete) | **Delete** /api/v1/videos/{id}/passwords/{videoPasswordId} | Delete a video password



## ApiV1VideosIdPasswordsGet

> VideoPasswordList ApiV1VideosIdPasswordsGet(ctx, id).Start(start).Count(count).Sort(sort).Execute()

List video passwords



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
	id := openapiclient._api_v1_videos_ownership__id__accept_post_id_parameter{Int32: new(int32)} // ApiV1VideosOwnershipIdAcceptPostIdParameter | The object id, uuid or short uuid
	start := int32(56) // int32 | Offset used to paginate results (optional)
	count := int32(56) // int32 | Number of items to return (optional) (default to 15)
	sort := "-createdAt" // string | Sort column (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VideoPasswordsAPI.ApiV1VideosIdPasswordsGet(context.Background(), id).Start(start).Count(count).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideoPasswordsAPI.ApiV1VideosIdPasswordsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1VideosIdPasswordsGet`: VideoPasswordList
	fmt.Fprintf(os.Stdout, "Response from `VideoPasswordsAPI.ApiV1VideosIdPasswordsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**ApiV1VideosOwnershipIdAcceptPostIdParameter**](.md) | The object id, uuid or short uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1VideosIdPasswordsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **start** | **int32** | Offset used to paginate results | 
 **count** | **int32** | Number of items to return | [default to 15]
 **sort** | **string** | Sort column | 

### Return type

[**VideoPasswordList**](VideoPasswordList.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1VideosIdPasswordsPut

> ApiV1VideosIdPasswordsPut(ctx, id).ApiV1VideosIdPasswordsPutRequest(apiV1VideosIdPasswordsPutRequest).Execute()

Update video passwords



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
	id := openapiclient._api_v1_videos_ownership__id__accept_post_id_parameter{Int32: new(int32)} // ApiV1VideosOwnershipIdAcceptPostIdParameter | The object id, uuid or short uuid
	apiV1VideosIdPasswordsPutRequest := *openapiclient.NewApiV1VideosIdPasswordsPutRequest() // ApiV1VideosIdPasswordsPutRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VideoPasswordsAPI.ApiV1VideosIdPasswordsPut(context.Background(), id).ApiV1VideosIdPasswordsPutRequest(apiV1VideosIdPasswordsPutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideoPasswordsAPI.ApiV1VideosIdPasswordsPut``: %v\n", err)
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

Other parameters are passed through a pointer to a apiApiV1VideosIdPasswordsPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiV1VideosIdPasswordsPutRequest** | [**ApiV1VideosIdPasswordsPutRequest**](ApiV1VideosIdPasswordsPutRequest.md) |  | 

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


## ApiV1VideosIdPasswordsVideoPasswordIdDelete

> ApiV1VideosIdPasswordsVideoPasswordIdDelete(ctx, id, videoPasswordId).Execute()

Delete a video password



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
	id := openapiclient._api_v1_videos_ownership__id__accept_post_id_parameter{Int32: new(int32)} // ApiV1VideosOwnershipIdAcceptPostIdParameter | The object id, uuid or short uuid
	videoPasswordId := int32(56) // int32 | The video password id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VideoPasswordsAPI.ApiV1VideosIdPasswordsVideoPasswordIdDelete(context.Background(), id, videoPasswordId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideoPasswordsAPI.ApiV1VideosIdPasswordsVideoPasswordIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**ApiV1VideosOwnershipIdAcceptPostIdParameter**](.md) | The object id, uuid or short uuid | 
**videoPasswordId** | **int32** | The video password id | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1VideosIdPasswordsVideoPasswordIdDeleteRequest struct via the builder pattern


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

