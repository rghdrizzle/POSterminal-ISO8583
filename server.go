package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"strconv"

	"github.com/moov-io/iso8583"
	"github.com/moov-io/iso8583/field"
	"github.com/moov-io/iso8583/specs"
)

type RequestIsoFields struct{
  MTI                  *field.String `index:"0"`
  ProcessingCode       *field.Numeric `index:"3"`
	TransmissionDateTime *field.String `index:"7"`
	STAN                 *field.String `index:"11"`
	ProductAmount        *field.String `index:"8"`
  ResponseCode         *field.String `index:"39"`
  AccountNumber        *field.String `index:"2"`
}
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
  data := &RequestIsoFields{}
  var responseCode string
  var resMessage []byte
  var customerBalance int
  defer conn.Close()
  for{
    buf := make([]byte, 1024)
    _, err:= conn.Read(buf)
    
    if err!=nil{
      if opErr, ok := err.(*net.OpError); ok && opErr.Op == "read" {
				fmt.Println("Connection closed by client:", err)
			}
      fmt.Println(err)
      return
    }
    Reqisomessage:= string(buf)
    isomessage := iso8583.NewMessage(specs.Spec87ASCII)
    err=isomessage.Unpack(buf)
    if err!=nil{
      fmt.Println(err)
    }
    err = isomessage.Unmarshal(data)
    if err!=nil{
      fmt.Println(err)
      }
    fmt.Println("ISO message received from client:"+Reqisomessage)
    fmt.Println(isomessage.GetMTI())
      amount,_:= strconv.Atoi(data.ProductAmount.Value())
      switch(data.MTI.Value()){
      case "0200":
        fmt.Println("Executing Financial transaction-Balance enquiry")
        customerBalance = GetCustomerBalance()
        if customerBalance>=amount{
          responseCode="00"
        }else{
          responseCode = "09"
        }
        responseisomessage := iso8583.NewMessage(specs.Spec87ASCII)
        responseisomessage.MTI(data.MTI.Value())
        accNo ,_:= isomessage.GetString(2)
        proCode ,_ :=isomessage.GetString(3)
        err := responseisomessage.Field(2,accNo)
        handleError(err)
        err = responseisomessage.Field(3,proCode)
        handleError(err)
        err = responseisomessage.Field(7,data.TransmissionDateTime.Value())
        handleError(err)
        err = responseisomessage.Field(8,data.ProductAmount.Value())
        handleError(err)
        err= responseisomessage.Field(11,data.STAN.Value())
        handleError(err)
        err= responseisomessage.Field(39,responseCode)
        handleError(err)

        resMessage, err = responseisomessage.Pack()
        handleError(err)
      default:
        responseCode="09"
    }

    
    _,err = conn.Write([]byte(resMessage))
    if err!=nil{
      fmt.Println(err)
    }
    fmt.Println("Iso response sent"+string(resMessage))

  }
}
func GetCustomerBalance() int{
  balance := rand.Intn(100000)

  return balance
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

