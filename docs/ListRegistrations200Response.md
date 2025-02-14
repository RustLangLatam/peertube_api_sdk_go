# ListRegistrations200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int32** |  | [optional] 
**Data** | Pointer to [**[]UserRegistration**](UserRegistration.md) |  | [optional] 

## Methods

### NewListRegistrations200Response

`func NewListRegistrations200Response() *ListRegistrations200Response`

NewListRegistrations200Response instantiates a new ListRegistrations200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListRegistrations200ResponseWithDefaults

`func NewListRegistrations200ResponseWithDefaults() *ListRegistrations200Response`

NewListRegistrations200ResponseWithDefaults instantiates a new ListRegistrations200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *ListRegistrations200Response) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ListRegistrations200Response) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ListRegistrations200Response) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *ListRegistrations200Response) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetData

`func (o *ListRegistrations200Response) GetData() []UserRegistration`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListRegistrations200Response) GetDataOk() (*[]UserRegistration, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListRegistrations200Response) SetData(v []UserRegistration)`

SetData sets Data field to given value.

### HasData

`func (o *ListRegistrations200Response) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


