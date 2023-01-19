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

// PurchaseOrder # The PurchaseOrder Object ### Description The `PurchaseOrder` object is a record of request for a product or service between a buyer and seller.  ### Usage Example Fetch from the `LIST PurchaseOrders` endpoint and view a company's purchase orders.
type PurchaseOrder struct {
	Id *string `json:"id,omitempty"`
	// The third-party API ID of the matching object.
	RemoteId NullableString `json:"remote_id,omitempty"`
	RemoteData []RemoteData `json:"remote_data,omitempty"`
	// The purchase order's status.
	Status NullablePurchaseOrderStatusEnum `json:"status,omitempty"`
	// The purchase order's issue date.
	IssueDate NullableTime `json:"issue_date,omitempty"`
	// The purchase order's delivery date.
	DeliveryDate NullableTime `json:"delivery_date,omitempty"`
	// The purchase order's delivery address.
	DeliveryAddress NullableString `json:"delivery_address,omitempty"`
	// The contact making the purchase order.
	Customer NullableString `json:"customer,omitempty"`
	// The party fulfilling the purchase order.
	Vendor NullableString `json:"vendor,omitempty"`
	// A memo attached to the purchase order.
	Memo NullableString `json:"memo,omitempty"`
	// The purchase order's total amount.
	TotalAmount NullableFloat32 `json:"total_amount,omitempty"`
	// The purchase order's currency.
	Currency NullableCurrencyEnum `json:"currency,omitempty"`
	// The purchase order's exchange rate.
	ExchangeRate NullableFloat64 `json:"exchange_rate,omitempty"`
	LineItems *[]PurchaseOrderLineItem `json:"line_items,omitempty"`
	// When the third party's purchase order note was created.
	RemoteCreatedAt NullableTime `json:"remote_created_at,omitempty"`
	// When the third party's purchase order note was updated.
	RemoteUpdatedAt NullableTime `json:"remote_updated_at,omitempty"`
	// Indicates whether or not this object has been deleted by third party webhooks.
	RemoteWasDeleted *bool `json:"remote_was_deleted,omitempty"`
	FieldMappings map[string]interface{} `json:"field_mappings,omitempty"`
	// raw json response by property name
	ResponseRaw map[string]json.RawMessage `json:"-"`
}

// NewPurchaseOrder instantiates a new PurchaseOrder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPurchaseOrder() *PurchaseOrder {
	this := PurchaseOrder{}
	return &this
}

// NewPurchaseOrderWithDefaults instantiates a new PurchaseOrder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPurchaseOrderWithDefaults() *PurchaseOrder {
	this := PurchaseOrder{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PurchaseOrder) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PurchaseOrder) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PurchaseOrder) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PurchaseOrder) SetId(v string) {
	o.Id = &v
}

// GetRemoteId returns the RemoteId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PurchaseOrder) GetRemoteId() string {
	if o == nil || o.RemoteId.Get() == nil {
		var ret string
		return ret
	}
	return *o.RemoteId.Get()
}

// GetRemoteIdOk returns a tuple with the RemoteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PurchaseOrder) GetRemoteIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RemoteId.Get(), o.RemoteId.IsSet()
}

// HasRemoteId returns a boolean if a field has been set.
func (o *PurchaseOrder) HasRemoteId() bool {
	if o != nil && o.RemoteId.IsSet() {
		return true
	}

	return false
}

// SetRemoteId gets a reference to the given NullableString and assigns it to the RemoteId field.
func (o *PurchaseOrder) SetRemoteId(v string) {
	o.RemoteId.Set(&v)
}
// SetRemoteIdNil sets the value for RemoteId to be an explicit nil
func (o *PurchaseOrder) SetRemoteIdNil() {
	o.RemoteId.Set(nil)
}

// UnsetRemoteId ensures that no value is present for RemoteId, not even an explicit nil
func (o *PurchaseOrder) UnsetRemoteId() {
	o.RemoteId.Unset()
}

// GetRemoteData returns the RemoteData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PurchaseOrder) GetRemoteData() []RemoteData {
	if o == nil  {
		var ret []RemoteData
		return ret
	}
	return o.RemoteData
}

