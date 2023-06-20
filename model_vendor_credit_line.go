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

// VendorCreditLine # The VendorCreditLine Object ### Description The `VendorCreditLine` object is used to represent a vendor credit's line items.  ### Usage Example Fetch from the `GET VendorCredit` endpoint and view the vendor credit's line items.
type VendorCreditLine struct {
	// The third-party API ID of the matching object.
	RemoteId NullableString `json:"remote_id,omitempty"`
	// The full value of the credit.
	NetAmount NullableFloat64 `json:"net_amount,omitempty"`
	// The line's associated tracking category.
	TrackingCategory NullableString `json:"tracking_category,omitempty"`
	// The line's associated tracking categories.
	TrackingCategories []string `json:"tracking_categories"`
	// The line's description.
	Description NullableString `json:"description,omitempty"`
	// The line's account.
	Account NullableString `json:"account,omitempty"`
	// The company the line belongs to.
	Company NullableString `json:"company,omitempty"`
	// The vendor credit line item's exchange rate.
	ExchangeRate NullableFloat64 `json:"exchange_rate,omitempty"`
	// This is the datetime that this object was last updated by Merge
	ModifiedAt *time.Time `json:"modified_at,omitempty"`
	// raw json response by property name
	ResponseRaw map[string]json.RawMessage `json:"-"`
}

// NewVendorCreditLine instantiates a new VendorCreditLine object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVendorCreditLine(trackingCategories []string) *VendorCreditLine {
	this := VendorCreditLine{}
	this.TrackingCategories = trackingCategories
	return &this
}

// NewVendorCreditLineWithDefaults instantiates a new VendorCreditLine object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVendorCreditLineWithDefaults() *VendorCreditLine {
	this := VendorCreditLine{}
	return &this
}

// GetRemoteId returns the RemoteId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VendorCreditLine) GetRemoteId() string {
	if o == nil || o.RemoteId.Get() == nil {
		var ret string
		return ret
	}
	return *o.RemoteId.Get()
}

// GetRemoteIdOk returns a tuple with the RemoteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VendorCreditLine) GetRemoteIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RemoteId.Get(), o.RemoteId.IsSet()
}

// HasRemoteId returns a boolean if a field has been set.
func (o *VendorCreditLine) HasRemoteId() bool {
	if o != nil && o.RemoteId.IsSet() {
		return true
	}

	return false
}

// SetRemoteId gets a reference to the given NullableString and assigns it to the RemoteId field.
func (o *VendorCreditLine) SetRemoteId(v string) {
	o.RemoteId.Set(&v)
}
// SetRemoteIdNil sets the value for RemoteId to be an explicit nil
func (o *VendorCreditLine) SetRemoteIdNil() {
	o.RemoteId.Set(nil)
}

// UnsetRemoteId ensures that no value is present for RemoteId, not even an explicit nil
func (o *VendorCreditLine) UnsetRemoteId() {
	o.RemoteId.Unset()
}

// GetNetAmount returns the NetAmount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VendorCreditLine) GetNetAmount() float64 {
	if o == nil || o.NetAmount.Get() == nil {
		var ret float64
		return ret
	}
	return *o.NetAmount.Get()
}

// GetNetAmountOk returns a tuple with the NetAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VendorCreditLine) GetNetAmountOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.NetAmount.Get(), o.NetAmount.IsSet()
}

// HasNetAmount returns a boolean if a field has been set.
func (o *VendorCreditLine) HasNetAmount() bool {
	if o != nil && o.NetAmount.IsSet() {
		return true
	}

	return false
}

// SetNetAmount gets a reference to the given NullableFloat64 and assigns it to the NetAmount field.
func (o *VendorCreditLine) SetNetAmount(v float64) {
	o.NetAmount.Set(&v)
}
// SetNetAmountNil sets the value for NetAmount to be an explicit nil
func (o *VendorCreditLine) SetNetAmountNil() {
	o.NetAmount.Set(nil)
}

// UnsetNetAmount ensures that no value is present for NetAmount, not even an explicit nil
func (o *VendorCreditLine) UnsetNetAmount() {
	o.NetAmount.Unset()
}

