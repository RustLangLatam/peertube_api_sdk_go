/*
PeerTube

The PeerTube API is built on HTTP(S) and is RESTful. You can use your favorite HTTP/REST library for your programming language to use PeerTube. The spec API is fully compatible with [openapi-generator](https://github.com/OpenAPITools/openapi-generator/wiki/API-client-generator-HOWTO) which generates a client SDK in the language of your choice - we generate some client SDKs automatically:  - [Python](https://framagit.org/framasoft/peertube/clients/python) - [Go](https://framagit.org/framasoft/peertube/clients/go) - [Kotlin](https://framagit.org/framasoft/peertube/clients/kotlin)  See the [REST API quick start](https://docs.joinpeertube.org/api/rest-getting-started) for a few examples of using the PeerTube API.  # Authentication  When you sign up for an account on a PeerTube instance, you are given the possibility to generate sessions on it, and authenticate there using an access token. Only __one access token can currently be used at a time__.  ## Roles  Accounts are given permissions based on their role. There are three roles on PeerTube: Administrator, Moderator, and User. See the [roles guide](https://docs.joinpeertube.org/admin/managing-users#roles) for a detail of their permissions.  # Errors  The API uses standard HTTP status codes to indicate the success or failure of the API call, completed by a [RFC7807-compliant](https://tools.ietf.org/html/rfc7807) response body.  ``` HTTP 1.1 404 Not Found Content-Type: application/problem+json; charset=utf-8  {   \"detail\": \"Video not found\",   \"docs\": \"https://docs.joinpeertube.org/api-rest-reference.html#operation/getVideo\",   \"status\": 404,   \"title\": \"Not Found\",   \"type\": \"about:blank\" } ```  We provide error `type` (following RFC7807) and `code` (internal PeerTube code) values for [a growing number of cases](https://github.com/Chocobozzz/PeerTube/blob/develop/packages/models/src/server/server-error-code.enum.ts), but it is still optional. Types are used to disambiguate errors that bear the same status code and are non-obvious:  ``` HTTP 1.1 403 Forbidden Content-Type: application/problem+json; charset=utf-8  {   \"detail\": \"Cannot get this video regarding follow constraints\",   \"docs\": \"https://docs.joinpeertube.org/api-rest-reference.html#operation/getVideo\",   \"status\": 403,   \"title\": \"Forbidden\",   \"type\": \"https://docs.joinpeertube.org/api-rest-reference.html#section/Errors/does_not_respect_follow_constraints\" } ```  Here a 403 error could otherwise mean that the video is private or blocklisted.  ### Validation errors  Each parameter is evaluated on its own against a set of rules before the route validator proceeds with potential testing involving parameter combinations. Errors coming from validation errors appear earlier and benefit from a more detailed error description:  ``` HTTP 1.1 400 Bad Request Content-Type: application/problem+json; charset=utf-8  {   \"detail\": \"Incorrect request parameters: id\",   \"docs\": \"https://docs.joinpeertube.org/api-rest-reference.html#operation/getVideo\",   \"instance\": \"/api/v1/videos/9c9de5e8-0a1e-484a-b099-e80766180\",   \"invalid-params\": {     \"id\": {       \"location\": \"params\",       \"msg\": \"Invalid value\",       \"param\": \"id\",       \"value\": \"9c9de5e8-0a1e-484a-b099-e80766180\"     }   },   \"status\": 400,   \"title\": \"Bad Request\",   \"type\": \"about:blank\" } ```  Where `id` is the name of the field concerned by the error, within the route definition. `invalid-params.<field>.location` can be either 'params', 'body', 'header', 'query' or 'cookies', and `invalid-params.<field>.value` reports the value that didn't pass validation whose `invalid-params.<field>.msg` is about.  ### Deprecated error fields  Some fields could be included with previous versions. They are still included but their use is deprecated: - `error`: superseded by `detail`  # Rate limits  We are rate-limiting all endpoints of PeerTube's API. Custom values can be set by administrators:  | Endpoint (prefix: `/api/v1`) | Calls         | Time frame   | |------------------------------|---------------|--------------| | `/_*`                         | 50            | 10 seconds   | | `POST /users/token`          | 15            | 5 minutes    | | `POST /users/register`       | 2<sup>*</sup> | 5 minutes    | | `POST /users/ask-send-verify-email` | 3      | 5 minutes    |  Depending on the endpoint, <sup>*</sup>failed requests are not taken into account. A service limit is announced by a `429 Too Many Requests` status code.  You can get details about the current state of your rate limit by reading the following headers:  | Header                  | Description                                                | |-------------------------|------------------------------------------------------------| | `X-RateLimit-Limit`     | Number of max requests allowed in the current time period  | | `X-RateLimit-Remaining` | Number of remaining requests in the current time period    | | `X-RateLimit-Reset`     | Timestamp of end of current time period as UNIX timestamp  | | `Retry-After`           | Seconds to delay after the first `429` is received         |  # CORS  This API features [Cross-Origin Resource Sharing (CORS)](https://fetch.spec.whatwg.org/), allowing cross-domain communication from the browser for some routes:  | Endpoint                    | |------------------------- ---| | `/api/_*`                    | | `/download/_*`               | | `/lazy-static/_*`            | | `/.well-known/webfinger`    |  In addition, all routes serving ActivityPub are CORS-enabled for all origins.

API version: 7.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"
	"github.com/RustLangLatam/peertube_api_sdk_go/utils"
)

// checks if the ServerConfigAboutInstance type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ServerConfigAboutInstance{}

// ServerConfigAboutInstance struct for ServerConfigAboutInstance
type ServerConfigAboutInstance struct {
	Name                  *string      `json:"name,omitempty"`
	ShortDescription      *string      `json:"shortDescription,omitempty"`
	Description           *string      `json:"description,omitempty"`
	Terms                 *string      `json:"terms,omitempty"`
	CodeOfConduct         *string      `json:"codeOfConduct,omitempty"`
	HardwareInformation   *string      `json:"hardwareInformation,omitempty"`
	CreationReason        *string      `json:"creationReason,omitempty"`
	ModerationInformation *string      `json:"moderationInformation,omitempty"`
	Administrator         *string      `json:"administrator,omitempty"`
	MaintenanceLifetime   *string      `json:"maintenanceLifetime,omitempty"`
	BusinessModel         *string      `json:"businessModel,omitempty"`
	Languages             []string     `json:"languages,omitempty"`
	Categories            []int32      `json:"categories,omitempty"`
	Avatars               []ActorImage `json:"avatars,omitempty"`
	Banners               []ActorImage `json:"banners,omitempty"`
}

// NewServerConfigAboutInstance instantiates a new ServerConfigAboutInstance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerConfigAboutInstance() *ServerConfigAboutInstance {
	this := ServerConfigAboutInstance{}
	return &this
}

// NewServerConfigAboutInstanceWithDefaults instantiates a new ServerConfigAboutInstance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerConfigAboutInstanceWithDefaults() *ServerConfigAboutInstance {
	this := ServerConfigAboutInstance{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ServerConfigAboutInstance) GetName() string {
	if o == nil || utils.IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfigAboutInstance) GetNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ServerConfigAboutInstance) HasName() bool {
	if o != nil && !utils.IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ServerConfigAboutInstance) SetName(v string) {
	o.Name = &v
}

// GetShortDescription returns the ShortDescription field value if set, zero value otherwise.
func (o *ServerConfigAboutInstance) GetShortDescription() string {
	if o == nil || utils.IsNil(o.ShortDescription) {
		var ret string
		return ret
	}
	return *o.ShortDescription
}

// GetShortDescriptionOk returns a tuple with the ShortDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfigAboutInstance) GetShortDescriptionOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ShortDescription) {
		return nil, false
	}
	return o.ShortDescription, true
}

// HasShortDescription returns a boolean if a field has been set.
func (o *ServerConfigAboutInstance) HasShortDescription() bool {
	if o != nil && !utils.IsNil(o.ShortDescription) {
		return true
	}

	return false
}

// SetShortDescription gets a reference to the given string and assigns it to the ShortDescription field.
func (o *ServerConfigAboutInstance) SetShortDescription(v string) {
	o.ShortDescription = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ServerConfigAboutInstance) GetDescription() string {
	if o == nil || utils.IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfigAboutInstance) GetDescriptionOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ServerConfigAboutInstance) HasDescription() bool {
	if o != nil && !utils.IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ServerConfigAboutInstance) SetDescription(v string) {
	o.Description = &v
}

// GetTerms returns the Terms field value if set, zero value otherwise.
func (o *ServerConfigAboutInstance) GetTerms() string {
	if o == nil || utils.IsNil(o.Terms) {
		var ret string
		return ret
	}
	return *o.Terms
}

// GetTermsOk returns a tuple with the Terms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfigAboutInstance) GetTermsOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Terms) {
		return nil, false
	}
	return o.Terms, true
}

// HasTerms returns a boolean if a field has been set.
func (o *ServerConfigAboutInstance) HasTerms() bool {
	if o != nil && !utils.IsNil(o.Terms) {
		return true
	}

	return false
}

// SetTerms gets a reference to the given string and assigns it to the Terms field.
func (o *ServerConfigAboutInstance) SetTerms(v string) {
	o.Terms = &v
}

// GetCodeOfConduct returns the CodeOfConduct field value if set, zero value otherwise.
func (o *ServerConfigAboutInstance) GetCodeOfConduct() string {
	if o == nil || utils.IsNil(o.CodeOfConduct) {
		var ret string
		return ret
	}
	return *o.CodeOfConduct
}

// GetCodeOfConductOk returns a tuple with the CodeOfConduct field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfigAboutInstance) GetCodeOfConductOk() (*string, bool) {
	if o == nil || utils.IsNil(o.CodeOfConduct) {
		return nil, false
	}
	return o.CodeOfConduct, true
}

// HasCodeOfConduct returns a boolean if a field has been set.
func (o *ServerConfigAboutInstance) HasCodeOfConduct() bool {
	if o != nil && !utils.IsNil(o.CodeOfConduct) {
		return true
	}

	return false
}

// SetCodeOfConduct gets a reference to the given string and assigns it to the CodeOfConduct field.
func (o *ServerConfigAboutInstance) SetCodeOfConduct(v string) {
	o.CodeOfConduct = &v
}

// GetHardwareInformation returns the HardwareInformation field value if set, zero value otherwise.
func (o *ServerConfigAboutInstance) GetHardwareInformation() string {
	if o == nil || utils.IsNil(o.HardwareInformation) {
		var ret string
		return ret
	}
	return *o.HardwareInformation
}

// GetHardwareInformationOk returns a tuple with the HardwareInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfigAboutInstance) GetHardwareInformationOk() (*string, bool) {
	if o == nil || utils.IsNil(o.HardwareInformation) {
		return nil, false
	}
	return o.HardwareInformation, true
}

// HasHardwareInformation returns a boolean if a field has been set.
func (o *ServerConfigAboutInstance) HasHardwareInformation() bool {
	if o != nil && !utils.IsNil(o.HardwareInformation) {
		return true
	}

	return false
}

// SetHardwareInformation gets a reference to the given string and assigns it to the HardwareInformation field.
func (o *ServerConfigAboutInstance) SetHardwareInformation(v string) {
	o.HardwareInformation = &v
}

// GetCreationReason returns the CreationReason field value if set, zero value otherwise.
func (o *ServerConfigAboutInstance) GetCreationReason() string {
	if o == nil || utils.IsNil(o.CreationReason) {
		var ret string
		return ret
	}
	return *o.CreationReason
}

// GetCreationReasonOk returns a tuple with the CreationReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfigAboutInstance) GetCreationReasonOk() (*string, bool) {
	if o == nil || utils.IsNil(o.CreationReason) {
		return nil, false
	}
	return o.CreationReason, true
}

// HasCreationReason returns a boolean if a field has been set.
func (o *ServerConfigAboutInstance) HasCreationReason() bool {
	if o != nil && !utils.IsNil(o.CreationReason) {
		return true
	}

	return false
}

// SetCreationReason gets a reference to the given string and assigns it to the CreationReason field.
func (o *ServerConfigAboutInstance) SetCreationReason(v string) {
	o.CreationReason = &v
}

// GetModerationInformation returns the ModerationInformation field value if set, zero value otherwise.
func (o *ServerConfigAboutInstance) GetModerationInformation() string {
	if o == nil || utils.IsNil(o.ModerationInformation) {
		var ret string
		return ret
	}
	return *o.ModerationInformation
}

// GetModerationInformationOk returns a tuple with the ModerationInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfigAboutInstance) GetModerationInformationOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ModerationInformation) {
		return nil, false
	}
	return o.ModerationInformation, true
}

// HasModerationInformation returns a boolean if a field has been set.
func (o *ServerConfigAboutInstance) HasModerationInformation() bool {
	if o != nil && !utils.IsNil(o.ModerationInformation) {
		return true
	}

	return false
}

// SetModerationInformation gets a reference to the given string and assigns it to the ModerationInformation field.
func (o *ServerConfigAboutInstance) SetModerationInformation(v string) {
	o.ModerationInformation = &v
}

// GetAdministrator returns the Administrator field value if set, zero value otherwise.
func (o *ServerConfigAboutInstance) GetAdministrator() string {
	if o == nil || utils.IsNil(o.Administrator) {
		var ret string
		return ret
	}
	return *o.Administrator
}

// GetAdministratorOk returns a tuple with the Administrator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfigAboutInstance) GetAdministratorOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Administrator) {
		return nil, false
	}
	return o.Administrator, true
}

// HasAdministrator returns a boolean if a field has been set.
func (o *ServerConfigAboutInstance) HasAdministrator() bool {
	if o != nil && !utils.IsNil(o.Administrator) {
		return true
	}

	return false
}

// SetAdministrator gets a reference to the given string and assigns it to the Administrator field.
func (o *ServerConfigAboutInstance) SetAdministrator(v string) {
	o.Administrator = &v
}

// GetMaintenanceLifetime returns the MaintenanceLifetime field value if set, zero value otherwise.
func (o *ServerConfigAboutInstance) GetMaintenanceLifetime() string {
	if o == nil || utils.IsNil(o.MaintenanceLifetime) {
		var ret string
		return ret
	}
	return *o.MaintenanceLifetime
}

// GetMaintenanceLifetimeOk returns a tuple with the MaintenanceLifetime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfigAboutInstance) GetMaintenanceLifetimeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.MaintenanceLifetime) {
		return nil, false
	}
	return o.MaintenanceLifetime, true
}

// HasMaintenanceLifetime returns a boolean if a field has been set.
func (o *ServerConfigAboutInstance) HasMaintenanceLifetime() bool {
	if o != nil && !utils.IsNil(o.MaintenanceLifetime) {
		return true
	}

	return false
}

// SetMaintenanceLifetime gets a reference to the given string and assigns it to the MaintenanceLifetime field.
func (o *ServerConfigAboutInstance) SetMaintenanceLifetime(v string) {
	o.MaintenanceLifetime = &v
}

// GetBusinessModel returns the BusinessModel field value if set, zero value otherwise.
func (o *ServerConfigAboutInstance) GetBusinessModel() string {
	if o == nil || utils.IsNil(o.BusinessModel) {
		var ret string
		return ret
	}
	return *o.BusinessModel
}

// GetBusinessModelOk returns a tuple with the BusinessModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfigAboutInstance) GetBusinessModelOk() (*string, bool) {
	if o == nil || utils.IsNil(o.BusinessModel) {
		return nil, false
	}
	return o.BusinessModel, true
}

// HasBusinessModel returns a boolean if a field has been set.
func (o *ServerConfigAboutInstance) HasBusinessModel() bool {
	if o != nil && !utils.IsNil(o.BusinessModel) {
		return true
	}

	return false
}

// SetBusinessModel gets a reference to the given string and assigns it to the BusinessModel field.
func (o *ServerConfigAboutInstance) SetBusinessModel(v string) {
	o.BusinessModel = &v
}

// GetLanguages returns the Languages field value if set, zero value otherwise.
func (o *ServerConfigAboutInstance) GetLanguages() []string {
	if o == nil || utils.IsNil(o.Languages) {
		var ret []string
		return ret
	}
	return o.Languages
}

// GetLanguagesOk returns a tuple with the Languages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfigAboutInstance) GetLanguagesOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.Languages) {
		return nil, false
	}
	return o.Languages, true
}

// HasLanguages returns a boolean if a field has been set.
func (o *ServerConfigAboutInstance) HasLanguages() bool {
	if o != nil && !utils.IsNil(o.Languages) {
		return true
	}

	return false
}

// SetLanguages gets a reference to the given []string and assigns it to the Languages field.
func (o *ServerConfigAboutInstance) SetLanguages(v []string) {
	o.Languages = v
}

// GetCategories returns the Categories field value if set, zero value otherwise.
func (o *ServerConfigAboutInstance) GetCategories() []int32 {
	if o == nil || utils.IsNil(o.Categories) {
		var ret []int32
		return ret
	}
	return o.Categories
}

// GetCategoriesOk returns a tuple with the Categories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfigAboutInstance) GetCategoriesOk() ([]int32, bool) {
	if o == nil || utils.IsNil(o.Categories) {
		return nil, false
	}
	return o.Categories, true
}

// HasCategories returns a boolean if a field has been set.
func (o *ServerConfigAboutInstance) HasCategories() bool {
	if o != nil && !utils.IsNil(o.Categories) {
		return true
	}

	return false
}

// SetCategories gets a reference to the given []int32 and assigns it to the Categories field.
func (o *ServerConfigAboutInstance) SetCategories(v []int32) {
	o.Categories = v
}

// GetAvatars returns the Avatars field value if set, zero value otherwise.
func (o *ServerConfigAboutInstance) GetAvatars() []ActorImage {
	if o == nil || utils.IsNil(o.Avatars) {
		var ret []ActorImage
		return ret
	}
	return o.Avatars
}

// GetAvatarsOk returns a tuple with the Avatars field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfigAboutInstance) GetAvatarsOk() ([]ActorImage, bool) {
	if o == nil || utils.IsNil(o.Avatars) {
		return nil, false
	}
	return o.Avatars, true
}

// HasAvatars returns a boolean if a field has been set.
func (o *ServerConfigAboutInstance) HasAvatars() bool {
	if o != nil && !utils.IsNil(o.Avatars) {
		return true
	}

	return false
}

// SetAvatars gets a reference to the given []ActorImage and assigns it to the Avatars field.
func (o *ServerConfigAboutInstance) SetAvatars(v []ActorImage) {
	o.Avatars = v
}

// GetBanners returns the Banners field value if set, zero value otherwise.
func (o *ServerConfigAboutInstance) GetBanners() []ActorImage {
	if o == nil || utils.IsNil(o.Banners) {
		var ret []ActorImage
		return ret
	}
	return o.Banners
}

// GetBannersOk returns a tuple with the Banners field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfigAboutInstance) GetBannersOk() ([]ActorImage, bool) {
	if o == nil || utils.IsNil(o.Banners) {
		return nil, false
	}
	return o.Banners, true
}

// HasBanners returns a boolean if a field has been set.
func (o *ServerConfigAboutInstance) HasBanners() bool {
	if o != nil && !utils.IsNil(o.Banners) {
		return true
	}

	return false
}

// SetBanners gets a reference to the given []ActorImage and assigns it to the Banners field.
func (o *ServerConfigAboutInstance) SetBanners(v []ActorImage) {
	o.Banners = v
}

func (o ServerConfigAboutInstance) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServerConfigAboutInstance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !utils.IsNil(o.ShortDescription) {
		toSerialize["shortDescription"] = o.ShortDescription
	}
	if !utils.IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !utils.IsNil(o.Terms) {
		toSerialize["terms"] = o.Terms
	}
	if !utils.IsNil(o.CodeOfConduct) {
		toSerialize["codeOfConduct"] = o.CodeOfConduct
	}
	if !utils.IsNil(o.HardwareInformation) {
		toSerialize["hardwareInformation"] = o.HardwareInformation
	}
	if !utils.IsNil(o.CreationReason) {
		toSerialize["creationReason"] = o.CreationReason
	}
	if !utils.IsNil(o.ModerationInformation) {
		toSerialize["moderationInformation"] = o.ModerationInformation
	}
	if !utils.IsNil(o.Administrator) {
		toSerialize["administrator"] = o.Administrator
	}
	if !utils.IsNil(o.MaintenanceLifetime) {
		toSerialize["maintenanceLifetime"] = o.MaintenanceLifetime
	}
	if !utils.IsNil(o.BusinessModel) {
		toSerialize["businessModel"] = o.BusinessModel
	}
	if !utils.IsNil(o.Languages) {
		toSerialize["languages"] = o.Languages
	}
	if !utils.IsNil(o.Categories) {
		toSerialize["categories"] = o.Categories
	}
	if !utils.IsNil(o.Avatars) {
		toSerialize["avatars"] = o.Avatars
	}
	if !utils.IsNil(o.Banners) {
		toSerialize["banners"] = o.Banners
	}
	return toSerialize, nil
}

type NullableServerConfigAboutInstance struct {
	value *ServerConfigAboutInstance
	isSet bool
}

func (v NullableServerConfigAboutInstance) Get() *ServerConfigAboutInstance {
	return v.value
}

func (v *NullableServerConfigAboutInstance) Set(val *ServerConfigAboutInstance) {
	v.value = val
	v.isSet = true
}

func (v NullableServerConfigAboutInstance) IsSet() bool {
	return v.isSet
}

func (v *NullableServerConfigAboutInstance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerConfigAboutInstance(val *ServerConfigAboutInstance) *NullableServerConfigAboutInstance {
	return &NullableServerConfigAboutInstance{value: val, isSet: true}
}

func (v NullableServerConfigAboutInstance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerConfigAboutInstance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
