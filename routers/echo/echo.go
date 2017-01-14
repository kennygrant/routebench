package echo

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

// Router is a wrapper for the echo router
type Router struct {
	echo   *echo.Echo
	router *echo.Router
}

// ParamRegexp returns true if this router supports regexp params
// otherwise params will be passed as :user_id not {user_id:[0-9]+}
func (r *Router) ParamRegexp() bool {
	return false
}

// Name is the name of this router wrapper.
func (r *Router) Name() string {
	return "echo mux"
}

// Setup sets up the router
func (r *Router) Setup() error {
	r.echo = echo.New()
	r.router = echo.NewRouter(r.echo)
	return nil
}

// Add adds this route to the router
func (r *Router) Add(method string, pattern string, h http.HandlerFunc) {
	switch method {
	case http.MethodPost:
		r.echo.POST(pattern, wrap(h))
	default:
		r.echo.GET(pattern, wrap(h))
	}

}

// Serve calls the ServeHTTP function for the router with this request.
func (r *Router) Serve(w http.ResponseWriter, req *http.Request) {
	r.echo.ServeHTTP(w, req)
}

func wrap(h http.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
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
