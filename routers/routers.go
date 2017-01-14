// Package routers contains a minimal set of routers to test
// with lots of other optional routers
package routers

import (
	"fmt"
	"net/http"
	"runtime"

	"github.com/kennygrant/routebench/routes"
)

// This pkg global stores the routers which should only mutate on init.
var routers []Router

// Router defines an abstract wrapper around routers which supports setup,
// adding routes, and handling requests
type Router interface {
	Name() string
	Setup() error
	Add(string, string, http.HandlerFunc)
	Serve(w http.ResponseWriter, r *http.Request)
	ParseHandler(w http.ResponseWriter, r *http.Request)
	ParamRegexp() bool
}

// All returns the list of routers to test
func All() []Router {
	return routers
}

// Setup sets up these routers by addng the routes given to the router
// and recording memory usage before and after doing so.
func Setup(routes []routes.Route, handler http.HandlerFunc) {

	// for each mux, call a setup function
	println("\nRouter Allocations                            Heap Allocations\n")

	// Set up simple patterns on all the routes before benchmarking starts
	for i, r := range routes {
		routes[i].PatternSimple = r.SimplifyPattern()
	}

	for _, router := range routers {

		if handler == nil {
			handler = router.ParseHandler
		}

		// Count heap alloc
		start := allocStat()

		err := router.Setup()
		if err != nil {
			fmt.Printf("%s: error setting up router:%s", router.Name(), err)
		}

		for _, r := range routes {

			if router.ParamRegexp() {
				router.Add(r.Method, r.Pattern, handler)
			} else {
				router.Add(r.Method, r.PatternSimple, handler)
			}
		}

		// Count heap alloc
		netAllocs := allocStat() - start
		fmt.Printf("Router:%s         	              %d Bytes\n", paddedName(router.Name()), netAllocs)
	}

	// Print a header line for the results
	println("\nRouter         	           	            Runs 	     Time/op	           Bytes/op          Allocs/op")
	println("")

}

func allocStat() uint64 {
	m := new(runtime.MemStats)
	runtime.GC()
	runtime.ReadMemStats(m)
	return m.HeapAlloc
}

func paddedName(name string) string {
	if len(name) < 10 {
		return name + "    "
	}
	return name
}

// TestHandler is a dummy web handler which simply writes 200 ok and hello world
func TestHandler(w http.ResponseWriter, r *http.Request) {
	//	w.WriteHeader(200)
	//	w.Write([]byte("hello world"))
}

// NullHandler is a dummy web handler which does nothing
func NullHandler(w http.ResponseWriter, r *http.Request) {

}
