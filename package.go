package packagecloud

import (
	"time"
)

// PackageFragment struct
type PackageFragment struct {
	Name              string    `json:"name"`
	CreatedAt         time.Time `json:"crated_at"`
	DistroVersion     string    `json:"distro_verison"`
	Version           string    `json:"version"`
	Release           string    `json:"release"`
	Epoch             string    `json:"epoch"`
	Private           bool      `json:"private"`
	Type              string    `json:"type"`
	UploaderName      string    `json:"uploader_name"`
	RepositoryHTMLURL string    `json:"repository_html_url"`
	PackageURL        string    `json:"package_url"`
	PackageHTMLURL    string    `json:"package_html_url"`
}

// PackageDetails struct
type PackageDetails struct {
	Name           string    `json:"name"`
	DistroVersion  string    `json:"distro_version"`
	Architecture   string    `json:"architecture"`
	Repository     string    `json:"repository"`
	Size           int       `json:"size"`
	Summary        string    `json:"summary"`
	CreatedAt      time.Time `json:"created_at"`
	Filename       string    `json:"filename"`
	Description    string    `json:"description"`
	Md5Sum         string    `json:"md5sum"`
	Sha1Sum        string    `json:"sha1sum"`
	Sha256Sum      string    `json:"sha256sum"`
	Sha512Sum      string    `json:"sha512sum"`
	UploaderName   string    `json:"uploader_name"`
	Licenses       []string  `json:"licenses"`
	Version        string    `json:"version"`
	Release        string    `json:"release"`
	Epoch          int       `json:"epoch"`
	RepositoryURL  string    `json:"repository_url"`
	VersionsURL    string    `json:"versions_url"`
	PackageURL     string    `json:"package_url"`
	PackageHTMLURL string    `json:"package_html_url"`
	SelfURL        string    `json:"self_url"`
}

// PackageVersion struct
type PackageVersion struct {
	Name              string `json:"name"`
	VersionsCount     int    `json:"versions_count"`
	VersionsURL       string `json:"versions_url"`
	RepositoryURL     string `json:"repository_url"`
	RepositoryHTMLURL string `json:"repository_html_url"`
}
