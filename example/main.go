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

	var tokens []*packagecloud.MasterToken
	tokens, err = client.ListMasterTokens(os.Args[1], os.Args[2])
	if err != nil {
		fmt.Println(err)
	}
	for _, token := range tokens {
		fmt.Printf("Token: %s\tValue: %s\tURI: %s\n", token.Name, token.Value, token.Paths.Self)
	}

	newMasterToken, err := client.CreateMasterToken(os.Args[1], os.Args[2], os.Args[3])
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("New Token:")
	fmt.Printf("Token: %s\tValue: %s\tURI: %s\n", newMasterToken.Name, newMasterToken.Value, newMasterToken.Paths.Self)

}
