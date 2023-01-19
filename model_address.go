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

// Address # The Address Object ### Description The `Address` object is used to represent a contact's or company's address.  ### Usage Example Fetch from the `GET CompanyInfo` endpoint and view the company's addresses.
type Address struct {
	// The address type.
	Type NullableAddressTypeEnum `json:"type,omitempty"`
	// Line 1 of the address's street.
	Street1 NullableString `json:"street_1,omitempty"`
	// Line 2 of the address's street.
	Street2 NullableString `json:"street_2,omitempty"`
	// The address's city.
	City NullableString `json:"city,omitempty"`
	// The address's state or region.
	State interface{} `json:"state,omitempty"`
	// The address's state or region.
	CountrySubdivision NullableString `json:"country_subdivision,omitempty"`
	// The address's country.
	Country NullableCountryEnum `json:"country,omitempty"`
	// The address's zip code.
	ZipCode NullableString `json:"zip_code,omitempty"`
	// raw json response by property name
	ResponseRaw map[string]json.RawMessage `json:"-"`
}

// NewAddress instantiates a new Address object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddress() *Address {
	this := Address{}
	return &this
}

// NewAddressWithDefaults instantiates a new Address object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressWithDefaults() *Address {
	this := Address{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Address) GetType() AddressTypeEnum {
	if o == nil || o.Type.Get() == nil {
		var ret AddressTypeEnum
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Address) GetTypeOk() (*AddressTypeEnum, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *Address) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableAddressTypeEnum and assigns it to the Type field.
func (o *Address) SetType(v AddressTypeEnum) {
	o.Type.Set(&v)
}
// SetTypeNil sets the value for Type to be an explicit nil
func (o *Address) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *Address) UnsetType() {
	o.Type.Unset()
}

// GetStreet1 returns the Street1 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Address) GetStreet1() string {
	if o == nil || o.Street1.Get() == nil {
		var ret string
		return ret
	}
	return *o.Street1.Get()
}

// GetStreet1Ok returns a tuple with the Street1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Address) GetStreet1Ok() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Street1.Get(), o.Street1.IsSet()
}

// HasStreet1 returns a boolean if a field has been set.
func (o *Address) HasStreet1() bool {
	if o != nil && o.Street1.IsSet() {
		return true
	}

	return false
}

// SetStreet1 gets a reference to the given NullableString and assigns it to the Street1 field.
func (o *Address) SetStreet1(v string) {
	o.Street1.Set(&v)
}
// SetStreet1Nil sets the value for Street1 to be an explicit nil
func (o *Address) SetStreet1Nil() {
	o.Street1.Set(nil)
}

// UnsetStreet1 ensures that no value is present for Street1, not even an explicit nil
func (o *Address) UnsetStreet1() {
	o.Street1.Unset()
}

// GetStreet2 returns the Street2 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Address) GetStreet2() string {
	if o == nil || o.Street2.Get() == nil {
		var ret string
		return ret
	}
	return *o.Street2.Get()
}

// GetStreet2Ok returns a tuple with the Street2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Address) GetStreet2Ok() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Street2.Get(), o.Street2.IsSet()
}

// HasStreet2 returns a boolean if a field has been set.
func (o *Address) HasStreet2() bool {
	if o != nil && o.Street2.IsSet() {
		return true
	}

	return false
}

// SetStreet2 gets a reference to the given NullableString and assigns it to the Street2 field.
func (o *Address) SetStreet2(v string) {
	o.Street2.Set(&v)
}
// SetStreet2Nil sets the value for Street2 to be an explicit nil
func (o *Address) SetStreet2Nil() {
	o.Street2.Set(nil)
}

// UnsetStreet2 ensures that no value is present for Street2, not even an explicit nil
func (o *Address) UnsetStreet2() {
	o.Street2.Unset()
}

// GetCity returns the City field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Address) GetCity() string {
	if o == nil || o.City.Get() == nil {
		var ret string
		return ret
	}
	return *o.City.Get()
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Address) GetCityOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.City.Get(), o.City.IsSet()
}

// HasCity returns a boolean if a field has been set.
func (o *Address) HasCity() bool {
	if o != nil && o.City.IsSet() {
		return true
	}

	return false
}

// SetCity gets a reference to the given NullableString and assigns it to the City field.
func (o *Address) SetCity(v string) {
	o.City.Set(&v)
}
// SetCityNil sets the value for City to be an explicit nil
func (o *Address) SetCityNil() {
	o.City.Set(nil)
}

// UnsetCity ensures that no value is present for City, not even an explicit nil
func (o *Address) UnsetCity() {
	o.City.Unset()
}

// GetState returns the State field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Address) GetState() interface{} {
	if o == nil  {
		var ret interface{}
		return ret
	}
	return o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Address) GetStateOk() (*interface{}, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return &o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *Address) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given interface{} and assigns it to the State field.
func (o *Address) SetState(v interface{}) {
	o.State = v
}

