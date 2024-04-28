package main

import (
  "net"
  "log"
  "fmt"
)

func StartServer(){
    ln , err := net.Listen("tcp",":8080")
    if err !=nil{
      log.Fatal(err)
    }
    for {
      _, err := ln.Accept();
      if err!=nil{
        log.Fatal(err)
      }
      fmt.Println("Connected")
      
      // run go routine for handling connection
    }

}
