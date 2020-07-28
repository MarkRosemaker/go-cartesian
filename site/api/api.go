package api

import "net/http"

// AddEndpoints adds the api endpoints.
func AddEndpoints() {
	http.HandleFunc("/api/points", points)
}
