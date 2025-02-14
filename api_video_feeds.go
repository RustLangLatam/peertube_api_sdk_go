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

// VideoFeedsAPIService VideoFeedsAPI service
type VideoFeedsAPIService service

type ApiGetSyndicatedCommentsRequest struct {
	ctx              context.Context
	ApiService       *VideoFeedsAPIService
	format           string
	videoId          *string
	accountId        *string
	accountName      *string
	videoChannelId   *string
	videoChannelName *string
}

// limit listing comments to a specific video
func (r ApiGetSyndicatedCommentsRequest) VideoId(videoId string) ApiGetSyndicatedCommentsRequest {
	r.videoId = &videoId
	return r
}

// limit listing comments to videos of a specific account
func (r ApiGetSyndicatedCommentsRequest) AccountId(accountId string) ApiGetSyndicatedCommentsRequest {
	r.accountId = &accountId
	return r
}

// limit listing comments to videos of a specific account
func (r ApiGetSyndicatedCommentsRequest) AccountName(accountName string) ApiGetSyndicatedCommentsRequest {
	r.accountName = &accountName
	return r
}

// limit listing comments to videos of a specific video channel
func (r ApiGetSyndicatedCommentsRequest) VideoChannelId(videoChannelId string) ApiGetSyndicatedCommentsRequest {
	r.videoChannelId = &videoChannelId
	return r
}

// limit listing comments to videos of a specific video channel
func (r ApiGetSyndicatedCommentsRequest) VideoChannelName(videoChannelName string) ApiGetSyndicatedCommentsRequest {
	r.videoChannelName = &videoChannelName
	return r
}

func (r ApiGetSyndicatedCommentsRequest) Execute() ([]VideoCommentsForXMLInner, *http.Response, error) {
	return r.ApiService.GetSyndicatedCommentsExecute(r)
}

/*
GetSyndicatedComments Comments on videos feeds

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param format format expected (we focus on making `rss` the most featureful ; it serves [Media RSS](https://www.rssboard.org/media-rss))
	@return ApiGetSyndicatedCommentsRequest
*/
func (a *VideoFeedsAPIService) GetSyndicatedComments(ctx context.Context, format string) ApiGetSyndicatedCommentsRequest {
	return ApiGetSyndicatedCommentsRequest{
		ApiService: a,
		ctx:        ctx,
		format:     format,
	}
}

// Execute executes the request
//
//	@return []VideoCommentsForXMLInner
func (a *VideoFeedsAPIService) GetSyndicatedCommentsExecute(r ApiGetSyndicatedCommentsRequest) ([]VideoCommentsForXMLInner, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []VideoCommentsForXMLInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VideoFeedsAPIService.GetSyndicatedComments")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/feeds/video-comments.{format}"
	localVarPath = strings.Replace(localVarPath, "{"+"format"+"}", url.PathEscape(parameterValueToString(r.format, "format")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.videoId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "videoId", r.videoId, "form", "")
	}
	if r.accountId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "accountId", r.accountId, "form", "")
	}
	if r.accountName != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "accountName", r.accountName, "form", "")
	}
	if r.videoChannelId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "videoChannelId", r.videoChannelId, "form", "")
	}
	if r.videoChannelName != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "videoChannelName", r.videoChannelName, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/xml", "application/rss+xml", "text/xml", "application/atom+xml", "application/json"}

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

type ApiGetSyndicatedSubscriptionVideosRequest struct {
	ctx              context.Context
	ApiService       *VideoFeedsAPIService
	format           string
	accountId        *string
	token            *string
	sort             *string
	nsfw             *string
	isLocal          *bool
	include          *int32
	privacyOneOf     *VideoPrivacySet
	hasHLSFiles      *bool
	hasWebVideoFiles *bool
}

// limit listing to a specific account
func (r ApiGetSyndicatedSubscriptionVideosRequest) AccountId(accountId string) ApiGetSyndicatedSubscriptionVideosRequest {
	r.accountId = &accountId
	return r
}

// private token allowing access
func (r ApiGetSyndicatedSubscriptionVideosRequest) Token(token string) ApiGetSyndicatedSubscriptionVideosRequest {
	r.token = &token
	return r
}

// Sort column
func (r ApiGetSyndicatedSubscriptionVideosRequest) Sort(sort string) ApiGetSyndicatedSubscriptionVideosRequest {
	r.sort = &sort
	return r
}

// whether to include nsfw videos, if any
func (r ApiGetSyndicatedSubscriptionVideosRequest) Nsfw(nsfw string) ApiGetSyndicatedSubscriptionVideosRequest {
	r.nsfw = &nsfw
	return r
}

// **PeerTube &gt;&#x3D; 4.0** Display only local or remote objects
func (r ApiGetSyndicatedSubscriptionVideosRequest) IsLocal(isLocal bool) ApiGetSyndicatedSubscriptionVideosRequest {
	r.isLocal = &isLocal
	return r
}

