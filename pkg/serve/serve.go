package serve

import (
	"log"
	"net/http"
	"os"
)

func ListenAndServe(handler http.Handler) {
	port := "8080"
	if s := os.Getenv("PORT"); s != "" {
		port = s
	}

	if err := http.ListenAndServe(":"+port, handler); err != nil {
		log.Fatalf("http.ListenAndServe: %v", err)
	}
}
