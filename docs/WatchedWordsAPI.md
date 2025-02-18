# \WatchedWordsAPI

All URIs are relative to *https://peertube2.cpy.re*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV1WatchedWordsAccountsAccountNameListsGet**](WatchedWordsAPI.md#ApiV1WatchedWordsAccountsAccountNameListsGet) | **Get** /api/v1/watched-words/accounts/{accountName}/lists | List account watched words
[**ApiV1WatchedWordsAccountsAccountNameListsListIdDelete**](WatchedWordsAPI.md#ApiV1WatchedWordsAccountsAccountNameListsListIdDelete) | **Delete** /api/v1/watched-words/accounts/{accountName}/lists/{listId} | Delete account watched words
[**ApiV1WatchedWordsAccountsAccountNameListsListIdPut**](WatchedWordsAPI.md#ApiV1WatchedWordsAccountsAccountNameListsListIdPut) | **Put** /api/v1/watched-words/accounts/{accountName}/lists/{listId} | Update account watched words
[**ApiV1WatchedWordsAccountsAccountNameListsPost**](WatchedWordsAPI.md#ApiV1WatchedWordsAccountsAccountNameListsPost) | **Post** /api/v1/watched-words/accounts/{accountName}/lists | Add account watched words
[**ApiV1WatchedWordsServerListsGet**](WatchedWordsAPI.md#ApiV1WatchedWordsServerListsGet) | **Get** /api/v1/watched-words/server/lists | List server watched words
[**ApiV1WatchedWordsServerListsListIdDelete**](WatchedWordsAPI.md#ApiV1WatchedWordsServerListsListIdDelete) | **Delete** /api/v1/watched-words/server/lists/{listId} | Delete server watched words
[**ApiV1WatchedWordsServerListsListIdPut**](WatchedWordsAPI.md#ApiV1WatchedWordsServerListsListIdPut) | **Put** /api/v1/watched-words/server/lists/{listId} | Update server watched words
[**ApiV1WatchedWordsServerListsPost**](WatchedWordsAPI.md#ApiV1WatchedWordsServerListsPost) | **Post** /api/v1/watched-words/server/lists | Add server watched words



## ApiV1WatchedWordsAccountsAccountNameListsGet

> ApiV1WatchedWordsAccountsAccountNameListsGet200Response ApiV1WatchedWordsAccountsAccountNameListsGet(ctx, accountName).Execute()

List account watched words



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
	accountName := "accountName_example" // string | account name to list watched words

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WatchedWordsAPI.ApiV1WatchedWordsAccountsAccountNameListsGet(context.Background(), accountName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WatchedWordsAPI.ApiV1WatchedWordsAccountsAccountNameListsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1WatchedWordsAccountsAccountNameListsGet`: ApiV1WatchedWordsAccountsAccountNameListsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `WatchedWordsAPI.ApiV1WatchedWordsAccountsAccountNameListsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountName** | **string** | account name to list watched words | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1WatchedWordsAccountsAccountNameListsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiV1WatchedWordsAccountsAccountNameListsGet200Response**](ApiV1WatchedWordsAccountsAccountNameListsGet200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1WatchedWordsAccountsAccountNameListsListIdDelete

> ApiV1WatchedWordsAccountsAccountNameListsListIdDelete(ctx, accountName, listId).Execute()

Delete account watched words



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
	accountName := "accountName_example" // string | 
	listId := "listId_example" // string | list of watched words to delete

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WatchedWordsAPI.ApiV1WatchedWordsAccountsAccountNameListsListIdDelete(context.Background(), accountName, listId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WatchedWordsAPI.ApiV1WatchedWordsAccountsAccountNameListsListIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountName** | **string** |  | 
**listId** | **string** | list of watched words to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1WatchedWordsAccountsAccountNameListsListIdDeleteRequest struct via the builder pattern


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


## ApiV1WatchedWordsAccountsAccountNameListsListIdPut

> ApiV1WatchedWordsAccountsAccountNameListsListIdPut(ctx, accountName, listId).ApiV1WatchedWordsAccountsAccountNameListsPostRequest(apiV1WatchedWordsAccountsAccountNameListsPostRequest).Execute()

Update account watched words



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
	accountName := "accountName_example" // string | 
	listId := "listId_example" // string | list of watched words to update
	apiV1WatchedWordsAccountsAccountNameListsPostRequest := *openapiclient.NewApiV1WatchedWordsAccountsAccountNameListsPostRequest() // ApiV1WatchedWordsAccountsAccountNameListsPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WatchedWordsAPI.ApiV1WatchedWordsAccountsAccountNameListsListIdPut(context.Background(), accountName, listId).ApiV1WatchedWordsAccountsAccountNameListsPostRequest(apiV1WatchedWordsAccountsAccountNameListsPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WatchedWordsAPI.ApiV1WatchedWordsAccountsAccountNameListsListIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountName** | **string** |  | 
**listId** | **string** | list of watched words to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1WatchedWordsAccountsAccountNameListsListIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apiV1WatchedWordsAccountsAccountNameListsPostRequest** | [**ApiV1WatchedWordsAccountsAccountNameListsPostRequest**](ApiV1WatchedWordsAccountsAccountNameListsPostRequest.md) |  | 

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


## ApiV1WatchedWordsAccountsAccountNameListsPost

> ApiV1WatchedWordsAccountsAccountNameListsPost200Response ApiV1WatchedWordsAccountsAccountNameListsPost(ctx, accountName).ApiV1WatchedWordsAccountsAccountNameListsPostRequest(apiV1WatchedWordsAccountsAccountNameListsPostRequest).Execute()

Add account watched words



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
	accountName := "accountName_example" // string | 
	apiV1WatchedWordsAccountsAccountNameListsPostRequest := *openapiclient.NewApiV1WatchedWordsAccountsAccountNameListsPostRequest() // ApiV1WatchedWordsAccountsAccountNameListsPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WatchedWordsAPI.ApiV1WatchedWordsAccountsAccountNameListsPost(context.Background(), accountName).ApiV1WatchedWordsAccountsAccountNameListsPostRequest(apiV1WatchedWordsAccountsAccountNameListsPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WatchedWordsAPI.ApiV1WatchedWordsAccountsAccountNameListsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1WatchedWordsAccountsAccountNameListsPost`: ApiV1WatchedWordsAccountsAccountNameListsPost200Response
	fmt.Fprintf(os.Stdout, "Response from `WatchedWordsAPI.ApiV1WatchedWordsAccountsAccountNameListsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1WatchedWordsAccountsAccountNameListsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiV1WatchedWordsAccountsAccountNameListsPostRequest** | [**ApiV1WatchedWordsAccountsAccountNameListsPostRequest**](ApiV1WatchedWordsAccountsAccountNameListsPostRequest.md) |  | 

### Return type

[**ApiV1WatchedWordsAccountsAccountNameListsPost200Response**](ApiV1WatchedWordsAccountsAccountNameListsPost200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1WatchedWordsServerListsGet

> ApiV1WatchedWordsAccountsAccountNameListsGet200Response ApiV1WatchedWordsServerListsGet(ctx).Execute()

List server watched words



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WatchedWordsAPI.ApiV1WatchedWordsServerListsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WatchedWordsAPI.ApiV1WatchedWordsServerListsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1WatchedWordsServerListsGet`: ApiV1WatchedWordsAccountsAccountNameListsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `WatchedWordsAPI.ApiV1WatchedWordsServerListsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1WatchedWordsServerListsGetRequest struct via the builder pattern


### Return type

[**ApiV1WatchedWordsAccountsAccountNameListsGet200Response**](ApiV1WatchedWordsAccountsAccountNameListsGet200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1WatchedWordsServerListsListIdDelete

> ApiV1WatchedWordsServerListsListIdDelete(ctx, listId).Execute()

Delete server watched words



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
	listId := "listId_example" // string | list of watched words to delete

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WatchedWordsAPI.ApiV1WatchedWordsServerListsListIdDelete(context.Background(), listId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WatchedWordsAPI.ApiV1WatchedWordsServerListsListIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **string** | list of watched words to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1WatchedWordsServerListsListIdDeleteRequest struct via the builder pattern


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


## ApiV1WatchedWordsServerListsListIdPut

> ApiV1WatchedWordsServerListsListIdPut(ctx, listId).ApiV1WatchedWordsAccountsAccountNameListsPostRequest(apiV1WatchedWordsAccountsAccountNameListsPostRequest).Execute()

Update server watched words



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
	listId := "listId_example" // string | list of watched words to update
	apiV1WatchedWordsAccountsAccountNameListsPostRequest := *openapiclient.NewApiV1WatchedWordsAccountsAccountNameListsPostRequest() // ApiV1WatchedWordsAccountsAccountNameListsPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WatchedWordsAPI.ApiV1WatchedWordsServerListsListIdPut(context.Background(), listId).ApiV1WatchedWordsAccountsAccountNameListsPostRequest(apiV1WatchedWordsAccountsAccountNameListsPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WatchedWordsAPI.ApiV1WatchedWordsServerListsListIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **string** | list of watched words to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1WatchedWordsServerListsListIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiV1WatchedWordsAccountsAccountNameListsPostRequest** | [**ApiV1WatchedWordsAccountsAccountNameListsPostRequest**](ApiV1WatchedWordsAccountsAccountNameListsPostRequest.md) |  | 

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


## ApiV1WatchedWordsServerListsPost

> ApiV1WatchedWordsAccountsAccountNameListsPost200Response ApiV1WatchedWordsServerListsPost(ctx).ApiV1WatchedWordsAccountsAccountNameListsPostRequest(apiV1WatchedWordsAccountsAccountNameListsPostRequest).Execute()

Add server watched words



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
	apiV1WatchedWordsAccountsAccountNameListsPostRequest := *openapiclient.NewApiV1WatchedWordsAccountsAccountNameListsPostRequest() // ApiV1WatchedWordsAccountsAccountNameListsPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WatchedWordsAPI.ApiV1WatchedWordsServerListsPost(context.Background()).ApiV1WatchedWordsAccountsAccountNameListsPostRequest(apiV1WatchedWordsAccountsAccountNameListsPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WatchedWordsAPI.ApiV1WatchedWordsServerListsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1WatchedWordsServerListsPost`: ApiV1WatchedWordsAccountsAccountNameListsPost200Response
	fmt.Fprintf(os.Stdout, "Response from `WatchedWordsAPI.ApiV1WatchedWordsServerListsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1WatchedWordsServerListsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiV1WatchedWordsAccountsAccountNameListsPostRequest** | [**ApiV1WatchedWordsAccountsAccountNameListsPostRequest**](ApiV1WatchedWordsAccountsAccountNameListsPostRequest.md) |  | 

### Return type

[**ApiV1WatchedWordsAccountsAccountNameListsPost200Response**](ApiV1WatchedWordsAccountsAccountNameListsPost200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

