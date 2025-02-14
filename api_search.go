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
	"reflect"
	"time"
)

// SearchAPIService SearchAPI service
type SearchAPIService service

type ApiSearchChannelsRequest struct {
	ctx          context.Context
	ApiService   *SearchAPIService
	search       *string
	start        *int32
	count        *int32
	searchTarget *string
	sort         *string
	host         *string
	handles      *[]string
}

// String to search. If the user can make a remote URI search, and the string is an URI then the PeerTube instance will fetch the remote object and add it to its database. Then, you can use the REST API to fetch the complete channel information and interact with it.
func (r ApiSearchChannelsRequest) Search(search string) ApiSearchChannelsRequest {
	r.search = &search
	return r
}

// Offset used to paginate results
func (r ApiSearchChannelsRequest) Start(start int32) ApiSearchChannelsRequest {
	r.start = &start
	return r
}

// Number of items to return
func (r ApiSearchChannelsRequest) Count(count int32) ApiSearchChannelsRequest {
	r.count = &count
	return r
}

// If the administrator enabled search index support, you can override the default search target.  **Warning**: If you choose to make an index search, PeerTube will get results from a third party service. It means the instance may not yet know the objects you fetched. If you want to load video/channel information:   * If the current user has the ability to make a remote URI search (this information is available in the config endpoint),   then reuse the search API to make a search using the object URI so PeerTube instance fetches the remote object and fill its database.   After that, you can use the classic REST API endpoints to fetch the complete object or interact with it   * If the current user doesn&#39;t have the ability to make a remote URI search, then redirect the user on the origin instance or fetch   the data from the origin instance API
func (r ApiSearchChannelsRequest) SearchTarget(searchTarget string) ApiSearchChannelsRequest {
	r.searchTarget = &searchTarget
	return r
}

// Sort column
func (r ApiSearchChannelsRequest) Sort(sort string) ApiSearchChannelsRequest {
	r.sort = &sort
	return r
}

// Find elements owned by this host
func (r ApiSearchChannelsRequest) Host(host string) ApiSearchChannelsRequest {
	r.host = &host
	return r
}

// Find elements with these handles
func (r ApiSearchChannelsRequest) Handles(handles []string) ApiSearchChannelsRequest {
	r.handles = &handles
	return r
}

func (r ApiSearchChannelsRequest) Execute() (*VideoChannelList, *http.Response, error) {
	return r.ApiService.SearchChannelsExecute(r)
}

/*
SearchChannels Search channels

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiSearchChannelsRequest
*/
func (a *SearchAPIService) SearchChannels(ctx context.Context) ApiSearchChannelsRequest {
	return ApiSearchChannelsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return VideoChannelList
func (a *SearchAPIService) SearchChannelsExecute(r ApiSearchChannelsRequest) (*VideoChannelList, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *VideoChannelList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SearchAPIService.SearchChannels")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/search/video-channels"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.search == nil {
		return localVarReturnValue, nil, reportError("search is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "search", r.search, "form", "")
	if r.start != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start", r.start, "form", "")
	}
	if r.count != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "count", r.count, "form", "")
	} else {
		var defaultValue int32 = 15
		r.count = &defaultValue
	}
	if r.searchTarget != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "searchTarget", r.searchTarget, "form", "")
	}
	if r.sort != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort", r.sort, "form", "")
	}
	if r.host != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "host", r.host, "form", "")
	}
	if r.handles != nil {
		t := *r.handles
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "handles", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "handles", t, "form", "multi")
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

type ApiSearchPlaylistsRequest struct {
	ctx          context.Context
	ApiService   *SearchAPIService
	search       *string
	start        *int32
	count        *int32
	searchTarget *string
	sort         *string
	host         *string
	uuids        *[]string
}

// String to search. If the user can make a remote URI search, and the string is an URI then the PeerTube instance will fetch the remote object and add it to its database. Then, you can use the REST API to fetch the complete playlist information and interact with it.
func (r ApiSearchPlaylistsRequest) Search(search string) ApiSearchPlaylistsRequest {
	r.search = &search
	return r
}

// Offset used to paginate results
func (r ApiSearchPlaylistsRequest) Start(start int32) ApiSearchPlaylistsRequest {
	r.start = &start
	return r
}

// Number of items to return
func (r ApiSearchPlaylistsRequest) Count(count int32) ApiSearchPlaylistsRequest {
	r.count = &count
	return r
}

