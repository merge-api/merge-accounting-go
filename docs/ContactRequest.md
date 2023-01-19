# ContactRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RemoteId** | Pointer to **NullableString** | The third-party API ID of the matching object. | [optional] 
**Name** | Pointer to **NullableString** | The contact&#39;s name. | [optional] 
**IsSupplier** | Pointer to **NullableBool** | Whether the contact is a supplier. | [optional] 
**IsCustomer** | Pointer to **NullableBool** | Whether the contact is a customer. | [optional] 
**EmailAddress** | Pointer to **NullableString** | The contact&#39;s email address. | [optional] 
**TaxNumber** | Pointer to **NullableString** | The contact&#39;s tax number. | [optional] 
**Status** | Pointer to [**NullableStatus7d1Enum**](Status7d1Enum.md) | The contact&#39;s status | [optional] 
**Currency** | Pointer to **NullableString** | The currency the contact&#39;s transactions are in. | [optional] 
**RemoteUpdatedAt** | Pointer to **NullableTime** | When the third party&#39;s contact was updated. | [optional] 
**Company** | Pointer to **NullableString** | The company the contact belongs to. | [optional] 
**Addresses** | Pointer to **[]string** | &#x60;Address&#x60; object IDs for the given &#x60;Contacts&#x60; object. | [optional] 
**PhoneNumbers** | Pointer to [**[]AccountingPhoneNumberRequest**](AccountingPhoneNumberRequest.md) | &#x60;AccountingPhoneNumber&#x60; object for the given &#x60;Contacts&#x60; object. | [optional] 
**IntegrationParams** | Pointer to **map[string]interface{}** |  | [optional] 
**LinkedAccountParams** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewContactRequest

`func NewContactRequest() *ContactRequest`

NewContactRequest instantiates a new ContactRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactRequestWithDefaults

`func NewContactRequestWithDefaults() *ContactRequest`

NewContactRequestWithDefaults instantiates a new ContactRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRemoteId

`func (o *ContactRequest) GetRemoteId() string`

GetRemoteId returns the RemoteId field if non-nil, zero value otherwise.

### GetRemoteIdOk

`func (o *ContactRequest) GetRemoteIdOk() (*string, bool)`

GetRemoteIdOk returns a tuple with the RemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteId

`func (o *ContactRequest) SetRemoteId(v string)`

SetRemoteId sets RemoteId field to given value.

### HasRemoteId

`func (o *ContactRequest) HasRemoteId() bool`

HasRemoteId returns a boolean if a field has been set.

### SetRemoteIdNil

`func (o *ContactRequest) SetRemoteIdNil(b bool)`

 SetRemoteIdNil sets the value for RemoteId to be an explicit nil

### UnsetRemoteId
`func (o *ContactRequest) UnsetRemoteId()`

UnsetRemoteId ensures that no value is present for RemoteId, not even an explicit nil
### GetName

`func (o *ContactRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ContactRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ContactRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ContactRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ContactRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ContactRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetIsSupplier

`func (o *ContactRequest) GetIsSupplier() bool`

GetIsSupplier returns the IsSupplier field if non-nil, zero value otherwise.

### GetIsSupplierOk

`func (o *ContactRequest) GetIsSupplierOk() (*bool, bool)`

GetIsSupplierOk returns a tuple with the IsSupplier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSupplier

`func (o *ContactRequest) SetIsSupplier(v bool)`

SetIsSupplier sets IsSupplier field to given value.

### HasIsSupplier

`func (o *ContactRequest) HasIsSupplier() bool`

HasIsSupplier returns a boolean if a field has been set.

### SetIsSupplierNil

`func (o *ContactRequest) SetIsSupplierNil(b bool)`

 SetIsSupplierNil sets the value for IsSupplier to be an explicit nil

### UnsetIsSupplier
`func (o *ContactRequest) UnsetIsSupplier()`

UnsetIsSupplier ensures that no value is present for IsSupplier, not even an explicit nil
### GetIsCustomer

`func (o *ContactRequest) GetIsCustomer() bool`

GetIsCustomer returns the IsCustomer field if non-nil, zero value otherwise.

### GetIsCustomerOk

`func (o *ContactRequest) GetIsCustomerOk() (*bool, bool)`

GetIsCustomerOk returns a tuple with the IsCustomer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCustomer

`func (o *ContactRequest) SetIsCustomer(v bool)`

SetIsCustomer sets IsCustomer field to given value.

### HasIsCustomer

`func (o *ContactRequest) HasIsCustomer() bool`

HasIsCustomer returns a boolean if a field has been set.

### SetIsCustomerNil

`func (o *ContactRequest) SetIsCustomerNil(b bool)`

 SetIsCustomerNil sets the value for IsCustomer to be an explicit nil

### UnsetIsCustomer
`func (o *ContactRequest) UnsetIsCustomer()`

UnsetIsCustomer ensures that no value is present for IsCustomer, not even an explicit nil
### GetEmailAddress

`func (o *ContactRequest) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *ContactRequest) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *ContactRequest) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.

