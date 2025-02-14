# AbuseMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**ByModerator** | Pointer to **bool** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Account** | Pointer to [**AccountSummary**](AccountSummary.md) |  | [optional] 

## Methods

### NewAbuseMessage

`func NewAbuseMessage() *AbuseMessage`

NewAbuseMessage instantiates a new AbuseMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAbuseMessageWithDefaults

`func NewAbuseMessageWithDefaults() *AbuseMessage`

NewAbuseMessageWithDefaults instantiates a new AbuseMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AbuseMessage) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AbuseMessage) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AbuseMessage) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AbuseMessage) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMessage

`func (o *AbuseMessage) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AbuseMessage) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AbuseMessage) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *AbuseMessage) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetByModerator

`func (o *AbuseMessage) GetByModerator() bool`

GetByModerator returns the ByModerator field if non-nil, zero value otherwise.

### GetByModeratorOk

`func (o *AbuseMessage) GetByModeratorOk() (*bool, bool)`

GetByModeratorOk returns a tuple with the ByModerator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByModerator

`func (o *AbuseMessage) SetByModerator(v bool)`

SetByModerator sets ByModerator field to given value.

### HasByModerator

`func (o *AbuseMessage) HasByModerator() bool`

HasByModerator returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AbuseMessage) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AbuseMessage) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AbuseMessage) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AbuseMessage) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetAccount

`func (o *AbuseMessage) GetAccount() AccountSummary`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *AbuseMessage) GetAccountOk() (*AccountSummary, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *AbuseMessage) SetAccount(v AccountSummary)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *AbuseMessage) HasAccount() bool`

HasAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


