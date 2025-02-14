# UserRegistrationAcceptOrReject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ModerationResponse** | **string** | Moderation response to send to the user | 
**PreventEmailDelivery** | Pointer to **bool** | Set it to true if you don&#39;t want PeerTube to send an email to the user | [optional] 

## Methods

### NewUserRegistrationAcceptOrReject

`func NewUserRegistrationAcceptOrReject(moderationResponse string, ) *UserRegistrationAcceptOrReject`

NewUserRegistrationAcceptOrReject instantiates a new UserRegistrationAcceptOrReject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserRegistrationAcceptOrRejectWithDefaults

`func NewUserRegistrationAcceptOrRejectWithDefaults() *UserRegistrationAcceptOrReject`

NewUserRegistrationAcceptOrRejectWithDefaults instantiates a new UserRegistrationAcceptOrReject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModerationResponse

`func (o *UserRegistrationAcceptOrReject) GetModerationResponse() string`

GetModerationResponse returns the ModerationResponse field if non-nil, zero value otherwise.

### GetModerationResponseOk

`func (o *UserRegistrationAcceptOrReject) GetModerationResponseOk() (*string, bool)`

GetModerationResponseOk returns a tuple with the ModerationResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModerationResponse

`func (o *UserRegistrationAcceptOrReject) SetModerationResponse(v string)`

SetModerationResponse sets ModerationResponse field to given value.


### GetPreventEmailDelivery

`func (o *UserRegistrationAcceptOrReject) GetPreventEmailDelivery() bool`

GetPreventEmailDelivery returns the PreventEmailDelivery field if non-nil, zero value otherwise.

### GetPreventEmailDeliveryOk

`func (o *UserRegistrationAcceptOrReject) GetPreventEmailDeliveryOk() (*bool, bool)`

GetPreventEmailDeliveryOk returns a tuple with the PreventEmailDelivery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreventEmailDelivery

`func (o *UserRegistrationAcceptOrReject) SetPreventEmailDelivery(v bool)`

SetPreventEmailDelivery sets PreventEmailDelivery field to given value.

### HasPreventEmailDelivery

`func (o *UserRegistrationAcceptOrReject) HasPreventEmailDelivery() bool`

HasPreventEmailDelivery returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


