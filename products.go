package gonube

import (
	"net/url"
	"strconv"
)

// Product data
type Product struct {
	Attributes     []map[string]string `json:"attributes"`
	Categories     []Category          `json:"categories"`
	CreatedAt      string              `json:"created_at"`
	Description    map[string]string   `json:"description"`
	Handle         map[string]string   `json:"handle"`
	ID             int                 `json:"id"`
	Images         []ProductImage      `json:"images"`
	Name           map[string]string   `json:"name"`
	Brand          interface{}         `json:"brand"`
	SeoTitle       map[string]string   `json:"seo_title"`
	SeoDescription map[string]string   `json:"seo_description"`
	Published      bool                `json:"published"`
	FreeShipping   bool                `json:"free_shipping"`
	UpdatedAt      string              `json:"updated_at"`
	Variants       []ProductVariant    `json:"variants"`
}

// ProductImage data
type ProductImage struct {
	ID        int    `json:"id"`
	Src       string `json:"src"`
	Position  int    `json:"position"`
	ProductID int    `json:"product_id"`
}

// ProductVariant data
type ProductVariant struct {
	ID               int                 `json:"id"`
	PromotionalPrice string              `json:"promotional_price"`
	CreatedAt        string              `json:"created_at"`
	Depth            interface{}         `json:"depth"`
	Height           interface{}         `json:"height"`
	Values           []map[string]string `json:"values"`
	Price            string              `json:"price"`
	ProductID        int                 `json:"product_id"`
	StockManagement  bool                `json:"stock_management"`
	Stock            int                 `json:"stock"`
	Sku              string              `json:"sku"`
	UpdatedAt        string              `json:"updated_at"`
	Weight           string              `json:"weight"`
	Width            interface{}         `json:"width"`
}

// Products API client
type Products struct {
	Client
}

// All retrieves all Lists
func (c Products) All(listParams url.Values) ([]Product, error) {
	r := make([]Product, 0)
	err := c.Client.Request("GET", "/products", listParams, nil, &r)
	return r, err
}

// Get product by ID
func (c Products) Get(id int, listParams url.Values) (Product, error) {
	var p Product
	err := c.Client.Request("GET", "products/"+strconv.Itoa(id), listParams, nil, &p)
	return p, err
}
