package main

import (
	"net/http"

	"github.com/vdjurovic/go-webapp-sample/controllers"
)

// func extractPathParam(writer http.ResponseWriter, request *http.Request, router httprouter.Params) {
// 	fmt.Fprintf(writer, "parameter value: %s", router.ByName("name"))
// }

func main() {
	router := controllers.InitRoutes()
	http.ListenAndServe(":8080", router)
}
