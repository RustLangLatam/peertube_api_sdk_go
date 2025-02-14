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
	"os"
	"strings"
)

// LiveVideosAPIService LiveVideosAPI service
type LiveVideosAPIService service

type ApiAddLiveRequest struct {
	ctx             context.Context
	ApiService      *LiveVideosAPIService
	channelId       *int32
	name            *string
	saveReplay      *bool
	replaySettings  *LiveVideoReplaySettings
	permanentLive   *bool
	latencyMode     *LiveVideoLatencyMode
	thumbnailfile   *os.File
	previewfile     *os.File
	privacy         *VideoPrivacySet
	category        *int32
	licence         *int32
	language        *string
	description     *string
	support         *string
	nsfw            *bool
	tags            *[]string
	commentsEnabled *bool
	commentsPolicy  *VideoCommentsPolicySet
	downloadEnabled *bool
}

// Channel id that will contain this live video
func (r ApiAddLiveRequest) ChannelId(channelId int32) ApiAddLiveRequest {
	r.channelId = &channelId
	return r
}

// Live video/replay name
func (r ApiAddLiveRequest) Name(name string) ApiAddLiveRequest {
	r.name = &name
	return r
}

func (r ApiAddLiveRequest) SaveReplay(saveReplay bool) ApiAddLiveRequest {
	r.saveReplay = &saveReplay
	return r
}

func (r ApiAddLiveRequest) ReplaySettings(replaySettings LiveVideoReplaySettings) ApiAddLiveRequest {
	r.replaySettings = &replaySettings
	return r
}

// User can stream multiple times in a permanent live
func (r ApiAddLiveRequest) PermanentLive(permanentLive bool) ApiAddLiveRequest {
	r.permanentLive = &permanentLive
	return r
}

func (r ApiAddLiveRequest) LatencyMode(latencyMode LiveVideoLatencyMode) ApiAddLiveRequest {
	r.latencyMode = &latencyMode
	return r
}

// Live video/replay thumbnail file
func (r ApiAddLiveRequest) Thumbnailfile(thumbnailfile *os.File) ApiAddLiveRequest {
	r.thumbnailfile = thumbnailfile
	return r
}

// Live video/replay preview file
func (r ApiAddLiveRequest) Previewfile(previewfile *os.File) ApiAddLiveRequest {
	r.previewfile = previewfile
	return r
}

func (r ApiAddLiveRequest) Privacy(privacy VideoPrivacySet) ApiAddLiveRequest {
	r.privacy = &privacy
	return r
}

// category id of the video (see [/videos/categories](#operation/getCategories))
func (r ApiAddLiveRequest) Category(category int32) ApiAddLiveRequest {
	r.category = &category
	return r
}

// licence id of the video (see [/videos/licences](#operation/getLicences))
func (r ApiAddLiveRequest) Licence(licence int32) ApiAddLiveRequest {
	r.licence = &licence
	return r
}

// language id of the video (see [/videos/languages](#operation/getLanguages))
func (r ApiAddLiveRequest) Language(language string) ApiAddLiveRequest {
	r.language = &language
	return r
}

// Live video/replay description
func (r ApiAddLiveRequest) Description(description string) ApiAddLiveRequest {
	r.description = &description
	return r
}

// A text tell the audience how to support the creator
func (r ApiAddLiveRequest) Support(support string) ApiAddLiveRequest {
	r.support = &support
	return r
}

// Whether or not this live video/replay contains sensitive content
func (r ApiAddLiveRequest) Nsfw(nsfw bool) ApiAddLiveRequest {
	r.nsfw = &nsfw
	return r
}

// Live video/replay tags (maximum 5 tags each between 2 and 30 characters)
func (r ApiAddLiveRequest) Tags(tags []string) ApiAddLiveRequest {
	r.tags = &tags
	return r
}

