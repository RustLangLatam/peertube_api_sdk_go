# \ChannelsSyncAPI

All URIs are relative to *https://peertube2.cpy.re*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddVideoChannelSync**](ChannelsSyncAPI.md#AddVideoChannelSync) | **Post** /api/v1/video-channel-syncs | Create a synchronization for a video channel
[**DelVideoChannelSync**](ChannelsSyncAPI.md#DelVideoChannelSync) | **Delete** /api/v1/video-channel-syncs/{channelSyncId} | Delete a video channel synchronization
[**TriggerVideoChannelSync**](ChannelsSyncAPI.md#TriggerVideoChannelSync) | **Post** /api/v1/video-channel-syncs/{channelSyncId}/sync | Triggers the channel synchronization job, fetching all the videos from the remote channel



## AddVideoChannelSync

> AddVideoChannelSync200Response AddVideoChannelSync(ctx).VideoChannelSyncCreate(videoChannelSyncCreate).Execute()

Create a synchronization for a video channel

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
	videoChannelSyncCreate := *openapiclient.NewVideoChannelSyncCreate() // VideoChannelSyncCreate |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChannelsSyncAPI.AddVideoChannelSync(context.Background()).VideoChannelSyncCreate(videoChannelSyncCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelsSyncAPI.AddVideoChannelSync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddVideoChannelSync`: AddVideoChannelSync200Response
	fmt.Fprintf(os.Stdout, "Response from `ChannelsSyncAPI.AddVideoChannelSync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddVideoChannelSyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **videoChannelSyncCreate** | [**VideoChannelSyncCreate**](VideoChannelSyncCreate.md) |  | 

### Return type

[**AddVideoChannelSync200Response**](AddVideoChannelSync200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DelVideoChannelSync

> DelVideoChannelSync(ctx, channelSyncId).Execute()

Delete a video channel synchronization

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
	channelSyncId := int32(56) // int32 | Channel Sync id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ChannelsSyncAPI.DelVideoChannelSync(context.Background(), channelSyncId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelsSyncAPI.DelVideoChannelSync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelSyncId** | **int32** | Channel Sync id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDelVideoChannelSyncRequest struct via the builder pattern


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


## TriggerVideoChannelSync

> TriggerVideoChannelSync(ctx, channelSyncId).Execute()

Triggers the channel synchronization job, fetching all the videos from the remote channel

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
	channelSyncId := int32(56) // int32 | Channel Sync id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ChannelsSyncAPI.TriggerVideoChannelSync(context.Background(), channelSyncId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelsSyncAPI.TriggerVideoChannelSync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelSyncId** | **int32** | Channel Sync id | 

### Other Parameters

Other parameters are passed through a pointer to a apiTriggerVideoChannelSyncRequest struct via the builder pattern


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

