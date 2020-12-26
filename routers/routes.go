package routers

import (
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)

	GetUserRoutes(router)
	GetPostRoutes(router)
	return router
}
