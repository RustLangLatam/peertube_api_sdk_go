# UserRegistrationState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The registration state (Pending &#x3D; &#x60;1&#x60;, Rejected &#x3D; &#x60;2&#x60;, Accepted &#x3D; &#x60;3&#x60;) | [optional] 
**Label** | Pointer to **string** |  | [optional] 

## Methods

### NewUserRegistrationState

`func NewUserRegistrationState() *UserRegistrationState`

NewUserRegistrationState instantiates a new UserRegistrationState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserRegistrationStateWithDefaults

`func NewUserRegistrationStateWithDefaults() *UserRegistrationState`

NewUserRegistrationStateWithDefaults instantiates a new UserRegistrationState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserRegistrationState) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserRegistrationState) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserRegistrationState) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *UserRegistrationState) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabel

`func (o *UserRegistrationState) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *UserRegistrationState) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *UserRegistrationState) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *UserRegistrationState) HasLabel() bool`

HasLabel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


