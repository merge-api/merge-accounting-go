# Item

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**RemoteId** | Pointer to **NullableString** | The third-party API ID of the matching object. | [optional] 
**Name** | Pointer to **NullableString** | The item&#39;s name. | [optional] 
**Status** | Pointer to [**NullableStatus7d1Enum**](Status7d1Enum.md) | The item&#39;s status.  * &#x60;ACTIVE&#x60; - ACTIVE * &#x60;ARCHIVED&#x60; - ARCHIVED | [optional] 
**UnitPrice** | Pointer to **NullableFloat64** | The item&#39;s unit price. | [optional] 
**PurchasePrice** | Pointer to **NullableFloat64** | The price at which the item is purchased from a vendor. | [optional] 
**PurchaseAccount** | Pointer to **NullableString** | References the default account used to record a purchase of the item. | [optional] 
**SalesAccount** | Pointer to **NullableString** | References the default account used to record a sale. | [optional] 
**Company** | Pointer to **NullableString** | The company the item belongs to. | [optional] 
**RemoteUpdatedAt** | Pointer to **NullableTime** | When the third party&#39;s item note was updated. | [optional] 
**RemoteWasDeleted** | Pointer to **bool** | Indicates whether or not this object has been deleted by third party webhooks. | [optional] [readonly] 
**FieldMappings** | Pointer to **map[string]interface{}** |  | [optional] [readonly] 
**ModifiedAt** | Pointer to **time.Time** | This is the datetime that this object was last updated by Merge | [optional] [readonly] 
**RemoteData** | Pointer to [**[]RemoteData**](RemoteData.md) |  | [optional] [readonly] 

## Methods

### NewItem

`func NewItem() *Item`

NewItem instantiates a new Item object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemWithDefaults

`func NewItemWithDefaults() *Item`

NewItemWithDefaults instantiates a new Item object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Item) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Item) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Item) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Item) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRemoteId

`func (o *Item) GetRemoteId() string`

GetRemoteId returns the RemoteId field if non-nil, zero value otherwise.

### GetRemoteIdOk

`func (o *Item) GetRemoteIdOk() (*string, bool)`

GetRemoteIdOk returns a tuple with the RemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteId

`func (o *Item) SetRemoteId(v string)`

SetRemoteId sets RemoteId field to given value.

### HasRemoteId

`func (o *Item) HasRemoteId() bool`

HasRemoteId returns a boolean if a field has been set.

### SetRemoteIdNil

`func (o *Item) SetRemoteIdNil(b bool)`

 SetRemoteIdNil sets the value for RemoteId to be an explicit nil

### UnsetRemoteId
`func (o *Item) UnsetRemoteId()`

UnsetRemoteId ensures that no value is present for RemoteId, not even an explicit nil
### GetName

`func (o *Item) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Item) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Item) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Item) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Item) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Item) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetStatus

`func (o *Item) GetStatus() Status7d1Enum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Item) GetStatusOk() (*Status7d1Enum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Item) SetStatus(v Status7d1Enum)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Item) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *Item) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *Item) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetUnitPrice

`func (o *Item) GetUnitPrice() float64`

GetUnitPrice returns the UnitPrice field if non-nil, zero value otherwise.

### GetUnitPriceOk

`func (o *Item) GetUnitPriceOk() (*float64, bool)`

GetUnitPriceOk returns a tuple with the UnitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPrice

`func (o *Item) SetUnitPrice(v float64)`

SetUnitPrice sets UnitPrice field to given value.

### HasUnitPrice

`func (o *Item) HasUnitPrice() bool`

HasUnitPrice returns a boolean if a field has been set.

### SetUnitPriceNil

`func (o *Item) SetUnitPriceNil(b bool)`

 SetUnitPriceNil sets the value for UnitPrice to be an explicit nil

### UnsetUnitPrice
`func (o *Item) UnsetUnitPrice()`

UnsetUnitPrice ensures that no value is present for UnitPrice, not even an explicit nil
### GetPurchasePrice

`func (o *Item) GetPurchasePrice() float64`

GetPurchasePrice returns the PurchasePrice field if non-nil, zero value otherwise.

### GetPurchasePriceOk

`func (o *Item) GetPurchasePriceOk() (*float64, bool)`

GetPurchasePriceOk returns a tuple with the PurchasePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchasePrice

`func (o *Item) SetPurchasePrice(v float64)`

SetPurchasePrice sets PurchasePrice field to given value.

### HasPurchasePrice

`func (o *Item) HasPurchasePrice() bool`

HasPurchasePrice returns a boolean if a field has been set.

### SetPurchasePriceNil

`func (o *Item) SetPurchasePriceNil(b bool)`

 SetPurchasePriceNil sets the value for PurchasePrice to be an explicit nil

### UnsetPurchasePrice
`func (o *Item) UnsetPurchasePrice()`

UnsetPurchasePrice ensures that no value is present for PurchasePrice, not even an explicit nil
### GetPurchaseAccount

`func (o *Item) GetPurchaseAccount() string`

GetPurchaseAccount returns the PurchaseAccount field if non-nil, zero value otherwise.

### GetPurchaseAccountOk

`func (o *Item) GetPurchaseAccountOk() (*string, bool)`

GetPurchaseAccountOk returns a tuple with the PurchaseAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseAccount

`func (o *Item) SetPurchaseAccount(v string)`

SetPurchaseAccount sets PurchaseAccount field to given value.

### HasPurchaseAccount

`func (o *Item) HasPurchaseAccount() bool`

HasPurchaseAccount returns a boolean if a field has been set.

