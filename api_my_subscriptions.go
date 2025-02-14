/*
PeerTube

The PeerTube API is built on HTTP(S) and is RESTful. You can use your favorite HTTP/REST library for your programming language to use PeerTube. The spec API is fully compatible with [openapi-generator](https://github.com/OpenAPITools/openapi-generator/wiki/API-client-generator-HOWTO) which generates a client SDK in the language of your choice - we generate some client SDKs automatically:  - [Python](https://framagit.org/framasoft/peertube/clients/python) - [Go](https://framagit.org/framasoft/peertube/clients/go) - [Kotlin](https://framagit.org/framasoft/peertube/clients/kotlin)  See the [REST API quick start](https://docs.joinpeertube.org/api/rest-getting-started) for a few examples of using the PeerTube API.  # Authentication  When you sign up for an account on a PeerTube instance, you are given the possibility to generate sessions on it, and authenticate there using an access token. Only __one access token can currently be used at a time__.  ## Roles  Accounts are given permissions based on their role. There are three roles on PeerTube: Administrator, Moderator, and User. See the [roles guide](https://docs.joinpeertube.org/admin/managing-users#roles) for a detail of their permissions.  # Errors  The API uses standard HTTP status codes to indicate the success or failure of the API call, completed by a [RFC7807-compliant](https://tools.ietf.org/html/rfc7807) response body.  ``` HTTP 1.1 404 Not Found Content-Type: application/problem+json; charset=utf-8  {   \"detail\": \"Video not found\",   \"docs\": \"https://docs.joinpeertube.org/api-rest-reference.html#operation/getVideo\",   \"status\": 404,   \"title\": \"Not Found\",   \"type\": \"about:blank\" } ```  We provide error `type` (following RFC7807) and `code` (internal PeerTube code) values for [a growing number of cases](https://github.com/Chocobozzz/PeerTube/blob/develop/packages/models/src/server/server-error-code.enum.ts), but it is still optional. Types are used to disambiguate errors that bear the same status code and are non-obvious:  ``` HTTP 1.1 403 Forbidden Content-Type: application/problem+json; charset=utf-8  {   \"detail\": \"Cannot get this video regarding follow constraints\",   \"docs\": \"https://docs.joinpeertube.org/api-rest-reference.html#operation/getVideo\",   \"status\": 403,   \"title\": \"Forbidden\",   \"type\": \"https://docs.joinpeertube.org/api-rest-reference.html#section/Errors/does_not_respect_follow_constraints\" } ```  Here a 403 error could otherwise mean that the video is private or blocklisted.  ### Validation errors  Each parameter is evaluated on its own against a set of rules before the route validator proceeds with potential testing involving parameter combinations. Errors coming from validation errors appear earlier and benefit from a more detailed error description:  ``` HTTP 1.1 400 Bad Request Content-Type: application/problem+json; charset=utf-8  {   \"detail\": \"Incorrect request parameters: id\",   \"docs\": \"https://docs.joinpeertube.org/api-rest-reference.html#operation/getVideo\",   \"instance\": \"/api/v1/videos/9c9de5e8-0a1e-484a-b099-e80766180\",   \"invalid-params\": {     \"id\": {       \"location\": \"params\",       \"msg\": \"Invalid value\",       \"param\": \"id\",       \"value\": \"9c9de5e8-0a1e-484a-b099-e80766180\"     }   },   \"status\": 400,   \"title\": \"Bad Request\",   \"type\": \"about:blank\" } ```  Where `id` is the name of the field concerned by the error, within the route definition. `invalid-params.<field>.location` can be either 'params', 'body', 'header', 'query' or 'cookies', and `invalid-params.<field>.value` reports the value that didn't pass validation whose `invalid-params.<field>.msg` is about.  ### Deprecated error fields  Some fields could be included with previous versions. They are still included but their use is deprecated: - `error`: superseded by `detail`  # Rate limits  We are rate-limiting all endpoints of PeerTube's API. Custom values can be set by administrators:  | Endpoint (prefix: `/api/v1`) | Calls         | Time frame   | |------------------------------|---------------|--------------| | `/_*`                         | 50            | 10 seconds   | | `POST /users/token`          | 15            | 5 minutes    | | `POST /users/register`       | 2<sup>*</sup> | 5 minutes    | | `POST /users/ask-send-verify-email` | 3      | 5 minutes    |  Depending on the endpoint, <sup>*</sup>failed requests are not taken into account. A service limit is announced by a `429 Too Many Requests` status code.  You can get details about the current state of your rate limit by reading the following headers:  | Header                  | Description                                                | |-------------------------|------------------------------------------------------------| | `X-RateLimit-Limit`     | Number of max requests allowed in the current time period  | | `X-RateLimit-Remaining` | Number of remaining requests in the current time period    | | `X-RateLimit-Reset`     | Timestamp of end of current time period as UNIX timestamp  | | `Retry-After`           | Seconds to delay after the first `429` is received         |  # CORS  This API features [Cross-Origin Resource Sharing (CORS)](https://fetch.spec.whatwg.org/), allowing cross-domain communication from the browser for some routes:  | Endpoint                    | |------------------------- ---| | `/api/_*`                    | | `/download/_*`               | | `/lazy-static/_*`            | | `/.well-known/webfinger`    |  In addition, all routes serving ActivityPub are CORS-enabled for all origins.

API version: 7.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package peertube_api_sdk

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strings"
)

// MySubscriptionsAPIService MySubscriptionsAPI service
type MySubscriptionsAPIService service

type ApiApiV1UsersMeSubscriptionsExistGetRequest struct {
	ctx        context.Context
	ApiService *MySubscriptionsAPIService
	uris       *[]string
}

// list of uris to check if each is part of the user subscriptions
func (r ApiApiV1UsersMeSubscriptionsExistGetRequest) Uris(uris []string) ApiApiV1UsersMeSubscriptionsExistGetRequest {
	r.uris = &uris
	return r
}

func (r ApiApiV1UsersMeSubscriptionsExistGetRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.ApiV1UsersMeSubscriptionsExistGetExecute(r)
}

/*
ApiV1UsersMeSubscriptionsExistGet Get if subscriptions exist for my user

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiApiV1UsersMeSubscriptionsExistGetRequest
*/
func (a *MySubscriptionsAPIService) ApiV1UsersMeSubscriptionsExistGet(ctx context.Context) ApiApiV1UsersMeSubscriptionsExistGetRequest {
	return ApiApiV1UsersMeSubscriptionsExistGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return map[string]interface{}
func (a *MySubscriptionsAPIService) ApiV1UsersMeSubscriptionsExistGetExecute(r ApiApiV1UsersMeSubscriptionsExistGetRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MySubscriptionsAPIService.ApiV1UsersMeSubscriptionsExistGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/users/me/subscriptions/exist"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.uris == nil {
		return localVarReturnValue, nil, reportError("uris is required and must be specified")
	}

	{
		t := *r.uris
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "uris", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "uris", t, "form", "multi")
		}
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiApiV1UsersMeSubscriptionsGetRequest struct {
	ctx        context.Context
	ApiService *MySubscriptionsAPIService
	start      *int32
	count      *int32
	sort       *string
}

