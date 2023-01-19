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

// PurchaseOrderLineItem # The PurchaseOrderLineItem Object ### Description The `PurchaseOrderLineItem` object is used to represent a purchase order's line item.  ### Usage Example Fetch from the `GET PurchaseOrder` endpoint and view a company's purchase orders.
type PurchaseOrderLineItem struct {
	// A description of the good being purchased.
	Description NullableString `json:"description,omitempty"`
	// The line item's unit price.
	UnitPrice NullableFloat32 `json:"unit_price,omitempty"`
	// The line item's quantity.
	Quantity NullableFloat32 `json:"quantity,omitempty"`
	Item NullableString `json:"item,omitempty"`
	// The purchase order line item's account.
	Account NullableString `json:"account,omitempty"`
	// The purchase order line item's associated tracking category.
	TrackingCategory NullableString `json:"tracking_category,omitempty"`
	// The purchase order line item's associated tracking categories.
	TrackingCategories []string `json:"tracking_categories"`
	// The purchase order line item's tax amount.
	TaxAmount NullableFloat64 `json:"tax_amount,omitempty"`
	// The purchase order line item's total amount.
	TotalLineAmount NullableFloat64 `json:"total_line_amount,omitempty"`
	// The purchase order line item's currency.
	Currency NullableCurrencyEnum `json:"currency,omitempty"`
	// The purchase order line item's exchange rate.
	ExchangeRate NullableFloat64 `json:"exchange_rate,omitempty"`
	// The company the purchase order line item belongs to.
	Company NullableString `json:"company,omitempty"`
	// raw json response by property name
	ResponseRaw map[string]json.RawMessage `json:"-"`
}

// NewPurchaseOrderLineItem instantiates a new PurchaseOrderLineItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPurchaseOrderLineItem(trackingCategories []string) *PurchaseOrderLineItem {
	this := PurchaseOrderLineItem{}
	this.TrackingCategories = trackingCategories
	return &this
}

// NewPurchaseOrderLineItemWithDefaults instantiates a new PurchaseOrderLineItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPurchaseOrderLineItemWithDefaults() *PurchaseOrderLineItem {
	this := PurchaseOrderLineItem{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PurchaseOrderLineItem) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PurchaseOrderLineItem) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *PurchaseOrderLineItem) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *PurchaseOrderLineItem) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *PurchaseOrderLineItem) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *PurchaseOrderLineItem) UnsetDescription() {
	o.Description.Unset()
}

// GetUnitPrice returns the UnitPrice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PurchaseOrderLineItem) GetUnitPrice() float32 {
	if o == nil || o.UnitPrice.Get() == nil {
		var ret float32
		return ret
	}
	return *o.UnitPrice.Get()
}

// GetUnitPriceOk returns a tuple with the UnitPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PurchaseOrderLineItem) GetUnitPriceOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UnitPrice.Get(), o.UnitPrice.IsSet()
}

// HasUnitPrice returns a boolean if a field has been set.
func (o *PurchaseOrderLineItem) HasUnitPrice() bool {
	if o != nil && o.UnitPrice.IsSet() {
		return true
	}

	return false
}

// SetUnitPrice gets a reference to the given NullableFloat32 and assigns it to the UnitPrice field.
func (o *PurchaseOrderLineItem) SetUnitPrice(v float32) {
	o.UnitPrice.Set(&v)
}
// SetUnitPriceNil sets the value for UnitPrice to be an explicit nil
func (o *PurchaseOrderLineItem) SetUnitPriceNil() {
	o.UnitPrice.Set(nil)
}

// UnsetUnitPrice ensures that no value is present for UnitPrice, not even an explicit nil
func (o *PurchaseOrderLineItem) UnsetUnitPrice() {
	o.UnitPrice.Unset()
}

// GetQuantity returns the Quantity field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PurchaseOrderLineItem) GetQuantity() float32 {
	if o == nil || o.Quantity.Get() == nil {
		var ret float32
		return ret
	}
	return *o.Quantity.Get()
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PurchaseOrderLineItem) GetQuantityOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Quantity.Get(), o.Quantity.IsSet()
}

