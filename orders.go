package gonube

import "net/url"

// Order data
type Order struct {
	CancelReason            interface{}    `json:"cancel_reason"`
	CreatedAt               string         `json:"created_at"`
	Currency                string         `json:"currency"`
	Gateway                 string         `json:"gateway"`
	ID                      int            `json:"id"`
	LandingSite             string         `json:"landing_site"`
	Language                string         `json:"language"`
	LocationID              interface{}    `json:"location_id"`
	Name                    string         `json:"name"`
	Note                    interface{}    `json:"note"`
	Number                  int            `json:"number"`
	OwnerNote               interface{}    `json:"owner_note"`
	PaymentStatus           string         `json:"payment_status"`
	Shipping                string         `json:"shipping"`
	ShippingStatus          string         `json:"shipping_status"`
	ShippingTrackingNumber  interface{}    `json:"shipping_tracking_number"`
	ShippingTrackingURL     interface{}    `json:"shipping_tracking_url"`
	ShippingMinDays         int            `json:"shipping_min_days"`
	ShippingMaxDays         int            `json:"shipping_max_days"`
	ShippingCostOwner       string         `json:"shipping_cost_owner"`
	ShippingCostCustomer    string         `json:"shipping_cost_customer"`
	ShippingOption          string         `json:"shipping_option"`
	ShippingOptionCode      string         `json:"shipping_option_code"`
	ShippingOptionReference interface{}    `json:"shipping_option_reference"`
	Status                  string         `json:"status"`
	Subtotal                string         `json:"subtotal"`
	Token                   string         `json:"token"`
	Discount                string         `json:"discount"`
	Price                   string         `json:"price"`
	PriceUsd                string         `json:"price_usd"`
	Weight                  string         `json:"weight"`
	UpdatedAt               string         `json:"updated_at"`
	ShippedAt               string         `json:"shipped_at"`
	Coupon                  []OrderCoupon  `json:"coupon"`
	Products                []OrderProduct `json:"products"`
	BillingAddress          string         `json:"billing_address"`
	BillingCity             string         `json:"billing_city"`
	BillingCountry          string         `json:"billing_country"`
	BillingDefault          bool           `json:"billing_default"`
	BillingFloor            interface{}    `json:"billing_floor"`
	BillingLocality         interface{}    `json:"billing_locality"`
	BillingNumber           string         `json:"billing_number"`
	BillingPhone            string         `json:"billing_phone"`
	BillingProvince         string         `json:"billing_province"`
	BillingZipcode          string         `json:"billing_zipcode"`
	Extra                   OrderExtra     `json:"extra"`
	ShippingPickupType      string         `json:"shipping_pickup_type"`
	ShippingStoreBranchName interface{}    `json:"shipping_store_branch_name"`
	ShippingAddress         Address        `json:"shipping_address"`
	Customer                Customer       `json:"customer"`
}

// OrderProduct data
type OrderProduct struct {
	Depth        interface{} `json:"depth"`
	Height       interface{} `json:"height"`
	Price        string      `json:"price"`
	ProductID    int         `json:"product_id"`
	Quantity     string      `json:"quantity"`
	FreeShipping bool        `json:"free_shipping"`
	VariantID    string      `json:"variant_id"`
	Weight       string      `json:"weight"`
	Width        interface{} `json:"width"`
}

// OrderCoupon data
type OrderCoupon struct {
	Code string `json:"code"`
}

// OrderExtra data
type OrderExtra struct {
	GiftWrap string `json:"gift-wrap"`
}

// Orders API client
type Orders struct {
	Client
}

// All retrieves all Lists
func (c Orders) All(listParams url.Values) ([]Order, error) {
	r := make([]Order, 0)
	err := c.Client.Request("GET", "/orders", listParams, nil, &r)
	return r, err
}
