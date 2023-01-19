# ExpenseLine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RemoteId** | Pointer to **NullableString** | The third-party API ID of the matching object. | [optional] 
**Item** | Pointer to **NullableString** | The line&#39;s item. | [optional] 
**NetAmount** | Pointer to **NullableFloat32** | The line&#39;s net amount. | [optional] 
**TrackingCategory** | Pointer to **NullableString** |  | [optional] 
**TrackingCategories** | Pointer to **[]string** |  | [optional] 
**Company** | Pointer to **NullableString** | The company the line belongs to. | [optional] 
**Account** | Pointer to **NullableString** | The expense&#39;s payment account. | [optional] 
**Description** | Pointer to **NullableString** | The description of the item that was purchased by the company. | [optional] 

## Methods

### NewExpenseLine

`func NewExpenseLine() *ExpenseLine`

NewExpenseLine instantiates a new ExpenseLine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExpenseLineWithDefaults

`func NewExpenseLineWithDefaults() *ExpenseLine`

NewExpenseLineWithDefaults instantiates a new ExpenseLine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRemoteId

`func (o *ExpenseLine) GetRemoteId() string`

GetRemoteId returns the RemoteId field if non-nil, zero value otherwise.

### GetRemoteIdOk

`func (o *ExpenseLine) GetRemoteIdOk() (*string, bool)`

GetRemoteIdOk returns a tuple with the RemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteId

`func (o *ExpenseLine) SetRemoteId(v string)`

SetRemoteId sets RemoteId field to given value.

### HasRemoteId

`func (o *ExpenseLine) HasRemoteId() bool`

HasRemoteId returns a boolean if a field has been set.

### SetRemoteIdNil

`func (o *ExpenseLine) SetRemoteIdNil(b bool)`

 SetRemoteIdNil sets the value for RemoteId to be an explicit nil

### UnsetRemoteId
`func (o *ExpenseLine) UnsetRemoteId()`

UnsetRemoteId ensures that no value is present for RemoteId, not even an explicit nil
### GetItem

`func (o *ExpenseLine) GetItem() string`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *ExpenseLine) GetItemOk() (*string, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *ExpenseLine) SetItem(v string)`

SetItem sets Item field to given value.

### HasItem

`func (o *ExpenseLine) HasItem() bool`

HasItem returns a boolean if a field has been set.

### SetItemNil

`func (o *ExpenseLine) SetItemNil(b bool)`

 SetItemNil sets the value for Item to be an explicit nil

### UnsetItem
`func (o *ExpenseLine) UnsetItem()`

UnsetItem ensures that no value is present for Item, not even an explicit nil
### GetNetAmount

`func (o *ExpenseLine) GetNetAmount() float32`

GetNetAmount returns the NetAmount field if non-nil, zero value otherwise.

### GetNetAmountOk

`func (o *ExpenseLine) GetNetAmountOk() (*float32, bool)`

GetNetAmountOk returns a tuple with the NetAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetAmount

`func (o *ExpenseLine) SetNetAmount(v float32)`

SetNetAmount sets NetAmount field to given value.

### HasNetAmount

`func (o *ExpenseLine) HasNetAmount() bool`

HasNetAmount returns a boolean if a field has been set.

### SetNetAmountNil

`func (o *ExpenseLine) SetNetAmountNil(b bool)`

 SetNetAmountNil sets the value for NetAmount to be an explicit nil

### UnsetNetAmount
`func (o *ExpenseLine) UnsetNetAmount()`

UnsetNetAmount ensures that no value is present for NetAmount, not even an explicit nil
### GetTrackingCategory

`func (o *ExpenseLine) GetTrackingCategory() string`

GetTrackingCategory returns the TrackingCategory field if non-nil, zero value otherwise.

### GetTrackingCategoryOk

`func (o *ExpenseLine) GetTrackingCategoryOk() (*string, bool)`

GetTrackingCategoryOk returns a tuple with the TrackingCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingCategory

`func (o *ExpenseLine) SetTrackingCategory(v string)`

SetTrackingCategory sets TrackingCategory field to given value.

### HasTrackingCategory

`func (o *ExpenseLine) HasTrackingCategory() bool`

HasTrackingCategory returns a boolean if a field has been set.

### SetTrackingCategoryNil

`func (o *ExpenseLine) SetTrackingCategoryNil(b bool)`

 SetTrackingCategoryNil sets the value for TrackingCategory to be an explicit nil

### UnsetTrackingCategory
`func (o *ExpenseLine) UnsetTrackingCategory()`

UnsetTrackingCategory ensures that no value is present for TrackingCategory, not even an explicit nil
### GetTrackingCategories

`func (o *ExpenseLine) GetTrackingCategories() []string`

GetTrackingCategories returns the TrackingCategories field if non-nil, zero value otherwise.

### GetTrackingCategoriesOk

`func (o *ExpenseLine) GetTrackingCategoriesOk() (*[]string, bool)`

GetTrackingCategoriesOk returns a tuple with the TrackingCategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingCategories

`func (o *ExpenseLine) SetTrackingCategories(v []string)`

SetTrackingCategories sets TrackingCategories field to given value.

### HasTrackingCategories

`func (o *ExpenseLine) HasTrackingCategories() bool`

HasTrackingCategories returns a boolean if a field has been set.

### GetCompany

`func (o *ExpenseLine) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *ExpenseLine) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *ExpenseLine) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *ExpenseLine) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### SetCompanyNil

`func (o *ExpenseLine) SetCompanyNil(b bool)`

 SetCompanyNil sets the value for Company to be an explicit nil

### UnsetCompany
`func (o *ExpenseLine) UnsetCompany()`

UnsetCompany ensures that no value is present for Company, not even an explicit nil
### GetAccount

`func (o *ExpenseLine) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ExpenseLine) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ExpenseLine) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ExpenseLine) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### SetAccountNil

`func (o *ExpenseLine) SetAccountNil(b bool)`

 SetAccountNil sets the value for Account to be an explicit nil

### UnsetAccount
`func (o *ExpenseLine) UnsetAccount()`

UnsetAccount ensures that no value is present for Account, not even an explicit nil
### GetDescription

`func (o *ExpenseLine) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ExpenseLine) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ExpenseLine) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ExpenseLine) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ExpenseLine) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ExpenseLine) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


