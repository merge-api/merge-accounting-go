# TaxRate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**RemoteId** | Pointer to **NullableString** | The third-party API ID of the matching object. | [optional] 
**RemoteData** | Pointer to [**[]RemoteData**](RemoteData.md) |  | [optional] [readonly] 
**Description** | Pointer to **NullableString** | The tax rate&#39;s description. | [optional] 
**TotalTaxRate** | Pointer to **NullableFloat32** | The tax rate&#39;s total tax rate. | [optional] 
**EffectiveTaxRate** | Pointer to **NullableFloat32** | The tax rate&#39;s effective tax rate. | [optional] 
**Company** | Pointer to **NullableString** | The company the tax rate belongs to. | [optional] 
**RemoteWasDeleted** | Pointer to **bool** | Indicates whether or not this object has been deleted by third party webhooks. | [optional] [readonly] 
**FieldMappings** | Pointer to **map[string]interface{}** |  | [optional] [readonly] 

## Methods

### NewTaxRate

`func NewTaxRate() *TaxRate`

NewTaxRate instantiates a new TaxRate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxRateWithDefaults

`func NewTaxRateWithDefaults() *TaxRate`

NewTaxRateWithDefaults instantiates a new TaxRate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TaxRate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TaxRate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TaxRate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TaxRate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRemoteId

`func (o *TaxRate) GetRemoteId() string`

GetRemoteId returns the RemoteId field if non-nil, zero value otherwise.

### GetRemoteIdOk

`func (o *TaxRate) GetRemoteIdOk() (*string, bool)`

GetRemoteIdOk returns a tuple with the RemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteId

`func (o *TaxRate) SetRemoteId(v string)`

SetRemoteId sets RemoteId field to given value.

### HasRemoteId

`func (o *TaxRate) HasRemoteId() bool`

HasRemoteId returns a boolean if a field has been set.

### SetRemoteIdNil

`func (o *TaxRate) SetRemoteIdNil(b bool)`

 SetRemoteIdNil sets the value for RemoteId to be an explicit nil

### UnsetRemoteId
`func (o *TaxRate) UnsetRemoteId()`

UnsetRemoteId ensures that no value is present for RemoteId, not even an explicit nil
### GetRemoteData

`func (o *TaxRate) GetRemoteData() []RemoteData`

GetRemoteData returns the RemoteData field if non-nil, zero value otherwise.

### GetRemoteDataOk

`func (o *TaxRate) GetRemoteDataOk() (*[]RemoteData, bool)`

GetRemoteDataOk returns a tuple with the RemoteData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteData

`func (o *TaxRate) SetRemoteData(v []RemoteData)`

SetRemoteData sets RemoteData field to given value.

### HasRemoteData

`func (o *TaxRate) HasRemoteData() bool`

HasRemoteData returns a boolean if a field has been set.

### SetRemoteDataNil

`func (o *TaxRate) SetRemoteDataNil(b bool)`

 SetRemoteDataNil sets the value for RemoteData to be an explicit nil

### UnsetRemoteData
`func (o *TaxRate) UnsetRemoteData()`

UnsetRemoteData ensures that no value is present for RemoteData, not even an explicit nil
### GetDescription

