# ServerConfigCustomInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**ShortDescription** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Terms** | Pointer to **string** |  | [optional] 
**CodeOfConduct** | Pointer to **string** |  | [optional] 
**CreationReason** | Pointer to **string** |  | [optional] 
**ModerationInformation** | Pointer to **string** |  | [optional] 
**Administrator** | Pointer to **string** |  | [optional] 
**MaintenanceLifetime** | Pointer to **string** |  | [optional] 
**BusinessModel** | Pointer to **string** |  | [optional] 
**HardwareInformation** | Pointer to **string** |  | [optional] 
**Languages** | Pointer to **[]string** |  | [optional] 
**Categories** | Pointer to **[]float32** |  | [optional] 
**IsNSFW** | Pointer to **bool** |  | [optional] 
**DefaultNSFWPolicy** | Pointer to **string** |  | [optional] 
**ServerCountry** | Pointer to **string** |  | [optional] 
**Support** | Pointer to [**ServerConfigInstanceSupport**](ServerConfigInstanceSupport.md) |  | [optional] 
**Social** | Pointer to [**ServerConfigInstanceSocial**](ServerConfigInstanceSocial.md) |  | [optional] 
**DefaultClientRoute** | Pointer to **string** |  | [optional] 
**Customizations** | Pointer to [**ServerConfigInstanceCustomizations**](ServerConfigInstanceCustomizations.md) |  | [optional] 

## Methods

### NewServerConfigCustomInstance

`func NewServerConfigCustomInstance() *ServerConfigCustomInstance`

NewServerConfigCustomInstance instantiates a new ServerConfigCustomInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerConfigCustomInstanceWithDefaults

`func NewServerConfigCustomInstanceWithDefaults() *ServerConfigCustomInstance`

NewServerConfigCustomInstanceWithDefaults instantiates a new ServerConfigCustomInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ServerConfigCustomInstance) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServerConfigCustomInstance) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServerConfigCustomInstance) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ServerConfigCustomInstance) HasName() bool`

HasName returns a boolean if a field has been set.

### GetShortDescription

`func (o *ServerConfigCustomInstance) GetShortDescription() string`

GetShortDescription returns the ShortDescription field if non-nil, zero value otherwise.

### GetShortDescriptionOk

`func (o *ServerConfigCustomInstance) GetShortDescriptionOk() (*string, bool)`

GetShortDescriptionOk returns a tuple with the ShortDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortDescription

`func (o *ServerConfigCustomInstance) SetShortDescription(v string)`

SetShortDescription sets ShortDescription field to given value.

### HasShortDescription

`func (o *ServerConfigCustomInstance) HasShortDescription() bool`

HasShortDescription returns a boolean if a field has been set.

### GetDescription

`func (o *ServerConfigCustomInstance) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ServerConfigCustomInstance) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ServerConfigCustomInstance) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ServerConfigCustomInstance) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTerms

`func (o *ServerConfigCustomInstance) GetTerms() string`

GetTerms returns the Terms field if non-nil, zero value otherwise.

### GetTermsOk

`func (o *ServerConfigCustomInstance) GetTermsOk() (*string, bool)`

GetTermsOk returns a tuple with the Terms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerms

`func (o *ServerConfigCustomInstance) SetTerms(v string)`

SetTerms sets Terms field to given value.

### HasTerms

`func (o *ServerConfigCustomInstance) HasTerms() bool`

HasTerms returns a boolean if a field has been set.

### GetCodeOfConduct

`func (o *ServerConfigCustomInstance) GetCodeOfConduct() string`

GetCodeOfConduct returns the CodeOfConduct field if non-nil, zero value otherwise.

### GetCodeOfConductOk

`func (o *ServerConfigCustomInstance) GetCodeOfConductOk() (*string, bool)`

GetCodeOfConductOk returns a tuple with the CodeOfConduct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeOfConduct

`func (o *ServerConfigCustomInstance) SetCodeOfConduct(v string)`

SetCodeOfConduct sets CodeOfConduct field to given value.

### HasCodeOfConduct

`func (o *ServerConfigCustomInstance) HasCodeOfConduct() bool`

HasCodeOfConduct returns a boolean if a field has been set.

### GetCreationReason

`func (o *ServerConfigCustomInstance) GetCreationReason() string`

GetCreationReason returns the CreationReason field if non-nil, zero value otherwise.

### GetCreationReasonOk

`func (o *ServerConfigCustomInstance) GetCreationReasonOk() (*string, bool)`

GetCreationReasonOk returns a tuple with the CreationReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationReason

`func (o *ServerConfigCustomInstance) SetCreationReason(v string)`

SetCreationReason sets CreationReason field to given value.

### HasCreationReason

`func (o *ServerConfigCustomInstance) HasCreationReason() bool`

