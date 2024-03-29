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

// JournalLine # The JournalLine Object ### Description The `JournalLine` object is used to represent a journal entry's line items.  ### Usage Example Fetch from the `GET JournalEntry` endpoint and view the journal entry's line items.
type JournalLine struct {
	// The third-party API ID of the matching object.
	RemoteId NullableString `json:"remote_id,omitempty"`
	Account NullableString `json:"account,omitempty"`
	// The value of the line item including taxes and other fees.
	NetAmount NullableFloat64 `json:"net_amount,omitempty"`
	TrackingCategory NullableString `json:"tracking_category,omitempty"`
	TrackingCategories *[]string `json:"tracking_categories,omitempty"`
	Contact NullableString `json:"contact,omitempty"`
	// The line's description.
	Description NullableString `json:"description,omitempty"`
	// The journal line item's exchange rate.
	ExchangeRate NullableFloat64 `json:"exchange_rate,omitempty"`
	// This is the datetime that this object was last updated by Merge
	ModifiedAt *time.Time `json:"modified_at,omitempty"`
	// raw json response by property name
	ResponseRaw map[string]json.RawMessage `json:"-"`
}

// NewJournalLine instantiates a new JournalLine object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJournalLine() *JournalLine {
	this := JournalLine{}
	return &this
}

// NewJournalLineWithDefaults instantiates a new JournalLine object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJournalLineWithDefaults() *JournalLine {
	this := JournalLine{}
	return &this
}

// GetRemoteId returns the RemoteId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JournalLine) GetRemoteId() string {
	if o == nil || o.RemoteId.Get() == nil {
		var ret string
		return ret
	}
	return *o.RemoteId.Get()
}

// GetRemoteIdOk returns a tuple with the RemoteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JournalLine) GetRemoteIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RemoteId.Get(), o.RemoteId.IsSet()
}

// HasRemoteId returns a boolean if a field has been set.
func (o *JournalLine) HasRemoteId() bool {
	if o != nil && o.RemoteId.IsSet() {
		return true
	}

	return false
}

// SetRemoteId gets a reference to the given NullableString and assigns it to the RemoteId field.
func (o *JournalLine) SetRemoteId(v string) {
	o.RemoteId.Set(&v)
}
// SetRemoteIdNil sets the value for RemoteId to be an explicit nil
func (o *JournalLine) SetRemoteIdNil() {
	o.RemoteId.Set(nil)
}

// UnsetRemoteId ensures that no value is present for RemoteId, not even an explicit nil
func (o *JournalLine) UnsetRemoteId() {
	o.RemoteId.Unset()
}

// GetAccount returns the Account field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JournalLine) GetAccount() string {
	if o == nil || o.Account.Get() == nil {
		var ret string
		return ret
	}
	return *o.Account.Get()
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JournalLine) GetAccountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Account.Get(), o.Account.IsSet()
}

// HasAccount returns a boolean if a field has been set.
func (o *JournalLine) HasAccount() bool {
	if o != nil && o.Account.IsSet() {
		return true
	}

	return false
}

// SetAccount gets a reference to the given NullableString and assigns it to the Account field.
func (o *JournalLine) SetAccount(v string) {
	o.Account.Set(&v)
}
// SetAccountNil sets the value for Account to be an explicit nil
func (o *JournalLine) SetAccountNil() {
	o.Account.Set(nil)
}

// UnsetAccount ensures that no value is present for Account, not even an explicit nil
func (o *JournalLine) UnsetAccount() {
	o.Account.Unset()
}

// GetNetAmount returns the NetAmount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JournalLine) GetNetAmount() float64 {
	if o == nil || o.NetAmount.Get() == nil {
		var ret float64
		return ret
	}
	return *o.NetAmount.Get()
}

// GetNetAmountOk returns a tuple with the NetAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JournalLine) GetNetAmountOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.NetAmount.Get(), o.NetAmount.IsSet()
}

// HasNetAmount returns a boolean if a field has been set.
func (o *JournalLine) HasNetAmount() bool {
	if o != nil && o.NetAmount.IsSet() {
		return true
	}

	return false
}

