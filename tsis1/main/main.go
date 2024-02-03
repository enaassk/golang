package main

import (
	"net/http"
	handlers "tsis1/enh/handlers"

	"github.com/gorilla/mux"
)

func main() {
    r := mux.NewRouter()

    r.HandleFunc("/api/members", handlers.GetMembers).Methods("GET")
    r.HandleFunc("/api/member/{name}", handlers.GetMember).Methods("GET")
    r.HandleFunc("/health", healthcheck.HealthCheck).Methods("GET")

    log.Println("Server is running on :8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}


