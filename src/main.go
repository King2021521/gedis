package main

import (
	"tcp"
	"fmt"
	"template"
	"net"
)

func main() {
	conn:=getConn()
	result:=template.Sdiff(conn,"set1","set2","set3")
	fmt.Println(result,"111")
}

func getConn() *net.TCPConn{
	var config = tcp.ConnConfig{"127.0.0.1:6379","root"}
	pool,err:=tcp.NewConnPool(1, config)
	if err!=nil{
		fmt.Println(err)
	}

	conn,_:=tcp.GetConn(pool)
	return conn
}
