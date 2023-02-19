package main

import "net/http"

func runApplication(port string) error {
	http.HandleFunc("/time", timeHandler)
	return http.ListenAndServe(port, nil)
}
