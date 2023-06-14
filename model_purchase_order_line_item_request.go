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
)

// PurchaseOrderLineItemRequest # The PurchaseOrderLineItem Object ### Description The `PurchaseOrderLineItem` object is used to represent a purchase order's line item.  ### Usage Example Fetch from the `GET PurchaseOrder` endpoint and view a company's purchase orders.
type PurchaseOrderLineItemRequest struct {
	// The third-party API ID of the matching object.
	RemoteId NullableString `json:"remote_id,omitempty"`
	// A description of the good being purchased.
	Description NullableString `json:"description,omitempty"`
	// The line item's unit price.
	UnitPrice NullableFloat64 `json:"unit_price,omitempty"`
	// The line item's quantity.
	Quantity NullableFloat64 `json:"quantity,omitempty"`
	Item NullableString `json:"item,omitempty"`
	// The purchase order line item's account.
	Account NullableString `json:"account,omitempty"`
	// The purchase order line item's associated tracking category.
	TrackingCategory NullableString `json:"tracking_category,omitempty"`
	// The purchase order line item's associated tracking categories.
	TrackingCategories []string `json:"tracking_categories"`
	// The purchase order line item's tax amount.
	TaxAmount NullableFloat64 `json:"tax_amount,omitempty"`
	// The purchase order line item's total amount.
	TotalLineAmount NullableFloat64 `json:"total_line_amount,omitempty"`
	// The purchase order line item's currency.  * `XUA` - ADB Unit of Account * `AFN` - Afghan Afghani * `AFA` - Afghan Afghani (1927–2002) * `ALL` - Albanian Lek * `ALK` - Albanian Lek (1946–1965) * `DZD` - Algerian Dinar * `ADP` - Andorran Peseta * `AOA` - Angolan Kwanza * `AOK` - Angolan Kwanza (1977–1991) * `AON` - Angolan New Kwanza (1990–2000) * `AOR` - Angolan Readjusted Kwanza (1995–1999) * `ARA` - Argentine Austral * `ARS` - Argentine Peso * `ARM` - Argentine Peso (1881–1970) * `ARP` - Argentine Peso (1983–1985) * `ARL` - Argentine Peso Ley (1970–1983) * `AMD` - Armenian Dram * `AWG` - Aruban Florin * `AUD` - Australian Dollar * `ATS` - Austrian Schilling * `AZN` - Azerbaijani Manat * `AZM` - Azerbaijani Manat (1993–2006) * `BSD` - Bahamian Dollar * `BHD` - Bahraini Dinar * `BDT` - Bangladeshi Taka * `BBD` - Barbadian Dollar * `BYN` - Belarusian Ruble * `BYB` - Belarusian Ruble (1994–1999) * `BYR` - Belarusian Ruble (2000–2016) * `BEF` - Belgian Franc * `BEC` - Belgian Franc (convertible) * `BEL` - Belgian Franc (financial) * `BZD` - Belize Dollar * `BMD` - Bermudan Dollar * `BTN` - Bhutanese Ngultrum * `BOB` - Bolivian Boliviano * `BOL` - Bolivian Boliviano (1863–1963) * `BOV` - Bolivian Mvdol * `BOP` - Bolivian Peso * `BAM` - Bosnia-Herzegovina Convertible Mark * `BAD` - Bosnia-Herzegovina Dinar (1992–1994) * `BAN` - Bosnia-Herzegovina New Dinar (1994–1997) * `BWP` - Botswanan Pula * `BRC` - Brazilian Cruzado (1986–1989) * `BRZ` - Brazilian Cruzeiro (1942–1967) * `BRE` - Brazilian Cruzeiro (1990–1993) * `BRR` - Brazilian Cruzeiro (1993–1994) * `BRN` - Brazilian New Cruzado (1989–1990) * `BRB` - Brazilian New Cruzeiro (1967–1986) * `BRL` - Brazilian Real * `GBP` - British Pound * `BND` - Brunei Dollar * `BGL` - Bulgarian Hard Lev * `BGN` - Bulgarian Lev * `BGO` - Bulgarian Lev (1879–1952) * `BGM` - Bulgarian Socialist Lev * `BUK` - Burmese Kyat * `BIF` - Burundian Franc * `XPF` - CFP Franc * `KHR` - Cambodian Riel * `CAD` - Canadian Dollar * `CVE` - Cape Verdean Escudo * `KYD` - Cayman Islands Dollar * `XAF` - Central African CFA Franc * `CLE` - Chilean Escudo * `CLP` - Chilean Peso * `CLF` - Chilean Unit of Account (UF) * `CNX` - Chinese People’s Bank Dollar * `CNY` - Chinese Yuan * `CNH` - Chinese Yuan (offshore) * `COP` - Colombian Peso * `COU` - Colombian Real Value Unit * `KMF` - Comorian Franc * `CDF` - Congolese Franc * `CRC` - Costa Rican Colón * `HRD` - Croatian Dinar * `HRK` - Croatian Kuna * `CUC` - Cuban Convertible Peso * `CUP` - Cuban Peso * `CYP` - Cypriot Pound * `CZK` - Czech Koruna * `CSK` - Czechoslovak Hard Koruna * `DKK` - Danish Krone * `DJF` - Djiboutian Franc * `DOP` - Dominican Peso * `NLG` - Dutch Guilder * `XCD` - East Caribbean Dollar * `DDM` - East German Mark * `ECS` - Ecuadorian Sucre * `ECV` - Ecuadorian Unit of Constant Value * `EGP` - Egyptian Pound * `GQE` - Equatorial Guinean Ekwele * `ERN` - Eritrean Nakfa * `EEK` - Estonian Kroon * `ETB` - Ethiopian Birr * `EUR` - Euro * `XBA` - European Composite Unit * `XEU` - European Currency Unit * `XBB` - European Monetary Unit * `XBC` - European Unit of Account (XBC) * `XBD` - European Unit of Account (XBD) * `FKP` - Falkland Islands Pound * `FJD` - Fijian Dollar * `FIM` - Finnish Markka * `FRF` - French Franc * `XFO` - French Gold Franc * `XFU` - French UIC-Franc * `GMD` - Gambian Dalasi * `GEK` - Georgian Kupon Larit * `GEL` - Georgian Lari * `DEM` - German Mark * `GHS` - Ghanaian Cedi * `GHC` - Ghanaian Cedi (1979–2007) * `GIP` - Gibraltar Pound * `XAU` - Gold * `GRD` - Greek Drachma * `GTQ` - Guatemalan Quetzal * `GWP` - Guinea-Bissau Peso * `GNF` - Guinean Franc * `GNS` - Guinean Syli * `GYD` - Guyanaese Dollar * `HTG` - Haitian Gourde * `HNL` - Honduran Lempira * `HKD` - Hong Kong Dollar * `HUF` - Hungarian Forint * `IMP` - IMP * `ISK` - Icelandic Króna * `ISJ` - Icelandic Króna (1918–1981) * `INR` - Indian Rupee * `IDR` - Indonesian Rupiah * `IRR` - Iranian Rial * `IQD` - Iraqi Dinar * `IEP` - Irish Pound * `ILS` - Israeli New Shekel * `ILP` - Israeli Pound * `ILR` - Israeli Shekel (1980–1985) * `ITL` - Italian Lira * `JMD` - Jamaican Dollar * `JPY` - Japanese Yen * `JOD` - Jordanian Dinar * `KZT` - Kazakhstani Tenge * `KES` - Kenyan Shilling * `KWD` - Kuwaiti Dinar * `KGS` - Kyrgystani Som * `LAK` - Laotian Kip * `LVL` - Latvian Lats * `LVR` - Latvian Ruble * `LBP` - Lebanese Pound * `LSL` - Lesotho Loti * `LRD` - Liberian Dollar * `LYD` - Libyan Dinar * `LTL` - Lithuanian Litas * `LTT` - Lithuanian Talonas * `LUL` - Luxembourg Financial Franc * `LUC` - Luxembourgian Convertible Franc * `LUF` - Luxembourgian Franc * `MOP` - Macanese Pataca * `MKD` - Macedonian Denar * `MKN` - Macedonian Denar (1992–1993) * `MGA` - Malagasy Ariary * `MGF` - Malagasy Franc * `MWK` - Malawian Kwacha * `MYR` - Malaysian Ringgit * `MVR` - Maldivian Rufiyaa * `MVP` - Maldivian Rupee (1947–1981) * `MLF` - Malian Franc * `MTL` - Maltese Lira * `MTP` - Maltese Pound * `MRU` - Mauritanian Ouguiya * `MRO` - Mauritanian Ouguiya (1973–2017) * `MUR` - Mauritian Rupee * `MXV` - Mexican Investment Unit * `MXN` - Mexican Peso * `MXP` - Mexican Silver Peso (1861–1992) * `MDC` - Moldovan Cupon * `MDL` - Moldovan Leu * `MCF` - Monegasque Franc * `MNT` - Mongolian Tugrik * `MAD` - Moroccan Dirham * `MAF` - Moroccan Franc * `MZE` - Mozambican Escudo * `MZN` - Mozambican Metical * `MZM` - Mozambican Metical (1980–2006) * `MMK` - Myanmar Kyat * `NAD` - Namibian Dollar * `NPR` - Nepalese Rupee * `ANG` - Netherlands Antillean Guilder * `TWD` - New Taiwan Dollar * `NZD` - New Zealand Dollar * `NIO` - Nicaraguan Córdoba * `NIC` - Nicaraguan Córdoba (1988–1991) * `NGN` - Nigerian Naira * `KPW` - North Korean Won * `NOK` - Norwegian Krone * `OMR` - Omani Rial * `PKR` - Pakistani Rupee * `XPD` - Palladium * `PAB` - Panamanian Balboa * `PGK` - Papua New Guinean Kina * `PYG` - Paraguayan Guarani * `PEI` - Peruvian Inti * `PEN` - Peruvian Sol * `PES` - Peruvian Sol (1863–1965) * `PHP` - Philippine Peso * `XPT` - Platinum * `PLN` - Polish Zloty * `PLZ` - Polish Zloty (1950–1995) * `PTE` - Portuguese Escudo * `GWE` - Portuguese Guinea Escudo * `QAR` - Qatari Rial * `XRE` - RINET Funds * `RHD` - Rhodesian Dollar * `RON` - Romanian Leu * `ROL` - Romanian Leu (1952–2006) * `RUB` - Russian Ruble * `RUR` - Russian Ruble (1991–1998) * `RWF` - Rwandan Franc * `SVC` - Salvadoran Colón * `WST` - Samoan Tala * `SAR` - Saudi Riyal * `RSD` - Serbian Dinar * `CSD` - Serbian Dinar (2002–2006) * `SCR` - Seychellois Rupee * `SLL` - Sierra Leonean Leone * `XAG` - Silver * `SGD` - Singapore Dollar * `SKK` - Slovak Koruna * `SIT` - Slovenian Tolar * `SBD` - Solomon Islands Dollar * `SOS` - Somali Shilling * `ZAR` - South African Rand * `ZAL` - South African Rand (financial) * `KRH` - South Korean Hwan (1953–1962) * `KRW` - South Korean Won * `KRO` - South Korean Won (1945–1953) * `SSP` - South Sudanese Pound * `SUR` - Soviet Rouble * `ESP` - Spanish Peseta * `ESA` - Spanish Peseta (A account) * `ESB` - Spanish Peseta (convertible account) * `XDR` - Special Drawing Rights * `LKR` - Sri Lankan Rupee * `SHP` - St. Helena Pound * `XSU` - Sucre * `SDD` - Sudanese Dinar (1992–2007) * `SDG` - Sudanese Pound * `SDP` - Sudanese Pound (1957–1998) * `SRD` - Surinamese Dollar * `SRG` - Surinamese Guilder * `SZL` - Swazi Lilangeni * `SEK` - Swedish Krona * `CHF` - Swiss Franc * `SYP` - Syrian Pound * `STN` - São Tomé & Príncipe Dobra * `STD` - São Tomé & Príncipe Dobra (1977–2017) * `TVD` - TVD * `TJR` - Tajikistani Ruble * `TJS` - Tajikistani Somoni * `TZS` - Tanzanian Shilling * `XTS` - Testing Currency Code * `THB` - Thai Baht * `XXX` - The codes assigned for transactions where no currency is involved * `TPE` - Timorese Escudo * `TOP` - Tongan Paʻanga * `TTD` - Trinidad & Tobago Dollar * `TND` - Tunisian Dinar * `TRY` - Turkish Lira * `TRL` - Turkish Lira (1922–2005) * `TMT` - Turkmenistani Manat * `TMM` - Turkmenistani Manat (1993–2009) * `USD` - US Dollar * `USN` - US Dollar (Next day) * `USS` - US Dollar (Same day) * `UGX` - Ugandan Shilling * `UGS` - Ugandan Shilling (1966–1987) * `UAH` - Ukrainian Hryvnia * `UAK` - Ukrainian Karbovanets * `AED` - United Arab Emirates Dirham * `UYW` - Uruguayan Nominal Wage Index Unit * `UYU` - Uruguayan Peso * `UYP` - Uruguayan Peso (1975–1993) * `UYI` - Uruguayan Peso (Indexed Units) * `UZS` - Uzbekistani Som * `VUV` - Vanuatu Vatu * `VES` - Venezuelan Bolívar * `VEB` - Venezuelan Bolívar (1871–2008) * `VEF` - Venezuelan Bolívar (2008–2018) * `VND` - Vietnamese Dong * `VNN` - Vietnamese Dong (1978–1985) * `CHE` - WIR Euro * `CHW` - WIR Franc * `XOF` - West African CFA Franc * `YDD` - Yemeni Dinar * `YER` - Yemeni Rial * `YUN` - Yugoslavian Convertible Dinar (1990–1992) * `YUD` - Yugoslavian Hard Dinar (1966–1990) * `YUM` - Yugoslavian New Dinar (1994–2002) * `YUR` - Yugoslavian Reformed Dinar (1992–1993) * `ZWN` - ZWN * `ZRN` - Zairean New Zaire (1993–1998) * `ZRZ` - Zairean Zaire (1971–1993) * `ZMW` - Zambian Kwacha * `ZMK` - Zambian Kwacha (1968–2012) * `ZWD` - Zimbabwean Dollar (1980–2008) * `ZWR` - Zimbabwean Dollar (2008) * `ZWL` - Zimbabwean Dollar (2009)
	Currency NullableCurrencyEnum `json:"currency,omitempty"`
	// The purchase order line item's exchange rate.
	ExchangeRate NullableFloat64 `json:"exchange_rate,omitempty"`
	// The company the purchase order line item belongs to.
	Company NullableString `json:"company,omitempty"`
	IntegrationParams map[string]interface{} `json:"integration_params,omitempty"`
	LinkedAccountParams map[string]interface{} `json:"linked_account_params,omitempty"`
	// raw json response by property name
	ResponseRaw map[string]json.RawMessage `json:"-"`
}

