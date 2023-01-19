# ExpenseRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RemoteId** | Pointer to **NullableString** | The third-party API ID of the matching object. | [optional] 
**TransactionDate** | Pointer to **NullableTime** | When the transaction occurred. | [optional] 
**RemoteCreatedAt** | Pointer to **NullableTime** | When the expense was created. | [optional] 
**Account** | Pointer to **NullableString** | The expense&#39;s payment account. | [optional] 
**Contact** | Pointer to **NullableString** | The expense&#39;s contact. | [optional] 
**TotalAmount** | Pointer to **NullableFloat32** | The expense&#39;s total amount. | [optional] 
**Currency** | Pointer to [**NullableCurrencyEnum**](CurrencyEnum.md) | The expense&#39;s currency. | [optional] 
**ExchangeRate** | Pointer to **NullableFloat64** | The expense&#39;s exchange rate. | [optional] 
**Company** | Pointer to **NullableString** | The company the expense belongs to. | [optional] 
**Memo** | Pointer to **NullableString** | The expense&#39;s private note. | [optional] 
**Lines** | Pointer to [**[]ExpenseLineRequest**](ExpenseLineRequest.md) |  | [optional] 
**IntegrationParams** | Pointer to **map[string]interface{}** |  | [optional] 
**LinkedAccountParams** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewExpenseRequest

`func NewExpenseRequest() *ExpenseRequest`

NewExpenseRequest instantiates a new ExpenseRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExpenseRequestWithDefaults

`func NewExpenseRequestWithDefaults() *ExpenseRequest`

NewExpenseRequestWithDefaults instantiates a new ExpenseRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRemoteId

`func (o *ExpenseRequest) GetRemoteId() string`

GetRemoteId returns the RemoteId field if non-nil, zero value otherwise.

### GetRemoteIdOk

`func (o *ExpenseRequest) GetRemoteIdOk() (*string, bool)`

GetRemoteIdOk returns a tuple with the RemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteId

`func (o *ExpenseRequest) SetRemoteId(v string)`

SetRemoteId sets RemoteId field to given value.

### HasRemoteId

`func (o *ExpenseRequest) HasRemoteId() bool`

HasRemoteId returns a boolean if a field has been set.

### SetRemoteIdNil

`func (o *ExpenseRequest) SetRemoteIdNil(b bool)`

 SetRemoteIdNil sets the value for RemoteId to be an explicit nil

### UnsetRemoteId
`func (o *ExpenseRequest) UnsetRemoteId()`

UnsetRemoteId ensures that no value is present for RemoteId, not even an explicit nil
### GetTransactionDate

`func (o *ExpenseRequest) GetTransactionDate() time.Time`

GetTransactionDate returns the TransactionDate field if non-nil, zero value otherwise.

### GetTransactionDateOk

`func (o *ExpenseRequest) GetTransactionDateOk() (*time.Time, bool)`

GetTransactionDateOk returns a tuple with the TransactionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionDate

`func (o *ExpenseRequest) SetTransactionDate(v time.Time)`

SetTransactionDate sets TransactionDate field to given value.

### HasTransactionDate

`func (o *ExpenseRequest) HasTransactionDate() bool`

HasTransactionDate returns a boolean if a field has been set.

### SetTransactionDateNil

`func (o *ExpenseRequest) SetTransactionDateNil(b bool)`

 SetTransactionDateNil sets the value for TransactionDate to be an explicit nil

### UnsetTransactionDate
`func (o *ExpenseRequest) UnsetTransactionDate()`

UnsetTransactionDate ensures that no value is present for TransactionDate, not even an explicit nil
### GetRemoteCreatedAt

`func (o *ExpenseRequest) GetRemoteCreatedAt() time.Time`

GetRemoteCreatedAt returns the RemoteCreatedAt field if non-nil, zero value otherwise.

### GetRemoteCreatedAtOk

`func (o *ExpenseRequest) GetRemoteCreatedAtOk() (*time.Time, bool)`

