package librenms

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	HostURL = "http://librenms.rukas.me"
)

type Client struct {
	HostURL    string
	apiKey     *string
	HTTPClient *http.Client
}

func NewClient(host *string, apiKey *string) (*Client, error) {
	c := Client{
		HostURL: HostURL,
		apiKey:  apiKey,
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}

	if host != nil {
		c.HostURL = *host
	}

	return &c, nil
}

func (c *Client) doRequest(req *http.Request) ([]byte, error) {
	req.Header.Set("X-Auth-Token", *c.apiKey)

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	return body, err
}
