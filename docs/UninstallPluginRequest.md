# UninstallPluginRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NpmName** | **string** | name of the plugin/theme in its package.json | 

## Methods

### NewUninstallPluginRequest

`func NewUninstallPluginRequest(npmName string, ) *UninstallPluginRequest`

NewUninstallPluginRequest instantiates a new UninstallPluginRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUninstallPluginRequestWithDefaults

`func NewUninstallPluginRequestWithDefaults() *UninstallPluginRequest`

NewUninstallPluginRequestWithDefaults instantiates a new UninstallPluginRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNpmName

`func (o *UninstallPluginRequest) GetNpmName() string`

GetNpmName returns the NpmName field if non-nil, zero value otherwise.

### GetNpmNameOk

`func (o *UninstallPluginRequest) GetNpmNameOk() (*string, bool)`

GetNpmNameOk returns a tuple with the NpmName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNpmName

`func (o *UninstallPluginRequest) SetNpmName(v string)`

SetNpmName sets NpmName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


