package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func rootFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my Login/Register API Demo")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", rootFunc)
	router.HandleFunc("/login", loginFunc).Methods("POST")
	router.HandleFunc("/register", registerFunc).Methods("POST")
	router.HandleFunc("/getall", getAllUsers).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}
