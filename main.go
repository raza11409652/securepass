package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/raza11409652/securepass/routers"
)

func main() {
	r := mux.NewRouter().StrictSlash(true)
	routers.RouterMUX(r)
	routers.RouterContent(r)
	log.Fatal(http.ListenAndServe(":8080", r))
}