// **Only administrators and moderators can use this parameter**  Include additional videos in results (can be combined using bitwise or operator) - &#x60;0&#x60; NONE - &#x60;1&#x60; NOT_PUBLISHED_STATE - &#x60;2&#x60; BLACKLISTED - &#x60;4&#x60; BLOCKED_OWNER - &#x60;8&#x60; FILES - &#x60;16&#x60; CAPTIONS - &#x60;32&#x60; VIDEO SOURCE
func (r ApiGetSyndicatedSubscriptionVideosRequest) Include(include int32) ApiGetSyndicatedSubscriptionVideosRequest {
	r.include = &include
	return r
}

// **PeerTube &gt;&#x3D; 4.0** Display only videos in this specific privacy/privacies
func (r ApiGetSyndicatedSubscriptionVideosRequest) PrivacyOneOf(privacyOneOf VideoPrivacySet) ApiGetSyndicatedSubscriptionVideosRequest {
	r.privacyOneOf = &privacyOneOf
	return r
}

// **PeerTube &gt;&#x3D; 4.0** Display only videos that have HLS files
func (r ApiGetSyndicatedSubscriptionVideosRequest) HasHLSFiles(hasHLSFiles bool) ApiGetSyndicatedSubscriptionVideosRequest {
	r.hasHLSFiles = &hasHLSFiles
	return r
}

// **PeerTube &gt;&#x3D; 6.0** Display only videos that have Web Video files
func (r ApiGetSyndicatedSubscriptionVideosRequest) HasWebVideoFiles(hasWebVideoFiles bool) ApiGetSyndicatedSubscriptionVideosRequest {
	r.hasWebVideoFiles = &hasWebVideoFiles
	return r
}

func (r ApiGetSyndicatedSubscriptionVideosRequest) Execute() ([]VideosForXMLInner, *http.Response, error) {
	return r.ApiService.GetSyndicatedSubscriptionVideosExecute(r)
}

/*
GetSyndicatedSubscriptionVideos Videos of subscriptions feeds

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param format format expected (we focus on making `rss` the most featureful ; it serves [Media RSS](https://www.rssboard.org/media-rss))
	@return ApiGetSyndicatedSubscriptionVideosRequest
*/
func (a *VideoFeedsAPIService) GetSyndicatedSubscriptionVideos(ctx context.Context, format string) ApiGetSyndicatedSubscriptionVideosRequest {
	return ApiGetSyndicatedSubscriptionVideosRequest{
		ApiService: a,
		ctx:        ctx,
		format:     format,
	}
}

