package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	r "golang.org/x/time/rate"
)

var rate, burst, port int

func init() {
	flag.IntVar(&port, "port", 4000, "port to listen on")
	flag.IntVar(&rate, "rate", 1000, "rate (tokens per second)")
	flag.IntVar(&burst, "burst", 100, "maximum burst size")

	flag.Parse()
}

func limitMiddlewareFn(limiter *r.Limiter) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if !limiter.Allow() {
				http.Error(w,
					http.StatusText(http.StatusTooManyRequests),
					http.StatusTooManyRequests)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}

func main() {
	limiter := r.NewLimiter(r.Limit(rate), burst)
	limit := limitMiddlewareFn(limiter)

	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)

	log.Printf("Listening on :%d \n", port)
	http.ListenAndServe(fmt.Sprint(":", port), limit(mux))
}
