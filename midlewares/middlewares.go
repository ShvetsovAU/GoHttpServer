package midlewares

import (
	"net/http"
	"time"
	"log"
	"sync"
	"github.com/shvetsovau/GoHttpServer/errors"
)

var once sync.Once

func ContentTypeHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Content-Type") != "application/vnd.api+json" {
			panic(errors.NewHttp(415, "Content-Type header must be set to: 'application/vnd.api+json'"))
		}

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

func DetailedLoggingHandler(next http.Handler) http.Handler {
	onceBody := func() {
		log.Println("--------------------------------------------------------")
	}
	fn := func(w http.ResponseWriter, r *http.Request) {
		once.Do(onceBody)
		t1 := time.Now()
		next.ServeHTTP(w, r)
		t2 := time.Now()

		log.Printf("[%v] %q %v\n", r.Method, r.URL, t2.Sub(t1))
		for k, v := range r.Header {
			log.Printf("Header[%q] = %q\n", k, v)
		}
		log.Printf("Host = %q\n", r.Host)
		log.Printf("RemoteAddr = %q\n", r.RemoteAddr)
		if err := r.ParseForm(); err != nil {
			log.Print(err)
		}
		for k, v := range r.Form {
			log.Printf("Form[%q] = %q\n", k, v)
		}
		log.Println("--------------------------------------------------------")
	}
	return http.HandlerFunc(fn)
}

func LoggingHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		t1 := time.Now()
		next.ServeHTTP(w, r)
		t2 := time.Now()
		log.Printf("[%v] %q %v\n", r.Method, r.URL, t2.Sub(t1))
	}

	return http.HandlerFunc(fn)
}