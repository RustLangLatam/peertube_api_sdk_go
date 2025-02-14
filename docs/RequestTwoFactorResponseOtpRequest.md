# RequestTwoFactorResponseOtpRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestToken** | Pointer to **string** | The token to send to confirm this request | [optional] 
**Secret** | Pointer to **string** | The OTP secret | [optional] 
**Uri** | Pointer to **string** | The OTP URI | [optional] 

## Methods

### NewRequestTwoFactorResponseOtpRequest

`func NewRequestTwoFactorResponseOtpRequest() *RequestTwoFactorResponseOtpRequest`

NewRequestTwoFactorResponseOtpRequest instantiates a new RequestTwoFactorResponseOtpRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestTwoFactorResponseOtpRequestWithDefaults

`func NewRequestTwoFactorResponseOtpRequestWithDefaults() *RequestTwoFactorResponseOtpRequest`

NewRequestTwoFactorResponseOtpRequestWithDefaults instantiates a new RequestTwoFactorResponseOtpRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestToken

`func (o *RequestTwoFactorResponseOtpRequest) GetRequestToken() string`

GetRequestToken returns the RequestToken field if non-nil, zero value otherwise.

### GetRequestTokenOk

`func (o *RequestTwoFactorResponseOtpRequest) GetRequestTokenOk() (*string, bool)`

GetRequestTokenOk returns a tuple with the RequestToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestToken

`func (o *RequestTwoFactorResponseOtpRequest) SetRequestToken(v string)`

SetRequestToken sets RequestToken field to given value.

### HasRequestToken

`func (o *RequestTwoFactorResponseOtpRequest) HasRequestToken() bool`

HasRequestToken returns a boolean if a field has been set.

### GetSecret

`func (o *RequestTwoFactorResponseOtpRequest) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *RequestTwoFactorResponseOtpRequest) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *RequestTwoFactorResponseOtpRequest) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *RequestTwoFactorResponseOtpRequest) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetUri

`func (o *RequestTwoFactorResponseOtpRequest) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *RequestTwoFactorResponseOtpRequest) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *RequestTwoFactorResponseOtpRequest) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *RequestTwoFactorResponseOtpRequest) HasUri() bool`

HasUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


