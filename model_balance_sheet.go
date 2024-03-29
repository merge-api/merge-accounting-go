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

// BalanceSheet # The BalanceSheet Object ### Description The `BalanceSheet` object shows a company’s assets, liabilities, and equity. Assets should be equal to liability and equity combined. This shows the company’s financial health at a specific point in time.  ### Usage Example Fetch from the `LIST BalanceSheets` endpoint and view a company's balance sheets.
type BalanceSheet struct {
	Id *string `json:"id,omitempty"`
	// The third-party API ID of the matching object.
	RemoteId NullableString `json:"remote_id,omitempty"`
	// The balance sheet's name.
	Name NullableString `json:"name,omitempty"`
	// The balance sheet's currency.  * `XUA` - ADB Unit of Account * `AFN` - Afghan Afghani * `AFA` - Afghan Afghani (1927–2002) * `ALL` - Albanian Lek * `ALK` - Albanian Lek (1946–1965) * `DZD` - Algerian Dinar * `ADP` - Andorran Peseta * `AOA` - Angolan Kwanza * `AOK` - Angolan Kwanza (1977–1991) * `AON` - Angolan New Kwanza (1990–2000) * `AOR` - Angolan Readjusted Kwanza (1995–1999) * `ARA` - Argentine Austral * `ARS` - Argentine Peso * `ARM` - Argentine Peso (1881–1970) * `ARP` - Argentine Peso (1983–1985) * `ARL` - Argentine Peso Ley (1970–1983) * `AMD` - Armenian Dram * `AWG` - Aruban Florin * `AUD` - Australian Dollar * `ATS` - Austrian Schilling * `AZN` - Azerbaijani Manat * `AZM` - Azerbaijani Manat (1993–2006) * `BSD` - Bahamian Dollar * `BHD` - Bahraini Dinar * `BDT` - Bangladeshi Taka * `BBD` - Barbadian Dollar * `BYN` - Belarusian Ruble * `BYB` - Belarusian Ruble (1994–1999) * `BYR` - Belarusian Ruble (2000–2016) * `BEF` - Belgian Franc * `BEC` - Belgian Franc (convertible) * `BEL` - Belgian Franc (financial) * `BZD` - Belize Dollar * `BMD` - Bermudan Dollar * `BTN` - Bhutanese Ngultrum * `BOB` - Bolivian Boliviano * `BOL` - Bolivian Boliviano (1863–1963) * `BOV` - Bolivian Mvdol * `BOP` - Bolivian Peso * `BAM` - Bosnia-Herzegovina Convertible Mark * `BAD` - Bosnia-Herzegovina Dinar (1992–1994) * `BAN` - Bosnia-Herzegovina New Dinar (1994–1997) * `BWP` - Botswanan Pula * `BRC` - Brazilian Cruzado (1986–1989) * `BRZ` - Brazilian Cruzeiro (1942–1967) * `BRE` - Brazilian Cruzeiro (1990–1993) * `BRR` - Brazilian Cruzeiro (1993–1994) * `BRN` - Brazilian New Cruzado (1989–1990) * `BRB` - Brazilian New Cruzeiro (1967–1986) * `BRL` - Brazilian Real * `GBP` - British Pound * `BND` - Brunei Dollar * `BGL` - Bulgarian Hard Lev * `BGN` - Bulgarian Lev * `BGO` - Bulgarian Lev (1879–1952) * `BGM` - Bulgarian Socialist Lev * `BUK` - Burmese Kyat * `BIF` - Burundian Franc * `XPF` - CFP Franc * `KHR` - Cambodian Riel * `CAD` - Canadian Dollar * `CVE` - Cape Verdean Escudo * `KYD` - Cayman Islands Dollar * `XAF` - Central African CFA Franc * `CLE` - Chilean Escudo * `CLP` - Chilean Peso * `CLF` - Chilean Unit of Account (UF) * `CNX` - Chinese People’s Bank Dollar * `CNY` - Chinese Yuan * `CNH` - Chinese Yuan (offshore) * `COP` - Colombian Peso * `COU` - Colombian Real Value Unit * `KMF` - Comorian Franc * `CDF` - Congolese Franc * `CRC` - Costa Rican Colón * `HRD` - Croatian Dinar * `HRK` - Croatian Kuna * `CUC` - Cuban Convertible Peso * `CUP` - Cuban Peso * `CYP` - Cypriot Pound * `CZK` - Czech Koruna * `CSK` - Czechoslovak Hard Koruna * `DKK` - Danish Krone * `DJF` - Djiboutian Franc * `DOP` - Dominican Peso * `NLG` - Dutch Guilder * `XCD` - East Caribbean Dollar * `DDM` - East German Mark * `ECS` - Ecuadorian Sucre * `ECV` - Ecuadorian Unit of Constant Value * `EGP` - Egyptian Pound * `GQE` - Equatorial Guinean Ekwele * `ERN` - Eritrean Nakfa * `EEK` - Estonian Kroon * `ETB` - Ethiopian Birr * `EUR` - Euro * `XBA` - European Composite Unit * `XEU` - European Currency Unit * `XBB` - European Monetary Unit * `XBC` - European Unit of Account (XBC) * `XBD` - European Unit of Account (XBD) * `FKP` - Falkland Islands Pound * `FJD` - Fijian Dollar * `FIM` - Finnish Markka * `FRF` - French Franc * `XFO` - French Gold Franc * `XFU` - French UIC-Franc * `GMD` - Gambian Dalasi * `GEK` - Georgian Kupon Larit * `GEL` - Georgian Lari * `DEM` - German Mark * `GHS` - Ghanaian Cedi * `GHC` - Ghanaian Cedi (1979–2007) * `GIP` - Gibraltar Pound * `XAU` - Gold * `GRD` - Greek Drachma * `GTQ` - Guatemalan Quetzal * `GWP` - Guinea-Bissau Peso * `GNF` - Guinean Franc * `GNS` - Guinean Syli * `GYD` - Guyanaese Dollar * `HTG` - Haitian Gourde * `HNL` - Honduran Lempira * `HKD` - Hong Kong Dollar * `HUF` - Hungarian Forint * `IMP` - IMP * `ISK` - Icelandic Króna * `ISJ` - Icelandic Króna (1918–1981) * `INR` - Indian Rupee * `IDR` - Indonesian Rupiah * `IRR` - Iranian Rial * `IQD` - Iraqi Dinar * `IEP` - Irish Pound * `ILS` - Israeli New Shekel * `ILP` - Israeli Pound * `ILR` - Israeli Shekel (1980–1985) * `ITL` - Italian Lira * `JMD` - Jamaican Dollar * `JPY` - Japanese Yen * `JOD` - Jordanian Dinar * `KZT` - Kazakhstani Tenge * `KES` - Kenyan Shilling * `KWD` - Kuwaiti Dinar * `KGS` - Kyrgystani Som * `LAK` - Laotian Kip * `LVL` - Latvian Lats * `LVR` - Latvian Ruble * `LBP` - Lebanese Pound * `LSL` - Lesotho Loti * `LRD` - Liberian Dollar * `LYD` - Libyan Dinar * `LTL` - Lithuanian Litas * `LTT` - Lithuanian Talonas * `LUL` - Luxembourg Financial Franc * `LUC` - Luxembourgian Convertible Franc * `LUF` - Luxembourgian Franc * `MOP` - Macanese Pataca * `MKD` - Macedonian Denar * `MKN` - Macedonian Denar (1992–1993) * `MGA` - Malagasy Ariary * `MGF` - Malagasy Franc * `MWK` - Malawian Kwacha * `MYR` - Malaysian Ringgit * `MVR` - Maldivian Rufiyaa * `MVP` - Maldivian Rupee (1947–1981) * `MLF` - Malian Franc * `MTL` - Maltese Lira * `MTP` - Maltese Pound * `MRU` - Mauritanian Ouguiya * `MRO` - Mauritanian Ouguiya (1973–2017) * `MUR` - Mauritian Rupee * `MXV` - Mexican Investment Unit * `MXN` - Mexican Peso * `MXP` - Mexican Silver Peso (1861–1992) * `MDC` - Moldovan Cupon * `MDL` - Moldovan Leu * `MCF` - Monegasque Franc * `MNT` - Mongolian Tugrik * `MAD` - Moroccan Dirham * `MAF` - Moroccan Franc * `MZE` - Mozambican Escudo * `MZN` - Mozambican Metical * `MZM` - Mozambican Metical (1980–2006) * `MMK` - Myanmar Kyat * `NAD` - Namibian Dollar * `NPR` - Nepalese Rupee * `ANG` - Netherlands Antillean Guilder * `TWD` - New Taiwan Dollar * `NZD` - New Zealand Dollar * `NIO` - Nicaraguan Córdoba * `NIC` - Nicaraguan Córdoba (1988–1991) * `NGN` - Nigerian Naira * `KPW` - North Korean Won * `NOK` - Norwegian Krone * `OMR` - Omani Rial * `PKR` - Pakistani Rupee * `XPD` - Palladium * `PAB` - Panamanian Balboa * `PGK` - Papua New Guinean Kina * `PYG` - Paraguayan Guarani * `PEI` - Peruvian Inti * `PEN` - Peruvian Sol * `PES` - Peruvian Sol (1863–1965) * `PHP` - Philippine Peso * `XPT` - Platinum * `PLN` - Polish Zloty * `PLZ` - Polish Zloty (1950–1995) * `PTE` - Portuguese Escudo * `GWE` - Portuguese Guinea Escudo * `QAR` - Qatari Rial * `XRE` - RINET Funds * `RHD` - Rhodesian Dollar * `RON` - Romanian Leu * `ROL` - Romanian Leu (1952–2006) * `RUB` - Russian Ruble * `RUR` - Russian Ruble (1991–1998) * `RWF` - Rwandan Franc * `SVC` - Salvadoran Colón * `WST` - Samoan Tala * `SAR` - Saudi Riyal * `RSD` - Serbian Dinar * `CSD` - Serbian Dinar (2002–2006) * `SCR` - Seychellois Rupee * `SLL` - Sierra Leonean Leone * `XAG` - Silver * `SGD` - Singapore Dollar * `SKK` - Slovak Koruna * `SIT` - Slovenian Tolar * `SBD` - Solomon Islands Dollar * `SOS` - Somali Shilling * `ZAR` - South African Rand * `ZAL` - South African Rand (financial) * `KRH` - South Korean Hwan (1953–1962) * `KRW` - South Korean Won * `KRO` - South Korean Won (1945–1953) * `SSP` - South Sudanese Pound * `SUR` - Soviet Rouble * `ESP` - Spanish Peseta * `ESA` - Spanish Peseta (A account) * `ESB` - Spanish Peseta (convertible account) * `XDR` - Special Drawing Rights * `LKR` - Sri Lankan Rupee * `SHP` - St. Helena Pound * `XSU` - Sucre * `SDD` - Sudanese Dinar (1992–2007) * `SDG` - Sudanese Pound * `SDP` - Sudanese Pound (1957–1998) * `SRD` - Surinamese Dollar * `SRG` - Surinamese Guilder * `SZL` - Swazi Lilangeni * `SEK` - Swedish Krona * `CHF` - Swiss Franc * `SYP` - Syrian Pound * `STN` - São Tomé & Príncipe Dobra * `STD` - São Tomé & Príncipe Dobra (1977–2017) * `TVD` - TVD * `TJR` - Tajikistani Ruble * `TJS` - Tajikistani Somoni * `TZS` - Tanzanian Shilling * `XTS` - Testing Currency Code * `THB` - Thai Baht * `XXX` - The codes assigned for transactions where no currency is involved * `TPE` - Timorese Escudo * `TOP` - Tongan Paʻanga * `TTD` - Trinidad & Tobago Dollar * `TND` - Tunisian Dinar * `TRY` - Turkish Lira * `TRL` - Turkish Lira (1922–2005) * `TMT` - Turkmenistani Manat * `TMM` - Turkmenistani Manat (1993–2009) * `USD` - US Dollar * `USN` - US Dollar (Next day) * `USS` - US Dollar (Same day) * `UGX` - Ugandan Shilling * `UGS` - Ugandan Shilling (1966–1987) * `UAH` - Ukrainian Hryvnia * `UAK` - Ukrainian Karbovanets * `AED` - United Arab Emirates Dirham * `UYW` - Uruguayan Nominal Wage Index Unit * `UYU` - Uruguayan Peso * `UYP` - Uruguayan Peso (1975–1993) * `UYI` - Uruguayan Peso (Indexed Units) * `UZS` - Uzbekistani Som * `VUV` - Vanuatu Vatu * `VES` - Venezuelan Bolívar * `VEB` - Venezuelan Bolívar (1871–2008) * `VEF` - Venezuelan Bolívar (2008–2018) * `VND` - Vietnamese Dong * `VNN` - Vietnamese Dong (1978–1985) * `CHE` - WIR Euro * `CHW` - WIR Franc * `XOF` - West African CFA Franc * `YDD` - Yemeni Dinar * `YER` - Yemeni Rial * `YUN` - Yugoslavian Convertible Dinar (1990–1992) * `YUD` - Yugoslavian Hard Dinar (1966–1990) * `YUM` - Yugoslavian New Dinar (1994–2002) * `YUR` - Yugoslavian Reformed Dinar (1992–1993) * `ZWN` - ZWN * `ZRN` - Zairean New Zaire (1993–1998) * `ZRZ` - Zairean Zaire (1971–1993) * `ZMW` - Zambian Kwacha * `ZMK` - Zambian Kwacha (1968–2012) * `ZWD` - Zimbabwean Dollar (1980–2008) * `ZWR` - Zimbabwean Dollar (2008) * `ZWL` - Zimbabwean Dollar (2009)
	Currency NullableCurrencyEnum `json:"currency,omitempty"`
	// `Company` object for the given `BalanceSheet` object.
	Company NullableString `json:"company,omitempty"`
	// The balance sheet's date. The balance sheet data will reflect the company's financial position this point in time.
	Date NullableTime `json:"date,omitempty"`
	// The balance sheet's net assets.
	NetAssets NullableFloat64 `json:"net_assets,omitempty"`
	Assets *[]ReportItem `json:"assets,omitempty"`
	Liabilities *[]ReportItem `json:"liabilities,omitempty"`
	Equity *[]ReportItem `json:"equity,omitempty"`
	// The time that balance sheet was generated by the accounting system.
	RemoteGeneratedAt NullableTime `json:"remote_generated_at,omitempty"`
	// Indicates whether or not this object has been deleted by third party webhooks.
	RemoteWasDeleted *bool `json:"remote_was_deleted,omitempty"`
	FieldMappings map[string]interface{} `json:"field_mappings,omitempty"`
	// This is the datetime that this object was last updated by Merge
	ModifiedAt *time.Time `json:"modified_at,omitempty"`
	RemoteData []RemoteData `json:"remote_data,omitempty"`
	// raw json response by property name
	ResponseRaw map[string]json.RawMessage `json:"-"`
}

