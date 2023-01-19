# PaymentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RemoteId** | Pointer to **NullableString** | The third-party API ID of the matching object. | [optional] 
**TransactionDate** | Pointer to **NullableTime** | The payment&#39;s transaction date. | [optional] 
**Contact** | Pointer to **NullableString** | The supplier, or customer involved in the payment. | [optional] 
**Account** | Pointer to **NullableString** | The supplier’s or customer’s account in which the payment is made. | [optional] 
**Currency** | Pointer to [**NullableCurrencyEnum**](CurrencyEnum.md) | The payment&#39;s currency. | [optional] 
**ExchangeRate** | Pointer to **NullableFloat64** | The payment&#39;s exchange rate. | [optional] 
**Company** | Pointer to **NullableString** | The company the payment belongs to. | [optional] 
**TotalAmount** | Pointer to **NullableFloat32** | The total amount of money being paid to the supplier, or customer, after taxes. | [optional] 
**RemoteUpdatedAt** | Pointer to **NullableTime** | When the third party&#39;s payment entry was updated. | [optional] 
**IntegrationParams** | Pointer to **map[string]interface{}** |  | [optional] 
**LinkedAccountParams** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewPaymentRequest

`func NewPaymentRequest() *PaymentRequest`

NewPaymentRequest instantiates a new PaymentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentRequestWithDefaults

`func NewPaymentRequestWithDefaults() *PaymentRequest`

NewPaymentRequestWithDefaults instantiates a new PaymentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRemoteId

`func (o *PaymentRequest) GetRemoteId() string`

GetRemoteId returns the RemoteId field if non-nil, zero value otherwise.

### GetRemoteIdOk

`func (o *PaymentRequest) GetRemoteIdOk() (*string, bool)`

GetRemoteIdOk returns a tuple with the RemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteId

`func (o *PaymentRequest) SetRemoteId(v string)`

SetRemoteId sets RemoteId field to given value.

### HasRemoteId

`func (o *PaymentRequest) HasRemoteId() bool`

HasRemoteId returns a boolean if a field has been set.

### SetRemoteIdNil

`func (o *PaymentRequest) SetRemoteIdNil(b bool)`

 SetRemoteIdNil sets the value for RemoteId to be an explicit nil

### UnsetRemoteId
`func (o *PaymentRequest) UnsetRemoteId()`

UnsetRemoteId ensures that no value is present for RemoteId, not even an explicit nil
### GetTransactionDate

`func (o *PaymentRequest) GetTransactionDate() time.Time`

GetTransactionDate returns the TransactionDate field if non-nil, zero value otherwise.

### GetTransactionDateOk

`func (o *PaymentRequest) GetTransactionDateOk() (*time.Time, bool)`

GetTransactionDateOk returns a tuple with the TransactionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionDate

`func (o *PaymentRequest) SetTransactionDate(v time.Time)`

SetTransactionDate sets TransactionDate field to given value.

### HasTransactionDate

`func (o *PaymentRequest) HasTransactionDate() bool`

HasTransactionDate returns a boolean if a field has been set.

### SetTransactionDateNil

`func (o *PaymentRequest) SetTransactionDateNil(b bool)`

 SetTransactionDateNil sets the value for TransactionDate to be an explicit nil

### UnsetTransactionDate
`func (o *PaymentRequest) UnsetTransactionDate()`

UnsetTransactionDate ensures that no value is present for TransactionDate, not even an explicit nil
### GetContact

`func (o *PaymentRequest) GetContact() string`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *PaymentRequest) GetContactOk() (*string, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *PaymentRequest) SetContact(v string)`

SetContact sets Contact field to given value.

### HasContact

`func (o *PaymentRequest) HasContact() bool`

HasContact returns a boolean if a field has been set.

### SetContactNil

`func (o *PaymentRequest) SetContactNil(b bool)`

 SetContactNil sets the value for Contact to be an explicit nil

### UnsetContact
`func (o *PaymentRequest) UnsetContact()`

UnsetContact ensures that no value is present for Contact, not even an explicit nil
### GetAccount

`func (o *PaymentRequest) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *PaymentRequest) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *PaymentRequest) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *PaymentRequest) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### SetAccountNil

`func (o *PaymentRequest) SetAccountNil(b bool)`

 SetAccountNil sets the value for Account to be an explicit nil

### UnsetAccount
`func (o *PaymentRequest) UnsetAccount()`

UnsetAccount ensures that no value is present for Account, not even an explicit nil
### GetCurrency

`func (o *PaymentRequest) GetCurrency() CurrencyEnum`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *PaymentRequest) GetCurrencyOk() (*CurrencyEnum, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *PaymentRequest) SetCurrency(v CurrencyEnum)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *PaymentRequest) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### SetCurrencyNil

`func (o *PaymentRequest) SetCurrencyNil(b bool)`

 SetCurrencyNil sets the value for Currency to be an explicit nil

### UnsetCurrency
`func (o *PaymentRequest) UnsetCurrency()`

UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
### GetExchangeRate

`func (o *PaymentRequest) GetExchangeRate() float64`

GetExchangeRate returns the ExchangeRate field if non-nil, zero value otherwise.

### GetExchangeRateOk

`func (o *PaymentRequest) GetExchangeRateOk() (*float64, bool)`

GetExchangeRateOk returns a tuple with the ExchangeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeRate

`func (o *PaymentRequest) SetExchangeRate(v float64)`

