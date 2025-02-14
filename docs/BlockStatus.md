# BlockStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accounts** | Pointer to [**map[string]BlockStatusAccountsValue**](BlockStatusAccountsValue.md) |  | [optional] 
**Hosts** | Pointer to [**map[string]BlockStatusHostsValue**](BlockStatusHostsValue.md) |  | [optional] 

## Methods

### NewBlockStatus

`func NewBlockStatus() *BlockStatus`

NewBlockStatus instantiates a new BlockStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlockStatusWithDefaults

`func NewBlockStatusWithDefaults() *BlockStatus`

NewBlockStatusWithDefaults instantiates a new BlockStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccounts

`func (o *BlockStatus) GetAccounts() map[string]BlockStatusAccountsValue`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *BlockStatus) GetAccountsOk() (*map[string]BlockStatusAccountsValue, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *BlockStatus) SetAccounts(v map[string]BlockStatusAccountsValue)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *BlockStatus) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### GetHosts

`func (o *BlockStatus) GetHosts() map[string]BlockStatusHostsValue`

GetHosts returns the Hosts field if non-nil, zero value otherwise.

### GetHostsOk

`func (o *BlockStatus) GetHostsOk() (*map[string]BlockStatusHostsValue, bool)`

GetHostsOk returns a tuple with the Hosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosts

`func (o *BlockStatus) SetHosts(v map[string]BlockStatusHostsValue)`

SetHosts sets Hosts field to given value.

### HasHosts

`func (o *BlockStatus) HasHosts() bool`

HasHosts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


