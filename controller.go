package main

import "net/http"

func runApplication() error {
	http.HandleFunc("/time", timeHandler)

	const port string = ":8795"
	return http.ListenAndServe(port, nil)
}