// SetNetAmount gets a reference to the given NullableFloat64 and assigns it to the NetAmount field.
func (o *JournalLine) SetNetAmount(v float64) {
	o.NetAmount.Set(&v)
}
// SetNetAmountNil sets the value for NetAmount to be an explicit nil
func (o *JournalLine) SetNetAmountNil() {
	o.NetAmount.Set(nil)
}

// UnsetNetAmount ensures that no value is present for NetAmount, not even an explicit nil
func (o *JournalLine) UnsetNetAmount() {
	o.NetAmount.Unset()
}

// GetTrackingCategory returns the TrackingCategory field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JournalLine) GetTrackingCategory() string {
	if o == nil || o.TrackingCategory.Get() == nil {
		var ret string
		return ret
	}
	return *o.TrackingCategory.Get()
}

// GetTrackingCategoryOk returns a tuple with the TrackingCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JournalLine) GetTrackingCategoryOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TrackingCategory.Get(), o.TrackingCategory.IsSet()
}

// HasTrackingCategory returns a boolean if a field has been set.
func (o *JournalLine) HasTrackingCategory() bool {
	if o != nil && o.TrackingCategory.IsSet() {
		return true
	}

	return false
}

// SetTrackingCategory gets a reference to the given NullableString and assigns it to the TrackingCategory field.
func (o *JournalLine) SetTrackingCategory(v string) {
	o.TrackingCategory.Set(&v)
}
// SetTrackingCategoryNil sets the value for TrackingCategory to be an explicit nil
func (o *JournalLine) SetTrackingCategoryNil() {
	o.TrackingCategory.Set(nil)
}

// UnsetTrackingCategory ensures that no value is present for TrackingCategory, not even an explicit nil
func (o *JournalLine) UnsetTrackingCategory() {
	o.TrackingCategory.Unset()
}

// GetTrackingCategories returns the TrackingCategories field value if set, zero value otherwise.
func (o *JournalLine) GetTrackingCategories() []string {
	if o == nil || o.TrackingCategories == nil {
		var ret []string
		return ret
	}
	return *o.TrackingCategories
}

// GetTrackingCategoriesOk returns a tuple with the TrackingCategories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JournalLine) GetTrackingCategoriesOk() (*[]string, bool) {
	if o == nil || o.TrackingCategories == nil {
		return nil, false
	}
	return o.TrackingCategories, true
}

// HasTrackingCategories returns a boolean if a field has been set.
func (o *JournalLine) HasTrackingCategories() bool {
	if o != nil && o.TrackingCategories != nil {
		return true
	}

	return false
}

// SetTrackingCategories gets a reference to the given []string and assigns it to the TrackingCategories field.
func (o *JournalLine) SetTrackingCategories(v []string) {
	o.TrackingCategories = &v
}

// GetContact returns the Contact field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JournalLine) GetContact() string {
	if o == nil || o.Contact.Get() == nil {
		var ret string
		return ret
	}
	return *o.Contact.Get()
}

// GetContactOk returns a tuple with the Contact field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JournalLine) GetContactOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Contact.Get(), o.Contact.IsSet()
}

// HasContact returns a boolean if a field has been set.
func (o *JournalLine) HasContact() bool {
	if o != nil && o.Contact.IsSet() {
		return true
	}

	return false
}

// SetContact gets a reference to the given NullableString and assigns it to the Contact field.
func (o *JournalLine) SetContact(v string) {
	o.Contact.Set(&v)
}
// SetContactNil sets the value for Contact to be an explicit nil
func (o *JournalLine) SetContactNil() {
	o.Contact.Set(nil)
}

// UnsetContact ensures that no value is present for Contact, not even an explicit nil
func (o *JournalLine) UnsetContact() {
	o.Contact.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JournalLine) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JournalLine) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *JournalLine) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *JournalLine) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *JournalLine) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *JournalLine) UnsetDescription() {
	o.Description.Unset()
}

