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

// AccountRequest # The Account Object ### Description The `Account` object is what companies use to track transactions. They can be both bank accounts or a general ledger account (also called a chart of accounts).  ### Usage Example Fetch from the `LIST Accounts` endpoint and view a company's accounts.
type AccountRequest struct {
	// The third-party API ID of the matching object.
	RemoteId NullableString `json:"remote_id,omitempty"`
	// The account's name.
	Name NullableString `json:"name,omitempty"`
	// The account's description.
	Description NullableString `json:"description,omitempty"`
	// The account's broadest grouping.
	Classification NullableClassificationEnum `json:"classification,omitempty"`
	// The account's type is a narrower and more specific grouping within the account's classification.
	Type NullableString `json:"type,omitempty"`
	// The account's status.
	Status NullableAccountStatusEnum `json:"status,omitempty"`
	// The account's current balance.
	CurrentBalance NullableFloat32 `json:"current_balance,omitempty"`
	// The account's currency.
	Currency NullableCurrencyEnum `json:"currency,omitempty"`
	// The account's number.
	AccountNumber NullableString `json:"account_number,omitempty"`
	// ID of the parent account.
	ParentAccount NullableString `json:"parent_account,omitempty"`
	// The company the account belongs to.
	Company NullableString `json:"company,omitempty"`
	IntegrationParams map[string]interface{} `json:"integration_params,omitempty"`
	LinkedAccountParams map[string]interface{} `json:"linked_account_params,omitempty"`
	// raw json response by property name
	ResponseRaw map[string]json.RawMessage `json:"-"`
}

// NewAccountRequest instantiates a new AccountRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountRequest() *AccountRequest {
	this := AccountRequest{}
	return &this
}

// NewAccountRequestWithDefaults instantiates a new AccountRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountRequestWithDefaults() *AccountRequest {
	this := AccountRequest{}
	return &this
}

// GetRemoteId returns the RemoteId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccountRequest) GetRemoteId() string {
	if o == nil || o.RemoteId.Get() == nil {
		var ret string
		return ret
	}
	return *o.RemoteId.Get()
}

// GetRemoteIdOk returns a tuple with the RemoteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountRequest) GetRemoteIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RemoteId.Get(), o.RemoteId.IsSet()
}

// HasRemoteId returns a boolean if a field has been set.
func (o *AccountRequest) HasRemoteId() bool {
	if o != nil && o.RemoteId.IsSet() {
		return true
	}

	return false
}

// SetRemoteId gets a reference to the given NullableString and assigns it to the RemoteId field.
func (o *AccountRequest) SetRemoteId(v string) {
	o.RemoteId.Set(&v)
}
// SetRemoteIdNil sets the value for RemoteId to be an explicit nil
func (o *AccountRequest) SetRemoteIdNil() {
	o.RemoteId.Set(nil)
}

// UnsetRemoteId ensures that no value is present for RemoteId, not even an explicit nil
func (o *AccountRequest) UnsetRemoteId() {
	o.RemoteId.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccountRequest) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountRequest) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *AccountRequest) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *AccountRequest) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *AccountRequest) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *AccountRequest) UnsetName() {
	o.Name.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccountRequest) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountRequest) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *AccountRequest) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *AccountRequest) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *AccountRequest) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *AccountRequest) UnsetDescription() {
	o.Description.Unset()
}

// GetClassification returns the Classification field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccountRequest) GetClassification() ClassificationEnum {
	if o == nil || o.Classification.Get() == nil {
		var ret ClassificationEnum
		return ret
	}
	return *o.Classification.Get()
}

// GetClassificationOk returns a tuple with the Classification field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountRequest) GetClassificationOk() (*ClassificationEnum, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Classification.Get(), o.Classification.IsSet()
}

// HasClassification returns a boolean if a field has been set.
func (o *AccountRequest) HasClassification() bool {
	if o != nil && o.Classification.IsSet() {
		return true
	}

	return false
}

// SetClassification gets a reference to the given NullableClassificationEnum and assigns it to the Classification field.
func (o *AccountRequest) SetClassification(v ClassificationEnum) {
	o.Classification.Set(&v)
}
// SetClassificationNil sets the value for Classification to be an explicit nil
func (o *AccountRequest) SetClassificationNil() {
	o.Classification.Set(nil)
}

// UnsetClassification ensures that no value is present for Classification, not even an explicit nil
func (o *AccountRequest) UnsetClassification() {
	o.Classification.Unset()
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccountRequest) GetType() string {
	if o == nil || o.Type.Get() == nil {
		var ret string
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountRequest) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *AccountRequest) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableString and assigns it to the Type field.
func (o *AccountRequest) SetType(v string) {
	o.Type.Set(&v)
}
// SetTypeNil sets the value for Type to be an explicit nil
func (o *AccountRequest) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *AccountRequest) UnsetType() {
	o.Type.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccountRequest) GetStatus() AccountStatusEnum {
	if o == nil || o.Status.Get() == nil {
		var ret AccountStatusEnum
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountRequest) GetStatusOk() (*AccountStatusEnum, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *AccountRequest) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableAccountStatusEnum and assigns it to the Status field.
func (o *AccountRequest) SetStatus(v AccountStatusEnum) {
	o.Status.Set(&v)
}
// SetStatusNil sets the value for Status to be an explicit nil
func (o *AccountRequest) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *AccountRequest) UnsetStatus() {
	o.Status.Unset()
}