// GetCountrySubdivision returns the CountrySubdivision field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Address) GetCountrySubdivision() string {
	if o == nil || o.CountrySubdivision.Get() == nil {
		var ret string
		return ret
	}
	return *o.CountrySubdivision.Get()
}

// GetCountrySubdivisionOk returns a tuple with the CountrySubdivision field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Address) GetCountrySubdivisionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CountrySubdivision.Get(), o.CountrySubdivision.IsSet()
}

// HasCountrySubdivision returns a boolean if a field has been set.
func (o *Address) HasCountrySubdivision() bool {
	if o != nil && o.CountrySubdivision.IsSet() {
		return true
	}

	return false
}

// SetCountrySubdivision gets a reference to the given NullableString and assigns it to the CountrySubdivision field.
func (o *Address) SetCountrySubdivision(v string) {
	o.CountrySubdivision.Set(&v)
}
// SetCountrySubdivisionNil sets the value for CountrySubdivision to be an explicit nil
func (o *Address) SetCountrySubdivisionNil() {
	o.CountrySubdivision.Set(nil)
}

// UnsetCountrySubdivision ensures that no value is present for CountrySubdivision, not even an explicit nil
func (o *Address) UnsetCountrySubdivision() {
	o.CountrySubdivision.Unset()
}

// GetCountry returns the Country field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Address) GetCountry() CountryEnum {
	if o == nil || o.Country.Get() == nil {
		var ret CountryEnum
		return ret
	}
	return *o.Country.Get()
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Address) GetCountryOk() (*CountryEnum, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Country.Get(), o.Country.IsSet()
}

// HasCountry returns a boolean if a field has been set.
func (o *Address) HasCountry() bool {
	if o != nil && o.Country.IsSet() {
		return true
	}

	return false
}

// SetCountry gets a reference to the given NullableCountryEnum and assigns it to the Country field.
func (o *Address) SetCountry(v CountryEnum) {
	o.Country.Set(&v)
}
// SetCountryNil sets the value for Country to be an explicit nil
func (o *Address) SetCountryNil() {
	o.Country.Set(nil)
}

// UnsetCountry ensures that no value is present for Country, not even an explicit nil
func (o *Address) UnsetCountry() {
	o.Country.Unset()
}

// GetZipCode returns the ZipCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Address) GetZipCode() string {
	if o == nil || o.ZipCode.Get() == nil {
		var ret string
		return ret
	}
	return *o.ZipCode.Get()
}

// GetZipCodeOk returns a tuple with the ZipCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Address) GetZipCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ZipCode.Get(), o.ZipCode.IsSet()
}

// HasZipCode returns a boolean if a field has been set.
func (o *Address) HasZipCode() bool {
	if o != nil && o.ZipCode.IsSet() {
		return true
	}

	return false
}

// SetZipCode gets a reference to the given NullableString and assigns it to the ZipCode field.
func (o *Address) SetZipCode(v string) {
	o.ZipCode.Set(&v)
}
// SetZipCodeNil sets the value for ZipCode to be an explicit nil
func (o *Address) SetZipCodeNil() {
	o.ZipCode.Set(nil)
}

// UnsetZipCode ensures that no value is present for ZipCode, not even an explicit nil
func (o *Address) UnsetZipCode() {
	o.ZipCode.Unset()
}

func (o Address) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}
	if o.Street1.IsSet() {
		toSerialize["street_1"] = o.Street1.Get()
	}
	if o.Street2.IsSet() {
		toSerialize["street_2"] = o.Street2.Get()
	}
	if o.City.IsSet() {
		toSerialize["city"] = o.City.Get()
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.CountrySubdivision.IsSet() {
		toSerialize["country_subdivision"] = o.CountrySubdivision.Get()
	}
	if o.Country.IsSet() {
		toSerialize["country"] = o.Country.Get()
	}
	if o.ZipCode.IsSet() {
		toSerialize["zip_code"] = o.ZipCode.Get()
	}
	return json.Marshal(toSerialize)
}

func (v *Address) UnmarshalJSON(src []byte) error {
    type AddressUnmarshalTarget Address

	var intermediateResult AddressUnmarshalTarget
	var err1 = json.Unmarshal(src, &intermediateResult)
    if err1 != nil {
        return err1
    }
    var err2 = json.Unmarshal(src, &intermediateResult.ResponseRaw)
	if err2 != nil {
		return err2
	}

	*v = Address(intermediateResult)
	return nil
}
type NullableAddress struct {
	value *Address
	isSet bool
}

func (v NullableAddress) Get() *Address {
	return v.value
}

func (v *NullableAddress) Set(val *Address) {
	v.value = val
	v.isSet = true
}

func (v NullableAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddress(val *Address) *NullableAddress {
	return &NullableAddress{value: val, isSet: true}
}

func (v NullableAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	var err1 = json.Unmarshal(src, &v.value)
    if err1 != nil {
        return err1
    }
    return json.Unmarshal(src, &v.value.ResponseRaw)
}