// If the administrator enabled search index support, you can override the default search target.  **Warning**: If you choose to make an index search, PeerTube will get results from a third party service. It means the instance may not yet know the objects you fetched. If you want to load video/channel information:   * If the current user has the ability to make a remote URI search (this information is available in the config endpoint),   then reuse the search API to make a search using the object URI so PeerTube instance fetches the remote object and fill its database.   After that, you can use the classic REST API endpoints to fetch the complete object or interact with it   * If the current user doesn&#39;t have the ability to make a remote URI search, then redirect the user on the origin instance or fetch   the data from the origin instance API
func (r ApiSearchPlaylistsRequest) SearchTarget(searchTarget string) ApiSearchPlaylistsRequest {
	r.searchTarget = &searchTarget
	return r
}

// Sort column
func (r ApiSearchPlaylistsRequest) Sort(sort string) ApiSearchPlaylistsRequest {
	r.sort = &sort
	return r
}

// Find elements owned by this host
func (r ApiSearchPlaylistsRequest) Host(host string) ApiSearchPlaylistsRequest {
	r.host = &host
	return r
}

// Find elements with specific UUIDs
func (r ApiSearchPlaylistsRequest) Uuids(uuids []string) ApiSearchPlaylistsRequest {
	r.uuids = &uuids
	return r
}

func (r ApiSearchPlaylistsRequest) Execute() (*ApiV1VideoChannelsChannelHandleVideoPlaylistsGet200Response, *http.Response, error) {
	return r.ApiService.SearchPlaylistsExecute(r)
}

/*
SearchPlaylists Search playlists

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiSearchPlaylistsRequest
*/
func (a *SearchAPIService) SearchPlaylists(ctx context.Context) ApiSearchPlaylistsRequest {
	return ApiSearchPlaylistsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ApiV1VideoChannelsChannelHandleVideoPlaylistsGet200Response
func (a *SearchAPIService) SearchPlaylistsExecute(r ApiSearchPlaylistsRequest) (*ApiV1VideoChannelsChannelHandleVideoPlaylistsGet200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ApiV1VideoChannelsChannelHandleVideoPlaylistsGet200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SearchAPIService.SearchPlaylists")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/search/video-playlists"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.search == nil {
		return localVarReturnValue, nil, reportError("search is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "search", r.search, "form", "")
	if r.start != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start", r.start, "form", "")
	}
	if r.count != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "count", r.count, "form", "")
	} else {
		var defaultValue int32 = 15
		r.count = &defaultValue
	}
	if r.searchTarget != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "searchTarget", r.searchTarget, "form", "")
	}
	if r.sort != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort", r.sort, "form", "")
	}
	if r.host != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "host", r.host, "form", "")
	}
	if r.uuids != nil {
		t := *r.uuids
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "uuids", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "uuids", t, "form", "multi")
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

type ApiSearchVideosRequest struct {
	ctx                          context.Context
	ApiService                   *SearchAPIService
	search                       *string
	categoryOneOf                *GetAccountVideosCategoryOneOfParameter
	isLive                       *bool
	tagsOneOf                    *GetAccountVideosTagsOneOfParameter
	tagsAllOf                    *GetAccountVideosTagsAllOfParameter
	licenceOneOf                 *GetAccountVideosLicenceOneOfParameter
	languageOneOf                *GetAccountVideosLanguageOneOfParameter
	autoTagOneOf                 *GetAccountVideosTagsAllOfParameter
	nsfw                         *string
	isLocal                      *bool
	include                      *int32
	privacyOneOf                 *VideoPrivacySet
	uuids                        *[]string
	hasHLSFiles                  *bool
	hasWebVideoFiles             *bool
	skipCount                    *string
	start                        *int32
	count                        *int32
	searchTarget                 *string
	sort                         *string
	excludeAlreadyWatched        *bool
	host                         *string
	startDate                    *time.Time
	endDate                      *time.Time
	originallyPublishedStartDate *time.Time
	originallyPublishedEndDate   *time.Time
	durationMin                  *int32
	durationMax                  *int32
}

// String to search. If the user can make a remote URI search, and the string is an URI then the PeerTube instance will fetch the remote object and add it to its database. Then, you can use the REST API to fetch the complete video information and interact with it.
func (r ApiSearchVideosRequest) Search(search string) ApiSearchVideosRequest {
	r.search = &search
	return r
}

// category id of the video (see [/videos/categories](#operation/getCategories))
func (r ApiSearchVideosRequest) CategoryOneOf(categoryOneOf GetAccountVideosCategoryOneOfParameter) ApiSearchVideosRequest {
	r.categoryOneOf = &categoryOneOf
	return r
}

