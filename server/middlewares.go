// Package server handles the HTTP Server behavior
package server

import (
	"log"
	"net/http"
	"time"
)

type middleware func(http.Handler) http.Handler

func withMiddlewares(next http.Handler, middlewares ...middleware) http.Handler {
	for i := len(middlewares) - 1; i >= 0; i-- {
		next = middlewares[i](next)
	}

	return next
}

func loggerMiddleware(logger *log.Logger) middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			logger.Printf("%s - %s - %s\n", time.Now().UTC(), req.Method, req.RequestURI)
			defer logger.Printf("%s - %s - %s - Completed", time.Now().UTC(), req.Method, req.RequestURI)
			next.ServeHTTP(res, req)
		})
	}
}
