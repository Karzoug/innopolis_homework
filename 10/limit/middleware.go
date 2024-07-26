package limit

import (
	"log"
	"net"
	"net/http"
	"strings"
)

var xForwardedFor = http.CanonicalHeaderKey("X-Forwarded-For")
var xRealIP = http.CanonicalHeaderKey("X-Real-IP")

func MiddlewareFn(cfg config) func(http.Handler) http.Handler {
	l := newLimiter(cfg)
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if rip := realIP(r); rip == "" {
				r.RemoteAddr = rip
			}

			ip, _, err := net.SplitHostPort(r.RemoteAddr)
			if err != nil {
				log.Print(err.Error())
				http.Error(w,
					http.StatusText(http.StatusInternalServerError),
					http.StatusInternalServerError)
				return
			}
			limiter := l.getVisitor(ip)
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

func realIP(r *http.Request) string {
	var ip string

	if xff := r.Header.Get(xForwardedFor); xff != "" {
		i := strings.Index(xff, ", ")
		if i == -1 {
			i = len(xff)
		}
		ip = xff[:i]
	} else if xrip := r.Header.Get(xRealIP); xrip != "" {
		ip = xrip
	}

	return ip
}