// GetExchangeRate returns the ExchangeRate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JournalLine) GetExchangeRate() float64 {
	if o == nil || o.ExchangeRate.Get() == nil {
		var ret float64
		return ret
	}
	return *o.ExchangeRate.Get()
}

// GetExchangeRateOk returns a tuple with the ExchangeRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JournalLine) GetExchangeRateOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ExchangeRate.Get(), o.ExchangeRate.IsSet()
}

// HasExchangeRate returns a boolean if a field has been set.
func (o *JournalLine) HasExchangeRate() bool {
	if o != nil && o.ExchangeRate.IsSet() {
		return true
	}

	return false
}

// SetExchangeRate gets a reference to the given NullableFloat64 and assigns it to the ExchangeRate field.
func (o *JournalLine) SetExchangeRate(v float64) {
	o.ExchangeRate.Set(&v)
}
// SetExchangeRateNil sets the value for ExchangeRate to be an explicit nil
func (o *JournalLine) SetExchangeRateNil() {
	o.ExchangeRate.Set(nil)
}

// UnsetExchangeRate ensures that no value is present for ExchangeRate, not even an explicit nil
func (o *JournalLine) UnsetExchangeRate() {
	o.ExchangeRate.Unset()
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise.
func (o *JournalLine) GetModifiedAt() time.Time {
	if o == nil || o.ModifiedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JournalLine) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil || o.ModifiedAt == nil {
		return nil, false
	}
	return o.ModifiedAt, true
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *JournalLine) HasModifiedAt() bool {
	if o != nil && o.ModifiedAt != nil {
		return true
	}

	return false
}

// SetModifiedAt gets a reference to the given time.Time and assigns it to the ModifiedAt field.
func (o *JournalLine) SetModifiedAt(v time.Time) {
	o.ModifiedAt = &v
}

func (o JournalLine) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RemoteId.IsSet() {
		toSerialize["remote_id"] = o.RemoteId.Get()
	}
	if o.Account.IsSet() {
		toSerialize["account"] = o.Account.Get()
	}
	if o.NetAmount.IsSet() {
		toSerialize["net_amount"] = o.NetAmount.Get()
	}
	if o.TrackingCategory.IsSet() {
		toSerialize["tracking_category"] = o.TrackingCategory.Get()
	}
	if o.TrackingCategories != nil {
		toSerialize["tracking_categories"] = o.TrackingCategories
	}
	if o.Contact.IsSet() {
		toSerialize["contact"] = o.Contact.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.ExchangeRate.IsSet() {
		toSerialize["exchange_rate"] = o.ExchangeRate.Get()
	}
	if o.ModifiedAt != nil {
		toSerialize["modified_at"] = o.ModifiedAt
	}
	return json.Marshal(toSerialize)
}

func (v *JournalLine) UnmarshalJSON(src []byte) error {
    type JournalLineUnmarshalTarget JournalLine

	var intermediateResult JournalLineUnmarshalTarget
	var err1 = json.Unmarshal(src, &intermediateResult)
    if err1 != nil {
        return err1
    }
    var err2 = json.Unmarshal(src, &intermediateResult.ResponseRaw)
	if err2 != nil {
		return err2
	}

	*v = JournalLine(intermediateResult)
	return nil
}
type NullableJournalLine struct {
	value *JournalLine
	isSet bool
}

func (v NullableJournalLine) Get() *JournalLine {
	return v.value
}

func (v *NullableJournalLine) Set(val *JournalLine) {
	v.value = val
	v.isSet = true
}

func (v NullableJournalLine) IsSet() bool {
	return v.isSet
}

func (v *NullableJournalLine) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJournalLine(val *JournalLine) *NullableJournalLine {
	return &NullableJournalLine{value: val, isSet: true}
}

func (v NullableJournalLine) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJournalLine) UnmarshalJSON(src []byte) error {
	v.isSet = true
	var err1 = json.Unmarshal(src, &v.value)
    if err1 != nil {
        return err1
    }
    return json.Unmarshal(src, &v.value.ResponseRaw)
}


