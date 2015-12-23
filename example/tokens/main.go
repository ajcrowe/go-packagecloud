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

	user, repo, tokenName := os.Args[1], os.Args[2], os.Args[3]

	// list all master tokens accociated with the user and repo
	tokens, _, err := client.ListMasterTokens(user, repo)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("master tokens:\n--------------")
	for _, token := range tokens {
		fmt.Printf("name: %s\tvalue: %s\turi: %s\n", token.Name, token.Value, token.Paths.Self)
	}

	// create new master token
	newMasterToken, _, err := client.CreateMasterToken(user, repo, tokenName)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("\nnew master token:\n-----------------")
	fmt.Printf("name: %s\tvalue: %s\turi: %s\n\n", newMasterToken.Name, newMasterToken.Value, newMasterToken.Paths.Self)

	// create new read only token in master token
	newReadToken, _, err := client.CreateReadToken(user, repo, newMasterToken.Paths.Self, "testreader")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("new read only token:\n-------------------")
	fmt.Printf("name: %s\tvalue: %s\n\n", newReadToken.Name, newReadToken.Value)

	// delete master token
	_, err = client.DestroyMasterToken(user, repo, newMasterToken.Paths.Self)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("deleted master token")

}
