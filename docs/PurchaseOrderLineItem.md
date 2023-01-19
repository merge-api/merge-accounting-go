# PurchaseOrderLineItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** | A description of the good being purchased. | [optional] 
**UnitPrice** | Pointer to **NullableFloat32** | The line item&#39;s unit price. | [optional] 
**Quantity** | Pointer to **NullableFloat32** | The line item&#39;s quantity. | [optional] 
**Item** | Pointer to **NullableString** |  | [optional] 
**Account** | Pointer to **NullableString** | The purchase order line item&#39;s account. | [optional] 
**TrackingCategory** | Pointer to **NullableString** | The purchase order line item&#39;s associated tracking category. | [optional] 
**TrackingCategories** | **[]string** | The purchase order line item&#39;s associated tracking categories. | 
**TaxAmount** | Pointer to **NullableFloat64** | The purchase order line item&#39;s tax amount. | [optional] 
**TotalLineAmount** | Pointer to **NullableFloat64** | The purchase order line item&#39;s total amount. | [optional] 
**Currency** | Pointer to [**NullableCurrencyEnum**](CurrencyEnum.md) | The purchase order line item&#39;s currency. | [optional] 
**ExchangeRate** | Pointer to **NullableFloat64** | The purchase order line item&#39;s exchange rate. | [optional] 
**Company** | Pointer to **NullableString** | The company the purchase order line item belongs to. | [optional] 

## Methods

### NewPurchaseOrderLineItem

`func NewPurchaseOrderLineItem(trackingCategories []string, ) *PurchaseOrderLineItem`

NewPurchaseOrderLineItem instantiates a new PurchaseOrderLineItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPurchaseOrderLineItemWithDefaults

`func NewPurchaseOrderLineItemWithDefaults() *PurchaseOrderLineItem`

NewPurchaseOrderLineItemWithDefaults instantiates a new PurchaseOrderLineItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *PurchaseOrderLineItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PurchaseOrderLineItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PurchaseOrderLineItem) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PurchaseOrderLineItem) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *PurchaseOrderLineItem) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PurchaseOrderLineItem) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetUnitPrice

`func (o *PurchaseOrderLineItem) GetUnitPrice() float32`

GetUnitPrice returns the UnitPrice field if non-nil, zero value otherwise.

### GetUnitPriceOk

`func (o *PurchaseOrderLineItem) GetUnitPriceOk() (*float32, bool)`

GetUnitPriceOk returns a tuple with the UnitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPrice

`func (o *PurchaseOrderLineItem) SetUnitPrice(v float32)`

SetUnitPrice sets UnitPrice field to given value.

### HasUnitPrice

`func (o *PurchaseOrderLineItem) HasUnitPrice() bool`

HasUnitPrice returns a boolean if a field has been set.

### SetUnitPriceNil

`func (o *PurchaseOrderLineItem) SetUnitPriceNil(b bool)`

 SetUnitPriceNil sets the value for UnitPrice to be an explicit nil

### UnsetUnitPrice
`func (o *PurchaseOrderLineItem) UnsetUnitPrice()`

UnsetUnitPrice ensures that no value is present for UnitPrice, not even an explicit nil
### GetQuantity

`func (o *PurchaseOrderLineItem) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *PurchaseOrderLineItem) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *PurchaseOrderLineItem) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *PurchaseOrderLineItem) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### SetQuantityNil

`func (o *PurchaseOrderLineItem) SetQuantityNil(b bool)`

 SetQuantityNil sets the value for Quantity to be an explicit nil

### UnsetQuantity
`func (o *PurchaseOrderLineItem) UnsetQuantity()`

UnsetQuantity ensures that no value is present for Quantity, not even an explicit nil
### GetItem

`func (o *PurchaseOrderLineItem) GetItem() string`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *PurchaseOrderLineItem) GetItemOk() (*string, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *PurchaseOrderLineItem) SetItem(v string)`

SetItem sets Item field to given value.

### HasItem

`func (o *PurchaseOrderLineItem) HasItem() bool`

HasItem returns a boolean if a field has been set.

### SetItemNil

`func (o *PurchaseOrderLineItem) SetItemNil(b bool)`

 SetItemNil sets the value for Item to be an explicit nil

### UnsetItem
`func (o *PurchaseOrderLineItem) UnsetItem()`

UnsetItem ensures that no value is present for Item, not even an explicit nil
### GetAccount

