# Account

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**RemoteId** | Pointer to **NullableString** | The third-party API ID of the matching object. | [optional] 
**RemoteData** | Pointer to [**[]RemoteData**](RemoteData.md) |  | [optional] [readonly] 
**Name** | Pointer to **NullableString** | The account&#39;s name. | [optional] 
**Description** | Pointer to **NullableString** | The account&#39;s description. | [optional] 
**Classification** | Pointer to [**NullableClassificationEnum**](ClassificationEnum.md) | The account&#39;s broadest grouping. | [optional] 
**Type** | Pointer to **NullableString** | The account&#39;s type is a narrower and more specific grouping within the account&#39;s classification. | [optional] 
**Status** | Pointer to [**NullableAccountStatusEnum**](AccountStatusEnum.md) | The account&#39;s status. | [optional] 
**CurrentBalance** | Pointer to **NullableFloat32** | The account&#39;s current balance. | [optional] 
**Currency** | Pointer to [**NullableCurrencyEnum**](CurrencyEnum.md) | The account&#39;s currency. | [optional] 
**AccountNumber** | Pointer to **NullableString** | The account&#39;s number. | [optional] 
**ParentAccount** | Pointer to **NullableString** | ID of the parent account. | [optional] 
**Company** | Pointer to **NullableString** | The company the account belongs to. | [optional] 
**RemoteWasDeleted** | Pointer to **bool** | Indicates whether or not this object has been deleted by third party webhooks. | [optional] [readonly] 
**FieldMappings** | Pointer to **map[string]interface{}** |  | [optional] [readonly] 

## Methods

### NewAccount

`func NewAccount() *Account`

NewAccount instantiates a new Account object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountWithDefaults

`func NewAccountWithDefaults() *Account`

NewAccountWithDefaults instantiates a new Account object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Account) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Account) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Account) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Account) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRemoteId

`func (o *Account) GetRemoteId() string`

GetRemoteId returns the RemoteId field if non-nil, zero value otherwise.

### GetRemoteIdOk

`func (o *Account) GetRemoteIdOk() (*string, bool)`

GetRemoteIdOk returns a tuple with the RemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteId

`func (o *Account) SetRemoteId(v string)`

SetRemoteId sets RemoteId field to given value.

### HasRemoteId

`func (o *Account) HasRemoteId() bool`

HasRemoteId returns a boolean if a field has been set.

### SetRemoteIdNil

`func (o *Account) SetRemoteIdNil(b bool)`

 SetRemoteIdNil sets the value for RemoteId to be an explicit nil

### UnsetRemoteId
`func (o *Account) UnsetRemoteId()`

UnsetRemoteId ensures that no value is present for RemoteId, not even an explicit nil
### GetRemoteData

`func (o *Account) GetRemoteData() []RemoteData`

GetRemoteData returns the RemoteData field if non-nil, zero value otherwise.

### GetRemoteDataOk

`func (o *Account) GetRemoteDataOk() (*[]RemoteData, bool)`

GetRemoteDataOk returns a tuple with the RemoteData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteData

`func (o *Account) SetRemoteData(v []RemoteData)`

SetRemoteData sets RemoteData field to given value.

### HasRemoteData

`func (o *Account) HasRemoteData() bool`

HasRemoteData returns a boolean if a field has been set.

### SetRemoteDataNil

`func (o *Account) SetRemoteDataNil(b bool)`

 SetRemoteDataNil sets the value for RemoteData to be an explicit nil

### UnsetRemoteData
`func (o *Account) UnsetRemoteData()`

UnsetRemoteData ensures that no value is present for RemoteData, not even an explicit nil
### GetName

`func (o *Account) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Account) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Account) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Account) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Account) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Account) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *Account) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Account) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Account) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Account) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *Account) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Account) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetClassification

`func (o *Account) GetClassification() ClassificationEnum`

GetClassification returns the Classification field if non-nil, zero value otherwise.

### GetClassificationOk

`func (o *Account) GetClassificationOk() (*ClassificationEnum, bool)`

GetClassificationOk returns a tuple with the Classification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassification

`func (o *Account) SetClassification(v ClassificationEnum)`

SetClassification sets Classification field to given value.

### HasClassification

`func (o *Account) HasClassification() bool`

HasClassification returns a boolean if a field has been set.

### SetClassificationNil

`func (o *Account) SetClassificationNil(b bool)`

 SetClassificationNil sets the value for Classification to be an explicit nil

### UnsetClassification
`func (o *Account) UnsetClassification()`

UnsetClassification ensures that no value is present for Classification, not even an explicit nil
### GetType

`func (o *Account) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Account) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Account) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Account) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *Account) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *Account) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetStatus

