package main

import (
       "testing"
         "errors"
          "time"
           "fmt"
           "net"
         )



type Conn interface {

	        Read(b []byte) (n int, err error)   
   		Write(b []byte) (n int, err error) 
   		Close() error
   		LocalAddr() net.Addr   	
   		RemoteAddr() net.Addr   	
   		SetDeadline(t time.Time) error
  		SetReadDeadline(t time.Time) error
   	        SetWriteDeadline(t time.Time) error
   	}
type conn struct{
 Check int
}

   func (c *conn) Read(b []byte) (int, error) {
           err := errors.New("read function error")
           return c.Check,err
   	}
   	

   func (c *conn) Write(b []byte) (int, error) {
           err := errors.New("write function error")
           return c.Check,err
   	}
  
   func (c *conn) Close() error {
         return errors.New("close function error")
   	}
   	
   
   func (c *conn) LocalAddr() net.Addr {
          return nil
  	}
   
   func (c *conn) RemoteAddr() net.Addr {
           return nil
  	}
   	

   func (c *conn) SetDeadline(t time.Time) error {
            return errors.New("setdeadline function error")
   	}
  
   func (c *conn) SetReadDeadline(t time.Time) error {
            return errors.New("setReadDeadline function error")
   	}
   	

   func (c *conn) SetWriteDeadline(t time.Time) error {
           return errors.New("SetWriteDeadline function error")
   	}
func TestConnection(t *testing.T){
   myCon:=&conn{Check:0}
   myChan:=make(chan error,1)
   HandleRequest(myCon,myChan)
   err:=<-myChan
   if err==nil{
       t.Fatalf("Connection error found",errors.New("This is the test"))
   }
}
func TestRightServers(t *testing.T){
   
 myChanel:=make(chan error,1)

  go func(port string,myChanel chan error) {

    err:=Server(port)      
    myChanel<-err
    }(":9000",myChanel)
  select{
          case <-time.After(3 * time.Second): 
                    fmt.Println("timeout")
          case err:=<-myChanel:
               if err==nil{
               t.Fatalf("server error found",errors.New("This is the test"))
        }
                       
   }
}

func TestWrongServers(t *testing.T){
   
 myChanel:=make(chan error,1)

  go func(port string,myChanel chan error) {

    err:=Server(port)      
    myChanel<-err
    }(":abcd",myChanel)
  select{
          case <-time.After(3 * time.Second): 
                    fmt.Println("timeout")
          case err:=<-myChanel:
               if err==nil{
               t.Fatalf("server error found",errors.New("This is the test"))
        }
                       
   }
  
  }

