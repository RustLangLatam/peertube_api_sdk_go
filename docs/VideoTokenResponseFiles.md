# VideoTokenResponseFiles

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | Pointer to **string** |  | [optional] 
**Expires** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewVideoTokenResponseFiles

`func NewVideoTokenResponseFiles() *VideoTokenResponseFiles`

NewVideoTokenResponseFiles instantiates a new VideoTokenResponseFiles object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVideoTokenResponseFilesWithDefaults

`func NewVideoTokenResponseFilesWithDefaults() *VideoTokenResponseFiles`

NewVideoTokenResponseFilesWithDefaults instantiates a new VideoTokenResponseFiles object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToken

`func (o *VideoTokenResponseFiles) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *VideoTokenResponseFiles) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *VideoTokenResponseFiles) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *VideoTokenResponseFiles) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetExpires

`func (o *VideoTokenResponseFiles) GetExpires() time.Time`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *VideoTokenResponseFiles) GetExpiresOk() (*time.Time, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *VideoTokenResponseFiles) SetExpires(v time.Time)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *VideoTokenResponseFiles) HasExpires() bool`

HasExpires returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


