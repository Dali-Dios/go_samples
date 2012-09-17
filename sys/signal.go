package main

import (
    "fmt"
    "os"
    "os/signal"
    "time"
)

func main() {

    go func() {
        c := make(chan os.Signal, 1)
        signal.Notify(c, os.Interrupt)
        for {
            sig := <- c
            fmt.Printf("CTRL-C: %v\n", sig)
        }
    }()

    //select {}// wait forever
    for {
        time.Sleep(1e9)
    }
}
