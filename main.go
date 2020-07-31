package main

import (
	"github.com/MarkRosemaker/go-cartesian/api/points"
	"github.com/MarkRosemaker/go-server/server/api"

	"github.com/MarkRosemaker/go-server/server"
)

func main() {
	o := server.Options{
		ContentSource: "site",
		Endpoints: api.Endpoints{
			api.BaseEndpoint{
				URL:          "/api/points",
				InitFunc:     points.Init,
				ResponseFunc: points.Respond}},
		Verbose: true,
	}

	server.Run(o)
}
