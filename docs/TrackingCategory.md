# TrackingCategory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | The tracking category&#39;s name. | [optional] 
**Status** | Pointer to [**NullableStatus7d1Enum**](Status7d1Enum.md) | The tracking category&#39;s status.  * &#x60;ACTIVE&#x60; - ACTIVE * &#x60;ARCHIVED&#x60; - ARCHIVED | [optional] 
**CategoryType** | Pointer to [**NullableCategoryTypeEnum**](CategoryTypeEnum.md) | The tracking category’s type.  * &#x60;CLASS&#x60; - CLASS * &#x60;DEPARTMENT&#x60; - DEPARTMENT | [optional] 
**ParentCategory** | Pointer to **NullableString** | ID of the parent tracking category. | [optional] 
**Company** | Pointer to **NullableString** | The company the tracking category belongs to. | [optional] 
**RemoteWasDeleted** | Pointer to **bool** | Indicates whether or not this object has been deleted by third party webhooks. | [optional] [readonly] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**RemoteId** | Pointer to **NullableString** | The third-party API ID of the matching object. | [optional] 
**FieldMappings** | Pointer to **map[string]interface{}** |  | [optional] [readonly] 
**ModifiedAt** | Pointer to **time.Time** | This is the datetime that this object was last updated by Merge | [optional] [readonly] 
**RemoteData** | Pointer to [**[]RemoteData**](RemoteData.md) |  | [optional] [readonly] 

## Methods

### NewTrackingCategory

`func NewTrackingCategory() *TrackingCategory`

NewTrackingCategory instantiates a new TrackingCategory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrackingCategoryWithDefaults

`func NewTrackingCategoryWithDefaults() *TrackingCategory`

NewTrackingCategoryWithDefaults instantiates a new TrackingCategory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TrackingCategory) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TrackingCategory) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TrackingCategory) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TrackingCategory) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *TrackingCategory) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *TrackingCategory) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetStatus

`func (o *TrackingCategory) GetStatus() Status7d1Enum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TrackingCategory) GetStatusOk() (*Status7d1Enum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TrackingCategory) SetStatus(v Status7d1Enum)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TrackingCategory) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *TrackingCategory) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *TrackingCategory) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetCategoryType

`func (o *TrackingCategory) GetCategoryType() CategoryTypeEnum`

GetCategoryType returns the CategoryType field if non-nil, zero value otherwise.

### GetCategoryTypeOk

`func (o *TrackingCategory) GetCategoryTypeOk() (*CategoryTypeEnum, bool)`

GetCategoryTypeOk returns a tuple with the CategoryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryType

`func (o *TrackingCategory) SetCategoryType(v CategoryTypeEnum)`

SetCategoryType sets CategoryType field to given value.

### HasCategoryType

`func (o *TrackingCategory) HasCategoryType() bool`

HasCategoryType returns a boolean if a field has been set.

### SetCategoryTypeNil

`func (o *TrackingCategory) SetCategoryTypeNil(b bool)`

 SetCategoryTypeNil sets the value for CategoryType to be an explicit nil

### UnsetCategoryType
`func (o *TrackingCategory) UnsetCategoryType()`

UnsetCategoryType ensures that no value is present for CategoryType, not even an explicit nil
### GetParentCategory

`func (o *TrackingCategory) GetParentCategory() string`

GetParentCategory returns the ParentCategory field if non-nil, zero value otherwise.

### GetParentCategoryOk

`func (o *TrackingCategory) GetParentCategoryOk() (*string, bool)`

GetParentCategoryOk returns a tuple with the ParentCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentCategory

`func (o *TrackingCategory) SetParentCategory(v string)`

SetParentCategory sets ParentCategory field to given value.

### HasParentCategory

`func (o *TrackingCategory) HasParentCategory() bool`

HasParentCategory returns a boolean if a field has been set.

### SetParentCategoryNil

`func (o *TrackingCategory) SetParentCategoryNil(b bool)`

 SetParentCategoryNil sets the value for ParentCategory to be an explicit nil

### UnsetParentCategory
`func (o *TrackingCategory) UnsetParentCategory()`

UnsetParentCategory ensures that no value is present for ParentCategory, not even an explicit nil
### GetCompany

`func (o *TrackingCategory) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *TrackingCategory) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *TrackingCategory) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *TrackingCategory) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### SetCompanyNil