### SetPurchaseAccountNil

`func (o *Item) SetPurchaseAccountNil(b bool)`

 SetPurchaseAccountNil sets the value for PurchaseAccount to be an explicit nil

### UnsetPurchaseAccount
`func (o *Item) UnsetPurchaseAccount()`

UnsetPurchaseAccount ensures that no value is present for PurchaseAccount, not even an explicit nil
### GetSalesAccount

`func (o *Item) GetSalesAccount() string`

GetSalesAccount returns the SalesAccount field if non-nil, zero value otherwise.

### GetSalesAccountOk

`func (o *Item) GetSalesAccountOk() (*string, bool)`

GetSalesAccountOk returns a tuple with the SalesAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesAccount

`func (o *Item) SetSalesAccount(v string)`

SetSalesAccount sets SalesAccount field to given value.

### HasSalesAccount

`func (o *Item) HasSalesAccount() bool`

HasSalesAccount returns a boolean if a field has been set.

### SetSalesAccountNil

`func (o *Item) SetSalesAccountNil(b bool)`

 SetSalesAccountNil sets the value for SalesAccount to be an explicit nil

### UnsetSalesAccount
`func (o *Item) UnsetSalesAccount()`

UnsetSalesAccount ensures that no value is present for SalesAccount, not even an explicit nil
### GetCompany

`func (o *Item) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *Item) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *Item) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *Item) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### SetCompanyNil

`func (o *Item) SetCompanyNil(b bool)`

 SetCompanyNil sets the value for Company to be an explicit nil

### UnsetCompany
`func (o *Item) UnsetCompany()`

UnsetCompany ensures that no value is present for Company, not even an explicit nil
### GetRemoteUpdatedAt

`func (o *Item) GetRemoteUpdatedAt() time.Time`

GetRemoteUpdatedAt returns the RemoteUpdatedAt field if non-nil, zero value otherwise.

### GetRemoteUpdatedAtOk

`func (o *Item) GetRemoteUpdatedAtOk() (*time.Time, bool)`

GetRemoteUpdatedAtOk returns a tuple with the RemoteUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteUpdatedAt

`func (o *Item) SetRemoteUpdatedAt(v time.Time)`

SetRemoteUpdatedAt sets RemoteUpdatedAt field to given value.

### HasRemoteUpdatedAt

`func (o *Item) HasRemoteUpdatedAt() bool`

HasRemoteUpdatedAt returns a boolean if a field has been set.

### SetRemoteUpdatedAtNil

`func (o *Item) SetRemoteUpdatedAtNil(b bool)`

 SetRemoteUpdatedAtNil sets the value for RemoteUpdatedAt to be an explicit nil

### UnsetRemoteUpdatedAt
`func (o *Item) UnsetRemoteUpdatedAt()`

UnsetRemoteUpdatedAt ensures that no value is present for RemoteUpdatedAt, not even an explicit nil
### GetRemoteWasDeleted

`func (o *Item) GetRemoteWasDeleted() bool`

GetRemoteWasDeleted returns the RemoteWasDeleted field if non-nil, zero value otherwise.

### GetRemoteWasDeletedOk

`func (o *Item) GetRemoteWasDeletedOk() (*bool, bool)`

GetRemoteWasDeletedOk returns a tuple with the RemoteWasDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteWasDeleted

`func (o *Item) SetRemoteWasDeleted(v bool)`

SetRemoteWasDeleted sets RemoteWasDeleted field to given value.

### HasRemoteWasDeleted

`func (o *Item) HasRemoteWasDeleted() bool`

HasRemoteWasDeleted returns a boolean if a field has been set.

### GetFieldMappings

`func (o *Item) GetFieldMappings() map[string]interface{}`

GetFieldMappings returns the FieldMappings field if non-nil, zero value otherwise.

### GetFieldMappingsOk

`func (o *Item) GetFieldMappingsOk() (*map[string]interface{}, bool)`

GetFieldMappingsOk returns a tuple with the FieldMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldMappings

`func (o *Item) SetFieldMappings(v map[string]interface{})`

SetFieldMappings sets FieldMappings field to given value.

### HasFieldMappings

`func (o *Item) HasFieldMappings() bool`

HasFieldMappings returns a boolean if a field has been set.

### SetFieldMappingsNil

`func (o *Item) SetFieldMappingsNil(b bool)`

 SetFieldMappingsNil sets the value for FieldMappings to be an explicit nil

### UnsetFieldMappings
`func (o *Item) UnsetFieldMappings()`

UnsetFieldMappings ensures that no value is present for FieldMappings, not even an explicit nil
### GetModifiedAt

`func (o *Item) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *Item) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *Item) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *Item) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetRemoteData

`func (o *Item) GetRemoteData() []RemoteData`

GetRemoteData returns the RemoteData field if non-nil, zero value otherwise.

### GetRemoteDataOk

`func (o *Item) GetRemoteDataOk() (*[]RemoteData, bool)`

GetRemoteDataOk returns a tuple with the RemoteData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteData

`func (o *Item) SetRemoteData(v []RemoteData)`

SetRemoteData sets RemoteData field to given value.

### HasRemoteData

`func (o *Item) HasRemoteData() bool`

HasRemoteData returns a boolean if a field has been set.

### SetRemoteDataNil

`func (o *Item) SetRemoteDataNil(b bool)`

 SetRemoteDataNil sets the value for RemoteData to be an explicit nil

### UnsetRemoteData
`func (o *Item) UnsetRemoteData()`

UnsetRemoteData ensures that no value is present for RemoteData, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


