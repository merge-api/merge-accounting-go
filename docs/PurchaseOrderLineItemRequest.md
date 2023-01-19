# PurchaseOrderLineItemRequest

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
**IntegrationParams** | Pointer to **map[string]interface{}** |  | [optional] 
**LinkedAccountParams** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewPurchaseOrderLineItemRequest

`func NewPurchaseOrderLineItemRequest(trackingCategories []string, ) *PurchaseOrderLineItemRequest`

NewPurchaseOrderLineItemRequest instantiates a new PurchaseOrderLineItemRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPurchaseOrderLineItemRequestWithDefaults

`func NewPurchaseOrderLineItemRequestWithDefaults() *PurchaseOrderLineItemRequest`

NewPurchaseOrderLineItemRequestWithDefaults instantiates a new PurchaseOrderLineItemRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *PurchaseOrderLineItemRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PurchaseOrderLineItemRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PurchaseOrderLineItemRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PurchaseOrderLineItemRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *PurchaseOrderLineItemRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PurchaseOrderLineItemRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetUnitPrice

`func (o *PurchaseOrderLineItemRequest) GetUnitPrice() float32`

GetUnitPrice returns the UnitPrice field if non-nil, zero value otherwise.

### GetUnitPriceOk

`func (o *PurchaseOrderLineItemRequest) GetUnitPriceOk() (*float32, bool)`

GetUnitPriceOk returns a tuple with the UnitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPrice

`func (o *PurchaseOrderLineItemRequest) SetUnitPrice(v float32)`

SetUnitPrice sets UnitPrice field to given value.

### HasUnitPrice

`func (o *PurchaseOrderLineItemRequest) HasUnitPrice() bool`

HasUnitPrice returns a boolean if a field has been set.

### SetUnitPriceNil

`func (o *PurchaseOrderLineItemRequest) SetUnitPriceNil(b bool)`

 SetUnitPriceNil sets the value for UnitPrice to be an explicit nil

### UnsetUnitPrice
`func (o *PurchaseOrderLineItemRequest) UnsetUnitPrice()`

UnsetUnitPrice ensures that no value is present for UnitPrice, not even an explicit nil
### GetQuantity

`func (o *PurchaseOrderLineItemRequest) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *PurchaseOrderLineItemRequest) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *PurchaseOrderLineItemRequest) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *PurchaseOrderLineItemRequest) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### SetQuantityNil

`func (o *PurchaseOrderLineItemRequest) SetQuantityNil(b bool)`

 SetQuantityNil sets the value for Quantity to be an explicit nil

### UnsetQuantity
`func (o *PurchaseOrderLineItemRequest) UnsetQuantity()`

UnsetQuantity ensures that no value is present for Quantity, not even an explicit nil
### GetItem

`func (o *PurchaseOrderLineItemRequest) GetItem() string`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *PurchaseOrderLineItemRequest) GetItemOk() (*string, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *PurchaseOrderLineItemRequest) SetItem(v string)`

SetItem sets Item field to given value.

### HasItem

`func (o *PurchaseOrderLineItemRequest) HasItem() bool`

HasItem returns a boolean if a field has been set.

### SetItemNil

`func (o *PurchaseOrderLineItemRequest) SetItemNil(b bool)`

 SetItemNil sets the value for Item to be an explicit nil

### UnsetItem
`func (o *PurchaseOrderLineItemRequest) UnsetItem()`

UnsetItem ensures that no value is present for Item, not even an explicit nil
### GetAccount

