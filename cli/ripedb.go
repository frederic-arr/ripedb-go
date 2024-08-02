package main

import (
	"fmt"
	"os"

	"github.com/frederic-arr/ripedb-go/ripedb"
)

func main() {
	client := ripedb.RipeAnonymousClient{
		Endpoint: ripedb.RIPE_PROD_ENDPOINT,
	}

	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Println("Usage: ripedb <resource> <key>")
		os.Exit(1)
	}

	resource := args[0]
	key := args[1]

	obj, err := ripedb.Lookup(&client, resource, key)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, attr := range obj.Attributes.Attribute {
		fmt.Printf("%-20s%s\n", attr.Name+":", attr.Value)
	}
}
