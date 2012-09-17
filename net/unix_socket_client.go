package main

import (
    "net"
    "log"
    "fmt"
)

func main() {
    i := 0
    for {
        i ++;
        sock := "/tmp/go_sample.sock"
        c, err := net.Dial("unix", sock)
        if err != nil {
            log.Fatal(err)
        }
        var s string
        s = fmt.Sprintf("Hello Unix Socket: %d", i);
        _, err = c.Write([]byte(s))
        if err != nil {
            log.Fatalf("Write Error: %s", err)
        }

        b := make([]byte, 128)
        _, err = c.Read(b)
        if err != nil {
            log.Fatalf("Read Error: %s", err)
        }
        if i % 1000 == 0 {
            fmt.Printf("SERVER ANSWER: %s\n", b)
        }
        c.Close()
    }
}
