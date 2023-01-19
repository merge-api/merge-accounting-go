# Payment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**RemoteId** | Pointer to **NullableString** | The third-party API ID of the matching object. | [optional] 
**RemoteData** | Pointer to [**[]RemoteData**](RemoteData.md) |  | [optional] [readonly] 
**TransactionDate** | Pointer to **NullableTime** | The payment&#39;s transaction date. | [optional] 
**Contact** | Pointer to **NullableString** | The supplier, or customer involved in the payment. | [optional] 
**Account** | Pointer to **NullableString** | The supplier’s or customer’s account in which the payment is made. | [optional] 
**Currency** | Pointer to [**NullableCurrencyEnum**](CurrencyEnum.md) | The payment&#39;s currency. | [optional] 
**ExchangeRate** | Pointer to **NullableFloat64** | The payment&#39;s exchange rate. | [optional] 
**Company** | Pointer to **NullableString** | The company the payment belongs to. | [optional] 
**TotalAmount** | Pointer to **NullableFloat32** | The total amount of money being paid to the supplier, or customer, after taxes. | [optional] 
**RemoteUpdatedAt** | Pointer to **NullableTime** | When the third party&#39;s payment entry was updated. | [optional] 
**RemoteWasDeleted** | Pointer to **bool** | Indicates whether or not this object has been deleted by third party webhooks. | [optional] [readonly] 
**FieldMappings** | Pointer to **map[string]interface{}** |  | [optional] [readonly] 

## Methods

### NewPayment

`func NewPayment() *Payment`

NewPayment instantiates a new Payment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentWithDefaults

`func NewPaymentWithDefaults() *Payment`

NewPaymentWithDefaults instantiates a new Payment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Payment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Payment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Payment) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Payment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRemoteId

`func (o *Payment) GetRemoteId() string`

GetRemoteId returns the RemoteId field if non-nil, zero value otherwise.

### GetRemoteIdOk

`func (o *Payment) GetRemoteIdOk() (*string, bool)`

GetRemoteIdOk returns a tuple with the RemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteId

`func (o *Payment) SetRemoteId(v string)`

SetRemoteId sets RemoteId field to given value.

### HasRemoteId

`func (o *Payment) HasRemoteId() bool`

HasRemoteId returns a boolean if a field has been set.

### SetRemoteIdNil

`func (o *Payment) SetRemoteIdNil(b bool)`

 SetRemoteIdNil sets the value for RemoteId to be an explicit nil

### UnsetRemoteId
`func (o *Payment) UnsetRemoteId()`

UnsetRemoteId ensures that no value is present for RemoteId, not even an explicit nil
### GetRemoteData

`func (o *Payment) GetRemoteData() []RemoteData`

GetRemoteData returns the RemoteData field if non-nil, zero value otherwise.

### GetRemoteDataOk

`func (o *Payment) GetRemoteDataOk() (*[]RemoteData, bool)`

GetRemoteDataOk returns a tuple with the RemoteData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteData

`func (o *Payment) SetRemoteData(v []RemoteData)`

SetRemoteData sets RemoteData field to given value.

### HasRemoteData

`func (o *Payment) HasRemoteData() bool`

HasRemoteData returns a boolean if a field has been set.

### SetRemoteDataNil

`func (o *Payment) SetRemoteDataNil(b bool)`

 SetRemoteDataNil sets the value for RemoteData to be an explicit nil

### UnsetRemoteData
`func (o *Payment) UnsetRemoteData()`

UnsetRemoteData ensures that no value is present for RemoteData, not even an explicit nil
### GetTransactionDate

`func (o *Payment) GetTransactionDate() time.Time`

GetTransactionDate returns the TransactionDate field if non-nil, zero value otherwise.

### GetTransactionDateOk

`func (o *Payment) GetTransactionDateOk() (*time.Time, bool)`

GetTransactionDateOk returns a tuple with the TransactionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionDate

`func (o *Payment) SetTransactionDate(v time.Time)`

SetTransactionDate sets TransactionDate field to given value.

### HasTransactionDate

`func (o *Payment) HasTransactionDate() bool`

HasTransactionDate returns a boolean if a field has been set.

### SetTransactionDateNil

`func (o *Payment) SetTransactionDateNil(b bool)`

 SetTransactionDateNil sets the value for TransactionDate to be an explicit nil

### UnsetTransactionDate
`func (o *Payment) UnsetTransactionDate()`

UnsetTransactionDate ensures that no value is present for TransactionDate, not even an explicit nil
### GetContact

`func (o *Payment) GetContact() string`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *Payment) GetContactOk() (*string, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *Payment) SetContact(v string)`

SetContact sets Contact field to given value.

### HasContact

`func (o *Payment) HasContact() bool`

HasContact returns a boolean if a field has been set.

### SetContactNil

`func (o *Payment) SetContactNil(b bool)`

 SetContactNil sets the value for Contact to be an explicit nil

### UnsetContact
`func (o *Payment) UnsetContact()`

UnsetContact ensures that no value is present for Contact, not even an explicit nil
### GetAccount

`func (o *Payment) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *Payment) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *Payment) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *Payment) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### SetAccountNil

`func (o *Payment) SetAccountNil(b bool)`

 SetAccountNil sets the value for Account to be an explicit nil

### UnsetAccount
`func (o *Payment) UnsetAccount()`

UnsetAccount ensures that no value is present for Account, not even an explicit nil
### GetCurrency

