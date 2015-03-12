package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	"github.com/BurntSushi/toml"
)

// conf/app.ini
// models/models.go
// modules/base/template.go
// modules/base/base.go
// modules/middleware/context.go
// modules/settings/settings.go
// routers/routers.go
// templates/helpers/
// template/layout.tmpl

var (
	NewDefaultFiles = []string{
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

type Project struct {
	Name     string
	Location string
	Config   ProjectConfig
}

type ProjectConfig struct {
	App    app
	Readme readme
	Git    git
}

type app struct {
	Directories []string
}

type readme struct {
	Enabled bool
	Format  string
}

type git struct {
	Enabled bool
}

func (self *Project) Init() {
	if _, err := toml.DecodeFile("templates/new/config.toml", &self.Config); err != nil {
		// handle error
	}

	fmt.Println(self.Config)
}

func (self *Project) Build() {

	main := path.Join(self.Location, self.Name)

	if err := os.MkdirAll(main, 0777); err != nil {
		log.Fatal(err)
	}

	for _, dir := range self.Config.App.Directories {
		if err := os.MkdirAll(path.Join(main, dir), 0777); err != nil {
			log.Fatal(err)
		}
	}

	template := Template{
		OutputPath:   path.Join(main, self.Name+".go"),
		TemplatePath: "./templates/new/files/main.tmpl",
		Context:      self,
	}

	// we can do something fancy here like ask to replace the file.
	if !template.Exist() {
		if err := template.Write(); err != nil {
			log.Fatal(err)
		}
	}

	for _, file := range NewDefaultFiles {
		template := Template{
			OutputPath:   path.Join(main, file),
			TemplatePath: path.Join("./templates/new/files/", strings.Replace(file, "/", "_", -1)) + ".tmpl",
			Context:      self,
		}

		// we can do something fancy here like ask to replace the file.
		if !template.Exist() {
			if err := template.Write(); err != nil {
				log.Fatal(err)
			}
		}
	}

	if self.Config.Readme.Enabled {
		// ...
	}

	if self.Config.Git.Enabled {
		// ...
	}
}
