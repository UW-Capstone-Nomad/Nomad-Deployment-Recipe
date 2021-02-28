package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

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
			// if c.NArg() > 0 {
			//   fileUrl = c.Args().Get(0)
			// }
			// if c.String("install") == "https://raw.githubusercontent.com/UW-Capstone-Nomad/Nomad-Deployment-Recipe/main/test/wordpress.nomad" {
			//   DownloadFile("./", fileUrl)
			//   fmt.Println("Downloading", fileUrl)
			// } else {
			//   fmt.Println("Invalid", fileUrl)
			// }
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

func DownloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
