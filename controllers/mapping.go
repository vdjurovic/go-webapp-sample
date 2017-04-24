package controllers

import "github.com/julienschmidt/httprouter"

func registerMappings(router *httprouter.Router) {
	router.POST("/users", userRegistration)
	router.GET("/users", getAllUsers)
}

// InitRoutes initializes routes for application
func InitRoutes() *httprouter.Router {
	router := httprouter.New()
	registerMappings(router)
	return router
}