`func (o *TrackingCategory) SetCompanyNil(b bool)`

 SetCompanyNil sets the value for Company to be an explicit nil

### UnsetCompany
`func (o *TrackingCategory) UnsetCompany()`

UnsetCompany ensures that no value is present for Company, not even an explicit nil
### GetRemoteWasDeleted

`func (o *TrackingCategory) GetRemoteWasDeleted() bool`

GetRemoteWasDeleted returns the RemoteWasDeleted field if non-nil, zero value otherwise.

### GetRemoteWasDeletedOk

`func (o *TrackingCategory) GetRemoteWasDeletedOk() (*bool, bool)`

GetRemoteWasDeletedOk returns a tuple with the RemoteWasDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteWasDeleted

`func (o *TrackingCategory) SetRemoteWasDeleted(v bool)`

SetRemoteWasDeleted sets RemoteWasDeleted field to given value.

### HasRemoteWasDeleted

`func (o *TrackingCategory) HasRemoteWasDeleted() bool`

HasRemoteWasDeleted returns a boolean if a field has been set.

### GetId

`func (o *TrackingCategory) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TrackingCategory) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TrackingCategory) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TrackingCategory) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRemoteId

`func (o *TrackingCategory) GetRemoteId() string`

GetRemoteId returns the RemoteId field if non-nil, zero value otherwise.

### GetRemoteIdOk

`func (o *TrackingCategory) GetRemoteIdOk() (*string, bool)`

GetRemoteIdOk returns a tuple with the RemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteId

`func (o *TrackingCategory) SetRemoteId(v string)`

SetRemoteId sets RemoteId field to given value.

### HasRemoteId

`func (o *TrackingCategory) HasRemoteId() bool`

HasRemoteId returns a boolean if a field has been set.

### SetRemoteIdNil

`func (o *TrackingCategory) SetRemoteIdNil(b bool)`

 SetRemoteIdNil sets the value for RemoteId to be an explicit nil

### UnsetRemoteId
`func (o *TrackingCategory) UnsetRemoteId()`

UnsetRemoteId ensures that no value is present for RemoteId, not even an explicit nil
### GetFieldMappings

`func (o *TrackingCategory) GetFieldMappings() map[string]interface{}`

GetFieldMappings returns the FieldMappings field if non-nil, zero value otherwise.

### GetFieldMappingsOk

`func (o *TrackingCategory) GetFieldMappingsOk() (*map[string]interface{}, bool)`

GetFieldMappingsOk returns a tuple with the FieldMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldMappings

`func (o *TrackingCategory) SetFieldMappings(v map[string]interface{})`

SetFieldMappings sets FieldMappings field to given value.

### HasFieldMappings

`func (o *TrackingCategory) HasFieldMappings() bool`

HasFieldMappings returns a boolean if a field has been set.

### SetFieldMappingsNil

`func (o *TrackingCategory) SetFieldMappingsNil(b bool)`

 SetFieldMappingsNil sets the value for FieldMappings to be an explicit nil

### UnsetFieldMappings
`func (o *TrackingCategory) UnsetFieldMappings()`

UnsetFieldMappings ensures that no value is present for FieldMappings, not even an explicit nil
### GetModifiedAt

`func (o *TrackingCategory) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *TrackingCategory) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *TrackingCategory) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *TrackingCategory) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetRemoteData

`func (o *TrackingCategory) GetRemoteData() []RemoteData`

GetRemoteData returns the RemoteData field if non-nil, zero value otherwise.

### GetRemoteDataOk

`func (o *TrackingCategory) GetRemoteDataOk() (*[]RemoteData, bool)`

GetRemoteDataOk returns a tuple with the RemoteData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteData

`func (o *TrackingCategory) SetRemoteData(v []RemoteData)`

SetRemoteData sets RemoteData field to given value.

### HasRemoteData

`func (o *TrackingCategory) HasRemoteData() bool`

HasRemoteData returns a boolean if a field has been set.

### SetRemoteDataNil

`func (o *TrackingCategory) SetRemoteDataNil(b bool)`

 SetRemoteDataNil sets the value for RemoteData to be an explicit nil

### UnsetRemoteData
`func (o *TrackingCategory) UnsetRemoteData()`

UnsetRemoteData ensures that no value is present for RemoteData, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


