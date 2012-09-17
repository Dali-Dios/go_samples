package main

import (
    "net"
    "log"
    "fmt"
    "os"
)

func main() {
    sock := "/tmp/go_sample.sock"
    os.Remove(sock)
    l, err := net.Listen("unix", sock)
    if err != nil {
        log.Fatalf("Listen Error: %s", err)
    }
    // Loop to wait for a connection
    for {
        conn, err := l.Accept()
        //conn.SetReadTimeout(5e9)
        if err != nil {
            log.Fatalf("Accept Error: %s", err)
        }
        go func(c net.Conn) {
            // serve connection
            b := make([]byte, 128)
            n, err := c.Read(b)

            // If client not sending anything after net.Dial, will EOF here
            if err != nil && fmt.Sprintf("%s", err) != "EOF" {
                log.Fatalf("Read Error: %s", err)
            }
            s := fmt.Sprintf("RECV [%d bytes]: %s\n", n, string(b))
            //fmt.Print(s)
            c.Write([]byte(s))
            c.Close()
        }(conn)
    }
}
