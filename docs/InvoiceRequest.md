# InvoiceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RemoteId** | Pointer to **NullableString** | The third-party API ID of the matching object. | [optional] 
**Type** | Pointer to [**NullableInvoiceTypeEnum**](InvoiceTypeEnum.md) | Whether the invoice is an accounts receivable or accounts payable. Accounts payable invoices are commonly referred to as Bills. | [optional] 
**Contact** | Pointer to **NullableString** | The invoice&#39;s contact. | [optional] 
**Number** | Pointer to **NullableString** | The invoice&#39;s number. | [optional] 
**IssueDate** | Pointer to **NullableTime** | The invoice&#39;s issue date. | [optional] 
**DueDate** | Pointer to **NullableTime** | The invoice&#39;s due date. | [optional] 
**PaidOnDate** | Pointer to **NullableTime** | The invoice&#39;s paid date. | [optional] 
**Memo** | Pointer to **NullableString** | The invoice&#39;s private note. | [optional] 
**Company** | Pointer to **NullableString** | The company the invoice belongs to. | [optional] 
**Currency** | Pointer to [**NullableCurrencyEnum**](CurrencyEnum.md) | The invoice&#39;s currency. | [optional] 
**ExchangeRate** | Pointer to **NullableFloat64** | The invoice&#39;s exchange rate. | [optional] 
**TotalDiscount** | Pointer to **NullableFloat32** | The total discounts applied to the total cost. | [optional] 
**SubTotal** | Pointer to **NullableFloat32** | The total amount being paid before taxes. | [optional] 
**TotalTaxAmount** | Pointer to **NullableFloat32** | The total amount being paid in taxes. | [optional] 
**TotalAmount** | Pointer to **NullableFloat32** | The invoice&#39;s total amount. | [optional] 
**Balance** | Pointer to **NullableFloat32** | The invoice&#39;s remaining balance. | [optional] 
**RemoteUpdatedAt** | Pointer to **NullableTime** | When the third party&#39;s invoice entry was updated. | [optional] 
**Payments** | Pointer to **[]string** | Array of &#x60;Payment&#x60; object IDs. | [optional] 
**LineItems** | Pointer to [**[]InvoiceLineItemRequest**](InvoiceLineItemRequest.md) |  | [optional] 
**IntegrationParams** | Pointer to **map[string]interface{}** |  | [optional] 
**LinkedAccountParams** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewInvoiceRequest

`func NewInvoiceRequest() *InvoiceRequest`

NewInvoiceRequest instantiates a new InvoiceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceRequestWithDefaults

`func NewInvoiceRequestWithDefaults() *InvoiceRequest`

NewInvoiceRequestWithDefaults instantiates a new InvoiceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRemoteId

`func (o *InvoiceRequest) GetRemoteId() string`

GetRemoteId returns the RemoteId field if non-nil, zero value otherwise.

### GetRemoteIdOk

`func (o *InvoiceRequest) GetRemoteIdOk() (*string, bool)`

GetRemoteIdOk returns a tuple with the RemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteId

`func (o *InvoiceRequest) SetRemoteId(v string)`

SetRemoteId sets RemoteId field to given value.

### HasRemoteId

`func (o *InvoiceRequest) HasRemoteId() bool`

HasRemoteId returns a boolean if a field has been set.

### SetRemoteIdNil

`func (o *InvoiceRequest) SetRemoteIdNil(b bool)`

 SetRemoteIdNil sets the value for RemoteId to be an explicit nil

### UnsetRemoteId
`func (o *InvoiceRequest) UnsetRemoteId()`

UnsetRemoteId ensures that no value is present for RemoteId, not even an explicit nil
### GetType

`func (o *InvoiceRequest) GetType() InvoiceTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InvoiceRequest) GetTypeOk() (*InvoiceTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InvoiceRequest) SetType(v InvoiceTypeEnum)`

SetType sets Type field to given value.

### HasType