HasCreationReason returns a boolean if a field has been set.

### GetModerationInformation

`func (o *ServerConfigCustomInstance) GetModerationInformation() string`

GetModerationInformation returns the ModerationInformation field if non-nil, zero value otherwise.

### GetModerationInformationOk

`func (o *ServerConfigCustomInstance) GetModerationInformationOk() (*string, bool)`

GetModerationInformationOk returns a tuple with the ModerationInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModerationInformation

`func (o *ServerConfigCustomInstance) SetModerationInformation(v string)`

SetModerationInformation sets ModerationInformation field to given value.

### HasModerationInformation

`func (o *ServerConfigCustomInstance) HasModerationInformation() bool`

HasModerationInformation returns a boolean if a field has been set.

### GetAdministrator

`func (o *ServerConfigCustomInstance) GetAdministrator() string`

GetAdministrator returns the Administrator field if non-nil, zero value otherwise.

### GetAdministratorOk

`func (o *ServerConfigCustomInstance) GetAdministratorOk() (*string, bool)`

GetAdministratorOk returns a tuple with the Administrator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdministrator

`func (o *ServerConfigCustomInstance) SetAdministrator(v string)`

SetAdministrator sets Administrator field to given value.

### HasAdministrator

`func (o *ServerConfigCustomInstance) HasAdministrator() bool`

HasAdministrator returns a boolean if a field has been set.

### GetMaintenanceLifetime

`func (o *ServerConfigCustomInstance) GetMaintenanceLifetime() string`

GetMaintenanceLifetime returns the MaintenanceLifetime field if non-nil, zero value otherwise.

### GetMaintenanceLifetimeOk

`func (o *ServerConfigCustomInstance) GetMaintenanceLifetimeOk() (*string, bool)`

GetMaintenanceLifetimeOk returns a tuple with the MaintenanceLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceLifetime

`func (o *ServerConfigCustomInstance) SetMaintenanceLifetime(v string)`

SetMaintenanceLifetime sets MaintenanceLifetime field to given value.

### HasMaintenanceLifetime

`func (o *ServerConfigCustomInstance) HasMaintenanceLifetime() bool`

HasMaintenanceLifetime returns a boolean if a field has been set.

### GetBusinessModel

`func (o *ServerConfigCustomInstance) GetBusinessModel() string`

GetBusinessModel returns the BusinessModel field if non-nil, zero value otherwise.

### GetBusinessModelOk

`func (o *ServerConfigCustomInstance) GetBusinessModelOk() (*string, bool)`

GetBusinessModelOk returns a tuple with the BusinessModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessModel

`func (o *ServerConfigCustomInstance) SetBusinessModel(v string)`

SetBusinessModel sets BusinessModel field to given value.

### HasBusinessModel

`func (o *ServerConfigCustomInstance) HasBusinessModel() bool`

HasBusinessModel returns a boolean if a field has been set.

### GetHardwareInformation

`func (o *ServerConfigCustomInstance) GetHardwareInformation() string`

GetHardwareInformation returns the HardwareInformation field if non-nil, zero value otherwise.

### GetHardwareInformationOk

`func (o *ServerConfigCustomInstance) GetHardwareInformationOk() (*string, bool)`

GetHardwareInformationOk returns a tuple with the HardwareInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardwareInformation

`func (o *ServerConfigCustomInstance) SetHardwareInformation(v string)`

SetHardwareInformation sets HardwareInformation field to given value.

### HasHardwareInformation

`func (o *ServerConfigCustomInstance) HasHardwareInformation() bool`

HasHardwareInformation returns a boolean if a field has been set.

### GetLanguages

`func (o *ServerConfigCustomInstance) GetLanguages() []string`

GetLanguages returns the Languages field if non-nil, zero value otherwise.

### GetLanguagesOk

`func (o *ServerConfigCustomInstance) GetLanguagesOk() (*[]string, bool)`

GetLanguagesOk returns a tuple with the Languages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguages

`func (o *ServerConfigCustomInstance) SetLanguages(v []string)`

SetLanguages sets Languages field to given value.

### HasLanguages

`func (o *ServerConfigCustomInstance) HasLanguages() bool`

HasLanguages returns a boolean if a field has been set.

### GetCategories

`func (o *ServerConfigCustomInstance) GetCategories() []float32`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *ServerConfigCustomInstance) GetCategoriesOk() (*[]float32, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *ServerConfigCustomInstance) SetCategories(v []float32)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *ServerConfigCustomInstance) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetIsNSFW

`func (o *ServerConfigCustomInstance) GetIsNSFW() bool`

GetIsNSFW returns the IsNSFW field if non-nil, zero value otherwise.

### GetIsNSFWOk

`func (o *ServerConfigCustomInstance) GetIsNSFWOk() (*bool, bool)`

