package main

import (
    "fmt"
    "net"
    "log"
)

func main() {
    l, err := net.Listen("unix", "/tmp/go_sample.sock")
    if err != nil {
        log.Fatal(err)
    }
    // Loop to wait for a connection
    conn, err := l.Accept()
    if err != nil {
        log.Fatal(err)
    }
    go func(c net.Conn) {
        // serve connection
        c.Write("dsf\n")
        c.Close()
    }(conn)
}
