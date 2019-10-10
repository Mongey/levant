package template

import (
	"text/template"

	consul "github.com/hashicorp/consul/api"
)

// tmpl provides everything needed to fully render and job template using
// inbuilt functions.
type tmpl struct {
	consulClient    *consul.Client
	flagVariables   *map[string]string
	jobTemplateFile string
	variableFiles   []string
	errMissingKey   bool
}

const (
	jsonVarExtension      = ".json"
	terraformVarExtension = ".tf"
	yamlVarExtension      = ".yaml"
	ymlVarExtension       = ".yml"
	rightDelim            = "]]"
	leftDelim             = "[["
)

// newTemplate returns an empty template with default options set
func (t *tmpl) newTemplate() *template.Template {
	tmpl := template.New("jobTemplate")
	tmpl.Delims(leftDelim, rightDelim)

	if t.errMissingKey {
		tmpl.Option("missingkey=error")
	} else {
		tmpl.Option("missingkey=zero")
	}

	tmpl.Funcs(funcMap(t.consulClient))
	return tmpl
}