// Deprecated in 6.2, use commentsPolicy instead
func (r ApiAddLiveRequest) CommentsEnabled(commentsEnabled bool) ApiAddLiveRequest {
	r.commentsEnabled = &commentsEnabled
	return r
}

func (r ApiAddLiveRequest) CommentsPolicy(commentsPolicy VideoCommentsPolicySet) ApiAddLiveRequest {
	r.commentsPolicy = &commentsPolicy
	return r
}

// Enable or disable downloading for the replay of this live video
func (r ApiAddLiveRequest) DownloadEnabled(downloadEnabled bool) ApiAddLiveRequest {
	r.downloadEnabled = &downloadEnabled
	return r
}

func (r ApiAddLiveRequest) Execute() (*VideoUploadResponse, *http.Response, error) {
	return r.ApiService.AddLiveExecute(r)
}

/*
AddLive Create a live

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAddLiveRequest
*/
func (a *LiveVideosAPIService) AddLive(ctx context.Context) ApiAddLiveRequest {
	return ApiAddLiveRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return VideoUploadResponse
func (a *LiveVideosAPIService) AddLiveExecute(r ApiAddLiveRequest) (*VideoUploadResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *VideoUploadResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LiveVideosAPIService.AddLive")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/videos/live"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.channelId == nil {
		return localVarReturnValue, nil, reportError("channelId is required and must be specified")
	}
	if r.name == nil {
		return localVarReturnValue, nil, reportError("name is required and must be specified")
	}
	if strlen(*r.name) < 3 {
		return localVarReturnValue, nil, reportError("name must have at least 3 elements")
	}
	if strlen(*r.name) > 120 {
		return localVarReturnValue, nil, reportError("name must have less than 120 elements")
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
	parameterAddToHeaderOrQuery(localVarFormParams, "channelId", r.channelId, "", "")
	if r.saveReplay != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "saveReplay", r.saveReplay, "", "")
	}
	if r.replaySettings != nil {
		paramJson, err := parameterToJson(*r.replaySettings)
		if err != nil {
			return localVarReturnValue, nil, err
		}
		localVarFormParams.Add("replaySettings", paramJson)
	}
	if r.permanentLive != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "permanentLive", r.permanentLive, "", "")
	}
	if r.latencyMode != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "latencyMode", r.latencyMode, "", "")
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
	if r.support != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "support", r.support, "", "")
	}
	if r.nsfw != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "nsfw", r.nsfw, "", "")
	}
	parameterAddToHeaderOrQuery(localVarFormParams, "name", r.name, "", "")
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

type ApiApiV1VideosIdLiveSessionGetRequest struct {
	ctx                    context.Context
	ApiService             *LiveVideosAPIService
	id                     ApiV1VideosOwnershipIdAcceptPostIdParameter
	xPeertubeVideoPassword *string
}

// Required on password protected video
func (r ApiApiV1VideosIdLiveSessionGetRequest) XPeertubeVideoPassword(xPeertubeVideoPassword string) ApiApiV1VideosIdLiveSessionGetRequest {
	r.xPeertubeVideoPassword = &xPeertubeVideoPassword
	return r
}

func (r ApiApiV1VideosIdLiveSessionGetRequest) Execute() (*LiveVideoSessionResponse, *http.Response, error) {
	return r.ApiService.ApiV1VideosIdLiveSessionGetExecute(r)
}

