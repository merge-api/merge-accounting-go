# Expense

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**RemoteId** | Pointer to **NullableString** | The third-party API ID of the matching object. | [optional] 
**RemoteData** | Pointer to [**[]RemoteData**](RemoteData.md) |  | [optional] [readonly] 
**TransactionDate** | Pointer to **NullableTime** | When the transaction occurred. | [optional] 
**RemoteCreatedAt** | Pointer to **NullableTime** | When the expense was created. | [optional] 
**Account** | Pointer to **NullableString** | The expense&#39;s payment account. | [optional] 
**Contact** | Pointer to **NullableString** | The expense&#39;s contact. | [optional] 
**TotalAmount** | Pointer to **NullableFloat32** | The expense&#39;s total amount. | [optional] 
**Currency** | Pointer to [**NullableCurrencyEnum**](CurrencyEnum.md) | The expense&#39;s currency. | [optional] 
**ExchangeRate** | Pointer to **NullableFloat64** | The expense&#39;s exchange rate. | [optional] 
**Company** | Pointer to **NullableString** | The company the expense belongs to. | [optional] 
**Memo** | Pointer to **NullableString** | The expense&#39;s private note. | [optional] 
**Lines** | Pointer to [**[]ExpenseLine**](ExpenseLine.md) |  | [optional] 
**RemoteWasDeleted** | Pointer to **bool** | Indicates whether or not this object has been deleted by third party webhooks. | [optional] [readonly] 
**FieldMappings** | Pointer to **map[string]interface{}** |  | [optional] [readonly] 

## Methods

### NewExpense

`func NewExpense() *Expense`

NewExpense instantiates a new Expense object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExpenseWithDefaults

`func NewExpenseWithDefaults() *Expense`

NewExpenseWithDefaults instantiates a new Expense object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Expense) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Expense) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Expense) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Expense) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRemoteId

`func (o *Expense) GetRemoteId() string`

GetRemoteId returns the RemoteId field if non-nil, zero value otherwise.

### GetRemoteIdOk

`func (o *Expense) GetRemoteIdOk() (*string, bool)`

GetRemoteIdOk returns a tuple with the RemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteId

`func (o *Expense) SetRemoteId(v string)`

SetRemoteId sets RemoteId field to given value.

### HasRemoteId

`func (o *Expense) HasRemoteId() bool`

HasRemoteId returns a boolean if a field has been set.

### SetRemoteIdNil

`func (o *Expense) SetRemoteIdNil(b bool)`

 SetRemoteIdNil sets the value for RemoteId to be an explicit nil

### UnsetRemoteId
`func (o *Expense) UnsetRemoteId()`

UnsetRemoteId ensures that no value is present for RemoteId, not even an explicit nil
### GetRemoteData

`func (o *Expense) GetRemoteData() []RemoteData`

GetRemoteData returns the RemoteData field if non-nil, zero value otherwise.

### GetRemoteDataOk

`func (o *Expense) GetRemoteDataOk() (*[]RemoteData, bool)`

GetRemoteDataOk returns a tuple with the RemoteData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteData

`func (o *Expense) SetRemoteData(v []RemoteData)`

SetRemoteData sets RemoteData field to given value.

### HasRemoteData

`func (o *Expense) HasRemoteData() bool`

HasRemoteData returns a boolean if a field has been set.

### SetRemoteDataNil

`func (o *Expense) SetRemoteDataNil(b bool)`

 SetRemoteDataNil sets the value for RemoteData to be an explicit nil

### UnsetRemoteData
`func (o *Expense) UnsetRemoteData()`

UnsetRemoteData ensures that no value is present for RemoteData, not even an explicit nil
### GetTransactionDate

`func (o *Expense) GetTransactionDate() time.Time`

GetTransactionDate returns the TransactionDate field if non-nil, zero value otherwise.

### GetTransactionDateOk

