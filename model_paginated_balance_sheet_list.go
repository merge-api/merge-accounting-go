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

// PaginatedBalanceSheetList struct for PaginatedBalanceSheetList
type PaginatedBalanceSheetList struct {
	Next NullableString `json:"next,omitempty"`
	Previous NullableString `json:"previous,omitempty"`
	Results *[]BalanceSheet `json:"results,omitempty"`
	// raw json response by property name
	ResponseRaw map[string]json.RawMessage `json:"-"`
}

// NewPaginatedBalanceSheetList instantiates a new PaginatedBalanceSheetList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedBalanceSheetList() *PaginatedBalanceSheetList {
	this := PaginatedBalanceSheetList{}
	return &this
}

// NewPaginatedBalanceSheetListWithDefaults instantiates a new PaginatedBalanceSheetList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedBalanceSheetListWithDefaults() *PaginatedBalanceSheetList {
	this := PaginatedBalanceSheetList{}
	return &this
}

// GetNext returns the Next field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaginatedBalanceSheetList) GetNext() string {
	if o == nil || o.Next.Get() == nil {
		var ret string
		return ret
	}
	return *o.Next.Get()
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaginatedBalanceSheetList) GetNextOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Next.Get(), o.Next.IsSet()
}

// HasNext returns a boolean if a field has been set.
func (o *PaginatedBalanceSheetList) HasNext() bool {
	if o != nil && o.Next.IsSet() {
		return true
	}

	return false
}

// SetNext gets a reference to the given NullableString and assigns it to the Next field.
func (o *PaginatedBalanceSheetList) SetNext(v string) {
	o.Next.Set(&v)
}
// SetNextNil sets the value for Next to be an explicit nil
func (o *PaginatedBalanceSheetList) SetNextNil() {
	o.Next.Set(nil)
}

// UnsetNext ensures that no value is present for Next, not even an explicit nil
func (o *PaginatedBalanceSheetList) UnsetNext() {
	o.Next.Unset()
}

// GetPrevious returns the Previous field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaginatedBalanceSheetList) GetPrevious() string {
	if o == nil || o.Previous.Get() == nil {
		var ret string
		return ret
	}
	return *o.Previous.Get()
}

// GetPreviousOk returns a tuple with the Previous field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaginatedBalanceSheetList) GetPreviousOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Previous.Get(), o.Previous.IsSet()
}

// HasPrevious returns a boolean if a field has been set.
func (o *PaginatedBalanceSheetList) HasPrevious() bool {
	if o != nil && o.Previous.IsSet() {
		return true
	}

	return false
}

// SetPrevious gets a reference to the given NullableString and assigns it to the Previous field.
func (o *PaginatedBalanceSheetList) SetPrevious(v string) {
	o.Previous.Set(&v)
}
// SetPreviousNil sets the value for Previous to be an explicit nil
func (o *PaginatedBalanceSheetList) SetPreviousNil() {
	o.Previous.Set(nil)
}

// UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
func (o *PaginatedBalanceSheetList) UnsetPrevious() {
	o.Previous.Unset()
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *PaginatedBalanceSheetList) GetResults() []BalanceSheet {
	if o == nil || o.Results == nil {
		var ret []BalanceSheet
		return ret
	}
	return *o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedBalanceSheetList) GetResultsOk() (*[]BalanceSheet, bool) {
	if o == nil || o.Results == nil {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *PaginatedBalanceSheetList) HasResults() bool {
	if o != nil && o.Results != nil {
		return true
	}

	return false
}

// SetResults gets a reference to the given []BalanceSheet and assigns it to the Results field.
func (o *PaginatedBalanceSheetList) SetResults(v []BalanceSheet) {
	o.Results = &v
}

func (o PaginatedBalanceSheetList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Next.IsSet() {
		toSerialize["next"] = o.Next.Get()
	}
	if o.Previous.IsSet() {
		toSerialize["previous"] = o.Previous.Get()
	}
	if o.Results != nil {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

func (v *PaginatedBalanceSheetList) UnmarshalJSON(src []byte) error {
    type PaginatedBalanceSheetListUnmarshalTarget PaginatedBalanceSheetList

	var intermediateResult PaginatedBalanceSheetListUnmarshalTarget
	var err1 = json.Unmarshal(src, &intermediateResult)
    if err1 != nil {
        return err1
    }
    var err2 = json.Unmarshal(src, &intermediateResult.ResponseRaw)
	if err2 != nil {
		return err2
	}

	*v = PaginatedBalanceSheetList(intermediateResult)
	return nil
}
type NullablePaginatedBalanceSheetList struct {
	value *PaginatedBalanceSheetList
	isSet bool
}

func (v NullablePaginatedBalanceSheetList) Get() *PaginatedBalanceSheetList {
	return v.value
}

func (v *NullablePaginatedBalanceSheetList) Set(val *PaginatedBalanceSheetList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedBalanceSheetList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedBalanceSheetList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedBalanceSheetList(val *PaginatedBalanceSheetList) *NullablePaginatedBalanceSheetList {
	return &NullablePaginatedBalanceSheetList{value: val, isSet: true}
}

func (v NullablePaginatedBalanceSheetList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedBalanceSheetList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	var err1 = json.Unmarshal(src, &v.value)
    if err1 != nil {
        return err1
    }
    return json.Unmarshal(src, &v.value.ResponseRaw)
}