### HasEmailAddress

`func (o *ContactRequest) HasEmailAddress() bool`

HasEmailAddress returns a boolean if a field has been set.

### SetEmailAddressNil

`func (o *ContactRequest) SetEmailAddressNil(b bool)`

 SetEmailAddressNil sets the value for EmailAddress to be an explicit nil

### UnsetEmailAddress
`func (o *ContactRequest) UnsetEmailAddress()`

UnsetEmailAddress ensures that no value is present for EmailAddress, not even an explicit nil
### GetTaxNumber

`func (o *ContactRequest) GetTaxNumber() string`

GetTaxNumber returns the TaxNumber field if non-nil, zero value otherwise.

### GetTaxNumberOk

`func (o *ContactRequest) GetTaxNumberOk() (*string, bool)`

GetTaxNumberOk returns a tuple with the TaxNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxNumber

`func (o *ContactRequest) SetTaxNumber(v string)`

SetTaxNumber sets TaxNumber field to given value.

### HasTaxNumber

`func (o *ContactRequest) HasTaxNumber() bool`

HasTaxNumber returns a boolean if a field has been set.

### SetTaxNumberNil

`func (o *ContactRequest) SetTaxNumberNil(b bool)`

 SetTaxNumberNil sets the value for TaxNumber to be an explicit nil

### UnsetTaxNumber
`func (o *ContactRequest) UnsetTaxNumber()`

UnsetTaxNumber ensures that no value is present for TaxNumber, not even an explicit nil
### GetStatus

`func (o *ContactRequest) GetStatus() Status7d1Enum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ContactRequest) GetStatusOk() (*Status7d1Enum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ContactRequest) SetStatus(v Status7d1Enum)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ContactRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *ContactRequest) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *ContactRequest) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetCurrency

`func (o *ContactRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ContactRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ContactRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *ContactRequest) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### SetCurrencyNil

`func (o *ContactRequest) SetCurrencyNil(b bool)`

 SetCurrencyNil sets the value for Currency to be an explicit nil

### UnsetCurrency
`func (o *ContactRequest) UnsetCurrency()`

UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
### GetRemoteUpdatedAt

`func (o *ContactRequest) GetRemoteUpdatedAt() time.Time`

GetRemoteUpdatedAt returns the RemoteUpdatedAt field if non-nil, zero value otherwise.

### GetRemoteUpdatedAtOk

`func (o *ContactRequest) GetRemoteUpdatedAtOk() (*time.Time, bool)`

GetRemoteUpdatedAtOk returns a tuple with the RemoteUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteUpdatedAt

`func (o *ContactRequest) SetRemoteUpdatedAt(v time.Time)`

SetRemoteUpdatedAt sets RemoteUpdatedAt field to given value.

### HasRemoteUpdatedAt

`func (o *ContactRequest) HasRemoteUpdatedAt() bool`

HasRemoteUpdatedAt returns a boolean if a field has been set.

### SetRemoteUpdatedAtNil

`func (o *ContactRequest) SetRemoteUpdatedAtNil(b bool)`

 SetRemoteUpdatedAtNil sets the value for RemoteUpdatedAt to be an explicit nil

