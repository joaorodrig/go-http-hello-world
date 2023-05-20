package main

import (
	"fmt"
	"net/http"
	"os"
)

// Default listener port
var PORT string = "8080"

// OS Variable with port overwrite
var PORTENV string = "HELLO_PORT"

func main() {
	// Default port
	var listenerPort = PORT

	// Check if we have a port variable defined
	newPort, portOverwrite := os.LookupEnv(PORTENV)
	if portOverwrite {
		listenerPort = newPort
	}

	// Configure http process with handler and start
	fmt.Printf("Starting http server listening on port %s\n", listenerPort)
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":"+listenerPort, nil)
}
