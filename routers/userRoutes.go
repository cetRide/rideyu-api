package routers

import (
	controllers "github.com/cetRide/rideyu-api/controllers/users"
	"github.com/gorilla/mux"
)

func GetUserRoutes(r *mux.Router) {
	r.HandleFunc("/account/login", controllers.Login).Methods("GET")
	r.HandleFunc("/account/register", controllers.Register).Methods("POST")
	r.HandleFunc("/account/logout", controllers.Logout).Methods("GET")
}
