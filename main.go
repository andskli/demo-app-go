package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type server struct{}

func helloName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	fmt.Fprintf(w, "{\"person\": \"%v\"}", params["name"])
}

func main() {
	r := mux.NewRouter()

	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/{name}", helloName).Methods(http.MethodGet)

	// wrap this in a logging handler
	loggedRouter := handlers.LoggingHandler(os.Stdout, api)

	http.ListenAndServe(":8080", loggedRouter)
}