// GetRemoteDataOk returns a tuple with the RemoteData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PurchaseOrder) GetRemoteDataOk() (*[]RemoteData, bool) {
	if o == nil || o.RemoteData == nil {
		return nil, false
	}
	return &o.RemoteData, true
}

// HasRemoteData returns a boolean if a field has been set.
func (o *PurchaseOrder) HasRemoteData() bool {
	if o != nil && o.RemoteData != nil {
		return true
	}

	return false
}

// SetRemoteData gets a reference to the given []RemoteData and assigns it to the RemoteData field.
func (o *PurchaseOrder) SetRemoteData(v []RemoteData) {
	o.RemoteData = v
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PurchaseOrder) GetStatus() PurchaseOrderStatusEnum {
	if o == nil || o.Status.Get() == nil {
		var ret PurchaseOrderStatusEnum
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PurchaseOrder) GetStatusOk() (*PurchaseOrderStatusEnum, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *PurchaseOrder) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullablePurchaseOrderStatusEnum and assigns it to the Status field.
func (o *PurchaseOrder) SetStatus(v PurchaseOrderStatusEnum) {
	o.Status.Set(&v)
}
// SetStatusNil sets the value for Status to be an explicit nil
func (o *PurchaseOrder) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *PurchaseOrder) UnsetStatus() {
	o.Status.Unset()
}

// GetIssueDate returns the IssueDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PurchaseOrder) GetIssueDate() time.Time {
	if o == nil || o.IssueDate.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.IssueDate.Get()
}

// GetIssueDateOk returns a tuple with the IssueDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PurchaseOrder) GetIssueDateOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IssueDate.Get(), o.IssueDate.IsSet()
}

// HasIssueDate returns a boolean if a field has been set.
func (o *PurchaseOrder) HasIssueDate() bool {
	if o != nil && o.IssueDate.IsSet() {
		return true
	}

	return false
}

// SetIssueDate gets a reference to the given NullableTime and assigns it to the IssueDate field.
func (o *PurchaseOrder) SetIssueDate(v time.Time) {
	o.IssueDate.Set(&v)
}
// SetIssueDateNil sets the value for IssueDate to be an explicit nil
func (o *PurchaseOrder) SetIssueDateNil() {
	o.IssueDate.Set(nil)
}

// UnsetIssueDate ensures that no value is present for IssueDate, not even an explicit nil
func (o *PurchaseOrder) UnsetIssueDate() {
	o.IssueDate.Unset()
}

// GetDeliveryDate returns the DeliveryDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PurchaseOrder) GetDeliveryDate() time.Time {
	if o == nil || o.DeliveryDate.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.DeliveryDate.Get()
}

// GetDeliveryDateOk returns a tuple with the DeliveryDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PurchaseOrder) GetDeliveryDateOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DeliveryDate.Get(), o.DeliveryDate.IsSet()
}

// HasDeliveryDate returns a boolean if a field has been set.
func (o *PurchaseOrder) HasDeliveryDate() bool {
	if o != nil && o.DeliveryDate.IsSet() {
		return true
	}

	return false
}

// SetDeliveryDate gets a reference to the given NullableTime and assigns it to the DeliveryDate field.
func (o *PurchaseOrder) SetDeliveryDate(v time.Time) {
	o.DeliveryDate.Set(&v)
}
// SetDeliveryDateNil sets the value for DeliveryDate to be an explicit nil
func (o *PurchaseOrder) SetDeliveryDateNil() {
	o.DeliveryDate.Set(nil)
}

// UnsetDeliveryDate ensures that no value is present for DeliveryDate, not even an explicit nil
func (o *PurchaseOrder) UnsetDeliveryDate() {
	o.DeliveryDate.Unset()
}

// GetDeliveryAddress returns the DeliveryAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PurchaseOrder) GetDeliveryAddress() string {
	if o == nil || o.DeliveryAddress.Get() == nil {
		var ret string
		return ret
	}
	return *o.DeliveryAddress.Get()
}

// GetDeliveryAddressOk returns a tuple with the DeliveryAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PurchaseOrder) GetDeliveryAddressOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DeliveryAddress.Get(), o.DeliveryAddress.IsSet()
}

