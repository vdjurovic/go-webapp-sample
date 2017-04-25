package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/vdjurovic/go-webapp-sample/db"
	"github.com/vdjurovic/go-webapp-sample/model"
)

func userRegistration(writer http.ResponseWriter, request *http.Request, router httprouter.Params) {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		panic(err)
	}
	log.Println(string(body))
	var user model.User
	err = json.Unmarshal(body, &user)
	if err != nil {
		panic(err)
	}
	// another way to read request body
	// decoder := json.NewDecoder(request.Body)
	// var user model.User
	// err := decoder.Decode(&user)
	// if err != nil {
	// 	panic(err)
	// }
	// defer request.Body.Close()
	log.Println(user)

	id := db.SaveUser(user)
	log.Println(id)
	writer.WriteHeader(http.StatusCreated)

}

func getAllUsers(writer http.ResponseWriter, requset *http.Request, _ httprouter.Params) {
	fmt.Fprintf(writer, "No users found")
}
