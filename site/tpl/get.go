// Package tpl defines behavior that relates to templates.
// Since this is not really the point of the exercise, the implementation is very rudimentary.
package tpl

import (
	"html/template"
	"log"
)

// GetTemplate clones the base template and adds a parsed template from the path.
func GetTemplate(path string) *template.Template {
	// clone the base template
	t := template.Must(baseTemplate.Clone())

	// parse the file
	if _, err := t.ParseFiles(path); err != nil {
		log.Fatalf("error parsing template at %q: %s", path, err)
	}

	return t
}
