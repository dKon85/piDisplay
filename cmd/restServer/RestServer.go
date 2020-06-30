package restServer

import (
	"fmt"
	"github.com/gorilla/mux"
	"html"
	"log"
	"net/http"
	"piDisplay/cmd/tinker"
)

func StartRESTServer(address string){

	r := mux.NewRouter()
	r.HandleFunc("/activate", activate)
	r.HandleFunc("/deactivate", deactivate)
	r.HandleFunc("/append/{text}", appendLine)
	http.Handle("/", r)


	log.Printf("Listening on %q", address)
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
	tinker.ActivateDisplay()
}

func deactivate(w http.ResponseWriter, r *http.Request){
	tinker.DeactivateDisplay()
}

func appendLine(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	tinker.AppendText(vars["text"])
}
