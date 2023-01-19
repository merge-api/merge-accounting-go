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

// InvoiceLineItem # The InvoiceLineItem Object ### Description The `InvoiceLineItem` object represents an itemized record of goods and/or services sold to a customer. If type = accounts_payable, invoice is a bill, if type = accounts_receivable it's an invoice.  ### Usage Example Fetch from the `GET Invoice` endpoint and view the invoice's line items.
type InvoiceLineItem struct {
	Id *string `json:"id,omitempty"`
	// The third-party API ID of the matching object.
	RemoteId NullableString `json:"remote_id,omitempty"`
	// The line item's description.
	Description NullableString `json:"description,omitempty"`
	// The line item's unit price.
	UnitPrice NullableFloat32 `json:"unit_price,omitempty"`
	// The line item's quantity.
	Quantity NullableFloat32 `json:"quantity,omitempty"`
	// The line item's total amount.
	TotalAmount NullableFloat32 `json:"total_amount,omitempty"`
	// The line item's currency.
	Currency NullableCurrencyEnum `json:"currency,omitempty"`
	// The line item's exchange rate.
	ExchangeRate NullableFloat64 `json:"exchange_rate,omitempty"`
	Item NullableString `json:"item,omitempty"`
	Account NullableString `json:"account,omitempty"`
	TrackingCategory NullableString `json:"tracking_category,omitempty"`
	TrackingCategories *[]string `json:"tracking_categories,omitempty"`
	// The company the line item belongs to.
	Company NullableString `json:"company,omitempty"`
	FieldMappings map[string]interface{} `json:"field_mappings,omitempty"`
	// raw json response by property name
	ResponseRaw map[string]json.RawMessage `json:"-"`
}

// NewInvoiceLineItem instantiates a new InvoiceLineItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvoiceLineItem() *InvoiceLineItem {
	this := InvoiceLineItem{}
	return &this
}

// NewInvoiceLineItemWithDefaults instantiates a new InvoiceLineItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvoiceLineItemWithDefaults() *InvoiceLineItem {
	this := InvoiceLineItem{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InvoiceLineItem) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceLineItem) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InvoiceLineItem) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InvoiceLineItem) SetId(v string) {
	o.Id = &v
}

// GetRemoteId returns the RemoteId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InvoiceLineItem) GetRemoteId() string {
	if o == nil || o.RemoteId.Get() == nil {
		var ret string
		return ret
	}
	return *o.RemoteId.Get()
}

// GetRemoteIdOk returns a tuple with the RemoteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvoiceLineItem) GetRemoteIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RemoteId.Get(), o.RemoteId.IsSet()
}

// HasRemoteId returns a boolean if a field has been set.
func (o *InvoiceLineItem) HasRemoteId() bool {
	if o != nil && o.RemoteId.IsSet() {
		return true
	}

	return false
}

// SetRemoteId gets a reference to the given NullableString and assigns it to the RemoteId field.
func (o *InvoiceLineItem) SetRemoteId(v string) {
	o.RemoteId.Set(&v)
}
// SetRemoteIdNil sets the value for RemoteId to be an explicit nil
func (o *InvoiceLineItem) SetRemoteIdNil() {
	o.RemoteId.Set(nil)
}

// UnsetRemoteId ensures that no value is present for RemoteId, not even an explicit nil
func (o *InvoiceLineItem) UnsetRemoteId() {
	o.RemoteId.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InvoiceLineItem) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvoiceLineItem) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *InvoiceLineItem) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *InvoiceLineItem) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *InvoiceLineItem) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *InvoiceLineItem) UnsetDescription() {
	o.Description.Unset()
}

// GetUnitPrice returns the UnitPrice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InvoiceLineItem) GetUnitPrice() float32 {
	if o == nil || o.UnitPrice.Get() == nil {
		var ret float32
		return ret
	}
	return *o.UnitPrice.Get()
}

// GetUnitPriceOk returns a tuple with the UnitPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvoiceLineItem) GetUnitPriceOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UnitPrice.Get(), o.UnitPrice.IsSet()
}

// HasUnitPrice returns a boolean if a field has been set.
func (o *InvoiceLineItem) HasUnitPrice() bool {
	if o != nil && o.UnitPrice.IsSet() {
		return true
	}

	return false
}

// SetUnitPrice gets a reference to the given NullableFloat32 and assigns it to the UnitPrice field.
func (o *InvoiceLineItem) SetUnitPrice(v float32) {
	o.UnitPrice.Set(&v)
}
// SetUnitPriceNil sets the value for UnitPrice to be an explicit nil
func (o *InvoiceLineItem) SetUnitPriceNil() {
	o.UnitPrice.Set(nil)
}

