# Abuse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Reason** | Pointer to **string** |  | [optional] 
**PredefinedReasons** | Pointer to **[]string** |  | [optional] 
**ReporterAccount** | Pointer to [**Account**](Account.md) |  | [optional] 
**State** | Pointer to [**AbuseStateConstant**](AbuseStateConstant.md) |  | [optional] 
**ModerationComment** | Pointer to **string** |  | [optional] 
**Video** | Pointer to [**VideoInfo**](VideoInfo.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewAbuse

`func NewAbuse() *Abuse`

NewAbuse instantiates a new Abuse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAbuseWithDefaults

`func NewAbuseWithDefaults() *Abuse`

NewAbuseWithDefaults instantiates a new Abuse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Abuse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Abuse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Abuse) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Abuse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetReason

`func (o *Abuse) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *Abuse) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *Abuse) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *Abuse) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetPredefinedReasons

`func (o *Abuse) GetPredefinedReasons() []string`

GetPredefinedReasons returns the PredefinedReasons field if non-nil, zero value otherwise.

### GetPredefinedReasonsOk

`func (o *Abuse) GetPredefinedReasonsOk() (*[]string, bool)`

GetPredefinedReasonsOk returns a tuple with the PredefinedReasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredefinedReasons

`func (o *Abuse) SetPredefinedReasons(v []string)`

SetPredefinedReasons sets PredefinedReasons field to given value.

### HasPredefinedReasons

`func (o *Abuse) HasPredefinedReasons() bool`

HasPredefinedReasons returns a boolean if a field has been set.

### GetReporterAccount

`func (o *Abuse) GetReporterAccount() Account`

GetReporterAccount returns the ReporterAccount field if non-nil, zero value otherwise.

### GetReporterAccountOk

`func (o *Abuse) GetReporterAccountOk() (*Account, bool)`

GetReporterAccountOk returns a tuple with the ReporterAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReporterAccount

`func (o *Abuse) SetReporterAccount(v Account)`

SetReporterAccount sets ReporterAccount field to given value.

### HasReporterAccount

`func (o *Abuse) HasReporterAccount() bool`

HasReporterAccount returns a boolean if a field has been set.

### GetState

`func (o *Abuse) GetState() AbuseStateConstant`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Abuse) GetStateOk() (*AbuseStateConstant, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Abuse) SetState(v AbuseStateConstant)`

SetState sets State field to given value.

### HasState

`func (o *Abuse) HasState() bool`

HasState returns a boolean if a field has been set.

### GetModerationComment

`func (o *Abuse) GetModerationComment() string`

GetModerationComment returns the ModerationComment field if non-nil, zero value otherwise.

### GetModerationCommentOk

`func (o *Abuse) GetModerationCommentOk() (*string, bool)`

GetModerationCommentOk returns a tuple with the ModerationComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModerationComment

`func (o *Abuse) SetModerationComment(v string)`

SetModerationComment sets ModerationComment field to given value.

### HasModerationComment

`func (o *Abuse) HasModerationComment() bool`

HasModerationComment returns a boolean if a field has been set.

### GetVideo

`func (o *Abuse) GetVideo() VideoInfo`

GetVideo returns the Video field if non-nil, zero value otherwise.

### GetVideoOk

`func (o *Abuse) GetVideoOk() (*VideoInfo, bool)`

GetVideoOk returns a tuple with the Video field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideo

`func (o *Abuse) SetVideo(v VideoInfo)`

SetVideo sets Video field to given value.

### HasVideo

`func (o *Abuse) HasVideo() bool`

HasVideo returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Abuse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Abuse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Abuse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Abuse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


