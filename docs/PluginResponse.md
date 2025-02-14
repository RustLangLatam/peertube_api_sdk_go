# PluginResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int32** |  | [optional] 
**Data** | Pointer to [**[]Plugin**](Plugin.md) |  | [optional] 

## Methods

### NewPluginResponse

`func NewPluginResponse() *PluginResponse`

NewPluginResponse instantiates a new PluginResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPluginResponseWithDefaults

`func NewPluginResponseWithDefaults() *PluginResponse`

NewPluginResponseWithDefaults instantiates a new PluginResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *PluginResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *PluginResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *PluginResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *PluginResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetData

`func (o *PluginResponse) GetData() []Plugin`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PluginResponse) GetDataOk() (*[]Plugin, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PluginResponse) SetData(v []Plugin)`

SetData sets Data field to given value.

### HasData

`func (o *PluginResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


