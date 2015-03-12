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

// Macaron is a command line scaffold generator for Macaron web app.
package main

import (
	"os"
	"path/filepath"

	"github.com/Unknwon/com"
	"github.com/Unknwon/log"
	"github.com/codegangsta/cli"
)

const APP_VER = "0.0.1"

var (
	defaultGOPATHSrc string
)

func init() {
	log.Prefix = "[Macaron]"
	log.TimeFormat = "15:04:05"
}

func main() {
	app := cli.NewApp()

	app.Name = "macaron"
	app.Version = APP_VER
	app.Usage = "a command line scaffold generator for Macaron web app."
	app.Author = "macaron - https://github.com/Unknwon/macaron"
	app.Email = ""

	app.Commands = []cli.Command{
		cmdNew,
		{
			Name:      "scaffold",
			ShortName: "s",
			Usage:     "A scaffold in Macaron is a full set of model, router, view and logic.",
			Action:    NewMacaronScaffold,
		},
	}
	app.Run(os.Args)
}

func setup(c *cli.Context) {
	log.Info("App Version: %s", APP_VER)

	// Check GOPATHS.
	gopaths := com.GetGOPATHs()
	if len(gopaths) == 0 {
		log.Fatal("No GOPATH setting is available")
	}
	defaultGOPATHSrc = filepath.Join(gopaths[0], "src")

}
