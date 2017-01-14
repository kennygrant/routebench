// RouteBench is a flexible set of router benchmarks
// inspired by httprouter tests but easier to include/exclude routers
// and does not pull down large dependencies on install
// To control which routers are benchmarked, see /routers/load.go
// Usage: go test -bench=. -timeout=2m -benchtime 2s -benchmem
package main

func main() {
	println("Usage: go test -bench=. -timeout=2m -benchtime 2s -benchmem")
}

// All the action is in routebench_test.go
