package template

import (
	"net"
	"protocol"
	"strings"
	"fmt"
)

func Sadd(conn *net.TCPConn, set string, elements ...string) interface{} {
	bytes := make([][]byte, len(elements)+1)
	bytes[0] = protocol.SafeEncode(set)
	for i := 0; i < len(elements); i++ {
		bytes[i+1] = protocol.SafeEncode(elements[i])
	}
	result := SendCommand(conn, protocol.SADD, bytes...)
	return strings.ReplaceAll(strings.ReplaceAll(result, protocol.CRLF, protocol.BLANK), protocol.COLON_BYTE, protocol.BLANK)
}

func Smembers(conn *net.TCPConn, set string) interface{} {
	result := SendCommand(conn, protocol.SMEMBERS, protocol.SafeEncode(set))
	return HandleComplexResult(result)
}

func Srem(conn *net.TCPConn, set string, elements ...string) interface{} {
	bytes := make([][]byte, len(elements)+1)
	bytes[0] = protocol.SafeEncode(set)
	for i := 0; i < len(elements); i++ {
		bytes[i+1] = protocol.SafeEncode(elements[i])
	}
	result := SendCommand(conn, protocol.SREM, bytes...)
	return strings.ReplaceAll(strings.ReplaceAll(result, protocol.CRLF, protocol.BLANK), protocol.COLON_BYTE, protocol.BLANK)
}

func Sismember(conn *net.TCPConn, set string, value string) interface{} {
	result := SendCommand(conn, protocol.SISMEMBER, protocol.SafeEncode(set), protocol.SafeEncode(value))
	return strings.ReplaceAll(strings.ReplaceAll(result, protocol.CRLF, protocol.BLANK), protocol.COLON_BYTE, protocol.BLANK)
}

func Scard(conn *net.TCPConn, set string) interface{} {
	result := SendCommand(conn, protocol.SCARD, protocol.SafeEncode(set))
	return strings.ReplaceAll(strings.ReplaceAll(result, protocol.CRLF, protocol.BLANK), protocol.COLON_BYTE, protocol.BLANK)
}

func Srandmembers(conn *net.TCPConn, set string, count int64) interface{} {
	result := SendCommand(conn, protocol.SRANDMEMBER, protocol.SafeEncode(set), protocol.SafeEncodeInt(count))
	return HandleComplexResult(result)
}

func Spop(conn *net.TCPConn, set string) (interface{}, error) {
	result := SendCommand(conn, protocol.SPOP, protocol.SafeEncode(set))
	if strings.HasPrefix(result, protocol.NONEXIST) {
		return nil, nil
	}

	if !strings.HasPrefix(result, protocol.DOLLARBYTE) {
		return nil, fmt.Errorf(result)
	}
	array := strings.Split(result, protocol.CRLF)
	return array[1], nil
}

/**
 * 返回给定所有集合的差集
 */
func Sdiff(conn *net.TCPConn, sets ... string) interface{} {
	bytes := make([][]byte, len(sets))
	for i := 0; i < len(sets); i++ {
		bytes[i] = protocol.SafeEncode(sets[i])
	}
	result := SendCommand(conn, protocol.SDIFF, bytes...)
	return HandleComplexResult(result)
}
