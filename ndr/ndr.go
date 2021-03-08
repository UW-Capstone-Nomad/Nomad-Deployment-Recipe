package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	// curl "github.com/andelf/go-curl"
	//"github.com/urfave/cli/v2"
	"https://github.com/urfave/cli/blob/master/docs/v2/manual.md"
	"os/exec"

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
			runFormulae()
			return nil
		},
		//Lin
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
			DeleteFile("wordpress.nomad", fileUrl)
			fmt.Println("Deleting", fileUrl)
			stopFormulae()
			return nil
			
			
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func runFormulae() error{
	cmd := exec.Command("sudo", "python3", "/home/rtan/Work_dir/Nomad/ndr/formulae.py")
	out, err := cmd.Output()

	if err != nil {
    	println(err.Error())
    	return err
	}
	fmt.Println("string(out)")
	return nil
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







//Lin
func DeleteFile() { 
   
    // Removing file 
    // Using Remove() function 
    e := os.Remove("wordpress.nomad") 
    //if err != nil { 
    //    log.Fatal(e) 
    if err != nil {
        fmt.Println(err)
        return
    }
       fmt.Println("File wordpress.nomad successfully deleted")
  
   }  
}

func stopFormulae() error{
	cmd := exec.Command("nomad", "job", "stop", "wordpress.nomad")
	out, err := cmd.Output()

	if err != nil {
    	println(err.Error())
    	return err
	}
	fmt.Println("string(out)")
	return nil
}

