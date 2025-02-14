# ListUserExports200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**State** | Pointer to [**ListUserExports200ResponseState**](ListUserExports200ResponseState.md) |  | [optional] 
**Size** | Pointer to **int32** | Size of the archive file in bytes | [optional] 
**PrivateDownloadUrl** | Pointer to **string** | This URL already contains the JWT token, so no additional authentication credentials are required | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**ExpiresOn** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewListUserExports200Response

`func NewListUserExports200Response() *ListUserExports200Response`

NewListUserExports200Response instantiates a new ListUserExports200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListUserExports200ResponseWithDefaults

`func NewListUserExports200ResponseWithDefaults() *ListUserExports200Response`

NewListUserExports200ResponseWithDefaults instantiates a new ListUserExports200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListUserExports200Response) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListUserExports200Response) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListUserExports200Response) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ListUserExports200Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetState

`func (o *ListUserExports200Response) GetState() ListUserExports200ResponseState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ListUserExports200Response) GetStateOk() (*ListUserExports200ResponseState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ListUserExports200Response) SetState(v ListUserExports200ResponseState)`

SetState sets State field to given value.

### HasState

`func (o *ListUserExports200Response) HasState() bool`

HasState returns a boolean if a field has been set.

### GetSize

`func (o *ListUserExports200Response) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ListUserExports200Response) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ListUserExports200Response) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *ListUserExports200Response) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetPrivateDownloadUrl

`func (o *ListUserExports200Response) GetPrivateDownloadUrl() string`

GetPrivateDownloadUrl returns the PrivateDownloadUrl field if non-nil, zero value otherwise.

### GetPrivateDownloadUrlOk

`func (o *ListUserExports200Response) GetPrivateDownloadUrlOk() (*string, bool)`

GetPrivateDownloadUrlOk returns a tuple with the PrivateDownloadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateDownloadUrl

`func (o *ListUserExports200Response) SetPrivateDownloadUrl(v string)`

SetPrivateDownloadUrl sets PrivateDownloadUrl field to given value.

### HasPrivateDownloadUrl

`func (o *ListUserExports200Response) HasPrivateDownloadUrl() bool`

HasPrivateDownloadUrl returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ListUserExports200Response) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ListUserExports200Response) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ListUserExports200Response) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ListUserExports200Response) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetExpiresOn

`func (o *ListUserExports200Response) GetExpiresOn() time.Time`

GetExpiresOn returns the ExpiresOn field if non-nil, zero value otherwise.

### GetExpiresOnOk

`func (o *ListUserExports200Response) GetExpiresOnOk() (*time.Time, bool)`

GetExpiresOnOk returns a tuple with the ExpiresOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresOn

`func (o *ListUserExports200Response) SetExpiresOn(v time.Time)`

SetExpiresOn sets ExpiresOn field to given value.

### HasExpiresOn

`func (o *ListUserExports200Response) HasExpiresOn() bool`

HasExpiresOn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


