# \VideoRatesAPI

All URIs are relative to *https://peertube2.cpy.re*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV1VideosIdRatePut**](VideoRatesAPI.md#ApiV1VideosIdRatePut) | **Put** /api/v1/videos/{id}/rate | Like/dislike a video



## ApiV1VideosIdRatePut

> ApiV1VideosIdRatePut(ctx, id).ApiV1VideosIdRatePutRequest(apiV1VideosIdRatePutRequest).XPeertubeVideoPassword(xPeertubeVideoPassword).Execute()

Like/dislike a video

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
	apiV1VideosIdRatePutRequest := *openapiclient.NewApiV1VideosIdRatePutRequest("Rating_example") // ApiV1VideosIdRatePutRequest |  (optional)
	xPeertubeVideoPassword := "xPeertubeVideoPassword_example" // string | Required on password protected video (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VideoRatesAPI.ApiV1VideosIdRatePut(context.Background(), id).ApiV1VideosIdRatePutRequest(apiV1VideosIdRatePutRequest).XPeertubeVideoPassword(xPeertubeVideoPassword).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideoRatesAPI.ApiV1VideosIdRatePut``: %v\n", err)
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

Other parameters are passed through a pointer to a apiApiV1VideosIdRatePutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiV1VideosIdRatePutRequest** | [**ApiV1VideosIdRatePutRequest**](ApiV1VideosIdRatePutRequest.md) |  | 
 **xPeertubeVideoPassword** | **string** | Required on password protected video | 

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

