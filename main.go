package main

import (
	"MyGram/config"
	"MyGram/router"
)

func main() {
	config.StartDB()
	router.Routes()
}
