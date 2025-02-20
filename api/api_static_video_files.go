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
	"io"
	"net/http"
	"net/url"
	"strings"
)

// StaticVideoFilesAPIService StaticVideoFilesAPI service
type StaticVideoFilesAPIService service

type ApiStaticStreamingPlaylistsHlsFilenameGetRequest struct {
	ctx        context.Context
	ApiService *StaticVideoFilesAPIService
	filename   string
}

func (r ApiStaticStreamingPlaylistsHlsFilenameGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.StaticStreamingPlaylistsHlsFilenameGetExecute(r)
}

/*
StaticStreamingPlaylistsHlsFilenameGet Get public HLS video file

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param filename Filename
	@return ApiStaticStreamingPlaylistsHlsFilenameGetRequest
*/
func (a *StaticVideoFilesAPIService) StaticStreamingPlaylistsHlsFilenameGet(ctx context.Context, filename string) ApiStaticStreamingPlaylistsHlsFilenameGetRequest {
	return ApiStaticStreamingPlaylistsHlsFilenameGetRequest{
		ApiService: a,
		ctx:        ctx,
		filename:   filename,
	}
}

// Execute executes the request
func (a *StaticVideoFilesAPIService) StaticStreamingPlaylistsHlsFilenameGetExecute(r ApiStaticStreamingPlaylistsHlsFilenameGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StaticVideoFilesAPIService.StaticStreamingPlaylistsHlsFilenameGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/static/streaming-playlists/hls/{filename}"
	localVarPath = strings.Replace(localVarPath, "{"+"filename"+"}", url.PathEscape(parameterValueToString(r.filename, "filename")), -1)

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

type ApiStaticStreamingPlaylistsHlsPrivateFilenameGetRequest struct {
	ctx                    context.Context
	ApiService             *StaticVideoFilesAPIService
	filename               string
	videoFileToken         *string
	reinjectVideoFileToken *bool
}

// Video file token [generated](#operation/requestVideoToken) by PeerTube so you don&#39;t need to provide an OAuth token in the request header.
func (r ApiStaticStreamingPlaylistsHlsPrivateFilenameGetRequest) VideoFileToken(videoFileToken string) ApiStaticStreamingPlaylistsHlsPrivateFilenameGetRequest {
	r.videoFileToken = &videoFileToken
	return r
}

// Ask the server to reinject videoFileToken in URLs in m3u8 playlist
func (r ApiStaticStreamingPlaylistsHlsPrivateFilenameGetRequest) ReinjectVideoFileToken(reinjectVideoFileToken bool) ApiStaticStreamingPlaylistsHlsPrivateFilenameGetRequest {
	r.reinjectVideoFileToken = &reinjectVideoFileToken
	return r
}

func (r ApiStaticStreamingPlaylistsHlsPrivateFilenameGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.StaticStreamingPlaylistsHlsPrivateFilenameGetExecute(r)
}

/*
StaticStreamingPlaylistsHlsPrivateFilenameGet Get private HLS video file

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param filename Filename
	@return ApiStaticStreamingPlaylistsHlsPrivateFilenameGetRequest
*/
func (a *StaticVideoFilesAPIService) StaticStreamingPlaylistsHlsPrivateFilenameGet(ctx context.Context, filename string) ApiStaticStreamingPlaylistsHlsPrivateFilenameGetRequest {
	return ApiStaticStreamingPlaylistsHlsPrivateFilenameGetRequest{
		ApiService: a,
		ctx:        ctx,
		filename:   filename,
	}
}

// Execute executes the request
func (a *StaticVideoFilesAPIService) StaticStreamingPlaylistsHlsPrivateFilenameGetExecute(r ApiStaticStreamingPlaylistsHlsPrivateFilenameGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StaticVideoFilesAPIService.StaticStreamingPlaylistsHlsPrivateFilenameGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/static/streaming-playlists/hls/private/{filename}"
	localVarPath = strings.Replace(localVarPath, "{"+"filename"+"}", url.PathEscape(parameterValueToString(r.filename, "filename")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.videoFileToken != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "videoFileToken", r.videoFileToken, "form", "")
	}
	if r.reinjectVideoFileToken != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "reinjectVideoFileToken", r.reinjectVideoFileToken, "form", "")
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

type ApiStaticWebVideosFilenameGetRequest struct {
	ctx        context.Context
	ApiService *StaticVideoFilesAPIService
	filename   string
}

func (r ApiStaticWebVideosFilenameGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.StaticWebVideosFilenameGetExecute(r)
}

/*
StaticWebVideosFilenameGet Get public Web Video file

**PeerTube >= 6.0**

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param filename Filename
	@return ApiStaticWebVideosFilenameGetRequest
*/
func (a *StaticVideoFilesAPIService) StaticWebVideosFilenameGet(ctx context.Context, filename string) ApiStaticWebVideosFilenameGetRequest {
	return ApiStaticWebVideosFilenameGetRequest{
		ApiService: a,
		ctx:        ctx,
		filename:   filename,
	}
}

// Execute executes the request
func (a *StaticVideoFilesAPIService) StaticWebVideosFilenameGetExecute(r ApiStaticWebVideosFilenameGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StaticVideoFilesAPIService.StaticWebVideosFilenameGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/static/web-videos/{filename}"
	localVarPath = strings.Replace(localVarPath, "{"+"filename"+"}", url.PathEscape(parameterValueToString(r.filename, "filename")), -1)

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

type ApiStaticWebVideosPrivateFilenameGetRequest struct {
	ctx            context.Context
	ApiService     *StaticVideoFilesAPIService
	filename       string
	videoFileToken *string
}

// Video file token [generated](#operation/requestVideoToken) by PeerTube so you don&#39;t need to provide an OAuth token in the request header.
func (r ApiStaticWebVideosPrivateFilenameGetRequest) VideoFileToken(videoFileToken string) ApiStaticWebVideosPrivateFilenameGetRequest {
	r.videoFileToken = &videoFileToken
	return r
}

func (r ApiStaticWebVideosPrivateFilenameGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.StaticWebVideosPrivateFilenameGetExecute(r)
}

/*
StaticWebVideosPrivateFilenameGet Get private Web Video file

**PeerTube >= 6.0**

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param filename Filename
	@return ApiStaticWebVideosPrivateFilenameGetRequest
*/
func (a *StaticVideoFilesAPIService) StaticWebVideosPrivateFilenameGet(ctx context.Context, filename string) ApiStaticWebVideosPrivateFilenameGetRequest {
	return ApiStaticWebVideosPrivateFilenameGetRequest{
		ApiService: a,
		ctx:        ctx,
		filename:   filename,
	}
}

// Execute executes the request
func (a *StaticVideoFilesAPIService) StaticWebVideosPrivateFilenameGetExecute(r ApiStaticWebVideosPrivateFilenameGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StaticVideoFilesAPIService.StaticWebVideosPrivateFilenameGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/static/web-videos/private/{filename}"
	localVarPath = strings.Replace(localVarPath, "{"+"filename"+"}", url.PathEscape(parameterValueToString(r.filename, "filename")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.videoFileToken != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "videoFileToken", r.videoFileToken, "form", "")
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