GetRemoteCreatedAtOk returns a tuple with the RemoteCreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteCreatedAt

`func (o *ExpenseRequest) SetRemoteCreatedAt(v time.Time)`

SetRemoteCreatedAt sets RemoteCreatedAt field to given value.

### HasRemoteCreatedAt

`func (o *ExpenseRequest) HasRemoteCreatedAt() bool`

HasRemoteCreatedAt returns a boolean if a field has been set.

### SetRemoteCreatedAtNil

`func (o *ExpenseRequest) SetRemoteCreatedAtNil(b bool)`

 SetRemoteCreatedAtNil sets the value for RemoteCreatedAt to be an explicit nil

### UnsetRemoteCreatedAt
`func (o *ExpenseRequest) UnsetRemoteCreatedAt()`

UnsetRemoteCreatedAt ensures that no value is present for RemoteCreatedAt, not even an explicit nil
### GetAccount

`func (o *ExpenseRequest) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ExpenseRequest) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ExpenseRequest) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ExpenseRequest) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### SetAccountNil

`func (o *ExpenseRequest) SetAccountNil(b bool)`

 SetAccountNil sets the value for Account to be an explicit nil

### UnsetAccount
`func (o *ExpenseRequest) UnsetAccount()`

UnsetAccount ensures that no value is present for Account, not even an explicit nil
### GetContact

`func (o *ExpenseRequest) GetContact() string`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *ExpenseRequest) GetContactOk() (*string, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *ExpenseRequest) SetContact(v string)`

SetContact sets Contact field to given value.

### HasContact

`func (o *ExpenseRequest) HasContact() bool`

HasContact returns a boolean if a field has been set.

### SetContactNil

`func (o *ExpenseRequest) SetContactNil(b bool)`

 SetContactNil sets the value for Contact to be an explicit nil

### UnsetContact
`func (o *ExpenseRequest) UnsetContact()`

UnsetContact ensures that no value is present for Contact, not even an explicit nil
### GetTotalAmount

`func (o *ExpenseRequest) GetTotalAmount() float32`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *ExpenseRequest) GetTotalAmountOk() (*float32, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *ExpenseRequest) SetTotalAmount(v float32)`

SetTotalAmount sets TotalAmount field to given value.

### HasTotalAmount

`func (o *ExpenseRequest) HasTotalAmount() bool`

HasTotalAmount returns a boolean if a field has been set.

### SetTotalAmountNil

`func (o *ExpenseRequest) SetTotalAmountNil(b bool)`

 SetTotalAmountNil sets the value for TotalAmount to be an explicit nil

### UnsetTotalAmount
`func (o *ExpenseRequest) UnsetTotalAmount()`

UnsetTotalAmount ensures that no value is present for TotalAmount, not even an explicit nil
### GetCurrency

`func (o *ExpenseRequest) GetCurrency() CurrencyEnum`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ExpenseRequest) GetCurrencyOk() (*CurrencyEnum, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ExpenseRequest) SetCurrency(v CurrencyEnum)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *ExpenseRequest) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### SetCurrencyNil

`func (o *ExpenseRequest) SetCurrencyNil(b bool)`

 SetCurrencyNil sets the value for Currency to be an explicit nil

### UnsetCurrency
`func (o *ExpenseRequest) UnsetCurrency()`

UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
### GetExchangeRate

`func (o *ExpenseRequest) GetExchangeRate() float64`

GetExchangeRate returns the ExchangeRate field if non-nil, zero value otherwise.

### GetExchangeRateOk

`func (o *ExpenseRequest) GetExchangeRateOk() (*float64, bool)`

GetExchangeRateOk returns a tuple with the ExchangeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeRate

`func (o *ExpenseRequest) SetExchangeRate(v float64)`

SetExchangeRate sets ExchangeRate field to given value.

### HasExchangeRate

`func (o *ExpenseRequest) HasExchangeRate() bool`

HasExchangeRate returns a boolean if a field has been set.

### SetExchangeRateNil

`func (o *ExpenseRequest) SetExchangeRateNil(b bool)`

 SetExchangeRateNil sets the value for ExchangeRate to be an explicit nil

### UnsetExchangeRate
`func (o *ExpenseRequest) UnsetExchangeRate()`

UnsetExchangeRate ensures that no value is present for ExchangeRate, not even an explicit nil
### GetCompany

`func (o *ExpenseRequest) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *ExpenseRequest) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *ExpenseRequest) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *ExpenseRequest) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### SetCompanyNil

