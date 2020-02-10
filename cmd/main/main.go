package main

import (
	"log"
	"net/http"
	"os"

	"github.com/apstndb/go-httptoolbox"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/read", httptoolbox.ReadContent)
	mux.HandleFunc("/exec", httptoolbox.Exec)
	mux.HandleFunc("/dmesg", httptoolbox.ExecDmesg)
	mux.HandleFunc("/envs", httptoolbox.GetEnvs)
	mux.HandleFunc("/headers", httptoolbox.GetHeaders)
	mux.HandleFunc("/write_headers", httptoolbox.WriteHeaders)
	mux.HandleFunc("/write_envs", httptoolbox.WriteEnvs)
	mux.HandleFunc("/tokeninfo", httptoolbox.TokenInfo)
	mux.HandleFunc("/email", httptoolbox.Email)
	mux.HandleFunc("/metadata", httptoolbox.Metadata)
	mux.HandleFunc("/", httptoolbox.DumpRequest)
	listenAndServe(mux)
}

func listenAndServe(handler http.Handler) {
	port := "8080"
	if s := os.Getenv("PORT"); s != "" {
		port = s
	}

	if err := http.ListenAndServe(":"+port, handler); err != nil {
		log.Fatalf("http.listenAndServe: %v", err)
	}
}