// NewBalanceSheet instantiates a new BalanceSheet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBalanceSheet() *BalanceSheet {
	this := BalanceSheet{}
	return &this
}

// NewBalanceSheetWithDefaults instantiates a new BalanceSheet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBalanceSheetWithDefaults() *BalanceSheet {
	this := BalanceSheet{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BalanceSheet) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BalanceSheet) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BalanceSheet) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BalanceSheet) SetId(v string) {
	o.Id = &v
}

// GetRemoteId returns the RemoteId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BalanceSheet) GetRemoteId() string {
	if o == nil || o.RemoteId.Get() == nil {
		var ret string
		return ret
	}
	return *o.RemoteId.Get()
}

// GetRemoteIdOk returns a tuple with the RemoteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BalanceSheet) GetRemoteIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RemoteId.Get(), o.RemoteId.IsSet()
}

// HasRemoteId returns a boolean if a field has been set.
func (o *BalanceSheet) HasRemoteId() bool {
	if o != nil && o.RemoteId.IsSet() {
		return true
	}

	return false
}

// SetRemoteId gets a reference to the given NullableString and assigns it to the RemoteId field.
func (o *BalanceSheet) SetRemoteId(v string) {
	o.RemoteId.Set(&v)
}
// SetRemoteIdNil sets the value for RemoteId to be an explicit nil
func (o *BalanceSheet) SetRemoteIdNil() {
	o.RemoteId.Set(nil)
}

