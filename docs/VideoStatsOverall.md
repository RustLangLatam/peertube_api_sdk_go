# VideoStatsOverall

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AverageWatchTime** | Pointer to **float32** |  | [optional] 
**TotalWatchTime** | Pointer to **float32** |  | [optional] 
**ViewersPeak** | Pointer to **float32** |  | [optional] 
**ViewersPeakDate** | Pointer to **time.Time** |  | [optional] 
**Countries** | Pointer to [**[]VideoStatsOverallCountriesInner**](VideoStatsOverallCountriesInner.md) |  | [optional] 

## Methods

### NewVideoStatsOverall

`func NewVideoStatsOverall() *VideoStatsOverall`

NewVideoStatsOverall instantiates a new VideoStatsOverall object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVideoStatsOverallWithDefaults

`func NewVideoStatsOverallWithDefaults() *VideoStatsOverall`

NewVideoStatsOverallWithDefaults instantiates a new VideoStatsOverall object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAverageWatchTime

`func (o *VideoStatsOverall) GetAverageWatchTime() float32`

GetAverageWatchTime returns the AverageWatchTime field if non-nil, zero value otherwise.

### GetAverageWatchTimeOk

`func (o *VideoStatsOverall) GetAverageWatchTimeOk() (*float32, bool)`

GetAverageWatchTimeOk returns a tuple with the AverageWatchTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageWatchTime

`func (o *VideoStatsOverall) SetAverageWatchTime(v float32)`

SetAverageWatchTime sets AverageWatchTime field to given value.

### HasAverageWatchTime

`func (o *VideoStatsOverall) HasAverageWatchTime() bool`

HasAverageWatchTime returns a boolean if a field has been set.

### GetTotalWatchTime

`func (o *VideoStatsOverall) GetTotalWatchTime() float32`

GetTotalWatchTime returns the TotalWatchTime field if non-nil, zero value otherwise.

### GetTotalWatchTimeOk

`func (o *VideoStatsOverall) GetTotalWatchTimeOk() (*float32, bool)`

GetTotalWatchTimeOk returns a tuple with the TotalWatchTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalWatchTime

`func (o *VideoStatsOverall) SetTotalWatchTime(v float32)`

SetTotalWatchTime sets TotalWatchTime field to given value.

### HasTotalWatchTime

`func (o *VideoStatsOverall) HasTotalWatchTime() bool`

HasTotalWatchTime returns a boolean if a field has been set.

### GetViewersPeak

`func (o *VideoStatsOverall) GetViewersPeak() float32`

GetViewersPeak returns the ViewersPeak field if non-nil, zero value otherwise.

### GetViewersPeakOk

`func (o *VideoStatsOverall) GetViewersPeakOk() (*float32, bool)`

GetViewersPeakOk returns a tuple with the ViewersPeak field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewersPeak

`func (o *VideoStatsOverall) SetViewersPeak(v float32)`

SetViewersPeak sets ViewersPeak field to given value.

### HasViewersPeak

`func (o *VideoStatsOverall) HasViewersPeak() bool`

HasViewersPeak returns a boolean if a field has been set.

### GetViewersPeakDate

`func (o *VideoStatsOverall) GetViewersPeakDate() time.Time`

GetViewersPeakDate returns the ViewersPeakDate field if non-nil, zero value otherwise.

### GetViewersPeakDateOk

`func (o *VideoStatsOverall) GetViewersPeakDateOk() (*time.Time, bool)`

GetViewersPeakDateOk returns a tuple with the ViewersPeakDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewersPeakDate

`func (o *VideoStatsOverall) SetViewersPeakDate(v time.Time)`

SetViewersPeakDate sets ViewersPeakDate field to given value.

### HasViewersPeakDate

`func (o *VideoStatsOverall) HasViewersPeakDate() bool`

HasViewersPeakDate returns a boolean if a field has been set.

### GetCountries

`func (o *VideoStatsOverall) GetCountries() []VideoStatsOverallCountriesInner`

GetCountries returns the Countries field if non-nil, zero value otherwise.

### GetCountriesOk

`func (o *VideoStatsOverall) GetCountriesOk() (*[]VideoStatsOverallCountriesInner, bool)`

GetCountriesOk returns a tuple with the Countries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountries

`func (o *VideoStatsOverall) SetCountries(v []VideoStatsOverallCountriesInner)`

SetCountries sets Countries field to given value.

### HasCountries

`func (o *VideoStatsOverall) HasCountries() bool`

HasCountries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


