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

// VendorCredit # The VendorCredit Object ### Description The `VendorCredit` object is an accounts receivable transaction used to show that a customer is owed a gift or refund. A vendor credit will contain information on the amount of credit owed to the customer, the vendor that owes the credit, and the account.  ### Usage Example Fetch from the `GET VendorCredit` endpoint and view a company's vendor credits.
type VendorCredit struct {
	Id *string `json:"id,omitempty"`
	// The third-party API ID of the matching object.
	RemoteId NullableString `json:"remote_id,omitempty"`
	RemoteData []RemoteData `json:"remote_data,omitempty"`
	// The vendor credit's number.
	Number NullableString `json:"number,omitempty"`
	// The vendor credit's transaction date.
	TransactionDate NullableTime `json:"transaction_date,omitempty"`
	// The vendor that owes the gift or refund.
	Vendor NullableString `json:"vendor,omitempty"`
	// The vendor credit's total amount.
	TotalAmount NullableFloat32 `json:"total_amount,omitempty"`
	// The vendor credit's currency.
	Currency NullableCurrencyEnum `json:"currency,omitempty"`
	// The vendor credit's exchange rate.
	ExchangeRate NullableFloat64 `json:"exchange_rate,omitempty"`
	// The company the vendor credit belongs to.
	Company NullableString `json:"company,omitempty"`
	Lines *[]VendorCreditLine `json:"lines,omitempty"`
	// Indicates whether or not this object has been deleted by third party webhooks.
	RemoteWasDeleted *bool `json:"remote_was_deleted,omitempty"`
	FieldMappings map[string]interface{} `json:"field_mappings,omitempty"`
	// raw json response by property name
	ResponseRaw map[string]json.RawMessage `json:"-"`
}

// NewVendorCredit instantiates a new VendorCredit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVendorCredit() *VendorCredit {
	this := VendorCredit{}
	return &this
}

// NewVendorCreditWithDefaults instantiates a new VendorCredit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVendorCreditWithDefaults() *VendorCredit {
	this := VendorCredit{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *VendorCredit) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VendorCredit) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *VendorCredit) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *VendorCredit) SetId(v string) {
	o.Id = &v
}

// GetRemoteId returns the RemoteId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VendorCredit) GetRemoteId() string {
	if o == nil || o.RemoteId.Get() == nil {
		var ret string
		return ret
	}
	return *o.RemoteId.Get()
}

// GetRemoteIdOk returns a tuple with the RemoteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VendorCredit) GetRemoteIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RemoteId.Get(), o.RemoteId.IsSet()
}

// HasRemoteId returns a boolean if a field has been set.
func (o *VendorCredit) HasRemoteId() bool {
	if o != nil && o.RemoteId.IsSet() {
		return true
	}

	return false
}

// SetRemoteId gets a reference to the given NullableString and assigns it to the RemoteId field.
func (o *VendorCredit) SetRemoteId(v string) {
	o.RemoteId.Set(&v)
}
// SetRemoteIdNil sets the value for RemoteId to be an explicit nil
func (o *VendorCredit) SetRemoteIdNil() {
	o.RemoteId.Set(nil)
}

// UnsetRemoteId ensures that no value is present for RemoteId, not even an explicit nil
func (o *VendorCredit) UnsetRemoteId() {
	o.RemoteId.Unset()
}

// GetRemoteData returns the RemoteData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VendorCredit) GetRemoteData() []RemoteData {
	if o == nil  {
		var ret []RemoteData
		return ret
	}
	return o.RemoteData
}

// GetRemoteDataOk returns a tuple with the RemoteData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VendorCredit) GetRemoteDataOk() (*[]RemoteData, bool) {
	if o == nil || o.RemoteData == nil {
		return nil, false
	}
	return &o.RemoteData, true
}

// HasRemoteData returns a boolean if a field has been set.
func (o *VendorCredit) HasRemoteData() bool {
	if o != nil && o.RemoteData != nil {
		return true
	}

	return false
}

// SetRemoteData gets a reference to the given []RemoteData and assigns it to the RemoteData field.
func (o *VendorCredit) SetRemoteData(v []RemoteData) {
	o.RemoteData = v
}

// GetNumber returns the Number field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VendorCredit) GetNumber() string {
	if o == nil || o.Number.Get() == nil {
		var ret string
		return ret
	}
	return *o.Number.Get()
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VendorCredit) GetNumberOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Number.Get(), o.Number.IsSet()
}

