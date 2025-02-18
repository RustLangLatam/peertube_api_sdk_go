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
	"os"
	"strings"
	"time"
)

// VideoAPIService VideoAPI service
type VideoAPIService service

type ApiAddViewRequest struct {
	ctx              context.Context
	ApiService       *VideoAPIService
	userViewingVideo *models.UserViewingVideo
	id               models.ApiV1VideosOwnershipIdAcceptPostIdParameter
}

func (r ApiAddViewRequest) UserViewingVideo(userViewingVideo models.UserViewingVideo) ApiAddViewRequest {
	r.userViewingVideo = &userViewingVideo
	return r
}

func (r ApiAddViewRequest) Execute() (*http.Response, error) {
	return r.ApiService.AddViewExecute(r)
}

/*
AddView Notify user is watching a video

Call this endpoint regularly (every 5-10 seconds for example) to notify the server the user is watching the video. After a while, PeerTube will increase video's viewers counter. If the user is authenticated, PeerTube will also store the current player time.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id The object id, uuid or short uuid
	@return ApiAddViewRequest
*/
func (a *VideoAPIService) AddView(ctx context.Context, id models.ApiV1VideosOwnershipIdAcceptPostIdParameter) ApiAddViewRequest {
	return ApiAddViewRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
func (a *VideoAPIService) AddViewExecute(r ApiAddViewRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VideoAPIService.AddView")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/videos/{id}/views"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.userViewingVideo == nil {
		return nil, utils.ReportError("userViewingVideo is required and must be specified")
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
	localVarPostBody = r.userViewingVideo
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

type ApiApiV1VideosIdStudioEditPostRequest struct {
	ctx        context.Context
	ApiService *VideoAPIService
	id         models.ApiV1VideosOwnershipIdAcceptPostIdParameter
}

func (r ApiApiV1VideosIdStudioEditPostRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiV1VideosIdStudioEditPostExecute(r)
}

/*
ApiV1VideosIdStudioEditPost Create a studio task

Create a task to edit a video  (cut, add intro/outro etc)

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id The object id, uuid or short uuid
	@return ApiApiV1VideosIdStudioEditPostRequest
*/
func (a *VideoAPIService) ApiV1VideosIdStudioEditPost(ctx context.Context, id models.ApiV1VideosOwnershipIdAcceptPostIdParameter) ApiApiV1VideosIdStudioEditPostRequest {
	return ApiApiV1VideosIdStudioEditPostRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
func (a *VideoAPIService) ApiV1VideosIdStudioEditPostExecute(r ApiApiV1VideosIdStudioEditPostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VideoAPIService.ApiV1VideosIdStudioEditPost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/videos/{id}/studio/edit"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/x-www-form-urlencoded"}

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

type ApiApiV1VideosIdWatchingPutRequest struct {
	ctx              context.Context
	ApiService       *VideoAPIService
	userViewingVideo *models.UserViewingVideo
	id               models.ApiV1VideosOwnershipIdAcceptPostIdParameter
}

func (r ApiApiV1VideosIdWatchingPutRequest) UserViewingVideo(userViewingVideo models.UserViewingVideo) ApiApiV1VideosIdWatchingPutRequest {
	r.userViewingVideo = &userViewingVideo
	return r
}

func (r ApiApiV1VideosIdWatchingPutRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiV1VideosIdWatchingPutExecute(r)
}

/*
ApiV1VideosIdWatchingPut Set watching progress of a video

This endpoint has been deprecated. Use `/videos/{id}/views` instead

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id The object id, uuid or short uuid
	@return ApiApiV1VideosIdWatchingPutRequest

Deprecated
*/
func (a *VideoAPIService) ApiV1VideosIdWatchingPut(ctx context.Context, id models.ApiV1VideosOwnershipIdAcceptPostIdParameter) ApiApiV1VideosIdWatchingPutRequest {
	return ApiApiV1VideosIdWatchingPutRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
// Deprecated
func (a *VideoAPIService) ApiV1VideosIdWatchingPutExecute(r ApiApiV1VideosIdWatchingPutRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPut
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VideoAPIService.ApiV1VideosIdWatchingPut")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/videos/{id}/watching"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.userViewingVideo == nil {
		return nil, utils.ReportError("userViewingVideo is required and must be specified")
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
	localVarPostBody = r.userViewingVideo
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

type ApiDelVideoRequest struct {
	ctx        context.Context
	ApiService *VideoAPIService
	id         models.ApiV1VideosOwnershipIdAcceptPostIdParameter
}

func (r ApiDelVideoRequest) Execute() (*http.Response, error) {
	return r.ApiService.DelVideoExecute(r)
}

/*
DelVideo Delete a video

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id The object id, uuid or short uuid
	@return ApiDelVideoRequest
*/
func (a *VideoAPIService) DelVideo(ctx context.Context, id models.ApiV1VideosOwnershipIdAcceptPostIdParameter) ApiDelVideoRequest {
	return ApiDelVideoRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
func (a *VideoAPIService) DelVideoExecute(r ApiDelVideoRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VideoAPIService.DelVideo")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/videos/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

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

type ApiDeleteVideoSourceFileRequest struct {
	ctx        context.Context
	ApiService *VideoAPIService
	id         models.ApiV1VideosOwnershipIdAcceptPostIdParameter
}

func (r ApiDeleteVideoSourceFileRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteVideoSourceFileExecute(r)
}

/*
DeleteVideoSourceFile Delete video source file

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id The object id, uuid or short uuid
	@return ApiDeleteVideoSourceFileRequest
*/
func (a *VideoAPIService) DeleteVideoSourceFile(ctx context.Context, id models.ApiV1VideosOwnershipIdAcceptPostIdParameter) ApiDeleteVideoSourceFileRequest {
	return ApiDeleteVideoSourceFileRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
func (a *VideoAPIService) DeleteVideoSourceFileExecute(r ApiDeleteVideoSourceFileRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VideoAPIService.DeleteVideoSourceFile")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/videos/{id}/source/file"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

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

type ApiGetCategoriesRequest struct {
	ctx        context.Context
	ApiService *VideoAPIService
}

func (r ApiGetCategoriesRequest) Execute() ([]string, *http.Response, error) {
	return r.ApiService.GetCategoriesExecute(r)
}

/*
GetCategories List available video categories

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetCategoriesRequest
*/
func (a *VideoAPIService) GetCategories(ctx context.Context) ApiGetCategoriesRequest {
	return ApiGetCategoriesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return []string
func (a *VideoAPIService) GetCategoriesExecute(r ApiGetCategoriesRequest) ([]string, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VideoAPIService.GetCategories")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/videos/categories"

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

type ApiGetLanguagesRequest struct {
	ctx        context.Context
	ApiService *VideoAPIService
}

func (r ApiGetLanguagesRequest) Execute() ([]string, *http.Response, error) {
	return r.ApiService.GetLanguagesExecute(r)
}

/*
GetLanguages List available video languages

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetLanguagesRequest
*/
func (a *VideoAPIService) GetLanguages(ctx context.Context) ApiGetLanguagesRequest {
	return ApiGetLanguagesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return []string
func (a *VideoAPIService) GetLanguagesExecute(r ApiGetLanguagesRequest) ([]string, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VideoAPIService.GetLanguages")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/videos/languages"

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

type ApiGetLicencesRequest struct {
	ctx        context.Context
	ApiService *VideoAPIService
}

func (r ApiGetLicencesRequest) Execute() ([]string, *http.Response, error) {
	return r.ApiService.GetLicencesExecute(r)
}

/*
GetLicences List available video licences

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetLicencesRequest
*/
func (a *VideoAPIService) GetLicences(ctx context.Context) ApiGetLicencesRequest {
	return ApiGetLicencesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return []string
func (a *VideoAPIService) GetLicencesExecute(r ApiGetLicencesRequest) ([]string, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VideoAPIService.GetLicences")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/videos/licences"

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

type ApiGetVideoRequest struct {
	ctx                    context.Context
	ApiService             *VideoAPIService
	id                     models.ApiV1VideosOwnershipIdAcceptPostIdParameter
	xPeertubeVideoPassword *string
}

// Required on password protected video
func (r ApiGetVideoRequest) XPeertubeVideoPassword(xPeertubeVideoPassword string) ApiGetVideoRequest {
	r.xPeertubeVideoPassword = &xPeertubeVideoPassword
	return r
}

func (r ApiGetVideoRequest) Execute() (*models.VideoDetails, *http.Response, error) {
	return r.ApiService.GetVideoExecute(r)
}

/*
GetVideo Get a video

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id The object id, uuid or short uuid
	@return ApiGetVideoRequest
*/
func (a *VideoAPIService) GetVideo(ctx context.Context, id models.ApiV1VideosOwnershipIdAcceptPostIdParameter) ApiGetVideoRequest {
	return ApiGetVideoRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return VideoDetails
func (a *VideoAPIService) GetVideoExecute(r ApiGetVideoRequest) (*models.VideoDetails, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.VideoDetails
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VideoAPIService.GetVideo")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/videos/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(&r.id, "id")), -1)

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
	if r.xPeertubeVideoPassword != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "x-peertube-video-password", r.xPeertubeVideoPassword, "simple", "")
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

type ApiGetVideoDescRequest struct {
	ctx                    context.Context
	ApiService             *VideoAPIService
	id                     models.ApiV1VideosOwnershipIdAcceptPostIdParameter
	xPeertubeVideoPassword *string
}

// Required on password protected video
func (r ApiGetVideoDescRequest) XPeertubeVideoPassword(xPeertubeVideoPassword string) ApiGetVideoDescRequest {
	r.xPeertubeVideoPassword = &xPeertubeVideoPassword
	return r
}

func (r ApiGetVideoDescRequest) Execute() (string, *http.Response, error) {
	return r.ApiService.GetVideoDescExecute(r)
}

/*
GetVideoDesc Get complete video description

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id The object id, uuid or short uuid
	@return ApiGetVideoDescRequest
*/
func (a *VideoAPIService) GetVideoDesc(ctx context.Context, id models.ApiV1VideosOwnershipIdAcceptPostIdParameter) ApiGetVideoDescRequest {
	return ApiGetVideoDescRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return string
func (a *VideoAPIService) GetVideoDescExecute(r ApiGetVideoDescRequest) (string, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VideoAPIService.GetVideoDesc")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/videos/{id}/description"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

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
	if r.xPeertubeVideoPassword != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "x-peertube-video-password", r.xPeertubeVideoPassword, "simple", "")
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

type ApiGetVideoPrivacyPoliciesRequest struct {
	ctx        context.Context
	ApiService *VideoAPIService
}

func (r ApiGetVideoPrivacyPoliciesRequest) Execute() ([]string, *http.Response, error) {
	return r.ApiService.GetVideoPrivacyPoliciesExecute(r)
}

/*
GetVideoPrivacyPolicies List available video privacy policies

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetVideoPrivacyPoliciesRequest
*/
func (a *VideoAPIService) GetVideoPrivacyPolicies(ctx context.Context) ApiGetVideoPrivacyPoliciesRequest {
	return ApiGetVideoPrivacyPoliciesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return []string
func (a *VideoAPIService) GetVideoPrivacyPoliciesExecute(r ApiGetVideoPrivacyPoliciesRequest) ([]string, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VideoAPIService.GetVideoPrivacyPolicies")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/videos/privacies"

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

type ApiGetVideoSourceRequest struct {
	ctx        context.Context
	ApiService *VideoAPIService
	id         models.ApiV1VideosOwnershipIdAcceptPostIdParameter
}

func (r ApiGetVideoSourceRequest) Execute() (*models.VideoSource, *http.Response, error) {
	return r.ApiService.GetVideoSourceExecute(r)
}

/*
GetVideoSource Get video source file metadata

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id The object id, uuid or short uuid
	@return ApiGetVideoSourceRequest
*/
func (a *VideoAPIService) GetVideoSource(ctx context.Context, id models.ApiV1VideosOwnershipIdAcceptPostIdParameter) ApiGetVideoSourceRequest {
	return ApiGetVideoSourceRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return VideoSource
func (a *VideoAPIService) GetVideoSourceExecute(r ApiGetVideoSourceRequest) (*models.VideoSource, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.VideoSource
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VideoAPIService.GetVideoSource")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/videos/{id}/source"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

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

type ApiGetVideosRequest struct {
	ctx                   context.Context
	ApiService            *VideoAPIService
	categoryOneOf         *models.GetAccountVideosCategoryOneOfParameter
	isLive                *bool
	tagsOneOf             *models.GetAccountVideosTagsOneOfParameter
	tagsAllOf             *models.GetAccountVideosTagsAllOfParameter
	licenceOneOf          *models.GetAccountVideosLicenceOneOfParameter
	languageOneOf         *models.GetAccountVideosLanguageOneOfParameter
	autoTagOneOf          *models.GetAccountVideosTagsAllOfParameter
	nsfw                  *string
	isLocal               *bool
	include               *int32
	privacyOneOf          *models.VideoPrivacySet
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
func (r ApiGetVideosRequest) CategoryOneOf(categoryOneOf models.GetAccountVideosCategoryOneOfParameter) ApiGetVideosRequest {
	r.categoryOneOf = &categoryOneOf
	return r
}

// whether or not the video is a live
func (r ApiGetVideosRequest) IsLive(isLive bool) ApiGetVideosRequest {
	r.isLive = &isLive
	return r
}

// tag(s) of the video
func (r ApiGetVideosRequest) TagsOneOf(tagsOneOf models.GetAccountVideosTagsOneOfParameter) ApiGetVideosRequest {
	r.tagsOneOf = &tagsOneOf
	return r
}

// tag(s) of the video, where all should be present in the video
func (r ApiGetVideosRequest) TagsAllOf(tagsAllOf models.GetAccountVideosTagsAllOfParameter) ApiGetVideosRequest {
	r.tagsAllOf = &tagsAllOf
	return r
}

// licence id of the video (see [/videos/licences](#operation/getLicences))
func (r ApiGetVideosRequest) LicenceOneOf(licenceOneOf models.GetAccountVideosLicenceOneOfParameter) ApiGetVideosRequest {
	r.licenceOneOf = &licenceOneOf
	return r
}

// language id of the video (see [/videos/languages](#operation/getLanguages)). Use &#x60;_unknown&#x60; to filter on videos that don&#39;t have a video language
func (r ApiGetVideosRequest) LanguageOneOf(languageOneOf models.GetAccountVideosLanguageOneOfParameter) ApiGetVideosRequest {
	r.languageOneOf = &languageOneOf
	return r
}

// **PeerTube &gt;&#x3D; 6.2** **Admins and moderators only** filter on videos that contain one of these automatic tags
func (r ApiGetVideosRequest) AutoTagOneOf(autoTagOneOf models.GetAccountVideosTagsAllOfParameter) ApiGetVideosRequest {
	r.autoTagOneOf = &autoTagOneOf
	return r
}

// whether to include nsfw videos, if any
func (r ApiGetVideosRequest) Nsfw(nsfw string) ApiGetVideosRequest {
	r.nsfw = &nsfw
	return r
}

// **PeerTube &gt;&#x3D; 4.0** Display only local or remote objects
func (r ApiGetVideosRequest) IsLocal(isLocal bool) ApiGetVideosRequest {
	r.isLocal = &isLocal
	return r
}

// **Only administrators and moderators can use this parameter**  Include additional videos in results (can be combined using bitwise or operator) - &#x60;0&#x60; NONE - &#x60;1&#x60; NOT_PUBLISHED_STATE - &#x60;2&#x60; BLACKLISTED - &#x60;4&#x60; BLOCKED_OWNER - &#x60;8&#x60; FILES - &#x60;16&#x60; CAPTIONS - &#x60;32&#x60; VIDEO SOURCE
func (r ApiGetVideosRequest) Include(include int32) ApiGetVideosRequest {
	r.include = &include
	return r
}

// **PeerTube &gt;&#x3D; 4.0** Display only videos in this specific privacy/privacies
func (r ApiGetVideosRequest) PrivacyOneOf(privacyOneOf models.VideoPrivacySet) ApiGetVideosRequest {
	r.privacyOneOf = &privacyOneOf
	return r
}

// **PeerTube &gt;&#x3D; 4.0** Display only videos that have HLS files
func (r ApiGetVideosRequest) HasHLSFiles(hasHLSFiles bool) ApiGetVideosRequest {
	r.hasHLSFiles = &hasHLSFiles
	return r
}

// **PeerTube &gt;&#x3D; 6.0** Display only videos that have Web Video files
func (r ApiGetVideosRequest) HasWebVideoFiles(hasWebVideoFiles bool) ApiGetVideosRequest {
	r.hasWebVideoFiles = &hasWebVideoFiles
	return r
}

// if you don&#39;t need the &#x60;total&#x60; in the response
func (r ApiGetVideosRequest) SkipCount(skipCount string) ApiGetVideosRequest {
	r.skipCount = &skipCount
	return r
}

// Offset used to paginate results
func (r ApiGetVideosRequest) Start(start int32) ApiGetVideosRequest {
	r.start = &start
	return r
}

// Number of items to return
func (r ApiGetVideosRequest) Count(count int32) ApiGetVideosRequest {
	r.count = &count
	return r
}

func (r ApiGetVideosRequest) Sort(sort string) ApiGetVideosRequest {
	r.sort = &sort
	return r
}

// Whether or not to exclude videos that are in the user&#39;s video history
func (r ApiGetVideosRequest) ExcludeAlreadyWatched(excludeAlreadyWatched bool) ApiGetVideosRequest {
	r.excludeAlreadyWatched = &excludeAlreadyWatched
	return r
}

// Plain text search, applied to various parts of the model depending on endpoint
func (r ApiGetVideosRequest) Search(search string) ApiGetVideosRequest {
	r.search = &search
	return r
}

func (r ApiGetVideosRequest) Execute() (*models.VideoListResponse, *http.Response, error) {
	return r.ApiService.GetVideosExecute(r)
}

/*
GetVideos List videos

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetVideosRequest
*/
func (a *VideoAPIService) GetVideos(ctx context.Context) ApiGetVideosRequest {
	return ApiGetVideosRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return VideoListResponse
func (a *VideoAPIService) GetVideosExecute(r ApiGetVideosRequest) (*models.VideoListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.VideoListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VideoAPIService.GetVideos")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/videos"

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

type ApiListVideoStoryboardsRequest struct {
	ctx        context.Context
	ApiService *VideoAPIService
	id         models.ApiV1VideosOwnershipIdAcceptPostIdParameter
}

func (r ApiListVideoStoryboardsRequest) Execute() (*models.ListVideoStoryboards200Response, *http.Response, error) {
	return r.ApiService.ListVideoStoryboardsExecute(r)
}

/*
ListVideoStoryboards List storyboards of a video

**PeerTube >= 6.0**

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id The object id, uuid or short uuid
	@return ApiListVideoStoryboardsRequest
*/
func (a *VideoAPIService) ListVideoStoryboards(ctx context.Context, id models.ApiV1VideosOwnershipIdAcceptPostIdParameter) ApiListVideoStoryboardsRequest {
	return ApiListVideoStoryboardsRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return ListVideoStoryboards200Response
func (a *VideoAPIService) ListVideoStoryboardsExecute(r ApiListVideoStoryboardsRequest) (*models.ListVideoStoryboards200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.ListVideoStoryboards200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VideoAPIService.ListVideoStoryboards")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/videos/{id}/storyboards"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

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

type ApiPutVideoRequest struct {
	ctx                   context.Context
	ApiService            *VideoAPIService
	id                    models.ApiV1VideosOwnershipIdAcceptPostIdParameter
	thumbnailfile         *os.File
	previewfile           *os.File
	category              *int32
	licence               *int32
	language              *string
	privacy               *models.VideoPrivacySet
	description           *string
	waitTranscoding       *string
	support               *string
	nsfw                  *bool
	name                  *string
	tags                  *[]string
	commentsEnabled       *bool
	commentsPolicy        *models.VideoCommentsPolicySet
	downloadEnabled       *bool
	originallyPublishedAt *time.Time
	scheduleUpdate        *models.VideoScheduledUpdate
	videoPasswords        *models.VideoPassword
}

// Video thumbnail file
func (r ApiPutVideoRequest) Thumbnailfile(thumbnailfile *os.File) ApiPutVideoRequest {
	r.thumbnailfile = thumbnailfile
	return r
}

// Video preview file
func (r ApiPutVideoRequest) Previewfile(previewfile *os.File) ApiPutVideoRequest {
	r.previewfile = previewfile
	return r
}

// category id of the video (see [/videos/categories](#operation/getCategories))
func (r ApiPutVideoRequest) Category(category int32) ApiPutVideoRequest {
	r.category = &category
	return r
}

// licence id of the video (see [/videos/licences](#operation/getLicences))
func (r ApiPutVideoRequest) Licence(licence int32) ApiPutVideoRequest {
	r.licence = &licence
	return r
}

// language id of the video (see [/videos/languages](#operation/getLanguages))
func (r ApiPutVideoRequest) Language(language string) ApiPutVideoRequest {
	r.language = &language
	return r
}

func (r ApiPutVideoRequest) Privacy(privacy models.VideoPrivacySet) ApiPutVideoRequest {
	r.privacy = &privacy
	return r
}

// Video description
func (r ApiPutVideoRequest) Description(description string) ApiPutVideoRequest {
	r.description = &description
	return r
}

// Whether or not we wait transcoding before publish the video
func (r ApiPutVideoRequest) WaitTranscoding(waitTranscoding string) ApiPutVideoRequest {
	r.waitTranscoding = &waitTranscoding
	return r
}

// A text tell the audience how to support the video creator
func (r ApiPutVideoRequest) Support(support string) ApiPutVideoRequest {
	r.support = &support
	return r
}

// Whether or not this video contains sensitive content
func (r ApiPutVideoRequest) Nsfw(nsfw bool) ApiPutVideoRequest {
	r.nsfw = &nsfw
	return r
}

// Video name
func (r ApiPutVideoRequest) Name(name string) ApiPutVideoRequest {
	r.name = &name
	return r
}

// Video tags (maximum 5 tags each between 2 and 30 characters)
func (r ApiPutVideoRequest) Tags(tags []string) ApiPutVideoRequest {
	r.tags = &tags
	return r
}

// Deprecated in 6.2, use commentsPolicy instead
func (r ApiPutVideoRequest) CommentsEnabled(commentsEnabled bool) ApiPutVideoRequest {
	r.commentsEnabled = &commentsEnabled
	return r
}

func (r ApiPutVideoRequest) CommentsPolicy(commentsPolicy models.VideoCommentsPolicySet) ApiPutVideoRequest {
	r.commentsPolicy = &commentsPolicy
	return r
}

// Enable or disable downloading for this video
func (r ApiPutVideoRequest) DownloadEnabled(downloadEnabled bool) ApiPutVideoRequest {
	r.downloadEnabled = &downloadEnabled
	return r
}

// Date when the content was originally published
func (r ApiPutVideoRequest) OriginallyPublishedAt(originallyPublishedAt time.Time) ApiPutVideoRequest {
	r.originallyPublishedAt = &originallyPublishedAt
	return r
}

func (r ApiPutVideoRequest) ScheduleUpdate(scheduleUpdate models.VideoScheduledUpdate) ApiPutVideoRequest {
	r.scheduleUpdate = &scheduleUpdate
	return r
}

func (r ApiPutVideoRequest) VideoPasswords(videoPasswords models.VideoPassword) ApiPutVideoRequest {
	r.videoPasswords = &videoPasswords
	return r
}

func (r ApiPutVideoRequest) Execute() (*http.Response, error) {
	return r.ApiService.PutVideoExecute(r)
}

/*
PutVideo Update a video

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id The object id, uuid or short uuid
	@return ApiPutVideoRequest
*/
func (a *VideoAPIService) PutVideo(ctx context.Context, id models.ApiV1VideosOwnershipIdAcceptPostIdParameter) ApiPutVideoRequest {
	return ApiPutVideoRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
func (a *VideoAPIService) PutVideoExecute(r ApiPutVideoRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPut
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VideoAPIService.PutVideo")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/videos/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

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
	var thumbnailfileLocalVarFormFileName string
	var thumbnailfileLocalVarFileName string
	var thumbnailfileLocalVarFileBytes []byte

	thumbnailfileLocalVarFormFileName = "thumbnailfile"
	thumbnailfileLocalVarFile := r.thumbnailfile

	if thumbnailfileLocalVarFile != nil {
		fbs, _ := io.ReadAll(thumbnailfileLocalVarFile)

		thumbnailfileLocalVarFileBytes = fbs
		thumbnailfileLocalVarFileName = thumbnailfileLocalVarFile.Name()
		thumbnailfileLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: thumbnailfileLocalVarFileBytes, fileName: thumbnailfileLocalVarFileName, formFileName: thumbnailfileLocalVarFormFileName})
	}
	var previewfileLocalVarFormFileName string
	var previewfileLocalVarFileName string
	var previewfileLocalVarFileBytes []byte

	previewfileLocalVarFormFileName = "previewfile"
	previewfileLocalVarFile := r.previewfile

	if previewfileLocalVarFile != nil {
		fbs, _ := io.ReadAll(previewfileLocalVarFile)

		previewfileLocalVarFileBytes = fbs
		previewfileLocalVarFileName = previewfileLocalVarFile.Name()
		previewfileLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: previewfileLocalVarFileBytes, fileName: previewfileLocalVarFileName, formFileName: previewfileLocalVarFormFileName})
	}
	if r.category != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "category", r.category, "", "")
	}
	if r.licence != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "licence", r.licence, "", "")
	}
	if r.language != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "language", r.language, "", "")
	}
	if r.privacy != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "privacy", r.privacy, "", "")
	}
	if r.description != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "description", r.description, "", "")
	}
	if r.waitTranscoding != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "waitTranscoding", r.waitTranscoding, "", "")
	}
	if r.support != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "support", r.support, "", "")
	}
	if r.nsfw != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "nsfw", r.nsfw, "", "")
	}
	if r.name != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "name", r.name, "", "")
	}
	if r.tags != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "tags", r.tags, "", "csv")
	}
	if r.commentsEnabled != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "commentsEnabled", r.commentsEnabled, "", "")
	}
	if r.commentsPolicy != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "commentsPolicy", r.commentsPolicy, "", "")
	}
	if r.downloadEnabled != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "downloadEnabled", r.downloadEnabled, "", "")
	}
	if r.originallyPublishedAt != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "originallyPublishedAt", r.originallyPublishedAt, "", "")
	}
	if r.scheduleUpdate != nil {
		paramJson, err := parameterToJson(*r.scheduleUpdate)
		if err != nil {
			return nil, err
		}
		localVarFormParams.Add("scheduleUpdate", paramJson)
	}
	if r.videoPasswords != nil {
		paramJson, err := parameterToJson(*r.videoPasswords)
		if err != nil {
			return nil, err
		}
		localVarFormParams.Add("videoPasswords", paramJson)
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

type ApiRequestVideoTokenRequest struct {
	ctx                    context.Context
	ApiService             *VideoAPIService
	id                     models.ApiV1VideosOwnershipIdAcceptPostIdParameter
	xPeertubeVideoPassword *string
}

// Required on password protected video
func (r ApiRequestVideoTokenRequest) XPeertubeVideoPassword(xPeertubeVideoPassword string) ApiRequestVideoTokenRequest {
	r.xPeertubeVideoPassword = &xPeertubeVideoPassword
	return r
}

func (r ApiRequestVideoTokenRequest) Execute() (*models.VideoTokenResponse, *http.Response, error) {
	return r.ApiService.RequestVideoTokenExecute(r)
}

/*
RequestVideoToken Request video token

Request special tokens that expire quickly to use them in some context (like accessing private static files)

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id The object id, uuid or short uuid
	@return ApiRequestVideoTokenRequest
*/
func (a *VideoAPIService) RequestVideoToken(ctx context.Context, id models.ApiV1VideosOwnershipIdAcceptPostIdParameter) ApiRequestVideoTokenRequest {
	return ApiRequestVideoTokenRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return VideoTokenResponse
func (a *VideoAPIService) RequestVideoTokenExecute(r ApiRequestVideoTokenRequest) (*models.VideoTokenResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.VideoTokenResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VideoAPIService.RequestVideoToken")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/videos/{id}/token"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

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
	if r.xPeertubeVideoPassword != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "x-peertube-video-password", r.xPeertubeVideoPassword, "simple", "")
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

type ApiUploadLegacyRequest struct {
	ctx                   context.Context
	ApiService            *VideoAPIService
	name                  *string
	channelId             *int32
	videofile             *os.File
	privacy               *models.VideoPrivacySet
	category              *int32
	licence               *int32
	language              *string
	description           *string
	waitTranscoding       *bool
	generateTranscription *bool
	support               *string
	nsfw                  *bool
	tags                  *[]string
	commentsEnabled       *bool
	commentsPolicy        *models.VideoCommentsPolicySet
	downloadEnabled       *bool
	originallyPublishedAt *time.Time
	scheduleUpdate        *models.VideoScheduledUpdate
	thumbnailfile         *os.File
	previewfile           *os.File
	videoPasswords        *[]string
}

// Video name
func (r ApiUploadLegacyRequest) Name(name string) ApiUploadLegacyRequest {
	r.name = &name
	return r
}

// Channel id that will contain this video
func (r ApiUploadLegacyRequest) ChannelId(channelId int32) ApiUploadLegacyRequest {
	r.channelId = &channelId
	return r
}

// Video file
func (r ApiUploadLegacyRequest) Videofile(videofile *os.File) ApiUploadLegacyRequest {
	r.videofile = videofile
	return r
}

func (r ApiUploadLegacyRequest) Privacy(privacy models.VideoPrivacySet) ApiUploadLegacyRequest {
	r.privacy = &privacy
	return r
}

// category id of the video (see [/videos/categories](#operation/getCategories))
func (r ApiUploadLegacyRequest) Category(category int32) ApiUploadLegacyRequest {
	r.category = &category
	return r
}

// licence id of the video (see [/videos/licences](#operation/getLicences))
func (r ApiUploadLegacyRequest) Licence(licence int32) ApiUploadLegacyRequest {
	r.licence = &licence
	return r
}

// language id of the video (see [/videos/languages](#operation/getLanguages))
func (r ApiUploadLegacyRequest) Language(language string) ApiUploadLegacyRequest {
	r.language = &language
	return r
}

// Video description
func (r ApiUploadLegacyRequest) Description(description string) ApiUploadLegacyRequest {
	r.description = &description
	return r
}

// Whether or not we wait transcoding before publish the video
func (r ApiUploadLegacyRequest) WaitTranscoding(waitTranscoding bool) ApiUploadLegacyRequest {
	r.waitTranscoding = &waitTranscoding
	return r
}

// **PeerTube &gt;&#x3D; 6.2** If enabled by the admin, automatically generate a subtitle of the video
func (r ApiUploadLegacyRequest) GenerateTranscription(generateTranscription bool) ApiUploadLegacyRequest {
	r.generateTranscription = &generateTranscription
	return r
}

// A text tell the audience how to support the video creator
func (r ApiUploadLegacyRequest) Support(support string) ApiUploadLegacyRequest {
	r.support = &support
	return r
}

// Whether or not this video contains sensitive content
func (r ApiUploadLegacyRequest) Nsfw(nsfw bool) ApiUploadLegacyRequest {
	r.nsfw = &nsfw
	return r
}

// Video tags (maximum 5 tags each between 2 and 30 characters)
func (r ApiUploadLegacyRequest) Tags(tags []string) ApiUploadLegacyRequest {
	r.tags = &tags
	return r
}

// Deprecated in 6.2, use commentsPolicy instead
func (r ApiUploadLegacyRequest) CommentsEnabled(commentsEnabled bool) ApiUploadLegacyRequest {
	r.commentsEnabled = &commentsEnabled
	return r
}

func (r ApiUploadLegacyRequest) CommentsPolicy(commentsPolicy models.VideoCommentsPolicySet) ApiUploadLegacyRequest {
	r.commentsPolicy = &commentsPolicy
	return r
}

// Enable or disable downloading for this video
func (r ApiUploadLegacyRequest) DownloadEnabled(downloadEnabled bool) ApiUploadLegacyRequest {
	r.downloadEnabled = &downloadEnabled
	return r
}

// Date when the content was originally published
func (r ApiUploadLegacyRequest) OriginallyPublishedAt(originallyPublishedAt time.Time) ApiUploadLegacyRequest {
	r.originallyPublishedAt = &originallyPublishedAt
	return r
}

func (r ApiUploadLegacyRequest) ScheduleUpdate(scheduleUpdate models.VideoScheduledUpdate) ApiUploadLegacyRequest {
	r.scheduleUpdate = &scheduleUpdate
	return r
}

// Video thumbnail file
func (r ApiUploadLegacyRequest) Thumbnailfile(thumbnailfile *os.File) ApiUploadLegacyRequest {
	r.thumbnailfile = thumbnailfile
	return r
}

// Video preview file
func (r ApiUploadLegacyRequest) Previewfile(previewfile *os.File) ApiUploadLegacyRequest {
	r.previewfile = previewfile
	return r
}

func (r ApiUploadLegacyRequest) VideoPasswords(videoPasswords []string) ApiUploadLegacyRequest {
	r.videoPasswords = &videoPasswords
	return r
}

func (r ApiUploadLegacyRequest) Execute() (*models.VideoUploadResponse, *http.Response, error) {
	return r.ApiService.UploadLegacyExecute(r)
}

/*
UploadLegacy Upload a video

Uses a single request to upload a video.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiUploadLegacyRequest
*/
func (a *VideoAPIService) UploadLegacy(ctx context.Context) ApiUploadLegacyRequest {
	return ApiUploadLegacyRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return VideoUploadResponse
func (a *VideoAPIService) UploadLegacyExecute(r ApiUploadLegacyRequest) (*models.VideoUploadResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.VideoUploadResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VideoAPIService.UploadLegacy")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/videos/upload"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.name == nil {
		return localVarReturnValue, nil, utils.ReportError("name is required and must be specified")
	}
	if strlen(*r.name) < 3 {
		return localVarReturnValue, nil, utils.ReportError("name must have at least 3 elements")
	}
	if strlen(*r.name) > 120 {
		return localVarReturnValue, nil, utils.ReportError("name must have less than 120 elements")
	}
	if r.channelId == nil {
		return localVarReturnValue, nil, utils.ReportError("channelId is required and must be specified")
	}
	if *r.channelId < 1 {
		return localVarReturnValue, nil, utils.ReportError("channelId must be greater than 1")
	}
	if r.videofile == nil {
		return localVarReturnValue, nil, utils.ReportError("videofile is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

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
	parameterAddToHeaderOrQuery(localVarFormParams, "name", r.name, "", "")
	parameterAddToHeaderOrQuery(localVarFormParams, "channelId", r.channelId, "", "")
	if r.privacy != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "privacy", r.privacy, "", "")
	}
	if r.category != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "category", r.category, "", "")
	}
	if r.licence != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "licence", r.licence, "", "")
	}
	if r.language != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "language", r.language, "", "")
	}
	if r.description != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "description", r.description, "", "")
	}
	if r.waitTranscoding != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "waitTranscoding", r.waitTranscoding, "", "")
	}
	if r.generateTranscription != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "generateTranscription", r.generateTranscription, "", "")
	}
	if r.support != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "support", r.support, "", "")
	}
	if r.nsfw != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "nsfw", r.nsfw, "", "")
	}
	if r.tags != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "tags", r.tags, "", "csv")
	}
	if r.commentsEnabled != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "commentsEnabled", r.commentsEnabled, "", "")
	}
	if r.commentsPolicy != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "commentsPolicy", r.commentsPolicy, "", "")
	}
	if r.downloadEnabled != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "downloadEnabled", r.downloadEnabled, "", "")
	}
	if r.originallyPublishedAt != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "originallyPublishedAt", r.originallyPublishedAt, "", "")
	}
	if r.scheduleUpdate != nil {
		paramJson, err := parameterToJson(*r.scheduleUpdate)
		if err != nil {
			return localVarReturnValue, nil, err
		}
		localVarFormParams.Add("scheduleUpdate", paramJson)
	}
	var thumbnailfileLocalVarFormFileName string
	var thumbnailfileLocalVarFileName string
	var thumbnailfileLocalVarFileBytes []byte

	thumbnailfileLocalVarFormFileName = "thumbnailfile"
	thumbnailfileLocalVarFile := r.thumbnailfile

	if thumbnailfileLocalVarFile != nil {
		fbs, _ := io.ReadAll(thumbnailfileLocalVarFile)

		thumbnailfileLocalVarFileBytes = fbs
		thumbnailfileLocalVarFileName = thumbnailfileLocalVarFile.Name()
		thumbnailfileLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: thumbnailfileLocalVarFileBytes, fileName: thumbnailfileLocalVarFileName, formFileName: thumbnailfileLocalVarFormFileName})
	}
	var previewfileLocalVarFormFileName string
	var previewfileLocalVarFileName string
	var previewfileLocalVarFileBytes []byte

	previewfileLocalVarFormFileName = "previewfile"
	previewfileLocalVarFile := r.previewfile

	if previewfileLocalVarFile != nil {
		fbs, _ := io.ReadAll(previewfileLocalVarFile)

		previewfileLocalVarFileBytes = fbs
		previewfileLocalVarFileName = previewfileLocalVarFile.Name()
		previewfileLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: previewfileLocalVarFileBytes, fileName: previewfileLocalVarFileName, formFileName: previewfileLocalVarFormFileName})
	}
	if r.videoPasswords != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "videoPasswords", r.videoPasswords, "", "csv")
	}
	var videofileLocalVarFormFileName string
	var videofileLocalVarFileName string
	var videofileLocalVarFileBytes []byte

	videofileLocalVarFormFileName = "videofile"
	videofileLocalVarFile := r.videofile

	if videofileLocalVarFile != nil {
		fbs, _ := io.ReadAll(videofileLocalVarFile)

		videofileLocalVarFileBytes = fbs
		videofileLocalVarFileName = videofileLocalVarFile.Name()
		videofileLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: videofileLocalVarFileBytes, fileName: videofileLocalVarFileName, formFileName: videofileLocalVarFormFileName})
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

