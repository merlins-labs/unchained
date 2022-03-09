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

// BlockSearchResponseResult struct for BlockSearchResponseResult
type BlockSearchResponseResult struct {
	Blocks []BlockComplete `json:"blocks"`
	TotalCount int32 `json:"total_count"`
}

// NewBlockSearchResponseResult instantiates a new BlockSearchResponseResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlockSearchResponseResult(blocks []BlockComplete, totalCount int32) *BlockSearchResponseResult {
	this := BlockSearchResponseResult{}
	this.Blocks = blocks
	this.TotalCount = totalCount
	return &this
}

// NewBlockSearchResponseResultWithDefaults instantiates a new BlockSearchResponseResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlockSearchResponseResultWithDefaults() *BlockSearchResponseResult {
	this := BlockSearchResponseResult{}
	return &this
}

// GetBlocks returns the Blocks field value
func (o *BlockSearchResponseResult) GetBlocks() []BlockComplete {
	if o == nil {
		var ret []BlockComplete
		return ret
	}

	return o.Blocks
}

// GetBlocksOk returns a tuple with the Blocks field value
// and a boolean to check if the value has been set.
func (o *BlockSearchResponseResult) GetBlocksOk() ([]BlockComplete, bool) {
	if o == nil {
		return nil, false
	}
	return o.Blocks, true
}

// SetBlocks sets field value
func (o *BlockSearchResponseResult) SetBlocks(v []BlockComplete) {
	o.Blocks = v
}

// GetTotalCount returns the TotalCount field value
func (o *BlockSearchResponseResult) GetTotalCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value
// and a boolean to check if the value has been set.
func (o *BlockSearchResponseResult) GetTotalCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalCount, true
}

// SetTotalCount sets field value
func (o *BlockSearchResponseResult) SetTotalCount(v int32) {
	o.TotalCount = v
}

func (o BlockSearchResponseResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["blocks"] = o.Blocks
	}
	if true {
		toSerialize["total_count"] = o.TotalCount
	}
	return json.Marshal(toSerialize)
}

type NullableBlockSearchResponseResult struct {
	value *BlockSearchResponseResult
	isSet bool
}

func (v NullableBlockSearchResponseResult) Get() *BlockSearchResponseResult {
	return v.value
}

func (v *NullableBlockSearchResponseResult) Set(val *BlockSearchResponseResult) {
	v.value = val
	v.isSet = true
}

func (v NullableBlockSearchResponseResult) IsSet() bool {
	return v.isSet
}

func (v *NullableBlockSearchResponseResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlockSearchResponseResult(val *BlockSearchResponseResult) *NullableBlockSearchResponseResult {
	return &NullableBlockSearchResponseResult{value: val, isSet: true}
}

func (v NullableBlockSearchResponseResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlockSearchResponseResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


