package template

import (
	"net"
	"protocol"
	"strings"
	"fmt"
)

func Lpush(conn *net.TCPConn, list string, elements ...string) interface{} {
	return push(conn, list, protocol.LPUSH, elements)
}

func Rpush(conn *net.TCPConn, list string, elements ...string) interface{} {
	return push(conn, list, protocol.RPUSH, elements)
}

func push(conn *net.TCPConn, list string, cmd string, elements []string) interface{} {
	bytes := make([][]byte, len(elements)+1)
	bytes[0] = protocol.SafeEncode(list)
	for i := 0; i < len(elements); i++ {
		bytes[i+1] = protocol.SafeEncode(elements[i])
	}
	result := SendCommand(conn, cmd, bytes...)
	return strings.ReplaceAll(strings.ReplaceAll(result, protocol.CRLF, protocol.BLANK), protocol.COLON_BYTE, protocol.BLANK)
}

func Lrange(conn *net.TCPConn, list string, start int64, end int64) interface{} {
	result := SendCommand(conn, protocol.LRANGE, protocol.SafeEncode(list), protocol.SafeEncodeInt(start), protocol.SafeEncodeInt(end))
	return HandleComplexResult(result)
}

func Lpop(conn *net.TCPConn, list string) (interface{}, error) {
	return pop(conn, list, protocol.LPOP)
}

func Rpop(conn *net.TCPConn, list string) (interface{}, error) {
	return pop(conn, list, protocol.RPOP)
}

func pop(conn *net.TCPConn, list string, cmd string) (interface{}, error) {
	result := SendCommand(conn, cmd, protocol.SafeEncode(list))
	if strings.HasPrefix(result, protocol.NONEXIST) {
		return nil, nil
	}

	if !strings.HasPrefix(result, protocol.DOLLARBYTE) {
		return nil, fmt.Errorf(result)
	}
	array := strings.Split(result, protocol.CRLF)
	return array[1], nil
}

func Llen(conn *net.TCPConn, list string) interface{} {
	result := SendCommand(conn, protocol.LLEN, protocol.SafeEncode(list))
	return strings.ReplaceAll(strings.ReplaceAll(result, protocol.CRLF, protocol.BLANK), protocol.COLON_BYTE, protocol.BLANK)
}

/**
 *  count > 0 : 从表头开始向表尾搜索，移除与 VALUE 相等的元素，数量为 COUNT 。
 *  count < 0 : 从表尾开始向表头搜索，移除与 VALUE 相等的元素，数量为 COUNT 的绝对值。
 *  count = 0 : 移除表中所有与 VALUE 相等的值。
 */
func Lrem(conn *net.TCPConn, list string, count int64, value string) interface{} {
	result := SendCommand(conn, protocol.LREM, protocol.SafeEncode(list), protocol.SafeEncodeInt(count), protocol.SafeEncode(value))
	return strings.ReplaceAll(strings.ReplaceAll(result, protocol.CRLF, protocol.BLANK), protocol.COLON_BYTE, protocol.BLANK)
}

func Lindex(conn *net.TCPConn, list string, pos int64) (interface{}, error) {
	result := SendCommand(conn, protocol.LINDEX, protocol.SafeEncode(list), protocol.SafeEncodeInt(pos))
	if strings.HasPrefix(result, protocol.NONEXIST) {
		return nil, nil
	}

	if !strings.HasPrefix(result, protocol.DOLLARBYTE) {
		return nil, fmt.Errorf(result)
	}
	array := strings.Split(result, protocol.CRLF)
	return array[1], nil
}

func Lset(conn *net.TCPConn, list string, pos int64, value string) (interface{}, error) {
	result := SendCommand(conn, protocol.LSET, protocol.SafeEncode(list), protocol.SafeEncodeInt(pos), protocol.SafeEncode(value))
	if result != protocol.OK {
		return nil, fmt.Errorf(result)
	}
	return strings.ReplaceAll(strings.ReplaceAll(result, protocol.CRLF, protocol.BLANK), protocol.PLUSBYTE, protocol.BLANK), nil
}

/**
 * target 目标元素
 */
func LinsertBefore(conn *net.TCPConn, list string, target string, value string) interface{} {
	result := SendCommand(conn, protocol.LINSERT, protocol.SafeEncode(list), protocol.SafeEncode("BEFORE"), protocol.SafeEncode(target), protocol.SafeEncode(value))
	return strings.ReplaceAll(strings.ReplaceAll(result, protocol.CRLF, protocol.BLANK), protocol.COLON_BYTE, protocol.BLANK)
}

/**
 * target 目标元素
 */
func LinsertAfter(conn *net.TCPConn, list string, target string, value string) interface{} {
	result := SendCommand(conn, protocol.LINSERT, protocol.SafeEncode(list), protocol.SafeEncode("AFTER"), protocol.SafeEncode(target), protocol.SafeEncode(value))
	return strings.ReplaceAll(strings.ReplaceAll(result, protocol.CRLF, protocol.BLANK), protocol.COLON_BYTE, protocol.BLANK)
}
