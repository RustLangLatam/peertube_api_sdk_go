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
	"strings"
)

// InstanceFollowsAPIService InstanceFollowsAPI service
type InstanceFollowsAPIService service

type ApiApiV1ServerFollowersGetRequest struct {
	ctx        context.Context
	ApiService *InstanceFollowsAPIService
	state      *string
	actorType  *string
	start      *int32
	count      *int32
	sort       *string
}

func (r ApiApiV1ServerFollowersGetRequest) State(state string) ApiApiV1ServerFollowersGetRequest {
	r.state = &state
	return r
}

func (r ApiApiV1ServerFollowersGetRequest) ActorType(actorType string) ApiApiV1ServerFollowersGetRequest {
	r.actorType = &actorType
	return r
}

// Offset used to paginate results
func (r ApiApiV1ServerFollowersGetRequest) Start(start int32) ApiApiV1ServerFollowersGetRequest {
	r.start = &start
	return r
}

// Number of items to return
func (r ApiApiV1ServerFollowersGetRequest) Count(count int32) ApiApiV1ServerFollowersGetRequest {
	r.count = &count
	return r
}

// Sort column
func (r ApiApiV1ServerFollowersGetRequest) Sort(sort string) ApiApiV1ServerFollowersGetRequest {
	r.sort = &sort
	return r
}

func (r ApiApiV1ServerFollowersGetRequest) Execute() (*models.GetAccountFollowers200Response, *http.Response, error) {
	return r.ApiService.ApiV1ServerFollowersGetExecute(r)
}

/*
ApiV1ServerFollowersGet List instances following the server

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiApiV1ServerFollowersGetRequest
*/
func (a *InstanceFollowsAPIService) ApiV1ServerFollowersGet(ctx context.Context) ApiApiV1ServerFollowersGetRequest {
	return ApiApiV1ServerFollowersGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetAccountFollowers200Response
func (a *InstanceFollowsAPIService) ApiV1ServerFollowersGetExecute(r ApiApiV1ServerFollowersGetRequest) (*models.GetAccountFollowers200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.GetAccountFollowers200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InstanceFollowsAPIService.ApiV1ServerFollowersGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/server/followers"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.state != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "state", r.state, "form", "")
	}
	if r.actorType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "actorType", r.actorType, "form", "")
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

type ApiApiV1ServerFollowersNameWithHostAcceptPostRequest struct {
	ctx          context.Context
	ApiService   *InstanceFollowsAPIService
	nameWithHost string
}

func (r ApiApiV1ServerFollowersNameWithHostAcceptPostRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiV1ServerFollowersNameWithHostAcceptPostExecute(r)
}

/*
ApiV1ServerFollowersNameWithHostAcceptPost Accept a pending follower to your server

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param nameWithHost The remote actor handle to remove from your followers
	@return ApiApiV1ServerFollowersNameWithHostAcceptPostRequest
*/
func (a *InstanceFollowsAPIService) ApiV1ServerFollowersNameWithHostAcceptPost(ctx context.Context, nameWithHost string) ApiApiV1ServerFollowersNameWithHostAcceptPostRequest {
	return ApiApiV1ServerFollowersNameWithHostAcceptPostRequest{
		ApiService:   a,
		ctx:          ctx,
		nameWithHost: nameWithHost,
	}
}

