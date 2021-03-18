package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	// curl "github.com/andelf/go-curl"
	"github.com/urfave/cli/v2"
	//"https://github.com/urfave/cli/blob/master/docs/v2/manual.md"
	"os/exec"

)

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "uninstall",
				Value: "package",
				Usage: "delete a nomad package from local, and stop the nomad job from cloud",
			},
		},
		//Lin
		Action: func(c *cli.Context) error {
			fileUrl := c.String("uninstall")
			DeleteFile("wordpress.nomad", fileUrl)
			
			fmt.Println("Deleting", fileUrl)
			//stopFormulae()
			return nil
			},
			
			
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
//Lin
func DeleteFile(filepath string, url string) error {
    // Removing file 
    // Using Remove() function 
    err := os.Remove("wordpress.nomad") 
    //if err != nil { 
    //    log.Fatal(e) 
    if err != nil {
        fmt.Println(err)
        return err
    }
       fmt.Println("File wordpress.nomad successfully deleted")
	   return nil
   }  


func stopFormulae() error{
	cmd := exec.Command("nomad", "job", "stop", "wordpress.nomad")
	_, err := cmd.Output()

	if err != nil {
    	println(err.Error())
    	return err
	}
	fmt.Println("string(out)")
	return nil
}