`func (o *InvoiceRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *InvoiceRequest) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *InvoiceRequest) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetContact

`func (o *InvoiceRequest) GetContact() string`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *InvoiceRequest) GetContactOk() (*string, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *InvoiceRequest) SetContact(v string)`

SetContact sets Contact field to given value.

### HasContact

`func (o *InvoiceRequest) HasContact() bool`

HasContact returns a boolean if a field has been set.

### SetContactNil

`func (o *InvoiceRequest) SetContactNil(b bool)`

 SetContactNil sets the value for Contact to be an explicit nil

### UnsetContact
`func (o *InvoiceRequest) UnsetContact()`

UnsetContact ensures that no value is present for Contact, not even an explicit nil
### GetNumber

`func (o *InvoiceRequest) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *InvoiceRequest) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *InvoiceRequest) SetNumber(v string)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *InvoiceRequest) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### SetNumberNil

`func (o *InvoiceRequest) SetNumberNil(b bool)`

 SetNumberNil sets the value for Number to be an explicit nil

### UnsetNumber
`func (o *InvoiceRequest) UnsetNumber()`

UnsetNumber ensures that no value is present for Number, not even an explicit nil
### GetIssueDate

`func (o *InvoiceRequest) GetIssueDate() time.Time`

GetIssueDate returns the IssueDate field if non-nil, zero value otherwise.

### GetIssueDateOk

`func (o *InvoiceRequest) GetIssueDateOk() (*time.Time, bool)`

GetIssueDateOk returns a tuple with the IssueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueDate

`func (o *InvoiceRequest) SetIssueDate(v time.Time)`

SetIssueDate sets IssueDate field to given value.

### HasIssueDate

`func (o *InvoiceRequest) HasIssueDate() bool`

HasIssueDate returns a boolean if a field has been set.

### SetIssueDateNil

`func (o *InvoiceRequest) SetIssueDateNil(b bool)`

 SetIssueDateNil sets the value for IssueDate to be an explicit nil

### UnsetIssueDate
`func (o *InvoiceRequest) UnsetIssueDate()`

UnsetIssueDate ensures that no value is present for IssueDate, not even an explicit nil
### GetDueDate

`func (o *InvoiceRequest) GetDueDate() time.Time`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *InvoiceRequest) GetDueDateOk() (*time.Time, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *InvoiceRequest) SetDueDate(v time.Time)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *InvoiceRequest) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### SetDueDateNil

`func (o *InvoiceRequest) SetDueDateNil(b bool)`

 SetDueDateNil sets the value for DueDate to be an explicit nil

### UnsetDueDate
`func (o *InvoiceRequest) UnsetDueDate()`

UnsetDueDate ensures that no value is present for DueDate, not even an explicit nil
### GetPaidOnDate

`func (o *InvoiceRequest) GetPaidOnDate() time.Time`

GetPaidOnDate returns the PaidOnDate field if non-nil, zero value otherwise.

### GetPaidOnDateOk

`func (o *InvoiceRequest) GetPaidOnDateOk() (*time.Time, bool)`

GetPaidOnDateOk returns a tuple with the PaidOnDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidOnDate

`func (o *InvoiceRequest) SetPaidOnDate(v time.Time)`

SetPaidOnDate sets PaidOnDate field to given value.

### HasPaidOnDate

`func (o *InvoiceRequest) HasPaidOnDate() bool`

HasPaidOnDate returns a boolean if a field has been set.

### SetPaidOnDateNil

`func (o *InvoiceRequest) SetPaidOnDateNil(b bool)`

 SetPaidOnDateNil sets the value for PaidOnDate to be an explicit nil

### UnsetPaidOnDate
`func (o *InvoiceRequest) UnsetPaidOnDate()`

UnsetPaidOnDate ensures that no value is present for PaidOnDate, not even an explicit nil
### GetMemo

`func (o *InvoiceRequest) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *InvoiceRequest) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *InvoiceRequest) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *InvoiceRequest) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### SetMemoNil

