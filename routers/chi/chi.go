package chi

import (
	"fmt"
	"net/http"

	"github.com/pressly/chi"
)

// Router is a wrapper for the chi mux
type Router struct {
	router *chi.Mux
}

// ParamRegexp returns true if this router supports regexp params
// otherwise params will be passed as :user_id not {user_id:[0-9]+}
func (r *Router) ParamRegexp() bool {
	return false
}

// Name is the name of this router wrapper.
func (r *Router) Name() string {
	return "chi mux"
}

// Setup sets up the router
func (r *Router) Setup() error {
	r.router = chi.NewMux()
	return nil
}

// Add adds this route to the router
func (r *Router) Add(method string, pattern string, h http.HandlerFunc) {

	if method == http.MethodPost {
		r.router.Post(pattern, h)
	} else {
		r.router.Get(pattern, h)
	}

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
		return
	}
}
