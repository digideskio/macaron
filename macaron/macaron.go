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
			Usage:     "Create a new empty macaron project.",
			Action:    NewMacaronProject,
		},
		{
			Name:      "scaffold",
			ShortName: "s",
			Usage:     "A scaffold in Macaron is a full set of model, router, views and logic.",
			Action:    NewMacaronScaffold,
		},
	}

	app.Run(os.Args)
}