// NewPurchaseOrderLineItemRequest instantiates a new PurchaseOrderLineItemRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPurchaseOrderLineItemRequest(trackingCategories []string) *PurchaseOrderLineItemRequest {
	this := PurchaseOrderLineItemRequest{}
	this.TrackingCategories = trackingCategories
	return &this
}

// NewPurchaseOrderLineItemRequestWithDefaults instantiates a new PurchaseOrderLineItemRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPurchaseOrderLineItemRequestWithDefaults() *PurchaseOrderLineItemRequest {
	this := PurchaseOrderLineItemRequest{}
	return &this
}

// GetRemoteId returns the RemoteId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PurchaseOrderLineItemRequest) GetRemoteId() string {
	if o == nil || o.RemoteId.Get() == nil {
		var ret string
		return ret
	}
	return *o.RemoteId.Get()
}

// GetRemoteIdOk returns a tuple with the RemoteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PurchaseOrderLineItemRequest) GetRemoteIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RemoteId.Get(), o.RemoteId.IsSet()
}

// HasRemoteId returns a boolean if a field has been set.
func (o *PurchaseOrderLineItemRequest) HasRemoteId() bool {
	if o != nil && o.RemoteId.IsSet() {
		return true
	}

	return false
}

// SetRemoteId gets a reference to the given NullableString and assigns it to the RemoteId field.
func (o *PurchaseOrderLineItemRequest) SetRemoteId(v string) {
	o.RemoteId.Set(&v)
}
// SetRemoteIdNil sets the value for RemoteId to be an explicit nil
func (o *PurchaseOrderLineItemRequest) SetRemoteIdNil() {
	o.RemoteId.Set(nil)
}

