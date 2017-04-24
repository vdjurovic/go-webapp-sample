package controllers

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func userRegistration(writer http.ResponseWriter, request *http.Request, router httprouter.Params) {

}

func getAllUsers(writer http.ResponseWriter, requset *http.Request, _ httprouter.Params) {
	fmt.Fprintf(writer, "No users found")
}
