go-packagecloud
===============

go-packagecloud is a Go client library for using the [Package Cloud](https://packagecloud.io) API.

GoDoc: [![GoDoc](https://godoc.org/github.com/ajcrowe/go-packagecloud?status.svg)](https://godoc.org/github.com/ajcrowe/go-packagecloud)

## Usage

```go
import "github.com/ajcrowe/go-packagecloud"
```

Create a new client to call the various services. For example to list 
repositories for the given API token:

```go
client := packagecloud.NewClient("token")
repos, _, err := client.ListRepositories()
```

## Development

This is currently under development and missing a number of functions.

* Repositories `complete`
* Master Tokens `complete`
* Read Tokens `complete`
* Licenses `complete`
* Packages `todo`

