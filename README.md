Cartesian API
=============

- [Cartesian API](#cartesian-api)
	- [Usage](#usage)
	- [Coding Challenge Implementation](#coding-challenge-implementation)
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

## Coding Challenge Implementation

### Summary

`main.go` starts a server (using my repository [`go-server`](https://github.com/markrosemaker/go-server)) that initializes the API endpoint and connects a handler with the API route.

The initialization is defined in [`go-cartesian/api/points/init.go`](https://github.com/MarkRosemaker/go-cartesian/blob/master/api/points/init.go). As per requirement, `Init` reads a list of points from 'data/points.json' into memory.

In [`go-cartesian/api/points/points.go`](https://github.com/MarkRosemaker/go-cartesian/blob/master/api/points/points.go), `func Respond(req *http.Request) interface{}` calculates the response to the request:

- The form is parsed to get the origin point as a [`point.Point`](https://github.com/MarkRosemaker/go-cartesian/blob/master/math/point/point.go) and the radius (a.k.a. distance) as an `int`.
- A slice of [`point.WithDistance`](https://github.com/MarkRosemaker/go-cartesian/blob/master/math/point/with-distance.go) (i.e. a [`points.WithDistance`](https://github.com/MarkRosemaker/go-cartesian/blob/master/math/points/with-distance.go)) is created, which holds both the points and their distances to the origin point.
- This slice is then sorted and points that are too far removed are filtered out. [See below](#sort-and-cut-vs-collect-and-sort) for two implementations.
- The result is a list of just the points, which are extracted from the filtered and sorted `points.WithDistance` list.

The resulting `interface{}` is encoded into JSON and returned to the user. This implementation of the handler can be viewed at [`go-server/server/api/endpoint_base.go`](https://github.com/MarkRosemaker/go-server/blob/master/server/api/endpoint_base.go).

If an error occurred, `Respond` will return an [`api.Error` (from `go-server`)](https://github.com/MarkRosemaker/go-server/blob/master/server/api/error.go). An HTTP status code is stored in such an object and the handler will write that code in the header.

### Sort and Cut vs. Collect and Sort

I've [implemented two methods](https://github.com/MarkRosemaker/go-cartesian/blob/master/math/points/with-distance.go) to filter the points.

- `sortCut` sorts the entire list and then cuts away all points that are not within the radius ("distance").
- `collectSort` first collects all points that are within the radius ("distance") and then sorts the collection.

Assuming we have a very, very large list of points instead of our 100 points, it would be beneficial to test which method is more efficient. In a real-world application, we might be able to work with some heuristics.

## Additions

For various reasons, like my enjoyment of the project and a desire to learn, I've done a bit more than what was required.

### go-server

Since my implementation of the server left the scope of the project and I wanted to continue programming it for myself, I decided to move this part of the code to a separate repository called [`go-server`](https://github.com/markrosemaker/go-server).

### Pages for Your Convenience

The server hosts templates and files from the folder `site`.

- At http://localhost:8080/valid, you can test the implementation with a form. The form will not allow invalid inputs.
- At http://localhost:8080/invalid, you can test out what happens if the input is invalid.
- At http://localhost:8080/too-slow, you can test out what happens if the request takes too long.
- At http://localhost:8080/visualize/, the points are visualized.
