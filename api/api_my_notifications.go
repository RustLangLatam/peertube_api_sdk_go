/*
PeerTube

The PeerTube API is built on HTTP(S) and is RESTful. You can use your favorite HTTP/REST library for your programming language to use PeerTube. The spec API is fully compatible with [openapi-generator](https://github.com/OpenAPITools/openapi-generator/wiki/API-client-generator-HOWTO) which generates a client SDK in the language of your choice - we generate some client SDKs automatically:  - [Python](https://framagit.org/framasoft/peertube/clients/python) - [Go](https://framagit.org/framasoft/peertube/clients/go) - [Kotlin](https://framagit.org/framasoft/peertube/clients/kotlin)  See the [REST API quick start](https://docs.joinpeertube.org/api/rest-getting-started) for a few examples of using the PeerTube API.  # Authentication  When you sign up for an account on a PeerTube instance, you are given the possibility to generate sessions on it, and authenticate there using an access token. Only __one access token can currently be used at a time__.  ## Roles  Accounts are given permissions based on their role. There are three roles on PeerTube: Administrator, Moderator, and User. See the [roles guide](https://docs.joinpeertube.org/admin/managing-users#roles) for a detail of their permissions.  # Errors  The API uses standard HTTP status codes to indicate the success or failure of the API call, completed by a [RFC7807-compliant](https://tools.ietf.org/html/rfc7807) response body.  ``` HTTP 1.1 404 Not Found Content-Type: application/problem+json; charset=utf-8  {   \"detail\": \"Video not found\",   \"docs\": \"https://docs.joinpeertube.org/api-rest-reference.html#operation/getVideo\",   \"status\": 404,   \"title\": \"Not Found\",   \"type\": \"about:blank\" } ```  We provide error `type` (following RFC7807) and `code` (internal PeerTube code) values for [a growing number of cases](https://github.com/Chocobozzz/PeerTube/blob/develop/packages/models/src/server/server-error-code.enum.ts), but it is still optional. Types are used to disambiguate errors that bear the same status code and are non-obvious:  ``` HTTP 1.1 403 Forbidden Content-Type: application/problem+json; charset=utf-8  {   \"detail\": \"Cannot get this video regarding follow constraints\",   \"docs\": \"https://docs.joinpeertube.org/api-rest-reference.html#operation/getVideo\",   \"status\": 403,   \"title\": \"Forbidden\",   \"type\": \"https://docs.joinpeertube.org/api-rest-reference.html#section/Errors/does_not_respect_follow_constraints\" } ```  Here a 403 error could otherwise mean that the video is private or blocklisted.  ### Validation errors  Each parameter is evaluated on its own against a set of rules before the route validator proceeds with potential testing involving parameter combinations. Errors coming from validation errors appear earlier and benefit from a more detailed error description:  ``` HTTP 1.1 400 Bad Request Content-Type: application/problem+json; charset=utf-8  {   \"detail\": \"Incorrect request parameters: id\",   \"docs\": \"https://docs.joinpeertube.org/api-rest-reference.html#operation/getVideo\",   \"instance\": \"/api/v1/videos/9c9de5e8-0a1e-484a-b099-e80766180\",   \"invalid-params\": {     \"id\": {       \"location\": \"params\",       \"msg\": \"Invalid value\",       \"param\": \"id\",       \"value\": \"9c9de5e8-0a1e-484a-b099-e80766180\"     }   },   \"status\": 400,   \"title\": \"Bad Request\",   \"type\": \"about:blank\" } ```  Where `id` is the name of the field concerned by the error, within the route definition. `invalid-params.<field>.location` can be either 'params', 'body', 'header', 'query' or 'cookies', and `invalid-params.<field>.value` reports the value that didn't pass validation whose `invalid-params.<field>.msg` is about.  ### Deprecated error fields  Some fields could be included with previous versions. They are still included but their use is deprecated: - `error`: superseded by `detail`  # Rate limits  We are rate-limiting all endpoints of PeerTube's API. Custom values can be set by administrators:  | Endpoint (prefix: `/api/v1`) | Calls         | Time frame   | |------------------------------|---------------|--------------| | `/_*`                         | 50            | 10 seconds   | | `POST /users/token`          | 15            | 5 minutes    | | `POST /users/register`       | 2<sup>*</sup> | 5 minutes    | | `POST /users/ask-send-verify-email` | 3      | 5 minutes    |  Depending on the endpoint, <sup>*</sup>failed requests are not taken into account. A service limit is announced by a `429 Too Many Requests` status code.  You can get details about the current state of your rate limit by reading the following headers:  | Header                  | Description                                                | |-------------------------|------------------------------------------------------------| | `X-RateLimit-Limit`     | Number of max requests allowed in the current time period  | | `X-RateLimit-Remaining` | Number of remaining requests in the current time period    | | `X-RateLimit-Reset`     | Timestamp of end of current time period as UNIX timestamp  | | `Retry-After`           | Seconds to delay after the first `429` is received         |  # CORS  This API features [Cross-Origin Resource Sharing (CORS)](https://fetch.spec.whatwg.org/), allowing cross-domain communication from the browser for some routes:  | Endpoint                    | |------------------------- ---| | `/api/_*`                    | | `/download/_*`               | | `/lazy-static/_*`            | | `/.well-known/webfinger`    |  In addition, all routes serving ActivityPub are CORS-enabled for all origins.

API version: 7.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"github.com/RustLangLatam/peertube_api_sdk_go/models"
	"io"
	"net/http"
	"net/url"
)

// MyNotificationsAPIService MyNotificationsAPI service
type MyNotificationsAPIService service

type ApiApiV1UsersMeNotificationSettingsPutRequest struct {
	ctx                                        context.Context
	ApiService                                 *MyNotificationsAPIService
	apiV1UsersMeNotificationSettingsPutRequest *models.ApiV1UsersMeNotificationSettingsPutRequest
}

func (r ApiApiV1UsersMeNotificationSettingsPutRequest) ApiV1UsersMeNotificationSettingsPutRequest(apiV1UsersMeNotificationSettingsPutRequest models.ApiV1UsersMeNotificationSettingsPutRequest) ApiApiV1UsersMeNotificationSettingsPutRequest {
	r.apiV1UsersMeNotificationSettingsPutRequest = &apiV1UsersMeNotificationSettingsPutRequest
	return r
}

func (r ApiApiV1UsersMeNotificationSettingsPutRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiV1UsersMeNotificationSettingsPutExecute(r)
}

/*
ApiV1UsersMeNotificationSettingsPut Update my notification settings

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiApiV1UsersMeNotificationSettingsPutRequest
*/
func (a *MyNotificationsAPIService) ApiV1UsersMeNotificationSettingsPut(ctx context.Context) ApiApiV1UsersMeNotificationSettingsPutRequest {
	return ApiApiV1UsersMeNotificationSettingsPutRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *MyNotificationsAPIService) ApiV1UsersMeNotificationSettingsPutExecute(r ApiApiV1UsersMeNotificationSettingsPutRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPut
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MyNotificationsAPIService.ApiV1UsersMeNotificationSettingsPut")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/users/me/notification-settings"

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
	localVarPostBody = r.apiV1UsersMeNotificationSettingsPutRequest
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

type ApiApiV1UsersMeNotificationsGetRequest struct {
	ctx        context.Context
	ApiService *MyNotificationsAPIService
	unread     *bool
	start      *int32
	count      *int32
	sort       *string
}

// only list unread notifications
func (r ApiApiV1UsersMeNotificationsGetRequest) Unread(unread bool) ApiApiV1UsersMeNotificationsGetRequest {
	r.unread = &unread
	return r
}

// Offset used to paginate results
func (r ApiApiV1UsersMeNotificationsGetRequest) Start(start int32) ApiApiV1UsersMeNotificationsGetRequest {
	r.start = &start
	return r
}

// Number of items to return
func (r ApiApiV1UsersMeNotificationsGetRequest) Count(count int32) ApiApiV1UsersMeNotificationsGetRequest {
	r.count = &count
	return r
}

// Sort column
func (r ApiApiV1UsersMeNotificationsGetRequest) Sort(sort string) ApiApiV1UsersMeNotificationsGetRequest {
	r.sort = &sort
	return r
}

func (r ApiApiV1UsersMeNotificationsGetRequest) Execute() (*models.NotificationListResponse, *http.Response, error) {
	return r.ApiService.ApiV1UsersMeNotificationsGetExecute(r)
}

/*
ApiV1UsersMeNotificationsGet List my notifications

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiApiV1UsersMeNotificationsGetRequest
*/
func (a *MyNotificationsAPIService) ApiV1UsersMeNotificationsGet(ctx context.Context) ApiApiV1UsersMeNotificationsGetRequest {
	return ApiApiV1UsersMeNotificationsGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return NotificationListResponse
func (a *MyNotificationsAPIService) ApiV1UsersMeNotificationsGetExecute(r ApiApiV1UsersMeNotificationsGetRequest) (*models.NotificationListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.NotificationListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MyNotificationsAPIService.ApiV1UsersMeNotificationsGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/users/me/notifications"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.unread != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "unread", r.unread, "form", "")
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

type ApiApiV1UsersMeNotificationsReadAllPostRequest struct {
	ctx        context.Context
	ApiService *MyNotificationsAPIService
}

func (r ApiApiV1UsersMeNotificationsReadAllPostRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiV1UsersMeNotificationsReadAllPostExecute(r)
}

/*
ApiV1UsersMeNotificationsReadAllPost Mark all my notification as read

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiApiV1UsersMeNotificationsReadAllPostRequest
*/
func (a *MyNotificationsAPIService) ApiV1UsersMeNotificationsReadAllPost(ctx context.Context) ApiApiV1UsersMeNotificationsReadAllPostRequest {
	return ApiApiV1UsersMeNotificationsReadAllPostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *MyNotificationsAPIService) ApiV1UsersMeNotificationsReadAllPostExecute(r ApiApiV1UsersMeNotificationsReadAllPostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MyNotificationsAPIService.ApiV1UsersMeNotificationsReadAllPost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/users/me/notifications/read-all"

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

type ApiApiV1UsersMeNotificationsReadPostRequest struct {
	ctx                                      context.Context
	ApiService                               *MyNotificationsAPIService
	apiV1UsersMeNotificationsReadPostRequest *models.ApiV1UsersMeNotificationsReadPostRequest
}

func (r ApiApiV1UsersMeNotificationsReadPostRequest) ApiV1UsersMeNotificationsReadPostRequest(apiV1UsersMeNotificationsReadPostRequest models.ApiV1UsersMeNotificationsReadPostRequest) ApiApiV1UsersMeNotificationsReadPostRequest {
	r.apiV1UsersMeNotificationsReadPostRequest = &apiV1UsersMeNotificationsReadPostRequest
	return r
}

func (r ApiApiV1UsersMeNotificationsReadPostRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiV1UsersMeNotificationsReadPostExecute(r)
}

/*
ApiV1UsersMeNotificationsReadPost Mark notifications as read by their id

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiApiV1UsersMeNotificationsReadPostRequest
*/
func (a *MyNotificationsAPIService) ApiV1UsersMeNotificationsReadPost(ctx context.Context) ApiApiV1UsersMeNotificationsReadPostRequest {
	return ApiApiV1UsersMeNotificationsReadPostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *MyNotificationsAPIService) ApiV1UsersMeNotificationsReadPostExecute(r ApiApiV1UsersMeNotificationsReadPostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MyNotificationsAPIService.ApiV1UsersMeNotificationsReadPost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/users/me/notifications/read"

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
	localVarPostBody = r.apiV1UsersMeNotificationsReadPostRequest
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
