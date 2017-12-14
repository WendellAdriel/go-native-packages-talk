// Package server handles the HTTP Server behavior
package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/WendellAdriel/go-native-packages-talk/config"
)

// Server encapsulates the Application Server and the Application Logger
type Server struct {
	logger *log.Logger
	mux    *http.ServeMux
}

// NewServer creates a new Application Server with the given Logger
func NewServer(logger *log.Logger) *Server {
	httpServer := &Server{logger: logger, mux: http.NewServeMux()}

	httpServer.mux.Handle("/", withMiddlewares(http.HandlerFunc(httpServer.home), loggerMiddleware(logger)))
	httpServer.mux.Handle("/hello/", withMiddlewares(http.HandlerFunc(httpServer.hello), loggerMiddleware(logger)))

	return httpServer
}

// GracefulShutdown shutdowns gracefully the server using the Context Package
func GracefulShutdown(app *http.Server, logger *log.Logger) {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	logger.Println("Shutting down server...")
	if err := app.Shutdown(ctx); err != nil {
		logger.Fatal("Error while shutting down the server")
	}
	logger.Printf("Server stopped gracefully at %s\n", time.Now().UTC())
}

// DefaultLogger creates a default Logger for the Application
func DefaultLogger() *log.Logger {
	return log.New(os.Stdout, fmt.Sprintf("%s: ", config.AppName), 0)
}

func (s *Server) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Server", config.AppName)
	s.mux.ServeHTTP(res, req)
}
