package main

import (
	"fmt"
	"bytes"
	"io/ioutil"
         "io"
	"net/http"
	"time"
         "encoding/json"
	"sync"
)

type Lang struct {
    Name string
    Url string
    Bytes  int64
    Time time.Duration
}



func Crawl(Print func(Lang), l Lang, wg *sync.WaitGroup){
  Print(l)
 defer wg.Done()
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
   var wt sync.WaitGroup
   wt.Add(3)
  t:=time.Now()
 python := Lang{Name:"Python",Url:"https://www.python.org/"}

 go Crawl(Print,python,&wt)


 ruby := Lang{Name:"Ruby",Url:"https://www.ruby-lang.org/en/"}

 go Crawl(Print,ruby,&wt)


 golang := Lang{Name:"Golang",Url:"https://golang.org/"}

 go Crawl(Print,golang,&wt)

 
	wt.Wait()
 elapsed:=time.Since(t)

  fmt.Println("Total Time",elapsed)
}

