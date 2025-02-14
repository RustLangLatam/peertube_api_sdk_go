# VideoScheduledUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Privacy** | Pointer to [**VideoPrivacySet**](VideoPrivacySet.md) |  | [optional] 
**UpdateAt** | **time.Time** | When to update the video | 

## Methods

### NewVideoScheduledUpdate

`func NewVideoScheduledUpdate(updateAt time.Time, ) *VideoScheduledUpdate`

NewVideoScheduledUpdate instantiates a new VideoScheduledUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVideoScheduledUpdateWithDefaults

`func NewVideoScheduledUpdateWithDefaults() *VideoScheduledUpdate`

NewVideoScheduledUpdateWithDefaults instantiates a new VideoScheduledUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrivacy

`func (o *VideoScheduledUpdate) GetPrivacy() VideoPrivacySet`

GetPrivacy returns the Privacy field if non-nil, zero value otherwise.

### GetPrivacyOk

`func (o *VideoScheduledUpdate) GetPrivacyOk() (*VideoPrivacySet, bool)`

GetPrivacyOk returns a tuple with the Privacy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacy

`func (o *VideoScheduledUpdate) SetPrivacy(v VideoPrivacySet)`

SetPrivacy sets Privacy field to given value.

### HasPrivacy

`func (o *VideoScheduledUpdate) HasPrivacy() bool`

HasPrivacy returns a boolean if a field has been set.

### GetUpdateAt

`func (o *VideoScheduledUpdate) GetUpdateAt() time.Time`

GetUpdateAt returns the UpdateAt field if non-nil, zero value otherwise.

### GetUpdateAtOk

`func (o *VideoScheduledUpdate) GetUpdateAtOk() (*time.Time, bool)`

GetUpdateAtOk returns a tuple with the UpdateAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateAt

`func (o *VideoScheduledUpdate) SetUpdateAt(v time.Time)`

SetUpdateAt sets UpdateAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


