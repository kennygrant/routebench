package gorilla

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Router is a wrapper for the gorilla mux.
type Router struct {
	router *mux.Router
}

// ParamRegexp returns true if this router supports regexp params
// otherwise params will be passed as :user_id not {user_id:[0-9]+}
func (r *Router) ParamRegexp() bool {
	return true
}

// Name is the name of this router wrapper.
func (r *Router) Name() string {
	return "gorilla mux"
}

// Setup sets up the router
func (r *Router) Setup() error {
	r.router = mux.NewRouter()
	return nil
}

// Add adds this route to the router
func (r *Router) Add(method string, pattern string, h http.HandlerFunc) {
	r.router.Handle(pattern, h)
}

// Serve calls the ServeHTTP function for the router with this request.
func (r *Router) Serve(w http.ResponseWriter, req *http.Request) {
	r.router.ServeHTTP(w, req)
}

// ParseHandler does nothing for std mux - could do something? attempt to get query strings at least?
func (r *Router) ParseHandler(w http.ResponseWriter, req *http.Request) {
	params := req.URL.Query()
	if len(params) > 0 {
		w.WriteHeader(200)
		s := fmt.Sprintf("%v", params)
		w.Write([]byte(s))
	}
}