// UnsetUnitPrice ensures that no value is present for UnitPrice, not even an explicit nil
func (o *InvoiceLineItem) UnsetUnitPrice() {
	o.UnitPrice.Unset()
}

// GetQuantity returns the Quantity field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InvoiceLineItem) GetQuantity() float32 {
	if o == nil || o.Quantity.Get() == nil {
		var ret float32
		return ret
	}
	return *o.Quantity.Get()
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvoiceLineItem) GetQuantityOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Quantity.Get(), o.Quantity.IsSet()
}

// HasQuantity returns a boolean if a field has been set.
func (o *InvoiceLineItem) HasQuantity() bool {
	if o != nil && o.Quantity.IsSet() {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given NullableFloat32 and assigns it to the Quantity field.
func (o *InvoiceLineItem) SetQuantity(v float32) {
	o.Quantity.Set(&v)
}
// SetQuantityNil sets the value for Quantity to be an explicit nil
func (o *InvoiceLineItem) SetQuantityNil() {
	o.Quantity.Set(nil)
}

// UnsetQuantity ensures that no value is present for Quantity, not even an explicit nil
func (o *InvoiceLineItem) UnsetQuantity() {
	o.Quantity.Unset()
}

// GetTotalAmount returns the TotalAmount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InvoiceLineItem) GetTotalAmount() float32 {
	if o == nil || o.TotalAmount.Get() == nil {
		var ret float32
		return ret
	}
	return *o.TotalAmount.Get()
}

// GetTotalAmountOk returns a tuple with the TotalAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvoiceLineItem) GetTotalAmountOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TotalAmount.Get(), o.TotalAmount.IsSet()
}

// HasTotalAmount returns a boolean if a field has been set.
func (o *InvoiceLineItem) HasTotalAmount() bool {
	if o != nil && o.TotalAmount.IsSet() {
		return true
	}

	return false
}

// SetTotalAmount gets a reference to the given NullableFloat32 and assigns it to the TotalAmount field.
func (o *InvoiceLineItem) SetTotalAmount(v float32) {
	o.TotalAmount.Set(&v)
}
// SetTotalAmountNil sets the value for TotalAmount to be an explicit nil
func (o *InvoiceLineItem) SetTotalAmountNil() {
	o.TotalAmount.Set(nil)
}

// UnsetTotalAmount ensures that no value is present for TotalAmount, not even an explicit nil
func (o *InvoiceLineItem) UnsetTotalAmount() {
	o.TotalAmount.Unset()
}

// GetCurrency returns the Currency field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InvoiceLineItem) GetCurrency() CurrencyEnum {
	if o == nil || o.Currency.Get() == nil {
		var ret CurrencyEnum
		return ret
	}
	return *o.Currency.Get()
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvoiceLineItem) GetCurrencyOk() (*CurrencyEnum, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Currency.Get(), o.Currency.IsSet()
}

// HasCurrency returns a boolean if a field has been set.
func (o *InvoiceLineItem) HasCurrency() bool {
	if o != nil && o.Currency.IsSet() {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given NullableCurrencyEnum and assigns it to the Currency field.
func (o *InvoiceLineItem) SetCurrency(v CurrencyEnum) {
	o.Currency.Set(&v)
}
// SetCurrencyNil sets the value for Currency to be an explicit nil
func (o *InvoiceLineItem) SetCurrencyNil() {
	o.Currency.Set(nil)
}

// UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
func (o *InvoiceLineItem) UnsetCurrency() {
	o.Currency.Unset()
}

// GetExchangeRate returns the ExchangeRate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InvoiceLineItem) GetExchangeRate() float64 {
	if o == nil || o.ExchangeRate.Get() == nil {
		var ret float64
		return ret
	}
	return *o.ExchangeRate.Get()
}

// GetExchangeRateOk returns a tuple with the ExchangeRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvoiceLineItem) GetExchangeRateOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ExchangeRate.Get(), o.ExchangeRate.IsSet()
}

// HasExchangeRate returns a boolean if a field has been set.
func (o *InvoiceLineItem) HasExchangeRate() bool {
	if o != nil && o.ExchangeRate.IsSet() {
		return true
	}

	return false
}

// SetExchangeRate gets a reference to the given NullableFloat64 and assigns it to the ExchangeRate field.
func (o *InvoiceLineItem) SetExchangeRate(v float64) {
	o.ExchangeRate.Set(&v)
}
// SetExchangeRateNil sets the value for ExchangeRate to be an explicit nil
func (o *InvoiceLineItem) SetExchangeRateNil() {
	o.ExchangeRate.Set(nil)
}

