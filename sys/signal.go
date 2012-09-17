package main

import (
    "fmt"
    "os"
    "os/signal"
    "time"
)

func main() {
    sigFlag := 0
    go func() {
        c := make(chan os.Signal, 1)
        signal.Notify(c, os.Interrupt)
        for {
            sig := <- c
            sigFlag ++
            if sigFlag > 1 {
                os.Exit(0)
            }
            fmt.Printf("CTRL-C: %v, once more within 1s to exit.\n", sig)
            time.AfterFunc(1e9, func() {
                sigFlag = 0
            })
        }
    }()

    //select {}// wait forever
    for {
        time.Sleep(1e9)
    }
}
