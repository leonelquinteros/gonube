package gonube

import "os"

var (
	accessToken  string
	userID       string
	clientID     string
	clientSecret string
)

func init() {
	// Set environment variables configuration.
	if os.Getenv("TIENDANUBE_ACCESS_TOKEN") != "" {
		accessToken = os.Getenv("TIENDANUBE_ACCESS_TOKEN")
	}
	if os.Getenv("TIENDANUBE_USER_ID") != "" {
		userID = os.Getenv("TIENDANUBE_USER_ID")
	}
	if os.Getenv("TIENDANUBE_CLIENT_ID") != "" {
		clientID = os.Getenv("TIENDANUBE_CLIENT_ID")
	}
	if os.Getenv("TIENDANUBE_CLIENT_SECRET") != "" {
		clientSecret = os.Getenv("TIENDANUBE_CLIENT_SECRET")
	}
}
