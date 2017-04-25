package main

import (
	"net/http"

	"github.com/vdjurovic/go-webapp-sample/controllers"
	"github.com/vdjurovic/go-webapp-sample/db"
)

// func extractPathParam(writer http.ResponseWriter, request *http.Request, router httprouter.Params) {
// 	fmt.Fprintf(writer, "parameter value: %s", router.ByName("name"))
// }

func main() {
	db.InitDB()
	router := controllers.InitRoutes()
	http.ListenAndServe(":8080", router)
}
