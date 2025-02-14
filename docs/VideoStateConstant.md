# VideoStateConstant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The video state: - &#x60;1&#x60;: Published - &#x60;2&#x60;: To transcode - &#x60;3&#x60;: To import - &#x60;4&#x60;: Waiting for live stream - &#x60;5&#x60;: Live ended - &#x60;6&#x60;: To move to an external storage (object storage...) - &#x60;7&#x60;: Transcoding failed - &#x60;8&#x60;: Moving to an external storage failed - &#x60;9&#x60;: To edit using studio edition feature  | [optional] 
**Label** | Pointer to **string** |  | [optional] 

## Methods

### NewVideoStateConstant

`func NewVideoStateConstant() *VideoStateConstant`

NewVideoStateConstant instantiates a new VideoStateConstant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVideoStateConstantWithDefaults

`func NewVideoStateConstantWithDefaults() *VideoStateConstant`

NewVideoStateConstantWithDefaults instantiates a new VideoStateConstant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VideoStateConstant) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VideoStateConstant) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VideoStateConstant) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *VideoStateConstant) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabel

`func (o *VideoStateConstant) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *VideoStateConstant) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *VideoStateConstant) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *VideoStateConstant) HasLabel() bool`

HasLabel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


