# ServerConfigImport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Videos** | Pointer to [**ServerConfigImportVideos**](ServerConfigImportVideos.md) |  | [optional] 
**VideoChannelSynchronization** | Pointer to [**ServerConfigEmail**](ServerConfigEmail.md) |  | [optional] 
**Users** | Pointer to [**ServerConfigEmail**](ServerConfigEmail.md) |  | [optional] 

## Methods

### NewServerConfigImport

`func NewServerConfigImport() *ServerConfigImport`

NewServerConfigImport instantiates a new ServerConfigImport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerConfigImportWithDefaults

`func NewServerConfigImportWithDefaults() *ServerConfigImport`

NewServerConfigImportWithDefaults instantiates a new ServerConfigImport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVideos

`func (o *ServerConfigImport) GetVideos() ServerConfigImportVideos`

GetVideos returns the Videos field if non-nil, zero value otherwise.

### GetVideosOk

`func (o *ServerConfigImport) GetVideosOk() (*ServerConfigImportVideos, bool)`

GetVideosOk returns a tuple with the Videos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideos

`func (o *ServerConfigImport) SetVideos(v ServerConfigImportVideos)`

SetVideos sets Videos field to given value.

### HasVideos

`func (o *ServerConfigImport) HasVideos() bool`

HasVideos returns a boolean if a field has been set.

### GetVideoChannelSynchronization

`func (o *ServerConfigImport) GetVideoChannelSynchronization() ServerConfigEmail`

GetVideoChannelSynchronization returns the VideoChannelSynchronization field if non-nil, zero value otherwise.

### GetVideoChannelSynchronizationOk

`func (o *ServerConfigImport) GetVideoChannelSynchronizationOk() (*ServerConfigEmail, bool)`

GetVideoChannelSynchronizationOk returns a tuple with the VideoChannelSynchronization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoChannelSynchronization

`func (o *ServerConfigImport) SetVideoChannelSynchronization(v ServerConfigEmail)`

SetVideoChannelSynchronization sets VideoChannelSynchronization field to given value.

### HasVideoChannelSynchronization

`func (o *ServerConfigImport) HasVideoChannelSynchronization() bool`

HasVideoChannelSynchronization returns a boolean if a field has been set.

### GetUsers

`func (o *ServerConfigImport) GetUsers() ServerConfigEmail`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *ServerConfigImport) GetUsersOk() (*ServerConfigEmail, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *ServerConfigImport) SetUsers(v ServerConfigEmail)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *ServerConfigImport) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


