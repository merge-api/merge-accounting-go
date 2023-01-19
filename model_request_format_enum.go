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
	"fmt"
)

// RequestFormatEnum the model 'RequestFormatEnum'
type RequestFormatEnum string

// apologies but this is to get around an import error
var _ = fmt.Printf
// List of RequestFormatEnum
const (
    REQUESTFORMATENUM_MERGE_NONSTANDARD_VALUE RequestFormatEnum = "MERGE_NONSTANDARD_VALUE"
    
	REQUESTFORMATENUM_JSON RequestFormatEnum = "JSON"
	REQUESTFORMATENUM_XML RequestFormatEnum = "XML"
	REQUESTFORMATENUM_MULTIPART RequestFormatEnum = "MULTIPART"
)

var allowedRequestFormatEnumEnumValues = []RequestFormatEnum{
	"JSON",
	"XML",
	"MULTIPART",
}

func (v *RequestFormatEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RequestFormatEnum(value)
	for _, existing := range allowedRequestFormatEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}
    *v = REQUESTFORMATENUM_MERGE_NONSTANDARD_VALUE
    return nil
}

// NewRequestFormatEnumFromValue returns a pointer to a valid RequestFormatEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRequestFormatEnumFromValue(v string) (*RequestFormatEnum, error) {
	ev := RequestFormatEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
        ev := REQUESTFORMATENUM_MERGE_NONSTANDARD_VALUE
        return &ev, nil
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RequestFormatEnum) IsValid() bool {
	for _, existing := range allowedRequestFormatEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RequestFormatEnum value
func (v RequestFormatEnum) Ptr() *RequestFormatEnum {
	return &v
}

type NullableRequestFormatEnum struct {
	value *RequestFormatEnum
	isSet bool
}

func (v NullableRequestFormatEnum) Get() *RequestFormatEnum {
	return v.value
}

func (v *NullableRequestFormatEnum) Set(val *RequestFormatEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestFormatEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestFormatEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestFormatEnum(val *RequestFormatEnum) *NullableRequestFormatEnum {
	return &NullableRequestFormatEnum{value: val, isSet: true}
}

func (v NullableRequestFormatEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestFormatEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

