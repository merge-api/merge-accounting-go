# VendorCreditLine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RemoteId** | Pointer to **NullableString** | The third-party API ID of the matching object. | [optional] 
**NetAmount** | Pointer to **NullableFloat32** | The full value of the credit. | [optional] 
**TrackingCategory** | Pointer to **NullableString** | The line&#39;s associated tracking category. | [optional] 
**TrackingCategories** | **[]string** | The line&#39;s associated tracking categories. | 
**Description** | Pointer to **NullableString** | The line&#39;s description. | [optional] 
**Account** | Pointer to **NullableString** | The line&#39;s account. | [optional] 
**Company** | Pointer to **NullableString** | The company the line belongs to. | [optional] 

## Methods

### NewVendorCreditLine

`func NewVendorCreditLine(trackingCategories []string, ) *VendorCreditLine`

NewVendorCreditLine instantiates a new VendorCreditLine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVendorCreditLineWithDefaults

`func NewVendorCreditLineWithDefaults() *VendorCreditLine`

NewVendorCreditLineWithDefaults instantiates a new VendorCreditLine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRemoteId

`func (o *VendorCreditLine) GetRemoteId() string`

GetRemoteId returns the RemoteId field if non-nil, zero value otherwise.

### GetRemoteIdOk

`func (o *VendorCreditLine) GetRemoteIdOk() (*string, bool)`

GetRemoteIdOk returns a tuple with the RemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteId

`func (o *VendorCreditLine) SetRemoteId(v string)`

SetRemoteId sets RemoteId field to given value.

### HasRemoteId

`func (o *VendorCreditLine) HasRemoteId() bool`

HasRemoteId returns a boolean if a field has been set.

### SetRemoteIdNil

`func (o *VendorCreditLine) SetRemoteIdNil(b bool)`

 SetRemoteIdNil sets the value for RemoteId to be an explicit nil

### UnsetRemoteId
`func (o *VendorCreditLine) UnsetRemoteId()`

UnsetRemoteId ensures that no value is present for RemoteId, not even an explicit nil
### GetNetAmount

`func (o *VendorCreditLine) GetNetAmount() float32`

GetNetAmount returns the NetAmount field if non-nil, zero value otherwise.

### GetNetAmountOk

`func (o *VendorCreditLine) GetNetAmountOk() (*float32, bool)`

GetNetAmountOk returns a tuple with the NetAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetAmount

`func (o *VendorCreditLine) SetNetAmount(v float32)`

SetNetAmount sets NetAmount field to given value.

### HasNetAmount

`func (o *VendorCreditLine) HasNetAmount() bool`

HasNetAmount returns a boolean if a field has been set.

### SetNetAmountNil

`func (o *VendorCreditLine) SetNetAmountNil(b bool)`

 SetNetAmountNil sets the value for NetAmount to be an explicit nil

### UnsetNetAmount
`func (o *VendorCreditLine) UnsetNetAmount()`

UnsetNetAmount ensures that no value is present for NetAmount, not even an explicit nil
### GetTrackingCategory

`func (o *VendorCreditLine) GetTrackingCategory() string`

GetTrackingCategory returns the TrackingCategory field if non-nil, zero value otherwise.

### GetTrackingCategoryOk

`func (o *VendorCreditLine) GetTrackingCategoryOk() (*string, bool)`

GetTrackingCategoryOk returns a tuple with the TrackingCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingCategory

`func (o *VendorCreditLine) SetTrackingCategory(v string)`

SetTrackingCategory sets TrackingCategory field to given value.

### HasTrackingCategory

`func (o *VendorCreditLine) HasTrackingCategory() bool`

HasTrackingCategory returns a boolean if a field has been set.

### SetTrackingCategoryNil

`func (o *VendorCreditLine) SetTrackingCategoryNil(b bool)`

 SetTrackingCategoryNil sets the value for TrackingCategory to be an explicit nil

### UnsetTrackingCategory
`func (o *VendorCreditLine) UnsetTrackingCategory()`

UnsetTrackingCategory ensures that no value is present for TrackingCategory, not even an explicit nil
### GetTrackingCategories

`func (o *VendorCreditLine) GetTrackingCategories() []string`

GetTrackingCategories returns the TrackingCategories field if non-nil, zero value otherwise.

### GetTrackingCategoriesOk

`func (o *VendorCreditLine) GetTrackingCategoriesOk() (*[]string, bool)`

GetTrackingCategoriesOk returns a tuple with the TrackingCategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingCategories

`func (o *VendorCreditLine) SetTrackingCategories(v []string)`

SetTrackingCategories sets TrackingCategories field to given value.


### GetDescription

`func (o *VendorCreditLine) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VendorCreditLine) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VendorCreditLine) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VendorCreditLine) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *VendorCreditLine) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *VendorCreditLine) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetAccount

`func (o *VendorCreditLine) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *VendorCreditLine) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *VendorCreditLine) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *VendorCreditLine) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### SetAccountNil

`func (o *VendorCreditLine) SetAccountNil(b bool)`

 SetAccountNil sets the value for Account to be an explicit nil

### UnsetAccount
`func (o *VendorCreditLine) UnsetAccount()`

UnsetAccount ensures that no value is present for Account, not even an explicit nil
### GetCompany

`func (o *VendorCreditLine) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *VendorCreditLine) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *VendorCreditLine) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *VendorCreditLine) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### SetCompanyNil

`func (o *VendorCreditLine) SetCompanyNil(b bool)`

 SetCompanyNil sets the value for Company to be an explicit nil

### UnsetCompany
`func (o *VendorCreditLine) UnsetCompany()`

UnsetCompany ensures that no value is present for Company, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


