package api

import (
	"context"
	"encoding/json"
)

// Role is a Role
type Role struct {
	ID          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Created     *Time  `json:"created,omitempty"`
	Description string `json:"description,omitempty"`
	Modified    *Time  `json:"modified,omitempty"`
}

// GetRoles gets all Passbolt Roles
func (c *Client) GetRoles(ctx context.Context) ([]Role, error) {
	msg, err := c.DoCustomRequest(ctx, "GET", "/roles.json", "v2", nil, nil)
	if err != nil {
		return nil, err
	}

	var roles []Role
	err = json.Unmarshal(msg.Body, &roles)
	if err != nil {
		return nil, err
	}
	return roles, nil
}
