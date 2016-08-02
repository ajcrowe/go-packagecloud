package packagecloud

// Master Token API Methods
// https://packagecloud.io/docs/api#resource_master_tokens

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

// MasterToken struct see https://packagecloud.io/docs/api#resource_master_tokens
type MasterToken struct {
	// User defined name for this token
	Name string `json:"name"`

	// Value of this token as an alphanumeric string
	Value string `json:"value,omitempty"`

	// array of read tokens accociated with this master token
	ReadTokens []ReadToken `json:"read_tokens,omitempty"`

	// path[self] the uri including the id of the master token
	// this is required when deleting a token
	Paths MasterTokenPath `json:"paths,omitempty"`
}

// MasterTokenPath represents the returned 'paths' hash
type MasterTokenPath struct {
	Self string `json:"self,omitempty"`
}

// MasterTokens is a slice of MasterToken structs
type MasterTokens []MasterToken

// GetTokenByName returns the first matching master token from the slice
func (tokens MasterTokens) GetTokenByName(name string) (*MasterToken, error) {
	for _, token := range tokens {
		if token.Name == name {
			return &token, nil
		}
	}
	return nil, errors.New("packagecloud: master token not found")
}

type newMasterToken struct {
	Name string `json:"name"`
}

type newMasterTokenRequest struct {
	MasterToken newMasterToken `json:"master_token"`
}

// MasterTokenResp represents the response when creating a new MasterToken.
type masterTokenResp struct {
	ID           int       `json:"id"`
	RepositoryID int       `json:"repository_id"`
	Name         string    `json:"name"`
	Value        string    `json:"value"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// parse returns a properly formed MasterToken struct with
// expected paths[self] attribute
func (mt masterTokenResp) parse(user, repo string) MasterToken {
	return MasterToken{
		Name:  mt.Name,
		Value: mt.Value,
		Paths: MasterTokenPath{
			Self: fmt.Sprintf("/api/v1/repos/%s/%s/master_tokens/%d", user, repo, mt.ID),
		},
	}
}

// ListMasterTokens returns a slice of pointers to MasterToken structs
func (c *Client) ListMasterTokens(user, repo string) (MasterTokens, *http.Response, error) {
	var tokens MasterTokens
	// Construct URL for request
	reqURL := createURI("repos", user, repo, "master_tokens")

	// Create HTTP request
	req, err := c.NewRequest("GET", reqURL.String(), "", nil)
	if err != nil {
		return tokens, nil, err
	}

	// Dd request
	resp, err := c.do(req, http.StatusOK, &tokens)
	if err != nil {
		return tokens, resp, err
	}
	return tokens, resp, nil

}

// CreateMasterToken returns a newly created MasterToken struct
// The strings user, repo and name are required.
func (c *Client) CreateMasterToken(user, repo, name string) (MasterToken, *http.Response, error) {
	var token MasterToken
	// Construct URL for request
	reqURL := createURI("repos", user, repo, "master_tokens")

	// create json body
	data, err := json.Marshal(&newMasterTokenRequest{
		MasterToken: newMasterToken{
			Name: name,
		},
	})
	if err != nil {
		return token, nil, err
	}

	// Create HTTP request
	req, err := c.NewRequest("POST", reqURL.String(), "", bytes.NewReader(data))
	if err != nil {
		return token, nil, err
	}
	// Do request
	var tokenResp masterTokenResp
	resp, err := c.do(req, http.StatusCreated, &tokenResp)
	if err != nil {
		return token, resp, err
	}
	token = tokenResp.parse(user, repo)
	return token, resp, nil
}

// DestroyMasterToken removes the master token by using the Paths.Self
// field in the MasterToken struct.
func (c *Client) DestroyMasterToken(user, repo, path string) (*http.Response, error) {
	// Create HTTP request
	reqURL := createURIFromPath(path)
	req, err := c.NewRequest("DELETE", reqURL.String(), "", nil)
	if err != nil {
		return nil, err
	}
	// Do request
	resp, err := c.do(req, http.StatusNoContent, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
