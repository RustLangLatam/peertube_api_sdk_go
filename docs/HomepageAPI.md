# \HomepageAPI

All URIs are relative to *https://peertube2.cpy.re*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV1CustomPagesHomepageInstanceGet**](HomepageAPI.md#ApiV1CustomPagesHomepageInstanceGet) | **Get** /api/v1/custom-pages/homepage/instance | Get instance custom homepage
[**ApiV1CustomPagesHomepageInstancePut**](HomepageAPI.md#ApiV1CustomPagesHomepageInstancePut) | **Put** /api/v1/custom-pages/homepage/instance | Set instance custom homepage



## ApiV1CustomPagesHomepageInstanceGet

> CustomHomepage ApiV1CustomPagesHomepageInstanceGet(ctx).Execute()

Get instance custom homepage

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
	resp, r, err := apiClient.HomepageAPI.ApiV1CustomPagesHomepageInstanceGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HomepageAPI.ApiV1CustomPagesHomepageInstanceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1CustomPagesHomepageInstanceGet`: CustomHomepage
	fmt.Fprintf(os.Stdout, "Response from `HomepageAPI.ApiV1CustomPagesHomepageInstanceGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1CustomPagesHomepageInstanceGetRequest struct via the builder pattern


### Return type

[**CustomHomepage**](CustomHomepage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1CustomPagesHomepageInstancePut

> ApiV1CustomPagesHomepageInstancePut(ctx).ApiV1CustomPagesHomepageInstancePutRequest(apiV1CustomPagesHomepageInstancePutRequest).Execute()

Set instance custom homepage

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
	apiV1CustomPagesHomepageInstancePutRequest := *openapiclient.NewApiV1CustomPagesHomepageInstancePutRequest() // ApiV1CustomPagesHomepageInstancePutRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.HomepageAPI.ApiV1CustomPagesHomepageInstancePut(context.Background()).ApiV1CustomPagesHomepageInstancePutRequest(apiV1CustomPagesHomepageInstancePutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HomepageAPI.ApiV1CustomPagesHomepageInstancePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1CustomPagesHomepageInstancePutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiV1CustomPagesHomepageInstancePutRequest** | [**ApiV1CustomPagesHomepageInstancePutRequest**](ApiV1CustomPagesHomepageInstancePutRequest.md) |  | 

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

