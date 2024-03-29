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

// MetaResponse struct for MetaResponse
type MetaResponse struct {
	RequestSchema map[string]interface{} `json:"request_schema"`
	RemoteFieldClasses *map[string]interface{} `json:"remote_field_classes,omitempty"`
	Status *LinkedAccountStatus `json:"status,omitempty"`
	HasConditionalParams bool `json:"has_conditional_params"`
	HasRequiredLinkedAccountParams bool `json:"has_required_linked_account_params"`
	// raw json response by property name
	ResponseRaw map[string]json.RawMessage `json:"-"`
}

// NewMetaResponse instantiates a new MetaResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetaResponse(requestSchema map[string]interface{}, hasConditionalParams bool, hasRequiredLinkedAccountParams bool) *MetaResponse {
	this := MetaResponse{}
	this.RequestSchema = requestSchema
	this.HasConditionalParams = hasConditionalParams
	this.HasRequiredLinkedAccountParams = hasRequiredLinkedAccountParams
	return &this
}

// NewMetaResponseWithDefaults instantiates a new MetaResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetaResponseWithDefaults() *MetaResponse {
	this := MetaResponse{}
	return &this
}

// GetRequestSchema returns the RequestSchema field value
func (o *MetaResponse) GetRequestSchema() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.RequestSchema
}

// GetRequestSchemaOk returns a tuple with the RequestSchema field value
// and a boolean to check if the value has been set.
func (o *MetaResponse) GetRequestSchemaOk() (*map[string]interface{}, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestSchema, true
}

// SetRequestSchema sets field value
func (o *MetaResponse) SetRequestSchema(v map[string]interface{}) {
	o.RequestSchema = v
}

// GetRemoteFieldClasses returns the RemoteFieldClasses field value if set, zero value otherwise.
func (o *MetaResponse) GetRemoteFieldClasses() map[string]interface{} {
	if o == nil || o.RemoteFieldClasses == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.RemoteFieldClasses
}

// GetRemoteFieldClassesOk returns a tuple with the RemoteFieldClasses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaResponse) GetRemoteFieldClassesOk() (*map[string]interface{}, bool) {
	if o == nil || o.RemoteFieldClasses == nil {
		return nil, false
	}
	return o.RemoteFieldClasses, true
}

// HasRemoteFieldClasses returns a boolean if a field has been set.
func (o *MetaResponse) HasRemoteFieldClasses() bool {
	if o != nil && o.RemoteFieldClasses != nil {
		return true
	}

	return false
}

// SetRemoteFieldClasses gets a reference to the given map[string]interface{} and assigns it to the RemoteFieldClasses field.
func (o *MetaResponse) SetRemoteFieldClasses(v map[string]interface{}) {
	o.RemoteFieldClasses = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *MetaResponse) GetStatus() LinkedAccountStatus {
	if o == nil || o.Status == nil {
		var ret LinkedAccountStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaResponse) GetStatusOk() (*LinkedAccountStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *MetaResponse) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given LinkedAccountStatus and assigns it to the Status field.
func (o *MetaResponse) SetStatus(v LinkedAccountStatus) {
	o.Status = &v
}

// GetHasConditionalParams returns the HasConditionalParams field value
func (o *MetaResponse) GetHasConditionalParams() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasConditionalParams
}

// GetHasConditionalParamsOk returns a tuple with the HasConditionalParams field value
// and a boolean to check if the value has been set.
func (o *MetaResponse) GetHasConditionalParamsOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.HasConditionalParams, true
}

// SetHasConditionalParams sets field value
func (o *MetaResponse) SetHasConditionalParams(v bool) {
	o.HasConditionalParams = v
}

// GetHasRequiredLinkedAccountParams returns the HasRequiredLinkedAccountParams field value
func (o *MetaResponse) GetHasRequiredLinkedAccountParams() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasRequiredLinkedAccountParams
}

// GetHasRequiredLinkedAccountParamsOk returns a tuple with the HasRequiredLinkedAccountParams field value
// and a boolean to check if the value has been set.
func (o *MetaResponse) GetHasRequiredLinkedAccountParamsOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.HasRequiredLinkedAccountParams, true
}

// SetHasRequiredLinkedAccountParams sets field value
func (o *MetaResponse) SetHasRequiredLinkedAccountParams(v bool) {
	o.HasRequiredLinkedAccountParams = v
}

func (o MetaResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["request_schema"] = o.RequestSchema
	}
	if o.RemoteFieldClasses != nil {
		toSerialize["remote_field_classes"] = o.RemoteFieldClasses
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["has_conditional_params"] = o.HasConditionalParams
	}
	if true {
		toSerialize["has_required_linked_account_params"] = o.HasRequiredLinkedAccountParams
	}
	return json.Marshal(toSerialize)
}

func (v *MetaResponse) UnmarshalJSON(src []byte) error {
    type MetaResponseUnmarshalTarget MetaResponse

	var intermediateResult MetaResponseUnmarshalTarget
	var err1 = json.Unmarshal(src, &intermediateResult)
    if err1 != nil {
        return err1
    }
    var err2 = json.Unmarshal(src, &intermediateResult.ResponseRaw)
	if err2 != nil {
		return err2
	}

	*v = MetaResponse(intermediateResult)
	return nil
}
type NullableMetaResponse struct {
	value *MetaResponse
	isSet bool
}

func (v NullableMetaResponse) Get() *MetaResponse {
	return v.value
}

func (v *NullableMetaResponse) Set(val *MetaResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMetaResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMetaResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetaResponse(val *MetaResponse) *NullableMetaResponse {
	return &NullableMetaResponse{value: val, isSet: true}
}

func (v NullableMetaResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetaResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	var err1 = json.Unmarshal(src, &v.value)
    if err1 != nil {
        return err1
    }
    return json.Unmarshal(src, &v.value.ResponseRaw)
}


