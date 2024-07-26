package main

import (
	"10/limit"
	"flag"
	"fmt"
	"log"
	"net/http"
)

var rate, burst, port int

func init() {
	flag.IntVar(&port, "port", 4000, "port to listen on")
	flag.IntVar(&rate, "rate", 1000, "rate (tokens per second) per ip")
	flag.IntVar(&burst, "burst", 100, "maximum burst size per ip")

	flag.Parse()
}

func main() {
	cfg := limit.NewConfig(rate, burst)
	limit := limit.MiddlewareFn(cfg)

	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)

	log.Printf("Listening on :%d \n", port)
	http.ListenAndServe(fmt.Sprint(":", port), limit(mux))
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}
