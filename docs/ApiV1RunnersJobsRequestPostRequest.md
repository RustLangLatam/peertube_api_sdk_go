# ApiV1RunnersJobsRequestPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RunnerToken** | **string** |  | 
**JobTypes** | Pointer to **[]string** | Filter jobs depending on their types | [optional] 

## Methods

### NewApiV1RunnersJobsRequestPostRequest

`func NewApiV1RunnersJobsRequestPostRequest(runnerToken string, ) *ApiV1RunnersJobsRequestPostRequest`

NewApiV1RunnersJobsRequestPostRequest instantiates a new ApiV1RunnersJobsRequestPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV1RunnersJobsRequestPostRequestWithDefaults

`func NewApiV1RunnersJobsRequestPostRequestWithDefaults() *ApiV1RunnersJobsRequestPostRequest`

NewApiV1RunnersJobsRequestPostRequestWithDefaults instantiates a new ApiV1RunnersJobsRequestPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRunnerToken

`func (o *ApiV1RunnersJobsRequestPostRequest) GetRunnerToken() string`

GetRunnerToken returns the RunnerToken field if non-nil, zero value otherwise.

### GetRunnerTokenOk

`func (o *ApiV1RunnersJobsRequestPostRequest) GetRunnerTokenOk() (*string, bool)`

GetRunnerTokenOk returns a tuple with the RunnerToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunnerToken

`func (o *ApiV1RunnersJobsRequestPostRequest) SetRunnerToken(v string)`

SetRunnerToken sets RunnerToken field to given value.


### GetJobTypes

`func (o *ApiV1RunnersJobsRequestPostRequest) GetJobTypes() []string`

GetJobTypes returns the JobTypes field if non-nil, zero value otherwise.

### GetJobTypesOk

`func (o *ApiV1RunnersJobsRequestPostRequest) GetJobTypesOk() (*[]string, bool)`

GetJobTypesOk returns a tuple with the JobTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTypes

`func (o *ApiV1RunnersJobsRequestPostRequest) SetJobTypes(v []string)`

SetJobTypes sets JobTypes field to given value.

### HasJobTypes

`func (o *ApiV1RunnersJobsRequestPostRequest) HasJobTypes() bool`

HasJobTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


