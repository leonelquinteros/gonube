package gonube

import (
	"fmt"
)

// ErrorResponse represents the error response format from Mailchimp's API
type ErrorResponse struct {
	Code         int      `json:"code"`
	Message      string   `json:"message"`
	Description  string   `json:"description"`
	ErrorMessage string   `json:"error"`
	Src          []string `json:"src"`
}

// Error implementation
func (er ErrorResponse) Error() string {
	return fmt.Sprintf("Error: %s - %d - %s: %s - %+v", er.ErrorMessage, er.Code, er.Message, er.Description, er.Src)
}
