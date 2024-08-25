package main

import (
	"fmt"
	"os"

	"github.com/alecthomas/kong"
	"github.com/frederic-arr/ripedb-go/ripedb"
)

type Context struct {
	Debug bool
}

var CLI struct {
	Debug bool   `help:"Enable debug mode."`
	Get   GetCmd `cmd:"" help:"Fetch a resource from the RIPE database."`
}

type GetCmd struct {
	Resource string `arg:"" name:"resource" help:"The resource of the resource to fetch."`
	Key      string `arg:"" name:"key" help:"The key of the resource to fetch."`
	Format   bool   `default:"true" negatable:"" short:"f" help:"Format the output or return the resource in its original formatting (including spaces, end-of-lines)."`
}

func (c *GetCmd) Run(ctx *Context) error {
	client := ripedb.NewRipeAnonymousClient()

	resource := c.Resource
	key := c.Key

	resp, err := client.Get(resource, key)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	header := "# This is the RIPE Database search service.\n# The objects are in RPSL (RFC 2622) format.\n# The RIPE Database is subject to Terms and Conditions."
	fmt.Printf("\033[90m%s\033[0m\n", header)

	if resp.ErrorMessages != nil {
		for _, errorMessage := range resp.ErrorMessages.ErrorMessage {
			if errorMessage.Text != nil {
				fmt.Printf("\033[31m%s\033[0m\n", *errorMessage.Text)
			}
		}

		os.Exit(1)
	}

	obj := resp.Objects.Object[0]
	longest := 0
	for _, attr := range obj.Attributes.Attribute {
		if len(attr.Name) > longest {
			longest = len(attr.Name)
		}
	}

	for _, attr := range obj.Attributes.Attribute {
		fmt.Printf("%-*s", longest+2, attr.Name+":")
		if attr.ReferencedType != nil {
			link := fmt.Sprintf("https://apps.db.ripe.net/db-web-ui/lookup?source=ripe&type=%s&key=%s", *attr.ReferencedType, attr.Value)
			fmt.Printf("\033]8;;%s\033\\\033[34m%s\033[0m\033]8;;\033\\", link, attr.Value)
		} else {
			fmt.Printf("%s", attr.Value)
		}

		if attr.Comment != nil {
			// Print in grey
			fmt.Printf(" \033[90m# %s\033[0m", *attr.Comment)
		}

		fmt.Println("")
	}

	return nil
}

func main() {
	ctx := kong.Parse(&CLI)
	err := ctx.Run(&Context{Debug: CLI.Debug})
	ctx.FatalIfErrorf(err)
}