// GetCurrentBalance returns the CurrentBalance field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccountRequest) GetCurrentBalance() float32 {
	if o == nil || o.CurrentBalance.Get() == nil {
		var ret float32
		return ret
	}
	return *o.CurrentBalance.Get()
}

// GetCurrentBalanceOk returns a tuple with the CurrentBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountRequest) GetCurrentBalanceOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CurrentBalance.Get(), o.CurrentBalance.IsSet()
}

// HasCurrentBalance returns a boolean if a field has been set.
func (o *AccountRequest) HasCurrentBalance() bool {
	if o != nil && o.CurrentBalance.IsSet() {
		return true
	}

	return false
}

// SetCurrentBalance gets a reference to the given NullableFloat32 and assigns it to the CurrentBalance field.
func (o *AccountRequest) SetCurrentBalance(v float32) {
	o.CurrentBalance.Set(&v)
}
// SetCurrentBalanceNil sets the value for CurrentBalance to be an explicit nil
func (o *AccountRequest) SetCurrentBalanceNil() {
	o.CurrentBalance.Set(nil)
}

// UnsetCurrentBalance ensures that no value is present for CurrentBalance, not even an explicit nil
func (o *AccountRequest) UnsetCurrentBalance() {
	o.CurrentBalance.Unset()
}

// GetCurrency returns the Currency field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccountRequest) GetCurrency() CurrencyEnum {
	if o == nil || o.Currency.Get() == nil {
		var ret CurrencyEnum
		return ret
	}
	return *o.Currency.Get()
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountRequest) GetCurrencyOk() (*CurrencyEnum, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Currency.Get(), o.Currency.IsSet()
}

// HasCurrency returns a boolean if a field has been set.
func (o *AccountRequest) HasCurrency() bool {
	if o != nil && o.Currency.IsSet() {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given NullableCurrencyEnum and assigns it to the Currency field.
func (o *AccountRequest) SetCurrency(v CurrencyEnum) {
	o.Currency.Set(&v)
}
// SetCurrencyNil sets the value for Currency to be an explicit nil
func (o *AccountRequest) SetCurrencyNil() {
	o.Currency.Set(nil)
}

// UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
func (o *AccountRequest) UnsetCurrency() {
	o.Currency.Unset()
}

// GetAccountNumber returns the AccountNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccountRequest) GetAccountNumber() string {
	if o == nil || o.AccountNumber.Get() == nil {
		var ret string
		return ret
	}
	return *o.AccountNumber.Get()
}

// GetAccountNumberOk returns a tuple with the AccountNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountRequest) GetAccountNumberOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AccountNumber.Get(), o.AccountNumber.IsSet()
}

// HasAccountNumber returns a boolean if a field has been set.
func (o *AccountRequest) HasAccountNumber() bool {
	if o != nil && o.AccountNumber.IsSet() {
		return true
	}

	return false
}

// SetAccountNumber gets a reference to the given NullableString and assigns it to the AccountNumber field.
func (o *AccountRequest) SetAccountNumber(v string) {
	o.AccountNumber.Set(&v)
}
// SetAccountNumberNil sets the value for AccountNumber to be an explicit nil
func (o *AccountRequest) SetAccountNumberNil() {
	o.AccountNumber.Set(nil)
}

// UnsetAccountNumber ensures that no value is present for AccountNumber, not even an explicit nil
func (o *AccountRequest) UnsetAccountNumber() {
	o.AccountNumber.Unset()
}

// GetParentAccount returns the ParentAccount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccountRequest) GetParentAccount() string {
	if o == nil || o.ParentAccount.Get() == nil {
		var ret string
		return ret
	}
	return *o.ParentAccount.Get()
}

// GetParentAccountOk returns a tuple with the ParentAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountRequest) GetParentAccountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ParentAccount.Get(), o.ParentAccount.IsSet()
}

// HasParentAccount returns a boolean if a field has been set.
func (o *AccountRequest) HasParentAccount() bool {
	if o != nil && o.ParentAccount.IsSet() {
		return true
	}

	return false
}

// SetParentAccount gets a reference to the given NullableString and assigns it to the ParentAccount field.
func (o *AccountRequest) SetParentAccount(v string) {
	o.ParentAccount.Set(&v)
}
// SetParentAccountNil sets the value for ParentAccount to be an explicit nil
func (o *AccountRequest) SetParentAccountNil() {
	o.ParentAccount.Set(nil)
}

// UnsetParentAccount ensures that no value is present for ParentAccount, not even an explicit nil
func (o *AccountRequest) UnsetParentAccount() {
	o.ParentAccount.Unset()
}

