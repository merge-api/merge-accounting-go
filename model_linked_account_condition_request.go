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

// LinkedAccountConditionRequest struct for LinkedAccountConditionRequest
type LinkedAccountConditionRequest struct {
	// The ID indicating which condition schema to use for a specific condition.
	ConditionSchemaId string `json:"condition_schema_id"`
	// The operator for a specific condition.
	Operator string `json:"operator"`
	// The value for a specific condition.
	Value interface{} `json:"value"`
	// raw json response by property name
	ResponseRaw map[string]json.RawMessage `json:"-"`
}

// NewLinkedAccountConditionRequest instantiates a new LinkedAccountConditionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkedAccountConditionRequest(conditionSchemaId string, operator string, value interface{}) *LinkedAccountConditionRequest {
	this := LinkedAccountConditionRequest{}
	this.ConditionSchemaId = conditionSchemaId
	this.Operator = operator
	this.Value = value
	return &this
}

// NewLinkedAccountConditionRequestWithDefaults instantiates a new LinkedAccountConditionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkedAccountConditionRequestWithDefaults() *LinkedAccountConditionRequest {
	this := LinkedAccountConditionRequest{}
	return &this
}

// GetConditionSchemaId returns the ConditionSchemaId field value
func (o *LinkedAccountConditionRequest) GetConditionSchemaId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConditionSchemaId
}

// GetConditionSchemaIdOk returns a tuple with the ConditionSchemaId field value
// and a boolean to check if the value has been set.
func (o *LinkedAccountConditionRequest) GetConditionSchemaIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ConditionSchemaId, true
}

// SetConditionSchemaId sets field value
func (o *LinkedAccountConditionRequest) SetConditionSchemaId(v string) {
	o.ConditionSchemaId = v
}

// GetOperator returns the Operator field value
func (o *LinkedAccountConditionRequest) GetOperator() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value
// and a boolean to check if the value has been set.
func (o *LinkedAccountConditionRequest) GetOperatorOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Operator, true
}

// SetOperator sets field value
func (o *LinkedAccountConditionRequest) SetOperator(v string) {
	o.Operator = v
}

// GetValue returns the Value field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *LinkedAccountConditionRequest) GetValue() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LinkedAccountConditionRequest) GetValueOk() (*interface{}, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *LinkedAccountConditionRequest) SetValue(v interface{}) {
	o.Value = v
}

func (o LinkedAccountConditionRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["condition_schema_id"] = o.ConditionSchemaId
	}
	if true {
		toSerialize["operator"] = o.Operator
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

func (v *LinkedAccountConditionRequest) UnmarshalJSON(src []byte) error {
    type LinkedAccountConditionRequestUnmarshalTarget LinkedAccountConditionRequest

	var intermediateResult LinkedAccountConditionRequestUnmarshalTarget
	var err1 = json.Unmarshal(src, &intermediateResult)
    if err1 != nil {
        return err1
    }
    var err2 = json.Unmarshal(src, &intermediateResult.ResponseRaw)
	if err2 != nil {
		return err2
	}

	*v = LinkedAccountConditionRequest(intermediateResult)
	return nil
}
type NullableLinkedAccountConditionRequest struct {
	value *LinkedAccountConditionRequest
	isSet bool
}

func (v NullableLinkedAccountConditionRequest) Get() *LinkedAccountConditionRequest {
	return v.value
}

func (v *NullableLinkedAccountConditionRequest) Set(val *LinkedAccountConditionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkedAccountConditionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkedAccountConditionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkedAccountConditionRequest(val *LinkedAccountConditionRequest) *NullableLinkedAccountConditionRequest {
	return &NullableLinkedAccountConditionRequest{value: val, isSet: true}
}

func (v NullableLinkedAccountConditionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkedAccountConditionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	var err1 = json.Unmarshal(src, &v.value)
    if err1 != nil {
        return err1
    }
    return json.Unmarshal(src, &v.value.ResponseRaw)
}