// whether or not the video is a live
func (r ApiSearchVideosRequest) IsLive(isLive bool) ApiSearchVideosRequest {
	r.isLive = &isLive
	return r
}

// tag(s) of the video
func (r ApiSearchVideosRequest) TagsOneOf(tagsOneOf GetAccountVideosTagsOneOfParameter) ApiSearchVideosRequest {
	r.tagsOneOf = &tagsOneOf
	return r
}

// tag(s) of the video, where all should be present in the video
func (r ApiSearchVideosRequest) TagsAllOf(tagsAllOf GetAccountVideosTagsAllOfParameter) ApiSearchVideosRequest {
	r.tagsAllOf = &tagsAllOf
	return r
}

// licence id of the video (see [/videos/licences](#operation/getLicences))
func (r ApiSearchVideosRequest) LicenceOneOf(licenceOneOf GetAccountVideosLicenceOneOfParameter) ApiSearchVideosRequest {
	r.licenceOneOf = &licenceOneOf
	return r
}

// language id of the video (see [/videos/languages](#operation/getLanguages)). Use &#x60;_unknown&#x60; to filter on videos that don&#39;t have a video language
func (r ApiSearchVideosRequest) LanguageOneOf(languageOneOf GetAccountVideosLanguageOneOfParameter) ApiSearchVideosRequest {
	r.languageOneOf = &languageOneOf
	return r
}

// **PeerTube &gt;&#x3D; 6.2** **Admins and moderators only** filter on videos that contain one of these automatic tags
func (r ApiSearchVideosRequest) AutoTagOneOf(autoTagOneOf GetAccountVideosTagsAllOfParameter) ApiSearchVideosRequest {
	r.autoTagOneOf = &autoTagOneOf
	return r
}

// whether to include nsfw videos, if any
func (r ApiSearchVideosRequest) Nsfw(nsfw string) ApiSearchVideosRequest {
	r.nsfw = &nsfw
	return r
}

// **PeerTube &gt;&#x3D; 4.0** Display only local or remote objects
func (r ApiSearchVideosRequest) IsLocal(isLocal bool) ApiSearchVideosRequest {
	r.isLocal = &isLocal
	return r
}

// **Only administrators and moderators can use this parameter**  Include additional videos in results (can be combined using bitwise or operator) - &#x60;0&#x60; NONE - &#x60;1&#x60; NOT_PUBLISHED_STATE - &#x60;2&#x60; BLACKLISTED - &#x60;4&#x60; BLOCKED_OWNER - &#x60;8&#x60; FILES - &#x60;16&#x60; CAPTIONS - &#x60;32&#x60; VIDEO SOURCE
func (r ApiSearchVideosRequest) Include(include int32) ApiSearchVideosRequest {
	r.include = &include
	return r
}

// **PeerTube &gt;&#x3D; 4.0** Display only videos in this specific privacy/privacies
func (r ApiSearchVideosRequest) PrivacyOneOf(privacyOneOf VideoPrivacySet) ApiSearchVideosRequest {
	r.privacyOneOf = &privacyOneOf
	return r
}

// Find elements with specific UUIDs
func (r ApiSearchVideosRequest) Uuids(uuids []string) ApiSearchVideosRequest {
	r.uuids = &uuids
	return r
}

// **PeerTube &gt;&#x3D; 4.0** Display only videos that have HLS files
func (r ApiSearchVideosRequest) HasHLSFiles(hasHLSFiles bool) ApiSearchVideosRequest {
	r.hasHLSFiles = &hasHLSFiles
	return r
}

// **PeerTube &gt;&#x3D; 6.0** Display only videos that have Web Video files
func (r ApiSearchVideosRequest) HasWebVideoFiles(hasWebVideoFiles bool) ApiSearchVideosRequest {
	r.hasWebVideoFiles = &hasWebVideoFiles
	return r
}

// if you don&#39;t need the &#x60;total&#x60; in the response
func (r ApiSearchVideosRequest) SkipCount(skipCount string) ApiSearchVideosRequest {
	r.skipCount = &skipCount
	return r
}

// Offset used to paginate results
func (r ApiSearchVideosRequest) Start(start int32) ApiSearchVideosRequest {
	r.start = &start
	return r
}

// Number of items to return
func (r ApiSearchVideosRequest) Count(count int32) ApiSearchVideosRequest {
	r.count = &count
	return r
}

