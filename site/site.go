package site

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/MarkRosemaker/go-cartesian/site/page"
)

const tmplSource string = "templates"

// InitPages initializes the pages from a folder of html templates.
func InitPages() error {
	// clean up template source
	src := filepath.Clean(tmplSource)

	// fix seperator on some systems
	if filepath.Separator != '/' {
		src = strings.ReplaceAll(src, string(filepath.Separator), "/")
	}

	// check first if the template folder actually exists (otherwise panic below)
	if _, err := os.Stat(src); os.IsNotExist(err) {
		return fmt.Errorf("can't find template source at %q", src)
	}

	// walk files in tmplSource to find template files
	return filepath.Walk(src,
		func(path string, info os.FileInfo, err error) error {
			// ignore folders
			if info.IsDir() {
				return nil
			}

			url := path

			// fix seperator on some systems
			if filepath.Separator != '/' {
				url = strings.ReplaceAll(url, string(filepath.Separator), "/")
			}

			// trim prefix to treat the template folder as its own file system
			url = strings.TrimPrefix(url, src)

			if strings.HasSuffix(path, "index.html") {
				// an .html file which we interpret as a template

				// remove suffix 'index.html'
				url = strings.TrimSuffix(url, "index.html")

				// create the page from the file path
				p := page.New(path)

				// serve the page when called
				http.HandleFunc(url, p.Serve)
			} else {
				// implement other behavior here to serve custom 404 pages, files, etc.
				log.Printf("file %q not used", path)
			}
			return nil
		})
}
