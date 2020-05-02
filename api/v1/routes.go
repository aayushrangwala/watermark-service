package v1

import (
	"net/http"

	v1 "github.com/aayushrangwala/User-Microservice/pkg/services/v1"
)

// Route struct defines the route mapping
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes is the list of route mapping from path to handler
type Routes []Route

var routes = Routes{
	Route{
		"GetUserProfile",
		"GET",
		"/user/profile",
		v1.GetUserProfile,
	},
	Route{
		"GetMicroserviceName",
		"GET",
		"/microservice/name",
		v1.GetMicroserviceName,
	},
	Route{
		"GetUserEmail",
		"GET",
		"/user/email",
		v1.GetUserEmail,
	},
}
