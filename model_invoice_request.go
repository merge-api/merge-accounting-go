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

// InvoiceRequest # The Invoice Object     ### Description     The `Invoice` object represents an itemized record of goods and/or services sold to a customer. If type = accounts_payable `Invoice` is a bill, if type = accounts_receivable it's an invoice.      ### Usage Example     Fetch from the `LIST Invoices` endpoint and view a company's invoices.
type InvoiceRequest struct {
	// The third-party API ID of the matching object.
	RemoteId NullableString `json:"remote_id,omitempty"`
	// Whether the invoice is an accounts receivable or accounts payable. Accounts payable invoices are commonly referred to as Bills.
	Type NullableInvoiceTypeEnum `json:"type,omitempty"`
	// The invoice's contact.
	Contact NullableString `json:"contact,omitempty"`
	// The invoice's number.
	Number NullableString `json:"number,omitempty"`
	// The invoice's issue date.
	IssueDate NullableTime `json:"issue_date,omitempty"`
	// The invoice's due date.
	DueDate NullableTime `json:"due_date,omitempty"`
	// The invoice's paid date.
	PaidOnDate NullableTime `json:"paid_on_date,omitempty"`
	// The invoice's private note.
	Memo NullableString `json:"memo,omitempty"`
	// The company the invoice belongs to.
	Company NullableString `json:"company,omitempty"`
	// The invoice's currency.
	Currency NullableCurrencyEnum `json:"currency,omitempty"`
	// The invoice's exchange rate.
	ExchangeRate NullableFloat64 `json:"exchange_rate,omitempty"`
	// The total discounts applied to the total cost.
	TotalDiscount NullableFloat32 `json:"total_discount,omitempty"`
	// The total amount being paid before taxes.
	SubTotal NullableFloat32 `json:"sub_total,omitempty"`
	// The total amount being paid in taxes.
	TotalTaxAmount NullableFloat32 `json:"total_tax_amount,omitempty"`
	// The invoice's total amount.
	TotalAmount NullableFloat32 `json:"total_amount,omitempty"`
	// The invoice's remaining balance.
	Balance NullableFloat32 `json:"balance,omitempty"`
	// When the third party's invoice entry was updated.
	RemoteUpdatedAt NullableTime `json:"remote_updated_at,omitempty"`
	// Array of `Payment` object IDs.
	Payments *[]string `json:"payments,omitempty"`
	LineItems *[]InvoiceLineItemRequest `json:"line_items,omitempty"`
	IntegrationParams map[string]interface{} `json:"integration_params,omitempty"`
	LinkedAccountParams map[string]interface{} `json:"linked_account_params,omitempty"`
	// raw json response by property name
	ResponseRaw map[string]json.RawMessage `json:"-"`
}

// NewInvoiceRequest instantiates a new InvoiceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvoiceRequest() *InvoiceRequest {
	this := InvoiceRequest{}
	return &this
}

// NewInvoiceRequestWithDefaults instantiates a new InvoiceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvoiceRequestWithDefaults() *InvoiceRequest {
	this := InvoiceRequest{}
	return &this
}

// GetRemoteId returns the RemoteId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InvoiceRequest) GetRemoteId() string {
	if o == nil || o.RemoteId.Get() == nil {
		var ret string
		return ret
	}
	return *o.RemoteId.Get()
}

// GetRemoteIdOk returns a tuple with the RemoteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvoiceRequest) GetRemoteIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RemoteId.Get(), o.RemoteId.IsSet()
}

// HasRemoteId returns a boolean if a field has been set.
func (o *InvoiceRequest) HasRemoteId() bool {
	if o != nil && o.RemoteId.IsSet() {
		return true
	}

	return false
}

// SetRemoteId gets a reference to the given NullableString and assigns it to the RemoteId field.
func (o *InvoiceRequest) SetRemoteId(v string) {
	o.RemoteId.Set(&v)
}
// SetRemoteIdNil sets the value for RemoteId to be an explicit nil
func (o *InvoiceRequest) SetRemoteIdNil() {
	o.RemoteId.Set(nil)
}

