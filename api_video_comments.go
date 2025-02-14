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
	"strings"
)

// VideoCommentsAPIService VideoCommentsAPI service
type VideoCommentsAPIService service

type ApiApiV1UsersMeVideosCommentsGetRequest struct {
	ctx             context.Context
	ApiService      *VideoCommentsAPIService
	search          *string
	searchAccount   *string
	searchVideo     *string
	videoId         *int32
	videoChannelId  *int32
	autoTagOneOf    *GetAccountVideosTagsAllOfParameter
	isHeldForReview *bool
}

// Plain text search, applied to various parts of the model depending on endpoint
func (r ApiApiV1UsersMeVideosCommentsGetRequest) Search(search string) ApiApiV1UsersMeVideosCommentsGetRequest {
	r.search = &search
	return r
}

// Filter comments by searching on the account
func (r ApiApiV1UsersMeVideosCommentsGetRequest) SearchAccount(searchAccount string) ApiApiV1UsersMeVideosCommentsGetRequest {
	r.searchAccount = &searchAccount
	return r
}

// Filter comments by searching on the video
func (r ApiApiV1UsersMeVideosCommentsGetRequest) SearchVideo(searchVideo string) ApiApiV1UsersMeVideosCommentsGetRequest {
	r.searchVideo = &searchVideo
	return r
}

// Limit results on this specific video
func (r ApiApiV1UsersMeVideosCommentsGetRequest) VideoId(videoId int32) ApiApiV1UsersMeVideosCommentsGetRequest {
	r.videoId = &videoId
	return r
}

// Limit results on this specific video channel
func (r ApiApiV1UsersMeVideosCommentsGetRequest) VideoChannelId(videoChannelId int32) ApiApiV1UsersMeVideosCommentsGetRequest {
	r.videoChannelId = &videoChannelId
	return r
}

// **PeerTube &gt;&#x3D; 6.2** filter on comments that contain one of these automatic tags
func (r ApiApiV1UsersMeVideosCommentsGetRequest) AutoTagOneOf(autoTagOneOf GetAccountVideosTagsAllOfParameter) ApiApiV1UsersMeVideosCommentsGetRequest {
	r.autoTagOneOf = &autoTagOneOf
	return r
}

// only display comments that are held for review
func (r ApiApiV1UsersMeVideosCommentsGetRequest) IsHeldForReview(isHeldForReview bool) ApiApiV1UsersMeVideosCommentsGetRequest {
	r.isHeldForReview = &isHeldForReview
	return r
}

func (r ApiApiV1UsersMeVideosCommentsGetRequest) Execute() (*ApiV1UsersMeVideosCommentsGet200Response, *http.Response, error) {
	return r.ApiService.ApiV1UsersMeVideosCommentsGetExecute(r)
}

/*
ApiV1UsersMeVideosCommentsGet List comments on user's videos

**PeerTube >= 6.2**

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiApiV1UsersMeVideosCommentsGetRequest
*/
func (a *VideoCommentsAPIService) ApiV1UsersMeVideosCommentsGet(ctx context.Context) ApiApiV1UsersMeVideosCommentsGetRequest {
	return ApiApiV1UsersMeVideosCommentsGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ApiV1UsersMeVideosCommentsGet200Response
func (a *VideoCommentsAPIService) ApiV1UsersMeVideosCommentsGetExecute(r ApiApiV1UsersMeVideosCommentsGetRequest) (*ApiV1UsersMeVideosCommentsGet200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ApiV1UsersMeVideosCommentsGet200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VideoCommentsAPIService.ApiV1UsersMeVideosCommentsGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/users/me/videos/comments"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.search != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search", r.search, "form", "")
	}
	if r.searchAccount != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "searchAccount", r.searchAccount, "form", "")
	}
	if r.searchVideo != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "searchVideo", r.searchVideo, "form", "")
	}
	if r.videoId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "videoId", r.videoId, "form", "")
	}
	if r.videoChannelId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "videoChannelId", r.videoChannelId, "form", "")
	}
	if r.autoTagOneOf != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "autoTagOneOf", r.autoTagOneOf, "form", "")
	}
	if r.isHeldForReview != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "isHeldForReview", r.isHeldForReview, "form", "")
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

