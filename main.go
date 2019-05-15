package main

import (
	"net/http"

	"github.com/apstndb/go-httptoolbox/pkg/handlers"
	"github.com/apstndb/go-httptoolbox/pkg/serve"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/read", handlers.ReadContent)
	mux.HandleFunc("/exec", handlers.Exec)
	mux.HandleFunc("/dmesg", handlers.ExecDmesg)
	mux.HandleFunc("/envs", handlers.GetEnvs)
	mux.HandleFunc("/headers", handlers.GetHeaders)
	serve.ListenAndServe(mux)
}
