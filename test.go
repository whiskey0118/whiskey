package main

import "github.com/astaxie/beego/orm"

func main() {
	maxIdle := 30
	maxConn := 30
	orm.RegisterDataBase("default", "mysql", "root:root@/orm_test?charset=utf8", maxIdle, maxConn)

}
