# GetVideoCaptions200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int32** |  | [optional] 
**Data** | Pointer to [**[]VideoCaption**](VideoCaption.md) |  | [optional] 

## Methods

### NewGetVideoCaptions200Response

`func NewGetVideoCaptions200Response() *GetVideoCaptions200Response`

NewGetVideoCaptions200Response instantiates a new GetVideoCaptions200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetVideoCaptions200ResponseWithDefaults

`func NewGetVideoCaptions200ResponseWithDefaults() *GetVideoCaptions200Response`

NewGetVideoCaptions200ResponseWithDefaults instantiates a new GetVideoCaptions200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *GetVideoCaptions200Response) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetVideoCaptions200Response) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetVideoCaptions200Response) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GetVideoCaptions200Response) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetData

`func (o *GetVideoCaptions200Response) GetData() []VideoCaption`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetVideoCaptions200Response) GetDataOk() (*[]VideoCaption, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetVideoCaptions200Response) SetData(v []VideoCaption)`

SetData sets Data field to given value.

### HasData

`func (o *GetVideoCaptions200Response) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


