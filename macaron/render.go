package main

import (
	"bytes"
	"io/ioutil"
	"path"
	"text/template"

	"github.com/Unknwon/com"
)

type Template struct {
	Filename string
	Location string
	Template string
	Data     interface{}
}

func (self *Template) Render() error {
	template_data, err := ioutil.ReadFile(self.Template)

	if err != nil {
		return err
	}

	tmpl, err := template.New(self.Filename).Parse(string(template_data))

	if err != nil {
		return err
	}

	var output bytes.Buffer

	if err := tmpl.Execute(&output, self.Data); err != nil {
		return err
	}

	if err := com.WriteFile(path.Join(self.Location, self.Filename), output.Bytes()); err != nil {

	}

	return nil
}
