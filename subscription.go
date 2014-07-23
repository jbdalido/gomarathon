package gomarathon

import (
	"fmt"
	"net/url"
)

// Register a new callback url
func (c *Client) RegisterCallbackURL(uri string) (*Response, error) {
	options := &RequestOptions{
		Path:   "eventSubscriptions",
		Method: "POST",
		Params: &Parameters{
			CallbackURL: url.QueryEscape(uri),
		},
	}
	r, err := c.request(options)
	if r != nil {
		if r.Code == 201 {
			return r, nil
		}
	}
	return nil, err
}

// Get all registered callback url
func (c *Client) GetEventSubscriptions() (*Response, error) {
	options := &RequestOptions{
		Path: "eventSubscriptions",
	}
	r, err := c.request(options)
	if err != nil {
		return nil, err
	}
	if r.Code != 200 {
		return nil, fmt.Errorf("Request Error")
	}
	return r, nil
}

// Delete a particular callback url
func (c *Client) DeleteCallbackURL(uri string) (*Response, error) {
	options := &RequestOptions{
		Path:   "eventSubscriptions",
		Method: "DELETE",
		Params: &Parameters{
			CallbackURL: url.QueryEscape(uri),
		},
	}
	r, err := c.request(options)
	if r != nil {
		if r.Code == 204 {
			return r, nil
		}
	}
	return nil, err
}
