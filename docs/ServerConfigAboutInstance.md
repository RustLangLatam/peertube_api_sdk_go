# ServerConfigAboutInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**ShortDescription** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Terms** | Pointer to **string** |  | [optional] 
**CodeOfConduct** | Pointer to **string** |  | [optional] 
**HardwareInformation** | Pointer to **string** |  | [optional] 
**CreationReason** | Pointer to **string** |  | [optional] 
**ModerationInformation** | Pointer to **string** |  | [optional] 
**Administrator** | Pointer to **string** |  | [optional] 
**MaintenanceLifetime** | Pointer to **string** |  | [optional] 
**BusinessModel** | Pointer to **string** |  | [optional] 
**Languages** | Pointer to **[]string** |  | [optional] 
**Categories** | Pointer to **[]int32** |  | [optional] 
**Avatars** | Pointer to [**[]ActorImage**](ActorImage.md) |  | [optional] 
**Banners** | Pointer to [**[]ActorImage**](ActorImage.md) |  | [optional] 

## Methods

### NewServerConfigAboutInstance

`func NewServerConfigAboutInstance() *ServerConfigAboutInstance`

NewServerConfigAboutInstance instantiates a new ServerConfigAboutInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerConfigAboutInstanceWithDefaults

`func NewServerConfigAboutInstanceWithDefaults() *ServerConfigAboutInstance`

NewServerConfigAboutInstanceWithDefaults instantiates a new ServerConfigAboutInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ServerConfigAboutInstance) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServerConfigAboutInstance) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServerConfigAboutInstance) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ServerConfigAboutInstance) HasName() bool`

HasName returns a boolean if a field has been set.

### GetShortDescription

`func (o *ServerConfigAboutInstance) GetShortDescription() string`

GetShortDescription returns the ShortDescription field if non-nil, zero value otherwise.

### GetShortDescriptionOk

`func (o *ServerConfigAboutInstance) GetShortDescriptionOk() (*string, bool)`

GetShortDescriptionOk returns a tuple with the ShortDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortDescription

`func (o *ServerConfigAboutInstance) SetShortDescription(v string)`

SetShortDescription sets ShortDescription field to given value.

### HasShortDescription

`func (o *ServerConfigAboutInstance) HasShortDescription() bool`

HasShortDescription returns a boolean if a field has been set.

### GetDescription

`func (o *ServerConfigAboutInstance) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ServerConfigAboutInstance) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ServerConfigAboutInstance) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ServerConfigAboutInstance) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTerms

`func (o *ServerConfigAboutInstance) GetTerms() string`

GetTerms returns the Terms field if non-nil, zero value otherwise.

### GetTermsOk

`func (o *ServerConfigAboutInstance) GetTermsOk() (*string, bool)`

GetTermsOk returns a tuple with the Terms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerms

`func (o *ServerConfigAboutInstance) SetTerms(v string)`

SetTerms sets Terms field to given value.

### HasTerms

`func (o *ServerConfigAboutInstance) HasTerms() bool`

HasTerms returns a boolean if a field has been set.

### GetCodeOfConduct

`func (o *ServerConfigAboutInstance) GetCodeOfConduct() string`

GetCodeOfConduct returns the CodeOfConduct field if non-nil, zero value otherwise.

### GetCodeOfConductOk

`func (o *ServerConfigAboutInstance) GetCodeOfConductOk() (*string, bool)`

GetCodeOfConductOk returns a tuple with the CodeOfConduct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeOfConduct

`func (o *ServerConfigAboutInstance) SetCodeOfConduct(v string)`

SetCodeOfConduct sets CodeOfConduct field to given value.

### HasCodeOfConduct

`func (o *ServerConfigAboutInstance) HasCodeOfConduct() bool`

HasCodeOfConduct returns a boolean if a field has been set.

### GetHardwareInformation

`func (o *ServerConfigAboutInstance) GetHardwareInformation() string`

GetHardwareInformation returns the HardwareInformation field if non-nil, zero value otherwise.

### GetHardwareInformationOk

`func (o *ServerConfigAboutInstance) GetHardwareInformationOk() (*string, bool)`

GetHardwareInformationOk returns a tuple with the HardwareInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardwareInformation

`func (o *ServerConfigAboutInstance) SetHardwareInformation(v string)`

SetHardwareInformation sets HardwareInformation field to given value.

### HasHardwareInformation

`func (o *ServerConfigAboutInstance) HasHardwareInformation() bool`

HasHardwareInformation returns a boolean if a field has been set.

### GetCreationReason

`func (o *ServerConfigAboutInstance) GetCreationReason() string`

GetCreationReason returns the CreationReason field if non-nil, zero value otherwise.

### GetCreationReasonOk

`func (o *ServerConfigAboutInstance) GetCreationReasonOk() (*string, bool)`

GetCreationReasonOk returns a tuple with the CreationReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationReason

`func (o *ServerConfigAboutInstance) SetCreationReason(v string)`

SetCreationReason sets CreationReason field to given value.

### HasCreationReason

