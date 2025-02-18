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

// AccountsAPIService AccountsAPI service
type AccountsAPIService service

type ApiApiV1AccountsNameRatingsGetRequest struct {
	ctx        context.Context
	ApiService *AccountsAPIService
	name       string
	start      *int32
	count      *int32
	sort       *string
	rating     *string
}

// Offset used to paginate results
func (r ApiApiV1AccountsNameRatingsGetRequest) Start(start int32) ApiApiV1AccountsNameRatingsGetRequest {
	r.start = &start
	return r
}

// Number of items to return
func (r ApiApiV1AccountsNameRatingsGetRequest) Count(count int32) ApiApiV1AccountsNameRatingsGetRequest {
	r.count = &count
	return r
}

// Sort column
func (r ApiApiV1AccountsNameRatingsGetRequest) Sort(sort string) ApiApiV1AccountsNameRatingsGetRequest {
	r.sort = &sort
	return r
}

// Optionally filter which ratings to retrieve
func (r ApiApiV1AccountsNameRatingsGetRequest) Rating(rating string) ApiApiV1AccountsNameRatingsGetRequest {
	r.rating = &rating
	return r
}

func (r ApiApiV1AccountsNameRatingsGetRequest) Execute() ([]models.VideoRating, *http.Response, error) {
	return r.ApiService.ApiV1AccountsNameRatingsGetExecute(r)
}

/*
ApiV1AccountsNameRatingsGet List ratings of an account

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param name The username or handle of the account
	@return ApiApiV1AccountsNameRatingsGetRequest
*/
func (a *AccountsAPIService) ApiV1AccountsNameRatingsGet(ctx context.Context, name string) ApiApiV1AccountsNameRatingsGetRequest {
	return ApiApiV1AccountsNameRatingsGetRequest{
		ApiService: a,
		ctx:        ctx,
		name:       name,
	}
}

// Execute executes the request
//
//	@return []VideoRating
func (a *AccountsAPIService) ApiV1AccountsNameRatingsGetExecute(r ApiApiV1AccountsNameRatingsGetRequest) ([]models.VideoRating, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []models.VideoRating
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccountsAPIService.ApiV1AccountsNameRatingsGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/accounts/{name}/ratings"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(parameterValueToString(r.name, "name")), -1)

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
	if r.rating != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "rating", r.rating, "form", "")
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

type ApiApiV1AccountsNameVideoChannelSyncsGetRequest struct {
	ctx        context.Context
	ApiService *AccountsAPIService
	name       string
	start      *int32
	count      *int32
	sort       *string
}

// Offset used to paginate results
func (r ApiApiV1AccountsNameVideoChannelSyncsGetRequest) Start(start int32) ApiApiV1AccountsNameVideoChannelSyncsGetRequest {
	r.start = &start
	return r
}

// Number of items to return
func (r ApiApiV1AccountsNameVideoChannelSyncsGetRequest) Count(count int32) ApiApiV1AccountsNameVideoChannelSyncsGetRequest {
	r.count = &count
	return r
}

// Sort column
func (r ApiApiV1AccountsNameVideoChannelSyncsGetRequest) Sort(sort string) ApiApiV1AccountsNameVideoChannelSyncsGetRequest {
	r.sort = &sort
	return r
}

func (r ApiApiV1AccountsNameVideoChannelSyncsGetRequest) Execute() (*models.VideoChannelSyncList, *http.Response, error) {
	return r.ApiService.ApiV1AccountsNameVideoChannelSyncsGetExecute(r)
}

/*
ApiV1AccountsNameVideoChannelSyncsGet List the synchronizations of video channels of an account

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param name The username or handle of the account
	@return ApiApiV1AccountsNameVideoChannelSyncsGetRequest
*/
func (a *AccountsAPIService) ApiV1AccountsNameVideoChannelSyncsGet(ctx context.Context, name string) ApiApiV1AccountsNameVideoChannelSyncsGetRequest {
	return ApiApiV1AccountsNameVideoChannelSyncsGetRequest{
		ApiService: a,
		ctx:        ctx,
		name:       name,
	}
}

