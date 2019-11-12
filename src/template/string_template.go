package template

import (
	"strings"
	"net"
	"protocol"
	"fmt"
)

/**
 * 认证权限
 */
func Auth(conn *net.TCPConn, pwd string) (interface{}, error) {
	result := SendCommand(conn, protocol.AUTH, protocol.SafeEncode(pwd))
	if result != protocol.OK {
		return nil, fmt.Errorf(result)
	}
	return result, nil
}

func Set(conn *net.TCPConn, key string, value string) (interface{}, error) {
	result := SendCommand(conn, protocol.SET, protocol.SafeEncode(key), protocol.SafeEncode(value))
	if result != protocol.OK {
		return nil, fmt.Errorf(result)
	}
	return strings.ReplaceAll(strings.ReplaceAll(result, protocol.CRLF, protocol.BLANK), protocol.PLUSBYTE, protocol.BLANK), nil
}

func Get(conn *net.TCPConn, key string) (interface{}, error) {
	result := SendCommand(conn, protocol.GET, protocol.SafeEncode(key))
	if strings.HasPrefix(result, protocol.NONEXIST) {
		return nil, nil
	}

	if !strings.HasPrefix(result, protocol.DOLLARBYTE) {
		return nil, fmt.Errorf(result)
	}
	array := strings.Split(result, protocol.CRLF)
	return array[1], nil
}

func Mset(conn *net.TCPConn, keyvalues ...string) (interface{}, error) {
	bytes := make([][]byte, len(keyvalues))
	for i := 0; i < len(keyvalues); i++ {
		bytes[i] = protocol.SafeEncode(keyvalues[i])
	}
	result := SendCommand(conn, protocol.MSET, bytes...)
	if result != protocol.OK {
		return nil, fmt.Errorf(result)
	}
	return strings.ReplaceAll(strings.ReplaceAll(result, protocol.CRLF, protocol.BLANK), protocol.PLUSBYTE, protocol.BLANK), nil
}

func Mget(conn *net.TCPConn, keys ...string) interface{} {
	bytes := make([][]byte, len(keys))
	for i := 0; i < len(keys); i++ {
		bytes[i] = protocol.SafeEncode(keys[i])
	}
	result := SendCommand(conn, protocol.MGET, bytes...)
	return HandleComplexResult(result)
}

func Setnx(conn *net.TCPConn, key string, value string) interface{} {
	return SendCommand(conn, protocol.SETNX, protocol.SafeEncode(key), protocol.SafeEncode(value))
}

func Incr(conn *net.TCPConn, key string) interface{} {
	result := SendCommand(conn, protocol.INCR, protocol.SafeEncode(key))
	return strings.ReplaceAll(
		strings.ReplaceAll(
			result, protocol.CRLF, protocol.BLANK),
		protocol.COLON_BYTE, protocol.BLANK)
}

func Decr(conn *net.TCPConn, key string) interface{} {
	result := SendCommand(conn, protocol.DECR, protocol.SafeEncode(key))
	return strings.ReplaceAll(strings.ReplaceAll(result, protocol.CRLF, protocol.BLANK), protocol.COLON_BYTE, protocol.BLANK)
}

func Setex(conn *net.TCPConn, key string, time int64, value string) (interface{}, error) {
	result := SendCommand(conn, protocol.SETEX, protocol.SafeEncode(key), protocol.SafeEncodeInt(time), protocol.SafeEncode(value))
	if result != protocol.OK {
		return nil, fmt.Errorf(result)
	}
	return strings.ReplaceAll(result, protocol.CRLF, protocol.BLANK), nil
}