// HasQuantity returns a boolean if a field has been set.
func (o *PurchaseOrderLineItem) HasQuantity() bool {
	if o != nil && o.Quantity.IsSet() {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given NullableFloat32 and assigns it to the Quantity field.
func (o *PurchaseOrderLineItem) SetQuantity(v float32) {
	o.Quantity.Set(&v)
}
// SetQuantityNil sets the value for Quantity to be an explicit nil
func (o *PurchaseOrderLineItem) SetQuantityNil() {
	o.Quantity.Set(nil)
}

// UnsetQuantity ensures that no value is present for Quantity, not even an explicit nil
func (o *PurchaseOrderLineItem) UnsetQuantity() {
	o.Quantity.Unset()
}

// GetItem returns the Item field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PurchaseOrderLineItem) GetItem() string {
	if o == nil || o.Item.Get() == nil {
		var ret string
		return ret
	}
	return *o.Item.Get()
}

// GetItemOk returns a tuple with the Item field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PurchaseOrderLineItem) GetItemOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Item.Get(), o.Item.IsSet()
}

// HasItem returns a boolean if a field has been set.
func (o *PurchaseOrderLineItem) HasItem() bool {
	if o != nil && o.Item.IsSet() {
		return true
	}

	return false
}

// SetItem gets a reference to the given NullableString and assigns it to the Item field.
func (o *PurchaseOrderLineItem) SetItem(v string) {
	o.Item.Set(&v)
}
// SetItemNil sets the value for Item to be an explicit nil
func (o *PurchaseOrderLineItem) SetItemNil() {
	o.Item.Set(nil)
}

// UnsetItem ensures that no value is present for Item, not even an explicit nil
func (o *PurchaseOrderLineItem) UnsetItem() {
	o.Item.Unset()
}

// GetAccount returns the Account field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PurchaseOrderLineItem) GetAccount() string {
	if o == nil || o.Account.Get() == nil {
		var ret string
		return ret
	}
	return *o.Account.Get()
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PurchaseOrderLineItem) GetAccountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Account.Get(), o.Account.IsSet()
}

// HasAccount returns a boolean if a field has been set.
func (o *PurchaseOrderLineItem) HasAccount() bool {
	if o != nil && o.Account.IsSet() {
		return true
	}

	return false
}

// SetAccount gets a reference to the given NullableString and assigns it to the Account field.
func (o *PurchaseOrderLineItem) SetAccount(v string) {
	o.Account.Set(&v)
}
// SetAccountNil sets the value for Account to be an explicit nil
func (o *PurchaseOrderLineItem) SetAccountNil() {
	o.Account.Set(nil)
}

// UnsetAccount ensures that no value is present for Account, not even an explicit nil
func (o *PurchaseOrderLineItem) UnsetAccount() {
	o.Account.Unset()
}

// GetTrackingCategory returns the TrackingCategory field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PurchaseOrderLineItem) GetTrackingCategory() string {
	if o == nil || o.TrackingCategory.Get() == nil {
		var ret string
		return ret
	}
	return *o.TrackingCategory.Get()
}

// GetTrackingCategoryOk returns a tuple with the TrackingCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PurchaseOrderLineItem) GetTrackingCategoryOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TrackingCategory.Get(), o.TrackingCategory.IsSet()
}

// HasTrackingCategory returns a boolean if a field has been set.
func (o *PurchaseOrderLineItem) HasTrackingCategory() bool {
	if o != nil && o.TrackingCategory.IsSet() {
		return true
	}

	return false
}

// SetTrackingCategory gets a reference to the given NullableString and assigns it to the TrackingCategory field.
func (o *PurchaseOrderLineItem) SetTrackingCategory(v string) {
	o.TrackingCategory.Set(&v)
}
// SetTrackingCategoryNil sets the value for TrackingCategory to be an explicit nil
func (o *PurchaseOrderLineItem) SetTrackingCategoryNil() {
	o.TrackingCategory.Set(nil)
}

// UnsetTrackingCategory ensures that no value is present for TrackingCategory, not even an explicit nil
func (o *PurchaseOrderLineItem) UnsetTrackingCategory() {
	o.TrackingCategory.Unset()
}

// GetTrackingCategories returns the TrackingCategories field value
func (o *PurchaseOrderLineItem) GetTrackingCategories() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.TrackingCategories
}

// GetTrackingCategoriesOk returns a tuple with the TrackingCategories field value
// and a boolean to check if the value has been set.
func (o *PurchaseOrderLineItem) GetTrackingCategoriesOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TrackingCategories, true
}

