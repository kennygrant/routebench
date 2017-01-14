// Package routes contains lists of routes to test in the routebench_tests.go test file
package routes

import (
	"fmt"
	"math/rand"
	"regexp"
)

// Route defines a route
type Route struct {
	Method        string // The method to request
	Pattern       string // The pattern to set up matching on params like {id:\d+}
	PatternSimple string // A simpler pattern using params like :id
	Path          string // an example of an actual path (including query string)
}

var re = regexp.MustCompile(`/{([^:]*):([^}]*)}`)

// SimplifyPattern returns a simplified pattern for routers which don't support
// regexp routes, and prefer just tokens like :id
func (r *Route) SimplifyPattern() string {
	// Note things like expensive wildcard matches are not tested here
	// so these types of routers get a performance advantage from that
	return re.ReplaceAllString(r.Pattern, `/:$1`)
}

// FuzzPath generates a new path from the pattern we have
func (r *Route) FuzzPath() string {
	b := make([]byte, 4)
	rand.Read(b)
	param := fmt.Sprintf("/%s", b)
	return re.ReplaceAllString(r.Pattern, param)
}

// NumericFuzzPath generates a new path from the pattern we have using numbers.
func (r *Route) NumericFuzzPath() string {
	n := rand.Int31()
	param := fmt.Sprintf("/%d", n)
	return re.ReplaceAllString(r.Pattern, param)
}