// Execute executes the request
func (a *InstanceFollowsAPIService) ApiV1ServerFollowersNameWithHostAcceptPostExecute(r ApiApiV1ServerFollowersNameWithHostAcceptPostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InstanceFollowsAPIService.ApiV1ServerFollowersNameWithHostAcceptPost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/server/followers/{nameWithHost}/accept"
	localVarPath = strings.Replace(localVarPath, "{"+"nameWithHost"+"}", url.PathEscape(parameterValueToString(r.nameWithHost, "nameWithHost")), -1)

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

type ApiApiV1ServerFollowersNameWithHostDeleteRequest struct {
	ctx          context.Context
	ApiService   *InstanceFollowsAPIService
	nameWithHost string
}

func (r ApiApiV1ServerFollowersNameWithHostDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiV1ServerFollowersNameWithHostDeleteExecute(r)
}

/*
ApiV1ServerFollowersNameWithHostDelete Remove or reject a follower to your server

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param nameWithHost The remote actor handle to remove from your followers
	@return ApiApiV1ServerFollowersNameWithHostDeleteRequest
*/
func (a *InstanceFollowsAPIService) ApiV1ServerFollowersNameWithHostDelete(ctx context.Context, nameWithHost string) ApiApiV1ServerFollowersNameWithHostDeleteRequest {
	return ApiApiV1ServerFollowersNameWithHostDeleteRequest{
		ApiService:   a,
		ctx:          ctx,
		nameWithHost: nameWithHost,
	}
}

// Execute executes the request
func (a *InstanceFollowsAPIService) ApiV1ServerFollowersNameWithHostDeleteExecute(r ApiApiV1ServerFollowersNameWithHostDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InstanceFollowsAPIService.ApiV1ServerFollowersNameWithHostDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/server/followers/{nameWithHost}"
	localVarPath = strings.Replace(localVarPath, "{"+"nameWithHost"+"}", url.PathEscape(parameterValueToString(r.nameWithHost, "nameWithHost")), -1)

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

type ApiApiV1ServerFollowersNameWithHostRejectPostRequest struct {
	ctx          context.Context
	ApiService   *InstanceFollowsAPIService
	nameWithHost string
}

func (r ApiApiV1ServerFollowersNameWithHostRejectPostRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiV1ServerFollowersNameWithHostRejectPostExecute(r)
}

/*
ApiV1ServerFollowersNameWithHostRejectPost Reject a pending follower to your server

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param nameWithHost The remote actor handle to remove from your followers
	@return ApiApiV1ServerFollowersNameWithHostRejectPostRequest
*/
func (a *InstanceFollowsAPIService) ApiV1ServerFollowersNameWithHostRejectPost(ctx context.Context, nameWithHost string) ApiApiV1ServerFollowersNameWithHostRejectPostRequest {
	return ApiApiV1ServerFollowersNameWithHostRejectPostRequest{
		ApiService:   a,
		ctx:          ctx,
		nameWithHost: nameWithHost,
	}
}

// Execute executes the request
func (a *InstanceFollowsAPIService) ApiV1ServerFollowersNameWithHostRejectPostExecute(r ApiApiV1ServerFollowersNameWithHostRejectPostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InstanceFollowsAPIService.ApiV1ServerFollowersNameWithHostRejectPost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/server/followers/{nameWithHost}/reject"
	localVarPath = strings.Replace(localVarPath, "{"+"nameWithHost"+"}", url.PathEscape(parameterValueToString(r.nameWithHost, "nameWithHost")), -1)

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

type ApiApiV1ServerFollowingGetRequest struct {
	ctx        context.Context
	ApiService *InstanceFollowsAPIService
	state      *string
	actorType  *string
	start      *int32
	count      *int32
	sort       *string
}

func (r ApiApiV1ServerFollowingGetRequest) State(state string) ApiApiV1ServerFollowingGetRequest {
	r.state = &state
	return r
}

func (r ApiApiV1ServerFollowingGetRequest) ActorType(actorType string) ApiApiV1ServerFollowingGetRequest {
	r.actorType = &actorType
	return r
}

// Offset used to paginate results
func (r ApiApiV1ServerFollowingGetRequest) Start(start int32) ApiApiV1ServerFollowingGetRequest {
	r.start = &start
	return r
}

// Number of items to return
func (r ApiApiV1ServerFollowingGetRequest) Count(count int32) ApiApiV1ServerFollowingGetRequest {
	r.count = &count
	return r
}

// Sort column
func (r ApiApiV1ServerFollowingGetRequest) Sort(sort string) ApiApiV1ServerFollowingGetRequest {
	r.sort = &sort
	return r
}

func (r ApiApiV1ServerFollowingGetRequest) Execute() (*models.GetAccountFollowers200Response, *http.Response, error) {
	return r.ApiService.ApiV1ServerFollowingGetExecute(r)
}

/*
ApiV1ServerFollowingGet List instances followed by the server

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiApiV1ServerFollowingGetRequest
*/
func (a *InstanceFollowsAPIService) ApiV1ServerFollowingGet(ctx context.Context) ApiApiV1ServerFollowingGetRequest {
	return ApiApiV1ServerFollowingGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetAccountFollowers200Response
func (a *InstanceFollowsAPIService) ApiV1ServerFollowingGetExecute(r ApiApiV1ServerFollowingGetRequest) (*models.GetAccountFollowers200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.GetAccountFollowers200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InstanceFollowsAPIService.ApiV1ServerFollowingGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/server/following"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.state != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "state", r.state, "form", "")
	}
	if r.actorType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "actorType", r.actorType, "form", "")
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

type ApiApiV1ServerFollowingHostOrHandleDeleteRequest struct {
	ctx          context.Context
	ApiService   *InstanceFollowsAPIService
	hostOrHandle string
}

func (r ApiApiV1ServerFollowingHostOrHandleDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiV1ServerFollowingHostOrHandleDeleteExecute(r)
}

/*
ApiV1ServerFollowingHostOrHandleDelete Unfollow an actor (PeerTube instance, channel or account)

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param hostOrHandle The hostOrHandle to unfollow
	@return ApiApiV1ServerFollowingHostOrHandleDeleteRequest
*/
func (a *InstanceFollowsAPIService) ApiV1ServerFollowingHostOrHandleDelete(ctx context.Context, hostOrHandle string) ApiApiV1ServerFollowingHostOrHandleDeleteRequest {
	return ApiApiV1ServerFollowingHostOrHandleDeleteRequest{
		ApiService:   a,
		ctx:          ctx,
		hostOrHandle: hostOrHandle,
	}
}

// Execute executes the request
func (a *InstanceFollowsAPIService) ApiV1ServerFollowingHostOrHandleDeleteExecute(r ApiApiV1ServerFollowingHostOrHandleDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InstanceFollowsAPIService.ApiV1ServerFollowingHostOrHandleDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/server/following/{hostOrHandle}"
	localVarPath = strings.Replace(localVarPath, "{"+"hostOrHandle"+"}", url.PathEscape(parameterValueToString(r.hostOrHandle, "hostOrHandle")), -1)

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

type ApiApiV1ServerFollowingPostRequest struct {
	ctx                             context.Context
	ApiService                      *InstanceFollowsAPIService
	apiV1ServerFollowingPostRequest *models.ApiV1ServerFollowingPostRequest
}

func (r ApiApiV1ServerFollowingPostRequest) ApiV1ServerFollowingPostRequest(apiV1ServerFollowingPostRequest models.ApiV1ServerFollowingPostRequest) ApiApiV1ServerFollowingPostRequest {
	r.apiV1ServerFollowingPostRequest = &apiV1ServerFollowingPostRequest
	return r
}

func (r ApiApiV1ServerFollowingPostRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiV1ServerFollowingPostExecute(r)
}

/*
ApiV1ServerFollowingPost Follow a list of actors (PeerTube instance, channel or account)

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiApiV1ServerFollowingPostRequest
*/
func (a *InstanceFollowsAPIService) ApiV1ServerFollowingPost(ctx context.Context) ApiApiV1ServerFollowingPostRequest {
	return ApiApiV1ServerFollowingPostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *InstanceFollowsAPIService) ApiV1ServerFollowingPostExecute(r ApiApiV1ServerFollowingPostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InstanceFollowsAPIService.ApiV1ServerFollowingPost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/server/following"

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
	localVarPostBody = r.apiV1ServerFollowingPostRequest
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
