# VideoChannelUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **interface{}** | Channel display name | [optional] 
**Description** | Pointer to **interface{}** | Channel description | [optional] 
**Support** | Pointer to **interface{}** | How to support/fund the channel | [optional] 
**BulkVideosSupportUpdate** | Pointer to **bool** | Update the support field for all videos of this channel | [optional] 

## Methods

### NewVideoChannelUpdate

`func NewVideoChannelUpdate() *VideoChannelUpdate`

NewVideoChannelUpdate instantiates a new VideoChannelUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVideoChannelUpdateWithDefaults

`func NewVideoChannelUpdateWithDefaults() *VideoChannelUpdate`

NewVideoChannelUpdateWithDefaults instantiates a new VideoChannelUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *VideoChannelUpdate) GetDisplayName() interface{}`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *VideoChannelUpdate) GetDisplayNameOk() (*interface{}, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *VideoChannelUpdate) SetDisplayName(v interface{})`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *VideoChannelUpdate) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *VideoChannelUpdate) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *VideoChannelUpdate) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetDescription

`func (o *VideoChannelUpdate) GetDescription() interface{}`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VideoChannelUpdate) GetDescriptionOk() (*interface{}, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VideoChannelUpdate) SetDescription(v interface{})`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VideoChannelUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *VideoChannelUpdate) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *VideoChannelUpdate) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetSupport

`func (o *VideoChannelUpdate) GetSupport() interface{}`

GetSupport returns the Support field if non-nil, zero value otherwise.

### GetSupportOk

`func (o *VideoChannelUpdate) GetSupportOk() (*interface{}, bool)`

GetSupportOk returns a tuple with the Support field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupport

`func (o *VideoChannelUpdate) SetSupport(v interface{})`

SetSupport sets Support field to given value.

### HasSupport

`func (o *VideoChannelUpdate) HasSupport() bool`

HasSupport returns a boolean if a field has been set.

### SetSupportNil

`func (o *VideoChannelUpdate) SetSupportNil(b bool)`

 SetSupportNil sets the value for Support to be an explicit nil

### UnsetSupport
`func (o *VideoChannelUpdate) UnsetSupport()`

UnsetSupport ensures that no value is present for Support, not even an explicit nil
### GetBulkVideosSupportUpdate

`func (o *VideoChannelUpdate) GetBulkVideosSupportUpdate() bool`

GetBulkVideosSupportUpdate returns the BulkVideosSupportUpdate field if non-nil, zero value otherwise.

### GetBulkVideosSupportUpdateOk

`func (o *VideoChannelUpdate) GetBulkVideosSupportUpdateOk() (*bool, bool)`

GetBulkVideosSupportUpdateOk returns a tuple with the BulkVideosSupportUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBulkVideosSupportUpdate

`func (o *VideoChannelUpdate) SetBulkVideosSupportUpdate(v bool)`

SetBulkVideosSupportUpdate sets BulkVideosSupportUpdate field to given value.

### HasBulkVideosSupportUpdate

`func (o *VideoChannelUpdate) HasBulkVideosSupportUpdate() bool`

HasBulkVideosSupportUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


