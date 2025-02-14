# VideoRedundancy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 
**Redundancies** | Pointer to [**VideoRedundancyRedundancies**](VideoRedundancyRedundancies.md) |  | [optional] 

## Methods

### NewVideoRedundancy

`func NewVideoRedundancy() *VideoRedundancy`

NewVideoRedundancy instantiates a new VideoRedundancy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVideoRedundancyWithDefaults

`func NewVideoRedundancyWithDefaults() *VideoRedundancy`

NewVideoRedundancyWithDefaults instantiates a new VideoRedundancy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VideoRedundancy) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VideoRedundancy) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VideoRedundancy) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *VideoRedundancy) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *VideoRedundancy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VideoRedundancy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VideoRedundancy) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VideoRedundancy) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUrl

`func (o *VideoRedundancy) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *VideoRedundancy) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *VideoRedundancy) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *VideoRedundancy) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetUuid

`func (o *VideoRedundancy) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *VideoRedundancy) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *VideoRedundancy) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *VideoRedundancy) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetRedundancies

`func (o *VideoRedundancy) GetRedundancies() VideoRedundancyRedundancies`

GetRedundancies returns the Redundancies field if non-nil, zero value otherwise.

### GetRedundanciesOk

`func (o *VideoRedundancy) GetRedundanciesOk() (*VideoRedundancyRedundancies, bool)`

GetRedundanciesOk returns a tuple with the Redundancies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedundancies

`func (o *VideoRedundancy) SetRedundancies(v VideoRedundancyRedundancies)`

SetRedundancies sets Redundancies field to given value.

### HasRedundancies

`func (o *VideoRedundancy) HasRedundancies() bool`

HasRedundancies returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


