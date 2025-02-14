# \StatsAPI

All URIs are relative to *https://peertube2.cpy.re*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV1MetricsPlaybackPost**](StatsAPI.md#ApiV1MetricsPlaybackPost) | **Post** /api/v1/metrics/playback | Create playback metrics
[**GetInstanceStats**](StatsAPI.md#GetInstanceStats) | **Get** /api/v1/server/stats | Get instance stats



## ApiV1MetricsPlaybackPost

> ApiV1MetricsPlaybackPost(ctx).PlaybackMetricCreate(playbackMetricCreate).Execute()

Create playback metrics



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
	playbackMetricCreate := *openapiclient.NewPlaybackMetricCreate("PlayerMode_example", false, float32(123), float32(123), float32(123), float32(123), float32(123), openapiclient._api_v1_videos_ownership__id__accept_post_id_parameter{Int32: new(int32)}) // PlaybackMetricCreate |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StatsAPI.ApiV1MetricsPlaybackPost(context.Background()).PlaybackMetricCreate(playbackMetricCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatsAPI.ApiV1MetricsPlaybackPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1MetricsPlaybackPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **playbackMetricCreate** | [**PlaybackMetricCreate**](PlaybackMetricCreate.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInstanceStats

> ServerStats GetInstanceStats(ctx).Execute()

Get instance stats



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
	resp, r, err := apiClient.StatsAPI.GetInstanceStats(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatsAPI.GetInstanceStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInstanceStats`: ServerStats
	fmt.Fprintf(os.Stdout, "Response from `StatsAPI.GetInstanceStats`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetInstanceStatsRequest struct via the builder pattern


### Return type

[**ServerStats**](ServerStats.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

