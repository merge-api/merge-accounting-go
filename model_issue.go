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
	"time"
)

// Issue struct for Issue
type Issue struct {
	Id *string `json:"id,omitempty"`
	// Status of the issue. Options: ('ONGOING', 'RESOLVED')  * `ONGOING` - ONGOING * `RESOLVED` - RESOLVED
	Status *IssueStatusEnum `json:"status,omitempty"`
	ErrorDescription string `json:"error_description"`
	EndUser *map[string]interface{} `json:"end_user,omitempty"`
	FirstIncidentTime NullableTime `json:"first_incident_time,omitempty"`
	LastIncidentTime NullableTime `json:"last_incident_time,omitempty"`
	IsMuted *bool `json:"is_muted,omitempty"`
	ErrorDetails *[]string `json:"error_details,omitempty"`
	// raw json response by property name
	ResponseRaw map[string]json.RawMessage `json:"-"`
}

// NewIssue instantiates a new Issue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIssue(errorDescription string) *Issue {
	this := Issue{}
	this.ErrorDescription = errorDescription
	return &this
}

// NewIssueWithDefaults instantiates a new Issue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIssueWithDefaults() *Issue {
	this := Issue{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Issue) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Issue) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Issue) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Issue) SetId(v string) {
	o.Id = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Issue) GetStatus() IssueStatusEnum {
	if o == nil || o.Status == nil {
		var ret IssueStatusEnum
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Issue) GetStatusOk() (*IssueStatusEnum, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Issue) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given IssueStatusEnum and assigns it to the Status field.
func (o *Issue) SetStatus(v IssueStatusEnum) {
	o.Status = &v
}

// GetErrorDescription returns the ErrorDescription field value
func (o *Issue) GetErrorDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ErrorDescription
}

// GetErrorDescriptionOk returns a tuple with the ErrorDescription field value
// and a boolean to check if the value has been set.
func (o *Issue) GetErrorDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ErrorDescription, true
}

// SetErrorDescription sets field value
func (o *Issue) SetErrorDescription(v string) {
	o.ErrorDescription = v
}

// GetEndUser returns the EndUser field value if set, zero value otherwise.
func (o *Issue) GetEndUser() map[string]interface{} {
	if o == nil || o.EndUser == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.EndUser
}

// GetEndUserOk returns a tuple with the EndUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Issue) GetEndUserOk() (*map[string]interface{}, bool) {
	if o == nil || o.EndUser == nil {
		return nil, false
	}
	return o.EndUser, true
}

// HasEndUser returns a boolean if a field has been set.
func (o *Issue) HasEndUser() bool {
	if o != nil && o.EndUser != nil {
		return true
	}

	return false
}

// SetEndUser gets a reference to the given map[string]interface{} and assigns it to the EndUser field.
func (o *Issue) SetEndUser(v map[string]interface{}) {
	o.EndUser = &v
}

// GetFirstIncidentTime returns the FirstIncidentTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Issue) GetFirstIncidentTime() time.Time {
	if o == nil || o.FirstIncidentTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.FirstIncidentTime.Get()
}

// GetFirstIncidentTimeOk returns a tuple with the FirstIncidentTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Issue) GetFirstIncidentTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.FirstIncidentTime.Get(), o.FirstIncidentTime.IsSet()
}

// HasFirstIncidentTime returns a boolean if a field has been set.
func (o *Issue) HasFirstIncidentTime() bool {
	if o != nil && o.FirstIncidentTime.IsSet() {
		return true
	}

	return false
}

// SetFirstIncidentTime gets a reference to the given NullableTime and assigns it to the FirstIncidentTime field.
func (o *Issue) SetFirstIncidentTime(v time.Time) {
	o.FirstIncidentTime.Set(&v)
}
// SetFirstIncidentTimeNil sets the value for FirstIncidentTime to be an explicit nil
func (o *Issue) SetFirstIncidentTimeNil() {
	o.FirstIncidentTime.Set(nil)
}

// UnsetFirstIncidentTime ensures that no value is present for FirstIncidentTime, not even an explicit nil
func (o *Issue) UnsetFirstIncidentTime() {
	o.FirstIncidentTime.Unset()
}

// GetLastIncidentTime returns the LastIncidentTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Issue) GetLastIncidentTime() time.Time {
	if o == nil || o.LastIncidentTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.LastIncidentTime.Get()
}

// GetLastIncidentTimeOk returns a tuple with the LastIncidentTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Issue) GetLastIncidentTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LastIncidentTime.Get(), o.LastIncidentTime.IsSet()
}

