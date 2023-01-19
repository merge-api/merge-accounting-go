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

// ModelOperation # The ModelOperation Object ### Description The `ModelOperation` object is used to represent the operations that are currently supported for a given model.  ### Usage Example View what operations are supported for the `Candidate` endpoint.
type ModelOperation struct {
	ModelName string `json:"model_name"`
	AvailableOperations []string `json:"available_operations"`
	RequiredPostParameters []string `json:"required_post_parameters"`
	SupportedFields []string `json:"supported_fields"`
	// raw json response by property name
	ResponseRaw map[string]json.RawMessage `json:"-"`
}

// NewModelOperation instantiates a new ModelOperation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelOperation(modelName string, availableOperations []string, requiredPostParameters []string, supportedFields []string) *ModelOperation {
	this := ModelOperation{}
	this.ModelName = modelName
	this.AvailableOperations = availableOperations
	this.RequiredPostParameters = requiredPostParameters
	this.SupportedFields = supportedFields
	return &this
}

// NewModelOperationWithDefaults instantiates a new ModelOperation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelOperationWithDefaults() *ModelOperation {
	this := ModelOperation{}
	return &this
}

// GetModelName returns the ModelName field value
func (o *ModelOperation) GetModelName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ModelName
}

// GetModelNameOk returns a tuple with the ModelName field value
// and a boolean to check if the value has been set.
func (o *ModelOperation) GetModelNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ModelName, true
}

// SetModelName sets field value
func (o *ModelOperation) SetModelName(v string) {
	o.ModelName = v
}

// GetAvailableOperations returns the AvailableOperations field value
func (o *ModelOperation) GetAvailableOperations() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AvailableOperations
}

// GetAvailableOperationsOk returns a tuple with the AvailableOperations field value
// and a boolean to check if the value has been set.
func (o *ModelOperation) GetAvailableOperationsOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AvailableOperations, true
}

// SetAvailableOperations sets field value
func (o *ModelOperation) SetAvailableOperations(v []string) {
	o.AvailableOperations = v
}

// GetRequiredPostParameters returns the RequiredPostParameters field value
func (o *ModelOperation) GetRequiredPostParameters() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.RequiredPostParameters
}

// GetRequiredPostParametersOk returns a tuple with the RequiredPostParameters field value
// and a boolean to check if the value has been set.
func (o *ModelOperation) GetRequiredPostParametersOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequiredPostParameters, true
}

// SetRequiredPostParameters sets field value
func (o *ModelOperation) SetRequiredPostParameters(v []string) {
	o.RequiredPostParameters = v
}

// GetSupportedFields returns the SupportedFields field value
func (o *ModelOperation) GetSupportedFields() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.SupportedFields
}

// GetSupportedFieldsOk returns a tuple with the SupportedFields field value
// and a boolean to check if the value has been set.
func (o *ModelOperation) GetSupportedFieldsOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SupportedFields, true
}

// SetSupportedFields sets field value
func (o *ModelOperation) SetSupportedFields(v []string) {
	o.SupportedFields = v
}

func (o ModelOperation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["model_name"] = o.ModelName
	}
	if true {
		toSerialize["available_operations"] = o.AvailableOperations
	}
	if true {
		toSerialize["required_post_parameters"] = o.RequiredPostParameters
	}
	if true {
		toSerialize["supported_fields"] = o.SupportedFields
	}
	return json.Marshal(toSerialize)
}

func (v *ModelOperation) UnmarshalJSON(src []byte) error {
    type ModelOperationUnmarshalTarget ModelOperation

	var intermediateResult ModelOperationUnmarshalTarget
	var err1 = json.Unmarshal(src, &intermediateResult)
    if err1 != nil {
        return err1
    }
    var err2 = json.Unmarshal(src, &intermediateResult.ResponseRaw)
	if err2 != nil {
		return err2
	}

	*v = ModelOperation(intermediateResult)
	return nil
}
type NullableModelOperation struct {
	value *ModelOperation
	isSet bool
}

func (v NullableModelOperation) Get() *ModelOperation {
	return v.value
}

func (v *NullableModelOperation) Set(val *ModelOperation) {
	v.value = val
	v.isSet = true
}

func (v NullableModelOperation) IsSet() bool {
	return v.isSet
}

func (v *NullableModelOperation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelOperation(val *ModelOperation) *NullableModelOperation {
	return &NullableModelOperation{value: val, isSet: true}
}

func (v NullableModelOperation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelOperation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	var err1 = json.Unmarshal(src, &v.value)
    if err1 != nil {
        return err1
    }
    return json.Unmarshal(src, &v.value.ResponseRaw)
}


