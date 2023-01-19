# JournalLine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RemoteId** | Pointer to **NullableString** | The third-party API ID of the matching object. | [optional] 
**Account** | Pointer to **NullableString** |  | [optional] 
**NetAmount** | Pointer to **NullableFloat32** | The value of the line item including taxes and other fees. | [optional] 
**TrackingCategory** | Pointer to **NullableString** |  | [optional] 
**TrackingCategories** | Pointer to **[]string** |  | [optional] 
**Contact** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** | The line&#39;s description. | [optional] 

## Methods

### NewJournalLine

`func NewJournalLine() *JournalLine`

NewJournalLine instantiates a new JournalLine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJournalLineWithDefaults

`func NewJournalLineWithDefaults() *JournalLine`

NewJournalLineWithDefaults instantiates a new JournalLine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRemoteId

`func (o *JournalLine) GetRemoteId() string`

GetRemoteId returns the RemoteId field if non-nil, zero value otherwise.

### GetRemoteIdOk

`func (o *JournalLine) GetRemoteIdOk() (*string, bool)`

GetRemoteIdOk returns a tuple with the RemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteId

`func (o *JournalLine) SetRemoteId(v string)`

SetRemoteId sets RemoteId field to given value.

### HasRemoteId

`func (o *JournalLine) HasRemoteId() bool`

HasRemoteId returns a boolean if a field has been set.

### SetRemoteIdNil

`func (o *JournalLine) SetRemoteIdNil(b bool)`

 SetRemoteIdNil sets the value for RemoteId to be an explicit nil

### UnsetRemoteId
`func (o *JournalLine) UnsetRemoteId()`

UnsetRemoteId ensures that no value is present for RemoteId, not even an explicit nil
### GetAccount

`func (o *JournalLine) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *JournalLine) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *JournalLine) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *JournalLine) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### SetAccountNil

`func (o *JournalLine) SetAccountNil(b bool)`

 SetAccountNil sets the value for Account to be an explicit nil

### UnsetAccount
`func (o *JournalLine) UnsetAccount()`

UnsetAccount ensures that no value is present for Account, not even an explicit nil
### GetNetAmount

`func (o *JournalLine) GetNetAmount() float32`

GetNetAmount returns the NetAmount field if non-nil, zero value otherwise.

### GetNetAmountOk

`func (o *JournalLine) GetNetAmountOk() (*float32, bool)`

GetNetAmountOk returns a tuple with the NetAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetAmount

`func (o *JournalLine) SetNetAmount(v float32)`

SetNetAmount sets NetAmount field to given value.

### HasNetAmount

`func (o *JournalLine) HasNetAmount() bool`

HasNetAmount returns a boolean if a field has been set.

### SetNetAmountNil

`func (o *JournalLine) SetNetAmountNil(b bool)`

 SetNetAmountNil sets the value for NetAmount to be an explicit nil

### UnsetNetAmount
`func (o *JournalLine) UnsetNetAmount()`

UnsetNetAmount ensures that no value is present for NetAmount, not even an explicit nil
### GetTrackingCategory

`func (o *JournalLine) GetTrackingCategory() string`

GetTrackingCategory returns the TrackingCategory field if non-nil, zero value otherwise.

### GetTrackingCategoryOk

`func (o *JournalLine) GetTrackingCategoryOk() (*string, bool)`

GetTrackingCategoryOk returns a tuple with the TrackingCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingCategory

`func (o *JournalLine) SetTrackingCategory(v string)`

SetTrackingCategory sets TrackingCategory field to given value.

### HasTrackingCategory

`func (o *JournalLine) HasTrackingCategory() bool`

HasTrackingCategory returns a boolean if a field has been set.

### SetTrackingCategoryNil

`func (o *JournalLine) SetTrackingCategoryNil(b bool)`

 SetTrackingCategoryNil sets the value for TrackingCategory to be an explicit nil

### UnsetTrackingCategory
`func (o *JournalLine) UnsetTrackingCategory()`

UnsetTrackingCategory ensures that no value is present for TrackingCategory, not even an explicit nil
### GetTrackingCategories

`func (o *JournalLine) GetTrackingCategories() []string`

GetTrackingCategories returns the TrackingCategories field if non-nil, zero value otherwise.

### GetTrackingCategoriesOk

`func (o *JournalLine) GetTrackingCategoriesOk() (*[]string, bool)`

GetTrackingCategoriesOk returns a tuple with the TrackingCategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingCategories

`func (o *JournalLine) SetTrackingCategories(v []string)`

SetTrackingCategories sets TrackingCategories field to given value.

### HasTrackingCategories

`func (o *JournalLine) HasTrackingCategories() bool`

HasTrackingCategories returns a boolean if a field has been set.

### GetContact

`func (o *JournalLine) GetContact() string`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *JournalLine) GetContactOk() (*string, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *JournalLine) SetContact(v string)`

SetContact sets Contact field to given value.

### HasContact

`func (o *JournalLine) HasContact() bool`

HasContact returns a boolean if a field has been set.

### SetContactNil

`func (o *JournalLine) SetContactNil(b bool)`

 SetContactNil sets the value for Contact to be an explicit nil

### UnsetContact
`func (o *JournalLine) UnsetContact()`

UnsetContact ensures that no value is present for Contact, not even an explicit nil
### GetDescription

`func (o *JournalLine) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *JournalLine) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *JournalLine) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *JournalLine) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *JournalLine) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *JournalLine) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


