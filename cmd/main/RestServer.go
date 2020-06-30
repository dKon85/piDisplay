package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"html"
	"log"
	"net/http"
)

func StartRESTServer(address string){

	r := mux.NewRouter()
	r.HandleFunc("/activate", activate)
	r.HandleFunc("/deactivate", deactivate)
	r.HandleFunc("/write/{text}", appendLine)
	http.Handle("/", r)


	log.Println("Listening on %q", address)
	log.Fatal(http.ListenAndServe(address, nil))
}

func handleRootAccess(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Fprintf(w, "Hello GET, %q", html.EscapeString(r.URL.Path))
	case "POST":
		fmt.Fprintf(w, "Hello POST, %q", html.EscapeString(r.URL.Path))
	}

}

func activate(w http.ResponseWriter, r *http.Request){
	ActivateDisplay()
}

func deactivate(w http.ResponseWriter, r *http.Request){
	DeactivateDisplay()
}

func appendLine(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	WriteLine(vars["text"], 0)
}