// SetTrackingCategories sets field value
func (o *PurchaseOrderLineItem) SetTrackingCategories(v []string) {
	o.TrackingCategories = v
}

// GetTaxAmount returns the TaxAmount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PurchaseOrderLineItem) GetTaxAmount() float64 {
	if o == nil || o.TaxAmount.Get() == nil {
		var ret float64
		return ret
	}
	return *o.TaxAmount.Get()
}

// GetTaxAmountOk returns a tuple with the TaxAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PurchaseOrderLineItem) GetTaxAmountOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TaxAmount.Get(), o.TaxAmount.IsSet()
}

// HasTaxAmount returns a boolean if a field has been set.
func (o *PurchaseOrderLineItem) HasTaxAmount() bool {
	if o != nil && o.TaxAmount.IsSet() {
		return true
	}

	return false
}

// SetTaxAmount gets a reference to the given NullableFloat64 and assigns it to the TaxAmount field.
func (o *PurchaseOrderLineItem) SetTaxAmount(v float64) {
	o.TaxAmount.Set(&v)
}
// SetTaxAmountNil sets the value for TaxAmount to be an explicit nil
func (o *PurchaseOrderLineItem) SetTaxAmountNil() {
	o.TaxAmount.Set(nil)
}

// UnsetTaxAmount ensures that no value is present for TaxAmount, not even an explicit nil
func (o *PurchaseOrderLineItem) UnsetTaxAmount() {
	o.TaxAmount.Unset()
}

// GetTotalLineAmount returns the TotalLineAmount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PurchaseOrderLineItem) GetTotalLineAmount() float64 {
	if o == nil || o.TotalLineAmount.Get() == nil {
		var ret float64
		return ret
	}
	return *o.TotalLineAmount.Get()
}

// GetTotalLineAmountOk returns a tuple with the TotalLineAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PurchaseOrderLineItem) GetTotalLineAmountOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TotalLineAmount.Get(), o.TotalLineAmount.IsSet()
}

// HasTotalLineAmount returns a boolean if a field has been set.
func (o *PurchaseOrderLineItem) HasTotalLineAmount() bool {
	if o != nil && o.TotalLineAmount.IsSet() {
		return true
	}

	return false
}

// SetTotalLineAmount gets a reference to the given NullableFloat64 and assigns it to the TotalLineAmount field.
func (o *PurchaseOrderLineItem) SetTotalLineAmount(v float64) {
	o.TotalLineAmount.Set(&v)
}
// SetTotalLineAmountNil sets the value for TotalLineAmount to be an explicit nil
func (o *PurchaseOrderLineItem) SetTotalLineAmountNil() {
	o.TotalLineAmount.Set(nil)
}

// UnsetTotalLineAmount ensures that no value is present for TotalLineAmount, not even an explicit nil
func (o *PurchaseOrderLineItem) UnsetTotalLineAmount() {
	o.TotalLineAmount.Unset()
}

// GetCurrency returns the Currency field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PurchaseOrderLineItem) GetCurrency() CurrencyEnum {
	if o == nil || o.Currency.Get() == nil {
		var ret CurrencyEnum
		return ret
	}
	return *o.Currency.Get()
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PurchaseOrderLineItem) GetCurrencyOk() (*CurrencyEnum, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Currency.Get(), o.Currency.IsSet()
}

// HasCurrency returns a boolean if a field has been set.
func (o *PurchaseOrderLineItem) HasCurrency() bool {
	if o != nil && o.Currency.IsSet() {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given NullableCurrencyEnum and assigns it to the Currency field.
func (o *PurchaseOrderLineItem) SetCurrency(v CurrencyEnum) {
	o.Currency.Set(&v)
}
// SetCurrencyNil sets the value for Currency to be an explicit nil
func (o *PurchaseOrderLineItem) SetCurrencyNil() {
	o.Currency.Set(nil)
}

// UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
func (o *PurchaseOrderLineItem) UnsetCurrency() {
	o.Currency.Unset()
}

// GetExchangeRate returns the ExchangeRate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PurchaseOrderLineItem) GetExchangeRate() float64 {
	if o == nil || o.ExchangeRate.Get() == nil {
		var ret float64
		return ret
	}
	return *o.ExchangeRate.Get()
}

// GetExchangeRateOk returns a tuple with the ExchangeRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PurchaseOrderLineItem) GetExchangeRateOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ExchangeRate.Get(), o.ExchangeRate.IsSet()
}

