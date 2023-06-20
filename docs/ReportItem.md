# ReportItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RemoteId** | Pointer to **NullableString** | The third-party API ID of the matching object. | [optional] 
**Name** | Pointer to **NullableString** | The report item&#39;s name. | [optional] 
**Value** | Pointer to **NullableFloat64** | The report item&#39;s value. | [optional] 
**SubItems** | Pointer to **map[string]interface{}** |  | [optional] [readonly] 
**Company** | Pointer to **NullableString** | The company the report item belongs to. | [optional] 
**ModifiedAt** | Pointer to **time.Time** | This is the datetime that this object was last updated by Merge | [optional] [readonly] 

## Methods

### NewReportItem

`func NewReportItem() *ReportItem`

NewReportItem instantiates a new ReportItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportItemWithDefaults

`func NewReportItemWithDefaults() *ReportItem`

NewReportItemWithDefaults instantiates a new ReportItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRemoteId

`func (o *ReportItem) GetRemoteId() string`

GetRemoteId returns the RemoteId field if non-nil, zero value otherwise.

### GetRemoteIdOk

`func (o *ReportItem) GetRemoteIdOk() (*string, bool)`

GetRemoteIdOk returns a tuple with the RemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteId

`func (o *ReportItem) SetRemoteId(v string)`

SetRemoteId sets RemoteId field to given value.

### HasRemoteId

`func (o *ReportItem) HasRemoteId() bool`

HasRemoteId returns a boolean if a field has been set.

### SetRemoteIdNil

`func (o *ReportItem) SetRemoteIdNil(b bool)`

 SetRemoteIdNil sets the value for RemoteId to be an explicit nil

### UnsetRemoteId
`func (o *ReportItem) UnsetRemoteId()`

UnsetRemoteId ensures that no value is present for RemoteId, not even an explicit nil
### GetName

`func (o *ReportItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ReportItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ReportItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ReportItem) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ReportItem) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ReportItem) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetValue

`func (o *ReportItem) GetValue() float64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ReportItem) GetValueOk() (*float64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ReportItem) SetValue(v float64)`

SetValue sets Value field to given value.

### HasValue

`func (o *ReportItem) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *ReportItem) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *ReportItem) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetSubItems

`func (o *ReportItem) GetSubItems() map[string]interface{}`

GetSubItems returns the SubItems field if non-nil, zero value otherwise.

### GetSubItemsOk

`func (o *ReportItem) GetSubItemsOk() (*map[string]interface{}, bool)`

GetSubItemsOk returns a tuple with the SubItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubItems

`func (o *ReportItem) SetSubItems(v map[string]interface{})`

SetSubItems sets SubItems field to given value.

### HasSubItems

`func (o *ReportItem) HasSubItems() bool`

HasSubItems returns a boolean if a field has been set.

### GetCompany

`func (o *ReportItem) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *ReportItem) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *ReportItem) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *ReportItem) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### SetCompanyNil

`func (o *ReportItem) SetCompanyNil(b bool)`

 SetCompanyNil sets the value for Company to be an explicit nil

### UnsetCompany
`func (o *ReportItem) UnsetCompany()`

UnsetCompany ensures that no value is present for Company, not even an explicit nil
### GetModifiedAt

`func (o *ReportItem) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *ReportItem) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *ReportItem) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *ReportItem) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