// HasNumber returns a boolean if a field has been set.
func (o *VendorCredit) HasNumber() bool {
	if o != nil && o.Number.IsSet() {
		return true
	}

	return false
}

// SetNumber gets a reference to the given NullableString and assigns it to the Number field.
func (o *VendorCredit) SetNumber(v string) {
	o.Number.Set(&v)
}
// SetNumberNil sets the value for Number to be an explicit nil
func (o *VendorCredit) SetNumberNil() {
	o.Number.Set(nil)
}

// UnsetNumber ensures that no value is present for Number, not even an explicit nil
func (o *VendorCredit) UnsetNumber() {
	o.Number.Unset()
}

// GetTransactionDate returns the TransactionDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VendorCredit) GetTransactionDate() time.Time {
	if o == nil || o.TransactionDate.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.TransactionDate.Get()
}

// GetTransactionDateOk returns a tuple with the TransactionDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VendorCredit) GetTransactionDateOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TransactionDate.Get(), o.TransactionDate.IsSet()
}

// HasTransactionDate returns a boolean if a field has been set.
func (o *VendorCredit) HasTransactionDate() bool {
	if o != nil && o.TransactionDate.IsSet() {
		return true
	}

	return false
}

// SetTransactionDate gets a reference to the given NullableTime and assigns it to the TransactionDate field.
func (o *VendorCredit) SetTransactionDate(v time.Time) {
	o.TransactionDate.Set(&v)
}
// SetTransactionDateNil sets the value for TransactionDate to be an explicit nil
func (o *VendorCredit) SetTransactionDateNil() {
	o.TransactionDate.Set(nil)
}

// UnsetTransactionDate ensures that no value is present for TransactionDate, not even an explicit nil
func (o *VendorCredit) UnsetTransactionDate() {
	o.TransactionDate.Unset()
}

// GetVendor returns the Vendor field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VendorCredit) GetVendor() string {
	if o == nil || o.Vendor.Get() == nil {
		var ret string
		return ret
	}
	return *o.Vendor.Get()
}

// GetVendorOk returns a tuple with the Vendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VendorCredit) GetVendorOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Vendor.Get(), o.Vendor.IsSet()
}

// HasVendor returns a boolean if a field has been set.
func (o *VendorCredit) HasVendor() bool {
	if o != nil && o.Vendor.IsSet() {
		return true
	}

	return false
}

// SetVendor gets a reference to the given NullableString and assigns it to the Vendor field.
func (o *VendorCredit) SetVendor(v string) {
	o.Vendor.Set(&v)
}
// SetVendorNil sets the value for Vendor to be an explicit nil
func (o *VendorCredit) SetVendorNil() {
	o.Vendor.Set(nil)
}

// UnsetVendor ensures that no value is present for Vendor, not even an explicit nil
func (o *VendorCredit) UnsetVendor() {
	o.Vendor.Unset()
}

// GetTotalAmount returns the TotalAmount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VendorCredit) GetTotalAmount() float32 {
	if o == nil || o.TotalAmount.Get() == nil {
		var ret float32
		return ret
	}
	return *o.TotalAmount.Get()
}

// GetTotalAmountOk returns a tuple with the TotalAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VendorCredit) GetTotalAmountOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TotalAmount.Get(), o.TotalAmount.IsSet()
}

// HasTotalAmount returns a boolean if a field has been set.
func (o *VendorCredit) HasTotalAmount() bool {
	if o != nil && o.TotalAmount.IsSet() {
		return true
	}

	return false
}

// SetTotalAmount gets a reference to the given NullableFloat32 and assigns it to the TotalAmount field.
func (o *VendorCredit) SetTotalAmount(v float32) {
	o.TotalAmount.Set(&v)
}
// SetTotalAmountNil sets the value for TotalAmount to be an explicit nil
func (o *VendorCredit) SetTotalAmountNil() {
	o.TotalAmount.Set(nil)
}

// UnsetTotalAmount ensures that no value is present for TotalAmount, not even an explicit nil
func (o *VendorCredit) UnsetTotalAmount() {
	o.TotalAmount.Unset()
}

