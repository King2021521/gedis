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
func Auth(pwd string, conn *net.TCPConn) (interface{}, error) {
	result := sendCommand(conn, protocol.AUTH, SafeEncode(pwd))
	if result != protocol.OK {
		return nil, fmt.Errorf(result)
	}
	return result, nil
}

func Set(key string, value string, conn *net.TCPConn) (interface{}, error) {
	result := sendCommand(conn, protocol.SET, SafeEncode(key), SafeEncode(value))
	if result != protocol.OK {
		return nil, fmt.Errorf(result)
	}
	return strings.ReplaceAll(result, protocol.CRLF, protocol.BLANK), nil
}

func Get(key string, conn *net.TCPConn) (interface{}, error) {
	result := sendCommand(conn, protocol.GET, SafeEncode(key))
	if strings.HasPrefix(result, protocol.NONEXIST) {
		return nil, nil
	}

	if !strings.HasPrefix(result, protocol.DOLLARBYTE) {
		return nil, fmt.Errorf(result)
	}
	array := strings.Split(result, protocol.CRLF)
	return array[1], nil
}

func Expire(key string, value int64, conn *net.TCPConn) interface{} {
	result := sendCommand(conn, protocol.EXPIRE, SafeEncode(key), encodeInt(value))
	return strings.ReplaceAll(strings.ReplaceAll(result, protocol.CRLF, protocol.BLANK), protocol.COLON_BYTE, protocol.BLANK)
}

func Mset(conn *net.TCPConn, keyvalues ...string) (interface{}, error) {
	bytes := make([][]byte, len(keyvalues))
	for i := 0; i < len(keyvalues); i++ {
		bytes[i] = SafeEncode(keyvalues[i])
	}
	result := sendCommand(conn, protocol.MSET, bytes...)
	if result != protocol.OK {
		return nil, fmt.Errorf(result)
	}
	return strings.ReplaceAll(result, protocol.CRLF, protocol.BLANK), nil
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
		if array[i] == protocol.NONEXIST {
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

func Del(key string, conn *net.TCPConn) interface{} {
	result := sendCommand(conn, protocol.DEL, SafeEncode(key))
	return strings.ReplaceAll(strings.ReplaceAll(result, protocol.CRLF, protocol.BLANK), protocol.COLON_BYTE, protocol.BLANK)
}

func Setnx(key string, value string, conn *net.TCPConn) interface{} {
	return sendCommand(conn, protocol.SETNX, SafeEncode(key), SafeEncode(value))
}

func Incr(key string, conn *net.TCPConn) interface{} {
	result := sendCommand(conn, protocol.INCR, SafeEncode(key))
	return strings.ReplaceAll(
		strings.ReplaceAll(
			result, protocol.CRLF, protocol.BLANK),
		protocol.COLON_BYTE, protocol.BLANK)
}

func Decr(key string, conn *net.TCPConn) interface{} {
	result := sendCommand(conn, protocol.DECR, SafeEncode(key))
	return strings.ReplaceAll(strings.ReplaceAll(result, protocol.CRLF, protocol.BLANK), protocol.COLON_BYTE, protocol.BLANK)
}

func Setex(key string, time int64, value string, conn *net.TCPConn) (interface{}, error) {
	result := sendCommand(conn, protocol.SETEX, SafeEncode(key), encodeInt(time), SafeEncode(value))
	if result != protocol.OK {
		return nil, fmt.Errorf(result)
	}
	return strings.ReplaceAll(result, protocol.CRLF, protocol.BLANK), nil
}

func Ttl(key string, conn *net.TCPConn) (int, error) {
	result := sendCommand(conn, protocol.TTL, SafeEncode(key))
	ttl := strings.ReplaceAll(strings.ReplaceAll(result, protocol.CRLF, protocol.BLANK), protocol.COLON_BYTE, protocol.BLANK)
	return strconv.Atoi(ttl)
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
