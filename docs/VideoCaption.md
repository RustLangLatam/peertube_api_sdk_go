# VideoCaption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Language** | Pointer to [**VideoConstantStringLanguage**](VideoConstantStringLanguage.md) |  | [optional] 
**CaptionPath** | Pointer to **string** |  | [optional] 

## Methods

### NewVideoCaption

`func NewVideoCaption() *VideoCaption`

NewVideoCaption instantiates a new VideoCaption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVideoCaptionWithDefaults

`func NewVideoCaptionWithDefaults() *VideoCaption`

NewVideoCaptionWithDefaults instantiates a new VideoCaption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLanguage

`func (o *VideoCaption) GetLanguage() VideoConstantStringLanguage`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *VideoCaption) GetLanguageOk() (*VideoConstantStringLanguage, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *VideoCaption) SetLanguage(v VideoConstantStringLanguage)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *VideoCaption) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetCaptionPath

`func (o *VideoCaption) GetCaptionPath() string`

GetCaptionPath returns the CaptionPath field if non-nil, zero value otherwise.

### GetCaptionPathOk

`func (o *VideoCaption) GetCaptionPathOk() (*string, bool)`

GetCaptionPathOk returns a tuple with the CaptionPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptionPath

`func (o *VideoCaption) SetCaptionPath(v string)`

SetCaptionPath sets CaptionPath field to given value.

### HasCaptionPath

`func (o *VideoCaption) HasCaptionPath() bool`

HasCaptionPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


