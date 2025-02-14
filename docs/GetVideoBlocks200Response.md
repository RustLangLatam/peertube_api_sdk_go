# GetVideoBlocks200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int32** |  | [optional] 
**Data** | Pointer to [**[]VideoBlacklist**](VideoBlacklist.md) |  | [optional] 

## Methods

### NewGetVideoBlocks200Response

`func NewGetVideoBlocks200Response() *GetVideoBlocks200Response`

NewGetVideoBlocks200Response instantiates a new GetVideoBlocks200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetVideoBlocks200ResponseWithDefaults

`func NewGetVideoBlocks200ResponseWithDefaults() *GetVideoBlocks200Response`

NewGetVideoBlocks200ResponseWithDefaults instantiates a new GetVideoBlocks200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *GetVideoBlocks200Response) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetVideoBlocks200Response) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetVideoBlocks200Response) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GetVideoBlocks200Response) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetData

`func (o *GetVideoBlocks200Response) GetData() []VideoBlacklist`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetVideoBlocks200Response) GetDataOk() (*[]VideoBlacklist, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetVideoBlocks200Response) SetData(v []VideoBlacklist)`

SetData sets Data field to given value.

### HasData

`func (o *GetVideoBlocks200Response) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


