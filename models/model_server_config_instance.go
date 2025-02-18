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

// checks if the ServerConfigInstance type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ServerConfigInstance{}

// ServerConfigInstance struct for ServerConfigInstance
type ServerConfigInstance struct {
	Name               *string                             `json:"name,omitempty"`
	ShortDescription   *string                             `json:"shortDescription,omitempty"`
	DefaultClientRoute *string                             `json:"defaultClientRoute,omitempty"`
	IsNSFW             *bool                               `json:"isNSFW,omitempty"`
	DefaultNSFWPolicy  *string                             `json:"defaultNSFWPolicy,omitempty"`
	ServerCountry      *string                             `json:"serverCountry,omitempty"`
	Support            *ServerConfigInstanceSupport        `json:"support,omitempty"`
	Social             *ServerConfigInstanceSocial         `json:"social,omitempty"`
	Customizations     *ServerConfigInstanceCustomizations `json:"customizations,omitempty"`
	Avatars            []ActorImage                        `json:"avatars,omitempty"`
	Banners            []ActorImage                        `json:"banners,omitempty"`
}

// NewServerConfigInstance instantiates a new ServerConfigInstance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerConfigInstance() *ServerConfigInstance {
	this := ServerConfigInstance{}
	return &this
}

// NewServerConfigInstanceWithDefaults instantiates a new ServerConfigInstance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerConfigInstanceWithDefaults() *ServerConfigInstance {
	this := ServerConfigInstance{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ServerConfigInstance) GetName() string {
	if o == nil || utils.IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfigInstance) GetNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ServerConfigInstance) HasName() bool {
	if o != nil && !utils.IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ServerConfigInstance) SetName(v string) {
	o.Name = &v
}

// GetShortDescription returns the ShortDescription field value if set, zero value otherwise.
func (o *ServerConfigInstance) GetShortDescription() string {
	if o == nil || utils.IsNil(o.ShortDescription) {
		var ret string
		return ret
	}
	return *o.ShortDescription
}

// GetShortDescriptionOk returns a tuple with the ShortDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfigInstance) GetShortDescriptionOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ShortDescription) {
		return nil, false
	}
	return o.ShortDescription, true
}

// HasShortDescription returns a boolean if a field has been set.
func (o *ServerConfigInstance) HasShortDescription() bool {
	if o != nil && !utils.IsNil(o.ShortDescription) {
		return true
	}

	return false
}

// SetShortDescription gets a reference to the given string and assigns it to the ShortDescription field.
func (o *ServerConfigInstance) SetShortDescription(v string) {
	o.ShortDescription = &v
}

// GetDefaultClientRoute returns the DefaultClientRoute field value if set, zero value otherwise.
func (o *ServerConfigInstance) GetDefaultClientRoute() string {
	if o == nil || utils.IsNil(o.DefaultClientRoute) {
		var ret string
		return ret
	}
	return *o.DefaultClientRoute
}

// GetDefaultClientRouteOk returns a tuple with the DefaultClientRoute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfigInstance) GetDefaultClientRouteOk() (*string, bool) {
	if o == nil || utils.IsNil(o.DefaultClientRoute) {
		return nil, false
	}
	return o.DefaultClientRoute, true
}

// HasDefaultClientRoute returns a boolean if a field has been set.
func (o *ServerConfigInstance) HasDefaultClientRoute() bool {
	if o != nil && !utils.IsNil(o.DefaultClientRoute) {
		return true
	}

	return false
}

// SetDefaultClientRoute gets a reference to the given string and assigns it to the DefaultClientRoute field.
func (o *ServerConfigInstance) SetDefaultClientRoute(v string) {
	o.DefaultClientRoute = &v
}

// GetIsNSFW returns the IsNSFW field value if set, zero value otherwise.
func (o *ServerConfigInstance) GetIsNSFW() bool {
	if o == nil || utils.IsNil(o.IsNSFW) {
		var ret bool
		return ret
	}
	return *o.IsNSFW
}

// GetIsNSFWOk returns a tuple with the IsNSFW field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfigInstance) GetIsNSFWOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.IsNSFW) {
		return nil, false
	}
	return o.IsNSFW, true
}

// HasIsNSFW returns a boolean if a field has been set.
func (o *ServerConfigInstance) HasIsNSFW() bool {
	if o != nil && !utils.IsNil(o.IsNSFW) {
		return true
	}

	return false
}

