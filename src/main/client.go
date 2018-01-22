package main

import (
    "server"
    "net/rpc"
    "log"
    "fmt"
)

func main() {
    client, err := rpc.DialHTTP("tcp", "localhost" + ":6666")
    if err != nil {
        log.Fatal("dialing:", err)
    }
    args := &server.Args{7, 8}
    var reply int
    err = client.Call("Arith.Multiply", args, &reply)
    if err != nil {
        log.Fatal("arith error:", err)
    }
    fmt.Printf("Arith: %d*%d=%d\n", args.A, args.B, reply)

    divargs := &server.Args{10, 4}
    var q server.Quotient
    err = client.Call("Arith.Divide", divargs, &q)
    fmt.Printf("Arith: %d/%d=%d\n", divargs.A, divargs.B, q.Quo)
    fmt.Printf("Arith: %d%%%d=%d\n", divargs.A, divargs.B, q.Rem)
}
