package main

import (
	"tcp"
	"fmt"
	"template"
	"net"
)

func main() {
	conn:=getConn()
	result1,_:=template.Set(conn,"viewhigh","望海")
	fmt.Println("设置结果",result1)
	result2:=template.Expire(conn,"viewhigh", 1000)
	fmt.Println("expire结果",result2)
	result3:=template.Ttl(conn,"viewhigh")
	fmt.Println("ttl结果",result3)
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
