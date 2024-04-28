package main

import (
    "net"
    "fmt"
  )

func main(){
  _, err:= net.Dial("tcp","localhost:8080")
  if err!=nil{
    fmt.Println(err)
  }
}
