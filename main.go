package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// ArticleHandler is a function handler
func ArticleHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, "Category is %v\n", vars["category"])
	fmt.Fprintf(w, "ID is: %v\n", vars["id"])

}

func main() {

	// Creating a new router
	fmt.Println("Server is running ion port 8000...")

	r := mux.NewRouter()

	r.HandleFunc("/articles/{category}/{id:[0-9]+}", ArticleHandler)

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())

}
