# VideoChannelSync

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**State** | Pointer to [**VideoChannelSyncState**](VideoChannelSyncState.md) |  | [optional] 
**ExternalChannelUrl** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**LastSyncAt** | Pointer to **NullableTime** |  | [optional] 
**Channel** | Pointer to [**VideoChannel**](VideoChannel.md) |  | [optional] 

## Methods

### NewVideoChannelSync

`func NewVideoChannelSync() *VideoChannelSync`

NewVideoChannelSync instantiates a new VideoChannelSync object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVideoChannelSyncWithDefaults

`func NewVideoChannelSyncWithDefaults() *VideoChannelSync`

NewVideoChannelSyncWithDefaults instantiates a new VideoChannelSync object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VideoChannelSync) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VideoChannelSync) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VideoChannelSync) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *VideoChannelSync) HasId() bool`

HasId returns a boolean if a field has been set.

### GetState

`func (o *VideoChannelSync) GetState() VideoChannelSyncState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *VideoChannelSync) GetStateOk() (*VideoChannelSyncState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *VideoChannelSync) SetState(v VideoChannelSyncState)`

SetState sets State field to given value.

### HasState

`func (o *VideoChannelSync) HasState() bool`

HasState returns a boolean if a field has been set.

### GetExternalChannelUrl

`func (o *VideoChannelSync) GetExternalChannelUrl() string`

GetExternalChannelUrl returns the ExternalChannelUrl field if non-nil, zero value otherwise.

### GetExternalChannelUrlOk

`func (o *VideoChannelSync) GetExternalChannelUrlOk() (*string, bool)`

GetExternalChannelUrlOk returns a tuple with the ExternalChannelUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalChannelUrl

`func (o *VideoChannelSync) SetExternalChannelUrl(v string)`

SetExternalChannelUrl sets ExternalChannelUrl field to given value.

### HasExternalChannelUrl

`func (o *VideoChannelSync) HasExternalChannelUrl() bool`

HasExternalChannelUrl returns a boolean if a field has been set.

### GetCreatedAt

`func (o *VideoChannelSync) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *VideoChannelSync) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *VideoChannelSync) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *VideoChannelSync) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetLastSyncAt

`func (o *VideoChannelSync) GetLastSyncAt() time.Time`

GetLastSyncAt returns the LastSyncAt field if non-nil, zero value otherwise.

### GetLastSyncAtOk

`func (o *VideoChannelSync) GetLastSyncAtOk() (*time.Time, bool)`

GetLastSyncAtOk returns a tuple with the LastSyncAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSyncAt

`func (o *VideoChannelSync) SetLastSyncAt(v time.Time)`

SetLastSyncAt sets LastSyncAt field to given value.

### HasLastSyncAt

`func (o *VideoChannelSync) HasLastSyncAt() bool`

HasLastSyncAt returns a boolean if a field has been set.

### SetLastSyncAtNil

`func (o *VideoChannelSync) SetLastSyncAtNil(b bool)`

 SetLastSyncAtNil sets the value for LastSyncAt to be an explicit nil

### UnsetLastSyncAt
`func (o *VideoChannelSync) UnsetLastSyncAt()`

UnsetLastSyncAt ensures that no value is present for LastSyncAt, not even an explicit nil
### GetChannel

`func (o *VideoChannelSync) GetChannel() VideoChannel`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *VideoChannelSync) GetChannelOk() (*VideoChannel, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *VideoChannelSync) SetChannel(v VideoChannel)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *VideoChannelSync) HasChannel() bool`

HasChannel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


