package fragmenta

import (
	"fmt"
	"net/http"

	"github.com/fragmenta/mux"
)

// Router is a wrapper for fragmenta/mux
type Router struct {
	router *mux.Mux
}

// ParamRegexp returns true if this router supports regexp params
// otherwise params will be passed as :user_id not {user_id:[0-9]+}
func (r *Router) ParamRegexp() bool {
	return true
}

// Name is the name of this router wrapper.
func (r *Router) Name() string {
	return "fragmenta mux"
}

// Setup sets up the router
func (r *Router) Setup() error {
	// caching is useful in real-world apps and so is on by default
	// if tests are not sophisticated (keep trying one path, no varying params)
	// this could distort results. I think most of the table tests would not be affected though.
	// static tests would reflect real-world results when hitting the same URLs repeatedly.
	// This trades off some memory for speed by caching mappings of URLS -> Handlers.
	// in a real app, many identical hits for the same URL will be recived (esp. / , but also
	// say /users/1 etc)
	// mux.MaxCacheEntries = 0

	r.router = mux.New()
	return nil
}

// Add adds this route to the router
func (r *Router) Add(method string, pattern string, h http.HandlerFunc) {

	switch method {
	case http.MethodPost:
		r.router.Post(pattern, wrap(h))
	default:
		r.router.Get(pattern, wrap(h))
	}

}

// Serve calls the ServeHTTP function for the router with this request.
func (r *Router) Serve(w http.ResponseWriter, req *http.Request) {
	r.router.ServeHTTP(w, req)
}

func wrap(h http.HandlerFunc) mux.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		h(w, r)
		return nil
	}
}

func handler(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// ParseHandler does nothing for std mux - could do something? attempt to get query strings at least?
func (r *Router) ParseHandler(w http.ResponseWriter, req *http.Request) {
	params, err := mux.ParamsWithMux(r.router, req)
	if err != nil {
		w.WriteHeader(200)
		s := fmt.Sprintf("%v", params)
		w.Write([]byte(s))
		return
	}

	w.WriteHeader(500)
	w.Write([]byte("error"))
}
