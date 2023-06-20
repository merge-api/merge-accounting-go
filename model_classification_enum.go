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

// ClassificationEnum * `ASSET` - ASSET * `EQUITY` - EQUITY * `EXPENSE` - EXPENSE * `LIABILITY` - LIABILITY * `REVENUE` - REVENUE
type ClassificationEnum string

// apologies but this is to get around an import error
var _ = fmt.Printf
// List of ClassificationEnum
const (
    CLASSIFICATIONENUM_MERGE_NONSTANDARD_VALUE ClassificationEnum = "MERGE_NONSTANDARD_VALUE"
    
	CLASSIFICATIONENUM_ASSET ClassificationEnum = "ASSET"
	CLASSIFICATIONENUM_EQUITY ClassificationEnum = "EQUITY"
	CLASSIFICATIONENUM_EXPENSE ClassificationEnum = "EXPENSE"
	CLASSIFICATIONENUM_LIABILITY ClassificationEnum = "LIABILITY"
	CLASSIFICATIONENUM_REVENUE ClassificationEnum = "REVENUE"
)

var allowedClassificationEnumEnumValues = []ClassificationEnum{
	"ASSET",
	"EQUITY",
	"EXPENSE",
	"LIABILITY",
	"REVENUE",
}

func (v *ClassificationEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ClassificationEnum(value)
	for _, existing := range allowedClassificationEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}
    *v = CLASSIFICATIONENUM_MERGE_NONSTANDARD_VALUE
    return nil
}

// NewClassificationEnumFromValue returns a pointer to a valid ClassificationEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewClassificationEnumFromValue(v string) (*ClassificationEnum, error) {
	ev := ClassificationEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
        ev := CLASSIFICATIONENUM_MERGE_NONSTANDARD_VALUE
        return &ev, nil
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ClassificationEnum) IsValid() bool {
	for _, existing := range allowedClassificationEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ClassificationEnum value
func (v ClassificationEnum) Ptr() *ClassificationEnum {
	return &v
}

type NullableClassificationEnum struct {
	value *ClassificationEnum
	isSet bool
}

func (v NullableClassificationEnum) Get() *ClassificationEnum {
	return v.value
}

func (v *NullableClassificationEnum) Set(val *ClassificationEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableClassificationEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableClassificationEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClassificationEnum(val *ClassificationEnum) *NullableClassificationEnum {
	return &NullableClassificationEnum{value: val, isSet: true}
}

func (v NullableClassificationEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClassificationEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

