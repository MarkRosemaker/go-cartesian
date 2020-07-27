package page

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/MarkRosemaker/go-cartesian/site/pagedata"
	"github.com/MarkRosemaker/go-cartesian/site/tpl"
)

// A Page struct contains information for quick serving of the page.
type Page struct {
	tpl     *template.Template
	tplName string
}

// Serve serves the page.
func (p Page) Serve(w http.ResponseWriter, req *http.Request) {

	// parse any forms (POST and/or queries in URL)
	if err := req.ParseForm(); err != nil {
		log.Printf("error parsing form on %q: %s", req.URL.Path, err)
	}

	p.executeTemplate(w, req)
}

func (p Page) executeTemplate(w http.ResponseWriter, req *http.Request) {
	err := p.tpl.ExecuteTemplate(w, p.tplName, pagedata.New(req))
	if err != nil {
		log.Printf("couldn't execute template %q: %s", req.URL.Path, err)
	}
}

// New creates a new page.
func New(path string) *Page {
	return &Page{
		tpl.GetTemplate(path),
		filepath.Base(path)}
}
