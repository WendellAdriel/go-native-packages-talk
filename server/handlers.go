// Package server handles the HTTP Server behavior
package server

import (
	"fmt"
	"net/http"
	"strings"
)

func (s *Server) home(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Hello, world."))
}

func (s *Server) hello(res http.ResponseWriter, req *http.Request) {
	greeting := fmt.Sprintf("Hello, %s.", strings.TrimPrefix(req.URL.Path, "/hello/"))
	res.Write([]byte(greeting))
}