// UnsetRemoteId ensures that no value is present for RemoteId, not even an explicit nil
func (o *PurchaseOrderLineItemRequest) UnsetRemoteId() {
	o.RemoteId.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PurchaseOrderLineItemRequest) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PurchaseOrderLineItemRequest) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *PurchaseOrderLineItemRequest) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *PurchaseOrderLineItemRequest) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *PurchaseOrderLineItemRequest) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *PurchaseOrderLineItemRequest) UnsetDescription() {
	o.Description.Unset()
}

// GetUnitPrice returns the UnitPrice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PurchaseOrderLineItemRequest) GetUnitPrice() float64 {
	if o == nil || o.UnitPrice.Get() == nil {
		var ret float64
		return ret
	}
	return *o.UnitPrice.Get()
}

// GetUnitPriceOk returns a tuple with the UnitPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PurchaseOrderLineItemRequest) GetUnitPriceOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UnitPrice.Get(), o.UnitPrice.IsSet()
}

// HasUnitPrice returns a boolean if a field has been set.
func (o *PurchaseOrderLineItemRequest) HasUnitPrice() bool {
	if o != nil && o.UnitPrice.IsSet() {
		return true
	}

	return false
}

// SetUnitPrice gets a reference to the given NullableFloat64 and assigns it to the UnitPrice field.
func (o *PurchaseOrderLineItemRequest) SetUnitPrice(v float64) {
	o.UnitPrice.Set(&v)
}
// SetUnitPriceNil sets the value for UnitPrice to be an explicit nil
func (o *PurchaseOrderLineItemRequest) SetUnitPriceNil() {
	o.UnitPrice.Set(nil)
}

