package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/cyse7125-sp25-team01/webapp-hello-world/database"
	"github.com/cyse7125-sp25-team01/webapp-hello-world/handlers"
)

func main() {
	database.InitDB()

	http.HandleFunc("/healthz", handlers.HealthCheckHandler)

	port := "8080"
	fmt.Println("Server is running on port", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
