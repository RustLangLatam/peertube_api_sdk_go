# ServerConfigTranscoding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hls** | Pointer to [**ServerConfigEmail**](ServerConfigEmail.md) |  | [optional] 
**WebVideos** | Pointer to [**ServerConfigEmail**](ServerConfigEmail.md) |  | [optional] 
**EnabledResolutions** | Pointer to **[]int32** |  | [optional] 

## Methods

### NewServerConfigTranscoding

`func NewServerConfigTranscoding() *ServerConfigTranscoding`

NewServerConfigTranscoding instantiates a new ServerConfigTranscoding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerConfigTranscodingWithDefaults

`func NewServerConfigTranscodingWithDefaults() *ServerConfigTranscoding`

NewServerConfigTranscodingWithDefaults instantiates a new ServerConfigTranscoding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHls

`func (o *ServerConfigTranscoding) GetHls() ServerConfigEmail`

GetHls returns the Hls field if non-nil, zero value otherwise.

### GetHlsOk

`func (o *ServerConfigTranscoding) GetHlsOk() (*ServerConfigEmail, bool)`

GetHlsOk returns a tuple with the Hls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHls

`func (o *ServerConfigTranscoding) SetHls(v ServerConfigEmail)`

SetHls sets Hls field to given value.

### HasHls

`func (o *ServerConfigTranscoding) HasHls() bool`

HasHls returns a boolean if a field has been set.

### GetWebVideos

`func (o *ServerConfigTranscoding) GetWebVideos() ServerConfigEmail`

GetWebVideos returns the WebVideos field if non-nil, zero value otherwise.

### GetWebVideosOk

`func (o *ServerConfigTranscoding) GetWebVideosOk() (*ServerConfigEmail, bool)`

GetWebVideosOk returns a tuple with the WebVideos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebVideos

`func (o *ServerConfigTranscoding) SetWebVideos(v ServerConfigEmail)`

SetWebVideos sets WebVideos field to given value.

### HasWebVideos

`func (o *ServerConfigTranscoding) HasWebVideos() bool`

HasWebVideos returns a boolean if a field has been set.

### GetEnabledResolutions

`func (o *ServerConfigTranscoding) GetEnabledResolutions() []int32`

GetEnabledResolutions returns the EnabledResolutions field if non-nil, zero value otherwise.

### GetEnabledResolutionsOk

`func (o *ServerConfigTranscoding) GetEnabledResolutionsOk() (*[]int32, bool)`

GetEnabledResolutionsOk returns a tuple with the EnabledResolutions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledResolutions

`func (o *ServerConfigTranscoding) SetEnabledResolutions(v []int32)`

SetEnabledResolutions sets EnabledResolutions field to given value.

### HasEnabledResolutions

`func (o *ServerConfigTranscoding) HasEnabledResolutions() bool`

HasEnabledResolutions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


