package main

import (
	"go-vanilla/server"
	"log"
	"net/http"
)

func main() {
	logger := log.Default()
	server := server.New("4000")

	http.HandleFunc("/", server.GetMessage)
	http.HandleFunc("/health", server.HealthCheck)

	logger.Printf("Server is running on http://localhost:%s \n", server.Port)
	logger.Fatal(http.ListenAndServe(":" + server.Port, nil))	
}