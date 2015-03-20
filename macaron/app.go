// Copyright 2015 Macaron Authors
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package main

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/Unknwon/com"
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
	Name        string   `toml:"-"`
	Location    string   `toml:"-"`
	Directories []string `toml:"directories"`
	Readme      readme
	Git         git
}

func (app *App) init(data []byte) {
	if _, err := toml.Decode(string(data), &app); err != nil {
		log.Fatal("Fail to parse config template: %v", err)
	}
}

func (app *App) build() {
	appPath := path.Join(app.Location, app.Name)

	if com.IsExist(appPath) {
		fmt.Print("App directory already exists, do you want to overwrite?(y/n): ")
		var answer string
		fmt.Scan(&answer)
		if strings.ToLower(answer) != "y" {
			fmt.Println("Existed directory is untouched.")
			return
		}
	}

	log.Info("Creating app directory...")
	log.Info("--> %s", appPath)
	if err := os.MkdirAll(appPath, os.ModePerm); err != nil {
		log.Fatal("Fail to create app directory: %v", err)
	}

	log.Info("Creating subdirectories...")
	for _, dir := range app.Directories {
		log.Info("--> %s", dir)
		if err := os.MkdirAll(path.Join(appPath, dir), os.ModePerm); err != nil {
			log.Fatal("Fail to create directory '%s': %v", dir, err)
		}
	}

	t := Template{
		OutputPath: path.Join(appPath, app.Name+".go"),
		Data:       MustAsset("templates/new/files/app.tmpl"),
		Context:    app,
	}

	// we can do something fancy here like ask to replace the file.
	if !t.exist() {
		if err := t.write(); err != nil {
			log.Fatal("Fail to generate file '%s': %v", t.OutputPath, err)
		}
	}
	return

	// for _, file := range newDefaultFiles {
	// 	template := Template{
	// 		OutputPath:   path.Join(main, file),
	// 		TemplatePath: path.Join("./templates/new/files/", strings.Replace(file, "/", "_", -1)) + ".tmpl",
	// 		Context:      app,
	// 	}

	// 	// we can do something fancy here like ask to replace the file.
	// 	if !template.Exist() {
	// 		if err := template.Write(); err != nil {
	// 			log.Fatal("Error: ", err)
	// 		}
	// 	}
	// }

	// if app.Readme.Enabled {
	// 	// ...
	// }

	// if app.Git.Enabled {
	// 	// ...
	// }
}
