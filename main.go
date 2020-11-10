package main

import (
	"flag"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"os"
	"whiskey/web"
)

var (
	addr = flag.String("addr", "127.0.0.1:8080", "http")
)

func init() {
	dbLink := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", "root", "Maozedong@123", "120.78.160.193", "8888", "whiskey")
	err := orm.RegisterDriver("mysql", orm.DRMySQL)
	if err != nil {
		log.Fatal("err", err)
		os.Exit(2)
		return
	}
	orm.RegisterDataBase("default", "mysql", dbLink)
}

func main() {
	flag.Parse()
	http.HandleFunc("/postDevice", web.PostDevice)
	http.HandleFunc("/read", web.ReadMessage)
	err := http.ListenAndServe("localhost:8089", nil)
	if err != nil {
		log.Fatal("listen 8080 err:", err)
		return
	}
}
