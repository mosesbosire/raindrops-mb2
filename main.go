package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"bitbucket.org/armakuni/raindrops-mb2/controller"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := controller.Start()
	http.Handle("/", router)
	fmt.Printf("Sarting web server: http://127.0.0.1:%s", port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	if err != nil {
		log.Fatal("Could not start web server")
	}
}
