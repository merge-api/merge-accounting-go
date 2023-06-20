# Invoice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Type** | Pointer to [**NullableInvoiceTypeEnum**](InvoiceTypeEnum.md) | Whether the invoice is an accounts receivable or accounts payable. If &#x60;type&#x60; is &#x60;accounts_payable&#x60;, the invoice is a bill. If &#x60;type&#x60; is &#x60;accounts_receivable&#x60;, it is an invoice.  * &#x60;ACCOUNTS_RECEIVABLE&#x60; - ACCOUNTS_RECEIVABLE * &#x60;ACCOUNTS_PAYABLE&#x60; - ACCOUNTS_PAYABLE | [optional] 
**Contact** | Pointer to **NullableString** | The invoice&#39;s contact. | [optional] 
**Number** | Pointer to **NullableString** | The invoice&#39;s number. | [optional] 
**IssueDate** | Pointer to **NullableTime** | The invoice&#39;s issue date. | [optional] 
**DueDate** | Pointer to **NullableTime** | The invoice&#39;s due date. | [optional] 
**PaidOnDate** | Pointer to **NullableTime** | The invoice&#39;s paid date. | [optional] 
**Memo** | Pointer to **NullableString** | The invoice&#39;s private note. | [optional] 
**Company** | Pointer to **NullableString** | The company the invoice belongs to. | [optional] 
**Currency** | Pointer to [**NullableCurrencyEnum**](CurrencyEnum.md) | The invoice&#39;s currency.  * &#x60;XUA&#x60; - ADB Unit of Account * &#x60;AFN&#x60; - Afghan Afghani * &#x60;AFA&#x60; - Afghan Afghani (1927–2002) * &#x60;ALL&#x60; - Albanian Lek * &#x60;ALK&#x60; - Albanian Lek (1946–1965) * &#x60;DZD&#x60; - Algerian Dinar * &#x60;ADP&#x60; - Andorran Peseta * &#x60;AOA&#x60; - Angolan Kwanza * &#x60;AOK&#x60; - Angolan Kwanza (1977–1991) * &#x60;AON&#x60; - Angolan New Kwanza (1990–2000) * &#x60;AOR&#x60; - Angolan Readjusted Kwanza (1995–1999) * &#x60;ARA&#x60; - Argentine Austral * &#x60;ARS&#x60; - Argentine Peso * &#x60;ARM&#x60; - Argentine Peso (1881–1970) * &#x60;ARP&#x60; - Argentine Peso (1983–1985) * &#x60;ARL&#x60; - Argentine Peso Ley (1970–1983) * &#x60;AMD&#x60; - Armenian Dram * &#x60;AWG&#x60; - Aruban Florin * &#x60;AUD&#x60; - Australian Dollar * &#x60;ATS&#x60; - Austrian Schilling * &#x60;AZN&#x60; - Azerbaijani Manat * &#x60;AZM&#x60; - Azerbaijani Manat (1993–2006) * &#x60;BSD&#x60; - Bahamian Dollar * &#x60;BHD&#x60; - Bahraini Dinar * &#x60;BDT&#x60; - Bangladeshi Taka * &#x60;BBD&#x60; - Barbadian Dollar * &#x60;BYN&#x60; - Belarusian Ruble * &#x60;BYB&#x60; - Belarusian Ruble (1994–1999) * &#x60;BYR&#x60; - Belarusian Ruble (2000–2016) * &#x60;BEF&#x60; - Belgian Franc * &#x60;BEC&#x60; - Belgian Franc (convertible) * &#x60;BEL&#x60; - Belgian Franc (financial) * &#x60;BZD&#x60; - Belize Dollar * &#x60;BMD&#x60; - Bermudan Dollar * &#x60;BTN&#x60; - Bhutanese Ngultrum * &#x60;BOB&#x60; - Bolivian Boliviano * &#x60;BOL&#x60; - Bolivian Boliviano (1863–1963) * &#x60;BOV&#x60; - Bolivian Mvdol * &#x60;BOP&#x60; - Bolivian Peso * &#x60;BAM&#x60; - Bosnia-Herzegovina Convertible Mark * &#x60;BAD&#x60; - Bosnia-Herzegovina Dinar (1992–1994) * &#x60;BAN&#x60; - Bosnia-Herzegovina New Dinar (1994–1997) * &#x60;BWP&#x60; - Botswanan Pula * &#x60;BRC&#x60; - Brazilian Cruzado (1986–1989) * &#x60;BRZ&#x60; - Brazilian Cruzeiro (1942–1967) * &#x60;BRE&#x60; - Brazilian Cruzeiro (1990–1993) * &#x60;BRR&#x60; - Brazilian Cruzeiro (1993–1994) * &#x60;BRN&#x60; - Brazilian New Cruzado (1989–1990) * &#x60;BRB&#x60; - Brazilian New Cruzeiro (1967–1986) * &#x60;BRL&#x60; - Brazilian Real * &#x60;GBP&#x60; - British Pound * &#x60;BND&#x60; - Brunei Dollar * &#x60;BGL&#x60; - Bulgarian Hard Lev * &#x60;BGN&#x60; - Bulgarian Lev * &#x60;BGO&#x60; - Bulgarian Lev (1879–1952) * &#x60;BGM&#x60; - Bulgarian Socialist Lev * &#x60;BUK&#x60; - Burmese Kyat * &#x60;BIF&#x60; - Burundian Franc * &#x60;XPF&#x60; - CFP Franc * &#x60;KHR&#x60; - Cambodian Riel * &#x60;CAD&#x60; - Canadian Dollar * &#x60;CVE&#x60; - Cape Verdean Escudo * &#x60;KYD&#x60; - Cayman Islands Dollar * &#x60;XAF&#x60; - Central African CFA Franc * &#x60;CLE&#x60; - Chilean Escudo * &#x60;CLP&#x60; - Chilean Peso * &#x60;CLF&#x60; - Chilean Unit of Account (UF) * &#x60;CNX&#x60; - Chinese People’s Bank Dollar * &#x60;CNY&#x60; - Chinese Yuan * &#x60;CNH&#x60; - Chinese Yuan (offshore) * &#x60;COP&#x60; - Colombian Peso * &#x60;COU&#x60; - Colombian Real Value Unit * &#x60;KMF&#x60; - Comorian Franc * &#x60;CDF&#x60; - Congolese Franc * &#x60;CRC&#x60; - Costa Rican Colón * &#x60;HRD&#x60; - Croatian Dinar * &#x60;HRK&#x60; - Croatian Kuna * &#x60;CUC&#x60; - Cuban Convertible Peso * &#x60;CUP&#x60; - Cuban Peso * &#x60;CYP&#x60; - Cypriot Pound * &#x60;CZK&#x60; - Czech Koruna * &#x60;CSK&#x60; - Czechoslovak Hard Koruna * &#x60;DKK&#x60; - Danish Krone * &#x60;DJF&#x60; - Djiboutian Franc * &#x60;DOP&#x60; - Dominican Peso * &#x60;NLG&#x60; - Dutch Guilder * &#x60;XCD&#x60; - East Caribbean Dollar * &#x60;DDM&#x60; - East German Mark * &#x60;ECS&#x60; - Ecuadorian Sucre * &#x60;ECV&#x60; - Ecuadorian Unit of Constant Value * &#x60;EGP&#x60; - Egyptian Pound * &#x60;GQE&#x60; - Equatorial Guinean Ekwele * &#x60;ERN&#x60; - Eritrean Nakfa * &#x60;EEK&#x60; - Estonian Kroon * &#x60;ETB&#x60; - Ethiopian Birr * &#x60;EUR&#x60; - Euro * &#x60;XBA&#x60; - European Composite Unit * &#x60;XEU&#x60; - European Currency Unit * &#x60;XBB&#x60; - European Monetary Unit * &#x60;XBC&#x60; - European Unit of Account (XBC) * &#x60;XBD&#x60; - European Unit of Account (XBD) * &#x60;FKP&#x60; - Falkland Islands Pound * &#x60;FJD&#x60; - Fijian Dollar * &#x60;FIM&#x60; - Finnish Markka * &#x60;FRF&#x60; - French Franc * &#x60;XFO&#x60; - French Gold Franc * &#x60;XFU&#x60; - French UIC-Franc * &#x60;GMD&#x60; - Gambian Dalasi * &#x60;GEK&#x60; - Georgian Kupon Larit * &#x60;GEL&#x60; - Georgian Lari * &#x60;DEM&#x60; - German Mark * &#x60;GHS&#x60; - Ghanaian Cedi * &#x60;GHC&#x60; - Ghanaian Cedi (1979–2007) * &#x60;GIP&#x60; - Gibraltar Pound * &#x60;XAU&#x60; - Gold * &#x60;GRD&#x60; - Greek Drachma * &#x60;GTQ&#x60; - Guatemalan Quetzal * &#x60;GWP&#x60; - Guinea-Bissau Peso * &#x60;GNF&#x60; - Guinean Franc * &#x60;GNS&#x60; - Guinean Syli * &#x60;GYD&#x60; - Guyanaese Dollar * &#x60;HTG&#x60; - Haitian Gourde * &#x60;HNL&#x60; - Honduran Lempira * &#x60;HKD&#x60; - Hong Kong Dollar * &#x60;HUF&#x60; - Hungarian Forint * &#x60;IMP&#x60; - IMP * &#x60;ISK&#x60; - Icelandic Króna * &#x60;ISJ&#x60; - Icelandic Króna (1918–1981) * &#x60;INR&#x60; - Indian Rupee * &#x60;IDR&#x60; - Indonesian Rupiah * &#x60;IRR&#x60; - Iranian Rial * &#x60;IQD&#x60; - Iraqi Dinar * &#x60;IEP&#x60; - Irish Pound * &#x60;ILS&#x60; - Israeli New Shekel * &#x60;ILP&#x60; - Israeli Pound * &#x60;ILR&#x60; - Israeli Shekel (1980–1985) * &#x60;ITL&#x60; - Italian Lira * &#x60;JMD&#x60; - Jamaican Dollar * &#x60;JPY&#x60; - Japanese Yen * &#x60;JOD&#x60; - Jordanian Dinar * &#x60;KZT&#x60; - Kazakhstani Tenge * &#x60;KES&#x60; - Kenyan Shilling * &#x60;KWD&#x60; - Kuwaiti Dinar * &#x60;KGS&#x60; - Kyrgystani Som * &#x60;LAK&#x60; - Laotian Kip * &#x60;LVL&#x60; - Latvian Lats * &#x60;LVR&#x60; - Latvian Ruble * &#x60;LBP&#x60; - Lebanese Pound * &#x60;LSL&#x60; - Lesotho Loti * &#x60;LRD&#x60; - Liberian Dollar * &#x60;LYD&#x60; - Libyan Dinar * &#x60;LTL&#x60; - Lithuanian Litas * &#x60;LTT&#x60; - Lithuanian Talonas * &#x60;LUL&#x60; - Luxembourg Financial Franc * &#x60;LUC&#x60; - Luxembourgian Convertible Franc * &#x60;LUF&#x60; - Luxembourgian Franc * &#x60;MOP&#x60; - Macanese Pataca * &#x60;MKD&#x60; - Macedonian Denar * &#x60;MKN&#x60; - Macedonian Denar (1992–1993) * &#x60;MGA&#x60; - Malagasy Ariary * &#x60;MGF&#x60; - Malagasy Franc * &#x60;MWK&#x60; - Malawian Kwacha * &#x60;MYR&#x60; - Malaysian Ringgit * &#x60;MVR&#x60; - Maldivian Rufiyaa * &#x60;MVP&#x60; - Maldivian Rupee (1947–1981) * &#x60;MLF&#x60; - Malian Franc * &#x60;MTL&#x60; - Maltese Lira * &#x60;MTP&#x60; - Maltese Pound * &#x60;MRU&#x60; - Mauritanian Ouguiya * &#x60;MRO&#x60; - Mauritanian Ouguiya (1973–2017) * &#x60;MUR&#x60; - Mauritian Rupee * &#x60;MXV&#x60; - Mexican Investment Unit * &#x60;MXN&#x60; - Mexican Peso * &#x60;MXP&#x60; - Mexican Silver Peso (1861–1992) * &#x60;MDC&#x60; - Moldovan Cupon * &#x60;MDL&#x60; - Moldovan Leu * &#x60;MCF&#x60; - Monegasque Franc * &#x60;MNT&#x60; - Mongolian Tugrik * &#x60;MAD&#x60; - Moroccan Dirham * &#x60;MAF&#x60; - Moroccan Franc * &#x60;MZE&#x60; - Mozambican Escudo * &#x60;MZN&#x60; - Mozambican Metical * &#x60;MZM&#x60; - Mozambican Metical (1980–2006) * &#x60;MMK&#x60; - Myanmar Kyat * &#x60;NAD&#x60; - Namibian Dollar * &#x60;NPR&#x60; - Nepalese Rupee * &#x60;ANG&#x60; - Netherlands Antillean Guilder * &#x60;TWD&#x60; - New Taiwan Dollar * &#x60;NZD&#x60; - New Zealand Dollar * &#x60;NIO&#x60; - Nicaraguan Córdoba * &#x60;NIC&#x60; - Nicaraguan Córdoba (1988–1991) * &#x60;NGN&#x60; - Nigerian Naira * &#x60;KPW&#x60; - North Korean Won * &#x60;NOK&#x60; - Norwegian Krone * &#x60;OMR&#x60; - Omani Rial * &#x60;PKR&#x60; - Pakistani Rupee * &#x60;XPD&#x60; - Palladium * &#x60;PAB&#x60; - Panamanian Balboa * &#x60;PGK&#x60; - Papua New Guinean Kina * &#x60;PYG&#x60; - Paraguayan Guarani * &#x60;PEI&#x60; - Peruvian Inti * &#x60;PEN&#x60; - Peruvian Sol * &#x60;PES&#x60; - Peruvian Sol (1863–1965) * &#x60;PHP&#x60; - Philippine Peso * &#x60;XPT&#x60; - Platinum * &#x60;PLN&#x60; - Polish Zloty * &#x60;PLZ&#x60; - Polish Zloty (1950–1995) * &#x60;PTE&#x60; - Portuguese Escudo * &#x60;GWE&#x60; - Portuguese Guinea Escudo * &#x60;QAR&#x60; - Qatari Rial * &#x60;XRE&#x60; - RINET Funds * &#x60;RHD&#x60; - Rhodesian Dollar * &#x60;RON&#x60; - Romanian Leu * &#x60;ROL&#x60; - Romanian Leu (1952–2006) * &#x60;RUB&#x60; - Russian Ruble * &#x60;RUR&#x60; - Russian Ruble (1991–1998) * &#x60;RWF&#x60; - Rwandan Franc * &#x60;SVC&#x60; - Salvadoran Colón * &#x60;WST&#x60; - Samoan Tala * &#x60;SAR&#x60; - Saudi Riyal * &#x60;RSD&#x60; - Serbian Dinar * &#x60;CSD&#x60; - Serbian Dinar (2002–2006) * &#x60;SCR&#x60; - Seychellois Rupee * &#x60;SLL&#x60; - Sierra Leonean Leone * &#x60;XAG&#x60; - Silver * &#x60;SGD&#x60; - Singapore Dollar * &#x60;SKK&#x60; - Slovak Koruna * &#x60;SIT&#x60; - Slovenian Tolar * &#x60;SBD&#x60; - Solomon Islands Dollar * &#x60;SOS&#x60; - Somali Shilling * &#x60;ZAR&#x60; - South African Rand * &#x60;ZAL&#x60; - South African Rand (financial) * &#x60;KRH&#x60; - South Korean Hwan (1953–1962) * &#x60;KRW&#x60; - South Korean Won * &#x60;KRO&#x60; - South Korean Won (1945–1953) * &#x60;SSP&#x60; - South Sudanese Pound * &#x60;SUR&#x60; - Soviet Rouble * &#x60;ESP&#x60; - Spanish Peseta * &#x60;ESA&#x60; - Spanish Peseta (A account) * &#x60;ESB&#x60; - Spanish Peseta (convertible account) * &#x60;XDR&#x60; - Special Drawing Rights * &#x60;LKR&#x60; - Sri Lankan Rupee * &#x60;SHP&#x60; - St. Helena Pound * &#x60;XSU&#x60; - Sucre * &#x60;SDD&#x60; - Sudanese Dinar (1992–2007) * &#x60;SDG&#x60; - Sudanese Pound * &#x60;SDP&#x60; - Sudanese Pound (1957–1998) * &#x60;SRD&#x60; - Surinamese Dollar * &#x60;SRG&#x60; - Surinamese Guilder * &#x60;SZL&#x60; - Swazi Lilangeni * &#x60;SEK&#x60; - Swedish Krona * &#x60;CHF&#x60; - Swiss Franc * &#x60;SYP&#x60; - Syrian Pound * &#x60;STN&#x60; - São Tomé &amp; Príncipe Dobra * &#x60;STD&#x60; - São Tomé &amp; Príncipe Dobra (1977–2017) * &#x60;TVD&#x60; - TVD * &#x60;TJR&#x60; - Tajikistani Ruble * &#x60;TJS&#x60; - Tajikistani Somoni * &#x60;TZS&#x60; - Tanzanian Shilling * &#x60;XTS&#x60; - Testing Currency Code * &#x60;THB&#x60; - Thai Baht * &#x60;XXX&#x60; - The codes assigned for transactions where no currency is involved * &#x60;TPE&#x60; - Timorese Escudo * &#x60;TOP&#x60; - Tongan Paʻanga * &#x60;TTD&#x60; - Trinidad &amp; Tobago Dollar * &#x60;TND&#x60; - Tunisian Dinar * &#x60;TRY&#x60; - Turkish Lira * &#x60;TRL&#x60; - Turkish Lira (1922–2005) * &#x60;TMT&#x60; - Turkmenistani Manat * &#x60;TMM&#x60; - Turkmenistani Manat (1993–2009) * &#x60;USD&#x60; - US Dollar * &#x60;USN&#x60; - US Dollar (Next day) * &#x60;USS&#x60; - US Dollar (Same day) * &#x60;UGX&#x60; - Ugandan Shilling * &#x60;UGS&#x60; - Ugandan Shilling (1966–1987) * &#x60;UAH&#x60; - Ukrainian Hryvnia * &#x60;UAK&#x60; - Ukrainian Karbovanets * &#x60;AED&#x60; - United Arab Emirates Dirham * &#x60;UYW&#x60; - Uruguayan Nominal Wage Index Unit * &#x60;UYU&#x60; - Uruguayan Peso * &#x60;UYP&#x60; - Uruguayan Peso (1975–1993) * &#x60;UYI&#x60; - Uruguayan Peso (Indexed Units) * &#x60;UZS&#x60; - Uzbekistani Som * &#x60;VUV&#x60; - Vanuatu Vatu * &#x60;VES&#x60; - Venezuelan Bolívar * &#x60;VEB&#x60; - Venezuelan Bolívar (1871–2008) * &#x60;VEF&#x60; - Venezuelan Bolívar (2008–2018) * &#x60;VND&#x60; - Vietnamese Dong * &#x60;VNN&#x60; - Vietnamese Dong (1978–1985) * &#x60;CHE&#x60; - WIR Euro * &#x60;CHW&#x60; - WIR Franc * &#x60;XOF&#x60; - West African CFA Franc * &#x60;YDD&#x60; - Yemeni Dinar * &#x60;YER&#x60; - Yemeni Rial * &#x60;YUN&#x60; - Yugoslavian Convertible Dinar (1990–1992) * &#x60;YUD&#x60; - Yugoslavian Hard Dinar (1966–1990) * &#x60;YUM&#x60; - Yugoslavian New Dinar (1994–2002) * &#x60;YUR&#x60; - Yugoslavian Reformed Dinar (1992–1993) * &#x60;ZWN&#x60; - ZWN * &#x60;ZRN&#x60; - Zairean New Zaire (1993–1998) * &#x60;ZRZ&#x60; - Zairean Zaire (1971–1993) * &#x60;ZMW&#x60; - Zambian Kwacha * &#x60;ZMK&#x60; - Zambian Kwacha (1968–2012) * &#x60;ZWD&#x60; - Zimbabwean Dollar (1980–2008) * &#x60;ZWR&#x60; - Zimbabwean Dollar (2008) * &#x60;ZWL&#x60; - Zimbabwean Dollar (2009) | [optional] 
**ExchangeRate** | Pointer to **NullableFloat64** | The invoice&#39;s exchange rate. | [optional] 
**TotalDiscount** | Pointer to **NullableFloat64** | The total discounts applied to the total cost. | [optional] 
**SubTotal** | Pointer to **NullableFloat64** | The total amount being paid before taxes. | [optional] 
**TotalTaxAmount** | Pointer to **NullableFloat64** | The total amount being paid in taxes. | [optional] 
**TotalAmount** | Pointer to **NullableFloat64** | The invoice&#39;s total amount. | [optional] 
**Balance** | Pointer to **NullableFloat64** | The invoice&#39;s remaining balance. | [optional] 
**RemoteUpdatedAt** | Pointer to **NullableTime** | When the third party&#39;s invoice entry was updated. | [optional] 
**TrackingCategories** | Pointer to **[]string** |  | [optional] 
**Payments** | Pointer to **[]string** | Array of &#x60;Payment&#x60; object IDs. | [optional] 
**LineItems** | Pointer to [**[]InvoiceLineItem**](InvoiceLineItem.md) |  | [optional] [readonly] 
**RemoteWasDeleted** | Pointer to **bool** |  | [optional] [readonly] 
**RemoteId** | Pointer to **NullableString** | The third-party API ID of the matching object. | [optional] 
**FieldMappings** | Pointer to **map[string]interface{}** |  | [optional] [readonly] 
**ModifiedAt** | Pointer to **time.Time** | This is the datetime that this object was last updated by Merge | [optional] [readonly] 
**RemoteData** | Pointer to [**[]RemoteData**](RemoteData.md) |  | [optional] [readonly] 

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

