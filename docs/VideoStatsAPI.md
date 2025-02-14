# \VideoStatsAPI

All URIs are relative to *https://peertube2.cpy.re*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV1VideosIdStatsOverallGet**](VideoStatsAPI.md#ApiV1VideosIdStatsOverallGet) | **Get** /api/v1/videos/{id}/stats/overall | Get overall stats of a video
[**ApiV1VideosIdStatsRetentionGet**](VideoStatsAPI.md#ApiV1VideosIdStatsRetentionGet) | **Get** /api/v1/videos/{id}/stats/retention | Get retention stats of a video
[**ApiV1VideosIdStatsTimeseriesMetricGet**](VideoStatsAPI.md#ApiV1VideosIdStatsTimeseriesMetricGet) | **Get** /api/v1/videos/{id}/stats/timeseries/{metric} | Get timeserie stats of a video



## ApiV1VideosIdStatsOverallGet

> VideoStatsOverall ApiV1VideosIdStatsOverallGet(ctx, id).StartDate(startDate).EndDate(endDate).Execute()

Get overall stats of a video

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/RustLangLatam/peertube_api_sdk"
)

func main() {
	id := openapiclient._api_v1_videos_ownership__id__accept_post_id_parameter{Int32: new(int32)} // ApiV1VideosOwnershipIdAcceptPostIdParameter | The object id, uuid or short uuid
	startDate := time.Now() // time.Time | Filter stats by start date (optional)
	endDate := time.Now() // time.Time | Filter stats by end date (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VideoStatsAPI.ApiV1VideosIdStatsOverallGet(context.Background(), id).StartDate(startDate).EndDate(endDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideoStatsAPI.ApiV1VideosIdStatsOverallGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1VideosIdStatsOverallGet`: VideoStatsOverall
	fmt.Fprintf(os.Stdout, "Response from `VideoStatsAPI.ApiV1VideosIdStatsOverallGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**ApiV1VideosOwnershipIdAcceptPostIdParameter**](.md) | The object id, uuid or short uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1VideosIdStatsOverallGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startDate** | **time.Time** | Filter stats by start date | 
 **endDate** | **time.Time** | Filter stats by end date | 

### Return type

[**VideoStatsOverall**](VideoStatsOverall.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1VideosIdStatsRetentionGet

> VideoStatsRetention ApiV1VideosIdStatsRetentionGet(ctx, id).Execute()

Get retention stats of a video

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
	resp, r, err := apiClient.VideoStatsAPI.ApiV1VideosIdStatsRetentionGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideoStatsAPI.ApiV1VideosIdStatsRetentionGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1VideosIdStatsRetentionGet`: VideoStatsRetention
	fmt.Fprintf(os.Stdout, "Response from `VideoStatsAPI.ApiV1VideosIdStatsRetentionGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**ApiV1VideosOwnershipIdAcceptPostIdParameter**](.md) | The object id, uuid or short uuid | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1VideosIdStatsRetentionGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VideoStatsRetention**](VideoStatsRetention.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1VideosIdStatsTimeseriesMetricGet

> VideoStatsTimeserie ApiV1VideosIdStatsTimeseriesMetricGet(ctx, id, metric).StartDate(startDate).EndDate(endDate).Execute()

Get timeserie stats of a video

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/RustLangLatam/peertube_api_sdk"
)

func main() {
	id := openapiclient._api_v1_videos_ownership__id__accept_post_id_parameter{Int32: new(int32)} // ApiV1VideosOwnershipIdAcceptPostIdParameter | The object id, uuid or short uuid
	metric := "metric_example" // string | The metric to get
	startDate := time.Now() // time.Time | Filter stats by start date (optional)
	endDate := time.Now() // time.Time | Filter stats by end date (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VideoStatsAPI.ApiV1VideosIdStatsTimeseriesMetricGet(context.Background(), id, metric).StartDate(startDate).EndDate(endDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideoStatsAPI.ApiV1VideosIdStatsTimeseriesMetricGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1VideosIdStatsTimeseriesMetricGet`: VideoStatsTimeserie
	fmt.Fprintf(os.Stdout, "Response from `VideoStatsAPI.ApiV1VideosIdStatsTimeseriesMetricGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**ApiV1VideosOwnershipIdAcceptPostIdParameter**](.md) | The object id, uuid or short uuid | 
**metric** | **string** | The metric to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1VideosIdStatsTimeseriesMetricGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **startDate** | **time.Time** | Filter stats by start date | 
 **endDate** | **time.Time** | Filter stats by end date | 

### Return type

[**VideoStatsTimeserie**](VideoStatsTimeserie.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

