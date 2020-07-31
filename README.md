Cartesian API
=============

- [Cartesian API](#cartesian-api)
	- [Usage](#usage)
	- [Implementation](#implementation)
		- [Summary](#summary)
		- [Sort and Cut vs. Collect and Sort](#sort-and-cut-vs-collect-and-sort)
	- [Additions](#additions)
		- [go-server](#go-server)
		- [Pages for Your Convenience](#pages-for-your-convenience)

## Usage

Install the files as usual:

`go get github.com/MarkRosemaker/go-cartesian`

Compile and run in the repository folder.

The API route is then available at http://localhost:8080/api/points.

## Implementation

### Summary

`main.go` starts a server (using my repository [`go-server`](https://github.com/markrosemaker/go-server)) that initializes the API endpoint and connects a handler with the API route.

The initialization is defined in [`go-cartesian/api/points/init.go`](https://github.com/MarkRosemaker/go-cartesian/blob/master/api/points/init.go). As per requirement, `Init` reads a list of points from 'data/points.json' into memory.

In [`go-cartesian/api/points/points.go`](https://github.com/MarkRosemaker/go-cartesian/blob/master/api/points/points.go), we define `func Respond(req *http.Request) interface{}`.

The resulting `interface{}` is encoded into JSON and returned to the user. This implementation of the handler can be viewed at [`go-server/server/api/endpoint_base.go`](https://github.com/MarkRosemaker/go-server/blob/master/server/api/endpoint_base.go).

If an error occurred, `Respond` will return an [`api.Error` (from `go-server`)](https://github.com/MarkRosemaker/go-server/blob/master/server/api/error.go). Since we store an HTTP status code in such an object, the handler can write the status code in the header.

### Sort and Cut vs. Collect and Sort

I've implemented two methods to filter the points.

- `sortCut` sorts the entire list and then cuts away all points that are not within the radius ("distance").
- `collectSort` first collects all points that are within the radius ("distance") and then sorts the collection.

Assuming we have a very, very large list of points instead of our 100 points, it would be beneficial to test which method is more efficient. The most efficient method may depend on the size of the list.

## Additions

### go-server

Since my implementation of the server left the scope of the project and I wanted to continue programming it for myself, I decided to move this part of the code to a separate repository called `go-server`.

### Pages for Your Convenience

The server hosts templates and files from the folder `site`.

At http://localhost:8080/safe, you can test the implementation with a form. The form will not allow unsafe inputs.

At http://localhost:8080/invalid, you can test out what happens if the input is invalid.

At http://localhost:8080/too-slow, you can test out what happens if the request takes too long.

At http://localhost:8080/visualize/, the points are visualized.
