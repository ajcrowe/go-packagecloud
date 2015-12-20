// Master Token API Methods
// https://packagecloud.io/docs/api#resource_master_tokens
//
package packagecloud

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type MasterToken struct {
	// User defined name for this token
	Name string `json:"name"`

	// Value of this token as an alphanumeric string
	Value string `json:"value,omitempty"`

	// array of read tokens accociated with this master token
	ReadTokens []ReadToken `json:"read_tokens,omitempty"`

	// path[self] the uri including the id of the master token
	// this is required when deleting a token
	Paths struct {
		Self string `json:"self,omitempty"`
	} `json:"paths,omitempty"`
}

type MasterTokenRequest struct {
	MasterToken NewMasterToken `json:"master_token"`
}

type NewMasterToken struct {
	Name string `json:"name"`
}

// ListMasterTokens returns a slice of pointer to MasterToken structs
func (c *Client) ListMasterTokens(user, repo string) ([]*MasterToken, error) {
	var tokens []*MasterToken
	// Construct URL for request
	reqUrl := createUri("repos", user, repo, "master_tokens")

	// Create HTTP request
	req, err := c.NewRequest("GET", reqUrl.String(), nil)
	if err != nil {
		return tokens, err
	}

	//
	resp, err := c.do(req, http.StatusOK, &tokens)
	if err != nil {
		fmt.Printf("packagecloud: Error bad response code: %s", resp.StatusCode)
		return tokens, err
	}
	return tokens, nil

}

// CreateMasterToken returns a newly created MasterToken struct
// The strings user, repo and name are required.
func (c *Client) CreateMasterToken(user, repo, name string) (MasterToken, error) {
	var token MasterToken
	// Construct URL for request
	reqUrl := createUri("repos", user, repo, "master_tokens")

	// create json body
	data, err := json.Marshal(&MasterTokenRequest{
		MasterToken: NewMasterToken{
			Name: name,
		},
	})
	fmt.Println(string(data))
	if err != nil {
		fmt.Printf("packagecloud: Error marshalling body data: %s\n", err.Error())
		return token, err
	}

	// Create HTTP request
	req, err := c.NewRequest("POST", reqUrl.String(), bytes.NewReader(data))
	if err != nil {
		return token, err
	}
	// Do request
	resp, err := c.do(req, http.StatusCreated, &token)
	if err != nil {
		fmt.Printf("packagecloud: Error bad response code: %d\n", resp.StatusCode)
		return token, err
	}
	return token, nil
}

// DestroyMasterToken removes the master token by using the Paths.Self
// field in the MasterToken struct.
func (c *Client) DestroyMasterToken(user, repo, tokenPath string) error {
	// Create HTTP request
	reqUrl := fmt.Sprintf("%s/%s", BaseURL, tokenPath)
	req, err := c.NewRequest("DELETE", reqUrl, nil)
	if err != nil {
		return err
	}
	// Do request
	resp, err := c.do(req, http.StatusNoContent, nil)
	if err != nil {
		fmt.Printf("packagecloud: Error bad response code: %d\n", resp.StatusCode)
		return err
	}
	return nil
}
