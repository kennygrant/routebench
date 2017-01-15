package routes

// CMS is a selection of routes typical of a simple CMS application
// a mix of static and dynamic routes
var CMS = []Route{
	{Method: "GET", Pattern: "/", Path: "/"},
	// Users resource
	{Method: "GET", Pattern: "/users", Path: "/users"},
	{Method: "GET", Pattern: "/users/create", Path: "/users/create"},
	{Method: "POST", Pattern: "/users/create", Path: "/users/create"},
	{Method: "GET", Pattern: "/users/:id/update", Path: "/users/1/update"},
	{Method: "POST", Pattern: "/users/:id/update", Path: "/users/1/update"},
	{Method: "POST", Pattern: "/users/:id/destroy", Path: "/users/1/destroy"},
	{Method: "GET", Pattern: "/users/:id", Path: "/users/1"},
	// Pages resource
	{Method: "GET", Pattern: "/pages", Path: "/pages"},
	{Method: "GET", Pattern: "/pages/create", Path: "/pages/create"},
	{Method: "POST", Pattern: "/pages/create", Path: "/pages/create"},
	{Method: "GET", Pattern: "/pages/:id/update", Path: "/pages/1/update"},
	{Method: "POST", Pattern: "/pages/:id/update", Path: "/pages/1/update"},
	{Method: "POST", Pattern: "/pages/:id/destroy", Path: "/pages/1/destroy"},
	{Method: "GET", Pattern: "/pages/:id", Path: "/pages/1"},
	// Tags resource
	{Method: "GET", Pattern: "/tags", Path: "/tags"},
	{Method: "GET", Pattern: "/tags/create", Path: "/tags/create"},
	{Method: "POST", Pattern: "/tags/create", Path: "/tags/create"},
	{Method: "GET", Pattern: "/tags/:id/update", Path: "/tags/1/update"},
	{Method: "POST", Pattern: "/tags/:id/update", Path: "/tags/1/update"},
	{Method: "POST", Pattern: "/tags/:id/destroy", Path: "/tags/1/destroy"},
	{Method: "GET", Pattern: "/tags/:id", Path: "/tags/1"},
	// Posts resource
	{Method: "GET", Pattern: "/posts", Path: "/posts"},
	{Method: "GET", Pattern: "/posts/create", Path: "/posts/create"},
	{Method: "POST", Pattern: "/posts/create", Path: "/posts/create"},
	{Method: "GET", Pattern: "/posts/:id/update", Path: "/posts/1/update"},
	{Method: "POST", Pattern: "/posts/:id/update", Path: "/posts/1/update"},
	{Method: "POST", Pattern: "/posts/:id/destroy", Path: "/posts/1/destroy"},
	{Method: "GET", Pattern: "/posts/:id", Path: "/posts/1"},
	// Images resource
	{Method: "GET", Pattern: "/images", Path: "/images"},
	{Method: "GET", Pattern: "/images/create", Path: "/images/create"},
	{Method: "POST", Pattern: "/images/create", Path: "/images/create"},
	{Method: "GET", Pattern: "/images/:id/update", Path: "/images/1/update"},
	{Method: "POST", Pattern: "/images/:id/update", Path: "/images/1/update"},
	{Method: "POST", Pattern: "/images/:id/destroy", Path: "/images/1/destroy"},
	{Method: "GET", Pattern: "/images/:id", Path: "/images/1"},
}
