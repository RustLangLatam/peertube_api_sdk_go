# VideoChannelCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | **interface{}** | Channel display name | 
**Name** | **string** | username of the channel to create | 
**Description** | Pointer to **interface{}** | Channel description | [optional] 
**Support** | Pointer to **interface{}** | How to support/fund the channel | [optional] 

## Methods

### NewVideoChannelCreate

`func NewVideoChannelCreate(displayName interface{}, name string, ) *VideoChannelCreate`

NewVideoChannelCreate instantiates a new VideoChannelCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVideoChannelCreateWithDefaults

`func NewVideoChannelCreateWithDefaults() *VideoChannelCreate`

NewVideoChannelCreateWithDefaults instantiates a new VideoChannelCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *VideoChannelCreate) GetDisplayName() interface{}`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *VideoChannelCreate) GetDisplayNameOk() (*interface{}, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *VideoChannelCreate) SetDisplayName(v interface{})`

SetDisplayName sets DisplayName field to given value.


### SetDisplayNameNil

`func (o *VideoChannelCreate) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *VideoChannelCreate) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetName

`func (o *VideoChannelCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VideoChannelCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VideoChannelCreate) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *VideoChannelCreate) GetDescription() interface{}`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VideoChannelCreate) GetDescriptionOk() (*interface{}, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VideoChannelCreate) SetDescription(v interface{})`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VideoChannelCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *VideoChannelCreate) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *VideoChannelCreate) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetSupport

`func (o *VideoChannelCreate) GetSupport() interface{}`

GetSupport returns the Support field if non-nil, zero value otherwise.

### GetSupportOk

`func (o *VideoChannelCreate) GetSupportOk() (*interface{}, bool)`

GetSupportOk returns a tuple with the Support field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupport

`func (o *VideoChannelCreate) SetSupport(v interface{})`

SetSupport sets Support field to given value.

### HasSupport

`func (o *VideoChannelCreate) HasSupport() bool`

HasSupport returns a boolean if a field has been set.

### SetSupportNil

`func (o *VideoChannelCreate) SetSupportNil(b bool)`

 SetSupportNil sets the value for Support to be an explicit nil

### UnsetSupport
`func (o *VideoChannelCreate) UnsetSupport()`

UnsetSupport ensures that no value is present for Support, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


