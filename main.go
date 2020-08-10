package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"github.com/ignite2405/login-regis/getusers"
	"github.com/ignite2405/login-regis/login"
	"github.com/ignite2405/login-regis/register"
)

func rootFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my Login/Register API Demo")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", rootFunc)
	router.HandleFunc("/login", login.LoginFunc).Methods("POST")
	router.HandleFunc("/register", register.RegisterFunc).Methods("POST")
	router.HandleFunc("/getall", getusers.GetAllUsers).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}
