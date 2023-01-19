# Address

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**NullableAddressTypeEnum**](AddressTypeEnum.md) | The address type. | [optional] 
**Street1** | Pointer to **NullableString** | Line 1 of the address&#39;s street. | [optional] 
**Street2** | Pointer to **NullableString** | Line 2 of the address&#39;s street. | [optional] 
**City** | Pointer to **NullableString** | The address&#39;s city. | [optional] 
**State** | Pointer to **interface{}** | The address&#39;s state or region. | [optional] [readonly] 
**CountrySubdivision** | Pointer to **NullableString** | The address&#39;s state or region. | [optional] 
**Country** | Pointer to [**NullableCountryEnum**](CountryEnum.md) | The address&#39;s country. | [optional] 
**ZipCode** | Pointer to **NullableString** | The address&#39;s zip code. | [optional] 

## Methods

### NewAddress

`func NewAddress() *Address`

NewAddress instantiates a new Address object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressWithDefaults

`func NewAddressWithDefaults() *Address`

NewAddressWithDefaults instantiates a new Address object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Address) GetType() AddressTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Address) GetTypeOk() (*AddressTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Address) SetType(v AddressTypeEnum)`

SetType sets Type field to given value.

### HasType

`func (o *Address) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *Address) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *Address) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetStreet1

`func (o *Address) GetStreet1() string`

GetStreet1 returns the Street1 field if non-nil, zero value otherwise.

### GetStreet1Ok

`func (o *Address) GetStreet1Ok() (*string, bool)`

GetStreet1Ok returns a tuple with the Street1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreet1

`func (o *Address) SetStreet1(v string)`

SetStreet1 sets Street1 field to given value.

### HasStreet1

`func (o *Address) HasStreet1() bool`

HasStreet1 returns a boolean if a field has been set.

### SetStreet1Nil

`func (o *Address) SetStreet1Nil(b bool)`

 SetStreet1Nil sets the value for Street1 to be an explicit nil

### UnsetStreet1
`func (o *Address) UnsetStreet1()`

UnsetStreet1 ensures that no value is present for Street1, not even an explicit nil
### GetStreet2

`func (o *Address) GetStreet2() string`

GetStreet2 returns the Street2 field if non-nil, zero value otherwise.

### GetStreet2Ok

`func (o *Address) GetStreet2Ok() (*string, bool)`

GetStreet2Ok returns a tuple with the Street2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreet2

`func (o *Address) SetStreet2(v string)`

SetStreet2 sets Street2 field to given value.

### HasStreet2

`func (o *Address) HasStreet2() bool`

HasStreet2 returns a boolean if a field has been set.

### SetStreet2Nil

`func (o *Address) SetStreet2Nil(b bool)`

 SetStreet2Nil sets the value for Street2 to be an explicit nil

### UnsetStreet2
`func (o *Address) UnsetStreet2()`

UnsetStreet2 ensures that no value is present for Street2, not even an explicit nil
### GetCity

`func (o *Address) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *Address) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *Address) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *Address) HasCity() bool`

HasCity returns a boolean if a field has been set.

### SetCityNil

`func (o *Address) SetCityNil(b bool)`

 SetCityNil sets the value for City to be an explicit nil

### UnsetCity
`func (o *Address) UnsetCity()`

UnsetCity ensures that no value is present for City, not even an explicit nil
### GetState

`func (o *Address) GetState() interface{}`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Address) GetStateOk() (*interface{}, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Address) SetState(v interface{})`

SetState sets State field to given value.

### HasState

`func (o *Address) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *Address) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *Address) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetCountrySubdivision

`func (o *Address) GetCountrySubdivision() string`

GetCountrySubdivision returns the CountrySubdivision field if non-nil, zero value otherwise.

### GetCountrySubdivisionOk

`func (o *Address) GetCountrySubdivisionOk() (*string, bool)`

GetCountrySubdivisionOk returns a tuple with the CountrySubdivision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountrySubdivision

`func (o *Address) SetCountrySubdivision(v string)`

SetCountrySubdivision sets CountrySubdivision field to given value.

### HasCountrySubdivision

`func (o *Address) HasCountrySubdivision() bool`

HasCountrySubdivision returns a boolean if a field has been set.

### SetCountrySubdivisionNil

`func (o *Address) SetCountrySubdivisionNil(b bool)`

 SetCountrySubdivisionNil sets the value for CountrySubdivision to be an explicit nil

### UnsetCountrySubdivision
`func (o *Address) UnsetCountrySubdivision()`

UnsetCountrySubdivision ensures that no value is present for CountrySubdivision, not even an explicit nil
### GetCountry

`func (o *Address) GetCountry() CountryEnum`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *Address) GetCountryOk() (*CountryEnum, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *Address) SetCountry(v CountryEnum)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *Address) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountryNil

`func (o *Address) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *Address) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil
### GetZipCode

`func (o *Address) GetZipCode() string`

GetZipCode returns the ZipCode field if non-nil, zero value otherwise.

### GetZipCodeOk

`func (o *Address) GetZipCodeOk() (*string, bool)`

GetZipCodeOk returns a tuple with the ZipCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipCode

`func (o *Address) SetZipCode(v string)`

SetZipCode sets ZipCode field to given value.

### HasZipCode

`func (o *Address) HasZipCode() bool`

HasZipCode returns a boolean if a field has been set.

### SetZipCodeNil

`func (o *Address) SetZipCodeNil(b bool)`

 SetZipCodeNil sets the value for ZipCode to be an explicit nil

### UnsetZipCode
`func (o *Address) UnsetZipCode()`

UnsetZipCode ensures that no value is present for ZipCode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