// UnsetRemoteId ensures that no value is present for RemoteId, not even an explicit nil
func (o *InvoiceRequest) UnsetRemoteId() {
	o.RemoteId.Unset()
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InvoiceRequest) GetType() InvoiceTypeEnum {
	if o == nil || o.Type.Get() == nil {
		var ret InvoiceTypeEnum
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvoiceRequest) GetTypeOk() (*InvoiceTypeEnum, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *InvoiceRequest) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableInvoiceTypeEnum and assigns it to the Type field.
func (o *InvoiceRequest) SetType(v InvoiceTypeEnum) {
	o.Type.Set(&v)
}
// SetTypeNil sets the value for Type to be an explicit nil
func (o *InvoiceRequest) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *InvoiceRequest) UnsetType() {
	o.Type.Unset()
}

// GetContact returns the Contact field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InvoiceRequest) GetContact() string {
	if o == nil || o.Contact.Get() == nil {
		var ret string
		return ret
	}
	return *o.Contact.Get()
}

// GetContactOk returns a tuple with the Contact field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvoiceRequest) GetContactOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Contact.Get(), o.Contact.IsSet()
}

// HasContact returns a boolean if a field has been set.
func (o *InvoiceRequest) HasContact() bool {
	if o != nil && o.Contact.IsSet() {
		return true
	}

	return false
}

// SetContact gets a reference to the given NullableString and assigns it to the Contact field.
func (o *InvoiceRequest) SetContact(v string) {
	o.Contact.Set(&v)
}
// SetContactNil sets the value for Contact to be an explicit nil
func (o *InvoiceRequest) SetContactNil() {
	o.Contact.Set(nil)
}

// UnsetContact ensures that no value is present for Contact, not even an explicit nil
func (o *InvoiceRequest) UnsetContact() {
	o.Contact.Unset()
}

// GetNumber returns the Number field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InvoiceRequest) GetNumber() string {
	if o == nil || o.Number.Get() == nil {
		var ret string
		return ret
	}
	return *o.Number.Get()
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvoiceRequest) GetNumberOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Number.Get(), o.Number.IsSet()
}

// HasNumber returns a boolean if a field has been set.
func (o *InvoiceRequest) HasNumber() bool {
	if o != nil && o.Number.IsSet() {
		return true
	}

	return false
}

// SetNumber gets a reference to the given NullableString and assigns it to the Number field.
func (o *InvoiceRequest) SetNumber(v string) {
	o.Number.Set(&v)
}
// SetNumberNil sets the value for Number to be an explicit nil
func (o *InvoiceRequest) SetNumberNil() {
	o.Number.Set(nil)
}

// UnsetNumber ensures that no value is present for Number, not even an explicit nil
func (o *InvoiceRequest) UnsetNumber() {
	o.Number.Unset()
}

// GetIssueDate returns the IssueDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InvoiceRequest) GetIssueDate() time.Time {
	if o == nil || o.IssueDate.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.IssueDate.Get()
}

// GetIssueDateOk returns a tuple with the IssueDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvoiceRequest) GetIssueDateOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IssueDate.Get(), o.IssueDate.IsSet()
}

// HasIssueDate returns a boolean if a field has been set.
func (o *InvoiceRequest) HasIssueDate() bool {
	if o != nil && o.IssueDate.IsSet() {
		return true
	}

	return false
}

// SetIssueDate gets a reference to the given NullableTime and assigns it to the IssueDate field.
func (o *InvoiceRequest) SetIssueDate(v time.Time) {
	o.IssueDate.Set(&v)
}
// SetIssueDateNil sets the value for IssueDate to be an explicit nil
func (o *InvoiceRequest) SetIssueDateNil() {
	o.IssueDate.Set(nil)
}

// UnsetIssueDate ensures that no value is present for IssueDate, not even an explicit nil
func (o *InvoiceRequest) UnsetIssueDate() {
	o.IssueDate.Unset()
}

// GetDueDate returns the DueDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InvoiceRequest) GetDueDate() time.Time {
	if o == nil || o.DueDate.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.DueDate.Get()
}

// GetDueDateOk returns a tuple with the DueDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvoiceRequest) GetDueDateOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DueDate.Get(), o.DueDate.IsSet()
}

// HasDueDate returns a boolean if a field has been set.
func (o *InvoiceRequest) HasDueDate() bool {
	if o != nil && o.DueDate.IsSet() {
		return true
	}

	return false
}