type ApiUploadResumableRequest struct {
	ctx           context.Context
	ApiService    *VideoAPIService
	uploadId      *string
	contentRange  *string
	contentLength *float32
	body          *os.File
}

// Created session id to proceed with. If you didn&#39;t send chunks in the last hour, it is not valid anymore and you need to initialize a new upload.
func (r ApiUploadResumableRequest) UploadId(uploadId string) ApiUploadResumableRequest {
	r.uploadId = &uploadId
	return r
}

// Specifies the bytes in the file that the request is uploading.  For example, a value of &#x60;bytes 0-262143/1000000&#x60; shows that the request is sending the first 262144 bytes (256 x 1024) in a 2,469,036 byte file.
func (r ApiUploadResumableRequest) ContentRange(contentRange string) ApiUploadResumableRequest {
	r.contentRange = &contentRange
	return r
}

// Size of the chunk that the request is sending.  Remember that larger chunks are more efficient. PeerTube&#39;s web client uses chunks varying from 1048576 bytes (~1MB) and increases or reduces size depending on connection health.
func (r ApiUploadResumableRequest) ContentLength(contentLength float32) ApiUploadResumableRequest {
	r.contentLength = &contentLength
	return r
}

func (r ApiUploadResumableRequest) Body(body *os.File) ApiUploadResumableRequest {
	r.body = body
	return r
}

