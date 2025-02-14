# ApiV1RunnersJobsJobUUIDAbortPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RunnerToken** | **string** |  | 
**JobToken** | **string** |  | 
**Reason** | **string** | Why the runner aborts this job | 

## Methods

### NewApiV1RunnersJobsJobUUIDAbortPostRequest

`func NewApiV1RunnersJobsJobUUIDAbortPostRequest(runnerToken string, jobToken string, reason string, ) *ApiV1RunnersJobsJobUUIDAbortPostRequest`

NewApiV1RunnersJobsJobUUIDAbortPostRequest instantiates a new ApiV1RunnersJobsJobUUIDAbortPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV1RunnersJobsJobUUIDAbortPostRequestWithDefaults

`func NewApiV1RunnersJobsJobUUIDAbortPostRequestWithDefaults() *ApiV1RunnersJobsJobUUIDAbortPostRequest`

NewApiV1RunnersJobsJobUUIDAbortPostRequestWithDefaults instantiates a new ApiV1RunnersJobsJobUUIDAbortPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRunnerToken

`func (o *ApiV1RunnersJobsJobUUIDAbortPostRequest) GetRunnerToken() string`

GetRunnerToken returns the RunnerToken field if non-nil, zero value otherwise.

### GetRunnerTokenOk

`func (o *ApiV1RunnersJobsJobUUIDAbortPostRequest) GetRunnerTokenOk() (*string, bool)`

GetRunnerTokenOk returns a tuple with the RunnerToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunnerToken

`func (o *ApiV1RunnersJobsJobUUIDAbortPostRequest) SetRunnerToken(v string)`

SetRunnerToken sets RunnerToken field to given value.


### GetJobToken

`func (o *ApiV1RunnersJobsJobUUIDAbortPostRequest) GetJobToken() string`

GetJobToken returns the JobToken field if non-nil, zero value otherwise.

### GetJobTokenOk

`func (o *ApiV1RunnersJobsJobUUIDAbortPostRequest) GetJobTokenOk() (*string, bool)`

GetJobTokenOk returns a tuple with the JobToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobToken

`func (o *ApiV1RunnersJobsJobUUIDAbortPostRequest) SetJobToken(v string)`

SetJobToken sets JobToken field to given value.


### GetReason

`func (o *ApiV1RunnersJobsJobUUIDAbortPostRequest) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ApiV1RunnersJobsJobUUIDAbortPostRequest) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ApiV1RunnersJobsJobUUIDAbortPostRequest) SetReason(v string)`

SetReason sets Reason field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


