# CreditNoteLineItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Item** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** | The credit note line item&#39;s name. | [optional] 
**Description** | Pointer to **NullableString** | The description of the item that is owed. | [optional] 
**Quantity** | Pointer to **NullableFloat64** | The credit note line item&#39;s quantity. | [optional] 
**Memo** | Pointer to **NullableString** | The credit note line item&#39;s memo. | [optional] 
**UnitPrice** | Pointer to **NullableFloat64** | The credit note line item&#39;s unit price. | [optional] 
**TaxRate** | Pointer to **NullableString** | The credit note line item&#39;s tax rate. | [optional] 
**TotalLineAmount** | Pointer to **NullableFloat64** | The credit note line item&#39;s total. | [optional] 
**TrackingCategory** | Pointer to **NullableString** | The credit note line item&#39;s associated tracking category. | [optional] 
**TrackingCategories** | **[]string** | The credit note line item&#39;s associated tracking categories. | 
**Account** | Pointer to **NullableString** | The credit note line item&#39;s account. | [optional] 
**Company** | Pointer to **NullableString** | The company the credit note line item belongs to. | [optional] 
**RemoteId** | Pointer to **NullableString** | The third-party API ID of the matching object. | [optional] 

## Methods

### NewCreditNoteLineItem

`func NewCreditNoteLineItem(trackingCategories []string, ) *CreditNoteLineItem`

NewCreditNoteLineItem instantiates a new CreditNoteLineItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreditNoteLineItemWithDefaults

`func NewCreditNoteLineItemWithDefaults() *CreditNoteLineItem`

NewCreditNoteLineItemWithDefaults instantiates a new CreditNoteLineItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItem

`func (o *CreditNoteLineItem) GetItem() string`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *CreditNoteLineItem) GetItemOk() (*string, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *CreditNoteLineItem) SetItem(v string)`

SetItem sets Item field to given value.

### HasItem

`func (o *CreditNoteLineItem) HasItem() bool`

HasItem returns a boolean if a field has been set.

### SetItemNil

`func (o *CreditNoteLineItem) SetItemNil(b bool)`

 SetItemNil sets the value for Item to be an explicit nil

### UnsetItem
`func (o *CreditNoteLineItem) UnsetItem()`

UnsetItem ensures that no value is present for Item, not even an explicit nil
### GetName

`func (o *CreditNoteLineItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreditNoteLineItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreditNoteLineItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreditNoteLineItem) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CreditNoteLineItem) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CreditNoteLineItem) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *CreditNoteLineItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreditNoteLineItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreditNoteLineItem) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreditNoteLineItem) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CreditNoteLineItem) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CreditNoteLineItem) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetQuantity

`func (o *CreditNoteLineItem) GetQuantity() float64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *CreditNoteLineItem) GetQuantityOk() (*float64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *CreditNoteLineItem) SetQuantity(v float64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *CreditNoteLineItem) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### SetQuantityNil

`func (o *CreditNoteLineItem) SetQuantityNil(b bool)`

 SetQuantityNil sets the value for Quantity to be an explicit nil

### UnsetQuantity
`func (o *CreditNoteLineItem) UnsetQuantity()`

UnsetQuantity ensures that no value is present for Quantity, not even an explicit nil
### GetMemo

`func (o *CreditNoteLineItem) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *CreditNoteLineItem) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *CreditNoteLineItem) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *CreditNoteLineItem) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### SetMemoNil

`func (o *CreditNoteLineItem) SetMemoNil(b bool)`

 SetMemoNil sets the value for Memo to be an explicit nil

### UnsetMemo
`func (o *CreditNoteLineItem) UnsetMemo()`

UnsetMemo ensures that no value is present for Memo, not even an explicit nil
### GetUnitPrice

`func (o *CreditNoteLineItem) GetUnitPrice() float64`

GetUnitPrice returns the UnitPrice field if non-nil, zero value otherwise.

### GetUnitPriceOk

`func (o *CreditNoteLineItem) GetUnitPriceOk() (*float64, bool)`

GetUnitPriceOk returns a tuple with the UnitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPrice

`func (o *CreditNoteLineItem) SetUnitPrice(v float64)`

SetUnitPrice sets UnitPrice field to given value.

### HasUnitPrice

`func (o *CreditNoteLineItem) HasUnitPrice() bool`

HasUnitPrice returns a boolean if a field has been set.

### SetUnitPriceNil

`func (o *CreditNoteLineItem) SetUnitPriceNil(b bool)`

 SetUnitPriceNil sets the value for UnitPrice to be an explicit nil

### UnsetUnitPrice
`func (o *CreditNoteLineItem) UnsetUnitPrice()`

UnsetUnitPrice ensures that no value is present for UnitPrice, not even an explicit nil
### GetTaxRate

`func (o *CreditNoteLineItem) GetTaxRate() string`

GetTaxRate returns the TaxRate field if non-nil, zero value otherwise.

### GetTaxRateOk

`func (o *CreditNoteLineItem) GetTaxRateOk() (*string, bool)`

GetTaxRateOk returns a tuple with the TaxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRate

`func (o *CreditNoteLineItem) SetTaxRate(v string)`

SetTaxRate sets TaxRate field to given value.

### HasTaxRate