// Offset used to paginate results
func (r ApiApiV1UsersMeSubscriptionsGetRequest) Start(start int32) ApiApiV1UsersMeSubscriptionsGetRequest {
	r.start = &start
	return r
}

// Number of items to return
func (r ApiApiV1UsersMeSubscriptionsGetRequest) Count(count int32) ApiApiV1UsersMeSubscriptionsGetRequest {
	r.count = &count
	return r
}

// Sort column
func (r ApiApiV1UsersMeSubscriptionsGetRequest) Sort(sort string) ApiApiV1UsersMeSubscriptionsGetRequest {
	r.sort = &sort
	return r
}

func (r ApiApiV1UsersMeSubscriptionsGetRequest) Execute() (*VideoChannelList, *http.Response, error) {
	return r.ApiService.ApiV1UsersMeSubscriptionsGetExecute(r)
}

/*
ApiV1UsersMeSubscriptionsGet Get my user subscriptions

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiApiV1UsersMeSubscriptionsGetRequest
*/
func (a *MySubscriptionsAPIService) ApiV1UsersMeSubscriptionsGet(ctx context.Context) ApiApiV1UsersMeSubscriptionsGetRequest {
	return ApiApiV1UsersMeSubscriptionsGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return VideoChannelList
func (a *MySubscriptionsAPIService) ApiV1UsersMeSubscriptionsGetExecute(r ApiApiV1UsersMeSubscriptionsGetRequest) (*VideoChannelList, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *VideoChannelList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MySubscriptionsAPIService.ApiV1UsersMeSubscriptionsGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/users/me/subscriptions"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.start != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start", r.start, "form", "")
	}
	if r.count != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "count", r.count, "form", "")
	} else {
		var defaultValue int32 = 15
		r.count = &defaultValue
	}
	if r.sort != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort", r.sort, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiApiV1UsersMeSubscriptionsPostRequest struct {
	ctx                                  context.Context
	ApiService                           *MySubscriptionsAPIService
	apiV1UsersMeSubscriptionsPostRequest *ApiV1UsersMeSubscriptionsPostRequest
}

