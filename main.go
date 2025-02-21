package main

import (
	"fmt"
	"log"
	"net/http"

	"go-app/database"
	"go-app/handlers"
)

func main() {
	database.InitDB()

	http.HandleFunc("/healthz", handlers.HealthCheckHandler)

	port := "8080"
	fmt.Println("Server is running on port", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
