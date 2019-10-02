package interfaces

import "encoding/json"

func UnmarshalSellerInfo(data []byte) (SellerInfo, error) {
	var r SellerInfo
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SellerInfo) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type SellerInfo struct {
	ID               *int64            `json:"id,omitempty"`
	Nickname         *string           `json:"nickname,omitempty"`
	RegistrationDate *string           `json:"registration_date,omitempty"`
	CountryID        *string           `json:"country_id,omitempty"`
	Address          *Address          `json:"address,omitempty"`
	UserType         *string           `json:"user_type,omitempty"`
	Tags             []string          `json:"tags"`
	Logo             interface{}       `json:"logo"`
	Points           *int64            `json:"points,omitempty"`
	SiteID           *string           `json:"site_id,omitempty"`
	Permalink        *string           `json:"permalink,omitempty"`
	SellerReputation *SellerReputation `json:"seller_reputation,omitempty"`
	BuyerReputation  *BuyerReputation  `json:"buyer_reputation,omitempty"`
	Status           *Status           `json:"status,omitempty"`
}

type Address struct {
	City  *string `json:"city,omitempty"`
	State *string `json:"state,omitempty"`
}

type BuyerReputation struct {
	Tags []interface{} `json:"tags"`
}

type SellerReputation struct {
	LevelID           *string       `json:"level_id,omitempty"`
	PowerSellerStatus interface{}   `json:"power_seller_status"`
	Transactions      *Transactions `json:"transactions,omitempty"`
}

type Transactions struct {
	Canceled  *int64   `json:"canceled,omitempty"`
	Completed *int64   `json:"completed,omitempty"`
	Period    *string  `json:"period,omitempty"`
	Ratings   *Ratings `json:"ratings,omitempty"`
	Total     *int64   `json:"total,omitempty"`
}

type Ratings struct {
	Negative *int64 `json:"negative,omitempty"`
	Neutral  *int64 `json:"neutral,omitempty"`
	Positive *int64 `json:"positive,omitempty"`
}

type Status struct {
	SiteStatus *string `json:"site_status,omitempty"`
}
