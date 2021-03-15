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
				Name:  "info",
				Value: "package",
				Usage: "display a nomad package from source",
			},
		},
		Action: func(c *cli.Context) error {
			fileUrl := c.String("info")
			fileUrl += ".nomad"
			files, _ := filepath.Glob("*")
			flag := false
			for _, num := range files {
				if strings.HasSuffix(num, ".nomad") {
					if num == fileUrl {
						fmt.Println(num)
						fmt.Println("Installed")
						flag = true
						break;
					} 
						
				}
			}
			if (!flag) {
				fmt.Println("Do not have this package!")
			}
			return nil
		},
	}
	// ./ndr -info wordpress
	

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
