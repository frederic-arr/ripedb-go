// Copyright (c) The RIPE DB Go Client Authors.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"
	"log"
	"log/slog"
	"os"

	"github.com/alecthomas/kong"
	"github.com/frederic-arr/ripedb-go/ripedb"
	"github.com/frederic-arr/ripedb-go/ripedb/models"
	"github.com/frederic-arr/rpsl-go"
)

type Context struct {
	Debug bool
}

var CLI struct {
	Debug    bool      `help:"Enable debug mode."`
	User     *string   `env:"RIPEDB_USER" help:"The user to use for authentication."`
	Password *string   `env:"RIPEDB_PASSWORD" help:"The password to use for authentication."`
	Key      *string   `env:"RIPEDB_KEYFILE" help:"The key to use for authentication."`
	Cert     *string   `env:"RIPEDB_CERTFILE" help:"The certificate to use for authentication."`
	Endpoint *string   `env:"RIPEDB_ENDPOINT" help:"The endpoint of the database."`
	Source   *string   `env:"RIPEDB_SOURCE" help:"The source of the database."`
	Get      GetCmd    `cmd:"" help:"Fetch a resource from the RIPE database."`
	Upsert   UpsertCmd `cmd:"" help:"Create or update a resource from the RIPE database."`
	Delete   DeleteCmd `cmd:"" help:"Delete a resource from the RIPE database."`
}

type GetCmd struct {
	Resource string `arg:"" name:"resource" help:"The resource of the resource to fetch."`
	Key      string `arg:"" name:"key" help:"The key of the resource to fetch."`
	Format   bool   `default:"true" negatable:"" short:"f" help:"Format the output or return the resource in its original formatting (including spaces, end-of-lines)."`
}

type UpsertCmd struct {
	Resource string `arg:"" name:"resource" help:"The resource of the resource to update."`
	Key      string `arg:"" name:"key" help:"The key of the resource to update."`
	Input    string `arg:"" name:"input" help:"RPSL object file with the new resource content." type:"path"`
	Format   bool   `default:"true" negatable:"" short:"f" help:"Format the output or return the resource in its original formatting (including spaces, end-of-lines)."`
}

type DeleteCmd struct {
	Resource string `arg:"" name:"resource" help:"The resource of the resource to delete."`
	Key      string `arg:"" name:"key" help:"The key of the resource to delete."`
	Format   bool   `default:"true" negatable:"" short:"f" help:"Format the output or return the resource in its original formatting (including spaces, end-of-lines)."`
}

func formatResponse(resp *models.Resource) {
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
			fmt.Printf(" \033[90m# %s\033[0m", *attr.Comment)
		}

		fmt.Println("")
	}
}

func (c *UpsertCmd) Run(ctx *Context, client *ripedb.RipeClient) error {
	resource := c.Resource
	key := c.Key
	input := c.Input

	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
        if err := file.Close(); err != nil {
            slog.Error("failed to close file", "error", err)
        }
    }()

	raw := []byte{}
	buf := make([]byte, 1024)
	for {
		n, err := file.Read(buf)
		if n == 0 {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		raw = append(raw, buf[:n]...)
	}

	obj, err := rpsl.Parse(string(raw))
	if err != nil {
		log.Fatal(err)
	}

	data := models.NewResourceFromRpslObject(obj)
	resp, err := (*client).PutResource(resource, key, data)
	if err != nil {
		log.Fatal(err)
	}

	formatResponse(resp)
	return nil
}

func (c *GetCmd) Run(ctx *Context, client *ripedb.RipeClient) error {
	resource := c.Resource
	key := c.Key

	resp, err := (*client).GetResource(resource, key)
	if err != nil {
		log.Fatal(err)
	}

	formatResponse(resp)
	return nil
}

func (c *DeleteCmd) Run(ctx *Context, client *ripedb.RipeClient) error {
	resource := c.Resource
	key := c.Key

	resp, err := (*client).DeleteResource(resource, key)
	if err != nil {
		log.Fatal(err)
	}

	formatResponse(resp)
	return nil
}

func main() {
	ctx := kong.Parse(&CLI)

	opts := ripedb.RipeClientOptions{
		Endpoint: CLI.Endpoint,
		Source:   CLI.Source,
		User:     CLI.User,
		Password: CLI.Password,
	}

	if CLI.Key != nil || CLI.Cert != nil {
		if CLI.Key == nil || CLI.Cert == nil {
			log.Fatal("both key and cert must be provided")
		}

		cert, err := os.ReadFile(*CLI.Cert)
		if err != nil {
			log.Fatal("error reading certificate:", err)
		}

		key, err := os.ReadFile(*CLI.Key)
		if err != nil {
			log.Fatal("error reading private key:", err)
		}

		opts.Certificate = &cert
		opts.Key = &key
	}

	client, err := ripedb.NewRipeClient(&opts)
	ctx.FatalIfErrorf(err)

	err = ctx.Run(&Context{Debug: CLI.Debug}, client)
	ctx.FatalIfErrorf(err)
}
