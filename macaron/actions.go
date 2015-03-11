package main

import (
	"fmt"
	"os"
	"path/filepath"

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

	location, err := filepath.Abs(args.Get(1))

	if err != nil {
		os.Exit(1)
	}

	if !com.IsExist(location) {
		// error - not a path..
		os.Exit(1)
	}

	if err := PathInsideGoPath(location); err != nil {
		// show error - you need to be inside your $GOPATH
		os.Exit(1)
	}

	// macaron new project_name
	project := Project{
		Name:     name,
		Location: location,
	}

	fmt.Println("name:", name)
	fmt.Println("location:", location)

	project.Init()
	project.Build()

	fmt.Println("working")
}

func NewMacaronScaffold(cli *cli.Context) {
	// macaron add user
}
