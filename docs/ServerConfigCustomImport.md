# ServerConfigCustomImport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Videos** | Pointer to [**ServerConfigImportVideos**](ServerConfigImportVideos.md) |  | [optional] 
**VideoChannelSynchronization** | Pointer to [**ServerConfigEmail**](ServerConfigEmail.md) |  | [optional] 

## Methods

### NewServerConfigCustomImport

`func NewServerConfigCustomImport() *ServerConfigCustomImport`

NewServerConfigCustomImport instantiates a new ServerConfigCustomImport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerConfigCustomImportWithDefaults

`func NewServerConfigCustomImportWithDefaults() *ServerConfigCustomImport`

NewServerConfigCustomImportWithDefaults instantiates a new ServerConfigCustomImport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVideos

`func (o *ServerConfigCustomImport) GetVideos() ServerConfigImportVideos`

GetVideos returns the Videos field if non-nil, zero value otherwise.

### GetVideosOk

`func (o *ServerConfigCustomImport) GetVideosOk() (*ServerConfigImportVideos, bool)`

GetVideosOk returns a tuple with the Videos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideos

`func (o *ServerConfigCustomImport) SetVideos(v ServerConfigImportVideos)`

SetVideos sets Videos field to given value.

### HasVideos

`func (o *ServerConfigCustomImport) HasVideos() bool`

HasVideos returns a boolean if a field has been set.

### GetVideoChannelSynchronization

`func (o *ServerConfigCustomImport) GetVideoChannelSynchronization() ServerConfigEmail`

GetVideoChannelSynchronization returns the VideoChannelSynchronization field if non-nil, zero value otherwise.

### GetVideoChannelSynchronizationOk

`func (o *ServerConfigCustomImport) GetVideoChannelSynchronizationOk() (*ServerConfigEmail, bool)`

GetVideoChannelSynchronizationOk returns a tuple with the VideoChannelSynchronization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoChannelSynchronization

`func (o *ServerConfigCustomImport) SetVideoChannelSynchronization(v ServerConfigEmail)`

SetVideoChannelSynchronization sets VideoChannelSynchronization field to given value.

### HasVideoChannelSynchronization

`func (o *ServerConfigCustomImport) HasVideoChannelSynchronization() bool`

HasVideoChannelSynchronization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


