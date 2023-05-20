package main3


import (
    "fmt"
    "os"
    "net/http"
    "net/http/httputil"
)


// Default listener port
var PORT string = "8080" 


// OS Variable with port overwrite
var PORTENV string = "HELLO_PORT"


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

