# ImportVideosInChannelCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalChannelUrl** | **string** |  | 
**VideoChannelSyncId** | Pointer to **int32** | If part of a channel sync process, specify its id to assign video imports to this channel synchronization | [optional] 

## Methods

### NewImportVideosInChannelCreate

`func NewImportVideosInChannelCreate(externalChannelUrl string, ) *ImportVideosInChannelCreate`

NewImportVideosInChannelCreate instantiates a new ImportVideosInChannelCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportVideosInChannelCreateWithDefaults

`func NewImportVideosInChannelCreateWithDefaults() *ImportVideosInChannelCreate`

NewImportVideosInChannelCreateWithDefaults instantiates a new ImportVideosInChannelCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalChannelUrl

`func (o *ImportVideosInChannelCreate) GetExternalChannelUrl() string`

GetExternalChannelUrl returns the ExternalChannelUrl field if non-nil, zero value otherwise.

### GetExternalChannelUrlOk

`func (o *ImportVideosInChannelCreate) GetExternalChannelUrlOk() (*string, bool)`

GetExternalChannelUrlOk returns a tuple with the ExternalChannelUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalChannelUrl

`func (o *ImportVideosInChannelCreate) SetExternalChannelUrl(v string)`

SetExternalChannelUrl sets ExternalChannelUrl field to given value.


### GetVideoChannelSyncId

`func (o *ImportVideosInChannelCreate) GetVideoChannelSyncId() int32`

GetVideoChannelSyncId returns the VideoChannelSyncId field if non-nil, zero value otherwise.

### GetVideoChannelSyncIdOk

`func (o *ImportVideosInChannelCreate) GetVideoChannelSyncIdOk() (*int32, bool)`

GetVideoChannelSyncIdOk returns a tuple with the VideoChannelSyncId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoChannelSyncId

`func (o *ImportVideosInChannelCreate) SetVideoChannelSyncId(v int32)`

SetVideoChannelSyncId sets VideoChannelSyncId field to given value.

### HasVideoChannelSyncId

`func (o *ImportVideosInChannelCreate) HasVideoChannelSyncId() bool`

HasVideoChannelSyncId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


