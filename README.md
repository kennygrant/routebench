# RouteBench 

Routebench is a set of packages for benchmarking, similar to the tests from httprouter, but making it a little easier to add/remove routers.

These route tests stress-test routers with larger route tables than you would normally see in anything but the largest monolithic application. 

The routers have very different characteristics, and the slower ones typically do more matching of more complex params (apart from stdlib, which is just slow), some trade speed for being able to easily prioritise routes or have certain patterns of nested routes. 

Finally, a time of 1860451 ns per route say is only 1.860451ms - this is a perfectly respectable time for handling a route and handing off to the handler. While it is desirable to have faster routing, a difference between 0.01ms and 0.1ms is probably not important for your app when it does real work, and anything below 1ms I would consider acceptable for most applications.

To use perform:

```Go 
go get github.com/kennygrant/routebench
```

cd to $GOPATH/github.com/kennygrant/routebench then run the benchmarks:

```Go 
go test -bench=Bench -timeout=2m -benchtime 2s -benchmem

Router Allocations                    Heap Allocations

Router:stdlib mux         	               23504 Bytes
Router:gorilla mux         	              656512 Bytes
Router:fragmenta mux         	             67808 Bytes
Router:httprouter mux         	           23640 Bytes

Router         	           	            Runs 	     Time/op	           Bytes/op          Allocs/op

BenchmarkStatic/stdlib_mux-4         	    2000	   1869980 ns/op	   20958 B/op	     537 allocs/op
BenchmarkStatic/gorilla_mux-4        	    2000	   1788148 ns/op	  115648 B/op	    1578 allocs/op
BenchmarkStatic/fragmenta_mux-4      	  200000	     10866 ns/op	       0 B/op	       0 allocs/op
BenchmarkStatic/httprouter_mux-4     	  200000	     15893 ns/op	       0 B/op	       0 allocs/op


Router Allocations                   Heap Allocations

Router:stdlib mux         	              43600 Bytes
Router:gorilla mux         	            1849592 Bytes
Router:fragmenta mux         	           581216 Bytes
Router:httprouter mux         	          46160 Bytes

Router         	           	              Runs 	     Time/op	         Bytes/op          Allocs/op

BenchmarkGithub/stdlib_mux-4         	    1000	   4230743 ns/op	   26419 B/op	     829 allocs/op
BenchmarkGithub/gorilla_mux-4        	     300	   9120427 ns/op	  246784 B/op	    2590 allocs/op
BenchmarkGithub/fragmenta_mux-4      	  100000	     33328 ns/op	     361 B/op	       6 allocs/op
BenchmarkGithub/httprouter_mux-4     	   30000	    100419 ns/op	   24970 B/op	     296 allocs/op


Router Allocations                   Heap Allocations

Router:stdlib mux         	              42160 Bytes
Router:gorilla mux         	            1849640 Bytes
Router:fragmenta mux         	           581216 Bytes
Router:httprouter mux         	          46160 Bytes

Router         	           	              Runs 	     Time/op	         Bytes/op          Allocs/op

BenchmarkGithubNumeric/stdlib_mux-4  	    1000	   4111546 ns/op	   31668 B/op	     901 allocs/op
BenchmarkGithubNumeric/gorilla_mux-4 	     200	  12433868 ns/op	  246784 B/op	    2590 allocs/op
BenchmarkGithubNumeric/fragmenta_mux-4   10000	    279320 ns/op	   10737 B/op	     130 allocs/op
BenchmarkGithubNumeric/httprouter_mux-4  30000	     96666 ns/op	   23759 B/op	     296 allocs/op


Router Allocations                   Heap Allocations

Router:stdlib mux         	              42160 Bytes
Router:gorilla mux         	            1849464 Bytes
Router:fragmenta mux         	           581216 Bytes
Router:httprouter mux         	          46160 Bytes

Router         	           	                        Runs 	     Time/op	         Bytes/op          Allocs/op

BenchmarkGithubFuzz/stdlib_mux-4               	    1000	   4108527 ns/op	   31684 B/op	     902 allocs/op
BenchmarkGithubFuzz/gorilla_mux-4              	     200	  12854291 ns/op	  246784 B/op	    2590 allocs/op
BenchmarkGithubFuzz/fragmenta_mux-4            	   10000	    280920 ns/op	   10737 B/op	     130 allocs/op
BenchmarkGithubFuzz/httprouter_mux-4           	   30000	     97366 ns/op	   23791 B/op	     296 allocs/op

PASS
ok  	github.com/kennygrant/routebench	59.952s
```

To add a new router to benchmark, add a new adapter under routebench/routers, add to load.go (commented out by default please) and send a pull request. The adapter packages avoid pulling every router at once and let you easily choose which ones to benchmark. 

I may try to look into automatically updating results somewhere with a CI setup. 

To remove routers comment out the import and the load in routers/load.go. 

Tables of routes were based on those used by httprouter benchmarks, but adapted with different parameter placeholders. 