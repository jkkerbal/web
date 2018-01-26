package main

import (
	"github.com/gorilla/mux"
	"github.com/jkkerbal/web/controllers"

	"net/http"
)

func Router() *mux.Router {

	r := mux.NewRouter()

	r.HandleFunc("/login/", controllers.Login)
	r.HandleFunc("/login", controllers.Login)
	r.HandleFunc("/register/", controllers.Register)
	r.HandleFunc("/register", controllers.Register)

	return r

}

func main() {

	router := Router()

	http.ListenAndServe(":8080", router)

}