// SetDueDate gets a reference to the given NullableTime and assigns it to the DueDate field.
func (o *InvoiceRequest) SetDueDate(v time.Time) {
	o.DueDate.Set(&v)
}
// SetDueDateNil sets the value for DueDate to be an explicit nil
func (o *InvoiceRequest) SetDueDateNil() {
	o.DueDate.Set(nil)
}

// UnsetDueDate ensures that no value is present for DueDate, not even an explicit nil
func (o *InvoiceRequest) UnsetDueDate() {
	o.DueDate.Unset()
}

// GetPaidOnDate returns the PaidOnDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InvoiceRequest) GetPaidOnDate() time.Time {
	if o == nil || o.PaidOnDate.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.PaidOnDate.Get()
}

// GetPaidOnDateOk returns a tuple with the PaidOnDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvoiceRequest) GetPaidOnDateOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PaidOnDate.Get(), o.PaidOnDate.IsSet()
}

// HasPaidOnDate returns a boolean if a field has been set.
func (o *InvoiceRequest) HasPaidOnDate() bool {
	if o != nil && o.PaidOnDate.IsSet() {
		return true
	}

	return false
}

// SetPaidOnDate gets a reference to the given NullableTime and assigns it to the PaidOnDate field.
func (o *InvoiceRequest) SetPaidOnDate(v time.Time) {
	o.PaidOnDate.Set(&v)
}
// SetPaidOnDateNil sets the value for PaidOnDate to be an explicit nil
func (o *InvoiceRequest) SetPaidOnDateNil() {
	o.PaidOnDate.Set(nil)
}

// UnsetPaidOnDate ensures that no value is present for PaidOnDate, not even an explicit nil
func (o *InvoiceRequest) UnsetPaidOnDate() {
	o.PaidOnDate.Unset()
}

// GetMemo returns the Memo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InvoiceRequest) GetMemo() string {
	if o == nil || o.Memo.Get() == nil {
		var ret string
		return ret
	}
	return *o.Memo.Get()
}

// GetMemoOk returns a tuple with the Memo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvoiceRequest) GetMemoOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Memo.Get(), o.Memo.IsSet()
}

// HasMemo returns a boolean if a field has been set.
func (o *InvoiceRequest) HasMemo() bool {
	if o != nil && o.Memo.IsSet() {
		return true
	}

	return false
}

// SetMemo gets a reference to the given NullableString and assigns it to the Memo field.
func (o *InvoiceRequest) SetMemo(v string) {
	o.Memo.Set(&v)
}
// SetMemoNil sets the value for Memo to be an explicit nil
func (o *InvoiceRequest) SetMemoNil() {
	o.Memo.Set(nil)
}

// UnsetMemo ensures that no value is present for Memo, not even an explicit nil
func (o *InvoiceRequest) UnsetMemo() {
	o.Memo.Unset()
}

// GetCompany returns the Company field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InvoiceRequest) GetCompany() string {
	if o == nil || o.Company.Get() == nil {
		var ret string
		return ret
	}
	return *o.Company.Get()
}

// GetCompanyOk returns a tuple with the Company field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvoiceRequest) GetCompanyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Company.Get(), o.Company.IsSet()
}

// HasCompany returns a boolean if a field has been set.
func (o *InvoiceRequest) HasCompany() bool {
	if o != nil && o.Company.IsSet() {
		return true
	}

	return false
}

// SetCompany gets a reference to the given NullableString and assigns it to the Company field.
func (o *InvoiceRequest) SetCompany(v string) {
	o.Company.Set(&v)
}
// SetCompanyNil sets the value for Company to be an explicit nil
func (o *InvoiceRequest) SetCompanyNil() {
	o.Company.Set(nil)
}

// UnsetCompany ensures that no value is present for Company, not even an explicit nil
func (o *InvoiceRequest) UnsetCompany() {
	o.Company.Unset()
}

// GetCurrency returns the Currency field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InvoiceRequest) GetCurrency() CurrencyEnum {
	if o == nil || o.Currency.Get() == nil {
		var ret CurrencyEnum
		return ret
	}
	return *o.Currency.Get()
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvoiceRequest) GetCurrencyOk() (*CurrencyEnum, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Currency.Get(), o.Currency.IsSet()
}

