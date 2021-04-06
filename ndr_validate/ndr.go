package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"github.com/urfave/cli/v2"
	"github.com/hashicorp/nomad/api"
)

func main() {
	type Client struct {
		c     *api.Client
		jobs  *api.Jobs
		alloc *api.Allocations
	}
	
	// NewClient returns a new client configured with the default values for Nomad
	// See: https://godoc.org/github.com/hashicorp/nomad/api#DefaultConfig
	func NewClient() (co *Client, err error) {
		co = &Client{}
		co.c, err = api.NewClient(api.DefaultConfig())
		if err != nil {
			return nil, err
		}
		co.jobs = co.c.Jobs()
		co.alloc = co.c.Allocations()
		return
	}
	


	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "validate",
				Value: "package",
				Usage: "a nomad package from source",
			},
		},
		Action: func(c *cli.Context) error {
			//fileUrl := c.String("validate")
			client, err := NewClient()
		
			res, _, err := co.jobs.Validate(job, nil)
			if err != nil {
				fmt.Println(num)
			}
			fmt.Println(res.Error)
			fmt.Println(res.Warnings)
			fmt.Println(res.ValidationErrors)
		
			return nil

		},
	}
	
	

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}


	
}
