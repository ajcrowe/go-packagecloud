package main

import (
	"fmt"
	"github.com/ajcrowe/go-packagecloud"
	"os"
)

func main() {
	client, err := packagecloud.NewClient("")
	if err != nil {
		fmt.Println(err)
		return
	}

	user, repoName := os.Args[1], os.Args[2]

	// list all repos for the api key
	repos, resp, err := client.ListRepositories()
	if err != nil {
		fmt.Println(err)
		fmt.Print(resp)
	}

	fmt.Println("repositories:\n-------------")
	for _, repo := range repos {
		fmt.Printf("name: %s\turl: %s\tprivate: %t\n", repo.Name, repo.URL, repo.Private)
	}

	repo, _, err := client.CreateRepository(user, repoName, true)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("\nrepository:\n-----------")
	fmt.Printf("name: %s\turl: %s\n", repo.Name, repo.Paths.Self)

}
