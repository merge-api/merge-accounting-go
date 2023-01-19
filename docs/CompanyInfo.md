# CompanyInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**RemoteId** | Pointer to **NullableString** | The third-party API ID of the matching object. | [optional] 
**RemoteData** | Pointer to [**[]RemoteData**](RemoteData.md) |  | [optional] [readonly] 
**Name** | Pointer to **NullableString** | The company&#39;s name. | [optional] 
**LegalName** | Pointer to **NullableString** | The company&#39;s legal name. | [optional] 
**TaxNumber** | Pointer to **NullableString** | The company&#39;s tax number. | [optional] 
**FiscalYearEndMonth** | Pointer to **NullableInt32** | The company&#39;s fiscal year end month. | [optional] 
**FiscalYearEndDay** | Pointer to **NullableInt32** | The company&#39;s fiscal year end day. | [optional] 
**Currency** | Pointer to [**NullableCurrencyEnum**](CurrencyEnum.md) | The currency set in the company&#39;s accounting platform. | [optional] 
**RemoteCreatedAt** | Pointer to **NullableTime** | When the third party&#39;s company was created. | [optional] 
**Urls** | Pointer to **[]string** | The company&#39;s urls. | [optional] 
**Addresses** | Pointer to [**[]Address**](Address.md) |  | [optional] 
**PhoneNumbers** | Pointer to [**[]AccountingPhoneNumber**](AccountingPhoneNumber.md) |  | [optional] 
**RemoteWasDeleted** | Pointer to **bool** | Indicates whether or not this object has been deleted by third party webhooks. | [optional] [readonly] 
**FieldMappings** | Pointer to **map[string]interface{}** |  | [optional] [readonly] 

## Methods

### NewCompanyInfo

`func NewCompanyInfo() *CompanyInfo`

NewCompanyInfo instantiates a new CompanyInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyInfoWithDefaults

`func NewCompanyInfoWithDefaults() *CompanyInfo`

NewCompanyInfoWithDefaults instantiates a new CompanyInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CompanyInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CompanyInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CompanyInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CompanyInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRemoteId

`func (o *CompanyInfo) GetRemoteId() string`

GetRemoteId returns the RemoteId field if non-nil, zero value otherwise.

### GetRemoteIdOk

`func (o *CompanyInfo) GetRemoteIdOk() (*string, bool)`

GetRemoteIdOk returns a tuple with the RemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteId

`func (o *CompanyInfo) SetRemoteId(v string)`

SetRemoteId sets RemoteId field to given value.

### HasRemoteId

`func (o *CompanyInfo) HasRemoteId() bool`

HasRemoteId returns a boolean if a field has been set.

### SetRemoteIdNil

`func (o *CompanyInfo) SetRemoteIdNil(b bool)`

 SetRemoteIdNil sets the value for RemoteId to be an explicit nil

### UnsetRemoteId
`func (o *CompanyInfo) UnsetRemoteId()`

UnsetRemoteId ensures that no value is present for RemoteId, not even an explicit nil
### GetRemoteData

`func (o *CompanyInfo) GetRemoteData() []RemoteData`

GetRemoteData returns the RemoteData field if non-nil, zero value otherwise.

### GetRemoteDataOk

`func (o *CompanyInfo) GetRemoteDataOk() (*[]RemoteData, bool)`

GetRemoteDataOk returns a tuple with the RemoteData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteData

`func (o *CompanyInfo) SetRemoteData(v []RemoteData)`

SetRemoteData sets RemoteData field to given value.

### HasRemoteData

`func (o *CompanyInfo) HasRemoteData() bool`

HasRemoteData returns a boolean if a field has been set.

### SetRemoteDataNil

`func (o *CompanyInfo) SetRemoteDataNil(b bool)`

 SetRemoteDataNil sets the value for RemoteData to be an explicit nil

### UnsetRemoteData
`func (o *CompanyInfo) UnsetRemoteData()`

UnsetRemoteData ensures that no value is present for RemoteData, not even an explicit nil
### GetName

