# \PluginsAPI

All URIs are relative to *https://peertube2.cpy.re*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddPlugin**](PluginsAPI.md#AddPlugin) | **Post** /api/v1/plugins/install | Install a plugin
[**ApiV1PluginsNpmNamePublicSettingsGet**](PluginsAPI.md#ApiV1PluginsNpmNamePublicSettingsGet) | **Get** /api/v1/plugins/{npmName}/public-settings | Get a plugin&#39;s public settings
[**ApiV1PluginsNpmNameRegisteredSettingsGet**](PluginsAPI.md#ApiV1PluginsNpmNameRegisteredSettingsGet) | **Get** /api/v1/plugins/{npmName}/registered-settings | Get a plugin&#39;s registered settings
[**ApiV1PluginsNpmNameSettingsPut**](PluginsAPI.md#ApiV1PluginsNpmNameSettingsPut) | **Put** /api/v1/plugins/{npmName}/settings | Set a plugin&#39;s settings
[**GetAvailablePlugins**](PluginsAPI.md#GetAvailablePlugins) | **Get** /api/v1/plugins/available | List available plugins
[**GetPlugin**](PluginsAPI.md#GetPlugin) | **Get** /api/v1/plugins/{npmName} | Get a plugin
[**GetPlugins**](PluginsAPI.md#GetPlugins) | **Get** /api/v1/plugins | List plugins
[**UninstallPlugin**](PluginsAPI.md#UninstallPlugin) | **Post** /api/v1/plugins/uninstall | Uninstall a plugin
[**UpdatePlugin**](PluginsAPI.md#UpdatePlugin) | **Post** /api/v1/plugins/update | Update a plugin



## AddPlugin

> AddPlugin(ctx).AddPluginRequest(addPluginRequest).Execute()

Install a plugin

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/RustLangLatam/peertube_api_sdk"
)

func main() {
	addPluginRequest := openapiclient.addPlugin_request{AddPluginRequestOneOf: openapiclient.NewAddPluginRequestOneOf("peertube-plugin-auth-ldap")} // AddPluginRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PluginsAPI.AddPlugin(context.Background()).AddPluginRequest(addPluginRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PluginsAPI.AddPlugin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddPluginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addPluginRequest** | [**AddPluginRequest**](AddPluginRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1PluginsNpmNamePublicSettingsGet

> map[string]interface{} ApiV1PluginsNpmNamePublicSettingsGet(ctx, npmName).Execute()

Get a plugin's public settings

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/RustLangLatam/peertube_api_sdk"
)

func main() {
	npmName := "peertube-plugin-auth-ldap" // string | name of the plugin/theme on npmjs.com or in its package.json

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PluginsAPI.ApiV1PluginsNpmNamePublicSettingsGet(context.Background(), npmName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PluginsAPI.ApiV1PluginsNpmNamePublicSettingsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1PluginsNpmNamePublicSettingsGet`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `PluginsAPI.ApiV1PluginsNpmNamePublicSettingsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**npmName** | **string** | name of the plugin/theme on npmjs.com or in its package.json | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1PluginsNpmNamePublicSettingsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1PluginsNpmNameRegisteredSettingsGet

> map[string]interface{} ApiV1PluginsNpmNameRegisteredSettingsGet(ctx, npmName).Execute()

Get a plugin's registered settings

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/RustLangLatam/peertube_api_sdk"
)

func main() {
	npmName := "peertube-plugin-auth-ldap" // string | name of the plugin/theme on npmjs.com or in its package.json

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PluginsAPI.ApiV1PluginsNpmNameRegisteredSettingsGet(context.Background(), npmName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PluginsAPI.ApiV1PluginsNpmNameRegisteredSettingsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1PluginsNpmNameRegisteredSettingsGet`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `PluginsAPI.ApiV1PluginsNpmNameRegisteredSettingsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**npmName** | **string** | name of the plugin/theme on npmjs.com or in its package.json | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1PluginsNpmNameRegisteredSettingsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1PluginsNpmNameSettingsPut

> ApiV1PluginsNpmNameSettingsPut(ctx, npmName).ApiV1PluginsNpmNameSettingsPutRequest(apiV1PluginsNpmNameSettingsPutRequest).Execute()

Set a plugin's settings

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/RustLangLatam/peertube_api_sdk"
)

func main() {
	npmName := "peertube-plugin-auth-ldap" // string | name of the plugin/theme on npmjs.com or in its package.json
	apiV1PluginsNpmNameSettingsPutRequest := *openapiclient.NewApiV1PluginsNpmNameSettingsPutRequest() // ApiV1PluginsNpmNameSettingsPutRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PluginsAPI.ApiV1PluginsNpmNameSettingsPut(context.Background(), npmName).ApiV1PluginsNpmNameSettingsPutRequest(apiV1PluginsNpmNameSettingsPutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PluginsAPI.ApiV1PluginsNpmNameSettingsPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**npmName** | **string** | name of the plugin/theme on npmjs.com or in its package.json | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1PluginsNpmNameSettingsPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiV1PluginsNpmNameSettingsPutRequest** | [**ApiV1PluginsNpmNameSettingsPutRequest**](ApiV1PluginsNpmNameSettingsPutRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAvailablePlugins

> PluginResponse GetAvailablePlugins(ctx).Search(search).PluginType(pluginType).CurrentPeerTubeEngine(currentPeerTubeEngine).Start(start).Count(count).Sort(sort).Execute()

List available plugins

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/RustLangLatam/peertube_api_sdk"
)

func main() {
	search := "search_example" // string |  (optional)
	pluginType := int32(56) // int32 |  (optional)
	currentPeerTubeEngine := "currentPeerTubeEngine_example" // string |  (optional)
	start := int32(56) // int32 | Offset used to paginate results (optional)
	count := int32(56) // int32 | Number of items to return (optional) (default to 15)
	sort := "-createdAt" // string | Sort column (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PluginsAPI.GetAvailablePlugins(context.Background()).Search(search).PluginType(pluginType).CurrentPeerTubeEngine(currentPeerTubeEngine).Start(start).Count(count).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PluginsAPI.GetAvailablePlugins``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAvailablePlugins`: PluginResponse
	fmt.Fprintf(os.Stdout, "Response from `PluginsAPI.GetAvailablePlugins`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAvailablePluginsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **string** |  | 
 **pluginType** | **int32** |  | 
 **currentPeerTubeEngine** | **string** |  | 
 **start** | **int32** | Offset used to paginate results | 
 **count** | **int32** | Number of items to return | [default to 15]
 **sort** | **string** | Sort column | 

### Return type

[**PluginResponse**](PluginResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlugin

> Plugin GetPlugin(ctx, npmName).Execute()

Get a plugin

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/RustLangLatam/peertube_api_sdk"
)

func main() {
	npmName := "peertube-plugin-auth-ldap" // string | name of the plugin/theme on npmjs.com or in its package.json

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PluginsAPI.GetPlugin(context.Background(), npmName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PluginsAPI.GetPlugin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlugin`: Plugin
	fmt.Fprintf(os.Stdout, "Response from `PluginsAPI.GetPlugin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**npmName** | **string** | name of the plugin/theme on npmjs.com or in its package.json | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPluginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Plugin**](Plugin.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlugins

> PluginResponse GetPlugins(ctx).PluginType(pluginType).Uninstalled(uninstalled).Start(start).Count(count).Sort(sort).Execute()

List plugins

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/RustLangLatam/peertube_api_sdk"
)

func main() {
	pluginType := int32(56) // int32 |  (optional)
	uninstalled := true // bool |  (optional)
	start := int32(56) // int32 | Offset used to paginate results (optional)
	count := int32(56) // int32 | Number of items to return (optional) (default to 15)
	sort := "-createdAt" // string | Sort column (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PluginsAPI.GetPlugins(context.Background()).PluginType(pluginType).Uninstalled(uninstalled).Start(start).Count(count).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PluginsAPI.GetPlugins``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlugins`: PluginResponse
	fmt.Fprintf(os.Stdout, "Response from `PluginsAPI.GetPlugins`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPluginsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pluginType** | **int32** |  | 
 **uninstalled** | **bool** |  | 
 **start** | **int32** | Offset used to paginate results | 
 **count** | **int32** | Number of items to return | [default to 15]
 **sort** | **string** | Sort column | 

### Return type

[**PluginResponse**](PluginResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UninstallPlugin

> UninstallPlugin(ctx).UninstallPluginRequest(uninstallPluginRequest).Execute()

Uninstall a plugin

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/RustLangLatam/peertube_api_sdk"
)

func main() {
	uninstallPluginRequest := *openapiclient.NewUninstallPluginRequest("peertube-plugin-auth-ldap") // UninstallPluginRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PluginsAPI.UninstallPlugin(context.Background()).UninstallPluginRequest(uninstallPluginRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PluginsAPI.UninstallPlugin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUninstallPluginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uninstallPluginRequest** | [**UninstallPluginRequest**](UninstallPluginRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePlugin

> UpdatePlugin(ctx).AddPluginRequest(addPluginRequest).Execute()

Update a plugin

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/RustLangLatam/peertube_api_sdk"
)

func main() {
	addPluginRequest := openapiclient.addPlugin_request{AddPluginRequestOneOf: openapiclient.NewAddPluginRequestOneOf("peertube-plugin-auth-ldap")} // AddPluginRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PluginsAPI.UpdatePlugin(context.Background()).AddPluginRequest(addPluginRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PluginsAPI.UpdatePlugin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePluginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addPluginRequest** | [**AddPluginRequest**](AddPluginRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

