package go_http_hello_world

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

// Handler for the http server
func HelloServer(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		fmt.Fprintf(w, "Web Server Alive!\n")

		// Get all request data and echo back
		requestDump, err := httputil.DumpRequest(r, true)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(requestDump))
	} else {
		fmt.Fprintf(w, "Path: '%s'\n", r.URL.Path[1:])
	}
}