`func (o *CompanyInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CompanyInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CompanyInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CompanyInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CompanyInfo) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CompanyInfo) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetLegalName

`func (o *CompanyInfo) GetLegalName() string`

GetLegalName returns the LegalName field if non-nil, zero value otherwise.

### GetLegalNameOk

`func (o *CompanyInfo) GetLegalNameOk() (*string, bool)`

GetLegalNameOk returns a tuple with the LegalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalName

`func (o *CompanyInfo) SetLegalName(v string)`

SetLegalName sets LegalName field to given value.

### HasLegalName

`func (o *CompanyInfo) HasLegalName() bool`

HasLegalName returns a boolean if a field has been set.

### SetLegalNameNil

`func (o *CompanyInfo) SetLegalNameNil(b bool)`

 SetLegalNameNil sets the value for LegalName to be an explicit nil

### UnsetLegalName
`func (o *CompanyInfo) UnsetLegalName()`

UnsetLegalName ensures that no value is present for LegalName, not even an explicit nil
### GetTaxNumber

`func (o *CompanyInfo) GetTaxNumber() string`

GetTaxNumber returns the TaxNumber field if non-nil, zero value otherwise.

### GetTaxNumberOk

`func (o *CompanyInfo) GetTaxNumberOk() (*string, bool)`

GetTaxNumberOk returns a tuple with the TaxNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxNumber

`func (o *CompanyInfo) SetTaxNumber(v string)`

SetTaxNumber sets TaxNumber field to given value.

### HasTaxNumber

`func (o *CompanyInfo) HasTaxNumber() bool`

HasTaxNumber returns a boolean if a field has been set.

### SetTaxNumberNil

`func (o *CompanyInfo) SetTaxNumberNil(b bool)`

 SetTaxNumberNil sets the value for TaxNumber to be an explicit nil

### UnsetTaxNumber
`func (o *CompanyInfo) UnsetTaxNumber()`

UnsetTaxNumber ensures that no value is present for TaxNumber, not even an explicit nil
### GetFiscalYearEndMonth

`func (o *CompanyInfo) GetFiscalYearEndMonth() int32`

GetFiscalYearEndMonth returns the FiscalYearEndMonth field if non-nil, zero value otherwise.

### GetFiscalYearEndMonthOk

`func (o *CompanyInfo) GetFiscalYearEndMonthOk() (*int32, bool)`

GetFiscalYearEndMonthOk returns a tuple with the FiscalYearEndMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiscalYearEndMonth

`func (o *CompanyInfo) SetFiscalYearEndMonth(v int32)`

SetFiscalYearEndMonth sets FiscalYearEndMonth field to given value.

### HasFiscalYearEndMonth

`func (o *CompanyInfo) HasFiscalYearEndMonth() bool`

HasFiscalYearEndMonth returns a boolean if a field has been set.

### SetFiscalYearEndMonthNil

`func (o *CompanyInfo) SetFiscalYearEndMonthNil(b bool)`

 SetFiscalYearEndMonthNil sets the value for FiscalYearEndMonth to be an explicit nil

### UnsetFiscalYearEndMonth
`func (o *CompanyInfo) UnsetFiscalYearEndMonth()`

UnsetFiscalYearEndMonth ensures that no value is present for FiscalYearEndMonth, not even an explicit nil
### GetFiscalYearEndDay

`func (o *CompanyInfo) GetFiscalYearEndDay() int32`

GetFiscalYearEndDay returns the FiscalYearEndDay field if non-nil, zero value otherwise.

### GetFiscalYearEndDayOk

`func (o *CompanyInfo) GetFiscalYearEndDayOk() (*int32, bool)`

GetFiscalYearEndDayOk returns a tuple with the FiscalYearEndDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiscalYearEndDay

`func (o *CompanyInfo) SetFiscalYearEndDay(v int32)`

SetFiscalYearEndDay sets FiscalYearEndDay field to given value.

### HasFiscalYearEndDay

`func (o *CompanyInfo) HasFiscalYearEndDay() bool`

HasFiscalYearEndDay returns a boolean if a field has been set.

### SetFiscalYearEndDayNil

`func (o *CompanyInfo) SetFiscalYearEndDayNil(b bool)`

 SetFiscalYearEndDayNil sets the value for FiscalYearEndDay to be an explicit nil

### UnsetFiscalYearEndDay
`func (o *CompanyInfo) UnsetFiscalYearEndDay()`

UnsetFiscalYearEndDay ensures that no value is present for FiscalYearEndDay, not even an explicit nil
### GetCurrency

`func (o *CompanyInfo) GetCurrency() CurrencyEnum`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CompanyInfo) GetCurrencyOk() (*CurrencyEnum, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CompanyInfo) SetCurrency(v CurrencyEnum)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *CompanyInfo) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### SetCurrencyNil

`func (o *CompanyInfo) SetCurrencyNil(b bool)`

 SetCurrencyNil sets the value for Currency to be an explicit nil

### UnsetCurrency
`func (o *CompanyInfo) UnsetCurrency()`

UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
### GetRemoteCreatedAt

`func (o *CompanyInfo) GetRemoteCreatedAt() time.Time`

GetRemoteCreatedAt returns the RemoteCreatedAt field if non-nil, zero value otherwise.

### GetRemoteCreatedAtOk

`func (o *CompanyInfo) GetRemoteCreatedAtOk() (*time.Time, bool)`

GetRemoteCreatedAtOk returns a tuple with the RemoteCreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteCreatedAt

`func (o *CompanyInfo) SetRemoteCreatedAt(v time.Time)`

SetRemoteCreatedAt sets RemoteCreatedAt field to given value.

### HasRemoteCreatedAt

`func (o *CompanyInfo) HasRemoteCreatedAt() bool`

HasRemoteCreatedAt returns a boolean if a field has been set.

### SetRemoteCreatedAtNil

`func (o *CompanyInfo) SetRemoteCreatedAtNil(b bool)`

 SetRemoteCreatedAtNil sets the value for RemoteCreatedAt to be an explicit nil

### UnsetRemoteCreatedAt
`func (o *CompanyInfo) UnsetRemoteCreatedAt()`

UnsetRemoteCreatedAt ensures that no value is present for RemoteCreatedAt, not even an explicit nil
### GetUrls

`func (o *CompanyInfo) GetUrls() []string`

GetUrls returns the Urls field if non-nil, zero value otherwise.

### GetUrlsOk

`func (o *CompanyInfo) GetUrlsOk() (*[]string, bool)`

GetUrlsOk returns a tuple with the Urls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrls

`func (o *CompanyInfo) SetUrls(v []string)`

SetUrls sets Urls field to given value.

### HasUrls

`func (o *CompanyInfo) HasUrls() bool`

HasUrls returns a boolean if a field has been set.

### SetUrlsNil

`func (o *CompanyInfo) SetUrlsNil(b bool)`

 SetUrlsNil sets the value for Urls to be an explicit nil

### UnsetUrls
`func (o *CompanyInfo) UnsetUrls()`

UnsetUrls ensures that no value is present for Urls, not even an explicit nil
### GetAddresses

`func (o *CompanyInfo) GetAddresses() []Address`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *CompanyInfo) GetAddressesOk() (*[]Address, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *CompanyInfo) SetAddresses(v []Address)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *CompanyInfo) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.

### GetPhoneNumbers

`func (o *CompanyInfo) GetPhoneNumbers() []AccountingPhoneNumber`

GetPhoneNumbers returns the PhoneNumbers field if non-nil, zero value otherwise.

### GetPhoneNumbersOk

`func (o *CompanyInfo) GetPhoneNumbersOk() (*[]AccountingPhoneNumber, bool)`

GetPhoneNumbersOk returns a tuple with the PhoneNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumbers

`func (o *CompanyInfo) SetPhoneNumbers(v []AccountingPhoneNumber)`

SetPhoneNumbers sets PhoneNumbers field to given value.

### HasPhoneNumbers

`func (o *CompanyInfo) HasPhoneNumbers() bool`

HasPhoneNumbers returns a boolean if a field has been set.

### GetRemoteWasDeleted

`func (o *CompanyInfo) GetRemoteWasDeleted() bool`

GetRemoteWasDeleted returns the RemoteWasDeleted field if non-nil, zero value otherwise.

### GetRemoteWasDeletedOk

`func (o *CompanyInfo) GetRemoteWasDeletedOk() (*bool, bool)`

GetRemoteWasDeletedOk returns a tuple with the RemoteWasDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteWasDeleted

`func (o *CompanyInfo) SetRemoteWasDeleted(v bool)`

SetRemoteWasDeleted sets RemoteWasDeleted field to given value.

### HasRemoteWasDeleted

`func (o *CompanyInfo) HasRemoteWasDeleted() bool`

HasRemoteWasDeleted returns a boolean if a field has been set.

### GetFieldMappings

`func (o *CompanyInfo) GetFieldMappings() map[string]interface{}`

GetFieldMappings returns the FieldMappings field if non-nil, zero value otherwise.

### GetFieldMappingsOk

`func (o *CompanyInfo) GetFieldMappingsOk() (*map[string]interface{}, bool)`

GetFieldMappingsOk returns a tuple with the FieldMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldMappings

`func (o *CompanyInfo) SetFieldMappings(v map[string]interface{})`

SetFieldMappings sets FieldMappings field to given value.

### HasFieldMappings

`func (o *CompanyInfo) HasFieldMappings() bool`

HasFieldMappings returns a boolean if a field has been set.

### SetFieldMappingsNil

`func (o *CompanyInfo) SetFieldMappingsNil(b bool)`

 SetFieldMappingsNil sets the value for FieldMappings to be an explicit nil

### UnsetFieldMappings
`func (o *CompanyInfo) UnsetFieldMappings()`

UnsetFieldMappings ensures that no value is present for FieldMappings, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


