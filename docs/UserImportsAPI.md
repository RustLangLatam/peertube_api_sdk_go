# \UserImportsAPI

All URIs are relative to *https://peertube2.cpy.re*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLatestUserImport**](UserImportsAPI.md#GetLatestUserImport) | **Get** /api/v1/users/{userId}/imports/latest | Get latest user import



## GetLatestUserImport

> GetLatestUserImport200Response GetLatestUserImport(ctx, userId).Execute()

Get latest user import



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
	userId := int32(56) // int32 | User id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserImportsAPI.GetLatestUserImport(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserImportsAPI.GetLatestUserImport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLatestUserImport`: GetLatestUserImport200Response
	fmt.Fprintf(os.Stdout, "Response from `UserImportsAPI.GetLatestUserImport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int32** | User id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLatestUserImportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetLatestUserImport200Response**](GetLatestUserImport200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

