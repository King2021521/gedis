package main

import (
	"tcp"
	"fmt"
	"template"
	"net"
)

func main() {
	conn:=getConn()
	result1,_:=template.Mget(conn,"name","hh")
	fmt.Println("结果",result1)
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
