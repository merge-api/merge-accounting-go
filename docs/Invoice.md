# Invoice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**RemoteId** | Pointer to **NullableString** | The third-party API ID of the matching object. | [optional] 
**RemoteData** | Pointer to [**[]RemoteData**](RemoteData.md) |  | [optional] [readonly] 
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
**LineItems** | Pointer to [**[]InvoiceLineItem**](InvoiceLineItem.md) |  | [optional] 
**RemoteWasDeleted** | Pointer to **bool** |  | [optional] [readonly] 
**FieldMappings** | Pointer to **map[string]interface{}** |  | [optional] [readonly] 

## Methods

### NewInvoice

`func NewInvoice() *Invoice`

NewInvoice instantiates a new Invoice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceWithDefaults

`func NewInvoiceWithDefaults() *Invoice`

NewInvoiceWithDefaults instantiates a new Invoice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Invoice) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Invoice) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Invoice) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Invoice) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRemoteId

`func (o *Invoice) GetRemoteId() string`

GetRemoteId returns the RemoteId field if non-nil, zero value otherwise.

### GetRemoteIdOk

`func (o *Invoice) GetRemoteIdOk() (*string, bool)`

GetRemoteIdOk returns a tuple with the RemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteId

`func (o *Invoice) SetRemoteId(v string)`

SetRemoteId sets RemoteId field to given value.

### HasRemoteId

`func (o *Invoice) HasRemoteId() bool`

HasRemoteId returns a boolean if a field has been set.

### SetRemoteIdNil

`func (o *Invoice) SetRemoteIdNil(b bool)`

 SetRemoteIdNil sets the value for RemoteId to be an explicit nil

### UnsetRemoteId
`func (o *Invoice) UnsetRemoteId()`

UnsetRemoteId ensures that no value is present for RemoteId, not even an explicit nil
### GetRemoteData

`func (o *Invoice) GetRemoteData() []RemoteData`

GetRemoteData returns the RemoteData field if non-nil, zero value otherwise.

### GetRemoteDataOk

`func (o *Invoice) GetRemoteDataOk() (*[]RemoteData, bool)`

GetRemoteDataOk returns a tuple with the RemoteData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteData

`func (o *Invoice) SetRemoteData(v []RemoteData)`

SetRemoteData sets RemoteData field to given value.

### HasRemoteData

`func (o *Invoice) HasRemoteData() bool`

HasRemoteData returns a boolean if a field has been set.

### SetRemoteDataNil

`func (o *Invoice) SetRemoteDataNil(b bool)`

 SetRemoteDataNil sets the value for RemoteData to be an explicit nil

### UnsetRemoteData
`func (o *Invoice) UnsetRemoteData()`

UnsetRemoteData ensures that no value is present for RemoteData, not even an explicit nil
### GetType

`func (o *Invoice) GetType() InvoiceTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Invoice) GetTypeOk() (*InvoiceTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Invoice) SetType(v InvoiceTypeEnum)`

SetType sets Type field to given value.

### HasType

`func (o *Invoice) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *Invoice) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *Invoice) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetContact

`func (o *Invoice) GetContact() string`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *Invoice) GetContactOk() (*string, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *Invoice) SetContact(v string)`

SetContact sets Contact field to given value.

### HasContact

`func (o *Invoice) HasContact() bool`

HasContact returns a boolean if a field has been set.

### SetContactNil

`func (o *Invoice) SetContactNil(b bool)`

 SetContactNil sets the value for Contact to be an explicit nil

### UnsetContact
`func (o *Invoice) UnsetContact()`

UnsetContact ensures that no value is present for Contact, not even an explicit nil
### GetNumber

