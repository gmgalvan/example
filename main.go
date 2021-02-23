package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
)

// main func x
func main() {
	port := "8080"
	if fromEnv := os.Getenv("PORT"); fromEnv != "" {
		port = fromEnv
	}

	server := http.NewServeMux()
	server.HandleFunc("/", hello)

	log.Printf("Server listening on port %s", port)
	err := http.ListenAndServe(":"+port, server)
	log.Fatal(err)
}

func hello(w http.ResponseWriter, r *http.RehosPushEventt, _ := os.Hostname()
	quest) {
	log.Printf("Serving request: %s", r.URL.Path)
	fmt.Fprintf(w, "Hello World\n")
	fmt.Fprintf(w, "Version: 0.0.1\n")
	addrs, _ := net.InterfaceAddrs()
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				fmt.Fprintf(w, "IP Address: %s\n", ipnet.IP.String())
			}
		}
	}
}
