package main

import (
	"github.com/MarkRosemaker/go-cartesian/points"
	"github.com/MarkRosemaker/go-server/server/api"

	"github.com/MarkRosemaker/go-server/server"
)

func main() {
	o := server.Options{
		ContentSource: "templates",
		Endpoints: api.Endpoints{
			api.Endpoint{
				URL:         "/api/points",
				HandlerFunc: points.HandlerFunc}},
		Verbose: true,
	}

	server.Run(o)
}
