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
	//testRange(protocol.SET,"name","zxm")
	//testPool(getConn())
	testMset(getConn())
}

func testRange(cmd string, a ...interface{}){
	var buffer bytes.Buffer
	buffer.WriteString(protocol.ASTERISKBYTE)
	buffer.WriteString(strconv.Itoa(len(a)+1))
	buffer.WriteString(protocol.CRLF)
	buffer.WriteString(protocol.DOLLARBYTE)
	buffer.WriteString(strconv.Itoa(len(cmd)))
	buffer.WriteString(protocol.CRLF)
	buffer.WriteString(cmd)
	buffer.WriteString(protocol.CRLF)

	for _, arg := range a {
		buffer.WriteString(protocol.DOLLARBYTE)
		buffer.WriteString(strconv.Itoa(len(arg.(string))))
		buffer.WriteString(protocol.CRLF)
		buffer.WriteString(arg.(string))
		buffer.WriteString(protocol.CRLF)
	}
	fmt.Println(buffer.String())
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

func testMset(conn *net.TCPConn){
	r1:=template.Mset(conn,"lakers","james","sun","book","cc","$1")
	fmt.Println(r1)
	r2:=template.Mget(conn,"lakers","sun1","cc")
	fmt.Println("结果",r2,"len",len(r2),"cap",cap(r2))
}

func testPool(conn *net.TCPConn){
	sendResult1 := template.Set("country1", "China LiaoNing DaLian", conn)
	sendResult11 := template.Expire("country1", 10000, conn)
	fmt.Println("send result1:" , sendResult1)
	fmt.Println("send result11:" , sendResult11)
	result := template.Get("country1", conn)
	fmt.Println("get result country1:" + result)

	/*pool.PutConn(conn)

	conn1,_:=tcp.GetConn(pool)
	fmt.Println(conn1.RemoteAddr())
	sendResult1 := template.Set("name", "james", conn)
	fmt.Println("send result:" + sendResult1)
	result1 := template.Get("name", conn)
	fmt.Println("get result:" + result1)*/
	/*conn1,_:=tcp.GetConn(pool)
	fmt.Println(conn1.RemoteAddr())

	conn2,_:=tcp.GetConn(pool)
	fmt.Println(conn2.RemoteAddr())

	conn3,_:=tcp.GetConn(pool)
	fmt.Println(conn3.RemoteAddr())

	_, err = tcp.GetConn(pool)
	fmt.Println(err)

	size:=tcp.PoolSize(pool)
	fmt.Println("连接数",size)
	if size<1{
		pool.PutConn(conn3)
	}
	size1:=tcp.PoolSize(pool)
	fmt.Println("连接数",size1)

	conn4,_:=tcp.GetConn(pool)
	fmt.Println(conn4.RemoteAddr())*/
}

func testRedis(){
	conn := tcp.Connect("10.10.5.239:6379")
	authResult := template.Auth("123456", conn)
	fmt.Println("auth result:" + authResult)
	sendResult := template.Set("name", "james", conn)
	fmt.Println("send result:" + sendResult)
	result := template.Get("name", conn)
	fmt.Println("get result:" + result)
}