`func (o *Invoice) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *Invoice) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *Invoice) SetNumber(v string)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *Invoice) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### SetNumberNil

`func (o *Invoice) SetNumberNil(b bool)`

 SetNumberNil sets the value for Number to be an explicit nil

### UnsetNumber
`func (o *Invoice) UnsetNumber()`

UnsetNumber ensures that no value is present for Number, not even an explicit nil
### GetIssueDate

`func (o *Invoice) GetIssueDate() time.Time`

GetIssueDate returns the IssueDate field if non-nil, zero value otherwise.

### GetIssueDateOk

`func (o *Invoice) GetIssueDateOk() (*time.Time, bool)`

GetIssueDateOk returns a tuple with the IssueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueDate

`func (o *Invoice) SetIssueDate(v time.Time)`

SetIssueDate sets IssueDate field to given value.

### HasIssueDate

`func (o *Invoice) HasIssueDate() bool`

HasIssueDate returns a boolean if a field has been set.

### SetIssueDateNil

`func (o *Invoice) SetIssueDateNil(b bool)`

 SetIssueDateNil sets the value for IssueDate to be an explicit nil

### UnsetIssueDate
`func (o *Invoice) UnsetIssueDate()`

UnsetIssueDate ensures that no value is present for IssueDate, not even an explicit nil
### GetDueDate

`func (o *Invoice) GetDueDate() time.Time`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *Invoice) GetDueDateOk() (*time.Time, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *Invoice) SetDueDate(v time.Time)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *Invoice) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### SetDueDateNil

`func (o *Invoice) SetDueDateNil(b bool)`

 SetDueDateNil sets the value for DueDate to be an explicit nil

### UnsetDueDate
`func (o *Invoice) UnsetDueDate()`

UnsetDueDate ensures that no value is present for DueDate, not even an explicit nil
### GetPaidOnDate

`func (o *Invoice) GetPaidOnDate() time.Time`

GetPaidOnDate returns the PaidOnDate field if non-nil, zero value otherwise.

### GetPaidOnDateOk

`func (o *Invoice) GetPaidOnDateOk() (*time.Time, bool)`

GetPaidOnDateOk returns a tuple with the PaidOnDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidOnDate

`func (o *Invoice) SetPaidOnDate(v time.Time)`

SetPaidOnDate sets PaidOnDate field to given value.

### HasPaidOnDate

`func (o *Invoice) HasPaidOnDate() bool`

HasPaidOnDate returns a boolean if a field has been set.

### SetPaidOnDateNil

`func (o *Invoice) SetPaidOnDateNil(b bool)`

 SetPaidOnDateNil sets the value for PaidOnDate to be an explicit nil

### UnsetPaidOnDate
`func (o *Invoice) UnsetPaidOnDate()`

UnsetPaidOnDate ensures that no value is present for PaidOnDate, not even an explicit nil
### GetMemo

`func (o *Invoice) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *Invoice) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *Invoice) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *Invoice) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### SetMemoNil

`func (o *Invoice) SetMemoNil(b bool)`

 SetMemoNil sets the value for Memo to be an explicit nil

### UnsetMemo
`func (o *Invoice) UnsetMemo()`

UnsetMemo ensures that no value is present for Memo, not even an explicit nil
### GetCompany

