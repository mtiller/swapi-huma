package handlers

import (
	"context"
	"fmt"
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

func GetInjector(ctx context.Context) *do.Injector {
	injector := ctx.Value(doContextKey)
	if injector != nil {
		return injector.(*do.Injector)
	}

	panic(fmt.Errorf("No DI container found"))
}

func GetService[T any](ctx context.Context) T {
	injector := GetInjector(ctx)
	return do.MustInvoke[T](injector)
}
