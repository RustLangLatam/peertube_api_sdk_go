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
	"github.com/RustLangLatam/peertube_api_sdk_go/utils"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strings"
)

// AbusesAPIService AbusesAPI service
type AbusesAPIService service

type ApiApiV1AbusesAbuseIdDeleteRequest struct {
	ctx        context.Context
	ApiService *AbusesAPIService
	abuseId    int32
}

func (r ApiApiV1AbusesAbuseIdDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiV1AbusesAbuseIdDeleteExecute(r)
}

/*
ApiV1AbusesAbuseIdDelete Delete an abuse

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param abuseId Abuse id
	@return ApiApiV1AbusesAbuseIdDeleteRequest
*/
func (a *AbusesAPIService) ApiV1AbusesAbuseIdDelete(ctx context.Context, abuseId int32) ApiApiV1AbusesAbuseIdDeleteRequest {
	return ApiApiV1AbusesAbuseIdDeleteRequest{
		ApiService: a,
		ctx:        ctx,
		abuseId:    abuseId,
	}
}

// Execute executes the request
func (a *AbusesAPIService) ApiV1AbusesAbuseIdDeleteExecute(r ApiApiV1AbusesAbuseIdDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AbusesAPIService.ApiV1AbusesAbuseIdDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/abuses/{abuseId}"
	localVarPath = strings.Replace(localVarPath, "{"+"abuseId"+"}", url.PathEscape(parameterValueToString(r.abuseId, "abuseId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.abuseId < 1 {
		return nil, utils.ReportError("abuseId must be greater than 1")
	}

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

type ApiApiV1AbusesAbuseIdMessagesAbuseMessageIdDeleteRequest struct {
	ctx            context.Context
	ApiService     *AbusesAPIService
	abuseId        int32
	abuseMessageId int32
}

func (r ApiApiV1AbusesAbuseIdMessagesAbuseMessageIdDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiV1AbusesAbuseIdMessagesAbuseMessageIdDeleteExecute(r)
}

/*
ApiV1AbusesAbuseIdMessagesAbuseMessageIdDelete Delete an abuse message

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param abuseId Abuse id
	@param abuseMessageId Abuse message id
	@return ApiApiV1AbusesAbuseIdMessagesAbuseMessageIdDeleteRequest
*/
func (a *AbusesAPIService) ApiV1AbusesAbuseIdMessagesAbuseMessageIdDelete(ctx context.Context, abuseId int32, abuseMessageId int32) ApiApiV1AbusesAbuseIdMessagesAbuseMessageIdDeleteRequest {
	return ApiApiV1AbusesAbuseIdMessagesAbuseMessageIdDeleteRequest{
		ApiService:     a,
		ctx:            ctx,
		abuseId:        abuseId,
		abuseMessageId: abuseMessageId,
	}
}

// Execute executes the request
func (a *AbusesAPIService) ApiV1AbusesAbuseIdMessagesAbuseMessageIdDeleteExecute(r ApiApiV1AbusesAbuseIdMessagesAbuseMessageIdDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AbusesAPIService.ApiV1AbusesAbuseIdMessagesAbuseMessageIdDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/abuses/{abuseId}/messages/{abuseMessageId}"
	localVarPath = strings.Replace(localVarPath, "{"+"abuseId"+"}", url.PathEscape(parameterValueToString(r.abuseId, "abuseId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"abuseMessageId"+"}", url.PathEscape(parameterValueToString(r.abuseMessageId, "abuseMessageId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.abuseId < 1 {
		return nil, utils.ReportError("abuseId must be greater than 1")
	}
	if r.abuseMessageId < 1 {
		return nil, utils.ReportError("abuseMessageId must be greater than 1")
	}

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

type ApiApiV1AbusesAbuseIdMessagesGetRequest struct {
	ctx        context.Context
	ApiService *AbusesAPIService
	abuseId    int32
}

func (r ApiApiV1AbusesAbuseIdMessagesGetRequest) Execute() (*models.ApiV1AbusesAbuseIdMessagesGet200Response, *http.Response, error) {
	return r.ApiService.ApiV1AbusesAbuseIdMessagesGetExecute(r)
}

/*
ApiV1AbusesAbuseIdMessagesGet List messages of an abuse

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param abuseId Abuse id
	@return ApiApiV1AbusesAbuseIdMessagesGetRequest
*/
func (a *AbusesAPIService) ApiV1AbusesAbuseIdMessagesGet(ctx context.Context, abuseId int32) ApiApiV1AbusesAbuseIdMessagesGetRequest {
	return ApiApiV1AbusesAbuseIdMessagesGetRequest{
		ApiService: a,
		ctx:        ctx,
		abuseId:    abuseId,
	}
}

// Execute executes the request
//
//	@return ApiV1AbusesAbuseIdMessagesGet200Response
func (a *AbusesAPIService) ApiV1AbusesAbuseIdMessagesGetExecute(r ApiApiV1AbusesAbuseIdMessagesGetRequest) (*models.ApiV1AbusesAbuseIdMessagesGet200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.ApiV1AbusesAbuseIdMessagesGet200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AbusesAPIService.ApiV1AbusesAbuseIdMessagesGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/abuses/{abuseId}/messages"
	localVarPath = strings.Replace(localVarPath, "{"+"abuseId"+"}", url.PathEscape(parameterValueToString(r.abuseId, "abuseId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.abuseId < 1 {
		return localVarReturnValue, nil, utils.ReportError("abuseId must be greater than 1")
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

type ApiApiV1AbusesAbuseIdMessagesPostRequest struct {
	ctx                                   context.Context
	ApiService                            *AbusesAPIService
	apiV1AbusesAbuseIdMessagesPostRequest *models.ApiV1AbusesAbuseIdMessagesPostRequest
	abuseId                               int32
}

func (r ApiApiV1AbusesAbuseIdMessagesPostRequest) ApiV1AbusesAbuseIdMessagesPostRequest(apiV1AbusesAbuseIdMessagesPostRequest models.ApiV1AbusesAbuseIdMessagesPostRequest) ApiApiV1AbusesAbuseIdMessagesPostRequest {
	r.apiV1AbusesAbuseIdMessagesPostRequest = &apiV1AbusesAbuseIdMessagesPostRequest
	return r
}

func (r ApiApiV1AbusesAbuseIdMessagesPostRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiV1AbusesAbuseIdMessagesPostExecute(r)
}

/*
ApiV1AbusesAbuseIdMessagesPost Add message to an abuse

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param abuseId Abuse id
	@return ApiApiV1AbusesAbuseIdMessagesPostRequest
*/
func (a *AbusesAPIService) ApiV1AbusesAbuseIdMessagesPost(ctx context.Context, abuseId int32) ApiApiV1AbusesAbuseIdMessagesPostRequest {
	return ApiApiV1AbusesAbuseIdMessagesPostRequest{
		ApiService: a,
		ctx:        ctx,
		abuseId:    abuseId,
	}
}

// Execute executes the request
func (a *AbusesAPIService) ApiV1AbusesAbuseIdMessagesPostExecute(r ApiApiV1AbusesAbuseIdMessagesPostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AbusesAPIService.ApiV1AbusesAbuseIdMessagesPost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/abuses/{abuseId}/messages"
	localVarPath = strings.Replace(localVarPath, "{"+"abuseId"+"}", url.PathEscape(parameterValueToString(r.abuseId, "abuseId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.apiV1AbusesAbuseIdMessagesPostRequest == nil {
		return nil, utils.ReportError("apiV1AbusesAbuseIdMessagesPostRequest is required and must be specified")
	}
	if r.abuseId < 1 {
		return nil, utils.ReportError("abuseId must be greater than 1")
	}

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
	localVarPostBody = r.apiV1AbusesAbuseIdMessagesPostRequest
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

type ApiApiV1AbusesAbuseIdPutRequest struct {
	ctx                          context.Context
	ApiService                   *AbusesAPIService
	abuseId                      int32
	apiV1AbusesAbuseIdPutRequest *models.ApiV1AbusesAbuseIdPutRequest
}

func (r ApiApiV1AbusesAbuseIdPutRequest) ApiV1AbusesAbuseIdPutRequest(apiV1AbusesAbuseIdPutRequest models.ApiV1AbusesAbuseIdPutRequest) ApiApiV1AbusesAbuseIdPutRequest {
	r.apiV1AbusesAbuseIdPutRequest = &apiV1AbusesAbuseIdPutRequest
	return r
}

func (r ApiApiV1AbusesAbuseIdPutRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiV1AbusesAbuseIdPutExecute(r)
}

/*
ApiV1AbusesAbuseIdPut Update an abuse

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param abuseId Abuse id
	@return ApiApiV1AbusesAbuseIdPutRequest
*/
func (a *AbusesAPIService) ApiV1AbusesAbuseIdPut(ctx context.Context, abuseId int32) ApiApiV1AbusesAbuseIdPutRequest {
	return ApiApiV1AbusesAbuseIdPutRequest{
		ApiService: a,
		ctx:        ctx,
		abuseId:    abuseId,
	}
}

// Execute executes the request
func (a *AbusesAPIService) ApiV1AbusesAbuseIdPutExecute(r ApiApiV1AbusesAbuseIdPutRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPut
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AbusesAPIService.ApiV1AbusesAbuseIdPut")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/abuses/{abuseId}"
	localVarPath = strings.Replace(localVarPath, "{"+"abuseId"+"}", url.PathEscape(parameterValueToString(r.abuseId, "abuseId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.abuseId < 1 {
		return nil, utils.ReportError("abuseId must be greater than 1")
	}

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
	localVarPostBody = r.apiV1AbusesAbuseIdPutRequest
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

type ApiApiV1AbusesPostRequest struct {
	ctx                    context.Context
	ApiService             *AbusesAPIService
	apiV1AbusesPostRequest *models.ApiV1AbusesPostRequest
}

func (r ApiApiV1AbusesPostRequest) ApiV1AbusesPostRequest(apiV1AbusesPostRequest models.ApiV1AbusesPostRequest) ApiApiV1AbusesPostRequest {
	r.apiV1AbusesPostRequest = &apiV1AbusesPostRequest
	return r
}

func (r ApiApiV1AbusesPostRequest) Execute() (*models.ApiV1AbusesPost200Response, *http.Response, error) {
	return r.ApiService.ApiV1AbusesPostExecute(r)
}

/*
ApiV1AbusesPost Report an abuse

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiApiV1AbusesPostRequest
*/
func (a *AbusesAPIService) ApiV1AbusesPost(ctx context.Context) ApiApiV1AbusesPostRequest {
	return ApiApiV1AbusesPostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ApiV1AbusesPost200Response
func (a *AbusesAPIService) ApiV1AbusesPostExecute(r ApiApiV1AbusesPostRequest) (*models.ApiV1AbusesPost200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.ApiV1AbusesPost200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AbusesAPIService.ApiV1AbusesPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/abuses"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.apiV1AbusesPostRequest == nil {
		return localVarReturnValue, nil, utils.ReportError("apiV1AbusesPostRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.apiV1AbusesPostRequest
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

type ApiGetAbusesRequest struct {
	ctx                context.Context
	ApiService         *AbusesAPIService
	id                 *int32
	predefinedReason   *[]string
	search             *string
	state              *models.AbuseStateSet
	searchReporter     *string
	searchReportee     *string
	searchVideo        *string
	searchVideoChannel *string
	videoIs            *string
	filter             *string
	start              *int32
	count              *int32
	sort               *string
}

// only list the report with this id
func (r ApiGetAbusesRequest) Id(id int32) ApiGetAbusesRequest {
	r.id = &id
	return r
}

// predefined reason the listed reports should contain
func (r ApiGetAbusesRequest) PredefinedReason(predefinedReason []string) ApiGetAbusesRequest {
	r.predefinedReason = &predefinedReason
	return r
}

// plain search that will match with video titles, reporter names and more
func (r ApiGetAbusesRequest) Search(search string) ApiGetAbusesRequest {
	r.search = &search
	return r
}

func (r ApiGetAbusesRequest) State(state models.AbuseStateSet) ApiGetAbusesRequest {
	r.state = &state
	return r
}

// only list reports of a specific reporter
func (r ApiGetAbusesRequest) SearchReporter(searchReporter string) ApiGetAbusesRequest {
	r.searchReporter = &searchReporter
	return r
}

// only list reports of a specific reportee
func (r ApiGetAbusesRequest) SearchReportee(searchReportee string) ApiGetAbusesRequest {
	r.searchReportee = &searchReportee
	return r
}

// only list reports of a specific video
func (r ApiGetAbusesRequest) SearchVideo(searchVideo string) ApiGetAbusesRequest {
	r.searchVideo = &searchVideo
	return r
}

// only list reports of a specific video channel
func (r ApiGetAbusesRequest) SearchVideoChannel(searchVideoChannel string) ApiGetAbusesRequest {
	r.searchVideoChannel = &searchVideoChannel
	return r
}

// only list deleted or blocklisted videos
func (r ApiGetAbusesRequest) VideoIs(videoIs string) ApiGetAbusesRequest {
	r.videoIs = &videoIs
	return r
}

// only list account, comment or video reports
func (r ApiGetAbusesRequest) Filter(filter string) ApiGetAbusesRequest {
	r.filter = &filter
	return r
}

// Offset used to paginate results
func (r ApiGetAbusesRequest) Start(start int32) ApiGetAbusesRequest {
	r.start = &start
	return r
}

// Number of items to return
func (r ApiGetAbusesRequest) Count(count int32) ApiGetAbusesRequest {
	r.count = &count
	return r
}

// Sort abuses by criteria
func (r ApiGetAbusesRequest) Sort(sort string) ApiGetAbusesRequest {
	r.sort = &sort
	return r
}

func (r ApiGetAbusesRequest) Execute() (*models.GetMyAbuses200Response, *http.Response, error) {
	return r.ApiService.GetAbusesExecute(r)
}

/*
GetAbuses List abuses

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetAbusesRequest
*/
func (a *AbusesAPIService) GetAbuses(ctx context.Context) ApiGetAbusesRequest {
	return ApiGetAbusesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetMyAbuses200Response
func (a *AbusesAPIService) GetAbusesExecute(r ApiGetAbusesRequest) (*models.GetMyAbuses200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.GetMyAbuses200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AbusesAPIService.GetAbuses")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/abuses"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.id != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "id", r.id, "form", "")
	}
	if r.predefinedReason != nil {
		t := *r.predefinedReason
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "predefinedReason", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "predefinedReason", t, "form", "multi")
		}
	}
	if r.search != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search", r.search, "form", "")
	}
	if r.state != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "state", r.state, "form", "")
	}
	if r.searchReporter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "searchReporter", r.searchReporter, "form", "")
	}
	if r.searchReportee != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "searchReportee", r.searchReportee, "form", "")
	}
	if r.searchVideo != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "searchVideo", r.searchVideo, "form", "")
	}
	if r.searchVideoChannel != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "searchVideoChannel", r.searchVideoChannel, "form", "")
	}
	if r.videoIs != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "videoIs", r.videoIs, "form", "")
	}
	if r.filter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter", r.filter, "form", "")
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

type ApiGetMyAbusesRequest struct {
	ctx        context.Context
	ApiService *AbusesAPIService
	id         *int32
	state      *models.AbuseStateSet
	sort       *string
	start      *int32
	count      *int32
}

// only list the report with this id
func (r ApiGetMyAbusesRequest) Id(id int32) ApiGetMyAbusesRequest {
	r.id = &id
	return r
}

func (r ApiGetMyAbusesRequest) State(state models.AbuseStateSet) ApiGetMyAbusesRequest {
	r.state = &state
	return r
}

// Sort abuses by criteria
func (r ApiGetMyAbusesRequest) Sort(sort string) ApiGetMyAbusesRequest {
	r.sort = &sort
	return r
}

// Offset used to paginate results
func (r ApiGetMyAbusesRequest) Start(start int32) ApiGetMyAbusesRequest {
	r.start = &start
	return r
}

// Number of items to return
func (r ApiGetMyAbusesRequest) Count(count int32) ApiGetMyAbusesRequest {
	r.count = &count
	return r
}

func (r ApiGetMyAbusesRequest) Execute() (*models.GetMyAbuses200Response, *http.Response, error) {
	return r.ApiService.GetMyAbusesExecute(r)
}

/*
GetMyAbuses List my abuses

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetMyAbusesRequest
*/
func (a *AbusesAPIService) GetMyAbuses(ctx context.Context) ApiGetMyAbusesRequest {
	return ApiGetMyAbusesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetMyAbuses200Response
func (a *AbusesAPIService) GetMyAbusesExecute(r ApiGetMyAbusesRequest) (*models.GetMyAbuses200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.GetMyAbuses200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AbusesAPIService.GetMyAbuses")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/users/me/abuses"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.id != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "id", r.id, "form", "")
	}
	if r.state != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "state", r.state, "form", "")
	}
	if r.sort != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort", r.sort, "form", "")
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
