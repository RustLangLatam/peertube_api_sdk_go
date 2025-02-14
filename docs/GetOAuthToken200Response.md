# GetOAuthToken200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TokenType** | Pointer to **string** |  | [optional] 
**AccessToken** | Pointer to **string** | valid for 1 day | [optional] 
**RefreshToken** | Pointer to **string** | valid for 2 weeks | [optional] 
**ExpiresIn** | Pointer to **int32** |  | [optional] 
**RefreshTokenExpiresIn** | Pointer to **int32** |  | [optional] 

## Methods

### NewGetOAuthToken200Response

`func NewGetOAuthToken200Response() *GetOAuthToken200Response`

NewGetOAuthToken200Response instantiates a new GetOAuthToken200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOAuthToken200ResponseWithDefaults

`func NewGetOAuthToken200ResponseWithDefaults() *GetOAuthToken200Response`

NewGetOAuthToken200ResponseWithDefaults instantiates a new GetOAuthToken200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTokenType

`func (o *GetOAuthToken200Response) GetTokenType() string`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *GetOAuthToken200Response) GetTokenTypeOk() (*string, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *GetOAuthToken200Response) SetTokenType(v string)`

SetTokenType sets TokenType field to given value.

### HasTokenType

`func (o *GetOAuthToken200Response) HasTokenType() bool`

HasTokenType returns a boolean if a field has been set.

### GetAccessToken

`func (o *GetOAuthToken200Response) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *GetOAuthToken200Response) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *GetOAuthToken200Response) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *GetOAuthToken200Response) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### GetRefreshToken

`func (o *GetOAuthToken200Response) GetRefreshToken() string`

GetRefreshToken returns the RefreshToken field if non-nil, zero value otherwise.

### GetRefreshTokenOk

`func (o *GetOAuthToken200Response) GetRefreshTokenOk() (*string, bool)`

GetRefreshTokenOk returns a tuple with the RefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshToken

`func (o *GetOAuthToken200Response) SetRefreshToken(v string)`

SetRefreshToken sets RefreshToken field to given value.

### HasRefreshToken

`func (o *GetOAuthToken200Response) HasRefreshToken() bool`

HasRefreshToken returns a boolean if a field has been set.

### GetExpiresIn

`func (o *GetOAuthToken200Response) GetExpiresIn() int32`

GetExpiresIn returns the ExpiresIn field if non-nil, zero value otherwise.

### GetExpiresInOk

`func (o *GetOAuthToken200Response) GetExpiresInOk() (*int32, bool)`

GetExpiresInOk returns a tuple with the ExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresIn

`func (o *GetOAuthToken200Response) SetExpiresIn(v int32)`

SetExpiresIn sets ExpiresIn field to given value.

### HasExpiresIn

`func (o *GetOAuthToken200Response) HasExpiresIn() bool`

HasExpiresIn returns a boolean if a field has been set.

### GetRefreshTokenExpiresIn

`func (o *GetOAuthToken200Response) GetRefreshTokenExpiresIn() int32`

GetRefreshTokenExpiresIn returns the RefreshTokenExpiresIn field if non-nil, zero value otherwise.

### GetRefreshTokenExpiresInOk

`func (o *GetOAuthToken200Response) GetRefreshTokenExpiresInOk() (*int32, bool)`

GetRefreshTokenExpiresInOk returns a tuple with the RefreshTokenExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshTokenExpiresIn

`func (o *GetOAuthToken200Response) SetRefreshTokenExpiresIn(v int32)`

SetRefreshTokenExpiresIn sets RefreshTokenExpiresIn field to given value.

### HasRefreshTokenExpiresIn

`func (o *GetOAuthToken200Response) HasRefreshTokenExpiresIn() bool`

HasRefreshTokenExpiresIn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