`func (o *Invoice) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *Invoice) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *Invoice) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *Invoice) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### SetCompanyNil

`func (o *Invoice) SetCompanyNil(b bool)`

 SetCompanyNil sets the value for Company to be an explicit nil

### UnsetCompany
`func (o *Invoice) UnsetCompany()`

UnsetCompany ensures that no value is present for Company, not even an explicit nil
### GetCurrency

`func (o *Invoice) GetCurrency() CurrencyEnum`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Invoice) GetCurrencyOk() (*CurrencyEnum, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Invoice) SetCurrency(v CurrencyEnum)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *Invoice) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### SetCurrencyNil

`func (o *Invoice) SetCurrencyNil(b bool)`

 SetCurrencyNil sets the value for Currency to be an explicit nil

### UnsetCurrency
`func (o *Invoice) UnsetCurrency()`

UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
### GetExchangeRate

`func (o *Invoice) GetExchangeRate() float64`

GetExchangeRate returns the ExchangeRate field if non-nil, zero value otherwise.

### GetExchangeRateOk

`func (o *Invoice) GetExchangeRateOk() (*float64, bool)`

GetExchangeRateOk returns a tuple with the ExchangeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeRate

`func (o *Invoice) SetExchangeRate(v float64)`

SetExchangeRate sets ExchangeRate field to given value.

### HasExchangeRate

`func (o *Invoice) HasExchangeRate() bool`

HasExchangeRate returns a boolean if a field has been set.

### SetExchangeRateNil

`func (o *Invoice) SetExchangeRateNil(b bool)`

 SetExchangeRateNil sets the value for ExchangeRate to be an explicit nil

### UnsetExchangeRate
`func (o *Invoice) UnsetExchangeRate()`

UnsetExchangeRate ensures that no value is present for ExchangeRate, not even an explicit nil
### GetTotalDiscount

`func (o *Invoice) GetTotalDiscount() float32`

GetTotalDiscount returns the TotalDiscount field if non-nil, zero value otherwise.

### GetTotalDiscountOk

`func (o *Invoice) GetTotalDiscountOk() (*float32, bool)`

GetTotalDiscountOk returns a tuple with the TotalDiscount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDiscount

`func (o *Invoice) SetTotalDiscount(v float32)`

SetTotalDiscount sets TotalDiscount field to given value.

### HasTotalDiscount

`func (o *Invoice) HasTotalDiscount() bool`

HasTotalDiscount returns a boolean if a field has been set.

### SetTotalDiscountNil

`func (o *Invoice) SetTotalDiscountNil(b bool)`

 SetTotalDiscountNil sets the value for TotalDiscount to be an explicit nil

### UnsetTotalDiscount
`func (o *Invoice) UnsetTotalDiscount()`

UnsetTotalDiscount ensures that no value is present for TotalDiscount, not even an explicit nil
### GetSubTotal

`func (o *Invoice) GetSubTotal() float32`

GetSubTotal returns the SubTotal field if non-nil, zero value otherwise.

### GetSubTotalOk

`func (o *Invoice) GetSubTotalOk() (*float32, bool)`

GetSubTotalOk returns a tuple with the SubTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubTotal

`func (o *Invoice) SetSubTotal(v float32)`

SetSubTotal sets SubTotal field to given value.

### HasSubTotal

`func (o *Invoice) HasSubTotal() bool`

HasSubTotal returns a boolean if a field has been set.

### SetSubTotalNil

`func (o *Invoice) SetSubTotalNil(b bool)`

 SetSubTotalNil sets the value for SubTotal to be an explicit nil

### UnsetSubTotal
`func (o *Invoice) UnsetSubTotal()`

UnsetSubTotal ensures that no value is present for SubTotal, not even an explicit nil
### GetTotalTaxAmount

`func (o *Invoice) GetTotalTaxAmount() float32`

GetTotalTaxAmount returns the TotalTaxAmount field if non-nil, zero value otherwise.

### GetTotalTaxAmountOk

`func (o *Invoice) GetTotalTaxAmountOk() (*float32, bool)`

GetTotalTaxAmountOk returns a tuple with the TotalTaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTaxAmount

`func (o *Invoice) SetTotalTaxAmount(v float32)`

SetTotalTaxAmount sets TotalTaxAmount field to given value.

### HasTotalTaxAmount

`func (o *Invoice) HasTotalTaxAmount() bool`

HasTotalTaxAmount returns a boolean if a field has been set.

### SetTotalTaxAmountNil

`func (o *Invoice) SetTotalTaxAmountNil(b bool)`

 SetTotalTaxAmountNil sets the value for TotalTaxAmount to be an explicit nil

### UnsetTotalTaxAmount
`func (o *Invoice) UnsetTotalTaxAmount()`

UnsetTotalTaxAmount ensures that no value is present for TotalTaxAmount, not even an explicit nil
### GetTotalAmount

`func (o *Invoice) GetTotalAmount() float32`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *Invoice) GetTotalAmountOk() (*float32, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *Invoice) SetTotalAmount(v float32)`

SetTotalAmount sets TotalAmount field to given value.

### HasTotalAmount

`func (o *Invoice) HasTotalAmount() bool`

HasTotalAmount returns a boolean if a field has been set.

### SetTotalAmountNil

`func (o *Invoice) SetTotalAmountNil(b bool)`

 SetTotalAmountNil sets the value for TotalAmount to be an explicit nil

### UnsetTotalAmount
`func (o *Invoice) UnsetTotalAmount()`

UnsetTotalAmount ensures that no value is present for TotalAmount, not even an explicit nil
### GetBalance