`func (o *Payment) GetCurrency() CurrencyEnum`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Payment) GetCurrencyOk() (*CurrencyEnum, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Payment) SetCurrency(v CurrencyEnum)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *Payment) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### SetCurrencyNil

`func (o *Payment) SetCurrencyNil(b bool)`

 SetCurrencyNil sets the value for Currency to be an explicit nil

### UnsetCurrency
`func (o *Payment) UnsetCurrency()`

UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
### GetExchangeRate

`func (o *Payment) GetExchangeRate() float64`

GetExchangeRate returns the ExchangeRate field if non-nil, zero value otherwise.

### GetExchangeRateOk

`func (o *Payment) GetExchangeRateOk() (*float64, bool)`

GetExchangeRateOk returns a tuple with the ExchangeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeRate

`func (o *Payment) SetExchangeRate(v float64)`

SetExchangeRate sets ExchangeRate field to given value.

### HasExchangeRate

`func (o *Payment) HasExchangeRate() bool`

HasExchangeRate returns a boolean if a field has been set.

### SetExchangeRateNil

`func (o *Payment) SetExchangeRateNil(b bool)`

 SetExchangeRateNil sets the value for ExchangeRate to be an explicit nil

### UnsetExchangeRate
`func (o *Payment) UnsetExchangeRate()`

UnsetExchangeRate ensures that no value is present for ExchangeRate, not even an explicit nil
### GetCompany

`func (o *Payment) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *Payment) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *Payment) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *Payment) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### SetCompanyNil

`func (o *Payment) SetCompanyNil(b bool)`

 SetCompanyNil sets the value for Company to be an explicit nil

### UnsetCompany
`func (o *Payment) UnsetCompany()`

UnsetCompany ensures that no value is present for Company, not even an explicit nil
### GetTotalAmount

`func (o *Payment) GetTotalAmount() float32`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *Payment) GetTotalAmountOk() (*float32, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *Payment) SetTotalAmount(v float32)`

SetTotalAmount sets TotalAmount field to given value.

### HasTotalAmount

`func (o *Payment) HasTotalAmount() bool`

HasTotalAmount returns a boolean if a field has been set.

### SetTotalAmountNil

`func (o *Payment) SetTotalAmountNil(b bool)`

 SetTotalAmountNil sets the value for TotalAmount to be an explicit nil

### UnsetTotalAmount
`func (o *Payment) UnsetTotalAmount()`

UnsetTotalAmount ensures that no value is present for TotalAmount, not even an explicit nil
### GetRemoteUpdatedAt

`func (o *Payment) GetRemoteUpdatedAt() time.Time`

GetRemoteUpdatedAt returns the RemoteUpdatedAt field if non-nil, zero value otherwise.

### GetRemoteUpdatedAtOk

`func (o *Payment) GetRemoteUpdatedAtOk() (*time.Time, bool)`

GetRemoteUpdatedAtOk returns a tuple with the RemoteUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteUpdatedAt

`func (o *Payment) SetRemoteUpdatedAt(v time.Time)`

SetRemoteUpdatedAt sets RemoteUpdatedAt field to given value.

### HasRemoteUpdatedAt

`func (o *Payment) HasRemoteUpdatedAt() bool`

HasRemoteUpdatedAt returns a boolean if a field has been set.

### SetRemoteUpdatedAtNil

`func (o *Payment) SetRemoteUpdatedAtNil(b bool)`

 SetRemoteUpdatedAtNil sets the value for RemoteUpdatedAt to be an explicit nil

### UnsetRemoteUpdatedAt
`func (o *Payment) UnsetRemoteUpdatedAt()`

UnsetRemoteUpdatedAt ensures that no value is present for RemoteUpdatedAt, not even an explicit nil
### GetRemoteWasDeleted

`func (o *Payment) GetRemoteWasDeleted() bool`

GetRemoteWasDeleted returns the RemoteWasDeleted field if non-nil, zero value otherwise.

### GetRemoteWasDeletedOk

`func (o *Payment) GetRemoteWasDeletedOk() (*bool, bool)`

GetRemoteWasDeletedOk returns a tuple with the RemoteWasDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteWasDeleted

`func (o *Payment) SetRemoteWasDeleted(v bool)`

SetRemoteWasDeleted sets RemoteWasDeleted field to given value.

### HasRemoteWasDeleted

`func (o *Payment) HasRemoteWasDeleted() bool`

HasRemoteWasDeleted returns a boolean if a field has been set.

### GetFieldMappings

`func (o *Payment) GetFieldMappings() map[string]interface{}`

GetFieldMappings returns the FieldMappings field if non-nil, zero value otherwise.

### GetFieldMappingsOk

`func (o *Payment) GetFieldMappingsOk() (*map[string]interface{}, bool)`

GetFieldMappingsOk returns a tuple with the FieldMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldMappings

`func (o *Payment) SetFieldMappings(v map[string]interface{})`

SetFieldMappings sets FieldMappings field to given value.

### HasFieldMappings

`func (o *Payment) HasFieldMappings() bool`

HasFieldMappings returns a boolean if a field has been set.

### SetFieldMappingsNil

`func (o *Payment) SetFieldMappingsNil(b bool)`

 SetFieldMappingsNil sets the value for FieldMappings to be an explicit nil

### UnsetFieldMappings
`func (o *Payment) UnsetFieldMappings()`

UnsetFieldMappings ensures that no value is present for FieldMappings, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


