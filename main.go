package main

import (
  "fmt"
	"flag"
  "log"
	"net/http"
)

var port string
var pin uint

func init() {
  flag.StringVar(&port, "p", "8080", "Port to listen on")
	flag.StringVar(&port, "port", "8080", "Port to listen on")
  flag.UintVar(&pin, "g", 7, "GPIO pin number")
	flag.UintVar(&pin, "gpio", 7, "GPIO pin number")
}

func main() {
	flag.Parse()
	log.Printf("Listening on :%s and controlling GPIO %d", port, pin)

	http.HandleFunc("/healthz", healthzHandler)

	if err := http.ListenAndServe(":"+port, nil); err != http.ErrServerClosed {
		log.Fatalln("Could not start origin server:", err)
	}
}

func healthzHandler(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "ok")
}
