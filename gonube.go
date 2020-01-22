package gonube

import "os"

var (
	clientID     string
	clientSecret string
)

func init() {
	// Set environment variables configuration.
	if os.Getenv("TIENDANUBE_CLIENT_ID") != "" {
		clientID = os.Getenv("TIENDANUBE_CLIENT_ID")
	}
	if os.Getenv("TIENDANUBE_CLIENT_SECRET") != "" {
		clientSecret = os.Getenv("TIENDANUBE_CLIENT_SECRET")
	}
}