func (r ApiUploadResumableRequest) Execute() (*models.VideoUploadResponse, *http.Response, error) {
	return r.ApiService.UploadResumableExecute(r)
}

/*
UploadResumable Send chunk for the resumable upload of a video

Uses [a resumable protocol](https://github.com/kukhariev/node-uploadx/blob/master/proto.md) to continue, pause or resume the upload of a video

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiUploadResumableRequest
*/
func (a *VideoAPIService) UploadResumable(ctx context.Context) ApiUploadResumableRequest {
	return ApiUploadResumableRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return VideoUploadResponse
func (a *VideoAPIService) UploadResumableExecute(r ApiUploadResumableRequest) (*models.VideoUploadResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.VideoUploadResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VideoAPIService.UploadResumable")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/videos/upload-resumable"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.uploadId == nil {
		return localVarReturnValue, nil, utils.ReportError("uploadId is required and must be specified")
	}
	if r.contentRange == nil {
		return localVarReturnValue, nil, utils.ReportError("contentRange is required and must be specified")
	}
	if r.contentLength == nil {
		return localVarReturnValue, nil, utils.ReportError("contentLength is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "upload_id", r.uploadId, "form", "")
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/octet-stream"}

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
	parameterAddToHeaderOrQuery(localVarHeaderParams, "Content-Range", r.contentRange, "simple", "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "Content-Length", r.contentLength, "simple", "")
	// body params
	localVarPostBody = r.body
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

type ApiUploadResumableCancelRequest struct {
	ctx           context.Context
	ApiService    *VideoAPIService
	uploadId      *string
	contentLength *float32
}

// Created session id to proceed with. If you didn&#39;t send chunks in the last hour, it is not valid anymore and you need to initialize a new upload.
func (r ApiUploadResumableCancelRequest) UploadId(uploadId string) ApiUploadResumableCancelRequest {
	r.uploadId = &uploadId
	return r
}

func (r ApiUploadResumableCancelRequest) ContentLength(contentLength float32) ApiUploadResumableCancelRequest {
	r.contentLength = &contentLength
	return r
}

func (r ApiUploadResumableCancelRequest) Execute() (*http.Response, error) {
	return r.ApiService.UploadResumableCancelExecute(r)
}

/*
UploadResumableCancel Cancel the resumable upload of a video, deleting any data uploaded so far

Uses [a resumable protocol](https://github.com/kukhariev/node-uploadx/blob/master/proto.md) to cancel the upload of a video

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiUploadResumableCancelRequest
*/
func (a *VideoAPIService) UploadResumableCancel(ctx context.Context) ApiUploadResumableCancelRequest {
	return ApiUploadResumableCancelRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *VideoAPIService) UploadResumableCancelExecute(r ApiUploadResumableCancelRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VideoAPIService.UploadResumableCancel")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/videos/upload-resumable"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.uploadId == nil {
		return nil, utils.ReportError("uploadId is required and must be specified")
	}
	if r.contentLength == nil {
		return nil, utils.ReportError("contentLength is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "upload_id", r.uploadId, "form", "")
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
	parameterAddToHeaderOrQuery(localVarHeaderParams, "Content-Length", r.contentLength, "simple", "")
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

type ApiUploadResumableInitRequest struct {
	ctx                         context.Context
	ApiService                  *VideoAPIService
	xUploadContentLength        *float32
	xUploadContentType          *string
	videoUploadRequestResumable *models.VideoUploadRequestResumable
}

// Number of bytes that will be uploaded in subsequent requests. Set this value to the size of the file you are uploading.
func (r ApiUploadResumableInitRequest) XUploadContentLength(xUploadContentLength float32) ApiUploadResumableInitRequest {
	r.xUploadContentLength = &xUploadContentLength
	return r
}

// MIME type of the file that you are uploading. Depending on your instance settings, acceptable values might vary.
func (r ApiUploadResumableInitRequest) XUploadContentType(xUploadContentType string) ApiUploadResumableInitRequest {
	r.xUploadContentType = &xUploadContentType
	return r
}

func (r ApiUploadResumableInitRequest) VideoUploadRequestResumable(videoUploadRequestResumable models.VideoUploadRequestResumable) ApiUploadResumableInitRequest {
	r.videoUploadRequestResumable = &videoUploadRequestResumable
	return r
}

func (r ApiUploadResumableInitRequest) Execute() (*http.Response, error) {
	return r.ApiService.UploadResumableInitExecute(r)
}

/*
UploadResumableInit Initialize the resumable upload of a video

Uses [a resumable protocol](https://github.com/kukhariev/node-uploadx/blob/master/proto.md) to initialize the upload of a video

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiUploadResumableInitRequest
*/
func (a *VideoAPIService) UploadResumableInit(ctx context.Context) ApiUploadResumableInitRequest {
	return ApiUploadResumableInitRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *VideoAPIService) UploadResumableInitExecute(r ApiUploadResumableInitRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VideoAPIService.UploadResumableInit")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/videos/upload-resumable"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.xUploadContentLength == nil {
		return nil, utils.ReportError("xUploadContentLength is required and must be specified")
	}
	if r.xUploadContentType == nil {
		return nil, utils.ReportError("xUploadContentType is required and must be specified")
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
	parameterAddToHeaderOrQuery(localVarHeaderParams, "X-Upload-Content-Length", r.xUploadContentLength, "simple", "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "X-Upload-Content-Type", r.xUploadContentType, "simple", "")
	// body params
	localVarPostBody = r.videoUploadRequestResumable
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