// UnsetUnitPrice ensures that no value is present for UnitPrice, not even an explicit nil
func (o *PurchaseOrderLineItemRequest) UnsetUnitPrice() {
	o.UnitPrice.Unset()
}

// GetQuantity returns the Quantity field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PurchaseOrderLineItemRequest) GetQuantity() float64 {
	if o == nil || o.Quantity.Get() == nil {
		var ret float64
		return ret
	}
	return *o.Quantity.Get()
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PurchaseOrderLineItemRequest) GetQuantityOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Quantity.Get(), o.Quantity.IsSet()
}

// HasQuantity returns a boolean if a field has been set.
func (o *PurchaseOrderLineItemRequest) HasQuantity() bool {
	if o != nil && o.Quantity.IsSet() {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given NullableFloat64 and assigns it to the Quantity field.
func (o *PurchaseOrderLineItemRequest) SetQuantity(v float64) {
	o.Quantity.Set(&v)
}
// SetQuantityNil sets the value for Quantity to be an explicit nil
func (o *PurchaseOrderLineItemRequest) SetQuantityNil() {
	o.Quantity.Set(nil)
}

// UnsetQuantity ensures that no value is present for Quantity, not even an explicit nil
func (o *PurchaseOrderLineItemRequest) UnsetQuantity() {
	o.Quantity.Unset()
}

// GetItem returns the Item field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PurchaseOrderLineItemRequest) GetItem() string {
	if o == nil || o.Item.Get() == nil {
		var ret string
		return ret
	}
	return *o.Item.Get()
}

// GetItemOk returns a tuple with the Item field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PurchaseOrderLineItemRequest) GetItemOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Item.Get(), o.Item.IsSet()
}

