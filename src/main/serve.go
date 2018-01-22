package main

import(
    "net/rpc"
    "server"
    "net"
    "log"
    "net/http"
)

func main() {
    arith := new(server.Arith)
    rpc.Register(arith)
    rpc.HandleHTTP()
    l, e := net.Listen("tcp", ":6666")
    if e != nil {
        log.Fatal("listen error:", e)
    }
    http.Serve(l, nil)
}