// UnsetRemoteId ensures that no value is present for RemoteId, not even an explicit nil
func (o *BalanceSheet) UnsetRemoteId() {
	o.RemoteId.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BalanceSheet) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BalanceSheet) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *BalanceSheet) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *BalanceSheet) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *BalanceSheet) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *BalanceSheet) UnsetName() {
	o.Name.Unset()
}

// GetCurrency returns the Currency field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BalanceSheet) GetCurrency() CurrencyEnum {
	if o == nil || o.Currency.Get() == nil {
		var ret CurrencyEnum
		return ret
	}
	return *o.Currency.Get()
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BalanceSheet) GetCurrencyOk() (*CurrencyEnum, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Currency.Get(), o.Currency.IsSet()
}

// HasCurrency returns a boolean if a field has been set.
func (o *BalanceSheet) HasCurrency() bool {
	if o != nil && o.Currency.IsSet() {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given NullableCurrencyEnum and assigns it to the Currency field.
func (o *BalanceSheet) SetCurrency(v CurrencyEnum) {
	o.Currency.Set(&v)
}
// SetCurrencyNil sets the value for Currency to be an explicit nil
func (o *BalanceSheet) SetCurrencyNil() {
	o.Currency.Set(nil)
}

// UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
func (o *BalanceSheet) UnsetCurrency() {
	o.Currency.Unset()
}

// GetCompany returns the Company field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BalanceSheet) GetCompany() string {
	if o == nil || o.Company.Get() == nil {
		var ret string
		return ret
	}
	return *o.Company.Get()
}

// GetCompanyOk returns a tuple with the Company field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BalanceSheet) GetCompanyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Company.Get(), o.Company.IsSet()
}