`func (o *Invoice) GetBalance() float32`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *Invoice) GetBalanceOk() (*float32, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *Invoice) SetBalance(v float32)`

SetBalance sets Balance field to given value.

### HasBalance

`func (o *Invoice) HasBalance() bool`

HasBalance returns a boolean if a field has been set.

### SetBalanceNil

`func (o *Invoice) SetBalanceNil(b bool)`

 SetBalanceNil sets the value for Balance to be an explicit nil

### UnsetBalance
`func (o *Invoice) UnsetBalance()`

UnsetBalance ensures that no value is present for Balance, not even an explicit nil
### GetRemoteUpdatedAt

`func (o *Invoice) GetRemoteUpdatedAt() time.Time`

GetRemoteUpdatedAt returns the RemoteUpdatedAt field if non-nil, zero value otherwise.

### GetRemoteUpdatedAtOk

`func (o *Invoice) GetRemoteUpdatedAtOk() (*time.Time, bool)`

GetRemoteUpdatedAtOk returns a tuple with the RemoteUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteUpdatedAt

`func (o *Invoice) SetRemoteUpdatedAt(v time.Time)`

SetRemoteUpdatedAt sets RemoteUpdatedAt field to given value.

### HasRemoteUpdatedAt

`func (o *Invoice) HasRemoteUpdatedAt() bool`

HasRemoteUpdatedAt returns a boolean if a field has been set.

### SetRemoteUpdatedAtNil

`func (o *Invoice) SetRemoteUpdatedAtNil(b bool)`

 SetRemoteUpdatedAtNil sets the value for RemoteUpdatedAt to be an explicit nil

### UnsetRemoteUpdatedAt
`func (o *Invoice) UnsetRemoteUpdatedAt()`

UnsetRemoteUpdatedAt ensures that no value is present for RemoteUpdatedAt, not even an explicit nil
### GetPayments

`func (o *Invoice) GetPayments() []string`

GetPayments returns the Payments field if non-nil, zero value otherwise.

### GetPaymentsOk

`func (o *Invoice) GetPaymentsOk() (*[]string, bool)`

GetPaymentsOk returns a tuple with the Payments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayments

`func (o *Invoice) SetPayments(v []string)`

SetPayments sets Payments field to given value.

### HasPayments

`func (o *Invoice) HasPayments() bool`

HasPayments returns a boolean if a field has been set.

### GetLineItems

`func (o *Invoice) GetLineItems() []InvoiceLineItem`

GetLineItems returns the LineItems field if non-nil, zero value otherwise.

### GetLineItemsOk

`func (o *Invoice) GetLineItemsOk() (*[]InvoiceLineItem, bool)`

GetLineItemsOk returns a tuple with the LineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItems

`func (o *Invoice) SetLineItems(v []InvoiceLineItem)`

SetLineItems sets LineItems field to given value.

### HasLineItems

`func (o *Invoice) HasLineItems() bool`

HasLineItems returns a boolean if a field has been set.

### GetRemoteWasDeleted

`func (o *Invoice) GetRemoteWasDeleted() bool`

GetRemoteWasDeleted returns the RemoteWasDeleted field if non-nil, zero value otherwise.

### GetRemoteWasDeletedOk

`func (o *Invoice) GetRemoteWasDeletedOk() (*bool, bool)`

GetRemoteWasDeletedOk returns a tuple with the RemoteWasDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteWasDeleted

`func (o *Invoice) SetRemoteWasDeleted(v bool)`

SetRemoteWasDeleted sets RemoteWasDeleted field to given value.

### HasRemoteWasDeleted

`func (o *Invoice) HasRemoteWasDeleted() bool`

HasRemoteWasDeleted returns a boolean if a field has been set.

### GetFieldMappings

`func (o *Invoice) GetFieldMappings() map[string]interface{}`

GetFieldMappings returns the FieldMappings field if non-nil, zero value otherwise.

### GetFieldMappingsOk

`func (o *Invoice) GetFieldMappingsOk() (*map[string]interface{}, bool)`

GetFieldMappingsOk returns a tuple with the FieldMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldMappings

`func (o *Invoice) SetFieldMappings(v map[string]interface{})`

SetFieldMappings sets FieldMappings field to given value.

### HasFieldMappings

`func (o *Invoice) HasFieldMappings() bool`

HasFieldMappings returns a boolean if a field has been set.

### SetFieldMappingsNil

`func (o *Invoice) SetFieldMappingsNil(b bool)`

 SetFieldMappingsNil sets the value for FieldMappings to be an explicit nil

### UnsetFieldMappings
`func (o *Invoice) UnsetFieldMappings()`

UnsetFieldMappings ensures that no value is present for FieldMappings, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