func (r ApiApiV1UsersMeSubscriptionsPostRequest) ApiV1UsersMeSubscriptionsPostRequest(apiV1UsersMeSubscriptionsPostRequest ApiV1UsersMeSubscriptionsPostRequest) ApiApiV1UsersMeSubscriptionsPostRequest {
	r.apiV1UsersMeSubscriptionsPostRequest = &apiV1UsersMeSubscriptionsPostRequest
	return r
}

func (r ApiApiV1UsersMeSubscriptionsPostRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiV1UsersMeSubscriptionsPostExecute(r)
}

/*
ApiV1UsersMeSubscriptionsPost Add subscription to my user

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiApiV1UsersMeSubscriptionsPostRequest
*/
func (a *MySubscriptionsAPIService) ApiV1UsersMeSubscriptionsPost(ctx context.Context) ApiApiV1UsersMeSubscriptionsPostRequest {
	return ApiApiV1UsersMeSubscriptionsPostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *MySubscriptionsAPIService) ApiV1UsersMeSubscriptionsPostExecute(r ApiApiV1UsersMeSubscriptionsPostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MySubscriptionsAPIService.ApiV1UsersMeSubscriptionsPost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/users/me/subscriptions"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.apiV1UsersMeSubscriptionsPostRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiApiV1UsersMeSubscriptionsSubscriptionHandleDeleteRequest struct {
	ctx                context.Context
	ApiService         *MySubscriptionsAPIService
	subscriptionHandle string
}

func (r ApiApiV1UsersMeSubscriptionsSubscriptionHandleDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiV1UsersMeSubscriptionsSubscriptionHandleDeleteExecute(r)
}

/*
ApiV1UsersMeSubscriptionsSubscriptionHandleDelete Delete subscription of my user

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param subscriptionHandle The subscription handle
	@return ApiApiV1UsersMeSubscriptionsSubscriptionHandleDeleteRequest
*/
func (a *MySubscriptionsAPIService) ApiV1UsersMeSubscriptionsSubscriptionHandleDelete(ctx context.Context, subscriptionHandle string) ApiApiV1UsersMeSubscriptionsSubscriptionHandleDeleteRequest {
	return ApiApiV1UsersMeSubscriptionsSubscriptionHandleDeleteRequest{
		ApiService:         a,
		ctx:                ctx,
		subscriptionHandle: subscriptionHandle,
	}
}