// HasCompany returns a boolean if a field has been set.
func (o *BalanceSheet) HasCompany() bool {
	if o != nil && o.Company.IsSet() {
		return true
	}

	return false
}

// SetCompany gets a reference to the given NullableString and assigns it to the Company field.
func (o *BalanceSheet) SetCompany(v string) {
	o.Company.Set(&v)
}
// SetCompanyNil sets the value for Company to be an explicit nil
func (o *BalanceSheet) SetCompanyNil() {
	o.Company.Set(nil)
}

// UnsetCompany ensures that no value is present for Company, not even an explicit nil
func (o *BalanceSheet) UnsetCompany() {
	o.Company.Unset()
}

// GetDate returns the Date field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BalanceSheet) GetDate() time.Time {
	if o == nil || o.Date.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.Date.Get()
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BalanceSheet) GetDateOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Date.Get(), o.Date.IsSet()
}

// HasDate returns a boolean if a field has been set.
func (o *BalanceSheet) HasDate() bool {
	if o != nil && o.Date.IsSet() {
		return true
	}

	return false
}

// SetDate gets a reference to the given NullableTime and assigns it to the Date field.
func (o *BalanceSheet) SetDate(v time.Time) {
	o.Date.Set(&v)
}
// SetDateNil sets the value for Date to be an explicit nil
func (o *BalanceSheet) SetDateNil() {
	o.Date.Set(nil)
}

