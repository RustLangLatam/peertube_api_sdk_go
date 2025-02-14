# VideoResolutionConstant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Video resolution (&#x60;0&#x60;, &#x60;240&#x60;, &#x60;360&#x60;, &#x60;720&#x60;, &#x60;1080&#x60;, &#x60;1440&#x60; or &#x60;2160&#x60;)  &#x60;0&#x60; is used as a special value for stillimage videos dedicated to audio, a.k.a. audio-only videos.  | [optional] 
**Label** | Pointer to **string** |  | [optional] 

## Methods

### NewVideoResolutionConstant

`func NewVideoResolutionConstant() *VideoResolutionConstant`

NewVideoResolutionConstant instantiates a new VideoResolutionConstant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVideoResolutionConstantWithDefaults

`func NewVideoResolutionConstantWithDefaults() *VideoResolutionConstant`

NewVideoResolutionConstantWithDefaults instantiates a new VideoResolutionConstant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VideoResolutionConstant) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VideoResolutionConstant) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VideoResolutionConstant) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *VideoResolutionConstant) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabel

`func (o *VideoResolutionConstant) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *VideoResolutionConstant) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *VideoResolutionConstant) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *VideoResolutionConstant) HasLabel() bool`

HasLabel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