// GetCurrency returns the Currency field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VendorCredit) GetCurrency() CurrencyEnum {
	if o == nil || o.Currency.Get() == nil {
		var ret CurrencyEnum
		return ret
	}
	return *o.Currency.Get()
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VendorCredit) GetCurrencyOk() (*CurrencyEnum, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Currency.Get(), o.Currency.IsSet()
}

// HasCurrency returns a boolean if a field has been set.
func (o *VendorCredit) HasCurrency() bool {
	if o != nil && o.Currency.IsSet() {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given NullableCurrencyEnum and assigns it to the Currency field.
func (o *VendorCredit) SetCurrency(v CurrencyEnum) {
	o.Currency.Set(&v)
}
// SetCurrencyNil sets the value for Currency to be an explicit nil
func (o *VendorCredit) SetCurrencyNil() {
	o.Currency.Set(nil)
}

// UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
func (o *VendorCredit) UnsetCurrency() {
	o.Currency.Unset()
}

// GetExchangeRate returns the ExchangeRate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VendorCredit) GetExchangeRate() float64 {
	if o == nil || o.ExchangeRate.Get() == nil {
		var ret float64
		return ret
	}
	return *o.ExchangeRate.Get()
}

// GetExchangeRateOk returns a tuple with the ExchangeRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VendorCredit) GetExchangeRateOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ExchangeRate.Get(), o.ExchangeRate.IsSet()
}

// HasExchangeRate returns a boolean if a field has been set.
func (o *VendorCredit) HasExchangeRate() bool {
	if o != nil && o.ExchangeRate.IsSet() {
		return true
	}

	return false
}

// SetExchangeRate gets a reference to the given NullableFloat64 and assigns it to the ExchangeRate field.
func (o *VendorCredit) SetExchangeRate(v float64) {
	o.ExchangeRate.Set(&v)
}
// SetExchangeRateNil sets the value for ExchangeRate to be an explicit nil
func (o *VendorCredit) SetExchangeRateNil() {
	o.ExchangeRate.Set(nil)
}

// UnsetExchangeRate ensures that no value is present for ExchangeRate, not even an explicit nil
func (o *VendorCredit) UnsetExchangeRate() {
	o.ExchangeRate.Unset()
}

// GetCompany returns the Company field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VendorCredit) GetCompany() string {
	if o == nil || o.Company.Get() == nil {
		var ret string
		return ret
	}
	return *o.Company.Get()
}

// GetCompanyOk returns a tuple with the Company field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VendorCredit) GetCompanyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Company.Get(), o.Company.IsSet()
}

// HasCompany returns a boolean if a field has been set.
func (o *VendorCredit) HasCompany() bool {
	if o != nil && o.Company.IsSet() {
		return true
	}

	return false
}

// SetCompany gets a reference to the given NullableString and assigns it to the Company field.
func (o *VendorCredit) SetCompany(v string) {
	o.Company.Set(&v)
}
// SetCompanyNil sets the value for Company to be an explicit nil
func (o *VendorCredit) SetCompanyNil() {
	o.Company.Set(nil)
}

// UnsetCompany ensures that no value is present for Company, not even an explicit nil
func (o *VendorCredit) UnsetCompany() {
	o.Company.Unset()
}

// GetLines returns the Lines field value if set, zero value otherwise.
func (o *VendorCredit) GetLines() []VendorCreditLine {
	if o == nil || o.Lines == nil {
		var ret []VendorCreditLine
		return ret
	}
	return *o.Lines
}

// GetLinesOk returns a tuple with the Lines field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VendorCredit) GetLinesOk() (*[]VendorCreditLine, bool) {
	if o == nil || o.Lines == nil {
		return nil, false
	}
	return o.Lines, true
}

// HasLines returns a boolean if a field has been set.
func (o *VendorCredit) HasLines() bool {
	if o != nil && o.Lines != nil {
		return true
	}

	return false
}

// SetLines gets a reference to the given []VendorCreditLine and assigns it to the Lines field.
func (o *VendorCredit) SetLines(v []VendorCreditLine) {
	o.Lines = &v
}

// GetRemoteWasDeleted returns the RemoteWasDeleted field value if set, zero value otherwise.
func (o *VendorCredit) GetRemoteWasDeleted() bool {
	if o == nil || o.RemoteWasDeleted == nil {
		var ret bool
		return ret
	}
	return *o.RemoteWasDeleted
}

