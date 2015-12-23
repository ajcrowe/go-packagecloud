package main

import (
	"fmt"
	"github.com/ajcrowe/go-packagecloud"
)

func main() {
	client, err := packagecloud.NewClient("")
	if err != nil {
		fmt.Println(err)
		return
	}

	// list all repos for the api key
	distros := client.GetDistributions()
	if err != nil {
		fmt.Println(err)
	}

	for _, distro := range distros["deb"] {
		for _, version := range distro.Versions {
			fmt.Printf("id: %d\tname: %s/%s\n", version.Id, distro.IndexName, version.IndexName)
		}
	}
}