// GetCompany returns the Company field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccountRequest) GetCompany() string {
	if o == nil || o.Company.Get() == nil {
		var ret string
		return ret
	}
	return *o.Company.Get()
}

// GetCompanyOk returns a tuple with the Company field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountRequest) GetCompanyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Company.Get(), o.Company.IsSet()
}

// HasCompany returns a boolean if a field has been set.
func (o *AccountRequest) HasCompany() bool {
	if o != nil && o.Company.IsSet() {
		return true
	}

	return false
}

// SetCompany gets a reference to the given NullableString and assigns it to the Company field.
func (o *AccountRequest) SetCompany(v string) {
	o.Company.Set(&v)
}
// SetCompanyNil sets the value for Company to be an explicit nil
func (o *AccountRequest) SetCompanyNil() {
	o.Company.Set(nil)
}

// UnsetCompany ensures that no value is present for Company, not even an explicit nil
func (o *AccountRequest) UnsetCompany() {
	o.Company.Unset()
}

// GetIntegrationParams returns the IntegrationParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccountRequest) GetIntegrationParams() map[string]interface{} {
	if o == nil  {
		var ret map[string]interface{}
		return ret
	}
	return o.IntegrationParams
}

// GetIntegrationParamsOk returns a tuple with the IntegrationParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountRequest) GetIntegrationParamsOk() (*map[string]interface{}, bool) {
	if o == nil || o.IntegrationParams == nil {
		return nil, false
	}
	return &o.IntegrationParams, true
}

// HasIntegrationParams returns a boolean if a field has been set.
func (o *AccountRequest) HasIntegrationParams() bool {
	if o != nil && o.IntegrationParams != nil {
		return true
	}

	return false
}

// SetIntegrationParams gets a reference to the given map[string]interface{} and assigns it to the IntegrationParams field.
func (o *AccountRequest) SetIntegrationParams(v map[string]interface{}) {
	o.IntegrationParams = v
}

// GetLinkedAccountParams returns the LinkedAccountParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccountRequest) GetLinkedAccountParams() map[string]interface{} {
	if o == nil  {
		var ret map[string]interface{}
		return ret
	}
	return o.LinkedAccountParams
}

// GetLinkedAccountParamsOk returns a tuple with the LinkedAccountParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountRequest) GetLinkedAccountParamsOk() (*map[string]interface{}, bool) {
	if o == nil || o.LinkedAccountParams == nil {
		return nil, false
	}
	return &o.LinkedAccountParams, true
}

// HasLinkedAccountParams returns a boolean if a field has been set.
func (o *AccountRequest) HasLinkedAccountParams() bool {
	if o != nil && o.LinkedAccountParams != nil {
		return true
	}

	return false
}

// SetLinkedAccountParams gets a reference to the given map[string]interface{} and assigns it to the LinkedAccountParams field.
func (o *AccountRequest) SetLinkedAccountParams(v map[string]interface{}) {
	o.LinkedAccountParams = v
}

func (o AccountRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RemoteId.IsSet() {
		toSerialize["remote_id"] = o.RemoteId.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.Classification.IsSet() {
		toSerialize["classification"] = o.Classification.Get()
	}
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}
	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
	}
	if o.CurrentBalance.IsSet() {
		toSerialize["current_balance"] = o.CurrentBalance.Get()
	}
	if o.Currency.IsSet() {
		toSerialize["currency"] = o.Currency.Get()
	}
	if o.AccountNumber.IsSet() {
		toSerialize["account_number"] = o.AccountNumber.Get()
	}
	if o.ParentAccount.IsSet() {
		toSerialize["parent_account"] = o.ParentAccount.Get()
	}
	if o.Company.IsSet() {
		toSerialize["company"] = o.Company.Get()
	}
	if o.IntegrationParams != nil {
		toSerialize["integration_params"] = o.IntegrationParams
	}
	if o.LinkedAccountParams != nil {
		toSerialize["linked_account_params"] = o.LinkedAccountParams
	}
	return json.Marshal(toSerialize)
}

func (v *AccountRequest) UnmarshalJSON(src []byte) error {
    type AccountRequestUnmarshalTarget AccountRequest

	var intermediateResult AccountRequestUnmarshalTarget
	var err1 = json.Unmarshal(src, &intermediateResult)
    if err1 != nil {
        return err1
    }
    var err2 = json.Unmarshal(src, &intermediateResult.ResponseRaw)
	if err2 != nil {
		return err2
	}

	*v = AccountRequest(intermediateResult)
	return nil
}
type NullableAccountRequest struct {
	value *AccountRequest
	isSet bool
}

func (v NullableAccountRequest) Get() *AccountRequest {
	return v.value
}

func (v *NullableAccountRequest) Set(val *AccountRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountRequest(val *AccountRequest) *NullableAccountRequest {
	return &NullableAccountRequest{value: val, isSet: true}
}

func (v NullableAccountRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	var err1 = json.Unmarshal(src, &v.value)
    if err1 != nil {
        return err1
    }
    return json.Unmarshal(src, &v.value.ResponseRaw)
}


