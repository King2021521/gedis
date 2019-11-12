package template

import (
	"net"
	"protocol"
	"template/handler"
)

func Lpush(conn *net.TCPConn, list string, elements ...string) (interface{}, error) {
	return push(conn, list, protocol.LPUSH, elements)
}

func Rpush(conn *net.TCPConn, list string, elements ...string) (interface{}, error) {
	return push(conn, list, protocol.RPUSH, elements)
}

func push(conn *net.TCPConn, list string, cmd string, elements []string) (interface{}, error) {
	bytes := handler.HandleMultiBulkRequest(list, elements)
	result := SendCommand(conn, cmd, bytes...)
	return handler.HandleReply(result)
}

func Lrange(conn *net.TCPConn, list string, start int64, end int64) (interface{}, error) {
	result := SendCommand(conn, protocol.LRANGE, protocol.SafeEncode(list), protocol.SafeEncodeInt(start), protocol.SafeEncodeInt(end))
	return handler.HandleReply(result)
}

func Lpop(conn *net.TCPConn, list string) (interface{}, error) {
	return pop(conn, list, protocol.LPOP)
}

func Rpop(conn *net.TCPConn, list string) (interface{}, error) {
	return pop(conn, list, protocol.RPOP)
}

func pop(conn *net.TCPConn, list string, cmd string) (interface{}, error) {
	result := SendCommand(conn, cmd, protocol.SafeEncode(list))
	return handler.HandleReply(result)
}

func Llen(conn *net.TCPConn, list string) (interface{}, error) {
	result := SendCommand(conn, protocol.LLEN, protocol.SafeEncode(list))
	return handler.HandleReply(result)
}

/**
 *  count > 0 : 从表头开始向表尾搜索，移除与 VALUE 相等的元素，数量为 COUNT 。
 *  count < 0 : 从表尾开始向表头搜索，移除与 VALUE 相等的元素，数量为 COUNT 的绝对值。
 *  count = 0 : 移除表中所有与 VALUE 相等的值。
 */
func Lrem(conn *net.TCPConn, list string, count int64, value string) (interface{}, error) {
	result := SendCommand(conn, protocol.LREM, protocol.SafeEncode(list), protocol.SafeEncodeInt(count), protocol.SafeEncode(value))
	return handler.HandleReply(result)
}

func Lindex(conn *net.TCPConn, list string, pos int64) (interface{}, error) {
	result := SendCommand(conn, protocol.LINDEX, protocol.SafeEncode(list), protocol.SafeEncodeInt(pos))
	return handler.HandleReply(result)
}

func Lset(conn *net.TCPConn, list string, pos int64, value string) (interface{}, error) {
	result := SendCommand(conn, protocol.LSET, protocol.SafeEncode(list), protocol.SafeEncodeInt(pos), protocol.SafeEncode(value))
	return handler.HandleReply(result)
}

/**
 * target 目标元素
 */
func LinsertBefore(conn *net.TCPConn, list string, target string, value string) (interface{}, error) {
	result := SendCommand(conn, protocol.LINSERT, protocol.SafeEncode(list), protocol.SafeEncode("BEFORE"), protocol.SafeEncode(target), protocol.SafeEncode(value))
	return handler.HandleReply(result)
}

/**
 * target 目标元素
 */
func LinsertAfter(conn *net.TCPConn, list string, target string, value string) (interface{}, error) {
	result := SendCommand(conn, protocol.LINSERT, protocol.SafeEncode(list), protocol.SafeEncode("AFTER"), protocol.SafeEncode(target), protocol.SafeEncode(value))
	return handler.HandleReply(result)
}
