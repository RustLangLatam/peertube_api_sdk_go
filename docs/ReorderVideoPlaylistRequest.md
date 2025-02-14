# ReorderVideoPlaylistRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartPosition** | **int32** | Start position of the element to reorder | 
**InsertAfterPosition** | **int32** | New position for the block to reorder, to add the block before the first element | 
**ReorderLength** | Pointer to **int32** | How many element from &#x60;startPosition&#x60; to reorder | [optional] 

## Methods

### NewReorderVideoPlaylistRequest

`func NewReorderVideoPlaylistRequest(startPosition int32, insertAfterPosition int32, ) *ReorderVideoPlaylistRequest`

NewReorderVideoPlaylistRequest instantiates a new ReorderVideoPlaylistRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReorderVideoPlaylistRequestWithDefaults

`func NewReorderVideoPlaylistRequestWithDefaults() *ReorderVideoPlaylistRequest`

NewReorderVideoPlaylistRequestWithDefaults instantiates a new ReorderVideoPlaylistRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartPosition

`func (o *ReorderVideoPlaylistRequest) GetStartPosition() int32`

GetStartPosition returns the StartPosition field if non-nil, zero value otherwise.

### GetStartPositionOk

`func (o *ReorderVideoPlaylistRequest) GetStartPositionOk() (*int32, bool)`

GetStartPositionOk returns a tuple with the StartPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartPosition

`func (o *ReorderVideoPlaylistRequest) SetStartPosition(v int32)`

SetStartPosition sets StartPosition field to given value.


### GetInsertAfterPosition

`func (o *ReorderVideoPlaylistRequest) GetInsertAfterPosition() int32`

GetInsertAfterPosition returns the InsertAfterPosition field if non-nil, zero value otherwise.

### GetInsertAfterPositionOk

`func (o *ReorderVideoPlaylistRequest) GetInsertAfterPositionOk() (*int32, bool)`

GetInsertAfterPositionOk returns a tuple with the InsertAfterPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsertAfterPosition

`func (o *ReorderVideoPlaylistRequest) SetInsertAfterPosition(v int32)`

SetInsertAfterPosition sets InsertAfterPosition field to given value.


### GetReorderLength

`func (o *ReorderVideoPlaylistRequest) GetReorderLength() int32`

GetReorderLength returns the ReorderLength field if non-nil, zero value otherwise.

### GetReorderLengthOk

`func (o *ReorderVideoPlaylistRequest) GetReorderLengthOk() (*int32, bool)`

GetReorderLengthOk returns a tuple with the ReorderLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReorderLength

`func (o *ReorderVideoPlaylistRequest) SetReorderLength(v int32)`

SetReorderLength sets ReorderLength field to given value.

### HasReorderLength

`func (o *ReorderVideoPlaylistRequest) HasReorderLength() bool`

HasReorderLength returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


