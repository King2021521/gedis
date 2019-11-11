package template

import (
	"net"
	"protocol"
	"fmt"
	"strings"
)

func Hset(conn *net.TCPConn, hash string, key string, value string) interface{} {
	result := SendCommand(conn, protocol.HSET, protocol.SafeEncode(hash), protocol.SafeEncode(key), protocol.SafeEncode(value))
	return strings.ReplaceAll(strings.ReplaceAll(result, protocol.CRLF, protocol.BLANK), protocol.COLON_BYTE, protocol.BLANK)
}

func Hget(conn *net.TCPConn, hash string, key string) (interface{}, error) {
	result := SendCommand(conn, protocol.HGET, protocol.SafeEncode(hash), protocol.SafeEncode(key))
	if strings.HasPrefix(result, protocol.NONEXIST) {
		return nil, nil
	}

	if !strings.HasPrefix(result, protocol.DOLLARBYTE) {
		return nil, fmt.Errorf(result)
	}
	array := strings.Split(result, protocol.CRLF)
	return array[1], nil
}

func Hmset(conn *net.TCPConn, hash string, keyvalues ...string) (interface{}, error) {
	bytes := make([][]byte, len(keyvalues)+1)
	bytes[0] = protocol.SafeEncode(hash)
	for i := 0; i < len(keyvalues); i++ {
		bytes[i+1] = protocol.SafeEncode(keyvalues[i])
	}
	result := SendCommand(conn, protocol.HMSET, bytes...)
	if result != protocol.OK {
		return nil, fmt.Errorf(result)
	}
	return strings.ReplaceAll(strings.ReplaceAll(result, protocol.CRLF, protocol.BLANK), protocol.PLUSBYTE, protocol.BLANK), nil
}

func Hmget(conn *net.TCPConn, hash string, keys ...string) interface{} {
	bytes := make([][]byte, len(keys)+1)
	bytes[0] = protocol.SafeEncode(hash)
	for i := 0; i < len(keys); i++ {
		bytes[i+1] = protocol.SafeEncode(keys[i])
	}
	result := SendCommand(conn, protocol.HMGET, bytes...)
	return HandleComplexResult(result)
}

func Hgetall(conn *net.TCPConn, hash string) interface{} {
	result := SendCommand(conn, protocol.HGETALL, protocol.SafeEncode(hash))
	return HandleComplexResult(result)
}

func Hexists(conn *net.TCPConn, hash string, key string) interface{} {
	result := SendCommand(conn, protocol.HEXISTS, protocol.SafeEncode(hash), protocol.SafeEncode(key))
	return strings.ReplaceAll(strings.ReplaceAll(result, protocol.CRLF, protocol.BLANK), protocol.COLON_BYTE, protocol.BLANK)
}

func Hdel(conn *net.TCPConn, hash string, key string) interface{} {
	result := SendCommand(conn, protocol.HDEL, protocol.SafeEncode(hash), protocol.SafeEncode(key))
	return strings.ReplaceAll(strings.ReplaceAll(result, protocol.CRLF, protocol.BLANK), protocol.COLON_BYTE, protocol.BLANK)
}

func Hkeys(conn *net.TCPConn, hash string) interface{} {
	result := SendCommand(conn, protocol.HKEYS, protocol.SafeEncode(hash))
	return HandleComplexResult(result)
}

func Hvalues(conn *net.TCPConn, hash string) interface{} {
	result := SendCommand(conn, protocol.HVALS, protocol.SafeEncode(hash))
	return HandleComplexResult(result)
}

func Hlen(conn *net.TCPConn, hash string) interface{} {
	result := SendCommand(conn, protocol.HLEN, protocol.SafeEncode(hash))
	return strings.ReplaceAll(strings.ReplaceAll(result, protocol.CRLF, protocol.BLANK), protocol.COLON_BYTE, protocol.BLANK)
}