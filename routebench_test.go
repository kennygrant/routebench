// To control which routers are benchmarked, see /routers/load.go
// For benchmarks:
// go test -bench=Bench -timeout=2m -benchtime 2s -benchmem
package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kennygrant/routebench/routers"
	"github.com/kennygrant/routebench/routes"
)

// To control which routers are benchmarked, see /routers/load.go

// TestAddRouters tests our router setup is are functional
func TestAddRouters(t *testing.T) {

	routers.Load()
	if len(routers.All()) == 0 {
		t.Fatalf("No routers loaded")
	}

}

// BenchTestStatic tests the routes in routes/static_routes.go
func BenchmarkStatic(b *testing.B) {

	// Add the muxes we want to consider
	routers.Load()

	// Call setup on the routers
	routers.Setup(routes.Static, routers.NullHandler)

	// Set up a recorder (unused)
	w := httptest.NewRecorder()

	// Set up a list of requests for these routes
	// Set up a fake request and recorder
	var requests []*http.Request
	for _, route := range routes.Static {
		r := httptest.NewRequest(route.Method, route.Path, nil)
		requests = append(requests, r)
	}

	// Now handle routes and benchmark
	for _, router := range routers.All() {
		// This Run will not return until the parallel tests finish.

		b.Run(router.Name(), func(b *testing.B) {
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				for _, r := range requests {
					router.Serve(w, r)
				}
			}
		})
	}
	println("")

}

// BenchmarkGithub tests the routes in routes/github_routes.go
func BenchmarkGithub(b *testing.B) {

	// Add the muxes we want to consider
	routers.Load()

	// Call setup on the routers
	routers.Setup(routes.GithubAPI, routers.NullHandler)

	// Set up a recorder (unused)
	w := httptest.NewRecorder()

	// Set up a list of requests for these routes
	// Set up a fake request and recorder
	var requests []*http.Request
	for _, route := range routes.GithubAPI {
		r := httptest.NewRequest(route.Method, route.Path, nil)
		requests = append(requests, r)
	}

	// Now handle routes and benchmark
	for _, router := range routers.All() {
		// This Run will not return until the parallel tests finish.
		b.Run(router.Name(), func(b *testing.B) {
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				for _, r := range requests {
					router.Serve(w, r)
				}
			}
		})
	}
	println("")

}

// BenchmarkGithubNumeric tests the routes in routes/github_routes.go
// with numeric params inserted
func BenchmarkGithubNumeric(b *testing.B) {

	// Add the muxes we want to consider
	routers.Load()

	// Call setup on the routers
	routers.Setup(routes.GithubAPI, routers.NullHandler)

	// Set up a recorder (unused)
	w := httptest.NewRecorder()

	// Set up a list of requests for these routes
	// Set up a fake request and recorder
	var requests []*http.Request
	for _, route := range routes.GithubAPI {
		r := httptest.NewRequest(route.Method, route.NumericFuzzPath(), nil)
		requests = append(requests, r)
	}

	// Now handle routes and benchmark
	for _, router := range routers.All() {
		// This Run will not return until the parallel tests finish.
		b.Run(router.Name(), func(b *testing.B) {
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				for _, r := range requests {
					router.Serve(w, r)
				}
			}
		})
	}
	println("")

}

// BenchmarkGithubFuzz tests the routes in routes/github_routes.go
// with random fuzzing inserted - let's see what breaks!
func BenchmarkGithubFuzz(b *testing.B) {

	// Add the muxes we want to consider
	routers.Load()

	// Call setup on the routers
	routers.Setup(routes.GithubAPI, routers.NullHandler)

	// Set up a recorder (unused)
	w := httptest.NewRecorder()

	// Set up a list of requests for these routes
	// Set up a fake request and recorder
	var requests []*http.Request
	for _, route := range routes.GithubAPI {
		r := httptest.NewRequest(route.Method, route.NumericFuzzPath(), nil)
		requests = append(requests, r)
	}

	// Now handle routes and benchmark
	for _, router := range routers.All() {
		// This Run will not return until the parallel tests finish.
		b.Run(router.Name(), func(b *testing.B) {
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				for _, r := range requests {
					router.Serve(w, r)
				}
			}
		})
	}
	println("")

}

// TestPaths sets up the github paths and logs them to stdout
// Usage: go test -v -run=TestPaths to check paths
func TestPaths(t *testing.T) {
	// Add the muxes we want to consider
	routers.Load()

	// Call setup on the routers
	routers.Setup(routes.GithubAPI, routers.NullHandler)

	for _, route := range routes.GithubAPI {
		t.Logf("PAT:%s SPAT:%s FUZZ:%s", route.Pattern, route.PatternSimple, route.NumericFuzzPath())
	}

}

// TestParse tests parsing a path on all routers
// Usage: go test -v -run=TestParse
func TestParse(t *testing.T) {

}

// TestParseFuzz tests parsing a fuzzed path on all routers
// Usage: go test -v -run=TestParseFuzz
func TestParseFuzz(t *testing.T) {

}

// TODO:
// write individual parse tests
// write some fuzzing tests which change params and introduce unusual params
