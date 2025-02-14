# VideoImportStateConstant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The video import state (Pending &#x3D; &#x60;1&#x60;, Success &#x3D; &#x60;2&#x60;, Failed &#x3D; &#x60;3&#x60;) | [optional] 
**Label** | Pointer to **string** |  | [optional] 

## Methods

### NewVideoImportStateConstant

`func NewVideoImportStateConstant() *VideoImportStateConstant`

NewVideoImportStateConstant instantiates a new VideoImportStateConstant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVideoImportStateConstantWithDefaults

`func NewVideoImportStateConstantWithDefaults() *VideoImportStateConstant`

NewVideoImportStateConstantWithDefaults instantiates a new VideoImportStateConstant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VideoImportStateConstant) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VideoImportStateConstant) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VideoImportStateConstant) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *VideoImportStateConstant) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabel

`func (o *VideoImportStateConstant) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *VideoImportStateConstant) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *VideoImportStateConstant) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *VideoImportStateConstant) HasLabel() bool`

HasLabel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


