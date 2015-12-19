package main

import (
	"github.com/ajcrowe/go-packagecloud"
	"log"
	"os"
)

func main() {
	client, err := packagecloud.NewClient("")
	var tokens []*packagecloud.MasterToken
	tokens, err = client.ListMasterTokens(os.Args[1], os.Args[2])
	if err != nil {
		log.Println(err)
	}
	for _, token := range tokens {
		log.Println(token)
	}
}