`func (o *Expense) GetTransactionDateOk() (*time.Time, bool)`

GetTransactionDateOk returns a tuple with the TransactionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionDate

`func (o *Expense) SetTransactionDate(v time.Time)`

SetTransactionDate sets TransactionDate field to given value.

### HasTransactionDate

`func (o *Expense) HasTransactionDate() bool`

HasTransactionDate returns a boolean if a field has been set.

### SetTransactionDateNil

`func (o *Expense) SetTransactionDateNil(b bool)`

 SetTransactionDateNil sets the value for TransactionDate to be an explicit nil

### UnsetTransactionDate
`func (o *Expense) UnsetTransactionDate()`

UnsetTransactionDate ensures that no value is present for TransactionDate, not even an explicit nil
### GetRemoteCreatedAt

`func (o *Expense) GetRemoteCreatedAt() time.Time`

GetRemoteCreatedAt returns the RemoteCreatedAt field if non-nil, zero value otherwise.

### GetRemoteCreatedAtOk

`func (o *Expense) GetRemoteCreatedAtOk() (*time.Time, bool)`

GetRemoteCreatedAtOk returns a tuple with the RemoteCreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteCreatedAt

`func (o *Expense) SetRemoteCreatedAt(v time.Time)`

SetRemoteCreatedAt sets RemoteCreatedAt field to given value.

### HasRemoteCreatedAt

`func (o *Expense) HasRemoteCreatedAt() bool`

HasRemoteCreatedAt returns a boolean if a field has been set.

### SetRemoteCreatedAtNil

`func (o *Expense) SetRemoteCreatedAtNil(b bool)`

 SetRemoteCreatedAtNil sets the value for RemoteCreatedAt to be an explicit nil

### UnsetRemoteCreatedAt
`func (o *Expense) UnsetRemoteCreatedAt()`

UnsetRemoteCreatedAt ensures that no value is present for RemoteCreatedAt, not even an explicit nil
### GetAccount

`func (o *Expense) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *Expense) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *Expense) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *Expense) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### SetAccountNil

`func (o *Expense) SetAccountNil(b bool)`

 SetAccountNil sets the value for Account to be an explicit nil

### UnsetAccount
`func (o *Expense) UnsetAccount()`

UnsetAccount ensures that no value is present for Account, not even an explicit nil
### GetContact

`func (o *Expense) GetContact() string`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *Expense) GetContactOk() (*string, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *Expense) SetContact(v string)`

SetContact sets Contact field to given value.

### HasContact

`func (o *Expense) HasContact() bool`

HasContact returns a boolean if a field has been set.

### SetContactNil

`func (o *Expense) SetContactNil(b bool)`

 SetContactNil sets the value for Contact to be an explicit nil

### UnsetContact
`func (o *Expense) UnsetContact()`

UnsetContact ensures that no value is present for Contact, not even an explicit nil
### GetTotalAmount

