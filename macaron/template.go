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
	"bytes"
	"io/ioutil"
	"os"
	"text/template"

	"github.com/Unknwon/com"
)

type Template struct {
	Filename   string
	OutputPath string
	Data       []byte
	Context    interface{}
}

func (t *Template) exist() bool {
	return com.IsExist(t.OutputPath)
}

func (t *Template) render() (*bytes.Buffer, error) {
	var output bytes.Buffer

	tmpl, err := template.New(t.OutputPath).Parse(string(t.Data))
	if err != nil {
		return nil, err
	}

	if err := tmpl.Execute(&output, t.Context); err != nil {
		return nil, err
	}

	return &output, nil
}

func (t *Template) write() error {
	output, err := t.render()
	if err != nil {
		return err
	}

	return ioutil.WriteFile(t.OutputPath, output.Bytes(), os.ModePerm)
}

// generateFile generates new file from given template data and context to path.
func generateFile(path string, data []byte, context interface{}) error {
	t := Template{
		OutputPath: path,
		Data:       data,
		Context:    context,
	}

	// we can do something fancy here like ask to replace the file.
	if t.exist() {
		return nil
	}
	return t.write()
}
