package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "install",
				Value: "package",
				Usage: "pull a nomad package from source",
			},
		},
		Action: func(c *cli.Context) error {
			fileUrl := c.String("install")
			DownloadFile("wordpress.nomad", fileUrl)
			fmt.Println("Downloading", fileUrl)
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}