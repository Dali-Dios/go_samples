package main

import (
    //"net"
    "log"
    "bytes"
    //"time"
    "encoding/binary"
)

func main() {
    /* conn, err := net.Dial("tcp", "www.baidu.com:80")
    if (err != nil) {
        fmt.Println("connect failed")
        return
    }

    fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
    reader := bufio.NewReader(conn)
    for {
        line,_,err := reader.ReadLine();
        if err != nil {
            break
        }
        fmt.Println(string(line));
    }*/
    //conn, err := net.DialTimeout("udp", "210.72.145.44:37", 1e9)
    // err != nil {
    //    log.Fatalf("Dial Error: %s", err)
    //}
    //conn.SetDeadline(1e9)
    buf := []byte{0, 0, 0, 65, 0, 0, 0, 65}
    //var buf2 [4]byte
    //conn.Write(buf)

    //b := make([]byte, 128)
    //_, err = conn.Read(b)
    //if err != nil {
    //    log.Fatalf("Read Error: %s", err)
    //}
    var ts uint32
    newBuf := bytes.NewBuffer(buf)
    err := binary.Read(newBuf, binary.BigEndian, &ts)
    if err != nil {
        log.Fatalf("binary.Read Error: %s", err)
    }
    log.Printf("time=%d\n", ts)
    return
    
    //var pi float64
    //b := []byte{0x18, 0x2d, 0x44, 0x54, 0xfb, 0x21, 0x09, 0x40}
    //buf := bytes.
}
