package gomarathon

import (
	"fmt"
)

// RegisterCallbackURL register a new callback url
func (c *Client) RegisterCallbackURL(uri string) (*Response, error) {
	options := &RequestOptions{
		Path:   "eventSubscriptions",
		Method: "POST",
		Params: &Parameters{
			CallbackURL: uri,
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

// GetEventSubscriptions gets all registered callback url
func (c *Client) GetEventSubscriptions() (*Response, error) {
	options := &RequestOptions{
		Path: "eventSubscriptions",
	}
	r, err := c.request(options)
	if err != nil {
		return nil, err
	}
	if r.Code != 200 {
		return nil, fmt.Errorf("request error")
	}
	return r, nil
}

// DeleteCallbackURL delete a particular callback url
func (c *Client) DeleteCallbackURL(uri string) (*Response, error) {
	options := &RequestOptions{
		Path:   "eventSubscriptions",
		Method: "DELETE",
		Params: &Parameters{
			CallbackURL: uri,
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
