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
	"os"
	"strings"
)

// VideoChannelsAPIService VideoChannelsAPI service
type VideoChannelsAPIService service

type ApiAddVideoChannelRequest struct {
	ctx                context.Context
	ApiService         *VideoChannelsAPIService
	videoChannelCreate *models.VideoChannelCreate
}

func (r ApiAddVideoChannelRequest) VideoChannelCreate(videoChannelCreate models.VideoChannelCreate) ApiAddVideoChannelRequest {
	r.videoChannelCreate = &videoChannelCreate
	return r
}

func (r ApiAddVideoChannelRequest) Execute() (*models.AddVideoChannel200Response, *http.Response, error) {
	return r.ApiService.AddVideoChannelExecute(r)
}

/*
AddVideoChannel Create a video channel

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAddVideoChannelRequest
*/
func (a *VideoChannelsAPIService) AddVideoChannel(ctx context.Context) ApiAddVideoChannelRequest {
	return ApiAddVideoChannelRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AddVideoChannel200Response
func (a *VideoChannelsAPIService) AddVideoChannelExecute(r ApiAddVideoChannelRequest) (*models.AddVideoChannel200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.AddVideoChannel200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VideoChannelsAPIService.AddVideoChannel")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/video-channels"

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
	localVarPostBody = r.videoChannelCreate
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

type ApiApiV1VideoChannelsChannelHandleAvatarDeleteRequest struct {
	ctx           context.Context
	ApiService    *VideoChannelsAPIService
	channelHandle string
}

func (r ApiApiV1VideoChannelsChannelHandleAvatarDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiV1VideoChannelsChannelHandleAvatarDeleteExecute(r)
}

/*
ApiV1VideoChannelsChannelHandleAvatarDelete Delete channel avatar

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param channelHandle The video channel handle
	@return ApiApiV1VideoChannelsChannelHandleAvatarDeleteRequest
*/
func (a *VideoChannelsAPIService) ApiV1VideoChannelsChannelHandleAvatarDelete(ctx context.Context, channelHandle string) ApiApiV1VideoChannelsChannelHandleAvatarDeleteRequest {
	return ApiApiV1VideoChannelsChannelHandleAvatarDeleteRequest{
		ApiService:    a,
		ctx:           ctx,
		channelHandle: channelHandle,
	}
}

// Execute executes the request
func (a *VideoChannelsAPIService) ApiV1VideoChannelsChannelHandleAvatarDeleteExecute(r ApiApiV1VideoChannelsChannelHandleAvatarDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VideoChannelsAPIService.ApiV1VideoChannelsChannelHandleAvatarDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/video-channels/{channelHandle}/avatar"
	localVarPath = strings.Replace(localVarPath, "{"+"channelHandle"+"}", url.PathEscape(parameterValueToString(r.channelHandle, "channelHandle")), -1)

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

type ApiApiV1VideoChannelsChannelHandleAvatarPickPostRequest struct {
	ctx           context.Context
	ApiService    *VideoChannelsAPIService
	channelHandle string
	avatarfile    *os.File
}

// The file to upload.
func (r ApiApiV1VideoChannelsChannelHandleAvatarPickPostRequest) Avatarfile(avatarfile *os.File) ApiApiV1VideoChannelsChannelHandleAvatarPickPostRequest {
	r.avatarfile = avatarfile
	return r
}

func (r ApiApiV1VideoChannelsChannelHandleAvatarPickPostRequest) Execute() (*models.ApiV1UsersMeAvatarPickPost200Response, *http.Response, error) {
	return r.ApiService.ApiV1VideoChannelsChannelHandleAvatarPickPostExecute(r)
}

/*
ApiV1VideoChannelsChannelHandleAvatarPickPost Update channel avatar

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param channelHandle The video channel handle
	@return ApiApiV1VideoChannelsChannelHandleAvatarPickPostRequest
*/
func (a *VideoChannelsAPIService) ApiV1VideoChannelsChannelHandleAvatarPickPost(ctx context.Context, channelHandle string) ApiApiV1VideoChannelsChannelHandleAvatarPickPostRequest {
	return ApiApiV1VideoChannelsChannelHandleAvatarPickPostRequest{
		ApiService:    a,
		ctx:           ctx,
		channelHandle: channelHandle,
	}
}

// Execute executes the request
//
//	@return ApiV1UsersMeAvatarPickPost200Response
func (a *VideoChannelsAPIService) ApiV1VideoChannelsChannelHandleAvatarPickPostExecute(r ApiApiV1VideoChannelsChannelHandleAvatarPickPostRequest) (*models.ApiV1UsersMeAvatarPickPost200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.ApiV1UsersMeAvatarPickPost200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VideoChannelsAPIService.ApiV1VideoChannelsChannelHandleAvatarPickPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/video-channels/{channelHandle}/avatar/pick"
	localVarPath = strings.Replace(localVarPath, "{"+"channelHandle"+"}", url.PathEscape(parameterValueToString(r.channelHandle, "channelHandle")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	var avatarfileLocalVarFormFileName string
	var avatarfileLocalVarFileName string
	var avatarfileLocalVarFileBytes []byte

	avatarfileLocalVarFormFileName = "avatarfile"
	avatarfileLocalVarFile := r.avatarfile

	if avatarfileLocalVarFile != nil {
		fbs, _ := io.ReadAll(avatarfileLocalVarFile)

		avatarfileLocalVarFileBytes = fbs
		avatarfileLocalVarFileName = avatarfileLocalVarFile.Name()
		avatarfileLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: avatarfileLocalVarFileBytes, fileName: avatarfileLocalVarFileName, formFileName: avatarfileLocalVarFormFileName})
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

type ApiApiV1VideoChannelsChannelHandleBannerDeleteRequest struct {
	ctx           context.Context
	ApiService    *VideoChannelsAPIService
	channelHandle string
}

func (r ApiApiV1VideoChannelsChannelHandleBannerDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiV1VideoChannelsChannelHandleBannerDeleteExecute(r)
}

/*
ApiV1VideoChannelsChannelHandleBannerDelete Delete channel banner

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param channelHandle The video channel handle
	@return ApiApiV1VideoChannelsChannelHandleBannerDeleteRequest
*/
func (a *VideoChannelsAPIService) ApiV1VideoChannelsChannelHandleBannerDelete(ctx context.Context, channelHandle string) ApiApiV1VideoChannelsChannelHandleBannerDeleteRequest {
	return ApiApiV1VideoChannelsChannelHandleBannerDeleteRequest{
		ApiService:    a,
		ctx:           ctx,
		channelHandle: channelHandle,
	}
}

// Execute executes the request
func (a *VideoChannelsAPIService) ApiV1VideoChannelsChannelHandleBannerDeleteExecute(r ApiApiV1VideoChannelsChannelHandleBannerDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VideoChannelsAPIService.ApiV1VideoChannelsChannelHandleBannerDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/video-channels/{channelHandle}/banner"
	localVarPath = strings.Replace(localVarPath, "{"+"channelHandle"+"}", url.PathEscape(parameterValueToString(r.channelHandle, "channelHandle")), -1)

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

type ApiApiV1VideoChannelsChannelHandleBannerPickPostRequest struct {
	ctx           context.Context
	ApiService    *VideoChannelsAPIService
	channelHandle string
	bannerfile    *os.File
}

// The file to upload.
func (r ApiApiV1VideoChannelsChannelHandleBannerPickPostRequest) Bannerfile(bannerfile *os.File) ApiApiV1VideoChannelsChannelHandleBannerPickPostRequest {
	r.bannerfile = bannerfile
	return r
}

func (r ApiApiV1VideoChannelsChannelHandleBannerPickPostRequest) Execute() (*models.ApiV1VideoChannelsChannelHandleBannerPickPost200Response, *http.Response, error) {
	return r.ApiService.ApiV1VideoChannelsChannelHandleBannerPickPostExecute(r)
}

/*
ApiV1VideoChannelsChannelHandleBannerPickPost Update channel banner

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param channelHandle The video channel handle
	@return ApiApiV1VideoChannelsChannelHandleBannerPickPostRequest
*/
func (a *VideoChannelsAPIService) ApiV1VideoChannelsChannelHandleBannerPickPost(ctx context.Context, channelHandle string) ApiApiV1VideoChannelsChannelHandleBannerPickPostRequest {
	return ApiApiV1VideoChannelsChannelHandleBannerPickPostRequest{
		ApiService:    a,
		ctx:           ctx,
		channelHandle: channelHandle,
	}
}

// Execute executes the request
//
//	@return ApiV1VideoChannelsChannelHandleBannerPickPost200Response
func (a *VideoChannelsAPIService) ApiV1VideoChannelsChannelHandleBannerPickPostExecute(r ApiApiV1VideoChannelsChannelHandleBannerPickPostRequest) (*models.ApiV1VideoChannelsChannelHandleBannerPickPost200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.ApiV1VideoChannelsChannelHandleBannerPickPost200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VideoChannelsAPIService.ApiV1VideoChannelsChannelHandleBannerPickPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/video-channels/{channelHandle}/banner/pick"
	localVarPath = strings.Replace(localVarPath, "{"+"channelHandle"+"}", url.PathEscape(parameterValueToString(r.channelHandle, "channelHandle")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	var bannerfileLocalVarFormFileName string
	var bannerfileLocalVarFileName string
	var bannerfileLocalVarFileBytes []byte

	bannerfileLocalVarFormFileName = "bannerfile"
	bannerfileLocalVarFile := r.bannerfile

	if bannerfileLocalVarFile != nil {
		fbs, _ := io.ReadAll(bannerfileLocalVarFile)

		bannerfileLocalVarFileBytes = fbs
		bannerfileLocalVarFileName = bannerfileLocalVarFile.Name()
		bannerfileLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: bannerfileLocalVarFileBytes, fileName: bannerfileLocalVarFileName, formFileName: bannerfileLocalVarFormFileName})
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

type ApiApiV1VideoChannelsChannelHandleImportVideosPostRequest struct {
	ctx                         context.Context
	ApiService                  *VideoChannelsAPIService
	channelHandle               string
	importVideosInChannelCreate *models.ImportVideosInChannelCreate
}

func (r ApiApiV1VideoChannelsChannelHandleImportVideosPostRequest) ImportVideosInChannelCreate(importVideosInChannelCreate models.ImportVideosInChannelCreate) ApiApiV1VideoChannelsChannelHandleImportVideosPostRequest {
	r.importVideosInChannelCreate = &importVideosInChannelCreate
	return r
}

func (r ApiApiV1VideoChannelsChannelHandleImportVideosPostRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiV1VideoChannelsChannelHandleImportVideosPostExecute(r)
}

/*
ApiV1VideoChannelsChannelHandleImportVideosPost Import videos in channel

Import a remote channel/playlist videos into a channel

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param channelHandle The video channel handle
	@return ApiApiV1VideoChannelsChannelHandleImportVideosPostRequest
*/
func (a *VideoChannelsAPIService) ApiV1VideoChannelsChannelHandleImportVideosPost(ctx context.Context, channelHandle string) ApiApiV1VideoChannelsChannelHandleImportVideosPostRequest {
	return ApiApiV1VideoChannelsChannelHandleImportVideosPostRequest{
		ApiService:    a,
		ctx:           ctx,
		channelHandle: channelHandle,
	}
}

// Execute executes the request
func (a *VideoChannelsAPIService) ApiV1VideoChannelsChannelHandleImportVideosPostExecute(r ApiApiV1VideoChannelsChannelHandleImportVideosPostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VideoChannelsAPIService.ApiV1VideoChannelsChannelHandleImportVideosPost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/video-channels/{channelHandle}/import-videos"
	localVarPath = strings.Replace(localVarPath, "{"+"channelHandle"+"}", url.PathEscape(parameterValueToString(r.channelHandle, "channelHandle")), -1)

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
	localVarPostBody = r.importVideosInChannelCreate
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

type ApiApiV1VideoChannelsChannelHandleVideoPlaylistsGetRequest struct {
	ctx           context.Context
	ApiService    *VideoChannelsAPIService
	channelHandle string
	start         *int32
	count         *int32
	sort          *string
	playlistType  *models.VideoPlaylistTypeSet
}

// Offset used to paginate results
func (r ApiApiV1VideoChannelsChannelHandleVideoPlaylistsGetRequest) Start(start int32) ApiApiV1VideoChannelsChannelHandleVideoPlaylistsGetRequest {
	r.start = &start
	return r
}

// Number of items to return
func (r ApiApiV1VideoChannelsChannelHandleVideoPlaylistsGetRequest) Count(count int32) ApiApiV1VideoChannelsChannelHandleVideoPlaylistsGetRequest {
	r.count = &count
	return r
}

// Sort column
func (r ApiApiV1VideoChannelsChannelHandleVideoPlaylistsGetRequest) Sort(sort string) ApiApiV1VideoChannelsChannelHandleVideoPlaylistsGetRequest {
	r.sort = &sort
	return r
}

func (r ApiApiV1VideoChannelsChannelHandleVideoPlaylistsGetRequest) PlaylistType(playlistType models.VideoPlaylistTypeSet) ApiApiV1VideoChannelsChannelHandleVideoPlaylistsGetRequest {
	r.playlistType = &playlistType
	return r
}

func (r ApiApiV1VideoChannelsChannelHandleVideoPlaylistsGetRequest) Execute() (*models.ApiV1VideoChannelsChannelHandleVideoPlaylistsGet200Response, *http.Response, error) {
	return r.ApiService.ApiV1VideoChannelsChannelHandleVideoPlaylistsGetExecute(r)
}

/*
ApiV1VideoChannelsChannelHandleVideoPlaylistsGet List playlists of a channel

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param channelHandle The video channel handle
	@return ApiApiV1VideoChannelsChannelHandleVideoPlaylistsGetRequest
*/
func (a *VideoChannelsAPIService) ApiV1VideoChannelsChannelHandleVideoPlaylistsGet(ctx context.Context, channelHandle string) ApiApiV1VideoChannelsChannelHandleVideoPlaylistsGetRequest {
	return ApiApiV1VideoChannelsChannelHandleVideoPlaylistsGetRequest{
		ApiService:    a,
		ctx:           ctx,
		channelHandle: channelHandle,
	}
}

// Execute executes the request
//
//	@return ApiV1VideoChannelsChannelHandleVideoPlaylistsGet200Response
func (a *VideoChannelsAPIService) ApiV1VideoChannelsChannelHandleVideoPlaylistsGetExecute(r ApiApiV1VideoChannelsChannelHandleVideoPlaylistsGetRequest) (*models.ApiV1VideoChannelsChannelHandleVideoPlaylistsGet200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.ApiV1VideoChannelsChannelHandleVideoPlaylistsGet200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VideoChannelsAPIService.ApiV1VideoChannelsChannelHandleVideoPlaylistsGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/video-channels/{channelHandle}/video-playlists"
	localVarPath = strings.Replace(localVarPath, "{"+"channelHandle"+"}", url.PathEscape(parameterValueToString(r.channelHandle, "channelHandle")), -1)

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

type ApiDelVideoChannelRequest struct {
	ctx           context.Context
	ApiService    *VideoChannelsAPIService
	channelHandle string
}

func (r ApiDelVideoChannelRequest) Execute() (*http.Response, error) {
	return r.ApiService.DelVideoChannelExecute(r)
}

/*
DelVideoChannel Delete a video channel

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param channelHandle The video channel handle
	@return ApiDelVideoChannelRequest
*/
func (a *VideoChannelsAPIService) DelVideoChannel(ctx context.Context, channelHandle string) ApiDelVideoChannelRequest {
	return ApiDelVideoChannelRequest{
		ApiService:    a,
		ctx:           ctx,
		channelHandle: channelHandle,
	}
}

// Execute executes the request
func (a *VideoChannelsAPIService) DelVideoChannelExecute(r ApiDelVideoChannelRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VideoChannelsAPIService.DelVideoChannel")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/video-channels/{channelHandle}"
	localVarPath = strings.Replace(localVarPath, "{"+"channelHandle"+"}", url.PathEscape(parameterValueToString(r.channelHandle, "channelHandle")), -1)

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

type ApiGetVideoChannelRequest struct {
	ctx           context.Context
	ApiService    *VideoChannelsAPIService
	channelHandle string
}

func (r ApiGetVideoChannelRequest) Execute() (*models.VideoChannel, *http.Response, error) {
	return r.ApiService.GetVideoChannelExecute(r)
}

/*
GetVideoChannel Get a video channel

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param channelHandle The video channel handle
	@return ApiGetVideoChannelRequest
*/
func (a *VideoChannelsAPIService) GetVideoChannel(ctx context.Context, channelHandle string) ApiGetVideoChannelRequest {
	return ApiGetVideoChannelRequest{
		ApiService:    a,
		ctx:           ctx,
		channelHandle: channelHandle,
	}
}

// Execute executes the request
//
//	@return VideoChannel
func (a *VideoChannelsAPIService) GetVideoChannelExecute(r ApiGetVideoChannelRequest) (*models.VideoChannel, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.VideoChannel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VideoChannelsAPIService.GetVideoChannel")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/video-channels/{channelHandle}"
	localVarPath = strings.Replace(localVarPath, "{"+"channelHandle"+"}", url.PathEscape(parameterValueToString(r.channelHandle, "channelHandle")), -1)

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

type ApiGetVideoChannelFollowersRequest struct {
	ctx           context.Context
	ApiService    *VideoChannelsAPIService
	channelHandle string
	start         *int32
	count         *int32
	sort          *string
	search        *string
}

// Offset used to paginate results
func (r ApiGetVideoChannelFollowersRequest) Start(start int32) ApiGetVideoChannelFollowersRequest {
	r.start = &start
	return r
}

// Number of items to return
func (r ApiGetVideoChannelFollowersRequest) Count(count int32) ApiGetVideoChannelFollowersRequest {
	r.count = &count
	return r
}

// Sort followers by criteria
func (r ApiGetVideoChannelFollowersRequest) Sort(sort string) ApiGetVideoChannelFollowersRequest {
	r.sort = &sort
	return r
}

// Plain text search, applied to various parts of the model depending on endpoint
func (r ApiGetVideoChannelFollowersRequest) Search(search string) ApiGetVideoChannelFollowersRequest {
	r.search = &search
	return r
}

func (r ApiGetVideoChannelFollowersRequest) Execute() (*models.GetAccountFollowers200Response, *http.Response, error) {
	return r.ApiService.GetVideoChannelFollowersExecute(r)
}

/*
GetVideoChannelFollowers List followers of a video channel

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param channelHandle The video channel handle
	@return ApiGetVideoChannelFollowersRequest
*/
func (a *VideoChannelsAPIService) GetVideoChannelFollowers(ctx context.Context, channelHandle string) ApiGetVideoChannelFollowersRequest {
	return ApiGetVideoChannelFollowersRequest{
		ApiService:    a,
		ctx:           ctx,
		channelHandle: channelHandle,
	}
}

// Execute executes the request
//
//	@return GetAccountFollowers200Response
func (a *VideoChannelsAPIService) GetVideoChannelFollowersExecute(r ApiGetVideoChannelFollowersRequest) (*models.GetAccountFollowers200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.GetAccountFollowers200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VideoChannelsAPIService.GetVideoChannelFollowers")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/video-channels/{channelHandle}/followers"
	localVarPath = strings.Replace(localVarPath, "{"+"channelHandle"+"}", url.PathEscape(parameterValueToString(r.channelHandle, "channelHandle")), -1)

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

type ApiGetVideoChannelVideosRequest struct {
	ctx                   context.Context
	ApiService            *VideoChannelsAPIService
	channelHandle         string
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
func (r ApiGetVideoChannelVideosRequest) CategoryOneOf(categoryOneOf models.GetAccountVideosCategoryOneOfParameter) ApiGetVideoChannelVideosRequest {
	r.categoryOneOf = &categoryOneOf
	return r
}

// whether or not the video is a live
func (r ApiGetVideoChannelVideosRequest) IsLive(isLive bool) ApiGetVideoChannelVideosRequest {
	r.isLive = &isLive
	return r
}

// tag(s) of the video
func (r ApiGetVideoChannelVideosRequest) TagsOneOf(tagsOneOf models.GetAccountVideosTagsOneOfParameter) ApiGetVideoChannelVideosRequest {
	r.tagsOneOf = &tagsOneOf
	return r
}

// tag(s) of the video, where all should be present in the video
func (r ApiGetVideoChannelVideosRequest) TagsAllOf(tagsAllOf models.GetAccountVideosTagsAllOfParameter) ApiGetVideoChannelVideosRequest {
	r.tagsAllOf = &tagsAllOf
	return r
}

// licence id of the video (see [/videos/licences](#operation/getLicences))
func (r ApiGetVideoChannelVideosRequest) LicenceOneOf(licenceOneOf models.GetAccountVideosLicenceOneOfParameter) ApiGetVideoChannelVideosRequest {
	r.licenceOneOf = &licenceOneOf
	return r
}

// language id of the video (see [/videos/languages](#operation/getLanguages)). Use &#x60;_unknown&#x60; to filter on videos that don&#39;t have a video language
func (r ApiGetVideoChannelVideosRequest) LanguageOneOf(languageOneOf models.GetAccountVideosLanguageOneOfParameter) ApiGetVideoChannelVideosRequest {
	r.languageOneOf = &languageOneOf
	return r
}

// **PeerTube &gt;&#x3D; 6.2** **Admins and moderators only** filter on videos that contain one of these automatic tags
func (r ApiGetVideoChannelVideosRequest) AutoTagOneOf(autoTagOneOf models.GetAccountVideosTagsAllOfParameter) ApiGetVideoChannelVideosRequest {
	r.autoTagOneOf = &autoTagOneOf
	return r
}

// whether to include nsfw videos, if any
func (r ApiGetVideoChannelVideosRequest) Nsfw(nsfw string) ApiGetVideoChannelVideosRequest {
	r.nsfw = &nsfw
	return r
}

// **PeerTube &gt;&#x3D; 4.0** Display only local or remote objects
func (r ApiGetVideoChannelVideosRequest) IsLocal(isLocal bool) ApiGetVideoChannelVideosRequest {
	r.isLocal = &isLocal
	return r
}

// **Only administrators and moderators can use this parameter**  Include additional videos in results (can be combined using bitwise or operator) - &#x60;0&#x60; NONE - &#x60;1&#x60; NOT_PUBLISHED_STATE - &#x60;2&#x60; BLACKLISTED - &#x60;4&#x60; BLOCKED_OWNER - &#x60;8&#x60; FILES - &#x60;16&#x60; CAPTIONS - &#x60;32&#x60; VIDEO SOURCE
func (r ApiGetVideoChannelVideosRequest) Include(include int32) ApiGetVideoChannelVideosRequest {
	r.include = &include
	return r
}

// **PeerTube &gt;&#x3D; 4.0** Display only videos in this specific privacy/privacies
func (r ApiGetVideoChannelVideosRequest) PrivacyOneOf(privacyOneOf models.VideoPrivacySet) ApiGetVideoChannelVideosRequest {
	r.privacyOneOf = &privacyOneOf
	return r
}

// **PeerTube &gt;&#x3D; 4.0** Display only videos that have HLS files
func (r ApiGetVideoChannelVideosRequest) HasHLSFiles(hasHLSFiles bool) ApiGetVideoChannelVideosRequest {
	r.hasHLSFiles = &hasHLSFiles
	return r
}

// **PeerTube &gt;&#x3D; 6.0** Display only videos that have Web Video files
func (r ApiGetVideoChannelVideosRequest) HasWebVideoFiles(hasWebVideoFiles bool) ApiGetVideoChannelVideosRequest {
	r.hasWebVideoFiles = &hasWebVideoFiles
	return r
}

// if you don&#39;t need the &#x60;total&#x60; in the response
func (r ApiGetVideoChannelVideosRequest) SkipCount(skipCount string) ApiGetVideoChannelVideosRequest {
	r.skipCount = &skipCount
	return r
}

// Offset used to paginate results
func (r ApiGetVideoChannelVideosRequest) Start(start int32) ApiGetVideoChannelVideosRequest {
	r.start = &start
	return r
}

// Number of items to return
func (r ApiGetVideoChannelVideosRequest) Count(count int32) ApiGetVideoChannelVideosRequest {
	r.count = &count
	return r
}

func (r ApiGetVideoChannelVideosRequest) Sort(sort string) ApiGetVideoChannelVideosRequest {
	r.sort = &sort
	return r
}

// Whether or not to exclude videos that are in the user&#39;s video history
func (r ApiGetVideoChannelVideosRequest) ExcludeAlreadyWatched(excludeAlreadyWatched bool) ApiGetVideoChannelVideosRequest {
	r.excludeAlreadyWatched = &excludeAlreadyWatched
	return r
}

// Plain text search, applied to various parts of the model depending on endpoint
func (r ApiGetVideoChannelVideosRequest) Search(search string) ApiGetVideoChannelVideosRequest {
	r.search = &search
	return r
}

func (r ApiGetVideoChannelVideosRequest) Execute() (*models.VideoListResponse, *http.Response, error) {
	return r.ApiService.GetVideoChannelVideosExecute(r)
}

/*
GetVideoChannelVideos List videos of a video channel

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param channelHandle The video channel handle
	@return ApiGetVideoChannelVideosRequest
*/
func (a *VideoChannelsAPIService) GetVideoChannelVideos(ctx context.Context, channelHandle string) ApiGetVideoChannelVideosRequest {
	return ApiGetVideoChannelVideosRequest{
		ApiService:    a,
		ctx:           ctx,
		channelHandle: channelHandle,
	}
}

// Execute executes the request
//
//	@return VideoListResponse
func (a *VideoChannelsAPIService) GetVideoChannelVideosExecute(r ApiGetVideoChannelVideosRequest) (*models.VideoListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.VideoListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VideoChannelsAPIService.GetVideoChannelVideos")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/video-channels/{channelHandle}/videos"
	localVarPath = strings.Replace(localVarPath, "{"+"channelHandle"+"}", url.PathEscape(parameterValueToString(r.channelHandle, "channelHandle")), -1)

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

type ApiGetVideoChannelsRequest struct {
	ctx        context.Context
	ApiService *VideoChannelsAPIService
	start      *int32
	count      *int32
	sort       *string
}

// Offset used to paginate results
func (r ApiGetVideoChannelsRequest) Start(start int32) ApiGetVideoChannelsRequest {
	r.start = &start
	return r
}

// Number of items to return
func (r ApiGetVideoChannelsRequest) Count(count int32) ApiGetVideoChannelsRequest {
	r.count = &count
	return r
}

// Sort column
func (r ApiGetVideoChannelsRequest) Sort(sort string) ApiGetVideoChannelsRequest {
	r.sort = &sort
	return r
}

func (r ApiGetVideoChannelsRequest) Execute() (*models.VideoChannelList, *http.Response, error) {
	return r.ApiService.GetVideoChannelsExecute(r)
}

/*
GetVideoChannels List video channels

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetVideoChannelsRequest
*/
func (a *VideoChannelsAPIService) GetVideoChannels(ctx context.Context) ApiGetVideoChannelsRequest {
	return ApiGetVideoChannelsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return VideoChannelList
func (a *VideoChannelsAPIService) GetVideoChannelsExecute(r ApiGetVideoChannelsRequest) (*models.VideoChannelList, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.VideoChannelList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VideoChannelsAPIService.GetVideoChannels")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/video-channels"

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

type ApiPutVideoChannelRequest struct {
	ctx                context.Context
	ApiService         *VideoChannelsAPIService
	channelHandle      string
	videoChannelUpdate *models.VideoChannelUpdate
}

func (r ApiPutVideoChannelRequest) VideoChannelUpdate(videoChannelUpdate models.VideoChannelUpdate) ApiPutVideoChannelRequest {
	r.videoChannelUpdate = &videoChannelUpdate
	return r
}

func (r ApiPutVideoChannelRequest) Execute() (*http.Response, error) {
	return r.ApiService.PutVideoChannelExecute(r)
}

/*
PutVideoChannel Update a video channel

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param channelHandle The video channel handle
	@return ApiPutVideoChannelRequest
*/
func (a *VideoChannelsAPIService) PutVideoChannel(ctx context.Context, channelHandle string) ApiPutVideoChannelRequest {
	return ApiPutVideoChannelRequest{
		ApiService:    a,
		ctx:           ctx,
		channelHandle: channelHandle,
	}
}

// Execute executes the request
func (a *VideoChannelsAPIService) PutVideoChannelExecute(r ApiPutVideoChannelRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPut
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VideoChannelsAPIService.PutVideoChannel")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/video-channels/{channelHandle}"
	localVarPath = strings.Replace(localVarPath, "{"+"channelHandle"+"}", url.PathEscape(parameterValueToString(r.channelHandle, "channelHandle")), -1)

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
	localVarPostBody = r.videoChannelUpdate
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