// HasCurrency returns a boolean if a field has been set.
func (o *InvoiceRequest) HasCurrency() bool {
	if o != nil && o.Currency.IsSet() {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given NullableCurrencyEnum and assigns it to the Currency field.
func (o *InvoiceRequest) SetCurrency(v CurrencyEnum) {
	o.Currency.Set(&v)
}
// SetCurrencyNil sets the value for Currency to be an explicit nil
func (o *InvoiceRequest) SetCurrencyNil() {
	o.Currency.Set(nil)
}

// UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
func (o *InvoiceRequest) UnsetCurrency() {
	o.Currency.Unset()
}

// GetExchangeRate returns the ExchangeRate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InvoiceRequest) GetExchangeRate() float64 {
	if o == nil || o.ExchangeRate.Get() == nil {
		var ret float64
		return ret
	}
	return *o.ExchangeRate.Get()
}

// GetExchangeRateOk returns a tuple with the ExchangeRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvoiceRequest) GetExchangeRateOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ExchangeRate.Get(), o.ExchangeRate.IsSet()
}

// HasExchangeRate returns a boolean if a field has been set.
func (o *InvoiceRequest) HasExchangeRate() bool {
	if o != nil && o.ExchangeRate.IsSet() {
		return true
	}

	return false
}

// SetExchangeRate gets a reference to the given NullableFloat64 and assigns it to the ExchangeRate field.
func (o *InvoiceRequest) SetExchangeRate(v float64) {
	o.ExchangeRate.Set(&v)
}
// SetExchangeRateNil sets the value for ExchangeRate to be an explicit nil
func (o *InvoiceRequest) SetExchangeRateNil() {
	o.ExchangeRate.Set(nil)
}

// UnsetExchangeRate ensures that no value is present for ExchangeRate, not even an explicit nil
func (o *InvoiceRequest) UnsetExchangeRate() {
	o.ExchangeRate.Unset()
}

// GetTotalDiscount returns the TotalDiscount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InvoiceRequest) GetTotalDiscount() float32 {
	if o == nil || o.TotalDiscount.Get() == nil {
		var ret float32
		return ret
	}
	return *o.TotalDiscount.Get()
}

// GetTotalDiscountOk returns a tuple with the TotalDiscount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvoiceRequest) GetTotalDiscountOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TotalDiscount.Get(), o.TotalDiscount.IsSet()
}

// HasTotalDiscount returns a boolean if a field has been set.
func (o *InvoiceRequest) HasTotalDiscount() bool {
	if o != nil && o.TotalDiscount.IsSet() {
		return true
	}

	return false
}

// SetTotalDiscount gets a reference to the given NullableFloat32 and assigns it to the TotalDiscount field.
func (o *InvoiceRequest) SetTotalDiscount(v float32) {
	o.TotalDiscount.Set(&v)
}
// SetTotalDiscountNil sets the value for TotalDiscount to be an explicit nil
func (o *InvoiceRequest) SetTotalDiscountNil() {
	o.TotalDiscount.Set(nil)
}

// UnsetTotalDiscount ensures that no value is present for TotalDiscount, not even an explicit nil
func (o *InvoiceRequest) UnsetTotalDiscount() {
	o.TotalDiscount.Unset()
}

// GetSubTotal returns the SubTotal field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InvoiceRequest) GetSubTotal() float32 {
	if o == nil || o.SubTotal.Get() == nil {
		var ret float32
		return ret
	}
	return *o.SubTotal.Get()
}

// GetSubTotalOk returns a tuple with the SubTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvoiceRequest) GetSubTotalOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SubTotal.Get(), o.SubTotal.IsSet()
}

// HasSubTotal returns a boolean if a field has been set.
func (o *InvoiceRequest) HasSubTotal() bool {
	if o != nil && o.SubTotal.IsSet() {
		return true
	}

	return false
}

// SetSubTotal gets a reference to the given NullableFloat32 and assigns it to the SubTotal field.
func (o *InvoiceRequest) SetSubTotal(v float32) {
	o.SubTotal.Set(&v)
}
// SetSubTotalNil sets the value for SubTotal to be an explicit nil
func (o *InvoiceRequest) SetSubTotalNil() {
	o.SubTotal.Set(nil)
}

