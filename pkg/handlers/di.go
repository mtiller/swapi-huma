package handlers

import (
	"context"
	"net/http"

	"github.com/samber/do"
)

type contextKey string

const doContextKey contextKey = "swapi-middleware-injector"

func Inject(injector *do.Injector) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			r = r.WithContext(context.WithValue(r.Context(), doContextKey, injector))

			next.ServeHTTP(w, r)
		})
	}
}