`func (o *PurchaseOrderLineItemRequest) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *PurchaseOrderLineItemRequest) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *PurchaseOrderLineItemRequest) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *PurchaseOrderLineItemRequest) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### SetAccountNil

`func (o *PurchaseOrderLineItemRequest) SetAccountNil(b bool)`

 SetAccountNil sets the value for Account to be an explicit nil

### UnsetAccount
`func (o *PurchaseOrderLineItemRequest) UnsetAccount()`

UnsetAccount ensures that no value is present for Account, not even an explicit nil
### GetTrackingCategory

`func (o *PurchaseOrderLineItemRequest) GetTrackingCategory() string`

GetTrackingCategory returns the TrackingCategory field if non-nil, zero value otherwise.

### GetTrackingCategoryOk

`func (o *PurchaseOrderLineItemRequest) GetTrackingCategoryOk() (*string, bool)`

GetTrackingCategoryOk returns a tuple with the TrackingCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingCategory

`func (o *PurchaseOrderLineItemRequest) SetTrackingCategory(v string)`

SetTrackingCategory sets TrackingCategory field to given value.

### HasTrackingCategory

`func (o *PurchaseOrderLineItemRequest) HasTrackingCategory() bool`

HasTrackingCategory returns a boolean if a field has been set.

### SetTrackingCategoryNil

`func (o *PurchaseOrderLineItemRequest) SetTrackingCategoryNil(b bool)`

 SetTrackingCategoryNil sets the value for TrackingCategory to be an explicit nil

### UnsetTrackingCategory
`func (o *PurchaseOrderLineItemRequest) UnsetTrackingCategory()`

UnsetTrackingCategory ensures that no value is present for TrackingCategory, not even an explicit nil
### GetTrackingCategories

`func (o *PurchaseOrderLineItemRequest) GetTrackingCategories() []string`

GetTrackingCategories returns the TrackingCategories field if non-nil, zero value otherwise.

### GetTrackingCategoriesOk

`func (o *PurchaseOrderLineItemRequest) GetTrackingCategoriesOk() (*[]string, bool)`

GetTrackingCategoriesOk returns a tuple with the TrackingCategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingCategories

`func (o *PurchaseOrderLineItemRequest) SetTrackingCategories(v []string)`

SetTrackingCategories sets TrackingCategories field to given value.


### GetTaxAmount

`func (o *PurchaseOrderLineItemRequest) GetTaxAmount() float64`

GetTaxAmount returns the TaxAmount field if non-nil, zero value otherwise.

### GetTaxAmountOk

`func (o *PurchaseOrderLineItemRequest) GetTaxAmountOk() (*float64, bool)`

GetTaxAmountOk returns a tuple with the TaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxAmount

`func (o *PurchaseOrderLineItemRequest) SetTaxAmount(v float64)`

SetTaxAmount sets TaxAmount field to given value.

### HasTaxAmount

`func (o *PurchaseOrderLineItemRequest) HasTaxAmount() bool`

HasTaxAmount returns a boolean if a field has been set.

### SetTaxAmountNil

`func (o *PurchaseOrderLineItemRequest) SetTaxAmountNil(b bool)`

 SetTaxAmountNil sets the value for TaxAmount to be an explicit nil

### UnsetTaxAmount
`func (o *PurchaseOrderLineItemRequest) UnsetTaxAmount()`

UnsetTaxAmount ensures that no value is present for TaxAmount, not even an explicit nil
### GetTotalLineAmount

`func (o *PurchaseOrderLineItemRequest) GetTotalLineAmount() float64`

GetTotalLineAmount returns the TotalLineAmount field if non-nil, zero value otherwise.

### GetTotalLineAmountOk

`func (o *PurchaseOrderLineItemRequest) GetTotalLineAmountOk() (*float64, bool)`

GetTotalLineAmountOk returns a tuple with the TotalLineAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalLineAmount

`func (o *PurchaseOrderLineItemRequest) SetTotalLineAmount(v float64)`

SetTotalLineAmount sets TotalLineAmount field to given value.

### HasTotalLineAmount

`func (o *PurchaseOrderLineItemRequest) HasTotalLineAmount() bool`

HasTotalLineAmount returns a boolean if a field has been set.

### SetTotalLineAmountNil

`func (o *PurchaseOrderLineItemRequest) SetTotalLineAmountNil(b bool)`

 SetTotalLineAmountNil sets the value for TotalLineAmount to be an explicit nil

### UnsetTotalLineAmount
`func (o *PurchaseOrderLineItemRequest) UnsetTotalLineAmount()`

UnsetTotalLineAmount ensures that no value is present for TotalLineAmount, not even an explicit nil
### GetCurrency

`func (o *PurchaseOrderLineItemRequest) GetCurrency() CurrencyEnum`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *PurchaseOrderLineItemRequest) GetCurrencyOk() (*CurrencyEnum, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *PurchaseOrderLineItemRequest) SetCurrency(v CurrencyEnum)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *PurchaseOrderLineItemRequest) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### SetCurrencyNil

`func (o *PurchaseOrderLineItemRequest) SetCurrencyNil(b bool)`

 SetCurrencyNil sets the value for Currency to be an explicit nil

### UnsetCurrency
`func (o *PurchaseOrderLineItemRequest) UnsetCurrency()`

UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
### GetExchangeRate

`func (o *PurchaseOrderLineItemRequest) GetExchangeRate() float64`

GetExchangeRate returns the ExchangeRate field if non-nil, zero value otherwise.

### GetExchangeRateOk

`func (o *PurchaseOrderLineItemRequest) GetExchangeRateOk() (*float64, bool)`

GetExchangeRateOk returns a tuple with the ExchangeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeRate

`func (o *PurchaseOrderLineItemRequest) SetExchangeRate(v float64)`

SetExchangeRate sets ExchangeRate field to given value.

### HasExchangeRate

`func (o *PurchaseOrderLineItemRequest) HasExchangeRate() bool`

HasExchangeRate returns a boolean if a field has been set.

### SetExchangeRateNil

`func (o *PurchaseOrderLineItemRequest) SetExchangeRateNil(b bool)`

 SetExchangeRateNil sets the value for ExchangeRate to be an explicit nil

### UnsetExchangeRate
`func (o *PurchaseOrderLineItemRequest) UnsetExchangeRate()`

UnsetExchangeRate ensures that no value is present for ExchangeRate, not even an explicit nil
### GetCompany

`func (o *PurchaseOrderLineItemRequest) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *PurchaseOrderLineItemRequest) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *PurchaseOrderLineItemRequest) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *PurchaseOrderLineItemRequest) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### SetCompanyNil

