package main

import (
	"log"
	"practical/config"
	"practical/routing"

	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	handleRequests()
}

func handleRequests() {

	r := mux.NewRouter()
	routing.BookWordsRouting(r)

	serveMux := http.NewServeMux()

	serveMux.Handle("/", r)

	log.Println("Starting server...")
	config, err := config.GetConfig()
	if err != nil {
		println(err.Error())
	} else {
		log.Panic(http.ListenAndServe(config.Production.PORT, serveMux))
	}

}
