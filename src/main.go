package main

import (
	"tcp"
	"fmt"
	"net"
)

func main() {
	/*conn:=getConn()
	result1,_:=template.Mget(conn,"name","hh")
	fmt.Println("结果",result1)*/
}

func getConn() *net.TCPConn{
	var config = tcp.ConnConfig{"127.0.0.1:6379","root"}
	pool:=tcp.NewSingleConnPool(1, config)
	fmt.Println("------",pool)
	conn,_:=tcp.GetConn(pool)
	return conn
}
