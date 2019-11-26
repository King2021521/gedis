package main

import (
	"tcp"
	"fmt"
	"net"
	"bytes"
	"encoding/binary"
	. "cluster"
)

func main() {
	/*conn:=getConn()
	result1,_:=template.Mget(conn,"name","hh")
	fmt.Println("结果",result1)*/
	//mData := []byte{0x01,0x02,0x03,0x04}
	mData :=[]byte("zhangxiaomin")
	checksum := Crc16(mData)

	fmt.Printf("check sum:%d \n",GetHashSlot(checksum))

	int16buf := new(bytes.Buffer)

	binary.Write(int16buf,binary.LittleEndian,checksum)
	fmt.Printf("write buf is: %+X \n",int16buf.Bytes())

	fmt.Printf("output-before:%X \n", mData)
	mData = append(mData,int16buf.Bytes()...)

	fmt.Printf("output-after:%X \n", mData)

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
