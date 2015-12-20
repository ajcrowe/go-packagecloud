package packagecloud

import (
	"errors"
	"time"
)

type Repository struct {
	Name              string    `json:"name"`
	CreatedAt         time.Time `json:"created_at"`
	URL               string    `json:"url"`
	LastPushHuman     string    `json:"last_push_human"`
	PackageCountHuman string    `json:"package_count_human"`
	Private           bool      `json:"private"`
	FQName            string    `json:"fqname"`
}

//
func (c *Client) ListRepositories() []*Repository {
	var repos []*Repository
	return repos
}

func (c *Client) GetRepository(user, name string) Repository {
	var repo Repository
	return repo
}

func (c *Client) CreateRepository(name string, private bool) (string, error) {
	var url string
	return url, nil
}

// Not Implemented in API Yet
func (c *Client) DeleteRepository(name string) error {
	return errors.New("packagecloud: DeleteRepository() not implemented in API")
}