// HasExchangeRate returns a boolean if a field has been set.
func (o *PurchaseOrderLineItem) HasExchangeRate() bool {
	if o != nil && o.ExchangeRate.IsSet() {
		return true
	}

	return false
}

// SetExchangeRate gets a reference to the given NullableFloat64 and assigns it to the ExchangeRate field.
func (o *PurchaseOrderLineItem) SetExchangeRate(v float64) {
	o.ExchangeRate.Set(&v)
}
// SetExchangeRateNil sets the value for ExchangeRate to be an explicit nil
func (o *PurchaseOrderLineItem) SetExchangeRateNil() {
	o.ExchangeRate.Set(nil)
}

// UnsetExchangeRate ensures that no value is present for ExchangeRate, not even an explicit nil
func (o *PurchaseOrderLineItem) UnsetExchangeRate() {
	o.ExchangeRate.Unset()
}

// GetCompany returns the Company field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PurchaseOrderLineItem) GetCompany() string {
	if o == nil || o.Company.Get() == nil {
		var ret string
		return ret
	}
	return *o.Company.Get()
}

// GetCompanyOk returns a tuple with the Company field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PurchaseOrderLineItem) GetCompanyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Company.Get(), o.Company.IsSet()
}

// HasCompany returns a boolean if a field has been set.
func (o *PurchaseOrderLineItem) HasCompany() bool {
	if o != nil && o.Company.IsSet() {
		return true
	}

	return false
}

// SetCompany gets a reference to the given NullableString and assigns it to the Company field.
func (o *PurchaseOrderLineItem) SetCompany(v string) {
	o.Company.Set(&v)
}
// SetCompanyNil sets the value for Company to be an explicit nil
func (o *PurchaseOrderLineItem) SetCompanyNil() {
	o.Company.Set(nil)
}

// UnsetCompany ensures that no value is present for Company, not even an explicit nil
func (o *PurchaseOrderLineItem) UnsetCompany() {
	o.Company.Unset()
}

func (o PurchaseOrderLineItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.UnitPrice.IsSet() {
		toSerialize["unit_price"] = o.UnitPrice.Get()
	}
	if o.Quantity.IsSet() {
		toSerialize["quantity"] = o.Quantity.Get()
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
	if true {
		toSerialize["tracking_categories"] = o.TrackingCategories
	}
	if o.TaxAmount.IsSet() {
		toSerialize["tax_amount"] = o.TaxAmount.Get()
	}
	if o.TotalLineAmount.IsSet() {
		toSerialize["total_line_amount"] = o.TotalLineAmount.Get()
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
	return json.Marshal(toSerialize)
}

func (v *PurchaseOrderLineItem) UnmarshalJSON(src []byte) error {
    type PurchaseOrderLineItemUnmarshalTarget PurchaseOrderLineItem

	var intermediateResult PurchaseOrderLineItemUnmarshalTarget
	var err1 = json.Unmarshal(src, &intermediateResult)
    if err1 != nil {
        return err1
    }
    var err2 = json.Unmarshal(src, &intermediateResult.ResponseRaw)
	if err2 != nil {
		return err2
	}

	*v = PurchaseOrderLineItem(intermediateResult)
	return nil
}
type NullablePurchaseOrderLineItem struct {
	value *PurchaseOrderLineItem
	isSet bool
}

func (v NullablePurchaseOrderLineItem) Get() *PurchaseOrderLineItem {
	return v.value
}

func (v *NullablePurchaseOrderLineItem) Set(val *PurchaseOrderLineItem) {
	v.value = val
	v.isSet = true
}

func (v NullablePurchaseOrderLineItem) IsSet() bool {
	return v.isSet
}

func (v *NullablePurchaseOrderLineItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePurchaseOrderLineItem(val *PurchaseOrderLineItem) *NullablePurchaseOrderLineItem {
	return &NullablePurchaseOrderLineItem{value: val, isSet: true}
}

func (v NullablePurchaseOrderLineItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePurchaseOrderLineItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	var err1 = json.Unmarshal(src, &v.value)
    if err1 != nil {
        return err1
    }
    return json.Unmarshal(src, &v.value.ResponseRaw)
}

