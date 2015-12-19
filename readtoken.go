package packagecloud

import (
//"time"
)

type ReadToken struct {
	// Unique Id for the read token
	Id int `json:"id"`

	// Name of the read token
	Name string `json:"name"`

	// Read token value used for access
	Value string `json:"value"`
}

// ListReadTokens returns a slice of pointers to ReadToken structs which are
// accociated with the masterToken.
func (c *Client) ListReadTokens(user, repo, masterToken string) ([]*ReadToken, error) {
	var tokens []*ReadToken
	return tokens, nil
}

// CreateReadToken creates a new read token for the specified master token value.
func (c *Client) CreateReadToken(user, repo, masterToken, name string) (ReadToken, error) {
	var token ReadToken
	return token, nil
}

// DestroyReadToken deletes the read token with specified id from the master token
func (c *Client) DestroyReadToken(user, repo, masterToken string, id int) error {
	return nil
}
