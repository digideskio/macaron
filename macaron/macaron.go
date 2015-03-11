package main

import (
	"os"

	"github.com/codegangsta/cli"
)

const (
	Name        = "macaron"
	Version     = "0.1.0"
	Description = "Macaron project generator."
	Author      = "macaron - https://github.com/Unknwon/macaron"
)

func main() {
	app := cli.NewApp()

	app.Name = Name
	app.Version = Version
	app.Usage = Description
	app.Author = Author
	app.Email = ""

	app.Commands = []cli.Command{
		{
			Name:      "new",
			ShortName: "n",
			Usage:     "Create a new macaron project.",
			Action:    NewMacaronProject,
		},
	}

	app.Run(os.Args)
}
