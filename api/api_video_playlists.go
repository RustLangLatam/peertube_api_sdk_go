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
	"reflect"
	"strings"
)

// VideoPlaylistsAPIService VideoPlaylistsAPI service
type VideoPlaylistsAPIService service

type ApiAddPlaylistRequest struct {
	ctx            context.Context
	ApiService     *VideoPlaylistsAPIService
	displayName    *string
	thumbnailfile  *os.File
	privacy        *models.VideoPlaylistPrivacySet
	description    *string
	videoChannelId *int32
}

// Video playlist display name
func (r ApiAddPlaylistRequest) DisplayName(displayName string) ApiAddPlaylistRequest {
	r.displayName = &displayName
	return r
}

// Video playlist thumbnail file
func (r ApiAddPlaylistRequest) Thumbnailfile(thumbnailfile *os.File) ApiAddPlaylistRequest {
	r.thumbnailfile = thumbnailfile
	return r
}

func (r ApiAddPlaylistRequest) Privacy(privacy models.VideoPlaylistPrivacySet) ApiAddPlaylistRequest {
	r.privacy = &privacy
	return r
}

// Video playlist description
func (r ApiAddPlaylistRequest) Description(description string) ApiAddPlaylistRequest {
	r.description = &description
	return r
}

// Video channel in which the playlist will be published
func (r ApiAddPlaylistRequest) VideoChannelId(videoChannelId int32) ApiAddPlaylistRequest {
	r.videoChannelId = &videoChannelId
	return r
}

func (r ApiAddPlaylistRequest) Execute() (*models.AddPlaylist200Response, *http.Response, error) {
	return r.ApiService.AddPlaylistExecute(r)
}

/*
AddPlaylist Create a video playlist

If the video playlist is set as public, `videoChannelId` is mandatory.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAddPlaylistRequest
*/
func (a *VideoPlaylistsAPIService) AddPlaylist(ctx context.Context) ApiAddPlaylistRequest {
	return ApiAddPlaylistRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AddPlaylist200Response
func (a *VideoPlaylistsAPIService) AddPlaylistExecute(r ApiAddPlaylistRequest) (*models.AddPlaylist200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.AddPlaylist200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VideoPlaylistsAPIService.AddPlaylist")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/video-playlists"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.displayName == nil {
		return localVarReturnValue, nil, utils.ReportError("displayName is required and must be specified")
	}
	if strlen(*r.displayName) < 1 {
		return localVarReturnValue, nil, utils.ReportError("displayName must have at least 1 elements")
	}
	if strlen(*r.displayName) > 120 {
		return localVarReturnValue, nil, utils.ReportError("displayName must have less than 120 elements")
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
	parameterAddToHeaderOrQuery(localVarFormParams, "displayName", r.displayName, "", "")
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
	if r.privacy != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "privacy", r.privacy, "", "")
	}
	if r.description != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "description", r.description, "", "")
	}
	if r.videoChannelId != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "videoChannelId", r.videoChannelId, "", "")
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

type ApiAddVideoPlaylistVideoRequest struct {
	ctx                          context.Context
	ApiService                   *VideoPlaylistsAPIService
	playlistId                   int32
	addVideoPlaylistVideoRequest *models.AddVideoPlaylistVideoRequest
}

func (r ApiAddVideoPlaylistVideoRequest) AddVideoPlaylistVideoRequest(addVideoPlaylistVideoRequest models.AddVideoPlaylistVideoRequest) ApiAddVideoPlaylistVideoRequest {
	r.addVideoPlaylistVideoRequest = &addVideoPlaylistVideoRequest
	return r
}

func (r ApiAddVideoPlaylistVideoRequest) Execute() (*models.AddVideoPlaylistVideo200Response, *http.Response, error) {
	return r.ApiService.AddVideoPlaylistVideoExecute(r)
}

/*
AddVideoPlaylistVideo Add a video in a playlist

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param playlistId Playlist id
	@return ApiAddVideoPlaylistVideoRequest
*/
func (a *VideoPlaylistsAPIService) AddVideoPlaylistVideo(ctx context.Context, playlistId int32) ApiAddVideoPlaylistVideoRequest {
	return ApiAddVideoPlaylistVideoRequest{
		ApiService: a,
		ctx:        ctx,
		playlistId: playlistId,
	}
}