`func (o *CreditNoteLineItem) HasTaxRate() bool`

HasTaxRate returns a boolean if a field has been set.

### SetTaxRateNil

`func (o *CreditNoteLineItem) SetTaxRateNil(b bool)`

 SetTaxRateNil sets the value for TaxRate to be an explicit nil

### UnsetTaxRate
`func (o *CreditNoteLineItem) UnsetTaxRate()`

UnsetTaxRate ensures that no value is present for TaxRate, not even an explicit nil
### GetTotalLineAmount

`func (o *CreditNoteLineItem) GetTotalLineAmount() float64`

GetTotalLineAmount returns the TotalLineAmount field if non-nil, zero value otherwise.

### GetTotalLineAmountOk

`func (o *CreditNoteLineItem) GetTotalLineAmountOk() (*float64, bool)`

GetTotalLineAmountOk returns a tuple with the TotalLineAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalLineAmount

`func (o *CreditNoteLineItem) SetTotalLineAmount(v float64)`

SetTotalLineAmount sets TotalLineAmount field to given value.

### HasTotalLineAmount

`func (o *CreditNoteLineItem) HasTotalLineAmount() bool`

HasTotalLineAmount returns a boolean if a field has been set.

### SetTotalLineAmountNil

`func (o *CreditNoteLineItem) SetTotalLineAmountNil(b bool)`

 SetTotalLineAmountNil sets the value for TotalLineAmount to be an explicit nil

### UnsetTotalLineAmount
`func (o *CreditNoteLineItem) UnsetTotalLineAmount()`

UnsetTotalLineAmount ensures that no value is present for TotalLineAmount, not even an explicit nil
### GetTrackingCategory

`func (o *CreditNoteLineItem) GetTrackingCategory() string`

GetTrackingCategory returns the TrackingCategory field if non-nil, zero value otherwise.

### GetTrackingCategoryOk

`func (o *CreditNoteLineItem) GetTrackingCategoryOk() (*string, bool)`

GetTrackingCategoryOk returns a tuple with the TrackingCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingCategory

`func (o *CreditNoteLineItem) SetTrackingCategory(v string)`

SetTrackingCategory sets TrackingCategory field to given value.

### HasTrackingCategory

`func (o *CreditNoteLineItem) HasTrackingCategory() bool`

HasTrackingCategory returns a boolean if a field has been set.

### SetTrackingCategoryNil

`func (o *CreditNoteLineItem) SetTrackingCategoryNil(b bool)`

 SetTrackingCategoryNil sets the value for TrackingCategory to be an explicit nil

### UnsetTrackingCategory
`func (o *CreditNoteLineItem) UnsetTrackingCategory()`

UnsetTrackingCategory ensures that no value is present for TrackingCategory, not even an explicit nil
### GetTrackingCategories

`func (o *CreditNoteLineItem) GetTrackingCategories() []string`

GetTrackingCategories returns the TrackingCategories field if non-nil, zero value otherwise.

### GetTrackingCategoriesOk

`func (o *CreditNoteLineItem) GetTrackingCategoriesOk() (*[]string, bool)`

GetTrackingCategoriesOk returns a tuple with the TrackingCategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingCategories

`func (o *CreditNoteLineItem) SetTrackingCategories(v []string)`

SetTrackingCategories sets TrackingCategories field to given value.


### GetAccount

`func (o *CreditNoteLineItem) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *CreditNoteLineItem) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *CreditNoteLineItem) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *CreditNoteLineItem) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### SetAccountNil

`func (o *CreditNoteLineItem) SetAccountNil(b bool)`

 SetAccountNil sets the value for Account to be an explicit nil

### UnsetAccount
`func (o *CreditNoteLineItem) UnsetAccount()`

UnsetAccount ensures that no value is present for Account, not even an explicit nil
### GetCompany

`func (o *CreditNoteLineItem) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *CreditNoteLineItem) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *CreditNoteLineItem) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *CreditNoteLineItem) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### SetCompanyNil

`func (o *CreditNoteLineItem) SetCompanyNil(b bool)`

 SetCompanyNil sets the value for Company to be an explicit nil

### UnsetCompany
`func (o *CreditNoteLineItem) UnsetCompany()`

UnsetCompany ensures that no value is present for Company, not even an explicit nil
### GetRemoteId

`func (o *CreditNoteLineItem) GetRemoteId() string`

GetRemoteId returns the RemoteId field if non-nil, zero value otherwise.

### GetRemoteIdOk

`func (o *CreditNoteLineItem) GetRemoteIdOk() (*string, bool)`

GetRemoteIdOk returns a tuple with the RemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteId

`func (o *CreditNoteLineItem) SetRemoteId(v string)`

SetRemoteId sets RemoteId field to given value.

### HasRemoteId

`func (o *CreditNoteLineItem) HasRemoteId() bool`

HasRemoteId returns a boolean if a field has been set.

### SetRemoteIdNil

`func (o *CreditNoteLineItem) SetRemoteIdNil(b bool)`

 SetRemoteIdNil sets the value for RemoteId to be an explicit nil

### UnsetRemoteId
`func (o *CreditNoteLineItem) UnsetRemoteId()`

UnsetRemoteId ensures that no value is present for RemoteId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


