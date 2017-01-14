package buffalo

import (
	"fmt"
	"net/http"

	"github.com/gobuffalo/buffalo"
)

// Router is a wrapper for the buffalo mux
type Router struct {
	router *buffalo.App
}

// ParamRegexp returns true if this router supports regexp params
// otherwise params will be passed as :user_id not {user_id:[0-9]+}
func (r *Router) ParamRegexp() bool {
	return true
}

// Name is the name of this router wrapper.
func (r *Router) Name() string {
	return "buffalo mux"
}

// Setup sets up the router
func (r *Router) Setup() error {
	r.router = buffalo.New(buffalo.Options{})
	return nil
}

// Add adds this route to the router
func (r *Router) Add(method string, pattern string, h http.HandlerFunc) {

	if method == http.MethodPost {
		r.router.POST(pattern, wrap(h))
	} else {
		r.router.GET(pattern, wrap(h))
	}

}

// Serve calls the ServeHTTP function for the router with this request.
func (r *Router) Serve(w http.ResponseWriter, req *http.Request) {
	r.router.ServeHTTP(w, req)
}

func wrap(h http.HandlerFunc) buffalo.Handler {
	return func(c buffalo.Context) error {
		h(c.Response(), c.Request())
		return nil
	}
}

// ParseHandler does nothing for std mux - could do something? attempt to get query strings at least?
func (r *Router) ParseHandler(w http.ResponseWriter, req *http.Request) {
	params := req.URL.Query()
	if len(params) > 0 {
		w.WriteHeader(200)
		s := fmt.Sprintf("%v", params)
		w.Write([]byte(s))
		return
	}
}
