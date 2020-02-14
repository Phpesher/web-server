package main

import (
	Config "EcderGo/app/config"
	indexHandler "EcderGo/app/handls"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", indexHandler.IndexHandler)

	conf := Config.InitConfig()

	port := conf.ListeningPort

	// Rout
	conf.Routes[0] = "/"

	fmt.Printf("Listening on port: %s \n", port)
	fmt.Printf("http://localhost:%s", port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