// UnsetSubTotal ensures that no value is present for SubTotal, not even an explicit nil
func (o *InvoiceRequest) UnsetSubTotal() {
	o.SubTotal.Unset()
}

// GetTotalTaxAmount returns the TotalTaxAmount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InvoiceRequest) GetTotalTaxAmount() float32 {
	if o == nil || o.TotalTaxAmount.Get() == nil {
		var ret float32
		return ret
	}
	return *o.TotalTaxAmount.Get()
}

// GetTotalTaxAmountOk returns a tuple with the TotalTaxAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvoiceRequest) GetTotalTaxAmountOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TotalTaxAmount.Get(), o.TotalTaxAmount.IsSet()
}

// HasTotalTaxAmount returns a boolean if a field has been set.
func (o *InvoiceRequest) HasTotalTaxAmount() bool {
	if o != nil && o.TotalTaxAmount.IsSet() {
		return true
	}

	return false
}

// SetTotalTaxAmount gets a reference to the given NullableFloat32 and assigns it to the TotalTaxAmount field.
func (o *InvoiceRequest) SetTotalTaxAmount(v float32) {
	o.TotalTaxAmount.Set(&v)
}
// SetTotalTaxAmountNil sets the value for TotalTaxAmount to be an explicit nil
func (o *InvoiceRequest) SetTotalTaxAmountNil() {
	o.TotalTaxAmount.Set(nil)
}

// UnsetTotalTaxAmount ensures that no value is present for TotalTaxAmount, not even an explicit nil
func (o *InvoiceRequest) UnsetTotalTaxAmount() {
	o.TotalTaxAmount.Unset()
}

// GetTotalAmount returns the TotalAmount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InvoiceRequest) GetTotalAmount() float32 {
	if o == nil || o.TotalAmount.Get() == nil {
		var ret float32
		return ret
	}
	return *o.TotalAmount.Get()
}

// GetTotalAmountOk returns a tuple with the TotalAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvoiceRequest) GetTotalAmountOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TotalAmount.Get(), o.TotalAmount.IsSet()
}

// HasTotalAmount returns a boolean if a field has been set.
func (o *InvoiceRequest) HasTotalAmount() bool {
	if o != nil && o.TotalAmount.IsSet() {
		return true
	}

	return false
}

// SetTotalAmount gets a reference to the given NullableFloat32 and assigns it to the TotalAmount field.
func (o *InvoiceRequest) SetTotalAmount(v float32) {
	o.TotalAmount.Set(&v)
}
// SetTotalAmountNil sets the value for TotalAmount to be an explicit nil
func (o *InvoiceRequest) SetTotalAmountNil() {
	o.TotalAmount.Set(nil)
}

// UnsetTotalAmount ensures that no value is present for TotalAmount, not even an explicit nil
func (o *InvoiceRequest) UnsetTotalAmount() {
	o.TotalAmount.Unset()
}

// GetBalance returns the Balance field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InvoiceRequest) GetBalance() float32 {
	if o == nil || o.Balance.Get() == nil {
		var ret float32
		return ret
	}
	return *o.Balance.Get()
}

// GetBalanceOk returns a tuple with the Balance field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvoiceRequest) GetBalanceOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Balance.Get(), o.Balance.IsSet()
}

// HasBalance returns a boolean if a field has been set.
func (o *InvoiceRequest) HasBalance() bool {
	if o != nil && o.Balance.IsSet() {
		return true
	}

	return false
}

// SetBalance gets a reference to the given NullableFloat32 and assigns it to the Balance field.
func (o *InvoiceRequest) SetBalance(v float32) {
	o.Balance.Set(&v)
}
// SetBalanceNil sets the value for Balance to be an explicit nil
func (o *InvoiceRequest) SetBalanceNil() {
	o.Balance.Set(nil)
}

// UnsetBalance ensures that no value is present for Balance, not even an explicit nil
func (o *InvoiceRequest) UnsetBalance() {
	o.Balance.Unset()
}

// GetRemoteUpdatedAt returns the RemoteUpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InvoiceRequest) GetRemoteUpdatedAt() time.Time {
	if o == nil || o.RemoteUpdatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.RemoteUpdatedAt.Get()
}