`func (o *Expense) GetTotalAmount() float32`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *Expense) GetTotalAmountOk() (*float32, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *Expense) SetTotalAmount(v float32)`

SetTotalAmount sets TotalAmount field to given value.

### HasTotalAmount

`func (o *Expense) HasTotalAmount() bool`

HasTotalAmount returns a boolean if a field has been set.

### SetTotalAmountNil

`func (o *Expense) SetTotalAmountNil(b bool)`

 SetTotalAmountNil sets the value for TotalAmount to be an explicit nil

### UnsetTotalAmount
`func (o *Expense) UnsetTotalAmount()`

UnsetTotalAmount ensures that no value is present for TotalAmount, not even an explicit nil
### GetCurrency

`func (o *Expense) GetCurrency() CurrencyEnum`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Expense) GetCurrencyOk() (*CurrencyEnum, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Expense) SetCurrency(v CurrencyEnum)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *Expense) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### SetCurrencyNil

`func (o *Expense) SetCurrencyNil(b bool)`

 SetCurrencyNil sets the value for Currency to be an explicit nil

### UnsetCurrency
`func (o *Expense) UnsetCurrency()`

UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
### GetExchangeRate

`func (o *Expense) GetExchangeRate() float64`

GetExchangeRate returns the ExchangeRate field if non-nil, zero value otherwise.

### GetExchangeRateOk

`func (o *Expense) GetExchangeRateOk() (*float64, bool)`

GetExchangeRateOk returns a tuple with the ExchangeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeRate

`func (o *Expense) SetExchangeRate(v float64)`

SetExchangeRate sets ExchangeRate field to given value.

### HasExchangeRate

`func (o *Expense) HasExchangeRate() bool`

HasExchangeRate returns a boolean if a field has been set.

### SetExchangeRateNil

`func (o *Expense) SetExchangeRateNil(b bool)`

 SetExchangeRateNil sets the value for ExchangeRate to be an explicit nil

### UnsetExchangeRate
`func (o *Expense) UnsetExchangeRate()`

UnsetExchangeRate ensures that no value is present for ExchangeRate, not even an explicit nil
### GetCompany

`func (o *Expense) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *Expense) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *Expense) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *Expense) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### SetCompanyNil

`func (o *Expense) SetCompanyNil(b bool)`

 SetCompanyNil sets the value for Company to be an explicit nil

### UnsetCompany
`func (o *Expense) UnsetCompany()`

UnsetCompany ensures that no value is present for Company, not even an explicit nil
### GetMemo

`func (o *Expense) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *Expense) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *Expense) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *Expense) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### SetMemoNil

`func (o *Expense) SetMemoNil(b bool)`

 SetMemoNil sets the value for Memo to be an explicit nil

### UnsetMemo
`func (o *Expense) UnsetMemo()`

UnsetMemo ensures that no value is present for Memo, not even an explicit nil
### GetLines

`func (o *Expense) GetLines() []ExpenseLine`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *Expense) GetLinesOk() (*[]ExpenseLine, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *Expense) SetLines(v []ExpenseLine)`

SetLines sets Lines field to given value.

### HasLines

`func (o *Expense) HasLines() bool`

HasLines returns a boolean if a field has been set.

### GetRemoteWasDeleted

`func (o *Expense) GetRemoteWasDeleted() bool`

GetRemoteWasDeleted returns the RemoteWasDeleted field if non-nil, zero value otherwise.

### GetRemoteWasDeletedOk

`func (o *Expense) GetRemoteWasDeletedOk() (*bool, bool)`

GetRemoteWasDeletedOk returns a tuple with the RemoteWasDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteWasDeleted

`func (o *Expense) SetRemoteWasDeleted(v bool)`

SetRemoteWasDeleted sets RemoteWasDeleted field to given value.

### HasRemoteWasDeleted

`func (o *Expense) HasRemoteWasDeleted() bool`

HasRemoteWasDeleted returns a boolean if a field has been set.

### GetFieldMappings

`func (o *Expense) GetFieldMappings() map[string]interface{}`

GetFieldMappings returns the FieldMappings field if non-nil, zero value otherwise.

### GetFieldMappingsOk

`func (o *Expense) GetFieldMappingsOk() (*map[string]interface{}, bool)`

GetFieldMappingsOk returns a tuple with the FieldMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldMappings

`func (o *Expense) SetFieldMappings(v map[string]interface{})`

SetFieldMappings sets FieldMappings field to given value.

### HasFieldMappings

`func (o *Expense) HasFieldMappings() bool`

HasFieldMappings returns a boolean if a field has been set.

### SetFieldMappingsNil

`func (o *Expense) SetFieldMappingsNil(b bool)`

 SetFieldMappingsNil sets the value for FieldMappings to be an explicit nil

### UnsetFieldMappings
`func (o *Expense) UnsetFieldMappings()`

UnsetFieldMappings ensures that no value is present for FieldMappings, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


