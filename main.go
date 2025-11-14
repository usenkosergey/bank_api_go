package main

import (
	"bank_api_go/controllers"
	"fmt"
	"log"
	"net/http"
)

func main() {

	fmt.Println("------")
	router := controllers.CreateRouter()

	log.Println("Server starting on :9099")
	log.Fatal(http.ListenAndServe(":9099", router))
}
