package gonube

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"path"
	"strconv"
)

const (
	clientName = "Gonube SDK (https://github.com/leonelquinteros/gonube)"
	apiHost    = "https://api.tiendanube.com/v1"
)

// AuthResponse data
type AuthResponse struct {
	AccessToken      string `json:"access_token"`
	TokenType        string `json:"token_type"`
	Scope            string `json:"scope"`
	UserID           uint   `json:"user_id"`
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

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

// GetAccessToken exchanges an authorization code obtained through OAuth2 for a permanent access_token.
func (c Client) GetAccessToken(code string) (AuthResponse, error) {
	var r AuthResponse
	params := url.Values{}
	params.Set("client_id", c.config.ClientID)
	params.Set("client_secret", c.config.ClientSecret)
	params.Set("grant_type", "authorization_code")
	params.Set("code", code)
	data := params.Encode()

	if c.config.Debug {
		log.Printf("Sending auth request to https://www.tiendanube.com/apps/authorize/token?grant_type=authorization_code with payload: %s", data)
	}

	authRequest, err := http.NewRequest("POST", "https://www.tiendanube.com/apps/authorize/token?grant_type=authorization_code", bytes.NewBufferString(data))
	if err != nil {
		return r, err
	}
	authRequest.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Perform request
	client := &http.Client{Transport: c.Transport}
	resp, err := client.Do(authRequest)
	if err != nil {
		return r, err
	}
	defer resp.Body.Close()

	// Read response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return r, err
	}
	if c.config.Debug {
		log.Printf("Got auth response: %s", body)
	}
	err = json.Unmarshal(body, &r)
	if err != nil {
		return r, err
	}

	// Validate response
	if r.Error != "" {
		return r, fmt.Errorf("%s: %s", r.Error, r.ErrorDescription)
	}

	// Set current client credentials
	c.config.AccessToken = r.AccessToken
	c.config.UserID = strconv.FormatUint(uint64(r.UserID), 10)

	return r, nil
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
	uri := base.String()

	// Parse params
	if params == nil {
		params = url.Values{}
	}
	encodedParams := params.Encode()
	if encodedParams != "" {
		uri += "?" + params.Encode()
	}

	// Debug
	if c.config.Debug {
		log.Printf("NEW REQUEST TO %s with payload: %+v", base.String(), data)
	}

	// Create request
	var req *http.Request
	if data != nil {
		var eData []byte
		eData, err = json.Marshal(data)
		if err != nil {
			return err
		}
		req, err = http.NewRequest(method, uri, bytes.NewBuffer(eData))
		if err != nil {
			return err
		}
	} else {
		req, err = http.NewRequest(method, uri, nil)
		if err != nil {
			return err
		}
	}

	req.Header.Set("Content-Type", "application/json")

	// Set Auth
	req.Header.Set("User-Agent", clientName)
	req.Header.Set("Authentication", "bearer "+c.config.AccessToken)

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
		log.Printf("RESPONSE FROM %s: %s", base.String(), body)
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

// Stores returns a Stores API client
func (c Client) Stores() Stores {
	return Stores{
		Client: c,
	}
}

// Products returns a Products API client
func (c Client) Products() Products {
	return Products{
		Client: c,
	}
}

// Orders returns a Orders API client
func (c Client) Orders() Orders {
	return Orders{
		Client: c,
	}
}