// HasDeliveryAddress returns a boolean if a field has been set.
func (o *PurchaseOrder) HasDeliveryAddress() bool {
	if o != nil && o.DeliveryAddress.IsSet() {
		return true
	}

	return false
}

// SetDeliveryAddress gets a reference to the given NullableString and assigns it to the DeliveryAddress field.
func (o *PurchaseOrder) SetDeliveryAddress(v string) {
	o.DeliveryAddress.Set(&v)
}
// SetDeliveryAddressNil sets the value for DeliveryAddress to be an explicit nil
func (o *PurchaseOrder) SetDeliveryAddressNil() {
	o.DeliveryAddress.Set(nil)
}

// UnsetDeliveryAddress ensures that no value is present for DeliveryAddress, not even an explicit nil
func (o *PurchaseOrder) UnsetDeliveryAddress() {
	o.DeliveryAddress.Unset()
}

// GetCustomer returns the Customer field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PurchaseOrder) GetCustomer() string {
	if o == nil || o.Customer.Get() == nil {
		var ret string
		return ret
	}
	return *o.Customer.Get()
}

// GetCustomerOk returns a tuple with the Customer field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PurchaseOrder) GetCustomerOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Customer.Get(), o.Customer.IsSet()
}

// HasCustomer returns a boolean if a field has been set.
func (o *PurchaseOrder) HasCustomer() bool {
	if o != nil && o.Customer.IsSet() {
		return true
	}

	return false
}

// SetCustomer gets a reference to the given NullableString and assigns it to the Customer field.
func (o *PurchaseOrder) SetCustomer(v string) {
	o.Customer.Set(&v)
}
// SetCustomerNil sets the value for Customer to be an explicit nil
func (o *PurchaseOrder) SetCustomerNil() {
	o.Customer.Set(nil)
}

// UnsetCustomer ensures that no value is present for Customer, not even an explicit nil
func (o *PurchaseOrder) UnsetCustomer() {
	o.Customer.Unset()
}

// GetVendor returns the Vendor field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PurchaseOrder) GetVendor() string {
	if o == nil || o.Vendor.Get() == nil {
		var ret string
		return ret
	}
	return *o.Vendor.Get()
}

// GetVendorOk returns a tuple with the Vendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PurchaseOrder) GetVendorOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Vendor.Get(), o.Vendor.IsSet()
}

// HasVendor returns a boolean if a field has been set.
func (o *PurchaseOrder) HasVendor() bool {
	if o != nil && o.Vendor.IsSet() {
		return true
	}

	return false
}

// SetVendor gets a reference to the given NullableString and assigns it to the Vendor field.
func (o *PurchaseOrder) SetVendor(v string) {
	o.Vendor.Set(&v)
}
// SetVendorNil sets the value for Vendor to be an explicit nil
func (o *PurchaseOrder) SetVendorNil() {
	o.Vendor.Set(nil)
}

// UnsetVendor ensures that no value is present for Vendor, not even an explicit nil
func (o *PurchaseOrder) UnsetVendor() {
	o.Vendor.Unset()
}

// GetMemo returns the Memo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PurchaseOrder) GetMemo() string {
	if o == nil || o.Memo.Get() == nil {
		var ret string
		return ret
	}
	return *o.Memo.Get()
}

// GetMemoOk returns a tuple with the Memo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PurchaseOrder) GetMemoOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Memo.Get(), o.Memo.IsSet()
}

// HasMemo returns a boolean if a field has been set.
func (o *PurchaseOrder) HasMemo() bool {
	if o != nil && o.Memo.IsSet() {
		return true
	}

	return false
}

// SetMemo gets a reference to the given NullableString and assigns it to the Memo field.
func (o *PurchaseOrder) SetMemo(v string) {
	o.Memo.Set(&v)
}
// SetMemoNil sets the value for Memo to be an explicit nil
func (o *PurchaseOrder) SetMemoNil() {
	o.Memo.Set(nil)
}

