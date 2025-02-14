# VideoPasswordList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int32** |  | [optional] 
**Data** | Pointer to [**[]VideoPassword**](VideoPassword.md) |  | [optional] 

## Methods

### NewVideoPasswordList

`func NewVideoPasswordList() *VideoPasswordList`

NewVideoPasswordList instantiates a new VideoPasswordList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVideoPasswordListWithDefaults

`func NewVideoPasswordListWithDefaults() *VideoPasswordList`

NewVideoPasswordListWithDefaults instantiates a new VideoPasswordList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *VideoPasswordList) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *VideoPasswordList) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *VideoPasswordList) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *VideoPasswordList) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetData

`func (o *VideoPasswordList) GetData() []VideoPassword`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *VideoPasswordList) GetDataOk() (*[]VideoPassword, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *VideoPasswordList) SetData(v []VideoPassword)`

SetData sets Data field to given value.

### HasData

`func (o *VideoPasswordList) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


