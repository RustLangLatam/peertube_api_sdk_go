# ServerConfigVideo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Image** | Pointer to [**ServerConfigVideoImage**](ServerConfigVideoImage.md) |  | [optional] 
**File** | Pointer to [**ServerConfigVideoFile**](ServerConfigVideoFile.md) |  | [optional] 

## Methods

### NewServerConfigVideo

`func NewServerConfigVideo() *ServerConfigVideo`

NewServerConfigVideo instantiates a new ServerConfigVideo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerConfigVideoWithDefaults

`func NewServerConfigVideoWithDefaults() *ServerConfigVideo`

NewServerConfigVideoWithDefaults instantiates a new ServerConfigVideo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImage

`func (o *ServerConfigVideo) GetImage() ServerConfigVideoImage`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *ServerConfigVideo) GetImageOk() (*ServerConfigVideoImage, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *ServerConfigVideo) SetImage(v ServerConfigVideoImage)`

SetImage sets Image field to given value.

### HasImage

`func (o *ServerConfigVideo) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetFile

`func (o *ServerConfigVideo) GetFile() ServerConfigVideoFile`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *ServerConfigVideo) GetFileOk() (*ServerConfigVideoFile, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *ServerConfigVideo) SetFile(v ServerConfigVideoFile)`

SetFile sets File field to given value.

### HasFile

`func (o *ServerConfigVideo) HasFile() bool`

HasFile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