// If the administrator enabled search index support, you can override the default search target.  **Warning**: If you choose to make an index search, PeerTube will get results from a third party service. It means the instance may not yet know the objects you fetched. If you want to load video/channel information:   * If the current user has the ability to make a remote URI search (this information is available in the config endpoint),   then reuse the search API to make a search using the object URI so PeerTube instance fetches the remote object and fill its database.   After that, you can use the classic REST API endpoints to fetch the complete object or interact with it   * If the current user doesn&#39;t have the ability to make a remote URI search, then redirect the user on the origin instance or fetch   the data from the origin instance API
func (r ApiSearchVideosRequest) SearchTarget(searchTarget string) ApiSearchVideosRequest {
	r.searchTarget = &searchTarget
	return r
}

// Sort videos by criteria (prefixing with &#x60;-&#x60; means &#x60;DESC&#x60; order):
func (r ApiSearchVideosRequest) Sort(sort string) ApiSearchVideosRequest {
	r.sort = &sort
	return r
}

// Whether or not to exclude videos that are in the user&#39;s video history
func (r ApiSearchVideosRequest) ExcludeAlreadyWatched(excludeAlreadyWatched bool) ApiSearchVideosRequest {
	r.excludeAlreadyWatched = &excludeAlreadyWatched
	return r
}

// Find elements owned by this host
func (r ApiSearchVideosRequest) Host(host string) ApiSearchVideosRequest {
	r.host = &host
	return r
}

// Get videos that are published after this date
func (r ApiSearchVideosRequest) StartDate(startDate time.Time) ApiSearchVideosRequest {
	r.startDate = &startDate
	return r
}

// Get videos that are published before this date
func (r ApiSearchVideosRequest) EndDate(endDate time.Time) ApiSearchVideosRequest {
	r.endDate = &endDate
	return r
}

// Get videos that are originally published after this date
func (r ApiSearchVideosRequest) OriginallyPublishedStartDate(originallyPublishedStartDate time.Time) ApiSearchVideosRequest {
	r.originallyPublishedStartDate = &originallyPublishedStartDate
	return r
}

// Get videos that are originally published before this date
func (r ApiSearchVideosRequest) OriginallyPublishedEndDate(originallyPublishedEndDate time.Time) ApiSearchVideosRequest {
	r.originallyPublishedEndDate = &originallyPublishedEndDate
	return r
}

// Get videos that have this minimum duration
func (r ApiSearchVideosRequest) DurationMin(durationMin int32) ApiSearchVideosRequest {
	r.durationMin = &durationMin
	return r
}

// Get videos that have this maximum duration
func (r ApiSearchVideosRequest) DurationMax(durationMax int32) ApiSearchVideosRequest {
	r.durationMax = &durationMax
	return r
}

func (r ApiSearchVideosRequest) Execute() (*VideoListResponse, *http.Response, error) {
	return r.ApiService.SearchVideosExecute(r)
}

/*
SearchVideos Search videos

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiSearchVideosRequest
*/
func (a *SearchAPIService) SearchVideos(ctx context.Context) ApiSearchVideosRequest {
	return ApiSearchVideosRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return VideoListResponse
func (a *SearchAPIService) SearchVideosExecute(r ApiSearchVideosRequest) (*VideoListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *VideoListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SearchAPIService.SearchVideos")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/search/videos"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.search == nil {
		return localVarReturnValue, nil, reportError("search is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "search", r.search, "form", "")
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
	if r.uuids != nil {
		t := *r.uuids
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "uuids", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "uuids", t, "form", "multi")
		}
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
	if r.searchTarget != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "searchTarget", r.searchTarget, "form", "")
	}
	if r.sort != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort", r.sort, "form", "")
	}
	if r.excludeAlreadyWatched != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "excludeAlreadyWatched", r.excludeAlreadyWatched, "form", "")
	}
	if r.host != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "host", r.host, "form", "")
	}
	if r.startDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "startDate", r.startDate, "form", "")
	}
	if r.endDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "endDate", r.endDate, "form", "")
	}
	if r.originallyPublishedStartDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "originallyPublishedStartDate", r.originallyPublishedStartDate, "form", "")
	}
	if r.originallyPublishedEndDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "originallyPublishedEndDate", r.originallyPublishedEndDate, "form", "")
	}
	if r.durationMin != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "durationMin", r.durationMin, "form", "")
	}
	if r.durationMax != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "durationMax", r.durationMax, "form", "")
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