// HasItem returns a boolean if a field has been set.
func (o *PurchaseOrderLineItemRequest) HasItem() bool {
	if o != nil && o.Item.IsSet() {
		return true
	}

	return false
}

// SetItem gets a reference to the given NullableString and assigns it to the Item field.
func (o *PurchaseOrderLineItemRequest) SetItem(v string) {
	o.Item.Set(&v)
}
// SetItemNil sets the value for Item to be an explicit nil
func (o *PurchaseOrderLineItemRequest) SetItemNil() {
	o.Item.Set(nil)
}

// UnsetItem ensures that no value is present for Item, not even an explicit nil
func (o *PurchaseOrderLineItemRequest) UnsetItem() {
	o.Item.Unset()
}

// GetAccount returns the Account field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PurchaseOrderLineItemRequest) GetAccount() string {
	if o == nil || o.Account.Get() == nil {
		var ret string
		return ret
	}
	return *o.Account.Get()
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PurchaseOrderLineItemRequest) GetAccountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Account.Get(), o.Account.IsSet()
}

// HasAccount returns a boolean if a field has been set.
func (o *PurchaseOrderLineItemRequest) HasAccount() bool {
	if o != nil && o.Account.IsSet() {
		return true
	}

	return false
}

// SetAccount gets a reference to the given NullableString and assigns it to the Account field.
func (o *PurchaseOrderLineItemRequest) SetAccount(v string) {
	o.Account.Set(&v)
}
// SetAccountNil sets the value for Account to be an explicit nil
func (o *PurchaseOrderLineItemRequest) SetAccountNil() {
	o.Account.Set(nil)
}

// UnsetAccount ensures that no value is present for Account, not even an explicit nil
func (o *PurchaseOrderLineItemRequest) UnsetAccount() {
	o.Account.Unset()
}

// GetTrackingCategory returns the TrackingCategory field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PurchaseOrderLineItemRequest) GetTrackingCategory() string {
	if o == nil || o.TrackingCategory.Get() == nil {
		var ret string
		return ret
	}
	return *o.TrackingCategory.Get()
}