type ApiApiV1VideosCommentsGetRequest struct {
	ctx            context.Context
	ApiService     *VideoCommentsAPIService
	search         *string
	searchAccount  *string
	searchVideo    *string
	videoId        *int32
	videoChannelId *int32
	autoTagOneOf   *GetAccountVideosTagsAllOfParameter
	isLocal        *bool
	onLocalVideo   *bool
}

// Plain text search, applied to various parts of the model depending on endpoint
func (r ApiApiV1VideosCommentsGetRequest) Search(search string) ApiApiV1VideosCommentsGetRequest {
	r.search = &search
	return r
}

// Filter comments by searching on the account
func (r ApiApiV1VideosCommentsGetRequest) SearchAccount(searchAccount string) ApiApiV1VideosCommentsGetRequest {
	r.searchAccount = &searchAccount
	return r
}

// Filter comments by searching on the video
func (r ApiApiV1VideosCommentsGetRequest) SearchVideo(searchVideo string) ApiApiV1VideosCommentsGetRequest {
	r.searchVideo = &searchVideo
	return r
}

// Limit results on this specific video
func (r ApiApiV1VideosCommentsGetRequest) VideoId(videoId int32) ApiApiV1VideosCommentsGetRequest {
	r.videoId = &videoId
	return r
}

// Limit results on this specific video channel
func (r ApiApiV1VideosCommentsGetRequest) VideoChannelId(videoChannelId int32) ApiApiV1VideosCommentsGetRequest {
	r.videoChannelId = &videoChannelId
	return r
}

// **PeerTube &gt;&#x3D; 6.2** filter on comments that contain one of these automatic tags
func (r ApiApiV1VideosCommentsGetRequest) AutoTagOneOf(autoTagOneOf GetAccountVideosTagsAllOfParameter) ApiApiV1VideosCommentsGetRequest {
	r.autoTagOneOf = &autoTagOneOf
	return r
}

// **PeerTube &gt;&#x3D; 4.0** Display only local or remote objects
func (r ApiApiV1VideosCommentsGetRequest) IsLocal(isLocal bool) ApiApiV1VideosCommentsGetRequest {
	r.isLocal = &isLocal
	return r
}

// Display only objects of local or remote videos
func (r ApiApiV1VideosCommentsGetRequest) OnLocalVideo(onLocalVideo bool) ApiApiV1VideosCommentsGetRequest {
	r.onLocalVideo = &onLocalVideo
	return r
}

func (r ApiApiV1VideosCommentsGetRequest) Execute() (*ApiV1UsersMeVideosCommentsGet200Response, *http.Response, error) {
	return r.ApiService.ApiV1VideosCommentsGetExecute(r)
}

/*
ApiV1VideosCommentsGet List instance comments

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiApiV1VideosCommentsGetRequest
*/
func (a *VideoCommentsAPIService) ApiV1VideosCommentsGet(ctx context.Context) ApiApiV1VideosCommentsGetRequest {
	return ApiApiV1VideosCommentsGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ApiV1UsersMeVideosCommentsGet200Response
func (a *VideoCommentsAPIService) ApiV1VideosCommentsGetExecute(r ApiApiV1VideosCommentsGetRequest) (*ApiV1UsersMeVideosCommentsGet200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ApiV1UsersMeVideosCommentsGet200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VideoCommentsAPIService.ApiV1VideosCommentsGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/videos/comments"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.search != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search", r.search, "form", "")
	}
	if r.searchAccount != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "searchAccount", r.searchAccount, "form", "")
	}
	if r.searchVideo != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "searchVideo", r.searchVideo, "form", "")
	}
	if r.videoId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "videoId", r.videoId, "form", "")
	}
	if r.videoChannelId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "videoChannelId", r.videoChannelId, "form", "")
	}
	if r.autoTagOneOf != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "autoTagOneOf", r.autoTagOneOf, "form", "")
	}
	if r.isLocal != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "isLocal", r.isLocal, "form", "")
	}
	if r.onLocalVideo != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "onLocalVideo", r.onLocalVideo, "form", "")
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

