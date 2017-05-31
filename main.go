package main

import (
	"log"
	"stress_test/conf"
	"stress_test/lib"
)

func main() {
	// lib.Get("/")
	config := conf.InitConf()
	// lib.Get(config, "/")
	if config.RequestType == "post" {
		log.Println("POST")
		lib.Post(config)
	}
	if config.RequestType == "get" {
		log.Println("GET")
		lib.Get(config)
	}

}
