package gonube

// Address data
type Address struct {
	Address   string      `json:"address"`
	City      string      `json:"city"`
	Country   string      `json:"country"`
	CreatedAt string      `json:"created_at"`
	Default   bool        `json:"default"`
	Floor     interface{} `json:"floor"`
	ID        int         `json:"id"`
	Locality  interface{} `json:"locality"`
	Number    string      `json:"number"`
	Phone     string      `json:"phone"`
	Province  string      `json:"province"`
	UpdatedAt string      `json:"updated_at"`
	Zipcode   string      `json:"zipcode"`
}
