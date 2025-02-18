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

// WatchedWordsAPIService WatchedWordsAPI service
type WatchedWordsAPIService service

type ApiApiV1WatchedWordsAccountsAccountNameListsGetRequest struct {
	ctx         context.Context
	ApiService  *WatchedWordsAPIService
	accountName string
}

func (r ApiApiV1WatchedWordsAccountsAccountNameListsGetRequest) Execute() (*models.ApiV1WatchedWordsAccountsAccountNameListsGet200Response, *http.Response, error) {
	return r.ApiService.ApiV1WatchedWordsAccountsAccountNameListsGetExecute(r)
}

/*
ApiV1WatchedWordsAccountsAccountNameListsGet List account watched words

**PeerTube >= 6.2**

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountName account name to list watched words
	@return ApiApiV1WatchedWordsAccountsAccountNameListsGetRequest
*/
func (a *WatchedWordsAPIService) ApiV1WatchedWordsAccountsAccountNameListsGet(ctx context.Context, accountName string) ApiApiV1WatchedWordsAccountsAccountNameListsGetRequest {
	return ApiApiV1WatchedWordsAccountsAccountNameListsGetRequest{
		ApiService:  a,
		ctx:         ctx,
		accountName: accountName,
	}
}

// Execute executes the request
//
//	@return ApiV1WatchedWordsAccountsAccountNameListsGet200Response
func (a *WatchedWordsAPIService) ApiV1WatchedWordsAccountsAccountNameListsGetExecute(r ApiApiV1WatchedWordsAccountsAccountNameListsGetRequest) (*models.ApiV1WatchedWordsAccountsAccountNameListsGet200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.ApiV1WatchedWordsAccountsAccountNameListsGet200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WatchedWordsAPIService.ApiV1WatchedWordsAccountsAccountNameListsGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/watched-words/accounts/{accountName}/lists"
	localVarPath = strings.Replace(localVarPath, "{"+"accountName"+"}", url.PathEscape(parameterValueToString(r.accountName, "accountName")), -1)

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

type ApiApiV1WatchedWordsAccountsAccountNameListsListIdDeleteRequest struct {
	ctx         context.Context
	ApiService  *WatchedWordsAPIService
	accountName string
	listId      string
}

func (r ApiApiV1WatchedWordsAccountsAccountNameListsListIdDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiV1WatchedWordsAccountsAccountNameListsListIdDeleteExecute(r)
}

/*
ApiV1WatchedWordsAccountsAccountNameListsListIdDelete Delete account watched words

**PeerTube >= 6.2**

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountName
	@param listId list of watched words to delete
	@return ApiApiV1WatchedWordsAccountsAccountNameListsListIdDeleteRequest
*/
func (a *WatchedWordsAPIService) ApiV1WatchedWordsAccountsAccountNameListsListIdDelete(ctx context.Context, accountName string, listId string) ApiApiV1WatchedWordsAccountsAccountNameListsListIdDeleteRequest {
	return ApiApiV1WatchedWordsAccountsAccountNameListsListIdDeleteRequest{
		ApiService:  a,
		ctx:         ctx,
		accountName: accountName,
		listId:      listId,
	}
}

// Execute executes the request
func (a *WatchedWordsAPIService) ApiV1WatchedWordsAccountsAccountNameListsListIdDeleteExecute(r ApiApiV1WatchedWordsAccountsAccountNameListsListIdDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WatchedWordsAPIService.ApiV1WatchedWordsAccountsAccountNameListsListIdDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/watched-words/accounts/{accountName}/lists/{listId}"
	localVarPath = strings.Replace(localVarPath, "{"+"accountName"+"}", url.PathEscape(parameterValueToString(r.accountName, "accountName")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"listId"+"}", url.PathEscape(parameterValueToString(r.listId, "listId")), -1)

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

type ApiApiV1WatchedWordsAccountsAccountNameListsListIdPutRequest struct {
	ctx                                                  context.Context
	ApiService                                           *WatchedWordsAPIService
	accountName                                          string
	listId                                               string
	apiV1WatchedWordsAccountsAccountNameListsPostRequest *models.ApiV1WatchedWordsAccountsAccountNameListsPostRequest
}

func (r ApiApiV1WatchedWordsAccountsAccountNameListsListIdPutRequest) ApiV1WatchedWordsAccountsAccountNameListsPostRequest(apiV1WatchedWordsAccountsAccountNameListsPostRequest models.ApiV1WatchedWordsAccountsAccountNameListsPostRequest) ApiApiV1WatchedWordsAccountsAccountNameListsListIdPutRequest {
	r.apiV1WatchedWordsAccountsAccountNameListsPostRequest = &apiV1WatchedWordsAccountsAccountNameListsPostRequest
	return r
}

func (r ApiApiV1WatchedWordsAccountsAccountNameListsListIdPutRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiV1WatchedWordsAccountsAccountNameListsListIdPutExecute(r)
}

/*
ApiV1WatchedWordsAccountsAccountNameListsListIdPut Update account watched words

**PeerTube >= 6.2**

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountName
	@param listId list of watched words to update
	@return ApiApiV1WatchedWordsAccountsAccountNameListsListIdPutRequest
*/
func (a *WatchedWordsAPIService) ApiV1WatchedWordsAccountsAccountNameListsListIdPut(ctx context.Context, accountName string, listId string) ApiApiV1WatchedWordsAccountsAccountNameListsListIdPutRequest {
	return ApiApiV1WatchedWordsAccountsAccountNameListsListIdPutRequest{
		ApiService:  a,
		ctx:         ctx,
		accountName: accountName,
		listId:      listId,
	}
}

// Execute executes the request
func (a *WatchedWordsAPIService) ApiV1WatchedWordsAccountsAccountNameListsListIdPutExecute(r ApiApiV1WatchedWordsAccountsAccountNameListsListIdPutRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPut
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WatchedWordsAPIService.ApiV1WatchedWordsAccountsAccountNameListsListIdPut")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/watched-words/accounts/{accountName}/lists/{listId}"
	localVarPath = strings.Replace(localVarPath, "{"+"accountName"+"}", url.PathEscape(parameterValueToString(r.accountName, "accountName")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"listId"+"}", url.PathEscape(parameterValueToString(r.listId, "listId")), -1)

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
	localVarPostBody = r.apiV1WatchedWordsAccountsAccountNameListsPostRequest
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

type ApiApiV1WatchedWordsAccountsAccountNameListsPostRequest struct {
	ctx                                                  context.Context
	ApiService                                           *WatchedWordsAPIService
	accountName                                          string
	apiV1WatchedWordsAccountsAccountNameListsPostRequest *models.ApiV1WatchedWordsAccountsAccountNameListsPostRequest
}

func (r ApiApiV1WatchedWordsAccountsAccountNameListsPostRequest) ApiV1WatchedWordsAccountsAccountNameListsPostRequest(apiV1WatchedWordsAccountsAccountNameListsPostRequest models.ApiV1WatchedWordsAccountsAccountNameListsPostRequest) ApiApiV1WatchedWordsAccountsAccountNameListsPostRequest {
	r.apiV1WatchedWordsAccountsAccountNameListsPostRequest = &apiV1WatchedWordsAccountsAccountNameListsPostRequest
	return r
}

func (r ApiApiV1WatchedWordsAccountsAccountNameListsPostRequest) Execute() (*models.ApiV1WatchedWordsAccountsAccountNameListsPost200Response, *http.Response, error) {
	return r.ApiService.ApiV1WatchedWordsAccountsAccountNameListsPostExecute(r)
}

/*
ApiV1WatchedWordsAccountsAccountNameListsPost Add account watched words

**PeerTube >= 6.2**

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param accountName
	@return ApiApiV1WatchedWordsAccountsAccountNameListsPostRequest
*/
func (a *WatchedWordsAPIService) ApiV1WatchedWordsAccountsAccountNameListsPost(ctx context.Context, accountName string) ApiApiV1WatchedWordsAccountsAccountNameListsPostRequest {
	return ApiApiV1WatchedWordsAccountsAccountNameListsPostRequest{
		ApiService:  a,
		ctx:         ctx,
		accountName: accountName,
	}
}

// Execute executes the request
//
//	@return ApiV1WatchedWordsAccountsAccountNameListsPost200Response
func (a *WatchedWordsAPIService) ApiV1WatchedWordsAccountsAccountNameListsPostExecute(r ApiApiV1WatchedWordsAccountsAccountNameListsPostRequest) (*models.ApiV1WatchedWordsAccountsAccountNameListsPost200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.ApiV1WatchedWordsAccountsAccountNameListsPost200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WatchedWordsAPIService.ApiV1WatchedWordsAccountsAccountNameListsPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/watched-words/accounts/{accountName}/lists"
	localVarPath = strings.Replace(localVarPath, "{"+"accountName"+"}", url.PathEscape(parameterValueToString(r.accountName, "accountName")), -1)

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
	localVarPostBody = r.apiV1WatchedWordsAccountsAccountNameListsPostRequest
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

type ApiApiV1WatchedWordsServerListsGetRequest struct {
	ctx        context.Context
	ApiService *WatchedWordsAPIService
}

func (r ApiApiV1WatchedWordsServerListsGetRequest) Execute() (*models.ApiV1WatchedWordsAccountsAccountNameListsGet200Response, *http.Response, error) {
	return r.ApiService.ApiV1WatchedWordsServerListsGetExecute(r)
}

/*
ApiV1WatchedWordsServerListsGet List server watched words

**PeerTube >= 6.2**

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiApiV1WatchedWordsServerListsGetRequest
*/
func (a *WatchedWordsAPIService) ApiV1WatchedWordsServerListsGet(ctx context.Context) ApiApiV1WatchedWordsServerListsGetRequest {
	return ApiApiV1WatchedWordsServerListsGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ApiV1WatchedWordsAccountsAccountNameListsGet200Response
func (a *WatchedWordsAPIService) ApiV1WatchedWordsServerListsGetExecute(r ApiApiV1WatchedWordsServerListsGetRequest) (*models.ApiV1WatchedWordsAccountsAccountNameListsGet200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.ApiV1WatchedWordsAccountsAccountNameListsGet200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WatchedWordsAPIService.ApiV1WatchedWordsServerListsGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/watched-words/server/lists"

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

type ApiApiV1WatchedWordsServerListsListIdDeleteRequest struct {
	ctx        context.Context
	ApiService *WatchedWordsAPIService
	listId     string
}

func (r ApiApiV1WatchedWordsServerListsListIdDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiV1WatchedWordsServerListsListIdDeleteExecute(r)
}

/*
ApiV1WatchedWordsServerListsListIdDelete Delete server watched words

**PeerTube >= 6.2**

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param listId list of watched words to delete
	@return ApiApiV1WatchedWordsServerListsListIdDeleteRequest
*/
func (a *WatchedWordsAPIService) ApiV1WatchedWordsServerListsListIdDelete(ctx context.Context, listId string) ApiApiV1WatchedWordsServerListsListIdDeleteRequest {
	return ApiApiV1WatchedWordsServerListsListIdDeleteRequest{
		ApiService: a,
		ctx:        ctx,
		listId:     listId,
	}
}

// Execute executes the request
func (a *WatchedWordsAPIService) ApiV1WatchedWordsServerListsListIdDeleteExecute(r ApiApiV1WatchedWordsServerListsListIdDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WatchedWordsAPIService.ApiV1WatchedWordsServerListsListIdDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/watched-words/server/lists/{listId}"
	localVarPath = strings.Replace(localVarPath, "{"+"listId"+"}", url.PathEscape(parameterValueToString(r.listId, "listId")), -1)

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

type ApiApiV1WatchedWordsServerListsListIdPutRequest struct {
	ctx                                                  context.Context
	ApiService                                           *WatchedWordsAPIService
	listId                                               string
	apiV1WatchedWordsAccountsAccountNameListsPostRequest *models.ApiV1WatchedWordsAccountsAccountNameListsPostRequest
}

func (r ApiApiV1WatchedWordsServerListsListIdPutRequest) ApiV1WatchedWordsAccountsAccountNameListsPostRequest(apiV1WatchedWordsAccountsAccountNameListsPostRequest models.ApiV1WatchedWordsAccountsAccountNameListsPostRequest) ApiApiV1WatchedWordsServerListsListIdPutRequest {
	r.apiV1WatchedWordsAccountsAccountNameListsPostRequest = &apiV1WatchedWordsAccountsAccountNameListsPostRequest
	return r
}

func (r ApiApiV1WatchedWordsServerListsListIdPutRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiV1WatchedWordsServerListsListIdPutExecute(r)
}

/*
ApiV1WatchedWordsServerListsListIdPut Update server watched words

**PeerTube >= 6.2**

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param listId list of watched words to update
	@return ApiApiV1WatchedWordsServerListsListIdPutRequest
*/
func (a *WatchedWordsAPIService) ApiV1WatchedWordsServerListsListIdPut(ctx context.Context, listId string) ApiApiV1WatchedWordsServerListsListIdPutRequest {
	return ApiApiV1WatchedWordsServerListsListIdPutRequest{
		ApiService: a,
		ctx:        ctx,
		listId:     listId,
	}
}

// Execute executes the request
func (a *WatchedWordsAPIService) ApiV1WatchedWordsServerListsListIdPutExecute(r ApiApiV1WatchedWordsServerListsListIdPutRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPut
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WatchedWordsAPIService.ApiV1WatchedWordsServerListsListIdPut")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/watched-words/server/lists/{listId}"
	localVarPath = strings.Replace(localVarPath, "{"+"listId"+"}", url.PathEscape(parameterValueToString(r.listId, "listId")), -1)

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
	localVarPostBody = r.apiV1WatchedWordsAccountsAccountNameListsPostRequest
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

type ApiApiV1WatchedWordsServerListsPostRequest struct {
	ctx                                                  context.Context
	ApiService                                           *WatchedWordsAPIService
	apiV1WatchedWordsAccountsAccountNameListsPostRequest *models.ApiV1WatchedWordsAccountsAccountNameListsPostRequest
}

func (r ApiApiV1WatchedWordsServerListsPostRequest) ApiV1WatchedWordsAccountsAccountNameListsPostRequest(apiV1WatchedWordsAccountsAccountNameListsPostRequest models.ApiV1WatchedWordsAccountsAccountNameListsPostRequest) ApiApiV1WatchedWordsServerListsPostRequest {
	r.apiV1WatchedWordsAccountsAccountNameListsPostRequest = &apiV1WatchedWordsAccountsAccountNameListsPostRequest
	return r
}

func (r ApiApiV1WatchedWordsServerListsPostRequest) Execute() (*models.ApiV1WatchedWordsAccountsAccountNameListsPost200Response, *http.Response, error) {
	return r.ApiService.ApiV1WatchedWordsServerListsPostExecute(r)
}

/*
ApiV1WatchedWordsServerListsPost Add server watched words

**PeerTube >= 6.2**

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiApiV1WatchedWordsServerListsPostRequest
*/
func (a *WatchedWordsAPIService) ApiV1WatchedWordsServerListsPost(ctx context.Context) ApiApiV1WatchedWordsServerListsPostRequest {
	return ApiApiV1WatchedWordsServerListsPostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ApiV1WatchedWordsAccountsAccountNameListsPost200Response
func (a *WatchedWordsAPIService) ApiV1WatchedWordsServerListsPostExecute(r ApiApiV1WatchedWordsServerListsPostRequest) (*models.ApiV1WatchedWordsAccountsAccountNameListsPost200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *models.ApiV1WatchedWordsAccountsAccountNameListsPost200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WatchedWordsAPIService.ApiV1WatchedWordsServerListsPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/watched-words/server/lists"

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
	localVarPostBody = r.apiV1WatchedWordsAccountsAccountNameListsPostRequest
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
