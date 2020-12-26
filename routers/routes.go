package routers

import (
	"net/http"

	"github.com/cetRide/rideyu-api/middlewares"
	"github.com/gorilla/mux"

	controllers "github.com/cetRide/rideyu-api/controllers/users"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)

	router.Use(middlewares.JwtAuthentication)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	return router
}

type Routes []Route

var routes = Routes{

	Route{
		"SignUp",
		"POST",
		"/account/register",
		controllers.Register,
	},
	Route{
		"Login",
		"GET",
		"/account/login",
		controllers.Login,
	},
	Route{
		"Logout",
		"GET",
		"/account/logout",
		controllers.Logout,
	},
}