`func (o *ServerConfigAboutInstance) HasCreationReason() bool`

HasCreationReason returns a boolean if a field has been set.

### GetModerationInformation

`func (o *ServerConfigAboutInstance) GetModerationInformation() string`

GetModerationInformation returns the ModerationInformation field if non-nil, zero value otherwise.

### GetModerationInformationOk

`func (o *ServerConfigAboutInstance) GetModerationInformationOk() (*string, bool)`

GetModerationInformationOk returns a tuple with the ModerationInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModerationInformation

`func (o *ServerConfigAboutInstance) SetModerationInformation(v string)`

SetModerationInformation sets ModerationInformation field to given value.

### HasModerationInformation

`func (o *ServerConfigAboutInstance) HasModerationInformation() bool`

HasModerationInformation returns a boolean if a field has been set.

### GetAdministrator

`func (o *ServerConfigAboutInstance) GetAdministrator() string`

GetAdministrator returns the Administrator field if non-nil, zero value otherwise.

### GetAdministratorOk

`func (o *ServerConfigAboutInstance) GetAdministratorOk() (*string, bool)`

GetAdministratorOk returns a tuple with the Administrator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdministrator

`func (o *ServerConfigAboutInstance) SetAdministrator(v string)`

SetAdministrator sets Administrator field to given value.

### HasAdministrator

`func (o *ServerConfigAboutInstance) HasAdministrator() bool`

HasAdministrator returns a boolean if a field has been set.

### GetMaintenanceLifetime

`func (o *ServerConfigAboutInstance) GetMaintenanceLifetime() string`

GetMaintenanceLifetime returns the MaintenanceLifetime field if non-nil, zero value otherwise.

### GetMaintenanceLifetimeOk

`func (o *ServerConfigAboutInstance) GetMaintenanceLifetimeOk() (*string, bool)`

GetMaintenanceLifetimeOk returns a tuple with the MaintenanceLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceLifetime

`func (o *ServerConfigAboutInstance) SetMaintenanceLifetime(v string)`

SetMaintenanceLifetime sets MaintenanceLifetime field to given value.

### HasMaintenanceLifetime

`func (o *ServerConfigAboutInstance) HasMaintenanceLifetime() bool`

HasMaintenanceLifetime returns a boolean if a field has been set.

### GetBusinessModel

`func (o *ServerConfigAboutInstance) GetBusinessModel() string`

GetBusinessModel returns the BusinessModel field if non-nil, zero value otherwise.

### GetBusinessModelOk

`func (o *ServerConfigAboutInstance) GetBusinessModelOk() (*string, bool)`

GetBusinessModelOk returns a tuple with the BusinessModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessModel

`func (o *ServerConfigAboutInstance) SetBusinessModel(v string)`

SetBusinessModel sets BusinessModel field to given value.

### HasBusinessModel

`func (o *ServerConfigAboutInstance) HasBusinessModel() bool`

HasBusinessModel returns a boolean if a field has been set.

### GetLanguages

`func (o *ServerConfigAboutInstance) GetLanguages() []string`

GetLanguages returns the Languages field if non-nil, zero value otherwise.

### GetLanguagesOk

`func (o *ServerConfigAboutInstance) GetLanguagesOk() (*[]string, bool)`

GetLanguagesOk returns a tuple with the Languages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguages

`func (o *ServerConfigAboutInstance) SetLanguages(v []string)`

SetLanguages sets Languages field to given value.

### HasLanguages

`func (o *ServerConfigAboutInstance) HasLanguages() bool`

HasLanguages returns a boolean if a field has been set.

### GetCategories

`func (o *ServerConfigAboutInstance) GetCategories() []int32`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *ServerConfigAboutInstance) GetCategoriesOk() (*[]int32, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *ServerConfigAboutInstance) SetCategories(v []int32)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *ServerConfigAboutInstance) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetAvatars

`func (o *ServerConfigAboutInstance) GetAvatars() []ActorImage`

GetAvatars returns the Avatars field if non-nil, zero value otherwise.

### GetAvatarsOk

`func (o *ServerConfigAboutInstance) GetAvatarsOk() (*[]ActorImage, bool)`

GetAvatarsOk returns a tuple with the Avatars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatars

`func (o *ServerConfigAboutInstance) SetAvatars(v []ActorImage)`

SetAvatars sets Avatars field to given value.

### HasAvatars

`func (o *ServerConfigAboutInstance) HasAvatars() bool`

HasAvatars returns a boolean if a field has been set.

### GetBanners

`func (o *ServerConfigAboutInstance) GetBanners() []ActorImage`

GetBanners returns the Banners field if non-nil, zero value otherwise.

### GetBannersOk

`func (o *ServerConfigAboutInstance) GetBannersOk() (*[]ActorImage, bool)`

GetBannersOk returns a tuple with the Banners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBanners

`func (o *ServerConfigAboutInstance) SetBanners(v []ActorImage)`

SetBanners sets Banners field to given value.

### HasBanners

`func (o *ServerConfigAboutInstance) HasBanners() bool`

HasBanners returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


