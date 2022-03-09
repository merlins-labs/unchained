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

// BlockResultsResponseResultEvents struct for BlockResultsResponseResultEvents
type BlockResultsResponseResultEvents struct {
	Type *string `json:"type,omitempty"`
	Attributes []Event `json:"attributes,omitempty"`
}

// NewBlockResultsResponseResultEvents instantiates a new BlockResultsResponseResultEvents object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlockResultsResponseResultEvents() *BlockResultsResponseResultEvents {
	this := BlockResultsResponseResultEvents{}
	return &this
}

// NewBlockResultsResponseResultEventsWithDefaults instantiates a new BlockResultsResponseResultEvents object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlockResultsResponseResultEventsWithDefaults() *BlockResultsResponseResultEvents {
	this := BlockResultsResponseResultEvents{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *BlockResultsResponseResultEvents) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockResultsResponseResultEvents) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *BlockResultsResponseResultEvents) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *BlockResultsResponseResultEvents) SetType(v string) {
	o.Type = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *BlockResultsResponseResultEvents) GetAttributes() []Event {
	if o == nil || o.Attributes == nil {
		var ret []Event
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockResultsResponseResultEvents) GetAttributesOk() ([]Event, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *BlockResultsResponseResultEvents) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given []Event and assigns it to the Attributes field.
func (o *BlockResultsResponseResultEvents) SetAttributes(v []Event) {
	o.Attributes = v
}

func (o BlockResultsResponseResultEvents) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	return json.Marshal(toSerialize)
}

type NullableBlockResultsResponseResultEvents struct {
	value *BlockResultsResponseResultEvents
	isSet bool
}

func (v NullableBlockResultsResponseResultEvents) Get() *BlockResultsResponseResultEvents {
	return v.value
}

func (v *NullableBlockResultsResponseResultEvents) Set(val *BlockResultsResponseResultEvents) {
	v.value = val
	v.isSet = true
}

func (v NullableBlockResultsResponseResultEvents) IsSet() bool {
	return v.isSet
}

func (v *NullableBlockResultsResponseResultEvents) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlockResultsResponseResultEvents(val *BlockResultsResponseResultEvents) *NullableBlockResultsResponseResultEvents {
	return &NullableBlockResultsResponseResultEvents{value: val, isSet: true}
}

func (v NullableBlockResultsResponseResultEvents) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlockResultsResponseResultEvents) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


