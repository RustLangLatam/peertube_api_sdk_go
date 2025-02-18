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

// VideoImportsAPIService VideoImportsAPI service
type VideoImportsAPIService service

type ApiApiV1VideosImportsIdCancelPostRequest struct {
	ctx        context.Context
	ApiService *VideoImportsAPIService
	id         int32
}

func (r ApiApiV1VideosImportsIdCancelPostRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiV1VideosImportsIdCancelPostExecute(r)
}

/*
ApiV1VideosImportsIdCancelPost Cancel video import

Cancel a pending video import

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id Entity id
	@return ApiApiV1VideosImportsIdCancelPostRequest
*/
func (a *VideoImportsAPIService) ApiV1VideosImportsIdCancelPost(ctx context.Context, id int32) ApiApiV1VideosImportsIdCancelPostRequest {
	return ApiApiV1VideosImportsIdCancelPostRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
func (a *VideoImportsAPIService) ApiV1VideosImportsIdCancelPostExecute(r ApiApiV1VideosImportsIdCancelPostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VideoImportsAPIService.ApiV1VideosImportsIdCancelPost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/videos/imports/{id}/cancel"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.id < 1 {
		return nil, utils.ReportError("id must be greater than 1")
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

type ApiApiV1VideosImportsIdDeleteRequest struct {
	ctx        context.Context
	ApiService *VideoImportsAPIService
	id         int32
}

func (r ApiApiV1VideosImportsIdDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiV1VideosImportsIdDeleteExecute(r)
}

/*
ApiV1VideosImportsIdDelete Delete video import

Delete ended video import

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id Entity id
	@return ApiApiV1VideosImportsIdDeleteRequest
*/
func (a *VideoImportsAPIService) ApiV1VideosImportsIdDelete(ctx context.Context, id int32) ApiApiV1VideosImportsIdDeleteRequest {
	return ApiApiV1VideosImportsIdDeleteRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
func (a *VideoImportsAPIService) ApiV1VideosImportsIdDeleteExecute(r ApiApiV1VideosImportsIdDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VideoImportsAPIService.ApiV1VideosImportsIdDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/videos/imports/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.id < 1 {
		return nil, utils.ReportError("id must be greater than 1")
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

type ApiImportVideoRequest struct {
	ctx                   context.Context
	ApiService            *VideoImportsAPIService
	name                  *string
	channelId             *int32
	targetUrl             *string
	magnetUri             *string
	torrentfile           **os.File
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
func (r ApiImportVideoRequest) Name(name string) ApiImportVideoRequest {
	r.name = &name
	return r
}

// Channel id that will contain this video
func (r ApiImportVideoRequest) ChannelId(channelId int32) ApiImportVideoRequest {
	r.channelId = &channelId
	return r
}

// remote URL where to find the import&#39;s source video
func (r ApiImportVideoRequest) TargetUrl(targetUrl string) ApiImportVideoRequest {
	r.targetUrl = &targetUrl
	return r
}

// magnet URI allowing to resolve the import&#39;s source video
func (r ApiImportVideoRequest) MagnetUri(magnetUri string) ApiImportVideoRequest {
	r.magnetUri = &magnetUri
	return r
}

// Torrent file containing only the video file
func (r ApiImportVideoRequest) Torrentfile(torrentfile *os.File) ApiImportVideoRequest {
	r.torrentfile = &torrentfile
	return r
}

func (r ApiImportVideoRequest) Privacy(privacy models.VideoPrivacySet) ApiImportVideoRequest {
	r.privacy = &privacy
	return r
}

// category id of the video (see [/videos/categories](#operation/getCategories))
func (r ApiImportVideoRequest) Category(category int32) ApiImportVideoRequest {
	r.category = &category
	return r
}

// licence id of the video (see [/videos/licences](#operation/getLicences))
func (r ApiImportVideoRequest) Licence(licence int32) ApiImportVideoRequest {
	r.licence = &licence
	return r
}

// language id of the video (see [/videos/languages](#operation/getLanguages))
func (r ApiImportVideoRequest) Language(language string) ApiImportVideoRequest {
	r.language = &language
	return r
}

// Video description
func (r ApiImportVideoRequest) Description(description string) ApiImportVideoRequest {
	r.description = &description
	return r
}

// Whether or not we wait transcoding before publish the video
func (r ApiImportVideoRequest) WaitTranscoding(waitTranscoding bool) ApiImportVideoRequest {
	r.waitTranscoding = &waitTranscoding
	return r
}

// **PeerTube &gt;&#x3D; 6.2** If enabled by the admin, automatically generate a subtitle of the video
func (r ApiImportVideoRequest) GenerateTranscription(generateTranscription bool) ApiImportVideoRequest {
	r.generateTranscription = &generateTranscription
	return r
}

// A text tell the audience how to support the video creator
func (r ApiImportVideoRequest) Support(support string) ApiImportVideoRequest {
	r.support = &support
	return r
}

// Whether or not this video contains sensitive content
func (r ApiImportVideoRequest) Nsfw(nsfw bool) ApiImportVideoRequest {
	r.nsfw = &nsfw
	return r
}

// Video tags (maximum 5 tags each between 2 and 30 characters)
func (r ApiImportVideoRequest) Tags(tags []string) ApiImportVideoRequest {
	r.tags = &tags
	return r
}

// Deprecated in 6.2, use commentsPolicy instead
func (r ApiImportVideoRequest) CommentsEnabled(commentsEnabled bool) ApiImportVideoRequest {
	r.commentsEnabled = &commentsEnabled
	return r
}

func (r ApiImportVideoRequest) CommentsPolicy(commentsPolicy models.VideoCommentsPolicySet) ApiImportVideoRequest {
	r.commentsPolicy = &commentsPolicy
	return r
}

// Enable or disable downloading for this video
func (r ApiImportVideoRequest) DownloadEnabled(downloadEnabled bool) ApiImportVideoRequest {
	r.downloadEnabled = &downloadEnabled
	return r
}

// Date when the content was originally published
func (r ApiImportVideoRequest) OriginallyPublishedAt(originallyPublishedAt time.Time) ApiImportVideoRequest {
	r.originallyPublishedAt = &originallyPublishedAt
	return r
}

func (r ApiImportVideoRequest) ScheduleUpdate(scheduleUpdate models.VideoScheduledUpdate) ApiImportVideoRequest {
	r.scheduleUpdate = &scheduleUpdate
	return r
}

// Video thumbnail file
func (r ApiImportVideoRequest) Thumbnailfile(thumbnailfile *os.File) ApiImportVideoRequest {
	r.thumbnailfile = thumbnailfile
	return r
}

// Video preview file
func (r ApiImportVideoRequest) Previewfile(previewfile *os.File) ApiImportVideoRequest {
	r.previewfile = previewfile
	return r
}

func (r ApiImportVideoRequest) VideoPasswords(videoPasswords []string) ApiImportVideoRequest {
	r.videoPasswords = &videoPasswords
	return r
}

func (r ApiImportVideoRequest) Execute() (*models.VideoUploadResponse, *http.Response, error) {
	return r.ApiService.ImportVideoExecute(r)
}

/*
ImportVideo Import a video

Import a torrent or magnetURI or HTTP resource (if enabled by the instance administrator)

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiImportVideoRequest
*/
func (a *VideoImportsAPIService) ImportVideo(ctx context.Context) ApiImportVideoRequest {
	return ApiImportVideoRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return VideoUploadResponse
func (a *VideoImportsAPIService) ImportVideoExecute(r ApiImportVideoRequest) (*models.VideoUploadResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.VideoUploadResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VideoImportsAPIService.ImportVideo")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/videos/imports"

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
	if r.targetUrl != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "targetUrl", r.targetUrl, "", "")
	}
	if r.magnetUri != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "magnetUri", r.magnetUri, "", "")
	}
	if r.torrentfile != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "torrentfile", r.torrentfile, "form", "")
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
