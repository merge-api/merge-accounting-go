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

// CreditNoteStatusEnum * `SUBMITTED` - SUBMITTED * `AUTHORIZED` - AUTHORIZED * `PAID` - PAID
type CreditNoteStatusEnum string

// apologies but this is to get around an import error
var _ = fmt.Printf
// List of CreditNoteStatusEnum
const (
    CREDITNOTESTATUSENUM_MERGE_NONSTANDARD_VALUE CreditNoteStatusEnum = "MERGE_NONSTANDARD_VALUE"
    
	CREDITNOTESTATUSENUM_SUBMITTED CreditNoteStatusEnum = "SUBMITTED"
	CREDITNOTESTATUSENUM_AUTHORIZED CreditNoteStatusEnum = "AUTHORIZED"
	CREDITNOTESTATUSENUM_PAID CreditNoteStatusEnum = "PAID"
)

var allowedCreditNoteStatusEnumEnumValues = []CreditNoteStatusEnum{
	"SUBMITTED",
	"AUTHORIZED",
	"PAID",
}

func (v *CreditNoteStatusEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CreditNoteStatusEnum(value)
	for _, existing := range allowedCreditNoteStatusEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}
    *v = CREDITNOTESTATUSENUM_MERGE_NONSTANDARD_VALUE
    return nil
}

// NewCreditNoteStatusEnumFromValue returns a pointer to a valid CreditNoteStatusEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreditNoteStatusEnumFromValue(v string) (*CreditNoteStatusEnum, error) {
	ev := CreditNoteStatusEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
        ev := CREDITNOTESTATUSENUM_MERGE_NONSTANDARD_VALUE
        return &ev, nil
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreditNoteStatusEnum) IsValid() bool {
	for _, existing := range allowedCreditNoteStatusEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CreditNoteStatusEnum value
func (v CreditNoteStatusEnum) Ptr() *CreditNoteStatusEnum {
	return &v
}

type NullableCreditNoteStatusEnum struct {
	value *CreditNoteStatusEnum
	isSet bool
}

func (v NullableCreditNoteStatusEnum) Get() *CreditNoteStatusEnum {
	return v.value
}

func (v *NullableCreditNoteStatusEnum) Set(val *CreditNoteStatusEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditNoteStatusEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditNoteStatusEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditNoteStatusEnum(val *CreditNoteStatusEnum) *NullableCreditNoteStatusEnum {
	return &NullableCreditNoteStatusEnum{value: val, isSet: true}
}

func (v NullableCreditNoteStatusEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditNoteStatusEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