`func (o *PurchaseOrderLineItemRequest) SetCompanyNil(b bool)`

 SetCompanyNil sets the value for Company to be an explicit nil

### UnsetCompany
`func (o *PurchaseOrderLineItemRequest) UnsetCompany()`

UnsetCompany ensures that no value is present for Company, not even an explicit nil
### GetIntegrationParams

`func (o *PurchaseOrderLineItemRequest) GetIntegrationParams() map[string]interface{}`

GetIntegrationParams returns the IntegrationParams field if non-nil, zero value otherwise.

### GetIntegrationParamsOk

`func (o *PurchaseOrderLineItemRequest) GetIntegrationParamsOk() (*map[string]interface{}, bool)`

GetIntegrationParamsOk returns a tuple with the IntegrationParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationParams

`func (o *PurchaseOrderLineItemRequest) SetIntegrationParams(v map[string]interface{})`

SetIntegrationParams sets IntegrationParams field to given value.

### HasIntegrationParams

`func (o *PurchaseOrderLineItemRequest) HasIntegrationParams() bool`

HasIntegrationParams returns a boolean if a field has been set.

### SetIntegrationParamsNil

`func (o *PurchaseOrderLineItemRequest) SetIntegrationParamsNil(b bool)`

 SetIntegrationParamsNil sets the value for IntegrationParams to be an explicit nil

### UnsetIntegrationParams
`func (o *PurchaseOrderLineItemRequest) UnsetIntegrationParams()`

UnsetIntegrationParams ensures that no value is present for IntegrationParams, not even an explicit nil
### GetLinkedAccountParams

`func (o *PurchaseOrderLineItemRequest) GetLinkedAccountParams() map[string]interface{}`

GetLinkedAccountParams returns the LinkedAccountParams field if non-nil, zero value otherwise.

### GetLinkedAccountParamsOk

`func (o *PurchaseOrderLineItemRequest) GetLinkedAccountParamsOk() (*map[string]interface{}, bool)`

GetLinkedAccountParamsOk returns a tuple with the LinkedAccountParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedAccountParams

`func (o *PurchaseOrderLineItemRequest) SetLinkedAccountParams(v map[string]interface{})`

SetLinkedAccountParams sets LinkedAccountParams field to given value.

### HasLinkedAccountParams

`func (o *PurchaseOrderLineItemRequest) HasLinkedAccountParams() bool`

HasLinkedAccountParams returns a boolean if a field has been set.

### SetLinkedAccountParamsNil

`func (o *PurchaseOrderLineItemRequest) SetLinkedAccountParamsNil(b bool)`

 SetLinkedAccountParamsNil sets the value for LinkedAccountParams to be an explicit nil

### UnsetLinkedAccountParams
`func (o *PurchaseOrderLineItemRequest) UnsetLinkedAccountParams()`

UnsetLinkedAccountParams ensures that no value is present for LinkedAccountParams, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