// SetIsNSFW gets a reference to the given bool and assigns it to the IsNSFW field.
func (o *ServerConfigInstance) SetIsNSFW(v bool) {
	o.IsNSFW = &v
}

// GetDefaultNSFWPolicy returns the DefaultNSFWPolicy field value if set, zero value otherwise.
func (o *ServerConfigInstance) GetDefaultNSFWPolicy() string {
	if o == nil || utils.IsNil(o.DefaultNSFWPolicy) {
		var ret string
		return ret
	}
	return *o.DefaultNSFWPolicy
}

// GetDefaultNSFWPolicyOk returns a tuple with the DefaultNSFWPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfigInstance) GetDefaultNSFWPolicyOk() (*string, bool) {
	if o == nil || utils.IsNil(o.DefaultNSFWPolicy) {
		return nil, false
	}
	return o.DefaultNSFWPolicy, true
}

// HasDefaultNSFWPolicy returns a boolean if a field has been set.
func (o *ServerConfigInstance) HasDefaultNSFWPolicy() bool {
	if o != nil && !utils.IsNil(o.DefaultNSFWPolicy) {
		return true
	}

	return false
}

// SetDefaultNSFWPolicy gets a reference to the given string and assigns it to the DefaultNSFWPolicy field.
func (o *ServerConfigInstance) SetDefaultNSFWPolicy(v string) {
	o.DefaultNSFWPolicy = &v
}

// GetServerCountry returns the ServerCountry field value if set, zero value otherwise.
func (o *ServerConfigInstance) GetServerCountry() string {
	if o == nil || utils.IsNil(o.ServerCountry) {
		var ret string
		return ret
	}
	return *o.ServerCountry
}

// GetServerCountryOk returns a tuple with the ServerCountry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfigInstance) GetServerCountryOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ServerCountry) {
		return nil, false
	}
	return o.ServerCountry, true
}

// HasServerCountry returns a boolean if a field has been set.
func (o *ServerConfigInstance) HasServerCountry() bool {
	if o != nil && !utils.IsNil(o.ServerCountry) {
		return true
	}

	return false
}

// SetServerCountry gets a reference to the given string and assigns it to the ServerCountry field.
func (o *ServerConfigInstance) SetServerCountry(v string) {
	o.ServerCountry = &v
}

// GetSupport returns the Support field value if set, zero value otherwise.
func (o *ServerConfigInstance) GetSupport() ServerConfigInstanceSupport {
	if o == nil || utils.IsNil(o.Support) {
		var ret ServerConfigInstanceSupport
		return ret
	}
	return *o.Support
}

// GetSupportOk returns a tuple with the Support field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfigInstance) GetSupportOk() (*ServerConfigInstanceSupport, bool) {
	if o == nil || utils.IsNil(o.Support) {
		return nil, false
	}
	return o.Support, true
}

// HasSupport returns a boolean if a field has been set.
func (o *ServerConfigInstance) HasSupport() bool {
	if o != nil && !utils.IsNil(o.Support) {
		return true
	}

	return false
}

// SetSupport gets a reference to the given ServerConfigInstanceSupport and assigns it to the Support field.
func (o *ServerConfigInstance) SetSupport(v ServerConfigInstanceSupport) {
	o.Support = &v
}

// GetSocial returns the Social field value if set, zero value otherwise.
func (o *ServerConfigInstance) GetSocial() ServerConfigInstanceSocial {
	if o == nil || utils.IsNil(o.Social) {
		var ret ServerConfigInstanceSocial
		return ret
	}
	return *o.Social
}

// GetSocialOk returns a tuple with the Social field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfigInstance) GetSocialOk() (*ServerConfigInstanceSocial, bool) {
	if o == nil || utils.IsNil(o.Social) {
		return nil, false
	}
	return o.Social, true
}

// HasSocial returns a boolean if a field has been set.
func (o *ServerConfigInstance) HasSocial() bool {
	if o != nil && !utils.IsNil(o.Social) {
		return true
	}

	return false
}

// SetSocial gets a reference to the given ServerConfigInstanceSocial and assigns it to the Social field.
func (o *ServerConfigInstance) SetSocial(v ServerConfigInstanceSocial) {
	o.Social = &v
}

// GetCustomizations returns the Customizations field value if set, zero value otherwise.
func (o *ServerConfigInstance) GetCustomizations() ServerConfigInstanceCustomizations {
	if o == nil || utils.IsNil(o.Customizations) {
		var ret ServerConfigInstanceCustomizations
		return ret
	}
	return *o.Customizations
}

