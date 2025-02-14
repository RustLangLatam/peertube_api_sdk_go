# GetAccountFollowers200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int32** |  | [optional] 
**Data** | Pointer to [**[]Follow**](Follow.md) |  | [optional] 

## Methods

### NewGetAccountFollowers200Response

`func NewGetAccountFollowers200Response() *GetAccountFollowers200Response`

NewGetAccountFollowers200Response instantiates a new GetAccountFollowers200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAccountFollowers200ResponseWithDefaults

`func NewGetAccountFollowers200ResponseWithDefaults() *GetAccountFollowers200Response`

NewGetAccountFollowers200ResponseWithDefaults instantiates a new GetAccountFollowers200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *GetAccountFollowers200Response) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetAccountFollowers200Response) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetAccountFollowers200Response) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GetAccountFollowers200Response) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetData

`func (o *GetAccountFollowers200Response) GetData() []Follow`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetAccountFollowers200Response) GetDataOk() (*[]Follow, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetAccountFollowers200Response) SetData(v []Follow)`

SetData sets Data field to given value.

### HasData

`func (o *GetAccountFollowers200Response) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


