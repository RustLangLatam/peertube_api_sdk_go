# RunnerJobParent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**RunnerJobType**](RunnerJobType.md) |  | [optional] 
**State** | Pointer to [**RunnerJobStateConstant**](RunnerJobStateConstant.md) |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 

## Methods

### NewRunnerJobParent

`func NewRunnerJobParent() *RunnerJobParent`

NewRunnerJobParent instantiates a new RunnerJobParent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunnerJobParentWithDefaults

`func NewRunnerJobParentWithDefaults() *RunnerJobParent`

NewRunnerJobParentWithDefaults instantiates a new RunnerJobParent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RunnerJobParent) GetType() RunnerJobType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RunnerJobParent) GetTypeOk() (*RunnerJobType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RunnerJobParent) SetType(v RunnerJobType)`

SetType sets Type field to given value.

### HasType

`func (o *RunnerJobParent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetState

`func (o *RunnerJobParent) GetState() RunnerJobStateConstant`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *RunnerJobParent) GetStateOk() (*RunnerJobStateConstant, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *RunnerJobParent) SetState(v RunnerJobStateConstant)`

SetState sets State field to given value.

### HasState

`func (o *RunnerJobParent) HasState() bool`

HasState returns a boolean if a field has been set.

### GetUuid

`func (o *RunnerJobParent) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *RunnerJobParent) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *RunnerJobParent) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *RunnerJobParent) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


