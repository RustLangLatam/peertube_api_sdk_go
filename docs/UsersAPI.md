# \UsersAPI

All URIs are relative to *https://peertube2.cpy.re*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddUser**](UsersAPI.md#AddUser) | **Post** /api/v1/users | Create a user
[**ConfirmTwoFactorRequest**](UsersAPI.md#ConfirmTwoFactorRequest) | **Post** /api/v1/users/{id}/two-factor/confirm-request | Confirm two factor auth
[**DelUser**](UsersAPI.md#DelUser) | **Delete** /api/v1/users/{id} | Delete a user
[**DisableTwoFactor**](UsersAPI.md#DisableTwoFactor) | **Post** /api/v1/users/{id}/two-factor/disable | Disable two factor auth
[**GetUser**](UsersAPI.md#GetUser) | **Get** /api/v1/users/{id} | Get a user
[**GetUsers**](UsersAPI.md#GetUsers) | **Get** /api/v1/users | List users
[**PutUser**](UsersAPI.md#PutUser) | **Put** /api/v1/users/{id} | Update a user
[**RequestTwoFactor**](UsersAPI.md#RequestTwoFactor) | **Post** /api/v1/users/{id}/two-factor/request | Request two factor auth
[**ResendEmailToVerifyUser**](UsersAPI.md#ResendEmailToVerifyUser) | **Post** /api/v1/users/ask-send-verify-email | Resend user verification link
[**VerifyUser**](UsersAPI.md#VerifyUser) | **Post** /api/v1/users/{id}/verify-email | Verify a user



## AddUser

> AddUserResponse AddUser(ctx).AddUser(addUser).Execute()

Create a user

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
	addUser := *openapiclient.NewAddUser("chocobozzz", "Password_example", "Email_example", *openapiclient.NewUserRole()) // AddUser | If the smtp server is configured, you can leave the password empty and an email will be sent asking the user to set it first. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.AddUser(context.Background()).AddUser(addUser).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.AddUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddUser`: AddUserResponse
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.AddUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addUser** | [**AddUser**](AddUser.md) | If the smtp server is configured, you can leave the password empty and an email will be sent asking the user to set it first.  | 

### Return type

[**AddUserResponse**](AddUserResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConfirmTwoFactorRequest

> ConfirmTwoFactorRequest(ctx, id).ConfirmTwoFactorRequestRequest(confirmTwoFactorRequestRequest).Execute()

Confirm two factor auth



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
	id := int32(56) // int32 | Entity id
	confirmTwoFactorRequestRequest := *openapiclient.NewConfirmTwoFactorRequestRequest("RequestToken_example", "OtpToken_example") // ConfirmTwoFactorRequestRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UsersAPI.ConfirmTwoFactorRequest(context.Background(), id).ConfirmTwoFactorRequestRequest(confirmTwoFactorRequestRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.ConfirmTwoFactorRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Entity id | 

### Other Parameters

Other parameters are passed through a pointer to a apiConfirmTwoFactorRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **confirmTwoFactorRequestRequest** | [**ConfirmTwoFactorRequestRequest**](ConfirmTwoFactorRequestRequest.md) |  | 

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


## DelUser

> DelUser(ctx, id).Execute()

Delete a user

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
	id := int32(56) // int32 | Entity id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UsersAPI.DelUser(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.DelUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Entity id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDelUserRequest struct via the builder pattern


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


## DisableTwoFactor

> DisableTwoFactor(ctx, id).RequestTwoFactorRequest(requestTwoFactorRequest).Execute()

Disable two factor auth



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
	id := int32(56) // int32 | Entity id
	requestTwoFactorRequest := *openapiclient.NewRequestTwoFactorRequest() // RequestTwoFactorRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UsersAPI.DisableTwoFactor(context.Background(), id).RequestTwoFactorRequest(requestTwoFactorRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.DisableTwoFactor``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Entity id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisableTwoFactorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestTwoFactorRequest** | [**RequestTwoFactorRequest**](RequestTwoFactorRequest.md) |  | 

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


## GetUser

> GetUser200Response GetUser(ctx, id).WithStats(withStats).Execute()

Get a user

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
	id := int32(56) // int32 | Entity id
	withStats := true // bool | include statistics about the user (only available as a moderator/admin) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.GetUser(context.Background(), id).WithStats(withStats).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUser`: GetUser200Response
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Entity id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **withStats** | **bool** | include statistics about the user (only available as a moderator/admin) | 

### Return type

[**GetUser200Response**](GetUser200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsers

> []User GetUsers(ctx).Search(search).Blocked(blocked).Start(start).Count(count).Sort(sort).Execute()

List users

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
	search := "search_example" // string | Plain text search that will match with user usernames or emails (optional)
	blocked := true // bool | Filter results down to (un)banned users (optional)
	start := int32(56) // int32 | Offset used to paginate results (optional)
	count := int32(56) // int32 | Number of items to return (optional) (default to 15)
	sort := "sort_example" // string | Sort users by criteria (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.GetUsers(context.Background()).Search(search).Blocked(blocked).Start(start).Count(count).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUsers`: []User
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **string** | Plain text search that will match with user usernames or emails | 
 **blocked** | **bool** | Filter results down to (un)banned users | 
 **start** | **int32** | Offset used to paginate results | 
 **count** | **int32** | Number of items to return | [default to 15]
 **sort** | **string** | Sort users by criteria | 

### Return type

[**[]User**](User.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutUser

> PutUser(ctx, id).UpdateUser(updateUser).Execute()

Update a user

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
	updateUser := *openapiclient.NewUpdateUser() // UpdateUser | 
	id := int32(56) // int32 | Entity id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UsersAPI.PutUser(context.Background(), id).UpdateUser(updateUser).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.PutUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Entity id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateUser** | [**UpdateUser**](UpdateUser.md) |  | 


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


## RequestTwoFactor

> []RequestTwoFactorResponse RequestTwoFactor(ctx, id).RequestTwoFactorRequest(requestTwoFactorRequest).Execute()

Request two factor auth



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
	id := int32(56) // int32 | Entity id
	requestTwoFactorRequest := *openapiclient.NewRequestTwoFactorRequest() // RequestTwoFactorRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.RequestTwoFactor(context.Background(), id).RequestTwoFactorRequest(requestTwoFactorRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.RequestTwoFactor``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RequestTwoFactor`: []RequestTwoFactorResponse
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.RequestTwoFactor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Entity id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRequestTwoFactorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestTwoFactorRequest** | [**RequestTwoFactorRequest**](RequestTwoFactorRequest.md) |  | 

### Return type

[**[]RequestTwoFactorResponse**](RequestTwoFactorResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResendEmailToVerifyUser

> ResendEmailToVerifyUser(ctx).ResendEmailToVerifyUserRequest(resendEmailToVerifyUserRequest).Execute()

Resend user verification link

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
	resendEmailToVerifyUserRequest := *openapiclient.NewResendEmailToVerifyUserRequest("Email_example") // ResendEmailToVerifyUserRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UsersAPI.ResendEmailToVerifyUser(context.Background()).ResendEmailToVerifyUserRequest(resendEmailToVerifyUserRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.ResendEmailToVerifyUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiResendEmailToVerifyUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resendEmailToVerifyUserRequest** | [**ResendEmailToVerifyUserRequest**](ResendEmailToVerifyUserRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerifyUser

> VerifyUser(ctx, id).VerifyUserRequest(verifyUserRequest).Execute()

Verify a user



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
	id := int32(56) // int32 | Entity id
	verifyUserRequest := *openapiclient.NewVerifyUserRequest("VerificationString_example") // VerifyUserRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UsersAPI.VerifyUser(context.Background(), id).VerifyUserRequest(verifyUserRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.VerifyUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Entity id | 

### Other Parameters

Other parameters are passed through a pointer to a apiVerifyUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **verifyUserRequest** | [**VerifyUserRequest**](VerifyUserRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