// UnsetExchangeRate ensures that no value is present for ExchangeRate, not even an explicit nil
func (o *InvoiceLineItem) UnsetExchangeRate() {
	o.ExchangeRate.Unset()
}

// GetItem returns the Item field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InvoiceLineItem) GetItem() string {
	if o == nil || o.Item.Get() == nil {
		var ret string
		return ret
	}
	return *o.Item.Get()
}

// GetItemOk returns a tuple with the Item field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvoiceLineItem) GetItemOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Item.Get(), o.Item.IsSet()
}

// HasItem returns a boolean if a field has been set.
func (o *InvoiceLineItem) HasItem() bool {
	if o != nil && o.Item.IsSet() {
		return true
	}

	return false
}

// SetItem gets a reference to the given NullableString and assigns it to the Item field.
func (o *InvoiceLineItem) SetItem(v string) {
	o.Item.Set(&v)
}
// SetItemNil sets the value for Item to be an explicit nil
func (o *InvoiceLineItem) SetItemNil() {
	o.Item.Set(nil)
}

// UnsetItem ensures that no value is present for Item, not even an explicit nil
func (o *InvoiceLineItem) UnsetItem() {
	o.Item.Unset()
}

// GetAccount returns the Account field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InvoiceLineItem) GetAccount() string {
	if o == nil || o.Account.Get() == nil {
		var ret string
		return ret
	}
	return *o.Account.Get()
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvoiceLineItem) GetAccountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Account.Get(), o.Account.IsSet()
}

// HasAccount returns a boolean if a field has been set.
func (o *InvoiceLineItem) HasAccount() bool {
	if o != nil && o.Account.IsSet() {
		return true
	}

	return false
}

// SetAccount gets a reference to the given NullableString and assigns it to the Account field.
func (o *InvoiceLineItem) SetAccount(v string) {
	o.Account.Set(&v)
}
// SetAccountNil sets the value for Account to be an explicit nil
func (o *InvoiceLineItem) SetAccountNil() {
	o.Account.Set(nil)
}

// UnsetAccount ensures that no value is present for Account, not even an explicit nil
func (o *InvoiceLineItem) UnsetAccount() {
	o.Account.Unset()
}

// GetTrackingCategory returns the TrackingCategory field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InvoiceLineItem) GetTrackingCategory() string {
	if o == nil || o.TrackingCategory.Get() == nil {
		var ret string
		return ret
	}
	return *o.TrackingCategory.Get()
}

// GetTrackingCategoryOk returns a tuple with the TrackingCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvoiceLineItem) GetTrackingCategoryOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TrackingCategory.Get(), o.TrackingCategory.IsSet()
}

// HasTrackingCategory returns a boolean if a field has been set.
func (o *InvoiceLineItem) HasTrackingCategory() bool {
	if o != nil && o.TrackingCategory.IsSet() {
		return true
	}

	return false
}

// SetTrackingCategory gets a reference to the given NullableString and assigns it to the TrackingCategory field.
func (o *InvoiceLineItem) SetTrackingCategory(v string) {
	o.TrackingCategory.Set(&v)
}
// SetTrackingCategoryNil sets the value for TrackingCategory to be an explicit nil
func (o *InvoiceLineItem) SetTrackingCategoryNil() {
	o.TrackingCategory.Set(nil)
}

// UnsetTrackingCategory ensures that no value is present for TrackingCategory, not even an explicit nil
func (o *InvoiceLineItem) UnsetTrackingCategory() {
	o.TrackingCategory.Unset()
}

// GetTrackingCategories returns the TrackingCategories field value if set, zero value otherwise.
func (o *InvoiceLineItem) GetTrackingCategories() []string {
	if o == nil || o.TrackingCategories == nil {
		var ret []string
		return ret
	}
	return *o.TrackingCategories
}

// GetTrackingCategoriesOk returns a tuple with the TrackingCategories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceLineItem) GetTrackingCategoriesOk() (*[]string, bool) {
	if o == nil || o.TrackingCategories == nil {
		return nil, false
	}
	return o.TrackingCategories, true
}

// HasTrackingCategories returns a boolean if a field has been set.
func (o *InvoiceLineItem) HasTrackingCategories() bool {
	if o != nil && o.TrackingCategories != nil {
		return true
	}

	return false
}