// Execute executes the request
//
//	@return []VideosForXMLInner
func (a *VideoFeedsAPIService) GetSyndicatedSubscriptionVideosExecute(r ApiGetSyndicatedSubscriptionVideosRequest) ([]VideosForXMLInner, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []VideosForXMLInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VideoFeedsAPIService.GetSyndicatedSubscriptionVideos")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/feeds/subscriptions.{format}"
	localVarPath = strings.Replace(localVarPath, "{"+"format"+"}", url.PathEscape(parameterValueToString(r.format, "format")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.accountId == nil {
		return localVarReturnValue, nil, reportError("accountId is required and must be specified")
	}
	if r.token == nil {
		return localVarReturnValue, nil, reportError("token is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "accountId", r.accountId, "form", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "token", r.token, "form", "")
	if r.sort != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort", r.sort, "form", "")
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
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/xml", "application/rss+xml", "text/xml", "application/atom+xml", "application/json"}

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

type ApiGetSyndicatedVideosRequest struct {
	ctx              context.Context
	ApiService       *VideoFeedsAPIService
	format           string
	accountId        *string
	accountName      *string
	videoChannelId   *string
	videoChannelName *string
	sort             *string
	nsfw             *string
	isLocal          *bool
	include          *int32
	privacyOneOf     *VideoPrivacySet
	hasHLSFiles      *bool
	hasWebVideoFiles *bool
}

// limit listing to a specific account
func (r ApiGetSyndicatedVideosRequest) AccountId(accountId string) ApiGetSyndicatedVideosRequest {
	r.accountId = &accountId
	return r
}

// limit listing to a specific account
func (r ApiGetSyndicatedVideosRequest) AccountName(accountName string) ApiGetSyndicatedVideosRequest {
	r.accountName = &accountName
	return r
}

// limit listing to a specific video channel
func (r ApiGetSyndicatedVideosRequest) VideoChannelId(videoChannelId string) ApiGetSyndicatedVideosRequest {
	r.videoChannelId = &videoChannelId
	return r
}

// limit listing to a specific video channel
func (r ApiGetSyndicatedVideosRequest) VideoChannelName(videoChannelName string) ApiGetSyndicatedVideosRequest {
	r.videoChannelName = &videoChannelName
	return r
}

// Sort column
func (r ApiGetSyndicatedVideosRequest) Sort(sort string) ApiGetSyndicatedVideosRequest {
	r.sort = &sort
	return r
}

// whether to include nsfw videos, if any
func (r ApiGetSyndicatedVideosRequest) Nsfw(nsfw string) ApiGetSyndicatedVideosRequest {
	r.nsfw = &nsfw
	return r
}

// **PeerTube &gt;&#x3D; 4.0** Display only local or remote objects
func (r ApiGetSyndicatedVideosRequest) IsLocal(isLocal bool) ApiGetSyndicatedVideosRequest {
	r.isLocal = &isLocal
	return r
}

// **Only administrators and moderators can use this parameter**  Include additional videos in results (can be combined using bitwise or operator) - &#x60;0&#x60; NONE - &#x60;1&#x60; NOT_PUBLISHED_STATE - &#x60;2&#x60; BLACKLISTED - &#x60;4&#x60; BLOCKED_OWNER - &#x60;8&#x60; FILES - &#x60;16&#x60; CAPTIONS - &#x60;32&#x60; VIDEO SOURCE
func (r ApiGetSyndicatedVideosRequest) Include(include int32) ApiGetSyndicatedVideosRequest {
	r.include = &include
	return r
}

// **PeerTube &gt;&#x3D; 4.0** Display only videos in this specific privacy/privacies
func (r ApiGetSyndicatedVideosRequest) PrivacyOneOf(privacyOneOf VideoPrivacySet) ApiGetSyndicatedVideosRequest {
	r.privacyOneOf = &privacyOneOf
	return r
}

// **PeerTube &gt;&#x3D; 4.0** Display only videos that have HLS files
func (r ApiGetSyndicatedVideosRequest) HasHLSFiles(hasHLSFiles bool) ApiGetSyndicatedVideosRequest {
	r.hasHLSFiles = &hasHLSFiles
	return r
}

// **PeerTube &gt;&#x3D; 6.0** Display only videos that have Web Video files
func (r ApiGetSyndicatedVideosRequest) HasWebVideoFiles(hasWebVideoFiles bool) ApiGetSyndicatedVideosRequest {
	r.hasWebVideoFiles = &hasWebVideoFiles
	return r
}

func (r ApiGetSyndicatedVideosRequest) Execute() ([]VideosForXMLInner, *http.Response, error) {
	return r.ApiService.GetSyndicatedVideosExecute(r)
}

/*
GetSyndicatedVideos Common videos feeds

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param format format expected (we focus on making `rss` the most featureful ; it serves [Media RSS](https://www.rssboard.org/media-rss))
	@return ApiGetSyndicatedVideosRequest
*/
func (a *VideoFeedsAPIService) GetSyndicatedVideos(ctx context.Context, format string) ApiGetSyndicatedVideosRequest {
	return ApiGetSyndicatedVideosRequest{
		ApiService: a,
		ctx:        ctx,
		format:     format,
	}
}

// Execute executes the request
//
//	@return []VideosForXMLInner
func (a *VideoFeedsAPIService) GetSyndicatedVideosExecute(r ApiGetSyndicatedVideosRequest) ([]VideosForXMLInner, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []VideosForXMLInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VideoFeedsAPIService.GetSyndicatedVideos")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/feeds/videos.{format}"
	localVarPath = strings.Replace(localVarPath, "{"+"format"+"}", url.PathEscape(parameterValueToString(r.format, "format")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.accountId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "accountId", r.accountId, "form", "")
	}
	if r.accountName != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "accountName", r.accountName, "form", "")
	}
	if r.videoChannelId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "videoChannelId", r.videoChannelId, "form", "")
	}
	if r.videoChannelName != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "videoChannelName", r.videoChannelName, "form", "")
	}
	if r.sort != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort", r.sort, "form", "")
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
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/xml", "application/rss+xml", "text/xml", "application/atom+xml", "application/json"}

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

type ApiGetVideosPodcastFeedRequest struct {
	ctx            context.Context
	ApiService     *VideoFeedsAPIService
	videoChannelId *string
}

// Limit listing to a specific video channel
func (r ApiGetVideosPodcastFeedRequest) VideoChannelId(videoChannelId string) ApiGetVideosPodcastFeedRequest {
	r.videoChannelId = &videoChannelId
	return r
}

func (r ApiGetVideosPodcastFeedRequest) Execute() (*http.Response, error) {
	return r.ApiService.GetVideosPodcastFeedExecute(r)
}

/*
GetVideosPodcastFeed Videos podcast feed

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetVideosPodcastFeedRequest
*/
func (a *VideoFeedsAPIService) GetVideosPodcastFeed(ctx context.Context) ApiGetVideosPodcastFeedRequest {
	return ApiGetVideosPodcastFeedRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *VideoFeedsAPIService) GetVideosPodcastFeedExecute(r ApiGetVideosPodcastFeedRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VideoFeedsAPIService.GetVideosPodcastFeed")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/feeds/podcast/videos.xml"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.videoChannelId == nil {
		return nil, reportError("videoChannelId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "videoChannelId", r.videoChannelId, "form", "")
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
