/*
PeerTube

The PeerTube API is built on HTTP(S) and is RESTful. You can use your favorite HTTP/REST library for your programming language to use PeerTube. The spec API is fully compatible with [openapi-generator](https://github.com/OpenAPITools/openapi-generator/wiki/API-client-generator-HOWTO) which generates a client SDK in the language of your choice - we generate some client SDKs automatically:  - [Python](https://framagit.org/framasoft/peertube/clients/python) - [Go](https://framagit.org/framasoft/peertube/clients/go) - [Kotlin](https://framagit.org/framasoft/peertube/clients/kotlin)  See the [REST API quick start](https://docs.joinpeertube.org/api/rest-getting-started) for a few examples of using the PeerTube API.  # Authentication  When you sign up for an account on a PeerTube instance, you are given the possibility to generate sessions on it, and authenticate there using an access token. Only __one access token can currently be used at a time__.  ## Roles  Accounts are given permissions based on their role. There are three roles on PeerTube: Administrator, Moderator, and User. See the [roles guide](https://docs.joinpeertube.org/admin/managing-users#roles) for a detail of their permissions.  # Errors  The API uses standard HTTP status codes to indicate the success or failure of the API call, completed by a [RFC7807-compliant](https://tools.ietf.org/html/rfc7807) response body.  ``` HTTP 1.1 404 Not Found Content-Type: application/problem+json; charset=utf-8  {   \"detail\": \"Video not found\",   \"docs\": \"https://docs.joinpeertube.org/api-rest-reference.html#operation/getVideo\",   \"status\": 404,   \"title\": \"Not Found\",   \"type\": \"about:blank\" } ```  We provide error `type` (following RFC7807) and `code` (internal PeerTube code) values for [a growing number of cases](https://github.com/Chocobozzz/PeerTube/blob/develop/packages/models/src/server/server-error-code.enum.ts), but it is still optional. Types are used to disambiguate errors that bear the same status code and are non-obvious:  ``` HTTP 1.1 403 Forbidden Content-Type: application/problem+json; charset=utf-8  {   \"detail\": \"Cannot get this video regarding follow constraints\",   \"docs\": \"https://docs.joinpeertube.org/api-rest-reference.html#operation/getVideo\",   \"status\": 403,   \"title\": \"Forbidden\",   \"type\": \"https://docs.joinpeertube.org/api-rest-reference.html#section/Errors/does_not_respect_follow_constraints\" } ```  Here a 403 error could otherwise mean that the video is private or blocklisted.  ### Validation errors  Each parameter is evaluated on its own against a set of rules before the route validator proceeds with potential testing involving parameter combinations. Errors coming from validation errors appear earlier and benefit from a more detailed error description:  ``` HTTP 1.1 400 Bad Request Content-Type: application/problem+json; charset=utf-8  {   \"detail\": \"Incorrect request parameters: id\",   \"docs\": \"https://docs.joinpeertube.org/api-rest-reference.html#operation/getVideo\",   \"instance\": \"/api/v1/videos/9c9de5e8-0a1e-484a-b099-e80766180\",   \"invalid-params\": {     \"id\": {       \"location\": \"params\",       \"msg\": \"Invalid value\",       \"param\": \"id\",       \"value\": \"9c9de5e8-0a1e-484a-b099-e80766180\"     }   },   \"status\": 400,   \"title\": \"Bad Request\",   \"type\": \"about:blank\" } ```  Where `id` is the name of the field concerned by the error, within the route definition. `invalid-params.<field>.location` can be either 'params', 'body', 'header', 'query' or 'cookies', and `invalid-params.<field>.value` reports the value that didn't pass validation whose `invalid-params.<field>.msg` is about.  ### Deprecated error fields  Some fields could be included with previous versions. They are still included but their use is deprecated: - `error`: superseded by `detail`  # Rate limits  We are rate-limiting all endpoints of PeerTube's API. Custom values can be set by administrators:  | Endpoint (prefix: `/api/v1`) | Calls         | Time frame   | |------------------------------|---------------|--------------| | `/_*`                         | 50            | 10 seconds   | | `POST /users/token`          | 15            | 5 minutes    | | `POST /users/register`       | 2<sup>*</sup> | 5 minutes    | | `POST /users/ask-send-verify-email` | 3      | 5 minutes    |  Depending on the endpoint, <sup>*</sup>failed requests are not taken into account. A service limit is announced by a `429 Too Many Requests` status code.  You can get details about the current state of your rate limit by reading the following headers:  | Header                  | Description                                                | |-------------------------|------------------------------------------------------------| | `X-RateLimit-Limit`     | Number of max requests allowed in the current time period  | | `X-RateLimit-Remaining` | Number of remaining requests in the current time period    | | `X-RateLimit-Reset`     | Timestamp of end of current time period as UNIX timestamp  | | `Retry-After`           | Seconds to delay after the first `429` is received         |  # CORS  This API features [Cross-Origin Resource Sharing (CORS)](https://fetch.spec.whatwg.org/), allowing cross-domain communication from the browser for some routes:  | Endpoint                    | |------------------------- ---| | `/api/_*`                    | | `/download/_*`               | | `/lazy-static/_*`            | | `/.well-known/webfinger`    |  In addition, all routes serving ActivityPub are CORS-enabled for all origins.

API version: 7.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package peertube_api_sdk

import (
	"encoding/json"
	"time"
)

// checks if the Plugin type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Plugin{}

// Plugin struct for Plugin
type Plugin struct {
	Name *string `json:"name,omitempty"`
	// - `1`: PLUGIN - `2`: THEME
	Type           *int32                 `json:"type,omitempty"`
	LatestVersion  *string                `json:"latestVersion,omitempty"`
	Version        *string                `json:"version,omitempty"`
	Enabled        *bool                  `json:"enabled,omitempty"`
	Uninstalled    *bool                  `json:"uninstalled,omitempty"`
	PeertubeEngine *string                `json:"peertubeEngine,omitempty"`
	Description    *string                `json:"description,omitempty"`
	Homepage       *string                `json:"homepage,omitempty"`
	Settings       map[string]interface{} `json:"settings,omitempty"`
	CreatedAt      *time.Time             `json:"createdAt,omitempty"`
	UpdatedAt      *time.Time             `json:"updatedAt,omitempty"`
}

// NewPlugin instantiates a new Plugin object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlugin() *Plugin {
	this := Plugin{}
	return &this
}

// NewPluginWithDefaults instantiates a new Plugin object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPluginWithDefaults() *Plugin {
	this := Plugin{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Plugin) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Plugin) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Plugin) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Plugin) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Plugin) GetType() int32 {
	if o == nil || IsNil(o.Type) {
		var ret int32
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Plugin) GetTypeOk() (*int32, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Plugin) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given int32 and assigns it to the Type field.
func (o *Plugin) SetType(v int32) {
	o.Type = &v
}

// GetLatestVersion returns the LatestVersion field value if set, zero value otherwise.
func (o *Plugin) GetLatestVersion() string {
	if o == nil || IsNil(o.LatestVersion) {
		var ret string
		return ret
	}
	return *o.LatestVersion
}

// GetLatestVersionOk returns a tuple with the LatestVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Plugin) GetLatestVersionOk() (*string, bool) {
	if o == nil || IsNil(o.LatestVersion) {
		return nil, false
	}
	return o.LatestVersion, true
}

// HasLatestVersion returns a boolean if a field has been set.
func (o *Plugin) HasLatestVersion() bool {
	if o != nil && !IsNil(o.LatestVersion) {
		return true
	}

	return false
}

// SetLatestVersion gets a reference to the given string and assigns it to the LatestVersion field.
func (o *Plugin) SetLatestVersion(v string) {
	o.LatestVersion = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *Plugin) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Plugin) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *Plugin) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *Plugin) SetVersion(v string) {
	o.Version = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *Plugin) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Plugin) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *Plugin) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *Plugin) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetUninstalled returns the Uninstalled field value if set, zero value otherwise.
func (o *Plugin) GetUninstalled() bool {
	if o == nil || IsNil(o.Uninstalled) {
		var ret bool
		return ret
	}
	return *o.Uninstalled
}

// GetUninstalledOk returns a tuple with the Uninstalled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Plugin) GetUninstalledOk() (*bool, bool) {
	if o == nil || IsNil(o.Uninstalled) {
		return nil, false
	}
	return o.Uninstalled, true
}

// HasUninstalled returns a boolean if a field has been set.
func (o *Plugin) HasUninstalled() bool {
	if o != nil && !IsNil(o.Uninstalled) {
		return true
	}

	return false
}

// SetUninstalled gets a reference to the given bool and assigns it to the Uninstalled field.
func (o *Plugin) SetUninstalled(v bool) {
	o.Uninstalled = &v
}

// GetPeertubeEngine returns the PeertubeEngine field value if set, zero value otherwise.
func (o *Plugin) GetPeertubeEngine() string {
	if o == nil || IsNil(o.PeertubeEngine) {
		var ret string
		return ret
	}
	return *o.PeertubeEngine
}

// GetPeertubeEngineOk returns a tuple with the PeertubeEngine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Plugin) GetPeertubeEngineOk() (*string, bool) {
	if o == nil || IsNil(o.PeertubeEngine) {
		return nil, false
	}
	return o.PeertubeEngine, true
}

// HasPeertubeEngine returns a boolean if a field has been set.
func (o *Plugin) HasPeertubeEngine() bool {
	if o != nil && !IsNil(o.PeertubeEngine) {
		return true
	}

	return false
}

// SetPeertubeEngine gets a reference to the given string and assigns it to the PeertubeEngine field.
func (o *Plugin) SetPeertubeEngine(v string) {
	o.PeertubeEngine = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Plugin) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Plugin) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Plugin) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Plugin) SetDescription(v string) {
	o.Description = &v
}

// GetHomepage returns the Homepage field value if set, zero value otherwise.
func (o *Plugin) GetHomepage() string {
	if o == nil || IsNil(o.Homepage) {
		var ret string
		return ret
	}
	return *o.Homepage
}

// GetHomepageOk returns a tuple with the Homepage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Plugin) GetHomepageOk() (*string, bool) {
	if o == nil || IsNil(o.Homepage) {
		return nil, false
	}
	return o.Homepage, true
}

// HasHomepage returns a boolean if a field has been set.
func (o *Plugin) HasHomepage() bool {
	if o != nil && !IsNil(o.Homepage) {
		return true
	}

	return false
}

// SetHomepage gets a reference to the given string and assigns it to the Homepage field.
func (o *Plugin) SetHomepage(v string) {
	o.Homepage = &v
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *Plugin) GetSettings() map[string]interface{} {
	if o == nil || IsNil(o.Settings) {
		var ret map[string]interface{}
		return ret
	}
	return o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Plugin) GetSettingsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Settings) {
		return map[string]interface{}{}, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *Plugin) HasSettings() bool {
	if o != nil && !IsNil(o.Settings) {
		return true
	}

	return false
}

// SetSettings gets a reference to the given map[string]interface{} and assigns it to the Settings field.
func (o *Plugin) SetSettings(v map[string]interface{}) {
	o.Settings = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Plugin) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Plugin) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Plugin) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *Plugin) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Plugin) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Plugin) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Plugin) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *Plugin) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o Plugin) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Plugin) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.LatestVersion) {
		toSerialize["latestVersion"] = o.LatestVersion
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Uninstalled) {
		toSerialize["uninstalled"] = o.Uninstalled
	}
	if !IsNil(o.PeertubeEngine) {
		toSerialize["peertubeEngine"] = o.PeertubeEngine
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Homepage) {
		toSerialize["homepage"] = o.Homepage
	}
	if !IsNil(o.Settings) {
		toSerialize["settings"] = o.Settings
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	return toSerialize, nil
}

type NullablePlugin struct {
	value *Plugin
	isSet bool
}

func (v NullablePlugin) Get() *Plugin {
	return v.value
}

func (v *NullablePlugin) Set(val *Plugin) {
	v.value = val
	v.isSet = true
}

func (v NullablePlugin) IsSet() bool {
	return v.isSet
}

func (v *NullablePlugin) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlugin(val *Plugin) *NullablePlugin {
	return &NullablePlugin{value: val, isSet: true}
}

func (v NullablePlugin) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlugin) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