// SetTrackingCategories gets a reference to the given []string and assigns it to the TrackingCategories field.
func (o *InvoiceLineItem) SetTrackingCategories(v []string) {
	o.TrackingCategories = &v
}

// GetCompany returns the Company field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InvoiceLineItem) GetCompany() string {
	if o == nil || o.Company.Get() == nil {
		var ret string
		return ret
	}
	return *o.Company.Get()
}

// GetCompanyOk returns a tuple with the Company field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvoiceLineItem) GetCompanyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Company.Get(), o.Company.IsSet()
}

// HasCompany returns a boolean if a field has been set.
func (o *InvoiceLineItem) HasCompany() bool {
	if o != nil && o.Company.IsSet() {
		return true
	}

	return false
}

// SetCompany gets a reference to the given NullableString and assigns it to the Company field.
func (o *InvoiceLineItem) SetCompany(v string) {
	o.Company.Set(&v)
}
// SetCompanyNil sets the value for Company to be an explicit nil
func (o *InvoiceLineItem) SetCompanyNil() {
	o.Company.Set(nil)
}

// UnsetCompany ensures that no value is present for Company, not even an explicit nil
func (o *InvoiceLineItem) UnsetCompany() {
	o.Company.Unset()
}

// GetFieldMappings returns the FieldMappings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InvoiceLineItem) GetFieldMappings() map[string]interface{} {
	if o == nil  {
		var ret map[string]interface{}
		return ret
	}
	return o.FieldMappings
}

// GetFieldMappingsOk returns a tuple with the FieldMappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvoiceLineItem) GetFieldMappingsOk() (*map[string]interface{}, bool) {
	if o == nil || o.FieldMappings == nil {
		return nil, false
	}
	return &o.FieldMappings, true
}

// HasFieldMappings returns a boolean if a field has been set.
func (o *InvoiceLineItem) HasFieldMappings() bool {
	if o != nil && o.FieldMappings != nil {
		return true
	}

	return false
}

// SetFieldMappings gets a reference to the given map[string]interface{} and assigns it to the FieldMappings field.
func (o *InvoiceLineItem) SetFieldMappings(v map[string]interface{}) {
	o.FieldMappings = v
}

func (o InvoiceLineItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.RemoteId.IsSet() {
		toSerialize["remote_id"] = o.RemoteId.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.UnitPrice.IsSet() {
		toSerialize["unit_price"] = o.UnitPrice.Get()
	}
	if o.Quantity.IsSet() {
		toSerialize["quantity"] = o.Quantity.Get()
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
	if o.Item.IsSet() {
		toSerialize["item"] = o.Item.Get()
	}
	if o.Account.IsSet() {
		toSerialize["account"] = o.Account.Get()
	}
	if o.TrackingCategory.IsSet() {
		toSerialize["tracking_category"] = o.TrackingCategory.Get()
	}
	if o.TrackingCategories != nil {
		toSerialize["tracking_categories"] = o.TrackingCategories
	}
	if o.Company.IsSet() {
		toSerialize["company"] = o.Company.Get()
	}
	if o.FieldMappings != nil {
		toSerialize["field_mappings"] = o.FieldMappings
	}
	return json.Marshal(toSerialize)
}

func (v *InvoiceLineItem) UnmarshalJSON(src []byte) error {
    type InvoiceLineItemUnmarshalTarget InvoiceLineItem

	var intermediateResult InvoiceLineItemUnmarshalTarget
	var err1 = json.Unmarshal(src, &intermediateResult)
    if err1 != nil {
        return err1
    }
    var err2 = json.Unmarshal(src, &intermediateResult.ResponseRaw)
	if err2 != nil {
		return err2
	}

	*v = InvoiceLineItem(intermediateResult)
	return nil
}
type NullableInvoiceLineItem struct {
	value *InvoiceLineItem
	isSet bool
}

func (v NullableInvoiceLineItem) Get() *InvoiceLineItem {
	return v.value
}

func (v *NullableInvoiceLineItem) Set(val *InvoiceLineItem) {
	v.value = val
	v.isSet = true
}

func (v NullableInvoiceLineItem) IsSet() bool {
	return v.isSet
}

func (v *NullableInvoiceLineItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvoiceLineItem(val *InvoiceLineItem) *NullableInvoiceLineItem {
	return &NullableInvoiceLineItem{value: val, isSet: true}
}

func (v NullableInvoiceLineItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvoiceLineItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	var err1 = json.Unmarshal(src, &v.value)
    if err1 != nil {
        return err1
    }
    return json.Unmarshal(src, &v.value.ResponseRaw)
}


