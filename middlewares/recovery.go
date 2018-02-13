package middlewares

import (
	"net/http"
	"log"
	"fmt"
	"github.com/shvetsovau/GoHttpServer/errors"
)

func RecoverHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("panic: %+v", err)
				httpSendError(w, err)
			}
		}()
		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

func httpSendError(w http.ResponseWriter, err interface{}) {
	w.Header().Set("Content-Type", "application/vnd.api+json")
	w.WriteHeader(errors.GetStatusCode(err, 500))
	fmt.Fprint(w, err)
}