// UnsetDate ensures that no value is present for Date, not even an explicit nil
func (o *BalanceSheet) UnsetDate() {
	o.Date.Unset()
}

// GetNetAssets returns the NetAssets field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BalanceSheet) GetNetAssets() float64 {
	if o == nil || o.NetAssets.Get() == nil {
		var ret float64
		return ret
	}
	return *o.NetAssets.Get()
}

// GetNetAssetsOk returns a tuple with the NetAssets field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BalanceSheet) GetNetAssetsOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.NetAssets.Get(), o.NetAssets.IsSet()
}

// HasNetAssets returns a boolean if a field has been set.
func (o *BalanceSheet) HasNetAssets() bool {
	if o != nil && o.NetAssets.IsSet() {
		return true
	}

	return false
}

// SetNetAssets gets a reference to the given NullableFloat64 and assigns it to the NetAssets field.
func (o *BalanceSheet) SetNetAssets(v float64) {
	o.NetAssets.Set(&v)
}
// SetNetAssetsNil sets the value for NetAssets to be an explicit nil
func (o *BalanceSheet) SetNetAssetsNil() {
	o.NetAssets.Set(nil)
}

// UnsetNetAssets ensures that no value is present for NetAssets, not even an explicit nil
func (o *BalanceSheet) UnsetNetAssets() {
	o.NetAssets.Unset()
}

