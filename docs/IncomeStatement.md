# IncomeStatement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**RemoteId** | Pointer to **NullableString** | The third-party API ID of the matching object. | [optional] 
**RemoteData** | Pointer to [**[]RemoteData**](RemoteData.md) |  | [optional] [readonly] 
**Name** | Pointer to **NullableString** | The income statement&#39;s name. | [optional] 
**Currency** | Pointer to [**NullableCurrencyEnum**](CurrencyEnum.md) | The income statement&#39;s currency. | [optional] 
**Company** | Pointer to **NullableString** | The company the income statement belongs to. | [optional] 
**StartPeriod** | Pointer to **NullableTime** | The income statement&#39;s start period. | [optional] 
**EndPeriod** | Pointer to **NullableTime** | The income statement&#39;s end period. | [optional] 
**Income** | Pointer to [**[]ReportItem**](ReportItem.md) |  | [optional] [readonly] 
**CostOfSales** | Pointer to [**[]ReportItem**](ReportItem.md) |  | [optional] [readonly] 
**GrossProfit** | Pointer to **NullableFloat32** | The revenue minus the cost of sale. | [optional] 
**OperatingExpenses** | Pointer to [**[]ReportItem**](ReportItem.md) |  | [optional] [readonly] 
**NetOperatingIncome** | Pointer to **NullableFloat32** | The revenue minus the operating expenses. | [optional] 
**NonOperatingExpenses** | Pointer to [**[]ReportItem**](ReportItem.md) |  | [optional] [readonly] 
**NetIncome** | Pointer to **NullableFloat32** | The gross profit minus the total expenses. | [optional] 
**RemoteWasDeleted** | Pointer to **bool** | Indicates whether or not this object has been deleted by third party webhooks. | [optional] [readonly] 
**FieldMappings** | Pointer to **map[string]interface{}** |  | [optional] [readonly] 

## Methods

### NewIncomeStatement

`func NewIncomeStatement() *IncomeStatement`

NewIncomeStatement instantiates a new IncomeStatement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncomeStatementWithDefaults

`func NewIncomeStatementWithDefaults() *IncomeStatement`

NewIncomeStatementWithDefaults instantiates a new IncomeStatement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IncomeStatement) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IncomeStatement) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IncomeStatement) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IncomeStatement) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRemoteId

`func (o *IncomeStatement) GetRemoteId() string`

GetRemoteId returns the RemoteId field if non-nil, zero value otherwise.

### GetRemoteIdOk

`func (o *IncomeStatement) GetRemoteIdOk() (*string, bool)`

GetRemoteIdOk returns a tuple with the RemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteId

`func (o *IncomeStatement) SetRemoteId(v string)`

SetRemoteId sets RemoteId field to given value.

### HasRemoteId

`func (o *IncomeStatement) HasRemoteId() bool`

HasRemoteId returns a boolean if a field has been set.

### SetRemoteIdNil

`func (o *IncomeStatement) SetRemoteIdNil(b bool)`

 SetRemoteIdNil sets the value for RemoteId to be an explicit nil

### UnsetRemoteId
`func (o *IncomeStatement) UnsetRemoteId()`

UnsetRemoteId ensures that no value is present for RemoteId, not even an explicit nil
### GetRemoteData

`func (o *IncomeStatement) GetRemoteData() []RemoteData`

GetRemoteData returns the RemoteData field if non-nil, zero value otherwise.

### GetRemoteDataOk

`func (o *IncomeStatement) GetRemoteDataOk() (*[]RemoteData, bool)`

GetRemoteDataOk returns a tuple with the RemoteData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteData

`func (o *IncomeStatement) SetRemoteData(v []RemoteData)`

SetRemoteData sets RemoteData field to given value.

### HasRemoteData

`func (o *IncomeStatement) HasRemoteData() bool`

HasRemoteData returns a boolean if a field has been set.

### SetRemoteDataNil

`func (o *IncomeStatement) SetRemoteDataNil(b bool)`

 SetRemoteDataNil sets the value for RemoteData to be an explicit nil

### UnsetRemoteData
`func (o *IncomeStatement) UnsetRemoteData()`

UnsetRemoteData ensures that no value is present for RemoteData, not even an explicit nil
### GetName

`func (o *IncomeStatement) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IncomeStatement) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IncomeStatement) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IncomeStatement) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *IncomeStatement) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *IncomeStatement) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCurrency

`func (o *IncomeStatement) GetCurrency() CurrencyEnum`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *IncomeStatement) GetCurrencyOk() (*CurrencyEnum, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *IncomeStatement) SetCurrency(v CurrencyEnum)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *IncomeStatement) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### SetCurrencyNil

`func (o *IncomeStatement) SetCurrencyNil(b bool)`

 SetCurrencyNil sets the value for Currency to be an explicit nil

### UnsetCurrency
`func (o *IncomeStatement) UnsetCurrency()`

UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
### GetCompany

`func (o *IncomeStatement) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *IncomeStatement) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *IncomeStatement) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *IncomeStatement) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### SetCompanyNil

