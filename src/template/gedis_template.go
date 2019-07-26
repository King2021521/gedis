package template

import (
	"tcp"
	"fmt"
	"strings"
	"net"
	"protocol"
)

func Auth(pwd string, conn *net.TCPConn) string {
	var content = fmt.Sprintf(protocol.AuthCmdFormat, len(pwd), pwd)
	result := tcp.Sender(conn, content)
	return result
}

func Set(key string, value string, conn *net.TCPConn) string {
	var content = fmt.Sprintf(protocol.SetCmdFormat, len(key), key, len(value), value)
	result := tcp.Sender(conn, content)
	return result
}

func Get(key string, conn *net.TCPConn) string {
	var content = fmt.Sprintf(protocol.GetCmdFormat, len(key), key)
	result := tcp.Sender(conn, content)
	array:=strings.Split(result,protocol.Newline)
	return array[1]
}