// Execute executes the request
//
//	@return AddVideoPlaylistVideo200Response
func (a *VideoPlaylistsAPIService) AddVideoPlaylistVideoExecute(r ApiAddVideoPlaylistVideoRequest) (*models.AddVideoPlaylistVideo200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.AddVideoPlaylistVideo200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VideoPlaylistsAPIService.AddVideoPlaylistVideo")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/video-playlists/{playlistId}/videos"
	localVarPath = strings.Replace(localVarPath, "{"+"playlistId"+"}", url.PathEscape(parameterValueToString(r.playlistId, "playlistId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.playlistId < 1 {
		return localVarReturnValue, nil, utils.ReportError("playlistId must be greater than 1")
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
	localVarPostBody = r.addVideoPlaylistVideoRequest
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

type ApiApiV1UsersMeVideoPlaylistsVideosExistGetRequest struct {
	ctx        context.Context
	ApiService *VideoPlaylistsAPIService
	videoIds   *[]int32
}

// The video ids to check
func (r ApiApiV1UsersMeVideoPlaylistsVideosExistGetRequest) VideoIds(videoIds []int32) ApiApiV1UsersMeVideoPlaylistsVideosExistGetRequest {
	r.videoIds = &videoIds
	return r
}

func (r ApiApiV1UsersMeVideoPlaylistsVideosExistGetRequest) Execute() (*models.ApiV1UsersMeVideoPlaylistsVideosExistGet200Response, *http.Response, error) {
	return r.ApiService.ApiV1UsersMeVideoPlaylistsVideosExistGetExecute(r)
}

/*
ApiV1UsersMeVideoPlaylistsVideosExistGet Check video exists in my playlists

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiApiV1UsersMeVideoPlaylistsVideosExistGetRequest
*/
func (a *VideoPlaylistsAPIService) ApiV1UsersMeVideoPlaylistsVideosExistGet(ctx context.Context) ApiApiV1UsersMeVideoPlaylistsVideosExistGetRequest {
	return ApiApiV1UsersMeVideoPlaylistsVideosExistGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ApiV1UsersMeVideoPlaylistsVideosExistGet200Response
func (a *VideoPlaylistsAPIService) ApiV1UsersMeVideoPlaylistsVideosExistGetExecute(r ApiApiV1UsersMeVideoPlaylistsVideosExistGetRequest) (*models.ApiV1UsersMeVideoPlaylistsVideosExistGet200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.ApiV1UsersMeVideoPlaylistsVideosExistGet200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VideoPlaylistsAPIService.ApiV1UsersMeVideoPlaylistsVideosExistGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/users/me/video-playlists/videos-exist"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.videoIds == nil {
		return localVarReturnValue, nil, utils.ReportError("videoIds is required and must be specified")
	}

	{
		t := *r.videoIds
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "videoIds", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "videoIds", t, "form", "multi")
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

type ApiApiV1VideoPlaylistsPlaylistIdDeleteRequest struct {
	ctx        context.Context
	ApiService *VideoPlaylistsAPIService
	playlistId int32
}

func (r ApiApiV1VideoPlaylistsPlaylistIdDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiV1VideoPlaylistsPlaylistIdDeleteExecute(r)
}

/*
ApiV1VideoPlaylistsPlaylistIdDelete Delete a video playlist

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param playlistId Playlist id
	@return ApiApiV1VideoPlaylistsPlaylistIdDeleteRequest
*/
func (a *VideoPlaylistsAPIService) ApiV1VideoPlaylistsPlaylistIdDelete(ctx context.Context, playlistId int32) ApiApiV1VideoPlaylistsPlaylistIdDeleteRequest {
	return ApiApiV1VideoPlaylistsPlaylistIdDeleteRequest{
		ApiService: a,
		ctx:        ctx,
		playlistId: playlistId,
	}
}

// Execute executes the request
func (a *VideoPlaylistsAPIService) ApiV1VideoPlaylistsPlaylistIdDeleteExecute(r ApiApiV1VideoPlaylistsPlaylistIdDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VideoPlaylistsAPIService.ApiV1VideoPlaylistsPlaylistIdDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/video-playlists/{playlistId}"
	localVarPath = strings.Replace(localVarPath, "{"+"playlistId"+"}", url.PathEscape(parameterValueToString(r.playlistId, "playlistId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.playlistId < 1 {
		return nil, utils.ReportError("playlistId must be greater than 1")
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

type ApiApiV1VideoPlaylistsPlaylistIdGetRequest struct {
	ctx        context.Context
	ApiService *VideoPlaylistsAPIService
	playlistId int32
}

func (r ApiApiV1VideoPlaylistsPlaylistIdGetRequest) Execute() (*models.VideoPlaylist, *http.Response, error) {
	return r.ApiService.ApiV1VideoPlaylistsPlaylistIdGetExecute(r)
}

/*
ApiV1VideoPlaylistsPlaylistIdGet Get a video playlist

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param playlistId Playlist id
	@return ApiApiV1VideoPlaylistsPlaylistIdGetRequest
*/
func (a *VideoPlaylistsAPIService) ApiV1VideoPlaylistsPlaylistIdGet(ctx context.Context, playlistId int32) ApiApiV1VideoPlaylistsPlaylistIdGetRequest {
	return ApiApiV1VideoPlaylistsPlaylistIdGetRequest{
		ApiService: a,
		ctx:        ctx,
		playlistId: playlistId,
	}
}

// Execute executes the request
//
//	@return VideoPlaylist
func (a *VideoPlaylistsAPIService) ApiV1VideoPlaylistsPlaylistIdGetExecute(r ApiApiV1VideoPlaylistsPlaylistIdGetRequest) (*models.VideoPlaylist, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.VideoPlaylist
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VideoPlaylistsAPIService.ApiV1VideoPlaylistsPlaylistIdGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/video-playlists/{playlistId}"
	localVarPath = strings.Replace(localVarPath, "{"+"playlistId"+"}", url.PathEscape(parameterValueToString(r.playlistId, "playlistId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.playlistId < 1 {
		return localVarReturnValue, nil, utils.ReportError("playlistId must be greater than 1")
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

type ApiApiV1VideoPlaylistsPlaylistIdPutRequest struct {
	ctx            context.Context
	ApiService     *VideoPlaylistsAPIService
	playlistId     int32
	displayName    *string
	thumbnailfile  *os.File
	privacy        *models.VideoPlaylistPrivacySet
	description    *string
	videoChannelId *int32
}

// Video playlist display name
func (r ApiApiV1VideoPlaylistsPlaylistIdPutRequest) DisplayName(displayName string) ApiApiV1VideoPlaylistsPlaylistIdPutRequest {
	r.displayName = &displayName
	return r
}

// Video playlist thumbnail file
func (r ApiApiV1VideoPlaylistsPlaylistIdPutRequest) Thumbnailfile(thumbnailfile *os.File) ApiApiV1VideoPlaylistsPlaylistIdPutRequest {
	r.thumbnailfile = thumbnailfile
	return r
}

func (r ApiApiV1VideoPlaylistsPlaylistIdPutRequest) Privacy(privacy models.VideoPlaylistPrivacySet) ApiApiV1VideoPlaylistsPlaylistIdPutRequest {
	r.privacy = &privacy
	return r
}

// Video playlist description
func (r ApiApiV1VideoPlaylistsPlaylistIdPutRequest) Description(description string) ApiApiV1VideoPlaylistsPlaylistIdPutRequest {
	r.description = &description
	return r
}

// Video channel in which the playlist will be published
func (r ApiApiV1VideoPlaylistsPlaylistIdPutRequest) VideoChannelId(videoChannelId int32) ApiApiV1VideoPlaylistsPlaylistIdPutRequest {
	r.videoChannelId = &videoChannelId
	return r
}

func (r ApiApiV1VideoPlaylistsPlaylistIdPutRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiV1VideoPlaylistsPlaylistIdPutExecute(r)
}

/*
ApiV1VideoPlaylistsPlaylistIdPut Update a video playlist

If the video playlist is set as public, the playlist must have a assigned channel.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param playlistId Playlist id
	@return ApiApiV1VideoPlaylistsPlaylistIdPutRequest
*/
func (a *VideoPlaylistsAPIService) ApiV1VideoPlaylistsPlaylistIdPut(ctx context.Context, playlistId int32) ApiApiV1VideoPlaylistsPlaylistIdPutRequest {
	return ApiApiV1VideoPlaylistsPlaylistIdPutRequest{
		ApiService: a,
		ctx:        ctx,
		playlistId: playlistId,
	}
}

// Execute executes the request
func (a *VideoPlaylistsAPIService) ApiV1VideoPlaylistsPlaylistIdPutExecute(r ApiApiV1VideoPlaylistsPlaylistIdPutRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPut
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VideoPlaylistsAPIService.ApiV1VideoPlaylistsPlaylistIdPut")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/video-playlists/{playlistId}"
	localVarPath = strings.Replace(localVarPath, "{"+"playlistId"+"}", url.PathEscape(parameterValueToString(r.playlistId, "playlistId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.playlistId < 1 {
		return nil, utils.ReportError("playlistId must be greater than 1")
	}

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
	if r.displayName != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "displayName", r.displayName, "", "")
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
	if r.privacy != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "privacy", r.privacy, "", "")
	}
	if r.description != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "description", r.description, "", "")
	}
	if r.videoChannelId != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "videoChannelId", r.videoChannelId, "", "")
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

type ApiDelVideoPlaylistVideoRequest struct {
	ctx               context.Context
	ApiService        *VideoPlaylistsAPIService
	playlistId        int32
	playlistElementId int32
}

func (r ApiDelVideoPlaylistVideoRequest) Execute() (*http.Response, error) {
	return r.ApiService.DelVideoPlaylistVideoExecute(r)
}

/*
DelVideoPlaylistVideo Delete an element from a playlist

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param playlistId Playlist id
	@param playlistElementId Playlist element id
	@return ApiDelVideoPlaylistVideoRequest
*/
func (a *VideoPlaylistsAPIService) DelVideoPlaylistVideo(ctx context.Context, playlistId int32, playlistElementId int32) ApiDelVideoPlaylistVideoRequest {
	return ApiDelVideoPlaylistVideoRequest{
		ApiService:        a,
		ctx:               ctx,
		playlistId:        playlistId,
		playlistElementId: playlistElementId,
	}
}

// Execute executes the request
func (a *VideoPlaylistsAPIService) DelVideoPlaylistVideoExecute(r ApiDelVideoPlaylistVideoRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VideoPlaylistsAPIService.DelVideoPlaylistVideo")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/video-playlists/{playlistId}/videos/{playlistElementId}"
	localVarPath = strings.Replace(localVarPath, "{"+"playlistId"+"}", url.PathEscape(parameterValueToString(r.playlistId, "playlistId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"playlistElementId"+"}", url.PathEscape(parameterValueToString(r.playlistElementId, "playlistElementId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.playlistId < 1 {
		return nil, utils.ReportError("playlistId must be greater than 1")
	}
	if r.playlistElementId < 1 {
		return nil, utils.ReportError("playlistElementId must be greater than 1")
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

type ApiGetPlaylistPrivacyPoliciesRequest struct {
	ctx        context.Context
	ApiService *VideoPlaylistsAPIService
}

func (r ApiGetPlaylistPrivacyPoliciesRequest) Execute() ([]string, *http.Response, error) {
	return r.ApiService.GetPlaylistPrivacyPoliciesExecute(r)
}

/*
GetPlaylistPrivacyPolicies List available playlist privacy policies

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetPlaylistPrivacyPoliciesRequest
*/
func (a *VideoPlaylistsAPIService) GetPlaylistPrivacyPolicies(ctx context.Context) ApiGetPlaylistPrivacyPoliciesRequest {
	return ApiGetPlaylistPrivacyPoliciesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return []string
func (a *VideoPlaylistsAPIService) GetPlaylistPrivacyPoliciesExecute(r ApiGetPlaylistPrivacyPoliciesRequest) ([]string, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VideoPlaylistsAPIService.GetPlaylistPrivacyPolicies")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/video-playlists/privacies"

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

type ApiGetPlaylistsRequest struct {
	ctx          context.Context
	ApiService   *VideoPlaylistsAPIService
	start        *int32
	count        *int32
	sort         *string
	playlistType *models.VideoPlaylistTypeSet
}

// Offset used to paginate results
func (r ApiGetPlaylistsRequest) Start(start int32) ApiGetPlaylistsRequest {
	r.start = &start
	return r
}

// Number of items to return
func (r ApiGetPlaylistsRequest) Count(count int32) ApiGetPlaylistsRequest {
	r.count = &count
	return r
}

// Sort column
func (r ApiGetPlaylistsRequest) Sort(sort string) ApiGetPlaylistsRequest {
	r.sort = &sort
	return r
}

func (r ApiGetPlaylistsRequest) PlaylistType(playlistType models.VideoPlaylistTypeSet) ApiGetPlaylistsRequest {
	r.playlistType = &playlistType
	return r
}

func (r ApiGetPlaylistsRequest) Execute() (*models.ApiV1VideoChannelsChannelHandleVideoPlaylistsGet200Response, *http.Response, error) {
	return r.ApiService.GetPlaylistsExecute(r)
}

/*
GetPlaylists List video playlists

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetPlaylistsRequest
*/
func (a *VideoPlaylistsAPIService) GetPlaylists(ctx context.Context) ApiGetPlaylistsRequest {
	return ApiGetPlaylistsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ApiV1VideoChannelsChannelHandleVideoPlaylistsGet200Response
func (a *VideoPlaylistsAPIService) GetPlaylistsExecute(r ApiGetPlaylistsRequest) (*models.ApiV1VideoChannelsChannelHandleVideoPlaylistsGet200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.ApiV1VideoChannelsChannelHandleVideoPlaylistsGet200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VideoPlaylistsAPIService.GetPlaylists")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/video-playlists"

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

type ApiGetVideoPlaylistVideosRequest struct {
	ctx        context.Context
	ApiService *VideoPlaylistsAPIService
	playlistId int32
	start      *int32
	count      *int32
}

// Offset used to paginate results
func (r ApiGetVideoPlaylistVideosRequest) Start(start int32) ApiGetVideoPlaylistVideosRequest {
	r.start = &start
	return r
}

// Number of items to return
func (r ApiGetVideoPlaylistVideosRequest) Count(count int32) ApiGetVideoPlaylistVideosRequest {
	r.count = &count
	return r
}

func (r ApiGetVideoPlaylistVideosRequest) Execute() (*models.VideoListResponse, *http.Response, error) {
	return r.ApiService.GetVideoPlaylistVideosExecute(r)
}

/*
GetVideoPlaylistVideos List videos of a playlist

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param playlistId Playlist id
	@return ApiGetVideoPlaylistVideosRequest
*/
func (a *VideoPlaylistsAPIService) GetVideoPlaylistVideos(ctx context.Context, playlistId int32) ApiGetVideoPlaylistVideosRequest {
	return ApiGetVideoPlaylistVideosRequest{
		ApiService: a,
		ctx:        ctx,
		playlistId: playlistId,
	}
}

// Execute executes the request
//
//	@return VideoListResponse
func (a *VideoPlaylistsAPIService) GetVideoPlaylistVideosExecute(r ApiGetVideoPlaylistVideosRequest) (*models.VideoListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.VideoListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VideoPlaylistsAPIService.GetVideoPlaylistVideos")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/video-playlists/{playlistId}/videos"
	localVarPath = strings.Replace(localVarPath, "{"+"playlistId"+"}", url.PathEscape(parameterValueToString(r.playlistId, "playlistId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.playlistId < 1 {
		return localVarReturnValue, nil, utils.ReportError("playlistId must be greater than 1")
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

type ApiPutVideoPlaylistVideoRequest struct {
	ctx                          context.Context
	ApiService                   *VideoPlaylistsAPIService
	playlistId                   int32
	playlistElementId            int32
	putVideoPlaylistVideoRequest *models.PutVideoPlaylistVideoRequest
}

func (r ApiPutVideoPlaylistVideoRequest) PutVideoPlaylistVideoRequest(putVideoPlaylistVideoRequest models.PutVideoPlaylistVideoRequest) ApiPutVideoPlaylistVideoRequest {
	r.putVideoPlaylistVideoRequest = &putVideoPlaylistVideoRequest
	return r
}

func (r ApiPutVideoPlaylistVideoRequest) Execute() (*http.Response, error) {
	return r.ApiService.PutVideoPlaylistVideoExecute(r)
}

/*
PutVideoPlaylistVideo Update a playlist element

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param playlistId Playlist id
	@param playlistElementId Playlist element id
	@return ApiPutVideoPlaylistVideoRequest
*/
func (a *VideoPlaylistsAPIService) PutVideoPlaylistVideo(ctx context.Context, playlistId int32, playlistElementId int32) ApiPutVideoPlaylistVideoRequest {
	return ApiPutVideoPlaylistVideoRequest{
		ApiService:        a,
		ctx:               ctx,
		playlistId:        playlistId,
		playlistElementId: playlistElementId,
	}
}

// Execute executes the request
func (a *VideoPlaylistsAPIService) PutVideoPlaylistVideoExecute(r ApiPutVideoPlaylistVideoRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPut
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VideoPlaylistsAPIService.PutVideoPlaylistVideo")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/video-playlists/{playlistId}/videos/{playlistElementId}"
	localVarPath = strings.Replace(localVarPath, "{"+"playlistId"+"}", url.PathEscape(parameterValueToString(r.playlistId, "playlistId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"playlistElementId"+"}", url.PathEscape(parameterValueToString(r.playlistElementId, "playlistElementId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.playlistId < 1 {
		return nil, utils.ReportError("playlistId must be greater than 1")
	}
	if r.playlistElementId < 1 {
		return nil, utils.ReportError("playlistElementId must be greater than 1")
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
	localVarPostBody = r.putVideoPlaylistVideoRequest
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

type ApiReorderVideoPlaylistRequest struct {
	ctx                         context.Context
	ApiService                  *VideoPlaylistsAPIService
	playlistId                  int32
	reorderVideoPlaylistRequest *models.ReorderVideoPlaylistRequest
}

func (r ApiReorderVideoPlaylistRequest) ReorderVideoPlaylistRequest(reorderVideoPlaylistRequest models.ReorderVideoPlaylistRequest) ApiReorderVideoPlaylistRequest {
	r.reorderVideoPlaylistRequest = &reorderVideoPlaylistRequest
	return r
}

func (r ApiReorderVideoPlaylistRequest) Execute() (*http.Response, error) {
	return r.ApiService.ReorderVideoPlaylistExecute(r)
}

/*
ReorderVideoPlaylist Reorder a playlist

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param playlistId Playlist id
	@return ApiReorderVideoPlaylistRequest
*/
func (a *VideoPlaylistsAPIService) ReorderVideoPlaylist(ctx context.Context, playlistId int32) ApiReorderVideoPlaylistRequest {
	return ApiReorderVideoPlaylistRequest{
		ApiService: a,
		ctx:        ctx,
		playlistId: playlistId,
	}
}

// Execute executes the request
func (a *VideoPlaylistsAPIService) ReorderVideoPlaylistExecute(r ApiReorderVideoPlaylistRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VideoPlaylistsAPIService.ReorderVideoPlaylist")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/video-playlists/{playlistId}/videos/reorder"
	localVarPath = strings.Replace(localVarPath, "{"+"playlistId"+"}", url.PathEscape(parameterValueToString(r.playlistId, "playlistId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.playlistId < 1 {
		return nil, utils.ReportError("playlistId must be greater than 1")
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
	localVarPostBody = r.reorderVideoPlaylistRequest
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
