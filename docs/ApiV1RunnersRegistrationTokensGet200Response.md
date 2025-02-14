# ApiV1RunnersRegistrationTokensGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int32** |  | [optional] 
**Data** | Pointer to [**[]RunnerRegistrationToken**](RunnerRegistrationToken.md) |  | [optional] 

## Methods

### NewApiV1RunnersRegistrationTokensGet200Response

`func NewApiV1RunnersRegistrationTokensGet200Response() *ApiV1RunnersRegistrationTokensGet200Response`

NewApiV1RunnersRegistrationTokensGet200Response instantiates a new ApiV1RunnersRegistrationTokensGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV1RunnersRegistrationTokensGet200ResponseWithDefaults

`func NewApiV1RunnersRegistrationTokensGet200ResponseWithDefaults() *ApiV1RunnersRegistrationTokensGet200Response`

NewApiV1RunnersRegistrationTokensGet200ResponseWithDefaults instantiates a new ApiV1RunnersRegistrationTokensGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *ApiV1RunnersRegistrationTokensGet200Response) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ApiV1RunnersRegistrationTokensGet200Response) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ApiV1RunnersRegistrationTokensGet200Response) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *ApiV1RunnersRegistrationTokensGet200Response) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetData

`func (o *ApiV1RunnersRegistrationTokensGet200Response) GetData() []RunnerRegistrationToken`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ApiV1RunnersRegistrationTokensGet200Response) GetDataOk() (*[]RunnerRegistrationToken, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ApiV1RunnersRegistrationTokensGet200Response) SetData(v []RunnerRegistrationToken)`

SetData sets Data field to given value.

### HasData

`func (o *ApiV1RunnersRegistrationTokensGet200Response) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