`func (o *InvoiceRequest) SetMemoNil(b bool)`

 SetMemoNil sets the value for Memo to be an explicit nil

### UnsetMemo
`func (o *InvoiceRequest) UnsetMemo()`

UnsetMemo ensures that no value is present for Memo, not even an explicit nil
### GetCompany

`func (o *InvoiceRequest) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *InvoiceRequest) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *InvoiceRequest) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *InvoiceRequest) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### SetCompanyNil

`func (o *InvoiceRequest) SetCompanyNil(b bool)`

 SetCompanyNil sets the value for Company to be an explicit nil

### UnsetCompany
`func (o *InvoiceRequest) UnsetCompany()`

UnsetCompany ensures that no value is present for Company, not even an explicit nil
### GetCurrency

`func (o *InvoiceRequest) GetCurrency() CurrencyEnum`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *InvoiceRequest) GetCurrencyOk() (*CurrencyEnum, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *InvoiceRequest) SetCurrency(v CurrencyEnum)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *InvoiceRequest) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### SetCurrencyNil

`func (o *InvoiceRequest) SetCurrencyNil(b bool)`

 SetCurrencyNil sets the value for Currency to be an explicit nil

### UnsetCurrency
`func (o *InvoiceRequest) UnsetCurrency()`

UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
### GetExchangeRate

`func (o *InvoiceRequest) GetExchangeRate() float64`

GetExchangeRate returns the ExchangeRate field if non-nil, zero value otherwise.

### GetExchangeRateOk

`func (o *InvoiceRequest) GetExchangeRateOk() (*float64, bool)`

GetExchangeRateOk returns a tuple with the ExchangeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeRate

`func (o *InvoiceRequest) SetExchangeRate(v float64)`

SetExchangeRate sets ExchangeRate field to given value.

### HasExchangeRate

`func (o *InvoiceRequest) HasExchangeRate() bool`

HasExchangeRate returns a boolean if a field has been set.

### SetExchangeRateNil

`func (o *InvoiceRequest) SetExchangeRateNil(b bool)`

 SetExchangeRateNil sets the value for ExchangeRate to be an explicit nil

### UnsetExchangeRate
`func (o *InvoiceRequest) UnsetExchangeRate()`

UnsetExchangeRate ensures that no value is present for ExchangeRate, not even an explicit nil
### GetTotalDiscount

`func (o *InvoiceRequest) GetTotalDiscount() float32`

GetTotalDiscount returns the TotalDiscount field if non-nil, zero value otherwise.

### GetTotalDiscountOk

`func (o *InvoiceRequest) GetTotalDiscountOk() (*float32, bool)`

GetTotalDiscountOk returns a tuple with the TotalDiscount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDiscount

`func (o *InvoiceRequest) SetTotalDiscount(v float32)`

SetTotalDiscount sets TotalDiscount field to given value.

### HasTotalDiscount

`func (o *InvoiceRequest) HasTotalDiscount() bool`

HasTotalDiscount returns a boolean if a field has been set.

### SetTotalDiscountNil

`func (o *InvoiceRequest) SetTotalDiscountNil(b bool)`

 SetTotalDiscountNil sets the value for TotalDiscount to be an explicit nil

### UnsetTotalDiscount
`func (o *InvoiceRequest) UnsetTotalDiscount()`

UnsetTotalDiscount ensures that no value is present for TotalDiscount, not even an explicit nil
### GetSubTotal

`func (o *InvoiceRequest) GetSubTotal() float32`

GetSubTotal returns the SubTotal field if non-nil, zero value otherwise.

### GetSubTotalOk

`func (o *InvoiceRequest) GetSubTotalOk() (*float32, bool)`

GetSubTotalOk returns a tuple with the SubTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubTotal

`func (o *InvoiceRequest) SetSubTotal(v float32)`

SetSubTotal sets SubTotal field to given value.

### HasSubTotal

`func (o *InvoiceRequest) HasSubTotal() bool`

HasSubTotal returns a boolean if a field has been set.

### SetSubTotalNil

`func (o *InvoiceRequest) SetSubTotalNil(b bool)`

 SetSubTotalNil sets the value for SubTotal to be an explicit nil

### UnsetSubTotal
`func (o *InvoiceRequest) UnsetSubTotal()`

UnsetSubTotal ensures that no value is present for SubTotal, not even an explicit nil
### GetTotalTaxAmount

`func (o *InvoiceRequest) GetTotalTaxAmount() float32`

GetTotalTaxAmount returns the TotalTaxAmount field if non-nil, zero value otherwise.

### GetTotalTaxAmountOk

`func (o *InvoiceRequest) GetTotalTaxAmountOk() (*float32, bool)`

GetTotalTaxAmountOk returns a tuple with the TotalTaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTaxAmount

`func (o *InvoiceRequest) SetTotalTaxAmount(v float32)`

SetTotalTaxAmount sets TotalTaxAmount field to given value.

### HasTotalTaxAmount

`func (o *InvoiceRequest) HasTotalTaxAmount() bool`

HasTotalTaxAmount returns a boolean if a field has been set.

### SetTotalTaxAmountNil

`func (o *InvoiceRequest) SetTotalTaxAmountNil(b bool)`

 SetTotalTaxAmountNil sets the value for TotalTaxAmount to be an explicit nil

### UnsetTotalTaxAmount
`func (o *InvoiceRequest) UnsetTotalTaxAmount()`

UnsetTotalTaxAmount ensures that no value is present for TotalTaxAmount, not even an explicit nil
### GetTotalAmount

`func (o *InvoiceRequest) GetTotalAmount() float32`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *InvoiceRequest) GetTotalAmountOk() (*float32, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *InvoiceRequest) SetTotalAmount(v float32)`

