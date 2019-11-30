package main

import (
  "fmt"
	"flag"
  "log"
  "time"
	"net/http"

  "github.com/gorilla/mux"
  "github.com/stianeikeland/go-rpio/v4"
)

var port string
var bcm uint
var pin rpio.Pin

func init() {
  flag.StringVar(&port, "p", "443", "Port to listen on")
	flag.StringVar(&port, "port", "443", "Port to listen on")
  flag.UintVar(&bcm, "g", 4, "GPIO pin number")
	flag.UintVar(&bcm, "gpio", 4, "GPIO pin number")
}

func main() {
  defer rpio.Close()

	flag.Parse()

  if err := rpio.Open(); err != nil {
    log.Fatalln("Could not open rpio:", err)
  }

  pin = rpio.Pin(bcm)
  pin.Output()

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

  log.Printf("Listening on :%s and controlling GPIO %d", port, pin)
	if err := srv.ListenAndServeTLS("./certs/tls-cert.pem", "./certs/tls-key.pem"); err != http.ErrServerClosed {
		log.Fatalln("Could not start origin server:", err)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
  log.Println("home")
  http.ServeFile(w, r, "./static/index.html")
}

func pushHandler(w http.ResponseWriter, r *http.Request) {
  log.Println("push")

  go func() {
    pin.High()
    time.Sleep(500 * time.Millisecond)
    pin.Low()
  }()

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "ok")
}

func healthzHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "ok")
}
