package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type handleAllBody struct {
	WhoAmI  string              `json:"whoami"`
	Host    string              `json:"host"`
	Path    string              `json:"path"`
	Headers map[string][]string `json:"headers"`
}

func main() {
	listenAddress := os.Getenv("LISTEN_ADDRESS")

	mux := http.NewServeMux()

	mux.HandleFunc("GET /*", handleAll)

	log.Printf("http server %v listening at %v\n", os.Getenv("WHOAMI"), listenAddress)
	err := http.ListenAndServe(listenAddress, mux)
	if err != nil {
		panic(err)
	}
}

func handleAll(w http.ResponseWriter, r *http.Request) {
	log.Printf("got you %v\n", os.Getenv("WHOAMI"))
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")

	body := &handleAllBody{
		Host:    r.Host,
		Path:    r.URL.Path,
		WhoAmI:  os.Getenv("WHOAMI"),
		Headers: r.Header,
	}

	if err := json.NewEncoder(w).Encode(body); err != nil {
		panic(err)
	}
}
