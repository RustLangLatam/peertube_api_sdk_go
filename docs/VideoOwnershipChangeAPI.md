# \VideoOwnershipChangeAPI

All URIs are relative to *https://peertube2.cpy.re*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV1VideosIdGiveOwnershipPost**](VideoOwnershipChangeAPI.md#ApiV1VideosIdGiveOwnershipPost) | **Post** /api/v1/videos/{id}/give-ownership | Request ownership change
[**ApiV1VideosOwnershipGet**](VideoOwnershipChangeAPI.md#ApiV1VideosOwnershipGet) | **Get** /api/v1/videos/ownership | List video ownership changes
[**ApiV1VideosOwnershipIdAcceptPost**](VideoOwnershipChangeAPI.md#ApiV1VideosOwnershipIdAcceptPost) | **Post** /api/v1/videos/ownership/{id}/accept | Accept ownership change request
[**ApiV1VideosOwnershipIdRefusePost**](VideoOwnershipChangeAPI.md#ApiV1VideosOwnershipIdRefusePost) | **Post** /api/v1/videos/ownership/{id}/refuse | Refuse ownership change request



## ApiV1VideosIdGiveOwnershipPost

> ApiV1VideosIdGiveOwnershipPost(ctx, id).Username(username).Execute()

Request ownership change

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
	username := "username_example" // string | 
	id := openapiclient._api_v1_videos_ownership__id__accept_post_id_parameter{Int32: new(int32)} // ApiV1VideosOwnershipIdAcceptPostIdParameter | The object id, uuid or short uuid

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VideoOwnershipChangeAPI.ApiV1VideosIdGiveOwnershipPost(context.Background(), id).Username(username).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideoOwnershipChangeAPI.ApiV1VideosIdGiveOwnershipPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiApiV1VideosIdGiveOwnershipPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **username** | **string** |  | 


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


## ApiV1VideosOwnershipGet

> ApiV1VideosOwnershipGet(ctx).Execute()

List video ownership changes

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
	r, err := apiClient.VideoOwnershipChangeAPI.ApiV1VideosOwnershipGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideoOwnershipChangeAPI.ApiV1VideosOwnershipGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1VideosOwnershipGetRequest struct via the builder pattern


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


## ApiV1VideosOwnershipIdAcceptPost

> ApiV1VideosOwnershipIdAcceptPost(ctx, id).Execute()

Accept ownership change request

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VideoOwnershipChangeAPI.ApiV1VideosOwnershipIdAcceptPost(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideoOwnershipChangeAPI.ApiV1VideosOwnershipIdAcceptPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiApiV1VideosOwnershipIdAcceptPostRequest struct via the builder pattern


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


## ApiV1VideosOwnershipIdRefusePost

> ApiV1VideosOwnershipIdRefusePost(ctx, id).Execute()

Refuse ownership change request

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VideoOwnershipChangeAPI.ApiV1VideosOwnershipIdRefusePost(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideoOwnershipChangeAPI.ApiV1VideosOwnershipIdRefusePost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiApiV1VideosOwnershipIdRefusePostRequest struct via the builder pattern


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

