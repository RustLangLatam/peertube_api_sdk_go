# ServerConfigUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VideoQuota** | Pointer to **int32** | In bytes | [optional] 
**VideoQuotaDaily** | Pointer to **int32** | In bytes | [optional] 

## Methods

### NewServerConfigUser

`func NewServerConfigUser() *ServerConfigUser`

NewServerConfigUser instantiates a new ServerConfigUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerConfigUserWithDefaults

`func NewServerConfigUserWithDefaults() *ServerConfigUser`

NewServerConfigUserWithDefaults instantiates a new ServerConfigUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVideoQuota

`func (o *ServerConfigUser) GetVideoQuota() int32`

GetVideoQuota returns the VideoQuota field if non-nil, zero value otherwise.

### GetVideoQuotaOk

`func (o *ServerConfigUser) GetVideoQuotaOk() (*int32, bool)`

GetVideoQuotaOk returns a tuple with the VideoQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoQuota

`func (o *ServerConfigUser) SetVideoQuota(v int32)`

SetVideoQuota sets VideoQuota field to given value.

### HasVideoQuota

`func (o *ServerConfigUser) HasVideoQuota() bool`

HasVideoQuota returns a boolean if a field has been set.

### GetVideoQuotaDaily

`func (o *ServerConfigUser) GetVideoQuotaDaily() int32`

GetVideoQuotaDaily returns the VideoQuotaDaily field if non-nil, zero value otherwise.

### GetVideoQuotaDailyOk

`func (o *ServerConfigUser) GetVideoQuotaDailyOk() (*int32, bool)`

GetVideoQuotaDailyOk returns a tuple with the VideoQuotaDaily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoQuotaDaily

`func (o *ServerConfigUser) SetVideoQuotaDaily(v int32)`

SetVideoQuotaDaily sets VideoQuotaDaily field to given value.

### HasVideoQuotaDaily

`func (o *ServerConfigUser) HasVideoQuotaDaily() bool`

HasVideoQuotaDaily returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


