# \SessionAPI

All URIs are relative to *https://peertube2.cpy.re*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOAuthClient**](SessionAPI.md#GetOAuthClient) | **Get** /api/v1/oauth-clients/local | Login prerequisite
[**GetOAuthToken**](SessionAPI.md#GetOAuthToken) | **Post** /api/v1/users/token | Login
[**RevokeOAuthToken**](SessionAPI.md#RevokeOAuthToken) | **Post** /api/v1/users/revoke-token | Logout



## GetOAuthClient

> OAuthClient GetOAuthClient(ctx).Execute()

Login prerequisite



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
	resp, r, err := apiClient.SessionAPI.GetOAuthClient(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionAPI.GetOAuthClient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOAuthClient`: OAuthClient
	fmt.Fprintf(os.Stdout, "Response from `SessionAPI.GetOAuthClient`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetOAuthClientRequest struct via the builder pattern


### Return type

[**OAuthClient**](OAuthClient.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOAuthToken

> GetOAuthToken200Response GetOAuthToken(ctx).ClientId(clientId).ClientSecret(clientSecret).GrantType(grantType).Username(username).Password(password).RefreshToken(refreshToken).Execute()

Login



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
	clientId := "clientId_example" // string |  (optional)
	clientSecret := "clientSecret_example" // string |  (optional)
	grantType := "grantType_example" // string |  (optional) (default to "password")
	username := "username_example" // string | immutable name of the user, used to find or mention its actor (optional)
	password := "password_example" // string |  (optional)
	refreshToken := "refreshToken_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SessionAPI.GetOAuthToken(context.Background()).ClientId(clientId).ClientSecret(clientSecret).GrantType(grantType).Username(username).Password(password).RefreshToken(refreshToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionAPI.GetOAuthToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOAuthToken`: GetOAuthToken200Response
	fmt.Fprintf(os.Stdout, "Response from `SessionAPI.GetOAuthToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOAuthTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientId** | **string** |  | 
 **clientSecret** | **string** |  | 
 **grantType** | **string** |  | [default to &quot;password&quot;]
 **username** | **string** | immutable name of the user, used to find or mention its actor | 
 **password** | **string** |  | 
 **refreshToken** | **string** |  | 

### Return type

[**GetOAuthToken200Response**](GetOAuthToken200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RevokeOAuthToken

> RevokeOAuthToken(ctx).Execute()

Logout



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
	r, err := apiClient.SessionAPI.RevokeOAuthToken(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionAPI.RevokeOAuthToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiRevokeOAuthTokenRequest struct via the builder pattern


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

