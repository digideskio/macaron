package main

import (
	"fmt"

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
}
