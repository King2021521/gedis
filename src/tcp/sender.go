package tcp

import (
	"fmt"
	"os"
	"net"
)

func Sender(conn *net.TCPConn, content string) string{
	//给服务器发信息
	_, err := conn.Write([]byte(content))

	if err != nil {
		fmt.Println(conn.RemoteAddr().String(), "服务器反馈")
		os.Exit(1)
	}

	buffer := make([]byte, 1024)
	msg, err := conn.Read(buffer) //接受服务器信息
	if err != nil {
		fmt.Println(conn.RemoteAddr().String(), "服务器反馈"+err.Error())
		os.Exit(1)
	}
	return string(buffer[:msg])
}