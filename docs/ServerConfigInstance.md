# ServerConfigInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**ShortDescription** | Pointer to **string** |  | [optional] 
**DefaultClientRoute** | Pointer to **string** |  | [optional] 
**IsNSFW** | Pointer to **bool** |  | [optional] 
**DefaultNSFWPolicy** | Pointer to **string** |  | [optional] 
**ServerCountry** | Pointer to **string** |  | [optional] 
**Support** | Pointer to [**ServerConfigInstanceSupport**](ServerConfigInstanceSupport.md) |  | [optional] 
**Social** | Pointer to [**ServerConfigInstanceSocial**](ServerConfigInstanceSocial.md) |  | [optional] 
**Customizations** | Pointer to [**ServerConfigInstanceCustomizations**](ServerConfigInstanceCustomizations.md) |  | [optional] 
**Avatars** | Pointer to [**[]ActorImage**](ActorImage.md) |  | [optional] 
**Banners** | Pointer to [**[]ActorImage**](ActorImage.md) |  | [optional] 

## Methods

### NewServerConfigInstance

`func NewServerConfigInstance() *ServerConfigInstance`

NewServerConfigInstance instantiates a new ServerConfigInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerConfigInstanceWithDefaults

`func NewServerConfigInstanceWithDefaults() *ServerConfigInstance`

