# \VideoTranscodingAPI

All URIs are relative to *https://peertube2.cpy.re*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateVideoTranscoding**](VideoTranscodingAPI.md#CreateVideoTranscoding) | **Post** /api/v1/videos/{id}/transcoding | Create a transcoding job



## CreateVideoTranscoding

> CreateVideoTranscoding(ctx, id).CreateVideoTranscodingRequest(createVideoTranscodingRequest).Execute()

Create a transcoding job

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
	id := openapiclient._api_v1_videos_ownership__id__accept_post_id_parameter{Int32: new(int32)} // ApiV1VideosOwnershipIdAcceptPostIdParameter | The object id, uuid or short uuid
	createVideoTranscodingRequest := *openapiclient.NewCreateVideoTranscodingRequest("TranscodingType_example") // CreateVideoTranscodingRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VideoTranscodingAPI.CreateVideoTranscoding(context.Background(), id).CreateVideoTranscodingRequest(createVideoTranscodingRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideoTranscodingAPI.CreateVideoTranscoding``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCreateVideoTranscodingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createVideoTranscodingRequest** | [**CreateVideoTranscodingRequest**](CreateVideoTranscodingRequest.md) |  | 

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

