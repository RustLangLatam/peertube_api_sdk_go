# ApiV1AbusesAbuseIdPutRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to [**AbuseStateSet**](AbuseStateSet.md) |  | [optional] 
**ModerationComment** | Pointer to **string** | Update the report comment visible only to the moderation team | [optional] 

## Methods

### NewApiV1AbusesAbuseIdPutRequest

`func NewApiV1AbusesAbuseIdPutRequest() *ApiV1AbusesAbuseIdPutRequest`

NewApiV1AbusesAbuseIdPutRequest instantiates a new ApiV1AbusesAbuseIdPutRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV1AbusesAbuseIdPutRequestWithDefaults

`func NewApiV1AbusesAbuseIdPutRequestWithDefaults() *ApiV1AbusesAbuseIdPutRequest`

NewApiV1AbusesAbuseIdPutRequestWithDefaults instantiates a new ApiV1AbusesAbuseIdPutRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *ApiV1AbusesAbuseIdPutRequest) GetState() AbuseStateSet`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ApiV1AbusesAbuseIdPutRequest) GetStateOk() (*AbuseStateSet, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ApiV1AbusesAbuseIdPutRequest) SetState(v AbuseStateSet)`

SetState sets State field to given value.

### HasState

`func (o *ApiV1AbusesAbuseIdPutRequest) HasState() bool`

HasState returns a boolean if a field has been set.

### GetModerationComment

`func (o *ApiV1AbusesAbuseIdPutRequest) GetModerationComment() string`

GetModerationComment returns the ModerationComment field if non-nil, zero value otherwise.

### GetModerationCommentOk

`func (o *ApiV1AbusesAbuseIdPutRequest) GetModerationCommentOk() (*string, bool)`

GetModerationCommentOk returns a tuple with the ModerationComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModerationComment

`func (o *ApiV1AbusesAbuseIdPutRequest) SetModerationComment(v string)`

SetModerationComment sets ModerationComment field to given value.

### HasModerationComment

`func (o *ApiV1AbusesAbuseIdPutRequest) HasModerationComment() bool`

HasModerationComment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