// GetCustomizationsOk returns a tuple with the Customizations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfigInstance) GetCustomizationsOk() (*ServerConfigInstanceCustomizations, bool) {
	if o == nil || utils.IsNil(o.Customizations) {
		return nil, false
	}
	return o.Customizations, true
}

// HasCustomizations returns a boolean if a field has been set.
func (o *ServerConfigInstance) HasCustomizations() bool {
	if o != nil && !utils.IsNil(o.Customizations) {
		return true
	}

	return false
}

// SetCustomizations gets a reference to the given ServerConfigInstanceCustomizations and assigns it to the Customizations field.
func (o *ServerConfigInstance) SetCustomizations(v ServerConfigInstanceCustomizations) {
	o.Customizations = &v
}

// GetAvatars returns the Avatars field value if set, zero value otherwise.
func (o *ServerConfigInstance) GetAvatars() []ActorImage {
	if o == nil || utils.IsNil(o.Avatars) {
		var ret []ActorImage
		return ret
	}
	return o.Avatars
}

// GetAvatarsOk returns a tuple with the Avatars field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfigInstance) GetAvatarsOk() ([]ActorImage, bool) {
	if o == nil || utils.IsNil(o.Avatars) {
		return nil, false
	}
	return o.Avatars, true
}

// HasAvatars returns a boolean if a field has been set.
func (o *ServerConfigInstance) HasAvatars() bool {
	if o != nil && !utils.IsNil(o.Avatars) {
		return true
	}

	return false
}

// SetAvatars gets a reference to the given []ActorImage and assigns it to the Avatars field.
func (o *ServerConfigInstance) SetAvatars(v []ActorImage) {
	o.Avatars = v
}

// GetBanners returns the Banners field value if set, zero value otherwise.
func (o *ServerConfigInstance) GetBanners() []ActorImage {
	if o == nil || utils.IsNil(o.Banners) {
		var ret []ActorImage
		return ret
	}
	return o.Banners
}

// GetBannersOk returns a tuple with the Banners field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfigInstance) GetBannersOk() ([]ActorImage, bool) {
	if o == nil || utils.IsNil(o.Banners) {
		return nil, false
	}
	return o.Banners, true
}

// HasBanners returns a boolean if a field has been set.
func (o *ServerConfigInstance) HasBanners() bool {
	if o != nil && !utils.IsNil(o.Banners) {
		return true
	}

	return false
}

// SetBanners gets a reference to the given []ActorImage and assigns it to the Banners field.
func (o *ServerConfigInstance) SetBanners(v []ActorImage) {
	o.Banners = v
}

func (o ServerConfigInstance) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServerConfigInstance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !utils.IsNil(o.ShortDescription) {
		toSerialize["shortDescription"] = o.ShortDescription
	}
	if !utils.IsNil(o.DefaultClientRoute) {
		toSerialize["defaultClientRoute"] = o.DefaultClientRoute
	}
	if !utils.IsNil(o.IsNSFW) {
		toSerialize["isNSFW"] = o.IsNSFW
	}
	if !utils.IsNil(o.DefaultNSFWPolicy) {
		toSerialize["defaultNSFWPolicy"] = o.DefaultNSFWPolicy
	}
	if !utils.IsNil(o.ServerCountry) {
		toSerialize["serverCountry"] = o.ServerCountry
	}
	if !utils.IsNil(o.Support) {
		toSerialize["support"] = o.Support
	}
	if !utils.IsNil(o.Social) {
		toSerialize["social"] = o.Social
	}
	if !utils.IsNil(o.Customizations) {
		toSerialize["customizations"] = o.Customizations
	}
	if !utils.IsNil(o.Avatars) {
		toSerialize["avatars"] = o.Avatars
	}
	if !utils.IsNil(o.Banners) {
		toSerialize["banners"] = o.Banners
	}
	return toSerialize, nil
}

type NullableServerConfigInstance struct {
	value *ServerConfigInstance
	isSet bool
}

func (v NullableServerConfigInstance) Get() *ServerConfigInstance {
	return v.value
}

func (v *NullableServerConfigInstance) Set(val *ServerConfigInstance) {
	v.value = val
	v.isSet = true
}

func (v NullableServerConfigInstance) IsSet() bool {
	return v.isSet
}

func (v *NullableServerConfigInstance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerConfigInstance(val *ServerConfigInstance) *NullableServerConfigInstance {
	return &NullableServerConfigInstance{value: val, isSet: true}
}

func (v NullableServerConfigInstance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerConfigInstance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
