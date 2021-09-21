package taxjar

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

var (
	ErrEmptyAPIKey = errors.New("invalid API key, cannot be empty")
)

const (
	host = "https://api.sandbox.taxjar.com/v2/taxes"
)

func New(apiKey string) (client *Client, err error) {
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

func (c *Client) GetTaxes(request TaxRequest) (taxAmount string, err error) {
	err = c.request(request, &taxAmount)
	return
}

func (c *Client) request(request Request, response interface{}) (err error) {
	var req *http.Request
	body := getRequestBody(request)

	if req, err = http.NewRequest("post", host, body); err != nil {
		err = fmt.Errorf("error creating taxjar request: %v", err)
		return
	}

	return
}
