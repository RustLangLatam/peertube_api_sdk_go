# \InstanceRedundancyAPI

All URIs are relative to *https://peertube2.cpy.re*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV1ServerRedundancyHostPut**](InstanceRedundancyAPI.md#ApiV1ServerRedundancyHostPut) | **Put** /api/v1/server/redundancy/{host} | Update a server redundancy policy



## ApiV1ServerRedundancyHostPut

> ApiV1ServerRedundancyHostPut(ctx, host).ApiV1ServerRedundancyHostPutRequest(apiV1ServerRedundancyHostPutRequest).Execute()

Update a server redundancy policy

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
	host := "host_example" // string | server domain to mirror
	apiV1ServerRedundancyHostPutRequest := *openapiclient.NewApiV1ServerRedundancyHostPutRequest(false) // ApiV1ServerRedundancyHostPutRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InstanceRedundancyAPI.ApiV1ServerRedundancyHostPut(context.Background(), host).ApiV1ServerRedundancyHostPutRequest(apiV1ServerRedundancyHostPutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceRedundancyAPI.ApiV1ServerRedundancyHostPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**host** | **string** | server domain to mirror | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1ServerRedundancyHostPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiV1ServerRedundancyHostPutRequest** | [**ApiV1ServerRedundancyHostPutRequest**](ApiV1ServerRedundancyHostPutRequest.md) |  | 

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