`func (o *IncomeStatement) SetCompanyNil(b bool)`

 SetCompanyNil sets the value for Company to be an explicit nil

### UnsetCompany
`func (o *IncomeStatement) UnsetCompany()`

UnsetCompany ensures that no value is present for Company, not even an explicit nil
### GetStartPeriod

`func (o *IncomeStatement) GetStartPeriod() time.Time`

GetStartPeriod returns the StartPeriod field if non-nil, zero value otherwise.

### GetStartPeriodOk

`func (o *IncomeStatement) GetStartPeriodOk() (*time.Time, bool)`

GetStartPeriodOk returns a tuple with the StartPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartPeriod

`func (o *IncomeStatement) SetStartPeriod(v time.Time)`

SetStartPeriod sets StartPeriod field to given value.

### HasStartPeriod

`func (o *IncomeStatement) HasStartPeriod() bool`

HasStartPeriod returns a boolean if a field has been set.

### SetStartPeriodNil

`func (o *IncomeStatement) SetStartPeriodNil(b bool)`

 SetStartPeriodNil sets the value for StartPeriod to be an explicit nil

### UnsetStartPeriod
`func (o *IncomeStatement) UnsetStartPeriod()`

UnsetStartPeriod ensures that no value is present for StartPeriod, not even an explicit nil
### GetEndPeriod

`func (o *IncomeStatement) GetEndPeriod() time.Time`

GetEndPeriod returns the EndPeriod field if non-nil, zero value otherwise.

### GetEndPeriodOk

`func (o *IncomeStatement) GetEndPeriodOk() (*time.Time, bool)`

GetEndPeriodOk returns a tuple with the EndPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndPeriod

`func (o *IncomeStatement) SetEndPeriod(v time.Time)`

SetEndPeriod sets EndPeriod field to given value.

### HasEndPeriod

`func (o *IncomeStatement) HasEndPeriod() bool`

HasEndPeriod returns a boolean if a field has been set.

### SetEndPeriodNil

`func (o *IncomeStatement) SetEndPeriodNil(b bool)`

 SetEndPeriodNil sets the value for EndPeriod to be an explicit nil

### UnsetEndPeriod
`func (o *IncomeStatement) UnsetEndPeriod()`

UnsetEndPeriod ensures that no value is present for EndPeriod, not even an explicit nil
### GetIncome

`func (o *IncomeStatement) GetIncome() []ReportItem`

GetIncome returns the Income field if non-nil, zero value otherwise.

### GetIncomeOk

`func (o *IncomeStatement) GetIncomeOk() (*[]ReportItem, bool)`

GetIncomeOk returns a tuple with the Income field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncome

`func (o *IncomeStatement) SetIncome(v []ReportItem)`

SetIncome sets Income field to given value.

### HasIncome

`func (o *IncomeStatement) HasIncome() bool`

HasIncome returns a boolean if a field has been set.

### GetCostOfSales

`func (o *IncomeStatement) GetCostOfSales() []ReportItem`

GetCostOfSales returns the CostOfSales field if non-nil, zero value otherwise.

### GetCostOfSalesOk

`func (o *IncomeStatement) GetCostOfSalesOk() (*[]ReportItem, bool)`

GetCostOfSalesOk returns a tuple with the CostOfSales field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostOfSales

`func (o *IncomeStatement) SetCostOfSales(v []ReportItem)`

SetCostOfSales sets CostOfSales field to given value.

### HasCostOfSales

`func (o *IncomeStatement) HasCostOfSales() bool`

HasCostOfSales returns a boolean if a field has been set.

### GetGrossProfit

`func (o *IncomeStatement) GetGrossProfit() float32`

GetGrossProfit returns the GrossProfit field if non-nil, zero value otherwise.

### GetGrossProfitOk

`func (o *IncomeStatement) GetGrossProfitOk() (*float32, bool)`

GetGrossProfitOk returns a tuple with the GrossProfit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrossProfit

`func (o *IncomeStatement) SetGrossProfit(v float32)`

SetGrossProfit sets GrossProfit field to given value.

### HasGrossProfit

`func (o *IncomeStatement) HasGrossProfit() bool`

HasGrossProfit returns a boolean if a field has been set.

### SetGrossProfitNil

`func (o *IncomeStatement) SetGrossProfitNil(b bool)`

 SetGrossProfitNil sets the value for GrossProfit to be an explicit nil

### UnsetGrossProfit
`func (o *IncomeStatement) UnsetGrossProfit()`

UnsetGrossProfit ensures that no value is present for GrossProfit, not even an explicit nil
### GetOperatingExpenses

`func (o *IncomeStatement) GetOperatingExpenses() []ReportItem`

GetOperatingExpenses returns the OperatingExpenses field if non-nil, zero value otherwise.

### GetOperatingExpensesOk

`func (o *IncomeStatement) GetOperatingExpensesOk() (*[]ReportItem, bool)`

GetOperatingExpensesOk returns a tuple with the OperatingExpenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingExpenses

`func (o *IncomeStatement) SetOperatingExpenses(v []ReportItem)`

