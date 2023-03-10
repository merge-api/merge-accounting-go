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

// PaginatedTransactionList struct for PaginatedTransactionList
type PaginatedTransactionList struct {
	Next NullableString `json:"next,omitempty"`
	Previous NullableString `json:"previous,omitempty"`
	Results *[]Transaction `json:"results,omitempty"`
	// raw json response by property name
	ResponseRaw map[string]json.RawMessage `json:"-"`
}

// NewPaginatedTransactionList instantiates a new PaginatedTransactionList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedTransactionList() *PaginatedTransactionList {
	this := PaginatedTransactionList{}
	return &this
}

// NewPaginatedTransactionListWithDefaults instantiates a new PaginatedTransactionList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedTransactionListWithDefaults() *PaginatedTransactionList {
	this := PaginatedTransactionList{}
	return &this
}

// GetNext returns the Next field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaginatedTransactionList) GetNext() string {
	if o == nil || o.Next.Get() == nil {
		var ret string
		return ret
	}
	return *o.Next.Get()
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaginatedTransactionList) GetNextOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Next.Get(), o.Next.IsSet()
}

// HasNext returns a boolean if a field has been set.
func (o *PaginatedTransactionList) HasNext() bool {
	if o != nil && o.Next.IsSet() {
		return true
	}

	return false
}

// SetNext gets a reference to the given NullableString and assigns it to the Next field.
func (o *PaginatedTransactionList) SetNext(v string) {
	o.Next.Set(&v)
}
// SetNextNil sets the value for Next to be an explicit nil
func (o *PaginatedTransactionList) SetNextNil() {
	o.Next.Set(nil)
}

// UnsetNext ensures that no value is present for Next, not even an explicit nil
func (o *PaginatedTransactionList) UnsetNext() {
	o.Next.Unset()
}

// GetPrevious returns the Previous field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaginatedTransactionList) GetPrevious() string {
	if o == nil || o.Previous.Get() == nil {
		var ret string
		return ret
	}
	return *o.Previous.Get()
}

// GetPreviousOk returns a tuple with the Previous field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaginatedTransactionList) GetPreviousOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Previous.Get(), o.Previous.IsSet()
}

// HasPrevious returns a boolean if a field has been set.
func (o *PaginatedTransactionList) HasPrevious() bool {
	if o != nil && o.Previous.IsSet() {
		return true
	}

	return false
}

// SetPrevious gets a reference to the given NullableString and assigns it to the Previous field.
func (o *PaginatedTransactionList) SetPrevious(v string) {
	o.Previous.Set(&v)
}
// SetPreviousNil sets the value for Previous to be an explicit nil
func (o *PaginatedTransactionList) SetPreviousNil() {
	o.Previous.Set(nil)
}

// UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
func (o *PaginatedTransactionList) UnsetPrevious() {
	o.Previous.Unset()
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *PaginatedTransactionList) GetResults() []Transaction {
	if o == nil || o.Results == nil {
		var ret []Transaction
		return ret
	}
	return *o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedTransactionList) GetResultsOk() (*[]Transaction, bool) {
	if o == nil || o.Results == nil {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *PaginatedTransactionList) HasResults() bool {
	if o != nil && o.Results != nil {
		return true
	}

	return false
}

// SetResults gets a reference to the given []Transaction and assigns it to the Results field.
func (o *PaginatedTransactionList) SetResults(v []Transaction) {
	o.Results = &v
}

func (o PaginatedTransactionList) MarshalJSON() ([]byte, error) {
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

func (v *PaginatedTransactionList) UnmarshalJSON(src []byte) error {
    type PaginatedTransactionListUnmarshalTarget PaginatedTransactionList

	var intermediateResult PaginatedTransactionListUnmarshalTarget
	var err1 = json.Unmarshal(src, &intermediateResult)
    if err1 != nil {
        return err1
    }
    var err2 = json.Unmarshal(src, &intermediateResult.ResponseRaw)
	if err2 != nil {
		return err2
	}

	*v = PaginatedTransactionList(intermediateResult)
	return nil
}
type NullablePaginatedTransactionList struct {
	value *PaginatedTransactionList
	isSet bool
}

func (v NullablePaginatedTransactionList) Get() *PaginatedTransactionList {
	return v.value
}

func (v *NullablePaginatedTransactionList) Set(val *PaginatedTransactionList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedTransactionList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedTransactionList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedTransactionList(val *PaginatedTransactionList) *NullablePaginatedTransactionList {
	return &NullablePaginatedTransactionList{value: val, isSet: true}
}

func (v NullablePaginatedTransactionList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedTransactionList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	var err1 = json.Unmarshal(src, &v.value)
    if err1 != nil {
        return err1
    }
    return json.Unmarshal(src, &v.value.ResponseRaw)
}


