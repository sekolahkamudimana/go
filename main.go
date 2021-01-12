package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"log"
	"os"
)

func main() {
 	port := os.Getenv("PORT")

	route := mux.NewRouter()
	s := route.PathPrefix("/api").Subrouter() //Base Path

	//Routes

	s.HandleFunc("/createProfile", createProfile).Methods("POST")
	s.HandleFunc("/getAllUsers", getAllUsers).Methods("GET")
	s.HandleFunc("/getUserProfile", getUserProfile).Methods("POST")
	s.HandleFunc("/updateProfile", updateProfile).Methods("PUT")
	s.HandleFunc("/deleteProfile/{id}", deleteProfile).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":" + port, s)) // Run Server
}
