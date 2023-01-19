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

// LinkedAccountSelectiveSyncConfigurationRequest struct for LinkedAccountSelectiveSyncConfigurationRequest
type LinkedAccountSelectiveSyncConfigurationRequest struct {
	// The conditions belonging to a selective sync.
	LinkedAccountConditions []LinkedAccountConditionRequest `json:"linked_account_conditions"`
	// raw json response by property name
	ResponseRaw map[string]json.RawMessage `json:"-"`
}

// NewLinkedAccountSelectiveSyncConfigurationRequest instantiates a new LinkedAccountSelectiveSyncConfigurationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkedAccountSelectiveSyncConfigurationRequest(linkedAccountConditions []LinkedAccountConditionRequest) *LinkedAccountSelectiveSyncConfigurationRequest {
	this := LinkedAccountSelectiveSyncConfigurationRequest{}
	this.LinkedAccountConditions = linkedAccountConditions
	return &this
}

// NewLinkedAccountSelectiveSyncConfigurationRequestWithDefaults instantiates a new LinkedAccountSelectiveSyncConfigurationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkedAccountSelectiveSyncConfigurationRequestWithDefaults() *LinkedAccountSelectiveSyncConfigurationRequest {
	this := LinkedAccountSelectiveSyncConfigurationRequest{}
	return &this
}

// GetLinkedAccountConditions returns the LinkedAccountConditions field value
func (o *LinkedAccountSelectiveSyncConfigurationRequest) GetLinkedAccountConditions() []LinkedAccountConditionRequest {
	if o == nil {
		var ret []LinkedAccountConditionRequest
		return ret
	}

	return o.LinkedAccountConditions
}

// GetLinkedAccountConditionsOk returns a tuple with the LinkedAccountConditions field value
// and a boolean to check if the value has been set.
func (o *LinkedAccountSelectiveSyncConfigurationRequest) GetLinkedAccountConditionsOk() (*[]LinkedAccountConditionRequest, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LinkedAccountConditions, true
}

// SetLinkedAccountConditions sets field value
func (o *LinkedAccountSelectiveSyncConfigurationRequest) SetLinkedAccountConditions(v []LinkedAccountConditionRequest) {
	o.LinkedAccountConditions = v
}

func (o LinkedAccountSelectiveSyncConfigurationRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["linked_account_conditions"] = o.LinkedAccountConditions
	}
	return json.Marshal(toSerialize)
}

func (v *LinkedAccountSelectiveSyncConfigurationRequest) UnmarshalJSON(src []byte) error {
    type LinkedAccountSelectiveSyncConfigurationRequestUnmarshalTarget LinkedAccountSelectiveSyncConfigurationRequest

	var intermediateResult LinkedAccountSelectiveSyncConfigurationRequestUnmarshalTarget
	var err1 = json.Unmarshal(src, &intermediateResult)
    if err1 != nil {
        return err1
    }
    var err2 = json.Unmarshal(src, &intermediateResult.ResponseRaw)
	if err2 != nil {
		return err2
	}

	*v = LinkedAccountSelectiveSyncConfigurationRequest(intermediateResult)
	return nil
}
type NullableLinkedAccountSelectiveSyncConfigurationRequest struct {
	value *LinkedAccountSelectiveSyncConfigurationRequest
	isSet bool
}

func (v NullableLinkedAccountSelectiveSyncConfigurationRequest) Get() *LinkedAccountSelectiveSyncConfigurationRequest {
	return v.value
}

func (v *NullableLinkedAccountSelectiveSyncConfigurationRequest) Set(val *LinkedAccountSelectiveSyncConfigurationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkedAccountSelectiveSyncConfigurationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkedAccountSelectiveSyncConfigurationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkedAccountSelectiveSyncConfigurationRequest(val *LinkedAccountSelectiveSyncConfigurationRequest) *NullableLinkedAccountSelectiveSyncConfigurationRequest {
	return &NullableLinkedAccountSelectiveSyncConfigurationRequest{value: val, isSet: true}
}

func (v NullableLinkedAccountSelectiveSyncConfigurationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkedAccountSelectiveSyncConfigurationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	var err1 = json.Unmarshal(src, &v.value)
    if err1 != nil {
        return err1
    }
    return json.Unmarshal(src, &v.value.ResponseRaw)
}