// GetAssets returns the Assets field value if set, zero value otherwise.
func (o *BalanceSheet) GetAssets() []ReportItem {
	if o == nil || o.Assets == nil {
		var ret []ReportItem
		return ret
	}
	return *o.Assets
}

// GetAssetsOk returns a tuple with the Assets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BalanceSheet) GetAssetsOk() (*[]ReportItem, bool) {
	if o == nil || o.Assets == nil {
		return nil, false
	}
	return o.Assets, true
}

// HasAssets returns a boolean if a field has been set.
func (o *BalanceSheet) HasAssets() bool {
	if o != nil && o.Assets != nil {
		return true
	}

	return false
}

// SetAssets gets a reference to the given []ReportItem and assigns it to the Assets field.
func (o *BalanceSheet) SetAssets(v []ReportItem) {
	o.Assets = &v
}

// GetLiabilities returns the Liabilities field value if set, zero value otherwise.
func (o *BalanceSheet) GetLiabilities() []ReportItem {
	if o == nil || o.Liabilities == nil {
		var ret []ReportItem
		return ret
	}
	return *o.Liabilities
}

// GetLiabilitiesOk returns a tuple with the Liabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BalanceSheet) GetLiabilitiesOk() (*[]ReportItem, bool) {
	if o == nil || o.Liabilities == nil {
		return nil, false
	}
	return o.Liabilities, true
}

// HasLiabilities returns a boolean if a field has been set.
func (o *BalanceSheet) HasLiabilities() bool {
	if o != nil && o.Liabilities != nil {
		return true
	}

	return false
}

// SetLiabilities gets a reference to the given []ReportItem and assigns it to the Liabilities field.
func (o *BalanceSheet) SetLiabilities(v []ReportItem) {
	o.Liabilities = &v
}

// GetEquity returns the Equity field value if set, zero value otherwise.
func (o *BalanceSheet) GetEquity() []ReportItem {
	if o == nil || o.Equity == nil {
		var ret []ReportItem
		return ret
	}
	return *o.Equity
}

