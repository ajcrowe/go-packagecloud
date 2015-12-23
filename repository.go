package packagecloud

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

var (
	repoUri = fmt.Sprintf("/api/%s/repos", apiVersion)
)

// Repository struct this is as yet undocumented see here for current docs
// https://packagecloud.io/docs/api#object_Repository
type Repository struct {
	// User defined name for this repository
	Name string `json:"name"`

	// Path to repository in the format <username>/<reponame>
	Path string `json:"path"`

	// Undocumented
	RepoType interface{} `json:"repo_type"`

	// Time repository was initially created
	CreatedAt time.Time `json:"created_at"`

	// Time repository was last updated
	UpdatedAt time.Time `json:"updated_at"`

	// Path accociated with this repo
	Paths struct {
		CreatePackage     string `json:"create_package"`
		PackageContents   string `json:"package_contents"`
		MasterTokens      string `json:"master_tokens"`
		CreateMasterToken string `json:"create_master_token"`
		Self              string `json:"self"`
	} `json:"paths"`

	// Url to install script to setup a host to use this repository
	Urls struct {
		InstallScript string `json:"install_script"`
	} `json:"urls"`
}

// Respoitory item returned when querying all user accessible repos
// this is documented here https://packagecloud.io/docs/api#object_Repository
type RepositoryListItem struct {
	//User defined name for this repository
	Name string `json:"name"`

	// Time repository was initially created
	CreatedAt time.Time `json:"created_at"`

	// Full URL to repository (note this is not the api endpoint)
	URL string `json:"url"`

	// Last manual (human) push to this repository in human readable form
	// format: <int> <period> ago
	LastPushHuman string `json:"last_push_human"`

	// Total packages present in the repository in human readable form
	// format <int> packages
	PackageCountHuman string `json:"package_count_human"`

	// Whether repository is private
	Private bool `json:"private"`

	// fully qualified name in the form <user>/<reponame>
	FQName string `json:"fqname"`
}

// Repository list struct
type RepositoryList []RepositoryListItem

// GetRepoListItemByName retruns a specific RepositoryListItem struct by name from an existing RepositoryList
func (repos RepositoryList) GetRepoListItemByName(name string) (*RepositoryListItem, error) {
	for _, repo := range repos {
		if repo.Name == name {
			return &repo, nil
		}
	}
	return nil, errors.New("packagecloud: repository not found")
}

// ListRepositories returns a slice of ResositoryListItem structs
func (c *Client) ListRepositories() (RepositoryList, *http.Response, error) {
	var repos RepositoryList
	reqUrl := createUriFromPath(repoUri)

	// Create HTTP request
	req, err := c.NewRequest("GET", reqUrl.String(), "", nil)
	if err != nil {
		return repos, nil, err
	}

	// Do request
	resp, err := c.do(req, http.StatusOK, &repos)
	if err != nil {
		return repos, resp, err
	}
	return repos, resp, nil
}

func (c *Client) GetRepository(user, name string) (Repository, *http.Response, error) {
	var repo Repository
	reqUrl := createUriFromPath(fmt.Sprintf("%s/%s/%s", repoUri, user, name))

	// Create HTTP request
	req, err := c.NewRequest("GET", reqUrl.String(), "", nil)
	if err != nil {
		return repo, nil, err
	}

	// Do request
	resp, err := c.do(req, http.StatusOK, &repo)
	if err != nil {
		return repo, resp, err
	}
	return repo, resp, nil

}

// newRepositoryRequest provides the struct required for creation of a new repository
type newRepositoryRequest struct {
	Name    string `json:"name"`
	Private bool   `json:"private"`
}

// newRepositoryResponse provides the struct to marshall the reponse to a new repository request
type newRepositoryResponse struct {
	Url string `json:"url"`
}

func (c *Client) CreateRepository(user, name string, private bool) (Repository, *http.Response, error) {
	var repo Repository
	reqUrl := createUriFromPath(repoUri)

	data, err := json.Marshal(&newRepositoryRequest{
		Name:    name,
		Private: private,
	})
	if err != nil {
		return repo, nil, err
	}

	// Create HTTP request
	req, err := c.NewRequest("POST", reqUrl.String(), "", bytes.NewReader(data))
	if err != nil {
		return repo, nil, err
	}

	// Do request
	var repoResp newRepositoryResponse
	resp, err := c.do(req, http.StatusCreated, &repoResp)
	if err != nil {
		return repo, resp, err
	}

	// Retrieve the repo data
	repo, resp, err = c.GetRepository(user, name)
	if err != nil {
		return repo, resp, err
	}

	// return new repo struct
	return repo, resp, nil

}

// Not Implemented in API Yet
func (c *Client) DeleteRepository(name string) error {
	return errors.New("packagecloud: DeleteRepository() not implemented in API")
}