// Execute executes the request
//
//	@return VideoChannelSyncList
func (a *AccountsAPIService) ApiV1AccountsNameVideoChannelSyncsGetExecute(r ApiApiV1AccountsNameVideoChannelSyncsGetRequest) (*models.VideoChannelSyncList, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.VideoChannelSyncList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccountsAPIService.ApiV1AccountsNameVideoChannelSyncsGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/accounts/{name}/video-channel-syncs"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(parameterValueToString(r.name, "name")), -1)

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

type ApiApiV1AccountsNameVideoChannelsGetRequest struct {
	ctx        context.Context
	ApiService *AccountsAPIService
	name       string
	withStats  *bool
	start      *int32
	count      *int32
	sort       *string
}

// include daily view statistics for the last 30 days and total views (only if authentified as the account user)
func (r ApiApiV1AccountsNameVideoChannelsGetRequest) WithStats(withStats bool) ApiApiV1AccountsNameVideoChannelsGetRequest {
	r.withStats = &withStats
	return r
}

// Offset used to paginate results
func (r ApiApiV1AccountsNameVideoChannelsGetRequest) Start(start int32) ApiApiV1AccountsNameVideoChannelsGetRequest {
	r.start = &start
	return r
}

// Number of items to return
func (r ApiApiV1AccountsNameVideoChannelsGetRequest) Count(count int32) ApiApiV1AccountsNameVideoChannelsGetRequest {
	r.count = &count
	return r
}

// Sort column
func (r ApiApiV1AccountsNameVideoChannelsGetRequest) Sort(sort string) ApiApiV1AccountsNameVideoChannelsGetRequest {
	r.sort = &sort
	return r
}

func (r ApiApiV1AccountsNameVideoChannelsGetRequest) Execute() (*models.VideoChannelList, *http.Response, error) {
	return r.ApiService.ApiV1AccountsNameVideoChannelsGetExecute(r)
}

/*
ApiV1AccountsNameVideoChannelsGet List video channels of an account

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param name The username or handle of the account
	@return ApiApiV1AccountsNameVideoChannelsGetRequest
*/
func (a *AccountsAPIService) ApiV1AccountsNameVideoChannelsGet(ctx context.Context, name string) ApiApiV1AccountsNameVideoChannelsGetRequest {
	return ApiApiV1AccountsNameVideoChannelsGetRequest{
		ApiService: a,
		ctx:        ctx,
		name:       name,
	}
}

// Execute executes the request
//
//	@return VideoChannelList
func (a *AccountsAPIService) ApiV1AccountsNameVideoChannelsGetExecute(r ApiApiV1AccountsNameVideoChannelsGetRequest) (*models.VideoChannelList, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.VideoChannelList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccountsAPIService.ApiV1AccountsNameVideoChannelsGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/accounts/{name}/video-channels"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(parameterValueToString(r.name, "name")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.withStats != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "withStats", r.withStats, "form", "")
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

type ApiApiV1AccountsNameVideoPlaylistsGetRequest struct {
	ctx          context.Context
	ApiService   *AccountsAPIService
	name         string
	start        *int32
	count        *int32
	sort         *string
	search       *string
	playlistType *models.VideoPlaylistTypeSet
}

// Offset used to paginate results
func (r ApiApiV1AccountsNameVideoPlaylistsGetRequest) Start(start int32) ApiApiV1AccountsNameVideoPlaylistsGetRequest {
	r.start = &start
	return r
}

// Number of items to return
func (r ApiApiV1AccountsNameVideoPlaylistsGetRequest) Count(count int32) ApiApiV1AccountsNameVideoPlaylistsGetRequest {
	r.count = &count
	return r
}

// Sort column
func (r ApiApiV1AccountsNameVideoPlaylistsGetRequest) Sort(sort string) ApiApiV1AccountsNameVideoPlaylistsGetRequest {
	r.sort = &sort
	return r
}

// Plain text search, applied to various parts of the model depending on endpoint
func (r ApiApiV1AccountsNameVideoPlaylistsGetRequest) Search(search string) ApiApiV1AccountsNameVideoPlaylistsGetRequest {
	r.search = &search
	return r
}

func (r ApiApiV1AccountsNameVideoPlaylistsGetRequest) PlaylistType(playlistType models.VideoPlaylistTypeSet) ApiApiV1AccountsNameVideoPlaylistsGetRequest {
	r.playlistType = &playlistType
	return r
}

func (r ApiApiV1AccountsNameVideoPlaylistsGetRequest) Execute() (*models.ApiV1VideoChannelsChannelHandleVideoPlaylistsGet200Response, *http.Response, error) {
	return r.ApiService.ApiV1AccountsNameVideoPlaylistsGetExecute(r)
}

/*
ApiV1AccountsNameVideoPlaylistsGet List playlists of an account

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param name The username or handle of the account
	@return ApiApiV1AccountsNameVideoPlaylistsGetRequest
*/
func (a *AccountsAPIService) ApiV1AccountsNameVideoPlaylistsGet(ctx context.Context, name string) ApiApiV1AccountsNameVideoPlaylistsGetRequest {
	return ApiApiV1AccountsNameVideoPlaylistsGetRequest{
		ApiService: a,
		ctx:        ctx,
		name:       name,
	}
}

// Execute executes the request
//
//	@return ApiV1VideoChannelsChannelHandleVideoPlaylistsGet200Response
func (a *AccountsAPIService) ApiV1AccountsNameVideoPlaylistsGetExecute(r ApiApiV1AccountsNameVideoPlaylistsGetRequest) (*models.ApiV1VideoChannelsChannelHandleVideoPlaylistsGet200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.ApiV1VideoChannelsChannelHandleVideoPlaylistsGet200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccountsAPIService.ApiV1AccountsNameVideoPlaylistsGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/accounts/{name}/video-playlists"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(parameterValueToString(r.name, "name")), -1)

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
	if r.search != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search", r.search, "form", "")
	}
	if r.playlistType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "playlistType", r.playlistType, "form", "")
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

type ApiGetAccountRequest struct {
	ctx        context.Context
	ApiService *AccountsAPIService
	name       string
}

func (r ApiGetAccountRequest) Execute() (*models.Account, *http.Response, error) {
	return r.ApiService.GetAccountExecute(r)
}

/*
GetAccount Get an account

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param name The username or handle of the account
	@return ApiGetAccountRequest
*/
func (a *AccountsAPIService) GetAccount(ctx context.Context, name string) ApiGetAccountRequest {
	return ApiGetAccountRequest{
		ApiService: a,
		ctx:        ctx,
		name:       name,
	}
}

// Execute executes the request
//
//	@return Account
func (a *AccountsAPIService) GetAccountExecute(r ApiGetAccountRequest) (*models.Account, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.Account
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccountsAPIService.GetAccount")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/accounts/{name}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(parameterValueToString(r.name, "name")), -1)

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

type ApiGetAccountFollowersRequest struct {
	ctx        context.Context
	ApiService *AccountsAPIService
	name       string
	start      *int32
	count      *int32
	sort       *string
	search     *string
}

// Offset used to paginate results
func (r ApiGetAccountFollowersRequest) Start(start int32) ApiGetAccountFollowersRequest {
	r.start = &start
	return r
}

// Number of items to return
func (r ApiGetAccountFollowersRequest) Count(count int32) ApiGetAccountFollowersRequest {
	r.count = &count
	return r
}

// Sort followers by criteria
func (r ApiGetAccountFollowersRequest) Sort(sort string) ApiGetAccountFollowersRequest {
	r.sort = &sort
	return r
}

// Plain text search, applied to various parts of the model depending on endpoint
func (r ApiGetAccountFollowersRequest) Search(search string) ApiGetAccountFollowersRequest {
	r.search = &search
	return r
}

func (r ApiGetAccountFollowersRequest) Execute() (*models.GetAccountFollowers200Response, *http.Response, error) {
	return r.ApiService.GetAccountFollowersExecute(r)
}

/*
GetAccountFollowers List followers of an account

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param name The username or handle of the account
	@return ApiGetAccountFollowersRequest
*/
func (a *AccountsAPIService) GetAccountFollowers(ctx context.Context, name string) ApiGetAccountFollowersRequest {
	return ApiGetAccountFollowersRequest{
		ApiService: a,
		ctx:        ctx,
		name:       name,
	}
}

// Execute executes the request
//
//	@return GetAccountFollowers200Response
func (a *AccountsAPIService) GetAccountFollowersExecute(r ApiGetAccountFollowersRequest) (*models.GetAccountFollowers200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.GetAccountFollowers200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccountsAPIService.GetAccountFollowers")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/accounts/{name}/followers"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(parameterValueToString(r.name, "name")), -1)

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

type ApiGetAccountVideosRequest struct {
	ctx                   context.Context
	ApiService            *AccountsAPIService
	name                  string
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
func (r ApiGetAccountVideosRequest) CategoryOneOf(categoryOneOf models.GetAccountVideosCategoryOneOfParameter) ApiGetAccountVideosRequest {
	r.categoryOneOf = &categoryOneOf
	return r
}

// whether or not the video is a live
func (r ApiGetAccountVideosRequest) IsLive(isLive bool) ApiGetAccountVideosRequest {
	r.isLive = &isLive
	return r
}

// tag(s) of the video
func (r ApiGetAccountVideosRequest) TagsOneOf(tagsOneOf models.GetAccountVideosTagsOneOfParameter) ApiGetAccountVideosRequest {
	r.tagsOneOf = &tagsOneOf
	return r
}

// tag(s) of the video, where all should be present in the video
func (r ApiGetAccountVideosRequest) TagsAllOf(tagsAllOf models.GetAccountVideosTagsAllOfParameter) ApiGetAccountVideosRequest {
	r.tagsAllOf = &tagsAllOf
	return r
}

// licence id of the video (see [/videos/licences](#operation/getLicences))
func (r ApiGetAccountVideosRequest) LicenceOneOf(licenceOneOf models.GetAccountVideosLicenceOneOfParameter) ApiGetAccountVideosRequest {
	r.licenceOneOf = &licenceOneOf
	return r
}

// language id of the video (see [/videos/languages](#operation/getLanguages)). Use &#x60;_unknown&#x60; to filter on videos that don&#39;t have a video language
func (r ApiGetAccountVideosRequest) LanguageOneOf(languageOneOf models.GetAccountVideosLanguageOneOfParameter) ApiGetAccountVideosRequest {
	r.languageOneOf = &languageOneOf
	return r
}

// **PeerTube &gt;&#x3D; 6.2** **Admins and moderators only** filter on videos that contain one of these automatic tags
func (r ApiGetAccountVideosRequest) AutoTagOneOf(autoTagOneOf models.GetAccountVideosTagsAllOfParameter) ApiGetAccountVideosRequest {
	r.autoTagOneOf = &autoTagOneOf
	return r
}

// whether to include nsfw videos, if any
func (r ApiGetAccountVideosRequest) Nsfw(nsfw string) ApiGetAccountVideosRequest {
	r.nsfw = &nsfw
	return r
}

// **PeerTube &gt;&#x3D; 4.0** Display only local or remote objects
func (r ApiGetAccountVideosRequest) IsLocal(isLocal bool) ApiGetAccountVideosRequest {
	r.isLocal = &isLocal
	return r
}

// **Only administrators and moderators can use this parameter**  Include additional videos in results (can be combined using bitwise or operator) - &#x60;0&#x60; NONE - &#x60;1&#x60; NOT_PUBLISHED_STATE - &#x60;2&#x60; BLACKLISTED - &#x60;4&#x60; BLOCKED_OWNER - &#x60;8&#x60; FILES - &#x60;16&#x60; CAPTIONS - &#x60;32&#x60; VIDEO SOURCE
func (r ApiGetAccountVideosRequest) Include(include int32) ApiGetAccountVideosRequest {
	r.include = &include
	return r
}

// **PeerTube &gt;&#x3D; 4.0** Display only videos in this specific privacy/privacies
func (r ApiGetAccountVideosRequest) PrivacyOneOf(privacyOneOf models.VideoPrivacySet) ApiGetAccountVideosRequest {
	r.privacyOneOf = &privacyOneOf
	return r
}

// **PeerTube &gt;&#x3D; 4.0** Display only videos that have HLS files
func (r ApiGetAccountVideosRequest) HasHLSFiles(hasHLSFiles bool) ApiGetAccountVideosRequest {
	r.hasHLSFiles = &hasHLSFiles
	return r
}

// **PeerTube &gt;&#x3D; 6.0** Display only videos that have Web Video files
func (r ApiGetAccountVideosRequest) HasWebVideoFiles(hasWebVideoFiles bool) ApiGetAccountVideosRequest {
	r.hasWebVideoFiles = &hasWebVideoFiles
	return r
}

// if you don&#39;t need the &#x60;total&#x60; in the response
func (r ApiGetAccountVideosRequest) SkipCount(skipCount string) ApiGetAccountVideosRequest {
	r.skipCount = &skipCount
	return r
}

// Offset used to paginate results
func (r ApiGetAccountVideosRequest) Start(start int32) ApiGetAccountVideosRequest {
	r.start = &start
	return r
}

// Number of items to return
func (r ApiGetAccountVideosRequest) Count(count int32) ApiGetAccountVideosRequest {
	r.count = &count
	return r
}

func (r ApiGetAccountVideosRequest) Sort(sort string) ApiGetAccountVideosRequest {
	r.sort = &sort
	return r
}

// Whether or not to exclude videos that are in the user&#39;s video history
func (r ApiGetAccountVideosRequest) ExcludeAlreadyWatched(excludeAlreadyWatched bool) ApiGetAccountVideosRequest {
	r.excludeAlreadyWatched = &excludeAlreadyWatched
	return r
}

// Plain text search, applied to various parts of the model depending on endpoint
func (r ApiGetAccountVideosRequest) Search(search string) ApiGetAccountVideosRequest {
	r.search = &search
	return r
}

func (r ApiGetAccountVideosRequest) Execute() (*models.VideoListResponse, *http.Response, error) {
	return r.ApiService.GetAccountVideosExecute(r)
}

/*
GetAccountVideos List videos of an account

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param name The username or handle of the account
	@return ApiGetAccountVideosRequest
*/
func (a *AccountsAPIService) GetAccountVideos(ctx context.Context, name string) ApiGetAccountVideosRequest {
	return ApiGetAccountVideosRequest{
		ApiService: a,
		ctx:        ctx,
		name:       name,
	}
}

// Execute executes the request
//
//	@return VideoListResponse
func (a *AccountsAPIService) GetAccountVideosExecute(r ApiGetAccountVideosRequest) (*models.VideoListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.VideoListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccountsAPIService.GetAccountVideos")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/accounts/{name}/videos"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(parameterValueToString(r.name, "name")), -1)

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

type ApiGetAccountsRequest struct {
	ctx        context.Context
	ApiService *AccountsAPIService
	start      *int32
	count      *int32
	sort       *string
}

// Offset used to paginate results
func (r ApiGetAccountsRequest) Start(start int32) ApiGetAccountsRequest {
	r.start = &start
	return r
}

// Number of items to return
func (r ApiGetAccountsRequest) Count(count int32) ApiGetAccountsRequest {
	r.count = &count
	return r
}

// Sort column
func (r ApiGetAccountsRequest) Sort(sort string) ApiGetAccountsRequest {
	r.sort = &sort
	return r
}

func (r ApiGetAccountsRequest) Execute() (*models.GetAccounts200Response, *http.Response, error) {
	return r.ApiService.GetAccountsExecute(r)
}

/*
GetAccounts List accounts

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetAccountsRequest
*/
func (a *AccountsAPIService) GetAccounts(ctx context.Context) ApiGetAccountsRequest {
	return ApiGetAccountsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetAccounts200Response
func (a *AccountsAPIService) GetAccountsExecute(r ApiGetAccountsRequest) (*models.GetAccounts200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.GetAccounts200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccountsAPIService.GetAccounts")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/accounts"

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
