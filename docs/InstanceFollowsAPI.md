# \InstanceFollowsAPI

All URIs are relative to *https://peertube2.cpy.re*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV1ServerFollowersGet**](InstanceFollowsAPI.md#ApiV1ServerFollowersGet) | **Get** /api/v1/server/followers | List instances following the server
[**ApiV1ServerFollowersNameWithHostAcceptPost**](InstanceFollowsAPI.md#ApiV1ServerFollowersNameWithHostAcceptPost) | **Post** /api/v1/server/followers/{nameWithHost}/accept | Accept a pending follower to your server
[**ApiV1ServerFollowersNameWithHostDelete**](InstanceFollowsAPI.md#ApiV1ServerFollowersNameWithHostDelete) | **Delete** /api/v1/server/followers/{nameWithHost} | Remove or reject a follower to your server
[**ApiV1ServerFollowersNameWithHostRejectPost**](InstanceFollowsAPI.md#ApiV1ServerFollowersNameWithHostRejectPost) | **Post** /api/v1/server/followers/{nameWithHost}/reject | Reject a pending follower to your server
[**ApiV1ServerFollowingGet**](InstanceFollowsAPI.md#ApiV1ServerFollowingGet) | **Get** /api/v1/server/following | List instances followed by the server
[**ApiV1ServerFollowingHostOrHandleDelete**](InstanceFollowsAPI.md#ApiV1ServerFollowingHostOrHandleDelete) | **Delete** /api/v1/server/following/{hostOrHandle} | Unfollow an actor (PeerTube instance, channel or account)
[**ApiV1ServerFollowingPost**](InstanceFollowsAPI.md#ApiV1ServerFollowingPost) | **Post** /api/v1/server/following | Follow a list of actors (PeerTube instance, channel or account)



## ApiV1ServerFollowersGet

> GetAccountFollowers200Response ApiV1ServerFollowersGet(ctx).State(state).ActorType(actorType).Start(start).Count(count).Sort(sort).Execute()

List instances following the server

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
	state := "state_example" // string |  (optional)
	actorType := "actorType_example" // string |  (optional)
	start := int32(56) // int32 | Offset used to paginate results (optional)
	count := int32(56) // int32 | Number of items to return (optional) (default to 15)
	sort := "-createdAt" // string | Sort column (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstanceFollowsAPI.ApiV1ServerFollowersGet(context.Background()).State(state).ActorType(actorType).Start(start).Count(count).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceFollowsAPI.ApiV1ServerFollowersGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1ServerFollowersGet`: GetAccountFollowers200Response
	fmt.Fprintf(os.Stdout, "Response from `InstanceFollowsAPI.ApiV1ServerFollowersGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1ServerFollowersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **state** | **string** |  | 
 **actorType** | **string** |  | 
 **start** | **int32** | Offset used to paginate results | 
 **count** | **int32** | Number of items to return | [default to 15]
 **sort** | **string** | Sort column | 

### Return type

[**GetAccountFollowers200Response**](GetAccountFollowers200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1ServerFollowersNameWithHostAcceptPost

> ApiV1ServerFollowersNameWithHostAcceptPost(ctx, nameWithHost).Execute()

Accept a pending follower to your server

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
	nameWithHost := "nameWithHost_example" // string | The remote actor handle to remove from your followers

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InstanceFollowsAPI.ApiV1ServerFollowersNameWithHostAcceptPost(context.Background(), nameWithHost).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceFollowsAPI.ApiV1ServerFollowersNameWithHostAcceptPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameWithHost** | **string** | The remote actor handle to remove from your followers | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1ServerFollowersNameWithHostAcceptPostRequest struct via the builder pattern


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


## ApiV1ServerFollowersNameWithHostDelete

> ApiV1ServerFollowersNameWithHostDelete(ctx, nameWithHost).Execute()

Remove or reject a follower to your server

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
	nameWithHost := "nameWithHost_example" // string | The remote actor handle to remove from your followers

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InstanceFollowsAPI.ApiV1ServerFollowersNameWithHostDelete(context.Background(), nameWithHost).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceFollowsAPI.ApiV1ServerFollowersNameWithHostDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameWithHost** | **string** | The remote actor handle to remove from your followers | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1ServerFollowersNameWithHostDeleteRequest struct via the builder pattern


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


## ApiV1ServerFollowersNameWithHostRejectPost

> ApiV1ServerFollowersNameWithHostRejectPost(ctx, nameWithHost).Execute()

Reject a pending follower to your server

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
	nameWithHost := "nameWithHost_example" // string | The remote actor handle to remove from your followers

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InstanceFollowsAPI.ApiV1ServerFollowersNameWithHostRejectPost(context.Background(), nameWithHost).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceFollowsAPI.ApiV1ServerFollowersNameWithHostRejectPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nameWithHost** | **string** | The remote actor handle to remove from your followers | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1ServerFollowersNameWithHostRejectPostRequest struct via the builder pattern


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


## ApiV1ServerFollowingGet

> GetAccountFollowers200Response ApiV1ServerFollowingGet(ctx).State(state).ActorType(actorType).Start(start).Count(count).Sort(sort).Execute()

List instances followed by the server

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
	state := "state_example" // string |  (optional)
	actorType := "actorType_example" // string |  (optional)
	start := int32(56) // int32 | Offset used to paginate results (optional)
	count := int32(56) // int32 | Number of items to return (optional) (default to 15)
	sort := "-createdAt" // string | Sort column (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstanceFollowsAPI.ApiV1ServerFollowingGet(context.Background()).State(state).ActorType(actorType).Start(start).Count(count).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceFollowsAPI.ApiV1ServerFollowingGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1ServerFollowingGet`: GetAccountFollowers200Response
	fmt.Fprintf(os.Stdout, "Response from `InstanceFollowsAPI.ApiV1ServerFollowingGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1ServerFollowingGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **state** | **string** |  | 
 **actorType** | **string** |  | 
 **start** | **int32** | Offset used to paginate results | 
 **count** | **int32** | Number of items to return | [default to 15]
 **sort** | **string** | Sort column | 

### Return type

[**GetAccountFollowers200Response**](GetAccountFollowers200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1ServerFollowingHostOrHandleDelete

> ApiV1ServerFollowingHostOrHandleDelete(ctx, hostOrHandle).Execute()

Unfollow an actor (PeerTube instance, channel or account)

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
	hostOrHandle := "hostOrHandle_example" // string | The hostOrHandle to unfollow

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InstanceFollowsAPI.ApiV1ServerFollowingHostOrHandleDelete(context.Background(), hostOrHandle).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceFollowsAPI.ApiV1ServerFollowingHostOrHandleDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostOrHandle** | **string** | The hostOrHandle to unfollow | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1ServerFollowingHostOrHandleDeleteRequest struct via the builder pattern


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


## ApiV1ServerFollowingPost

> ApiV1ServerFollowingPost(ctx).ApiV1ServerFollowingPostRequest(apiV1ServerFollowingPostRequest).Execute()

Follow a list of actors (PeerTube instance, channel or account)

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
	apiV1ServerFollowingPostRequest := *openapiclient.NewApiV1ServerFollowingPostRequest() // ApiV1ServerFollowingPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InstanceFollowsAPI.ApiV1ServerFollowingPost(context.Background()).ApiV1ServerFollowingPostRequest(apiV1ServerFollowingPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceFollowsAPI.ApiV1ServerFollowingPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1ServerFollowingPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiV1ServerFollowingPostRequest** | [**ApiV1ServerFollowingPostRequest**](ApiV1ServerFollowingPostRequest.md) |  | 

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

