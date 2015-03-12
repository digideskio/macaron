package main

import (
	"bytes"
	"io/ioutil"
	"path"
	"text/template"

	"github.com/Unknwon/com"
)

type Template struct {
	Filename     string
	OutputPath   string
	TemplatePath string
	Context      interface{}
}

func (self *Template) Exist() bool {
	return com.IsExist(self.Path())
}

func (self *Template) Write() error {

	output, err := self.Render()

	if err != nil {
		return err
	}

	if err := com.WriteFile(self.Path(), output.Bytes()); err != nil {
		return err
	}

	return nil
}

func (self *Template) Render() (bytes.Buffer, error) {
	var output bytes.Buffer

	template_data, err := ioutil.ReadFile(self.TemplatePath)

	if err != nil {
		return output, err
	}

	tmpl, err := template.New(self.Filename).Parse(string(template_data))

	if err != nil {
		return output, err
	}

	if err := tmpl.Execute(&output, self.Context); err != nil {
		return output, err
	}

	return output, err
}

func (self *Template) Path() string {
	return path.Join(self.OutputPath, self.Filename)
}
