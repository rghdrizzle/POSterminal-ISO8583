package main

import (
    "net"
    "fmt"
    "bufio"
    "os"
    "log"
    "github.com/moov-io/iso8583"
    "github.com/moov-io/iso8583/specs"
  )
func main(){
  for{
    fmt.Println("---------NEW SESSION HAS STARTED-----------")
    fmt.Println("POS TERMINAL")
    fmt.Println("You have requested to purchase a product")
    fmt.Println("Enter the price of the product:(Please provide the cost in digits of 8, forexample if cost =100 then type 00000100)")
    scanner:=bufio.NewScanner(os.Stdin)
    scanner.Scan()
    productCost:=scanner.Text()
    fmt.Sprintf("Total cost:$%s",productCost)
    fmt.Println("enter y to continue:")
    scanner.Scan()
    if(scanner.Text()!="y"){
        fmt.Println("Transaction terminated terminated")
        continue
    }
    fmt.Println("Enter your account number:")
    scanner.Scan()
    accNo := scanner.Text()
    if(accNo=="" || len(accNo)!=11){
      fmt.Println("Invalid Account Number")
      fmt.Println("------------SESSION TERMINATED--------------")
      continue
    }
    customerValidity := IsTransactionValid(accNo,productCost)
    if(!customerValidity){
      fmt.Println("Insufficient funds to complete the Transaction")
      fmt.Println("---------------SESSION TERMINATED------------------")
      continue
    }
    fmt.Println("Customer is eligible to make Transaction")
    fmt.Println("Debitting from customer bank")
    fmt.Println("Transaction successfull")

  }

    
}
func IsTransactionValid(accNo string,productCost string)bool{
  MTI := "0200" // financial message according the iso8583 protocol 
  ProCode := "001000"
  TransDate := "0601021504"
  STAN := "000456"
  isomessage := iso8583.NewMessage(specs.Spec87ASCII)

  isomessage.MTI(MTI)
  err := isomessage.Field(2,accNo)
  handleError(err)
  err = isomessage.Field(3,ProCode)
  handleError(err)
  err = isomessage.Field(7,TransDate)
  handleError(err)
  err = isomessage.Field(8,productCost)
  handleError(err)
  err= isomessage.Field(11,STAN)
  handleError(err)
  requestMessage, err := isomessage.Pack()
  handleError(err)

  fmt.Println(string(requestMessage))
  iso8583.Describe(isomessage,os.Stdout)
  sendISOmsg(string(requestMessage))
  return true // placeholder for now

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

func handleError(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}
