package template

import (
	"net"
	"protocol"
	"strings"
	"fmt"
)

func Zadd(conn *net.TCPConn, zset string, scoresvalues ...string) interface{} {
	bytes := make([][]byte, len(scoresvalues)+1)
	bytes[0] = protocol.SafeEncode(zset)
	for i := 0; i < len(scoresvalues); i++ {
		bytes[i+1] = protocol.SafeEncode(scoresvalues[i])
	}
	result := SendCommand(conn, protocol.ZADD, bytes...)
	return strings.ReplaceAll(strings.ReplaceAll(result, protocol.CRLF, protocol.BLANK), protocol.COLON_BYTE, protocol.BLANK)
}

func Zscore(conn *net.TCPConn, zset string, value string) (interface{}, error) {
	result := SendCommand(conn, protocol.ZSCORE, protocol.SafeEncode(zset), protocol.SafeEncode(value))
	if strings.HasPrefix(result, protocol.NONEXIST) {
		return nil, nil
	}

	if !strings.HasPrefix(result, protocol.DOLLARBYTE) {
		return nil, fmt.Errorf(result)
	}
	array := strings.Split(result, protocol.CRLF)
	return array[1], nil
}

func Zrange(conn *net.TCPConn, zset string, start int64, end int64) interface{} {
	result := SendCommand(conn, protocol.ZRANGE, protocol.SafeEncode(zset), protocol.SafeEncodeInt(start), protocol.SafeEncodeInt(end))
	return HandleComplexResult(result)
}

func ZrangeWithScores(conn *net.TCPConn, zset string, start int64, end int64) interface{} {
	result := SendCommand(conn, protocol.ZRANGE, protocol.SafeEncode(zset), protocol.SafeEncodeInt(start), protocol.SafeEncodeInt(end), protocol.SafeEncode("WITHSCORES"))
	return HandleComplexResult(result)
}

func Zrevrange(conn *net.TCPConn, zset string, start int64, end int64) interface{} {
	result := SendCommand(conn, protocol.ZREVRANGE, protocol.SafeEncode(zset), protocol.SafeEncodeInt(start), protocol.SafeEncodeInt(end))
	return HandleComplexResult(result)
}

func ZrevrangeWithScores(conn *net.TCPConn, zset string, start int64, end int64) interface{} {
	result := SendCommand(conn, protocol.ZREVRANGE, protocol.SafeEncode(zset), protocol.SafeEncodeInt(start), protocol.SafeEncodeInt(end), protocol.SafeEncode("WITHSCORES"))
	return HandleComplexResult(result)
}

func Zcard(conn *net.TCPConn, zset string) interface{} {
	result := SendCommand(conn, protocol.ZCARD, protocol.SafeEncode(zset))
	return strings.ReplaceAll(strings.ReplaceAll(result, protocol.CRLF, protocol.BLANK), protocol.COLON_BYTE, protocol.BLANK)
}

func Zrem(conn *net.TCPConn, zset string, elements ...string) interface{} {
	bytes := make([][]byte, len(elements)+1)
	bytes[0] = protocol.SafeEncode(zset)
	for i := 0; i < len(elements); i++ {
		bytes[i+1] = protocol.SafeEncode(elements[i])
	}
	result := SendCommand(conn, protocol.ZREM, bytes...)
	return strings.ReplaceAll(strings.ReplaceAll(result, protocol.CRLF, protocol.BLANK), protocol.COLON_BYTE, protocol.BLANK)
}

func Zrank(conn *net.TCPConn, zset string, value string) interface{} {
	result := SendCommand(conn, protocol.ZRANK, protocol.SafeEncode(zset))
	return strings.ReplaceAll(strings.ReplaceAll(result, protocol.CRLF, protocol.BLANK), protocol.COLON_BYTE, protocol.BLANK)
}

func Zrevrank(conn *net.TCPConn, zset string, value string) interface{} {
	result := SendCommand(conn, protocol.ZREVRANK, protocol.SafeEncode(zset))
	return strings.ReplaceAll(strings.ReplaceAll(result, protocol.CRLF, protocol.BLANK), protocol.COLON_BYTE, protocol.BLANK)
}