// GetTrackingCategory returns the TrackingCategory field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VendorCreditLine) GetTrackingCategory() string {
	if o == nil || o.TrackingCategory.Get() == nil {
		var ret string
		return ret
	}
	return *o.TrackingCategory.Get()
}

// GetTrackingCategoryOk returns a tuple with the TrackingCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VendorCreditLine) GetTrackingCategoryOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TrackingCategory.Get(), o.TrackingCategory.IsSet()
}

// HasTrackingCategory returns a boolean if a field has been set.
func (o *VendorCreditLine) HasTrackingCategory() bool {
	if o != nil && o.TrackingCategory.IsSet() {
		return true
	}

	return false
}

// SetTrackingCategory gets a reference to the given NullableString and assigns it to the TrackingCategory field.
func (o *VendorCreditLine) SetTrackingCategory(v string) {
	o.TrackingCategory.Set(&v)
}
// SetTrackingCategoryNil sets the value for TrackingCategory to be an explicit nil
func (o *VendorCreditLine) SetTrackingCategoryNil() {
	o.TrackingCategory.Set(nil)
}

// UnsetTrackingCategory ensures that no value is present for TrackingCategory, not even an explicit nil
func (o *VendorCreditLine) UnsetTrackingCategory() {
	o.TrackingCategory.Unset()
}

// GetTrackingCategories returns the TrackingCategories field value
func (o *VendorCreditLine) GetTrackingCategories() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.TrackingCategories
}

// GetTrackingCategoriesOk returns a tuple with the TrackingCategories field value
// and a boolean to check if the value has been set.
func (o *VendorCreditLine) GetTrackingCategoriesOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TrackingCategories, true
}

// SetTrackingCategories sets field value
func (o *VendorCreditLine) SetTrackingCategories(v []string) {
	o.TrackingCategories = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VendorCreditLine) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VendorCreditLine) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *VendorCreditLine) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *VendorCreditLine) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *VendorCreditLine) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *VendorCreditLine) UnsetDescription() {
	o.Description.Unset()
}

// GetAccount returns the Account field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VendorCreditLine) GetAccount() string {
	if o == nil || o.Account.Get() == nil {
		var ret string
		return ret
	}
	return *o.Account.Get()
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VendorCreditLine) GetAccountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Account.Get(), o.Account.IsSet()
}

// HasAccount returns a boolean if a field has been set.
func (o *VendorCreditLine) HasAccount() bool {
	if o != nil && o.Account.IsSet() {
		return true
	}

	return false
}

// SetAccount gets a reference to the given NullableString and assigns it to the Account field.
func (o *VendorCreditLine) SetAccount(v string) {
	o.Account.Set(&v)
}
// SetAccountNil sets the value for Account to be an explicit nil
func (o *VendorCreditLine) SetAccountNil() {
	o.Account.Set(nil)
}

// UnsetAccount ensures that no value is present for Account, not even an explicit nil
func (o *VendorCreditLine) UnsetAccount() {
	o.Account.Unset()
}

// GetCompany returns the Company field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VendorCreditLine) GetCompany() string {
	if o == nil || o.Company.Get() == nil {
		var ret string
		return ret
	}
	return *o.Company.Get()
}

// GetCompanyOk returns a tuple with the Company field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VendorCreditLine) GetCompanyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Company.Get(), o.Company.IsSet()
}

// HasCompany returns a boolean if a field has been set.
func (o *VendorCreditLine) HasCompany() bool {
	if o != nil && o.Company.IsSet() {
		return true
	}

	return false
}

// SetCompany gets a reference to the given NullableString and assigns it to the Company field.
func (o *VendorCreditLine) SetCompany(v string) {
	o.Company.Set(&v)
}
// SetCompanyNil sets the value for Company to be an explicit nil
func (o *VendorCreditLine) SetCompanyNil() {
	o.Company.Set(nil)
}

// UnsetCompany ensures that no value is present for Company, not even an explicit nil
func (o *VendorCreditLine) UnsetCompany() {
	o.Company.Unset()
}

