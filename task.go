package gomarathon

import (
	//	"encoding/json"
	"fmt"
)

func (c *Client) ListTasks() (*Response, error) {
	options := &RequestOptions{
		Path: "tasks",
	}
	r, err := c.request(options)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (c *Client) GetAppTasks(appId string) (*Response, error) {
	options := &RequestOptions{
		Path: fmt.Sprintf("apps/%s/tasks", appId),
	}
	r, err := c.request(options)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (c *Client) KillTasks(appId string, host string, scale bool) (*Response, error) {
	options := &RequestOptions{
		Path:   fmt.Sprintf("apps/%s/tasks", appId),
		Method: "DELETE",
		Params: &Parameters{
			Host:  host,
			Scale: scale,
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

func (c *Client) KillTask(appId string, taskId string, scale bool) (*Response, error) {
	options := &RequestOptions{
		Path:   fmt.Sprintf("apps/%s/tasks/%s", appId, taskId),
		Method: "DELETE",
		Params: &Parameters{
			Scale: scale,
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