// GetEquityOk returns a tuple with the Equity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BalanceSheet) GetEquityOk() (*[]ReportItem, bool) {
	if o == nil || o.Equity == nil {
		return nil, false
	}
	return o.Equity, true
}

// HasEquity returns a boolean if a field has been set.
func (o *BalanceSheet) HasEquity() bool {
	if o != nil && o.Equity != nil {
		return true
	}

	return false
}

// SetEquity gets a reference to the given []ReportItem and assigns it to the Equity field.
func (o *BalanceSheet) SetEquity(v []ReportItem) {
	o.Equity = &v
}

// GetRemoteGeneratedAt returns the RemoteGeneratedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BalanceSheet) GetRemoteGeneratedAt() time.Time {
	if o == nil || o.RemoteGeneratedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.RemoteGeneratedAt.Get()
}

// GetRemoteGeneratedAtOk returns a tuple with the RemoteGeneratedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BalanceSheet) GetRemoteGeneratedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RemoteGeneratedAt.Get(), o.RemoteGeneratedAt.IsSet()
}

// HasRemoteGeneratedAt returns a boolean if a field has been set.
func (o *BalanceSheet) HasRemoteGeneratedAt() bool {
	if o != nil && o.RemoteGeneratedAt.IsSet() {
		return true
	}

	return false
}

// SetRemoteGeneratedAt gets a reference to the given NullableTime and assigns it to the RemoteGeneratedAt field.
func (o *BalanceSheet) SetRemoteGeneratedAt(v time.Time) {
	o.RemoteGeneratedAt.Set(&v)
}
// SetRemoteGeneratedAtNil sets the value for RemoteGeneratedAt to be an explicit nil
func (o *BalanceSheet) SetRemoteGeneratedAtNil() {
	o.RemoteGeneratedAt.Set(nil)
}

// UnsetRemoteGeneratedAt ensures that no value is present for RemoteGeneratedAt, not even an explicit nil
func (o *BalanceSheet) UnsetRemoteGeneratedAt() {
	o.RemoteGeneratedAt.Unset()
}

// GetRemoteWasDeleted returns the RemoteWasDeleted field value if set, zero value otherwise.
func (o *BalanceSheet) GetRemoteWasDeleted() bool {
	if o == nil || o.RemoteWasDeleted == nil {
		var ret bool
		return ret
	}
	return *o.RemoteWasDeleted
}

// GetRemoteWasDeletedOk returns a tuple with the RemoteWasDeleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BalanceSheet) GetRemoteWasDeletedOk() (*bool, bool) {
	if o == nil || o.RemoteWasDeleted == nil {
		return nil, false
	}
	return o.RemoteWasDeleted, true
}

// HasRemoteWasDeleted returns a boolean if a field has been set.
func (o *BalanceSheet) HasRemoteWasDeleted() bool {
	if o != nil && o.RemoteWasDeleted != nil {
		return true
	}

	return false
}

// SetRemoteWasDeleted gets a reference to the given bool and assigns it to the RemoteWasDeleted field.
func (o *BalanceSheet) SetRemoteWasDeleted(v bool) {
	o.RemoteWasDeleted = &v
}

// GetFieldMappings returns the FieldMappings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BalanceSheet) GetFieldMappings() map[string]interface{} {
	if o == nil  {
		var ret map[string]interface{}
		return ret
	}
	return o.FieldMappings
}

// GetFieldMappingsOk returns a tuple with the FieldMappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BalanceSheet) GetFieldMappingsOk() (*map[string]interface{}, bool) {
	if o == nil || o.FieldMappings == nil {
		return nil, false
	}
	return &o.FieldMappings, true
}

// HasFieldMappings returns a boolean if a field has been set.
func (o *BalanceSheet) HasFieldMappings() bool {
	if o != nil && o.FieldMappings != nil {
		return true
	}

	return false
}

