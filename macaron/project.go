package main

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/BurntSushi/toml"
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
		Filename: "app.go",
		Location: main,
		Template: "./templates/new/files/app.tmpl",
		Data:     self,
	}

	if err := template.Render(); err != nil {
		log.Fatal(err)
	}

	if self.Config.Readme.Enabled {
		// ...
	}

	if self.Config.Git.Enabled {
		// ...
	}
}