// GetExchangeRate returns the ExchangeRate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VendorCreditLine) GetExchangeRate() float64 {
	if o == nil || o.ExchangeRate.Get() == nil {
		var ret float64
		return ret
	}
	return *o.ExchangeRate.Get()
}

// GetExchangeRateOk returns a tuple with the ExchangeRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VendorCreditLine) GetExchangeRateOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ExchangeRate.Get(), o.ExchangeRate.IsSet()
}

// HasExchangeRate returns a boolean if a field has been set.
func (o *VendorCreditLine) HasExchangeRate() bool {
	if o != nil && o.ExchangeRate.IsSet() {
		return true
	}

	return false
}

// SetExchangeRate gets a reference to the given NullableFloat64 and assigns it to the ExchangeRate field.
func (o *VendorCreditLine) SetExchangeRate(v float64) {
	o.ExchangeRate.Set(&v)
}
// SetExchangeRateNil sets the value for ExchangeRate to be an explicit nil
func (o *VendorCreditLine) SetExchangeRateNil() {
	o.ExchangeRate.Set(nil)
}

// UnsetExchangeRate ensures that no value is present for ExchangeRate, not even an explicit nil
func (o *VendorCreditLine) UnsetExchangeRate() {
	o.ExchangeRate.Unset()
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise.
func (o *VendorCreditLine) GetModifiedAt() time.Time {
	if o == nil || o.ModifiedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VendorCreditLine) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil || o.ModifiedAt == nil {
		return nil, false
	}
	return o.ModifiedAt, true
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *VendorCreditLine) HasModifiedAt() bool {
	if o != nil && o.ModifiedAt != nil {
		return true
	}

	return false
}

// SetModifiedAt gets a reference to the given time.Time and assigns it to the ModifiedAt field.
func (o *VendorCreditLine) SetModifiedAt(v time.Time) {
	o.ModifiedAt = &v
}

func (o VendorCreditLine) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RemoteId.IsSet() {
		toSerialize["remote_id"] = o.RemoteId.Get()
	}
	if o.NetAmount.IsSet() {
		toSerialize["net_amount"] = o.NetAmount.Get()
	}
	if o.TrackingCategory.IsSet() {
		toSerialize["tracking_category"] = o.TrackingCategory.Get()
	}
	if true {
		toSerialize["tracking_categories"] = o.TrackingCategories
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.Account.IsSet() {
		toSerialize["account"] = o.Account.Get()
	}
	if o.Company.IsSet() {
		toSerialize["company"] = o.Company.Get()
	}
	if o.ExchangeRate.IsSet() {
		toSerialize["exchange_rate"] = o.ExchangeRate.Get()
	}
	if o.ModifiedAt != nil {
		toSerialize["modified_at"] = o.ModifiedAt
	}
	return json.Marshal(toSerialize)
}

func (v *VendorCreditLine) UnmarshalJSON(src []byte) error {
    type VendorCreditLineUnmarshalTarget VendorCreditLine

	var intermediateResult VendorCreditLineUnmarshalTarget
	var err1 = json.Unmarshal(src, &intermediateResult)
    if err1 != nil {
        return err1
    }
    var err2 = json.Unmarshal(src, &intermediateResult.ResponseRaw)
	if err2 != nil {
		return err2
	}

	*v = VendorCreditLine(intermediateResult)
	return nil
}
type NullableVendorCreditLine struct {
	value *VendorCreditLine
	isSet bool
}

func (v NullableVendorCreditLine) Get() *VendorCreditLine {
	return v.value
}

func (v *NullableVendorCreditLine) Set(val *VendorCreditLine) {
	v.value = val
	v.isSet = true
}

func (v NullableVendorCreditLine) IsSet() bool {
	return v.isSet
}

func (v *NullableVendorCreditLine) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVendorCreditLine(val *VendorCreditLine) *NullableVendorCreditLine {
	return &NullableVendorCreditLine{value: val, isSet: true}
}

func (v NullableVendorCreditLine) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVendorCreditLine) UnmarshalJSON(src []byte) error {
	v.isSet = true
	var err1 = json.Unmarshal(src, &v.value)
    if err1 != nil {
        return err1
    }
    return json.Unmarshal(src, &v.value.ResponseRaw)
}


