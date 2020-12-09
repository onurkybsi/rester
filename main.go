package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/onurkybsi/rester/src/controller"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/ping", controller.Ping).Methods("GET")

	fmt.Println("Listening on 8080...")

	http.ListenAndServe(":8080", router)
}
