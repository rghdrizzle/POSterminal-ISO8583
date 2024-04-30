package main

import (
  "net"
  "log"
  "fmt"
  "bufio"
  "os"
)

func StartServer(){
    ln , err := net.Listen("tcp",":8080")
    if err !=nil{
      log.Fatal(err)
    }
    for {
      conn, err := ln.Accept();
      if err!=nil{
        log.Fatal(err)
      }
      fmt.Println("Connected")
      
      go handleConnection(conn)
    }

}

func handleConnection(conn net.Conn){
  defer conn.Close()
  for{
    buf := make([]byte, 1024)
    _, err:= conn.Read(buf)
    if err!=nil{
      fmt.Println(err)
    }
    fmt.Println("Message received from client:"+string(buf))

    fmt.Println("Enter the message to send:")
    scanner:=bufio.NewScanner(os.Stdin)
    scanner.Scan()
    msg:=scanner.Text()
    _,err = conn.Write([]byte(msg))
    if err!=nil{
      fmt.Println(err)
    }

  }
}