// UnsetMemo ensures that no value is present for Memo, not even an explicit nil
func (o *PurchaseOrder) UnsetMemo() {
	o.Memo.Unset()
}

// GetTotalAmount returns the TotalAmount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PurchaseOrder) GetTotalAmount() float32 {
	if o == nil || o.TotalAmount.Get() == nil {
		var ret float32
		return ret
	}
	return *o.TotalAmount.Get()
}

// GetTotalAmountOk returns a tuple with the TotalAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PurchaseOrder) GetTotalAmountOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TotalAmount.Get(), o.TotalAmount.IsSet()
}

// HasTotalAmount returns a boolean if a field has been set.
func (o *PurchaseOrder) HasTotalAmount() bool {
	if o != nil && o.TotalAmount.IsSet() {
		return true
	}

	return false
}

// SetTotalAmount gets a reference to the given NullableFloat32 and assigns it to the TotalAmount field.
func (o *PurchaseOrder) SetTotalAmount(v float32) {
	o.TotalAmount.Set(&v)
}
// SetTotalAmountNil sets the value for TotalAmount to be an explicit nil
func (o *PurchaseOrder) SetTotalAmountNil() {
	o.TotalAmount.Set(nil)
}

// UnsetTotalAmount ensures that no value is present for TotalAmount, not even an explicit nil
func (o *PurchaseOrder) UnsetTotalAmount() {
	o.TotalAmount.Unset()
}

// GetCurrency returns the Currency field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PurchaseOrder) GetCurrency() CurrencyEnum {
	if o == nil || o.Currency.Get() == nil {
		var ret CurrencyEnum
		return ret
	}
	return *o.Currency.Get()
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PurchaseOrder) GetCurrencyOk() (*CurrencyEnum, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Currency.Get(), o.Currency.IsSet()
}

// HasCurrency returns a boolean if a field has been set.
func (o *PurchaseOrder) HasCurrency() bool {
	if o != nil && o.Currency.IsSet() {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given NullableCurrencyEnum and assigns it to the Currency field.
func (o *PurchaseOrder) SetCurrency(v CurrencyEnum) {
	o.Currency.Set(&v)
}
// SetCurrencyNil sets the value for Currency to be an explicit nil
func (o *PurchaseOrder) SetCurrencyNil() {
	o.Currency.Set(nil)
}

// UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
func (o *PurchaseOrder) UnsetCurrency() {
	o.Currency.Unset()
}

// GetExchangeRate returns the ExchangeRate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PurchaseOrder) GetExchangeRate() float64 {
	if o == nil || o.ExchangeRate.Get() == nil {
		var ret float64
		return ret
	}
	return *o.ExchangeRate.Get()
}

// GetExchangeRateOk returns a tuple with the ExchangeRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PurchaseOrder) GetExchangeRateOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ExchangeRate.Get(), o.ExchangeRate.IsSet()
}

// HasExchangeRate returns a boolean if a field has been set.
func (o *PurchaseOrder) HasExchangeRate() bool {
	if o != nil && o.ExchangeRate.IsSet() {
		return true
	}

	return false
}

// SetExchangeRate gets a reference to the given NullableFloat64 and assigns it to the ExchangeRate field.
func (o *PurchaseOrder) SetExchangeRate(v float64) {
	o.ExchangeRate.Set(&v)
}
// SetExchangeRateNil sets the value for ExchangeRate to be an explicit nil
func (o *PurchaseOrder) SetExchangeRateNil() {
	o.ExchangeRate.Set(nil)
}

// UnsetExchangeRate ensures that no value is present for ExchangeRate, not even an explicit nil
func (o *PurchaseOrder) UnsetExchangeRate() {
	o.ExchangeRate.Unset()
}

// GetLineItems returns the LineItems field value if set, zero value otherwise.
func (o *PurchaseOrder) GetLineItems() []PurchaseOrderLineItem {
	if o == nil || o.LineItems == nil {
		var ret []PurchaseOrderLineItem
		return ret
	}
	return *o.LineItems
}

// GetLineItemsOk returns a tuple with the LineItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PurchaseOrder) GetLineItemsOk() (*[]PurchaseOrderLineItem, bool) {
	if o == nil || o.LineItems == nil {
		return nil, false
	}
	return o.LineItems, true
}

