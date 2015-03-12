package main

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/BurntSushi/toml"
)

type readme struct {
	Enabled bool
	Format  string
}

type git struct {
	Enabled bool
}

type App struct {
	Name        string `toml:"-"`
	Location    string `toml:"-"`
	Directories []string
	Readme      readme
	Git         git
}

func (app *App) Init() {
	if _, err := toml.DecodeFile("templates/new/config.toml", &app); err != nil {
		// handle error
	}

	fmt.Println(app)
}

func (app *App) Build() {

	main := path.Join(app.Location, app.Name)

	if err := os.MkdirAll(main, 0777); err != nil {
		log.Fatal(err)
	}

	for _, dir := range app.Directories {
		if err := os.MkdirAll(path.Join(main, dir), 0777); err != nil {
			log.Fatal(err)
		}
	}

	template := Template{
		Filename:     app.Name + ".go",
		OutputPath:   main,
		TemplatePath: "./templates/new/files/app.tmpl",
		Context:      app,
	}

	// we can do something fancy here like ask to replace the file.
	if !template.Exist() {
		if err := template.Write(); err != nil {
			log.Fatal(err)
		}
	}

	if app.Readme.Enabled {
		// ...
	}

	if app.Git.Enabled {
		// ...
	}
}
