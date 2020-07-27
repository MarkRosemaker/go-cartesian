package server

import (
	"log"
	"net/http"

	"github.com/MarkRosemaker/go-cartesian/server/serverdata"
	"github.com/MarkRosemaker/go-cartesian/site"
)

// Start starts the server.
func Start() {
	// load points to memory
	if err := serverdata.Load(); err != nil {
		log.Fatalf("error loading server data: %s", err)
	}

	// initialize the pages of the site
	if err := site.InitPages(); err != nil {
		log.Fatalf("error initializing pages: %s", err)
	}

	// start the server
	log.Printf("starting server")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("couldn't start server: %s", err)
	}
}
