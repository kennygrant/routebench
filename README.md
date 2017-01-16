# RouteBench  [![GoDoc](https://godoc.org/github.com/kennygrant/routebench?status.svg)](https://godoc.org/github.com/kennygrant/routebench) [![Go Report Card](https://goreportcard.com/badge/github.com/kennygrant/routebench)](https://goreportcard.com/report/github.com/kennygrant/routebench) [![CircleCI](https://circleci.com/gh/kennygrant/routebench.svg?style=svg)](https://circleci.com/gh/kennygrant/routebench)

Routebench is a set of packages for benchmarking, similar to the tests from httprouter, but making it a little easier to add/remove routers.

These route tests stress-test routers with larger route tables than you would normally see in anything but the largest monolithic application. 

The routers have very different characteristics, and the slower ones typically do more matching of more complex params (apart from stdlib, which is just slow), some trade speed for being able to easily prioritise routes or have certain patterns of nested routes. 

Finally, a time of 1860451 ns per route say is only 1.860451ms - this is a perfectly respectable time for handling a route and handing off to the handler. While it is desirable to have faster routing, a difference between 0.01ms and 0.1ms is probably not important for your app when it does real work, and anything below 1ms I would consider acceptable for most applications.

TODO: Add a short report for each router on the supported features, so that the benchmarks also function as a simple comparison between routers for different uses.

To use perform:

```Go 
go get github.com/kennygrant/routebench
```

cd to $GOPATH/github.com/kennygrant/routebench then run the benchmarks:

```Go 
go test -bench=. -timeout=2m -benchtime 1s -benchmem

Router Allocations                            Heap Allocations

Router:stdlib mux         	              23984 Bytes
Router:gorilla mux         	              656512 Bytes
Router:fragmenta mux         	              67808 Bytes
Router:httprouter mux         	              23640 Bytes

Router         	           	            Runs 	     Time/op	           Bytes/op          Allocs/op

BenchmarkStatic/stdlib_mux-4         	    1000	   1946545 ns/op	   20619 B/op	     537 allocs/op
BenchmarkStatic/gorilla_mux-4        	    1000	   1846382 ns/op	  115648 B/op	    1578 allocs/op
BenchmarkStatic/fragmenta_mux-4      	  100000	     13969 ns/op	       0 B/op	       0 allocs/op
BenchmarkStatic/httprouter_mux-4     	  100000	     16240 ns/op	       0 B/op	       0 allocs/op


Router Allocations                            Heap Allocations

Router:stdlib mux         	              5104 Bytes
Router:gorilla mux         	              124240 Bytes
Router:fragmenta mux         	              47072 Bytes
Router:httprouter mux         	              3176 Bytes

Router         	           	            Runs 	     Time/op	           Bytes/op          Allocs/op

BenchmarkCMS/stdlib_mux-4            	    3000	    526839 ns/op	   17445 B/op	     537 allocs/op
BenchmarkCMS/gorilla_mux-4           	    2000	    882967 ns/op	   24193 B/op	     642 allocs/op
BenchmarkCMS/fragmenta_mux-4         	   10000	    130708 ns/op	   20544 B/op	     312 allocs/op
BenchmarkCMS/httprouter_mux-4        	    5000	    215592 ns/op	   11888 B/op	     630 allocs/op


Router Allocations                            Heap Allocations

Router:stdlib mux         	              59504 Bytes
Router:gorilla mux         	              1849320 Bytes
Router:fragmenta mux         	              581216 Bytes
Router:httprouter mux         	              46160 Bytes

Router         	           	            Runs 	     Time/op	           Bytes/op          Allocs/op

BenchmarkGithub/stdlib_mux-4         	     300	   4370083 ns/op	   30504 B/op	     829 allocs/op
BenchmarkGithub/gorilla_mux-4        	     200	   8628078 ns/op	  246784 B/op	    2590 allocs/op
BenchmarkGithub/fragmenta_mux-4      	    2000	    982664 ns/op	   11428 B/op	     136 allocs/op
BenchmarkGithub/httprouter_mux-4     	   20000	     98858 ns/op	   24608 B/op	     296 allocs/op


Router Allocations                            Heap Allocations

Router:stdlib mux         	              59504 Bytes
Router:gorilla mux         	              1849352 Bytes
Router:fragmenta mux         	              581216 Bytes
Router:httprouter mux         	              46160 Bytes

Router         	           	            Runs 	     Time/op	           Bytes/op          Allocs/op

BenchmarkGithubNumeric/stdlib_mux-4  	     300	   4926243 ns/op	   35751 B/op	     901 allocs/op
BenchmarkGithubNumeric/gorilla_mux-4 	     100	  13113776 ns/op	  246784 B/op	    2590 allocs/op
BenchmarkGithubNumeric/fragmenta_mux-4         	    5000	    356540 ns/op	   10883 B/op	     136 allocs/op
BenchmarkGithubNumeric/httprouter_mux-4        	   20000	    107344 ns/op	   23759 B/op	     296 allocs/op


Router Allocations                            Heap Allocations

Router:stdlib mux         	              58064 Bytes
Router:gorilla mux         	              1849576 Bytes
Router:fragmenta mux         	              581216 Bytes
Router:httprouter mux         	              46160 Bytes

Router         	           	            Runs 	     Time/op	           Bytes/op          Allocs/op

BenchmarkGithubFuzz/stdlib_mux-4               	     300	   4592686 ns/op	   35767 B/op	     902 allocs/op
BenchmarkGithubFuzz/gorilla_mux-4              	     100	  12931693 ns/op	  246784 B/op	    2590 allocs/op
BenchmarkGithubFuzz/fragmenta_mux-4            	    5000	    324911 ns/op	    7617 B/op	     136 allocs/op
BenchmarkGithubFuzz/httprouter_mux-4           	   10000	    101702 ns/op	   23791 B/op	     296 allocs/op

PASS
ok  	github.com/kennygrant/routebench	38.060s
```

Not all routers are shown in these stats, just the default enabled ones. Not all are enabled because that would make go get pull down every routing package in the universe (as the httprouter tests do at present). Enable the ones your'e interested in and test those specifically to see numbers or send a pull request if you don't see the one you want to test. 

You can also see the results on [!CircleCI](https://circleci.com/gh/kennygrant/routebench/3#build-timing/containers/0) - I should set this up to do a build with all routers automatically at some point. 

To add a new router to benchmark, add a new adapter under routebench/routers, add to load.go (commented out by default please) and send a pull request. The adapter packages avoid pulling every router at once and let you easily choose which ones to benchmark. 

I may try to look into automatically updating results somewhere with a CI setup. 

To remove routers comment out the import and the load in routers/load.go. 

Tables of routes were based on those used by httprouter benchmarks, but adapted with different parameter placeholders. 
