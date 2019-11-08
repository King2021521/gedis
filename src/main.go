package main

import (
	"tcp"
	"fmt"
	"template"
	"bytes"
	"protocol"
	"strconv"
	"net"
)

func main() {
	conn:=getConn()
	result,_:=template.Mset(conn,"root111","1111","asdasdad","asdadad")

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
