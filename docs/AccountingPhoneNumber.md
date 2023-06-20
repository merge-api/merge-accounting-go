# AccountingPhoneNumber

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | Pointer to **NullableString** | The phone number. | [optional] 
**Type** | Pointer to **NullableString** | The phone number&#39;s type. | [optional] 
**ModifiedAt** | Pointer to **time.Time** | This is the datetime that this object was last updated by Merge | [optional] [readonly] 

## Methods

### NewAccountingPhoneNumber

`func NewAccountingPhoneNumber() *AccountingPhoneNumber`

NewAccountingPhoneNumber instantiates a new AccountingPhoneNumber object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountingPhoneNumberWithDefaults

`func NewAccountingPhoneNumberWithDefaults() *AccountingPhoneNumber`

NewAccountingPhoneNumberWithDefaults instantiates a new AccountingPhoneNumber object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *AccountingPhoneNumber) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *AccountingPhoneNumber) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *AccountingPhoneNumber) SetNumber(v string)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *AccountingPhoneNumber) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### SetNumberNil

`func (o *AccountingPhoneNumber) SetNumberNil(b bool)`

 SetNumberNil sets the value for Number to be an explicit nil

### UnsetNumber
`func (o *AccountingPhoneNumber) UnsetNumber()`

UnsetNumber ensures that no value is present for Number, not even an explicit nil
### GetType

`func (o *AccountingPhoneNumber) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AccountingPhoneNumber) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AccountingPhoneNumber) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AccountingPhoneNumber) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *AccountingPhoneNumber) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *AccountingPhoneNumber) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetModifiedAt

`func (o *AccountingPhoneNumber) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *AccountingPhoneNumber) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *AccountingPhoneNumber) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *AccountingPhoneNumber) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