// GetRemoteUpdatedAtOk returns a tuple with the RemoteUpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvoiceRequest) GetRemoteUpdatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RemoteUpdatedAt.Get(), o.RemoteUpdatedAt.IsSet()
}

// HasRemoteUpdatedAt returns a boolean if a field has been set.
func (o *InvoiceRequest) HasRemoteUpdatedAt() bool {
	if o != nil && o.RemoteUpdatedAt.IsSet() {
		return true
	}

	return false
}

// SetRemoteUpdatedAt gets a reference to the given NullableTime and assigns it to the RemoteUpdatedAt field.
func (o *InvoiceRequest) SetRemoteUpdatedAt(v time.Time) {
	o.RemoteUpdatedAt.Set(&v)
}
// SetRemoteUpdatedAtNil sets the value for RemoteUpdatedAt to be an explicit nil
func (o *InvoiceRequest) SetRemoteUpdatedAtNil() {
	o.RemoteUpdatedAt.Set(nil)
}

// UnsetRemoteUpdatedAt ensures that no value is present for RemoteUpdatedAt, not even an explicit nil
func (o *InvoiceRequest) UnsetRemoteUpdatedAt() {
	o.RemoteUpdatedAt.Unset()
}

// GetPayments returns the Payments field value if set, zero value otherwise.
func (o *InvoiceRequest) GetPayments() []string {
	if o == nil || o.Payments == nil {
		var ret []string
		return ret
	}
	return *o.Payments
}

// GetPaymentsOk returns a tuple with the Payments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceRequest) GetPaymentsOk() (*[]string, bool) {
	if o == nil || o.Payments == nil {
		return nil, false
	}
	return o.Payments, true
}

// HasPayments returns a boolean if a field has been set.
func (o *InvoiceRequest) HasPayments() bool {
	if o != nil && o.Payments != nil {
		return true
	}

	return false
}

// SetPayments gets a reference to the given []string and assigns it to the Payments field.
func (o *InvoiceRequest) SetPayments(v []string) {
	o.Payments = &v
}

// GetLineItems returns the LineItems field value if set, zero value otherwise.
func (o *InvoiceRequest) GetLineItems() []InvoiceLineItemRequest {
	if o == nil || o.LineItems == nil {
		var ret []InvoiceLineItemRequest
		return ret
	}
	return *o.LineItems
}

// GetLineItemsOk returns a tuple with the LineItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceRequest) GetLineItemsOk() (*[]InvoiceLineItemRequest, bool) {
	if o == nil || o.LineItems == nil {
		return nil, false
	}
	return o.LineItems, true
}

// HasLineItems returns a boolean if a field has been set.
func (o *InvoiceRequest) HasLineItems() bool {
	if o != nil && o.LineItems != nil {
		return true
	}

	return false
}

// SetLineItems gets a reference to the given []InvoiceLineItemRequest and assigns it to the LineItems field.
func (o *InvoiceRequest) SetLineItems(v []InvoiceLineItemRequest) {
	o.LineItems = &v
}

// GetIntegrationParams returns the IntegrationParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InvoiceRequest) GetIntegrationParams() map[string]interface{} {
	if o == nil  {
		var ret map[string]interface{}
		return ret
	}
	return o.IntegrationParams
}

// GetIntegrationParamsOk returns a tuple with the IntegrationParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvoiceRequest) GetIntegrationParamsOk() (*map[string]interface{}, bool) {
	if o == nil || o.IntegrationParams == nil {
		return nil, false
	}
	return &o.IntegrationParams, true
}

// HasIntegrationParams returns a boolean if a field has been set.
func (o *InvoiceRequest) HasIntegrationParams() bool {
	if o != nil && o.IntegrationParams != nil {
		return true
	}

	return false
}

// SetIntegrationParams gets a reference to the given map[string]interface{} and assigns it to the IntegrationParams field.
func (o *InvoiceRequest) SetIntegrationParams(v map[string]interface{}) {
	o.IntegrationParams = v
}

// GetLinkedAccountParams returns the LinkedAccountParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InvoiceRequest) GetLinkedAccountParams() map[string]interface{} {
	if o == nil  {
		var ret map[string]interface{}
		return ret
	}
	return o.LinkedAccountParams
}

