# PaginatedCashFlowStatementList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Next** | Pointer to **NullableString** |  | [optional] 
**Previous** | Pointer to **NullableString** |  | [optional] 
**Results** | Pointer to [**[]CashFlowStatement**](CashFlowStatement.md) |  | [optional] 

## Methods

### NewPaginatedCashFlowStatementList

`func NewPaginatedCashFlowStatementList() *PaginatedCashFlowStatementList`

NewPaginatedCashFlowStatementList instantiates a new PaginatedCashFlowStatementList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedCashFlowStatementListWithDefaults

`func NewPaginatedCashFlowStatementListWithDefaults() *PaginatedCashFlowStatementList`

NewPaginatedCashFlowStatementListWithDefaults instantiates a new PaginatedCashFlowStatementList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNext

`func (o *PaginatedCashFlowStatementList) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *PaginatedCashFlowStatementList) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *PaginatedCashFlowStatementList) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *PaginatedCashFlowStatementList) HasNext() bool`

HasNext returns a boolean if a field has been set.

### SetNextNil

`func (o *PaginatedCashFlowStatementList) SetNextNil(b bool)`

 SetNextNil sets the value for Next to be an explicit nil

### UnsetNext
`func (o *PaginatedCashFlowStatementList) UnsetNext()`

UnsetNext ensures that no value is present for Next, not even an explicit nil
### GetPrevious

`func (o *PaginatedCashFlowStatementList) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *PaginatedCashFlowStatementList) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *PaginatedCashFlowStatementList) SetPrevious(v string)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *PaginatedCashFlowStatementList) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.

### SetPreviousNil

`func (o *PaginatedCashFlowStatementList) SetPreviousNil(b bool)`

 SetPreviousNil sets the value for Previous to be an explicit nil

### UnsetPrevious
`func (o *PaginatedCashFlowStatementList) UnsetPrevious()`

UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
### GetResults

`func (o *PaginatedCashFlowStatementList) GetResults() []CashFlowStatement`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedCashFlowStatementList) GetResultsOk() (*[]CashFlowStatement, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedCashFlowStatementList) SetResults(v []CashFlowStatement)`

SetResults sets Results field to given value.

### HasResults

`func (o *PaginatedCashFlowStatementList) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


