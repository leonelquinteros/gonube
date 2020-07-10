package gonube

// Store data
type Store struct {
	Address          interface{} `json:"address"`
	AdminLanguage    string      `json:"admin_language"`
	Blog             interface{} `json:"blog"`
	BusinessID       interface{} `json:"business_id"`
	BusinessName     interface{} `json:"business_name"`
	BusinessAddress  interface{} `json:"business_address"`
	ContactEmail     string      `json:"contact_email"`
	Country          string      `json:"country"`
	CreatedAt        string      `json:"created_at"`
	CustomerAccounts string      `json:"customer_accounts"`
	Description      struct {
		En string `json:"en"`
		Es string `json:"es"`
		Pt string `json:"pt"`
	} `json:"description"`
	Domains    []string `json:"domains"`
	Email      string   `json:"email"`
	Facebook   string   `json:"facebook"`
	GooglePlus string   `json:"google_plus"`
	ID         int      `json:"id"`
	Instagram  string   `json:"instagram"`
	Languages  struct {
		En struct {
			Currency string `json:"currency"`
			Active   bool   `json:"active"`
		} `json:"en"`
		Es struct {
			Currency string `json:"currency"`
			Active   bool   `json:"active"`
		} `json:"es"`
		Pt struct {
			Currency string `json:"currency"`
			Active   bool   `json:"active"`
		} `json:"pt"`
	} `json:"languages"`
	Logo         string `json:"logo"`
	MainCurrency string `json:"main_currency"`
	CurrentTheme string `json:"current_theme"`
	MainLanguage string `json:"main_language"`
	Name         struct {
		En string `json:"en"`
		Es string `json:"es"`
		Pt string `json:"pt"`
	} `json:"name"`
	OriginalDomain string      `json:"original_domain"`
	Phone          interface{} `json:"phone"`
	Pinterest      string      `json:"pinterest"`
	PlanName       string      `json:"plan_name"`
	Type           interface{} `json:"type"`
	Twitter        string      `json:"twitter"`
}

// Stores API client
type Stores struct {
	Client
}

// Get store
func (c Stores) Get() (Store, error) {
	var s Store
	err := c.Client.Request("GET", "store", nil, nil, &s)
	return s, err
}
