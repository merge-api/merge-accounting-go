# TransactionLineItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Memo** | Pointer to **NullableString** | An internal note used by the business to clarify purpose of the transaction. | [optional] 
**UnitPrice** | Pointer to **NullableFloat64** | The line item&#39;s unit price. | [optional] 
**Quantity** | Pointer to **NullableFloat64** | The line item&#39;s quantity. | [optional] 
**Item** | Pointer to **NullableString** |  | [optional] 
**Account** | Pointer to **NullableString** | The line item&#39;s account. | [optional] 
**TrackingCategory** | Pointer to **NullableString** | The line&#39;s associated tracking category. | [optional] 
**TrackingCategories** | **[]string** | The line&#39;s associated tracking categories. | 
**TotalLineAmount** | Pointer to **NullableFloat64** | The line item&#39;s total. | [optional] 
**TaxRate** | Pointer to **NullableString** | The line item&#39;s tax rate. | [optional] 
**Currency** | Pointer to [**NullableCurrencyEnum**](CurrencyEnum.md) | The line item&#39;s currency. | [optional] 
**ExchangeRate** | Pointer to **NullableFloat64** | The line item&#39;s exchange rate. | [optional] 
**Company** | Pointer to **NullableString** | The company the line belongs to. | [optional] 
**RemoteId** | Pointer to **NullableString** | The third-party API ID of the matching object. | [optional] 

## Methods

### NewTransactionLineItem

`func NewTransactionLineItem(trackingCategories []string, ) *TransactionLineItem`

NewTransactionLineItem instantiates a new TransactionLineItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionLineItemWithDefaults

`func NewTransactionLineItemWithDefaults() *TransactionLineItem`

NewTransactionLineItemWithDefaults instantiates a new TransactionLineItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMemo

`func (o *TransactionLineItem) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *TransactionLineItem) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *TransactionLineItem) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *TransactionLineItem) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### SetMemoNil

`func (o *TransactionLineItem) SetMemoNil(b bool)`

 SetMemoNil sets the value for Memo to be an explicit nil

### UnsetMemo
`func (o *TransactionLineItem) UnsetMemo()`

UnsetMemo ensures that no value is present for Memo, not even an explicit nil
### GetUnitPrice

`func (o *TransactionLineItem) GetUnitPrice() float64`

GetUnitPrice returns the UnitPrice field if non-nil, zero value otherwise.

### GetUnitPriceOk

`func (o *TransactionLineItem) GetUnitPriceOk() (*float64, bool)`

GetUnitPriceOk returns a tuple with the UnitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPrice

`func (o *TransactionLineItem) SetUnitPrice(v float64)`

SetUnitPrice sets UnitPrice field to given value.

### HasUnitPrice

`func (o *TransactionLineItem) HasUnitPrice() bool`

HasUnitPrice returns a boolean if a field has been set.

### SetUnitPriceNil

`func (o *TransactionLineItem) SetUnitPriceNil(b bool)`

 SetUnitPriceNil sets the value for UnitPrice to be an explicit nil

### UnsetUnitPrice
`func (o *TransactionLineItem) UnsetUnitPrice()`

UnsetUnitPrice ensures that no value is present for UnitPrice, not even an explicit nil
### GetQuantity

`func (o *TransactionLineItem) GetQuantity() float64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *TransactionLineItem) GetQuantityOk() (*float64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *TransactionLineItem) SetQuantity(v float64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *TransactionLineItem) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### SetQuantityNil

`func (o *TransactionLineItem) SetQuantityNil(b bool)`

 SetQuantityNil sets the value for Quantity to be an explicit nil

### UnsetQuantity
`func (o *TransactionLineItem) UnsetQuantity()`

UnsetQuantity ensures that no value is present for Quantity, not even an explicit nil
### GetItem

`func (o *TransactionLineItem) GetItem() string`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *TransactionLineItem) GetItemOk() (*string, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *TransactionLineItem) SetItem(v string)`

SetItem sets Item field to given value.

### HasItem

`func (o *TransactionLineItem) HasItem() bool`

HasItem returns a boolean if a field has been set.

### SetItemNil

`func (o *TransactionLineItem) SetItemNil(b bool)`

 SetItemNil sets the value for Item to be an explicit nil

### UnsetItem
`func (o *TransactionLineItem) UnsetItem()`

UnsetItem ensures that no value is present for Item, not even an explicit nil
### GetAccount

`func (o *TransactionLineItem) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *TransactionLineItem) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *TransactionLineItem) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *TransactionLineItem) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### SetAccountNil

`func (o *TransactionLineItem) SetAccountNil(b bool)`

 SetAccountNil sets the value for Account to be an explicit nil

### UnsetAccount
`func (o *TransactionLineItem) UnsetAccount()`

UnsetAccount ensures that no value is present for Account, not even an explicit nil
### GetTrackingCategory

`func (o *TransactionLineItem) GetTrackingCategory() string`

GetTrackingCategory returns the TrackingCategory field if non-nil, zero value otherwise.

### GetTrackingCategoryOk

`func (o *TransactionLineItem) GetTrackingCategoryOk() (*string, bool)`

GetTrackingCategoryOk returns a tuple with the TrackingCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingCategory

`func (o *TransactionLineItem) SetTrackingCategory(v string)`

SetTrackingCategory sets TrackingCategory field to given value.

### HasTrackingCategory

`func (o *TransactionLineItem) HasTrackingCategory() bool`

HasTrackingCategory returns a boolean if a field has been set.

### SetTrackingCategoryNil

`func (o *TransactionLineItem) SetTrackingCategoryNil(b bool)`

 SetTrackingCategoryNil sets the value for TrackingCategory to be an explicit nil