`func (o *Invoice) GetTotalDiscount() float64`

GetTotalDiscount returns the TotalDiscount field if non-nil, zero value otherwise.

### GetTotalDiscountOk

`func (o *Invoice) GetTotalDiscountOk() (*float64, bool)`

GetTotalDiscountOk returns a tuple with the TotalDiscount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDiscount

`func (o *Invoice) SetTotalDiscount(v float64)`

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

`func (o *Invoice) GetSubTotal() float64`

GetSubTotal returns the SubTotal field if non-nil, zero value otherwise.

### GetSubTotalOk

`func (o *Invoice) GetSubTotalOk() (*float64, bool)`

GetSubTotalOk returns a tuple with the SubTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubTotal

`func (o *Invoice) SetSubTotal(v float64)`

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

`func (o *Invoice) GetTotalTaxAmount() float64`

GetTotalTaxAmount returns the TotalTaxAmount field if non-nil, zero value otherwise.

### GetTotalTaxAmountOk

`func (o *Invoice) GetTotalTaxAmountOk() (*float64, bool)`

GetTotalTaxAmountOk returns a tuple with the TotalTaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTaxAmount

`func (o *Invoice) SetTotalTaxAmount(v float64)`

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

`func (o *Invoice) GetTotalAmount() float64`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *Invoice) GetTotalAmountOk() (*float64, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *Invoice) SetTotalAmount(v float64)`

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

`func (o *Invoice) GetBalance() float64`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *Invoice) GetBalanceOk() (*float64, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *Invoice) SetBalance(v float64)`

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
### GetTrackingCategories

`func (o *Invoice) GetTrackingCategories() []string`

GetTrackingCategories returns the TrackingCategories field if non-nil, zero value otherwise.

### GetTrackingCategoriesOk

`func (o *Invoice) GetTrackingCategoriesOk() (*[]string, bool)`

GetTrackingCategoriesOk returns a tuple with the TrackingCategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingCategories

`func (o *Invoice) SetTrackingCategories(v []string)`

SetTrackingCategories sets TrackingCategories field to given value.

### HasTrackingCategories

`func (o *Invoice) HasTrackingCategories() bool`

HasTrackingCategories returns a boolean if a field has been set.

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
### GetModifiedAt

`func (o *Invoice) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *Invoice) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *Invoice) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *Invoice) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

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

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