GetIsNSFWOk returns a tuple with the IsNSFW field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNSFW

`func (o *ServerConfigCustomInstance) SetIsNSFW(v bool)`

SetIsNSFW sets IsNSFW field to given value.

### HasIsNSFW

`func (o *ServerConfigCustomInstance) HasIsNSFW() bool`

HasIsNSFW returns a boolean if a field has been set.

### GetDefaultNSFWPolicy

`func (o *ServerConfigCustomInstance) GetDefaultNSFWPolicy() string`

GetDefaultNSFWPolicy returns the DefaultNSFWPolicy field if non-nil, zero value otherwise.

### GetDefaultNSFWPolicyOk

`func (o *ServerConfigCustomInstance) GetDefaultNSFWPolicyOk() (*string, bool)`

GetDefaultNSFWPolicyOk returns a tuple with the DefaultNSFWPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultNSFWPolicy

`func (o *ServerConfigCustomInstance) SetDefaultNSFWPolicy(v string)`

SetDefaultNSFWPolicy sets DefaultNSFWPolicy field to given value.

### HasDefaultNSFWPolicy

`func (o *ServerConfigCustomInstance) HasDefaultNSFWPolicy() bool`

HasDefaultNSFWPolicy returns a boolean if a field has been set.

### GetServerCountry

`func (o *ServerConfigCustomInstance) GetServerCountry() string`

GetServerCountry returns the ServerCountry field if non-nil, zero value otherwise.

### GetServerCountryOk

`func (o *ServerConfigCustomInstance) GetServerCountryOk() (*string, bool)`

GetServerCountryOk returns a tuple with the ServerCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerCountry

`func (o *ServerConfigCustomInstance) SetServerCountry(v string)`

SetServerCountry sets ServerCountry field to given value.

### HasServerCountry

`func (o *ServerConfigCustomInstance) HasServerCountry() bool`

HasServerCountry returns a boolean if a field has been set.

### GetSupport

`func (o *ServerConfigCustomInstance) GetSupport() ServerConfigInstanceSupport`

GetSupport returns the Support field if non-nil, zero value otherwise.

### GetSupportOk

`func (o *ServerConfigCustomInstance) GetSupportOk() (*ServerConfigInstanceSupport, bool)`

GetSupportOk returns a tuple with the Support field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupport

`func (o *ServerConfigCustomInstance) SetSupport(v ServerConfigInstanceSupport)`

SetSupport sets Support field to given value.

### HasSupport

`func (o *ServerConfigCustomInstance) HasSupport() bool`

HasSupport returns a boolean if a field has been set.

### GetSocial

`func (o *ServerConfigCustomInstance) GetSocial() ServerConfigInstanceSocial`

GetSocial returns the Social field if non-nil, zero value otherwise.

### GetSocialOk

`func (o *ServerConfigCustomInstance) GetSocialOk() (*ServerConfigInstanceSocial, bool)`

GetSocialOk returns a tuple with the Social field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocial

`func (o *ServerConfigCustomInstance) SetSocial(v ServerConfigInstanceSocial)`

SetSocial sets Social field to given value.

### HasSocial

`func (o *ServerConfigCustomInstance) HasSocial() bool`

HasSocial returns a boolean if a field has been set.

### GetDefaultClientRoute

`func (o *ServerConfigCustomInstance) GetDefaultClientRoute() string`

GetDefaultClientRoute returns the DefaultClientRoute field if non-nil, zero value otherwise.

### GetDefaultClientRouteOk

`func (o *ServerConfigCustomInstance) GetDefaultClientRouteOk() (*string, bool)`

GetDefaultClientRouteOk returns a tuple with the DefaultClientRoute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultClientRoute

`func (o *ServerConfigCustomInstance) SetDefaultClientRoute(v string)`

SetDefaultClientRoute sets DefaultClientRoute field to given value.

### HasDefaultClientRoute

`func (o *ServerConfigCustomInstance) HasDefaultClientRoute() bool`

HasDefaultClientRoute returns a boolean if a field has been set.

### GetCustomizations

`func (o *ServerConfigCustomInstance) GetCustomizations() ServerConfigInstanceCustomizations`

GetCustomizations returns the Customizations field if non-nil, zero value otherwise.

### GetCustomizationsOk

`func (o *ServerConfigCustomInstance) GetCustomizationsOk() (*ServerConfigInstanceCustomizations, bool)`

GetCustomizationsOk returns a tuple with the Customizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomizations

`func (o *ServerConfigCustomInstance) SetCustomizations(v ServerConfigInstanceCustomizations)`

SetCustomizations sets Customizations field to given value.

### HasCustomizations

`func (o *ServerConfigCustomInstance) HasCustomizations() bool`

HasCustomizations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


