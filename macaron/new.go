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
	"path/filepath"

	"github.com/Unknwon/com"
	"github.com/Unknwon/log"
	"github.com/codegangsta/cli"
)

var cmdNew = cli.Command{
	Name:   "new",
	Usage:  "Create a new empty macaron app.",
	Action: runNew,
	Flags:  []cli.Flag{},
}

func runNew(c *cli.Context) {
	setup(c)

	var (
		name     string
		location string
	)

	// App name must be given.
	if len(c.Args()) < 1 {
		fmt.Print("Please enter your app name: ")
		fmt.Scan(&name)
		if len(name) == 0 {
			log.Fatal("App name is missing or invalid(e.g. myapp)")
		}
	} else {
		name = c.Args().Get(0)
	}

	// Validate path if given(can't be empty, otherwise will not be parsed).
	if len(c.Args()) >= 2 {
		location = c.Args().Get(1)
	}
	if !filepath.IsAbs(location) {
		location = filepath.Join(defaultGOPATHSrc, location)
	}

	if !com.IsDir(location) {
		log.Fatal("Given path does not exist or is not a directory: %s", location)
	} else if err := pathInsideGOPATH(location); err != nil {
		log.Fatal("Invalid path: %v", err)
	}

	// Making a new app.
	app := &App{
		Name:     name,
		Location: location,
	}
	app.init(MustAsset("templates/new/config.toml"))
	app.build()
}
