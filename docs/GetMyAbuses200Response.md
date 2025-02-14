# GetMyAbuses200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int32** |  | [optional] 
**Data** | Pointer to [**[]Abuse**](Abuse.md) |  | [optional] 

## Methods

### NewGetMyAbuses200Response

`func NewGetMyAbuses200Response() *GetMyAbuses200Response`

NewGetMyAbuses200Response instantiates a new GetMyAbuses200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMyAbuses200ResponseWithDefaults

`func NewGetMyAbuses200ResponseWithDefaults() *GetMyAbuses200Response`

NewGetMyAbuses200ResponseWithDefaults instantiates a new GetMyAbuses200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *GetMyAbuses200Response) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetMyAbuses200Response) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetMyAbuses200Response) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GetMyAbuses200Response) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetData

`func (o *GetMyAbuses200Response) GetData() []Abuse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetMyAbuses200Response) GetDataOk() (*[]Abuse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetMyAbuses200Response) SetData(v []Abuse)`

SetData sets Data field to given value.

### HasData

`func (o *GetMyAbuses200Response) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


