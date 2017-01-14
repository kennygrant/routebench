package httprouter

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Silence can be turned on to not report problems
// tree-based routers have with certain urls
var Silence = false

// Router is a wrapper for httprouter
type Router struct {
	router *httprouter.Router
}

// ParamRegexp returns true if this router supports regexp params
// otherwise params will be passed as :user_id not {user_id:[0-9]+}
func (r *Router) ParamRegexp() bool {
	return false
}

// Name is the name of this router wrapper.
func (r *Router) Name() string {
	return "httprouter mux"
}

// Setup sets up the router
func (r *Router) Setup() error {
	r.router = httprouter.New()
	return nil
}

// Add adds this route to the router
func (r *Router) Add(method string, pattern string, h http.HandlerFunc) {
	defer func() {
		if recover() != nil && !Silence {
			fmt.Printf("%s: panic handling route:%s\n", r.Name(), pattern)
		}
	}()
	r.router.Handle(method, pattern, wrap(h))
}

// Serve calls the ServeHTTP function for the router with this request.
func (r *Router) Serve(w http.ResponseWriter, req *http.Request) {
	r.router.ServeHTTP(w, req)
}

func wrap(h http.HandlerFunc) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		h(w, r)
	}
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