// GetTrackingCategoryOk returns a tuple with the TrackingCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PurchaseOrderLineItemRequest) GetTrackingCategoryOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TrackingCategory.Get(), o.TrackingCategory.IsSet()
}

// HasTrackingCategory returns a boolean if a field has been set.
func (o *PurchaseOrderLineItemRequest) HasTrackingCategory() bool {
	if o != nil && o.TrackingCategory.IsSet() {
		return true
	}

	return false
}

// SetTrackingCategory gets a reference to the given NullableString and assigns it to the TrackingCategory field.
func (o *PurchaseOrderLineItemRequest) SetTrackingCategory(v string) {
	o.TrackingCategory.Set(&v)
}
// SetTrackingCategoryNil sets the value for TrackingCategory to be an explicit nil
func (o *PurchaseOrderLineItemRequest) SetTrackingCategoryNil() {
	o.TrackingCategory.Set(nil)
}

// UnsetTrackingCategory ensures that no value is present for TrackingCategory, not even an explicit nil
func (o *PurchaseOrderLineItemRequest) UnsetTrackingCategory() {
	o.TrackingCategory.Unset()
}

// GetTrackingCategories returns the TrackingCategories field value
func (o *PurchaseOrderLineItemRequest) GetTrackingCategories() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.TrackingCategories
}

// GetTrackingCategoriesOk returns a tuple with the TrackingCategories field value
// and a boolean to check if the value has been set.
func (o *PurchaseOrderLineItemRequest) GetTrackingCategoriesOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TrackingCategories, true
}

// SetTrackingCategories sets field value
func (o *PurchaseOrderLineItemRequest) SetTrackingCategories(v []string) {
	o.TrackingCategories = v
}

// GetTaxAmount returns the TaxAmount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PurchaseOrderLineItemRequest) GetTaxAmount() float64 {
	if o == nil || o.TaxAmount.Get() == nil {
		var ret float64
		return ret
	}
	return *o.TaxAmount.Get()
}

// GetTaxAmountOk returns a tuple with the TaxAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PurchaseOrderLineItemRequest) GetTaxAmountOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TaxAmount.Get(), o.TaxAmount.IsSet()
}

// HasTaxAmount returns a boolean if a field has been set.
func (o *PurchaseOrderLineItemRequest) HasTaxAmount() bool {
	if o != nil && o.TaxAmount.IsSet() {
		return true
	}

	return false
}

// SetTaxAmount gets a reference to the given NullableFloat64 and assigns it to the TaxAmount field.
func (o *PurchaseOrderLineItemRequest) SetTaxAmount(v float64) {
	o.TaxAmount.Set(&v)
}
// SetTaxAmountNil sets the value for TaxAmount to be an explicit nil
func (o *PurchaseOrderLineItemRequest) SetTaxAmountNil() {
	o.TaxAmount.Set(nil)
}

// UnsetTaxAmount ensures that no value is present for TaxAmount, not even an explicit nil
func (o *PurchaseOrderLineItemRequest) UnsetTaxAmount() {
	o.TaxAmount.Unset()
}

// GetTotalLineAmount returns the TotalLineAmount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PurchaseOrderLineItemRequest) GetTotalLineAmount() float64 {
	if o == nil || o.TotalLineAmount.Get() == nil {
		var ret float64
		return ret
	}
	return *o.TotalLineAmount.Get()
}

// GetTotalLineAmountOk returns a tuple with the TotalLineAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PurchaseOrderLineItemRequest) GetTotalLineAmountOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TotalLineAmount.Get(), o.TotalLineAmount.IsSet()
}

// HasTotalLineAmount returns a boolean if a field has been set.
func (o *PurchaseOrderLineItemRequest) HasTotalLineAmount() bool {
	if o != nil && o.TotalLineAmount.IsSet() {
		return true
	}

	return false
}

// SetTotalLineAmount gets a reference to the given NullableFloat64 and assigns it to the TotalLineAmount field.
func (o *PurchaseOrderLineItemRequest) SetTotalLineAmount(v float64) {
	o.TotalLineAmount.Set(&v)
}
// SetTotalLineAmountNil sets the value for TotalLineAmount to be an explicit nil
func (o *PurchaseOrderLineItemRequest) SetTotalLineAmountNil() {
	o.TotalLineAmount.Set(nil)
}

// UnsetTotalLineAmount ensures that no value is present for TotalLineAmount, not even an explicit nil
func (o *PurchaseOrderLineItemRequest) UnsetTotalLineAmount() {
	o.TotalLineAmount.Unset()
}