SetExchangeRate sets ExchangeRate field to given value.

### HasExchangeRate

`func (o *PaymentRequest) HasExchangeRate() bool`

HasExchangeRate returns a boolean if a field has been set.

### SetExchangeRateNil

`func (o *PaymentRequest) SetExchangeRateNil(b bool)`

 SetExchangeRateNil sets the value for ExchangeRate to be an explicit nil

### UnsetExchangeRate
`func (o *PaymentRequest) UnsetExchangeRate()`

UnsetExchangeRate ensures that no value is present for ExchangeRate, not even an explicit nil
### GetCompany

`func (o *PaymentRequest) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *PaymentRequest) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *PaymentRequest) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *PaymentRequest) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### SetCompanyNil

`func (o *PaymentRequest) SetCompanyNil(b bool)`

 SetCompanyNil sets the value for Company to be an explicit nil

### UnsetCompany
`func (o *PaymentRequest) UnsetCompany()`

UnsetCompany ensures that no value is present for Company, not even an explicit nil
### GetTotalAmount

`func (o *PaymentRequest) GetTotalAmount() float32`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *PaymentRequest) GetTotalAmountOk() (*float32, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *PaymentRequest) SetTotalAmount(v float32)`

SetTotalAmount sets TotalAmount field to given value.

### HasTotalAmount

`func (o *PaymentRequest) HasTotalAmount() bool`

HasTotalAmount returns a boolean if a field has been set.

### SetTotalAmountNil

`func (o *PaymentRequest) SetTotalAmountNil(b bool)`

 SetTotalAmountNil sets the value for TotalAmount to be an explicit nil

### UnsetTotalAmount
`func (o *PaymentRequest) UnsetTotalAmount()`

UnsetTotalAmount ensures that no value is present for TotalAmount, not even an explicit nil
### GetRemoteUpdatedAt

`func (o *PaymentRequest) GetRemoteUpdatedAt() time.Time`

GetRemoteUpdatedAt returns the RemoteUpdatedAt field if non-nil, zero value otherwise.

### GetRemoteUpdatedAtOk

`func (o *PaymentRequest) GetRemoteUpdatedAtOk() (*time.Time, bool)`

GetRemoteUpdatedAtOk returns a tuple with the RemoteUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteUpdatedAt

`func (o *PaymentRequest) SetRemoteUpdatedAt(v time.Time)`

SetRemoteUpdatedAt sets RemoteUpdatedAt field to given value.

### HasRemoteUpdatedAt

`func (o *PaymentRequest) HasRemoteUpdatedAt() bool`

HasRemoteUpdatedAt returns a boolean if a field has been set.

### SetRemoteUpdatedAtNil

`func (o *PaymentRequest) SetRemoteUpdatedAtNil(b bool)`

 SetRemoteUpdatedAtNil sets the value for RemoteUpdatedAt to be an explicit nil

### UnsetRemoteUpdatedAt
`func (o *PaymentRequest) UnsetRemoteUpdatedAt()`

UnsetRemoteUpdatedAt ensures that no value is present for RemoteUpdatedAt, not even an explicit nil
### GetIntegrationParams

`func (o *PaymentRequest) GetIntegrationParams() map[string]interface{}`

GetIntegrationParams returns the IntegrationParams field if non-nil, zero value otherwise.

### GetIntegrationParamsOk

`func (o *PaymentRequest) GetIntegrationParamsOk() (*map[string]interface{}, bool)`

GetIntegrationParamsOk returns a tuple with the IntegrationParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationParams

`func (o *PaymentRequest) SetIntegrationParams(v map[string]interface{})`

SetIntegrationParams sets IntegrationParams field to given value.

### HasIntegrationParams

`func (o *PaymentRequest) HasIntegrationParams() bool`

HasIntegrationParams returns a boolean if a field has been set.

### SetIntegrationParamsNil

`func (o *PaymentRequest) SetIntegrationParamsNil(b bool)`

 SetIntegrationParamsNil sets the value for IntegrationParams to be an explicit nil

### UnsetIntegrationParams
`func (o *PaymentRequest) UnsetIntegrationParams()`

UnsetIntegrationParams ensures that no value is present for IntegrationParams, not even an explicit nil
### GetLinkedAccountParams

`func (o *PaymentRequest) GetLinkedAccountParams() map[string]interface{}`

GetLinkedAccountParams returns the LinkedAccountParams field if non-nil, zero value otherwise.

### GetLinkedAccountParamsOk

`func (o *PaymentRequest) GetLinkedAccountParamsOk() (*map[string]interface{}, bool)`

GetLinkedAccountParamsOk returns a tuple with the LinkedAccountParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedAccountParams

`func (o *PaymentRequest) SetLinkedAccountParams(v map[string]interface{})`

SetLinkedAccountParams sets LinkedAccountParams field to given value.

### HasLinkedAccountParams

`func (o *PaymentRequest) HasLinkedAccountParams() bool`

HasLinkedAccountParams returns a boolean if a field has been set.

### SetLinkedAccountParamsNil

`func (o *PaymentRequest) SetLinkedAccountParamsNil(b bool)`

 SetLinkedAccountParamsNil sets the value for LinkedAccountParams to be an explicit nil

### UnsetLinkedAccountParams
`func (o *PaymentRequest) UnsetLinkedAccountParams()`

UnsetLinkedAccountParams ensures that no value is present for LinkedAccountParams, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


