# CategoryOverview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | Pointer to [**VideoConstantNumberCategory**](VideoConstantNumberCategory.md) |  | [optional] 
**Videos** | Pointer to [**[]Video**](Video.md) |  | [optional] 

## Methods

### NewCategoryOverview

`func NewCategoryOverview() *CategoryOverview`

NewCategoryOverview instantiates a new CategoryOverview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCategoryOverviewWithDefaults

`func NewCategoryOverviewWithDefaults() *CategoryOverview`

NewCategoryOverviewWithDefaults instantiates a new CategoryOverview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *CategoryOverview) GetCategory() VideoConstantNumberCategory`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *CategoryOverview) GetCategoryOk() (*VideoConstantNumberCategory, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *CategoryOverview) SetCategory(v VideoConstantNumberCategory)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *CategoryOverview) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetVideos

`func (o *CategoryOverview) GetVideos() []Video`

GetVideos returns the Videos field if non-nil, zero value otherwise.

### GetVideosOk

`func (o *CategoryOverview) GetVideosOk() (*[]Video, bool)`

GetVideosOk returns a tuple with the Videos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideos

`func (o *CategoryOverview) SetVideos(v []Video)`

SetVideos sets Videos field to given value.

### HasVideos

`func (o *CategoryOverview) HasVideos() bool`

HasVideos returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


