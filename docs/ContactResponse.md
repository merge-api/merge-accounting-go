# ContactResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Model** | [**Contact**](Contact.md) |  | 
**Warnings** | [**[]WarningValidationProblem**](WarningValidationProblem.md) |  | 
**Errors** | [**[]ErrorValidationProblem**](ErrorValidationProblem.md) |  | 
**Logs** | Pointer to [**[]DebugModeLog**](DebugModeLog.md) |  | [optional] 

## Methods

### NewContactResponse

`func NewContactResponse(model Contact, warnings []WarningValidationProblem, errors []ErrorValidationProblem, ) *ContactResponse`

NewContactResponse instantiates a new ContactResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactResponseWithDefaults

`func NewContactResponseWithDefaults() *ContactResponse`

NewContactResponseWithDefaults instantiates a new ContactResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModel

`func (o *ContactResponse) GetModel() Contact`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ContactResponse) GetModelOk() (*Contact, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ContactResponse) SetModel(v Contact)`

SetModel sets Model field to given value.


### GetWarnings

`func (o *ContactResponse) GetWarnings() []WarningValidationProblem`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *ContactResponse) GetWarningsOk() (*[]WarningValidationProblem, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *ContactResponse) SetWarnings(v []WarningValidationProblem)`

SetWarnings sets Warnings field to given value.


### GetErrors

`func (o *ContactResponse) GetErrors() []ErrorValidationProblem`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ContactResponse) GetErrorsOk() (*[]ErrorValidationProblem, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ContactResponse) SetErrors(v []ErrorValidationProblem)`

SetErrors sets Errors field to given value.


### GetLogs

`func (o *ContactResponse) GetLogs() []DebugModeLog`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *ContactResponse) GetLogsOk() (*[]DebugModeLog, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *ContactResponse) SetLogs(v []DebugModeLog)`

SetLogs sets Logs field to given value.

### HasLogs

`func (o *ContactResponse) HasLogs() bool`

HasLogs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


