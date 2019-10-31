package template

import (
	"strings"
	"net"
	"protocol"
	"bytes"
	"strconv"
	"fmt"
	"os"
)

func Auth(pwd string, conn *net.TCPConn) string {
	return sendCommand(conn, protocol.AUTH, pwd)
}

func Set(key string, value string, conn *net.TCPConn) string {
	return sendCommand(conn, protocol.SET, key, value)
}

func Get(key string, conn *net.TCPConn) string {
	result := sendCommand(conn, protocol.GET, key)
	array := strings.Split(result, protocol.CRLF)
	return array[1]
}

func sendCommand(conn *net.TCPConn, cmd string, a ...interface{}) string {
	var buffer bytes.Buffer
	buffer.WriteString(protocol.ASTERISKBYTE)
	buffer.WriteString(strconv.Itoa(len(a) + 1))
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
	return send(conn, buffer.String())
}

func send(conn *net.TCPConn, content string) string {
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
