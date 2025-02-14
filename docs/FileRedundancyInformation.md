# FileRedundancyInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**FileUrl** | Pointer to **string** |  | [optional] 
**Strategy** | Pointer to **string** |  | [optional] 
**Size** | Pointer to **int32** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**ExpiresOn** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewFileRedundancyInformation

`func NewFileRedundancyInformation() *FileRedundancyInformation`

NewFileRedundancyInformation instantiates a new FileRedundancyInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileRedundancyInformationWithDefaults

`func NewFileRedundancyInformationWithDefaults() *FileRedundancyInformation`

NewFileRedundancyInformationWithDefaults instantiates a new FileRedundancyInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FileRedundancyInformation) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FileRedundancyInformation) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FileRedundancyInformation) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *FileRedundancyInformation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetFileUrl

`func (o *FileRedundancyInformation) GetFileUrl() string`

GetFileUrl returns the FileUrl field if non-nil, zero value otherwise.

### GetFileUrlOk

`func (o *FileRedundancyInformation) GetFileUrlOk() (*string, bool)`

GetFileUrlOk returns a tuple with the FileUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileUrl

`func (o *FileRedundancyInformation) SetFileUrl(v string)`

SetFileUrl sets FileUrl field to given value.

### HasFileUrl

`func (o *FileRedundancyInformation) HasFileUrl() bool`

HasFileUrl returns a boolean if a field has been set.

### GetStrategy

`func (o *FileRedundancyInformation) GetStrategy() string`

GetStrategy returns the Strategy field if non-nil, zero value otherwise.

### GetStrategyOk

`func (o *FileRedundancyInformation) GetStrategyOk() (*string, bool)`

GetStrategyOk returns a tuple with the Strategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategy

`func (o *FileRedundancyInformation) SetStrategy(v string)`

SetStrategy sets Strategy field to given value.

### HasStrategy

`func (o *FileRedundancyInformation) HasStrategy() bool`

HasStrategy returns a boolean if a field has been set.

### GetSize

`func (o *FileRedundancyInformation) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *FileRedundancyInformation) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *FileRedundancyInformation) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *FileRedundancyInformation) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetCreatedAt

`func (o *FileRedundancyInformation) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FileRedundancyInformation) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FileRedundancyInformation) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *FileRedundancyInformation) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *FileRedundancyInformation) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *FileRedundancyInformation) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *FileRedundancyInformation) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *FileRedundancyInformation) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetExpiresOn

`func (o *FileRedundancyInformation) GetExpiresOn() time.Time`

GetExpiresOn returns the ExpiresOn field if non-nil, zero value otherwise.

### GetExpiresOnOk

`func (o *FileRedundancyInformation) GetExpiresOnOk() (*time.Time, bool)`

GetExpiresOnOk returns a tuple with the ExpiresOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresOn

`func (o *FileRedundancyInformation) SetExpiresOn(v time.Time)`

SetExpiresOn sets ExpiresOn field to given value.

### HasExpiresOn

`func (o *FileRedundancyInformation) HasExpiresOn() bool`

HasExpiresOn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


