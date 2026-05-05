package main

import (
	"log"
	"net/http"
	"os"

	"github.com/calebs-company/realtime/internal/handlers"
	"github.com/calebs-company/realtime/internal/hub"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "9000"
	}

	h := hub.New()
	go h.Run()

	mux := http.NewServeMux()
	mux.Handle("/ws", handlers.NewWSHandler(h))
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"status":"ok","service":"realtime"}`))
	})

	log.Printf("realtime service listening on :%s", port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatal(err)
	}
}
