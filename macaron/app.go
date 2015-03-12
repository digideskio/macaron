package main

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/Unknwon/log"
)

var (
	newDefaultFiles = []string{
		"conf/app.ini",
		"models/models.go",
		"modules/base/template.go",
		"modules/base/base.go",
		"modules/middleware/context.go",
		"modules/settings/settings.go",
		"routers/routers.go",
		"templates/layout.tmpl",
		"public/css/base.css",
		"public/js/base.js",
	}
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

	log.Debug("Main Path: %s", main)

	if err := os.MkdirAll(main, 0777); err != nil {
		log.Fatal("Error: ", err)
	}

	for _, dir := range app.Directories {
		if err := os.MkdirAll(path.Join(main, dir), 0777); err != nil {
			log.Fatal("Error: ", err)
		}
	}

	template := Template{
		OutputPath:   path.Join(main, app.Name+".go"),
		TemplatePath: "./templates/new/files/app.tmpl",
		Context:      app,
	}

	// we can do something fancy here like ask to replace the file.
	if !template.Exist() {
		if err := template.Write(); err != nil {
			log.Fatal("Error: ", err)
		}
	}

	for _, file := range newDefaultFiles {
		template := Template{
			OutputPath:   path.Join(main, file),
			TemplatePath: path.Join("./templates/new/files/", strings.Replace(file, "/", "_", -1)) + ".tmpl",
			Context:      app,
		}

		// we can do something fancy here like ask to replace the file.
		if !template.Exist() {
			if err := template.Write(); err != nil {
				log.Fatal("Error: ", err)
			}
		}
	}

	if app.Readme.Enabled {
		// ...
	}

	if app.Git.Enabled {
		// ...
	}
}
