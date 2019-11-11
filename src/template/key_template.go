package template

import (
	"net"
	"protocol"
	"strings"
	"strconv"
)

func Keys(conn *net.TCPConn, regex string) interface{} {
	result := SendCommand(conn, protocol.KEYS, protocol.SafeEncode(regex))
	return HandleComplexResult(result)
}

func Expire(key string, value int64, conn *net.TCPConn) interface{} {
	result := SendCommand(conn, protocol.EXPIRE, protocol.SafeEncode(key), protocol.SafeEncodeInt(value))
	return strings.ReplaceAll(strings.ReplaceAll(result, protocol.CRLF, protocol.BLANK), protocol.COLON_BYTE, protocol.BLANK)
}

func Del(key string, conn *net.TCPConn) interface{} {
	result := SendCommand(conn, protocol.DEL, protocol.SafeEncode(key))
	return strings.ReplaceAll(strings.ReplaceAll(result, protocol.CRLF, protocol.BLANK), protocol.COLON_BYTE, protocol.BLANK)
}

func Ttl(key string, conn *net.TCPConn) (int, error) {
	result := SendCommand(conn, protocol.TTL, protocol.SafeEncode(key))
	ttl := strings.ReplaceAll(strings.ReplaceAll(result, protocol.CRLF, protocol.BLANK), protocol.COLON_BYTE, protocol.BLANK)
	return strconv.Atoi(ttl)
}