// HasLastIncidentTime returns a boolean if a field has been set.
func (o *Issue) HasLastIncidentTime() bool {
	if o != nil && o.LastIncidentTime.IsSet() {
		return true
	}

	return false
}

// SetLastIncidentTime gets a reference to the given NullableTime and assigns it to the LastIncidentTime field.
func (o *Issue) SetLastIncidentTime(v time.Time) {
	o.LastIncidentTime.Set(&v)
}
// SetLastIncidentTimeNil sets the value for LastIncidentTime to be an explicit nil
func (o *Issue) SetLastIncidentTimeNil() {
	o.LastIncidentTime.Set(nil)
}

// UnsetLastIncidentTime ensures that no value is present for LastIncidentTime, not even an explicit nil
func (o *Issue) UnsetLastIncidentTime() {
	o.LastIncidentTime.Unset()
}

// GetIsMuted returns the IsMuted field value if set, zero value otherwise.
func (o *Issue) GetIsMuted() bool {
	if o == nil || o.IsMuted == nil {
		var ret bool
		return ret
	}
	return *o.IsMuted
}

// GetIsMutedOk returns a tuple with the IsMuted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Issue) GetIsMutedOk() (*bool, bool) {
	if o == nil || o.IsMuted == nil {
		return nil, false
	}
	return o.IsMuted, true
}

// HasIsMuted returns a boolean if a field has been set.
func (o *Issue) HasIsMuted() bool {
	if o != nil && o.IsMuted != nil {
		return true
	}

	return false
}

// SetIsMuted gets a reference to the given bool and assigns it to the IsMuted field.
func (o *Issue) SetIsMuted(v bool) {
	o.IsMuted = &v
}

// GetErrorDetails returns the ErrorDetails field value if set, zero value otherwise.
func (o *Issue) GetErrorDetails() []string {
	if o == nil || o.ErrorDetails == nil {
		var ret []string
		return ret
	}
	return *o.ErrorDetails
}

// GetErrorDetailsOk returns a tuple with the ErrorDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Issue) GetErrorDetailsOk() (*[]string, bool) {
	if o == nil || o.ErrorDetails == nil {
		return nil, false
	}
	return o.ErrorDetails, true
}

// HasErrorDetails returns a boolean if a field has been set.
func (o *Issue) HasErrorDetails() bool {
	if o != nil && o.ErrorDetails != nil {
		return true
	}

	return false
}

// SetErrorDetails gets a reference to the given []string and assigns it to the ErrorDetails field.
func (o *Issue) SetErrorDetails(v []string) {
	o.ErrorDetails = &v
}

func (o Issue) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["error_description"] = o.ErrorDescription
	}
	if o.EndUser != nil {
		toSerialize["end_user"] = o.EndUser
	}
	if o.FirstIncidentTime.IsSet() {
		toSerialize["first_incident_time"] = o.FirstIncidentTime.Get()
	}
	if o.LastIncidentTime.IsSet() {
		toSerialize["last_incident_time"] = o.LastIncidentTime.Get()
	}
	if o.IsMuted != nil {
		toSerialize["is_muted"] = o.IsMuted
	}
	if o.ErrorDetails != nil {
		toSerialize["error_details"] = o.ErrorDetails
	}
	return json.Marshal(toSerialize)
}

func (v *Issue) UnmarshalJSON(src []byte) error {
    type IssueUnmarshalTarget Issue

	var intermediateResult IssueUnmarshalTarget
	var err1 = json.Unmarshal(src, &intermediateResult)
    if err1 != nil {
        return err1
    }
    var err2 = json.Unmarshal(src, &intermediateResult.ResponseRaw)
	if err2 != nil {
		return err2
	}

	*v = Issue(intermediateResult)
	return nil
}
type NullableIssue struct {
	value *Issue
	isSet bool
}

func (v NullableIssue) Get() *Issue {
	return v.value
}

func (v *NullableIssue) Set(val *Issue) {
	v.value = val
	v.isSet = true
}

func (v NullableIssue) IsSet() bool {
	return v.isSet
}

func (v *NullableIssue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIssue(val *Issue) *NullableIssue {
	return &NullableIssue{value: val, isSet: true}
}

func (v NullableIssue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIssue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	var err1 = json.Unmarshal(src, &v.value)
    if err1 != nil {
        return err1
    }
    return json.Unmarshal(src, &v.value.ResponseRaw)
}


