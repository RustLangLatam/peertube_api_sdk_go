/*
PeerTube

The PeerTube API is built on HTTP(S) and is RESTful. You can use your favorite HTTP/REST library for your programming language to use PeerTube. The spec API is fully compatible with [openapi-generator](https://github.com/OpenAPITools/openapi-generator/wiki/API-client-generator-HOWTO) which generates a client SDK in the language of your choice - we generate some client SDKs automatically:  - [Python](https://framagit.org/framasoft/peertube/clients/python) - [Go](https://framagit.org/framasoft/peertube/clients/go) - [Kotlin](https://framagit.org/framasoft/peertube/clients/kotlin)  See the [REST API quick start](https://docs.joinpeertube.org/api/rest-getting-started) for a few examples of using the PeerTube API.  # Authentication  When you sign up for an account on a PeerTube instance, you are given the possibility to generate sessions on it, and authenticate there using an access token. Only __one access token can currently be used at a time__.  ## Roles  Accounts are given permissions based on their role. There are three roles on PeerTube: Administrator, Moderator, and User. See the [roles guide](https://docs.joinpeertube.org/admin/managing-users#roles) for a detail of their permissions.  # Errors  The API uses standard HTTP status codes to indicate the success or failure of the API call, completed by a [RFC7807-compliant](https://tools.ietf.org/html/rfc7807) response body.  ``` HTTP 1.1 404 Not Found Content-Type: application/problem+json; charset=utf-8  {   \"detail\": \"Video not found\",   \"docs\": \"https://docs.joinpeertube.org/api-rest-reference.html#operation/getVideo\",   \"status\": 404,   \"title\": \"Not Found\",   \"type\": \"about:blank\" } ```  We provide error `type` (following RFC7807) and `code` (internal PeerTube code) values for [a growing number of cases](https://github.com/Chocobozzz/PeerTube/blob/develop/packages/models/src/server/server-error-code.enum.ts), but it is still optional. Types are used to disambiguate errors that bear the same status code and are non-obvious:  ``` HTTP 1.1 403 Forbidden Content-Type: application/problem+json; charset=utf-8  {   \"detail\": \"Cannot get this video regarding follow constraints\",   \"docs\": \"https://docs.joinpeertube.org/api-rest-reference.html#operation/getVideo\",   \"status\": 403,   \"title\": \"Forbidden\",   \"type\": \"https://docs.joinpeertube.org/api-rest-reference.html#section/Errors/does_not_respect_follow_constraints\" } ```  Here a 403 error could otherwise mean that the video is private or blocklisted.  ### Validation errors  Each parameter is evaluated on its own against a set of rules before the route validator proceeds with potential testing involving parameter combinations. Errors coming from validation errors appear earlier and benefit from a more detailed error description:  ``` HTTP 1.1 400 Bad Request Content-Type: application/problem+json; charset=utf-8  {   \"detail\": \"Incorrect request parameters: id\",   \"docs\": \"https://docs.joinpeertube.org/api-rest-reference.html#operation/getVideo\",   \"instance\": \"/api/v1/videos/9c9de5e8-0a1e-484a-b099-e80766180\",   \"invalid-params\": {     \"id\": {       \"location\": \"params\",       \"msg\": \"Invalid value\",       \"param\": \"id\",       \"value\": \"9c9de5e8-0a1e-484a-b099-e80766180\"     }   },   \"status\": 400,   \"title\": \"Bad Request\",   \"type\": \"about:blank\" } ```  Where `id` is the name of the field concerned by the error, within the route definition. `invalid-params.<field>.location` can be either 'params', 'body', 'header', 'query' or 'cookies', and `invalid-params.<field>.value` reports the value that didn't pass validation whose `invalid-params.<field>.msg` is about.  ### Deprecated error fields  Some fields could be included with previous versions. They are still included but their use is deprecated: - `error`: superseded by `detail`  # Rate limits  We are rate-limiting all endpoints of PeerTube's API. Custom values can be set by administrators:  | Endpoint (prefix: `/api/v1`) | Calls         | Time frame   | |------------------------------|---------------|--------------| | `/_*`                         | 50            | 10 seconds   | | `POST /users/token`          | 15            | 5 minutes    | | `POST /users/register`       | 2<sup>*</sup> | 5 minutes    | | `POST /users/ask-send-verify-email` | 3      | 5 minutes    |  Depending on the endpoint, <sup>*</sup>failed requests are not taken into account. A service limit is announced by a `429 Too Many Requests` status code.  You can get details about the current state of your rate limit by reading the following headers:  | Header                  | Description                                                | |-------------------------|------------------------------------------------------------| | `X-RateLimit-Limit`     | Number of max requests allowed in the current time period  | | `X-RateLimit-Remaining` | Number of remaining requests in the current time period    | | `X-RateLimit-Reset`     | Timestamp of end of current time period as UNIX timestamp  | | `Retry-After`           | Seconds to delay after the first `429` is received         |  # CORS  This API features [Cross-Origin Resource Sharing (CORS)](https://fetch.spec.whatwg.org/), allowing cross-domain communication from the browser for some routes:  | Endpoint                    | |------------------------- ---| | `/api/_*`                    | | `/download/_*`               | | `/lazy-static/_*`            | | `/.well-known/webfinger`    |  In addition, all routes serving ActivityPub are CORS-enabled for all origins.

API version: 7.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package peertube_api_sdk

import (
	"encoding/json"
)

// checks if the ServerConfigCustomInstance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServerConfigCustomInstance{}

// ServerConfigCustomInstance struct for ServerConfigCustomInstance
type ServerConfigCustomInstance struct {
	Name                  *string                             `json:"name,omitempty"`
	ShortDescription      *string                             `json:"shortDescription,omitempty"`
	Description           *string                             `json:"description,omitempty"`
	Terms                 *string                             `json:"terms,omitempty"`
	CodeOfConduct         *string                             `json:"codeOfConduct,omitempty"`
	CreationReason        *string                             `json:"creationReason,omitempty"`
	ModerationInformation *string                             `json:"moderationInformation,omitempty"`
	Administrator         *string                             `json:"administrator,omitempty"`
	MaintenanceLifetime   *string                             `json:"maintenanceLifetime,omitempty"`
	BusinessModel         *string                             `json:"businessModel,omitempty"`
	HardwareInformation   *string                             `json:"hardwareInformation,omitempty"`
	Languages             []string                            `json:"languages,omitempty"`
	Categories            []float32                           `json:"categories,omitempty"`
	IsNSFW                *bool                               `json:"isNSFW,omitempty"`
	DefaultNSFWPolicy     *string                             `json:"defaultNSFWPolicy,omitempty"`
	ServerCountry         *string                             `json:"serverCountry,omitempty"`
	Support               *ServerConfigInstanceSupport        `json:"support,omitempty"`
	Social                *ServerConfigInstanceSocial         `json:"social,omitempty"`
	DefaultClientRoute    *string                             `json:"defaultClientRoute,omitempty"`
	Customizations        *ServerConfigInstanceCustomizations `json:"customizations,omitempty"`
}

// NewServerConfigCustomInstance instantiates a new ServerConfigCustomInstance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerConfigCustomInstance() *ServerConfigCustomInstance {
	this := ServerConfigCustomInstance{}
	return &this
}

// NewServerConfigCustomInstanceWithDefaults instantiates a new ServerConfigCustomInstance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerConfigCustomInstanceWithDefaults() *ServerConfigCustomInstance {
	this := ServerConfigCustomInstance{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ServerConfigCustomInstance) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfigCustomInstance) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ServerConfigCustomInstance) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ServerConfigCustomInstance) SetName(v string) {
	o.Name = &v
}

// GetShortDescription returns the ShortDescription field value if set, zero value otherwise.
func (o *ServerConfigCustomInstance) GetShortDescription() string {
	if o == nil || IsNil(o.ShortDescription) {
		var ret string
		return ret
	}
	return *o.ShortDescription
}

// GetShortDescriptionOk returns a tuple with the ShortDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfigCustomInstance) GetShortDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.ShortDescription) {
		return nil, false
	}
	return o.ShortDescription, true
}

// HasShortDescription returns a boolean if a field has been set.
func (o *ServerConfigCustomInstance) HasShortDescription() bool {
	if o != nil && !IsNil(o.ShortDescription) {
		return true
	}

	return false
}

// SetShortDescription gets a reference to the given string and assigns it to the ShortDescription field.
func (o *ServerConfigCustomInstance) SetShortDescription(v string) {
	o.ShortDescription = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ServerConfigCustomInstance) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfigCustomInstance) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ServerConfigCustomInstance) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ServerConfigCustomInstance) SetDescription(v string) {
	o.Description = &v
}

// GetTerms returns the Terms field value if set, zero value otherwise.
func (o *ServerConfigCustomInstance) GetTerms() string {
	if o == nil || IsNil(o.Terms) {
		var ret string
		return ret
	}
	return *o.Terms
}

// GetTermsOk returns a tuple with the Terms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfigCustomInstance) GetTermsOk() (*string, bool) {
	if o == nil || IsNil(o.Terms) {
		return nil, false
	}
	return o.Terms, true
}

// HasTerms returns a boolean if a field has been set.
func (o *ServerConfigCustomInstance) HasTerms() bool {
	if o != nil && !IsNil(o.Terms) {
		return true
	}

	return false
}

// SetTerms gets a reference to the given string and assigns it to the Terms field.
func (o *ServerConfigCustomInstance) SetTerms(v string) {
	o.Terms = &v
}

// GetCodeOfConduct returns the CodeOfConduct field value if set, zero value otherwise.
func (o *ServerConfigCustomInstance) GetCodeOfConduct() string {
	if o == nil || IsNil(o.CodeOfConduct) {
		var ret string
		return ret
	}
	return *o.CodeOfConduct
}

// GetCodeOfConductOk returns a tuple with the CodeOfConduct field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfigCustomInstance) GetCodeOfConductOk() (*string, bool) {
	if o == nil || IsNil(o.CodeOfConduct) {
		return nil, false
	}
	return o.CodeOfConduct, true
}

// HasCodeOfConduct returns a boolean if a field has been set.
func (o *ServerConfigCustomInstance) HasCodeOfConduct() bool {
	if o != nil && !IsNil(o.CodeOfConduct) {
		return true
	}

	return false
}

// SetCodeOfConduct gets a reference to the given string and assigns it to the CodeOfConduct field.
func (o *ServerConfigCustomInstance) SetCodeOfConduct(v string) {
	o.CodeOfConduct = &v
}

// GetCreationReason returns the CreationReason field value if set, zero value otherwise.
func (o *ServerConfigCustomInstance) GetCreationReason() string {
	if o == nil || IsNil(o.CreationReason) {
		var ret string
		return ret
	}
	return *o.CreationReason
}

// GetCreationReasonOk returns a tuple with the CreationReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfigCustomInstance) GetCreationReasonOk() (*string, bool) {
	if o == nil || IsNil(o.CreationReason) {
		return nil, false
	}
	return o.CreationReason, true
}

// HasCreationReason returns a boolean if a field has been set.
func (o *ServerConfigCustomInstance) HasCreationReason() bool {
	if o != nil && !IsNil(o.CreationReason) {
		return true
	}

	return false
}

// SetCreationReason gets a reference to the given string and assigns it to the CreationReason field.
func (o *ServerConfigCustomInstance) SetCreationReason(v string) {
	o.CreationReason = &v
}

// GetModerationInformation returns the ModerationInformation field value if set, zero value otherwise.
func (o *ServerConfigCustomInstance) GetModerationInformation() string {
	if o == nil || IsNil(o.ModerationInformation) {
		var ret string
		return ret
	}
	return *o.ModerationInformation
}

// GetModerationInformationOk returns a tuple with the ModerationInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfigCustomInstance) GetModerationInformationOk() (*string, bool) {
	if o == nil || IsNil(o.ModerationInformation) {
		return nil, false
	}
	return o.ModerationInformation, true
}

// HasModerationInformation returns a boolean if a field has been set.
func (o *ServerConfigCustomInstance) HasModerationInformation() bool {
	if o != nil && !IsNil(o.ModerationInformation) {
		return true
	}

	return false
}

// SetModerationInformation gets a reference to the given string and assigns it to the ModerationInformation field.
func (o *ServerConfigCustomInstance) SetModerationInformation(v string) {
	o.ModerationInformation = &v
}

// GetAdministrator returns the Administrator field value if set, zero value otherwise.
func (o *ServerConfigCustomInstance) GetAdministrator() string {
	if o == nil || IsNil(o.Administrator) {
		var ret string
		return ret
	}
	return *o.Administrator
}

// GetAdministratorOk returns a tuple with the Administrator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfigCustomInstance) GetAdministratorOk() (*string, bool) {
	if o == nil || IsNil(o.Administrator) {
		return nil, false
	}
	return o.Administrator, true
}

// HasAdministrator returns a boolean if a field has been set.
func (o *ServerConfigCustomInstance) HasAdministrator() bool {
	if o != nil && !IsNil(o.Administrator) {
		return true
	}

	return false
}

// SetAdministrator gets a reference to the given string and assigns it to the Administrator field.
func (o *ServerConfigCustomInstance) SetAdministrator(v string) {
	o.Administrator = &v
}

// GetMaintenanceLifetime returns the MaintenanceLifetime field value if set, zero value otherwise.
func (o *ServerConfigCustomInstance) GetMaintenanceLifetime() string {
	if o == nil || IsNil(o.MaintenanceLifetime) {
		var ret string
		return ret
	}
	return *o.MaintenanceLifetime
}

// GetMaintenanceLifetimeOk returns a tuple with the MaintenanceLifetime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfigCustomInstance) GetMaintenanceLifetimeOk() (*string, bool) {
	if o == nil || IsNil(o.MaintenanceLifetime) {
		return nil, false
	}
	return o.MaintenanceLifetime, true
}

// HasMaintenanceLifetime returns a boolean if a field has been set.
func (o *ServerConfigCustomInstance) HasMaintenanceLifetime() bool {
	if o != nil && !IsNil(o.MaintenanceLifetime) {
		return true
	}

	return false
}

// SetMaintenanceLifetime gets a reference to the given string and assigns it to the MaintenanceLifetime field.
func (o *ServerConfigCustomInstance) SetMaintenanceLifetime(v string) {
	o.MaintenanceLifetime = &v
}

// GetBusinessModel returns the BusinessModel field value if set, zero value otherwise.
func (o *ServerConfigCustomInstance) GetBusinessModel() string {
	if o == nil || IsNil(o.BusinessModel) {
		var ret string
		return ret
	}
	return *o.BusinessModel
}

// GetBusinessModelOk returns a tuple with the BusinessModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfigCustomInstance) GetBusinessModelOk() (*string, bool) {
	if o == nil || IsNil(o.BusinessModel) {
		return nil, false
	}
	return o.BusinessModel, true
}

// HasBusinessModel returns a boolean if a field has been set.
func (o *ServerConfigCustomInstance) HasBusinessModel() bool {
	if o != nil && !IsNil(o.BusinessModel) {
		return true
	}

	return false
}

// SetBusinessModel gets a reference to the given string and assigns it to the BusinessModel field.
func (o *ServerConfigCustomInstance) SetBusinessModel(v string) {
	o.BusinessModel = &v
}

// GetHardwareInformation returns the HardwareInformation field value if set, zero value otherwise.
func (o *ServerConfigCustomInstance) GetHardwareInformation() string {
	if o == nil || IsNil(o.HardwareInformation) {
		var ret string
		return ret
	}
	return *o.HardwareInformation
}

// GetHardwareInformationOk returns a tuple with the HardwareInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfigCustomInstance) GetHardwareInformationOk() (*string, bool) {
	if o == nil || IsNil(o.HardwareInformation) {
		return nil, false
	}
	return o.HardwareInformation, true
}

// HasHardwareInformation returns a boolean if a field has been set.
func (o *ServerConfigCustomInstance) HasHardwareInformation() bool {
	if o != nil && !IsNil(o.HardwareInformation) {
		return true
	}

	return false
}

// SetHardwareInformation gets a reference to the given string and assigns it to the HardwareInformation field.
func (o *ServerConfigCustomInstance) SetHardwareInformation(v string) {
	o.HardwareInformation = &v
}

// GetLanguages returns the Languages field value if set, zero value otherwise.
func (o *ServerConfigCustomInstance) GetLanguages() []string {
	if o == nil || IsNil(o.Languages) {
		var ret []string
		return ret
	}
	return o.Languages
}

// GetLanguagesOk returns a tuple with the Languages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfigCustomInstance) GetLanguagesOk() ([]string, bool) {
	if o == nil || IsNil(o.Languages) {
		return nil, false
	}
	return o.Languages, true
}

// HasLanguages returns a boolean if a field has been set.
func (o *ServerConfigCustomInstance) HasLanguages() bool {
	if o != nil && !IsNil(o.Languages) {
		return true
	}

	return false
}

// SetLanguages gets a reference to the given []string and assigns it to the Languages field.
func (o *ServerConfigCustomInstance) SetLanguages(v []string) {
	o.Languages = v
}

// GetCategories returns the Categories field value if set, zero value otherwise.
func (o *ServerConfigCustomInstance) GetCategories() []float32 {
	if o == nil || IsNil(o.Categories) {
		var ret []float32
		return ret
	}
	return o.Categories
}

// GetCategoriesOk returns a tuple with the Categories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfigCustomInstance) GetCategoriesOk() ([]float32, bool) {
	if o == nil || IsNil(o.Categories) {
		return nil, false
	}
	return o.Categories, true
}

// HasCategories returns a boolean if a field has been set.
func (o *ServerConfigCustomInstance) HasCategories() bool {
	if o != nil && !IsNil(o.Categories) {
		return true
	}

	return false
}

// SetCategories gets a reference to the given []float32 and assigns it to the Categories field.
func (o *ServerConfigCustomInstance) SetCategories(v []float32) {
	o.Categories = v
}

// GetIsNSFW returns the IsNSFW field value if set, zero value otherwise.
func (o *ServerConfigCustomInstance) GetIsNSFW() bool {
	if o == nil || IsNil(o.IsNSFW) {
		var ret bool
		return ret
	}
	return *o.IsNSFW
}

// GetIsNSFWOk returns a tuple with the IsNSFW field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfigCustomInstance) GetIsNSFWOk() (*bool, bool) {
	if o == nil || IsNil(o.IsNSFW) {
		return nil, false
	}
	return o.IsNSFW, true
}

// HasIsNSFW returns a boolean if a field has been set.
func (o *ServerConfigCustomInstance) HasIsNSFW() bool {
	if o != nil && !IsNil(o.IsNSFW) {
		return true
	}

	return false
}

// SetIsNSFW gets a reference to the given bool and assigns it to the IsNSFW field.
func (o *ServerConfigCustomInstance) SetIsNSFW(v bool) {
	o.IsNSFW = &v
}

// GetDefaultNSFWPolicy returns the DefaultNSFWPolicy field value if set, zero value otherwise.
func (o *ServerConfigCustomInstance) GetDefaultNSFWPolicy() string {
	if o == nil || IsNil(o.DefaultNSFWPolicy) {
		var ret string
		return ret
	}
	return *o.DefaultNSFWPolicy
}

// GetDefaultNSFWPolicyOk returns a tuple with the DefaultNSFWPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfigCustomInstance) GetDefaultNSFWPolicyOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultNSFWPolicy) {
		return nil, false
	}
	return o.DefaultNSFWPolicy, true
}

// HasDefaultNSFWPolicy returns a boolean if a field has been set.
func (o *ServerConfigCustomInstance) HasDefaultNSFWPolicy() bool {
	if o != nil && !IsNil(o.DefaultNSFWPolicy) {
		return true
	}

	return false
}

// SetDefaultNSFWPolicy gets a reference to the given string and assigns it to the DefaultNSFWPolicy field.
func (o *ServerConfigCustomInstance) SetDefaultNSFWPolicy(v string) {
	o.DefaultNSFWPolicy = &v
}

// GetServerCountry returns the ServerCountry field value if set, zero value otherwise.
func (o *ServerConfigCustomInstance) GetServerCountry() string {
	if o == nil || IsNil(o.ServerCountry) {
		var ret string
		return ret
	}
	return *o.ServerCountry
}

// GetServerCountryOk returns a tuple with the ServerCountry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfigCustomInstance) GetServerCountryOk() (*string, bool) {
	if o == nil || IsNil(o.ServerCountry) {
		return nil, false
	}
	return o.ServerCountry, true
}

// HasServerCountry returns a boolean if a field has been set.
func (o *ServerConfigCustomInstance) HasServerCountry() bool {
	if o != nil && !IsNil(o.ServerCountry) {
		return true
	}

	return false
}

// SetServerCountry gets a reference to the given string and assigns it to the ServerCountry field.
func (o *ServerConfigCustomInstance) SetServerCountry(v string) {
	o.ServerCountry = &v
}

// GetSupport returns the Support field value if set, zero value otherwise.
func (o *ServerConfigCustomInstance) GetSupport() ServerConfigInstanceSupport {
	if o == nil || IsNil(o.Support) {
		var ret ServerConfigInstanceSupport
		return ret
	}
	return *o.Support
}

// GetSupportOk returns a tuple with the Support field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfigCustomInstance) GetSupportOk() (*ServerConfigInstanceSupport, bool) {
	if o == nil || IsNil(o.Support) {
		return nil, false
	}
	return o.Support, true
}

// HasSupport returns a boolean if a field has been set.
func (o *ServerConfigCustomInstance) HasSupport() bool {
	if o != nil && !IsNil(o.Support) {
		return true
	}

	return false
}

// SetSupport gets a reference to the given ServerConfigInstanceSupport and assigns it to the Support field.
func (o *ServerConfigCustomInstance) SetSupport(v ServerConfigInstanceSupport) {
	o.Support = &v
}

// GetSocial returns the Social field value if set, zero value otherwise.
func (o *ServerConfigCustomInstance) GetSocial() ServerConfigInstanceSocial {
	if o == nil || IsNil(o.Social) {
		var ret ServerConfigInstanceSocial
		return ret
	}
	return *o.Social
}

// GetSocialOk returns a tuple with the Social field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfigCustomInstance) GetSocialOk() (*ServerConfigInstanceSocial, bool) {
	if o == nil || IsNil(o.Social) {
		return nil, false
	}
	return o.Social, true
}

// HasSocial returns a boolean if a field has been set.
func (o *ServerConfigCustomInstance) HasSocial() bool {
	if o != nil && !IsNil(o.Social) {
		return true
	}

	return false
}

// SetSocial gets a reference to the given ServerConfigInstanceSocial and assigns it to the Social field.
func (o *ServerConfigCustomInstance) SetSocial(v ServerConfigInstanceSocial) {
	o.Social = &v
}

// GetDefaultClientRoute returns the DefaultClientRoute field value if set, zero value otherwise.
func (o *ServerConfigCustomInstance) GetDefaultClientRoute() string {
	if o == nil || IsNil(o.DefaultClientRoute) {
		var ret string
		return ret
	}
	return *o.DefaultClientRoute
}

// GetDefaultClientRouteOk returns a tuple with the DefaultClientRoute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfigCustomInstance) GetDefaultClientRouteOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultClientRoute) {
		return nil, false
	}
	return o.DefaultClientRoute, true
}

// HasDefaultClientRoute returns a boolean if a field has been set.
func (o *ServerConfigCustomInstance) HasDefaultClientRoute() bool {
	if o != nil && !IsNil(o.DefaultClientRoute) {
		return true
	}

	return false
}

// SetDefaultClientRoute gets a reference to the given string and assigns it to the DefaultClientRoute field.
func (o *ServerConfigCustomInstance) SetDefaultClientRoute(v string) {
	o.DefaultClientRoute = &v
}

// GetCustomizations returns the Customizations field value if set, zero value otherwise.
func (o *ServerConfigCustomInstance) GetCustomizations() ServerConfigInstanceCustomizations {
	if o == nil || IsNil(o.Customizations) {
		var ret ServerConfigInstanceCustomizations
		return ret
	}
	return *o.Customizations
}

// GetCustomizationsOk returns a tuple with the Customizations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfigCustomInstance) GetCustomizationsOk() (*ServerConfigInstanceCustomizations, bool) {
	if o == nil || IsNil(o.Customizations) {
		return nil, false
	}
	return o.Customizations, true
}

// HasCustomizations returns a boolean if a field has been set.
func (o *ServerConfigCustomInstance) HasCustomizations() bool {
	if o != nil && !IsNil(o.Customizations) {
		return true
	}

	return false
}

// SetCustomizations gets a reference to the given ServerConfigInstanceCustomizations and assigns it to the Customizations field.
func (o *ServerConfigCustomInstance) SetCustomizations(v ServerConfigInstanceCustomizations) {
	o.Customizations = &v
}

func (o ServerConfigCustomInstance) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServerConfigCustomInstance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.ShortDescription) {
		toSerialize["shortDescription"] = o.ShortDescription
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Terms) {
		toSerialize["terms"] = o.Terms
	}
	if !IsNil(o.CodeOfConduct) {
		toSerialize["codeOfConduct"] = o.CodeOfConduct
	}
	if !IsNil(o.CreationReason) {
		toSerialize["creationReason"] = o.CreationReason
	}
	if !IsNil(o.ModerationInformation) {
		toSerialize["moderationInformation"] = o.ModerationInformation
	}
	if !IsNil(o.Administrator) {
		toSerialize["administrator"] = o.Administrator
	}
	if !IsNil(o.MaintenanceLifetime) {
		toSerialize["maintenanceLifetime"] = o.MaintenanceLifetime
	}
	if !IsNil(o.BusinessModel) {
		toSerialize["businessModel"] = o.BusinessModel
	}
	if !IsNil(o.HardwareInformation) {
		toSerialize["hardwareInformation"] = o.HardwareInformation
	}
	if !IsNil(o.Languages) {
		toSerialize["languages"] = o.Languages
	}
	if !IsNil(o.Categories) {
		toSerialize["categories"] = o.Categories
	}
	if !IsNil(o.IsNSFW) {
		toSerialize["isNSFW"] = o.IsNSFW
	}
	if !IsNil(o.DefaultNSFWPolicy) {
		toSerialize["defaultNSFWPolicy"] = o.DefaultNSFWPolicy
	}
	if !IsNil(o.ServerCountry) {
		toSerialize["serverCountry"] = o.ServerCountry
	}
	if !IsNil(o.Support) {
		toSerialize["support"] = o.Support
	}
	if !IsNil(o.Social) {
		toSerialize["social"] = o.Social
	}
	if !IsNil(o.DefaultClientRoute) {
		toSerialize["defaultClientRoute"] = o.DefaultClientRoute
	}
	if !IsNil(o.Customizations) {
		toSerialize["customizations"] = o.Customizations
	}
	return toSerialize, nil
}

type NullableServerConfigCustomInstance struct {
	value *ServerConfigCustomInstance
	isSet bool
}

func (v NullableServerConfigCustomInstance) Get() *ServerConfigCustomInstance {
	return v.value
}

func (v *NullableServerConfigCustomInstance) Set(val *ServerConfigCustomInstance) {
	v.value = val
	v.isSet = true
}

func (v NullableServerConfigCustomInstance) IsSet() bool {
	return v.isSet
}

func (v *NullableServerConfigCustomInstance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerConfigCustomInstance(val *ServerConfigCustomInstance) *NullableServerConfigCustomInstance {
	return &NullableServerConfigCustomInstance{value: val, isSet: true}
}

func (v NullableServerConfigCustomInstance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerConfigCustomInstance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
