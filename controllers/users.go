package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
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
	// decoder := json.NewDecoder(request.Body)
	// var user model.User
	// err := decoder.Decode(&user)
	// if err != nil {
	// 	panic(err)
	// }
	// defer request.Body.Close()
	log.Println(user)
	out, _ := json.Marshal(user)
	writer.Write(out)
}

func getAllUsers(writer http.ResponseWriter, requset *http.Request, _ httprouter.Params) {
	fmt.Fprintf(writer, "No users found")
}
