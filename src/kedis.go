package main

import (
        "fmt"
        "net"
       )   

func Serve(c net.Conn) {
}

func main() {
    fmt.Printf("hello world\n") 
    listen, err := net.Listen("tcp","127.0.0.1:8080")  
    if err != nil {
        fmt.Printf("server fault",err)
        return 
    }
    conn, err := listen.Accept()
    fmt.Printf("conected",conn,err,"\n")
}