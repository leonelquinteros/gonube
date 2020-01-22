package gonube

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"path"
)

const (
	clientName = "Gonube SDK (https://github.com/leonelquinteros/gonube)"
	apiHost    = "https://api.tiendanube.com/v1"
)

// ClientConfig object used for client creation
type ClientConfig struct {
	AccessToken  string
	UserID       string
	ClientID     string
	ClientSecret string
	Debug        bool
}

// NewClientConfig constructs a ClientConfig object with the environment variables set as default
func NewClientConfig() ClientConfig {
	return ClientConfig{
		AccessToken:  accessToken,
		UserID:       userID,
		ClientID:     clientID,
		ClientSecret: clientSecret,
	}
}

// Client object
type Client struct {
	config ClientConfig

	Transport http.RoundTripper
}

// New constructor from configuration
func New(cc ClientConfig) Client {
	return Client{
		config: cc,
	}
}

// Request executes any Tiendanube API method using the current client configuration
func (c Client) Request(method, endpoint string, params url.Values, data, response interface{}) error {
	// Parse URL
	base, err := url.Parse(apiHost)
	if err != nil {
		return err
	}
	base.Path = path.Join(base.Path, c.config.UserID, endpoint)
	// Handle root path redirect
	if endpoint == "" || endpoint == "/" {
		base.Path += "/"
	}

	// Parse params
	if params != nil {
		for k := range params {
			base.Query().Set(k, params.Get(k))
		}
	}

	// Parse data
	var eData []byte
	if data != nil {
		eData, err = json.Marshal(data)
		if err != nil {
			return err
		}
	}

	// Create request
	req, err := http.NewRequest(method, base.String(), bytes.NewBuffer(eData))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	// Set Auth
	log.Println(c.config.AccessToken)
	req.Header.Set("User-Agent", clientName)
	req.Header.Set("Authentication", "bearer "+c.config.AccessToken)

	// Debug
	if c.config.Debug {
		log.Printf("NEW REQUEST TO %s", base.String())
	}

	// Perform request
	client := &http.Client{Transport: c.Transport}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Read response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// Debug
	if c.config.Debug {
		log.Printf("RESPONSE FROM %s: \n%s", base.String(), body)
	}

	// Handle API errors
	if resp.StatusCode >= 400 {
		errResp := ErrorResponse{}
		err = json.Unmarshal(body, &errResp)
		if err != nil {
			return err
		}

		err = errResp
	} else {
		// Unmarshal into response
		if len(body) > 0 {
			err = json.Unmarshal(body, response)
		}
	}

	return err
}

// Orders returns a Orders API client
func (c Client) Orders() Orders {
	return Orders{
		Client: c,
	}
}