// GetRemoteWasDeletedOk returns a tuple with the RemoteWasDeleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VendorCredit) GetRemoteWasDeletedOk() (*bool, bool) {
	if o == nil || o.RemoteWasDeleted == nil {
		return nil, false
	}
	return o.RemoteWasDeleted, true
}

// HasRemoteWasDeleted returns a boolean if a field has been set.
func (o *VendorCredit) HasRemoteWasDeleted() bool {
	if o != nil && o.RemoteWasDeleted != nil {
		return true
	}

	return false
}

// SetRemoteWasDeleted gets a reference to the given bool and assigns it to the RemoteWasDeleted field.
func (o *VendorCredit) SetRemoteWasDeleted(v bool) {
	o.RemoteWasDeleted = &v
}

// GetFieldMappings returns the FieldMappings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VendorCredit) GetFieldMappings() map[string]interface{} {
	if o == nil  {
		var ret map[string]interface{}
		return ret
	}
	return o.FieldMappings
}

// GetFieldMappingsOk returns a tuple with the FieldMappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VendorCredit) GetFieldMappingsOk() (*map[string]interface{}, bool) {
	if o == nil || o.FieldMappings == nil {
		return nil, false
	}
	return &o.FieldMappings, true
}

// HasFieldMappings returns a boolean if a field has been set.
func (o *VendorCredit) HasFieldMappings() bool {
	if o != nil && o.FieldMappings != nil {
		return true
	}

	return false
}

// SetFieldMappings gets a reference to the given map[string]interface{} and assigns it to the FieldMappings field.
func (o *VendorCredit) SetFieldMappings(v map[string]interface{}) {
	o.FieldMappings = v
}

func (o VendorCredit) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.RemoteId.IsSet() {
		toSerialize["remote_id"] = o.RemoteId.Get()
	}
	if o.RemoteData != nil {
		toSerialize["remote_data"] = o.RemoteData
	}
	if o.Number.IsSet() {
		toSerialize["number"] = o.Number.Get()
	}
	if o.TransactionDate.IsSet() {
		toSerialize["transaction_date"] = o.TransactionDate.Get()
	}
	if o.Vendor.IsSet() {
		toSerialize["vendor"] = o.Vendor.Get()
	}
	if o.TotalAmount.IsSet() {
		toSerialize["total_amount"] = o.TotalAmount.Get()
	}
	if o.Currency.IsSet() {
		toSerialize["currency"] = o.Currency.Get()
	}
	if o.ExchangeRate.IsSet() {
		toSerialize["exchange_rate"] = o.ExchangeRate.Get()
	}
	if o.Company.IsSet() {
		toSerialize["company"] = o.Company.Get()
	}
	if o.Lines != nil {
		toSerialize["lines"] = o.Lines
	}
	if o.RemoteWasDeleted != nil {
		toSerialize["remote_was_deleted"] = o.RemoteWasDeleted
	}
	if o.FieldMappings != nil {
		toSerialize["field_mappings"] = o.FieldMappings
	}
	return json.Marshal(toSerialize)
}

func (v *VendorCredit) UnmarshalJSON(src []byte) error {
    type VendorCreditUnmarshalTarget VendorCredit

	var intermediateResult VendorCreditUnmarshalTarget
	var err1 = json.Unmarshal(src, &intermediateResult)
    if err1 != nil {
        return err1
    }
    var err2 = json.Unmarshal(src, &intermediateResult.ResponseRaw)
	if err2 != nil {
		return err2
	}

	*v = VendorCredit(intermediateResult)
	return nil
}
type NullableVendorCredit struct {
	value *VendorCredit
	isSet bool
}

func (v NullableVendorCredit) Get() *VendorCredit {
	return v.value
}

func (v *NullableVendorCredit) Set(val *VendorCredit) {
	v.value = val
	v.isSet = true
}

func (v NullableVendorCredit) IsSet() bool {
	return v.isSet
}

func (v *NullableVendorCredit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVendorCredit(val *VendorCredit) *NullableVendorCredit {
	return &NullableVendorCredit{value: val, isSet: true}
}

func (v NullableVendorCredit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVendorCredit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	var err1 = json.Unmarshal(src, &v.value)
    if err1 != nil {
        return err1
    }
    return json.Unmarshal(src, &v.value.ResponseRaw)
}