// GetLinkedAccountParamsOk returns a tuple with the LinkedAccountParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvoiceRequest) GetLinkedAccountParamsOk() (*map[string]interface{}, bool) {
	if o == nil || o.LinkedAccountParams == nil {
		return nil, false
	}
	return &o.LinkedAccountParams, true
}

// HasLinkedAccountParams returns a boolean if a field has been set.
func (o *InvoiceRequest) HasLinkedAccountParams() bool {
	if o != nil && o.LinkedAccountParams != nil {
		return true
	}

	return false
}

// SetLinkedAccountParams gets a reference to the given map[string]interface{} and assigns it to the LinkedAccountParams field.
func (o *InvoiceRequest) SetLinkedAccountParams(v map[string]interface{}) {
	o.LinkedAccountParams = v
}

func (o InvoiceRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RemoteId.IsSet() {
		toSerialize["remote_id"] = o.RemoteId.Get()
	}
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}
	if o.Contact.IsSet() {
		toSerialize["contact"] = o.Contact.Get()
	}
	if o.Number.IsSet() {
		toSerialize["number"] = o.Number.Get()
	}
	if o.IssueDate.IsSet() {
		toSerialize["issue_date"] = o.IssueDate.Get()
	}
	if o.DueDate.IsSet() {
		toSerialize["due_date"] = o.DueDate.Get()
	}
	if o.PaidOnDate.IsSet() {
		toSerialize["paid_on_date"] = o.PaidOnDate.Get()
	}
	if o.Memo.IsSet() {
		toSerialize["memo"] = o.Memo.Get()
	}
	if o.Company.IsSet() {
		toSerialize["company"] = o.Company.Get()
	}
	if o.Currency.IsSet() {
		toSerialize["currency"] = o.Currency.Get()
	}
	if o.ExchangeRate.IsSet() {
		toSerialize["exchange_rate"] = o.ExchangeRate.Get()
	}
	if o.TotalDiscount.IsSet() {
		toSerialize["total_discount"] = o.TotalDiscount.Get()
	}
	if o.SubTotal.IsSet() {
		toSerialize["sub_total"] = o.SubTotal.Get()
	}
	if o.TotalTaxAmount.IsSet() {
		toSerialize["total_tax_amount"] = o.TotalTaxAmount.Get()
	}
	if o.TotalAmount.IsSet() {
		toSerialize["total_amount"] = o.TotalAmount.Get()
	}
	if o.Balance.IsSet() {
		toSerialize["balance"] = o.Balance.Get()
	}
	if o.RemoteUpdatedAt.IsSet() {
		toSerialize["remote_updated_at"] = o.RemoteUpdatedAt.Get()
	}
	if o.Payments != nil {
		toSerialize["payments"] = o.Payments
	}
	if o.LineItems != nil {
		toSerialize["line_items"] = o.LineItems
	}
	if o.IntegrationParams != nil {
		toSerialize["integration_params"] = o.IntegrationParams
	}
	if o.LinkedAccountParams != nil {
		toSerialize["linked_account_params"] = o.LinkedAccountParams
	}
	return json.Marshal(toSerialize)
}

func (v *InvoiceRequest) UnmarshalJSON(src []byte) error {
    type InvoiceRequestUnmarshalTarget InvoiceRequest

	var intermediateResult InvoiceRequestUnmarshalTarget
	var err1 = json.Unmarshal(src, &intermediateResult)
    if err1 != nil {
        return err1
    }
    var err2 = json.Unmarshal(src, &intermediateResult.ResponseRaw)
	if err2 != nil {
		return err2
	}

	*v = InvoiceRequest(intermediateResult)
	return nil
}
type NullableInvoiceRequest struct {
	value *InvoiceRequest
	isSet bool
}

func (v NullableInvoiceRequest) Get() *InvoiceRequest {
	return v.value
}

func (v *NullableInvoiceRequest) Set(val *InvoiceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableInvoiceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableInvoiceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvoiceRequest(val *InvoiceRequest) *NullableInvoiceRequest {
	return &NullableInvoiceRequest{value: val, isSet: true}
}

func (v NullableInvoiceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvoiceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	var err1 = json.Unmarshal(src, &v.value)
    if err1 != nil {
        return err1
    }
    return json.Unmarshal(src, &v.value.ResponseRaw)
}


