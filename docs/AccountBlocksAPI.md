# \AccountBlocksAPI

All URIs are relative to *https://peertube2.cpy.re*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV1BlocklistStatusGet**](AccountBlocksAPI.md#ApiV1BlocklistStatusGet) | **Get** /api/v1/blocklist/status | Get block status of accounts/hosts
[**ApiV1ServerBlocklistAccountsAccountNameDelete**](AccountBlocksAPI.md#ApiV1ServerBlocklistAccountsAccountNameDelete) | **Delete** /api/v1/server/blocklist/accounts/{accountName} | Unblock an account by its handle
[**ApiV1ServerBlocklistAccountsGet**](AccountBlocksAPI.md#ApiV1ServerBlocklistAccountsGet) | **Get** /api/v1/server/blocklist/accounts | List account blocks
[**ApiV1ServerBlocklistAccountsPost**](AccountBlocksAPI.md#ApiV1ServerBlocklistAccountsPost) | **Post** /api/v1/server/blocklist/accounts | Block an account



## ApiV1BlocklistStatusGet

> BlockStatus ApiV1BlocklistStatusGet(ctx).Accounts(accounts).Hosts(hosts).Execute()

Get block status of accounts/hosts

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
	accounts := []string{"Inner_example"} // []string | Check if these accounts are blocked (optional)
	hosts := []string{"Inner_example"} // []string | Check if these hosts are blocked (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountBlocksAPI.ApiV1BlocklistStatusGet(context.Background()).Accounts(accounts).Hosts(hosts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountBlocksAPI.ApiV1BlocklistStatusGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1BlocklistStatusGet`: BlockStatus
	fmt.Fprintf(os.Stdout, "Response from `AccountBlocksAPI.ApiV1BlocklistStatusGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1BlocklistStatusGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accounts** | **[]string** | Check if these accounts are blocked | 
 **hosts** | **[]string** | Check if these hosts are blocked | 

### Return type

[**BlockStatus**](BlockStatus.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1ServerBlocklistAccountsAccountNameDelete

> ApiV1ServerBlocklistAccountsAccountNameDelete(ctx, accountName).Execute()

Unblock an account by its handle

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
	accountName := "accountName_example" // string | account to unblock, in the form `username@domain`

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AccountBlocksAPI.ApiV1ServerBlocklistAccountsAccountNameDelete(context.Background(), accountName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountBlocksAPI.ApiV1ServerBlocklistAccountsAccountNameDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountName** | **string** | account to unblock, in the form &#x60;username@domain&#x60; | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1ServerBlocklistAccountsAccountNameDeleteRequest struct via the builder pattern


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


## ApiV1ServerBlocklistAccountsGet

> ApiV1ServerBlocklistAccountsGet(ctx).Start(start).Count(count).Sort(sort).Execute()

List account blocks

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
	start := int32(56) // int32 | Offset used to paginate results (optional)
	count := int32(56) // int32 | Number of items to return (optional) (default to 15)
	sort := "-createdAt" // string | Sort column (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AccountBlocksAPI.ApiV1ServerBlocklistAccountsGet(context.Background()).Start(start).Count(count).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountBlocksAPI.ApiV1ServerBlocklistAccountsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1ServerBlocklistAccountsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **int32** | Offset used to paginate results | 
 **count** | **int32** | Number of items to return | [default to 15]
 **sort** | **string** | Sort column | 

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


## ApiV1ServerBlocklistAccountsPost

> ApiV1ServerBlocklistAccountsPost(ctx).ApiV1ServerBlocklistAccountsPostRequest(apiV1ServerBlocklistAccountsPostRequest).Execute()

Block an account

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
	apiV1ServerBlocklistAccountsPostRequest := *openapiclient.NewApiV1ServerBlocklistAccountsPostRequest("chocobozzz@example.org") // ApiV1ServerBlocklistAccountsPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AccountBlocksAPI.ApiV1ServerBlocklistAccountsPost(context.Background()).ApiV1ServerBlocklistAccountsPostRequest(apiV1ServerBlocklistAccountsPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountBlocksAPI.ApiV1ServerBlocklistAccountsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1ServerBlocklistAccountsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiV1ServerBlocklistAccountsPostRequest** | [**ApiV1ServerBlocklistAccountsPostRequest**](ApiV1ServerBlocklistAccountsPostRequest.md) |  | 

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

