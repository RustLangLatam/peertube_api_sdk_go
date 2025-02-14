# GetLatestUserImport200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**State** | Pointer to [**GetLatestUserImport200ResponseState**](GetLatestUserImport200ResponseState.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewGetLatestUserImport200Response

`func NewGetLatestUserImport200Response() *GetLatestUserImport200Response`

NewGetLatestUserImport200Response instantiates a new GetLatestUserImport200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetLatestUserImport200ResponseWithDefaults

`func NewGetLatestUserImport200ResponseWithDefaults() *GetLatestUserImport200Response`

NewGetLatestUserImport200ResponseWithDefaults instantiates a new GetLatestUserImport200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetLatestUserImport200Response) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetLatestUserImport200Response) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetLatestUserImport200Response) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GetLatestUserImport200Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetState

`func (o *GetLatestUserImport200Response) GetState() GetLatestUserImport200ResponseState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GetLatestUserImport200Response) GetStateOk() (*GetLatestUserImport200ResponseState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GetLatestUserImport200Response) SetState(v GetLatestUserImport200ResponseState)`

SetState sets State field to given value.

### HasState

`func (o *GetLatestUserImport200Response) HasState() bool`

HasState returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GetLatestUserImport200Response) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetLatestUserImport200Response) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetLatestUserImport200Response) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GetLatestUserImport200Response) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