`func (o *TaxRate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TaxRate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TaxRate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TaxRate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *TaxRate) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *TaxRate) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetTotalTaxRate

`func (o *TaxRate) GetTotalTaxRate() float32`

GetTotalTaxRate returns the TotalTaxRate field if non-nil, zero value otherwise.

### GetTotalTaxRateOk

`func (o *TaxRate) GetTotalTaxRateOk() (*float32, bool)`

GetTotalTaxRateOk returns a tuple with the TotalTaxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTaxRate

`func (o *TaxRate) SetTotalTaxRate(v float32)`

SetTotalTaxRate sets TotalTaxRate field to given value.

### HasTotalTaxRate

`func (o *TaxRate) HasTotalTaxRate() bool`

HasTotalTaxRate returns a boolean if a field has been set.

### SetTotalTaxRateNil

`func (o *TaxRate) SetTotalTaxRateNil(b bool)`

 SetTotalTaxRateNil sets the value for TotalTaxRate to be an explicit nil

### UnsetTotalTaxRate
`func (o *TaxRate) UnsetTotalTaxRate()`

UnsetTotalTaxRate ensures that no value is present for TotalTaxRate, not even an explicit nil
### GetEffectiveTaxRate

`func (o *TaxRate) GetEffectiveTaxRate() float32`

GetEffectiveTaxRate returns the EffectiveTaxRate field if non-nil, zero value otherwise.

### GetEffectiveTaxRateOk

`func (o *TaxRate) GetEffectiveTaxRateOk() (*float32, bool)`

GetEffectiveTaxRateOk returns a tuple with the EffectiveTaxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveTaxRate

`func (o *TaxRate) SetEffectiveTaxRate(v float32)`

SetEffectiveTaxRate sets EffectiveTaxRate field to given value.

### HasEffectiveTaxRate

`func (o *TaxRate) HasEffectiveTaxRate() bool`

HasEffectiveTaxRate returns a boolean if a field has been set.

### SetEffectiveTaxRateNil

`func (o *TaxRate) SetEffectiveTaxRateNil(b bool)`

 SetEffectiveTaxRateNil sets the value for EffectiveTaxRate to be an explicit nil

### UnsetEffectiveTaxRate
`func (o *TaxRate) UnsetEffectiveTaxRate()`

UnsetEffectiveTaxRate ensures that no value is present for EffectiveTaxRate, not even an explicit nil
### GetCompany

`func (o *TaxRate) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *TaxRate) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *TaxRate) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *TaxRate) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### SetCompanyNil

`func (o *TaxRate) SetCompanyNil(b bool)`

 SetCompanyNil sets the value for Company to be an explicit nil

### UnsetCompany
`func (o *TaxRate) UnsetCompany()`

UnsetCompany ensures that no value is present for Company, not even an explicit nil
### GetRemoteWasDeleted

`func (o *TaxRate) GetRemoteWasDeleted() bool`

GetRemoteWasDeleted returns the RemoteWasDeleted field if non-nil, zero value otherwise.

### GetRemoteWasDeletedOk

`func (o *TaxRate) GetRemoteWasDeletedOk() (*bool, bool)`

GetRemoteWasDeletedOk returns a tuple with the RemoteWasDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteWasDeleted

`func (o *TaxRate) SetRemoteWasDeleted(v bool)`

SetRemoteWasDeleted sets RemoteWasDeleted field to given value.

### HasRemoteWasDeleted

`func (o *TaxRate) HasRemoteWasDeleted() bool`

HasRemoteWasDeleted returns a boolean if a field has been set.

### GetFieldMappings

`func (o *TaxRate) GetFieldMappings() map[string]interface{}`

GetFieldMappings returns the FieldMappings field if non-nil, zero value otherwise.

### GetFieldMappingsOk

`func (o *TaxRate) GetFieldMappingsOk() (*map[string]interface{}, bool)`

GetFieldMappingsOk returns a tuple with the FieldMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldMappings

`func (o *TaxRate) SetFieldMappings(v map[string]interface{})`

SetFieldMappings sets FieldMappings field to given value.

### HasFieldMappings

`func (o *TaxRate) HasFieldMappings() bool`

HasFieldMappings returns a boolean if a field has been set.

### SetFieldMappingsNil

`func (o *TaxRate) SetFieldMappingsNil(b bool)`

 SetFieldMappingsNil sets the value for FieldMappings to be an explicit nil

### UnsetFieldMappings
`func (o *TaxRate) UnsetFieldMappings()`

UnsetFieldMappings ensures that no value is present for FieldMappings, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


