package taxjar

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
)

var (
	// ErrEmptyAPIKey is returned when a Client is initialized with an empty API key
	ErrEmptyAPIKey = errors.New("invalid API key, cannot be empty")
)

/*
curl https://api.goshippo.com/shipments/  \
    -H "Authorization: ShippoToken <API_TOKEN>" \
    -H "Content-Type: application/json"  \
    -d '{...}'

*/

const (
	host        = "https://api.taxjar.com"
	sandboxHost = "https://api.sandbox.taxjar.com"

	apiVersion = "v2"

	endpointTaxes = "/taxes"
)

var (
	ErrNotFound = errors.New("page not found")
)

// New initializes and returns a new TaxJar Client
func New(apiKey string) (client *Client, err error) {
	return newClient(host, apiKey)
}

// NewSandbox initializes and returns a new TarJar sandbox Client
func NewSandbox(apiKey string) (client *Client, err error) {
	return newClient(sandboxHost, apiKey)
}

// New initializes and returns a new Stripe Client
func newClient(host, apiKey string) (client *Client, err error) {
	if len(apiKey) == 0 {
		err = ErrEmptyAPIKey
		return
	}

	var c Client
	if c.u, err = url.Parse(host); err != nil {
		return
	}

	c.apiKey = apiKey
	client = &c
	return
}

type Client struct {
	hc http.Client
	u  *url.URL

	apiKey string
}

func (c *Client) GetTaxes(req TaxesRequest) (resp *TaxesResponse, err error) {
	var r taxesResponse
	u := c.getURL(endpointTaxes, nil)
	if err = c.request("POST", u, req, &r); err != nil {
		return
	}

	resp = &r.Tax
	return
}

func (c *Client) request(method, url string, request, response interface{}) (err error) {
	var body io.Reader
	if body, err = getRequestBody(request); err != nil {
		err = fmt.Errorf("error getting request body: %v", err)
		return
	}

	var req *http.Request
	if req, err = http.NewRequest(method, url, body); err != nil {
		err = fmt.Errorf("error creating request: %v", err)
		return
	}

	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	req.Header.Set("Content-Type", "application/json")

	var resp *http.Response
	if resp, err = c.hc.Do(req); err != nil {
		err = fmt.Errorf("error performing request: %v", err)
		return
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case 200, 201:
		return handleResponse(resp.Body, response)
	case 400:
		return handleError(resp.Body)
	case 404:
		return ErrNotFound
	default:
		return fmt.Errorf("Unexpected status code of: %d (url: <%s>, method: <%s>)", resp.StatusCode, url, method)
	}
}

func (c *Client) getURL(endpoint string, q url.Values) string {
	u := *c.u
	u.Path = path.Join(apiVersion, endpoint)

	if q != nil {
		u.RawQuery = q.Encode()
	}

	return u.String()
}
