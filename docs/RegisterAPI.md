# \RegisterAPI

All URIs are relative to *https://peertube2.cpy.re*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AcceptRegistration**](RegisterAPI.md#AcceptRegistration) | **Post** /api/v1/users/registrations/{registrationId}/accept | Accept registration
[**DeleteRegistration**](RegisterAPI.md#DeleteRegistration) | **Delete** /api/v1/users/registrations/{registrationId} | Delete registration
[**ListRegistrations**](RegisterAPI.md#ListRegistrations) | **Get** /api/v1/users/registrations | List registrations
[**RegisterUser**](RegisterAPI.md#RegisterUser) | **Post** /api/v1/users/register | Register a user
[**RejectRegistration**](RegisterAPI.md#RejectRegistration) | **Post** /api/v1/users/registrations/{registrationId}/reject | Reject registration
[**RequestRegistration**](RegisterAPI.md#RequestRegistration) | **Post** /api/v1/users/registrations/request | Request registration
[**ResendEmailToVerifyRegistration**](RegisterAPI.md#ResendEmailToVerifyRegistration) | **Post** /api/v1/users/registrations/ask-send-verify-email | Resend verification link to registration email
[**VerifyRegistrationEmail**](RegisterAPI.md#VerifyRegistrationEmail) | **Post** /api/v1/users/registrations/{registrationId}/verify-email | Verify a registration email



## AcceptRegistration

> AcceptRegistration(ctx, registrationId).UserRegistrationAcceptOrReject(userRegistrationAcceptOrReject).Execute()

Accept registration

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
	registrationId := int32(56) // int32 | Registration ID
	userRegistrationAcceptOrReject := *openapiclient.NewUserRegistrationAcceptOrReject("ModerationResponse_example") // UserRegistrationAcceptOrReject |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RegisterAPI.AcceptRegistration(context.Background(), registrationId).UserRegistrationAcceptOrReject(userRegistrationAcceptOrReject).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RegisterAPI.AcceptRegistration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registrationId** | **int32** | Registration ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAcceptRegistrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userRegistrationAcceptOrReject** | [**UserRegistrationAcceptOrReject**](UserRegistrationAcceptOrReject.md) |  | 

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


## DeleteRegistration

> DeleteRegistration(ctx, registrationId).Execute()

Delete registration



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
	registrationId := int32(56) // int32 | Registration ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RegisterAPI.DeleteRegistration(context.Background(), registrationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RegisterAPI.DeleteRegistration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registrationId** | **int32** | Registration ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRegistrationRequest struct via the builder pattern


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


## ListRegistrations

> ListRegistrations200Response ListRegistrations(ctx).Start(start).Count(count).Search(search).Sort(sort).Execute()

List registrations

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
	start := int32(56) // int32 | Offset used to paginate results (optional)
	count := int32(56) // int32 | Number of items to return (optional) (default to 15)
	search := "search_example" // string |  (optional)
	sort := "sort_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RegisterAPI.ListRegistrations(context.Background()).Start(start).Count(count).Search(search).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RegisterAPI.ListRegistrations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListRegistrations`: ListRegistrations200Response
	fmt.Fprintf(os.Stdout, "Response from `RegisterAPI.ListRegistrations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRegistrationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **int32** | Offset used to paginate results | 
 **count** | **int32** | Number of items to return | [default to 15]
 **search** | **string** |  | 
 **sort** | **string** |  | 

### Return type

[**ListRegistrations200Response**](ListRegistrations200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegisterUser

> RegisterUser(ctx).RegisterUser(registerUser).Execute()

Register a user



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
	registerUser := *openapiclient.NewRegisterUser("chocobozzz", "Password_example", "Email_example") // RegisterUser | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RegisterAPI.RegisterUser(context.Background()).RegisterUser(registerUser).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RegisterAPI.RegisterUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegisterUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **registerUser** | [**RegisterUser**](RegisterUser.md) |  | 

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


## RejectRegistration

> RejectRegistration(ctx, registrationId).UserRegistrationAcceptOrReject(userRegistrationAcceptOrReject).Execute()

Reject registration

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
	registrationId := int32(56) // int32 | Registration ID
	userRegistrationAcceptOrReject := *openapiclient.NewUserRegistrationAcceptOrReject("ModerationResponse_example") // UserRegistrationAcceptOrReject |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RegisterAPI.RejectRegistration(context.Background(), registrationId).UserRegistrationAcceptOrReject(userRegistrationAcceptOrReject).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RegisterAPI.RejectRegistration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registrationId** | **int32** | Registration ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRejectRegistrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userRegistrationAcceptOrReject** | [**UserRegistrationAcceptOrReject**](UserRegistrationAcceptOrReject.md) |  | 

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


## RequestRegistration

> UserRegistration RequestRegistration(ctx).UserRegistrationRequest(userRegistrationRequest).Execute()

Request registration



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
	userRegistrationRequest := *openapiclient.NewUserRegistrationRequest("chocobozzz", "Password_example", "Email_example", "RegistrationReason_example") // UserRegistrationRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RegisterAPI.RequestRegistration(context.Background()).UserRegistrationRequest(userRegistrationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RegisterAPI.RequestRegistration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RequestRegistration`: UserRegistration
	fmt.Fprintf(os.Stdout, "Response from `RegisterAPI.RequestRegistration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRequestRegistrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userRegistrationRequest** | [**UserRegistrationRequest**](UserRegistrationRequest.md) |  | 

### Return type

[**UserRegistration**](UserRegistration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResendEmailToVerifyRegistration

> ResendEmailToVerifyRegistration(ctx).ResendEmailToVerifyRegistrationRequest(resendEmailToVerifyRegistrationRequest).Execute()

Resend verification link to registration email

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
	resendEmailToVerifyRegistrationRequest := *openapiclient.NewResendEmailToVerifyRegistrationRequest("Email_example") // ResendEmailToVerifyRegistrationRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RegisterAPI.ResendEmailToVerifyRegistration(context.Background()).ResendEmailToVerifyRegistrationRequest(resendEmailToVerifyRegistrationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RegisterAPI.ResendEmailToVerifyRegistration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiResendEmailToVerifyRegistrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resendEmailToVerifyRegistrationRequest** | [**ResendEmailToVerifyRegistrationRequest**](ResendEmailToVerifyRegistrationRequest.md) |  | 

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


## VerifyRegistrationEmail

> VerifyRegistrationEmail(ctx, registrationId).VerifyRegistrationEmailRequest(verifyRegistrationEmailRequest).Execute()

Verify a registration email



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
	registrationId := int32(56) // int32 | Registration ID
	verifyRegistrationEmailRequest := *openapiclient.NewVerifyRegistrationEmailRequest("VerificationString_example") // VerifyRegistrationEmailRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RegisterAPI.VerifyRegistrationEmail(context.Background(), registrationId).VerifyRegistrationEmailRequest(verifyRegistrationEmailRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RegisterAPI.VerifyRegistrationEmail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registrationId** | **int32** | Registration ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiVerifyRegistrationEmailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **verifyRegistrationEmailRequest** | [**VerifyRegistrationEmailRequest**](VerifyRegistrationEmailRequest.md) |  | 

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

