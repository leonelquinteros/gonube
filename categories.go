package gonube

// Category data
type Category struct {
	CreatedAt     string            `json:"created_at"`
	Description   map[string]string `json:"description"`
	Handle        map[string]string `json:"handle"`
	ID            int               `json:"id"`
	Name          map[string]string `json:"name"`
	Parent        interface{}       `json:"parent"`
	Subcategories []interface{}     `json:"subcategories"`
	UpdatedAt     string            `json:"updated_at"`
}
