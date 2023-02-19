package main

import "log"

const port = ":8795"

func main() {
	if err := runApplication(port); err != nil {
		log.Panicf("Failed to run application: %v", err)
	}
}
