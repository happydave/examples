package main

import (
	"flag"
	"log"
	"os"
	"runtime"
)

var version string = "0"

func main() {
	log.SetOutput(os.Stdout)
	log.SetFlags(0)

	logVerbose := flag.Bool("v", false, "Verbose logging")
	flag.Parse()

	if *logVerbose {
		log.Printf("Verbose logging enabled")
	}

	log.Printf("Version: %s built with %s", version, runtime.Version())
}
