package gonube

// Customer data
type Customer struct {
	CreatedAt          string  `json:"created_at"`
	Email              string  `json:"email"`
	ID                 int     `json:"id"`
	LastOrderID        int     `json:"last_order_id"`
	Name               string  `json:"name"`
	Phone              string  `json:"phone"`
	TotalSpent         string  `json:"total_spent"`
	TotalSpentCurrency string  `json:"total_spent_currency"`
	UpdatedAt          string  `json:"updated_at"`
	DefaultAddress     Address `json:"default_address"`
}