// GetCurrency returns the Currency field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PurchaseOrderLineItemRequest) GetCurrency() CurrencyEnum {
	if o == nil || o.Currency.Get() == nil {
		var ret CurrencyEnum
		return ret
	}
	return *o.Currency.Get()
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PurchaseOrderLineItemRequest) GetCurrencyOk() (*CurrencyEnum, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Currency.Get(), o.Currency.IsSet()
}

// HasCurrency returns a boolean if a field has been set.
func (o *PurchaseOrderLineItemRequest) HasCurrency() bool {
	if o != nil && o.Currency.IsSet() {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given NullableCurrencyEnum and assigns it to the Currency field.
func (o *PurchaseOrderLineItemRequest) SetCurrency(v CurrencyEnum) {
	o.Currency.Set(&v)
}
// SetCurrencyNil sets the value for Currency to be an explicit nil
func (o *PurchaseOrderLineItemRequest) SetCurrencyNil() {
	o.Currency.Set(nil)
}

// UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
func (o *PurchaseOrderLineItemRequest) UnsetCurrency() {
	o.Currency.Unset()
}

// GetExchangeRate returns the ExchangeRate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PurchaseOrderLineItemRequest) GetExchangeRate() float64 {
	if o == nil || o.ExchangeRate.Get() == nil {
		var ret float64
		return ret
	}
	return *o.ExchangeRate.Get()
}

// GetExchangeRateOk returns a tuple with the ExchangeRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PurchaseOrderLineItemRequest) GetExchangeRateOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ExchangeRate.Get(), o.ExchangeRate.IsSet()
}

// HasExchangeRate returns a boolean if a field has been set.
func (o *PurchaseOrderLineItemRequest) HasExchangeRate() bool {
	if o != nil && o.ExchangeRate.IsSet() {
		return true
	}

	return false
}

// SetExchangeRate gets a reference to the given NullableFloat64 and assigns it to the ExchangeRate field.
func (o *PurchaseOrderLineItemRequest) SetExchangeRate(v float64) {
	o.ExchangeRate.Set(&v)
}
// SetExchangeRateNil sets the value for ExchangeRate to be an explicit nil
func (o *PurchaseOrderLineItemRequest) SetExchangeRateNil() {
	o.ExchangeRate.Set(nil)
}

// UnsetExchangeRate ensures that no value is present for ExchangeRate, not even an explicit nil
func (o *PurchaseOrderLineItemRequest) UnsetExchangeRate() {
	o.ExchangeRate.Unset()
}

// GetCompany returns the Company field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PurchaseOrderLineItemRequest) GetCompany() string {
	if o == nil || o.Company.Get() == nil {
		var ret string
		return ret
	}
	return *o.Company.Get()
}

// GetCompanyOk returns a tuple with the Company field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PurchaseOrderLineItemRequest) GetCompanyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Company.Get(), o.Company.IsSet()
}

// HasCompany returns a boolean if a field has been set.
func (o *PurchaseOrderLineItemRequest) HasCompany() bool {
	if o != nil && o.Company.IsSet() {
		return true
	}

	return false
}

// SetCompany gets a reference to the given NullableString and assigns it to the Company field.
func (o *PurchaseOrderLineItemRequest) SetCompany(v string) {
	o.Company.Set(&v)
}
// SetCompanyNil sets the value for Company to be an explicit nil
func (o *PurchaseOrderLineItemRequest) SetCompanyNil() {
	o.Company.Set(nil)
}

// UnsetCompany ensures that no value is present for Company, not even an explicit nil
func (o *PurchaseOrderLineItemRequest) UnsetCompany() {
	o.Company.Unset()
}

// GetIntegrationParams returns the IntegrationParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PurchaseOrderLineItemRequest) GetIntegrationParams() map[string]interface{} {
	if o == nil  {
		var ret map[string]interface{}
		return ret
	}
	return o.IntegrationParams
}

// GetIntegrationParamsOk returns a tuple with the IntegrationParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PurchaseOrderLineItemRequest) GetIntegrationParamsOk() (*map[string]interface{}, bool) {
	if o == nil || o.IntegrationParams == nil {
		return nil, false
	}
	return &o.IntegrationParams, true
}

// HasIntegrationParams returns a boolean if a field has been set.
func (o *PurchaseOrderLineItemRequest) HasIntegrationParams() bool {
	if o != nil && o.IntegrationParams != nil {
		return true
	}

	return false
}

