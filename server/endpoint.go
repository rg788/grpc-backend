package main

import (
	"fmt"

	"github.com/gorilla/mux"
)
func endPoint(){
	fmt.Printf("endpoint")
 r := mux.NewRouter()

  r.HandleFunc("/v1/ports",CreatePort).Methods("POST")
 r.HandleFunc("/v1/ports/{id}",UpdatePort).Methods("PUT")
 r.HandleFunc("/v1/ports/{id}",RetreivePort).Methods("GET")
 r.HandleFunc("/v1/ports/{id}",DeletePort).Methods("DELETE")
 r.HandleFunc("/v1/ports?page=1&count=10",ListPort).Methods("GET")


}