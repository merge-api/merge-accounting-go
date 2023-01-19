# AccountingPhoneNumberRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | Pointer to **NullableString** | The phone number. | [optional] 
**Type** | Pointer to **NullableString** | The phone number&#39;s type. | [optional] 
**IntegrationParams** | Pointer to **map[string]interface{}** |  | [optional] 
**LinkedAccountParams** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewAccountingPhoneNumberRequest

`func NewAccountingPhoneNumberRequest() *AccountingPhoneNumberRequest`

NewAccountingPhoneNumberRequest instantiates a new AccountingPhoneNumberRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountingPhoneNumberRequestWithDefaults

`func NewAccountingPhoneNumberRequestWithDefaults() *AccountingPhoneNumberRequest`

NewAccountingPhoneNumberRequestWithDefaults instantiates a new AccountingPhoneNumberRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *AccountingPhoneNumberRequest) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *AccountingPhoneNumberRequest) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *AccountingPhoneNumberRequest) SetNumber(v string)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *AccountingPhoneNumberRequest) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### SetNumberNil

`func (o *AccountingPhoneNumberRequest) SetNumberNil(b bool)`

 SetNumberNil sets the value for Number to be an explicit nil

### UnsetNumber
`func (o *AccountingPhoneNumberRequest) UnsetNumber()`

UnsetNumber ensures that no value is present for Number, not even an explicit nil
### GetType

`func (o *AccountingPhoneNumberRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AccountingPhoneNumberRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AccountingPhoneNumberRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AccountingPhoneNumberRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *AccountingPhoneNumberRequest) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *AccountingPhoneNumberRequest) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetIntegrationParams

`func (o *AccountingPhoneNumberRequest) GetIntegrationParams() map[string]interface{}`

GetIntegrationParams returns the IntegrationParams field if non-nil, zero value otherwise.

### GetIntegrationParamsOk

`func (o *AccountingPhoneNumberRequest) GetIntegrationParamsOk() (*map[string]interface{}, bool)`

GetIntegrationParamsOk returns a tuple with the IntegrationParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationParams

`func (o *AccountingPhoneNumberRequest) SetIntegrationParams(v map[string]interface{})`

SetIntegrationParams sets IntegrationParams field to given value.

### HasIntegrationParams

`func (o *AccountingPhoneNumberRequest) HasIntegrationParams() bool`

HasIntegrationParams returns a boolean if a field has been set.

### SetIntegrationParamsNil

`func (o *AccountingPhoneNumberRequest) SetIntegrationParamsNil(b bool)`

 SetIntegrationParamsNil sets the value for IntegrationParams to be an explicit nil

### UnsetIntegrationParams
`func (o *AccountingPhoneNumberRequest) UnsetIntegrationParams()`

UnsetIntegrationParams ensures that no value is present for IntegrationParams, not even an explicit nil
### GetLinkedAccountParams

`func (o *AccountingPhoneNumberRequest) GetLinkedAccountParams() map[string]interface{}`

GetLinkedAccountParams returns the LinkedAccountParams field if non-nil, zero value otherwise.

### GetLinkedAccountParamsOk

`func (o *AccountingPhoneNumberRequest) GetLinkedAccountParamsOk() (*map[string]interface{}, bool)`

GetLinkedAccountParamsOk returns a tuple with the LinkedAccountParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedAccountParams

`func (o *AccountingPhoneNumberRequest) SetLinkedAccountParams(v map[string]interface{})`

SetLinkedAccountParams sets LinkedAccountParams field to given value.

### HasLinkedAccountParams

`func (o *AccountingPhoneNumberRequest) HasLinkedAccountParams() bool`

HasLinkedAccountParams returns a boolean if a field has been set.

### SetLinkedAccountParamsNil

`func (o *AccountingPhoneNumberRequest) SetLinkedAccountParamsNil(b bool)`

 SetLinkedAccountParamsNil sets the value for LinkedAccountParams to be an explicit nil

### UnsetLinkedAccountParams
`func (o *AccountingPhoneNumberRequest) UnsetLinkedAccountParams()`

UnsetLinkedAccountParams ensures that no value is present for LinkedAccountParams, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