// SetIntegrationParams gets a reference to the given map[string]interface{} and assigns it to the IntegrationParams field.
func (o *PurchaseOrderLineItemRequest) SetIntegrationParams(v map[string]interface{}) {
	o.IntegrationParams = v
}

// GetLinkedAccountParams returns the LinkedAccountParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PurchaseOrderLineItemRequest) GetLinkedAccountParams() map[string]interface{} {
	if o == nil  {
		var ret map[string]interface{}
		return ret
	}
	return o.LinkedAccountParams
}

// GetLinkedAccountParamsOk returns a tuple with the LinkedAccountParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PurchaseOrderLineItemRequest) GetLinkedAccountParamsOk() (*map[string]interface{}, bool) {
	if o == nil || o.LinkedAccountParams == nil {
		return nil, false
	}
	return &o.LinkedAccountParams, true
}

// HasLinkedAccountParams returns a boolean if a field has been set.
func (o *PurchaseOrderLineItemRequest) HasLinkedAccountParams() bool {
	if o != nil && o.LinkedAccountParams != nil {
		return true
	}

	return false
}

// SetLinkedAccountParams gets a reference to the given map[string]interface{} and assigns it to the LinkedAccountParams field.
func (o *PurchaseOrderLineItemRequest) SetLinkedAccountParams(v map[string]interface{}) {
	o.LinkedAccountParams = v
}

func (o PurchaseOrderLineItemRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RemoteId.IsSet() {
		toSerialize["remote_id"] = o.RemoteId.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.UnitPrice.IsSet() {
		toSerialize["unit_price"] = o.UnitPrice.Get()
	}
	if o.Quantity.IsSet() {
		toSerialize["quantity"] = o.Quantity.Get()
	}
	if o.Item.IsSet() {
		toSerialize["item"] = o.Item.Get()
	}
	if o.Account.IsSet() {
		toSerialize["account"] = o.Account.Get()
	}
	if o.TrackingCategory.IsSet() {
		toSerialize["tracking_category"] = o.TrackingCategory.Get()
	}
	if true {
		toSerialize["tracking_categories"] = o.TrackingCategories
	}
	if o.TaxAmount.IsSet() {
		toSerialize["tax_amount"] = o.TaxAmount.Get()
	}
	if o.TotalLineAmount.IsSet() {
		toSerialize["total_line_amount"] = o.TotalLineAmount.Get()
	}
	if o.Currency.IsSet() {
		toSerialize["currency"] = o.Currency.Get()
	}
	if o.ExchangeRate.IsSet() {
		toSerialize["exchange_rate"] = o.ExchangeRate.Get()
	}
	if o.Company.IsSet() {
		toSerialize["company"] = o.Company.Get()
	}
	if o.IntegrationParams != nil {
		toSerialize["integration_params"] = o.IntegrationParams
	}
	if o.LinkedAccountParams != nil {
		toSerialize["linked_account_params"] = o.LinkedAccountParams
	}
	return json.Marshal(toSerialize)
}

func (v *PurchaseOrderLineItemRequest) UnmarshalJSON(src []byte) error {
    type PurchaseOrderLineItemRequestUnmarshalTarget PurchaseOrderLineItemRequest

	var intermediateResult PurchaseOrderLineItemRequestUnmarshalTarget
	var err1 = json.Unmarshal(src, &intermediateResult)
    if err1 != nil {
        return err1
    }
    var err2 = json.Unmarshal(src, &intermediateResult.ResponseRaw)
	if err2 != nil {
		return err2
	}

	*v = PurchaseOrderLineItemRequest(intermediateResult)
	return nil
}
type NullablePurchaseOrderLineItemRequest struct {
	value *PurchaseOrderLineItemRequest
	isSet bool
}

func (v NullablePurchaseOrderLineItemRequest) Get() *PurchaseOrderLineItemRequest {
	return v.value
}

func (v *NullablePurchaseOrderLineItemRequest) Set(val *PurchaseOrderLineItemRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePurchaseOrderLineItemRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePurchaseOrderLineItemRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePurchaseOrderLineItemRequest(val *PurchaseOrderLineItemRequest) *NullablePurchaseOrderLineItemRequest {
	return &NullablePurchaseOrderLineItemRequest{value: val, isSet: true}
}

func (v NullablePurchaseOrderLineItemRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePurchaseOrderLineItemRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	var err1 = json.Unmarshal(src, &v.value)
    if err1 != nil {
        return err1
    }
    return json.Unmarshal(src, &v.value.ResponseRaw)
}


