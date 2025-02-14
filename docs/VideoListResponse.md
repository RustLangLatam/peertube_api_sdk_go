# VideoListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int32** |  | [optional] 
**Data** | Pointer to [**[]Video**](Video.md) |  | [optional] 

## Methods

### NewVideoListResponse

`func NewVideoListResponse() *VideoListResponse`

NewVideoListResponse instantiates a new VideoListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVideoListResponseWithDefaults

`func NewVideoListResponseWithDefaults() *VideoListResponse`

NewVideoListResponseWithDefaults instantiates a new VideoListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *VideoListResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *VideoListResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *VideoListResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *VideoListResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetData

`func (o *VideoListResponse) GetData() []Video`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *VideoListResponse) GetDataOk() (*[]Video, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *VideoListResponse) SetData(v []Video)`

SetData sets Data field to given value.

### HasData

`func (o *VideoListResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


