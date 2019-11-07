package template

import (
	"strings"
	"net"
	"protocol"
	"bytes"
	"strconv"
	"fmt"
	"os"
	"github.com/emirpasic/gods/lists/arraylist"
)

/**
 * 认证权限
 */
func Auth(pwd string, conn *net.TCPConn) string {
	return sendCommand(conn, protocol.AUTH, SafeEncode(pwd))
}

func Set(key string, value string, conn *net.TCPConn) string {
	return sendCommand(conn, protocol.SET, SafeEncode(key), SafeEncode(value))
}

func Get(key string, conn *net.TCPConn) string {
	result := sendCommand(conn, protocol.GET, SafeEncode(key))
	array := strings.Split(result, protocol.CRLF)
	return array[1]
}

func Expire(key string, value int64, conn *net.TCPConn) string {
	return sendCommand(conn, protocol.EXPIRE, SafeEncode(key), encodeInt(value))
}

func Mset(conn *net.TCPConn, keyvalues ...string) string {
	bytes := make([][]byte, len(keyvalues))
	for i := 0; i < len(keyvalues); i++ {
		bytes[i] = SafeEncode(keyvalues[i])
	}
	return sendCommand(conn, protocol.MSET, bytes...)
}

func Mget(conn *net.TCPConn, keys ...string) []interface{} {
	bytes := make([][]byte, len(keys))
	for i := 0; i < len(keys); i++ {
		bytes[i] = SafeEncode(keys[i])
	}
	result := sendCommand(conn, protocol.MGET, bytes...)
	array := strings.Split(result, protocol.CRLF)
	results := arraylist.New()
	for i := 1; i < len(array)-1; i++ {
		if array[i] == "$-1" {
			results.Add(nil)
			continue
		}
		results.Add(array[i+1])
		i++
		if i > len(array)-2 {
			break
		}
	}
	return results.Values()
}

func sendCommand(conn *net.TCPConn, cmd string, a ...[]byte) string {
	var buffer bytes.Buffer
	buffer.Write(SafeEncode(protocol.ASTERISKBYTE))
	buffer.Write(SafeEncode(strconv.Itoa(len(a) + 1)))
	buffer.Write(SafeEncode(protocol.CRLF))
	buffer.Write(SafeEncode(protocol.DOLLARBYTE))
	buffer.Write(SafeEncode(strconv.Itoa(len(cmd))))
	buffer.Write(SafeEncode(protocol.CRLF))
	buffer.Write(SafeEncode(cmd))
	buffer.Write(SafeEncode(protocol.CRLF))

	for _, arg := range a {
		buffer.Write(SafeEncode(protocol.DOLLARBYTE))
		buffer.Write(SafeEncode(strconv.Itoa(len(arg))))
		buffer.Write(SafeEncode(protocol.CRLF))
		buffer.Write(arg)
		buffer.Write(SafeEncode(protocol.CRLF))
	}
	return send(conn, buffer)
}

func SafeEncode(arg string) []byte {
	return []byte(arg)
}

func encodeInt(arg int64) []byte {
	return SafeEncode(strconv.FormatInt(arg, 10))
}

func send(conn *net.TCPConn, content bytes.Buffer) string {
	//给服务器发信息
	_, err := conn.Write(content.Bytes())

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