// Execute executes the request
func (a *MySubscriptionsAPIService) ApiV1UsersMeSubscriptionsSubscriptionHandleDeleteExecute(r ApiApiV1UsersMeSubscriptionsSubscriptionHandleDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MySubscriptionsAPIService.ApiV1UsersMeSubscriptionsSubscriptionHandleDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/users/me/subscriptions/{subscriptionHandle}"
	localVarPath = strings.Replace(localVarPath, "{"+"subscriptionHandle"+"}", url.PathEscape(parameterValueToString(r.subscriptionHandle, "subscriptionHandle")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiApiV1UsersMeSubscriptionsSubscriptionHandleGetRequest struct {
	ctx                context.Context
	ApiService         *MySubscriptionsAPIService
	subscriptionHandle string
}

func (r ApiApiV1UsersMeSubscriptionsSubscriptionHandleGetRequest) Execute() (*VideoChannel, *http.Response, error) {
	return r.ApiService.ApiV1UsersMeSubscriptionsSubscriptionHandleGetExecute(r)
}

/*
ApiV1UsersMeSubscriptionsSubscriptionHandleGet Get subscription of my user

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param subscriptionHandle The subscription handle
	@return ApiApiV1UsersMeSubscriptionsSubscriptionHandleGetRequest
*/
func (a *MySubscriptionsAPIService) ApiV1UsersMeSubscriptionsSubscriptionHandleGet(ctx context.Context, subscriptionHandle string) ApiApiV1UsersMeSubscriptionsSubscriptionHandleGetRequest {
	return ApiApiV1UsersMeSubscriptionsSubscriptionHandleGetRequest{
		ApiService:         a,
		ctx:                ctx,
		subscriptionHandle: subscriptionHandle,
	}
}

// Execute executes the request
//
//	@return VideoChannel
func (a *MySubscriptionsAPIService) ApiV1UsersMeSubscriptionsSubscriptionHandleGetExecute(r ApiApiV1UsersMeSubscriptionsSubscriptionHandleGetRequest) (*VideoChannel, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *VideoChannel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MySubscriptionsAPIService.ApiV1UsersMeSubscriptionsSubscriptionHandleGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/users/me/subscriptions/{subscriptionHandle}"
	localVarPath = strings.Replace(localVarPath, "{"+"subscriptionHandle"+"}", url.PathEscape(parameterValueToString(r.subscriptionHandle, "subscriptionHandle")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiApiV1UsersMeSubscriptionsVideosGetRequest struct {
	ctx                   context.Context
	ApiService            *MySubscriptionsAPIService
	categoryOneOf         *GetAccountVideosCategoryOneOfParameter
	isLive                *bool
	tagsOneOf             *GetAccountVideosTagsOneOfParameter
	tagsAllOf             *GetAccountVideosTagsAllOfParameter
	licenceOneOf          *GetAccountVideosLicenceOneOfParameter
	languageOneOf         *GetAccountVideosLanguageOneOfParameter
	autoTagOneOf          *GetAccountVideosTagsAllOfParameter
	nsfw                  *string
	isLocal               *bool
	include               *int32
	privacyOneOf          *VideoPrivacySet
	hasHLSFiles           *bool
	hasWebVideoFiles      *bool
	skipCount             *string
	start                 *int32
	count                 *int32
	sort                  *string
	excludeAlreadyWatched *bool
	search                *string
}

// category id of the video (see [/videos/categories](#operation/getCategories))
func (r ApiApiV1UsersMeSubscriptionsVideosGetRequest) CategoryOneOf(categoryOneOf GetAccountVideosCategoryOneOfParameter) ApiApiV1UsersMeSubscriptionsVideosGetRequest {
	r.categoryOneOf = &categoryOneOf
	return r
}

// whether or not the video is a live
func (r ApiApiV1UsersMeSubscriptionsVideosGetRequest) IsLive(isLive bool) ApiApiV1UsersMeSubscriptionsVideosGetRequest {
	r.isLive = &isLive
	return r
}

// tag(s) of the video
func (r ApiApiV1UsersMeSubscriptionsVideosGetRequest) TagsOneOf(tagsOneOf GetAccountVideosTagsOneOfParameter) ApiApiV1UsersMeSubscriptionsVideosGetRequest {
	r.tagsOneOf = &tagsOneOf
	return r
}

// tag(s) of the video, where all should be present in the video
func (r ApiApiV1UsersMeSubscriptionsVideosGetRequest) TagsAllOf(tagsAllOf GetAccountVideosTagsAllOfParameter) ApiApiV1UsersMeSubscriptionsVideosGetRequest {
	r.tagsAllOf = &tagsAllOf
	return r
}

// licence id of the video (see [/videos/licences](#operation/getLicences))
func (r ApiApiV1UsersMeSubscriptionsVideosGetRequest) LicenceOneOf(licenceOneOf GetAccountVideosLicenceOneOfParameter) ApiApiV1UsersMeSubscriptionsVideosGetRequest {
	r.licenceOneOf = &licenceOneOf
	return r
}

// language id of the video (see [/videos/languages](#operation/getLanguages)). Use &#x60;_unknown&#x60; to filter on videos that don&#39;t have a video language
func (r ApiApiV1UsersMeSubscriptionsVideosGetRequest) LanguageOneOf(languageOneOf GetAccountVideosLanguageOneOfParameter) ApiApiV1UsersMeSubscriptionsVideosGetRequest {
	r.languageOneOf = &languageOneOf
	return r
}

// **PeerTube &gt;&#x3D; 6.2** **Admins and moderators only** filter on videos that contain one of these automatic tags
func (r ApiApiV1UsersMeSubscriptionsVideosGetRequest) AutoTagOneOf(autoTagOneOf GetAccountVideosTagsAllOfParameter) ApiApiV1UsersMeSubscriptionsVideosGetRequest {
	r.autoTagOneOf = &autoTagOneOf
	return r
}

// whether to include nsfw videos, if any
func (r ApiApiV1UsersMeSubscriptionsVideosGetRequest) Nsfw(nsfw string) ApiApiV1UsersMeSubscriptionsVideosGetRequest {
	r.nsfw = &nsfw
	return r
}

// **PeerTube &gt;&#x3D; 4.0** Display only local or remote objects
func (r ApiApiV1UsersMeSubscriptionsVideosGetRequest) IsLocal(isLocal bool) ApiApiV1UsersMeSubscriptionsVideosGetRequest {
	r.isLocal = &isLocal
	return r
}

// **Only administrators and moderators can use this parameter**  Include additional videos in results (can be combined using bitwise or operator) - &#x60;0&#x60; NONE - &#x60;1&#x60; NOT_PUBLISHED_STATE - &#x60;2&#x60; BLACKLISTED - &#x60;4&#x60; BLOCKED_OWNER - &#x60;8&#x60; FILES - &#x60;16&#x60; CAPTIONS - &#x60;32&#x60; VIDEO SOURCE
func (r ApiApiV1UsersMeSubscriptionsVideosGetRequest) Include(include int32) ApiApiV1UsersMeSubscriptionsVideosGetRequest {
	r.include = &include
	return r
}

// **PeerTube &gt;&#x3D; 4.0** Display only videos in this specific privacy/privacies
func (r ApiApiV1UsersMeSubscriptionsVideosGetRequest) PrivacyOneOf(privacyOneOf VideoPrivacySet) ApiApiV1UsersMeSubscriptionsVideosGetRequest {
	r.privacyOneOf = &privacyOneOf
	return r
}

// **PeerTube &gt;&#x3D; 4.0** Display only videos that have HLS files
func (r ApiApiV1UsersMeSubscriptionsVideosGetRequest) HasHLSFiles(hasHLSFiles bool) ApiApiV1UsersMeSubscriptionsVideosGetRequest {
	r.hasHLSFiles = &hasHLSFiles
	return r
}

// **PeerTube &gt;&#x3D; 6.0** Display only videos that have Web Video files
func (r ApiApiV1UsersMeSubscriptionsVideosGetRequest) HasWebVideoFiles(hasWebVideoFiles bool) ApiApiV1UsersMeSubscriptionsVideosGetRequest {
	r.hasWebVideoFiles = &hasWebVideoFiles
	return r
}

// if you don&#39;t need the &#x60;total&#x60; in the response
func (r ApiApiV1UsersMeSubscriptionsVideosGetRequest) SkipCount(skipCount string) ApiApiV1UsersMeSubscriptionsVideosGetRequest {
	r.skipCount = &skipCount
	return r
}

// Offset used to paginate results
func (r ApiApiV1UsersMeSubscriptionsVideosGetRequest) Start(start int32) ApiApiV1UsersMeSubscriptionsVideosGetRequest {
	r.start = &start
	return r
}

// Number of items to return
func (r ApiApiV1UsersMeSubscriptionsVideosGetRequest) Count(count int32) ApiApiV1UsersMeSubscriptionsVideosGetRequest {
	r.count = &count
	return r
}

func (r ApiApiV1UsersMeSubscriptionsVideosGetRequest) Sort(sort string) ApiApiV1UsersMeSubscriptionsVideosGetRequest {
	r.sort = &sort
	return r
}

// Whether or not to exclude videos that are in the user&#39;s video history
func (r ApiApiV1UsersMeSubscriptionsVideosGetRequest) ExcludeAlreadyWatched(excludeAlreadyWatched bool) ApiApiV1UsersMeSubscriptionsVideosGetRequest {
	r.excludeAlreadyWatched = &excludeAlreadyWatched
	return r
}

// Plain text search, applied to various parts of the model depending on endpoint
func (r ApiApiV1UsersMeSubscriptionsVideosGetRequest) Search(search string) ApiApiV1UsersMeSubscriptionsVideosGetRequest {
	r.search = &search
	return r
}

func (r ApiApiV1UsersMeSubscriptionsVideosGetRequest) Execute() (*VideoListResponse, *http.Response, error) {
	return r.ApiService.ApiV1UsersMeSubscriptionsVideosGetExecute(r)
}

/*
ApiV1UsersMeSubscriptionsVideosGet List videos of subscriptions of my user

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiApiV1UsersMeSubscriptionsVideosGetRequest
*/
func (a *MySubscriptionsAPIService) ApiV1UsersMeSubscriptionsVideosGet(ctx context.Context) ApiApiV1UsersMeSubscriptionsVideosGetRequest {
	return ApiApiV1UsersMeSubscriptionsVideosGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return VideoListResponse
func (a *MySubscriptionsAPIService) ApiV1UsersMeSubscriptionsVideosGetExecute(r ApiApiV1UsersMeSubscriptionsVideosGetRequest) (*VideoListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *VideoListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MySubscriptionsAPIService.ApiV1UsersMeSubscriptionsVideosGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/users/me/subscriptions/videos"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.categoryOneOf != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "categoryOneOf", r.categoryOneOf, "form", "")
	}
	if r.isLive != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "isLive", r.isLive, "form", "")
	}
	if r.tagsOneOf != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "tagsOneOf", r.tagsOneOf, "form", "")
	}
	if r.tagsAllOf != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "tagsAllOf", r.tagsAllOf, "form", "")
	}
	if r.licenceOneOf != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "licenceOneOf", r.licenceOneOf, "form", "")
	}
	if r.languageOneOf != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "languageOneOf", r.languageOneOf, "form", "")
	}
	if r.autoTagOneOf != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "autoTagOneOf", r.autoTagOneOf, "form", "")
	}
	if r.nsfw != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "nsfw", r.nsfw, "form", "")
	}
	if r.isLocal != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "isLocal", r.isLocal, "form", "")
	}
	if r.include != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "include", r.include, "form", "")
	}
	if r.privacyOneOf != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "privacyOneOf", r.privacyOneOf, "form", "")
	}
	if r.hasHLSFiles != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "hasHLSFiles", r.hasHLSFiles, "form", "")
	}
	if r.hasWebVideoFiles != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "hasWebVideoFiles", r.hasWebVideoFiles, "form", "")
	}
	if r.skipCount != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "skipCount", r.skipCount, "form", "")
	} else {
		var defaultValue string = "false"
		r.skipCount = &defaultValue
	}
	if r.start != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start", r.start, "form", "")
	}
	if r.count != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "count", r.count, "form", "")
	} else {
		var defaultValue int32 = 15
		r.count = &defaultValue
	}
	if r.sort != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort", r.sort, "form", "")
	}
	if r.excludeAlreadyWatched != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "excludeAlreadyWatched", r.excludeAlreadyWatched, "form", "")
	}
	if r.search != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search", r.search, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
