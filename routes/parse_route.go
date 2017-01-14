package routes

// Single defines a single route to parse (with params)
// Note the duplicate params
var Single = []Route{
	{Method: "GET", Pattern: "/users/{id:\\d+}/update", Path: "/users/1/update?id=12&foo=bar"},
}
