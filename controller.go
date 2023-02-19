package main

import (
	"log"
	"net/http"
)

func runApplication() error {
	http.HandleFunc("/time", timeHandler)

	const port string = ":8795"
	log.Print("log from chebelok")
	return http.ListenAndServe(port, nil)
}
