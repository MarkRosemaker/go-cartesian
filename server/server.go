package server

import (
	"log"
	"net/http"

	serverdata "github.com/MarkRosemaker/go-cartesian/server/data"
	"github.com/MarkRosemaker/go-cartesian/site"
	"github.com/MarkRosemaker/go-cartesian/site/api"
)

// Start starts the server.
func Start() error {
	// load points to memory
	if err := serverdata.Load(); err != nil {
		log.Fatalf("error loading server data: %s", err)
	}

	// add api endpoints
	api.AddEndpoints()

	// initialize the pages of the site
	if err := site.InitContent(); err != nil {
		log.Fatalf("error initializing pages: %s", err)
	}

	// start the server
	log.Printf("starting server")
	return http.ListenAndServe(":8080", nil)
}
