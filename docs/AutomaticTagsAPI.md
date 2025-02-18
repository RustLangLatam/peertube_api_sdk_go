# \AutomaticTagsAPI

All URIs are relative to *https://peertube2.cpy.re*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV1AutomaticTagsAccountsAccountNameAvailableGet**](AutomaticTagsAPI.md#ApiV1AutomaticTagsAccountsAccountNameAvailableGet) | **Get** /api/v1/automatic-tags/accounts/{accountName}/available | Get account available auto tags
[**ApiV1AutomaticTagsPoliciesAccountsAccountNameCommentsGet**](AutomaticTagsAPI.md#ApiV1AutomaticTagsPoliciesAccountsAccountNameCommentsGet) | **Get** /api/v1/automatic-tags/policies/accounts/{accountName}/comments | Get account auto tag policies on comments
[**ApiV1AutomaticTagsPoliciesAccountsAccountNameCommentsPut**](AutomaticTagsAPI.md#ApiV1AutomaticTagsPoliciesAccountsAccountNameCommentsPut) | **Put** /api/v1/automatic-tags/policies/accounts/{accountName}/comments | Update account auto tag policies on comments
[**ApiV1AutomaticTagsServerAvailableGet**](AutomaticTagsAPI.md#ApiV1AutomaticTagsServerAvailableGet) | **Get** /api/v1/automatic-tags/server/available | Get server available auto tags



## ApiV1AutomaticTagsAccountsAccountNameAvailableGet

> AutomaticTagAvailable ApiV1AutomaticTagsAccountsAccountNameAvailableGet(ctx, accountName).Execute()

Get account available auto tags



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
	accountName := "accountName_example" // string | account name to get auto tag policies

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutomaticTagsAPI.ApiV1AutomaticTagsAccountsAccountNameAvailableGet(context.Background(), accountName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutomaticTagsAPI.ApiV1AutomaticTagsAccountsAccountNameAvailableGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1AutomaticTagsAccountsAccountNameAvailableGet`: AutomaticTagAvailable
	fmt.Fprintf(os.Stdout, "Response from `AutomaticTagsAPI.ApiV1AutomaticTagsAccountsAccountNameAvailableGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountName** | **string** | account name to get auto tag policies | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1AutomaticTagsAccountsAccountNameAvailableGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AutomaticTagAvailable**](AutomaticTagAvailable.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1AutomaticTagsPoliciesAccountsAccountNameCommentsGet

> CommentAutoTagPolicies ApiV1AutomaticTagsPoliciesAccountsAccountNameCommentsGet(ctx, accountName).Execute()

Get account auto tag policies on comments



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
	accountName := "accountName_example" // string | account name to get auto tag policies

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutomaticTagsAPI.ApiV1AutomaticTagsPoliciesAccountsAccountNameCommentsGet(context.Background(), accountName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutomaticTagsAPI.ApiV1AutomaticTagsPoliciesAccountsAccountNameCommentsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1AutomaticTagsPoliciesAccountsAccountNameCommentsGet`: CommentAutoTagPolicies
	fmt.Fprintf(os.Stdout, "Response from `AutomaticTagsAPI.ApiV1AutomaticTagsPoliciesAccountsAccountNameCommentsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountName** | **string** | account name to get auto tag policies | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1AutomaticTagsPoliciesAccountsAccountNameCommentsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CommentAutoTagPolicies**](CommentAutoTagPolicies.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1AutomaticTagsPoliciesAccountsAccountNameCommentsPut

> ApiV1AutomaticTagsPoliciesAccountsAccountNameCommentsPut(ctx, accountName).ApiV1AutomaticTagsPoliciesAccountsAccountNameCommentsPutRequest(apiV1AutomaticTagsPoliciesAccountsAccountNameCommentsPutRequest).Execute()

Update account auto tag policies on comments



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
	accountName := "accountName_example" // string | account name to update auto tag policies
	apiV1AutomaticTagsPoliciesAccountsAccountNameCommentsPutRequest := *openapiclient.NewApiV1AutomaticTagsPoliciesAccountsAccountNameCommentsPutRequest() // ApiV1AutomaticTagsPoliciesAccountsAccountNameCommentsPutRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AutomaticTagsAPI.ApiV1AutomaticTagsPoliciesAccountsAccountNameCommentsPut(context.Background(), accountName).ApiV1AutomaticTagsPoliciesAccountsAccountNameCommentsPutRequest(apiV1AutomaticTagsPoliciesAccountsAccountNameCommentsPutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutomaticTagsAPI.ApiV1AutomaticTagsPoliciesAccountsAccountNameCommentsPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountName** | **string** | account name to update auto tag policies | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1AutomaticTagsPoliciesAccountsAccountNameCommentsPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiV1AutomaticTagsPoliciesAccountsAccountNameCommentsPutRequest** | [**ApiV1AutomaticTagsPoliciesAccountsAccountNameCommentsPutRequest**](ApiV1AutomaticTagsPoliciesAccountsAccountNameCommentsPutRequest.md) |  | 

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


## ApiV1AutomaticTagsServerAvailableGet

> AutomaticTagAvailable ApiV1AutomaticTagsServerAvailableGet(ctx).Execute()

Get server available auto tags



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
	resp, r, err := apiClient.AutomaticTagsAPI.ApiV1AutomaticTagsServerAvailableGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutomaticTagsAPI.ApiV1AutomaticTagsServerAvailableGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1AutomaticTagsServerAvailableGet`: AutomaticTagAvailable
	fmt.Fprintf(os.Stdout, "Response from `AutomaticTagsAPI.ApiV1AutomaticTagsServerAvailableGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1AutomaticTagsServerAvailableGetRequest struct via the builder pattern


### Return type

[**AutomaticTagAvailable**](AutomaticTagAvailable.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

