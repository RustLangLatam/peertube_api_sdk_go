# CreateVideoTranscodingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TranscodingType** | **string** |  | 
**ForceTranscoding** | Pointer to **bool** | If the video is stuck in transcoding state, do it anyway | [optional] [default to false]

## Methods

### NewCreateVideoTranscodingRequest

`func NewCreateVideoTranscodingRequest(transcodingType string, ) *CreateVideoTranscodingRequest`

NewCreateVideoTranscodingRequest instantiates a new CreateVideoTranscodingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateVideoTranscodingRequestWithDefaults

`func NewCreateVideoTranscodingRequestWithDefaults() *CreateVideoTranscodingRequest`

NewCreateVideoTranscodingRequestWithDefaults instantiates a new CreateVideoTranscodingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTranscodingType

`func (o *CreateVideoTranscodingRequest) GetTranscodingType() string`

GetTranscodingType returns the TranscodingType field if non-nil, zero value otherwise.

### GetTranscodingTypeOk

`func (o *CreateVideoTranscodingRequest) GetTranscodingTypeOk() (*string, bool)`

GetTranscodingTypeOk returns a tuple with the TranscodingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscodingType

`func (o *CreateVideoTranscodingRequest) SetTranscodingType(v string)`

SetTranscodingType sets TranscodingType field to given value.


### GetForceTranscoding

`func (o *CreateVideoTranscodingRequest) GetForceTranscoding() bool`

GetForceTranscoding returns the ForceTranscoding field if non-nil, zero value otherwise.

### GetForceTranscodingOk

`func (o *CreateVideoTranscodingRequest) GetForceTranscodingOk() (*bool, bool)`

GetForceTranscodingOk returns a tuple with the ForceTranscoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceTranscoding

`func (o *CreateVideoTranscodingRequest) SetForceTranscoding(v bool)`

SetForceTranscoding sets ForceTranscoding field to given value.

### HasForceTranscoding

`func (o *CreateVideoTranscodingRequest) HasForceTranscoding() bool`

HasForceTranscoding returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


