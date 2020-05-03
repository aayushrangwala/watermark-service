package services

import (
	v1 "github.com/aayushrangwala/User-Microservice/pkg/services/v1"
	"github.com/aayushrangwala/watermark-service/internal"
)



// Routes is the list of route mapping from path to handler
type Routes []internal.Route

var routes = Routes{
	{
		"GetUserProfile",
		"GET",
		"/user/profile",
		v1.GetUserProfile,
	},
	{
		"GetMicroserviceName",
		"GET",
		"/microservice/name",
		v1.GetMicroserviceName,
	},
	{
		"GetUserEmail",
		"GET",
		"/user/email",
		v1.GetUserEmail,
	},
}
