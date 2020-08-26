package main

//
////go程序设计语言 7.7实例 创建http服务器
//import (
//	"fmt"
//	"log"
//	"net/http"
//)
//
//func main() {
//	db := database{"shose":50, "socks":5}
//	err := http.ListenAndServe("localhost:8080",db)
//	if err != nil {
//		log.Fatal(err)
//	}
//}
//
//type dollars float32
//
//func (d dollars) String() string{
//	return fmt.Sprintf("$%.2f",d)
//}
//
////database类型实现了Handler接口
//type database map[string]dollars
//
//func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request)  {
//	switch req.URL.Path {
//	case "/list":
//		for item,price := range db{
//			fmt.Fprintf(w,"%s\n",item,price)
//		}
//	case "/price":
//		item := req.URL.Query().Get("item")
//		price,ok := db[item]
//		if !ok{
//			w.WriteHeader(http.StatusNotFound) //404
//			//http.Error(w, "msg",http.StatusNotFound)
//			fmt.Fprintf(w,"no such item:%q\n",item)
//			return
//		}
//		fmt.Fprintf(w,"%s\n",price)
//	default:
//		w.WriteHeader(http.StatusNotFound)
//		fmt.Fprintf(w,"no such page:%s\n",req.URL)
//	}
//
//}
