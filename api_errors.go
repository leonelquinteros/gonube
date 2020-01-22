package gonube

import (
	"fmt"
)

// ErrorResponse represents the error response format from Mailchimp's API
type ErrorResponse struct {
	ErrorMessage string   `json:"error"`
	Src          []string `json:"src"`
}

// Error implementation
func (er ErrorResponse) Error() string {
	return fmt.Sprintf("Error: %s - %+v", er.ErrorMessage, er.Src)
}
