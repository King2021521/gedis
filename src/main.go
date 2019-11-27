package main

import (
	"client"
	"fmt"
	"net"
)

func main() {
	var config = client.ConnConfig{"127.0.0.1:6379","root"}
	pool:= client.NewSingleConnPool(1, config)

	client1:=client.BuildClient(pool)
	client2:=client.BuildClient(pool)
	fmt.Println(client1,client2)
	//fmt.Println(client.Get("hh"))
}

func getConn() *net.TCPConn{
	var config = client.ConnConfig{"127.0.0.1:6379","root"}
	pool:= client.NewSingleConnPool(1, config)
	conn,_:= client.GetConn(pool)
	return conn
}
