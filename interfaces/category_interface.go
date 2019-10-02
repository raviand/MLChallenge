package interfaces

import "encoding/json"

func UnmarshalCategoryData(data []byte) (CategoryData, error) {
	var r CategoryData
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CategoryData) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type CategoryData struct {
	ID                       *string            `json:"id,omitempty"`
	Name                     *string            `json:"name,omitempty"`
	Picture                  *string            `json:"picture,omitempty"`
	Permalink                *string            `json:"permalink,omitempty"`
	TotalItemsInThisCategory *int64             `json:"total_items_in_this_category,omitempty"`
	PathFromRoot             []PathFromRoot     `json:"path_from_root"`
	ChildrenCategories       []ChildrenCategory `json:"children_categories"`
	AttributeTypes           *string            `json:"attribute_types,omitempty"`
	Settings                 *CategorySettings  `json:"settings,omitempty"`
	MetaCategID              interface{}        `json:"meta_categ_id"`
	Attributable             *bool              `json:"attributable,omitempty"`
}

type ChildrenCategory struct {
	ID                       *string `json:"id,omitempty"`
	Name                     *string `json:"name,omitempty"`
	TotalItemsInThisCategory *int64  `json:"total_items_in_this_category,omitempty"`
}

type PathFromRoot struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type CategorySettings struct {
	AdultContent            *bool         `json:"adult_content,omitempty"`
	BuyingAllowed           *bool         `json:"buying_allowed,omitempty"`
	BuyingModes             []string      `json:"buying_modes"`
	CatalogDomain           *string       `json:"catalog_domain,omitempty"`
	CoverageAreas           *string       `json:"coverage_areas,omitempty"`
	Currencies              []string      `json:"currencies"`
	Fragile                 *bool         `json:"fragile,omitempty"`
	ImmediatePayment        *string       `json:"immediate_payment,omitempty"`
	ItemConditions          []string      `json:"item_conditions"`
	ItemsReviewsAllowed     *bool         `json:"items_reviews_allowed,omitempty"`
	ListingAllowed          *bool         `json:"listing_allowed,omitempty"`
	MaxDescriptionLength    *int64        `json:"max_description_length,omitempty"`
	MaxPicturesPerItem      *int64        `json:"max_pictures_per_item,omitempty"`
	MaxPicturesPerItemVar   *int64        `json:"max_pictures_per_item_var,omitempty"`
	MaxSubTitleLength       *int64        `json:"max_sub_title_length,omitempty"`
	MaxTitleLength          *int64        `json:"max_title_length,omitempty"`
	MaximumPrice            interface{}   `json:"maximum_price"`
	MinimumPrice            *int64        `json:"minimum_price,omitempty"`
	MirrorCategory          interface{}   `json:"mirror_category"`
	MirrorMasterCategory    interface{}   `json:"mirror_master_category"`
	MirrorSlaveCategories   []interface{} `json:"mirror_slave_categories"`
	Price                   *string       `json:"price,omitempty"`
	ReservationAllowed      *string       `json:"reservation_allowed,omitempty"`
	Restrictions            []interface{} `json:"restrictions"`
	RoundedAddress          *bool         `json:"rounded_address,omitempty"`
	SellerContact           *string       `json:"seller_contact,omitempty"`
	ShippingModes           []string      `json:"shipping_modes"`
	ShippingOptions         []string      `json:"shipping_options"`
	ShippingProfile         *string       `json:"shipping_profile,omitempty"`
	ShowContactInformation  *bool         `json:"show_contact_information,omitempty"`
	SimpleShipping          *string       `json:"simple_shipping,omitempty"`
	Stock                   *string       `json:"stock,omitempty"`
	SubVertical             *string       `json:"sub_vertical,omitempty"`
	Subscribable            *bool         `json:"subscribable,omitempty"`
	Tags                    []interface{} `json:"tags"`
	Vertical                *string       `json:"vertical,omitempty"`
	VipSubdomain            *string       `json:"vip_subdomain,omitempty"`
	BuyerProtectionPrograms []string      `json:"buyer_protection_programs"`
}