SetOperatingExpenses sets OperatingExpenses field to given value.

### HasOperatingExpenses

`func (o *IncomeStatement) HasOperatingExpenses() bool`

HasOperatingExpenses returns a boolean if a field has been set.

### GetNetOperatingIncome

`func (o *IncomeStatement) GetNetOperatingIncome() float32`

GetNetOperatingIncome returns the NetOperatingIncome field if non-nil, zero value otherwise.

### GetNetOperatingIncomeOk

`func (o *IncomeStatement) GetNetOperatingIncomeOk() (*float32, bool)`

GetNetOperatingIncomeOk returns a tuple with the NetOperatingIncome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetOperatingIncome

`func (o *IncomeStatement) SetNetOperatingIncome(v float32)`

SetNetOperatingIncome sets NetOperatingIncome field to given value.

### HasNetOperatingIncome

`func (o *IncomeStatement) HasNetOperatingIncome() bool`

HasNetOperatingIncome returns a boolean if a field has been set.

### SetNetOperatingIncomeNil

`func (o *IncomeStatement) SetNetOperatingIncomeNil(b bool)`

 SetNetOperatingIncomeNil sets the value for NetOperatingIncome to be an explicit nil

### UnsetNetOperatingIncome
`func (o *IncomeStatement) UnsetNetOperatingIncome()`

UnsetNetOperatingIncome ensures that no value is present for NetOperatingIncome, not even an explicit nil
### GetNonOperatingExpenses

`func (o *IncomeStatement) GetNonOperatingExpenses() []ReportItem`

GetNonOperatingExpenses returns the NonOperatingExpenses field if non-nil, zero value otherwise.

### GetNonOperatingExpensesOk

`func (o *IncomeStatement) GetNonOperatingExpensesOk() (*[]ReportItem, bool)`

GetNonOperatingExpensesOk returns a tuple with the NonOperatingExpenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonOperatingExpenses

`func (o *IncomeStatement) SetNonOperatingExpenses(v []ReportItem)`

SetNonOperatingExpenses sets NonOperatingExpenses field to given value.

### HasNonOperatingExpenses

`func (o *IncomeStatement) HasNonOperatingExpenses() bool`

HasNonOperatingExpenses returns a boolean if a field has been set.

### GetNetIncome

`func (o *IncomeStatement) GetNetIncome() float32`

GetNetIncome returns the NetIncome field if non-nil, zero value otherwise.

### GetNetIncomeOk

`func (o *IncomeStatement) GetNetIncomeOk() (*float32, bool)`

GetNetIncomeOk returns a tuple with the NetIncome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetIncome

`func (o *IncomeStatement) SetNetIncome(v float32)`

SetNetIncome sets NetIncome field to given value.

### HasNetIncome

`func (o *IncomeStatement) HasNetIncome() bool`

HasNetIncome returns a boolean if a field has been set.

### SetNetIncomeNil

`func (o *IncomeStatement) SetNetIncomeNil(b bool)`

 SetNetIncomeNil sets the value for NetIncome to be an explicit nil

### UnsetNetIncome
`func (o *IncomeStatement) UnsetNetIncome()`

UnsetNetIncome ensures that no value is present for NetIncome, not even an explicit nil
### GetRemoteWasDeleted

`func (o *IncomeStatement) GetRemoteWasDeleted() bool`

GetRemoteWasDeleted returns the RemoteWasDeleted field if non-nil, zero value otherwise.

### GetRemoteWasDeletedOk

`func (o *IncomeStatement) GetRemoteWasDeletedOk() (*bool, bool)`

GetRemoteWasDeletedOk returns a tuple with the RemoteWasDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteWasDeleted

`func (o *IncomeStatement) SetRemoteWasDeleted(v bool)`

SetRemoteWasDeleted sets RemoteWasDeleted field to given value.

### HasRemoteWasDeleted

`func (o *IncomeStatement) HasRemoteWasDeleted() bool`

HasRemoteWasDeleted returns a boolean if a field has been set.

### GetFieldMappings

`func (o *IncomeStatement) GetFieldMappings() map[string]interface{}`

GetFieldMappings returns the FieldMappings field if non-nil, zero value otherwise.

### GetFieldMappingsOk

`func (o *IncomeStatement) GetFieldMappingsOk() (*map[string]interface{}, bool)`

GetFieldMappingsOk returns a tuple with the FieldMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldMappings

`func (o *IncomeStatement) SetFieldMappings(v map[string]interface{})`

SetFieldMappings sets FieldMappings field to given value.

### HasFieldMappings

`func (o *IncomeStatement) HasFieldMappings() bool`

HasFieldMappings returns a boolean if a field has been set.

### SetFieldMappingsNil

`func (o *IncomeStatement) SetFieldMappingsNil(b bool)`

 SetFieldMappingsNil sets the value for FieldMappings to be an explicit nil

### UnsetFieldMappings
`func (o *IncomeStatement) UnsetFieldMappings()`

UnsetFieldMappings ensures that no value is present for FieldMappings, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


