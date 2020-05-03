package internal

import (
	"net/http"
)

// Route struct defines the route mapping
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}
