package template

import (
	"net"
	"core"
	"fmt"
	"strings"
)

func Auth(pwd string, conn *net.TCPConn) string {
	var content = fmt.Sprintf("*2\r\n$4\r\nauth\r\n$%d\r\n%s\r\n", len(pwd), pwd)
	result := core.Sender(conn, content)
	return result
}

func Set(key string, value string, conn *net.TCPConn) string {
	var content = fmt.Sprintf("*3\r\n$3\r\nset\r\n$%d\r\n%s\r\n$%d\r\n%s\r\n", len(key), key, len(value), value)
	result := core.Sender(conn, content)
	return result
}

func Get(key string, conn *net.TCPConn) string {
	var content = fmt.Sprintf("*2\r\n$3\r\nget\r\n$%d\r\n%s\r\n", len(key), key)
	result := core.Sender(conn, content)
	array:=strings.Split(result,"\r\n")
	return array[1]
}
