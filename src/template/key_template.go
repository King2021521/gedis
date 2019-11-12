package template

import (
	"net"
	"protocol"
	"strings"
)

func Keys(conn *net.TCPConn, regex string) interface{} {
	result := SendCommand(conn, protocol.KEYS, protocol.SafeEncode(regex))
	return HandleComplexResult(result)
}

func Expire(conn *net.TCPConn, key string, value int64) interface{} {
	result := SendCommand(conn, protocol.EXPIRE, protocol.SafeEncode(key), protocol.SafeEncodeInt(value))
	return strings.ReplaceAll(strings.ReplaceAll(result, protocol.CRLF, protocol.BLANK), protocol.COLON_BYTE, protocol.BLANK)
}

func Del(conn *net.TCPConn, key string) interface{} {
	result := SendCommand(conn, protocol.DEL, protocol.SafeEncode(key))
	return strings.ReplaceAll(strings.ReplaceAll(result, protocol.CRLF, protocol.BLANK), protocol.COLON_BYTE, protocol.BLANK)
}

func Ttl(conn *net.TCPConn, key string) interface{} {
	result := SendCommand(conn, protocol.TTL, protocol.SafeEncode(key))
	return strings.ReplaceAll(strings.ReplaceAll(result, protocol.CRLF, protocol.BLANK), protocol.COLON_BYTE, protocol.BLANK)
}

