package wix

import (
	"fmt"
	"net/http"
	"time"
)

type Client struct {
	APIKey     string
	SiteID     string
	HTTPClient *http.Client
}

func NewClient(apiKey, siteID string) *Client {
	return &Client{
		APIKey: apiKey,
		SiteID: siteID,
		HTTPClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

func (c *Client) newRequest(method, url string) (*http.Request, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", c.APIKey)
	req.Header.Set("wix-site-id", c.SiteID)
	req.Header.Set("Content-Type", "application/json")

	return req, nil
}

func (c *Client) TestConnection() error {
	req, err := c.newRequest(
		http.MethodGet,
		"https://www.wixapis.com/blog/v3/draft-posts",
	)
	if err != nil {
		return err
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("Wix API returned HTTP %d", resp.StatusCode)
	}

	return nil
}
