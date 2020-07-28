package site

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/MarkRosemaker/go-cartesian/site/content"
)

const tmplSource string = "templates"

// InitContent initializes the content from a folder of templates and files.
func InitContent() error {
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

			// initialize the page or file
			var h http.Handler

			base := filepath.Base(path)
			name := strings.TrimSuffix(base, filepath.Ext(base))

			switch name {
			case "index":
				// remove base from url
				url = strings.TrimSuffix(url, base)

				// create page from template
				h, err = content.NewPage(path)
				if err != nil {
					return err
				}
			case "404":
				//  serve custom 404 page
				log.Printf("custom 404 pages not implemented yet")
			default:
				// just serve as standard file
				h, err = content.NewFile(path)
				if err != nil {
					return err
				}
			}

			// serve the page when called
			http.Handle(url, h)

			return nil
		})
}
