package main

import(
      "net"
       "fmt"
       "strings"
       
      )

func ChatResponse(msg string)string{
  var reqBuf string 
  switch strings.ToLower(msg) {
    case "hello": reqBuf="hi"
    case "name": reqBuf="Chitty"
    case "goodbye" :reqBuf="Bye"
    default: reqBuf="I am Busy"

    } 
  return reqBuf
}

func HandleRequest(conn net.Conn,myChan chan error){  //,myChan chan error
 for{
  buf := make([]byte, 1024)

  totLen, err := conn.Read(buf)
  if err != nil {
    myChan<-err
    return 
  }
  msg:=string(buf[:totLen-2])
  fmt.Print("Message Received:", msg+"\n")
 
   reqBuf:=ChatResponse(msg)

  _,err= conn.Write([]byte("new message :"+reqBuf+"\n"))
    if err != nil {
    myChan<-err
    return 
  }
  }
  myChan<-nil
  conn.Close()
}
func Server(port string)error{
 ln,err:= net.Listen("tcp",port)
  if err!=nil{
   return err
}
  defer ln.Close()
    myChan:=make(chan error)
 for{
   con,err:=ln.Accept()
   if err!=nil{
      return err
  }
    
   go HandleRequest(con,myChan) 
  err=<-myChan
 if err!=nil{
   panic(err)
   return err
  
 }
}
 return nil
}
