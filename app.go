package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strconv"

	"github.com/urfave/cli"
)

type Info struct {
	Name string
	Size int64
	Path string
}

func main() {

	app := cli.NewApp()
	app.Name = "MyCli"
	app.Usage = "finding zero byte file"
	app.Version = "1.0.0"

	// flags for option command
	var fileLocation string
	// flags for fileSize
	var fileSize string

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "location",
			Value:       "files/",
			Usage:       "specify files location",
			Destination: &fileLocation,
		},

		cli.StringFlag{
			Name:        "size",
			Value:       "0",
			Usage:       "specify minimum files size",
			Destination: &fileSize,
		},
	}

	//execute app
	app.Action = func(c *cli.Context) error {
		fmt.Println("file location ", fileLocation)
		fileSizeConv, err := strconv.ParseInt(fileSize, 10, 64)
		if err != nil {
			log.Fatal(err)
			return nil
		}
		fmt.Println("finding the file size equal to ", fileSize)
		printFiles(fileLocation, fileSizeConv)
		return nil
	}

	app.Run(os.Args)

}

// print specific files with name and size and location
func printFiles(directory string, fileSize int64) {
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.IsDir() {
			printFiles(path.Join(directory, file.Name()), fileSize)
		}

		if file.Size() == fileSize {
			fmt.Println("Name :", file.Name())
		}
	}

}