### UnsetTrackingCategory
`func (o *TransactionLineItem) UnsetTrackingCategory()`

UnsetTrackingCategory ensures that no value is present for TrackingCategory, not even an explicit nil
### GetTrackingCategories

`func (o *TransactionLineItem) GetTrackingCategories() []string`

GetTrackingCategories returns the TrackingCategories field if non-nil, zero value otherwise.

### GetTrackingCategoriesOk

`func (o *TransactionLineItem) GetTrackingCategoriesOk() (*[]string, bool)`

GetTrackingCategoriesOk returns a tuple with the TrackingCategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingCategories

`func (o *TransactionLineItem) SetTrackingCategories(v []string)`

SetTrackingCategories sets TrackingCategories field to given value.


### GetTotalLineAmount

`func (o *TransactionLineItem) GetTotalLineAmount() float64`

GetTotalLineAmount returns the TotalLineAmount field if non-nil, zero value otherwise.

### GetTotalLineAmountOk

`func (o *TransactionLineItem) GetTotalLineAmountOk() (*float64, bool)`

GetTotalLineAmountOk returns a tuple with the TotalLineAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalLineAmount

`func (o *TransactionLineItem) SetTotalLineAmount(v float64)`

SetTotalLineAmount sets TotalLineAmount field to given value.

### HasTotalLineAmount

`func (o *TransactionLineItem) HasTotalLineAmount() bool`

HasTotalLineAmount returns a boolean if a field has been set.

### SetTotalLineAmountNil

`func (o *TransactionLineItem) SetTotalLineAmountNil(b bool)`

 SetTotalLineAmountNil sets the value for TotalLineAmount to be an explicit nil

### UnsetTotalLineAmount
`func (o *TransactionLineItem) UnsetTotalLineAmount()`

UnsetTotalLineAmount ensures that no value is present for TotalLineAmount, not even an explicit nil
### GetTaxRate

`func (o *TransactionLineItem) GetTaxRate() string`

GetTaxRate returns the TaxRate field if non-nil, zero value otherwise.

### GetTaxRateOk

`func (o *TransactionLineItem) GetTaxRateOk() (*string, bool)`

GetTaxRateOk returns a tuple with the TaxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRate

`func (o *TransactionLineItem) SetTaxRate(v string)`

SetTaxRate sets TaxRate field to given value.

### HasTaxRate

`func (o *TransactionLineItem) HasTaxRate() bool`

HasTaxRate returns a boolean if a field has been set.

### SetTaxRateNil

`func (o *TransactionLineItem) SetTaxRateNil(b bool)`

 SetTaxRateNil sets the value for TaxRate to be an explicit nil

### UnsetTaxRate
`func (o *TransactionLineItem) UnsetTaxRate()`

UnsetTaxRate ensures that no value is present for TaxRate, not even an explicit nil
### GetCurrency

`func (o *TransactionLineItem) GetCurrency() CurrencyEnum`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *TransactionLineItem) GetCurrencyOk() (*CurrencyEnum, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *TransactionLineItem) SetCurrency(v CurrencyEnum)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *TransactionLineItem) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### SetCurrencyNil

`func (o *TransactionLineItem) SetCurrencyNil(b bool)`

 SetCurrencyNil sets the value for Currency to be an explicit nil

### UnsetCurrency
`func (o *TransactionLineItem) UnsetCurrency()`

UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
### GetExchangeRate

`func (o *TransactionLineItem) GetExchangeRate() float64`

GetExchangeRate returns the ExchangeRate field if non-nil, zero value otherwise.

### GetExchangeRateOk

`func (o *TransactionLineItem) GetExchangeRateOk() (*float64, bool)`

GetExchangeRateOk returns a tuple with the ExchangeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeRate

`func (o *TransactionLineItem) SetExchangeRate(v float64)`

SetExchangeRate sets ExchangeRate field to given value.

### HasExchangeRate

`func (o *TransactionLineItem) HasExchangeRate() bool`

HasExchangeRate returns a boolean if a field has been set.

### SetExchangeRateNil

`func (o *TransactionLineItem) SetExchangeRateNil(b bool)`

 SetExchangeRateNil sets the value for ExchangeRate to be an explicit nil

### UnsetExchangeRate
`func (o *TransactionLineItem) UnsetExchangeRate()`

UnsetExchangeRate ensures that no value is present for ExchangeRate, not even an explicit nil
### GetCompany

`func (o *TransactionLineItem) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *TransactionLineItem) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *TransactionLineItem) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *TransactionLineItem) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### SetCompanyNil

`func (o *TransactionLineItem) SetCompanyNil(b bool)`

 SetCompanyNil sets the value for Company to be an explicit nil

### UnsetCompany
`func (o *TransactionLineItem) UnsetCompany()`

UnsetCompany ensures that no value is present for Company, not even an explicit nil
### GetRemoteId

`func (o *TransactionLineItem) GetRemoteId() string`

GetRemoteId returns the RemoteId field if non-nil, zero value otherwise.

### GetRemoteIdOk

`func (o *TransactionLineItem) GetRemoteIdOk() (*string, bool)`

GetRemoteIdOk returns a tuple with the RemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteId

`func (o *TransactionLineItem) SetRemoteId(v string)`

SetRemoteId sets RemoteId field to given value.

### HasRemoteId

`func (o *TransactionLineItem) HasRemoteId() bool`

HasRemoteId returns a boolean if a field has been set.

### SetRemoteIdNil

`func (o *TransactionLineItem) SetRemoteIdNil(b bool)`

 SetRemoteIdNil sets the value for RemoteId to be an explicit nil

### UnsetRemoteId
`func (o *TransactionLineItem) UnsetRemoteId()`

UnsetRemoteId ensures that no value is present for RemoteId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


