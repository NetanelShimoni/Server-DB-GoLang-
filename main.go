package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)


func Router() *mux.Router  {
	router := mux.NewRouter()
	router.HandleFunc("/users", getUsers).Methods("GET")
	router.HandleFunc("/users", postUsers).Methods("POST")
	router.HandleFunc("/users/{id}", updateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", deleteUser).Methods("DELETE")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Println(err)
	}
return router
}

func main() {
	Router()

}