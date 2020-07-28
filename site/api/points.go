package api

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/MarkRosemaker/go-cartesian/points/point"
)

func points(w http.ResponseWriter, req *http.Request) {
	// ctx is the Context for this handler. Calling cancel closes the
	// ctx.Done channel, which is the cancellation signal for requests
	// started by this handler.
	var (
		ctx    context.Context
		cancel context.CancelFunc
	)
	timeout, err := time.ParseDuration(req.FormValue("timeout"))
	if err == nil {
		// The request has a timeout, so create a context that is
		// canceled automatically when the timeout expires.
		ctx, cancel = context.WithTimeout(context.Background(), timeout)
	} else {
		ctx, cancel = context.WithCancel(context.Background())
	}
	defer cancel() // Cancel ctx as soon as the function returns.

	// Store the time in ctx for use by code in other packages.
	// ctx = reqtime.NewContext(ctx, time.Now()) TODO

	// serve with correct MIME type
	w.Header().Set("Content-Type", "application/json")

	p, err := point.FromRequest(req)
	if err != nil {
		http.Error(w,
			fmt.Sprintf(`{"error":"%s"}`, err),
			http.StatusBadRequest)
		return
	}

	fmt.Println(ctx)
	fmt.Println(p)

}