// HasLineItems returns a boolean if a field has been set.
func (o *PurchaseOrder) HasLineItems() bool {
	if o != nil && o.LineItems != nil {
		return true
	}

	return false
}

// SetLineItems gets a reference to the given []PurchaseOrderLineItem and assigns it to the LineItems field.
func (o *PurchaseOrder) SetLineItems(v []PurchaseOrderLineItem) {
	o.LineItems = &v
}

// GetRemoteCreatedAt returns the RemoteCreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PurchaseOrder) GetRemoteCreatedAt() time.Time {
	if o == nil || o.RemoteCreatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.RemoteCreatedAt.Get()
}

// GetRemoteCreatedAtOk returns a tuple with the RemoteCreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PurchaseOrder) GetRemoteCreatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RemoteCreatedAt.Get(), o.RemoteCreatedAt.IsSet()
}

// HasRemoteCreatedAt returns a boolean if a field has been set.
func (o *PurchaseOrder) HasRemoteCreatedAt() bool {
	if o != nil && o.RemoteCreatedAt.IsSet() {
		return true
	}

	return false
}

// SetRemoteCreatedAt gets a reference to the given NullableTime and assigns it to the RemoteCreatedAt field.
func (o *PurchaseOrder) SetRemoteCreatedAt(v time.Time) {
	o.RemoteCreatedAt.Set(&v)
}
// SetRemoteCreatedAtNil sets the value for RemoteCreatedAt to be an explicit nil
func (o *PurchaseOrder) SetRemoteCreatedAtNil() {
	o.RemoteCreatedAt.Set(nil)
}

// UnsetRemoteCreatedAt ensures that no value is present for RemoteCreatedAt, not even an explicit nil
func (o *PurchaseOrder) UnsetRemoteCreatedAt() {
	o.RemoteCreatedAt.Unset()
}

// GetRemoteUpdatedAt returns the RemoteUpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PurchaseOrder) GetRemoteUpdatedAt() time.Time {
	if o == nil || o.RemoteUpdatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.RemoteUpdatedAt.Get()
}

// GetRemoteUpdatedAtOk returns a tuple with the RemoteUpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PurchaseOrder) GetRemoteUpdatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RemoteUpdatedAt.Get(), o.RemoteUpdatedAt.IsSet()
}

// HasRemoteUpdatedAt returns a boolean if a field has been set.
func (o *PurchaseOrder) HasRemoteUpdatedAt() bool {
	if o != nil && o.RemoteUpdatedAt.IsSet() {
		return true
	}

	return false
}

// SetRemoteUpdatedAt gets a reference to the given NullableTime and assigns it to the RemoteUpdatedAt field.
func (o *PurchaseOrder) SetRemoteUpdatedAt(v time.Time) {
	o.RemoteUpdatedAt.Set(&v)
}
// SetRemoteUpdatedAtNil sets the value for RemoteUpdatedAt to be an explicit nil
func (o *PurchaseOrder) SetRemoteUpdatedAtNil() {
	o.RemoteUpdatedAt.Set(nil)
}

// UnsetRemoteUpdatedAt ensures that no value is present for RemoteUpdatedAt, not even an explicit nil
func (o *PurchaseOrder) UnsetRemoteUpdatedAt() {
	o.RemoteUpdatedAt.Unset()
}

// GetRemoteWasDeleted returns the RemoteWasDeleted field value if set, zero value otherwise.
func (o *PurchaseOrder) GetRemoteWasDeleted() bool {
	if o == nil || o.RemoteWasDeleted == nil {
		var ret bool
		return ret
	}
	return *o.RemoteWasDeleted
}

// GetRemoteWasDeletedOk returns a tuple with the RemoteWasDeleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PurchaseOrder) GetRemoteWasDeletedOk() (*bool, bool) {
	if o == nil || o.RemoteWasDeleted == nil {
		return nil, false
	}
	return o.RemoteWasDeleted, true
}