type ApiApiV1VideosIdCommentThreadsGetRequest struct {
	ctx                    context.Context
	ApiService             *VideoCommentsAPIService
	id                     ApiV1VideosOwnershipIdAcceptPostIdParameter
	start                  *int32
	count                  *int32
	sort                   *string
	xPeertubeVideoPassword *string
}

// Offset used to paginate results
func (r ApiApiV1VideosIdCommentThreadsGetRequest) Start(start int32) ApiApiV1VideosIdCommentThreadsGetRequest {
	r.start = &start
	return r
}

// Number of items to return
func (r ApiApiV1VideosIdCommentThreadsGetRequest) Count(count int32) ApiApiV1VideosIdCommentThreadsGetRequest {
	r.count = &count
	return r
}

// Sort comments by criteria
func (r ApiApiV1VideosIdCommentThreadsGetRequest) Sort(sort string) ApiApiV1VideosIdCommentThreadsGetRequest {
	r.sort = &sort
	return r
}

// Required on password protected video
func (r ApiApiV1VideosIdCommentThreadsGetRequest) XPeertubeVideoPassword(xPeertubeVideoPassword string) ApiApiV1VideosIdCommentThreadsGetRequest {
	r.xPeertubeVideoPassword = &xPeertubeVideoPassword
	return r
}

func (r ApiApiV1VideosIdCommentThreadsGetRequest) Execute() (*CommentThreadResponse, *http.Response, error) {
	return r.ApiService.ApiV1VideosIdCommentThreadsGetExecute(r)
}