`func (o *Account) GetStatus() AccountStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Account) GetStatusOk() (*AccountStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Account) SetStatus(v AccountStatusEnum)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Account) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *Account) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *Account) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetCurrentBalance

`func (o *Account) GetCurrentBalance() float32`

GetCurrentBalance returns the CurrentBalance field if non-nil, zero value otherwise.

### GetCurrentBalanceOk

`func (o *Account) GetCurrentBalanceOk() (*float32, bool)`

GetCurrentBalanceOk returns a tuple with the CurrentBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentBalance

`func (o *Account) SetCurrentBalance(v float32)`

SetCurrentBalance sets CurrentBalance field to given value.

### HasCurrentBalance

`func (o *Account) HasCurrentBalance() bool`

HasCurrentBalance returns a boolean if a field has been set.

### SetCurrentBalanceNil

`func (o *Account) SetCurrentBalanceNil(b bool)`

 SetCurrentBalanceNil sets the value for CurrentBalance to be an explicit nil

### UnsetCurrentBalance
`func (o *Account) UnsetCurrentBalance()`

UnsetCurrentBalance ensures that no value is present for CurrentBalance, not even an explicit nil
### GetCurrency

`func (o *Account) GetCurrency() CurrencyEnum`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Account) GetCurrencyOk() (*CurrencyEnum, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Account) SetCurrency(v CurrencyEnum)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *Account) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### SetCurrencyNil

`func (o *Account) SetCurrencyNil(b bool)`

 SetCurrencyNil sets the value for Currency to be an explicit nil

### UnsetCurrency
`func (o *Account) UnsetCurrency()`

UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
### GetAccountNumber

`func (o *Account) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *Account) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *Account) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *Account) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### SetAccountNumberNil

`func (o *Account) SetAccountNumberNil(b bool)`

 SetAccountNumberNil sets the value for AccountNumber to be an explicit nil

### UnsetAccountNumber
`func (o *Account) UnsetAccountNumber()`

UnsetAccountNumber ensures that no value is present for AccountNumber, not even an explicit nil
### GetParentAccount

`func (o *Account) GetParentAccount() string`

GetParentAccount returns the ParentAccount field if non-nil, zero value otherwise.

### GetParentAccountOk

`func (o *Account) GetParentAccountOk() (*string, bool)`

GetParentAccountOk returns a tuple with the ParentAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentAccount

`func (o *Account) SetParentAccount(v string)`

SetParentAccount sets ParentAccount field to given value.

### HasParentAccount

`func (o *Account) HasParentAccount() bool`

HasParentAccount returns a boolean if a field has been set.

### SetParentAccountNil

`func (o *Account) SetParentAccountNil(b bool)`

 SetParentAccountNil sets the value for ParentAccount to be an explicit nil

### UnsetParentAccount
`func (o *Account) UnsetParentAccount()`

UnsetParentAccount ensures that no value is present for ParentAccount, not even an explicit nil
### GetCompany

`func (o *Account) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *Account) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *Account) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *Account) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### SetCompanyNil

`func (o *Account) SetCompanyNil(b bool)`

 SetCompanyNil sets the value for Company to be an explicit nil

### UnsetCompany
`func (o *Account) UnsetCompany()`

UnsetCompany ensures that no value is present for Company, not even an explicit nil
### GetRemoteWasDeleted

`func (o *Account) GetRemoteWasDeleted() bool`

GetRemoteWasDeleted returns the RemoteWasDeleted field if non-nil, zero value otherwise.

### GetRemoteWasDeletedOk

`func (o *Account) GetRemoteWasDeletedOk() (*bool, bool)`

GetRemoteWasDeletedOk returns a tuple with the RemoteWasDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteWasDeleted

`func (o *Account) SetRemoteWasDeleted(v bool)`

SetRemoteWasDeleted sets RemoteWasDeleted field to given value.

### HasRemoteWasDeleted

`func (o *Account) HasRemoteWasDeleted() bool`

HasRemoteWasDeleted returns a boolean if a field has been set.

### GetFieldMappings

`func (o *Account) GetFieldMappings() map[string]interface{}`

GetFieldMappings returns the FieldMappings field if non-nil, zero value otherwise.

### GetFieldMappingsOk

`func (o *Account) GetFieldMappingsOk() (*map[string]interface{}, bool)`

GetFieldMappingsOk returns a tuple with the FieldMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldMappings

`func (o *Account) SetFieldMappings(v map[string]interface{})`

SetFieldMappings sets FieldMappings field to given value.

### HasFieldMappings

`func (o *Account) HasFieldMappings() bool`

HasFieldMappings returns a boolean if a field has been set.

### SetFieldMappingsNil

`func (o *Account) SetFieldMappingsNil(b bool)`

 SetFieldMappingsNil sets the value for FieldMappings to be an explicit nil

### UnsetFieldMappings
`func (o *Account) UnsetFieldMappings()`

UnsetFieldMappings ensures that no value is present for FieldMappings, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