// SetFieldMappings gets a reference to the given map[string]interface{} and assigns it to the FieldMappings field.
func (o *BalanceSheet) SetFieldMappings(v map[string]interface{}) {
	o.FieldMappings = v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise.
func (o *BalanceSheet) GetModifiedAt() time.Time {
	if o == nil || o.ModifiedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BalanceSheet) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil || o.ModifiedAt == nil {
		return nil, false
	}
	return o.ModifiedAt, true
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *BalanceSheet) HasModifiedAt() bool {
	if o != nil && o.ModifiedAt != nil {
		return true
	}

	return false
}

// SetModifiedAt gets a reference to the given time.Time and assigns it to the ModifiedAt field.
func (o *BalanceSheet) SetModifiedAt(v time.Time) {
	o.ModifiedAt = &v
}

// GetRemoteData returns the RemoteData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BalanceSheet) GetRemoteData() []RemoteData {
	if o == nil  {
		var ret []RemoteData
		return ret
	}
	return o.RemoteData
}

// GetRemoteDataOk returns a tuple with the RemoteData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BalanceSheet) GetRemoteDataOk() (*[]RemoteData, bool) {
	if o == nil || o.RemoteData == nil {
		return nil, false
	}
	return &o.RemoteData, true
}

// HasRemoteData returns a boolean if a field has been set.
func (o *BalanceSheet) HasRemoteData() bool {
	if o != nil && o.RemoteData != nil {
		return true
	}

	return false
}

// SetRemoteData gets a reference to the given []RemoteData and assigns it to the RemoteData field.
func (o *BalanceSheet) SetRemoteData(v []RemoteData) {
	o.RemoteData = v
}

func (o BalanceSheet) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.RemoteId.IsSet() {
		toSerialize["remote_id"] = o.RemoteId.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Currency.IsSet() {
		toSerialize["currency"] = o.Currency.Get()
	}
	if o.Company.IsSet() {
		toSerialize["company"] = o.Company.Get()
	}
	if o.Date.IsSet() {
		toSerialize["date"] = o.Date.Get()
	}
	if o.NetAssets.IsSet() {
		toSerialize["net_assets"] = o.NetAssets.Get()
	}
	if o.Assets != nil {
		toSerialize["assets"] = o.Assets
	}
	if o.Liabilities != nil {
		toSerialize["liabilities"] = o.Liabilities
	}
	if o.Equity != nil {
		toSerialize["equity"] = o.Equity
	}
	if o.RemoteGeneratedAt.IsSet() {
		toSerialize["remote_generated_at"] = o.RemoteGeneratedAt.Get()
	}
	if o.RemoteWasDeleted != nil {
		toSerialize["remote_was_deleted"] = o.RemoteWasDeleted
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

func (v *BalanceSheet) UnmarshalJSON(src []byte) error {
    type BalanceSheetUnmarshalTarget BalanceSheet

	var intermediateResult BalanceSheetUnmarshalTarget
	var err1 = json.Unmarshal(src, &intermediateResult)
    if err1 != nil {
        return err1
    }
    var err2 = json.Unmarshal(src, &intermediateResult.ResponseRaw)
	if err2 != nil {
		return err2
	}

	*v = BalanceSheet(intermediateResult)
	return nil
}
type NullableBalanceSheet struct {
	value *BalanceSheet
	isSet bool
}

func (v NullableBalanceSheet) Get() *BalanceSheet {
	return v.value
}

func (v *NullableBalanceSheet) Set(val *BalanceSheet) {
	v.value = val
	v.isSet = true
}

func (v NullableBalanceSheet) IsSet() bool {
	return v.isSet
}

func (v *NullableBalanceSheet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBalanceSheet(val *BalanceSheet) *NullableBalanceSheet {
	return &NullableBalanceSheet{value: val, isSet: true}
}

func (v NullableBalanceSheet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBalanceSheet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	var err1 = json.Unmarshal(src, &v.value)
    if err1 != nil {
        return err1
    }
    return json.Unmarshal(src, &v.value.ResponseRaw)
}


