# Job

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Data** | Pointer to **map[string]interface{}** |  | [optional] 
**Error** | Pointer to **map[string]interface{}** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**FinishedOn** | Pointer to **time.Time** |  | [optional] 
**ProcessedOn** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewJob

`func NewJob() *Job`

NewJob instantiates a new Job object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobWithDefaults

`func NewJobWithDefaults() *Job`

NewJobWithDefaults instantiates a new Job object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Job) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Job) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Job) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Job) HasId() bool`

HasId returns a boolean if a field has been set.

### GetState

`func (o *Job) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Job) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Job) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Job) HasState() bool`

HasState returns a boolean if a field has been set.

### GetType

`func (o *Job) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Job) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Job) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Job) HasType() bool`

HasType returns a boolean if a field has been set.

### GetData

`func (o *Job) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *Job) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *Job) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *Job) HasData() bool`

HasData returns a boolean if a field has been set.

### GetError

`func (o *Job) GetError() map[string]interface{}`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *Job) GetErrorOk() (*map[string]interface{}, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *Job) SetError(v map[string]interface{})`

SetError sets Error field to given value.

### HasError

`func (o *Job) HasError() bool`

HasError returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Job) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Job) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Job) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Job) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetFinishedOn

`func (o *Job) GetFinishedOn() time.Time`

GetFinishedOn returns the FinishedOn field if non-nil, zero value otherwise.

### GetFinishedOnOk

`func (o *Job) GetFinishedOnOk() (*time.Time, bool)`

GetFinishedOnOk returns a tuple with the FinishedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedOn

`func (o *Job) SetFinishedOn(v time.Time)`

SetFinishedOn sets FinishedOn field to given value.

### HasFinishedOn

`func (o *Job) HasFinishedOn() bool`

HasFinishedOn returns a boolean if a field has been set.

### GetProcessedOn

`func (o *Job) GetProcessedOn() time.Time`

GetProcessedOn returns the ProcessedOn field if non-nil, zero value otherwise.

### GetProcessedOnOk

`func (o *Job) GetProcessedOnOk() (*time.Time, bool)`

GetProcessedOnOk returns a tuple with the ProcessedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessedOn

`func (o *Job) SetProcessedOn(v time.Time)`

SetProcessedOn sets ProcessedOn field to given value.

### HasProcessedOn

`func (o *Job) HasProcessedOn() bool`

HasProcessedOn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


