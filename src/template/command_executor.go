package template

import (
	"net"
	"bytes"
	"protocol"
	"strconv"
	"fmt"
	"os"
)

func SendCommand(conn *net.TCPConn, cmd string, a ...[]byte) string {
	var buffer bytes.Buffer
	buffer.Write(protocol.SafeEncode(protocol.ASTERISKBYTE))
	buffer.Write(protocol.SafeEncode(strconv.Itoa(len(a) + 1)))
	buffer.Write(protocol.SafeEncode(protocol.CRLF))
	buffer.Write(protocol.SafeEncode(protocol.DOLLARBYTE))
	buffer.Write(protocol.SafeEncode(strconv.Itoa(len(cmd))))
	buffer.Write(protocol.SafeEncode(protocol.CRLF))
	buffer.Write(protocol.SafeEncode(cmd))
	buffer.Write(protocol.SafeEncode(protocol.CRLF))

	for _, arg := range a {
		buffer.Write(protocol.SafeEncode(protocol.DOLLARBYTE))
		buffer.Write(protocol.SafeEncode(strconv.Itoa(len(arg))))
		buffer.Write(protocol.SafeEncode(protocol.CRLF))
		buffer.Write(arg)
		buffer.Write(protocol.SafeEncode(protocol.CRLF))
	}
	return send(conn, buffer)
}

func send(conn *net.TCPConn, content bytes.Buffer) string {
	//send to server
	_, err := conn.Write(content.Bytes())

	if err != nil {
		fmt.Println(conn.RemoteAddr().String(), "server response")
		os.Exit(1)
	}

	buffer := make([]byte, 1024)
	//receive server info
	msg, err := conn.Read(buffer)
	if err != nil {
		fmt.Println(conn.RemoteAddr().String(), "server response:"+err.Error())
		os.Exit(1)
	}
	return string(buffer[:msg])
}