NewServerConfigInstanceWithDefaults instantiates a new ServerConfigInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ServerConfigInstance) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServerConfigInstance) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServerConfigInstance) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ServerConfigInstance) HasName() bool`

HasName returns a boolean if a field has been set.

### GetShortDescription

`func (o *ServerConfigInstance) GetShortDescription() string`

GetShortDescription returns the ShortDescription field if non-nil, zero value otherwise.

### GetShortDescriptionOk

`func (o *ServerConfigInstance) GetShortDescriptionOk() (*string, bool)`

GetShortDescriptionOk returns a tuple with the ShortDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortDescription

`func (o *ServerConfigInstance) SetShortDescription(v string)`

SetShortDescription sets ShortDescription field to given value.

### HasShortDescription

`func (o *ServerConfigInstance) HasShortDescription() bool`

HasShortDescription returns a boolean if a field has been set.

### GetDefaultClientRoute

`func (o *ServerConfigInstance) GetDefaultClientRoute() string`

GetDefaultClientRoute returns the DefaultClientRoute field if non-nil, zero value otherwise.

### GetDefaultClientRouteOk

`func (o *ServerConfigInstance) GetDefaultClientRouteOk() (*string, bool)`

GetDefaultClientRouteOk returns a tuple with the DefaultClientRoute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultClientRoute

`func (o *ServerConfigInstance) SetDefaultClientRoute(v string)`

SetDefaultClientRoute sets DefaultClientRoute field to given value.

### HasDefaultClientRoute

`func (o *ServerConfigInstance) HasDefaultClientRoute() bool`

HasDefaultClientRoute returns a boolean if a field has been set.

### GetIsNSFW

`func (o *ServerConfigInstance) GetIsNSFW() bool`

GetIsNSFW returns the IsNSFW field if non-nil, zero value otherwise.

### GetIsNSFWOk

`func (o *ServerConfigInstance) GetIsNSFWOk() (*bool, bool)`

GetIsNSFWOk returns a tuple with the IsNSFW field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNSFW

`func (o *ServerConfigInstance) SetIsNSFW(v bool)`

SetIsNSFW sets IsNSFW field to given value.

### HasIsNSFW

`func (o *ServerConfigInstance) HasIsNSFW() bool`

HasIsNSFW returns a boolean if a field has been set.

### GetDefaultNSFWPolicy

`func (o *ServerConfigInstance) GetDefaultNSFWPolicy() string`

GetDefaultNSFWPolicy returns the DefaultNSFWPolicy field if non-nil, zero value otherwise.

### GetDefaultNSFWPolicyOk

`func (o *ServerConfigInstance) GetDefaultNSFWPolicyOk() (*string, bool)`

GetDefaultNSFWPolicyOk returns a tuple with the DefaultNSFWPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultNSFWPolicy

`func (o *ServerConfigInstance) SetDefaultNSFWPolicy(v string)`

SetDefaultNSFWPolicy sets DefaultNSFWPolicy field to given value.

### HasDefaultNSFWPolicy

`func (o *ServerConfigInstance) HasDefaultNSFWPolicy() bool`

HasDefaultNSFWPolicy returns a boolean if a field has been set.

### GetServerCountry

`func (o *ServerConfigInstance) GetServerCountry() string`

GetServerCountry returns the ServerCountry field if non-nil, zero value otherwise.

### GetServerCountryOk

`func (o *ServerConfigInstance) GetServerCountryOk() (*string, bool)`

GetServerCountryOk returns a tuple with the ServerCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerCountry

`func (o *ServerConfigInstance) SetServerCountry(v string)`

SetServerCountry sets ServerCountry field to given value.

### HasServerCountry

`func (o *ServerConfigInstance) HasServerCountry() bool`

HasServerCountry returns a boolean if a field has been set.

### GetSupport

`func (o *ServerConfigInstance) GetSupport() ServerConfigInstanceSupport`

GetSupport returns the Support field if non-nil, zero value otherwise.

### GetSupportOk

`func (o *ServerConfigInstance) GetSupportOk() (*ServerConfigInstanceSupport, bool)`

GetSupportOk returns a tuple with the Support field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupport

`func (o *ServerConfigInstance) SetSupport(v ServerConfigInstanceSupport)`

SetSupport sets Support field to given value.

### HasSupport

`func (o *ServerConfigInstance) HasSupport() bool`

HasSupport returns a boolean if a field has been set.

### GetSocial

`func (o *ServerConfigInstance) GetSocial() ServerConfigInstanceSocial`

GetSocial returns the Social field if non-nil, zero value otherwise.

### GetSocialOk

`func (o *ServerConfigInstance) GetSocialOk() (*ServerConfigInstanceSocial, bool)`

GetSocialOk returns a tuple with the Social field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocial

`func (o *ServerConfigInstance) SetSocial(v ServerConfigInstanceSocial)`

SetSocial sets Social field to given value.

### HasSocial

`func (o *ServerConfigInstance) HasSocial() bool`

HasSocial returns a boolean if a field has been set.

### GetCustomizations

`func (o *ServerConfigInstance) GetCustomizations() ServerConfigInstanceCustomizations`

GetCustomizations returns the Customizations field if non-nil, zero value otherwise.

### GetCustomizationsOk

`func (o *ServerConfigInstance) GetCustomizationsOk() (*ServerConfigInstanceCustomizations, bool)`

GetCustomizationsOk returns a tuple with the Customizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomizations

`func (o *ServerConfigInstance) SetCustomizations(v ServerConfigInstanceCustomizations)`

SetCustomizations sets Customizations field to given value.

### HasCustomizations

`func (o *ServerConfigInstance) HasCustomizations() bool`

HasCustomizations returns a boolean if a field has been set.

### GetAvatars

`func (o *ServerConfigInstance) GetAvatars() []ActorImage`

GetAvatars returns the Avatars field if non-nil, zero value otherwise.

### GetAvatarsOk

`func (o *ServerConfigInstance) GetAvatarsOk() (*[]ActorImage, bool)`

GetAvatarsOk returns a tuple with the Avatars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatars

`func (o *ServerConfigInstance) SetAvatars(v []ActorImage)`

SetAvatars sets Avatars field to given value.

### HasAvatars

`func (o *ServerConfigInstance) HasAvatars() bool`

HasAvatars returns a boolean if a field has been set.

### GetBanners

`func (o *ServerConfigInstance) GetBanners() []ActorImage`

GetBanners returns the Banners field if non-nil, zero value otherwise.

### GetBannersOk

`func (o *ServerConfigInstance) GetBannersOk() (*[]ActorImage, bool)`

GetBannersOk returns a tuple with the Banners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBanners

`func (o *ServerConfigInstance) SetBanners(v []ActorImage)`

SetBanners sets Banners field to given value.

### HasBanners

`func (o *ServerConfigInstance) HasBanners() bool`

HasBanners returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


