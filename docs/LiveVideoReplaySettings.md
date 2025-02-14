# LiveVideoReplaySettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Privacy** | Pointer to [**VideoPrivacySet**](VideoPrivacySet.md) |  | [optional] 

## Methods

### NewLiveVideoReplaySettings

`func NewLiveVideoReplaySettings() *LiveVideoReplaySettings`

NewLiveVideoReplaySettings instantiates a new LiveVideoReplaySettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLiveVideoReplaySettingsWithDefaults

`func NewLiveVideoReplaySettingsWithDefaults() *LiveVideoReplaySettings`

NewLiveVideoReplaySettingsWithDefaults instantiates a new LiveVideoReplaySettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrivacy

`func (o *LiveVideoReplaySettings) GetPrivacy() VideoPrivacySet`

GetPrivacy returns the Privacy field if non-nil, zero value otherwise.

### GetPrivacyOk

`func (o *LiveVideoReplaySettings) GetPrivacyOk() (*VideoPrivacySet, bool)`

GetPrivacyOk returns a tuple with the Privacy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacy

`func (o *LiveVideoReplaySettings) SetPrivacy(v VideoPrivacySet)`

SetPrivacy sets Privacy field to given value.

### HasPrivacy

`func (o *LiveVideoReplaySettings) HasPrivacy() bool`

HasPrivacy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


