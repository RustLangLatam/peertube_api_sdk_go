# ApiV1RunnersJobsGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int32** |  | [optional] 
**Data** | Pointer to [**[]RunnerJobAdmin**](RunnerJobAdmin.md) |  | [optional] 

## Methods

### NewApiV1RunnersJobsGet200Response

`func NewApiV1RunnersJobsGet200Response() *ApiV1RunnersJobsGet200Response`

NewApiV1RunnersJobsGet200Response instantiates a new ApiV1RunnersJobsGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV1RunnersJobsGet200ResponseWithDefaults

`func NewApiV1RunnersJobsGet200ResponseWithDefaults() *ApiV1RunnersJobsGet200Response`

NewApiV1RunnersJobsGet200ResponseWithDefaults instantiates a new ApiV1RunnersJobsGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *ApiV1RunnersJobsGet200Response) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ApiV1RunnersJobsGet200Response) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ApiV1RunnersJobsGet200Response) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *ApiV1RunnersJobsGet200Response) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetData

`func (o *ApiV1RunnersJobsGet200Response) GetData() []RunnerJobAdmin`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ApiV1RunnersJobsGet200Response) GetDataOk() (*[]RunnerJobAdmin, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ApiV1RunnersJobsGet200Response) SetData(v []RunnerJobAdmin)`

SetData sets Data field to given value.

### HasData

`func (o *ApiV1RunnersJobsGet200Response) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


