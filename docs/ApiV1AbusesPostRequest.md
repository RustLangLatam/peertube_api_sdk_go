# ApiV1AbusesPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reason** | **string** | Reason why the user reports this video | 
**PredefinedReasons** | Pointer to **[]string** | Reason categories that help triage reports | [optional] 
**Video** | Pointer to [**ApiV1AbusesPostRequestVideo**](ApiV1AbusesPostRequestVideo.md) |  | [optional] 
**Comment** | Pointer to [**ApiV1AbusesPostRequestComment**](ApiV1AbusesPostRequestComment.md) |  | [optional] 
**Account** | Pointer to [**ApiV1AbusesPostRequestAccount**](ApiV1AbusesPostRequestAccount.md) |  | [optional] 

## Methods

### NewApiV1AbusesPostRequest

`func NewApiV1AbusesPostRequest(reason string, ) *ApiV1AbusesPostRequest`

NewApiV1AbusesPostRequest instantiates a new ApiV1AbusesPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV1AbusesPostRequestWithDefaults

`func NewApiV1AbusesPostRequestWithDefaults() *ApiV1AbusesPostRequest`

NewApiV1AbusesPostRequestWithDefaults instantiates a new ApiV1AbusesPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReason

`func (o *ApiV1AbusesPostRequest) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ApiV1AbusesPostRequest) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ApiV1AbusesPostRequest) SetReason(v string)`

SetReason sets Reason field to given value.


### GetPredefinedReasons

`func (o *ApiV1AbusesPostRequest) GetPredefinedReasons() []string`

GetPredefinedReasons returns the PredefinedReasons field if non-nil, zero value otherwise.

### GetPredefinedReasonsOk

`func (o *ApiV1AbusesPostRequest) GetPredefinedReasonsOk() (*[]string, bool)`

GetPredefinedReasonsOk returns a tuple with the PredefinedReasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredefinedReasons

`func (o *ApiV1AbusesPostRequest) SetPredefinedReasons(v []string)`

SetPredefinedReasons sets PredefinedReasons field to given value.

### HasPredefinedReasons

`func (o *ApiV1AbusesPostRequest) HasPredefinedReasons() bool`

HasPredefinedReasons returns a boolean if a field has been set.

### GetVideo

`func (o *ApiV1AbusesPostRequest) GetVideo() ApiV1AbusesPostRequestVideo`

GetVideo returns the Video field if non-nil, zero value otherwise.

### GetVideoOk

`func (o *ApiV1AbusesPostRequest) GetVideoOk() (*ApiV1AbusesPostRequestVideo, bool)`

GetVideoOk returns a tuple with the Video field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideo

`func (o *ApiV1AbusesPostRequest) SetVideo(v ApiV1AbusesPostRequestVideo)`

SetVideo sets Video field to given value.

### HasVideo

`func (o *ApiV1AbusesPostRequest) HasVideo() bool`

HasVideo returns a boolean if a field has been set.

### GetComment

`func (o *ApiV1AbusesPostRequest) GetComment() ApiV1AbusesPostRequestComment`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ApiV1AbusesPostRequest) GetCommentOk() (*ApiV1AbusesPostRequestComment, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ApiV1AbusesPostRequest) SetComment(v ApiV1AbusesPostRequestComment)`

SetComment sets Comment field to given value.

### HasComment

`func (o *ApiV1AbusesPostRequest) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetAccount

`func (o *ApiV1AbusesPostRequest) GetAccount() ApiV1AbusesPostRequestAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ApiV1AbusesPostRequest) GetAccountOk() (*ApiV1AbusesPostRequestAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ApiV1AbusesPostRequest) SetAccount(v ApiV1AbusesPostRequestAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ApiV1AbusesPostRequest) HasAccount() bool`

HasAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


