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

// PluginsAPIService PluginsAPI service
type PluginsAPIService service

type ApiAddPluginRequest struct {
	ctx              context.Context
	ApiService       *PluginsAPIService
	addPluginRequest *models.AddPluginRequest
}

func (r ApiAddPluginRequest) AddPluginRequest(addPluginRequest models.AddPluginRequest) ApiAddPluginRequest {
	r.addPluginRequest = &addPluginRequest
	return r
}

func (r ApiAddPluginRequest) Execute() (*http.Response, error) {
	return r.ApiService.AddPluginExecute(r)
}

/*
AddPlugin Install a plugin

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAddPluginRequest
*/
func (a *PluginsAPIService) AddPlugin(ctx context.Context) ApiAddPluginRequest {
	return ApiAddPluginRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *PluginsAPIService) AddPluginExecute(r ApiAddPluginRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PluginsAPIService.AddPlugin")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/plugins/install"

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
	localVarPostBody = r.addPluginRequest
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

type ApiApiV1PluginsNpmNamePublicSettingsGetRequest struct {
	ctx        context.Context
	ApiService *PluginsAPIService
	npmName    string
}

func (r ApiApiV1PluginsNpmNamePublicSettingsGetRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.ApiV1PluginsNpmNamePublicSettingsGetExecute(r)
}

/*
ApiV1PluginsNpmNamePublicSettingsGet Get a plugin's public settings

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param npmName name of the plugin/theme on npmjs.com or in its package.json
	@return ApiApiV1PluginsNpmNamePublicSettingsGetRequest
*/
func (a *PluginsAPIService) ApiV1PluginsNpmNamePublicSettingsGet(ctx context.Context, npmName string) ApiApiV1PluginsNpmNamePublicSettingsGetRequest {
	return ApiApiV1PluginsNpmNamePublicSettingsGetRequest{
		ApiService: a,
		ctx:        ctx,
		npmName:    npmName,
	}
}

// Execute executes the request
//
//	@return map[string]interface{}
func (a *PluginsAPIService) ApiV1PluginsNpmNamePublicSettingsGetExecute(r ApiApiV1PluginsNpmNamePublicSettingsGetRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PluginsAPIService.ApiV1PluginsNpmNamePublicSettingsGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/plugins/{npmName}/public-settings"
	localVarPath = strings.Replace(localVarPath, "{"+"npmName"+"}", url.PathEscape(parameterValueToString(r.npmName, "npmName")), -1)

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

type ApiApiV1PluginsNpmNameRegisteredSettingsGetRequest struct {
	ctx        context.Context
	ApiService *PluginsAPIService
	npmName    string
}

func (r ApiApiV1PluginsNpmNameRegisteredSettingsGetRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.ApiV1PluginsNpmNameRegisteredSettingsGetExecute(r)
}

/*
ApiV1PluginsNpmNameRegisteredSettingsGet Get a plugin's registered settings

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param npmName name of the plugin/theme on npmjs.com or in its package.json
	@return ApiApiV1PluginsNpmNameRegisteredSettingsGetRequest
*/
func (a *PluginsAPIService) ApiV1PluginsNpmNameRegisteredSettingsGet(ctx context.Context, npmName string) ApiApiV1PluginsNpmNameRegisteredSettingsGetRequest {
	return ApiApiV1PluginsNpmNameRegisteredSettingsGetRequest{
		ApiService: a,
		ctx:        ctx,
		npmName:    npmName,
	}
}

// Execute executes the request
//
//	@return map[string]interface{}
func (a *PluginsAPIService) ApiV1PluginsNpmNameRegisteredSettingsGetExecute(r ApiApiV1PluginsNpmNameRegisteredSettingsGetRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PluginsAPIService.ApiV1PluginsNpmNameRegisteredSettingsGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/plugins/{npmName}/registered-settings"
	localVarPath = strings.Replace(localVarPath, "{"+"npmName"+"}", url.PathEscape(parameterValueToString(r.npmName, "npmName")), -1)

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

type ApiApiV1PluginsNpmNameSettingsPutRequest struct {
	ctx                                   context.Context
	ApiService                            *PluginsAPIService
	npmName                               string
	apiV1PluginsNpmNameSettingsPutRequest *models.ApiV1PluginsNpmNameSettingsPutRequest
}

func (r ApiApiV1PluginsNpmNameSettingsPutRequest) ApiV1PluginsNpmNameSettingsPutRequest(apiV1PluginsNpmNameSettingsPutRequest models.ApiV1PluginsNpmNameSettingsPutRequest) ApiApiV1PluginsNpmNameSettingsPutRequest {
	r.apiV1PluginsNpmNameSettingsPutRequest = &apiV1PluginsNpmNameSettingsPutRequest
	return r
}

func (r ApiApiV1PluginsNpmNameSettingsPutRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiV1PluginsNpmNameSettingsPutExecute(r)
}

/*
ApiV1PluginsNpmNameSettingsPut Set a plugin's settings

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param npmName name of the plugin/theme on npmjs.com or in its package.json
	@return ApiApiV1PluginsNpmNameSettingsPutRequest
*/
func (a *PluginsAPIService) ApiV1PluginsNpmNameSettingsPut(ctx context.Context, npmName string) ApiApiV1PluginsNpmNameSettingsPutRequest {
	return ApiApiV1PluginsNpmNameSettingsPutRequest{
		ApiService: a,
		ctx:        ctx,
		npmName:    npmName,
	}
}

// Execute executes the request
func (a *PluginsAPIService) ApiV1PluginsNpmNameSettingsPutExecute(r ApiApiV1PluginsNpmNameSettingsPutRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPut
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PluginsAPIService.ApiV1PluginsNpmNameSettingsPut")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/plugins/{npmName}/settings"
	localVarPath = strings.Replace(localVarPath, "{"+"npmName"+"}", url.PathEscape(parameterValueToString(r.npmName, "npmName")), -1)

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
	localVarPostBody = r.apiV1PluginsNpmNameSettingsPutRequest
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

type ApiGetAvailablePluginsRequest struct {
	ctx                   context.Context
	ApiService            *PluginsAPIService
	search                *string
	pluginType            *int32
	currentPeerTubeEngine *string
	start                 *int32
	count                 *int32
	sort                  *string
}

func (r ApiGetAvailablePluginsRequest) Search(search string) ApiGetAvailablePluginsRequest {
	r.search = &search
	return r
}

func (r ApiGetAvailablePluginsRequest) PluginType(pluginType int32) ApiGetAvailablePluginsRequest {
	r.pluginType = &pluginType
	return r
}

func (r ApiGetAvailablePluginsRequest) CurrentPeerTubeEngine(currentPeerTubeEngine string) ApiGetAvailablePluginsRequest {
	r.currentPeerTubeEngine = &currentPeerTubeEngine
	return r
}

// Offset used to paginate results
func (r ApiGetAvailablePluginsRequest) Start(start int32) ApiGetAvailablePluginsRequest {
	r.start = &start
	return r
}

// Number of items to return
func (r ApiGetAvailablePluginsRequest) Count(count int32) ApiGetAvailablePluginsRequest {
	r.count = &count
	return r
}

// Sort column
func (r ApiGetAvailablePluginsRequest) Sort(sort string) ApiGetAvailablePluginsRequest {
	r.sort = &sort
	return r
}

func (r ApiGetAvailablePluginsRequest) Execute() (*models.PluginResponse, *http.Response, error) {
	return r.ApiService.GetAvailablePluginsExecute(r)
}

/*
GetAvailablePlugins List available plugins

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetAvailablePluginsRequest
*/
func (a *PluginsAPIService) GetAvailablePlugins(ctx context.Context) ApiGetAvailablePluginsRequest {
	return ApiGetAvailablePluginsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return PluginResponse
func (a *PluginsAPIService) GetAvailablePluginsExecute(r ApiGetAvailablePluginsRequest) (*models.PluginResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.PluginResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PluginsAPIService.GetAvailablePlugins")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/plugins/available"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.search != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search", r.search, "form", "")
	}
	if r.pluginType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pluginType", r.pluginType, "form", "")
	}
	if r.currentPeerTubeEngine != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "currentPeerTubeEngine", r.currentPeerTubeEngine, "form", "")
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

type ApiGetPluginRequest struct {
	ctx        context.Context
	ApiService *PluginsAPIService
	npmName    string
}

func (r ApiGetPluginRequest) Execute() (*models.Plugin, *http.Response, error) {
	return r.ApiService.GetPluginExecute(r)
}

/*
GetPlugin Get a plugin

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param npmName name of the plugin/theme on npmjs.com or in its package.json
	@return ApiGetPluginRequest
*/
func (a *PluginsAPIService) GetPlugin(ctx context.Context, npmName string) ApiGetPluginRequest {
	return ApiGetPluginRequest{
		ApiService: a,
		ctx:        ctx,
		npmName:    npmName,
	}
}

// Execute executes the request
//
//	@return Plugin
func (a *PluginsAPIService) GetPluginExecute(r ApiGetPluginRequest) (*models.Plugin, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.Plugin
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PluginsAPIService.GetPlugin")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/plugins/{npmName}"
	localVarPath = strings.Replace(localVarPath, "{"+"npmName"+"}", url.PathEscape(parameterValueToString(r.npmName, "npmName")), -1)

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

type ApiGetPluginsRequest struct {
	ctx         context.Context
	ApiService  *PluginsAPIService
	pluginType  *int32
	uninstalled *bool
	start       *int32
	count       *int32
	sort        *string
}

func (r ApiGetPluginsRequest) PluginType(pluginType int32) ApiGetPluginsRequest {
	r.pluginType = &pluginType
	return r
}

func (r ApiGetPluginsRequest) Uninstalled(uninstalled bool) ApiGetPluginsRequest {
	r.uninstalled = &uninstalled
	return r
}

// Offset used to paginate results
func (r ApiGetPluginsRequest) Start(start int32) ApiGetPluginsRequest {
	r.start = &start
	return r
}

// Number of items to return
func (r ApiGetPluginsRequest) Count(count int32) ApiGetPluginsRequest {
	r.count = &count
	return r
}

// Sort column
func (r ApiGetPluginsRequest) Sort(sort string) ApiGetPluginsRequest {
	r.sort = &sort
	return r
}

func (r ApiGetPluginsRequest) Execute() (*models.PluginResponse, *http.Response, error) {
	return r.ApiService.GetPluginsExecute(r)
}

/*
GetPlugins List plugins

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetPluginsRequest
*/
func (a *PluginsAPIService) GetPlugins(ctx context.Context) ApiGetPluginsRequest {
	return ApiGetPluginsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return PluginResponse
func (a *PluginsAPIService) GetPluginsExecute(r ApiGetPluginsRequest) (*models.PluginResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.PluginResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PluginsAPIService.GetPlugins")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/plugins"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.pluginType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pluginType", r.pluginType, "form", "")
	}
	if r.uninstalled != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "uninstalled", r.uninstalled, "form", "")
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

type ApiUninstallPluginRequest struct {
	ctx                    context.Context
	ApiService             *PluginsAPIService
	uninstallPluginRequest *models.UninstallPluginRequest
}

func (r ApiUninstallPluginRequest) UninstallPluginRequest(uninstallPluginRequest models.UninstallPluginRequest) ApiUninstallPluginRequest {
	r.uninstallPluginRequest = &uninstallPluginRequest
	return r
}

func (r ApiUninstallPluginRequest) Execute() (*http.Response, error) {
	return r.ApiService.UninstallPluginExecute(r)
}

/*
UninstallPlugin Uninstall a plugin

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiUninstallPluginRequest
*/
func (a *PluginsAPIService) UninstallPlugin(ctx context.Context) ApiUninstallPluginRequest {
	return ApiUninstallPluginRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *PluginsAPIService) UninstallPluginExecute(r ApiUninstallPluginRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PluginsAPIService.UninstallPlugin")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/plugins/uninstall"

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
	localVarPostBody = r.uninstallPluginRequest
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

type ApiUpdatePluginRequest struct {
	ctx              context.Context
	ApiService       *PluginsAPIService
	addPluginRequest *models.AddPluginRequest
}

func (r ApiUpdatePluginRequest) AddPluginRequest(addPluginRequest models.AddPluginRequest) ApiUpdatePluginRequest {
	r.addPluginRequest = &addPluginRequest
	return r
}

func (r ApiUpdatePluginRequest) Execute() (*http.Response, error) {
	return r.ApiService.UpdatePluginExecute(r)
}

/*
UpdatePlugin Update a plugin

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiUpdatePluginRequest
*/
func (a *PluginsAPIService) UpdatePlugin(ctx context.Context) ApiUpdatePluginRequest {
	return ApiUpdatePluginRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *PluginsAPIService) UpdatePluginExecute(r ApiUpdatePluginRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PluginsAPIService.UpdatePlugin")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/plugins/update"

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
	localVarPostBody = r.addPluginRequest
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
