// Simple application to use as an example on a meetup
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/WendellAdriel/go-native-packages-talk/cmd"
	"github.com/WendellAdriel/go-native-packages-talk/config"
)

var (
	version bool
	help    bool
	start   bool
	port    string
)

func init() {
	flag.BoolVar(&version, "version", false, "Prints application version and info")
	flag.BoolVar(&version, "v", false, "Prints application version and info (shorthand)")
	flag.BoolVar(&help, "help", false, "Shows the available flags")
	flag.BoolVar(&help, "h", false, "Shows the available flags (shorthand)")
	flag.BoolVar(&start, "start", false, "Starts the HTTP server")
	flag.BoolVar(&start, "s", false, "Starts the HTTP server (shorthand)")
	flag.StringVar(&port, "port", config.DefaultPort, "The port to init the HTTP Server")
	flag.StringVar(&port, "p", config.DefaultPort, "The port to init the HTTP Server (shorthand)")

	flag.Usage = func() {
		flag.PrintDefaults()
	}

	if len(os.Args) < 2 {
		fmt.Println("You need to provide at least one flag.")
		flag.Usage()
		os.Exit(1)
	}

	flag.Parse()
}

func main() {
	if version {
		fmt.Printf("%s - %s - Created by: %s", config.AppName, config.AppVersion, config.AppAuthor)
		os.Exit(0)
	}

	if help {
		fmt.Printf("%s - %s - Created by: %s\n---\nFLAGS AVAILABLE:\n\n", config.AppName, config.AppVersion, config.AppAuthor)
		flag.Usage()
		os.Exit(0)
	}

	if start {
		cmd.StartServer(port)
	}
}
