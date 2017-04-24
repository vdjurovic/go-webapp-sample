package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func extractPathParam(writer http.ResponseWriter, request *http.Request, router httprouter.Params) {
	fmt.Fprintf(writer, "parameter value: %s", router.ByName("name"))
}

func main() {
	router := httprouter.New()
	router.GET("/some/:name/path", extractPathParam)
	http.ListenAndServe(":8080", router)
}