### UnsetRemoteUpdatedAt
`func (o *ContactRequest) UnsetRemoteUpdatedAt()`

UnsetRemoteUpdatedAt ensures that no value is present for RemoteUpdatedAt, not even an explicit nil
### GetCompany

`func (o *ContactRequest) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *ContactRequest) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *ContactRequest) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *ContactRequest) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### SetCompanyNil

`func (o *ContactRequest) SetCompanyNil(b bool)`

 SetCompanyNil sets the value for Company to be an explicit nil

### UnsetCompany
`func (o *ContactRequest) UnsetCompany()`

UnsetCompany ensures that no value is present for Company, not even an explicit nil
### GetAddresses

`func (o *ContactRequest) GetAddresses() []string`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *ContactRequest) GetAddressesOk() (*[]string, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *ContactRequest) SetAddresses(v []string)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *ContactRequest) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.

### GetPhoneNumbers

`func (o *ContactRequest) GetPhoneNumbers() []AccountingPhoneNumberRequest`

GetPhoneNumbers returns the PhoneNumbers field if non-nil, zero value otherwise.

### GetPhoneNumbersOk

`func (o *ContactRequest) GetPhoneNumbersOk() (*[]AccountingPhoneNumberRequest, bool)`

GetPhoneNumbersOk returns a tuple with the PhoneNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumbers

`func (o *ContactRequest) SetPhoneNumbers(v []AccountingPhoneNumberRequest)`

SetPhoneNumbers sets PhoneNumbers field to given value.

### HasPhoneNumbers

`func (o *ContactRequest) HasPhoneNumbers() bool`

HasPhoneNumbers returns a boolean if a field has been set.

### GetIntegrationParams

`func (o *ContactRequest) GetIntegrationParams() map[string]interface{}`

GetIntegrationParams returns the IntegrationParams field if non-nil, zero value otherwise.

### GetIntegrationParamsOk

`func (o *ContactRequest) GetIntegrationParamsOk() (*map[string]interface{}, bool)`

GetIntegrationParamsOk returns a tuple with the IntegrationParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationParams

`func (o *ContactRequest) SetIntegrationParams(v map[string]interface{})`

SetIntegrationParams sets IntegrationParams field to given value.

### HasIntegrationParams

`func (o *ContactRequest) HasIntegrationParams() bool`

HasIntegrationParams returns a boolean if a field has been set.

### SetIntegrationParamsNil

`func (o *ContactRequest) SetIntegrationParamsNil(b bool)`

 SetIntegrationParamsNil sets the value for IntegrationParams to be an explicit nil

### UnsetIntegrationParams
`func (o *ContactRequest) UnsetIntegrationParams()`

UnsetIntegrationParams ensures that no value is present for IntegrationParams, not even an explicit nil
### GetLinkedAccountParams

`func (o *ContactRequest) GetLinkedAccountParams() map[string]interface{}`

GetLinkedAccountParams returns the LinkedAccountParams field if non-nil, zero value otherwise.

### GetLinkedAccountParamsOk

`func (o *ContactRequest) GetLinkedAccountParamsOk() (*map[string]interface{}, bool)`

GetLinkedAccountParamsOk returns a tuple with the LinkedAccountParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedAccountParams

`func (o *ContactRequest) SetLinkedAccountParams(v map[string]interface{})`

SetLinkedAccountParams sets LinkedAccountParams field to given value.

### HasLinkedAccountParams

`func (o *ContactRequest) HasLinkedAccountParams() bool`

HasLinkedAccountParams returns a boolean if a field has been set.

### SetLinkedAccountParamsNil

`func (o *ContactRequest) SetLinkedAccountParamsNil(b bool)`

 SetLinkedAccountParamsNil sets the value for LinkedAccountParams to be an explicit nil

### UnsetLinkedAccountParams
`func (o *ContactRequest) UnsetLinkedAccountParams()`

UnsetLinkedAccountParams ensures that no value is present for LinkedAccountParams, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


