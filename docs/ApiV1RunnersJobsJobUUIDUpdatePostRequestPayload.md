# ApiV1RunnersJobsJobUUIDUpdatePostRequestPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**MasterPlaylistFile** | Pointer to ***os.File** |  | [optional] 
**ResolutionPlaylistFile** | Pointer to ***os.File** |  | [optional] 
**ResolutionPlaylistFilename** | Pointer to **string** |  | [optional] 
**VideoChunkFile** | Pointer to ***os.File** |  | [optional] 
**VideoChunkFilename** | Pointer to **string** |  | [optional] 

## Methods

### NewApiV1RunnersJobsJobUUIDUpdatePostRequestPayload

`func NewApiV1RunnersJobsJobUUIDUpdatePostRequestPayload() *ApiV1RunnersJobsJobUUIDUpdatePostRequestPayload`

NewApiV1RunnersJobsJobUUIDUpdatePostRequestPayload instantiates a new ApiV1RunnersJobsJobUUIDUpdatePostRequestPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV1RunnersJobsJobUUIDUpdatePostRequestPayloadWithDefaults

`func NewApiV1RunnersJobsJobUUIDUpdatePostRequestPayloadWithDefaults() *ApiV1RunnersJobsJobUUIDUpdatePostRequestPayload`

NewApiV1RunnersJobsJobUUIDUpdatePostRequestPayloadWithDefaults instantiates a new ApiV1RunnersJobsJobUUIDUpdatePostRequestPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ApiV1RunnersJobsJobUUIDUpdatePostRequestPayload) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApiV1RunnersJobsJobUUIDUpdatePostRequestPayload) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApiV1RunnersJobsJobUUIDUpdatePostRequestPayload) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ApiV1RunnersJobsJobUUIDUpdatePostRequestPayload) HasType() bool`

HasType returns a boolean if a field has been set.

### GetMasterPlaylistFile

`func (o *ApiV1RunnersJobsJobUUIDUpdatePostRequestPayload) GetMasterPlaylistFile() *os.File`

GetMasterPlaylistFile returns the MasterPlaylistFile field if non-nil, zero value otherwise.

### GetMasterPlaylistFileOk

`func (o *ApiV1RunnersJobsJobUUIDUpdatePostRequestPayload) GetMasterPlaylistFileOk() (**os.File, bool)`

GetMasterPlaylistFileOk returns a tuple with the MasterPlaylistFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterPlaylistFile

`func (o *ApiV1RunnersJobsJobUUIDUpdatePostRequestPayload) SetMasterPlaylistFile(v *os.File)`

SetMasterPlaylistFile sets MasterPlaylistFile field to given value.

### HasMasterPlaylistFile

`func (o *ApiV1RunnersJobsJobUUIDUpdatePostRequestPayload) HasMasterPlaylistFile() bool`

HasMasterPlaylistFile returns a boolean if a field has been set.

### GetResolutionPlaylistFile

`func (o *ApiV1RunnersJobsJobUUIDUpdatePostRequestPayload) GetResolutionPlaylistFile() *os.File`

GetResolutionPlaylistFile returns the ResolutionPlaylistFile field if non-nil, zero value otherwise.

### GetResolutionPlaylistFileOk

`func (o *ApiV1RunnersJobsJobUUIDUpdatePostRequestPayload) GetResolutionPlaylistFileOk() (**os.File, bool)`

GetResolutionPlaylistFileOk returns a tuple with the ResolutionPlaylistFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolutionPlaylistFile

`func (o *ApiV1RunnersJobsJobUUIDUpdatePostRequestPayload) SetResolutionPlaylistFile(v *os.File)`

SetResolutionPlaylistFile sets ResolutionPlaylistFile field to given value.

### HasResolutionPlaylistFile

`func (o *ApiV1RunnersJobsJobUUIDUpdatePostRequestPayload) HasResolutionPlaylistFile() bool`

HasResolutionPlaylistFile returns a boolean if a field has been set.

### GetResolutionPlaylistFilename

`func (o *ApiV1RunnersJobsJobUUIDUpdatePostRequestPayload) GetResolutionPlaylistFilename() string`

GetResolutionPlaylistFilename returns the ResolutionPlaylistFilename field if non-nil, zero value otherwise.

### GetResolutionPlaylistFilenameOk

`func (o *ApiV1RunnersJobsJobUUIDUpdatePostRequestPayload) GetResolutionPlaylistFilenameOk() (*string, bool)`

GetResolutionPlaylistFilenameOk returns a tuple with the ResolutionPlaylistFilename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolutionPlaylistFilename

`func (o *ApiV1RunnersJobsJobUUIDUpdatePostRequestPayload) SetResolutionPlaylistFilename(v string)`

SetResolutionPlaylistFilename sets ResolutionPlaylistFilename field to given value.

### HasResolutionPlaylistFilename

`func (o *ApiV1RunnersJobsJobUUIDUpdatePostRequestPayload) HasResolutionPlaylistFilename() bool`

HasResolutionPlaylistFilename returns a boolean if a field has been set.

### GetVideoChunkFile

`func (o *ApiV1RunnersJobsJobUUIDUpdatePostRequestPayload) GetVideoChunkFile() *os.File`

GetVideoChunkFile returns the VideoChunkFile field if non-nil, zero value otherwise.

### GetVideoChunkFileOk

`func (o *ApiV1RunnersJobsJobUUIDUpdatePostRequestPayload) GetVideoChunkFileOk() (**os.File, bool)`

GetVideoChunkFileOk returns a tuple with the VideoChunkFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoChunkFile

`func (o *ApiV1RunnersJobsJobUUIDUpdatePostRequestPayload) SetVideoChunkFile(v *os.File)`

SetVideoChunkFile sets VideoChunkFile field to given value.

### HasVideoChunkFile

`func (o *ApiV1RunnersJobsJobUUIDUpdatePostRequestPayload) HasVideoChunkFile() bool`

HasVideoChunkFile returns a boolean if a field has been set.

### GetVideoChunkFilename

`func (o *ApiV1RunnersJobsJobUUIDUpdatePostRequestPayload) GetVideoChunkFilename() string`

GetVideoChunkFilename returns the VideoChunkFilename field if non-nil, zero value otherwise.

### GetVideoChunkFilenameOk

`func (o *ApiV1RunnersJobsJobUUIDUpdatePostRequestPayload) GetVideoChunkFilenameOk() (*string, bool)`

GetVideoChunkFilenameOk returns a tuple with the VideoChunkFilename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoChunkFilename

`func (o *ApiV1RunnersJobsJobUUIDUpdatePostRequestPayload) SetVideoChunkFilename(v string)`

SetVideoChunkFilename sets VideoChunkFilename field to given value.

### HasVideoChunkFilename

`func (o *ApiV1RunnersJobsJobUUIDUpdatePostRequestPayload) HasVideoChunkFilename() bool`

HasVideoChunkFilename returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


