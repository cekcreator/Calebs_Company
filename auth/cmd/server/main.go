package main

import (
	"log"
	"net/http"
	"os"

	"github.com/calebs-company/auth/internal/handlers"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8001"
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/auth/token", handlers.IssueToken)
	mux.HandleFunc("/auth/verify", handlers.VerifyToken)
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"status":"ok","service":"auth"}`))
	})

	log.Printf("auth service listening on :%s", port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatal(err)
	}
}
