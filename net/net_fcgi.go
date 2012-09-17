package main

import (
    "fmt"
    "net"
    "net/http"
    "net/http/fcgi"
)

type FastCGIServer struct{}
func (s FastCGIServer) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    w.Header().Set("Content-Type", "text/html")
    str := fmt.Sprintf("<b>RequestURI</b>:%s<br />\n<b>User-Agent</b>:%s\n",
        req.URL.String(), req.UserAgent())
    w.Write([]byte(str))
}

func main() {
    s := new(FastCGIServer)
    l,_ := net.Listen("tcp", "127.0.0.1:9000")
    fcgi.Serve(l, s)
}
