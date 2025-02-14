# \UserExportsAPI

All URIs are relative to *https://peertube2.cpy.re*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteUserExport**](UserExportsAPI.md#DeleteUserExport) | **Delete** /api/v1/users/{userId}/exports/{id} | Delete a user export
[**ListUserExports**](UserExportsAPI.md#ListUserExports) | **Get** /api/v1/users/{userId}/exports | List user exports
[**RequestUserExport**](UserExportsAPI.md#RequestUserExport) | **Post** /api/v1/users/{userId}/exports/request | Request user export



## DeleteUserExport

> DeleteUserExport(ctx, userId, id).Execute()

Delete a user export



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
	userId := int32(56) // int32 | User id
	id := int32(56) // int32 | Entity id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UserExportsAPI.DeleteUserExport(context.Background(), userId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserExportsAPI.DeleteUserExport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int32** | User id | 
**id** | **int32** | Entity id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserExportRequest struct via the builder pattern


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


## ListUserExports

> ListUserExports200Response ListUserExports(ctx, userId).Execute()

List user exports



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
	userId := int32(56) // int32 | User id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserExportsAPI.ListUserExports(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserExportsAPI.ListUserExports``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListUserExports`: ListUserExports200Response
	fmt.Fprintf(os.Stdout, "Response from `UserExportsAPI.ListUserExports`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int32** | User id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListUserExportsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListUserExports200Response**](ListUserExports200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RequestUserExport

> RequestUserExport200Response RequestUserExport(ctx, userId).RequestUserExportRequest(requestUserExportRequest).Execute()

Request user export



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
	userId := int32(56) // int32 | User id
	requestUserExportRequest := *openapiclient.NewRequestUserExportRequest() // RequestUserExportRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserExportsAPI.RequestUserExport(context.Background(), userId).RequestUserExportRequest(requestUserExportRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserExportsAPI.RequestUserExport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RequestUserExport`: RequestUserExport200Response
	fmt.Fprintf(os.Stdout, "Response from `UserExportsAPI.RequestUserExport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int32** | User id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRequestUserExportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestUserExportRequest** | [**RequestUserExportRequest**](RequestUserExportRequest.md) |  | 

### Return type

[**RequestUserExport200Response**](RequestUserExport200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

