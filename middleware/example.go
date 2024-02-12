package middleware

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func ExampleMiddlewareFactory(test string) Middleware {
	return func(next httprouter.Handle) httprouter.Handle {
		return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
			w.Header().Set("X-TEST", test)
			next(w, r, ps)
		}
	}
}
