package main

import (
    "net"
    "fmt"
    "bufio"
    "os"
  )

func main(){
  conn, err:= net.Dial("tcp","localhost:8080")
  if err!=nil{
    fmt.Println(err)
  }
  defer conn.Close()
    fmt.Println("Enter the message to send:")
    scanner:=bufio.NewScanner(os.Stdin)
    scanner.Scan()
    msg:=scanner.Text()
    _,err= conn.Write([]byte(msg))
    if err!=nil{
      fmt.Println(err)
    }

     buf := make([]byte, 1024)
    _, err= conn.Read(buf)
    if err!=nil{
      fmt.Println(err)
    }
    fmt.Println(string(buf))

    
}
