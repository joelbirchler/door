package main

import (
  "fmt"
	"flag"
  "log"
  "time"
	"net/http"

  "github.com/gorilla/mux"
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

	r := mux.NewRouter()

  r.Path("/").HandlerFunc(homeHandler)
  r.Path("/healthz").HandlerFunc(healthzHandler)
  r.Path("/push").Methods("PUT").HandlerFunc(pushHandler)
  r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

  srv := &http.Server{
    Handler: r,
    Addr: "0.0.0.0:" + port,
    WriteTimeout: 10 * time.Second,
    ReadTimeout: 10 * time.Second,
  }

	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalln("Could not start origin server:", err)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
  http.ServeFile(w, r, "./static/index.html")
}

func pushHandler(w http.ResponseWriter, r *http.Request) {
  log.Println("push")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "ok")
}

func healthzHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "ok")
}
