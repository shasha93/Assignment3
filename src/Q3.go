package main

import (
	"fmt"
	"bytes"
	"io/ioutil"
         "io"
	"net/http"
	"time"
         "encoding/json"
	
)

type Lang struct {
    Name string
    Url string
    Bytes  int64
    Time time.Duration
}



func Crawl(Print func(Lang), l Lang){
  Print(l)
}

func Print(l Lang){
t:=time.Now() 

 res,err:=http.Get(l.Url)
if err!=nil{
  fmt.Println(err)
  }
  defer res.Body.Close()

  body,_:=ioutil.ReadAll(res.Body)
  l.Bytes,_ = io.Copy(ioutil.Discard,bytes.NewReader(body))
 l.Time = time.Since(t)
  fmt.Println("\nGO formatted")
 fmt.Printf("%+v \n",l)
 fmt.Println("\nJson formatted")
 jFormat, _ := json.Marshal(l)
 fmt.Printf(string(jFormat)+"\n")
}
func main() {

 
 python := Lang{Name:"Python",Url:"https://www.python.org/"}

 Crawl(Print,python)


 ruby := Lang{Name:"Ruby",Url:"https://www.ruby-lang.org/en/"}

 Crawl(Print,ruby)


 golang := Lang{Name:"Golang",Url:"https://golang.org/"}

 Crawl(Print,golang)

}

