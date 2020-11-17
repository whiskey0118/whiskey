package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	log.SetPrefix("<Read message>")
	//设置日志的格式
	log.SetFlags(log.LstdFlags | log.Ldate | log.Ltime | log.Lshortfile)
	log.Printf("this is a test")
	http.HandleFunc("/", testFun)
	http.ListenAndServe("127.0.0.1:8080", nil)

}

func testFun(w http.ResponseWriter, r *http.Request) {
	//var res []byte
	//_,err := r.Body.Read(res)
	//defer r.Body.Close()
	//if err != nil {
	//	log.Println("body read err: ",err)
	//}
	//fmt.Printf("%s",res)
	//
	//err = r.ParseForm()
	//fmt.Println(r.Form["test"][0])
	//if err != nil {
	//	fmt.Println(err)
	//}
	s, _ := ioutil.ReadAll(r.Body)
	fmt.Printf("%s", s)
	fmt.Fprintf(w, "%s", "this test")

}