/*
ApiV1VideosIdLiveSessionGet Get live session of a replay

If the video is a replay of a live, you can find the associated live session using this endpoint

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id The object id, uuid or short uuid
	@return ApiApiV1VideosIdLiveSessionGetRequest
*/
func (a *LiveVideosAPIService) ApiV1VideosIdLiveSessionGet(ctx context.Context, id ApiV1VideosOwnershipIdAcceptPostIdParameter) ApiApiV1VideosIdLiveSessionGetRequest {
	return ApiApiV1VideosIdLiveSessionGetRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return LiveVideoSessionResponse
func (a *LiveVideosAPIService) ApiV1VideosIdLiveSessionGetExecute(r ApiApiV1VideosIdLiveSessionGetRequest) (*LiveVideoSessionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *LiveVideoSessionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LiveVideosAPIService.ApiV1VideosIdLiveSessionGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/videos/{id}/live-session"
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

type ApiApiV1VideosLiveIdSessionsGetRequest struct {
	ctx        context.Context
	ApiService *LiveVideosAPIService
	id         ApiV1VideosOwnershipIdAcceptPostIdParameter
}

func (r ApiApiV1VideosLiveIdSessionsGetRequest) Execute() (*ApiV1VideosLiveIdSessionsGet200Response, *http.Response, error) {
	return r.ApiService.ApiV1VideosLiveIdSessionsGetExecute(r)
}

/*
ApiV1VideosLiveIdSessionsGet List live sessions

List all sessions created in a particular live

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id The object id, uuid or short uuid
	@return ApiApiV1VideosLiveIdSessionsGetRequest
*/
func (a *LiveVideosAPIService) ApiV1VideosLiveIdSessionsGet(ctx context.Context, id ApiV1VideosOwnershipIdAcceptPostIdParameter) ApiApiV1VideosLiveIdSessionsGetRequest {
	return ApiApiV1VideosLiveIdSessionsGetRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return ApiV1VideosLiveIdSessionsGet200Response
func (a *LiveVideosAPIService) ApiV1VideosLiveIdSessionsGetExecute(r ApiApiV1VideosLiveIdSessionsGetRequest) (*ApiV1VideosLiveIdSessionsGet200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ApiV1VideosLiveIdSessionsGet200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LiveVideosAPIService.ApiV1VideosLiveIdSessionsGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/videos/live/{id}/sessions"
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

type ApiGetLiveIdRequest struct {
	ctx        context.Context
	ApiService *LiveVideosAPIService
	id         ApiV1VideosOwnershipIdAcceptPostIdParameter
}

func (r ApiGetLiveIdRequest) Execute() (*LiveVideoResponse, *http.Response, error) {
	return r.ApiService.GetLiveIdExecute(r)
}

/*
GetLiveId Get information about a live

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id The object id, uuid or short uuid
	@return ApiGetLiveIdRequest
*/
func (a *LiveVideosAPIService) GetLiveId(ctx context.Context, id ApiV1VideosOwnershipIdAcceptPostIdParameter) ApiGetLiveIdRequest {
	return ApiGetLiveIdRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return LiveVideoResponse
func (a *LiveVideosAPIService) GetLiveIdExecute(r ApiGetLiveIdRequest) (*LiveVideoResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *LiveVideoResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LiveVideosAPIService.GetLiveId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/videos/live/{id}"
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

type ApiUpdateLiveIdRequest struct {
	ctx             context.Context
	ApiService      *LiveVideosAPIService
	id              ApiV1VideosOwnershipIdAcceptPostIdParameter
	liveVideoUpdate *LiveVideoUpdate
}

func (r ApiUpdateLiveIdRequest) LiveVideoUpdate(liveVideoUpdate LiveVideoUpdate) ApiUpdateLiveIdRequest {
	r.liveVideoUpdate = &liveVideoUpdate
	return r
}

func (r ApiUpdateLiveIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.UpdateLiveIdExecute(r)
}

/*
UpdateLiveId Update information about a live

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id The object id, uuid or short uuid
	@return ApiUpdateLiveIdRequest
*/
func (a *LiveVideosAPIService) UpdateLiveId(ctx context.Context, id ApiV1VideosOwnershipIdAcceptPostIdParameter) ApiUpdateLiveIdRequest {
	return ApiUpdateLiveIdRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
func (a *LiveVideosAPIService) UpdateLiveIdExecute(r ApiUpdateLiveIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPut
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LiveVideosAPIService.UpdateLiveId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/videos/live/{id}"
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
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.liveVideoUpdate
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
