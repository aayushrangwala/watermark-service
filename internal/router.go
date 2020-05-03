package internal

import (
	"net/http"

	"github.com/aayushrangwala/watermark-service/internal/util"

	"github.com/gorilla/mux"
)

// NewRouter is the function which creates the list of router
func NewRouter(routes []Route) *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = util.Logger(handler, route.Name)
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}