/*
ApiV1VideosIdCommentThreadsGet List threads of a video

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id The object id, uuid or short uuid
	@return ApiApiV1VideosIdCommentThreadsGetRequest
*/
func (a *VideoCommentsAPIService) ApiV1VideosIdCommentThreadsGet(ctx context.Context, id ApiV1VideosOwnershipIdAcceptPostIdParameter) ApiApiV1VideosIdCommentThreadsGetRequest {
	return ApiApiV1VideosIdCommentThreadsGetRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return CommentThreadResponse
func (a *VideoCommentsAPIService) ApiV1VideosIdCommentThreadsGetExecute(r ApiApiV1VideosIdCommentThreadsGetRequest) (*CommentThreadResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CommentThreadResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VideoCommentsAPIService.ApiV1VideosIdCommentThreadsGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/videos/{id}/comment-threads"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

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

type ApiApiV1VideosIdCommentThreadsPostRequest struct {
	ctx                                    context.Context
	ApiService                             *VideoCommentsAPIService
	id                                     ApiV1VideosOwnershipIdAcceptPostIdParameter
	apiV1VideosIdCommentThreadsPostRequest *ApiV1VideosIdCommentThreadsPostRequest
}

func (r ApiApiV1VideosIdCommentThreadsPostRequest) ApiV1VideosIdCommentThreadsPostRequest(apiV1VideosIdCommentThreadsPostRequest ApiV1VideosIdCommentThreadsPostRequest) ApiApiV1VideosIdCommentThreadsPostRequest {
	r.apiV1VideosIdCommentThreadsPostRequest = &apiV1VideosIdCommentThreadsPostRequest
	return r
}

func (r ApiApiV1VideosIdCommentThreadsPostRequest) Execute() (*CommentThreadPostResponse, *http.Response, error) {
	return r.ApiService.ApiV1VideosIdCommentThreadsPostExecute(r)
}

/*
ApiV1VideosIdCommentThreadsPost Create a thread

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id The object id, uuid or short uuid
	@return ApiApiV1VideosIdCommentThreadsPostRequest
*/
func (a *VideoCommentsAPIService) ApiV1VideosIdCommentThreadsPost(ctx context.Context, id ApiV1VideosOwnershipIdAcceptPostIdParameter) ApiApiV1VideosIdCommentThreadsPostRequest {
	return ApiApiV1VideosIdCommentThreadsPostRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return CommentThreadPostResponse
func (a *VideoCommentsAPIService) ApiV1VideosIdCommentThreadsPostExecute(r ApiApiV1VideosIdCommentThreadsPostRequest) (*CommentThreadPostResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CommentThreadPostResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VideoCommentsAPIService.ApiV1VideosIdCommentThreadsPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/videos/{id}/comment-threads"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.apiV1VideosIdCommentThreadsPostRequest
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

type ApiApiV1VideosIdCommentThreadsThreadIdGetRequest struct {
	ctx                    context.Context
	ApiService             *VideoCommentsAPIService
	id                     ApiV1VideosOwnershipIdAcceptPostIdParameter
	threadId               int32
	xPeertubeVideoPassword *string
}

// Required on password protected video
func (r ApiApiV1VideosIdCommentThreadsThreadIdGetRequest) XPeertubeVideoPassword(xPeertubeVideoPassword string) ApiApiV1VideosIdCommentThreadsThreadIdGetRequest {
	r.xPeertubeVideoPassword = &xPeertubeVideoPassword
	return r
}

func (r ApiApiV1VideosIdCommentThreadsThreadIdGetRequest) Execute() (*VideoCommentThreadTree, *http.Response, error) {
	return r.ApiService.ApiV1VideosIdCommentThreadsThreadIdGetExecute(r)
}

/*
ApiV1VideosIdCommentThreadsThreadIdGet Get a thread

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id The object id, uuid or short uuid
	@param threadId The thread id (root comment id)
	@return ApiApiV1VideosIdCommentThreadsThreadIdGetRequest
*/
func (a *VideoCommentsAPIService) ApiV1VideosIdCommentThreadsThreadIdGet(ctx context.Context, id ApiV1VideosOwnershipIdAcceptPostIdParameter, threadId int32) ApiApiV1VideosIdCommentThreadsThreadIdGetRequest {
	return ApiApiV1VideosIdCommentThreadsThreadIdGetRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
		threadId:   threadId,
	}
}

// Execute executes the request
//
//	@return VideoCommentThreadTree
func (a *VideoCommentsAPIService) ApiV1VideosIdCommentThreadsThreadIdGetExecute(r ApiApiV1VideosIdCommentThreadsThreadIdGetRequest) (*VideoCommentThreadTree, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *VideoCommentThreadTree
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VideoCommentsAPIService.ApiV1VideosIdCommentThreadsThreadIdGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/videos/{id}/comment-threads/{threadId}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"threadId"+"}", url.PathEscape(parameterValueToString(r.threadId, "threadId")), -1)

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

type ApiApiV1VideosIdCommentsCommentIdApprovePostRequest struct {
	ctx        context.Context
	ApiService *VideoCommentsAPIService
	id         ApiV1VideosOwnershipIdAcceptPostIdParameter
	commentId  int32
}

func (r ApiApiV1VideosIdCommentsCommentIdApprovePostRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiV1VideosIdCommentsCommentIdApprovePostExecute(r)
}

/*
ApiV1VideosIdCommentsCommentIdApprovePost Approve a comment

**PeerTube >= 6.2** Approve a comment that requires a review

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id The object id, uuid or short uuid
	@param commentId The comment id
	@return ApiApiV1VideosIdCommentsCommentIdApprovePostRequest
*/
func (a *VideoCommentsAPIService) ApiV1VideosIdCommentsCommentIdApprovePost(ctx context.Context, id ApiV1VideosOwnershipIdAcceptPostIdParameter, commentId int32) ApiApiV1VideosIdCommentsCommentIdApprovePostRequest {
	return ApiApiV1VideosIdCommentsCommentIdApprovePostRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
		commentId:  commentId,
	}
}

// Execute executes the request
func (a *VideoCommentsAPIService) ApiV1VideosIdCommentsCommentIdApprovePostExecute(r ApiApiV1VideosIdCommentsCommentIdApprovePostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VideoCommentsAPIService.ApiV1VideosIdCommentsCommentIdApprovePost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/videos/{id}/comments/{commentId}/approve"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"commentId"+"}", url.PathEscape(parameterValueToString(r.commentId, "commentId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.commentId < 1 {
		return nil, reportError("commentId must be greater than 1")
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

type ApiApiV1VideosIdCommentsCommentIdDeleteRequest struct {
	ctx        context.Context
	ApiService *VideoCommentsAPIService
	id         ApiV1VideosOwnershipIdAcceptPostIdParameter
	commentId  int32
}

func (r ApiApiV1VideosIdCommentsCommentIdDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiV1VideosIdCommentsCommentIdDeleteExecute(r)
}

/*
ApiV1VideosIdCommentsCommentIdDelete Delete a comment or a reply

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id The object id, uuid or short uuid
	@param commentId The comment id
	@return ApiApiV1VideosIdCommentsCommentIdDeleteRequest
*/
func (a *VideoCommentsAPIService) ApiV1VideosIdCommentsCommentIdDelete(ctx context.Context, id ApiV1VideosOwnershipIdAcceptPostIdParameter, commentId int32) ApiApiV1VideosIdCommentsCommentIdDeleteRequest {
	return ApiApiV1VideosIdCommentsCommentIdDeleteRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
		commentId:  commentId,
	}
}

// Execute executes the request
func (a *VideoCommentsAPIService) ApiV1VideosIdCommentsCommentIdDeleteExecute(r ApiApiV1VideosIdCommentsCommentIdDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VideoCommentsAPIService.ApiV1VideosIdCommentsCommentIdDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/videos/{id}/comments/{commentId}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"commentId"+"}", url.PathEscape(parameterValueToString(r.commentId, "commentId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.commentId < 1 {
		return nil, reportError("commentId must be greater than 1")
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

type ApiApiV1VideosIdCommentsCommentIdPostRequest struct {
	ctx                                    context.Context
	ApiService                             *VideoCommentsAPIService
	id                                     ApiV1VideosOwnershipIdAcceptPostIdParameter
	commentId                              int32
	apiV1VideosIdCommentThreadsPostRequest *ApiV1VideosIdCommentThreadsPostRequest
	xPeertubeVideoPassword                 *string
}

func (r ApiApiV1VideosIdCommentsCommentIdPostRequest) ApiV1VideosIdCommentThreadsPostRequest(apiV1VideosIdCommentThreadsPostRequest ApiV1VideosIdCommentThreadsPostRequest) ApiApiV1VideosIdCommentsCommentIdPostRequest {
	r.apiV1VideosIdCommentThreadsPostRequest = &apiV1VideosIdCommentThreadsPostRequest
	return r
}

// Required on password protected video
func (r ApiApiV1VideosIdCommentsCommentIdPostRequest) XPeertubeVideoPassword(xPeertubeVideoPassword string) ApiApiV1VideosIdCommentsCommentIdPostRequest {
	r.xPeertubeVideoPassword = &xPeertubeVideoPassword
	return r
}

func (r ApiApiV1VideosIdCommentsCommentIdPostRequest) Execute() (*CommentThreadPostResponse, *http.Response, error) {
	return r.ApiService.ApiV1VideosIdCommentsCommentIdPostExecute(r)
}

/*
ApiV1VideosIdCommentsCommentIdPost Reply to a thread of a video

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id The object id, uuid or short uuid
	@param commentId The comment id
	@return ApiApiV1VideosIdCommentsCommentIdPostRequest
*/
func (a *VideoCommentsAPIService) ApiV1VideosIdCommentsCommentIdPost(ctx context.Context, id ApiV1VideosOwnershipIdAcceptPostIdParameter, commentId int32) ApiApiV1VideosIdCommentsCommentIdPostRequest {
	return ApiApiV1VideosIdCommentsCommentIdPostRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
		commentId:  commentId,
	}
}

// Execute executes the request
//
//	@return CommentThreadPostResponse
func (a *VideoCommentsAPIService) ApiV1VideosIdCommentsCommentIdPostExecute(r ApiApiV1VideosIdCommentsCommentIdPostRequest) (*CommentThreadPostResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CommentThreadPostResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VideoCommentsAPIService.ApiV1VideosIdCommentsCommentIdPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/videos/{id}/comments/{commentId}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"commentId"+"}", url.PathEscape(parameterValueToString(r.commentId, "commentId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.commentId < 1 {
		return localVarReturnValue, nil, reportError("commentId must be greater than 1")
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
	if r.xPeertubeVideoPassword != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "x-peertube-video-password", r.xPeertubeVideoPassword, "simple", "")
	}
	// body params
	localVarPostBody = r.apiV1VideosIdCommentThreadsPostRequest
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
