/*
 * Merge Accounting API
 *
 * The unified API for building rich integrations with multiple Accounting & Finance platforms.
 *
 * API version: 1.0
 * Contact: hello@merge.dev
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package merge_accounting_client

import (
	"encoding/json"
)

// JournalEntryEndpointRequest struct for JournalEntryEndpointRequest
type JournalEntryEndpointRequest struct {
	Model JournalEntryRequest `json:"model"`
	// raw json response by property name
	ResponseRaw map[string]json.RawMessage `json:"-"`
}

// NewJournalEntryEndpointRequest instantiates a new JournalEntryEndpointRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJournalEntryEndpointRequest(model JournalEntryRequest) *JournalEntryEndpointRequest {
	this := JournalEntryEndpointRequest{}
	this.Model = model
	return &this
}

// NewJournalEntryEndpointRequestWithDefaults instantiates a new JournalEntryEndpointRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJournalEntryEndpointRequestWithDefaults() *JournalEntryEndpointRequest {
	this := JournalEntryEndpointRequest{}
	return &this
}

// GetModel returns the Model field value
func (o *JournalEntryEndpointRequest) GetModel() JournalEntryRequest {
	if o == nil {
		var ret JournalEntryRequest
		return ret
	}

	return o.Model
}

// GetModelOk returns a tuple with the Model field value
// and a boolean to check if the value has been set.
func (o *JournalEntryEndpointRequest) GetModelOk() (*JournalEntryRequest, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Model, true
}

// SetModel sets field value
func (o *JournalEntryEndpointRequest) SetModel(v JournalEntryRequest) {
	o.Model = v
}

func (o JournalEntryEndpointRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["model"] = o.Model
	}
	return json.Marshal(toSerialize)
}

func (v *JournalEntryEndpointRequest) UnmarshalJSON(src []byte) error {
    type JournalEntryEndpointRequestUnmarshalTarget JournalEntryEndpointRequest

	var intermediateResult JournalEntryEndpointRequestUnmarshalTarget
	var err1 = json.Unmarshal(src, &intermediateResult)
    if err1 != nil {
        return err1
    }
    var err2 = json.Unmarshal(src, &intermediateResult.ResponseRaw)
	if err2 != nil {
		return err2
	}

	*v = JournalEntryEndpointRequest(intermediateResult)
	return nil
}
type NullableJournalEntryEndpointRequest struct {
	value *JournalEntryEndpointRequest
	isSet bool
}

func (v NullableJournalEntryEndpointRequest) Get() *JournalEntryEndpointRequest {
	return v.value
}

func (v *NullableJournalEntryEndpointRequest) Set(val *JournalEntryEndpointRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableJournalEntryEndpointRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableJournalEntryEndpointRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJournalEntryEndpointRequest(val *JournalEntryEndpointRequest) *NullableJournalEntryEndpointRequest {
	return &NullableJournalEntryEndpointRequest{value: val, isSet: true}
}

func (v NullableJournalEntryEndpointRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJournalEntryEndpointRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	var err1 = json.Unmarshal(src, &v.value)
    if err1 != nil {
        return err1
    }
    return json.Unmarshal(src, &v.value.ResponseRaw)
}


