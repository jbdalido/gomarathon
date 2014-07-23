package gomarathon

import (
	"fmt"
	"net/url"
)

func (c *Client) RegisterCallbackUrl(uri string) (*Response, error) {
	options := &RequestOptions{
		Path:   "eventSubscriptions",
		Method: "POST",
		Params: &Parameters{
			CallBackUrl: url.QueryEscape(uri),
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

func (c *Client) DeleteCallbackUrl(uri string) (*Response, error) {
	options := &RequestOptions{
		Path:   "eventSubscriptions",
		Method: "DELETE",
		Params: &Parameters{
			CallBackUrl: url.QueryEscape(uri),
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
