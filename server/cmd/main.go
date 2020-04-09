package main

import (
	"log"
	"net/http"
	web "workingtimeweb/server/http"
	"workingtimeweb/server/storage"
)

const port = ":4444"

func main() {
	httpPort := port

	repo := storage.NewMongoRepository()
	webService := web.New(repo)

	log.Printf("Running on port %s\n", httpPort)
	log.Fatal(http.ListenAndServe(httpPort, webService.Router))
}
