package interfaces

import "encoding/json"

func UnmarshalSiteInfo(data []byte) (SiteInfo, error) {
	var r SiteInfo
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SiteInfo) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type SiteInfo struct {
	Categories         []Category `json:"categories"`
	CountryID          *string    `json:"country_id,omitempty"`
	Currencies         []Currency `json:"currencies"`
	DefaultCurrencyID  *string    `json:"default_currency_id,omitempty"`
	ID                 *string    `json:"id,omitempty"`
	ImmediatePayment   *string    `json:"immediate_payment,omitempty"`
	MercadopagoVersion *int64     `json:"mercadopago_version,omitempty"`
	Name               *string    `json:"name,omitempty"`
	PaymentMethodIDS   []string   `json:"payment_method_ids"`
	SaleFeesMode       *string    `json:"sale_fees_mode,omitempty"`
	Settings           *Settings  `json:"settings,omitempty"`
}

type Category struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type Currency struct {
	ID     *string `json:"id,omitempty"`
	Symbol *string `json:"symbol,omitempty"`
}

type Settings struct {
	IdentificationTypes      []string                  `json:"identification_types"`
	IdentificationTypesRules []IdentificationTypesRule `json:"identification_types_rules"`
	TaxpayerTypes            []string                  `json:"taxpayer_types"`
}

type IdentificationTypesRule struct {
	IdentificationType *string `json:"identification_type,omitempty"`
	Rules              []Rule  `json:"rules"`
}

type Rule struct {
	BeginsWith           *string  `json:"begins_with,omitempty"`
	EnabledTaxpayerTypes []string `json:"enabled_taxpayer_types"`
	MaxLength            *int64   `json:"max_length,omitempty"`
	MinLength            *int64   `json:"min_length,omitempty"`
	Type                 *string  `json:"type,omitempty"`
}
