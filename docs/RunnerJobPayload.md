# RunnerJobPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Input** | Pointer to [**VODAudioMergeTranscoding1Input**](VODAudioMergeTranscoding1Input.md) |  | [optional] 
**Output** | Pointer to [**VODWebVideoTranscoding1Output**](VODWebVideoTranscoding1Output.md) |  | [optional] 

## Methods

### NewRunnerJobPayload

`func NewRunnerJobPayload() *RunnerJobPayload`

NewRunnerJobPayload instantiates a new RunnerJobPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunnerJobPayloadWithDefaults

`func NewRunnerJobPayloadWithDefaults() *RunnerJobPayload`

NewRunnerJobPayloadWithDefaults instantiates a new RunnerJobPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInput

`func (o *RunnerJobPayload) GetInput() VODAudioMergeTranscoding1Input`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *RunnerJobPayload) GetInputOk() (*VODAudioMergeTranscoding1Input, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *RunnerJobPayload) SetInput(v VODAudioMergeTranscoding1Input)`

SetInput sets Input field to given value.

### HasInput

`func (o *RunnerJobPayload) HasInput() bool`

HasInput returns a boolean if a field has been set.

### GetOutput

`func (o *RunnerJobPayload) GetOutput() VODWebVideoTranscoding1Output`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *RunnerJobPayload) GetOutputOk() (*VODWebVideoTranscoding1Output, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *RunnerJobPayload) SetOutput(v VODWebVideoTranscoding1Output)`

SetOutput sets Output field to given value.

### HasOutput

`func (o *RunnerJobPayload) HasOutput() bool`

HasOutput returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