`func (o *PurchaseOrderLineItem) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *PurchaseOrderLineItem) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *PurchaseOrderLineItem) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *PurchaseOrderLineItem) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### SetAccountNil

`func (o *PurchaseOrderLineItem) SetAccountNil(b bool)`

 SetAccountNil sets the value for Account to be an explicit nil

### UnsetAccount
`func (o *PurchaseOrderLineItem) UnsetAccount()`

UnsetAccount ensures that no value is present for Account, not even an explicit nil
### GetTrackingCategory

`func (o *PurchaseOrderLineItem) GetTrackingCategory() string`

GetTrackingCategory returns the TrackingCategory field if non-nil, zero value otherwise.

### GetTrackingCategoryOk

`func (o *PurchaseOrderLineItem) GetTrackingCategoryOk() (*string, bool)`

GetTrackingCategoryOk returns a tuple with the TrackingCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingCategory

`func (o *PurchaseOrderLineItem) SetTrackingCategory(v string)`

SetTrackingCategory sets TrackingCategory field to given value.

### HasTrackingCategory

`func (o *PurchaseOrderLineItem) HasTrackingCategory() bool`

HasTrackingCategory returns a boolean if a field has been set.

### SetTrackingCategoryNil

`func (o *PurchaseOrderLineItem) SetTrackingCategoryNil(b bool)`

 SetTrackingCategoryNil sets the value for TrackingCategory to be an explicit nil

### UnsetTrackingCategory
`func (o *PurchaseOrderLineItem) UnsetTrackingCategory()`

UnsetTrackingCategory ensures that no value is present for TrackingCategory, not even an explicit nil
### GetTrackingCategories

`func (o *PurchaseOrderLineItem) GetTrackingCategories() []string`

GetTrackingCategories returns the TrackingCategories field if non-nil, zero value otherwise.

### GetTrackingCategoriesOk

`func (o *PurchaseOrderLineItem) GetTrackingCategoriesOk() (*[]string, bool)`

GetTrackingCategoriesOk returns a tuple with the TrackingCategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingCategories

`func (o *PurchaseOrderLineItem) SetTrackingCategories(v []string)`

SetTrackingCategories sets TrackingCategories field to given value.


### GetTaxAmount

`func (o *PurchaseOrderLineItem) GetTaxAmount() float64`

GetTaxAmount returns the TaxAmount field if non-nil, zero value otherwise.

### GetTaxAmountOk

`func (o *PurchaseOrderLineItem) GetTaxAmountOk() (*float64, bool)`

GetTaxAmountOk returns a tuple with the TaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxAmount

`func (o *PurchaseOrderLineItem) SetTaxAmount(v float64)`

SetTaxAmount sets TaxAmount field to given value.

### HasTaxAmount

`func (o *PurchaseOrderLineItem) HasTaxAmount() bool`

HasTaxAmount returns a boolean if a field has been set.

### SetTaxAmountNil

`func (o *PurchaseOrderLineItem) SetTaxAmountNil(b bool)`

 SetTaxAmountNil sets the value for TaxAmount to be an explicit nil

### UnsetTaxAmount
`func (o *PurchaseOrderLineItem) UnsetTaxAmount()`

UnsetTaxAmount ensures that no value is present for TaxAmount, not even an explicit nil
### GetTotalLineAmount

`func (o *PurchaseOrderLineItem) GetTotalLineAmount() float64`

GetTotalLineAmount returns the TotalLineAmount field if non-nil, zero value otherwise.

### GetTotalLineAmountOk

`func (o *PurchaseOrderLineItem) GetTotalLineAmountOk() (*float64, bool)`

GetTotalLineAmountOk returns a tuple with the TotalLineAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalLineAmount

`func (o *PurchaseOrderLineItem) SetTotalLineAmount(v float64)`

SetTotalLineAmount sets TotalLineAmount field to given value.

### HasTotalLineAmount

`func (o *PurchaseOrderLineItem) HasTotalLineAmount() bool`

HasTotalLineAmount returns a boolean if a field has been set.

### SetTotalLineAmountNil

`func (o *PurchaseOrderLineItem) SetTotalLineAmountNil(b bool)`

 SetTotalLineAmountNil sets the value for TotalLineAmount to be an explicit nil

### UnsetTotalLineAmount
`func (o *PurchaseOrderLineItem) UnsetTotalLineAmount()`

UnsetTotalLineAmount ensures that no value is present for TotalLineAmount, not even an explicit nil
### GetCurrency

`func (o *PurchaseOrderLineItem) GetCurrency() CurrencyEnum`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *PurchaseOrderLineItem) GetCurrencyOk() (*CurrencyEnum, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *PurchaseOrderLineItem) SetCurrency(v CurrencyEnum)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *PurchaseOrderLineItem) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### SetCurrencyNil

`func (o *PurchaseOrderLineItem) SetCurrencyNil(b bool)`

 SetCurrencyNil sets the value for Currency to be an explicit nil

### UnsetCurrency
`func (o *PurchaseOrderLineItem) UnsetCurrency()`

UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
### GetExchangeRate

`func (o *PurchaseOrderLineItem) GetExchangeRate() float64`

GetExchangeRate returns the ExchangeRate field if non-nil, zero value otherwise.

### GetExchangeRateOk

`func (o *PurchaseOrderLineItem) GetExchangeRateOk() (*float64, bool)`

GetExchangeRateOk returns a tuple with the ExchangeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeRate

`func (o *PurchaseOrderLineItem) SetExchangeRate(v float64)`

SetExchangeRate sets ExchangeRate field to given value.

### HasExchangeRate

`func (o *PurchaseOrderLineItem) HasExchangeRate() bool`

HasExchangeRate returns a boolean if a field has been set.

### SetExchangeRateNil

`func (o *PurchaseOrderLineItem) SetExchangeRateNil(b bool)`

 SetExchangeRateNil sets the value for ExchangeRate to be an explicit nil

### UnsetExchangeRate
`func (o *PurchaseOrderLineItem) UnsetExchangeRate()`

UnsetExchangeRate ensures that no value is present for ExchangeRate, not even an explicit nil
### GetCompany

`func (o *PurchaseOrderLineItem) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *PurchaseOrderLineItem) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *PurchaseOrderLineItem) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *PurchaseOrderLineItem) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### SetCompanyNil

`func (o *PurchaseOrderLineItem) SetCompanyNil(b bool)`

 SetCompanyNil sets the value for Company to be an explicit nil

### UnsetCompany
`func (o *PurchaseOrderLineItem) UnsetCompany()`

UnsetCompany ensures that no value is present for Company, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


