//@author zhangxiaomin
package client

import (
	"net"
	"bytes"
	"protocol"
	"strconv"
	"fmt"
)
/**
 * 通信协议处理：如redis操作：set key value，解析成redis理解的格式后就是：
 * *3\r\n$3\r\nset\r\n$3\r\nkey\r\n$5\r\nvalue\r\n
 */
func SendCommand(conn *net.TCPConn, cmd string, a ...[]byte) string {
	//前置处理命令前缀部分，格式固定
	var buffer bytes.Buffer
	buffer.Write(protocol.SafeEncode(protocol.ASTERISKBYTE))
	buffer.Write(protocol.SafeEncode(strconv.Itoa(len(a) + 1)))
	buffer.Write(protocol.SafeEncode(protocol.CRLF))
	buffer.Write(protocol.SafeEncode(protocol.DOLLARBYTE))
	buffer.Write(protocol.SafeEncode(strconv.Itoa(len(cmd))))
	buffer.Write(protocol.SafeEncode(protocol.CRLF))
	buffer.Write(protocol.SafeEncode(cmd))
	buffer.Write(protocol.SafeEncode(protocol.CRLF))

	//循环处理命令详情部分（key、value长度不一定固定，所以循环处理，并拼接）
	for _, arg := range a {
		buffer.Write(protocol.SafeEncode(protocol.DOLLARBYTE))
		buffer.Write(protocol.SafeEncode(strconv.Itoa(len(arg))))
		buffer.Write(protocol.SafeEncode(protocol.CRLF))
		buffer.Write(arg)
		buffer.Write(protocol.SafeEncode(protocol.CRLF))
	}
	return send(conn, buffer)
}

/**
 * tcp通信，请求redis sever
 */
func send(conn *net.TCPConn, content bytes.Buffer) string {
	//send to server
	_, err := conn.Write(content.Bytes())

	if err != nil {
		fmt.Println(conn.RemoteAddr().String(), "server response", err.Error())
		return err.Error()
	}

	buffer := make([]byte, 1024)
	//receive server info
	msg, err := conn.Read(buffer)
	if err != nil {
		fmt.Println(conn.RemoteAddr().String(), "server response:"+err.Error())
		return err.Error()
	}
	return string(buffer[:msg])
}