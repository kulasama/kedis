package main

import (
        "fmt"
        "net"
       )   

func RequestHandler(conn net.Conn) {  
    defer conn.Close() 
    b:=[]byte("hello world")
    conn.Write(b)
}

func main() {
    fmt.Printf("server init\n") 
    listen, err := net.Listen("tcp","127.0.0.1:8080")  
    if err != nil {
        fmt.Printf("server fault",err)
        return 
    }                           
    for {
        conn, err := listen.Accept()  
        if err != nil {
            fmt.Printf("accept error",err)
            //break
        }
        go RequestHandler(conn)
        fmt.Printf("conected",conn.RemoteAddr(),"\n")  
    }           
}