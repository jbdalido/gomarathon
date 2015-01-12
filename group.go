package gomarathon

import (
	"fmt"
)

/*
* Group requests
 */

// ListGroups is listing all groups
func (c *Client) ListGroups() (*Response, error) {
	return c.ListGroupsByCmd("")
}

// ListAppsByCmd list all apps by cmd filter
func (c *Client) ListGroupsByCmd(filter string) (*Response, error) {
	options := &RequestOptions{
		Path: fmt.Sprintf("groups"),
		Params: &Parameters{
			Cmd: filter,
		},
	}
	r, err := c.request(options)
	if err != nil {
		return nil, err
	}
	return r, nil
}

// GetGroup gets a single group
func (c *Client) GetGroup(groupID string) (*Response, error) {
	options := &RequestOptions{
		Path: fmt.Sprintf("groups/%s", groupID),
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

// CreateGroup Create a new Application
func (c *Client) CreateGroup(group *Group) (*Response, error) {
	fmt.Println("Creating Group")

	// TODO : VALIDATE DATAS
	options := &RequestOptions{
		Path:   "groups",
		Datas:  group,
		Method: "POST",
	}
	r, err := c.request(options)
	if r != nil {
		if r.Code == 201 {
			return r, nil
		}
	}

	fmt.Println("Error: ", err)

	return nil, err
}

// UpdateGroup update the group
func (c *Client) UpdateGroup(groupID string, group *Group) (*Response, error) {
	options := &RequestOptions{
		Path:   fmt.Sprintf("groups/%s", groupID),
		Datas:  group,
		Method: "PUT",
	}
	r, err := c.request(options)
	if r != nil {
		if r.Code == 204 {
			return r, nil
		}
	}
	return nil, err

}

// DeleteGroup delete this group from the cluster. Returns with version but when completed will respond
// with an event because this operation can take a long time.
func (c *Client) DeleteGroup(groupID string) (*Response, error) {
	options := &RequestOptions{
		Path:   fmt.Sprintf("groups/%s", groupID),
		Method: "DELETE",
	}
	r, err := c.request(options)
	if r != nil {
		if r.Code == 204 {
			return r, nil
		}
	}
	return nil, err
}
