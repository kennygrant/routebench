package routers

import (
	"github.com/kennygrant/routebench/routers/fragmenta"
	"github.com/kennygrant/routebench/routers/gorilla"
	"github.com/kennygrant/routebench/routers/httprouter"
	"github.com/kennygrant/routebench/routers/stdlibmux"
	// Uncomment to test more routers
	// Pull requests welcome for extra router adapters
	//"github.com/kennygrant/routebench/routers/gin"
	//"github.com/kennygrant/routebench/routers/buffalo"
	//"github.com/kennygrant/routebench/routers/chi"
	//"github.com/kennygrant/routebench/routers/echo"
)

// Load adds the muxes we want to consider to our list
func Load() {
	if len(routers) > 0 {
		routers = make([]Router, 0)
	}

	// Optional extras - uncomment to test
	// to add a router send a pull request with an adapter

	// routers = append(routers, &gin.Router{})
	// routers = append(routers, &echo.Router{})
	// routers = append(routers, &chi.Router{})
	// routers = append(routers, &buffalo.Router{})

	// A default set which is not too huge
	// I want to keep this small
	routers = append(routers, &stdlibmux.Router{})
	routers = append(routers, &gorilla.Router{})
	routers = append(routers, &fragmenta.Router{})
	routers = append(routers, &httprouter.Router{})

	// Turn off to see which github APIs these routers have problems storing
	// this does mean these routers do a little less work as they skip
	// about 10 routes - required for httprouter and gin
	httprouter.Silence = true

}
