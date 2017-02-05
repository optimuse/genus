package genus

import (
	"bytes"
	"errors"
	"io/ioutil"
	"text/template"
)

type Template struct {
	Name        string // template name
	Source      string // source path
	TargetDir   string // target directory
	Filename    string // filename of generated code
	SkipExists  bool   // skip generation if exist
	SkipFormat  bool   // skip go format
	rawTemplate []byte // rawTemplate data in bytes
	rawResult   []byte
}

// Set raw template data
func (tmpl *Template) SetRawTemplate(raw []byte) (data []byte) {
	tmpl.rawTemplate = raw
	return
}

// Load raw template data from file
func (tmpl *Template) loadFile() (data []byte, err error) {
	if tmpl.Source == "" {
		return nil, errors.New("Empty source path")
	}

	data, err = ioutil.ReadFile(tmpl.Source)
	if err != nil {
		return nil, err
	}

	tmpl.rawTemplate = data
	return
}

// Render template by context
func (tmpl *Template) render(context interface{}) (data []byte, err error) {
	parsed, parsedErr := template.New(tmpl.Name).Parse(string(tmpl.rawTemplate))
	if parsedErr != nil {
		return nil, parsedErr
	}

	buf := bytes.NewBuffer([]byte{})
	if execErr := parsed.Execute(buf, context); execErr != nil {
		return nil, execErr
	}
	data = buf.Bytes()

	tmpl.rawResult = data
	return
}