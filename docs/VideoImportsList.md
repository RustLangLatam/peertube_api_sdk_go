# VideoImportsList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int32** |  | [optional] 
**Data** | Pointer to [**[]VideoImport**](VideoImport.md) |  | [optional] 

## Methods

### NewVideoImportsList

`func NewVideoImportsList() *VideoImportsList`

NewVideoImportsList instantiates a new VideoImportsList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVideoImportsListWithDefaults

`func NewVideoImportsListWithDefaults() *VideoImportsList`

NewVideoImportsListWithDefaults instantiates a new VideoImportsList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *VideoImportsList) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *VideoImportsList) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *VideoImportsList) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *VideoImportsList) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetData

`func (o *VideoImportsList) GetData() []VideoImport`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *VideoImportsList) GetDataOk() (*[]VideoImport, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *VideoImportsList) SetData(v []VideoImport)`

SetData sets Data field to given value.

### HasData

`func (o *VideoImportsList) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


