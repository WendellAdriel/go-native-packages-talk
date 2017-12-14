// Package cmd has all the logic of the CLI flags of the application
package cmd

import (
	"fmt"
	"log"
	"net/http"

	"github.com/WendellAdriel/go-native-packages-talk/server"
)

// StartServer starts the HTTP Server
func StartServer(port string) {
	app, logger := setupServer(port)

	go func() {
		logger.Printf("Starting server at http://0.0.0.0:%s", port)

		if err := app.ListenAndServe(); err != nil {
			logger.Fatal(err)
		}
	}()

	server.GracefulShutdown(app, logger)
}

func setupServer(port string) (*http.Server, *log.Logger) {
	logger := server.DefaultLogger()
	httpServer := &http.Server{Addr: fmt.Sprintf(":%s", port), Handler: server.NewServer(logger)}

	return httpServer, logger
}
