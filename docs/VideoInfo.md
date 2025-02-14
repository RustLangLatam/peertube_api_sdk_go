# VideoInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** | title of the video | [optional] 
**State** | Pointer to [**VideoStateConstant**](VideoStateConstant.md) | represents the internal state of the video processing within the PeerTube instance | [optional] 

## Methods

### NewVideoInfo

`func NewVideoInfo() *VideoInfo`

NewVideoInfo instantiates a new VideoInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVideoInfoWithDefaults

`func NewVideoInfoWithDefaults() *VideoInfo`

NewVideoInfoWithDefaults instantiates a new VideoInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VideoInfo) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VideoInfo) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VideoInfo) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *VideoInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *VideoInfo) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *VideoInfo) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *VideoInfo) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *VideoInfo) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetName

`func (o *VideoInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VideoInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VideoInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VideoInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetState

`func (o *VideoInfo) GetState() VideoStateConstant`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *VideoInfo) GetStateOk() (*VideoStateConstant, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *VideoInfo) SetState(v VideoStateConstant)`

SetState sets State field to given value.

### HasState

`func (o *VideoInfo) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


