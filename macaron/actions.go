package main

import (
	"os"

	"github.com/Unknwon/com"
	"github.com/codegangsta/cli"
)

func NewMacaronProject(cli *cli.Context) {

	args := cli.Args()

	if len(args) < 2 {
		// error - need project name and location
		os.Exit(1)
	}

	name := args.Get(0)
	location := args.Get(1)

	if !com.IsExist(location) {
		// error - not a path..
		os.Exit(1)
	}

	// macaron new project_name
	project := Project{
		Name:     name,
		Location: location,
	}

	project.Init()
	project.Build()
}

func NewMacaronScaffold(cli *cli.Context) {
	// macaron add user
}
