package main

import (
	"log"
	"net/http"
	"os"
	objects "simple-store/object"
)

func main() {
	http.HandleFunc("/objects/", objects.Handler)
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
}