// HasRemoteWasDeleted returns a boolean if a field has been set.
func (o *PurchaseOrder) HasRemoteWasDeleted() bool {
	if o != nil && o.RemoteWasDeleted != nil {
		return true
	}

	return false
}

// SetRemoteWasDeleted gets a reference to the given bool and assigns it to the RemoteWasDeleted field.
func (o *PurchaseOrder) SetRemoteWasDeleted(v bool) {
	o.RemoteWasDeleted = &v
}

// GetFieldMappings returns the FieldMappings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PurchaseOrder) GetFieldMappings() map[string]interface{} {
	if o == nil  {
		var ret map[string]interface{}
		return ret
	}
	return o.FieldMappings
}

// GetFieldMappingsOk returns a tuple with the FieldMappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PurchaseOrder) GetFieldMappingsOk() (*map[string]interface{}, bool) {
	if o == nil || o.FieldMappings == nil {
		return nil, false
	}
	return &o.FieldMappings, true
}

// HasFieldMappings returns a boolean if a field has been set.
func (o *PurchaseOrder) HasFieldMappings() bool {
	if o != nil && o.FieldMappings != nil {
		return true
	}

	return false
}

// SetFieldMappings gets a reference to the given map[string]interface{} and assigns it to the FieldMappings field.
func (o *PurchaseOrder) SetFieldMappings(v map[string]interface{}) {
	o.FieldMappings = v
}

func (o PurchaseOrder) MarshalJSON() ([]byte, error) {
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
	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
	}
	if o.IssueDate.IsSet() {
		toSerialize["issue_date"] = o.IssueDate.Get()
	}
	if o.DeliveryDate.IsSet() {
		toSerialize["delivery_date"] = o.DeliveryDate.Get()
	}
	if o.DeliveryAddress.IsSet() {
		toSerialize["delivery_address"] = o.DeliveryAddress.Get()
	}
	if o.Customer.IsSet() {
		toSerialize["customer"] = o.Customer.Get()
	}
	if o.Vendor.IsSet() {
		toSerialize["vendor"] = o.Vendor.Get()
	}
	if o.Memo.IsSet() {
		toSerialize["memo"] = o.Memo.Get()
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
	if o.LineItems != nil {
		toSerialize["line_items"] = o.LineItems
	}
	if o.RemoteCreatedAt.IsSet() {
		toSerialize["remote_created_at"] = o.RemoteCreatedAt.Get()
	}
	if o.RemoteUpdatedAt.IsSet() {
		toSerialize["remote_updated_at"] = o.RemoteUpdatedAt.Get()
	}
	if o.RemoteWasDeleted != nil {
		toSerialize["remote_was_deleted"] = o.RemoteWasDeleted
	}
	if o.FieldMappings != nil {
		toSerialize["field_mappings"] = o.FieldMappings
	}
	return json.Marshal(toSerialize)
}

func (v *PurchaseOrder) UnmarshalJSON(src []byte) error {
    type PurchaseOrderUnmarshalTarget PurchaseOrder

	var intermediateResult PurchaseOrderUnmarshalTarget
	var err1 = json.Unmarshal(src, &intermediateResult)
    if err1 != nil {
        return err1
    }
    var err2 = json.Unmarshal(src, &intermediateResult.ResponseRaw)
	if err2 != nil {
		return err2
	}

	*v = PurchaseOrder(intermediateResult)
	return nil
}
type NullablePurchaseOrder struct {
	value *PurchaseOrder
	isSet bool
}

func (v NullablePurchaseOrder) Get() *PurchaseOrder {
	return v.value
}

func (v *NullablePurchaseOrder) Set(val *PurchaseOrder) {
	v.value = val
	v.isSet = true
}

func (v NullablePurchaseOrder) IsSet() bool {
	return v.isSet
}

func (v *NullablePurchaseOrder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePurchaseOrder(val *PurchaseOrder) *NullablePurchaseOrder {
	return &NullablePurchaseOrder{value: val, isSet: true}
}

func (v NullablePurchaseOrder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePurchaseOrder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	var err1 = json.Unmarshal(src, &v.value)
    if err1 != nil {
        return err1
    }
    return json.Unmarshal(src, &v.value.ResponseRaw)
}


