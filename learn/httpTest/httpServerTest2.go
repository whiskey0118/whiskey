package main

//go程序设计语言 7.7实例 创建http服务器
import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	db := databases{"shose": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)

	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

type databases map[string]dollars
type dollars int

func (db databases) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %d\n", item, price)
	}
	//fmt.Fprintf(w,"hahaha")
}

func (db databases) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		http.Error(w, "no such item %q \n", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "%d\n", price)

}
