package main

import (
    "net"
    "fmt"
    "bufio"
    "os"
    "strings"
  )

func main(){
  while(true){
    fmt.Println("---------NEW SESSION HAS STARTED-----------")
    fmt.Println("POS TERMINAL")
    fmt.Println("You have requested to purchase a product")
    fmt.Println("Enter the price of the product:")
    scanner:=bufio.NewScanner(os.Stdin)
    scanner.Scan()
    msg:=scanner.Text()
    fmt.Sprintf("Total cost:$%d",msg)
    fmt.Println("enter y to continue:")
    scanner.Scan()
    if(scanner.Text()!="y"){
        fmt.Println("Transaction terminated terminated")
        continue
    }
    fmt.Println("Enter your account number:")
    scanner.Scan()
    accNo := scanner.Text()
    if(strings.Equals("") || len(accNo)!=11){
      fmt.Println("Invalid Account Number")
      fmt.Println("------------SESSION TERMINATED--------------")
      continue
    }
  }

    
}
func sendISOmsg(msg string){
  conn, err:= net.Dial("tcp","localhost:8080")
  if err!=nil{
    fmt.Println(err)
  }
  defer conn.Close()
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
