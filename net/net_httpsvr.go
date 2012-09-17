package main

import (
    "net/http"
    "log"
)

func http_handler(w http.ResponseWriter, req *http.Request) {
    w.Header().Set("Content-Type", "text/plain")
    w.Write([]byte("This is an example server.\n"))
}

func main() {
    http.HandleFunc("/", http_handler)
    log.Printf("About to listen on 10080.\n")
    err := http.ListenAndServe(":10080", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}

