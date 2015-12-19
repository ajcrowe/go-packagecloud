// Master Token API Methods
// https://packagecloud.io/docs/api#resource_master_tokens
//
package packagecloud

import (
	"log"
	"net/http"
)

type MasterToken struct {
	// User defined name for this token
	Name string `json:"name"`

	// Value of this token as an alphanumeric string
	Value string `json:"value"`

	// array of read tokens accociated with this master token
	ReadTokens []ReadToken `json:"read_tokens"`

	// path[self] the uri including the id of the master token
	// this is required when deleting a token
	Paths struct {
		Self string `json:"self"`
	} `json:"paths"`
}

// ListMasterTokens returns a slice of pointer to MasterToken structs
func (c *Client) ListMasterTokens(user, repo string) ([]*MasterToken, error) {
	var tokens []*MasterToken
	// Construct url to qurey
	reqUrl := createUri("repos", user, repo, "master_tokens")

	// Create new request
	req, err := http.NewRequest("GET", reqUrl.String(), nil)
	req.SetBasicAuth(c.Token, "")
	if err != nil {
		return tokens, err
	}

	//
	resp, err := c.do(req, http.StatusOK, &tokens)
	if err != nil {
		log.Println(resp.StatusCode)
		return tokens, err
	}
	return tokens, nil

}

// CreateMasterToken returns a newly created MasterToken struct
// The strings user, repo and name are required.
func (c *Client) CreateMasterToken(user, repo, name string) (MasterToken, error) {
	var token MasterToken
	return token, nil
}

// DestroyMasterToken removes the master token by using the Paths.Self
// field in the MasterToken struct.
func (c *Client) DestroyMasterToken(user, repo, tokenPath string) error {
	return nil
}