SetTotalAmount sets TotalAmount field to given value.

### HasTotalAmount

`func (o *InvoiceRequest) HasTotalAmount() bool`

HasTotalAmount returns a boolean if a field has been set.

### SetTotalAmountNil

`func (o *InvoiceRequest) SetTotalAmountNil(b bool)`

 SetTotalAmountNil sets the value for TotalAmount to be an explicit nil

### UnsetTotalAmount
`func (o *InvoiceRequest) UnsetTotalAmount()`

UnsetTotalAmount ensures that no value is present for TotalAmount, not even an explicit nil
### GetBalance

`func (o *InvoiceRequest) GetBalance() float32`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *InvoiceRequest) GetBalanceOk() (*float32, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *InvoiceRequest) SetBalance(v float32)`

SetBalance sets Balance field to given value.

### HasBalance

`func (o *InvoiceRequest) HasBalance() bool`

HasBalance returns a boolean if a field has been set.

### SetBalanceNil

`func (o *InvoiceRequest) SetBalanceNil(b bool)`

 SetBalanceNil sets the value for Balance to be an explicit nil

### UnsetBalance
`func (o *InvoiceRequest) UnsetBalance()`

UnsetBalance ensures that no value is present for Balance, not even an explicit nil
### GetRemoteUpdatedAt

`func (o *InvoiceRequest) GetRemoteUpdatedAt() time.Time`

GetRemoteUpdatedAt returns the RemoteUpdatedAt field if non-nil, zero value otherwise.

### GetRemoteUpdatedAtOk

`func (o *InvoiceRequest) GetRemoteUpdatedAtOk() (*time.Time, bool)`

GetRemoteUpdatedAtOk returns a tuple with the RemoteUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteUpdatedAt

`func (o *InvoiceRequest) SetRemoteUpdatedAt(v time.Time)`

SetRemoteUpdatedAt sets RemoteUpdatedAt field to given value.

### HasRemoteUpdatedAt

`func (o *InvoiceRequest) HasRemoteUpdatedAt() bool`

HasRemoteUpdatedAt returns a boolean if a field has been set.

### SetRemoteUpdatedAtNil

`func (o *InvoiceRequest) SetRemoteUpdatedAtNil(b bool)`

 SetRemoteUpdatedAtNil sets the value for RemoteUpdatedAt to be an explicit nil

### UnsetRemoteUpdatedAt
`func (o *InvoiceRequest) UnsetRemoteUpdatedAt()`

UnsetRemoteUpdatedAt ensures that no value is present for RemoteUpdatedAt, not even an explicit nil
### GetPayments

`func (o *InvoiceRequest) GetPayments() []string`

GetPayments returns the Payments field if non-nil, zero value otherwise.

### GetPaymentsOk

`func (o *InvoiceRequest) GetPaymentsOk() (*[]string, bool)`

GetPaymentsOk returns a tuple with the Payments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayments

`func (o *InvoiceRequest) SetPayments(v []string)`

SetPayments sets Payments field to given value.

### HasPayments

`func (o *InvoiceRequest) HasPayments() bool`

HasPayments returns a boolean if a field has been set.

### GetLineItems

`func (o *InvoiceRequest) GetLineItems() []InvoiceLineItemRequest`

GetLineItems returns the LineItems field if non-nil, zero value otherwise.

### GetLineItemsOk

`func (o *InvoiceRequest) GetLineItemsOk() (*[]InvoiceLineItemRequest, bool)`

GetLineItemsOk returns a tuple with the LineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItems

`func (o *InvoiceRequest) SetLineItems(v []InvoiceLineItemRequest)`

SetLineItems sets LineItems field to given value.

### HasLineItems

`func (o *InvoiceRequest) HasLineItems() bool`

HasLineItems returns a boolean if a field has been set.

### GetIntegrationParams

`func (o *InvoiceRequest) GetIntegrationParams() map[string]interface{}`

GetIntegrationParams returns the IntegrationParams field if non-nil, zero value otherwise.

### GetIntegrationParamsOk

`func (o *InvoiceRequest) GetIntegrationParamsOk() (*map[string]interface{}, bool)`

GetIntegrationParamsOk returns a tuple with the IntegrationParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationParams

`func (o *InvoiceRequest) SetIntegrationParams(v map[string]interface{})`

SetIntegrationParams sets IntegrationParams field to given value.

### HasIntegrationParams

`func (o *InvoiceRequest) HasIntegrationParams() bool`

HasIntegrationParams returns a boolean if a field has been set.

### SetIntegrationParamsNil

`func (o *InvoiceRequest) SetIntegrationParamsNil(b bool)`

 SetIntegrationParamsNil sets the value for IntegrationParams to be an explicit nil

### UnsetIntegrationParams
`func (o *InvoiceRequest) UnsetIntegrationParams()`

UnsetIntegrationParams ensures that no value is present for IntegrationParams, not even an explicit nil
### GetLinkedAccountParams

`func (o *InvoiceRequest) GetLinkedAccountParams() map[string]interface{}`

GetLinkedAccountParams returns the LinkedAccountParams field if non-nil, zero value otherwise.

### GetLinkedAccountParamsOk

`func (o *InvoiceRequest) GetLinkedAccountParamsOk() (*map[string]interface{}, bool)`

GetLinkedAccountParamsOk returns a tuple with the LinkedAccountParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedAccountParams

`func (o *InvoiceRequest) SetLinkedAccountParams(v map[string]interface{})`

SetLinkedAccountParams sets LinkedAccountParams field to given value.

### HasLinkedAccountParams

`func (o *InvoiceRequest) HasLinkedAccountParams() bool`

HasLinkedAccountParams returns a boolean if a field has been set.

### SetLinkedAccountParamsNil

`func (o *InvoiceRequest) SetLinkedAccountParamsNil(b bool)`

 SetLinkedAccountParamsNil sets the value for LinkedAccountParams to be an explicit nil

### UnsetLinkedAccountParams
`func (o *InvoiceRequest) UnsetLinkedAccountParams()`

UnsetLinkedAccountParams ensures that no value is present for LinkedAccountParams, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


