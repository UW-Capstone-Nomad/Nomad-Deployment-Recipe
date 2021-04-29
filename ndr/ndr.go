package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	// curl "github.com/andelf/go-curl"
	"github.com/urfave/cli/v2"
	"strings"
	"os/exec"
	"io/ioutil"
)

func main() {
	app := cli.NewApp()
	app.Name = "Nomad deployment CLI"
	app.Usage = "Improve users' deployment experience"
	app.Commands = []*cli.Command{
		{
			Name:  "install",
			// Value: "package",
			Usage: "pull a nomad package from source",
			Action: func(c *cli.Context) error {
				fileUrl := c.Args().Get(0)
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
				// runFormulae()
				return nil
			},		
		},
		{
			Name:  "info",
			// Value: "package",
			Usage: "display a nomad package from source",
			Action: func(c *cli.Context) error {
				fileUrl := c.Args().Get(0)
				fileUrl += ".nomad"
				files, _ := filepath.Glob("*")
				flag := false
				for _, num := range files {
					
					if strings.HasSuffix(num, ".nomad") {
						
						if num == fileUrl {
							fmt.Println(num)
							fmt.Println("Installed")
							flag = true
							// break;
						} 
							
					}
					if strings.Compare(num, "Formulae.txt") == 0 {
						ReadFormulae(num)
						// break;	
					}
				}
				if (!flag) {
					fmt.Println("Do not have this package!")
				}
				return nil
			},
		},
		{
			Name:  "list",
			// Value: "package",
			Usage: "display a nomad package from source",
			Action: func(c *cli.Context) error {
				files, _ := filepath.Glob("*")
				for _, num := range files {
					if strings.HasSuffix(num, ".nomad") {
						fmt.Println(num)
					}
				}
				return nil
			},
		},
	  }

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func runFormulae() error{
	cmd := exec.Command("sudo", "python3", "/home/rtan/Work_dir/Nomad/ndr/formulae.py")
	_, err := cmd.Output()

	if err != nil {
    	println(err.Error())
    	return err
	}
	return nil
}
func ReadFormulae(filepath string) error {
	data, err := ioutil.ReadFile(filepath)
	// Get the data
	fmt.Println(string(data))
	return err
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
	fmt.Println(resp)
	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