`func (o *ExpenseRequest) SetCompanyNil(b bool)`

 SetCompanyNil sets the value for Company to be an explicit nil

### UnsetCompany
`func (o *ExpenseRequest) UnsetCompany()`

UnsetCompany ensures that no value is present for Company, not even an explicit nil
### GetMemo

`func (o *ExpenseRequest) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *ExpenseRequest) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *ExpenseRequest) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *ExpenseRequest) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### SetMemoNil

`func (o *ExpenseRequest) SetMemoNil(b bool)`

 SetMemoNil sets the value for Memo to be an explicit nil

### UnsetMemo
`func (o *ExpenseRequest) UnsetMemo()`

UnsetMemo ensures that no value is present for Memo, not even an explicit nil
### GetLines

`func (o *ExpenseRequest) GetLines() []ExpenseLineRequest`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *ExpenseRequest) GetLinesOk() (*[]ExpenseLineRequest, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *ExpenseRequest) SetLines(v []ExpenseLineRequest)`

SetLines sets Lines field to given value.

### HasLines

`func (o *ExpenseRequest) HasLines() bool`

HasLines returns a boolean if a field has been set.

### GetIntegrationParams

`func (o *ExpenseRequest) GetIntegrationParams() map[string]interface{}`

GetIntegrationParams returns the IntegrationParams field if non-nil, zero value otherwise.

### GetIntegrationParamsOk

`func (o *ExpenseRequest) GetIntegrationParamsOk() (*map[string]interface{}, bool)`

GetIntegrationParamsOk returns a tuple with the IntegrationParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationParams

`func (o *ExpenseRequest) SetIntegrationParams(v map[string]interface{})`

SetIntegrationParams sets IntegrationParams field to given value.

### HasIntegrationParams

`func (o *ExpenseRequest) HasIntegrationParams() bool`

HasIntegrationParams returns a boolean if a field has been set.

### SetIntegrationParamsNil

`func (o *ExpenseRequest) SetIntegrationParamsNil(b bool)`

 SetIntegrationParamsNil sets the value for IntegrationParams to be an explicit nil

### UnsetIntegrationParams
`func (o *ExpenseRequest) UnsetIntegrationParams()`

UnsetIntegrationParams ensures that no value is present for IntegrationParams, not even an explicit nil
### GetLinkedAccountParams

`func (o *ExpenseRequest) GetLinkedAccountParams() map[string]interface{}`

GetLinkedAccountParams returns the LinkedAccountParams field if non-nil, zero value otherwise.

### GetLinkedAccountParamsOk

`func (o *ExpenseRequest) GetLinkedAccountParamsOk() (*map[string]interface{}, bool)`

GetLinkedAccountParamsOk returns a tuple with the LinkedAccountParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedAccountParams

`func (o *ExpenseRequest) SetLinkedAccountParams(v map[string]interface{})`

SetLinkedAccountParams sets LinkedAccountParams field to given value.

### HasLinkedAccountParams

`func (o *ExpenseRequest) HasLinkedAccountParams() bool`

HasLinkedAccountParams returns a boolean if a field has been set.

### SetLinkedAccountParamsNil

`func (o *ExpenseRequest) SetLinkedAccountParamsNil(b bool)`

 SetLinkedAccountParamsNil sets the value for LinkedAccountParams to be an explicit nil

### UnsetLinkedAccountParams
`func (o *ExpenseRequest) UnsetLinkedAccountParams()`

UnsetLinkedAccountParams ensures that no value is present for LinkedAccountParams, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


