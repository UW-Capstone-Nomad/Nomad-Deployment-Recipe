package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "list",
				Value: "package",
				Usage: "display a nomad package from source",
			},
		},
		Action: func(c *cli.Context) error {
			files, _ := filepath.Glob("*")
			for _, num := range files {
				if strings.HasSuffix(num, ".nomad") {
					fmt.Println(num)
				}
			}
			return nil
		},
	}
	
	

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
