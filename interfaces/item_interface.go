package interfaces

import "encoding/json"

func UnmarshalItemInterface(data []byte) (ItemInterface, error) {
	var r ItemInterface
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ItemInterface) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type ItemInterface struct {
	AcceptsMercadopago           bool        `json:"accepts_mercadopago"`
	Attributes                   []Attribute `json:"attributes"`
	AutomaticRelist              bool        `json:"automatic_relist"`
	AvailableQuantity            int64       `json:"available_quantity"`
	BasePrice                    int64       `json:"base_price"`
	BuyingMode                   string      `json:"buying_mode"`
	CatalogListing               bool        `json:"catalog_listing"`
	CatalogProductID             interface{} `json:"catalog_product_id"`
	CategoryID                   string      `json:"category_id"`
	CategoryData                 *CategoryData
	Condition                    string        `json:"condition"`
	CoverageAreas                []interface{} `json:"coverage_areas"`
	CurrencyID                   string        `json:"currency_id"`
	DateCreated                  string        `json:"date_created"`
	DealIDS                      []interface{} `json:"deal_ids"`
	Descriptions                 []Description `json:"descriptions"`
	DifferentialPricing          interface{}   `json:"differential_pricing"`
	DomainID                     string        `json:"domain_id"`
	Geolocation                  Geolocation   `json:"geolocation"`
	Health                       float64       `json:"health"`
	ID                           string        `json:"id"`
	InitialQuantity              int64         `json:"initial_quantity"`
	InternationalDeliveryMode    string        `json:"international_delivery_mode"`
	LastUpdated                  string        `json:"last_updated"`
	ListingSource                string        `json:"listing_source"`
	ListingTypeID                string        `json:"listing_type_id"`
	Location                     Location      `json:"location"`
	NonMercadoPagoPaymentMethods []interface{} `json:"non_mercado_pago_payment_methods"`
	OfficialStoreID              interface{}   `json:"official_store_id"`
	OriginalPrice                interface{}   `json:"original_price"`
	ParentItemID                 interface{}   `json:"parent_item_id"`
	Permalink                    string        `json:"permalink"`
	Pictures                     []Picture     `json:"pictures"`
	Price                        int64         `json:"price"`
	SaleTerms                    []interface{} `json:"sale_terms"`
	SecureThumbnail              string        `json:"secure_thumbnail"`
	SellerAddress                SellerAddress `json:"seller_address"`
	SellerContact                interface{}   `json:"seller_contact"`
	SellerID                     int64         `json:"seller_id"`
	SellerData                   *SellerInfo
	Shipping                     Shipping `json:"shipping"`
	SiteID                       string   `json:"site_id"`
	SiteData                     *SiteInfo
	SoldQuantity                 int64         `json:"sold_quantity"`
	StartTime                    string        `json:"start_time"`
	Status                       string        `json:"status"`
	StopTime                     string        `json:"stop_time"`
	SubStatus                    []string      `json:"sub_status"`
	Subtitle                     interface{}   `json:"subtitle"`
	Tags                         []string      `json:"tags"`
	Thumbnail                    string        `json:"thumbnail"`
	Title                        string        `json:"title"`
	Variations                   []interface{} `json:"variations"`
	VideoID                      interface{}   `json:"video_id"`
	Warnings                     []interface{} `json:"warnings"`
	Warranty                     interface{}   `json:"warranty"`
}

type Attribute struct {
	AttributeGroupID   string      `json:"attribute_group_id"`
	AttributeGroupName string      `json:"attribute_group_name"`
	ID                 string      `json:"id"`
	Name               string      `json:"name"`
	ValueID            string      `json:"value_id"`
	ValueName          string      `json:"value_name"`
	ValueStruct        interface{} `json:"value_struct"`
}

type Description struct {
	ID string `json:"id"`
}

type Geolocation struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type Location struct {
}

type Picture struct {
	ID        string `json:"id"`
	MaxSize   string `json:"max_size"`
	Quality   string `json:"quality"`
	SecureURL string `json:"secure_url"`
	Size      string `json:"size"`
	URL       string `json:"url"`
}

type SellerAddress struct {
	City           City           `json:"city"`
	Country        City           `json:"country"`
	ID             int64          `json:"id"`
	Latitude       float64        `json:"latitude"`
	Longitude      float64        `json:"longitude"`
	SearchLocation SearchLocation `json:"search_location"`
	State          City           `json:"state"`
}

type City struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type SearchLocation struct {
	City  City `json:"city"`
	State City `json:"state"`
}

type Shipping struct {
	Dimensions   interface{}   `json:"dimensions"`
	FreeShipping bool          `json:"free_shipping"`
	LocalPickUp  bool          `json:"local_pick_up"`
	LogisticType string        `json:"logistic_type"`
	Methods      []interface{} `json:"methods"`
	Mode         string        `json:"mode"`
	StorePickUp  bool          `json:"store_pick_up"`
	Tags         []interface{} `json:"tags"`
}
