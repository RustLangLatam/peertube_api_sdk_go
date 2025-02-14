# SendClientLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** |  | 
**Url** | **string** | URL of the current user page | 
**Level** | **string** |  | 
**StackTrace** | Pointer to **string** | Stack trace of the error if there is one | [optional] 
**UserAgent** | Pointer to **string** | User agent of the web browser that sends the message | [optional] 
**Meta** | Pointer to **string** | Additional information regarding this log | [optional] 

## Methods

### NewSendClientLog

`func NewSendClientLog(message string, url string, level string, ) *SendClientLog`

NewSendClientLog instantiates a new SendClientLog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendClientLogWithDefaults

`func NewSendClientLogWithDefaults() *SendClientLog`

NewSendClientLogWithDefaults instantiates a new SendClientLog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *SendClientLog) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SendClientLog) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SendClientLog) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetUrl

`func (o *SendClientLog) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SendClientLog) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SendClientLog) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetLevel

`func (o *SendClientLog) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *SendClientLog) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *SendClientLog) SetLevel(v string)`

SetLevel sets Level field to given value.


### GetStackTrace

`func (o *SendClientLog) GetStackTrace() string`

GetStackTrace returns the StackTrace field if non-nil, zero value otherwise.

### GetStackTraceOk

`func (o *SendClientLog) GetStackTraceOk() (*string, bool)`

GetStackTraceOk returns a tuple with the StackTrace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackTrace

`func (o *SendClientLog) SetStackTrace(v string)`

SetStackTrace sets StackTrace field to given value.

### HasStackTrace

`func (o *SendClientLog) HasStackTrace() bool`

HasStackTrace returns a boolean if a field has been set.

### GetUserAgent

`func (o *SendClientLog) GetUserAgent() string`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *SendClientLog) GetUserAgentOk() (*string, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *SendClientLog) SetUserAgent(v string)`

SetUserAgent sets UserAgent field to given value.

### HasUserAgent

`func (o *SendClientLog) HasUserAgent() bool`

HasUserAgent returns a boolean if a field has been set.

### GetMeta

`func (o *SendClientLog) GetMeta() string`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *SendClientLog) GetMetaOk() (*string, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *SendClientLog) SetMeta(v string)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *SendClientLog) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


