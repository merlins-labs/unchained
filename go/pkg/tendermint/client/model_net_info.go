/*
Tendermint RPC

Tendermint supports the following RPC protocols:  * URI over HTTP * JSONRPC over HTTP * JSONRPC over websockets  ## Configuration  RPC can be configured by tuning parameters under `[rpc]` table in the `$TMHOME/config/config.toml` file or by using the `--rpc.X` command-line flags.  Default rpc listen address is `tcp://0.0.0.0:26657`. To set another address, set the `laddr` config parameter to desired value. CORS (Cross-Origin Resource Sharing) can be enabled by setting `cors_allowed_origins`, `cors_allowed_methods`, `cors_allowed_headers` config parameters.  ## Arguments  Arguments which expect strings or byte arrays may be passed as quoted strings, like `\"abc\"` or as `0x`-prefixed strings, like `0x616263`.  ## URI/HTTP  A REST like interface.      curl localhost:26657/block?height=5  ## JSONRPC/HTTP  JSONRPC requests can be POST'd to the root RPC endpoint via HTTP.      curl --header \"Content-Type: application/json\" --request POST --data '{\"method\": \"block\", \"params\": [\"5\"], \"id\": 1}' localhost:26657  ## JSONRPC/websockets  JSONRPC requests can be also made via websocket. The websocket endpoint is at `/websocket`, e.g. `localhost:26657/websocket`. Asynchronous RPC functions like event `subscribe` and `unsubscribe` are only available via websockets.  Example using https://github.com/hashrocket/ws:      ws ws://localhost:26657/websocket     > { \"jsonrpc\": \"2.0\", \"method\": \"subscribe\", \"params\": [\"tm.event='NewBlock'\"], \"id\": 1 } 

API version: Master
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetInfo struct for NetInfo
type NetInfo struct {
	Listening *bool `json:"listening,omitempty"`
	Listeners []string `json:"listeners,omitempty"`
	NPeers *string `json:"n_peers,omitempty"`
	Peers []Peer `json:"peers,omitempty"`
}

// NewNetInfo instantiates a new NetInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetInfo() *NetInfo {
	this := NetInfo{}
	return &this
}

// NewNetInfoWithDefaults instantiates a new NetInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetInfoWithDefaults() *NetInfo {
	this := NetInfo{}
	return &this
}

// GetListening returns the Listening field value if set, zero value otherwise.
func (o *NetInfo) GetListening() bool {
	if o == nil || o.Listening == nil {
		var ret bool
		return ret
	}
	return *o.Listening
}

// GetListeningOk returns a tuple with the Listening field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetInfo) GetListeningOk() (*bool, bool) {
	if o == nil || o.Listening == nil {
		return nil, false
	}
	return o.Listening, true
}

// HasListening returns a boolean if a field has been set.
func (o *NetInfo) HasListening() bool {
	if o != nil && o.Listening != nil {
		return true
	}

	return false
}

// SetListening gets a reference to the given bool and assigns it to the Listening field.
func (o *NetInfo) SetListening(v bool) {
	o.Listening = &v
}

// GetListeners returns the Listeners field value if set, zero value otherwise.
func (o *NetInfo) GetListeners() []string {
	if o == nil || o.Listeners == nil {
		var ret []string
		return ret
	}
	return o.Listeners
}

// GetListenersOk returns a tuple with the Listeners field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetInfo) GetListenersOk() ([]string, bool) {
	if o == nil || o.Listeners == nil {
		return nil, false
	}
	return o.Listeners, true
}

// HasListeners returns a boolean if a field has been set.
func (o *NetInfo) HasListeners() bool {
	if o != nil && o.Listeners != nil {
		return true
	}

	return false
}

// SetListeners gets a reference to the given []string and assigns it to the Listeners field.
func (o *NetInfo) SetListeners(v []string) {
	o.Listeners = v
}

// GetNPeers returns the NPeers field value if set, zero value otherwise.
func (o *NetInfo) GetNPeers() string {
	if o == nil || o.NPeers == nil {
		var ret string
		return ret
	}
	return *o.NPeers
}

// GetNPeersOk returns a tuple with the NPeers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetInfo) GetNPeersOk() (*string, bool) {
	if o == nil || o.NPeers == nil {
		return nil, false
	}
	return o.NPeers, true
}

// HasNPeers returns a boolean if a field has been set.
func (o *NetInfo) HasNPeers() bool {
	if o != nil && o.NPeers != nil {
		return true
	}

	return false
}

// SetNPeers gets a reference to the given string and assigns it to the NPeers field.
func (o *NetInfo) SetNPeers(v string) {
	o.NPeers = &v
}

// GetPeers returns the Peers field value if set, zero value otherwise.
func (o *NetInfo) GetPeers() []Peer {
	if o == nil || o.Peers == nil {
		var ret []Peer
		return ret
	}
	return o.Peers
}

// GetPeersOk returns a tuple with the Peers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetInfo) GetPeersOk() ([]Peer, bool) {
	if o == nil || o.Peers == nil {
		return nil, false
	}
	return o.Peers, true
}

// HasPeers returns a boolean if a field has been set.
func (o *NetInfo) HasPeers() bool {
	if o != nil && o.Peers != nil {
		return true
	}

	return false
}

// SetPeers gets a reference to the given []Peer and assigns it to the Peers field.
func (o *NetInfo) SetPeers(v []Peer) {
	o.Peers = v
}

func (o NetInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Listening != nil {
		toSerialize["listening"] = o.Listening
	}
	if o.Listeners != nil {
		toSerialize["listeners"] = o.Listeners
	}
	if o.NPeers != nil {
		toSerialize["n_peers"] = o.NPeers
	}
	if o.Peers != nil {
		toSerialize["peers"] = o.Peers
	}
	return json.Marshal(toSerialize)
}

type NullableNetInfo struct {
	value *NetInfo
	isSet bool
}

func (v NullableNetInfo) Get() *NetInfo {
	return v.value
}

func (v *NullableNetInfo) Set(val *NetInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableNetInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableNetInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetInfo(val *NetInfo) *NullableNetInfo {
	return &NullableNetInfo{value: val, isSet: true}
}

func (v NullableNetInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


