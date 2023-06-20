/*
 * Merge Accounting API
 *
 * The unified API for building rich integrations with multiple Accounting & Finance platforms.
 *
 * API version: 1.0
 * Contact: hello@merge.dev
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package merge_accounting_client

import (
	"encoding/json"
	"time"
)

// TrackingCategory # The TrackingCategory Object ### Description The `TrackingCategory` object is used to represent a company's tracking categories.  ### Usage Example Fetch from the `GET TrackingCategory` endpoint and view a company's tracking category.
type TrackingCategory struct {
	// The tracking category's name.
	Name NullableString `json:"name,omitempty"`
	// The tracking category's status.  * `ACTIVE` - ACTIVE * `ARCHIVED` - ARCHIVED
	Status NullableStatus7d1Enum `json:"status,omitempty"`
	// The tracking category’s type.  * `CLASS` - CLASS * `DEPARTMENT` - DEPARTMENT
	CategoryType NullableCategoryTypeEnum `json:"category_type,omitempty"`
	// ID of the parent tracking category.
	ParentCategory NullableString `json:"parent_category,omitempty"`
	// The company the tracking category belongs to.
	Company NullableString `json:"company,omitempty"`
	// Indicates whether or not this object has been deleted by third party webhooks.
	RemoteWasDeleted *bool `json:"remote_was_deleted,omitempty"`
	Id *string `json:"id,omitempty"`
	// The third-party API ID of the matching object.
	RemoteId NullableString `json:"remote_id,omitempty"`
	FieldMappings map[string]interface{} `json:"field_mappings,omitempty"`
	// This is the datetime that this object was last updated by Merge
	ModifiedAt *time.Time `json:"modified_at,omitempty"`
	RemoteData []RemoteData `json:"remote_data,omitempty"`
	// raw json response by property name
	ResponseRaw map[string]json.RawMessage `json:"-"`
}

// NewTrackingCategory instantiates a new TrackingCategory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTrackingCategory() *TrackingCategory {
	this := TrackingCategory{}
	return &this
}

// NewTrackingCategoryWithDefaults instantiates a new TrackingCategory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTrackingCategoryWithDefaults() *TrackingCategory {
	this := TrackingCategory{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TrackingCategory) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TrackingCategory) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *TrackingCategory) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *TrackingCategory) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *TrackingCategory) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *TrackingCategory) UnsetName() {
	o.Name.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TrackingCategory) GetStatus() Status7d1Enum {
	if o == nil || o.Status.Get() == nil {
		var ret Status7d1Enum
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TrackingCategory) GetStatusOk() (*Status7d1Enum, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *TrackingCategory) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableStatus7d1Enum and assigns it to the Status field.
func (o *TrackingCategory) SetStatus(v Status7d1Enum) {
	o.Status.Set(&v)
}
// SetStatusNil sets the value for Status to be an explicit nil
func (o *TrackingCategory) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *TrackingCategory) UnsetStatus() {
	o.Status.Unset()
}

// GetCategoryType returns the CategoryType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TrackingCategory) GetCategoryType() CategoryTypeEnum {
	if o == nil || o.CategoryType.Get() == nil {
		var ret CategoryTypeEnum
		return ret
	}
	return *o.CategoryType.Get()
}

// GetCategoryTypeOk returns a tuple with the CategoryType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TrackingCategory) GetCategoryTypeOk() (*CategoryTypeEnum, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CategoryType.Get(), o.CategoryType.IsSet()
}

// HasCategoryType returns a boolean if a field has been set.
func (o *TrackingCategory) HasCategoryType() bool {
	if o != nil && o.CategoryType.IsSet() {
		return true
	}

	return false
}

// SetCategoryType gets a reference to the given NullableCategoryTypeEnum and assigns it to the CategoryType field.
func (o *TrackingCategory) SetCategoryType(v CategoryTypeEnum) {
	o.CategoryType.Set(&v)
}
// SetCategoryTypeNil sets the value for CategoryType to be an explicit nil
func (o *TrackingCategory) SetCategoryTypeNil() {
	o.CategoryType.Set(nil)
}

// UnsetCategoryType ensures that no value is present for CategoryType, not even an explicit nil
func (o *TrackingCategory) UnsetCategoryType() {
	o.CategoryType.Unset()
}

// GetParentCategory returns the ParentCategory field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TrackingCategory) GetParentCategory() string {
	if o == nil || o.ParentCategory.Get() == nil {
		var ret string
		return ret
	}
	return *o.ParentCategory.Get()
}

// GetParentCategoryOk returns a tuple with the ParentCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TrackingCategory) GetParentCategoryOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ParentCategory.Get(), o.ParentCategory.IsSet()
}

// HasParentCategory returns a boolean if a field has been set.
func (o *TrackingCategory) HasParentCategory() bool {
	if o != nil && o.ParentCategory.IsSet() {
		return true
	}

	return false
}

// SetParentCategory gets a reference to the given NullableString and assigns it to the ParentCategory field.
func (o *TrackingCategory) SetParentCategory(v string) {
	o.ParentCategory.Set(&v)
}
// SetParentCategoryNil sets the value for ParentCategory to be an explicit nil
func (o *TrackingCategory) SetParentCategoryNil() {
	o.ParentCategory.Set(nil)
}

// UnsetParentCategory ensures that no value is present for ParentCategory, not even an explicit nil
func (o *TrackingCategory) UnsetParentCategory() {
	o.ParentCategory.Unset()
}

// GetCompany returns the Company field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TrackingCategory) GetCompany() string {
	if o == nil || o.Company.Get() == nil {
		var ret string
		return ret
	}
	return *o.Company.Get()
}

// GetCompanyOk returns a tuple with the Company field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TrackingCategory) GetCompanyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Company.Get(), o.Company.IsSet()
}

// HasCompany returns a boolean if a field has been set.
func (o *TrackingCategory) HasCompany() bool {
	if o != nil && o.Company.IsSet() {
		return true
	}

	return false
}

// SetCompany gets a reference to the given NullableString and assigns it to the Company field.
func (o *TrackingCategory) SetCompany(v string) {
	o.Company.Set(&v)
}
// SetCompanyNil sets the value for Company to be an explicit nil
func (o *TrackingCategory) SetCompanyNil() {
	o.Company.Set(nil)
}

// UnsetCompany ensures that no value is present for Company, not even an explicit nil
func (o *TrackingCategory) UnsetCompany() {
	o.Company.Unset()
}

// GetRemoteWasDeleted returns the RemoteWasDeleted field value if set, zero value otherwise.
func (o *TrackingCategory) GetRemoteWasDeleted() bool {
	if o == nil || o.RemoteWasDeleted == nil {
		var ret bool
		return ret
	}
	return *o.RemoteWasDeleted
}

// GetRemoteWasDeletedOk returns a tuple with the RemoteWasDeleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackingCategory) GetRemoteWasDeletedOk() (*bool, bool) {
	if o == nil || o.RemoteWasDeleted == nil {
		return nil, false
	}
	return o.RemoteWasDeleted, true
}

// HasRemoteWasDeleted returns a boolean if a field has been set.
func (o *TrackingCategory) HasRemoteWasDeleted() bool {
	if o != nil && o.RemoteWasDeleted != nil {
		return true
	}

	return false
}

// SetRemoteWasDeleted gets a reference to the given bool and assigns it to the RemoteWasDeleted field.
func (o *TrackingCategory) SetRemoteWasDeleted(v bool) {
	o.RemoteWasDeleted = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TrackingCategory) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackingCategory) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TrackingCategory) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TrackingCategory) SetId(v string) {
	o.Id = &v
}

// GetRemoteId returns the RemoteId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TrackingCategory) GetRemoteId() string {
	if o == nil || o.RemoteId.Get() == nil {
		var ret string
		return ret
	}
	return *o.RemoteId.Get()
}

// GetRemoteIdOk returns a tuple with the RemoteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TrackingCategory) GetRemoteIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RemoteId.Get(), o.RemoteId.IsSet()
}

// HasRemoteId returns a boolean if a field has been set.
func (o *TrackingCategory) HasRemoteId() bool {
	if o != nil && o.RemoteId.IsSet() {
		return true
	}

	return false
}

// SetRemoteId gets a reference to the given NullableString and assigns it to the RemoteId field.
func (o *TrackingCategory) SetRemoteId(v string) {
	o.RemoteId.Set(&v)
}
// SetRemoteIdNil sets the value for RemoteId to be an explicit nil
func (o *TrackingCategory) SetRemoteIdNil() {
	o.RemoteId.Set(nil)
}

// UnsetRemoteId ensures that no value is present for RemoteId, not even an explicit nil
func (o *TrackingCategory) UnsetRemoteId() {
	o.RemoteId.Unset()
}

// GetFieldMappings returns the FieldMappings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TrackingCategory) GetFieldMappings() map[string]interface{} {
	if o == nil  {
		var ret map[string]interface{}
		return ret
	}
	return o.FieldMappings
}

// GetFieldMappingsOk returns a tuple with the FieldMappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TrackingCategory) GetFieldMappingsOk() (*map[string]interface{}, bool) {
	if o == nil || o.FieldMappings == nil {
		return nil, false
	}
	return &o.FieldMappings, true
}

// HasFieldMappings returns a boolean if a field has been set.
func (o *TrackingCategory) HasFieldMappings() bool {
	if o != nil && o.FieldMappings != nil {
		return true
	}

	return false
}

// SetFieldMappings gets a reference to the given map[string]interface{} and assigns it to the FieldMappings field.
func (o *TrackingCategory) SetFieldMappings(v map[string]interface{}) {
	o.FieldMappings = v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise.
func (o *TrackingCategory) GetModifiedAt() time.Time {
	if o == nil || o.ModifiedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackingCategory) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil || o.ModifiedAt == nil {
		return nil, false
	}
	return o.ModifiedAt, true
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *TrackingCategory) HasModifiedAt() bool {
	if o != nil && o.ModifiedAt != nil {
		return true
	}

	return false
}

// SetModifiedAt gets a reference to the given time.Time and assigns it to the ModifiedAt field.
func (o *TrackingCategory) SetModifiedAt(v time.Time) {
	o.ModifiedAt = &v
}

// GetRemoteData returns the RemoteData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TrackingCategory) GetRemoteData() []RemoteData {
	if o == nil  {
		var ret []RemoteData
		return ret
	}
	return o.RemoteData
}

// GetRemoteDataOk returns a tuple with the RemoteData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TrackingCategory) GetRemoteDataOk() (*[]RemoteData, bool) {
	if o == nil || o.RemoteData == nil {
		return nil, false
	}
	return &o.RemoteData, true
}

// HasRemoteData returns a boolean if a field has been set.
func (o *TrackingCategory) HasRemoteData() bool {
	if o != nil && o.RemoteData != nil {
		return true
	}

	return false
}

// SetRemoteData gets a reference to the given []RemoteData and assigns it to the RemoteData field.
func (o *TrackingCategory) SetRemoteData(v []RemoteData) {
	o.RemoteData = v
}

func (o TrackingCategory) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
	}
	if o.CategoryType.IsSet() {
		toSerialize["category_type"] = o.CategoryType.Get()
	}
	if o.ParentCategory.IsSet() {
		toSerialize["parent_category"] = o.ParentCategory.Get()
	}
	if o.Company.IsSet() {
		toSerialize["company"] = o.Company.Get()
	}
	if o.RemoteWasDeleted != nil {
		toSerialize["remote_was_deleted"] = o.RemoteWasDeleted
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.RemoteId.IsSet() {
		toSerialize["remote_id"] = o.RemoteId.Get()
	}
	if o.FieldMappings != nil {
		toSerialize["field_mappings"] = o.FieldMappings
	}
	if o.ModifiedAt != nil {
		toSerialize["modified_at"] = o.ModifiedAt
	}
	if o.RemoteData != nil {
		toSerialize["remote_data"] = o.RemoteData
	}
	return json.Marshal(toSerialize)
}

func (v *TrackingCategory) UnmarshalJSON(src []byte) error {
    type TrackingCategoryUnmarshalTarget TrackingCategory

	var intermediateResult TrackingCategoryUnmarshalTarget
	var err1 = json.Unmarshal(src, &intermediateResult)
    if err1 != nil {
        return err1
    }
    var err2 = json.Unmarshal(src, &intermediateResult.ResponseRaw)
	if err2 != nil {
		return err2
	}

	*v = TrackingCategory(intermediateResult)
	return nil
}
type NullableTrackingCategory struct {
	value *TrackingCategory
	isSet bool
}

func (v NullableTrackingCategory) Get() *TrackingCategory {
	return v.value
}

func (v *NullableTrackingCategory) Set(val *TrackingCategory) {
	v.value = val
	v.isSet = true
}

func (v NullableTrackingCategory) IsSet() bool {
	return v.isSet
}

func (v *NullableTrackingCategory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrackingCategory(val *TrackingCategory) *NullableTrackingCategory {
	return &NullableTrackingCategory{value: val, isSet: true}
}

func (v NullableTrackingCategory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrackingCategory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	var err1 = json.Unmarshal(src, &v.value)
    if err1 != nil {
        return err1
    }
    return json.Unmarshal(src, &v.value.ResponseRaw)
}


