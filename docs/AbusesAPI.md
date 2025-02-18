# \AbusesAPI

All URIs are relative to *https://peertube2.cpy.re*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV1AbusesAbuseIdDelete**](AbusesAPI.md#ApiV1AbusesAbuseIdDelete) | **Delete** /api/v1/abuses/{abuseId} | Delete an abuse
[**ApiV1AbusesAbuseIdMessagesAbuseMessageIdDelete**](AbusesAPI.md#ApiV1AbusesAbuseIdMessagesAbuseMessageIdDelete) | **Delete** /api/v1/abuses/{abuseId}/messages/{abuseMessageId} | Delete an abuse message
[**ApiV1AbusesAbuseIdMessagesGet**](AbusesAPI.md#ApiV1AbusesAbuseIdMessagesGet) | **Get** /api/v1/abuses/{abuseId}/messages | List messages of an abuse
[**ApiV1AbusesAbuseIdMessagesPost**](AbusesAPI.md#ApiV1AbusesAbuseIdMessagesPost) | **Post** /api/v1/abuses/{abuseId}/messages | Add message to an abuse
[**ApiV1AbusesAbuseIdPut**](AbusesAPI.md#ApiV1AbusesAbuseIdPut) | **Put** /api/v1/abuses/{abuseId} | Update an abuse
[**ApiV1AbusesPost**](AbusesAPI.md#ApiV1AbusesPost) | **Post** /api/v1/abuses | Report an abuse
[**GetAbuses**](AbusesAPI.md#GetAbuses) | **Get** /api/v1/abuses | List abuses
[**GetMyAbuses**](AbusesAPI.md#GetMyAbuses) | **Get** /api/v1/users/me/abuses | List my abuses



## ApiV1AbusesAbuseIdDelete

> ApiV1AbusesAbuseIdDelete(ctx, abuseId).Execute()

Delete an abuse

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
	abuseId := int32(56) // int32 | Abuse id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AbusesAPI.ApiV1AbusesAbuseIdDelete(context.Background(), abuseId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AbusesAPI.ApiV1AbusesAbuseIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**abuseId** | **int32** | Abuse id | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1AbusesAbuseIdDeleteRequest struct via the builder pattern


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


## ApiV1AbusesAbuseIdMessagesAbuseMessageIdDelete

> ApiV1AbusesAbuseIdMessagesAbuseMessageIdDelete(ctx, abuseId, abuseMessageId).Execute()

Delete an abuse message

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
	abuseId := int32(56) // int32 | Abuse id
	abuseMessageId := int32(56) // int32 | Abuse message id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AbusesAPI.ApiV1AbusesAbuseIdMessagesAbuseMessageIdDelete(context.Background(), abuseId, abuseMessageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AbusesAPI.ApiV1AbusesAbuseIdMessagesAbuseMessageIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**abuseId** | **int32** | Abuse id | 
**abuseMessageId** | **int32** | Abuse message id | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1AbusesAbuseIdMessagesAbuseMessageIdDeleteRequest struct via the builder pattern


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


## ApiV1AbusesAbuseIdMessagesGet

> ApiV1AbusesAbuseIdMessagesGet200Response ApiV1AbusesAbuseIdMessagesGet(ctx, abuseId).Execute()

List messages of an abuse

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
	abuseId := int32(56) // int32 | Abuse id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AbusesAPI.ApiV1AbusesAbuseIdMessagesGet(context.Background(), abuseId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AbusesAPI.ApiV1AbusesAbuseIdMessagesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1AbusesAbuseIdMessagesGet`: ApiV1AbusesAbuseIdMessagesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `AbusesAPI.ApiV1AbusesAbuseIdMessagesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**abuseId** | **int32** | Abuse id | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1AbusesAbuseIdMessagesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiV1AbusesAbuseIdMessagesGet200Response**](ApiV1AbusesAbuseIdMessagesGet200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1AbusesAbuseIdMessagesPost

> ApiV1AbusesAbuseIdMessagesPost(ctx, abuseId).ApiV1AbusesAbuseIdMessagesPostRequest(apiV1AbusesAbuseIdMessagesPostRequest).Execute()

Add message to an abuse

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
	apiV1AbusesAbuseIdMessagesPostRequest := *openapiclient.NewApiV1AbusesAbuseIdMessagesPostRequest("Message_example") // ApiV1AbusesAbuseIdMessagesPostRequest | 
	abuseId := int32(56) // int32 | Abuse id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AbusesAPI.ApiV1AbusesAbuseIdMessagesPost(context.Background(), abuseId).ApiV1AbusesAbuseIdMessagesPostRequest(apiV1AbusesAbuseIdMessagesPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AbusesAPI.ApiV1AbusesAbuseIdMessagesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**abuseId** | **int32** | Abuse id | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1AbusesAbuseIdMessagesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiV1AbusesAbuseIdMessagesPostRequest** | [**ApiV1AbusesAbuseIdMessagesPostRequest**](ApiV1AbusesAbuseIdMessagesPostRequest.md) |  | 


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


## ApiV1AbusesAbuseIdPut

> ApiV1AbusesAbuseIdPut(ctx, abuseId).ApiV1AbusesAbuseIdPutRequest(apiV1AbusesAbuseIdPutRequest).Execute()

Update an abuse

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
	abuseId := int32(56) // int32 | Abuse id
	apiV1AbusesAbuseIdPutRequest := *openapiclient.NewApiV1AbusesAbuseIdPutRequest() // ApiV1AbusesAbuseIdPutRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AbusesAPI.ApiV1AbusesAbuseIdPut(context.Background(), abuseId).ApiV1AbusesAbuseIdPutRequest(apiV1AbusesAbuseIdPutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AbusesAPI.ApiV1AbusesAbuseIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**abuseId** | **int32** | Abuse id | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1AbusesAbuseIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiV1AbusesAbuseIdPutRequest** | [**ApiV1AbusesAbuseIdPutRequest**](ApiV1AbusesAbuseIdPutRequest.md) |  | 

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


## ApiV1AbusesPost

> ApiV1AbusesPost200Response ApiV1AbusesPost(ctx).ApiV1AbusesPostRequest(apiV1AbusesPostRequest).Execute()

Report an abuse

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
	apiV1AbusesPostRequest := *openapiclient.NewApiV1AbusesPostRequest("Reason_example") // ApiV1AbusesPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AbusesAPI.ApiV1AbusesPost(context.Background()).ApiV1AbusesPostRequest(apiV1AbusesPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AbusesAPI.ApiV1AbusesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1AbusesPost`: ApiV1AbusesPost200Response
	fmt.Fprintf(os.Stdout, "Response from `AbusesAPI.ApiV1AbusesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1AbusesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiV1AbusesPostRequest** | [**ApiV1AbusesPostRequest**](ApiV1AbusesPostRequest.md) |  | 

### Return type

[**ApiV1AbusesPost200Response**](ApiV1AbusesPost200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAbuses

> GetMyAbuses200Response GetAbuses(ctx).Id(id).PredefinedReason(predefinedReason).Search(search).State(state).SearchReporter(searchReporter).SearchReportee(searchReportee).SearchVideo(searchVideo).SearchVideoChannel(searchVideoChannel).VideoIs(videoIs).Filter(filter).Start(start).Count(count).Sort(sort).Execute()

List abuses

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
	id := int32(56) // int32 | only list the report with this id (optional)
	predefinedReason := []string{"PredefinedReason_example"} // []string | predefined reason the listed reports should contain (optional)
	search := "search_example" // string | plain search that will match with video titles, reporter names and more (optional)
	state := openapiclient.AbuseStateSet(1) // AbuseStateSet |  (optional)
	searchReporter := "searchReporter_example" // string | only list reports of a specific reporter (optional)
	searchReportee := "searchReportee_example" // string | only list reports of a specific reportee (optional)
	searchVideo := "searchVideo_example" // string | only list reports of a specific video (optional)
	searchVideoChannel := "searchVideoChannel_example" // string | only list reports of a specific video channel (optional)
	videoIs := "videoIs_example" // string | only list deleted or blocklisted videos (optional)
	filter := "filter_example" // string | only list account, comment or video reports (optional)
	start := int32(56) // int32 | Offset used to paginate results (optional)
	count := int32(56) // int32 | Number of items to return (optional) (default to 15)
	sort := "sort_example" // string | Sort abuses by criteria (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AbusesAPI.GetAbuses(context.Background()).Id(id).PredefinedReason(predefinedReason).Search(search).State(state).SearchReporter(searchReporter).SearchReportee(searchReportee).SearchVideo(searchVideo).SearchVideoChannel(searchVideoChannel).VideoIs(videoIs).Filter(filter).Start(start).Count(count).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AbusesAPI.GetAbuses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAbuses`: GetMyAbuses200Response
	fmt.Fprintf(os.Stdout, "Response from `AbusesAPI.GetAbuses`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAbusesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int32** | only list the report with this id | 
 **predefinedReason** | **[]string** | predefined reason the listed reports should contain | 
 **search** | **string** | plain search that will match with video titles, reporter names and more | 
 **state** | [**AbuseStateSet**](AbuseStateSet.md) |  | 
 **searchReporter** | **string** | only list reports of a specific reporter | 
 **searchReportee** | **string** | only list reports of a specific reportee | 
 **searchVideo** | **string** | only list reports of a specific video | 
 **searchVideoChannel** | **string** | only list reports of a specific video channel | 
 **videoIs** | **string** | only list deleted or blocklisted videos | 
 **filter** | **string** | only list account, comment or video reports | 
 **start** | **int32** | Offset used to paginate results | 
 **count** | **int32** | Number of items to return | [default to 15]
 **sort** | **string** | Sort abuses by criteria | 

### Return type

[**GetMyAbuses200Response**](GetMyAbuses200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMyAbuses

> GetMyAbuses200Response GetMyAbuses(ctx).Id(id).State(state).Sort(sort).Start(start).Count(count).Execute()

List my abuses

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
	id := int32(56) // int32 | only list the report with this id (optional)
	state := openapiclient.AbuseStateSet(1) // AbuseStateSet |  (optional)
	sort := "sort_example" // string | Sort abuses by criteria (optional)
	start := int32(56) // int32 | Offset used to paginate results (optional)
	count := int32(56) // int32 | Number of items to return (optional) (default to 15)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AbusesAPI.GetMyAbuses(context.Background()).Id(id).State(state).Sort(sort).Start(start).Count(count).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AbusesAPI.GetMyAbuses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMyAbuses`: GetMyAbuses200Response
	fmt.Fprintf(os.Stdout, "Response from `AbusesAPI.GetMyAbuses`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMyAbusesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int32** | only list the report with this id | 
 **state** | [**AbuseStateSet**](AbuseStateSet.md) |  | 
 **sort** | **string** | Sort abuses by criteria | 
 **start** | **int32** | Offset used to paginate results | 
 **count** | **int32** | Number of items to return | [default to 15]

### Return type

[**GetMyAbuses200Response**](GetMyAbuses200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

